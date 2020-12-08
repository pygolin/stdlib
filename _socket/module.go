package _socket
import (
	πg "github.com/pygolin/runtime"
	// _ "github.com/pygolin/stdlib/__go__/syscall"
	// _ "github.com/pygolin/stdlib/__go__/net"
	// _ "github.com/pygolin/stdlib/math"
)
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/_socket.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ß := πg.InternStr("")
		ßAF_INET := πg.InternStr("AF_INET")
		ßAF_INET6 := πg.InternStr("AF_INET6")
		ßAF_UNIX := πg.InternStr("AF_UNIX")
		ßAccept := πg.InternStr("Accept")
		ßAddr := πg.InternStr("Addr")
		ßBind := πg.InternStr("Bind")
		ßClose := πg.InternStr("Close")
		ßConnect := πg.InternStr("Connect")
		ßError := πg.InternStr("Error")
		ßGetpeername := πg.InternStr("Getpeername")
		ßGetsockname := πg.InternStr("Getsockname")
		ßIOError := πg.InternStr("IOError")
		ßIP := πg.InternStr("IP")
		ßIPv4 := πg.InternStr("IPv4")
		ßListen := πg.InternStr("Listen")
		ßLookupAddr := πg.InternStr("LookupAddr")
		ßLookupIP := πg.InternStr("LookupIP")
		ßName := πg.InternStr("Name")
		ßNone := πg.InternStr("None")
		ßNotImplementedError := πg.InternStr("NotImplementedError")
		ßParseIP := πg.InternStr("ParseIP")
		ßPort := πg.InternStr("Port")
		ßRecvfrom := πg.InternStr("Recvfrom")
		ßSHUT_RD := πg.InternStr("SHUT_RD")
		ßSHUT_RDWR := πg.InternStr("SHUT_RDWR")
		ßSHUT_WR := πg.InternStr("SHUT_WR")
		ßSOCK_DGRAM := πg.InternStr("SOCK_DGRAM")
		ßSOCK_RAW := πg.InternStr("SOCK_RAW")
		ßSOCK_SEQPACKET := πg.InternStr("SOCK_SEQPACKET")
		ßSOCK_STREAM := πg.InternStr("SOCK_STREAM")
		ßSOL_SOCKET := πg.InternStr("SOL_SOCKET")
		ßSO_RCVTIMEO := πg.InternStr("SO_RCVTIMEO")
		ßSO_REUSEADDR := πg.InternStr("SO_REUSEADDR")
		ßSec := πg.InternStr("Sec")
		ßSendto := πg.InternStr("Sendto")
		ßSetNonblock := πg.InternStr("SetNonblock")
		ßSetsockoptInt := πg.InternStr("SetsockoptInt")
		ßSetsockoptTimeval := πg.InternStr("SetsockoptTimeval")
		ßShutdown := πg.InternStr("Shutdown")
		ßSockaddrInet4 := πg.InternStr("SockaddrInet4")
		ßSockaddrInet6 := πg.InternStr("SockaddrInet6")
		ßSockaddrUnix := πg.InternStr("SockaddrUnix")
		ßSocket := πg.InternStr("Socket")
		ßString := πg.InternStr("String")
		ßTimespec := πg.InternStr("Timespec")
		ßTo4 := πg.InternStr("To4")
		ßTo6 := πg.InternStr("To6")
		ßUsec := πg.InternStr("Usec")
		ß_Accept := πg.InternStr("_Accept")
		ß_Bind := πg.InternStr("_Bind")
		ß_Close := πg.InternStr("_Close")
		ß_Connect := πg.InternStr("_Connect")
		ß_Getpeername := πg.InternStr("_Getpeername")
		ß_Getsockname := πg.InternStr("_Getsockname")
		ß_GetsockoptInt := πg.InternStr("_GetsockoptInt")
		ß_IP := πg.InternStr("_IP")
		ß_IPv4 := πg.InternStr("_IPv4")
		ß_Listen := πg.InternStr("_Listen")
		ß_LookupAddr := πg.InternStr("_LookupAddr")
		ß_LookupIP := πg.InternStr("_LookupIP")
		ß_ParseIP := πg.InternStr("_ParseIP")
		ß_Recvfrom := πg.InternStr("_Recvfrom")
		ß_Sendto := πg.InternStr("_Sendto")
		ß_SetNonblock := πg.InternStr("_SetNonblock")
		ß_SetsockoptInt := πg.InternStr("_SetsockoptInt")
		ß_SetsockoptTimeval := πg.InternStr("_SetsockoptTimeval")
		ß_Shutdown := πg.InternStr("_Shutdown")
		ß_SockaddrInet4 := πg.InternStr("_SockaddrInet4")
		ß_SockaddrInet6 := πg.InternStr("_SockaddrInet6")
		ß_SockaddrUnix := πg.InternStr("_SockaddrUnix")
		ß_Socket := πg.InternStr("_Socket")
		ß_Timespec := πg.InternStr("_Timespec")
		ß_Timeval := πg.InternStr("_Timeval")
		ß__init__ := πg.InternStr("__init__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß_fd := πg.InternStr("_fd")
		ß_get_address := πg.InternStr("_get_address")
		ß_parse_address := πg.InternStr("_parse_address")
		ßaccept := πg.InternStr("accept")
		ßbind := πg.InternStr("bind")
		ßbytearray := πg.InternStr("bytearray")
		ßchr := πg.InternStr("chr")
		ßclose := πg.InternStr("close")
		ßconnect := πg.InternStr("connect")
		ßconnect_ex := πg.InternStr("connect_ex")
		ßerror := πg.InternStr("error")
		ßfamily := πg.InternStr("family")
		ßfd := πg.InternStr("fd")
		ßfileno := πg.InternStr("fileno")
		ßfromfd := πg.InternStr("fromfd")
		ßgaierror := πg.InternStr("gaierror")
		ßgetaddrinfo := πg.InternStr("getaddrinfo")
		ßgetdefaulttimeout := πg.InternStr("getdefaulttimeout")
		ßgethostbyaddr := πg.InternStr("gethostbyaddr")
		ßgethostbyname := πg.InternStr("gethostbyname")
		ßgethostname := πg.InternStr("gethostname")
		ßgetnameinfo := πg.InternStr("getnameinfo")
		ßgetpeername := πg.InternStr("getpeername")
		ßgetprotobyname := πg.InternStr("getprotobyname")
		ßgetservbyname := πg.InternStr("getservbyname")
		ßgetservbyport := πg.InternStr("getservbyport")
		ßgetsockname := πg.InternStr("getsockname")
		ßgetsockopt := πg.InternStr("getsockopt")
		ßgettimeout := πg.InternStr("gettimeout")
		ßherror := πg.InternStr("herror")
		ßhtonl := πg.InternStr("htonl")
		ßhtons := πg.InternStr("htons")
		ßinet_aton := πg.InternStr("inet_aton")
		ßinet_ntoa := πg.InternStr("inet_ntoa")
		ßint := πg.InternStr("int")
		ßisinstance := πg.InternStr("isinstance")
		ßjoin := πg.InternStr("join")
		ßlen := πg.InternStr("len")
		ßlevel := πg.InternStr("level")
		ßlisten := πg.InternStr("listen")
		ßmath := πg.InternStr("math")
		ßmodf := πg.InternStr("modf")
		ßnew := πg.InternStr("new")
		ßntohl := πg.InternStr("ntohl")
		ßntohs := πg.InternStr("ntohs")
		ßobject := πg.InternStr("object")
		ßord := πg.InternStr("ord")
		ßproto := πg.InternStr("proto")
		ßrecv := πg.InternStr("recv")
		ßrecv_into := πg.InternStr("recv_into")
		ßrecvfrom := πg.InternStr("recvfrom")
		ßrecvfrom_into := πg.InternStr("recvfrom_into")
		ßsend := πg.InternStr("send")
		ßsendall := πg.InternStr("sendall")
		ßsendto := πg.InternStr("sendto")
		ßsetblocking := πg.InternStr("setblocking")
		ßsetdefaulttimeout := πg.InternStr("setdefaulttimeout")
		ßsetsockopt := πg.InternStr("setsockopt")
		ßsettimeout := πg.InternStr("settimeout")
		ßshutdown := πg.InternStr("shutdown")
		ßsocket := πg.InternStr("socket")
		ßsocketpair := πg.InternStr("socketpair")
		ßsplit := πg.InternStr("split")
		ßstr := πg.InternStr("str")
		ßtimeout := πg.InternStr("timeout")
		ßtype := πg.InternStr("type")
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
		var πTemp006 []πg.Param
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
			// line 15: from '__go__/net' import (
			πF.SetLineno(15)
			if πTemp002, πE = πg.ImportModule(πF, "__go__/net"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßIPv4); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_IPv4.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßIP); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_IP.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßLookupAddr); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_LookupAddr.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßLookupIP); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_LookupIP.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßParseIP); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_ParseIP.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 18: from '__go__/syscall' import (
			πF.SetLineno(18)
			if πTemp002, πE = πg.ImportModule(πF, "__go__/syscall"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßAF_INET); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßAF_INET.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßAF_INET6); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßAF_INET6.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßAF_UNIX); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßAF_UNIX.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßSHUT_RD); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßSHUT_RD.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßSHUT_RDWR); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßSHUT_RDWR.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßSHUT_WR); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßSHUT_WR.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßSO_REUSEADDR); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßSO_REUSEADDR.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßSOL_SOCKET); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßSOL_SOCKET.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßSOCK_DGRAM); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßSOCK_DGRAM.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßSOCK_RAW); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßSOCK_RAW.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßSOCK_SEQPACKET); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßSOCK_SEQPACKET.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßSOCK_STREAM); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßSOCK_STREAM.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßAccept); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_Accept.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßBind); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_Bind.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßClose); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_Close.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßConnect); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_Connect.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßGetpeername); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_Getpeername.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßGetsockname); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_Getsockname.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßListen); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_Listen.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßRecvfrom); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_Recvfrom.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßSendto); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_Sendto.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßSetNonblock); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_SetNonblock.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßSetsockoptInt); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_SetsockoptInt.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßSetsockoptTimeval); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_SetsockoptTimeval.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßShutdown); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_Shutdown.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßSockaddrInet4); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_SockaddrInet4.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßSockaddrInet6); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_SockaddrInet6.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßSockaddrUnix); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_SockaddrUnix.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßSocket); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_Socket.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßTimespec); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_Timespec.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 49: import math
			πF.SetLineno(49)
			if πTemp002, πE = πg.ImportModule(πF, "math"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßmath.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 77: class error(IOError):
			πF.SetLineno(77)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßIOError); πE != nil {
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
			_, πE = πg.NewCode("error", "/usr/lib/python2.7/_socket.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp004
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("error").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßerror.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 81: class gaierror(error):
			πF.SetLineno(81)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
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
			_, πE = πg.NewCode("gaierror", "/usr/lib/python2.7/_socket.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp004
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 82: pass
					πF.SetLineno(82)
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("gaierror").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßgaierror.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 85: class herror(error):
			πF.SetLineno(85)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
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
			_, πE = πg.NewCode("herror", "/usr/lib/python2.7/_socket.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp004
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 86: pass
					πF.SetLineno(86)
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("herror").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßherror.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 89: class timeout(error):
			πF.SetLineno(89)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
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
			_, πE = πg.NewCode("timeout", "/usr/lib/python2.7/_socket.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp004
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 90: pass
					πF.SetLineno(90)
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("timeout").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßtimeout.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 93: class socket(object):
			πF.SetLineno(93)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
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
			_, πE = πg.NewCode("socket", "/usr/lib/python2.7/_socket.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 95: def __init__(self, family=AF_INET, type=SOCK_STREAM, proto=0, fd=None):
					πF.SetLineno(95)
					πTemp002 = make([]πg.Param, 5)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßAF_INET); πE != nil {
						continue
					}
					πTemp002[1] = πg.Param{Name: "family", Def: πTemp003}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßSOCK_STREAM); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "type", Def: πTemp003}
					πTemp002[3] = πg.Param{Name: "proto", Def: πg.NewInt(0).ToObject()}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[4] = πg.Param{Name: "fd", Def: πTemp003}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/_socket.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µfamily *πg.Object = πArgs[1]; _ = µfamily
						var µtype *πg.Object = πArgs[2]; _ = µtype
						var µproto *πg.Object = πArgs[3]; _ = µproto
						var µfd *πg.Object = πArgs[4]; _ = µfd
						var µerr *πg.Object = πg.UnboundLocal; _ = µerr
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							if πE = πg.CheckLocal(πF, µfd, "fd"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µfd == πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 96: if fd is None:
							πF.SetLineno(96)
						Label1:
							// line 97: fd, err = _Socket(family, type, proto)
							πF.SetLineno(97)
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
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_Socket); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp005}}}, πTemp002); πE != nil {
								continue
							}
							µfd = πTemp001
							µerr = πTemp005
							if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, µerr); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label3
							}
							goto Label4
							// line 98: if err:
							πF.SetLineno(98)
						Label3:
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µerr, ßError, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp002
							if πTemp001, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 99: raise error(err.Error())
							πF.SetLineno(99)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							goto Label4
						Label4:
							goto Label2
						Label2:
							// line 100: self._fd = fd
							πF.SetLineno(100)
							if πE = πg.CheckLocal(πF, µfd, "fd"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µfd); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_fd, πTemp001); πE != nil {
								continue
							}
							// line 101: self.family = family
							πF.SetLineno(101)
							if πE = πg.CheckLocal(πF, µfamily, "family"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µfamily); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßfamily, πTemp001); πE != nil {
								continue
							}
							// line 102: self.type = type
							πF.SetLineno(102)
							if πE = πg.CheckLocal(πF, µtype, "type"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µtype); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßtype, πTemp001); πE != nil {
								continue
							}
							// line 103: self.proto = proto
							πF.SetLineno(103)
							if πE = πg.CheckLocal(πF, µproto, "proto"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µproto); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßproto, πTemp001); πE != nil {
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
					// line 105: def accept(self):
					πF.SetLineno(105)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("accept", "/usr/lib/python2.7/_socket.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µfd *πg.Object = πg.UnboundLocal; _ = µfd
						var µsockaddr *πg.Object = πg.UnboundLocal; _ = µsockaddr
						var µerr *πg.Object = πg.UnboundLocal; _ = µerr
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 106: fd, sockaddr, err = _Accept(self._fd)
							πF.SetLineno(106)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_fd, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_Accept); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}}}, πTemp003); πE != nil {
								continue
							}
							µfd = πTemp002
							µsockaddr = πTemp004
							µerr = πTemp005
							if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, µerr); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label1
							}
							goto Label2
							// line 107: if err:
							πF.SetLineno(107)
						Label1:
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µerr, ßError, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 108: raise error(err.Error())
							πF.SetLineno(108)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label2
						Label2:
							// line 109: return (socket(self.family, self.type, self.proto, fd),
							πF.SetLineno(109)
							πTemp001 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßfamily, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßproto, nil); πE != nil {
								continue
							}
							πTemp001[2] = πTemp003
							if πE = πg.CheckLocal(πF, µfd, "fd"); πE != nil {
								continue
							}
							πTemp001[3] = µfd
							if πTemp003, πE = πg.ResolveGlobal(πF, ßsocket); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µsockaddr, "sockaddr"); πE != nil {
								continue
							}
							πTemp001[0] = µsockaddr
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_get_address, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp002 = πg.NewTuple2(πTemp004, πTemp005).ToObject()
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
					if πE = πClass.SetItem(πF, ßaccept.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 112: def bind(self, address):
					πF.SetLineno(112)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "address", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("bind", "/usr/lib/python2.7/_socket.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µaddress *πg.Object = πArgs[1]; _ = µaddress
						var µsockaddr *πg.Object = πg.UnboundLocal; _ = µsockaddr
						var µerr *πg.Object = πg.UnboundLocal; _ = µerr
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
							// line 113: sockaddr = self._parse_address(address)
							πF.SetLineno(113)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µaddress, "address"); πE != nil {
								continue
							}
							πTemp001[0] = µaddress
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_parse_address, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µsockaddr = πTemp003
							// line 114: err = _Bind(self._fd, sockaddr)
							πF.SetLineno(114)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_fd, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µsockaddr, "sockaddr"); πE != nil {
								continue
							}
							πTemp001[1] = µsockaddr
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_Bind); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µerr = πTemp003
							if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µerr); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 115: if err:
							πF.SetLineno(115)
						Label1:
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µerr, ßError, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 116: raise error(err.Error())
							πF.SetLineno(116)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
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
					if πE = πClass.SetItem(πF, ßbind.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 118: def close(self):
					πF.SetLineno(118)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("close", "/usr/lib/python2.7/_socket.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 119: _Close(self._fd)
							πF.SetLineno(119)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_fd, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_Close); πE != nil {
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
					if πE = πClass.SetItem(πF, ßclose.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 121: def connect(self, address):
					πF.SetLineno(121)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "address", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("connect", "/usr/lib/python2.7/_socket.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µaddress *πg.Object = πArgs[1]; _ = µaddress
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
							// line 122: self.connect_ex(address)
							πF.SetLineno(122)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µaddress, "address"); πE != nil {
								continue
							}
							πTemp001[0] = µaddress
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßconnect_ex, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßconnect.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 124: def connect_ex(self, address):
					πF.SetLineno(124)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "address", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("connect_ex", "/usr/lib/python2.7/_socket.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µaddress *πg.Object = πArgs[1]; _ = µaddress
						var µaddr *πg.Object = πg.UnboundLocal; _ = µaddr
						var µerr *πg.Object = πg.UnboundLocal; _ = µerr
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
							// line 125: addr = self._parse_address(address)
							πF.SetLineno(125)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µaddress, "address"); πE != nil {
								continue
							}
							πTemp001[0] = µaddress
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_parse_address, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µaddr = πTemp003
							// line 126: err = _Connect(self._fd, addr)
							πF.SetLineno(126)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_fd, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µaddr, "addr"); πE != nil {
								continue
							}
							πTemp001[1] = µaddr
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_Connect); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µerr = πTemp003
							if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µerr); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 127: if err:
							πF.SetLineno(127)
						Label1:
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
								continue
							}
							πTemp001[0] = µerr
							if πTemp002, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 128: raise error(err)
							πF.SetLineno(128)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
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
					if πE = πClass.SetItem(πF, ßconnect_ex.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 130: def fileno(self):
					πF.SetLineno(130)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("fileno", "/usr/lib/python2.7/_socket.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 131: return self._fd
							πF.SetLineno(131)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_fd, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßfileno.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 133: def listen(self, backlog):
					πF.SetLineno(133)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "backlog", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("listen", "/usr/lib/python2.7/_socket.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µbacklog *πg.Object = πArgs[1]; _ = µbacklog
						var µerr *πg.Object = πg.UnboundLocal; _ = µerr
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
							// line 134: err = _Listen(self._fd, backlog)
							πF.SetLineno(134)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_fd, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µbacklog, "backlog"); πE != nil {
								continue
							}
							πTemp001[1] = µbacklog
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_Listen); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µerr = πTemp003
							if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µerr); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 135: if err:
							πF.SetLineno(135)
						Label1:
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µerr, ßError, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 136: raise error(err.Error())
							πF.SetLineno(136)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
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
					if πE = πClass.SetItem(πF, ßlisten.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 138: def getpeername(self):
					πF.SetLineno(138)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp010 = πg.NewFunction(πg.NewCode("getpeername", "/usr/lib/python2.7/_socket.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µsockaddr *πg.Object = πg.UnboundLocal; _ = µsockaddr
						var µerr *πg.Object = πg.UnboundLocal; _ = µerr
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
							// line 139: sockaddr, err = _Getpeername(self._fd)
							πF.SetLineno(139)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_fd, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_Getpeername); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp004}}}, πTemp003); πE != nil {
								continue
							}
							µsockaddr = πTemp002
							µerr = πTemp004
							if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, µerr); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							goto Label2
							// line 140: if err:
							πF.SetLineno(140)
						Label1:
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µerr, ßError, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 141: raise error(err.Error())
							πF.SetLineno(141)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label2
						Label2:
							// line 142: return self._get_address(sockaddr)
							πF.SetLineno(142)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µsockaddr, "sockaddr"); πE != nil {
								continue
							}
							πTemp001[0] = µsockaddr
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_get_address, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßgetpeername.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 144: def getsockname(self):
					πF.SetLineno(144)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp011 = πg.NewFunction(πg.NewCode("getsockname", "/usr/lib/python2.7/_socket.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µsockaddr *πg.Object = πg.UnboundLocal; _ = µsockaddr
						var µerr *πg.Object = πg.UnboundLocal; _ = µerr
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
							// line 145: sockaddr, err = _Getsockname(self._fd)
							πF.SetLineno(145)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_fd, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_Getsockname); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp004}}}, πTemp003); πE != nil {
								continue
							}
							µsockaddr = πTemp002
							µerr = πTemp004
							if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, µerr); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							goto Label2
							// line 146: if err:
							πF.SetLineno(146)
						Label1:
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µerr, ßError, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 147: raise error(err.Error())
							πF.SetLineno(147)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label2
						Label2:
							// line 148: return self._get_address(sockaddr)
							πF.SetLineno(148)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µsockaddr, "sockaddr"); πE != nil {
								continue
							}
							πTemp001[0] = µsockaddr
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_get_address, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßgetsockname.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 150: def getsockopt(self, level, optname, buflen=None):
					πF.SetLineno(150)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "level", Def: nil}
					πTemp002[2] = πg.Param{Name: "optname", Def: nil}
					if πTemp013, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[3] = πg.Param{Name: "buflen", Def: πTemp013}
					πTemp012 = πg.NewFunction(πg.NewCode("getsockopt", "/usr/lib/python2.7/_socket.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µlevel *πg.Object = πArgs[1]; _ = µlevel
						var µoptname *πg.Object = πArgs[2]; _ = µoptname
						var µbuflen *πg.Object = πArgs[3]; _ = µbuflen
						var µval *πg.Object = πg.UnboundLocal; _ = µval
						var µerr *πg.Object = πg.UnboundLocal; _ = µerr
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
							// line 151: val, err = _GetsockoptInt(self._fd, level, optname)
							πF.SetLineno(151)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_fd, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µlevel, "level"); πE != nil {
								continue
							}
							πTemp001[1] = µlevel
							if πE = πg.CheckLocal(πF, µoptname, "optname"); πE != nil {
								continue
							}
							πTemp001[2] = µoptname
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_GetsockoptInt); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp004}}}, πTemp003); πE != nil {
								continue
							}
							µval = πTemp002
							µerr = πTemp004
							if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, µerr); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							goto Label2
							// line 152: if err:
							πF.SetLineno(152)
						Label1:
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
								continue
							}
							πTemp001[0] = µerr
							if πTemp002, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 153: raise error(err)
							πF.SetLineno(153)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label2
						Label2:
							// line 154: return val
							πF.SetLineno(154)
							if πE = πg.CheckLocal(πF, µval, "val"); πE != nil {
								continue
							}
							πR = µval
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßgetsockopt.ToObject(), πTemp012); πE != nil {
						continue
					}
					// line 156: def recv(self, bufsize, flags=0):
					πF.SetLineno(156)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "bufsize", Def: nil}
					πTemp002[2] = πg.Param{Name: "flags", Def: πg.NewInt(0).ToObject()}
					πTemp013 = πg.NewFunction(πg.NewCode("recv", "/usr/lib/python2.7/_socket.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µbufsize *πg.Object = πArgs[1]; _ = µbufsize
						var µflags *πg.Object = πArgs[2]; _ = µflags
						var µdata *πg.Object = πg.UnboundLocal; _ = µdata
						var µ_ *πg.Object = πg.UnboundLocal; _ = µ_
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
							// line 157: data, _ = self.recvfrom(bufsize, flags)
							πF.SetLineno(157)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µbufsize, "bufsize"); πE != nil {
								continue
							}
							πTemp001[0] = µbufsize
							if πE = πg.CheckLocal(πF, µflags, "flags"); πE != nil {
								continue
							}
							πTemp001[1] = µflags
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßrecvfrom, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp004}}}, πTemp003); πE != nil {
								continue
							}
							µdata = πTemp002
							µ_ = πTemp004
							// line 158: return data
							πF.SetLineno(158)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πR = µdata
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßrecv.ToObject(), πTemp013); πE != nil {
						continue
					}
					// line 160: def recv_into(self, buffer, nbytes=0, flags=0):
					πF.SetLineno(160)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "buffer", Def: nil}
					πTemp002[2] = πg.Param{Name: "nbytes", Def: πg.NewInt(0).ToObject()}
					πTemp002[3] = πg.Param{Name: "flags", Def: πg.NewInt(0).ToObject()}
					πTemp014 = πg.NewFunction(πg.NewCode("recv_into", "/usr/lib/python2.7/_socket.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µbuffer *πg.Object = πArgs[1]; _ = µbuffer
						var µnbytes *πg.Object = πArgs[2]; _ = µnbytes
						var µflags *πg.Object = πArgs[3]; _ = µflags
						var µn *πg.Object = πg.UnboundLocal; _ = µn
						var µ_ *πg.Object = πg.UnboundLocal; _ = µ_
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
							// line 161: n, _ = self.recvfrom_into(buffer, nbytes, flags)
							πF.SetLineno(161)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µbuffer, "buffer"); πE != nil {
								continue
							}
							πTemp001[0] = µbuffer
							if πE = πg.CheckLocal(πF, µnbytes, "nbytes"); πE != nil {
								continue
							}
							πTemp001[1] = µnbytes
							if πE = πg.CheckLocal(πF, µflags, "flags"); πE != nil {
								continue
							}
							πTemp001[2] = µflags
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßrecvfrom_into, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp004}}}, πTemp003); πE != nil {
								continue
							}
							µn = πTemp002
							µ_ = πTemp004
							// line 162: return n
							πF.SetLineno(162)
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
					if πE = πClass.SetItem(πF, ßrecv_into.ToObject(), πTemp014); πE != nil {
						continue
					}
					// line 164: def recvfrom(self, bufsize, flags=0):
					πF.SetLineno(164)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "bufsize", Def: nil}
					πTemp002[2] = πg.Param{Name: "flags", Def: πg.NewInt(0).ToObject()}
					πTemp015 = πg.NewFunction(πg.NewCode("recvfrom", "/usr/lib/python2.7/_socket.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µbufsize *πg.Object = πArgs[1]; _ = µbufsize
						var µflags *πg.Object = πArgs[2]; _ = µflags
						var µbuffer *πg.Object = πg.UnboundLocal; _ = µbuffer
						var µn *πg.Object = πg.UnboundLocal; _ = µn
						var µaddr *πg.Object = πg.UnboundLocal; _ = µaddr
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
							// line 165: buffer = bytearray(bufsize)
							πF.SetLineno(165)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µbufsize, "bufsize"); πE != nil {
								continue
							}
							πTemp001[0] = µbufsize
							if πTemp002, πE = πg.ResolveGlobal(πF, ßbytearray); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µbuffer = πTemp003
							// line 166: n, addr = self.recvfrom_into(buffer, bufsize, flags)
							πF.SetLineno(166)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µbuffer, "buffer"); πE != nil {
								continue
							}
							πTemp001[0] = µbuffer
							if πE = πg.CheckLocal(πF, µbufsize, "bufsize"); πE != nil {
								continue
							}
							πTemp001[1] = µbufsize
							if πE = πg.CheckLocal(πF, µflags, "flags"); πE != nil {
								continue
							}
							πTemp001[2] = µflags
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßrecvfrom_into, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp004}}}, πTemp003); πE != nil {
								continue
							}
							µn = πTemp002
							µaddr = πTemp004
							// line 167: return str(buffer[:n]), addr
							πF.SetLineno(167)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.None, µn, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µbuffer, "buffer"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µbuffer, πTemp003); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.CheckLocal(πF, µaddr, "addr"); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(πTemp004, µaddr).ToObject()
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
					if πE = πClass.SetItem(πF, ßrecvfrom.ToObject(), πTemp015); πE != nil {
						continue
					}
					// line 169: def recvfrom_into(self, buffer, nbytes=0, flags=0):
					πF.SetLineno(169)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "buffer", Def: nil}
					πTemp002[2] = πg.Param{Name: "nbytes", Def: πg.NewInt(0).ToObject()}
					πTemp002[3] = πg.Param{Name: "flags", Def: πg.NewInt(0).ToObject()}
					πTemp016 = πg.NewFunction(πg.NewCode("recvfrom_into", "/usr/lib/python2.7/_socket.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µbuffer *πg.Object = πArgs[1]; _ = µbuffer
						var µnbytes *πg.Object = πArgs[2]; _ = µnbytes
						var µflags *πg.Object = πArgs[3]; _ = µflags
						var µn *πg.Object = πg.UnboundLocal; _ = µn
						var µ_ *πg.Object = πg.UnboundLocal; _ = µ_
						var µerr *πg.Object = πg.UnboundLocal; _ = µerr
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 170: n, _, err = _Recvfrom(self._fd, buffer, flags)
							πF.SetLineno(170)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_fd, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µbuffer, "buffer"); πE != nil {
								continue
							}
							πTemp001[1] = µbuffer
							if πE = πg.CheckLocal(πF, µflags, "flags"); πE != nil {
								continue
							}
							πTemp001[2] = µflags
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_Recvfrom); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}}}, πTemp003); πE != nil {
								continue
							}
							µn = πTemp002
							µ_ = πTemp004
							µerr = πTemp005
							if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, µerr); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label1
							}
							goto Label2
							// line 171: if err:
							πF.SetLineno(171)
						Label1:
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µerr, ßError, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 172: raise error(err.Error())
							πF.SetLineno(172)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label2
						Label2:
							// line 173: return n, None
							πF.SetLineno(173)
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(µn, πTemp003).ToObject()
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
					if πE = πClass.SetItem(πF, ßrecvfrom_into.ToObject(), πTemp016); πE != nil {
						continue
					}
					// line 175: def setsockopt(self, level, optname, value):
					πF.SetLineno(175)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "level", Def: nil}
					πTemp002[2] = πg.Param{Name: "optname", Def: nil}
					πTemp002[3] = πg.Param{Name: "value", Def: nil}
					πTemp017 = πg.NewFunction(πg.NewCode("setsockopt", "/usr/lib/python2.7/_socket.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µlevel *πg.Object = πArgs[1]; _ = µlevel
						var µoptname *πg.Object = πArgs[2]; _ = µoptname
						var µvalue *πg.Object = πArgs[3]; _ = µvalue
						var µerr *πg.Object = πg.UnboundLocal; _ = µerr
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
							// line 176: err = _SetsockoptInt(self._fd, level, optname, value)
							πF.SetLineno(176)
							πTemp001 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_fd, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µlevel, "level"); πE != nil {
								continue
							}
							πTemp001[1] = µlevel
							if πE = πg.CheckLocal(πF, µoptname, "optname"); πE != nil {
								continue
							}
							πTemp001[2] = µoptname
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp001[3] = µvalue
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_SetsockoptInt); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µerr = πTemp003
							if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µerr); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 177: if err:
							πF.SetLineno(177)
						Label1:
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µerr, ßError, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 178: raise error(err.Error())
							πF.SetLineno(178)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
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
					if πE = πClass.SetItem(πF, ßsetsockopt.ToObject(), πTemp017); πE != nil {
						continue
					}
					// line 180: def send(self, string, flags=0):
					πF.SetLineno(180)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "string", Def: nil}
					πTemp002[2] = πg.Param{Name: "flags", Def: πg.NewInt(0).ToObject()}
					πTemp018 = πg.NewFunction(πg.NewCode("send", "/usr/lib/python2.7/_socket.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µstring *πg.Object = πArgs[1]; _ = µstring
						var µflags *πg.Object = πArgs[2]; _ = µflags
						var µsockaddr *πg.Object = πg.UnboundLocal; _ = µsockaddr
						var µerr *πg.Object = πg.UnboundLocal; _ = µerr
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
							// line 181: sockaddr, err = _Getsockname(self._fd)
							πF.SetLineno(181)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_fd, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_Getsockname); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp004}}}, πTemp003); πE != nil {
								continue
							}
							µsockaddr = πTemp002
							µerr = πTemp004
							if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, µerr); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							goto Label2
							// line 182: if err:
							πF.SetLineno(182)
						Label1:
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
								continue
							}
							πTemp001[0] = µerr
							if πTemp002, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 183: raise error(err)
							πF.SetLineno(183)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label2
						Label2:
							// line 184: err = _Sendto(self._fd, string, flags, sockaddr)
							πF.SetLineno(184)
							πTemp001 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_fd, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
								continue
							}
							πTemp001[1] = µstring
							if πE = πg.CheckLocal(πF, µflags, "flags"); πE != nil {
								continue
							}
							πTemp001[2] = µflags
							if πE = πg.CheckLocal(πF, µsockaddr, "sockaddr"); πE != nil {
								continue
							}
							πTemp001[3] = µsockaddr
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_Sendto); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µerr = πTemp003
							if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, µerr); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label3
							}
							goto Label4
							// line 185: if err:
							πF.SetLineno(185)
						Label3:
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
								continue
							}
							πTemp001[0] = µerr
							if πTemp002, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 186: raise error(err)
							πF.SetLineno(186)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label4
						Label4:
							// line 187: return len(string)
							πF.SetLineno(187)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
								continue
							}
							πTemp001[0] = µstring
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
					if πE = πClass.SetItem(πF, ßsend.ToObject(), πTemp018); πE != nil {
						continue
					}
					// line 189: def sendto(self, string, flags_or_address, address=None):
					πF.SetLineno(189)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "string", Def: nil}
					πTemp002[2] = πg.Param{Name: "flags_or_address", Def: nil}
					if πTemp020, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[3] = πg.Param{Name: "address", Def: πTemp020}
					πTemp019 = πg.NewFunction(πg.NewCode("sendto", "/usr/lib/python2.7/_socket.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µstring *πg.Object = πArgs[1]; _ = µstring
						var µflags_or_address *πg.Object = πArgs[2]; _ = µflags_or_address
						var µaddress *πg.Object = πArgs[3]; _ = µaddress
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
							// line 190: raise NotImplementedError
							πF.SetLineno(190)
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
					if πE = πClass.SetItem(πF, ßsendto.ToObject(), πTemp019); πE != nil {
						continue
					}
					// line 192: def sendall(self, string, flags=0):
					πF.SetLineno(192)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "string", Def: nil}
					πTemp002[2] = πg.Param{Name: "flags", Def: πg.NewInt(0).ToObject()}
					πTemp020 = πg.NewFunction(πg.NewCode("sendall", "/usr/lib/python2.7/_socket.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µstring *πg.Object = πArgs[1]; _ = µstring
						var µflags *πg.Object = πArgs[2]; _ = µflags
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
							// line 193: return self.send(string, flags)
							πF.SetLineno(193)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
								continue
							}
							πTemp001[0] = µstring
							if πE = πg.CheckLocal(πF, µflags, "flags"); πE != nil {
								continue
							}
							πTemp001[1] = µflags
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßsend, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßsendall.ToObject(), πTemp020); πE != nil {
						continue
					}
					// line 195: def setblocking(self, flag):
					πF.SetLineno(195)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "flag", Def: nil}
					πTemp021 = πg.NewFunction(πg.NewCode("setblocking", "/usr/lib/python2.7/_socket.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µflag *πg.Object = πArgs[1]; _ = µflag
						var µerr *πg.Object = πg.UnboundLocal; _ = µerr
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
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
							// line 196: err = _SetNonblock(fd, int(not flag))
							πF.SetLineno(196)
							πTemp001 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßfd); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µflag, "flag"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µflag); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp004).ToObject()
							πTemp003[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πTemp001[1] = πTemp005
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_SetNonblock); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µerr = πTemp005
							if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µerr); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 197: if err:
							πF.SetLineno(197)
						Label1:
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µerr, ßError, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp005
							if πTemp002, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 198: raise error(err.Error())
							πF.SetLineno(198)
							πE = πF.Raise(πTemp005, nil, nil)
							continue
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
					if πE = πClass.SetItem(πF, ßsetblocking.ToObject(), πTemp021); πE != nil {
						continue
					}
					// line 200: def settimeout(self, value):
					πF.SetLineno(200)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "value", Def: nil}
					πTemp022 = πg.NewFunction(πg.NewCode("settimeout", "/usr/lib/python2.7/_socket.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µvalue *πg.Object = πArgs[1]; _ = µvalue
						var µtimevalue *πg.Object = πg.UnboundLocal; _ = µtimevalue
						var µtimeval *πg.Object = πg.UnboundLocal; _ = µtimeval
						var µfrac *πg.Object = πg.UnboundLocal; _ = µfrac
						var µinteger *πg.Object = πg.UnboundLocal; _ = µinteger
						var µerr *πg.Object = πg.UnboundLocal; _ = µerr
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
							πTemp001 = πg.GetBool(µvalue == πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 201: if value is None:
							πF.SetLineno(201)
						Label1:
							// line 202: timevalue = None
							πF.SetLineno(202)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							µtimevalue = πTemp001
							goto Label3
						Label2:
							// line 204: timeval = _Timeval.new()
							πF.SetLineno(204)
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_Timeval); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßnew, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µtimeval = πTemp001
							// line 205: frac, integer = math.modf(value)
							πF.SetLineno(205)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp004[0] = µvalue
							if πTemp001, πE = πg.ResolveGlobal(πF, ßmath); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßmodf, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp005}}}, πTemp001); πE != nil {
								continue
							}
							µfrac = πTemp002
							µinteger = πTemp005
							// line 206: timeval.Sec = int(integer)
							πF.SetLineno(206)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µinteger, "integer"); πE != nil {
								continue
							}
							πTemp004[0] = µinteger
							if πTemp001, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtimeval, "timeval"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µtimeval, ßSec, πTemp001); πE != nil {
								continue
							}
							// line 207: timeval.Usec = int(frac * 1000000.0)
							πF.SetLineno(207)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfrac, "frac"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mul(πF, µfrac, πg.NewFloat(1000000.0).ToObject()); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtimeval, "timeval"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µtimeval, ßUsec, πTemp001); πE != nil {
								continue
							}
							goto Label3
						Label3:
							// line 208: err = _SetsockoptTimeval(self._fd, level, SO_RCVTIMEO, timeval)
							πF.SetLineno(208)
							πTemp004 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_fd, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßlevel); πE != nil {
								continue
							}
							πTemp004[1] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßSO_RCVTIMEO); πE != nil {
								continue
							}
							πTemp004[2] = πTemp001
							if πE = πg.CheckLocal(πF, µtimeval, "timeval"); πE != nil {
								continue
							}
							πTemp004[3] = µtimeval
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_SetsockoptTimeval); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µerr = πTemp002
							if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, µerr); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label4
							}
							goto Label5
							// line 209: if err:
							πF.SetLineno(209)
						Label4:
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
								continue
							}
							πTemp004[0] = µerr
							if πTemp001, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 210: raise error(err)
							πF.SetLineno(210)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
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
					if πE = πClass.SetItem(πF, ßsettimeout.ToObject(), πTemp022); πE != nil {
						continue
					}
					// line 212: def gettimeout(self):
					πF.SetLineno(212)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp023 = πg.NewFunction(πg.NewCode("gettimeout", "/usr/lib/python2.7/_socket.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNotImplementedError); πE != nil {
								continue
							}
							// line 213: raise NotImplementedError
							πF.SetLineno(213)
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
					if πE = πClass.SetItem(πF, ßgettimeout.ToObject(), πTemp023); πE != nil {
						continue
					}
					// line 215: def shutdown(self, how):
					πF.SetLineno(215)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "how", Def: nil}
					πTemp024 = πg.NewFunction(πg.NewCode("shutdown", "/usr/lib/python2.7/_socket.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µhow *πg.Object = πArgs[1]; _ = µhow
						var µerr *πg.Object = πg.UnboundLocal; _ = µerr
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
							// line 216: err = _Shutdown(self._fd, how)
							πF.SetLineno(216)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_fd, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µhow, "how"); πE != nil {
								continue
							}
							πTemp001[1] = µhow
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_Shutdown); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µerr = πTemp003
							if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µerr); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 217: if err:
							πF.SetLineno(217)
						Label1:
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
								continue
							}
							πTemp001[0] = µerr
							if πTemp002, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 218: raise error(err)
							πF.SetLineno(218)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
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
					if πE = πClass.SetItem(πF, ßshutdown.ToObject(), πTemp024); πE != nil {
						continue
					}
					// line 220: def _parse_address(self, address):
					πF.SetLineno(220)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "address", Def: nil}
					πTemp025 = πg.NewFunction(πg.NewCode("_parse_address", "/usr/lib/python2.7/_socket.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µaddress *πg.Object = πArgs[1]; _ = µaddress
						var µsockaddr *πg.Object = πg.UnboundLocal; _ = µsockaddr
						var µhost *πg.Object = πg.UnboundLocal; _ = µhost
						var µport *πg.Object = πg.UnboundLocal; _ = µport
						var µip *πg.Object = πg.UnboundLocal; _ = µip
						var µips *πg.Object = πg.UnboundLocal; _ = µips
						var µerr *πg.Object = πg.UnboundLocal; _ = µerr
						var µconvert *πg.Object = πg.UnboundLocal; _ = µconvert
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
						var πTemp006 bool
						_ = πTemp006
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 13: goto Label13
							case 14: goto Label14
							default: panic("unexpected function state")
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßfamily, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßAF_UNIX); πE != nil {
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
							goto Label2
							// line 221: if self.family == AF_UNIX:
							πF.SetLineno(221)
						Label1:
							// line 222: sockaddr = _SockaddrUnix.new()
							πF.SetLineno(222)
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_SockaddrUnix); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßnew, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µsockaddr = πTemp001
							// line 223: sockaddr.Name, = address
							πF.SetLineno(223)
							if πE = πg.CheckLocal(πF, µaddress, "address"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µaddress); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µsockaddr, "sockaddr"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µsockaddr, ßName, πTemp001); πE != nil {
								continue
							}
							// line 224: return sockaddr
							πF.SetLineno(224)
							if πE = πg.CheckLocal(πF, µsockaddr, "sockaddr"); πE != nil {
								continue
							}
							πR = µsockaddr
							continue
							goto Label2
						Label2:
							// line 225: host, port = address
							πF.SetLineno(225)
							if πE = πg.CheckLocal(πF, µaddress, "address"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp002}}}, µaddress); πE != nil {
								continue
							}
							µhost = πTemp001
							µport = πTemp002
							if πE = πg.CheckLocal(πF, µhost, "host"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µhost); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp004).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label3
							}
							goto Label4
							// line 226: if not host:
							πF.SetLineno(226)
						Label3:
							// line 227: host = '127.0.0.1'
							πF.SetLineno(227)
							µhost = πg.NewStr("127.0.0.1").ToObject()
							goto Label4
						Label4:
							// line 228: ip = _ParseIP(host)
							πF.SetLineno(228)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µhost, "host"); πE != nil {
								continue
							}
							πTemp005[0] = µhost
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_ParseIP); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µip = πTemp002
							if πE = πg.CheckLocal(πF, µip, "ip"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µip); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label5
							}
							goto Label6
							// line 229: if ip:
							πF.SetLineno(229)
						Label5:
							// line 230: ips = [ip]
							πF.SetLineno(230)
							πTemp005 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µip, "ip"); πE != nil {
								continue
							}
							πTemp005[0] = µip
							πTemp001 = πg.NewList(πTemp005...).ToObject()
							µips = πTemp001
							goto Label7
						Label6:
							// line 232: ips, err = _LookupIP(host)
							πF.SetLineno(232)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µhost, "host"); πE != nil {
								continue
							}
							πTemp005[0] = µhost
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_LookupIP); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp003}}}, πTemp002); πE != nil {
								continue
							}
							µips = πTemp001
							µerr = πTemp003
							if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µerr); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label8
							}
							goto Label9
							// line 233: if err:
							πF.SetLineno(233)
						Label8:
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µerr, ßError, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp005[0] = πTemp002
							if πTemp001, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 234: raise error(err.Error())
							πF.SetLineno(234)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							goto Label9
						Label9:
							goto Label7
						Label7:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßfamily, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßAF_INET); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, πTemp002, πTemp003); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label10
							}
							goto Label11
							// line 235: if self.family == AF_INET:
							πF.SetLineno(235)
						Label10:
							// line 236: convert = _IP.To4
							πF.SetLineno(236)
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_IP); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßTo4, nil); πE != nil {
								continue
							}
							µconvert = πTemp002
							goto Label12
						Label11:
							// line 238: convert = _IP.To6
							πF.SetLineno(238)
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_IP); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßTo6, nil); πE != nil {
								continue
							}
							µconvert = πTemp002
							goto Label12
						Label12:
							if πE = πg.CheckLocal(πF, µips, "ips"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, µips); πE != nil {
								continue
							}
							πF.PushCheckpoint(14)
							πTemp004 = false
						Label13:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp004 {
								πF.PopCheckpoint()
								goto Label15
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
								µip = πTemp002
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(13)            
							// line 240: ip = convert(ip)
							πF.SetLineno(240)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µip, "ip"); πE != nil {
								continue
							}
							πTemp005[0] = µip
							if πE = πg.CheckLocal(πF, µconvert, "convert"); πE != nil {
								continue
							}
							if πTemp002, πE = µconvert.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µip = πTemp002
							if πE = πg.CheckLocal(πF, µip, "ip"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, µip); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label16
							}
							goto Label17
							// line 241: if ip:
							πF.SetLineno(241)
						Label16:
							// line 242: break
							πF.SetLineno(242)
							πTemp004 = true
							continue
							goto Label17
						Label17:
							continue
						Label14:
							if πE != nil || πR != nil {
								continue
							}
							πTemp005 = πF.MakeArgs(1)
							πTemp005[0] = πg.NewStr("cannot resolve address").ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 244: raise error('cannot resolve address')
							πF.SetLineno(244)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
						Label15:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßfamily, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßAF_INET); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, πTemp002, πTemp003); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label18
							}
							goto Label19
							// line 245: if self.family == AF_INET:
							πF.SetLineno(245)
						Label18:
							// line 246: sockaddr = _SockaddrInet4.new()
							πF.SetLineno(246)
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_SockaddrInet4); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßnew, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µsockaddr = πTemp001
							goto Label20
						Label19:
							// line 248: sockaddr = _SockaddrInet6.new()
							πF.SetLineno(248)
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_SockaddrInet6); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßnew, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µsockaddr = πTemp001
							goto Label20
						Label20:
							// line 249: sockaddr.Port = port
							πF.SetLineno(249)
							if πE = πg.CheckLocal(πF, µport, "port"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µport); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µsockaddr, "sockaddr"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µsockaddr, ßPort, πTemp001); πE != nil {
								continue
							}
							// line 250: sockaddr.Addr[:] = ip
							πF.SetLineno(250)
							if πE = πg.CheckLocal(πF, µip, "ip"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µip); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µsockaddr, "sockaddr"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µsockaddr, ßAddr, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.SetItem(πF, πTemp002, πTemp003, πTemp001); πE != nil {
								continue
							}
							// line 251: return sockaddr
							πF.SetLineno(251)
							if πE = πg.CheckLocal(πF, µsockaddr, "sockaddr"); πE != nil {
								continue
							}
							πR = µsockaddr
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß_parse_address.ToObject(), πTemp025); πE != nil {
						continue
					}
					// line 253: def _get_address(self, sockaddr):
					πF.SetLineno(253)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "sockaddr", Def: nil}
					πTemp026 = πg.NewFunction(πg.NewCode("_get_address", "/usr/lib/python2.7/_socket.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µsockaddr *πg.Object = πArgs[1]; _ = µsockaddr
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
						var πTemp007 *πg.Object
						_ = πTemp007
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µsockaddr, "sockaddr"); πE != nil {
								continue
							}
							πTemp001[0] = µsockaddr
							πTemp002 = πF.MakeArgs(1)
							if πTemp003, πE = πg.ResolveGlobal(πF, ß_SockaddrUnix); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßnew, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[1] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp005, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							goto Label2
							// line 254: if isinstance(sockaddr, type(_SockaddrUnix.new())):
							πF.SetLineno(254)
						Label1:
							// line 255: return (sockaddr.Name,)
							πF.SetLineno(255)
							if πE = πg.CheckLocal(πF, µsockaddr, "sockaddr"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µsockaddr, ßName, nil); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple1(πTemp004).ToObject()
							πR = πTemp003
							continue
							goto Label2
						Label2:
							// line 256: return _IPv4(*sockaddr.Addr).String(), sockaddr.Port
							πF.SetLineno(256)
							if πE = πg.CheckLocal(πF, µsockaddr, "sockaddr"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µsockaddr, ßAddr, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.ResolveGlobal(πF, ß_IPv4); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Invoke(πF, πTemp006, nil, πTemp004, nil, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp007, ßString, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µsockaddr, "sockaddr"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µsockaddr, ßPort, nil); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple2(πTemp006, πTemp004).ToObject()
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
					if πE = πClass.SetItem(πF, ß_get_address.ToObject(), πTemp026); πE != nil {
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("socket").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßsocket.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 259: def fromfd(fd, family, type, proto=None):
			πF.SetLineno(259)
			πTemp006 = make([]πg.Param, 4)
			πTemp006[0] = πg.Param{Name: "fd", Def: nil}
			πTemp006[1] = πg.Param{Name: "family", Def: nil}
			πTemp006[2] = πg.Param{Name: "type", Def: nil}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[3] = πg.Param{Name: "proto", Def: πTemp003}
			πTemp001 = πg.NewFunction(πg.NewCode("fromfd", "/usr/lib/python2.7/_socket.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µfd *πg.Object = πArgs[0]; _ = µfd
				var µfamily *πg.Object = πArgs[1]; _ = µfamily
				var µtype *πg.Object = πArgs[2]; _ = µtype
				var µproto *πg.Object = πArgs[3]; _ = µproto
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
					// line 260: return socket(family, type, proto, fd)
					πF.SetLineno(260)
					πTemp001 = πF.MakeArgs(4)
					if πE = πg.CheckLocal(πF, µfamily, "family"); πE != nil {
						continue
					}
					πTemp001[0] = µfamily
					if πE = πg.CheckLocal(πF, µtype, "type"); πE != nil {
						continue
					}
					πTemp001[1] = µtype
					if πE = πg.CheckLocal(πF, µproto, "proto"); πE != nil {
						continue
					}
					πTemp001[2] = µproto
					if πE = πg.CheckLocal(πF, µfd, "fd"); πE != nil {
						continue
					}
					πTemp001[3] = µfd
					if πTemp002, πE = πg.ResolveGlobal(πF, ßsocket); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßfromfd.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 263: def gethostbyname(hostname):
			πF.SetLineno(263)
			πTemp006 = make([]πg.Param, 1)
			πTemp006[0] = πg.Param{Name: "hostname", Def: nil}
			πTemp003 = πg.NewFunction(πg.NewCode("gethostbyname", "/usr/lib/python2.7/_socket.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µhostname *πg.Object = πArgs[0]; _ = µhostname
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
					// line 264: raise NotImplementedError
					πF.SetLineno(264)
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
			if πE = πF.Globals().SetItem(πF, ßgethostbyname.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 267: def gethostbyaddr(ipaddr):
			πF.SetLineno(267)
			πTemp006 = make([]πg.Param, 1)
			πTemp006[0] = πg.Param{Name: "ipaddr", Def: nil}
			πTemp005 = πg.NewFunction(πg.NewCode("gethostbyaddr", "/usr/lib/python2.7/_socket.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µipaddr *πg.Object = πArgs[0]; _ = µipaddr
				var µnames *πg.Object = πg.UnboundLocal; _ = µnames
				var µerr *πg.Object = πg.UnboundLocal; _ = µerr
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
				var πTemp006 *πg.Object
				_ = πTemp006
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 268: names, err = _LookupAddr(ipaddr)
					πF.SetLineno(268)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µipaddr, "ipaddr"); πE != nil {
						continue
					}
					πTemp001[0] = µipaddr
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_LookupAddr); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp004}}}, πTemp003); πE != nil {
						continue
					}
					µnames = πTemp002
					µerr = πTemp004
					if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, µerr); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label1
					}
					goto Label2
					// line 269: if err:
					πF.SetLineno(269)
				Label1:
					// line 270: return error(err)
					πF.SetLineno(270)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
						continue
					}
					πTemp001[0] = µerr
					if πTemp002, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
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
					// line 271: return names[0], [], [ipaddr]
					πF.SetLineno(271)
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µnames, "names"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µnames, πTemp003); πE != nil {
						continue
					}
					πTemp001 = make([]*πg.Object, 0)
					πTemp003 = πg.NewList(πTemp001...).ToObject()
					πTemp001 = make([]*πg.Object, 1)
					if πE = πg.CheckLocal(πF, µipaddr, "ipaddr"); πE != nil {
						continue
					}
					πTemp001[0] = µipaddr
					πTemp006 = πg.NewList(πTemp001...).ToObject()
					πTemp002 = πg.NewTuple3(πTemp004, πTemp003, πTemp006).ToObject()
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
			if πE = πF.Globals().SetItem(πF, ßgethostbyaddr.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 274: def gethostname():
			πF.SetLineno(274)
			πTemp006 = make([]πg.Param, 0)
			πTemp007 = πg.NewFunction(πg.NewCode("gethostname", "/usr/lib/python2.7/_socket.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 275: raise NotImplementedError
					πF.SetLineno(275)
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
			if πE = πF.Globals().SetItem(πF, ßgethostname.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 278: def getprotobyname(proto):
			πF.SetLineno(278)
			πTemp006 = make([]πg.Param, 1)
			πTemp006[0] = πg.Param{Name: "proto", Def: nil}
			πTemp008 = πg.NewFunction(πg.NewCode("getprotobyname", "/usr/lib/python2.7/_socket.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µproto *πg.Object = πArgs[0]; _ = µproto
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
					// line 279: raise NotImplementedError
					πF.SetLineno(279)
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
			if πE = πF.Globals().SetItem(πF, ßgetprotobyname.ToObject(), πTemp008); πE != nil {
				continue
			}
			// line 283: def getservbyname(servicename, protocolname=None):
			πF.SetLineno(283)
			πTemp006 = make([]πg.Param, 2)
			πTemp006[0] = πg.Param{Name: "servicename", Def: nil}
			if πTemp010, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[1] = πg.Param{Name: "protocolname", Def: πTemp010}
			πTemp009 = πg.NewFunction(πg.NewCode("getservbyname", "/usr/lib/python2.7/_socket.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µservicename *πg.Object = πArgs[0]; _ = µservicename
				var µprotocolname *πg.Object = πArgs[1]; _ = µprotocolname
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
					// line 284: raise NotImplementedError
					πF.SetLineno(284)
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
			if πE = πF.Globals().SetItem(πF, ßgetservbyname.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 287: def getservbyport(portnumber, protocolname=None):
			πF.SetLineno(287)
			πTemp006 = make([]πg.Param, 2)
			πTemp006[0] = πg.Param{Name: "portnumber", Def: nil}
			if πTemp011, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[1] = πg.Param{Name: "protocolname", Def: πTemp011}
			πTemp010 = πg.NewFunction(πg.NewCode("getservbyport", "/usr/lib/python2.7/_socket.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µportnumber *πg.Object = πArgs[0]; _ = µportnumber
				var µprotocolname *πg.Object = πArgs[1]; _ = µprotocolname
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
					// line 288: raise NotImplementedError
					πF.SetLineno(288)
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
			if πE = πF.Globals().SetItem(πF, ßgetservbyport.ToObject(), πTemp010); πE != nil {
				continue
			}
			// line 291: def socketpair(family=None, type=None, proto=None):
			πF.SetLineno(291)
			πTemp006 = make([]πg.Param, 3)
			if πTemp012, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[0] = πg.Param{Name: "family", Def: πTemp012}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[1] = πg.Param{Name: "type", Def: πTemp012}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[2] = πg.Param{Name: "proto", Def: πTemp012}
			πTemp011 = πg.NewFunction(πg.NewCode("socketpair", "/usr/lib/python2.7/_socket.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µfamily *πg.Object = πArgs[0]; _ = µfamily
				var µtype *πg.Object = πArgs[1]; _ = µtype
				var µproto *πg.Object = πArgs[2]; _ = µproto
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
					// line 292: raise NotImplementedError
					πF.SetLineno(292)
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
			if πE = πF.Globals().SetItem(πF, ßsocketpair.ToObject(), πTemp011); πE != nil {
				continue
			}
			// line 295: def ntohs(n):
			πF.SetLineno(295)
			πTemp006 = make([]πg.Param, 1)
			πTemp006[0] = πg.Param{Name: "n", Def: nil}
			πTemp012 = πg.NewFunction(πg.NewCode("ntohs", "/usr/lib/python2.7/_socket.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µn *πg.Object = πArgs[0]; _ = µn
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
					// line 296: raise NotImplementedError
					πF.SetLineno(296)
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
			if πE = πF.Globals().SetItem(πF, ßntohs.ToObject(), πTemp012); πE != nil {
				continue
			}
			// line 299: def ntohl(n):
			πF.SetLineno(299)
			πTemp006 = make([]πg.Param, 1)
			πTemp006[0] = πg.Param{Name: "n", Def: nil}
			πTemp013 = πg.NewFunction(πg.NewCode("ntohl", "/usr/lib/python2.7/_socket.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µn *πg.Object = πArgs[0]; _ = µn
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
					// line 300: raise NotImplementedError
					πF.SetLineno(300)
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
			if πE = πF.Globals().SetItem(πF, ßntohl.ToObject(), πTemp013); πE != nil {
				continue
			}
			// line 303: def htons(n):
			πF.SetLineno(303)
			πTemp006 = make([]πg.Param, 1)
			πTemp006[0] = πg.Param{Name: "n", Def: nil}
			πTemp014 = πg.NewFunction(πg.NewCode("htons", "/usr/lib/python2.7/_socket.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µn *πg.Object = πArgs[0]; _ = µn
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
					// line 304: raise NotImplementedError
					πF.SetLineno(304)
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
			if πE = πF.Globals().SetItem(πF, ßhtons.ToObject(), πTemp014); πE != nil {
				continue
			}
			// line 307: def htonl(n):
			πF.SetLineno(307)
			πTemp006 = make([]πg.Param, 1)
			πTemp006[0] = πg.Param{Name: "n", Def: nil}
			πTemp015 = πg.NewFunction(πg.NewCode("htonl", "/usr/lib/python2.7/_socket.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µn *πg.Object = πArgs[0]; _ = µn
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
					// line 308: raise NotImplementedError
					πF.SetLineno(308)
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
			if πE = πF.Globals().SetItem(πF, ßhtonl.ToObject(), πTemp015); πE != nil {
				continue
			}
			// line 312: def getaddrinfo(host, port, family=None, socktype=None, proto=None, flags=None):
			πF.SetLineno(312)
			πTemp006 = make([]πg.Param, 6)
			πTemp006[0] = πg.Param{Name: "host", Def: nil}
			πTemp006[1] = πg.Param{Name: "port", Def: nil}
			if πTemp017, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[2] = πg.Param{Name: "family", Def: πTemp017}
			if πTemp017, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[3] = πg.Param{Name: "socktype", Def: πTemp017}
			if πTemp017, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[4] = πg.Param{Name: "proto", Def: πTemp017}
			if πTemp017, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[5] = πg.Param{Name: "flags", Def: πTemp017}
			πTemp016 = πg.NewFunction(πg.NewCode("getaddrinfo", "/usr/lib/python2.7/_socket.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µhost *πg.Object = πArgs[0]; _ = µhost
				var µport *πg.Object = πArgs[1]; _ = µport
				var µfamily *πg.Object = πArgs[2]; _ = µfamily
				var µsocktype *πg.Object = πArgs[3]; _ = µsocktype
				var µproto *πg.Object = πArgs[4]; _ = µproto
				var µflags *πg.Object = πArgs[5]; _ = µflags
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
					// line 313: raise NotImplementedError
					πF.SetLineno(313)
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
			if πE = πF.Globals().SetItem(πF, ßgetaddrinfo.ToObject(), πTemp016); πE != nil {
				continue
			}
			// line 317: def getnameinfo(sockaddr, flags):
			πF.SetLineno(317)
			πTemp006 = make([]πg.Param, 2)
			πTemp006[0] = πg.Param{Name: "sockaddr", Def: nil}
			πTemp006[1] = πg.Param{Name: "flags", Def: nil}
			πTemp017 = πg.NewFunction(πg.NewCode("getnameinfo", "/usr/lib/python2.7/_socket.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µsockaddr *πg.Object = πArgs[0]; _ = µsockaddr
				var µflags *πg.Object = πArgs[1]; _ = µflags
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
					// line 318: raise NotImplementedError
					πF.SetLineno(318)
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
			if πE = πF.Globals().SetItem(πF, ßgetnameinfo.ToObject(), πTemp017); πE != nil {
				continue
			}
			// line 322: def inet_aton(ipaddr):
			πF.SetLineno(322)
			πTemp006 = make([]πg.Param, 1)
			πTemp006[0] = πg.Param{Name: "ipaddr", Def: nil}
			πTemp018 = πg.NewFunction(πg.NewCode("inet_aton", "/usr/lib/python2.7/_socket.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µipaddr *πg.Object = πArgs[0]; _ = µipaddr
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
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 323: return ''.join(chr(int(n)) for n in ipaddr.split('.'))
					πF.SetLineno(323)
					πTemp001 = πF.MakeArgs(1)
					πTemp003 = make([]πg.Param, 0)
					πTemp002 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/_socket.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 []*πg.Object
						_ = πTemp007
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
								πTemp002[0] = πg.NewStr(".").ToObject()
								if πE = πg.CheckLocal(πF, µipaddr, "ipaddr"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.GetAttr(πF, µipaddr, ßsplit, nil); πE != nil {
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
									µn = πTemp003
								}
								if πE != nil || !πTemp006 {
									continue
								}
								πF.PushCheckpoint(1)            
								// line 323: return ''.join(chr(int(n)) for n in ipaddr.split('.'))
								πF.SetLineno(323)
								πTemp002 = πF.MakeArgs(1)
								πTemp007 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
									continue
								}
								πTemp007[0] = µn
								if πTemp003, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
									continue
								}
								if πTemp004, πE = πTemp003.Call(πF, πTemp007, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp007)
								πTemp002[0] = πTemp004
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
					if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp004
					if πTemp004, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßinet_aton.ToObject(), πTemp018); πE != nil {
				continue
			}
			// line 327: def inet_ntoa(ipaddr):
			πF.SetLineno(327)
			πTemp006 = make([]πg.Param, 1)
			πTemp006[0] = πg.Param{Name: "ipaddr", Def: nil}
			πTemp019 = πg.NewFunction(πg.NewCode("inet_ntoa", "/usr/lib/python2.7/_socket.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µipaddr *πg.Object = πArgs[0]; _ = µipaddr
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
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 328: return ''.join(ord(c) for c in ipaddr)
					πF.SetLineno(328)
					πTemp001 = πF.MakeArgs(1)
					πTemp003 = make([]πg.Param, 0)
					πTemp002 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/_socket.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µc *πg.Object = πg.UnboundLocal; _ = µc
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
								if πE = πg.CheckLocal(πF, µipaddr, "ipaddr"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.Iter(πF, µipaddr); πE != nil {
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
								// line 328: return ''.join(ord(c) for c in ipaddr)
								πF.SetLineno(328)
								πTemp005 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
									continue
								}
								πTemp005[0] = µc
								if πTemp004, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
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
					if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp004
					if πTemp004, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßinet_ntoa.ToObject(), πTemp019); πE != nil {
				continue
			}
			// line 332: def getdefaulttimeout():
			πF.SetLineno(332)
			πTemp006 = make([]πg.Param, 0)
			πTemp020 = πg.NewFunction(πg.NewCode("getdefaulttimeout", "/usr/lib/python2.7/_socket.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 333: raise NotImplementedError
					πF.SetLineno(333)
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
			if πE = πF.Globals().SetItem(πF, ßgetdefaulttimeout.ToObject(), πTemp020); πE != nil {
				continue
			}
			// line 336: def setdefaulttimeout(to):
			πF.SetLineno(336)
			πTemp006 = make([]πg.Param, 1)
			πTemp006[0] = πg.Param{Name: "to", Def: nil}
			πTemp021 = πg.NewFunction(πg.NewCode("setdefaulttimeout", "/usr/lib/python2.7/_socket.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µto *πg.Object = πArgs[0]; _ = µto
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
					// line 337: raise NotImplementedError
					πF.SetLineno(337)
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
			if πE = πF.Globals().SetItem(πF, ßsetdefaulttimeout.ToObject(), πTemp021); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("_socket", Code)
}

