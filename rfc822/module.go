package rfc822
import (
	πg "github.com/pygolin/runtime"
	// _ "github.com/pygolin/stdlib/sys"
	// _ "github.com/pygolin/stdlib/warnings"
	// _ "github.com/pygolin/stdlib/os"
	// _ "github.com/pygolin/stdlib/time"
)
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/rfc822.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ß := πg.InternStr("")
		ß0 := πg.InternStr("0")
		ßADT := πg.InternStr("ADT")
		ßAST := πg.InternStr("AST")
		ßAddressList := πg.InternStr("AddressList")
		ßAddrlistClass := πg.InternStr("AddrlistClass")
		ßApr := πg.InternStr("Apr")
		ßAttributeError := πg.InternStr("AttributeError")
		ßAug := πg.InternStr("Aug")
		ßCDT := πg.InternStr("CDT")
		ßCR := πg.InternStr("CR")
		ßCST := πg.InternStr("CST")
		ßDate := πg.InternStr("Date")
		ßDec := πg.InternStr("Dec")
		ßEDT := πg.InternStr("EDT")
		ßEST := πg.InternStr("EST")
		ßFalse := πg.InternStr("False")
		ßFeb := πg.InternStr("Feb")
		ßFri := πg.InternStr("Fri")
		ßGMT := πg.InternStr("GMT")
		ßHOME := πg.InternStr("HOME")
		ßIOError := πg.InternStr("IOError")
		ßJan := πg.InternStr("Jan")
		ßJul := πg.InternStr("Jul")
		ßJun := πg.InternStr("Jun")
		ßKeyError := πg.InternStr("KeyError")
		ßLWS := πg.InternStr("LWS")
		ßMDT := πg.InternStr("MDT")
		ßMST := πg.InternStr("MST")
		ßMar := πg.InternStr("Mar")
		ßMay := πg.InternStr("May")
		ßMessage := πg.InternStr("Message")
		ßMon := πg.InternStr("Mon")
		ßNone := πg.InternStr("None")
		ßNov := πg.InternStr("Nov")
		ßOct := πg.InternStr("Oct")
		ßPDT := πg.InternStr("PDT")
		ßPST := πg.InternStr("PST")
		ßSat := πg.InternStr("Sat")
		ßSep := πg.InternStr("Sep")
		ßSun := πg.InternStr("Sun")
		ßThu := πg.InternStr("Thu")
		ßTue := πg.InternStr("Tue")
		ßUT := πg.InternStr("UT")
		ßUTC := πg.InternStr("UTC")
		ßValueError := πg.InternStr("ValueError")
		ßWed := πg.InternStr("Wed")
		ßZ := πg.InternStr("Z")
		ß__add__ := πg.InternStr("__add__")
		ß__all__ := πg.InternStr("__all__")
		ß__contains__ := πg.InternStr("__contains__")
		ß__delitem__ := πg.InternStr("__delitem__")
		ß__doc__ := πg.InternStr("__doc__")
		ß__getitem__ := πg.InternStr("__getitem__")
		ß__iadd__ := πg.InternStr("__iadd__")
		ß__init__ := πg.InternStr("__init__")
		ß__isub__ := πg.InternStr("__isub__")
		ß__iter__ := πg.InternStr("__iter__")
		ß__len__ := πg.InternStr("__len__")
		ß__main__ := πg.InternStr("__main__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß__setitem__ := πg.InternStr("__setitem__")
		ß__str__ := πg.InternStr("__str__")
		ß__sub__ := πg.InternStr("__sub__")
		ß_blanklines := πg.InternStr("_blanklines")
		ß_daynames := πg.InternStr("_daynames")
		ß_monthnames := πg.InternStr("_monthnames")
		ß_timezones := πg.InternStr("_timezones")
		ßaddresslist := πg.InternStr("addresslist")
		ßappend := πg.InternStr("append")
		ßapr := πg.InternStr("apr")
		ßapril := πg.InternStr("april")
		ßargv := πg.InternStr("argv")
		ßasctime := πg.InternStr("asctime")
		ßatomends := πg.InternStr("atomends")
		ßaug := πg.InternStr("aug")
		ßaugust := πg.InternStr("august")
		ßcommentlist := πg.InternStr("commentlist")
		ßdate := πg.InternStr("date")
		ßdec := πg.InternStr("dec")
		ßdecember := πg.InternStr("december")
		ßdict := πg.InternStr("dict")
		ßdivmod := πg.InternStr("divmod")
		ßdump_address_pair := πg.InternStr("dump_address_pair")
		ßendswith := πg.InternStr("endswith")
		ßenviron := πg.InternStr("environ")
		ßf := πg.InternStr("f")
		ßfeb := πg.InternStr("feb")
		ßfebruary := πg.InternStr("february")
		ßfield := πg.InternStr("field")
		ßfile := πg.InternStr("file")
		ßfind := πg.InternStr("find")
		ßformatdate := πg.InternStr("formatdate")
		ßfp := πg.InternStr("fp")
		ßfri := πg.InternStr("fri")
		ßfrom := πg.InternStr("from")
		ßget := πg.InternStr("get")
		ßgetaddr := πg.InternStr("getaddr")
		ßgetaddress := πg.InternStr("getaddress")
		ßgetaddrlist := πg.InternStr("getaddrlist")
		ßgetaddrspec := πg.InternStr("getaddrspec")
		ßgetallmatchingheaders := πg.InternStr("getallmatchingheaders")
		ßgetatom := πg.InternStr("getatom")
		ßgetcomment := πg.InternStr("getcomment")
		ßgetdate := πg.InternStr("getdate")
		ßgetdate_tz := πg.InternStr("getdate_tz")
		ßgetdelimited := πg.InternStr("getdelimited")
		ßgetdomain := πg.InternStr("getdomain")
		ßgetdomainliteral := πg.InternStr("getdomainliteral")
		ßgetfirstmatchingheader := πg.InternStr("getfirstmatchingheader")
		ßgetheader := πg.InternStr("getheader")
		ßgetheaders := πg.InternStr("getheaders")
		ßgetphraselist := πg.InternStr("getphraselist")
		ßgetquote := πg.InternStr("getquote")
		ßgetrawheader := πg.InternStr("getrawheader")
		ßgetrouteaddr := πg.InternStr("getrouteaddr")
		ßgmtime := πg.InternStr("gmtime")
		ßgotonext := πg.InternStr("gotonext")
		ßhas_key := πg.InternStr("has_key")
		ßhasattr := πg.InternStr("hasattr")
		ßheaders := πg.InternStr("headers")
		ßhh := πg.InternStr("hh")
		ßhhmm := πg.InternStr("hhmm")
		ßhhmmss := πg.InternStr("hhmmss")
		ßindex := πg.InternStr("index")
		ßint := πg.InternStr("int")
		ßiscomment := πg.InternStr("iscomment")
		ßisdigit := πg.InternStr("isdigit")
		ßisheader := πg.InternStr("isheader")
		ßislast := πg.InternStr("islast")
		ßisspace := πg.InternStr("isspace")
		ßitems := πg.InternStr("items")
		ßiter := πg.InternStr("iter")
		ßjan := πg.InternStr("jan")
		ßjanuary := πg.InternStr("january")
		ßjoin := πg.InternStr("join")
		ßjul := πg.InternStr("jul")
		ßjuly := πg.InternStr("july")
		ßjun := πg.InternStr("jun")
		ßjune := πg.InternStr("june")
		ßkeys := πg.InternStr("keys")
		ßlen := πg.InternStr("len")
		ßlocaltime := πg.InternStr("localtime")
		ßlower := πg.InternStr("lower")
		ßm := πg.InternStr("m")
		ßmap := πg.InternStr("map")
		ßmar := πg.InternStr("mar")
		ßmarch := πg.InternStr("march")
		ßmay := πg.InternStr("may")
		ßmktime := πg.InternStr("mktime")
		ßmktime_tz := πg.InternStr("mktime_tz")
		ßmm := πg.InternStr("mm")
		ßmon := πg.InternStr("mon")
		ßn := πg.InternStr("n")
		ßnov := πg.InternStr("nov")
		ßnovember := πg.InternStr("november")
		ßobject := πg.InternStr("object")
		ßoct := πg.InternStr("oct")
		ßoctober := πg.InternStr("october")
		ßopen := πg.InternStr("open")
		ßos := πg.InternStr("os")
		ßparseaddr := πg.InternStr("parseaddr")
		ßparsedate := πg.InternStr("parsedate")
		ßparsedate_tz := πg.InternStr("parsedate_tz")
		ßpath := πg.InternStr("path")
		ßphraseends := πg.InternStr("phraseends")
		ßpos := πg.InternStr("pos")
		ßquote := πg.InternStr("quote")
		ßr := πg.InternStr("r")
		ßrange := πg.InternStr("range")
		ßreadheaders := πg.InternStr("readheaders")
		ßreadline := πg.InternStr("readline")
		ßremove := πg.InternStr("remove")
		ßreplace := πg.InternStr("replace")
		ßreversed := πg.InternStr("reversed")
		ßrewindbody := πg.InternStr("rewindbody")
		ßrfind := πg.InternStr("rfind")
		ßsat := πg.InternStr("sat")
		ßseek := πg.InternStr("seek")
		ßseekable := πg.InternStr("seekable")
		ßsep := πg.InternStr("sep")
		ßseptember := πg.InternStr("september")
		ßsetdefault := πg.InternStr("setdefault")
		ßspecials := πg.InternStr("specials")
		ßsplit := πg.InternStr("split")
		ßss := πg.InternStr("ss")
		ßstartofbody := πg.InternStr("startofbody")
		ßstartofheaders := πg.InternStr("startofheaders")
		ßstartswith := πg.InternStr("startswith")
		ßstatus := πg.InternStr("status")
		ßstrip := πg.InternStr("strip")
		ßsubject := πg.InternStr("subject")
		ßsun := πg.InternStr("sun")
		ßsys := πg.InternStr("sys")
		ßtell := πg.InternStr("tell")
		ßthu := πg.InternStr("thu")
		ßtime := πg.InternStr("time")
		ßtimezone := πg.InternStr("timezone")
		ßto := πg.InternStr("to")
		ßtue := πg.InternStr("tue")
		ßtz := πg.InternStr("tz")
		ßunixfrom := πg.InternStr("unixfrom")
		ßunquote := πg.InternStr("unquote")
		ßunread := πg.InternStr("unread")
		ßupper := πg.InternStr("upper")
		ßvalues := πg.InternStr("values")
		ßwarnpy3k := πg.InternStr("warnpy3k")
		ßwed := πg.InternStr("wed")
		var πTemp001 *πg.Object
		_ = πTemp001
		var πTemp002 []*πg.Object
		_ = πTemp002
		var πTemp003 *πg.Object
		_ = πTemp003
		var πTemp004 πg.KWArgs
		_ = πTemp004
		var πTemp005 *πg.Dict
		_ = πTemp005
		var πTemp006 *πg.Object
		_ = πTemp006
		var πTemp007 []πg.Param
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
		var πTemp015 bool
		_ = πTemp015
		var πTemp016 *πg.Object
		_ = πTemp016
		var πTemp017 *πg.Object
		_ = πTemp017
		var πTemp018 []*πg.Object
		_ = πTemp018
		var πTemp019 bool
		_ = πTemp019
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			case 10: goto Label10
			case 11: goto Label11
			default: panic("unexpected function state")
			}
			// line 1: """RFC 2822 message manipulation.
			πF.SetLineno(1)
			// line 1: """RFC 2822 message manipulation.
			πF.SetLineno(1)
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("RFC 2822 message manipulation.\n\nNote: This is only a very rough sketch of a full RFC-822 parser; in particular\nthe tokenizing of addresses does not adhere to all the quoting rules.\n\nNote: RFC 2822 is a long awaited update to RFC 822.  This module should\nconform to RFC 2822, and is thus mis-named (it's not worth renaming it).  Some\neffort at RFC 2822 updates have been made, but a thorough audit has not been\nperformed.  Consider any RFC 2822 non-conformance to be a bug.\n\n    RFC 2822: http://www.faqs.org/rfcs/rfc2822.html\n    RFC 822 : http://www.faqs.org/rfcs/rfc822.html (obsolete)\n\nDirections for use:\n\nTo create a Message object: first open a file, e.g.:\n\n  fp = open(file, 'r')\n\nYou can use any other legal way of getting an open file object, e.g. use\nsys.stdin or call os.popen().  Then pass the open file object to the Message()\nconstructor:\n\n  m = Message(fp)\n\nThis class can work with any input object that supports a readline method.  If\nthe input object has seek and tell capability, the rewindbody method will\nwork; also illegal lines will be pushed back onto the input stream.  If the\ninput object lacks seek but has an `unread' method that can push back a line\nof input, Message will use that to push back illegal lines.  Thus this class\ncan be used to parse messages coming from a buffered stream.\n\nThe optional `seekable' argument is provided as a workaround for certain stdio\nlibraries in which tell() discards buffered data before discovering that the\nlseek() system call doesn't work.  For maximum portability, you should set the\nseekable argument to zero to prevent that initial \\code{tell} when passing in\nan unseekable object such as a file object created from a socket object.  If\nit is 1 on entry -- which it is by default -- the tell() method of the open\nfile object is called once; if this raises an exception, seekable is reset to\n0.  For other nonzero values of seekable, this test is not made.\n\nTo get the text of a particular header there are several methods:\n\n  str = m.getheader(name)\n  str = m.getrawheader(name)\n\nwhere name is the name of the header, e.g. 'Subject'.  The difference is that\ngetheader() strips the leading and trailing whitespace, while getrawheader()\ndoesn't.  Both functions retain embedded whitespace (including newlines)\nexactly as they are specified in the header, and leave the case of the text\nunchanged.\n\nFor addresses and address lists there are functions\n\n  realname, mailaddress = m.getaddr(name)\n  list = m.getaddrlist(name)\n\nwhere the latter returns a list of (realname, mailaddr) tuples.\n\nThere is also a method\n\n  time = m.getdate(name)\n\nwhich parses a Date-like field and returns a time-compatible tuple,\ni.e. a tuple such as returned by time.localtime() or accepted by\ntime.mktime().\n\nSee the class definition for lower level access methods.\n\nThere are also some utility functions here.\n").ToObject()); πE != nil {
				continue
			}
			// line 74: import time
			πF.SetLineno(74)
			if πTemp002, πE = πg.ImportModule(πF, "time"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßtime.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 76: from warnings import warnpy3k
			πF.SetLineno(76)
			if πTemp002, πE = πg.ImportModule(πF, "warnings"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßwarnpy3k); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßwarnpy3k.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 77: warnpy3k("in 3.x, rfc822 has been removed in favor of the email package",
			πF.SetLineno(77)
			πTemp002 = πF.MakeArgs(1)
			πTemp002[0] = πg.NewStr("in 3.x, rfc822 has been removed in favor of the email package").ToObject()
			πTemp004 = πg.KWArgs{
				{"stacklevel", πg.NewInt(2).ToObject()},
			}
			if πTemp001, πE = πg.ResolveGlobal(πF, ßwarnpy3k); πE != nil {
				continue
			}
			if πTemp003, πE = πTemp001.Call(πF, πTemp002, πTemp004); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			// line 80: __all__ = ["Message","AddressList","parsedate","parsedate_tz","mktime_tz"]
			πF.SetLineno(80)
			πTemp002 = make([]*πg.Object, 5)
			πTemp002[0] = ßMessage.ToObject()
			πTemp002[1] = ßAddressList.ToObject()
			πTemp002[2] = ßparsedate.ToObject()
			πTemp002[3] = ßparsedate_tz.ToObject()
			πTemp002[4] = ßmktime_tz.ToObject()
			πTemp001 = πg.NewList(πTemp002...).ToObject()
			if πE = πF.Globals().SetItem(πF, ß__all__.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 82: _blanklines = ('\r\n', '\n')            # Optimization for islast()
			πF.SetLineno(82)
			πTemp001 = πg.NewTuple2(πg.NewStr("\r\n").ToObject(), πg.NewStr("\n").ToObject()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_blanklines.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 85: class Message(object):
			πF.SetLineno(85)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp006, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp006
			πTemp005 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp005.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("Message", "/usr/lib/python2.7/rfc822.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp005
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 86: """Represents a single RFC 2822-compliant message."""
					πF.SetLineno(86)
					// line 86: """Represents a single RFC 2822-compliant message."""
					πF.SetLineno(86)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Represents a single RFC 2822-compliant message.").ToObject()); πE != nil {
						continue
					}
					// line 88: def __init__(self, fp, seekable = 1):
					πF.SetLineno(88)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "fp", Def: nil}
					πTemp002[2] = πg.Param{Name: "seekable", Def: πg.NewInt(1).ToObject()}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/rfc822.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µfp *πg.Object = πArgs[1]; _ = µfp
						var µseekable *πg.Object = πArgs[2]; _ = µseekable
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.BaseException
						_ = πTemp004
						var πTemp005 *πg.Traceback
						_ = πTemp005
						var πTemp006 *πg.Object
						_ = πTemp006
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 9: goto Label9
							case 4: goto Label4
							case 14: goto Label14
							default: panic("unexpected function state")
							}
							// line 89: """Initialize the class instance and read the headers."""
							πF.SetLineno(89)
							if πE = πg.CheckLocal(πF, µseekable, "seekable"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, µseekable, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							goto Label2
							// line 90: if seekable == 1:
							πF.SetLineno(90)
						Label1:
							// line 93: try:
							πF.SetLineno(93)
							πF.PushCheckpoint(4)
							// line 94: fp.tell()
							πF.SetLineno(94)
							if πE = πg.CheckLocal(πF, µfp, "fp"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µfp, ßtell, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πF.PopCheckpoint()
							goto Label3
						Label4:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp004, πTemp005 = πF.ExcInfo()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßAttributeError); πE != nil {
								continue
							}
							if πTemp006, πE = πg.ResolveGlobal(πF, ßIOError); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple2(πTemp003, πTemp006).ToObject()
							if πTemp002, πE = πg.IsInstance(πF, πTemp004.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label5
							}
							πE = πF.Raise(πTemp004.ToObject(), nil, πTemp005.ToObject())
							continue
							// line 95: except (AttributeError, IOError):
							πF.SetLineno(95)
						Label5:
							// line 96: seekable = 0
							πF.SetLineno(96)
							µseekable = πg.NewInt(0).ToObject()
							πF.RestoreExc(nil, nil)
							goto Label3
						Label3:
							goto Label2
						Label2:
							// line 97: self.fp = fp
							πF.SetLineno(97)
							if πE = πg.CheckLocal(πF, µfp, "fp"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µfp); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßfp, πTemp001); πE != nil {
								continue
							}
							// line 98: self.seekable = seekable
							πF.SetLineno(98)
							if πE = πg.CheckLocal(πF, µseekable, "seekable"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µseekable); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßseekable, πTemp001); πE != nil {
								continue
							}
							// line 99: self.startofheaders = None
							πF.SetLineno(99)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßstartofheaders, πTemp003); πE != nil {
								continue
							}
							// line 100: self.startofbody = None
							πF.SetLineno(100)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßstartofbody, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßseekable, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label6
							}
							goto Label7
							// line 102: if self.seekable:
							πF.SetLineno(102)
						Label6:
							// line 103: try:
							πF.SetLineno(103)
							πF.PushCheckpoint(9)
							// line 104: self.startofheaders = self.fp.tell()
							πF.SetLineno(104)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßfp, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßtell, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßstartofheaders, πTemp003); πE != nil {
								continue
							}
							πF.PopCheckpoint()
							goto Label8
						Label9:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp004, πTemp005 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßIOError); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsInstance(πF, πTemp004.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label10
							}
							πE = πF.Raise(πTemp004.ToObject(), nil, πTemp005.ToObject())
							continue
							// line 105: except IOError:
							πF.SetLineno(105)
						Label10:
							// line 106: self.seekable = 0
							πF.SetLineno(106)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßseekable, πTemp001); πE != nil {
								continue
							}
							πF.RestoreExc(nil, nil)
							goto Label8
						Label8:
							goto Label7
						Label7:
							// line 108: self.readheaders()
							πF.SetLineno(108)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßreadheaders, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßseekable, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label11
							}
							goto Label12
							// line 110: if self.seekable:
							πF.SetLineno(110)
						Label11:
							// line 111: try:
							πF.SetLineno(111)
							πF.PushCheckpoint(14)
							// line 112: self.startofbody = self.fp.tell()
							πF.SetLineno(112)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßfp, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßtell, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßstartofbody, πTemp003); πE != nil {
								continue
							}
							πF.PopCheckpoint()
							goto Label13
						Label14:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp004, πTemp005 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßIOError); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsInstance(πF, πTemp004.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label15
							}
							πE = πF.Raise(πTemp004.ToObject(), nil, πTemp005.ToObject())
							continue
							// line 113: except IOError:
							πF.SetLineno(113)
						Label15:
							// line 114: self.seekable = 0
							πF.SetLineno(114)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßseekable, πTemp001); πE != nil {
								continue
							}
							πF.RestoreExc(nil, nil)
							goto Label13
						Label13:
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
					if πE = πClass.SetItem(πF, ß__init__.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 89: """Initialize the class instance and read the headers."""
					πF.SetLineno(89)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewStr("Initialize the class instance and read the headers.").ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ß__init__); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp004, ß__doc__, πTemp003); πE != nil {
						continue
					}
					// line 116: def rewindbody(self):
					πF.SetLineno(116)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("rewindbody", "/usr/lib/python2.7/rfc822.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
							// line 117: """Rewind the file to the start of the body (if seekable)."""
							πF.SetLineno(117)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßseekable, nil); πE != nil {
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
							// line 118: if not self.seekable:
							πF.SetLineno(118)
						Label1:
							if πTemp001, πE = πg.ResolveGlobal(πF, ßIOError); πE != nil {
								continue
							}
							// line 119: raise IOError, "unseekable file"
							πF.SetLineno(119)
							πE = πF.Raise(πTemp001, πg.NewStr("unseekable file").ToObject(), nil)
							continue
							goto Label2
						Label2:
							// line 120: self.fp.seek(self.startofbody)
							πF.SetLineno(120)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstartofbody, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßfp, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßseek, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßrewindbody.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 117: """Rewind the file to the start of the body (if seekable)."""
					πF.SetLineno(117)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πg.NewStr("Rewind the file to the start of the body (if seekable).").ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßrewindbody); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp005, ß__doc__, πTemp004); πE != nil {
						continue
					}
					// line 122: def readheaders(self):
					πF.SetLineno(122)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("readheaders", "/usr/lib/python2.7/rfc822.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µlst *πg.Object = πg.UnboundLocal; _ = µlst
						var µheaderseen *πg.Object = πg.UnboundLocal; _ = µheaderseen
						var µfirstline *πg.Object = πg.UnboundLocal; _ = µfirstline
						var µstartofline *πg.Object = πg.UnboundLocal; _ = µstartofline
						var µunread *πg.Object = πg.UnboundLocal; _ = µunread
						var µtell *πg.Object = πg.UnboundLocal; _ = µtell
						var µline *πg.Object = πg.UnboundLocal; _ = µline
						var µx *πg.Object = πg.UnboundLocal; _ = µx
						var πTemp001 *πg.Dict
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
						var πTemp007 *πg.BaseException
						_ = πTemp007
						var πTemp008 *πg.Traceback
						_ = πTemp008
						var πTemp009 *πg.Object
						_ = πTemp009
						var πTemp010 *πg.Object
						_ = πTemp010
						var πTemp011 bool
						_ = πTemp011
						var πTemp012 *πg.Object
						_ = πTemp012
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 10: goto Label10
							case 4: goto Label4
							case 5: goto Label5
							default: panic("unexpected function state")
							}
							// line 123: """Read header lines.
							πF.SetLineno(123)
							// line 137: self.dict = {}
							πF.SetLineno(137)
							πTemp001 = πg.NewDict()
							πTemp002 = πTemp001.ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßdict, πTemp003); πE != nil {
								continue
							}
							// line 138: self.unixfrom = ''
							πF.SetLineno(138)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, ß.ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßunixfrom, πTemp002); πE != nil {
								continue
							}
							// line 139: self.headers = lst = []
							πF.SetLineno(139)
							πTemp004 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp004...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßheaders, πTemp003); πE != nil {
								continue
							}
							µlst = πTemp002
							// line 140: self.status = ''
							πF.SetLineno(140)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, ß.ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßstatus, πTemp002); πE != nil {
								continue
							}
							// line 141: headerseen = ""
							πF.SetLineno(141)
							µheaderseen = ß.ToObject()
							// line 142: firstline = 1
							πF.SetLineno(142)
							µfirstline = πg.NewInt(1).ToObject()
							// line 143: startofline = unread = tell = None
							πF.SetLineno(143)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							µstartofline = πTemp002
							µunread = πTemp002
							µtell = πTemp002
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßfp, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp002
							πTemp004[1] = ßunread.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßseekable, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label2
							}
							goto Label3
							// line 144: if hasattr(self.fp, 'unread'):
							πF.SetLineno(144)
						Label1:
							// line 145: unread = self.fp.unread
							πF.SetLineno(145)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßfp, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßunread, nil); πE != nil {
								continue
							}
							µunread = πTemp003
							goto Label3
							// line 146: elif self.seekable:
							πF.SetLineno(146)
						Label2:
							// line 147: tell = self.fp.tell
							πF.SetLineno(147)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßfp, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßtell, nil); πE != nil {
								continue
							}
							µtell = πTemp003
							goto Label3
						Label3:
							// line 148: while 1:
							πF.SetLineno(148)
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
							if πTemp006, πE = πg.IsTrue(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(4)            
							if πE = πg.CheckLocal(πF, µtell, "tell"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, µtell); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label7
							}
							goto Label8
							// line 149: if tell:
							πF.SetLineno(149)
						Label7:
							// line 150: try:
							πF.SetLineno(150)
							πF.PushCheckpoint(10)
							// line 151: startofline = tell()
							πF.SetLineno(151)
							if πE = πg.CheckLocal(πF, µtell, "tell"); πE != nil {
								continue
							}
							if πTemp002, πE = µtell.Call(πF, nil, nil); πE != nil {
								continue
							}
							µstartofline = πTemp002
							πF.PopCheckpoint()
							goto Label9
						Label10:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp007, πTemp008 = πF.ExcInfo()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßIOError); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsInstance(πF, πTemp007.ToObject(), πTemp002); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label11
							}
							πE = πF.Raise(πTemp007.ToObject(), nil, πTemp008.ToObject())
							continue
							// line 152: except IOError:
							πF.SetLineno(152)
						Label11:
							// line 153: startofline = tell = None
							πF.SetLineno(153)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							µstartofline = πTemp002
							µtell = πTemp002
							// line 154: self.seekable = 0
							πF.SetLineno(154)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßseekable, πTemp002); πE != nil {
								continue
							}
							πF.RestoreExc(nil, nil)
							goto Label9
						Label9:
							goto Label8
						Label8:
							// line 155: line = self.fp.readline()
							πF.SetLineno(155)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßfp, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßreadline, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							µline = πTemp002
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, µline); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp006).ToObject()
							if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label12
							}
							goto Label13
							// line 156: if not line:
							πF.SetLineno(156)
						Label12:
							// line 157: self.status = 'EOF in headers'
							πF.SetLineno(157)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewStr("EOF in headers").ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßstatus, πTemp002); πE != nil {
								continue
							}
							// line 158: break
							πF.SetLineno(158)
							πTemp005 = true
							continue
							goto Label13
						Label13:
							if πE = πg.CheckLocal(πF, µfirstline, "firstline"); πE != nil {
								continue
							}
							πTemp002 = µfirstline
							if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp006 {
								goto Label14
							}
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = πg.NewStr("From ").ToObject()
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µline, ßstartswith, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp002 = πTemp009
						Label14:
							if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label15
							}
							goto Label16
							// line 160: if firstline and line.startswith('From '):
							πF.SetLineno(160)
						Label15:
							// line 161: self.unixfrom = self.unixfrom + line
							πF.SetLineno(161)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßunixfrom, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp003, µline); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßunixfrom, πTemp003); πE != nil {
								continue
							}
							// line 162: continue
							πF.SetLineno(162)
							continue
							goto Label16
						Label16:
							// line 163: firstline = 0
							πF.SetLineno(163)
							µfirstline = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µheaderseen, "headerseen"); πE != nil {
								continue
							}
							πTemp002 = µheaderseen
							if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp006 {
								goto Label17
							}
							πTemp009 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetItem(πF, µline, πTemp009); πE != nil {
								continue
							}
							if πTemp011, πE = πg.Contains(πF, πg.NewStr(" \t").ToObject(), πTemp010); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(πTemp011).ToObject()
							πTemp002 = πTemp003
						Label17:
							if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label18
							}
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							πTemp004[0] = µline
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßiscomment, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label19
							}
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							πTemp004[0] = µline
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßislast, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label20
							}
							goto Label21
							// line 164: if headerseen and line[0] in ' \t':
							πF.SetLineno(164)
						Label18:
							// line 166: lst.append(line)
							πF.SetLineno(166)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							πTemp004[0] = µline
							if πE = πg.CheckLocal(πF, µlst, "lst"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µlst, ßappend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 167: x = (self.dict[headerseen] + "\n " + line.strip())
							πF.SetLineno(167)
							if πE = πg.CheckLocal(πF, µheaderseen, "headerseen"); πE != nil {
								continue
							}
							πTemp009 = µheaderseen
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp012, πE = πg.GetAttr(πF, µself, ßdict, nil); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetItem(πF, πTemp012, πTemp009); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, πTemp010, πg.NewStr("\n ").ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, µline, ßstrip, nil); πE != nil {
								continue
							}
							if πTemp010, πE = πTemp009.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp003, πTemp010); πE != nil {
								continue
							}
							µx = πTemp002
							// line 168: self.dict[headerseen] = x.strip()
							πF.SetLineno(168)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µx, ßstrip, nil); πE != nil {
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
							if πTemp009, πE = πg.GetAttr(πF, µself, ßdict, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µheaderseen, "headerseen"); πE != nil {
								continue
							}
							πTemp010 = µheaderseen
							if πE = πg.SetItem(πF, πTemp009, πTemp010, πTemp002); πE != nil {
								continue
							}
							// line 169: continue
							πF.SetLineno(169)
							continue
							goto Label21
							// line 170: elif self.iscomment(line):
							πF.SetLineno(170)
						Label19:
							// line 172: continue
							πF.SetLineno(172)
							continue
							goto Label21
							// line 173: elif self.islast(line):
							πF.SetLineno(173)
						Label20:
							// line 175: break
							πF.SetLineno(175)
							πTemp005 = true
							continue
							goto Label21
						Label21:
							// line 176: headerseen = self.isheader(line)
							πF.SetLineno(176)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							πTemp004[0] = µline
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßisheader, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µheaderseen = πTemp003
							if πE = πg.CheckLocal(πF, µheaderseen, "headerseen"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, µheaderseen); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label22
							}
							if πE = πg.CheckLocal(πF, µheaderseen, "headerseen"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(µheaderseen != πTemp003).ToObject()
							if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label23
							}
							goto Label24
							// line 177: if headerseen:
							πF.SetLineno(177)
						Label22:
							// line 179: lst.append(line)
							πF.SetLineno(179)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							πTemp004[0] = µline
							if πE = πg.CheckLocal(πF, µlst, "lst"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µlst, ßappend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 180: self.dict[headerseen] = line[len(headerseen)+1:].strip()
							πF.SetLineno(180)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µheaderseen, "headerseen"); πE != nil {
								continue
							}
							πTemp004[0] = µheaderseen
							if πTemp009, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp010, πE = πTemp009.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp003, πE = πg.Add(πF, πTemp010, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πTemp003, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µline, πTemp002); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßstrip, nil); πE != nil {
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
							if πTemp009, πE = πg.GetAttr(πF, µself, ßdict, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µheaderseen, "headerseen"); πE != nil {
								continue
							}
							πTemp010 = µheaderseen
							if πE = πg.SetItem(πF, πTemp009, πTemp010, πTemp002); πE != nil {
								continue
							}
							// line 181: continue
							πF.SetLineno(181)
							continue
							goto Label25
							// line 182: elif headerseen is not None:
							πF.SetLineno(182)
						Label23:
							// line 186: continue
							πF.SetLineno(186)
							continue
							goto Label25
						Label24:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdict, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp006).ToObject()
							if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label26
							}
							goto Label27
							// line 189: if not self.dict:
							πF.SetLineno(189)
						Label26:
							// line 190: self.status = 'No headers'
							πF.SetLineno(190)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewStr("No headers").ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßstatus, πTemp002); πE != nil {
								continue
							}
							goto Label28
						Label27:
							// line 192: self.status = 'Non-header line where header expected'
							πF.SetLineno(192)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewStr("Non-header line where header expected").ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßstatus, πTemp002); πE != nil {
								continue
							}
							goto Label28
						Label28:
							if πE = πg.CheckLocal(πF, µunread, "unread"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, µunread); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label29
							}
							if πE = πg.CheckLocal(πF, µtell, "tell"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, µtell); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label30
							}
							goto Label31
							// line 194: if unread:
							πF.SetLineno(194)
						Label29:
							// line 195: unread(line)
							πF.SetLineno(195)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							πTemp004[0] = µline
							if πE = πg.CheckLocal(πF, µunread, "unread"); πE != nil {
								continue
							}
							if πTemp002, πE = µunread.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label32
							// line 196: elif tell:
							πF.SetLineno(196)
						Label30:
							// line 197: self.fp.seek(startofline)
							πF.SetLineno(197)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µstartofline, "startofline"); πE != nil {
								continue
							}
							πTemp004[0] = µstartofline
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßfp, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßseek, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label32
						Label31:
							// line 199: self.status = self.status + '; bad seek'
							πF.SetLineno(199)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßstatus, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp003, πg.NewStr("; bad seek").ToObject()); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßstatus, πTemp003); πE != nil {
								continue
							}
							goto Label32
						Label32:
							// line 200: break
							πF.SetLineno(200)
							πTemp005 = true
							continue
							goto Label25
						Label25:
							continue
						Label5:
							if πE != nil || πR != nil {
								continue
							}
						Label6:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßreadheaders.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 123: """Read header lines.
					πF.SetLineno(123)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πg.NewStr("Read header lines.\n\n        Read header lines up to the entirely blank line that terminates them.\n        The (normally blank) line that ends the headers is skipped, but not\n        included in the returned list.  If a non-header line ends the headers,\n        (which is an error), an attempt is made to backspace over it; it is\n        never included in the returned list.\n\n        The variable self.status is set to the empty string if all went well,\n        otherwise it is an error message.  The variable self.headers is a\n        completely uninterpreted list of lines contained in the header (so\n        printing them will reproduce the header exactly as it appears in the\n        file).\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßreadheaders); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp006, ß__doc__, πTemp005); πE != nil {
						continue
					}
					// line 202: def isheader(self, line):
					πF.SetLineno(202)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "line", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("isheader", "/usr/lib/python2.7/rfc822.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µline *πg.Object = πArgs[1]; _ = µline
						var µi *πg.Object = πg.UnboundLocal; _ = µi
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
							// line 203: """Determine whether a given line is a legal header.
							πF.SetLineno(203)
							// line 209: i = line.find(':')
							πF.SetLineno(209)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr(":").ToObject()
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µline, ßfind, nil); πE != nil {
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
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GT(πF, µi, πTemp003); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 210: if i > -1:
							πF.SetLineno(210)
						Label1:
							// line 211: return line[:i].lower()
							πF.SetLineno(211)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.None, µi, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µline, πTemp002); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßlower, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πR = πTemp003
							continue
							goto Label2
						Label2:
							// line 212: return None
							πF.SetLineno(212)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
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
					if πE = πClass.SetItem(πF, ßisheader.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 203: """Determine whether a given line is a legal header.
					πF.SetLineno(203)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, πg.NewStr("Determine whether a given line is a legal header.\n\n        This method should return the header name, suitably canonicalized.\n        You may override this method in order to use Message parsing on tagged\n        data in RFC 2822-like formats with special header formats.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßisheader); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp007, ß__doc__, πTemp006); πE != nil {
						continue
					}
					// line 214: def islast(self, line):
					πF.SetLineno(214)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "line", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("islast", "/usr/lib/python2.7/rfc822.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µline *πg.Object = πArgs[1]; _ = µline
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
							// line 215: """Determine whether a line is a legal end of RFC 2822 headers.
							πF.SetLineno(215)
							// line 222: return line in _blanklines
							πF.SetLineno(222)
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_blanklines); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Contains(πF, πTemp002, µline); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp003).ToObject()
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
					if πE = πClass.SetItem(πF, ßislast.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 215: """Determine whether a line is a legal end of RFC 2822 headers.
					πF.SetLineno(215)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp007}, πg.NewStr("Determine whether a line is a legal end of RFC 2822 headers.\n\n        You may override this method if your application wants to bend the\n        rules, e.g. to strip trailing whitespace, or to recognize MH template\n        separators ('--------').  For convenience (e.g. for code reading from\n        sockets) a line consisting of \\r\\n also matches.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ßislast); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp008, ß__doc__, πTemp007); πE != nil {
						continue
					}
					// line 224: def iscomment(self, line):
					πF.SetLineno(224)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "line", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("iscomment", "/usr/lib/python2.7/rfc822.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µline *πg.Object = πArgs[1]; _ = µline
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 225: """Determine whether a line should be skipped entirely.
							πF.SetLineno(225)
							// line 231: return False
							πF.SetLineno(231)
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
					if πE = πClass.SetItem(πF, ßiscomment.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 225: """Determine whether a line should be skipped entirely.
					πF.SetLineno(225)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp008}, πg.NewStr("Determine whether a line should be skipped entirely.\n\n        You may override this method in order to use Message parsing on tagged\n        data in RFC 2822-like formats that support embedded comments or\n        free-text data.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßiscomment); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp009, ß__doc__, πTemp008); πE != nil {
						continue
					}
					// line 233: def getallmatchingheaders(self, name):
					πF.SetLineno(233)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "name", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("getallmatchingheaders", "/usr/lib/python2.7/rfc822.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µname *πg.Object = πArgs[1]; _ = µname
						var µn *πg.Object = πg.UnboundLocal; _ = µn
						var µlst *πg.Object = πg.UnboundLocal; _ = µlst
						var µhit *πg.Object = πg.UnboundLocal; _ = µhit
						var µline *πg.Object = πg.UnboundLocal; _ = µline
						var πTemp001 *πg.Object
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							// line 234: """Find all header lines matching a given header name.
							πF.SetLineno(234)
							// line 242: name = name.lower() + ':'
							πF.SetLineno(242)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µname, ßlower, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp003, πg.NewStr(":").ToObject()); πE != nil {
								continue
							}
							µname = πTemp001
							// line 243: n = len(name)
							πF.SetLineno(243)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp004[0] = µname
							if πTemp001, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µn = πTemp002
							// line 244: lst = []
							πF.SetLineno(244)
							πTemp004 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp004...).ToObject()
							µlst = πTemp001
							// line 245: hit = 0
							πF.SetLineno(245)
							µhit = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßheaders, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, πTemp002); πE != nil {
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
								µline = πTemp002
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(1)            
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.None, µn, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, µline, πTemp003); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp007, ßlower, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, πTemp007, µname); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label4
							}
							if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.NewInt(1).ToObject(), πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, µline, πTemp003); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp007, ßisspace, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp007); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp006).ToObject()
							if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label5
							}
							goto Label6
							// line 247: if line[:n].lower() == name:
							πF.SetLineno(247)
						Label4:
							// line 248: hit = 1
							πF.SetLineno(248)
							µhit = πg.NewInt(1).ToObject()
							goto Label6
							// line 249: elif not line[:1].isspace():
							πF.SetLineno(249)
						Label5:
							// line 250: hit = 0
							πF.SetLineno(250)
							µhit = πg.NewInt(0).ToObject()
							goto Label6
						Label6:
							if πE = πg.CheckLocal(πF, µhit, "hit"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, µhit); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label7
							}
							goto Label8
							// line 251: if hit:
							πF.SetLineno(251)
						Label7:
							// line 252: lst.append(line)
							πF.SetLineno(252)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							πTemp004[0] = µline
							if πE = πg.CheckLocal(πF, µlst, "lst"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µlst, ßappend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label8
						Label8:
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 253: return lst
							πF.SetLineno(253)
							if πE = πg.CheckLocal(πF, µlst, "lst"); πE != nil {
								continue
							}
							πR = µlst
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßgetallmatchingheaders.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 234: """Find all header lines matching a given header name.
					πF.SetLineno(234)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp009}, πg.NewStr("Find all header lines matching a given header name.\n\n        Look through the list of headers and find all lines matching a given\n        header name (and their continuation lines).  A list of the lines is\n        returned, without interpretation.  If the header does not occur, an\n        empty list is returned.  If the header occurs multiple times, all\n        occurrences are returned.  Case is not important in the header name.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßgetallmatchingheaders); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp010, ß__doc__, πTemp009); πE != nil {
						continue
					}
					// line 255: def getfirstmatchingheader(self, name):
					πF.SetLineno(255)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "name", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("getfirstmatchingheader", "/usr/lib/python2.7/rfc822.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µname *πg.Object = πArgs[1]; _ = µname
						var µn *πg.Object = πg.UnboundLocal; _ = µn
						var µlst *πg.Object = πg.UnboundLocal; _ = µlst
						var µhit *πg.Object = πg.UnboundLocal; _ = µhit
						var µline *πg.Object = πg.UnboundLocal; _ = µline
						var πTemp001 *πg.Object
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							// line 256: """Get the first header line matching name.
							πF.SetLineno(256)
							// line 261: name = name.lower() + ':'
							πF.SetLineno(261)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µname, ßlower, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp003, πg.NewStr(":").ToObject()); πE != nil {
								continue
							}
							µname = πTemp001
							// line 262: n = len(name)
							πF.SetLineno(262)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp004[0] = µname
							if πTemp001, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µn = πTemp002
							// line 263: lst = []
							πF.SetLineno(263)
							πTemp004 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp004...).ToObject()
							µlst = πTemp001
							// line 264: hit = 0
							πF.SetLineno(264)
							µhit = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßheaders, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, πTemp002); πE != nil {
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
								µline = πTemp002
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(1)            
							if πE = πg.CheckLocal(πF, µhit, "hit"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, µhit); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label4
							}
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.None, µn, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, µline, πTemp003); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp007, ßlower, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, πTemp007, µname); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label5
							}
							goto Label6
							// line 266: if hit:
							πF.SetLineno(266)
						Label4:
							if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.NewInt(1).ToObject(), πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, µline, πTemp003); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp007, ßisspace, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp007); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp006).ToObject()
							if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label7
							}
							goto Label8
							// line 267: if not line[:1].isspace():
							πF.SetLineno(267)
						Label7:
							// line 268: break
							πF.SetLineno(268)
							πTemp005 = true
							continue
							goto Label8
						Label8:
							goto Label6
							// line 269: elif line[:n].lower() == name:
							πF.SetLineno(269)
						Label5:
							// line 270: hit = 1
							πF.SetLineno(270)
							µhit = πg.NewInt(1).ToObject()
							goto Label6
						Label6:
							if πE = πg.CheckLocal(πF, µhit, "hit"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, µhit); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label9
							}
							goto Label10
							// line 271: if hit:
							πF.SetLineno(271)
						Label9:
							// line 272: lst.append(line)
							πF.SetLineno(272)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							πTemp004[0] = µline
							if πE = πg.CheckLocal(πF, µlst, "lst"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µlst, ßappend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label10
						Label10:
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 273: return lst
							πF.SetLineno(273)
							if πE = πg.CheckLocal(πF, µlst, "lst"); πE != nil {
								continue
							}
							πR = µlst
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßgetfirstmatchingheader.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 256: """Get the first header line matching name.
					πF.SetLineno(256)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp010}, πg.NewStr("Get the first header line matching name.\n\n        This is similar to getallmatchingheaders, but it returns only the\n        first matching header (and its continuation lines).\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ßgetfirstmatchingheader); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp011, ß__doc__, πTemp010); πE != nil {
						continue
					}
					// line 275: def getrawheader(self, name):
					πF.SetLineno(275)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "name", Def: nil}
					πTemp010 = πg.NewFunction(πg.NewCode("getrawheader", "/usr/lib/python2.7/rfc822.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µname *πg.Object = πArgs[1]; _ = µname
						var µlst *πg.Object = πg.UnboundLocal; _ = µlst
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
							// line 276: """A higher-level interface to getfirstmatchingheader().
							πF.SetLineno(276)
							// line 284: lst = self.getfirstmatchingheader(name)
							πF.SetLineno(284)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp001[0] = µname
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßgetfirstmatchingheader, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µlst = πTemp003
							if πE = πg.CheckLocal(πF, µlst, "lst"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µlst); πE != nil {
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
							// line 285: if not lst:
							πF.SetLineno(285)
						Label1:
							// line 286: return None
							πF.SetLineno(286)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πR = πTemp002
							continue
							goto Label2
						Label2:
							// line 287: lst[0] = lst[0][len(name) + 1:]
							πF.SetLineno(287)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp001[0] = µname
							if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp003, πE = πg.Add(πF, πTemp006, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πTemp003, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							πTemp005 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µlst, "lst"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µlst, πTemp005); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp006, πTemp002); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µlst, "lst"); πE != nil {
								continue
							}
							πTemp005 = πg.NewInt(0).ToObject()
							if πE = πg.SetItem(πF, µlst, πTemp005, πTemp002); πE != nil {
								continue
							}
							// line 288: return ''.join(lst)
							πF.SetLineno(288)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µlst, "lst"); πE != nil {
								continue
							}
							πTemp001[0] = µlst
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
					if πE = πClass.SetItem(πF, ßgetrawheader.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 276: """A higher-level interface to getfirstmatchingheader().
					πF.SetLineno(276)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πg.NewStr("A higher-level interface to getfirstmatchingheader().\n\n        Return a string containing the literal text of the header but with the\n        keyword stripped.  All leading, trailing and embedded whitespace is\n        kept in the string, however.  Return None if the header does not\n        occur.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp012, πE = πg.ResolveClass(πF, πClass, nil, ßgetrawheader); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp012, ß__doc__, πTemp011); πE != nil {
						continue
					}
					// line 290: def getheader(self, name, default=None):
					πF.SetLineno(290)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "name", Def: nil}
					if πTemp012, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "default", Def: πTemp012}
					πTemp011 = πg.NewFunction(πg.NewCode("getheader", "/usr/lib/python2.7/rfc822.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µname *πg.Object = πArgs[1]; _ = µname
						var µdefault *πg.Object = πArgs[2]; _ = µdefault
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
							// line 291: """Get the header value for a name.
							πF.SetLineno(291)
							// line 297: return self.dict.get(name.lower(), default)
							πF.SetLineno(297)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µname, ßlower, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µdefault, "default"); πE != nil {
								continue
							}
							πTemp001[1] = µdefault
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdict, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßget, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßgetheader.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 291: """Get the header value for a name.
					πF.SetLineno(291)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp012}, πg.NewStr("Get the header value for a name.\n\n        This is the normal interface: it returns a stripped version of the\n        header value for a given header name, or None if it doesn't exist.\n        This uses the dictionary version which finds the *last* such header.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp013, πE = πg.ResolveClass(πF, πClass, nil, ßgetheader); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp013, ß__doc__, πTemp012); πE != nil {
						continue
					}
					// line 298: get = getheader
					πF.SetLineno(298)
					if πTemp012, πE = πg.ResolveClass(πF, πClass, nil, ßgetheader); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßget.ToObject(), πTemp012); πE != nil {
						continue
					}
					// line 300: def getheaders(self, name):
					πF.SetLineno(300)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "name", Def: nil}
					πTemp012 = πg.NewFunction(πg.NewCode("getheaders", "/usr/lib/python2.7/rfc822.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µname *πg.Object = πArgs[1]; _ = µname
						var µresult *πg.Object = πg.UnboundLocal; _ = µresult
						var µcurrent *πg.Object = πg.UnboundLocal; _ = µcurrent
						var µhave_header *πg.Object = πg.UnboundLocal; _ = µhave_header
						var µs *πg.Object = πg.UnboundLocal; _ = µs
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							// line 301: """Get all values for a header.
							πF.SetLineno(301)
							// line 307: result = []
							πF.SetLineno(307)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µresult = πTemp002
							// line 308: current = ''
							πF.SetLineno(308)
							µcurrent = ß.ToObject()
							// line 309: have_header = 0
							πF.SetLineno(309)
							µhave_header = πg.NewInt(0).ToObject()
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp001[0] = µname
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßgetallmatchingheaders, nil); πE != nil {
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
								µs = πTemp003
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(1)            
							πTemp003 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µs, πTemp003); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßisspace, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label4
							}
							goto Label5
							// line 311: if s[0].isspace():
							πF.SetLineno(311)
						Label4:
							if πE = πg.CheckLocal(πF, µcurrent, "current"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, µcurrent); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label7
							}
							goto Label8
							// line 312: if current:
							πF.SetLineno(312)
						Label7:
							// line 313: current = "%s\n %s" % (current, s.strip())
							πF.SetLineno(313)
							if πE = πg.CheckLocal(πF, µcurrent, "current"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µs, ßstrip, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp007.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp004 = πg.NewTuple2(µcurrent, πTemp008).ToObject()
							if πTemp003, πE = πg.Mod(πF, πg.NewStr("%s\n %s").ToObject(), πTemp004); πE != nil {
								continue
							}
							µcurrent = πTemp003
							goto Label9
						Label8:
							// line 315: current = s.strip()
							πF.SetLineno(315)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µs, ßstrip, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							µcurrent = πTemp004
							goto Label9
						Label9:
							goto Label6
						Label5:
							if πE = πg.CheckLocal(πF, µhave_header, "have_header"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, µhave_header); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label10
							}
							goto Label11
							// line 317: if have_header:
							πF.SetLineno(317)
						Label10:
							// line 318: result.append(current)
							πF.SetLineno(318)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µcurrent, "current"); πE != nil {
								continue
							}
							πTemp001[0] = µcurrent
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µresult, ßappend, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label11
						Label11:
							// line 319: current = s[s.find(":") + 1:].strip()
							πF.SetLineno(319)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr(":").ToObject()
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µs, ßfind, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp007.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp004, πE = πg.Add(πF, πTemp008, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πTemp004, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µs, πTemp003); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßstrip, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							µcurrent = πTemp004
							// line 320: have_header = 1
							πF.SetLineno(320)
							µhave_header = πg.NewInt(1).ToObject()
							goto Label6
						Label6:
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							if πE = πg.CheckLocal(πF, µhave_header, "have_header"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, µhave_header); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label12
							}
							goto Label13
							// line 321: if have_header:
							πF.SetLineno(321)
						Label12:
							// line 322: result.append(current)
							πF.SetLineno(322)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µcurrent, "current"); πE != nil {
								continue
							}
							πTemp001[0] = µcurrent
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
							goto Label13
						Label13:
							// line 323: return result
							πF.SetLineno(323)
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
					if πE = πClass.SetItem(πF, ßgetheaders.ToObject(), πTemp012); πE != nil {
						continue
					}
					// line 301: """Get all values for a header.
					πF.SetLineno(301)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp013}, πg.NewStr("Get all values for a header.\n\n        This returns a list of values for headers given more than once; each\n        value in the result list is stripped in the same way as the result of\n        getheader().  If the header is not given, return an empty list.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp014, πE = πg.ResolveClass(πF, πClass, nil, ßgetheaders); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp014, ß__doc__, πTemp013); πE != nil {
						continue
					}
					// line 325: def getaddr(self, name):
					πF.SetLineno(325)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "name", Def: nil}
					πTemp013 = πg.NewFunction(πg.NewCode("getaddr", "/usr/lib/python2.7/rfc822.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µname *πg.Object = πArgs[1]; _ = µname
						var µalist *πg.Object = πg.UnboundLocal; _ = µalist
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 326: """Get a single address from a header, as a tuple.
							πF.SetLineno(326)
							// line 332: alist = self.getaddrlist(name)
							πF.SetLineno(332)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp001[0] = µname
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßgetaddrlist, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µalist = πTemp003
							if πE = πg.CheckLocal(πF, µalist, "alist"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µalist); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 333: if alist:
							πF.SetLineno(333)
						Label1:
							// line 334: return alist[0]
							πF.SetLineno(334)
							πTemp002 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µalist, "alist"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µalist, πTemp002); πE != nil {
								continue
							}
							πR = πTemp003
							continue
							goto Label3
						Label2:
							// line 336: return (None, None)
							πF.SetLineno(336)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(πTemp003, πTemp005).ToObject()
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
					if πE = πClass.SetItem(πF, ßgetaddr.ToObject(), πTemp013); πE != nil {
						continue
					}
					// line 326: """Get a single address from a header, as a tuple.
					πF.SetLineno(326)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp014}, πg.NewStr("Get a single address from a header, as a tuple.\n\n        An example return value:\n        ('Guido van Rossum', 'guido@cwi.nl')\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp015, πE = πg.ResolveClass(πF, πClass, nil, ßgetaddr); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp015, ß__doc__, πTemp014); πE != nil {
						continue
					}
					// line 338: def getaddrlist(self, name):
					πF.SetLineno(338)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "name", Def: nil}
					πTemp014 = πg.NewFunction(πg.NewCode("getaddrlist", "/usr/lib/python2.7/rfc822.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µname *πg.Object = πArgs[1]; _ = µname
						var µraw *πg.Object = πg.UnboundLocal; _ = µraw
						var µh *πg.Object = πg.UnboundLocal; _ = µh
						var µi *πg.Object = πg.UnboundLocal; _ = µi
						var µaddr *πg.Object = πg.UnboundLocal; _ = µaddr
						var µalladdrs *πg.Object = πg.UnboundLocal; _ = µalladdrs
						var µa *πg.Object = πg.UnboundLocal; _ = µa
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							// line 339: """Get a list of addresses from a header.
							πF.SetLineno(339)
							// line 345: raw = []
							πF.SetLineno(345)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µraw = πTemp002
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp001[0] = µname
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßgetallmatchingheaders, nil); πE != nil {
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
								µh = πTemp003
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(1)            
							πTemp004 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µh, "h"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, µh, πTemp004); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Contains(πF, πg.NewStr(" \t").ToObject(), πTemp007); πE != nil {
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
							// line 347: if h[0] in ' \t':
							πF.SetLineno(347)
						Label4:
							// line 348: raw.append(h)
							πF.SetLineno(348)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µh, "h"); πE != nil {
								continue
							}
							πTemp001[0] = µh
							if πE = πg.CheckLocal(πF, µraw, "raw"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µraw, ßappend, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label6
						Label5:
							if πE = πg.CheckLocal(πF, µraw, "raw"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, µraw); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label7
							}
							goto Label8
							// line 350: if raw:
							πF.SetLineno(350)
						Label7:
							// line 351: raw.append(', ')
							πF.SetLineno(351)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr(", ").ToObject()
							if πE = πg.CheckLocal(πF, µraw, "raw"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µraw, ßappend, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label8
						Label8:
							// line 352: i = h.find(':')
							πF.SetLineno(352)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr(":").ToObject()
							if πE = πg.CheckLocal(πF, µh, "h"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µh, ßfind, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µi = πTemp004
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GT(πF, µi, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label9
							}
							goto Label10
							// line 353: if i > 0:
							πF.SetLineno(353)
						Label9:
							// line 354: addr = h[i+1:]
							πF.SetLineno(354)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Add(πF, µi, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πTemp004, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µh, "h"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µh, πTemp003); πE != nil {
								continue
							}
							µaddr = πTemp004
							goto Label10
						Label10:
							// line 355: raw.append(addr)
							πF.SetLineno(355)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µaddr, "addr"); πE != nil {
								continue
							}
							πTemp001[0] = µaddr
							if πE = πg.CheckLocal(πF, µraw, "raw"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µraw, ßappend, nil); πE != nil {
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
							// line 356: alladdrs = ''.join(raw)
							πF.SetLineno(356)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µraw, "raw"); πE != nil {
								continue
							}
							πTemp001[0] = µraw
							if πTemp002, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µalladdrs = πTemp003
							// line 357: a = AddressList(alladdrs)
							πF.SetLineno(357)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µalladdrs, "alladdrs"); πE != nil {
								continue
							}
							πTemp001[0] = µalladdrs
							if πTemp002, πE = πg.ResolveGlobal(πF, ßAddressList); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µa = πTemp003
							// line 358: return a.addresslist
							πF.SetLineno(358)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µa, ßaddresslist, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßgetaddrlist.ToObject(), πTemp014); πE != nil {
						continue
					}
					// line 339: """Get a list of addresses from a header.
					πF.SetLineno(339)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp015}, πg.NewStr("Get a list of addresses from a header.\n\n        Retrieves a list of addresses from a header, where each address is a\n        tuple as returned by getaddr().  Scans all named headers, so it works\n        properly with multiple To: or Cc: headers for example.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp016, πE = πg.ResolveClass(πF, πClass, nil, ßgetaddrlist); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp016, ß__doc__, πTemp015); πE != nil {
						continue
					}
					// line 360: def getdate(self, name):
					πF.SetLineno(360)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "name", Def: nil}
					πTemp015 = πg.NewFunction(πg.NewCode("getdate", "/usr/lib/python2.7/rfc822.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µname *πg.Object = πArgs[1]; _ = µname
						var µdata *πg.Object = πg.UnboundLocal; _ = µdata
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
							// line 361: """Retrieve a date field from a header.
							πF.SetLineno(361)
							// line 366: try:
							πF.SetLineno(366)
							πF.PushCheckpoint(2)
							// line 367: data = self[name]
							πF.SetLineno(367)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp001 = µname
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µself, πTemp001); πE != nil {
								continue
							}
							µdata = πTemp002
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
							// line 368: except KeyError:
							πF.SetLineno(368)
						Label3:
							// line 369: return None
							πF.SetLineno(369)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							πF.RestoreExc(nil, nil)
							goto Label1
						Label1:
							// line 370: return parsedate(data)
							πF.SetLineno(370)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πTemp006[0] = µdata
							if πTemp001, πE = πg.ResolveGlobal(πF, ßparsedate); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
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
					if πE = πClass.SetItem(πF, ßgetdate.ToObject(), πTemp015); πE != nil {
						continue
					}
					// line 361: """Retrieve a date field from a header.
					πF.SetLineno(361)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp016}, πg.NewStr("Retrieve a date field from a header.\n\n        Retrieves a date field from the named header, returning a tuple\n        compatible with time.mktime().\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp017, πE = πg.ResolveClass(πF, πClass, nil, ßgetdate); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp017, ß__doc__, πTemp016); πE != nil {
						continue
					}
					// line 372: def getdate_tz(self, name):
					πF.SetLineno(372)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "name", Def: nil}
					πTemp016 = πg.NewFunction(πg.NewCode("getdate_tz", "/usr/lib/python2.7/rfc822.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µname *πg.Object = πArgs[1]; _ = µname
						var µdata *πg.Object = πg.UnboundLocal; _ = µdata
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
							// line 373: """Retrieve a date field from a header as a 10-tuple.
							πF.SetLineno(373)
							// line 378: try:
							πF.SetLineno(378)
							πF.PushCheckpoint(2)
							// line 379: data = self[name]
							πF.SetLineno(379)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp001 = µname
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µself, πTemp001); πE != nil {
								continue
							}
							µdata = πTemp002
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
							// line 380: except KeyError:
							πF.SetLineno(380)
						Label3:
							// line 381: return None
							πF.SetLineno(381)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							πF.RestoreExc(nil, nil)
							goto Label1
						Label1:
							// line 382: return parsedate_tz(data)
							πF.SetLineno(382)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πTemp006[0] = µdata
							if πTemp001, πE = πg.ResolveGlobal(πF, ßparsedate_tz); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
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
					if πE = πClass.SetItem(πF, ßgetdate_tz.ToObject(), πTemp016); πE != nil {
						continue
					}
					// line 373: """Retrieve a date field from a header as a 10-tuple.
					πF.SetLineno(373)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp017}, πg.NewStr("Retrieve a date field from a header as a 10-tuple.\n\n        The first 9 elements make up a tuple compatible with time.mktime(),\n        and the 10th is the offset of the poster's time zone from GMT/UTC.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp018, πE = πg.ResolveClass(πF, πClass, nil, ßgetdate_tz); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp018, ß__doc__, πTemp017); πE != nil {
						continue
					}
					// line 387: def __len__(self):
					πF.SetLineno(387)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp017 = πg.NewFunction(πg.NewCode("__len__", "/usr/lib/python2.7/rfc822.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 388: """Get the number of headers in a message."""
							πF.SetLineno(388)
							// line 389: return len(self.dict)
							πF.SetLineno(389)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdict, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__len__.ToObject(), πTemp017); πE != nil {
						continue
					}
					// line 388: """Get the number of headers in a message."""
					πF.SetLineno(388)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp018}, πg.NewStr("Get the number of headers in a message.").ToObject()); πE != nil {
						continue
					}
					if πTemp019, πE = πg.ResolveClass(πF, πClass, nil, ß__len__); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp019, ß__doc__, πTemp018); πE != nil {
						continue
					}
					// line 391: def __getitem__(self, name):
					πF.SetLineno(391)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "name", Def: nil}
					πTemp018 = πg.NewFunction(πg.NewCode("__getitem__", "/usr/lib/python2.7/rfc822.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µname *πg.Object = πArgs[1]; _ = µname
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
							// line 392: """Get a specific header, as from a dictionary."""
							πF.SetLineno(392)
							// line 393: return self.dict[name.lower()]
							πF.SetLineno(393)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µname, ßlower, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001 = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdict, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__getitem__.ToObject(), πTemp018); πE != nil {
						continue
					}
					// line 392: """Get a specific header, as from a dictionary."""
					πF.SetLineno(392)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp019}, πg.NewStr("Get a specific header, as from a dictionary.").ToObject()); πE != nil {
						continue
					}
					if πTemp020, πE = πg.ResolveClass(πF, πClass, nil, ß__getitem__); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp020, ß__doc__, πTemp019); πE != nil {
						continue
					}
					// line 395: def __setitem__(self, name, value):
					πF.SetLineno(395)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "name", Def: nil}
					πTemp002[2] = πg.Param{Name: "value", Def: nil}
					πTemp019 = πg.NewFunction(πg.NewCode("__setitem__", "/usr/lib/python2.7/rfc822.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µname *πg.Object = πArgs[1]; _ = µname
						var µvalue *πg.Object = πArgs[2]; _ = µvalue
						var µtext *πg.Object = πg.UnboundLocal; _ = µtext
						var µline *πg.Object = πg.UnboundLocal; _ = µline
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
						var πTemp007 bool
						_ = πTemp007
						var πTemp008 bool
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
							// line 396: """Set the value of a header.
							πF.SetLineno(396)
							// line 402: del self[name] # Won't fail if it doesn't exist
							πF.SetLineno(402)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp001 = µname
							if πE = πg.DelItem(πF, µself, πTemp001); πE != nil {
								continue
							}
							// line 403: self.dict[name.lower()] = value
							πF.SetLineno(403)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µvalue); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdict, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µname, ßlower, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003 = πTemp005
							if πE = πg.SetItem(πF, πTemp002, πTemp003, πTemp001); πE != nil {
								continue
							}
							// line 404: text = name + ": " + value
							πF.SetLineno(404)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, µname, πg.NewStr(": ").ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp002, µvalue); πE != nil {
								continue
							}
							µtext = πTemp001
							πTemp006 = πF.MakeArgs(1)
							πTemp006[0] = πg.NewStr("\n").ToObject()
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µtext, ßsplit, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πTemp001, πE = πg.Iter(πF, πTemp003); πE != nil {
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
								µline = πTemp002
							}
							if πE != nil || !πTemp008 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 406: self.headers.append(line + "\n")
							πF.SetLineno(406)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, µline, πg.NewStr("\n").ToObject()); πE != nil {
								continue
							}
							πTemp006[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßheaders, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
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
					if πE = πClass.SetItem(πF, ß__setitem__.ToObject(), πTemp019); πE != nil {
						continue
					}
					// line 396: """Set the value of a header.
					πF.SetLineno(396)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp020}, πg.NewStr("Set the value of a header.\n\n        Note: This is not a perfect inversion of __getitem__, because any\n        changed headers get stuck at the end of the raw-headers list rather\n        than where the altered header was.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp021, πE = πg.ResolveClass(πF, πClass, nil, ß__setitem__); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp021, ß__doc__, πTemp020); πE != nil {
						continue
					}
					// line 408: def __delitem__(self, name):
					πF.SetLineno(408)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "name", Def: nil}
					πTemp020 = πg.NewFunction(πg.NewCode("__delitem__", "/usr/lib/python2.7/rfc822.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µname *πg.Object = πArgs[1]; _ = µname
						var µn *πg.Object = πg.UnboundLocal; _ = µn
						var µlst *πg.Object = πg.UnboundLocal; _ = µlst
						var µhit *πg.Object = πg.UnboundLocal; _ = µhit
						var µi *πg.Object = πg.UnboundLocal; _ = µi
						var µline *πg.Object = πg.UnboundLocal; _ = µline
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
							case 11: goto Label11
							case 12: goto Label12
							case 3: goto Label3
							case 4: goto Label4
							default: panic("unexpected function state")
							}
							// line 409: """Delete all occurrences of a specific header, if it is present."""
							πF.SetLineno(409)
							// line 410: name = name.lower()
							πF.SetLineno(410)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µname, ßlower, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µname = πTemp002
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdict, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Contains(πF, πTemp003, µname); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp004).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
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
							// line 411: if not name in self.dict:
							πF.SetLineno(411)
						Label1:
							// line 412: return
							πF.SetLineno(412)
							πR = πg.None
							continue
							goto Label2
						Label2:
							// line 413: del self.dict[name]
							πF.SetLineno(413)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdict, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp002 = µname
							if πE = πg.DelItem(πF, πTemp001, πTemp002); πE != nil {
								continue
							}
							// line 414: name = name + ':'
							πF.SetLineno(414)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, µname, πg.NewStr(":").ToObject()); πE != nil {
								continue
							}
							µname = πTemp001
							// line 415: n = len(name)
							πF.SetLineno(415)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp005[0] = µname
							if πTemp001, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µn = πTemp002
							// line 416: lst = []
							πF.SetLineno(416)
							πTemp005 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp005...).ToObject()
							µlst = πTemp001
							// line 417: hit = 0
							πF.SetLineno(417)
							µhit = πg.NewInt(0).ToObject()
							πTemp005 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßheaders, nil); πE != nil {
								continue
							}
							πTemp006[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp005[0] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
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
								πTemp007 = !isStop
							} else {
								πTemp007 = true
								µi = πTemp002
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(3)            
							// line 419: line = self.headers[i]
							πF.SetLineno(419)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp002 = µi
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µself, ßheaders, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp008, πTemp002); πE != nil {
								continue
							}
							µline = πTemp003
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.None, µn, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, µline, πTemp003); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp008, ßlower, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, πTemp008, µname); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label6
							}
							if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.NewInt(1).ToObject(), πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, µline, πTemp003); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp008, ßisspace, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp008); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp007).ToObject()
							if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label7
							}
							goto Label8
							// line 420: if line[:n].lower() == name:
							πF.SetLineno(420)
						Label6:
							// line 421: hit = 1
							πF.SetLineno(421)
							µhit = πg.NewInt(1).ToObject()
							goto Label8
							// line 422: elif not line[:1].isspace():
							πF.SetLineno(422)
						Label7:
							// line 423: hit = 0
							πF.SetLineno(423)
							µhit = πg.NewInt(0).ToObject()
							goto Label8
						Label8:
							if πE = πg.CheckLocal(πF, µhit, "hit"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, µhit); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label9
							}
							goto Label10
							// line 424: if hit:
							πF.SetLineno(424)
						Label9:
							// line 425: lst.append(i)
							πF.SetLineno(425)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp005[0] = µi
							if πE = πg.CheckLocal(πF, µlst, "lst"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µlst, ßappend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							goto Label10
						Label10:
							continue
						Label4:
							if πE != nil || πR != nil {
								continue
							}
						Label5:
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µlst, "lst"); πE != nil {
								continue
							}
							πTemp005[0] = µlst
							if πTemp002, πE = πg.ResolveGlobal(πF, ßreversed); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp001, πE = πg.Iter(πF, πTemp003); πE != nil {
								continue
							}
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
							πF.PushCheckpoint(11)            
							// line 427: del self.headers[i]
							πF.SetLineno(427)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßheaders, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp003 = µi
							if πE = πg.DelItem(πF, πTemp002, πTemp003); πE != nil {
								continue
							}
							continue
						Label12:
							if πE != nil || πR != nil {
								continue
							}
						Label13:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß__delitem__.ToObject(), πTemp020); πE != nil {
						continue
					}
					// line 409: """Delete all occurrences of a specific header, if it is present."""
					πF.SetLineno(409)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp021}, πg.NewStr("Delete all occurrences of a specific header, if it is present.").ToObject()); πE != nil {
						continue
					}
					if πTemp022, πE = πg.ResolveClass(πF, πClass, nil, ß__delitem__); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp022, ß__doc__, πTemp021); πE != nil {
						continue
					}
					// line 429: def setdefault(self, name, default=""):
					πF.SetLineno(429)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "name", Def: nil}
					πTemp002[2] = πg.Param{Name: "default", Def: ß.ToObject()}
					πTemp021 = πg.NewFunction(πg.NewCode("setdefault", "/usr/lib/python2.7/rfc822.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µname *πg.Object = πArgs[1]; _ = µname
						var µdefault *πg.Object = πArgs[2]; _ = µdefault
						var µlowername *πg.Object = πg.UnboundLocal; _ = µlowername
						var µtext *πg.Object = πg.UnboundLocal; _ = µtext
						var µline *πg.Object = πg.UnboundLocal; _ = µline
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
							// line 430: lowername = name.lower()
							πF.SetLineno(430)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µname, ßlower, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µlowername = πTemp002
							if πE = πg.CheckLocal(πF, µlowername, "lowername"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdict, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Contains(πF, πTemp002, µlowername); πE != nil {
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
							// line 431: if lowername in self.dict:
							πF.SetLineno(431)
						Label1:
							// line 432: return self.dict[lowername]
							πF.SetLineno(432)
							if πE = πg.CheckLocal(πF, µlowername, "lowername"); πE != nil {
								continue
							}
							πTemp001 = µlowername
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßdict, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp004, πTemp001); πE != nil {
								continue
							}
							πR = πTemp002
							continue
							goto Label3
						Label2:
							// line 434: text = name + ": " + default
							πF.SetLineno(434)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, µname, πg.NewStr(": ").ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µdefault, "default"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp002, µdefault); πE != nil {
								continue
							}
							µtext = πTemp001
							πTemp005 = πF.MakeArgs(1)
							πTemp005[0] = πg.NewStr("\n").ToObject()
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µtext, ßsplit, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp001, πE = πg.Iter(πF, πTemp004); πE != nil {
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
								πTemp006 = !isStop
							} else {
								πTemp006 = true
								µline = πTemp002
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(4)            
							// line 436: self.headers.append(line + "\n")
							πF.SetLineno(436)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, µline, πg.NewStr("\n").ToObject()); πE != nil {
								continue
							}
							πTemp005[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßheaders, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							continue
						Label5:
							if πE != nil || πR != nil {
								continue
							}
						Label6:
							// line 437: self.dict[lowername] = default
							πF.SetLineno(437)
							if πE = πg.CheckLocal(πF, µdefault, "default"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µdefault); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdict, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µlowername, "lowername"); πE != nil {
								continue
							}
							πTemp004 = µlowername
							if πE = πg.SetItem(πF, πTemp002, πTemp004, πTemp001); πE != nil {
								continue
							}
							// line 438: return default
							πF.SetLineno(438)
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
					if πE = πClass.SetItem(πF, ßsetdefault.ToObject(), πTemp021); πE != nil {
						continue
					}
					// line 440: def has_key(self, name):
					πF.SetLineno(440)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "name", Def: nil}
					πTemp022 = πg.NewFunction(πg.NewCode("has_key", "/usr/lib/python2.7/rfc822.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µname *πg.Object = πArgs[1]; _ = µname
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
							// line 441: """Determine whether a message contains the named header."""
							πF.SetLineno(441)
							// line 442: return name.lower() in self.dict
							πF.SetLineno(442)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µname, ßlower, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdict, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Contains(πF, πTemp002, πTemp003); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp004).ToObject()
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
					if πE = πClass.SetItem(πF, ßhas_key.ToObject(), πTemp022); πE != nil {
						continue
					}
					// line 441: """Determine whether a message contains the named header."""
					πF.SetLineno(441)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp023}, πg.NewStr("Determine whether a message contains the named header.").ToObject()); πE != nil {
						continue
					}
					if πTemp024, πE = πg.ResolveClass(πF, πClass, nil, ßhas_key); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp024, ß__doc__, πTemp023); πE != nil {
						continue
					}
					// line 444: def __contains__(self, name):
					πF.SetLineno(444)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "name", Def: nil}
					πTemp023 = πg.NewFunction(πg.NewCode("__contains__", "/usr/lib/python2.7/rfc822.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µname *πg.Object = πArgs[1]; _ = µname
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
							// line 445: """Determine whether a message contains the named header."""
							πF.SetLineno(445)
							// line 446: return name.lower() in self.dict
							πF.SetLineno(446)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µname, ßlower, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdict, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Contains(πF, πTemp002, πTemp003); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp004).ToObject()
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
					if πE = πClass.SetItem(πF, ß__contains__.ToObject(), πTemp023); πE != nil {
						continue
					}
					// line 445: """Determine whether a message contains the named header."""
					πF.SetLineno(445)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp024}, πg.NewStr("Determine whether a message contains the named header.").ToObject()); πE != nil {
						continue
					}
					if πTemp025, πE = πg.ResolveClass(πF, πClass, nil, ß__contains__); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp025, ß__doc__, πTemp024); πE != nil {
						continue
					}
					// line 448: def __iter__(self):
					πF.SetLineno(448)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp024 = πg.NewFunction(πg.NewCode("__iter__", "/usr/lib/python2.7/rfc822.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 449: return iter(self.dict)
							πF.SetLineno(449)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdict, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
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
					if πE = πClass.SetItem(πF, ß__iter__.ToObject(), πTemp024); πE != nil {
						continue
					}
					// line 451: def keys(self):
					πF.SetLineno(451)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp025 = πg.NewFunction(πg.NewCode("keys", "/usr/lib/python2.7/rfc822.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 452: """Get all of a message's header field names."""
							πF.SetLineno(452)
							// line 453: return self.dict.keys()
							πF.SetLineno(453)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdict, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßkeys, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßkeys.ToObject(), πTemp025); πE != nil {
						continue
					}
					// line 452: """Get all of a message's header field names."""
					πF.SetLineno(452)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp026}, πg.NewStr("Get all of a message's header field names.").ToObject()); πE != nil {
						continue
					}
					if πTemp027, πE = πg.ResolveClass(πF, πClass, nil, ßkeys); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp027, ß__doc__, πTemp026); πE != nil {
						continue
					}
					// line 455: def values(self):
					πF.SetLineno(455)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp026 = πg.NewFunction(πg.NewCode("values", "/usr/lib/python2.7/rfc822.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 456: """Get all of a message's header field values."""
							πF.SetLineno(456)
							// line 457: return self.dict.values()
							πF.SetLineno(457)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdict, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßvalues, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßvalues.ToObject(), πTemp026); πE != nil {
						continue
					}
					// line 456: """Get all of a message's header field values."""
					πF.SetLineno(456)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp027}, πg.NewStr("Get all of a message's header field values.").ToObject()); πE != nil {
						continue
					}
					if πTemp028, πE = πg.ResolveClass(πF, πClass, nil, ßvalues); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp028, ß__doc__, πTemp027); πE != nil {
						continue
					}
					// line 459: def items(self):
					πF.SetLineno(459)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp027 = πg.NewFunction(πg.NewCode("items", "/usr/lib/python2.7/rfc822.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 460: """Get all of a message's headers.
							πF.SetLineno(460)
							// line 464: return self.dict.items()
							πF.SetLineno(464)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdict, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßitems, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßitems.ToObject(), πTemp027); πE != nil {
						continue
					}
					// line 460: """Get all of a message's headers.
					πF.SetLineno(460)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp028}, πg.NewStr("Get all of a message's headers.\n\n        Returns a list of name, value tuples.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp029, πE = πg.ResolveClass(πF, πClass, nil, ßitems); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp029, ß__doc__, πTemp028); πE != nil {
						continue
					}
					// line 466: def __str__(self):
					πF.SetLineno(466)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp028 = πg.NewFunction(πg.NewCode("__str__", "/usr/lib/python2.7/rfc822.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 467: return ''.join(self.headers)
							πF.SetLineno(467)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßheaders, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
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
					if πE = πClass.SetItem(πF, ß__str__.ToObject(), πTemp028); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp005.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp006, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("Message").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp005.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßMessage.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 477: def unquote(s):
			πF.SetLineno(477)
			πTemp007 = make([]πg.Param, 1)
			πTemp007[0] = πg.Param{Name: "s", Def: nil}
			πTemp001 = πg.NewFunction(πg.NewCode("unquote", "/usr/lib/python2.7/rfc822.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µs *πg.Object = πArgs[0]; _ = µs
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
				var πTemp006 []*πg.Object
				_ = πTemp006
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 478: """Remove quotes from a string."""
					πF.SetLineno(478)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp002[0] = µs
					if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.GT(πF, πTemp004, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label1
					}
					goto Label2
					// line 479: if len(s) > 1:
					πF.SetLineno(479)
				Label1:
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewStr("\"").ToObject()
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µs, ßstartswith, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001 = πTemp004
					if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp005 {
						goto Label3
					}
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewStr("\"").ToObject()
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µs, ßendswith, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001 = πTemp004
				Label3:
					if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label4
					}
					goto Label5
					// line 480: if s.startswith('"') and s.endswith('"'):
					πF.SetLineno(480)
				Label4:
					// line 481: return s[1:-1].replace('\\\\', '\\').replace('\\"', '"')
					πF.SetLineno(481)
					πTemp002 = πF.MakeArgs(2)
					πTemp002[0] = πg.NewStr("\\\"").ToObject()
					πTemp002[1] = πg.NewStr("\"").ToObject()
					πTemp006 = πF.MakeArgs(2)
					πTemp006[0] = πg.NewStr("\\\\").ToObject()
					πTemp006[1] = πg.NewStr("\\").ToObject()
					if πTemp003, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πTemp003, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µs, πTemp001); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, πTemp003, ßreplace, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πTemp001, πE = πg.GetAttr(πF, πTemp003, ßreplace, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πR = πTemp003
					continue
					goto Label5
				Label5:
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewStr("<").ToObject()
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µs, ßstartswith, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001 = πTemp004
					if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp005 {
						goto Label6
					}
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewStr(">").ToObject()
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µs, ßendswith, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001 = πTemp004
				Label6:
					if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label7
					}
					goto Label8
					// line 482: if s.startswith('<') and s.endswith('>'):
					πF.SetLineno(482)
				Label7:
					// line 483: return s[1:-1]
					πF.SetLineno(483)
					if πTemp003, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πTemp003, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µs, πTemp001); πE != nil {
						continue
					}
					πR = πTemp003
					continue
					goto Label8
				Label8:
					goto Label2
				Label2:
					// line 484: return s
					πF.SetLineno(484)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πR = µs
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßunquote.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 478: """Remove quotes from a string."""
			πF.SetLineno(478)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewStr("Remove quotes from a string.").ToObject()); πE != nil {
				continue
			}
			if πTemp006, πE = πg.ResolveGlobal(πF, ßunquote); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp006, ß__doc__, πTemp003); πE != nil {
				continue
			}
			// line 487: def quote(s):
			πF.SetLineno(487)
			πTemp007 = make([]πg.Param, 1)
			πTemp007[0] = πg.Param{Name: "s", Def: nil}
			πTemp003 = πg.NewFunction(πg.NewCode("quote", "/usr/lib/python2.7/rfc822.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µs *πg.Object = πArgs[0]; _ = µs
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
					// line 488: """Add quotes around a string."""
					πF.SetLineno(488)
					// line 489: return s.replace('\\', '\\\\').replace('"', '\\"')
					πF.SetLineno(489)
					πTemp001 = πF.MakeArgs(2)
					πTemp001[0] = πg.NewStr("\"").ToObject()
					πTemp001[1] = πg.NewStr("\\\"").ToObject()
					πTemp002 = πF.MakeArgs(2)
					πTemp002[0] = πg.NewStr("\\").ToObject()
					πTemp002[1] = πg.NewStr("\\\\").ToObject()
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µs, ßreplace, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßreplace, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßquote.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 488: """Add quotes around a string."""
			πF.SetLineno(488)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, πg.NewStr("Add quotes around a string.").ToObject()); πE != nil {
				continue
			}
			if πTemp008, πE = πg.ResolveGlobal(πF, ßquote); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp008, ß__doc__, πTemp006); πE != nil {
				continue
			}
			// line 492: def parseaddr(address):
			πF.SetLineno(492)
			πTemp007 = make([]πg.Param, 1)
			πTemp007[0] = πg.Param{Name: "address", Def: nil}
			πTemp006 = πg.NewFunction(πg.NewCode("parseaddr", "/usr/lib/python2.7/rfc822.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µaddress *πg.Object = πArgs[0]; _ = µaddress
				var µa *πg.Object = πg.UnboundLocal; _ = µa
				var µlst *πg.Object = πg.UnboundLocal; _ = µlst
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
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 493: """Parse an address into a (realname, mailaddr) tuple."""
					πF.SetLineno(493)
					// line 494: a = AddressList(address)
					πF.SetLineno(494)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µaddress, "address"); πE != nil {
						continue
					}
					πTemp001[0] = µaddress
					if πTemp002, πE = πg.ResolveGlobal(πF, ßAddressList); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µa = πTemp003
					// line 495: lst = a.addresslist
					πF.SetLineno(495)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µa, ßaddresslist, nil); πE != nil {
						continue
					}
					µlst = πTemp002
					if πE = πg.CheckLocal(πF, µlst, "lst"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, µlst); πE != nil {
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
					// line 496: if not lst:
					πF.SetLineno(496)
				Label1:
					// line 497: return (None, None)
					πF.SetLineno(497)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(πTemp003, πTemp005).ToObject()
					πR = πTemp002
					continue
					goto Label2
				Label2:
					// line 498: return lst[0]
					πF.SetLineno(498)
					πTemp002 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µlst, "lst"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µlst, πTemp002); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßparseaddr.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 493: """Parse an address into a (realname, mailaddr) tuple."""
			πF.SetLineno(493)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp008}, πg.NewStr("Parse an address into a (realname, mailaddr) tuple.").ToObject()); πE != nil {
				continue
			}
			if πTemp009, πE = πg.ResolveGlobal(πF, ßparseaddr); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp009, ß__doc__, πTemp008); πE != nil {
				continue
			}
			// line 501: class AddrlistClass(object):
			πF.SetLineno(501)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp010, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp010
			πTemp005 = πg.NewDict()
			if πTemp008, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp005.SetItem(πF, ß__module__.ToObject(), πTemp008); πE != nil {
				continue
			}
			_, πE = πg.NewCode("AddrlistClass", "/usr/lib/python2.7/rfc822.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp005
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 502: """Address parser class by Ben Escoto.
					πF.SetLineno(502)
					// line 502: """Address parser class by Ben Escoto.
					πF.SetLineno(502)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Address parser class by Ben Escoto.\n\n    To understand what this class does, it helps to have a copy of\n    RFC 2822 in front of you.\n\n    http://www.faqs.org/rfcs/rfc2822.html\n\n    Note: this class interface is deprecated and may be removed in the future.\n    Use rfc822.AddressList instead.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 513: def __init__(self, field):
					πF.SetLineno(513)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "field", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/rfc822.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µfield *πg.Object = πArgs[1]; _ = µfield
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 514: """Initialize a new instance.
							πF.SetLineno(514)
							// line 519: self.specials = '()<>@,:;.\"[]'
							πF.SetLineno(519)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewStr("()<>@,:;.\"[]").ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßspecials, πTemp001); πE != nil {
								continue
							}
							// line 520: self.pos = 0
							πF.SetLineno(520)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßpos, πTemp001); πE != nil {
								continue
							}
							// line 521: self.LWS = ' \t'
							πF.SetLineno(521)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewStr(" \t").ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßLWS, πTemp001); πE != nil {
								continue
							}
							// line 522: self.CR = '\r\n'
							πF.SetLineno(522)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewStr("\r\n").ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßCR, πTemp001); πE != nil {
								continue
							}
							// line 523: self.atomends = self.specials + self.LWS + self.CR
							πF.SetLineno(523)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßspecials, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßLWS, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp003, πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßCR, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp002, πTemp003); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßatomends, πTemp002); πE != nil {
								continue
							}
							// line 527: self.phraseends = self.atomends.replace('.', '')
							πF.SetLineno(527)
							πTemp005 = πF.MakeArgs(2)
							πTemp005[0] = πg.NewStr(".").ToObject()
							πTemp005[1] = ß.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßatomends, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßreplace, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßphraseends, πTemp002); πE != nil {
								continue
							}
							// line 528: self.field = field
							πF.SetLineno(528)
							if πE = πg.CheckLocal(πF, µfield, "field"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µfield); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßfield, πTemp001); πE != nil {
								continue
							}
							// line 529: self.commentlist = []
							πF.SetLineno(529)
							πTemp005 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp005...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßcommentlist, πTemp002); πE != nil {
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
					// line 514: """Initialize a new instance.
					πF.SetLineno(514)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewStr("Initialize a new instance.\n\n        `field' is an unparsed address header field, containing one or more\n        addresses.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ß__init__); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp004, ß__doc__, πTemp003); πE != nil {
						continue
					}
					// line 531: def gotonext(self):
					πF.SetLineno(531)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("gotonext", "/usr/lib/python2.7/rfc822.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πTemp001 bool
						_ = πTemp001
						var πTemp002 bool
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							// line 532: """Parse up to the start of the next address."""
							πF.SetLineno(532)
							// line 533: while self.pos < len(self.field):
							πF.SetLineno(533)
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
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßpos, nil); πE != nil {
								continue
							}
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßfield, nil); πE != nil {
								continue
							}
							πTemp005[0] = πTemp006
							if πTemp006, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp003, πE = πg.LT(πF, πTemp004, πTemp007); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πE != nil || !πTemp002 {
								continue
							}
							πF.PushCheckpoint(1)            
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßpos, nil); πE != nil {
								continue
							}
							πTemp004 = πTemp006
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßfield, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, πTemp007, πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßLWS, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Add(πF, πTemp007, πg.NewStr("\n\r").ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Contains(πF, πTemp004, πTemp006); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(πTemp002).ToObject()
							if πTemp002, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label4
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßpos, nil); πE != nil {
								continue
							}
							πTemp004 = πTemp006
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßfield, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, πTemp007, πTemp004); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Eq(πF, πTemp006, πg.NewStr("(").ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label5
							}
							goto Label6
							// line 534: if self.field[self.pos] in self.LWS + '\n\r':
							πF.SetLineno(534)
						Label4:
							// line 535: self.pos = self.pos + 1
							πF.SetLineno(535)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßpos, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, πTemp004, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßpos, πTemp004); πE != nil {
								continue
							}
							goto Label7
							// line 536: elif self.field[self.pos] == '(':
							πF.SetLineno(536)
						Label5:
							// line 537: self.commentlist.append(self.getcomment())
							πF.SetLineno(537)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßgetcomment, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp005[0] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßcommentlist, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßappend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							goto Label7
						Label6:
							// line 538: else: break
							πF.SetLineno(538)
							πTemp001 = true
							continue
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
					if πE = πClass.SetItem(πF, ßgotonext.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 532: """Parse up to the start of the next address."""
					πF.SetLineno(532)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πg.NewStr("Parse up to the start of the next address.").ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßgotonext); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp005, ß__doc__, πTemp004); πE != nil {
						continue
					}
					// line 540: def getaddrlist(self):
					πF.SetLineno(540)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("getaddrlist", "/usr/lib/python2.7/rfc822.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µresult *πg.Object = πg.UnboundLocal; _ = µresult
						var µad *πg.Object = πg.UnboundLocal; _ = µad
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
							// line 541: """Parse all addresses.
							πF.SetLineno(541)
							// line 545: result = []
							πF.SetLineno(545)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µresult = πTemp002
							// line 546: ad = self.getaddress()
							πF.SetLineno(546)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßgetaddress, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µad = πTemp003
							// line 547: while ad:
							πF.SetLineno(547)
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
							if πE = πg.CheckLocal(πF, µad, "ad"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, µad); πE != nil {
								continue
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 548: result += ad
							πF.SetLineno(548)
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µad, "ad"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IAdd(πF, µresult, µad); πE != nil {
								continue
							}
							µresult = πTemp002
							// line 549: ad = self.getaddress()
							πF.SetLineno(549)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßgetaddress, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µad = πTemp003
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 550: return result
							πF.SetLineno(550)
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
					if πE = πClass.SetItem(πF, ßgetaddrlist.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 541: """Parse all addresses.
					πF.SetLineno(541)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πg.NewStr("Parse all addresses.\n\n        Returns a list containing all of the addresses.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßgetaddrlist); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp006, ß__doc__, πTemp005); πE != nil {
						continue
					}
					// line 552: def getaddress(self):
					πF.SetLineno(552)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("getaddress", "/usr/lib/python2.7/rfc822.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µoldpos *πg.Object = πg.UnboundLocal; _ = µoldpos
						var µoldcl *πg.Object = πg.UnboundLocal; _ = µoldcl
						var µplist *πg.Object = πg.UnboundLocal; _ = µplist
						var µreturnlist *πg.Object = πg.UnboundLocal; _ = µreturnlist
						var µaddrspec *πg.Object = πg.UnboundLocal; _ = µaddrspec
						var µfieldlen *πg.Object = πg.UnboundLocal; _ = µfieldlen
						var µrouteaddr *πg.Object = πg.UnboundLocal; _ = µrouteaddr
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
						var πTemp007 []*πg.Object
						_ = πTemp007
						var πTemp008 bool
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
							case 10: goto Label10
							case 11: goto Label11
							default: panic("unexpected function state")
							}
							// line 553: """Parse the next address."""
							πF.SetLineno(553)
							// line 554: self.commentlist = []
							πF.SetLineno(554)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßcommentlist, πTemp003); πE != nil {
								continue
							}
							// line 555: self.gotonext()
							πF.SetLineno(555)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßgotonext, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 557: oldpos = self.pos
							πF.SetLineno(557)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßpos, nil); πE != nil {
								continue
							}
							µoldpos = πTemp002
							// line 558: oldcl = self.commentlist
							πF.SetLineno(558)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcommentlist, nil); πE != nil {
								continue
							}
							µoldcl = πTemp002
							// line 559: plist = self.getphraselist()
							πF.SetLineno(559)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßgetphraselist, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µplist = πTemp003
							// line 561: self.gotonext()
							πF.SetLineno(561)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßgotonext, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 562: returnlist = []
							πF.SetLineno(562)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µreturnlist = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßpos, nil); πE != nil {
								continue
							}
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßfield, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.GE(πF, πTemp003, πTemp005); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßpos, nil); πE != nil {
								continue
							}
							πTemp003 = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßfield, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, πTemp005, πTemp003); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Contains(πF, πg.NewStr(".@").ToObject(), πTemp004); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp006).ToObject()
							if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label2
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßpos, nil); πE != nil {
								continue
							}
							πTemp003 = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßfield, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, πTemp005, πTemp003); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, πTemp004, πg.NewStr(":").ToObject()); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label3
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßpos, nil); πE != nil {
								continue
							}
							πTemp003 = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßfield, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, πTemp005, πTemp003); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, πTemp004, πg.NewStr("<").ToObject()); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label4
							}
							if πE = πg.CheckLocal(πF, µplist, "plist"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, µplist); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label5
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßpos, nil); πE != nil {
								continue
							}
							πTemp003 = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßfield, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, πTemp005, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßspecials, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Contains(πF, πTemp003, πTemp004); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp006).ToObject()
							if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label6
							}
							goto Label7
							// line 564: if self.pos >= len(self.field):
							πF.SetLineno(564)
						Label1:
							if πE = πg.CheckLocal(πF, µplist, "plist"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, µplist); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label8
							}
							goto Label9
							// line 566: if plist:
							πF.SetLineno(566)
						Label8:
							// line 567: returnlist = [(' '.join(self.commentlist), plist[0])]
							πF.SetLineno(567)
							πTemp001 = make([]*πg.Object, 1)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßcommentlist, nil); πE != nil {
								continue
							}
							πTemp007[0] = πTemp003
							if πTemp003, πE = πg.GetAttr(πF, πg.NewStr(" ").ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp003 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µplist, "plist"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µplist, πTemp003); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(πTemp004, πTemp005).ToObject()
							πTemp001[0] = πTemp002
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µreturnlist = πTemp002
							goto Label9
						Label9:
							goto Label7
							// line 569: elif self.field[self.pos] in '.@':
							πF.SetLineno(569)
						Label2:
							// line 572: self.pos = oldpos
							πF.SetLineno(572)
							if πE = πg.CheckLocal(πF, µoldpos, "oldpos"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µoldpos); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßpos, πTemp002); πE != nil {
								continue
							}
							// line 573: self.commentlist = oldcl
							πF.SetLineno(573)
							if πE = πg.CheckLocal(πF, µoldcl, "oldcl"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µoldcl); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßcommentlist, πTemp002); πE != nil {
								continue
							}
							// line 574: addrspec = self.getaddrspec()
							πF.SetLineno(574)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßgetaddrspec, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µaddrspec = πTemp003
							// line 575: returnlist = [(' '.join(self.commentlist), addrspec)]
							πF.SetLineno(575)
							πTemp001 = make([]*πg.Object, 1)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßcommentlist, nil); πE != nil {
								continue
							}
							πTemp007[0] = πTemp003
							if πTemp003, πE = πg.GetAttr(πF, πg.NewStr(" ").ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							if πE = πg.CheckLocal(πF, µaddrspec, "addrspec"); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(πTemp004, µaddrspec).ToObject()
							πTemp001[0] = πTemp002
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µreturnlist = πTemp002
							goto Label7
							// line 577: elif self.field[self.pos] == ':':
							πF.SetLineno(577)
						Label3:
							// line 579: returnlist = []
							πF.SetLineno(579)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µreturnlist = πTemp002
							// line 581: fieldlen = len(self.field)
							πF.SetLineno(581)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßfield, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µfieldlen = πTemp003
							// line 582: self.pos += 1
							πF.SetLineno(582)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßpos, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IAdd(πF, πTemp002, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßpos, πTemp003); πE != nil {
								continue
							}
							// line 583: while self.pos < len(self.field):
							πF.SetLineno(583)
							πF.PushCheckpoint(11)
							πTemp006 = false
						Label10:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp006 {
								πF.PopCheckpoint()
								goto Label12
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßpos, nil); πE != nil {
								continue
							}
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßfield, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.LT(πF, πTemp003, πTemp005); πE != nil {
								continue
							}
							if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πE != nil || !πTemp008 {
								continue
							}
							πF.PushCheckpoint(10)            
							// line 584: self.gotonext()
							πF.SetLineno(584)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßgotonext, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßpos, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µfieldlen, "fieldlen"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.LT(πF, πTemp004, µfieldlen); πE != nil {
								continue
							}
							πTemp002 = πTemp003
							if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp008 {
								goto Label13
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßpos, nil); πE != nil {
								continue
							}
							πTemp004 = πTemp005
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, µself, ßfield, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, πTemp009, πTemp004); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Eq(πF, πTemp005, πg.NewStr(";").ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp003
						Label13:
							if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label14
							}
							goto Label15
							// line 585: if self.pos < fieldlen and self.field[self.pos] == ';':
							πF.SetLineno(585)
						Label14:
							// line 586: self.pos += 1
							πF.SetLineno(586)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßpos, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IAdd(πF, πTemp002, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßpos, πTemp003); πE != nil {
								continue
							}
							// line 587: break
							πF.SetLineno(587)
							πTemp006 = true
							continue
							goto Label15
						Label15:
							// line 588: returnlist = returnlist + self.getaddress()
							πF.SetLineno(588)
							if πE = πg.CheckLocal(πF, µreturnlist, "returnlist"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßgetaddress, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, µreturnlist, πTemp004); πE != nil {
								continue
							}
							µreturnlist = πTemp002
							continue
						Label11:
							if πE != nil || πR != nil {
								continue
							}
						Label12:
							goto Label7
							// line 590: elif self.field[self.pos] == '<':
							πF.SetLineno(590)
						Label4:
							// line 592: routeaddr = self.getrouteaddr()
							πF.SetLineno(592)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßgetrouteaddr, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µrouteaddr = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcommentlist, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label16
							}
							goto Label17
							// line 594: if self.commentlist:
							πF.SetLineno(594)
						Label16:
							// line 595: returnlist = [(' '.join(plist) + ' (' + \
							πF.SetLineno(595)
							πTemp001 = make([]*πg.Object, 1)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µplist, "plist"); πE != nil {
								continue
							}
							πTemp007[0] = µplist
							if πTemp009, πE = πg.GetAttr(πF, πg.NewStr(" ").ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp010, πE = πTemp009.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							if πTemp005, πE = πg.Add(πF, πTemp010, πg.NewStr(" (").ToObject()); πE != nil {
								continue
							}
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, µself, ßcommentlist, nil); πE != nil {
								continue
							}
							πTemp007[0] = πTemp009
							if πTemp009, πE = πg.GetAttr(πF, πg.NewStr(" ").ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp010, πE = πTemp009.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							if πTemp004, πE = πg.Add(πF, πTemp005, πTemp010); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, πTemp004, πg.NewStr(")").ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µrouteaddr, "routeaddr"); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(πTemp003, µrouteaddr).ToObject()
							πTemp001[0] = πTemp002
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µreturnlist = πTemp002
							goto Label18
						Label17:
							// line 597: else: returnlist = [(' '.join(plist), routeaddr)]
							πF.SetLineno(597)
							πTemp001 = make([]*πg.Object, 1)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µplist, "plist"); πE != nil {
								continue
							}
							πTemp007[0] = µplist
							if πTemp003, πE = πg.GetAttr(πF, πg.NewStr(" ").ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							if πE = πg.CheckLocal(πF, µrouteaddr, "routeaddr"); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(πTemp004, µrouteaddr).ToObject()
							πTemp001[0] = πTemp002
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µreturnlist = πTemp002
							goto Label18
						Label18:
							goto Label7
							// line 600: if plist:
							πF.SetLineno(600)
						Label5:
							// line 601: returnlist = [(' '.join(self.commentlist), plist[0])]
							πF.SetLineno(601)
							πTemp001 = make([]*πg.Object, 1)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßcommentlist, nil); πE != nil {
								continue
							}
							πTemp007[0] = πTemp003
							if πTemp003, πE = πg.GetAttr(πF, πg.NewStr(" ").ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp003 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µplist, "plist"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µplist, πTemp003); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(πTemp004, πTemp005).ToObject()
							πTemp001[0] = πTemp002
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µreturnlist = πTemp002
							goto Label7
							// line 602: elif self.field[self.pos] in self.specials:
							πF.SetLineno(602)
						Label6:
							// line 603: self.pos += 1
							πF.SetLineno(603)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßpos, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IAdd(πF, πTemp002, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßpos, πTemp003); πE != nil {
								continue
							}
							goto Label7
						Label7:
							// line 605: self.gotonext()
							πF.SetLineno(605)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßgotonext, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßpos, nil); πE != nil {
								continue
							}
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßfield, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp005
							if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp003, πE = πg.LT(πF, πTemp004, πTemp009); πE != nil {
								continue
							}
							πTemp002 = πTemp003
							if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp006 {
								goto Label19
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßpos, nil); πE != nil {
								continue
							}
							πTemp004 = πTemp005
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, µself, ßfield, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, πTemp009, πTemp004); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Eq(πF, πTemp005, πg.NewStr(",").ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp003
						Label19:
							if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label20
							}
							goto Label21
							// line 606: if self.pos < len(self.field) and self.field[self.pos] == ',':
							πF.SetLineno(606)
						Label20:
							// line 607: self.pos += 1
							πF.SetLineno(607)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßpos, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IAdd(πF, πTemp002, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßpos, πTemp003); πE != nil {
								continue
							}
							goto Label21
						Label21:
							// line 608: return returnlist
							πF.SetLineno(608)
							if πE = πg.CheckLocal(πF, µreturnlist, "returnlist"); πE != nil {
								continue
							}
							πR = µreturnlist
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßgetaddress.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 553: """Parse the next address."""
					πF.SetLineno(553)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, πg.NewStr("Parse the next address.").ToObject()); πE != nil {
						continue
					}
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßgetaddress); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp007, ß__doc__, πTemp006); πE != nil {
						continue
					}
					// line 610: def getrouteaddr(self):
					πF.SetLineno(610)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("getrouteaddr", "/usr/lib/python2.7/rfc822.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µexpectroute *πg.Object = πg.UnboundLocal; _ = µexpectroute
						var µadlist *πg.Object = πg.UnboundLocal; _ = µadlist
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 3: goto Label3
							case 4: goto Label4
							default: panic("unexpected function state")
							}
							// line 611: """Parse a route address (Return-path value).
							πF.SetLineno(611)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßpos, nil); πE != nil {
								continue
							}
							πTemp002 = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßfield, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp004, πTemp002); πE != nil {
								continue
							}
							if πTemp001, πE = πg.NE(πF, πTemp003, πg.NewStr("<").ToObject()); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							goto Label2
							// line 615: if self.field[self.pos] != '<':
							πF.SetLineno(615)
						Label1:
							// line 616: return
							πF.SetLineno(616)
							πR = πg.None
							continue
							goto Label2
						Label2:
							// line 618: expectroute = 0
							πF.SetLineno(618)
							µexpectroute = πg.NewInt(0).ToObject()
							// line 619: self.pos += 1
							πF.SetLineno(619)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßpos, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IAdd(πF, πTemp001, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßpos, πTemp002); πE != nil {
								continue
							}
							// line 620: self.gotonext()
							πF.SetLineno(620)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßgotonext, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 621: adlist = ""
							πF.SetLineno(621)
							µadlist = ß.ToObject()
							// line 622: while self.pos < len(self.field):
							πF.SetLineno(622)
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
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßpos, nil); πE != nil {
								continue
							}
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßfield, nil); πE != nil {
								continue
							}
							πTemp007[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							if πTemp001, πE = πg.LT(πF, πTemp002, πTemp004); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(3)            
							if πE = πg.CheckLocal(πF, µexpectroute, "expectroute"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, µexpectroute); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label6
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßpos, nil); πE != nil {
								continue
							}
							πTemp002 = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßfield, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp004, πTemp002); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, πTemp003, πg.NewStr(">").ToObject()); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label7
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßpos, nil); πE != nil {
								continue
							}
							πTemp002 = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßfield, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp004, πTemp002); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, πTemp003, πg.NewStr("@").ToObject()); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label8
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßpos, nil); πE != nil {
								continue
							}
							πTemp002 = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßfield, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp004, πTemp002); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, πTemp003, πg.NewStr(":").ToObject()); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label9
							}
							goto Label10
							// line 623: if expectroute:
							πF.SetLineno(623)
						Label6:
							// line 624: self.getdomain()
							πF.SetLineno(624)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßgetdomain, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 625: expectroute = 0
							πF.SetLineno(625)
							µexpectroute = πg.NewInt(0).ToObject()
							goto Label11
							// line 626: elif self.field[self.pos] == '>':
							πF.SetLineno(626)
						Label7:
							// line 627: self.pos += 1
							πF.SetLineno(627)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßpos, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IAdd(πF, πTemp001, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßpos, πTemp002); πE != nil {
								continue
							}
							// line 628: break
							πF.SetLineno(628)
							πTemp005 = true
							continue
							goto Label11
							// line 629: elif self.field[self.pos] == '@':
							πF.SetLineno(629)
						Label8:
							// line 630: self.pos += 1
							πF.SetLineno(630)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßpos, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IAdd(πF, πTemp001, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßpos, πTemp002); πE != nil {
								continue
							}
							// line 631: expectroute = 1
							πF.SetLineno(631)
							µexpectroute = πg.NewInt(1).ToObject()
							goto Label11
							// line 632: elif self.field[self.pos] == ':':
							πF.SetLineno(632)
						Label9:
							// line 633: self.pos += 1
							πF.SetLineno(633)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßpos, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IAdd(πF, πTemp001, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßpos, πTemp002); πE != nil {
								continue
							}
							goto Label11
						Label10:
							// line 635: adlist = self.getaddrspec()
							πF.SetLineno(635)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßgetaddrspec, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µadlist = πTemp002
							// line 636: self.pos += 1
							πF.SetLineno(636)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßpos, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IAdd(πF, πTemp001, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßpos, πTemp002); πE != nil {
								continue
							}
							// line 637: break
							πF.SetLineno(637)
							πTemp005 = true
							continue
							goto Label11
						Label11:
							// line 638: self.gotonext()
							πF.SetLineno(638)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßgotonext, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							continue
						Label4:
							if πE != nil || πR != nil {
								continue
							}
						Label5:
							// line 640: return adlist
							πF.SetLineno(640)
							if πE = πg.CheckLocal(πF, µadlist, "adlist"); πE != nil {
								continue
							}
							πR = µadlist
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßgetrouteaddr.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 611: """Parse a route address (Return-path value).
					πF.SetLineno(611)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp007}, πg.NewStr("Parse a route address (Return-path value).\n\n        This method just skips all the route stuff and returns the addrspec.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ßgetrouteaddr); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp008, ß__doc__, πTemp007); πE != nil {
						continue
					}
					// line 642: def getaddrspec(self):
					πF.SetLineno(642)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("getaddrspec", "/usr/lib/python2.7/rfc822.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µaslist *πg.Object = πg.UnboundLocal; _ = µaslist
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							// line 643: """Parse an RFC 2822 addr-spec."""
							πF.SetLineno(643)
							// line 644: aslist = []
							πF.SetLineno(644)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µaslist = πTemp002
							// line 646: self.gotonext()
							πF.SetLineno(646)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßgotonext, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 647: while self.pos < len(self.field):
							πF.SetLineno(647)
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
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßpos, nil); πE != nil {
								continue
							}
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßfield, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp006
							if πTemp006, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.LT(πF, πTemp003, πTemp007); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(1)            
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßpos, nil); πE != nil {
								continue
							}
							πTemp003 = πTemp006
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßfield, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, πTemp007, πTemp003); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, πTemp006, πg.NewStr(".").ToObject()); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label4
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßpos, nil); πE != nil {
								continue
							}
							πTemp003 = πTemp006
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßfield, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, πTemp007, πTemp003); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, πTemp006, πg.NewStr("\"").ToObject()); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label5
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßpos, nil); πE != nil {
								continue
							}
							πTemp003 = πTemp006
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßfield, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, πTemp007, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßatomends, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Contains(πF, πTemp003, πTemp006); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp005).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label6
							}
							goto Label7
							// line 648: if self.field[self.pos] == '.':
							πF.SetLineno(648)
						Label4:
							// line 649: aslist.append('.')
							πF.SetLineno(649)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr(".").ToObject()
							if πE = πg.CheckLocal(πF, µaslist, "aslist"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µaslist, ßappend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 650: self.pos += 1
							πF.SetLineno(650)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßpos, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IAdd(πF, πTemp002, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßpos, πTemp003); πE != nil {
								continue
							}
							goto Label8
							// line 651: elif self.field[self.pos] == '"':
							πF.SetLineno(651)
						Label5:
							// line 652: aslist.append('"%s"' % self.getquote())
							πF.SetLineno(652)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßgetquote, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("\"%s\"").ToObject(), πTemp006); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µaslist, "aslist"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µaslist, ßappend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label8
							// line 653: elif self.field[self.pos] in self.atomends:
							πF.SetLineno(653)
						Label6:
							// line 654: break
							πF.SetLineno(654)
							πTemp004 = true
							continue
							goto Label8
						Label7:
							// line 655: else: aslist.append(self.getatom())
							πF.SetLineno(655)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßgetatom, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µaslist, "aslist"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µaslist, ßappend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label8
						Label8:
							// line 656: self.gotonext()
							πF.SetLineno(656)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßgotonext, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßpos, nil); πE != nil {
								continue
							}
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßfield, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp007
							if πTemp007, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp007.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp003, πE = πg.GE(πF, πTemp006, πTemp008); πE != nil {
								continue
							}
							πTemp002 = πTemp003
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label9
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßpos, nil); πE != nil {
								continue
							}
							πTemp006 = πTemp007
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µself, ßfield, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, πTemp008, πTemp006); πE != nil {
								continue
							}
							if πTemp003, πE = πg.NE(πF, πTemp007, πg.NewStr("@").ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp003
						Label9:
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label10
							}
							goto Label11
							// line 658: if self.pos >= len(self.field) or self.field[self.pos] != '@':
							πF.SetLineno(658)
						Label10:
							// line 659: return ''.join(aslist)
							πF.SetLineno(659)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µaslist, "aslist"); πE != nil {
								continue
							}
							πTemp001[0] = µaslist
							if πTemp002, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πR = πTemp003
							continue
							goto Label11
						Label11:
							// line 661: aslist.append('@')
							πF.SetLineno(661)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("@").ToObject()
							if πE = πg.CheckLocal(πF, µaslist, "aslist"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µaslist, ßappend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 662: self.pos += 1
							πF.SetLineno(662)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßpos, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IAdd(πF, πTemp002, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßpos, πTemp003); πE != nil {
								continue
							}
							// line 663: self.gotonext()
							πF.SetLineno(663)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßgotonext, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 664: return ''.join(aslist) + self.getdomain()
							πF.SetLineno(664)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µaslist, "aslist"); πE != nil {
								continue
							}
							πTemp001[0] = µaslist
							if πTemp003, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßgetdomain, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp006, πTemp007); πE != nil {
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
					if πE = πClass.SetItem(πF, ßgetaddrspec.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 643: """Parse an RFC 2822 addr-spec."""
					πF.SetLineno(643)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp008}, πg.NewStr("Parse an RFC 2822 addr-spec.").ToObject()); πE != nil {
						continue
					}
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßgetaddrspec); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp009, ß__doc__, πTemp008); πE != nil {
						continue
					}
					// line 666: def getdomain(self):
					πF.SetLineno(666)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("getdomain", "/usr/lib/python2.7/rfc822.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µsdlist *πg.Object = πg.UnboundLocal; _ = µsdlist
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
							// line 667: """Get the complete domain name from an address."""
							πF.SetLineno(667)
							// line 668: sdlist = []
							πF.SetLineno(668)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µsdlist = πTemp002
							// line 669: while self.pos < len(self.field):
							πF.SetLineno(669)
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
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßpos, nil); πE != nil {
								continue
							}
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßfield, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp006
							if πTemp006, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.LT(πF, πTemp005, πTemp007); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πE != nil || !πTemp004 {
								continue
							}
							πF.PushCheckpoint(1)            
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßpos, nil); πE != nil {
								continue
							}
							πTemp005 = πTemp006
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßfield, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, πTemp007, πTemp005); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßLWS, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Contains(πF, πTemp005, πTemp006); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp004).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label4
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßpos, nil); πE != nil {
								continue
							}
							πTemp005 = πTemp006
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßfield, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, πTemp007, πTemp005); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, πTemp006, πg.NewStr("(").ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label5
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßpos, nil); πE != nil {
								continue
							}
							πTemp005 = πTemp006
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßfield, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, πTemp007, πTemp005); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, πTemp006, πg.NewStr("[").ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label6
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßpos, nil); πE != nil {
								continue
							}
							πTemp005 = πTemp006
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßfield, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, πTemp007, πTemp005); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, πTemp006, πg.NewStr(".").ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label7
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßpos, nil); πE != nil {
								continue
							}
							πTemp005 = πTemp006
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßfield, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, πTemp007, πTemp005); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßatomends, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Contains(πF, πTemp005, πTemp006); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp004).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label8
							}
							goto Label9
							// line 670: if self.field[self.pos] in self.LWS:
							πF.SetLineno(670)
						Label4:
							// line 671: self.pos += 1
							πF.SetLineno(671)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßpos, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IAdd(πF, πTemp002, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßpos, πTemp005); πE != nil {
								continue
							}
							goto Label10
							// line 672: elif self.field[self.pos] == '(':
							πF.SetLineno(672)
						Label5:
							// line 673: self.commentlist.append(self.getcomment())
							πF.SetLineno(673)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßgetcomment, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp005
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcommentlist, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label10
							// line 674: elif self.field[self.pos] == '[':
							πF.SetLineno(674)
						Label6:
							// line 675: sdlist.append(self.getdomainliteral())
							πF.SetLineno(675)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßgetdomainliteral, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp005
							if πE = πg.CheckLocal(πF, µsdlist, "sdlist"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µsdlist, ßappend, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label10
							// line 676: elif self.field[self.pos] == '.':
							πF.SetLineno(676)
						Label7:
							// line 677: self.pos += 1
							πF.SetLineno(677)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßpos, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IAdd(πF, πTemp002, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßpos, πTemp005); πE != nil {
								continue
							}
							// line 678: sdlist.append('.')
							πF.SetLineno(678)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr(".").ToObject()
							if πE = πg.CheckLocal(πF, µsdlist, "sdlist"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µsdlist, ßappend, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label10
							// line 679: elif self.field[self.pos] in self.atomends:
							πF.SetLineno(679)
						Label8:
							// line 680: break
							πF.SetLineno(680)
							πTemp003 = true
							continue
							goto Label10
						Label9:
							// line 681: else: sdlist.append(self.getatom())
							πF.SetLineno(681)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßgetatom, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp005
							if πE = πg.CheckLocal(πF, µsdlist, "sdlist"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µsdlist, ßappend, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label10
						Label10:
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 682: return ''.join(sdlist)
							πF.SetLineno(682)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µsdlist, "sdlist"); πE != nil {
								continue
							}
							πTemp001[0] = µsdlist
							if πTemp002, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßgetdomain.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 667: """Get the complete domain name from an address."""
					πF.SetLineno(667)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp009}, πg.NewStr("Get the complete domain name from an address.").ToObject()); πE != nil {
						continue
					}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßgetdomain); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp010, ß__doc__, πTemp009); πE != nil {
						continue
					}
					// line 684: def getdelimited(self, beginchar, endchars, allowcomments = 1):
					πF.SetLineno(684)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "beginchar", Def: nil}
					πTemp002[2] = πg.Param{Name: "endchars", Def: nil}
					πTemp002[3] = πg.Param{Name: "allowcomments", Def: πg.NewInt(1).ToObject()}
					πTemp009 = πg.NewFunction(πg.NewCode("getdelimited", "/usr/lib/python2.7/rfc822.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µbeginchar *πg.Object = πArgs[1]; _ = µbeginchar
						var µendchars *πg.Object = πArgs[2]; _ = µendchars
						var µallowcomments *πg.Object = πArgs[3]; _ = µallowcomments
						var µslist *πg.Object = πg.UnboundLocal; _ = µslist
						var µquote *πg.Object = πg.UnboundLocal; _ = µquote
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
						var πTemp007 bool
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
							// line 685: """Parse a header fragment delimited by special characters.
							πF.SetLineno(685)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßpos, nil); πE != nil {
								continue
							}
							πTemp002 = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßfield, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp004, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µbeginchar, "beginchar"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.NE(πF, πTemp003, µbeginchar); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							goto Label2
							// line 697: if self.field[self.pos] != beginchar:
							πF.SetLineno(697)
						Label1:
							// line 698: return ''
							πF.SetLineno(698)
							πR = ß.ToObject()
							continue
							goto Label2
						Label2:
							// line 700: slist = ['']
							πF.SetLineno(700)
							πTemp006 = make([]*πg.Object, 1)
							πTemp006[0] = ß.ToObject()
							πTemp001 = πg.NewList(πTemp006...).ToObject()
							µslist = πTemp001
							// line 701: quote = 0
							πF.SetLineno(701)
							µquote = πg.NewInt(0).ToObject()
							// line 702: self.pos += 1
							πF.SetLineno(702)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßpos, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IAdd(πF, πTemp001, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßpos, πTemp002); πE != nil {
								continue
							}
							// line 703: while self.pos < len(self.field):
							πF.SetLineno(703)
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
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßpos, nil); πE != nil {
								continue
							}
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßfield, nil); πE != nil {
								continue
							}
							πTemp006[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πTemp001, πE = πg.LT(πF, πTemp002, πTemp004); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(3)            
							if πE = πg.CheckLocal(πF, µquote, "quote"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, µquote, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label6
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßpos, nil); πE != nil {
								continue
							}
							πTemp002 = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßfield, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp004, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µendchars, "endchars"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Contains(πF, µendchars, πTemp003); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp007).ToObject()
							if πTemp007, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label7
							}
							if πE = πg.CheckLocal(πF, µallowcomments, "allowcomments"); πE != nil {
								continue
							}
							πTemp001 = µallowcomments
							if πTemp007, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp007 {
								goto Label8
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßpos, nil); πE != nil {
								continue
							}
							πTemp003 = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µself, ßfield, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, πTemp008, πTemp003); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, πTemp004, πg.NewStr("(").ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp002
						Label8:
							if πTemp007, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label9
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßpos, nil); πE != nil {
								continue
							}
							πTemp002 = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßfield, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp004, πTemp002); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, πTemp003, πg.NewStr("\\").ToObject()); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label10
							}
							goto Label11
							// line 704: if quote == 1:
							πF.SetLineno(704)
						Label6:
							// line 705: slist.append(self.field[self.pos])
							πF.SetLineno(705)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßpos, nil); πE != nil {
								continue
							}
							πTemp001 = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßfield, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp003, πTemp001); πE != nil {
								continue
							}
							πTemp006[0] = πTemp002
							if πE = πg.CheckLocal(πF, µslist, "slist"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µslist, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							// line 706: quote = 0
							πF.SetLineno(706)
							µquote = πg.NewInt(0).ToObject()
							goto Label12
							// line 707: elif self.field[self.pos] in endchars:
							πF.SetLineno(707)
						Label7:
							// line 708: self.pos += 1
							πF.SetLineno(708)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßpos, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IAdd(πF, πTemp001, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßpos, πTemp002); πE != nil {
								continue
							}
							// line 709: break
							πF.SetLineno(709)
							πTemp005 = true
							continue
							goto Label12
							// line 710: elif allowcomments and self.field[self.pos] == '(':
							πF.SetLineno(710)
						Label9:
							// line 711: slist.append(self.getcomment())
							πF.SetLineno(711)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßgetcomment, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp006[0] = πTemp002
							if πE = πg.CheckLocal(πF, µslist, "slist"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µslist, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							// line 712: continue        # have already advanced pos from getcomment
							πF.SetLineno(712)
							continue
							goto Label12
							// line 713: elif self.field[self.pos] == '\\':
							πF.SetLineno(713)
						Label10:
							// line 714: quote = 1
							πF.SetLineno(714)
							µquote = πg.NewInt(1).ToObject()
							goto Label12
						Label11:
							// line 716: slist.append(self.field[self.pos])
							πF.SetLineno(716)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßpos, nil); πE != nil {
								continue
							}
							πTemp001 = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßfield, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp003, πTemp001); πE != nil {
								continue
							}
							πTemp006[0] = πTemp002
							if πE = πg.CheckLocal(πF, µslist, "slist"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µslist, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							goto Label12
						Label12:
							// line 717: self.pos += 1
							πF.SetLineno(717)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßpos, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IAdd(πF, πTemp001, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßpos, πTemp002); πE != nil {
								continue
							}
							continue
						Label4:
							if πE != nil || πR != nil {
								continue
							}
						Label5:
							// line 719: return ''.join(slist)
							πF.SetLineno(719)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µslist, "slist"); πE != nil {
								continue
							}
							πTemp006[0] = µslist
							if πTemp001, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
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
					if πE = πClass.SetItem(πF, ßgetdelimited.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 685: """Parse a header fragment delimited by special characters.
					πF.SetLineno(685)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp010}, πg.NewStr("Parse a header fragment delimited by special characters.\n\n        `beginchar' is the start character for the fragment.  If self is not\n        looking at an instance of `beginchar' then getdelimited returns the\n        empty string.\n\n        `endchars' is a sequence of allowable end-delimiting characters.\n        Parsing stops when one of these is encountered.\n\n        If `allowcomments' is non-zero, embedded RFC 2822 comments are allowed\n        within the parsed fragment.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ßgetdelimited); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp011, ß__doc__, πTemp010); πE != nil {
						continue
					}
					// line 721: def getquote(self):
					πF.SetLineno(721)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp010 = πg.NewFunction(πg.NewCode("getquote", "/usr/lib/python2.7/rfc822.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 722: """Get a quote-delimited fragment from self's field."""
							πF.SetLineno(722)
							// line 723: return self.getdelimited('"', '"\r', 0)
							πF.SetLineno(723)
							πTemp001 = πF.MakeArgs(3)
							πTemp001[0] = πg.NewStr("\"").ToObject()
							πTemp001[1] = πg.NewStr("\"\r").ToObject()
							πTemp001[2] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßgetdelimited, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßgetquote.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 722: """Get a quote-delimited fragment from self's field."""
					πF.SetLineno(722)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πg.NewStr("Get a quote-delimited fragment from self's field.").ToObject()); πE != nil {
						continue
					}
					if πTemp012, πE = πg.ResolveClass(πF, πClass, nil, ßgetquote); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp012, ß__doc__, πTemp011); πE != nil {
						continue
					}
					// line 725: def getcomment(self):
					πF.SetLineno(725)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp011 = πg.NewFunction(πg.NewCode("getcomment", "/usr/lib/python2.7/rfc822.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 726: """Get a parenthesis-delimited fragment from self's field."""
							πF.SetLineno(726)
							// line 727: return self.getdelimited('(', ')\r', 1)
							πF.SetLineno(727)
							πTemp001 = πF.MakeArgs(3)
							πTemp001[0] = πg.NewStr("(").ToObject()
							πTemp001[1] = πg.NewStr(")\r").ToObject()
							πTemp001[2] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßgetdelimited, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßgetcomment.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 726: """Get a parenthesis-delimited fragment from self's field."""
					πF.SetLineno(726)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp012}, πg.NewStr("Get a parenthesis-delimited fragment from self's field.").ToObject()); πE != nil {
						continue
					}
					if πTemp013, πE = πg.ResolveClass(πF, πClass, nil, ßgetcomment); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp013, ß__doc__, πTemp012); πE != nil {
						continue
					}
					// line 729: def getdomainliteral(self):
					πF.SetLineno(729)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp012 = πg.NewFunction(πg.NewCode("getdomainliteral", "/usr/lib/python2.7/rfc822.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 730: """Parse an RFC 2822 domain-literal."""
							πF.SetLineno(730)
							// line 731: return '[%s]' % self.getdelimited('[', ']\r', 0)
							πF.SetLineno(731)
							πTemp002 = πF.MakeArgs(3)
							πTemp002[0] = πg.NewStr("[").ToObject()
							πTemp002[1] = πg.NewStr("]\r").ToObject()
							πTemp002[2] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßgetdelimited, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("[%s]").ToObject(), πTemp004); πE != nil {
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
					if πE = πClass.SetItem(πF, ßgetdomainliteral.ToObject(), πTemp012); πE != nil {
						continue
					}
					// line 730: """Parse an RFC 2822 domain-literal."""
					πF.SetLineno(730)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp013}, πg.NewStr("Parse an RFC 2822 domain-literal.").ToObject()); πE != nil {
						continue
					}
					if πTemp014, πE = πg.ResolveClass(πF, πClass, nil, ßgetdomainliteral); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp014, ß__doc__, πTemp013); πE != nil {
						continue
					}
					// line 733: def getatom(self, atomends=None):
					πF.SetLineno(733)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					if πTemp014, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[1] = πg.Param{Name: "atomends", Def: πTemp014}
					πTemp013 = πg.NewFunction(πg.NewCode("getatom", "/usr/lib/python2.7/rfc822.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µatomends *πg.Object = πArgs[1]; _ = µatomends
						var µatomlist *πg.Object = πg.UnboundLocal; _ = µatomlist
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
							case 3: goto Label3
							case 4: goto Label4
							default: panic("unexpected function state")
							}
							// line 734: """Parse an RFC 2822 atom.
							πF.SetLineno(734)
							// line 740: atomlist = ['']
							πF.SetLineno(740)
							πTemp001 = make([]*πg.Object, 1)
							πTemp001[0] = ß.ToObject()
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µatomlist = πTemp002
							if πE = πg.CheckLocal(πF, µatomends, "atomends"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(µatomends == πTemp003).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 741: if atomends is None:
							πF.SetLineno(741)
						Label1:
							// line 742: atomends = self.atomends
							πF.SetLineno(742)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßatomends, nil); πE != nil {
								continue
							}
							µatomends = πTemp002
							goto Label2
						Label2:
							// line 744: while self.pos < len(self.field):
							πF.SetLineno(744)
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
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßpos, nil); πE != nil {
								continue
							}
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßfield, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp006
							if πTemp006, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.LT(πF, πTemp003, πTemp007); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(3)            
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßpos, nil); πE != nil {
								continue
							}
							πTemp003 = πTemp006
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßfield, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, πTemp007, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µatomends, "atomends"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Contains(πF, µatomends, πTemp006); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp005).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label6
							}
							goto Label7
							// line 745: if self.field[self.pos] in atomends:
							πF.SetLineno(745)
						Label6:
							// line 746: break
							πF.SetLineno(746)
							πTemp004 = true
							continue
							goto Label8
						Label7:
							// line 747: else: atomlist.append(self.field[self.pos])
							πF.SetLineno(747)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßpos, nil); πE != nil {
								continue
							}
							πTemp002 = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßfield, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp006, πTemp002); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µatomlist, "atomlist"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µatomlist, ßappend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label8
						Label8:
							// line 748: self.pos += 1
							πF.SetLineno(748)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßpos, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IAdd(πF, πTemp002, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßpos, πTemp003); πE != nil {
								continue
							}
							continue
						Label4:
							if πE != nil || πR != nil {
								continue
							}
						Label5:
							// line 750: return ''.join(atomlist)
							πF.SetLineno(750)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µatomlist, "atomlist"); πE != nil {
								continue
							}
							πTemp001[0] = µatomlist
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
					if πE = πClass.SetItem(πF, ßgetatom.ToObject(), πTemp013); πE != nil {
						continue
					}
					// line 734: """Parse an RFC 2822 atom.
					πF.SetLineno(734)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp014}, πg.NewStr("Parse an RFC 2822 atom.\n\n        Optional atomends specifies a different set of end token delimiters\n        (the default is to use self.atomends).  This is used e.g. in\n        getphraselist() since phrase endings must not include the `.' (which\n        is legal in phrases).").ToObject()); πE != nil {
						continue
					}
					if πTemp015, πE = πg.ResolveClass(πF, πClass, nil, ßgetatom); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp015, ß__doc__, πTemp014); πE != nil {
						continue
					}
					// line 752: def getphraselist(self):
					πF.SetLineno(752)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp014 = πg.NewFunction(πg.NewCode("getphraselist", "/usr/lib/python2.7/rfc822.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µplist *πg.Object = πg.UnboundLocal; _ = µplist
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
							// line 753: """Parse a sequence of RFC 2822 phrases.
							πF.SetLineno(753)
							// line 759: plist = []
							πF.SetLineno(759)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µplist = πTemp002
							// line 761: while self.pos < len(self.field):
							πF.SetLineno(761)
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
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßpos, nil); πE != nil {
								continue
							}
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßfield, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp006
							if πTemp006, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.LT(πF, πTemp005, πTemp007); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πE != nil || !πTemp004 {
								continue
							}
							πF.PushCheckpoint(1)            
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßpos, nil); πE != nil {
								continue
							}
							πTemp005 = πTemp006
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßfield, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, πTemp007, πTemp005); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßLWS, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Contains(πF, πTemp005, πTemp006); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp004).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label4
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßpos, nil); πE != nil {
								continue
							}
							πTemp005 = πTemp006
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßfield, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, πTemp007, πTemp005); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, πTemp006, πg.NewStr("\"").ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label5
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßpos, nil); πE != nil {
								continue
							}
							πTemp005 = πTemp006
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßfield, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, πTemp007, πTemp005); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, πTemp006, πg.NewStr("(").ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label6
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßpos, nil); πE != nil {
								continue
							}
							πTemp005 = πTemp006
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßfield, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, πTemp007, πTemp005); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßphraseends, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Contains(πF, πTemp005, πTemp006); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp004).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label7
							}
							goto Label8
							// line 762: if self.field[self.pos] in self.LWS:
							πF.SetLineno(762)
						Label4:
							// line 763: self.pos += 1
							πF.SetLineno(763)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßpos, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IAdd(πF, πTemp002, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßpos, πTemp005); πE != nil {
								continue
							}
							goto Label9
							// line 764: elif self.field[self.pos] == '"':
							πF.SetLineno(764)
						Label5:
							// line 765: plist.append(self.getquote())
							πF.SetLineno(765)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßgetquote, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp005
							if πE = πg.CheckLocal(πF, µplist, "plist"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µplist, ßappend, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label9
							// line 766: elif self.field[self.pos] == '(':
							πF.SetLineno(766)
						Label6:
							// line 767: self.commentlist.append(self.getcomment())
							πF.SetLineno(767)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßgetcomment, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp005
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcommentlist, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label9
							// line 768: elif self.field[self.pos] in self.phraseends:
							πF.SetLineno(768)
						Label7:
							// line 769: break
							πF.SetLineno(769)
							πTemp003 = true
							continue
							goto Label9
						Label8:
							// line 771: plist.append(self.getatom(self.phraseends))
							πF.SetLineno(771)
							πTemp001 = πF.MakeArgs(1)
							πTemp008 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßphraseends, nil); πE != nil {
								continue
							}
							πTemp008[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßgetatom, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp001[0] = πTemp005
							if πE = πg.CheckLocal(πF, µplist, "plist"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µplist, ßappend, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label9
						Label9:
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 773: return plist
							πF.SetLineno(773)
							if πE = πg.CheckLocal(πF, µplist, "plist"); πE != nil {
								continue
							}
							πR = µplist
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßgetphraselist.ToObject(), πTemp014); πE != nil {
						continue
					}
					// line 753: """Parse a sequence of RFC 2822 phrases.
					πF.SetLineno(753)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp015}, πg.NewStr("Parse a sequence of RFC 2822 phrases.\n\n        A phrase is a sequence of words, which are in turn either RFC 2822\n        atoms or quoted-strings.  Phrases are canonicalized by squeezing all\n        runs of continuous whitespace into one space.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp016, πE = πg.ResolveClass(πF, πClass, nil, ßgetphraselist); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp016, ß__doc__, πTemp015); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp009, πE = πTemp005.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp009 == nil {
				πTemp009 = πg.TypeType.ToObject()
			}
			if πTemp010, πE = πTemp009.Call(πF, []*πg.Object{πg.NewStr("AddrlistClass").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp005.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßAddrlistClass.ToObject(), πTemp010); πE != nil {
				continue
			}
			// line 775: class AddressList(AddrlistClass):
			πF.SetLineno(775)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp010, πE = πg.ResolveGlobal(πF, ßAddrlistClass); πE != nil {
				continue
			}
			πTemp002[0] = πTemp010
			πTemp005 = πg.NewDict()
			if πTemp008, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp005.SetItem(πF, ß__module__.ToObject(), πTemp008); πE != nil {
				continue
			}
			_, πE = πg.NewCode("AddressList", "/usr/lib/python2.7/rfc822.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp005
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 776: """An AddressList encapsulates a list of parsed RFC 2822 addresses."""
					πF.SetLineno(776)
					// line 776: """An AddressList encapsulates a list of parsed RFC 2822 addresses."""
					πF.SetLineno(776)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("An AddressList encapsulates a list of parsed RFC 2822 addresses.").ToObject()); πE != nil {
						continue
					}
					// line 777: def __init__(self, field):
					πF.SetLineno(777)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "field", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/rfc822.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µfield *πg.Object = πArgs[1]; _ = µfield
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
							// line 778: AddrlistClass.__init__(self, field)
							πF.SetLineno(778)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µfield, "field"); πE != nil {
								continue
							}
							πTemp001[1] = µfield
							if πTemp002, πE = πg.ResolveGlobal(πF, ßAddrlistClass); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ß__init__, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.CheckLocal(πF, µfield, "field"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µfield); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 779: if field:
							πF.SetLineno(779)
						Label1:
							// line 780: self.addresslist = self.getaddrlist()
							πF.SetLineno(780)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßgetaddrlist, nil); πE != nil {
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
							if πE = πg.SetAttr(πF, µself, ßaddresslist, πTemp002); πE != nil {
								continue
							}
							goto Label3
						Label2:
							// line 782: self.addresslist = []
							πF.SetLineno(782)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßaddresslist, πTemp003); πE != nil {
								continue
							}
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
					if πE = πClass.SetItem(πF, ß__init__.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 784: def __len__(self):
					πF.SetLineno(784)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("__len__", "/usr/lib/python2.7/rfc822.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 785: return len(self.addresslist)
							πF.SetLineno(785)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßaddresslist, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__len__.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 787: def __str__(self):
					πF.SetLineno(787)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("__str__", "/usr/lib/python2.7/rfc822.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 788: return ", ".join(map(dump_address_pair, self.addresslist))
							πF.SetLineno(788)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(2)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßdump_address_pair); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßaddresslist, nil); πE != nil {
								continue
							}
							πTemp002[1] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßmap); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							if πTemp003, πE = πg.GetAttr(πF, πg.NewStr(", ").ToObject(), ßjoin, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__str__.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 790: def __add__(self, other):
					πF.SetLineno(790)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "other", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("__add__", "/usr/lib/python2.7/rfc822.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
						var µnewaddr *πg.Object = πg.UnboundLocal; _ = µnewaddr
						var µx *πg.Object = πg.UnboundLocal; _ = µx
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							// line 792: newaddr = AddressList(None)
							πF.SetLineno(792)
							πTemp001 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßAddressList); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µnewaddr = πTemp003
							// line 793: newaddr.addresslist = self.addresslist[:]
							πF.SetLineno(793)
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßaddresslist, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp004, πTemp002); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnewaddr, "newaddr"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µnewaddr, ßaddresslist, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µother, ßaddresslist, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, πTemp003); πE != nil {
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
								µx = πTemp003
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(1)            
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßaddresslist, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Contains(πF, πTemp007, µx); πE != nil {
								continue
							}
							πTemp004 = πg.GetBool(πTemp006).ToObject()
							if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp006).ToObject()
							if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label4
							}
							goto Label5
							// line 795: if not x in self.addresslist:
							πF.SetLineno(795)
						Label4:
							// line 796: newaddr.addresslist.append(x)
							πF.SetLineno(796)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp001[0] = µx
							if πE = πg.CheckLocal(πF, µnewaddr, "newaddr"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µnewaddr, ßaddresslist, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßappend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
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
							// line 797: return newaddr
							πF.SetLineno(797)
							if πE = πg.CheckLocal(πF, µnewaddr, "newaddr"); πE != nil {
								continue
							}
							πR = µnewaddr
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß__add__.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 799: def __iadd__(self, other):
					πF.SetLineno(799)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "other", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("__iadd__", "/usr/lib/python2.7/rfc822.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
						var µx *πg.Object = πg.UnboundLocal; _ = µx
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µother, ßaddresslist, nil); πE != nil {
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
								µx = πTemp002
							}
							if πE != nil || !πTemp004 {
								continue
							}
							πF.PushCheckpoint(1)            
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßaddresslist, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Contains(πF, πTemp006, µx); πE != nil {
								continue
							}
							πTemp005 = πg.GetBool(πTemp004).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp005); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp004).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label4
							}
							goto Label5
							// line 802: if not x in self.addresslist:
							πF.SetLineno(802)
						Label4:
							// line 803: self.addresslist.append(x)
							πF.SetLineno(803)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp007[0] = µx
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßaddresslist, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp005.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							goto Label5
						Label5:
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 804: return self
							πF.SetLineno(804)
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
					if πE = πClass.SetItem(πF, ß__iadd__.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 806: def __sub__(self, other):
					πF.SetLineno(806)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "other", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("__sub__", "/usr/lib/python2.7/rfc822.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
						var µnewaddr *πg.Object = πg.UnboundLocal; _ = µnewaddr
						var µx *πg.Object = πg.UnboundLocal; _ = µx
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
							// line 808: newaddr = AddressList(None)
							πF.SetLineno(808)
							πTemp001 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßAddressList); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µnewaddr = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßaddresslist, nil); πE != nil {
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
								µx = πTemp003
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(1)            
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µother, ßaddresslist, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Contains(πF, πTemp007, µx); πE != nil {
								continue
							}
							πTemp006 = πg.GetBool(πTemp005).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp005).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label4
							}
							goto Label5
							// line 810: if not x in other.addresslist:
							πF.SetLineno(810)
						Label4:
							// line 811: newaddr.addresslist.append(x)
							πF.SetLineno(811)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp001[0] = µx
							if πE = πg.CheckLocal(πF, µnewaddr, "newaddr"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µnewaddr, ßaddresslist, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp003, ßappend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
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
							// line 812: return newaddr
							πF.SetLineno(812)
							if πE = πg.CheckLocal(πF, µnewaddr, "newaddr"); πE != nil {
								continue
							}
							πR = µnewaddr
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß__sub__.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 814: def __isub__(self, other):
					πF.SetLineno(814)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "other", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("__isub__", "/usr/lib/python2.7/rfc822.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
						var µx *πg.Object = πg.UnboundLocal; _ = µx
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
						var πTemp006 []*πg.Object
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
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µother, ßaddresslist, nil); πE != nil {
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
								µx = πTemp002
							}
							if πE != nil || !πTemp004 {
								continue
							}
							πF.PushCheckpoint(1)            
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßaddresslist, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Contains(πF, πTemp005, µx); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp004).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label4
							}
							goto Label5
							// line 817: if x in self.addresslist:
							πF.SetLineno(817)
						Label4:
							// line 818: self.addresslist.remove(x)
							πF.SetLineno(818)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp006[0] = µx
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßaddresslist, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp002, ßremove, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp005.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							goto Label5
						Label5:
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 819: return self
							πF.SetLineno(819)
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
					if πE = πClass.SetItem(πF, ß__isub__.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 821: def __getitem__(self, index):
					πF.SetLineno(821)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "index", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("__getitem__", "/usr/lib/python2.7/rfc822.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µindex *πg.Object = πArgs[1]; _ = µindex
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
							// line 823: return self.addresslist[index]
							πF.SetLineno(823)
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							πTemp001 = µindex
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßaddresslist, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__getitem__.ToObject(), πTemp009); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp009, πE = πTemp005.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp009 == nil {
				πTemp009 = πg.TypeType.ToObject()
			}
			if πTemp010, πE = πTemp009.Call(πF, []*πg.Object{πg.NewStr("AddressList").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp005.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßAddressList.ToObject(), πTemp010); πE != nil {
				continue
			}
			// line 825: def dump_address_pair(pair):
			πF.SetLineno(825)
			πTemp007 = make([]πg.Param, 1)
			πTemp007[0] = πg.Param{Name: "pair", Def: nil}
			πTemp008 = πg.NewFunction(πg.NewCode("dump_address_pair", "/usr/lib/python2.7/rfc822.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µpair *πg.Object = πArgs[0]; _ = µpair
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
					// line 826: """Dump a (name, address) pair in a canonicalized form."""
					πF.SetLineno(826)
					πTemp001 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µpair, "pair"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetItem(πF, µpair, πTemp001); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label1
					}
					goto Label2
					// line 827: if pair[0]:
					πF.SetLineno(827)
				Label1:
					// line 828: return '"' + pair[0] + '" <' + pair[1] + '>'
					πF.SetLineno(828)
					πTemp006 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µpair, "pair"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetItem(πF, µpair, πTemp006); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Add(πF, πg.NewStr("\"").ToObject(), πTemp007); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Add(πF, πTemp005, πg.NewStr("\" <").ToObject()); πE != nil {
						continue
					}
					πTemp005 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µpair, "pair"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, µpair, πTemp005); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Add(πF, πTemp004, πTemp006); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, πTemp002, πg.NewStr(">").ToObject()); πE != nil {
						continue
					}
					πR = πTemp001
					continue
					goto Label3
				Label2:
					// line 830: return pair[1]
					πF.SetLineno(830)
					πTemp001 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µpair, "pair"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetItem(πF, µpair, πTemp001); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßdump_address_pair.ToObject(), πTemp008); πE != nil {
				continue
			}
			// line 826: """Dump a (name, address) pair in a canonicalized form."""
			πF.SetLineno(826)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp009}, πg.NewStr("Dump a (name, address) pair in a canonicalized form.").ToObject()); πE != nil {
				continue
			}
			if πTemp010, πE = πg.ResolveGlobal(πF, ßdump_address_pair); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp010, ß__doc__, πTemp009); πE != nil {
				continue
			}
			// line 834: _monthnames = ['jan', 'feb', 'mar', 'apr', 'may', 'jun', 'jul',
			πF.SetLineno(834)
			πTemp002 = make([]*πg.Object, 24)
			πTemp002[0] = ßjan.ToObject()
			πTemp002[1] = ßfeb.ToObject()
			πTemp002[2] = ßmar.ToObject()
			πTemp002[3] = ßapr.ToObject()
			πTemp002[4] = ßmay.ToObject()
			πTemp002[5] = ßjun.ToObject()
			πTemp002[6] = ßjul.ToObject()
			πTemp002[7] = ßaug.ToObject()
			πTemp002[8] = ßsep.ToObject()
			πTemp002[9] = ßoct.ToObject()
			πTemp002[10] = ßnov.ToObject()
			πTemp002[11] = ßdec.ToObject()
			πTemp002[12] = ßjanuary.ToObject()
			πTemp002[13] = ßfebruary.ToObject()
			πTemp002[14] = ßmarch.ToObject()
			πTemp002[15] = ßapril.ToObject()
			πTemp002[16] = ßmay.ToObject()
			πTemp002[17] = ßjune.ToObject()
			πTemp002[18] = ßjuly.ToObject()
			πTemp002[19] = ßaugust.ToObject()
			πTemp002[20] = ßseptember.ToObject()
			πTemp002[21] = ßoctober.ToObject()
			πTemp002[22] = ßnovember.ToObject()
			πTemp002[23] = ßdecember.ToObject()
			πTemp009 = πg.NewList(πTemp002...).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_monthnames.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 838: _daynames = ['mon', 'tue', 'wed', 'thu', 'fri', 'sat', 'sun']
			πF.SetLineno(838)
			πTemp002 = make([]*πg.Object, 7)
			πTemp002[0] = ßmon.ToObject()
			πTemp002[1] = ßtue.ToObject()
			πTemp002[2] = ßwed.ToObject()
			πTemp002[3] = ßthu.ToObject()
			πTemp002[4] = ßfri.ToObject()
			πTemp002[5] = ßsat.ToObject()
			πTemp002[6] = ßsun.ToObject()
			πTemp009 = πg.NewList(πTemp002...).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_daynames.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 846: _timezones = {'UT':0, 'UTC':0, 'GMT':0, 'Z':0,
			πF.SetLineno(846)
			πTemp005 = πg.NewDict()
			if πE = πTemp005.SetItem(πF, ßUT.ToObject(), πg.NewInt(0).ToObject()); πE != nil {
				continue
			}
			if πE = πTemp005.SetItem(πF, ßUTC.ToObject(), πg.NewInt(0).ToObject()); πE != nil {
				continue
			}
			if πE = πTemp005.SetItem(πF, ßGMT.ToObject(), πg.NewInt(0).ToObject()); πE != nil {
				continue
			}
			if πE = πTemp005.SetItem(πF, ßZ.ToObject(), πg.NewInt(0).ToObject()); πE != nil {
				continue
			}
			if πTemp009, πE = πg.Neg(πF, πg.NewInt(400).ToObject()); πE != nil {
				continue
			}
			if πE = πTemp005.SetItem(πF, ßAST.ToObject(), πTemp009); πE != nil {
				continue
			}
			if πTemp009, πE = πg.Neg(πF, πg.NewInt(300).ToObject()); πE != nil {
				continue
			}
			if πE = πTemp005.SetItem(πF, ßADT.ToObject(), πTemp009); πE != nil {
				continue
			}
			if πTemp009, πE = πg.Neg(πF, πg.NewInt(500).ToObject()); πE != nil {
				continue
			}
			if πE = πTemp005.SetItem(πF, ßEST.ToObject(), πTemp009); πE != nil {
				continue
			}
			if πTemp009, πE = πg.Neg(πF, πg.NewInt(400).ToObject()); πE != nil {
				continue
			}
			if πE = πTemp005.SetItem(πF, ßEDT.ToObject(), πTemp009); πE != nil {
				continue
			}
			if πTemp009, πE = πg.Neg(πF, πg.NewInt(600).ToObject()); πE != nil {
				continue
			}
			if πE = πTemp005.SetItem(πF, ßCST.ToObject(), πTemp009); πE != nil {
				continue
			}
			if πTemp009, πE = πg.Neg(πF, πg.NewInt(500).ToObject()); πE != nil {
				continue
			}
			if πE = πTemp005.SetItem(πF, ßCDT.ToObject(), πTemp009); πE != nil {
				continue
			}
			if πTemp009, πE = πg.Neg(πF, πg.NewInt(700).ToObject()); πE != nil {
				continue
			}
			if πE = πTemp005.SetItem(πF, ßMST.ToObject(), πTemp009); πE != nil {
				continue
			}
			if πTemp009, πE = πg.Neg(πF, πg.NewInt(600).ToObject()); πE != nil {
				continue
			}
			if πE = πTemp005.SetItem(πF, ßMDT.ToObject(), πTemp009); πE != nil {
				continue
			}
			if πTemp009, πE = πg.Neg(πF, πg.NewInt(800).ToObject()); πE != nil {
				continue
			}
			if πE = πTemp005.SetItem(πF, ßPST.ToObject(), πTemp009); πE != nil {
				continue
			}
			if πTemp009, πE = πg.Neg(πF, πg.NewInt(700).ToObject()); πE != nil {
				continue
			}
			if πE = πTemp005.SetItem(πF, ßPDT.ToObject(), πTemp009); πE != nil {
				continue
			}
			πTemp009 = πTemp005.ToObject()
			if πE = πF.Globals().SetItem(πF, ß_timezones.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 855: def parsedate_tz(data):
			πF.SetLineno(855)
			πTemp007 = make([]πg.Param, 1)
			πTemp007[0] = πg.Param{Name: "data", Def: nil}
			πTemp009 = πg.NewFunction(πg.NewCode("parsedate_tz", "/usr/lib/python2.7/rfc822.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µdata *πg.Object = πArgs[0]; _ = µdata
				var µi *πg.Object = πg.UnboundLocal; _ = µi
				var µstuff *πg.Object = πg.UnboundLocal; _ = µstuff
				var µs *πg.Object = πg.UnboundLocal; _ = µs
				var µdd *πg.Object = πg.UnboundLocal; _ = µdd
				var µmm *πg.Object = πg.UnboundLocal; _ = µmm
				var µyy *πg.Object = πg.UnboundLocal; _ = µyy
				var µtm *πg.Object = πg.UnboundLocal; _ = µtm
				var µtz *πg.Object = πg.UnboundLocal; _ = µtz
				var µthh *πg.Object = πg.UnboundLocal; _ = µthh
				var µtmm *πg.Object = πg.UnboundLocal; _ = µtmm
				var µtss *πg.Object = πg.UnboundLocal; _ = µtss
				var µtzoffset *πg.Object = πg.UnboundLocal; _ = µtzoffset
				var µtzsign *πg.Object = πg.UnboundLocal; _ = µtzsign
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
				var πTemp008 bool
				_ = πTemp008
				var πTemp009 []*πg.Object
				_ = πTemp009
				var πTemp010 *πg.BaseException
				_ = πTemp010
				var πTemp011 *πg.Traceback
				_ = πTemp011
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 41: goto Label41
					case 47: goto Label47
					default: panic("unexpected function state")
					}
					// line 856: """Convert a date string to a time tuple.
					πF.SetLineno(856)
					if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, µdata); πE != nil {
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
					// line 860: if not data:
					πF.SetLineno(860)
				Label1:
					// line 861: return None
					πF.SetLineno(861)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πR = πTemp001
					continue
					goto Label2
				Label2:
					// line 862: data = data.split()
					πF.SetLineno(862)
					if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µdata, ßsplit, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					µdata = πTemp003
					if πTemp005, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp004 = πTemp005
					πTemp006 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetItem(πF, µdata, πTemp006); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, πTemp007, πTemp004); πE != nil {
						continue
					}
					πTemp004 = πg.NewTuple2(πg.NewStr(",").ToObject(), πg.NewStr(".").ToObject()).ToObject()
					if πTemp008, πE = πg.Contains(πF, πTemp004, πTemp005); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(πTemp008).ToObject()
					πTemp001 = πTemp003
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label3
					}
					πTemp004 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µdata, πTemp004); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp005, ßlower, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ß_daynames); πE != nil {
						continue
					}
					if πTemp008, πE = πg.Contains(πF, πTemp004, πTemp005); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(πTemp008).ToObject()
					πTemp001 = πTemp003
				Label3:
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label4
					}
					goto Label5
					// line 863: if data[0][-1] in (',', '.') or data[0].lower() in _daynames:
					πF.SetLineno(863)
				Label4:
					// line 865: del data[0]
					πF.SetLineno(865)
					if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
						continue
					}
					πTemp001 = πg.NewInt(0).ToObject()
					if πE = πg.DelItem(πF, µdata, πTemp001); πE != nil {
						continue
					}
					goto Label6
				Label5:
					// line 868: i = data[0].rfind(',')
					πF.SetLineno(868)
					πTemp009 = πF.MakeArgs(1)
					πTemp009[0] = πg.NewStr(",").ToObject()
					πTemp001 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µdata, πTemp001); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, πTemp003, ßrfind, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp009, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp009)
					µi = πTemp003
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GE(πF, µi, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label7
					}
					goto Label8
					// line 869: if i >= 0:
					πF.SetLineno(869)
				Label7:
					// line 870: data[0] = data[0][i+1:]
					πF.SetLineno(870)
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Add(πF, µi, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πTemp003, πg.None, πg.None}, nil); πE != nil {
						continue
					}
					πTemp004 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µdata, πTemp004); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, πTemp005, πTemp001); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
						continue
					}
					πTemp004 = πg.NewInt(0).ToObject()
					if πE = πg.SetItem(πF, µdata, πTemp004, πTemp001); πE != nil {
						continue
					}
					goto Label8
				Label8:
					goto Label6
				Label6:
					πTemp009 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
						continue
					}
					πTemp009[0] = µdata
					if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp009, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp009)
					if πTemp001, πE = πg.Eq(πF, πTemp004, πg.NewInt(3).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label9
					}
					goto Label10
					// line 871: if len(data) == 3: # RFC 850 date, deprecated
					πF.SetLineno(871)
				Label9:
					// line 872: stuff = data[0].split('-')
					πF.SetLineno(872)
					πTemp009 = πF.MakeArgs(1)
					πTemp009[0] = πg.NewStr("-").ToObject()
					πTemp001 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µdata, πTemp001); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, πTemp003, ßsplit, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp009, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp009)
					µstuff = πTemp003
					πTemp009 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µstuff, "stuff"); πE != nil {
						continue
					}
					πTemp009[0] = µstuff
					if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp009, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp009)
					if πTemp001, πE = πg.Eq(πF, πTemp004, πg.NewInt(3).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label11
					}
					goto Label12
					// line 873: if len(stuff) == 3:
					πF.SetLineno(873)
				Label11:
					// line 874: data = stuff + data[1:]
					πF.SetLineno(874)
					if πE = πg.CheckLocal(πF, µstuff, "stuff"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µdata, πTemp003); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, µstuff, πTemp004); πE != nil {
						continue
					}
					µdata = πTemp001
					goto Label12
				Label12:
					goto Label10
				Label10:
					πTemp009 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
						continue
					}
					πTemp009[0] = µdata
					if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp009, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp009)
					if πTemp001, πE = πg.Eq(πF, πTemp004, πg.NewInt(4).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label13
					}
					goto Label14
					// line 875: if len(data) == 4:
					πF.SetLineno(875)
				Label13:
					// line 876: s = data[3]
					πF.SetLineno(876)
					πTemp001 = πg.NewInt(3).ToObject()
					if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µdata, πTemp001); πE != nil {
						continue
					}
					µs = πTemp003
					// line 877: i = s.find('+')
					πF.SetLineno(877)
					πTemp009 = πF.MakeArgs(1)
					πTemp009[0] = πg.NewStr("+").ToObject()
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µs, ßfind, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp009, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp009)
					µi = πTemp003
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GT(πF, µi, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label15
					}
					goto Label16
					// line 878: if i > 0:
					πF.SetLineno(878)
				Label15:
					// line 879: data[3:] = [s[:i], s[i+1:]]
					πF.SetLineno(879)
					πTemp009 = make([]*πg.Object, 2)
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.None, µi, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µs, πTemp001); πE != nil {
						continue
					}
					πTemp009[0] = πTemp003
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Add(πF, µi, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πTemp003, πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µs, πTemp001); πE != nil {
						continue
					}
					πTemp009[1] = πTemp003
					πTemp001 = πg.NewList(πTemp009...).ToObject()
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(3).ToObject(), πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.SetItem(πF, µdata, πTemp004, πTemp003); πE != nil {
						continue
					}
					goto Label17
				Label16:
					// line 881: data.append('') # Dummy tz
					πF.SetLineno(881)
					πTemp009 = πF.MakeArgs(1)
					πTemp009[0] = ß.ToObject()
					if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µdata, ßappend, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp009, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp009)
					goto Label17
				Label17:
					goto Label14
				Label14:
					πTemp009 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
						continue
					}
					πTemp009[0] = µdata
					if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp009, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp009)
					if πTemp001, πE = πg.LT(πF, πTemp004, πg.NewInt(5).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label18
					}
					goto Label19
					// line 882: if len(data) < 5:
					πF.SetLineno(882)
				Label18:
					// line 883: return None
					πF.SetLineno(883)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πR = πTemp001
					continue
					goto Label19
				Label19:
					// line 884: data = data[:5]
					πF.SetLineno(884)
					if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.NewInt(5).ToObject(), πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µdata, πTemp001); πE != nil {
						continue
					}
					µdata = πTemp003
					// line 885: [dd, mm, yy, tm, tz] = data
					πF.SetLineno(885)
					if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp006}}}, µdata); πE != nil {
						continue
					}
					µdd = πTemp001
					µmm = πTemp003
					µyy = πTemp004
					µtm = πTemp005
					µtz = πTemp006
					// line 886: mm = mm.lower()
					πF.SetLineno(886)
					if πE = πg.CheckLocal(πF, µmm, "mm"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µmm, ßlower, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					µmm = πTemp003
					if πE = πg.CheckLocal(πF, µmm, "mm"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ß_monthnames); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Contains(πF, πTemp004, µmm); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(πTemp002).ToObject()
					if πTemp002, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(!πTemp002).ToObject()
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label20
					}
					goto Label21
					// line 887: if not mm in _monthnames:
					πF.SetLineno(887)
				Label20:
					// line 888: dd, mm = mm, dd.lower()
					πF.SetLineno(888)
					if πE = πg.CheckLocal(πF, µmm, "mm"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µdd, "dd"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µdd, ßlower, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple2(µmm, πTemp004).ToObject()
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp004}}}, πTemp001); πE != nil {
						continue
					}
					µdd = πTemp003
					µmm = πTemp004
					if πE = πg.CheckLocal(πF, µmm, "mm"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ß_monthnames); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Contains(πF, πTemp004, µmm); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(πTemp002).ToObject()
					if πTemp002, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(!πTemp002).ToObject()
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label22
					}
					goto Label23
					// line 889: if not mm in _monthnames:
					πF.SetLineno(889)
				Label22:
					// line 890: return None
					πF.SetLineno(890)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πR = πTemp001
					continue
					goto Label23
				Label23:
					goto Label21
				Label21:
					// line 891: mm = _monthnames.index(mm)+1
					πF.SetLineno(891)
					πTemp009 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µmm, "mm"); πE != nil {
						continue
					}
					πTemp009[0] = µmm
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_monthnames); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßindex, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp009, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp009)
					if πTemp001, πE = πg.Add(πF, πTemp003, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µmm = πTemp001
					if πE = πg.CheckLocal(πF, µmm, "mm"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GT(πF, µmm, πg.NewInt(12).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label24
					}
					goto Label25
					// line 892: if mm > 12: mm = mm - 12
					πF.SetLineno(892)
				Label24:
					// line 892: if mm > 12: mm = mm - 12
					πF.SetLineno(892)
					if πE = πg.CheckLocal(πF, µmm, "mm"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Sub(πF, µmm, πg.NewInt(12).ToObject()); πE != nil {
						continue
					}
					µmm = πTemp001
					goto Label25
				Label25:
					if πTemp004, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp003 = πTemp004
					if πE = πg.CheckLocal(πF, µdd, "dd"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µdd, πTemp003); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, πTemp004, πg.NewStr(",").ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label26
					}
					goto Label27
					// line 893: if dd[-1] == ',':
					πF.SetLineno(893)
				Label26:
					// line 894: dd = dd[:-1]
					πF.SetLineno(894)
					if πTemp003, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πTemp003, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µdd, "dd"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µdd, πTemp001); πE != nil {
						continue
					}
					µdd = πTemp003
					goto Label27
				Label27:
					// line 895: i = yy.find(':')
					πF.SetLineno(895)
					πTemp009 = πF.MakeArgs(1)
					πTemp009[0] = πg.NewStr(":").ToObject()
					if πE = πg.CheckLocal(πF, µyy, "yy"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µyy, ßfind, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp009, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp009)
					µi = πTemp003
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GT(πF, µi, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label28
					}
					goto Label29
					// line 896: if i > 0:
					πF.SetLineno(896)
				Label28:
					// line 897: yy, tm = tm, yy
					πF.SetLineno(897)
					if πE = πg.CheckLocal(πF, µtm, "tm"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µyy, "yy"); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple2(µtm, µyy).ToObject()
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp004}}}, πTemp001); πE != nil {
						continue
					}
					µyy = πTemp003
					µtm = πTemp004
					goto Label29
				Label29:
					if πTemp004, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp003 = πTemp004
					if πE = πg.CheckLocal(πF, µyy, "yy"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µyy, πTemp003); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, πTemp004, πg.NewStr(",").ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label30
					}
					goto Label31
					// line 898: if yy[-1] == ',':
					πF.SetLineno(898)
				Label30:
					// line 899: yy = yy[:-1]
					πF.SetLineno(899)
					if πTemp003, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πTemp003, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µyy, "yy"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µyy, πTemp001); πE != nil {
						continue
					}
					µyy = πTemp003
					goto Label31
				Label31:
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µyy, "yy"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µyy, πTemp003); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßisdigit, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(!πTemp002).ToObject()
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label32
					}
					goto Label33
					// line 900: if not yy[0].isdigit():
					πF.SetLineno(900)
				Label32:
					// line 901: yy, tz = tz, yy
					πF.SetLineno(901)
					if πE = πg.CheckLocal(πF, µtz, "tz"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µyy, "yy"); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple2(µtz, µyy).ToObject()
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp004}}}, πTemp001); πE != nil {
						continue
					}
					µyy = πTemp003
					µtz = πTemp004
					goto Label33
				Label33:
					if πTemp004, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp003 = πTemp004
					if πE = πg.CheckLocal(πF, µtm, "tm"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µtm, πTemp003); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, πTemp004, πg.NewStr(",").ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label34
					}
					goto Label35
					// line 902: if tm[-1] == ',':
					πF.SetLineno(902)
				Label34:
					// line 903: tm = tm[:-1]
					πF.SetLineno(903)
					if πTemp003, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πTemp003, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µtm, "tm"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µtm, πTemp001); πE != nil {
						continue
					}
					µtm = πTemp003
					goto Label35
				Label35:
					// line 904: tm = tm.split(':')
					πF.SetLineno(904)
					πTemp009 = πF.MakeArgs(1)
					πTemp009[0] = πg.NewStr(":").ToObject()
					if πE = πg.CheckLocal(πF, µtm, "tm"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µtm, ßsplit, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp009, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp009)
					µtm = πTemp003
					πTemp009 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µtm, "tm"); πE != nil {
						continue
					}
					πTemp009[0] = µtm
					if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp009, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp009)
					if πTemp001, πE = πg.Eq(πF, πTemp004, πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label36
					}
					πTemp009 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µtm, "tm"); πE != nil {
						continue
					}
					πTemp009[0] = µtm
					if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp009, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp009)
					if πTemp001, πE = πg.Eq(πF, πTemp004, πg.NewInt(3).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label37
					}
					goto Label38
					// line 905: if len(tm) == 2:
					πF.SetLineno(905)
				Label36:
					// line 906: [thh, tmm] = tm
					πF.SetLineno(906)
					if πE = πg.CheckLocal(πF, µtm, "tm"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp003}}}, µtm); πE != nil {
						continue
					}
					µthh = πTemp001
					µtmm = πTemp003
					// line 907: tss = '0'
					πF.SetLineno(907)
					µtss = ß0.ToObject()
					goto Label39
					// line 908: elif len(tm) == 3:
					πF.SetLineno(908)
				Label37:
					// line 909: [thh, tmm, tss] = tm
					πF.SetLineno(909)
					if πE = πg.CheckLocal(πF, µtm, "tm"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp004}}}, µtm); πE != nil {
						continue
					}
					µthh = πTemp001
					µtmm = πTemp003
					µtss = πTemp004
					goto Label39
				Label38:
					// line 911: return None
					πF.SetLineno(911)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πR = πTemp001
					continue
					goto Label39
				Label39:
					// line 912: try:
					πF.SetLineno(912)
					πF.PushCheckpoint(41)
					// line 913: yy = int(yy)
					πF.SetLineno(913)
					πTemp009 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µyy, "yy"); πE != nil {
						continue
					}
					πTemp009[0] = µyy
					if πTemp001, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp009, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp009)
					µyy = πTemp003
					// line 914: dd = int(dd)
					πF.SetLineno(914)
					πTemp009 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µdd, "dd"); πE != nil {
						continue
					}
					πTemp009[0] = µdd
					if πTemp001, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp009, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp009)
					µdd = πTemp003
					// line 915: thh = int(thh)
					πF.SetLineno(915)
					πTemp009 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µthh, "thh"); πE != nil {
						continue
					}
					πTemp009[0] = µthh
					if πTemp001, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp009, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp009)
					µthh = πTemp003
					// line 916: tmm = int(tmm)
					πF.SetLineno(916)
					πTemp009 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µtmm, "tmm"); πE != nil {
						continue
					}
					πTemp009[0] = µtmm
					if πTemp001, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp009, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp009)
					µtmm = πTemp003
					// line 917: tss = int(tss)
					πF.SetLineno(917)
					πTemp009 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µtss, "tss"); πE != nil {
						continue
					}
					πTemp009[0] = µtss
					if πTemp001, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp009, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp009)
					µtss = πTemp003
					πF.PopCheckpoint()
					goto Label40
				Label41:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp010, πTemp011 = πF.ExcInfo()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsInstance(πF, πTemp010.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label42
					}
					πE = πF.Raise(πTemp010.ToObject(), nil, πTemp011.ToObject())
					continue
					// line 918: except ValueError:
					πF.SetLineno(918)
				Label42:
					// line 919: return None
					πF.SetLineno(919)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πR = πTemp001
					continue
					πF.RestoreExc(nil, nil)
					goto Label40
				Label40:
					// line 920: tzoffset = None
					πF.SetLineno(920)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					µtzoffset = πTemp001
					// line 921: tz = tz.upper()
					πF.SetLineno(921)
					if πE = πg.CheckLocal(πF, µtz, "tz"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µtz, ßupper, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					µtz = πTemp003
					if πE = πg.CheckLocal(πF, µtz, "tz"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_timezones); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Contains(πF, πTemp003, µtz); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(πTemp002).ToObject()
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label43
					}
					goto Label44
					// line 922: if tz in _timezones:
					πF.SetLineno(922)
				Label43:
					// line 923: tzoffset = _timezones[tz]
					πF.SetLineno(923)
					if πE = πg.CheckLocal(πF, µtz, "tz"); πE != nil {
						continue
					}
					πTemp001 = µtz
					if πTemp004, πE = πg.ResolveGlobal(πF, ß_timezones); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, πTemp004, πTemp001); πE != nil {
						continue
					}
					µtzoffset = πTemp003
					goto Label45
				Label44:
					// line 925: try:
					πF.SetLineno(925)
					πF.PushCheckpoint(47)
					// line 926: tzoffset = int(tz)
					πF.SetLineno(926)
					πTemp009 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µtz, "tz"); πE != nil {
						continue
					}
					πTemp009[0] = µtz
					if πTemp001, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp009, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp009)
					µtzoffset = πTemp003
					πF.PopCheckpoint()
					goto Label46
				Label47:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp010, πTemp011 = πF.ExcInfo()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsInstance(πF, πTemp010.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label48
					}
					πE = πF.Raise(πTemp010.ToObject(), nil, πTemp011.ToObject())
					continue
					// line 927: except ValueError:
					πF.SetLineno(927)
				Label48:
					// line 928: pass
					πF.SetLineno(928)
					πF.RestoreExc(nil, nil)
					goto Label46
				Label46:
					goto Label45
				Label45:
					if πE = πg.CheckLocal(πF, µtzoffset, "tzoffset"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, µtzoffset); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label49
					}
					goto Label50
					// line 930: if tzoffset:
					πF.SetLineno(930)
				Label49:
					if πE = πg.CheckLocal(πF, µtzoffset, "tzoffset"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.LT(πF, µtzoffset, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label51
					}
					goto Label52
					// line 931: if tzoffset < 0:
					πF.SetLineno(931)
				Label51:
					// line 932: tzsign = -1
					πF.SetLineno(932)
					if πTemp001, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µtzsign = πTemp001
					// line 933: tzoffset = -tzoffset
					πF.SetLineno(933)
					if πE = πg.CheckLocal(πF, µtzoffset, "tzoffset"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Neg(πF, µtzoffset); πE != nil {
						continue
					}
					µtzoffset = πTemp001
					goto Label53
				Label52:
					// line 935: tzsign = 1
					πF.SetLineno(935)
					µtzsign = πg.NewInt(1).ToObject()
					goto Label53
				Label53:
					// line 936: tzoffset = tzsign * ( (tzoffset//100)*3600 + (tzoffset % 100)*60)
					πF.SetLineno(936)
					if πE = πg.CheckLocal(πF, µtzsign, "tzsign"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µtzoffset, "tzoffset"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.FloorDiv(πF, µtzoffset, πg.NewInt(100).ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Mul(πF, πTemp005, πg.NewInt(3600).ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µtzoffset, "tzoffset"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.Mod(πF, µtzoffset, πg.NewInt(100).ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Mul(πF, πTemp006, πg.NewInt(60).ToObject()); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Add(πF, πTemp004, πTemp005); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Mul(πF, µtzsign, πTemp003); πE != nil {
						continue
					}
					µtzoffset = πTemp001
					goto Label50
				Label50:
					// line 937: return (yy, mm, dd, thh, tmm, tss, 0, 1, 0, tzoffset)
					πF.SetLineno(937)
					πTemp009 = make([]*πg.Object, 10)
					if πE = πg.CheckLocal(πF, µyy, "yy"); πE != nil {
						continue
					}
					πTemp009[0] = µyy
					if πE = πg.CheckLocal(πF, µmm, "mm"); πE != nil {
						continue
					}
					πTemp009[1] = µmm
					if πE = πg.CheckLocal(πF, µdd, "dd"); πE != nil {
						continue
					}
					πTemp009[2] = µdd
					if πE = πg.CheckLocal(πF, µthh, "thh"); πE != nil {
						continue
					}
					πTemp009[3] = µthh
					if πE = πg.CheckLocal(πF, µtmm, "tmm"); πE != nil {
						continue
					}
					πTemp009[4] = µtmm
					if πE = πg.CheckLocal(πF, µtss, "tss"); πE != nil {
						continue
					}
					πTemp009[5] = µtss
					πTemp009[6] = πg.NewInt(0).ToObject()
					πTemp009[7] = πg.NewInt(1).ToObject()
					πTemp009[8] = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µtzoffset, "tzoffset"); πE != nil {
						continue
					}
					πTemp009[9] = µtzoffset
					πTemp001 = πg.NewTuple(πTemp009...).ToObject()
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
			if πE = πF.Globals().SetItem(πF, ßparsedate_tz.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 856: """Convert a date string to a time tuple.
			πF.SetLineno(856)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp010}, πg.NewStr("Convert a date string to a time tuple.\n\n    Accounts for military timezones.\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp011, πE = πg.ResolveGlobal(πF, ßparsedate_tz); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp011, ß__doc__, πTemp010); πE != nil {
				continue
			}
			// line 940: def parsedate(data):
			πF.SetLineno(940)
			πTemp007 = make([]πg.Param, 1)
			πTemp007[0] = πg.Param{Name: "data", Def: nil}
			πTemp010 = πg.NewFunction(πg.NewCode("parsedate", "/usr/lib/python2.7/rfc822.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µdata *πg.Object = πArgs[0]; _ = µdata
				var µt *πg.Object = πg.UnboundLocal; _ = µt
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
					// line 941: """Convert a time string to a time tuple."""
					πF.SetLineno(941)
					// line 942: t = parsedate_tz(data)
					πF.SetLineno(942)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
						continue
					}
					πTemp001[0] = µdata
					if πTemp002, πE = πg.ResolveGlobal(πF, ßparsedate_tz); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µt = πTemp003
					if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(µt == πTemp003).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label1
					}
					goto Label2
					// line 943: if t is None:
					πF.SetLineno(943)
				Label1:
					// line 944: return t
					πF.SetLineno(944)
					if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
						continue
					}
					πR = µt
					continue
					goto Label2
				Label2:
					// line 945: return t[:9]
					πF.SetLineno(945)
					if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.NewInt(9).ToObject(), πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µt, πTemp002); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßparsedate.ToObject(), πTemp010); πE != nil {
				continue
			}
			// line 941: """Convert a time string to a time tuple."""
			πF.SetLineno(941)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πg.NewStr("Convert a time string to a time tuple.").ToObject()); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßparsedate); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp012, ß__doc__, πTemp011); πE != nil {
				continue
			}
			// line 948: def mktime_tz(data):
			πF.SetLineno(948)
			πTemp007 = make([]πg.Param, 1)
			πTemp007[0] = πg.Param{Name: "data", Def: nil}
			πTemp011 = πg.NewFunction(πg.NewCode("mktime_tz", "/usr/lib/python2.7/rfc822.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µdata *πg.Object = πArgs[0]; _ = µdata
				var µt *πg.Object = πg.UnboundLocal; _ = µt
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
					// line 949: """Turn a 10-tuple as returned by parsedate_tz() into a UTC timestamp."""
					πF.SetLineno(949)
					πTemp002 = πg.NewInt(9).ToObject()
					if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µdata, πTemp002); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(πTemp003 == πTemp002).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label1
					}
					goto Label2
					// line 950: if data[9] is None:
					πF.SetLineno(950)
				Label1:
					// line 952: return time.mktime(data[:8] + (-1,))
					πF.SetLineno(952)
					πTemp005 = πF.MakeArgs(1)
					if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.NewInt(8).ToObject(), πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µdata, πTemp002); πE != nil {
						continue
					}
					if πTemp006, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple1(πTemp006).ToObject()
					if πTemp001, πE = πg.Add(πF, πTemp003, πTemp002); πE != nil {
						continue
					}
					πTemp005[0] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßtime); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßmktime, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πR = πTemp001
					continue
					goto Label3
				Label2:
					// line 954: t = time.mktime(data[:8] + (0,))
					πF.SetLineno(954)
					πTemp005 = πF.MakeArgs(1)
					if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.NewInt(8).ToObject(), πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µdata, πTemp002); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple1(πg.NewInt(0).ToObject()).ToObject()
					if πTemp001, πE = πg.Add(πF, πTemp003, πTemp002); πE != nil {
						continue
					}
					πTemp005[0] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßtime); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßmktime, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					µt = πTemp001
					// line 955: return t - data[9] - time.timezone
					πF.SetLineno(955)
					if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
						continue
					}
					πTemp003 = πg.NewInt(9).ToObject()
					if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, µdata, πTemp003); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Sub(πF, µt, πTemp006); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßtime); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp003, ßtimezone, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Sub(πF, πTemp002, πTemp006); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßmktime_tz.ToObject(), πTemp011); πE != nil {
				continue
			}
			// line 949: """Turn a 10-tuple as returned by parsedate_tz() into a UTC timestamp."""
			πF.SetLineno(949)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp012}, πg.NewStr("Turn a 10-tuple as returned by parsedate_tz() into a UTC timestamp.").ToObject()); πE != nil {
				continue
			}
			if πTemp013, πE = πg.ResolveGlobal(πF, ßmktime_tz); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp013, ß__doc__, πTemp012); πE != nil {
				continue
			}
			// line 957: def formatdate(timeval=None):
			πF.SetLineno(957)
			πTemp007 = make([]πg.Param, 1)
			if πTemp013, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp007[0] = πg.Param{Name: "timeval", Def: πTemp013}
			πTemp012 = πg.NewFunction(πg.NewCode("formatdate", "/usr/lib/python2.7/rfc822.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µtimeval *πg.Object = πArgs[0]; _ = µtimeval
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
					// line 958: """Returns time format preferred for Internet standards.
					πF.SetLineno(958)
					if πE = πg.CheckLocal(πF, µtimeval, "timeval"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µtimeval == πTemp002).ToObject()
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label1
					}
					goto Label2
					// line 967: if timeval is None:
					πF.SetLineno(967)
				Label1:
					// line 968: timeval = time.time()
					πF.SetLineno(968)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßtime); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßtime, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					µtimeval = πTemp001
					goto Label2
				Label2:
					// line 969: timeval = time.gmtime(timeval)
					πF.SetLineno(969)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µtimeval, "timeval"); πE != nil {
						continue
					}
					πTemp004[0] = µtimeval
					if πTemp001, πE = πg.ResolveGlobal(πF, ßtime); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßgmtime, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µtimeval = πTemp001
					// line 970: return "%s, %02d %s %04d %02d:%02d:%02d GMT" % (
					πF.SetLineno(970)
					πTemp004 = make([]*πg.Object, 7)
					πTemp006 = πg.NewInt(6).ToObject()
					if πE = πg.CheckLocal(πF, µtimeval, "timeval"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetItem(πF, µtimeval, πTemp006); πE != nil {
						continue
					}
					πTemp005 = πTemp007
					πTemp008 = make([]*πg.Object, 7)
					πTemp008[0] = ßMon.ToObject()
					πTemp008[1] = ßTue.ToObject()
					πTemp008[2] = ßWed.ToObject()
					πTemp008[3] = ßThu.ToObject()
					πTemp008[4] = ßFri.ToObject()
					πTemp008[5] = ßSat.ToObject()
					πTemp008[6] = ßSun.ToObject()
					πTemp007 = πg.NewTuple(πTemp008...).ToObject()
					if πTemp006, πE = πg.GetItem(πF, πTemp007, πTemp005); πE != nil {
						continue
					}
					πTemp004[0] = πTemp006
					πTemp005 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µtimeval, "timeval"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, µtimeval, πTemp005); πE != nil {
						continue
					}
					πTemp004[1] = πTemp006
					πTemp007 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µtimeval, "timeval"); πE != nil {
						continue
					}
					if πTemp009, πE = πg.GetItem(πF, µtimeval, πTemp007); πE != nil {
						continue
					}
					if πTemp006, πE = πg.Sub(πF, πTemp009, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp005 = πTemp006
					πTemp008 = make([]*πg.Object, 12)
					πTemp008[0] = ßJan.ToObject()
					πTemp008[1] = ßFeb.ToObject()
					πTemp008[2] = ßMar.ToObject()
					πTemp008[3] = ßApr.ToObject()
					πTemp008[4] = ßMay.ToObject()
					πTemp008[5] = ßJun.ToObject()
					πTemp008[6] = ßJul.ToObject()
					πTemp008[7] = ßAug.ToObject()
					πTemp008[8] = ßSep.ToObject()
					πTemp008[9] = ßOct.ToObject()
					πTemp008[10] = ßNov.ToObject()
					πTemp008[11] = ßDec.ToObject()
					πTemp007 = πg.NewTuple(πTemp008...).ToObject()
					if πTemp006, πE = πg.GetItem(πF, πTemp007, πTemp005); πE != nil {
						continue
					}
					πTemp004[2] = πTemp006
					πTemp005 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µtimeval, "timeval"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, µtimeval, πTemp005); πE != nil {
						continue
					}
					πTemp004[3] = πTemp006
					πTemp005 = πg.NewInt(3).ToObject()
					if πE = πg.CheckLocal(πF, µtimeval, "timeval"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, µtimeval, πTemp005); πE != nil {
						continue
					}
					πTemp004[4] = πTemp006
					πTemp005 = πg.NewInt(4).ToObject()
					if πE = πg.CheckLocal(πF, µtimeval, "timeval"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, µtimeval, πTemp005); πE != nil {
						continue
					}
					πTemp004[5] = πTemp006
					πTemp005 = πg.NewInt(5).ToObject()
					if πE = πg.CheckLocal(πF, µtimeval, "timeval"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, µtimeval, πTemp005); πE != nil {
						continue
					}
					πTemp004[6] = πTemp006
					πTemp002 = πg.NewTuple(πTemp004...).ToObject()
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("%s, %02d %s %04d %02d:%02d:%02d GMT").ToObject(), πTemp002); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßformatdate.ToObject(), πTemp012); πE != nil {
				continue
			}
			// line 958: """Returns time format preferred for Internet standards.
			πF.SetLineno(958)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp013}, πg.NewStr("Returns time format preferred for Internet standards.\n\n    Sun, 06 Nov 1994 08:49:37 GMT  ; RFC 822, updated by RFC 1123\n\n    According to RFC 1123, day and month names must always be in\n    English.  If not for that, this code could use strftime().  It\n    can't because strftime() honors the locale and could generate\n    non-English names.\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp014, πE = πg.ResolveGlobal(πF, ßformatdate); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp014, ß__doc__, πTemp013); πE != nil {
				continue
			}
			if πTemp014, πE = πg.ResolveGlobal(πF, ß__name__); πE != nil {
				continue
			}
			if πTemp013, πE = πg.Eq(πF, πTemp014, ß__main__.ToObject()); πE != nil {
				continue
			}
			if πTemp015, πE = πg.IsTrue(πF, πTemp013); πE != nil {
				continue
			}
			if πTemp015 {
				goto Label1
			}
			goto Label2
			// line 982: if __name__ == '__main__':
			πF.SetLineno(982)
		Label1:
			// line 983: import sys, os
			πF.SetLineno(983)
			if πTemp002, πE = πg.ImportModule(πF, "sys"); πE != nil {
				continue
			}
			πTemp013 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßsys.ToObject(), πTemp013); πE != nil {
				continue
			}
			if πTemp002, πE = πg.ImportModule(πF, "os"); πE != nil {
				continue
			}
			πTemp013 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßos.ToObject(), πTemp013); πE != nil {
				continue
			}
			// line 984: file = os.path.join(os.environ['HOME'], 'Mail/inbox/1')
			πF.SetLineno(984)
			πTemp002 = πF.MakeArgs(2)
			πTemp013 = ßHOME.ToObject()
			if πTemp016, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
				continue
			}
			if πTemp017, πE = πg.GetAttr(πF, πTemp016, ßenviron, nil); πE != nil {
				continue
			}
			if πTemp014, πE = πg.GetItem(πF, πTemp017, πTemp013); πE != nil {
				continue
			}
			πTemp002[0] = πTemp014
			πTemp002[1] = πg.NewStr("Mail/inbox/1").ToObject()
			if πTemp013, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
				continue
			}
			if πTemp014, πE = πg.GetAttr(πF, πTemp013, ßpath, nil); πE != nil {
				continue
			}
			if πTemp013, πE = πg.GetAttr(πF, πTemp014, ßjoin, nil); πE != nil {
				continue
			}
			if πTemp014, πE = πTemp013.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πF.Globals().SetItem(πF, ßfile.ToObject(), πTemp014); πE != nil {
				continue
			}
			if πTemp013, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
				continue
			}
			if πTemp016, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
				continue
			}
			if πTemp017, πE = πg.GetAttr(πF, πTemp016, ßargv, nil); πE != nil {
				continue
			}
			if πTemp014, πE = πg.GetItem(πF, πTemp017, πTemp013); πE != nil {
				continue
			}
			if πTemp015, πE = πg.IsTrue(πF, πTemp014); πE != nil {
				continue
			}
			if πTemp015 {
				goto Label3
			}
			goto Label4
			// line 985: if sys.argv[1:]: file = sys.argv[1]
			πF.SetLineno(985)
		Label3:
			// line 985: if sys.argv[1:]: file = sys.argv[1]
			πF.SetLineno(985)
			πTemp013 = πg.NewInt(1).ToObject()
			if πTemp016, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
				continue
			}
			if πTemp017, πE = πg.GetAttr(πF, πTemp016, ßargv, nil); πE != nil {
				continue
			}
			if πTemp014, πE = πg.GetItem(πF, πTemp017, πTemp013); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßfile.ToObject(), πTemp014); πE != nil {
				continue
			}
			goto Label4
		Label4:
			// line 986: f = open(file, 'r')
			πF.SetLineno(986)
			πTemp002 = πF.MakeArgs(2)
			if πTemp013, πE = πg.ResolveGlobal(πF, ßfile); πE != nil {
				continue
			}
			πTemp002[0] = πTemp013
			πTemp002[1] = ßr.ToObject()
			if πTemp013, πE = πg.ResolveGlobal(πF, ßopen); πE != nil {
				continue
			}
			if πTemp014, πE = πTemp013.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πF.Globals().SetItem(πF, ßf.ToObject(), πTemp014); πE != nil {
				continue
			}
			// line 987: m = Message(f)
			πF.SetLineno(987)
			πTemp002 = πF.MakeArgs(1)
			if πTemp013, πE = πg.ResolveGlobal(πF, ßf); πE != nil {
				continue
			}
			πTemp002[0] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßMessage); πE != nil {
				continue
			}
			if πTemp014, πE = πTemp013.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πF.Globals().SetItem(πF, ßm.ToObject(), πTemp014); πE != nil {
				continue
			}
			// line 988: print 'From:', m.getaddr('from')
			πF.SetLineno(988)
			πTemp002 = make([]*πg.Object, 2)
			πTemp002[0] = πg.NewStr("From:").ToObject()
			πTemp018 = πF.MakeArgs(1)
			πTemp018[0] = ßfrom.ToObject()
			if πTemp013, πE = πg.ResolveGlobal(πF, ßm); πE != nil {
				continue
			}
			if πTemp014, πE = πg.GetAttr(πF, πTemp013, ßgetaddr, nil); πE != nil {
				continue
			}
			if πTemp013, πE = πTemp014.Call(πF, πTemp018, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp018)
			πTemp002[1] = πTemp013
			if πE = πg.Print(πF, πTemp002, true); πE != nil {
				continue
			}
			// line 989: print 'To:', m.getaddrlist('to')
			πF.SetLineno(989)
			πTemp002 = make([]*πg.Object, 2)
			πTemp002[0] = πg.NewStr("To:").ToObject()
			πTemp018 = πF.MakeArgs(1)
			πTemp018[0] = ßto.ToObject()
			if πTemp013, πE = πg.ResolveGlobal(πF, ßm); πE != nil {
				continue
			}
			if πTemp014, πE = πg.GetAttr(πF, πTemp013, ßgetaddrlist, nil); πE != nil {
				continue
			}
			if πTemp013, πE = πTemp014.Call(πF, πTemp018, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp018)
			πTemp002[1] = πTemp013
			if πE = πg.Print(πF, πTemp002, true); πE != nil {
				continue
			}
			// line 990: print 'Subject:', m.getheader('subject')
			πF.SetLineno(990)
			πTemp002 = make([]*πg.Object, 2)
			πTemp002[0] = πg.NewStr("Subject:").ToObject()
			πTemp018 = πF.MakeArgs(1)
			πTemp018[0] = ßsubject.ToObject()
			if πTemp013, πE = πg.ResolveGlobal(πF, ßm); πE != nil {
				continue
			}
			if πTemp014, πE = πg.GetAttr(πF, πTemp013, ßgetheader, nil); πE != nil {
				continue
			}
			if πTemp013, πE = πTemp014.Call(πF, πTemp018, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp018)
			πTemp002[1] = πTemp013
			if πE = πg.Print(πF, πTemp002, true); πE != nil {
				continue
			}
			// line 991: print 'Date:', m.getheader('date')
			πF.SetLineno(991)
			πTemp002 = make([]*πg.Object, 2)
			πTemp002[0] = πg.NewStr("Date:").ToObject()
			πTemp018 = πF.MakeArgs(1)
			πTemp018[0] = ßdate.ToObject()
			if πTemp013, πE = πg.ResolveGlobal(πF, ßm); πE != nil {
				continue
			}
			if πTemp014, πE = πg.GetAttr(πF, πTemp013, ßgetheader, nil); πE != nil {
				continue
			}
			if πTemp013, πE = πTemp014.Call(πF, πTemp018, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp018)
			πTemp002[1] = πTemp013
			if πE = πg.Print(πF, πTemp002, true); πE != nil {
				continue
			}
			// line 992: date = m.getdate_tz('date')
			πF.SetLineno(992)
			πTemp002 = πF.MakeArgs(1)
			πTemp002[0] = ßdate.ToObject()
			if πTemp013, πE = πg.ResolveGlobal(πF, ßm); πE != nil {
				continue
			}
			if πTemp014, πE = πg.GetAttr(πF, πTemp013, ßgetdate_tz, nil); πE != nil {
				continue
			}
			if πTemp013, πE = πTemp014.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πF.Globals().SetItem(πF, ßdate.ToObject(), πTemp013); πE != nil {
				continue
			}
			// line 993: tz = date[-1]
			πF.SetLineno(993)
			if πTemp014, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp013 = πTemp014
			if πTemp016, πE = πg.ResolveGlobal(πF, ßdate); πE != nil {
				continue
			}
			if πTemp014, πE = πg.GetItem(πF, πTemp016, πTemp013); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßtz.ToObject(), πTemp014); πE != nil {
				continue
			}
			// line 994: date = time.localtime(mktime_tz(date))
			πF.SetLineno(994)
			πTemp002 = πF.MakeArgs(1)
			πTemp018 = πF.MakeArgs(1)
			if πTemp013, πE = πg.ResolveGlobal(πF, ßdate); πE != nil {
				continue
			}
			πTemp018[0] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßmktime_tz); πE != nil {
				continue
			}
			if πTemp014, πE = πTemp013.Call(πF, πTemp018, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp018)
			πTemp002[0] = πTemp014
			if πTemp013, πE = πg.ResolveGlobal(πF, ßtime); πE != nil {
				continue
			}
			if πTemp014, πE = πg.GetAttr(πF, πTemp013, ßlocaltime, nil); πE != nil {
				continue
			}
			if πTemp013, πE = πTemp014.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πF.Globals().SetItem(πF, ßdate.ToObject(), πTemp013); πE != nil {
				continue
			}
			if πTemp013, πE = πg.ResolveGlobal(πF, ßdate); πE != nil {
				continue
			}
			if πTemp015, πE = πg.IsTrue(πF, πTemp013); πE != nil {
				continue
			}
			if πTemp015 {
				goto Label5
			}
			goto Label6
			// line 995: if date:
			πF.SetLineno(995)
		Label5:
			// line 996: print 'ParsedDate:', time.asctime(date),
			πF.SetLineno(996)
			πTemp002 = make([]*πg.Object, 2)
			πTemp002[0] = πg.NewStr("ParsedDate:").ToObject()
			πTemp018 = πF.MakeArgs(1)
			if πTemp013, πE = πg.ResolveGlobal(πF, ßdate); πE != nil {
				continue
			}
			πTemp018[0] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßtime); πE != nil {
				continue
			}
			if πTemp014, πE = πg.GetAttr(πF, πTemp013, ßasctime, nil); πE != nil {
				continue
			}
			if πTemp013, πE = πTemp014.Call(πF, πTemp018, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp018)
			πTemp002[1] = πTemp013
			if πE = πg.Print(πF, πTemp002, false); πE != nil {
				continue
			}
			// line 997: hhmmss = tz
			πF.SetLineno(997)
			if πTemp013, πE = πg.ResolveGlobal(πF, ßtz); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßhhmmss.ToObject(), πTemp013); πE != nil {
				continue
			}
			// line 998: hhmm, ss = divmod(hhmmss, 60)
			πF.SetLineno(998)
			πTemp002 = πF.MakeArgs(2)
			if πTemp013, πE = πg.ResolveGlobal(πF, ßhhmmss); πE != nil {
				continue
			}
			πTemp002[0] = πTemp013
			πTemp002[1] = πg.NewInt(60).ToObject()
			if πTemp013, πE = πg.ResolveGlobal(πF, ßdivmod); πE != nil {
				continue
			}
			if πTemp014, πE = πTemp013.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp013}, πg.TieTarget{Target: &πTemp016}}}, πTemp014); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßhhmm.ToObject(), πTemp013); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßss.ToObject(), πTemp016); πE != nil {
				continue
			}
			// line 999: hh, mm = divmod(hhmm, 60)
			πF.SetLineno(999)
			πTemp002 = πF.MakeArgs(2)
			if πTemp013, πE = πg.ResolveGlobal(πF, ßhhmm); πE != nil {
				continue
			}
			πTemp002[0] = πTemp013
			πTemp002[1] = πg.NewInt(60).ToObject()
			if πTemp013, πE = πg.ResolveGlobal(πF, ßdivmod); πE != nil {
				continue
			}
			if πTemp014, πE = πTemp013.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp013}, πg.TieTarget{Target: &πTemp016}}}, πTemp014); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßhh.ToObject(), πTemp013); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßmm.ToObject(), πTemp016); πE != nil {
				continue
			}
			// line 1000: print "%+03d%02d" % (hh, mm),
			πF.SetLineno(1000)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp016, πE = πg.ResolveGlobal(πF, ßhh); πE != nil {
				continue
			}
			if πTemp017, πE = πg.ResolveGlobal(πF, ßmm); πE != nil {
				continue
			}
			πTemp014 = πg.NewTuple2(πTemp016, πTemp017).ToObject()
			if πTemp013, πE = πg.Mod(πF, πg.NewStr("%+03d%02d").ToObject(), πTemp014); πE != nil {
				continue
			}
			πTemp002[0] = πTemp013
			if πE = πg.Print(πF, πTemp002, false); πE != nil {
				continue
			}
			if πTemp013, πE = πg.ResolveGlobal(πF, ßss); πE != nil {
				continue
			}
			if πTemp015, πE = πg.IsTrue(πF, πTemp013); πE != nil {
				continue
			}
			if πTemp015 {
				goto Label8
			}
			goto Label9
			// line 1001: if ss: print ".%02d" % ss,
			πF.SetLineno(1001)
		Label8:
			// line 1001: if ss: print ".%02d" % ss,
			πF.SetLineno(1001)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp014, πE = πg.ResolveGlobal(πF, ßss); πE != nil {
				continue
			}
			if πTemp013, πE = πg.Mod(πF, πg.NewStr(".%02d").ToObject(), πTemp014); πE != nil {
				continue
			}
			πTemp002[0] = πTemp013
			if πE = πg.Print(πF, πTemp002, false); πE != nil {
				continue
			}
			goto Label9
		Label9:
			// line 1002: print
			πF.SetLineno(1002)
			πTemp002 = make([]*πg.Object, 0)
			if πE = πg.Print(πF, πTemp002, true); πE != nil {
				continue
			}
			goto Label7
		Label6:
			// line 1004: print 'ParsedDate:', None
			πF.SetLineno(1004)
			πTemp002 = make([]*πg.Object, 2)
			πTemp002[0] = πg.NewStr("ParsedDate:").ToObject()
			if πTemp013, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp002[1] = πTemp013
			if πE = πg.Print(πF, πTemp002, true); πE != nil {
				continue
			}
			goto Label7
		Label7:
			// line 1005: m.rewindbody()
			πF.SetLineno(1005)
			if πTemp013, πE = πg.ResolveGlobal(πF, ßm); πE != nil {
				continue
			}
			if πTemp014, πE = πg.GetAttr(πF, πTemp013, ßrewindbody, nil); πE != nil {
				continue
			}
			if πTemp013, πE = πTemp014.Call(πF, nil, nil); πE != nil {
				continue
			}
			// line 1006: n = 0
			πF.SetLineno(1006)
			if πE = πF.Globals().SetItem(πF, ßn.ToObject(), πg.NewInt(0).ToObject()); πE != nil {
				continue
			}
			// line 1007: while f.readline():
			πF.SetLineno(1007)
			πF.PushCheckpoint(11)
			πTemp015 = false
		Label10:
			if πE != nil || πR != nil {
				continue
			}
			if πTemp015 {
				πF.PopCheckpoint()
				goto Label12
			}
			if πTemp013, πE = πg.ResolveGlobal(πF, ßf); πE != nil {
				continue
			}
			if πTemp014, πE = πg.GetAttr(πF, πTemp013, ßreadline, nil); πE != nil {
				continue
			}
			if πTemp013, πE = πTemp014.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πTemp019, πE = πg.IsTrue(πF, πTemp013); πE != nil {
				continue
			}
			if πE != nil || !πTemp019 {
				continue
			}
			πF.PushCheckpoint(10)            
			// line 1008: n += 1
			πF.SetLineno(1008)
			if πTemp013, πE = πg.ResolveGlobal(πF, ßn); πE != nil {
				continue
			}
			if πTemp014, πE = πg.IAdd(πF, πTemp013, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßn.ToObject(), πTemp014); πE != nil {
				continue
			}
			continue
		Label11:
			if πE != nil || πR != nil {
				continue
			}
		Label12:
			// line 1009: print 'Lines:', n
			πF.SetLineno(1009)
			πTemp002 = make([]*πg.Object, 2)
			πTemp002[0] = πg.NewStr("Lines:").ToObject()
			if πTemp013, πE = πg.ResolveGlobal(πF, ßn); πE != nil {
				continue
			}
			πTemp002[1] = πTemp013
			if πE = πg.Print(πF, πTemp002, true); πE != nil {
				continue
			}
			// line 1010: print '-'*70
			πF.SetLineno(1010)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp013, πE = πg.Mul(πF, πg.NewStr("-").ToObject(), πg.NewInt(70).ToObject()); πE != nil {
				continue
			}
			πTemp002[0] = πTemp013
			if πE = πg.Print(πF, πTemp002, true); πE != nil {
				continue
			}
			// line 1011: print 'len =', len(m)
			πF.SetLineno(1011)
			πTemp002 = make([]*πg.Object, 2)
			πTemp002[0] = πg.NewStr("len =").ToObject()
			πTemp018 = πF.MakeArgs(1)
			if πTemp013, πE = πg.ResolveGlobal(πF, ßm); πE != nil {
				continue
			}
			πTemp018[0] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
				continue
			}
			if πTemp014, πE = πTemp013.Call(πF, πTemp018, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp018)
			πTemp002[1] = πTemp014
			if πE = πg.Print(πF, πTemp002, true); πE != nil {
				continue
			}
			if πTemp014, πE = πg.ResolveGlobal(πF, ßm); πE != nil {
				continue
			}
			if πTemp015, πE = πg.Contains(πF, πTemp014, ßDate.ToObject()); πE != nil {
				continue
			}
			πTemp013 = πg.GetBool(πTemp015).ToObject()
			if πTemp015, πE = πg.IsTrue(πF, πTemp013); πE != nil {
				continue
			}
			if πTemp015 {
				goto Label13
			}
			goto Label14
			// line 1012: if 'Date' in m: print 'Date =', m['Date']
			πF.SetLineno(1012)
		Label13:
			// line 1012: if 'Date' in m: print 'Date =', m['Date']
			πF.SetLineno(1012)
			πTemp002 = make([]*πg.Object, 2)
			πTemp002[0] = πg.NewStr("Date =").ToObject()
			πTemp013 = ßDate.ToObject()
			if πTemp016, πE = πg.ResolveGlobal(πF, ßm); πE != nil {
				continue
			}
			if πTemp014, πE = πg.GetItem(πF, πTemp016, πTemp013); πE != nil {
				continue
			}
			πTemp002[1] = πTemp014
			if πE = πg.Print(πF, πTemp002, true); πE != nil {
				continue
			}
			goto Label14
		Label14:
			if πTemp014, πE = πg.ResolveGlobal(πF, ßm); πE != nil {
				continue
			}
			if πTemp015, πE = πg.Contains(πF, πTemp014, πg.NewStr("X-Nonsense").ToObject()); πE != nil {
				continue
			}
			πTemp013 = πg.GetBool(πTemp015).ToObject()
			if πTemp015, πE = πg.IsTrue(πF, πTemp013); πE != nil {
				continue
			}
			if πTemp015 {
				goto Label15
			}
			goto Label16
			// line 1013: if 'X-Nonsense' in m: pass
			πF.SetLineno(1013)
		Label15:
			// line 1013: if 'X-Nonsense' in m: pass
			πF.SetLineno(1013)
			goto Label16
		Label16:
			// line 1014: print 'keys =', m.keys()
			πF.SetLineno(1014)
			πTemp002 = make([]*πg.Object, 2)
			πTemp002[0] = πg.NewStr("keys =").ToObject()
			if πTemp013, πE = πg.ResolveGlobal(πF, ßm); πE != nil {
				continue
			}
			if πTemp014, πE = πg.GetAttr(πF, πTemp013, ßkeys, nil); πE != nil {
				continue
			}
			if πTemp013, πE = πTemp014.Call(πF, nil, nil); πE != nil {
				continue
			}
			πTemp002[1] = πTemp013
			if πE = πg.Print(πF, πTemp002, true); πE != nil {
				continue
			}
			// line 1015: print 'values =', m.values()
			πF.SetLineno(1015)
			πTemp002 = make([]*πg.Object, 2)
			πTemp002[0] = πg.NewStr("values =").ToObject()
			if πTemp013, πE = πg.ResolveGlobal(πF, ßm); πE != nil {
				continue
			}
			if πTemp014, πE = πg.GetAttr(πF, πTemp013, ßvalues, nil); πE != nil {
				continue
			}
			if πTemp013, πE = πTemp014.Call(πF, nil, nil); πE != nil {
				continue
			}
			πTemp002[1] = πTemp013
			if πE = πg.Print(πF, πTemp002, true); πE != nil {
				continue
			}
			// line 1016: print 'items =', m.items()
			πF.SetLineno(1016)
			πTemp002 = make([]*πg.Object, 2)
			πTemp002[0] = πg.NewStr("items =").ToObject()
			if πTemp013, πE = πg.ResolveGlobal(πF, ßm); πE != nil {
				continue
			}
			if πTemp014, πE = πg.GetAttr(πF, πTemp013, ßitems, nil); πE != nil {
				continue
			}
			if πTemp013, πE = πTemp014.Call(πF, nil, nil); πE != nil {
				continue
			}
			πTemp002[1] = πTemp013
			if πE = πg.Print(πF, πTemp002, true); πE != nil {
				continue
			}
			goto Label2
		Label2:
		}
		return nil, πE
	})
	πg.RegisterModule("rfc822", Code)
}

