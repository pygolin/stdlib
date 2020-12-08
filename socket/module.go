package socket
import (
	πg "github.com/pygolin/runtime"
	// _ "github.com/pygolin/stdlib/cStringIO"
	// _ "github.com/pygolin/stdlib/warnings"
	// _ "github.com/pygolin/stdlib/errno"
	// _ "github.com/pygolin/stdlib/functools"
	// _ "github.com/pygolin/stdlib/_socket"
	// _ "github.com/pygolin/stdlib/sys"
	// _ "github.com/pygolin/stdlib/os"
	// _ "github.com/pygolin/stdlib/types"
	// _ "github.com/pygolin/stdlib/StringIO"
)
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/socket.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ß := πg.InternStr("")
		ßAF_INET := πg.InternStr("AF_INET")
		ßEBADF := πg.InternStr("EBADF")
		ßEINTR := πg.InternStr("EINTR")
		ßFalse := πg.InternStr("False")
		ßImportError := πg.InternStr("ImportError")
		ßMethodType := πg.InternStr("MethodType")
		ßNone := πg.InternStr("None")
		ßSOCK_STREAM := πg.InternStr("SOCK_STREAM")
		ßSocketType := πg.InternStr("SocketType")
		ßStopIteration := πg.InternStr("StopIteration")
		ßStringIO := πg.InternStr("StringIO")
		ßTrue := πg.InternStr("True")
		ß_ := πg.InternStr("_")
		ß_GLOBAL_DEFAULT_TIMEOUT := πg.InternStr("_GLOBAL_DEFAULT_TIMEOUT")
		ß__all__ := πg.InternStr("__all__")
		ß__del__ := πg.InternStr("__del__")
		ß__dict__ := πg.InternStr("__dict__")
		ß__doc__ := πg.InternStr("__doc__")
		ß__getattr__ := πg.InternStr("__getattr__")
		ß__init__ := πg.InternStr("__init__")
		ß__iter__ := πg.InternStr("__iter__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß__slots__ := πg.InternStr("__slots__")
		ß__weakref__ := πg.InternStr("__weakref__")
		ß_close := πg.InternStr("_close")
		ß_closedsocket := πg.InternStr("_closedsocket")
		ß_delegate_methods := πg.InternStr("_delegate_methods")
		ß_dummy := πg.InternStr("_dummy")
		ß_fileobject := πg.InternStr("_fileobject")
		ß_getclosed := πg.InternStr("_getclosed")
		ß_k := πg.InternStr("_k")
		ß_m := πg.InternStr("_m")
		ß_rbuf := πg.InternStr("_rbuf")
		ß_rbufsize := πg.InternStr("_rbufsize")
		ß_realsocket := πg.InternStr("_realsocket")
		ß_sock := πg.InternStr("_sock")
		ß_socket := πg.InternStr("_socket")
		ß_socketmethods := πg.InternStr("_socketmethods")
		ß_socketobject := πg.InternStr("_socketobject")
		ß_v := πg.InternStr("_v")
		ß_wbuf := πg.InternStr("_wbuf")
		ß_wbuf_len := πg.InternStr("_wbuf_len")
		ß_wbufsize := πg.InternStr("_wbufsize")
		ßaccept := πg.InternStr("accept")
		ßappend := πg.InternStr("append")
		ßargs := πg.InternStr("args")
		ßbind := πg.InternStr("bind")
		ßbufsize := πg.InternStr("bufsize")
		ßclose := πg.InternStr("close")
		ßclosed := πg.InternStr("closed")
		ßconnect := πg.InternStr("connect")
		ßconnect_ex := πg.InternStr("connect_ex")
		ßcreate_connection := πg.InternStr("create_connection")
		ßdefault_bufsize := πg.InternStr("default_bufsize")
		ßdir := πg.InternStr("dir")
		ßdup := πg.InternStr("dup")
		ßendswith := πg.InternStr("endswith")
		ßerrno := πg.InternStr("errno")
		ßerror := πg.InternStr("error")
		ßextend := πg.InternStr("extend")
		ßfamily := πg.InternStr("family")
		ßfileno := πg.InternStr("fileno")
		ßfilter := πg.InternStr("filter")
		ßfind := πg.InternStr("find")
		ßflush := πg.InternStr("flush")
		ßgetaddrinfo := πg.InternStr("getaddrinfo")
		ßgetattr := πg.InternStr("getattr")
		ßgetfqdn := πg.InternStr("getfqdn")
		ßgethostbyaddr := πg.InternStr("gethostbyaddr")
		ßgethostname := πg.InternStr("gethostname")
		ßgetpeername := πg.InternStr("getpeername")
		ßgetsockname := πg.InternStr("getsockname")
		ßgetsockopt := πg.InternStr("getsockopt")
		ßgettimeout := πg.InternStr("gettimeout")
		ßgetvalue := πg.InternStr("getvalue")
		ßglobals := πg.InternStr("globals")
		ßinsert := πg.InternStr("insert")
		ßiteritems := πg.InternStr("iteritems")
		ßjoin := πg.InternStr("join")
		ßlen := πg.InternStr("len")
		ßlist := πg.InternStr("list")
		ßlisten := πg.InternStr("listen")
		ßm := πg.InternStr("m")
		ßmakefile := πg.InternStr("makefile")
		ßmap := πg.InternStr("map")
		ßmax := πg.InternStr("max")
		ßmeth := πg.InternStr("meth")
		ßmode := πg.InternStr("mode")
		ßname := πg.InternStr("name")
		ßnext := πg.InternStr("next")
		ßobject := πg.InternStr("object")
		ßos := πg.InternStr("os")
		ßp := πg.InternStr("p")
		ßpartial := πg.InternStr("partial")
		ßproperty := πg.InternStr("property")
		ßproto := πg.InternStr("proto")
		ßr := πg.InternStr("r")
		ßrb := πg.InternStr("rb")
		ßread := πg.InternStr("read")
		ßreadline := πg.InternStr("readline")
		ßreadlines := πg.InternStr("readlines")
		ßrecv := πg.InternStr("recv")
		ßrecv_into := πg.InternStr("recv_into")
		ßrecvfrom := πg.InternStr("recvfrom")
		ßrecvfrom_into := πg.InternStr("recvfrom_into")
		ßseek := πg.InternStr("seek")
		ßsend := πg.InternStr("send")
		ßsendall := πg.InternStr("sendall")
		ßsendto := πg.InternStr("sendto")
		ßsetattr := πg.InternStr("setattr")
		ßsetblocking := πg.InternStr("setblocking")
		ßsetsockopt := πg.InternStr("setsockopt")
		ßsettimeout := πg.InternStr("settimeout")
		ßshutdown := πg.InternStr("shutdown")
		ßsocket := πg.InternStr("socket")
		ßsoftspace := πg.InternStr("softspace")
		ßstartswith := πg.InternStr("startswith")
		ßstr := πg.InternStr("str")
		ßstrip := πg.InternStr("strip")
		ßsum := πg.InternStr("sum")
		ßsys := πg.InternStr("sys")
		ßtell := πg.InternStr("tell")
		ßtype := πg.InternStr("type")
		ßwarnings := πg.InternStr("warnings")
		ßwrite := πg.InternStr("write")
		ßwritelines := πg.InternStr("writelines")
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
		var πTemp010 *πg.BaseException
		_ = πTemp010
		var πTemp011 *πg.Traceback
		_ = πTemp011
		var πTemp012 []*πg.Object
		_ = πTemp012
		var πTemp013 []πg.Param
		_ = πTemp013
		var πTemp014 *πg.Dict
		_ = πTemp014
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			case 1: goto Label1
			case 2: goto Label2
			case 7: goto Label7
			case 10: goto Label10
			case 12: goto Label12
			case 13: goto Label13
			default: panic("unexpected function state")
			}
			// line 4: """\
			πF.SetLineno(4)
			// line 4: """\
			πF.SetLineno(4)
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("This module provides socket operations and some related functions.\nOn Unix, it supports IP (Internet Protocol) and Unix domain sockets.\nOn other systems, it only supports IP. Functions specific for a\nsocket are available as methods of the socket object.\n\nFunctions:\n\nsocket() -- create a new socket object\nsocketpair() -- create a pair of new socket objects [*]\nfromfd() -- create a socket object from an open file descriptor [*]\ngethostname() -- return the current hostname\ngethostbyname() -- map a hostname to its IP number\ngethostbyaddr() -- map an IP number or hostname to DNS info\ngetservbyname() -- map a service name and a protocol name to a port number\ngetprotobyname() -- map a protocol name (e.g. 'tcp') to a number\nntohs(), ntohl() -- convert 16, 32 bit int from network to host byte order\nhtons(), htonl() -- convert 16, 32 bit int from host to network byte order\ninet_aton() -- convert IP addr string (123.45.67.89) to 32-bit packed format\ninet_ntoa() -- convert 32-bit packed format IP to string (123.45.67.89)\nssl() -- secure socket layer support (only available if configured)\nsocket.getdefaulttimeout() -- get the default timeout value\nsocket.setdefaulttimeout() -- set the default timeout value\ncreate_connection() -- connects to an address, with an optional timeout and\n                       optional source address.\n\n [*] not available on all platforms!\n\nSpecial objects:\n\nSocketType -- type object for socket objects\nerror -- exception raised for I/O errors\nhas_ipv6 -- boolean value indicating if IPv6 is supported\n\nInteger constants:\n\nAF_INET, AF_UNIX -- socket domains (first argument to socket() call)\nSOCK_STREAM, SOCK_DGRAM, SOCK_RAW -- socket types (second argument)\n\nMany other constants may be defined; these may be used in calls to\nthe setsockopt() and getsockopt() methods.\n").ToObject()); πE != nil {
				continue
			}
			// line 47: import _socket
			πF.SetLineno(47)
			if πTemp002, πE = πg.ImportModule(πF, "_socket"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ß_socket.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp003, πE = πg.ResolveGlobal(πF, ß_socket); πE != nil {
				continue
			}
			if πTemp004, πE = πg.GetAttr(πF, πTemp003, ß__dict__, nil); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßiteritems, nil); πE != nil {
				continue
			}
			if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
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
				if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
					continue
				}
				if πE = πF.Globals().SetItem(πF, ß_k.ToObject(), πTemp004); πE != nil {
					continue
				}
				if πE = πF.Globals().SetItem(πF, ß_v.ToObject(), πTemp007); πE != nil {
					continue
				}
			}
			if πE != nil || !πTemp006 {
				continue
			}
			πF.PushCheckpoint(1)            
			πTemp002 = πF.MakeArgs(1)
			πTemp002[0] = ß_.ToObject()
			if πTemp004, πE = πg.ResolveGlobal(πF, ß_k); πE != nil {
				continue
			}
			if πTemp007, πE = πg.GetAttr(πF, πTemp004, ßstartswith, nil); πE != nil {
				continue
			}
			if πTemp004, πE = πTemp007.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
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
			// line 50: if not _k.startswith('_'):
			πF.SetLineno(50)
		Label4:
			// line 51: globals()[_k] = _v
			πF.SetLineno(51)
			if πTemp003, πE = πg.ResolveGlobal(πF, ß_v); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp003); πE != nil {
				continue
			}
			if πTemp007, πE = πg.ResolveGlobal(πF, ßglobals); πE != nil {
				continue
			}
			if πTemp008, πE = πTemp007.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πTemp009, πE = πg.ResolveGlobal(πF, ß_k); πE != nil {
				continue
			}
			πTemp007 = πTemp009
			if πE = πg.SetItem(πF, πTemp008, πTemp007, πTemp004); πE != nil {
				continue
			}
			goto Label5
		Label5:
			continue
		Label2:
			if πE != nil || πR != nil {
				continue
			}
		Label3:
			// line 52: from functools import partial
			πF.SetLineno(52)
			if πTemp002, πE = πg.ImportModule(πF, "functools"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßpartial); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßpartial.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 53: from types import MethodType
			πF.SetLineno(53)
			if πTemp002, πE = πg.ImportModule(πF, "types"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßMethodType); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßMethodType.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 89: import os, sys, warnings
			πF.SetLineno(89)
			if πTemp002, πE = πg.ImportModule(πF, "os"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßos.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp002, πE = πg.ImportModule(πF, "sys"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßsys.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp002, πE = πg.ImportModule(πF, "warnings"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßwarnings.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 91: try:
			πF.SetLineno(91)
			πF.PushCheckpoint(7)
			// line 92: from cStringIO import StringIO
			πF.SetLineno(92)
			if πTemp002, πE = πg.ImportModule(πF, "cStringIO"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßStringIO); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßStringIO.ToObject(), πTemp003); πE != nil {
				continue
			}
			πF.PopCheckpoint()
			goto Label6
		Label7:
			if πE == nil {
			  continue
			}
			πE = nil
			πTemp010, πTemp011 = πF.ExcInfo()
			if πTemp001, πE = πg.ResolveGlobal(πF, ßImportError); πE != nil {
				continue
			}
			if πTemp005, πE = πg.IsInstance(πF, πTemp010.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp005 {
				goto Label8
			}
			πE = πF.Raise(πTemp010.ToObject(), nil, πTemp011.ToObject())
			continue
			// line 93: except ImportError:
			πF.SetLineno(93)
		Label8:
			// line 94: from StringIO import StringIO
			πF.SetLineno(94)
			if πTemp002, πE = πg.ImportModule(πF, "StringIO"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßStringIO); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßStringIO.ToObject(), πTemp003); πE != nil {
				continue
			}
			πF.RestoreExc(nil, nil)
			goto Label6
		Label6:
			// line 96: try:
			πF.SetLineno(96)
			πF.PushCheckpoint(10)
			// line 97: import errno
			πF.SetLineno(97)
			if πTemp002, πE = πg.ImportModule(πF, "errno"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßerrno.ToObject(), πTemp001); πE != nil {
				continue
			}
			πF.PopCheckpoint()
			goto Label9
		Label10:
			if πE == nil {
			  continue
			}
			πE = nil
			πTemp010, πTemp011 = πF.ExcInfo()
			if πTemp001, πE = πg.ResolveGlobal(πF, ßImportError); πE != nil {
				continue
			}
			if πTemp005, πE = πg.IsInstance(πF, πTemp010.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp005 {
				goto Label11
			}
			πE = πF.Raise(πTemp010.ToObject(), nil, πTemp011.ToObject())
			continue
			// line 98: except ImportError:
			πF.SetLineno(98)
		Label11:
			// line 99: errno = None
			πF.SetLineno(99)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßerrno.ToObject(), πTemp001); πE != nil {
				continue
			}
			πF.RestoreExc(nil, nil)
			goto Label9
		Label9:
			// line 100: EBADF = getattr(errno, 'EBADF', 9)
			πF.SetLineno(100)
			πTemp002 = πF.MakeArgs(3)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßerrno); πE != nil {
				continue
			}
			πTemp002[0] = πTemp001
			πTemp002[1] = ßEBADF.ToObject()
			πTemp002[2] = πg.NewInt(9).ToObject()
			if πTemp001, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
				continue
			}
			if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πF.Globals().SetItem(πF, ßEBADF.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 101: EINTR = getattr(errno, 'EINTR', 4)
			πF.SetLineno(101)
			πTemp002 = πF.MakeArgs(3)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßerrno); πE != nil {
				continue
			}
			πTemp002[0] = πTemp001
			πTemp002[1] = ßEINTR.ToObject()
			πTemp002[2] = πg.NewInt(4).ToObject()
			if πTemp001, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
				continue
			}
			if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πF.Globals().SetItem(πF, ßEINTR.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 103: __all__ = ["getfqdn", "create_connection"]
			πF.SetLineno(103)
			πTemp002 = make([]*πg.Object, 2)
			πTemp002[0] = ßgetfqdn.ToObject()
			πTemp002[1] = ßcreate_connection.ToObject()
			πTemp001 = πg.NewList(πTemp002...).ToObject()
			if πE = πF.Globals().SetItem(πF, ß__all__.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 104: __all__.extend(dir(_socket))
			πF.SetLineno(104)
			πTemp002 = πF.MakeArgs(1)
			πTemp012 = πF.MakeArgs(1)
			if πTemp001, πE = πg.ResolveGlobal(πF, ß_socket); πE != nil {
				continue
			}
			πTemp012[0] = πTemp001
			if πTemp001, πE = πg.ResolveGlobal(πF, ßdir); πE != nil {
				continue
			}
			if πTemp003, πE = πTemp001.Call(πF, πTemp012, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp012)
			πTemp002[0] = πTemp003
			if πTemp001, πE = πg.ResolveGlobal(πF, ß__all__); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßextend, nil); πE != nil {
				continue
			}
			if πTemp001, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			// line 108: _realsocket = socket
			πF.SetLineno(108)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßsocket); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_realsocket.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 132: def getfqdn(name=''):
			πF.SetLineno(132)
			πTemp013 = make([]πg.Param, 1)
			πTemp013[0] = πg.Param{Name: "name", Def: ß.ToObject()}
			πTemp001 = πg.NewFunction(πg.NewCode("getfqdn", "/usr/lib/python2.7/socket.py", πTemp013, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µname *πg.Object = πArgs[0]; _ = µname
				var µhostname *πg.Object = πg.UnboundLocal; _ = µhostname
				var µaliases *πg.Object = πg.UnboundLocal; _ = µaliases
				var µipaddrs *πg.Object = πg.UnboundLocal; _ = µipaddrs
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
				var πTemp008 *πg.BaseException
				_ = πTemp008
				var πTemp009 *πg.Traceback
				_ = πTemp009
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 5: goto Label5
					case 6: goto Label6
					case 7: goto Label7
					default: panic("unexpected function state")
					}
					// line 133: """Get fully qualified domain name from name.
					πF.SetLineno(133)
					// line 141: name = name.strip()
					πF.SetLineno(141)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µname, ßstrip, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					µname = πTemp002
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, µname); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(!πTemp004).ToObject()
					πTemp001 = πTemp002
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label1
					}
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µname, πg.NewStr("0.0.0.0").ToObject()); πE != nil {
						continue
					}
					πTemp001 = πTemp002
				Label1:
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label2
					}
					goto Label3
					// line 142: if not name or name == '0.0.0.0':
					πF.SetLineno(142)
				Label2:
					// line 143: name = gethostname()
					πF.SetLineno(143)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßgethostname); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					µname = πTemp002
					goto Label3
				Label3:
					// line 144: try:
					πF.SetLineno(144)
					πF.PushCheckpoint(5)
					// line 145: hostname, aliases, ipaddrs = gethostbyaddr(name)
					πF.SetLineno(145)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp005[0] = µname
					if πTemp001, πE = πg.ResolveGlobal(πF, ßgethostbyaddr); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp006}, πg.TieTarget{Target: &πTemp007}}}, πTemp002); πE != nil {
						continue
					}
					µhostname = πTemp001
					µaliases = πTemp006
					µipaddrs = πTemp007
					πF.PopCheckpoint()
					// line 149: aliases.insert(0, hostname)
					πF.SetLineno(149)
					πTemp005 = πF.MakeArgs(2)
					πTemp005[0] = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µhostname, "hostname"); πE != nil {
						continue
					}
					πTemp005[1] = µhostname
					if πE = πg.CheckLocal(πF, µaliases, "aliases"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µaliases, ßinsert, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πE = πg.CheckLocal(πF, µaliases, "aliases"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Iter(πF, µaliases); πE != nil {
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
						πTemp004 = !isStop
					} else {
						πTemp004 = true
						µname = πTemp002
					}
					if πE != nil || !πTemp004 {
						continue
					}
					πF.PushCheckpoint(6)            
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Contains(πF, µname, πg.NewStr(".").ToObject()); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(πTemp004).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label9
					}
					goto Label10
					// line 151: if '.' in name:
					πF.SetLineno(151)
				Label9:
					// line 152: break
					πF.SetLineno(152)
					πTemp003 = true
					continue
					goto Label10
				Label10:
					continue
				Label7:
					if πE != nil || πR != nil {
						continue
					}
					// line 154: name = hostname
					πF.SetLineno(154)
					if πE = πg.CheckLocal(πF, µhostname, "hostname"); πE != nil {
						continue
					}
					µname = µhostname
				Label8:
					goto Label4
				Label5:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp008, πTemp009 = πF.ExcInfo()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsInstance(πF, πTemp008.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label11
					}
					πE = πF.Raise(πTemp008.ToObject(), nil, πTemp009.ToObject())
					continue
					// line 146: except error:
					πF.SetLineno(146)
				Label11:
					// line 147: pass
					πF.SetLineno(147)
					πF.RestoreExc(nil, nil)
					goto Label4
				Label4:
					// line 155: return name
					πF.SetLineno(155)
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
			if πE = πF.Globals().SetItem(πF, ßgetfqdn.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 133: """Get fully qualified domain name from name.
			πF.SetLineno(133)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewStr("Get fully qualified domain name from name.\n\n    An empty argument is interpreted as meaning the local host.\n\n    First the hostname returned by gethostbyaddr() is checked, then\n    possibly existing aliases. In case no FQDN is available, hostname\n    from gethostname() is returned.\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp004, πE = πg.ResolveGlobal(πF, ßgetfqdn); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp004, ß__doc__, πTemp003); πE != nil {
				continue
			}
			// line 158: _socketmethods = (
			πF.SetLineno(158)
			πTemp002 = make([]*πg.Object, 14)
			πTemp002[0] = ßbind.ToObject()
			πTemp002[1] = ßconnect.ToObject()
			πTemp002[2] = ßconnect_ex.ToObject()
			πTemp002[3] = ßfileno.ToObject()
			πTemp002[4] = ßlisten.ToObject()
			πTemp002[5] = ßgetpeername.ToObject()
			πTemp002[6] = ßgetsockname.ToObject()
			πTemp002[7] = ßgetsockopt.ToObject()
			πTemp002[8] = ßsetsockopt.ToObject()
			πTemp002[9] = ßsendall.ToObject()
			πTemp002[10] = ßsetblocking.ToObject()
			πTemp002[11] = ßsettimeout.ToObject()
			πTemp002[12] = ßgettimeout.ToObject()
			πTemp002[13] = ßshutdown.ToObject()
			πTemp003 = πg.NewTuple(πTemp002...).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_socketmethods.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 172: _delegate_methods = ("recv", "recvfrom", "recv_into", "recvfrom_into",
			πF.SetLineno(172)
			πTemp003 = πg.NewTuple6(ßrecv.ToObject(), ßrecvfrom.ToObject(), ßrecv_into.ToObject(), ßrecvfrom_into.ToObject(), ßsend.ToObject(), ßsendto.ToObject()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_delegate_methods.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 175: class _closedsocket(object):
			πF.SetLineno(175)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			πTemp014 = πg.NewDict()
			if πTemp003, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp014.SetItem(πF, ß__module__.ToObject(), πTemp003); πE != nil {
				continue
			}
			_, πE = πg.NewCode("_closedsocket", "/usr/lib/python2.7/socket.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp014
				_ = πClass
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 []πg.Param
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 176: __slots__ = []
					πF.SetLineno(176)
					πTemp001 = make([]*πg.Object, 0)
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					if πE = πClass.SetItem(πF, ß__slots__.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 177: def _dummy(*args):
					πF.SetLineno(177)
					πTemp003 = make([]πg.Param, 0)
					πTemp002 = πg.NewFunction(πg.NewCode("_dummy", "/usr/lib/python2.7/socket.py", πTemp003, πg.CodeFlagVarArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µargs *πg.Object = πArgs[0]; _ = µargs
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
							πTemp001 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßEBADF); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = πg.NewStr("Bad file descriptor").ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 178: raise error(EBADF, 'Bad file descriptor')
							πF.SetLineno(178)
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
					if πE = πClass.SetItem(πF, ß_dummy.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 180: send = recv = recv_into = sendto = recvfrom = recvfrom_into = _dummy
					πF.SetLineno(180)
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ß_dummy); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßsend.ToObject(), πTemp004); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßrecv.ToObject(), πTemp004); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßrecv_into.ToObject(), πTemp004); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßsendto.ToObject(), πTemp004); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßrecvfrom.ToObject(), πTemp004); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßrecvfrom_into.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 181: __getattr__ = _dummy
					πF.SetLineno(181)
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ß_dummy); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ß__getattr__.ToObject(), πTemp004); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp004, πE = πTemp014.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp004 == nil {
				πTemp004 = πg.TypeType.ToObject()
			}
			if πTemp007, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("_closedsocket").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp014.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_closedsocket.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 187: class _socketobject(object):
			πF.SetLineno(187)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			πTemp014 = πg.NewDict()
			if πTemp003, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp014.SetItem(πF, ß__module__.ToObject(), πTemp003); πE != nil {
				continue
			}
			_, πE = πg.NewCode("_socketobject", "/usr/lib/python2.7/socket.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp014
				_ = πClass
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
				var πTemp006 []πg.Param
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 *πg.Object
				_ = πTemp008
				var πTemp009 *πg.Object
				_ = πTemp009
				var πTemp010 πg.KWArgs
				_ = πTemp010
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 191: __slots__ = ["_sock", "__weakref__"] + list(_delegate_methods)
					πF.SetLineno(191)
					πTemp002 = make([]*πg.Object, 2)
					πTemp002[0] = ß_sock.ToObject()
					πTemp002[1] = ß__weakref__.ToObject()
					πTemp003 = πg.NewList(πTemp002...).ToObject()
					πTemp002 = πF.MakeArgs(1)
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ß_delegate_methods); πE != nil {
						continue
					}
					πTemp002[0] = πTemp004
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßlist); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.Add(πF, πTemp003, πTemp005); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ß__slots__.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 193: def __init__(self, family=AF_INET, type=SOCK_STREAM, proto=0, _sock=None):
					πF.SetLineno(193)
					πTemp006 = make([]πg.Param, 5)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßAF_INET); πE != nil {
						continue
					}
					πTemp006[1] = πg.Param{Name: "family", Def: πTemp003}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßSOCK_STREAM); πE != nil {
						continue
					}
					πTemp006[2] = πg.Param{Name: "type", Def: πTemp003}
					πTemp006[3] = πg.Param{Name: "proto", Def: πg.NewInt(0).ToObject()}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp006[4] = πg.Param{Name: "_sock", Def: πTemp003}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/socket.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µfamily *πg.Object = πArgs[1]; _ = µfamily
						var µtype *πg.Object = πArgs[2]; _ = µtype
						var µproto *πg.Object = πArgs[3]; _ = µproto
						var µ_sock *πg.Object = πArgs[4]; _ = µ_sock
						var µmethod *πg.Object = πg.UnboundLocal; _ = µmethod
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 3: goto Label3
							case 4: goto Label4
							default: panic("unexpected function state")
							}
							if πE = πg.CheckLocal(πF, µ_sock, "_sock"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µ_sock == πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 194: if _sock is None:
							πF.SetLineno(194)
						Label1:
							// line 195: _sock = _realsocket(family, type, proto)
							πF.SetLineno(195)
							πTemp004 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µfamily, "family"); πE != nil {
								continue
							}
							πTemp004[0] = µfamily
							if πE = πg.CheckLocal(πF, µtype, "type"); πE != nil {
								continue
							}
							πTemp004[1] = µtype
							if πE = πg.CheckLocal(πF, µproto, "proto"); πE != nil {
								continue
							}
							πTemp004[2] = µproto
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_realsocket); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µ_sock = πTemp002
							goto Label2
						Label2:
							// line 196: self._sock = _sock
							πF.SetLineno(196)
							if πE = πg.CheckLocal(πF, µ_sock, "_sock"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µ_sock); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_sock, πTemp001); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_delegate_methods); πE != nil {
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
								µmethod = πTemp002
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(3)            
							// line 198: setattr(self, method, getattr(_sock, method))
							πF.SetLineno(198)
							πTemp004 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp004[0] = µself
							if πE = πg.CheckLocal(πF, µmethod, "method"); πE != nil {
								continue
							}
							πTemp004[1] = µmethod
							πTemp006 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µ_sock, "_sock"); πE != nil {
								continue
							}
							πTemp006[0] = µ_sock
							if πE = πg.CheckLocal(πF, µmethod, "method"); πE != nil {
								continue
							}
							πTemp006[1] = µmethod
							if πTemp002, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp004[2] = πTemp007
							if πTemp002, πE = πg.ResolveGlobal(πF, ßsetattr); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							continue
						Label4:
							if πE != nil || πR != nil {
								continue
							}
						Label5:
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
					// line 200: def close(self, _closedsocket=_closedsocket,
					πF.SetLineno(200)
					πTemp006 = make([]πg.Param, 4)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ß_closedsocket); πE != nil {
						continue
					}
					πTemp006[1] = πg.Param{Name: "_closedsocket", Def: πTemp004}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ß_delegate_methods); πE != nil {
						continue
					}
					πTemp006[2] = πg.Param{Name: "_delegate_methods", Def: πTemp004}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßsetattr); πE != nil {
						continue
					}
					πTemp006[3] = πg.Param{Name: "setattr", Def: πTemp004}
					πTemp003 = πg.NewFunction(πg.NewCode("close", "/usr/lib/python2.7/socket.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µ_closedsocket *πg.Object = πArgs[1]; _ = µ_closedsocket
						var µ_delegate_methods *πg.Object = πArgs[2]; _ = µ_delegate_methods
						var µsetattr *πg.Object = πArgs[3]; _ = µsetattr
						var µdummy *πg.Object = πg.UnboundLocal; _ = µdummy
						var µmethod *πg.Object = πg.UnboundLocal; _ = µmethod
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							// line 203: self._sock = _closedsocket()
							πF.SetLineno(203)
							if πE = πg.CheckLocal(πF, µ_closedsocket, "_closedsocket"); πE != nil {
								continue
							}
							if πTemp001, πE = µ_closedsocket.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_sock, πTemp002); πE != nil {
								continue
							}
							// line 204: dummy = self._sock._dummy
							πF.SetLineno(204)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_sock, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ß_dummy, nil); πE != nil {
								continue
							}
							µdummy = πTemp002
							if πE = πg.CheckLocal(πF, µ_delegate_methods, "_delegate_methods"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, µ_delegate_methods); πE != nil {
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
								µmethod = πTemp002
							}
							if πE != nil || !πTemp004 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 206: setattr(self, method, dummy)
							πF.SetLineno(206)
							πTemp005 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp005[0] = µself
							if πE = πg.CheckLocal(πF, µmethod, "method"); πE != nil {
								continue
							}
							πTemp005[1] = µmethod
							if πE = πg.CheckLocal(πF, µdummy, "dummy"); πE != nil {
								continue
							}
							πTemp005[2] = µdummy
							if πE = πg.CheckLocal(πF, µsetattr, "setattr"); πE != nil {
								continue
							}
							if πTemp002, πE = µsetattr.Call(πF, πTemp005, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßclose.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 209: def accept(self):
					πF.SetLineno(209)
					πTemp006 = make([]πg.Param, 1)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("accept", "/usr/lib/python2.7/socket.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µsock *πg.Object = πg.UnboundLocal; _ = µsock
						var µaddr *πg.Object = πg.UnboundLocal; _ = µaddr
						var πTemp001 *πg.Object
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
							// line 210: sock, addr = self._sock.accept()
							πF.SetLineno(210)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_sock, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßaccept, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp003}}}, πTemp001); πE != nil {
								continue
							}
							µsock = πTemp002
							µaddr = πTemp003
							// line 211: return _socketobject(_sock=sock), addr
							πF.SetLineno(211)
							if πE = πg.CheckLocal(πF, µsock, "sock"); πE != nil {
								continue
							}
							πTemp004 = πg.KWArgs{
								{"_sock", µsock},
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_socketobject); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µaddr, "addr"); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple2(πTemp003, µaddr).ToObject()
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
					if πE = πClass.SetItem(πF, ßaccept.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 214: def dup(self):
					πF.SetLineno(214)
					πTemp006 = make([]πg.Param, 1)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("dup", "/usr/lib/python2.7/socket.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 πg.KWArgs
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
							// line 215: """dup() -> socket object
							πF.SetLineno(215)
							// line 218: return _socketobject(_sock=self._sock)
							πF.SetLineno(218)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_sock, nil); πE != nil {
								continue
							}
							πTemp002 = πg.KWArgs{
								{"_sock", πTemp001},
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_socketobject); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ßdup.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 215: """dup() -> socket object
					πF.SetLineno(215)
					// line 215: """dup() -> socket object
					πF.SetLineno(215)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("dup() -> socket object\n\n        Return a new socket object connected to the same system resource.").ToObject()); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp007}, πg.NewStr("dup() -> socket object\n\n        Return a new socket object connected to the same system resource.").ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ßdup); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp008, ß__doc__, πTemp007); πE != nil {
						continue
					}
					// line 220: def makefile(self, mode='r', bufsize=-1):
					πF.SetLineno(220)
					πTemp006 = make([]πg.Param, 3)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp006[1] = πg.Param{Name: "mode", Def: ßr.ToObject()}
					if πTemp008, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp006[2] = πg.Param{Name: "bufsize", Def: πTemp008}
					πTemp007 = πg.NewFunction(πg.NewCode("makefile", "/usr/lib/python2.7/socket.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µmode *πg.Object = πArgs[1]; _ = µmode
						var µbufsize *πg.Object = πArgs[2]; _ = µbufsize
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
							// line 221: """makefile([mode[, bufsize]]) -> file object
							πF.SetLineno(221)
							// line 225: return _fileobject(self._sock, mode, bufsize)
							πF.SetLineno(225)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_sock, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µmode, "mode"); πE != nil {
								continue
							}
							πTemp001[1] = µmode
							if πE = πg.CheckLocal(πF, µbufsize, "bufsize"); πE != nil {
								continue
							}
							πTemp001[2] = µbufsize
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_fileobject); πE != nil {
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
					if πE = πClass.SetItem(πF, ßmakefile.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 221: """makefile([mode[, bufsize]]) -> file object
					πF.SetLineno(221)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp008}, πg.NewStr("makefile([mode[, bufsize]]) -> file object\n\n        Return a regular file object corresponding to the socket.  The mode\n        and bufsize arguments are as for the built-in open() function.").ToObject()); πE != nil {
						continue
					}
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßmakefile); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp009, ß__doc__, πTemp008); πE != nil {
						continue
					}
					// line 227: family = property(lambda self: self._sock.family, doc="the socket family")
					πF.SetLineno(227)
					πTemp002 = πF.MakeArgs(1)
					πTemp006 = make([]πg.Param, 1)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("<lambda>", "/usr/lib/python2.7/socket.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 227: family = property(lambda self: self._sock.family, doc="the socket family")
							πF.SetLineno(227)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_sock, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßfamily, nil); πE != nil {
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
					πTemp002[0] = πTemp008
					πTemp010 = πg.KWArgs{
						{"doc", πg.NewStr("the socket family").ToObject()},
					}
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ßproperty); πE != nil {
						continue
					}
					if πTemp009, πE = πTemp008.Call(πF, πTemp002, πTemp010); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πE = πClass.SetItem(πF, ßfamily.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 228: type = property(lambda self: self._sock.type, doc="the socket type")
					πF.SetLineno(228)
					πTemp002 = πF.MakeArgs(1)
					πTemp006 = make([]πg.Param, 1)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("<lambda>", "/usr/lib/python2.7/socket.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 228: type = property(lambda self: self._sock.type, doc="the socket type")
							πF.SetLineno(228)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_sock, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßtype, nil); πE != nil {
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
					πTemp002[0] = πTemp008
					πTemp010 = πg.KWArgs{
						{"doc", πg.NewStr("the socket type").ToObject()},
					}
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ßproperty); πE != nil {
						continue
					}
					if πTemp009, πE = πTemp008.Call(πF, πTemp002, πTemp010); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πE = πClass.SetItem(πF, ßtype.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 229: proto = property(lambda self: self._sock.proto, doc="the socket protocol")
					πF.SetLineno(229)
					πTemp002 = πF.MakeArgs(1)
					πTemp006 = make([]πg.Param, 1)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("<lambda>", "/usr/lib/python2.7/socket.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 229: proto = property(lambda self: self._sock.proto, doc="the socket protocol")
							πF.SetLineno(229)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_sock, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßproto, nil); πE != nil {
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
					πTemp002[0] = πTemp008
					πTemp010 = πg.KWArgs{
						{"doc", πg.NewStr("the socket protocol").ToObject()},
					}
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ßproperty); πE != nil {
						continue
					}
					if πTemp009, πE = πTemp008.Call(πF, πTemp002, πTemp010); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πE = πClass.SetItem(πF, ßproto.ToObject(), πTemp009); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp004, πE = πTemp014.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp004 == nil {
				πTemp004 = πg.TypeType.ToObject()
			}
			if πTemp007, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("_socketobject").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp014.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_socketobject.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 231: def meth(name,self,*args):
			πF.SetLineno(231)
			πTemp013 = make([]πg.Param, 2)
			πTemp013[0] = πg.Param{Name: "name", Def: nil}
			πTemp013[1] = πg.Param{Name: "self", Def: nil}
			πTemp003 = πg.NewFunction(πg.NewCode("meth", "/usr/lib/python2.7/socket.py", πTemp013, πg.CodeFlagVarArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µname *πg.Object = πArgs[0]; _ = µname
				var µself *πg.Object = πArgs[1]; _ = µself
				var µargs *πg.Object = πArgs[2]; _ = µargs
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
					// line 232: return getattr(self._sock,name)(*args)
					πF.SetLineno(232)
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µself, ß_sock, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
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
					if πTemp002, πE = πg.Invoke(πF, πTemp003, nil, µargs, nil, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßmeth.ToObject(), πTemp003); πE != nil {
				continue
			}
			if πTemp007, πE = πg.ResolveGlobal(πF, ß_socketmethods); πE != nil {
				continue
			}
			if πTemp004, πE = πg.Iter(πF, πTemp007); πE != nil {
				continue
			}
			πF.PushCheckpoint(13)
			πTemp005 = false
		Label12:
			if πE != nil || πR != nil {
				continue
			}
			if πTemp005 {
				πF.PopCheckpoint()
				goto Label14
			}
			if πTemp007, πE = πg.Next(πF, πTemp004); πE != nil {
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
				if πE = πF.Globals().SetItem(πF, ß_m.ToObject(), πTemp007); πE != nil {
					continue
				}
			}
			if πE != nil || !πTemp006 {
				continue
			}
			πF.PushCheckpoint(12)            
			// line 235: p = partial(meth,_m)
			πF.SetLineno(235)
			πTemp002 = πF.MakeArgs(2)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßmeth); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			if πTemp007, πE = πg.ResolveGlobal(πF, ß_m); πE != nil {
				continue
			}
			πTemp002[1] = πTemp007
			if πTemp007, πE = πg.ResolveGlobal(πF, ßpartial); πE != nil {
				continue
			}
			if πTemp008, πE = πTemp007.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πF.Globals().SetItem(πF, ßp.ToObject(), πTemp008); πE != nil {
				continue
			}
			// line 236: p.__name__ = _m
			πF.SetLineno(236)
			if πTemp007, πE = πg.ResolveGlobal(πF, ß_m); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp008}, πTemp007); πE != nil {
				continue
			}
			if πTemp009, πE = πg.ResolveGlobal(πF, ßp); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp009, ß__name__, πTemp008); πE != nil {
				continue
			}
			// line 238: m = MethodType(p,None,_socketobject)
			πF.SetLineno(238)
			πTemp002 = πF.MakeArgs(3)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßp); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			if πTemp007, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp002[1] = πTemp007
			if πTemp007, πE = πg.ResolveGlobal(πF, ß_socketobject); πE != nil {
				continue
			}
			πTemp002[2] = πTemp007
			if πTemp007, πE = πg.ResolveGlobal(πF, ßMethodType); πE != nil {
				continue
			}
			if πTemp008, πE = πTemp007.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πF.Globals().SetItem(πF, ßm.ToObject(), πTemp008); πE != nil {
				continue
			}
			// line 239: setattr(_socketobject,_m,m)
			πF.SetLineno(239)
			πTemp002 = πF.MakeArgs(3)
			if πTemp007, πE = πg.ResolveGlobal(πF, ß_socketobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			if πTemp007, πE = πg.ResolveGlobal(πF, ß_m); πE != nil {
				continue
			}
			πTemp002[1] = πTemp007
			if πTemp007, πE = πg.ResolveGlobal(πF, ßm); πE != nil {
				continue
			}
			πTemp002[2] = πTemp007
			if πTemp007, πE = πg.ResolveGlobal(πF, ßsetattr); πE != nil {
				continue
			}
			if πTemp008, πE = πTemp007.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			continue
		Label13:
			if πE != nil || πR != nil {
				continue
			}
		Label14:
			// line 241: socket = SocketType = _socketobject
			πF.SetLineno(241)
			if πTemp004, πE = πg.ResolveGlobal(πF, ß_socketobject); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßsocket.ToObject(), πTemp004); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßSocketType.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 243: class _fileobject(object):
			πF.SetLineno(243)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp008, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp008
			πTemp014 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp014.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("_fileobject", "/usr/lib/python2.7/socket.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp014
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
				var πTemp006 πg.KWArgs
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
					// line 244: """Faux file object attached to a socket object."""
					πF.SetLineno(244)
					// line 244: """Faux file object attached to a socket object."""
					πF.SetLineno(244)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Faux file object attached to a socket object.").ToObject()); πE != nil {
						continue
					}
					// line 246: default_bufsize = 8192
					πF.SetLineno(246)
					if πE = πClass.SetItem(πF, ßdefault_bufsize.ToObject(), πg.NewInt(8192).ToObject()); πE != nil {
						continue
					}
					// line 247: name = "<socket>"
					πF.SetLineno(247)
					if πE = πClass.SetItem(πF, ßname.ToObject(), πg.NewStr("<socket>").ToObject()); πE != nil {
						continue
					}
					// line 249: __slots__ = ["mode", "bufsize", "softspace",
					πF.SetLineno(249)
					πTemp001 = make([]*πg.Object, 10)
					πTemp001[0] = ßmode.ToObject()
					πTemp001[1] = ßbufsize.ToObject()
					πTemp001[2] = ßsoftspace.ToObject()
					πTemp001[3] = ß_sock.ToObject()
					πTemp001[4] = ß_rbufsize.ToObject()
					πTemp001[5] = ß_wbufsize.ToObject()
					πTemp001[6] = ß_rbuf.ToObject()
					πTemp001[7] = ß_wbuf.ToObject()
					πTemp001[8] = ß_wbuf_len.ToObject()
					πTemp001[9] = ß_close.ToObject()
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					if πE = πClass.SetItem(πF, ß__slots__.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 254: def __init__(self, sock, mode='rb', bufsize=-1, close=False):
					πF.SetLineno(254)
					πTemp003 = make([]πg.Param, 5)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp003[1] = πg.Param{Name: "sock", Def: nil}
					πTemp003[2] = πg.Param{Name: "mode", Def: ßrb.ToObject()}
					if πTemp004, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp003[3] = πg.Param{Name: "bufsize", Def: πTemp004}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßFalse); πE != nil {
						continue
					}
					πTemp003[4] = πg.Param{Name: "close", Def: πTemp004}
					πTemp002 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/socket.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µsock *πg.Object = πArgs[1]; _ = µsock
						var µmode *πg.Object = πArgs[2]; _ = µmode
						var µbufsize *πg.Object = πArgs[3]; _ = µbufsize
						var µclose *πg.Object = πArgs[4]; _ = µclose
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
						_ = πTemp002
						var πTemp003 *πg.Object
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
							// line 255: self._sock = sock
							πF.SetLineno(255)
							if πE = πg.CheckLocal(πF, µsock, "sock"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µsock); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_sock, πTemp001); πE != nil {
								continue
							}
							// line 256: self.mode = mode # Not actually used in this version
							πF.SetLineno(256)
							if πE = πg.CheckLocal(πF, µmode, "mode"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µmode); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßmode, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µbufsize, "bufsize"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.LT(πF, µbufsize, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							goto Label2
							// line 257: if bufsize < 0:
							πF.SetLineno(257)
						Label1:
							// line 258: bufsize = self.default_bufsize
							πF.SetLineno(258)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdefault_bufsize, nil); πE != nil {
								continue
							}
							µbufsize = πTemp001
							goto Label2
						Label2:
							// line 259: self.bufsize = bufsize
							πF.SetLineno(259)
							if πE = πg.CheckLocal(πF, µbufsize, "bufsize"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µbufsize); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßbufsize, πTemp001); πE != nil {
								continue
							}
							// line 260: self.softspace = False
							πF.SetLineno(260)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßsoftspace, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µbufsize, "bufsize"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, µbufsize, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label3
							}
							if πE = πg.CheckLocal(πF, µbufsize, "bufsize"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, µbufsize, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label4
							}
							goto Label5
							// line 264: if bufsize == 0:
							πF.SetLineno(264)
						Label3:
							// line 265: self._rbufsize = 1
							πF.SetLineno(265)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_rbufsize, πTemp001); πE != nil {
								continue
							}
							goto Label6
							// line 266: elif bufsize == 1:
							πF.SetLineno(266)
						Label4:
							// line 267: self._rbufsize = self.default_bufsize
							πF.SetLineno(267)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdefault_bufsize, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_rbufsize, πTemp003); πE != nil {
								continue
							}
							goto Label6
						Label5:
							// line 269: self._rbufsize = bufsize
							πF.SetLineno(269)
							if πE = πg.CheckLocal(πF, µbufsize, "bufsize"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µbufsize); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_rbufsize, πTemp001); πE != nil {
								continue
							}
							goto Label6
						Label6:
							// line 270: self._wbufsize = bufsize
							πF.SetLineno(270)
							if πE = πg.CheckLocal(πF, µbufsize, "bufsize"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µbufsize); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_wbufsize, πTemp001); πE != nil {
								continue
							}
							// line 275: self._rbuf = StringIO()
							πF.SetLineno(275)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßStringIO); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_rbuf, πTemp001); πE != nil {
								continue
							}
							// line 276: self._wbuf = [] # A list of strings
							πF.SetLineno(276)
							πTemp004 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp004...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_wbuf, πTemp003); πE != nil {
								continue
							}
							// line 277: self._wbuf_len = 0
							πF.SetLineno(277)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_wbuf_len, πTemp001); πE != nil {
								continue
							}
							// line 278: self._close = close
							πF.SetLineno(278)
							if πE = πg.CheckLocal(πF, µclose, "close"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µclose); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_close, πTemp001); πE != nil {
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
					// line 280: def _getclosed(self):
					πF.SetLineno(280)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("_getclosed", "/usr/lib/python2.7/socket.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 281: return self._sock is None
							πF.SetLineno(281)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_sock, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp002 == πTemp003).ToObject()
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
					if πE = πClass.SetItem(πF, ß_getclosed.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 282: closed = property(_getclosed, doc="True if the file is closed")
					πF.SetLineno(282)
					πTemp001 = πF.MakeArgs(1)
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ß_getclosed); πE != nil {
						continue
					}
					πTemp001[0] = πTemp005
					πTemp006 = πg.KWArgs{
						{"doc", πg.NewStr("True if the file is closed").ToObject()},
					}
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßproperty); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp005.Call(πF, πTemp001, πTemp006); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πClass.SetItem(πF, ßclosed.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 284: def close(self):
					πF.SetLineno(284)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("close", "/usr/lib/python2.7/socket.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							default: panic("unexpected function state")
							}
							// line 285: try:
							πF.SetLineno(285)
							πF.PushCheckpoint(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_sock, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label2
							}
							goto Label3
							// line 286: if self._sock:
							πF.SetLineno(286)
						Label2:
							// line 287: self.flush()
							πF.SetLineno(287)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßflush, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							goto Label3
						Label3:
							πF.PopCheckpoint()
							goto Label1
						Label1:
							πTemp004, πTemp005 = πF.RestoreExc(nil, nil)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_close, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label4
							}
							goto Label5
							// line 289: if self._close:
							πF.SetLineno(289)
						Label4:
							// line 290: self._sock.close()
							πF.SetLineno(290)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_sock, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßclose, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							goto Label5
						Label5:
							// line 291: self._sock = None
							πF.SetLineno(291)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_sock, πTemp003); πE != nil {
								continue
							}
							if πTemp004 != nil {
								πE = πF.Raise(πTemp004.ToObject(), nil, πTemp005.ToObject())
								continue
							}
							if πR != nil {
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
					if πE = πClass.SetItem(πF, ßclose.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 293: def __del__(self):
					πF.SetLineno(293)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("__del__", "/usr/lib/python2.7/socket.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.BaseException
						_ = πTemp003
						var πTemp004 *πg.Traceback
						_ = πTemp004
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							// line 294: try:
							πF.SetLineno(294)
							πF.PushCheckpoint(2)
							// line 295: self.close()
							πF.SetLineno(295)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßclose, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
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
							goto Label3
							// line 296: except:
							πF.SetLineno(296)
						Label3:
							// line 298: pass
							πF.SetLineno(298)
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
					if πE = πClass.SetItem(πF, ß__del__.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 300: def flush(self):
					πF.SetLineno(300)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("flush", "/usr/lib/python2.7/socket.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µdata *πg.Object = πg.UnboundLocal; _ = µdata
						var µbuffer_size *πg.Object = πg.UnboundLocal; _ = µbuffer_size
						var µdata_size *πg.Object = πg.UnboundLocal; _ = µdata_size
						var µwrite_offset *πg.Object = πg.UnboundLocal; _ = µwrite_offset
						var µview *πg.Object = πg.UnboundLocal; _ = µview
						var µremainder *πg.Object = πg.UnboundLocal; _ = µremainder
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 bool
						_ = πTemp005
						var πTemp006 *πg.BaseException
						_ = πTemp006
						var πTemp007 *πg.Traceback
						_ = πTemp007
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 3: goto Label3
							case 4: goto Label4
							case 5: goto Label5
							default: panic("unexpected function state")
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_wbuf, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							goto Label2
							// line 301: if self._wbuf:
							πF.SetLineno(301)
						Label1:
							// line 302: data = "".join(self._wbuf)
							πF.SetLineno(302)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_wbuf, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πTemp001, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µdata = πTemp004
							// line 303: self._wbuf = []
							πF.SetLineno(303)
							πTemp003 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp003...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_wbuf, πTemp004); πE != nil {
								continue
							}
							// line 304: self._wbuf_len = 0
							πF.SetLineno(304)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_wbuf_len, πTemp001); πE != nil {
								continue
							}
							// line 305: buffer_size = max(self._rbufsize, self.default_bufsize)
							πF.SetLineno(305)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_rbufsize, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdefault_bufsize, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßmax); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µbuffer_size = πTemp004
							// line 306: data_size = len(data)
							πF.SetLineno(306)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πTemp003[0] = µdata
							if πTemp001, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µdata_size = πTemp004
							// line 307: write_offset = 0
							πF.SetLineno(307)
							µwrite_offset = πg.NewInt(0).ToObject()
							// line 308: view = data #memoryview(data)
							πF.SetLineno(308)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							µview = µdata
							// line 309: try:
							πF.SetLineno(309)
							πF.PushCheckpoint(3)
							// line 310: while write_offset < data_size:
							πF.SetLineno(310)
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
							if πE = πg.CheckLocal(πF, µwrite_offset, "write_offset"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µdata_size, "data_size"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.LT(πF, µwrite_offset, µdata_size); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(4)            
							// line 311: self._sock.sendall(view[write_offset:write_offset+buffer_size])
							πF.SetLineno(311)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µwrite_offset, "write_offset"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µwrite_offset, "write_offset"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µbuffer_size, "buffer_size"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Add(πF, µwrite_offset, µbuffer_size); πE != nil {
								continue
							}
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{µwrite_offset, πTemp004, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µview, "view"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µview, πTemp001); πE != nil {
								continue
							}
							πTemp003[0] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_sock, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßsendall, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 312: write_offset += buffer_size
							πF.SetLineno(312)
							if πE = πg.CheckLocal(πF, µwrite_offset, "write_offset"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µbuffer_size, "buffer_size"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.IAdd(πF, µwrite_offset, µbuffer_size); πE != nil {
								continue
							}
							µwrite_offset = πTemp001
							continue
						Label5:
							if πE != nil || πR != nil {
								continue
							}
						Label6:
							πF.PopCheckpoint()
							goto Label3
						Label3:
							πTemp006, πTemp007 = πF.RestoreExc(nil, nil)
							if πE = πg.CheckLocal(πF, µwrite_offset, "write_offset"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µdata_size, "data_size"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.LT(πF, µwrite_offset, µdata_size); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label7
							}
							goto Label8
							// line 314: if write_offset < data_size:
							πF.SetLineno(314)
						Label7:
							// line 315: remainder = data[write_offset:]
							πF.SetLineno(315)
							if πE = πg.CheckLocal(πF, µwrite_offset, "write_offset"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{µwrite_offset, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µdata, πTemp001); πE != nil {
								continue
							}
							µremainder = πTemp004
							// line 316: del view, data  # explicit free
							πF.SetLineno(316)
							if πE = πg.CheckLocal(πF, µview, "view"); πE != nil {
								continue
							}
							µview = πg.UnboundLocal
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							µdata = πg.UnboundLocal
							// line 317: self._wbuf.append(remainder)
							πF.SetLineno(317)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µremainder, "remainder"); πE != nil {
								continue
							}
							πTemp003[0] = µremainder
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_wbuf, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßappend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 318: self._wbuf_len = len(remainder)
							πF.SetLineno(318)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µremainder, "remainder"); πE != nil {
								continue
							}
							πTemp003[0] = µremainder
							if πTemp001, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_wbuf_len, πTemp001); πE != nil {
								continue
							}
							goto Label8
						Label8:
							if πTemp006 != nil {
								πE = πF.Raise(πTemp006.ToObject(), nil, πTemp007.ToObject())
								continue
							}
							if πR != nil {
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
					if πE = πClass.SetItem(πF, ßflush.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 320: def fileno(self):
					πF.SetLineno(320)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("fileno", "/usr/lib/python2.7/socket.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 321: return self._sock.fileno()
							πF.SetLineno(321)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_sock, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßfileno, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßfileno.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 323: def write(self, data):
					πF.SetLineno(323)
					πTemp003 = make([]πg.Param, 2)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp003[1] = πg.Param{Name: "data", Def: nil}
					πTemp010 = πg.NewFunction(πg.NewCode("write", "/usr/lib/python2.7/socket.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µdata *πg.Object = πArgs[1]; _ = µdata
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
						var πTemp008 bool
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
							// line 324: data = str(data) # XXX Should really reject non-string non-buffers
							πF.SetLineno(324)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πTemp001[0] = µdata
							if πTemp002, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µdata = πTemp003
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µdata); πE != nil {
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
							// line 325: if not data:
							πF.SetLineno(325)
						Label1:
							// line 326: return
							πF.SetLineno(326)
							πR = πg.None
							continue
							goto Label2
						Label2:
							// line 327: self._wbuf.append(data)
							πF.SetLineno(327)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πTemp001[0] = µdata
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_wbuf, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 328: self._wbuf_len += len(data)
							πF.SetLineno(328)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_wbuf_len, nil); πE != nil {
								continue
							}
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πTemp001[0] = µdata
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp003, πE = πg.IAdd(πF, πTemp002, πTemp005); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_wbuf_len, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ß_wbufsize, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Eq(πF, πTemp005, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp003
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label3
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ß_wbufsize, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Eq(πF, πTemp007, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp003 = πTemp005
							if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if !πTemp006 {
								goto Label4
							}
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.Contains(πF, µdata, πg.NewStr("\n").ToObject()); πE != nil {
								continue
							}
							πTemp005 = πg.GetBool(πTemp008).ToObject()
							πTemp003 = πTemp005
						Label4:
							πTemp002 = πTemp003
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label3
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ß_wbufsize, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GT(πF, πTemp007, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp003 = πTemp005
							if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if !πTemp006 {
								goto Label5
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ß_wbuf_len, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, µself, ß_wbufsize, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GE(πF, πTemp007, πTemp009); πE != nil {
								continue
							}
							πTemp003 = πTemp005
						Label5:
							πTemp002 = πTemp003
						Label3:
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label6
							}
							goto Label7
							// line 329: if (self._wbufsize == 0 or
							πF.SetLineno(329)
						Label6:
							// line 332: self.flush()
							πF.SetLineno(332)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßflush, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßwrite.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 334: def writelines(self, list):
					πF.SetLineno(334)
					πTemp003 = make([]πg.Param, 2)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp003[1] = πg.Param{Name: "list", Def: nil}
					πTemp011 = πg.NewFunction(πg.NewCode("writelines", "/usr/lib/python2.7/socket.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µlist *πg.Object = πArgs[1]; _ = µlist
						var µlines *πg.Object = πg.UnboundLocal; _ = µlines
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
						var πTemp006 bool
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
							// line 337: lines = filter(None, map(str, list))
							πF.SetLineno(337)
							πTemp001 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp003 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µlist, "list"); πE != nil {
								continue
							}
							πTemp003[1] = µlist
							if πTemp002, πE = πg.ResolveGlobal(πF, ßmap); πE != nil {
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
							µlines = πTemp004
							// line 338: self._wbuf_len += sum(map(len, lines))
							πF.SetLineno(338)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_wbuf_len, nil); πE != nil {
								continue
							}
							πTemp001 = πF.MakeArgs(1)
							πTemp003 = πF.MakeArgs(2)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							πTemp003[0] = πTemp004
							if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
								continue
							}
							πTemp003[1] = µlines
							if πTemp004, πE = πg.ResolveGlobal(πF, ßmap); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πTemp001[0] = πTemp005
							if πTemp004, πE = πg.ResolveGlobal(πF, ßsum); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp004, πE = πg.IAdd(πF, πTemp002, πTemp005); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_wbuf_len, πTemp004); πE != nil {
								continue
							}
							// line 339: self._wbuf.extend(lines)
							πF.SetLineno(339)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
								continue
							}
							πTemp001[0] = µlines
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_wbuf, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßextend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ß_wbufsize, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.LE(πF, πTemp005, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp004
							if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ß_wbuf_len, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ß_wbufsize, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GE(πF, πTemp005, πTemp007); πE != nil {
								continue
							}
							πTemp002 = πTemp004
						Label1:
							if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label2
							}
							goto Label3
							// line 340: if (self._wbufsize <= 1 or
							πF.SetLineno(340)
						Label2:
							// line 342: self.flush()
							πF.SetLineno(342)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßflush, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßwritelines.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 344: def read(self, size=-1):
					πF.SetLineno(344)
					πTemp003 = make([]πg.Param, 2)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					if πTemp013, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp003[1] = πg.Param{Name: "size", Def: πTemp013}
					πTemp012 = πg.NewFunction(πg.NewCode("read", "/usr/lib/python2.7/socket.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µsize *πg.Object = πArgs[1]; _ = µsize
						var µrbufsize *πg.Object = πg.UnboundLocal; _ = µrbufsize
						var µbuf *πg.Object = πg.UnboundLocal; _ = µbuf
						var µdata *πg.Object = πg.UnboundLocal; _ = µdata
						var µe *πg.Object = πg.UnboundLocal; _ = µe
						var µbuf_len *πg.Object = πg.UnboundLocal; _ = µbuf_len
						var µrv *πg.Object = πg.UnboundLocal; _ = µrv
						var µleft *πg.Object = πg.UnboundLocal; _ = µleft
						var µn *πg.Object = πg.UnboundLocal; _ = µn
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
						var πTemp006 *πg.BaseException
						_ = πTemp006
						var πTemp007 *πg.Traceback
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
							case 4: goto Label4
							case 5: goto Label5
							case 8: goto Label8
							case 16: goto Label16
							case 17: goto Label17
							case 20: goto Label20
							default: panic("unexpected function state")
							}
							// line 348: rbufsize = max(self._rbufsize, self.default_bufsize)
							πF.SetLineno(348)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_rbufsize, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdefault_bufsize, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßmax); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µrbufsize = πTemp003
							// line 352: buf = self._rbuf
							πF.SetLineno(352)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_rbuf, nil); πE != nil {
								continue
							}
							µbuf = πTemp002
							// line 353: buf.seek(0, 2)  # seek end
							πF.SetLineno(353)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = πg.NewInt(0).ToObject()
							πTemp001[1] = πg.NewInt(2).ToObject()
							if πE = πg.CheckLocal(πF, µbuf, "buf"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µbuf, ßseek, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
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
							// line 354: if size < 0:
							πF.SetLineno(354)
						Label1:
							// line 356: self._rbuf = StringIO()  # reset _rbuf.  we consume it via buf.
							πF.SetLineno(356)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßStringIO); πE != nil {
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
							if πE = πg.SetAttr(πF, µself, ß_rbuf, πTemp002); πE != nil {
								continue
							}
							// line 357: while True:
							πF.SetLineno(357)
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
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(4)            
							// line 358: try:
							πF.SetLineno(358)
							πF.PushCheckpoint(8)
							// line 359: data = self._sock.recv(rbufsize)
							πF.SetLineno(359)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µrbufsize, "rbufsize"); πE != nil {
								continue
							}
							πTemp001[0] = µrbufsize
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_sock, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßrecv, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µdata = πTemp002
							πF.PopCheckpoint()
							goto Label7
						Label8:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp006, πTemp007 = πF.ExcInfo()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsInstance(πF, πTemp006.ToObject(), πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label9
							}
							πE = πF.Raise(πTemp006.ToObject(), nil, πTemp007.ToObject())
							continue
							// line 360: except error, e:
							πF.SetLineno(360)
						Label9:
							µe = πTemp006.ToObject()
							πTemp003 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µe, "e"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, µe, ßargs, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, πTemp009, πTemp003); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßEINTR); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, πTemp008, πTemp003); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label10
							}
							goto Label11
							// line 361: if e.args[0] == EINTR:
							πF.SetLineno(361)
						Label10:
							// line 362: continue
							πF.SetLineno(362)
							continue
							goto Label11
						Label11:
							// line 363: raise
							πF.SetLineno(363)
							πE = πF.Raise(nil, nil, nil)
							continue
							πF.RestoreExc(nil, nil)
							goto Label7
						Label7:
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, µdata); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp005).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label12
							}
							goto Label13
							// line 364: if not data:
							πF.SetLineno(364)
						Label12:
							// line 365: break
							πF.SetLineno(365)
							πTemp004 = true
							continue
							goto Label13
						Label13:
							// line 366: buf.write(data)
							πF.SetLineno(366)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πTemp001[0] = µdata
							if πE = πg.CheckLocal(πF, µbuf, "buf"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µbuf, ßwrite, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							continue
						Label5:
							if πE != nil || πR != nil {
								continue
							}
						Label6:
							// line 367: return buf.getvalue()
							πF.SetLineno(367)
							if πE = πg.CheckLocal(πF, µbuf, "buf"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µbuf, ßgetvalue, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πR = πTemp003
							continue
							goto Label3
						Label2:
							// line 370: buf_len = buf.tell()
							πF.SetLineno(370)
							if πE = πg.CheckLocal(πF, µbuf, "buf"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µbuf, ßtell, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µbuf_len = πTemp003
							if πE = πg.CheckLocal(πF, µbuf_len, "buf_len"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µsize, "size"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GE(πF, µbuf_len, µsize); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label14
							}
							goto Label15
							// line 371: if buf_len >= size:
							πF.SetLineno(371)
						Label14:
							// line 373: buf.seek(0)
							πF.SetLineno(373)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µbuf, "buf"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µbuf, ßseek, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 374: rv = buf.read(size)
							πF.SetLineno(374)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µsize, "size"); πE != nil {
								continue
							}
							πTemp001[0] = µsize
							if πE = πg.CheckLocal(πF, µbuf, "buf"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µbuf, ßread, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µrv = πTemp003
							// line 375: self._rbuf = StringIO()
							πF.SetLineno(375)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßStringIO); πE != nil {
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
							if πE = πg.SetAttr(πF, µself, ß_rbuf, πTemp002); πE != nil {
								continue
							}
							// line 376: self._rbuf.write(buf.read())
							πF.SetLineno(376)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µbuf, "buf"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µbuf, ßread, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_rbuf, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßwrite, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 377: return rv
							πF.SetLineno(377)
							if πE = πg.CheckLocal(πF, µrv, "rv"); πE != nil {
								continue
							}
							πR = µrv
							continue
							goto Label15
						Label15:
							// line 379: self._rbuf = StringIO()  # reset _rbuf.  we consume it via buf.
							πF.SetLineno(379)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßStringIO); πE != nil {
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
							if πE = πg.SetAttr(πF, µself, ß_rbuf, πTemp002); πE != nil {
								continue
							}
							// line 380: while True:
							πF.SetLineno(380)
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
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(16)            
							// line 381: left = size - buf_len
							πF.SetLineno(381)
							if πE = πg.CheckLocal(πF, µsize, "size"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µbuf_len, "buf_len"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Sub(πF, µsize, µbuf_len); πE != nil {
								continue
							}
							µleft = πTemp002
							// line 387: try:
							πF.SetLineno(387)
							πF.PushCheckpoint(20)
							// line 388: data = self._sock.recv(left)
							πF.SetLineno(388)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µleft, "left"); πE != nil {
								continue
							}
							πTemp001[0] = µleft
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_sock, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßrecv, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µdata = πTemp002
							πF.PopCheckpoint()
							goto Label19
						Label20:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp006, πTemp007 = πF.ExcInfo()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsInstance(πF, πTemp006.ToObject(), πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label21
							}
							πE = πF.Raise(πTemp006.ToObject(), nil, πTemp007.ToObject())
							continue
							// line 389: except error, e:
							πF.SetLineno(389)
						Label21:
							µe = πTemp006.ToObject()
							πTemp003 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µe, "e"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, µe, ßargs, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, πTemp009, πTemp003); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßEINTR); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, πTemp008, πTemp003); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label22
							}
							goto Label23
							// line 390: if e.args[0] == EINTR:
							πF.SetLineno(390)
						Label22:
							// line 391: continue
							πF.SetLineno(391)
							continue
							goto Label23
						Label23:
							// line 392: raise
							πF.SetLineno(392)
							πE = πF.Raise(nil, nil, nil)
							continue
							πF.RestoreExc(nil, nil)
							goto Label19
						Label19:
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, µdata); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp005).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label24
							}
							goto Label25
							// line 393: if not data:
							πF.SetLineno(393)
						Label24:
							// line 394: break
							πF.SetLineno(394)
							πTemp004 = true
							continue
							goto Label25
						Label25:
							// line 395: n = len(data)
							πF.SetLineno(395)
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
							µn = πTemp003
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µsize, "size"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Eq(πF, µn, µsize); πE != nil {
								continue
							}
							πTemp002 = πTemp003
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp005 {
								goto Label26
							}
							if πE = πg.CheckLocal(πF, µbuf_len, "buf_len"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.IsTrue(πF, µbuf_len); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp010).ToObject()
							πTemp002 = πTemp003
						Label26:
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label27
							}
							goto Label28
							// line 396: if n == size and not buf_len:
							πF.SetLineno(396)
						Label27:
							// line 402: return data
							πF.SetLineno(402)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πR = µdata
							continue
							goto Label28
						Label28:
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µleft, "left"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, µn, µleft); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label29
							}
							goto Label30
							// line 403: if n == left:
							πF.SetLineno(403)
						Label29:
							// line 404: buf.write(data)
							πF.SetLineno(404)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πTemp001[0] = µdata
							if πE = πg.CheckLocal(πF, µbuf, "buf"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µbuf, ßwrite, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 405: del data  # explicit free
							πF.SetLineno(405)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							µdata = πg.UnboundLocal
							// line 406: break
							πF.SetLineno(406)
							πTemp004 = true
							continue
							goto Label30
						Label30:
							// line 407: assert n <= left, "recv(%d) returned %d bytes" % (left, n)
							πF.SetLineno(407)
							if πE = πg.CheckLocal(πF, µleft, "left"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple2(µleft, µn).ToObject()
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("recv(%d) returned %d bytes").ToObject(), πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µleft, "left"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.LE(πF, µn, µleft); πE != nil {
								continue
							}
							if πE = πg.Assert(πF, πTemp003, πTemp002); πE != nil {
								continue
							}
							// line 408: buf.write(data)
							πF.SetLineno(408)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πTemp001[0] = µdata
							if πE = πg.CheckLocal(πF, µbuf, "buf"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µbuf, ßwrite, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 409: buf_len += n
							πF.SetLineno(409)
							if πE = πg.CheckLocal(πF, µbuf_len, "buf_len"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IAdd(πF, µbuf_len, µn); πE != nil {
								continue
							}
							µbuf_len = πTemp002
							// line 410: del data  # explicit free
							πF.SetLineno(410)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							µdata = πg.UnboundLocal
							continue
						Label17:
							if πE != nil || πR != nil {
								continue
							}
						Label18:
							// line 412: return buf.getvalue()
							πF.SetLineno(412)
							if πE = πg.CheckLocal(πF, µbuf, "buf"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µbuf, ßgetvalue, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßread.ToObject(), πTemp012); πE != nil {
						continue
					}
					// line 414: def readline(self, size=-1):
					πF.SetLineno(414)
					πTemp003 = make([]πg.Param, 2)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					if πTemp014, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp003[1] = πg.Param{Name: "size", Def: πTemp014}
					πTemp013 = πg.NewFunction(πg.NewCode("readline", "/usr/lib/python2.7/socket.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µsize *πg.Object = πArgs[1]; _ = µsize
						var µbuf *πg.Object = πg.UnboundLocal; _ = µbuf
						var µbline *πg.Object = πg.UnboundLocal; _ = µbline
						var µbuffers *πg.Object = πg.UnboundLocal; _ = µbuffers
						var µdata *πg.Object = πg.UnboundLocal; _ = µdata
						var µrecv *πg.Object = πg.UnboundLocal; _ = µrecv
						var µe *πg.Object = πg.UnboundLocal; _ = µe
						var µnl *πg.Object = πg.UnboundLocal; _ = µnl
						var µbuf_len *πg.Object = πg.UnboundLocal; _ = µbuf_len
						var µrv *πg.Object = πg.UnboundLocal; _ = µrv
						var µleft *πg.Object = πg.UnboundLocal; _ = µleft
						var µn *πg.Object = πg.UnboundLocal; _ = µn
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
						var πTemp007 bool
						_ = πTemp007
						var πTemp008 bool
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
							case 39: goto Label39
							case 42: goto Label42
							case 11: goto Label11
							case 12: goto Label12
							case 15: goto Label15
							case 16: goto Label16
							case 17: goto Label17
							case 24: goto Label24
							case 25: goto Label25
							case 28: goto Label28
							default: panic("unexpected function state")
							}
							// line 415: buf = self._rbuf
							πF.SetLineno(415)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_rbuf, nil); πE != nil {
								continue
							}
							µbuf = πTemp001
							// line 416: buf.seek(0, 2)  # seek end
							πF.SetLineno(416)
							πTemp002 = πF.MakeArgs(2)
							πTemp002[0] = πg.NewInt(0).ToObject()
							πTemp002[1] = πg.NewInt(2).ToObject()
							if πE = πg.CheckLocal(πF, µbuf, "buf"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µbuf, ßseek, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πE = πg.CheckLocal(πF, µbuf, "buf"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µbuf, ßtell, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GT(πF, πTemp004, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							goto Label2
							// line 417: if buf.tell() > 0:
							πF.SetLineno(417)
						Label1:
							// line 419: buf.seek(0)
							πF.SetLineno(419)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µbuf, "buf"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µbuf, ßseek, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 420: bline = buf.readline(size)
							πF.SetLineno(420)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µsize, "size"); πE != nil {
								continue
							}
							πTemp002[0] = µsize
							if πE = πg.CheckLocal(πF, µbuf, "buf"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µbuf, ßreadline, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µbline = πTemp003
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("\n").ToObject()
							if πE = πg.CheckLocal(πF, µbline, "bline"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µbline, ßendswith, nil); πE != nil {
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
							if πTemp005 {
								goto Label3
							}
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µbline, "bline"); πE != nil {
								continue
							}
							πTemp002[0] = µbline
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πE = πg.CheckLocal(πF, µsize, "size"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Eq(πF, πTemp006, µsize); πE != nil {
								continue
							}
							πTemp001 = πTemp003
						Label3:
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label4
							}
							goto Label5
							// line 421: if bline.endswith('\n') or len(bline) == size:
							πF.SetLineno(421)
						Label4:
							// line 422: self._rbuf = StringIO()
							πF.SetLineno(422)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßStringIO); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_rbuf, πTemp001); πE != nil {
								continue
							}
							// line 423: self._rbuf.write(buf.read())
							πF.SetLineno(423)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µbuf, "buf"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µbuf, ßread, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_rbuf, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßwrite, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 424: return bline
							πF.SetLineno(424)
							if πE = πg.CheckLocal(πF, µbline, "bline"); πE != nil {
								continue
							}
							πR = µbline
							continue
							goto Label5
						Label5:
							// line 425: del bline
							πF.SetLineno(425)
							if πE = πg.CheckLocal(πF, µbline, "bline"); πE != nil {
								continue
							}
							µbline = πg.UnboundLocal
							goto Label2
						Label2:
							if πE = πg.CheckLocal(πF, µsize, "size"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.LT(πF, µsize, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label6
							}
							goto Label7
							// line 426: if size < 0:
							πF.SetLineno(426)
						Label6:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_rbufsize, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.LE(πF, πTemp003, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label9
							}
							goto Label10
							// line 428: if self._rbufsize <= 1:
							πF.SetLineno(428)
						Label9:
							// line 430: buf.seek(0)
							πF.SetLineno(430)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µbuf, "buf"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µbuf, ßseek, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 431: buffers = [buf.read()]
							πF.SetLineno(431)
							πTemp002 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µbuf, "buf"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µbuf, ßread, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							πTemp001 = πg.NewList(πTemp002...).ToObject()
							µbuffers = πTemp001
							// line 432: self._rbuf = StringIO()  # reset _rbuf.  we consume it via buf.
							πF.SetLineno(432)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßStringIO); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_rbuf, πTemp001); πE != nil {
								continue
							}
							// line 433: data = None
							πF.SetLineno(433)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							µdata = πTemp001
							// line 434: recv = self._sock.recv
							πF.SetLineno(434)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_sock, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßrecv, nil); πE != nil {
								continue
							}
							µrecv = πTemp003
							// line 435: while True:
							πF.SetLineno(435)
							πF.PushCheckpoint(12)
							πTemp005 = false
						Label11:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp005 {
								πF.PopCheckpoint()
								goto Label13
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
							πF.PushCheckpoint(11)            
							// line 436: try:
							πF.SetLineno(436)
							πF.PushCheckpoint(15)
							// line 437: while data != "\n":
							πF.SetLineno(437)
							πF.PushCheckpoint(17)
							πTemp007 = false
						Label16:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp007 {
								πF.PopCheckpoint()
								goto Label18
							}
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.NE(πF, µdata, πg.NewStr("\n").ToObject()); πE != nil {
								continue
							}
							if πTemp008, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πE != nil || !πTemp008 {
								continue
							}
							πF.PushCheckpoint(16)            
							// line 438: data = recv(1)
							πF.SetLineno(438)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µrecv, "recv"); πE != nil {
								continue
							}
							if πTemp001, πE = µrecv.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µdata = πTemp001
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.IsTrue(πF, µdata); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp008).ToObject()
							if πTemp008, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label19
							}
							goto Label20
							// line 439: if not data:
							πF.SetLineno(439)
						Label19:
							// line 440: break
							πF.SetLineno(440)
							πTemp007 = true
							continue
							goto Label20
						Label20:
							// line 441: buffers.append(data)
							πF.SetLineno(441)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πTemp002[0] = µdata
							if πE = πg.CheckLocal(πF, µbuffers, "buffers"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µbuffers, ßappend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							continue
						Label17:
							if πE != nil || πR != nil {
								continue
							}
						Label18:
							πF.PopCheckpoint()
							goto Label14
						Label15:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp009, πTemp010 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsInstance(πF, πTemp009.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label21
							}
							πE = πF.Raise(πTemp009.ToObject(), nil, πTemp010.ToObject())
							continue
							// line 442: except error, e:
							πF.SetLineno(442)
						Label21:
							µe = πTemp009.ToObject()
							πTemp003 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µe, "e"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µe, ßargs, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, πTemp006, πTemp003); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßEINTR); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, πTemp004, πTemp003); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label22
							}
							goto Label23
							// line 445: if e.args[0] == EINTR:
							πF.SetLineno(445)
						Label22:
							// line 446: continue
							πF.SetLineno(446)
							continue
							goto Label23
						Label23:
							// line 447: raise
							πF.SetLineno(447)
							πE = πF.Raise(nil, nil, nil)
							continue
							πF.RestoreExc(nil, nil)
							goto Label14
						Label14:
							// line 448: break
							πF.SetLineno(448)
							πTemp005 = true
							continue
							continue
						Label12:
							if πE != nil || πR != nil {
								continue
							}
						Label13:
							// line 449: return "".join(buffers)
							πF.SetLineno(449)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µbuffers, "buffers"); πE != nil {
								continue
							}
							πTemp002[0] = µbuffers
							if πTemp001, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πR = πTemp003
							continue
							goto Label10
						Label10:
							// line 451: buf.seek(0, 2)  # seek end
							πF.SetLineno(451)
							πTemp002 = πF.MakeArgs(2)
							πTemp002[0] = πg.NewInt(0).ToObject()
							πTemp002[1] = πg.NewInt(2).ToObject()
							if πE = πg.CheckLocal(πF, µbuf, "buf"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µbuf, ßseek, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 452: self._rbuf = StringIO()  # reset _rbuf.  we consume it via buf.
							πF.SetLineno(452)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßStringIO); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_rbuf, πTemp001); πE != nil {
								continue
							}
							// line 453: while True:
							πF.SetLineno(453)
							πF.PushCheckpoint(25)
							πTemp005 = false
						Label24:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp005 {
								πF.PopCheckpoint()
								goto Label26
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
							πF.PushCheckpoint(24)            
							// line 454: try:
							πF.SetLineno(454)
							πF.PushCheckpoint(28)
							// line 455: data = self._sock.recv(self._rbufsize)
							πF.SetLineno(455)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_rbufsize, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_sock, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßrecv, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µdata = πTemp001
							πF.PopCheckpoint()
							goto Label27
						Label28:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp009, πTemp010 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsInstance(πF, πTemp009.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label29
							}
							πE = πF.Raise(πTemp009.ToObject(), nil, πTemp010.ToObject())
							continue
							// line 456: except error, e:
							πF.SetLineno(456)
						Label29:
							µe = πTemp009.ToObject()
							πTemp003 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µe, "e"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µe, ßargs, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, πTemp006, πTemp003); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßEINTR); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, πTemp004, πTemp003); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label30
							}
							goto Label31
							// line 457: if e.args[0] == EINTR:
							πF.SetLineno(457)
						Label30:
							// line 458: continue
							πF.SetLineno(458)
							continue
							goto Label31
						Label31:
							// line 459: raise
							πF.SetLineno(459)
							πE = πF.Raise(nil, nil, nil)
							continue
							πF.RestoreExc(nil, nil)
							goto Label27
						Label27:
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, µdata); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp007).ToObject()
							if πTemp007, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label32
							}
							goto Label33
							// line 460: if not data:
							πF.SetLineno(460)
						Label32:
							// line 461: break
							πF.SetLineno(461)
							πTemp005 = true
							continue
							goto Label33
						Label33:
							// line 462: nl = data.find('\n')
							πF.SetLineno(462)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("\n").ToObject()
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µdata, ßfind, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µnl = πTemp003
							if πE = πg.CheckLocal(πF, µnl, "nl"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GE(πF, µnl, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label34
							}
							goto Label35
							// line 463: if nl >= 0:
							πF.SetLineno(463)
						Label34:
							// line 464: nl += 1
							πF.SetLineno(464)
							if πE = πg.CheckLocal(πF, µnl, "nl"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.IAdd(πF, µnl, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µnl = πTemp001
							// line 465: buf.write(data[:nl])
							πF.SetLineno(465)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnl, "nl"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.None, µnl, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µdata, πTemp001); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µbuf, "buf"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µbuf, ßwrite, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 466: self._rbuf.write(data[nl:])
							πF.SetLineno(466)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnl, "nl"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{µnl, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µdata, πTemp001); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_rbuf, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßwrite, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 467: del data
							πF.SetLineno(467)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							µdata = πg.UnboundLocal
							// line 468: break
							πF.SetLineno(468)
							πTemp005 = true
							continue
							goto Label35
						Label35:
							// line 469: buf.write(data)
							πF.SetLineno(469)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πTemp002[0] = µdata
							if πE = πg.CheckLocal(πF, µbuf, "buf"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µbuf, ßwrite, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							continue
						Label25:
							if πE != nil || πR != nil {
								continue
							}
						Label26:
							// line 470: return buf.getvalue()
							πF.SetLineno(470)
							if πE = πg.CheckLocal(πF, µbuf, "buf"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µbuf, ßgetvalue, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πR = πTemp003
							continue
							goto Label8
						Label7:
							// line 473: buf.seek(0, 2)  # seek end
							πF.SetLineno(473)
							πTemp002 = πF.MakeArgs(2)
							πTemp002[0] = πg.NewInt(0).ToObject()
							πTemp002[1] = πg.NewInt(2).ToObject()
							if πE = πg.CheckLocal(πF, µbuf, "buf"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µbuf, ßseek, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 474: buf_len = buf.tell()
							πF.SetLineno(474)
							if πE = πg.CheckLocal(πF, µbuf, "buf"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µbuf, ßtell, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µbuf_len = πTemp003
							if πE = πg.CheckLocal(πF, µbuf_len, "buf_len"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µsize, "size"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GE(πF, µbuf_len, µsize); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label36
							}
							goto Label37
							// line 475: if buf_len >= size:
							πF.SetLineno(475)
						Label36:
							// line 476: buf.seek(0)
							πF.SetLineno(476)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µbuf, "buf"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µbuf, ßseek, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 477: rv = buf.read(size)
							πF.SetLineno(477)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µsize, "size"); πE != nil {
								continue
							}
							πTemp002[0] = µsize
							if πE = πg.CheckLocal(πF, µbuf, "buf"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µbuf, ßread, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µrv = πTemp003
							// line 478: self._rbuf = StringIO()
							πF.SetLineno(478)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßStringIO); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_rbuf, πTemp001); πE != nil {
								continue
							}
							// line 479: self._rbuf.write(buf.read())
							πF.SetLineno(479)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µbuf, "buf"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µbuf, ßread, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_rbuf, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßwrite, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 480: return rv
							πF.SetLineno(480)
							if πE = πg.CheckLocal(πF, µrv, "rv"); πE != nil {
								continue
							}
							πR = µrv
							continue
							goto Label37
						Label37:
							// line 481: self._rbuf = StringIO()  # reset _rbuf.  we consume it via buf.
							πF.SetLineno(481)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßStringIO); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_rbuf, πTemp001); πE != nil {
								continue
							}
							// line 482: while True:
							πF.SetLineno(482)
							πF.PushCheckpoint(39)
							πTemp005 = false
						Label38:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp005 {
								πF.PopCheckpoint()
								goto Label40
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
							πF.PushCheckpoint(38)            
							// line 483: try:
							πF.SetLineno(483)
							πF.PushCheckpoint(42)
							// line 484: data = self._sock.recv(self._rbufsize)
							πF.SetLineno(484)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_rbufsize, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_sock, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßrecv, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µdata = πTemp001
							πF.PopCheckpoint()
							goto Label41
						Label42:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp009, πTemp010 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
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
							// line 485: except error, e:
							πF.SetLineno(485)
						Label43:
							µe = πTemp009.ToObject()
							πTemp003 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µe, "e"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µe, ßargs, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, πTemp006, πTemp003); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßEINTR); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, πTemp004, πTemp003); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label44
							}
							goto Label45
							// line 486: if e.args[0] == EINTR:
							πF.SetLineno(486)
						Label44:
							// line 487: continue
							πF.SetLineno(487)
							continue
							goto Label45
						Label45:
							// line 488: raise
							πF.SetLineno(488)
							πE = πF.Raise(nil, nil, nil)
							continue
							πF.RestoreExc(nil, nil)
							goto Label41
						Label41:
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, µdata); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp007).ToObject()
							if πTemp007, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label46
							}
							goto Label47
							// line 489: if not data:
							πF.SetLineno(489)
						Label46:
							// line 490: break
							πF.SetLineno(490)
							πTemp005 = true
							continue
							goto Label47
						Label47:
							// line 491: left = size - buf_len
							πF.SetLineno(491)
							if πE = πg.CheckLocal(πF, µsize, "size"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µbuf_len, "buf_len"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Sub(πF, µsize, µbuf_len); πE != nil {
								continue
							}
							µleft = πTemp001
							// line 493: nl = data.find('\n', 0, left)
							πF.SetLineno(493)
							πTemp002 = πF.MakeArgs(3)
							πTemp002[0] = πg.NewStr("\n").ToObject()
							πTemp002[1] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µleft, "left"); πE != nil {
								continue
							}
							πTemp002[2] = µleft
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µdata, ßfind, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µnl = πTemp003
							if πE = πg.CheckLocal(πF, µnl, "nl"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GE(πF, µnl, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label48
							}
							goto Label49
							// line 494: if nl >= 0:
							πF.SetLineno(494)
						Label48:
							// line 495: nl += 1
							πF.SetLineno(495)
							if πE = πg.CheckLocal(πF, µnl, "nl"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.IAdd(πF, µnl, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µnl = πTemp001
							// line 497: self._rbuf.write(data[nl:])
							πF.SetLineno(497)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnl, "nl"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{µnl, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µdata, πTemp001); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_rbuf, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßwrite, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πE = πg.CheckLocal(πF, µbuf_len, "buf_len"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, µbuf_len); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label50
							}
							goto Label51
							// line 498: if buf_len:
							πF.SetLineno(498)
						Label50:
							// line 499: buf.write(data[:nl])
							πF.SetLineno(499)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnl, "nl"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.None, µnl, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µdata, πTemp001); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µbuf, "buf"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µbuf, ßwrite, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 500: break
							πF.SetLineno(500)
							πTemp005 = true
							continue
							goto Label52
						Label51:
							// line 504: return data[:nl]
							πF.SetLineno(504)
							if πE = πg.CheckLocal(πF, µnl, "nl"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.None, µnl, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µdata, πTemp001); πE != nil {
								continue
							}
							πR = πTemp003
							continue
							goto Label52
						Label52:
							goto Label49
						Label49:
							// line 505: n = len(data)
							πF.SetLineno(505)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πTemp002[0] = µdata
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
							if πE = πg.CheckLocal(πF, µsize, "size"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Eq(πF, µn, µsize); πE != nil {
								continue
							}
							πTemp001 = πTemp003
							if πTemp007, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp007 {
								goto Label53
							}
							if πE = πg.CheckLocal(πF, µbuf_len, "buf_len"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.IsTrue(πF, µbuf_len); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp008).ToObject()
							πTemp001 = πTemp003
						Label53:
							if πTemp007, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label54
							}
							goto Label55
							// line 506: if n == size and not buf_len:
							πF.SetLineno(506)
						Label54:
							// line 509: return data
							πF.SetLineno(509)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πR = µdata
							continue
							goto Label55
						Label55:
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µleft, "left"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GE(πF, µn, µleft); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label56
							}
							goto Label57
							// line 510: if n >= left:
							πF.SetLineno(510)
						Label56:
							// line 511: buf.write(data[:left])
							πF.SetLineno(511)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µleft, "left"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.None, µleft, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µdata, πTemp001); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µbuf, "buf"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µbuf, ßwrite, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 512: self._rbuf.write(data[left:])
							πF.SetLineno(512)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µleft, "left"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{µleft, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µdata, πTemp001); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_rbuf, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßwrite, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 513: break
							πF.SetLineno(513)
							πTemp005 = true
							continue
							goto Label57
						Label57:
							// line 514: buf.write(data)
							πF.SetLineno(514)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πTemp002[0] = µdata
							if πE = πg.CheckLocal(πF, µbuf, "buf"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µbuf, ßwrite, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 515: buf_len += n
							πF.SetLineno(515)
							if πE = πg.CheckLocal(πF, µbuf_len, "buf_len"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.IAdd(πF, µbuf_len, µn); πE != nil {
								continue
							}
							µbuf_len = πTemp001
							continue
						Label39:
							if πE != nil || πR != nil {
								continue
							}
						Label40:
							// line 517: return buf.getvalue()
							πF.SetLineno(517)
							if πE = πg.CheckLocal(πF, µbuf, "buf"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µbuf, ßgetvalue, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
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
					if πE = πClass.SetItem(πF, ßreadline.ToObject(), πTemp013); πE != nil {
						continue
					}
					// line 519: def readlines(self, sizehint=0):
					πF.SetLineno(519)
					πTemp003 = make([]πg.Param, 2)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp003[1] = πg.Param{Name: "sizehint", Def: πg.NewInt(0).ToObject()}
					πTemp014 = πg.NewFunction(πg.NewCode("readlines", "/usr/lib/python2.7/socket.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µsizehint *πg.Object = πArgs[1]; _ = µsizehint
						var µtotal *πg.Object = πg.UnboundLocal; _ = µtotal
						var µlist *πg.Object = πg.UnboundLocal; _ = µlist
						var µline *πg.Object = πg.UnboundLocal; _ = µline
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							// line 520: total = 0
							πF.SetLineno(520)
							µtotal = πg.NewInt(0).ToObject()
							// line 521: list = []
							πF.SetLineno(521)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µlist = πTemp002
							// line 522: while True:
							πF.SetLineno(522)
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
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πE != nil || !πTemp004 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 523: line = self.readline()
							πF.SetLineno(523)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßreadline, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µline = πTemp005
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µline); πE != nil {
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
							// line 524: if not line:
							πF.SetLineno(524)
						Label4:
							// line 525: break
							πF.SetLineno(525)
							πTemp003 = true
							continue
							goto Label5
						Label5:
							// line 526: list.append(line)
							πF.SetLineno(526)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							πTemp001[0] = µline
							if πE = πg.CheckLocal(πF, µlist, "list"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µlist, ßappend, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 527: total += len(line)
							πF.SetLineno(527)
							if πE = πg.CheckLocal(πF, µtotal, "total"); πE != nil {
								continue
							}
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							πTemp001[0] = µline
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.IAdd(πF, µtotal, πTemp005); πE != nil {
								continue
							}
							µtotal = πTemp002
							if πE = πg.CheckLocal(πF, µsizehint, "sizehint"); πE != nil {
								continue
							}
							πTemp002 = µsizehint
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp004 {
								goto Label6
							}
							if πE = πg.CheckLocal(πF, µtotal, "total"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µsizehint, "sizehint"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GE(πF, µtotal, µsizehint); πE != nil {
								continue
							}
							πTemp002 = πTemp005
						Label6:
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label7
							}
							goto Label8
							// line 528: if sizehint and total >= sizehint:
							πF.SetLineno(528)
						Label7:
							// line 529: break
							πF.SetLineno(529)
							πTemp003 = true
							continue
							goto Label8
						Label8:
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 530: return list
							πF.SetLineno(530)
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
					if πE = πClass.SetItem(πF, ßreadlines.ToObject(), πTemp014); πE != nil {
						continue
					}
					// line 534: def __iter__(self):
					πF.SetLineno(534)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp015 = πg.NewFunction(πg.NewCode("__iter__", "/usr/lib/python2.7/socket.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 535: return self
							πF.SetLineno(535)
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
					if πE = πClass.SetItem(πF, ß__iter__.ToObject(), πTemp015); πE != nil {
						continue
					}
					// line 537: def next(self):
					πF.SetLineno(537)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp016 = πg.NewFunction(πg.NewCode("next", "/usr/lib/python2.7/socket.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µline *πg.Object = πg.UnboundLocal; _ = µline
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
							// line 538: line = self.readline()
							πF.SetLineno(538)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßreadline, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µline = πTemp002
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, µline); πE != nil {
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
							// line 539: if not line:
							πF.SetLineno(539)
						Label1:
							if πTemp001, πE = πg.ResolveGlobal(πF, ßStopIteration); πE != nil {
								continue
							}
							// line 540: raise StopIteration
							πF.SetLineno(540)
							πE = πF.Raise(πTemp001, nil, nil)
							continue
							goto Label2
						Label2:
							// line 541: return line
							πF.SetLineno(541)
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							πR = µline
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßnext.ToObject(), πTemp016); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp007, πE = πTemp014.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp007 == nil {
				πTemp007 = πg.TypeType.ToObject()
			}
			if πTemp008, πE = πTemp007.Call(πF, []*πg.Object{πg.NewStr("_fileobject").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp014.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_fileobject.ToObject(), πTemp008); πE != nil {
				continue
			}
			// line 543: _GLOBAL_DEFAULT_TIMEOUT = object()
			πF.SetLineno(543)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			if πTemp007, πE = πTemp004.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_GLOBAL_DEFAULT_TIMEOUT.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 545: def create_connection(address, timeout=_GLOBAL_DEFAULT_TIMEOUT,
			πF.SetLineno(545)
			πTemp013 = make([]πg.Param, 3)
			πTemp013[0] = πg.Param{Name: "address", Def: nil}
			if πTemp007, πE = πg.ResolveGlobal(πF, ß_GLOBAL_DEFAULT_TIMEOUT); πE != nil {
				continue
			}
			πTemp013[1] = πg.Param{Name: "timeout", Def: πTemp007}
			if πTemp007, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp013[2] = πg.Param{Name: "source_address", Def: πTemp007}
			πTemp004 = πg.NewFunction(πg.NewCode("create_connection", "/usr/lib/python2.7/socket.py", πTemp013, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µaddress *πg.Object = πArgs[0]; _ = µaddress
				var µtimeout *πg.Object = πArgs[1]; _ = µtimeout
				var µsource_address *πg.Object = πArgs[2]; _ = µsource_address
				var µhost *πg.Object = πg.UnboundLocal; _ = µhost
				var µport *πg.Object = πg.UnboundLocal; _ = µport
				var µerr *πg.Object = πg.UnboundLocal; _ = µerr
				var µres *πg.Object = πg.UnboundLocal; _ = µres
				var µaf *πg.Object = πg.UnboundLocal; _ = µaf
				var µsocktype *πg.Object = πg.UnboundLocal; _ = µsocktype
				var µproto *πg.Object = πg.UnboundLocal; _ = µproto
				var µcanonname *πg.Object = πg.UnboundLocal; _ = µcanonname
				var µsa *πg.Object = πg.UnboundLocal; _ = µsa
				var µsock *πg.Object = πg.UnboundLocal; _ = µsock
				var µ_ *πg.Object = πg.UnboundLocal; _ = µ_
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
				var πTemp006 bool
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
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 1: goto Label1
					case 2: goto Label2
					case 5: goto Label5
					default: panic("unexpected function state")
					}
					// line 547: """Connect to *address* and return the socket object.
					πF.SetLineno(547)
					// line 559: host, port = address
					πF.SetLineno(559)
					if πE = πg.CheckLocal(πF, µaddress, "address"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp002}}}, µaddress); πE != nil {
						continue
					}
					µhost = πTemp001
					µport = πTemp002
					// line 560: err = None
					πF.SetLineno(560)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					µerr = πTemp001
					πTemp003 = πF.MakeArgs(4)
					if πE = πg.CheckLocal(πF, µhost, "host"); πE != nil {
						continue
					}
					πTemp003[0] = µhost
					if πE = πg.CheckLocal(πF, µport, "port"); πE != nil {
						continue
					}
					πTemp003[1] = µport
					πTemp003[2] = πg.NewInt(0).ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßSOCK_STREAM); πE != nil {
						continue
					}
					πTemp003[3] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßgetaddrinfo); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
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
						µres = πTemp002
					}
					if πE != nil || !πTemp006 {
						continue
					}
					πF.PushCheckpoint(1)            
					// line 562: af, socktype, proto, canonname, sa = res
					πF.SetLineno(562)
					if πE = πg.CheckLocal(πF, µres, "res"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}, πg.TieTarget{Target: &πTemp008}, πg.TieTarget{Target: &πTemp009}}}, µres); πE != nil {
						continue
					}
					µaf = πTemp002
					µsocktype = πTemp004
					µproto = πTemp007
					µcanonname = πTemp008
					µsa = πTemp009
					// line 563: sock = None
					πF.SetLineno(563)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					µsock = πTemp002
					// line 564: try:
					πF.SetLineno(564)
					πF.PushCheckpoint(5)
					// line 565: sock = socket(af, socktype, proto)
					πF.SetLineno(565)
					πTemp003 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µaf, "af"); πE != nil {
						continue
					}
					πTemp003[0] = µaf
					if πE = πg.CheckLocal(πF, µsocktype, "socktype"); πE != nil {
						continue
					}
					πTemp003[1] = µsocktype
					if πE = πg.CheckLocal(πF, µproto, "proto"); πE != nil {
						continue
					}
					πTemp003[2] = µproto
					if πTemp002, πE = πg.ResolveGlobal(πF, ßsocket); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µsock = πTemp004
					if πE = πg.CheckLocal(πF, µtimeout, "timeout"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ß_GLOBAL_DEFAULT_TIMEOUT); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(µtimeout != πTemp004).ToObject()
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label6
					}
					goto Label7
					// line 566: if timeout is not _GLOBAL_DEFAULT_TIMEOUT:
					πF.SetLineno(566)
				Label6:
					// line 567: sock.settimeout(timeout)
					πF.SetLineno(567)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µtimeout, "timeout"); πE != nil {
						continue
					}
					πTemp003[0] = µtimeout
					if πE = πg.CheckLocal(πF, µsock, "sock"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µsock, ßsettimeout, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					goto Label7
				Label7:
					if πE = πg.CheckLocal(πF, µsource_address, "source_address"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, µsource_address); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label8
					}
					goto Label9
					// line 568: if source_address:
					πF.SetLineno(568)
				Label8:
					// line 569: sock.bind(source_address)
					πF.SetLineno(569)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µsource_address, "source_address"); πE != nil {
						continue
					}
					πTemp003[0] = µsource_address
					if πE = πg.CheckLocal(πF, µsock, "sock"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µsock, ßbind, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					goto Label9
				Label9:
					// line 570: sock.connect(sa)
					πF.SetLineno(570)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µsa, "sa"); πE != nil {
						continue
					}
					πTemp003[0] = µsa
					if πE = πg.CheckLocal(πF, µsock, "sock"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µsock, ßconnect, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					// line 571: return sock
					πF.SetLineno(571)
					if πE = πg.CheckLocal(πF, µsock, "sock"); πE != nil {
						continue
					}
					πR = µsock
					continue
					πF.PopCheckpoint()
					goto Label4
				Label5:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp010, πTemp011 = πF.ExcInfo()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsInstance(πF, πTemp010.ToObject(), πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label10
					}
					πE = πF.Raise(πTemp010.ToObject(), nil, πTemp011.ToObject())
					continue
					// line 573: except error as _:
					πF.SetLineno(573)
				Label10:
					µ_ = πTemp010.ToObject()
					// line 574: err = _
					πF.SetLineno(574)
					if πE = πg.CheckLocal(πF, µ_, "_"); πE != nil {
						continue
					}
					µerr = µ_
					if πE = πg.CheckLocal(πF, µsock, "sock"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(µsock != πTemp004).ToObject()
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label11
					}
					goto Label12
					// line 575: if sock is not None:
					πF.SetLineno(575)
				Label11:
					// line 576: sock.close()
					πF.SetLineno(576)
					if πE = πg.CheckLocal(πF, µsock, "sock"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µsock, ßclose, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					goto Label12
				Label12:
					πF.RestoreExc(nil, nil)
					goto Label4
				Label4:
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
					if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µerr != πTemp002).ToObject()
					if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label13
					}
					goto Label14
					// line 578: if err is not None:
					πF.SetLineno(578)
				Label13:
					if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
						continue
					}
					// line 579: raise err
					πF.SetLineno(579)
					πE = πF.Raise(µerr, nil, nil)
					continue
					goto Label15
				Label14:
					πTemp003 = πF.MakeArgs(1)
					πTemp003[0] = πg.NewStr("getaddrinfo returns an empty list").ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					// line 581: raise error("getaddrinfo returns an empty list")
					πF.SetLineno(581)
					πE = πF.Raise(πTemp002, nil, nil)
					continue
					goto Label15
				Label15:
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßcreate_connection.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 547: """Connect to *address* and return the socket object.
			πF.SetLineno(547)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp007}, πg.NewStr("Connect to *address* and return the socket object.\n\n    Convenience function.  Connect to *address* (a 2-tuple ``(host,\n    port)``) and return the socket object.  Passing the optional\n    *timeout* parameter will set the timeout on the socket instance\n    before attempting to connect.  If no *timeout* is supplied, the\n    global default timeout setting returned by :func:`getdefaulttimeout`\n    is used.  If *source_address* is set it must be a tuple of (host, port)\n    for the socket to bind as a source address before making the connection.\n    A host of '' or port 0 tells the OS to use the default.\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp008, πE = πg.ResolveGlobal(πF, ßcreate_connection); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp008, ß__doc__, πTemp007); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("socket", Code)
}

