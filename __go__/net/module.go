package net

import (
	mod "net"
	"reflect"

	pygolin_runtime "github.com/pygolin/runtime"
)

func fun(f *pygolin_runtime.Frame, _ []*pygolin_runtime.Object) (*pygolin_runtime.Object, *pygolin_runtime.BaseException) {
	if true {
		var x mod.AddrError
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "AddrError", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.Buffers
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Buffers", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.CIDRMask)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "CIDRMask", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.DNSConfigError
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "DNSConfigError", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.DNSError
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "DNSError", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.DefaultResolver)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "DefaultResolver", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Dial)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Dial", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.DialIP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "DialIP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.DialTCP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "DialTCP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.DialTimeout)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "DialTimeout", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.DialUDP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "DialUDP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.DialUnix)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "DialUnix", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.Dialer
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Dialer", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.ErrWriteToConnected)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ErrWriteToConnected", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.FileConn)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "FileConn", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.FileListener)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "FileListener", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.FilePacketConn)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "FilePacketConn", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.FlagBroadcast))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "FlagBroadcast", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.FlagLoopback))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "FlagLoopback", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.FlagMulticast))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "FlagMulticast", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.FlagPointToPoint))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "FlagPointToPoint", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.FlagUp))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "FlagUp", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.Flags
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Flags", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.HardwareAddr
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "HardwareAddr", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.IP
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "IP", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.IPAddr
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "IPAddr", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.IPConn
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "IPConn", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.IPMask
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "IPMask", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.IPNet
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "IPNet", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.IPv4)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPv4", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.IPv4Mask)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPv4Mask", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.IPv4allrouter)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPv4allrouter", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.IPv4allsys)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPv4allsys", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.IPv4bcast)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPv4bcast", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IPv4len))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPv4len", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.IPv4zero)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPv4zero", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.IPv6interfacelocalallnodes)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPv6interfacelocalallnodes", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IPv6len))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPv6len", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.IPv6linklocalallnodes)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPv6linklocalallnodes", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.IPv6linklocalallrouters)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPv6linklocalallrouters", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.IPv6loopback)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPv6loopback", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.IPv6unspecified)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPv6unspecified", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.IPv6zero)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPv6zero", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.Interface
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Interface", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.InterfaceAddrs)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "InterfaceAddrs", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.InterfaceByIndex)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "InterfaceByIndex", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.InterfaceByName)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "InterfaceByName", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Interfaces)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Interfaces", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.InvalidAddrError
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "InvalidAddrError", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.JoinHostPort)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "JoinHostPort", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Listen)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Listen", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.ListenConfig
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "ListenConfig", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.ListenIP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ListenIP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.ListenMulticastUDP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ListenMulticastUDP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.ListenPacket)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ListenPacket", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.ListenTCP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ListenTCP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.ListenUDP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ListenUDP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.ListenUnix)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ListenUnix", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.ListenUnixgram)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ListenUnixgram", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.LookupAddr)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "LookupAddr", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.LookupCNAME)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "LookupCNAME", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.LookupHost)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "LookupHost", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.LookupIP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "LookupIP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.LookupMX)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "LookupMX", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.LookupNS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "LookupNS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.LookupPort)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "LookupPort", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.LookupSRV)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "LookupSRV", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.LookupTXT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "LookupTXT", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.MX
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "MX", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.NS
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "NS", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.OpError
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "OpError", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.ParseCIDR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ParseCIDR", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.ParseError
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "ParseError", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.ParseIP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ParseIP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.ParseMAC)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ParseMAC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Pipe)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Pipe", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.ResolveIPAddr)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ResolveIPAddr", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.ResolveTCPAddr)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ResolveTCPAddr", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.ResolveUDPAddr)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ResolveUDPAddr", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.ResolveUnixAddr)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ResolveUnixAddr", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.Resolver
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Resolver", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.SRV
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "SRV", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.SplitHostPort)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SplitHostPort", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.TCPAddr
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "TCPAddr", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.TCPConn
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "TCPConn", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.TCPListener
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "TCPListener", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.UDPAddr
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "UDPAddr", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.UDPConn
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "UDPConn", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.UnixAddr
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "UnixAddr", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.UnixConn
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "UnixConn", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.UnixListener
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "UnixListener", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.UnknownNetworkError
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "UnknownNetworkError", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}

	return nil, nil
}

var Code = pygolin_runtime.NewCode("<module>", "net", nil, 0, fun)

func init() {
	pygolin_runtime.RegisterModule("__go__/net", Code)
}
