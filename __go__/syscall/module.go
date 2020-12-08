package syscall

import (
	"reflect"
	mod "syscall"

	pygolin_runtime "github.com/pygolin/runtime"
)

func fun(f *pygolin_runtime.Frame, _ []*pygolin_runtime.Object) (*pygolin_runtime.Object, *pygolin_runtime.BaseException) {
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.AF_ALG))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "AF_ALG", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.AF_APPLETALK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "AF_APPLETALK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.AF_ASH))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "AF_ASH", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.AF_ATMPVC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "AF_ATMPVC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.AF_ATMSVC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "AF_ATMSVC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.AF_AX25))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "AF_AX25", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.AF_BLUETOOTH))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "AF_BLUETOOTH", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.AF_BRIDGE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "AF_BRIDGE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.AF_CAIF))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "AF_CAIF", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.AF_CAN))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "AF_CAN", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.AF_DECnet))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "AF_DECnet", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.AF_ECONET))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "AF_ECONET", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.AF_FILE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "AF_FILE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.AF_IEEE802154))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "AF_IEEE802154", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.AF_INET))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "AF_INET", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.AF_INET6))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "AF_INET6", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.AF_IPX))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "AF_IPX", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.AF_IRDA))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "AF_IRDA", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.AF_ISDN))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "AF_ISDN", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.AF_IUCV))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "AF_IUCV", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.AF_KEY))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "AF_KEY", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.AF_LLC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "AF_LLC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.AF_LOCAL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "AF_LOCAL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.AF_MAX))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "AF_MAX", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.AF_NETBEUI))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "AF_NETBEUI", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.AF_NETLINK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "AF_NETLINK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.AF_NETROM))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "AF_NETROM", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.AF_PACKET))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "AF_PACKET", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.AF_PHONET))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "AF_PHONET", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.AF_PPPOX))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "AF_PPPOX", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.AF_RDS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "AF_RDS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.AF_ROSE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "AF_ROSE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.AF_ROUTE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "AF_ROUTE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.AF_RXRPC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "AF_RXRPC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.AF_SECURITY))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "AF_SECURITY", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.AF_SNA))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "AF_SNA", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.AF_TIPC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "AF_TIPC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.AF_UNIX))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "AF_UNIX", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.AF_UNSPEC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "AF_UNSPEC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.AF_WANPIPE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "AF_WANPIPE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.AF_X25))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "AF_X25", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ARPHRD_ADAPT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ARPHRD_ADAPT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ARPHRD_APPLETLK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ARPHRD_APPLETLK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ARPHRD_ARCNET))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ARPHRD_ARCNET", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ARPHRD_ASH))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ARPHRD_ASH", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ARPHRD_ATM))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ARPHRD_ATM", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ARPHRD_AX25))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ARPHRD_AX25", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ARPHRD_BIF))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ARPHRD_BIF", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ARPHRD_CHAOS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ARPHRD_CHAOS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ARPHRD_CISCO))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ARPHRD_CISCO", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ARPHRD_CSLIP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ARPHRD_CSLIP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ARPHRD_CSLIP6))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ARPHRD_CSLIP6", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ARPHRD_DDCMP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ARPHRD_DDCMP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ARPHRD_DLCI))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ARPHRD_DLCI", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ARPHRD_ECONET))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ARPHRD_ECONET", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ARPHRD_EETHER))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ARPHRD_EETHER", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ARPHRD_ETHER))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ARPHRD_ETHER", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ARPHRD_EUI64))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ARPHRD_EUI64", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ARPHRD_FCAL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ARPHRD_FCAL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ARPHRD_FCFABRIC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ARPHRD_FCFABRIC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ARPHRD_FCPL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ARPHRD_FCPL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ARPHRD_FCPP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ARPHRD_FCPP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ARPHRD_FDDI))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ARPHRD_FDDI", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ARPHRD_FRAD))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ARPHRD_FRAD", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ARPHRD_HDLC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ARPHRD_HDLC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ARPHRD_HIPPI))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ARPHRD_HIPPI", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ARPHRD_HWX25))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ARPHRD_HWX25", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ARPHRD_IEEE1394))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ARPHRD_IEEE1394", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ARPHRD_IEEE802))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ARPHRD_IEEE802", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ARPHRD_IEEE80211))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ARPHRD_IEEE80211", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ARPHRD_IEEE80211_PRISM))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ARPHRD_IEEE80211_PRISM", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ARPHRD_IEEE80211_RADIOTAP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ARPHRD_IEEE80211_RADIOTAP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ARPHRD_IEEE802154))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ARPHRD_IEEE802154", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ARPHRD_IEEE802154_PHY))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ARPHRD_IEEE802154_PHY", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ARPHRD_IEEE802_TR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ARPHRD_IEEE802_TR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ARPHRD_INFINIBAND))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ARPHRD_INFINIBAND", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ARPHRD_IPDDP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ARPHRD_IPDDP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ARPHRD_IPGRE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ARPHRD_IPGRE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ARPHRD_IRDA))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ARPHRD_IRDA", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ARPHRD_LAPB))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ARPHRD_LAPB", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ARPHRD_LOCALTLK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ARPHRD_LOCALTLK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ARPHRD_LOOPBACK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ARPHRD_LOOPBACK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ARPHRD_METRICOM))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ARPHRD_METRICOM", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ARPHRD_NETROM))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ARPHRD_NETROM", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ARPHRD_NONE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ARPHRD_NONE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ARPHRD_PIMREG))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ARPHRD_PIMREG", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ARPHRD_PPP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ARPHRD_PPP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ARPHRD_PRONET))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ARPHRD_PRONET", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ARPHRD_RAWHDLC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ARPHRD_RAWHDLC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ARPHRD_ROSE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ARPHRD_ROSE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ARPHRD_RSRVD))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ARPHRD_RSRVD", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ARPHRD_SIT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ARPHRD_SIT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ARPHRD_SKIP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ARPHRD_SKIP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ARPHRD_SLIP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ARPHRD_SLIP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ARPHRD_SLIP6))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ARPHRD_SLIP6", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ARPHRD_TUNNEL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ARPHRD_TUNNEL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ARPHRD_TUNNEL6))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ARPHRD_TUNNEL6", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ARPHRD_VOID))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ARPHRD_VOID", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ARPHRD_X25))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ARPHRD_X25", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Accept)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Accept", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Accept4)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Accept4", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Access)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Access", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Acct)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Acct", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Adjtimex)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Adjtimex", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.AttachLsf)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "AttachLsf", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.B0))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "B0", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.B1000000))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "B1000000", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.B110))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "B110", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.B115200))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "B115200", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.B1152000))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "B1152000", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.B1200))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "B1200", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.B134))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "B134", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.B150))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "B150", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.B1500000))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "B1500000", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.B1800))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "B1800", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.B19200))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "B19200", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.B200))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "B200", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.B2000000))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "B2000000", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.B230400))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "B230400", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.B2400))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "B2400", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.B2500000))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "B2500000", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.B300))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "B300", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.B3000000))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "B3000000", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.B3500000))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "B3500000", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.B38400))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "B38400", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.B4000000))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "B4000000", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.B460800))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "B460800", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.B4800))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "B4800", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.B50))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "B50", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.B500000))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "B500000", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.B57600))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "B57600", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.B576000))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "B576000", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.B600))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "B600", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.B75))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "B75", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.B921600))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "B921600", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.B9600))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "B9600", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.BPF_A))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BPF_A", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.BPF_ABS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BPF_ABS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.BPF_ADD))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BPF_ADD", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.BPF_ALU))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BPF_ALU", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.BPF_AND))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BPF_AND", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.BPF_B))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BPF_B", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.BPF_DIV))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BPF_DIV", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.BPF_H))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BPF_H", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.BPF_IMM))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BPF_IMM", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.BPF_IND))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BPF_IND", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.BPF_JA))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BPF_JA", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.BPF_JEQ))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BPF_JEQ", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.BPF_JGE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BPF_JGE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.BPF_JGT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BPF_JGT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.BPF_JMP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BPF_JMP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.BPF_JSET))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BPF_JSET", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.BPF_K))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BPF_K", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.BPF_LD))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BPF_LD", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.BPF_LDX))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BPF_LDX", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.BPF_LEN))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BPF_LEN", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.BPF_LSH))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BPF_LSH", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.BPF_MAJOR_VERSION))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BPF_MAJOR_VERSION", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.BPF_MAXINSNS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BPF_MAXINSNS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.BPF_MEM))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BPF_MEM", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.BPF_MEMWORDS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BPF_MEMWORDS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.BPF_MINOR_VERSION))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BPF_MINOR_VERSION", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.BPF_MISC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BPF_MISC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.BPF_MSH))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BPF_MSH", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.BPF_MUL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BPF_MUL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.BPF_NEG))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BPF_NEG", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.BPF_OR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BPF_OR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.BPF_RET))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BPF_RET", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.BPF_RSH))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BPF_RSH", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.BPF_ST))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BPF_ST", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.BPF_STX))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BPF_STX", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.BPF_SUB))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BPF_SUB", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.BPF_TAX))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BPF_TAX", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.BPF_TXA))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BPF_TXA", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.BPF_W))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BPF_W", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.BPF_X))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BPF_X", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.BRKINT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BRKINT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Bind)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Bind", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.BindToDevice)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BindToDevice", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.BytePtrFromString)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BytePtrFromString", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.ByteSliceFromString)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ByteSliceFromString", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.CLOCAL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "CLOCAL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.CLONE_CHILD_CLEARTID))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "CLONE_CHILD_CLEARTID", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.CLONE_CHILD_SETTID))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "CLONE_CHILD_SETTID", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.CLONE_DETACHED))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "CLONE_DETACHED", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.CLONE_FILES))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "CLONE_FILES", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.CLONE_FS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "CLONE_FS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint64(mod.CLONE_IO))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "CLONE_IO", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.CLONE_NEWIPC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "CLONE_NEWIPC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.CLONE_NEWNET))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "CLONE_NEWNET", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.CLONE_NEWNS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "CLONE_NEWNS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.CLONE_NEWPID))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "CLONE_NEWPID", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.CLONE_NEWUSER))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "CLONE_NEWUSER", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.CLONE_NEWUTS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "CLONE_NEWUTS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.CLONE_PARENT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "CLONE_PARENT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.CLONE_PARENT_SETTID))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "CLONE_PARENT_SETTID", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.CLONE_PTRACE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "CLONE_PTRACE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.CLONE_SETTLS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "CLONE_SETTLS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.CLONE_SIGHAND))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "CLONE_SIGHAND", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.CLONE_SYSVSEM))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "CLONE_SYSVSEM", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.CLONE_THREAD))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "CLONE_THREAD", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.CLONE_UNTRACED))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "CLONE_UNTRACED", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.CLONE_VFORK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "CLONE_VFORK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.CLONE_VM))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "CLONE_VM", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.CREAD))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "CREAD", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.CS5))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "CS5", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.CS6))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "CS6", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.CS7))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "CS7", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.CS8))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "CS8", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.CSIZE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "CSIZE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.CSTOPB))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "CSTOPB", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Chdir)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Chdir", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Chmod)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Chmod", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Chown)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Chown", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Chroot)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Chroot", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Clearenv)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Clearenv", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Close)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Close", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.CloseOnExec)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "CloseOnExec", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.CmsgLen)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "CmsgLen", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.CmsgSpace)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "CmsgSpace", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.Cmsghdr
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Cmsghdr", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Connect)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Connect", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Creat)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Creat", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.Credential
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Credential", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.DT_BLK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "DT_BLK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.DT_CHR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "DT_CHR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.DT_DIR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "DT_DIR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.DT_FIFO))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "DT_FIFO", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.DT_LNK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "DT_LNK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.DT_REG))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "DT_REG", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.DT_SOCK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "DT_SOCK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.DT_UNKNOWN))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "DT_UNKNOWN", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.DT_WHT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "DT_WHT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.DetachLsf)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "DetachLsf", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.Dirent
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Dirent", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Dup)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Dup", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Dup2)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Dup2", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Dup3)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Dup3", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.E2BIG))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "E2BIG", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.EACCES))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EACCES", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.EADDRINUSE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EADDRINUSE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.EADDRNOTAVAIL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EADDRNOTAVAIL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.EADV))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EADV", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.EAFNOSUPPORT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EAFNOSUPPORT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.EAGAIN))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EAGAIN", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.EALREADY))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EALREADY", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.EBADE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EBADE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.EBADF))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EBADF", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.EBADFD))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EBADFD", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.EBADMSG))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EBADMSG", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.EBADR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EBADR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.EBADRQC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EBADRQC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.EBADSLT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EBADSLT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.EBFONT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EBFONT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.EBUSY))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EBUSY", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ECANCELED))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ECANCELED", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ECHILD))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ECHILD", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ECHO))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ECHO", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ECHOCTL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ECHOCTL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ECHOE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ECHOE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ECHOK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ECHOK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ECHOKE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ECHOKE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ECHONL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ECHONL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ECHOPRT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ECHOPRT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ECHRNG))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ECHRNG", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ECOMM))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ECOMM", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ECONNABORTED))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ECONNABORTED", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ECONNREFUSED))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ECONNREFUSED", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ECONNRESET))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ECONNRESET", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.EDEADLK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EDEADLK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.EDEADLOCK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EDEADLOCK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.EDESTADDRREQ))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EDESTADDRREQ", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.EDOM))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EDOM", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.EDOTDOT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EDOTDOT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.EDQUOT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EDQUOT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.EEXIST))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EEXIST", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.EFAULT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EFAULT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.EFBIG))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EFBIG", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.EHOSTDOWN))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EHOSTDOWN", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.EHOSTUNREACH))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EHOSTUNREACH", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.EIDRM))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EIDRM", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.EILSEQ))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EILSEQ", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.EINPROGRESS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EINPROGRESS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.EINTR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EINTR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.EINVAL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EINVAL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.EIO))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EIO", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.EISCONN))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EISCONN", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.EISDIR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EISDIR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.EISNAM))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EISNAM", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.EKEYEXPIRED))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EKEYEXPIRED", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.EKEYREJECTED))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EKEYREJECTED", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.EKEYREVOKED))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EKEYREVOKED", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.EL2HLT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EL2HLT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.EL2NSYNC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EL2NSYNC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.EL3HLT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EL3HLT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.EL3RST))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EL3RST", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ELIBACC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ELIBACC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ELIBBAD))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ELIBBAD", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ELIBEXEC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ELIBEXEC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ELIBMAX))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ELIBMAX", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ELIBSCN))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ELIBSCN", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ELNRNG))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ELNRNG", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ELOOP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ELOOP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.EMEDIUMTYPE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EMEDIUMTYPE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.EMFILE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EMFILE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.EMLINK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EMLINK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.EMSGSIZE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EMSGSIZE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.EMULTIHOP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EMULTIHOP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ENAMETOOLONG))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ENAMETOOLONG", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ENAVAIL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ENAVAIL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ENETDOWN))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ENETDOWN", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ENETRESET))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ENETRESET", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ENETUNREACH))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ENETUNREACH", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ENFILE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ENFILE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ENOANO))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ENOANO", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ENOBUFS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ENOBUFS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ENOCSI))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ENOCSI", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ENODATA))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ENODATA", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ENODEV))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ENODEV", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ENOENT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ENOENT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ENOEXEC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ENOEXEC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ENOKEY))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ENOKEY", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ENOLCK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ENOLCK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ENOLINK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ENOLINK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ENOMEDIUM))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ENOMEDIUM", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ENOMEM))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ENOMEM", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ENOMSG))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ENOMSG", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ENONET))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ENONET", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ENOPKG))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ENOPKG", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ENOPROTOOPT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ENOPROTOOPT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ENOSPC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ENOSPC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ENOSR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ENOSR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ENOSTR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ENOSTR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ENOSYS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ENOSYS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ENOTBLK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ENOTBLK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ENOTCONN))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ENOTCONN", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ENOTDIR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ENOTDIR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ENOTEMPTY))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ENOTEMPTY", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ENOTNAM))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ENOTNAM", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ENOTRECOVERABLE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ENOTRECOVERABLE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ENOTSOCK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ENOTSOCK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ENOTSUP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ENOTSUP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ENOTTY))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ENOTTY", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ENOTUNIQ))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ENOTUNIQ", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ENXIO))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ENXIO", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.EOPNOTSUPP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EOPNOTSUPP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.EOVERFLOW))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EOVERFLOW", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.EOWNERDEAD))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EOWNERDEAD", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.EPERM))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EPERM", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.EPFNOSUPPORT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EPFNOSUPPORT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.EPIPE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EPIPE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.EPOLLERR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EPOLLERR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(int64(mod.EPOLLET))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EPOLLET", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.EPOLLHUP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EPOLLHUP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.EPOLLIN))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EPOLLIN", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.EPOLLMSG))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EPOLLMSG", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.EPOLLONESHOT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EPOLLONESHOT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.EPOLLOUT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EPOLLOUT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.EPOLLPRI))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EPOLLPRI", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.EPOLLRDBAND))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EPOLLRDBAND", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.EPOLLRDHUP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EPOLLRDHUP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.EPOLLRDNORM))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EPOLLRDNORM", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.EPOLLWRBAND))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EPOLLWRBAND", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.EPOLLWRNORM))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EPOLLWRNORM", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.EPOLL_CLOEXEC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EPOLL_CLOEXEC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.EPOLL_CTL_ADD))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EPOLL_CTL_ADD", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.EPOLL_CTL_DEL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EPOLL_CTL_DEL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.EPOLL_CTL_MOD))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EPOLL_CTL_MOD", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.EPOLL_NONBLOCK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EPOLL_NONBLOCK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.EPROTO))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EPROTO", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.EPROTONOSUPPORT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EPROTONOSUPPORT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.EPROTOTYPE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EPROTOTYPE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ERANGE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ERANGE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.EREMCHG))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EREMCHG", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.EREMOTE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EREMOTE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.EREMOTEIO))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EREMOTEIO", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ERESTART))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ERESTART", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ERFKILL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ERFKILL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.EROFS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EROFS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ESHUTDOWN))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ESHUTDOWN", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ESOCKTNOSUPPORT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ESOCKTNOSUPPORT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ESPIPE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ESPIPE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ESRCH))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ESRCH", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ESRMNT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ESRMNT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ESTALE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ESTALE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ESTRPIPE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ESTRPIPE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ETH_P_1588))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ETH_P_1588", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ETH_P_8021Q))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ETH_P_8021Q", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ETH_P_802_2))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ETH_P_802_2", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ETH_P_802_3))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ETH_P_802_3", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ETH_P_AARP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ETH_P_AARP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ETH_P_ALL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ETH_P_ALL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ETH_P_AOE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ETH_P_AOE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ETH_P_ARCNET))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ETH_P_ARCNET", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ETH_P_ARP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ETH_P_ARP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ETH_P_ATALK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ETH_P_ATALK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ETH_P_ATMFATE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ETH_P_ATMFATE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ETH_P_ATMMPOA))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ETH_P_ATMMPOA", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ETH_P_AX25))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ETH_P_AX25", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ETH_P_BPQ))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ETH_P_BPQ", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ETH_P_CAIF))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ETH_P_CAIF", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ETH_P_CAN))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ETH_P_CAN", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ETH_P_CONTROL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ETH_P_CONTROL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ETH_P_CUST))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ETH_P_CUST", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ETH_P_DDCMP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ETH_P_DDCMP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ETH_P_DEC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ETH_P_DEC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ETH_P_DIAG))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ETH_P_DIAG", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ETH_P_DNA_DL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ETH_P_DNA_DL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ETH_P_DNA_RC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ETH_P_DNA_RC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ETH_P_DNA_RT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ETH_P_DNA_RT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ETH_P_DSA))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ETH_P_DSA", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ETH_P_ECONET))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ETH_P_ECONET", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ETH_P_EDSA))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ETH_P_EDSA", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ETH_P_FCOE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ETH_P_FCOE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ETH_P_FIP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ETH_P_FIP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ETH_P_HDLC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ETH_P_HDLC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ETH_P_IEEE802154))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ETH_P_IEEE802154", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ETH_P_IEEEPUP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ETH_P_IEEEPUP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ETH_P_IEEEPUPAT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ETH_P_IEEEPUPAT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ETH_P_IP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ETH_P_IP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ETH_P_IPV6))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ETH_P_IPV6", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ETH_P_IPX))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ETH_P_IPX", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ETH_P_IRDA))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ETH_P_IRDA", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ETH_P_LAT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ETH_P_LAT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ETH_P_LINK_CTL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ETH_P_LINK_CTL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ETH_P_LOCALTALK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ETH_P_LOCALTALK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ETH_P_LOOP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ETH_P_LOOP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ETH_P_MOBITEX))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ETH_P_MOBITEX", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ETH_P_MPLS_MC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ETH_P_MPLS_MC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ETH_P_MPLS_UC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ETH_P_MPLS_UC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ETH_P_PAE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ETH_P_PAE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ETH_P_PAUSE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ETH_P_PAUSE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ETH_P_PHONET))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ETH_P_PHONET", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ETH_P_PPPTALK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ETH_P_PPPTALK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ETH_P_PPP_DISC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ETH_P_PPP_DISC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ETH_P_PPP_MP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ETH_P_PPP_MP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ETH_P_PPP_SES))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ETH_P_PPP_SES", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ETH_P_PUP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ETH_P_PUP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ETH_P_PUPAT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ETH_P_PUPAT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ETH_P_RARP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ETH_P_RARP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ETH_P_SCA))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ETH_P_SCA", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ETH_P_SLOW))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ETH_P_SLOW", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ETH_P_SNAP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ETH_P_SNAP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ETH_P_TEB))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ETH_P_TEB", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ETH_P_TIPC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ETH_P_TIPC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ETH_P_TRAILER))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ETH_P_TRAILER", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ETH_P_TR_802_2))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ETH_P_TR_802_2", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ETH_P_WAN_PPP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ETH_P_WAN_PPP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ETH_P_WCCP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ETH_P_WCCP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ETH_P_X25))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ETH_P_X25", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ETIME))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ETIME", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ETIMEDOUT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ETIMEDOUT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ETOOMANYREFS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ETOOMANYREFS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ETXTBSY))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ETXTBSY", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.EUCLEAN))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EUCLEAN", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.EUNATCH))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EUNATCH", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.EUSERS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EUSERS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.EWOULDBLOCK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EWOULDBLOCK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.EXDEV))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EXDEV", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.EXFULL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EXFULL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Environ)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Environ", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.EpollCreate)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EpollCreate", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.EpollCreate1)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EpollCreate1", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.EpollCtl)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EpollCtl", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.EpollEvent
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "EpollEvent", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.EpollWait)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EpollWait", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.Errno
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Errno", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Exec)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Exec", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Exit)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Exit", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.FD_CLOEXEC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "FD_CLOEXEC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.FD_SETSIZE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "FD_SETSIZE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.FLUSHO))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "FLUSHO", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.F_DUPFD))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "F_DUPFD", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.F_DUPFD_CLOEXEC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "F_DUPFD_CLOEXEC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.F_EXLCK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "F_EXLCK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.F_GETFD))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "F_GETFD", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.F_GETFL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "F_GETFL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.F_GETLEASE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "F_GETLEASE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.F_GETLK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "F_GETLK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.F_GETLK64))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "F_GETLK64", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.F_GETOWN))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "F_GETOWN", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.F_GETOWN_EX))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "F_GETOWN_EX", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.F_GETPIPE_SZ))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "F_GETPIPE_SZ", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.F_GETSIG))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "F_GETSIG", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.F_LOCK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "F_LOCK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.F_NOTIFY))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "F_NOTIFY", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.F_OK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "F_OK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.F_RDLCK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "F_RDLCK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.F_SETFD))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "F_SETFD", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.F_SETFL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "F_SETFL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.F_SETLEASE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "F_SETLEASE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.F_SETLK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "F_SETLK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.F_SETLK64))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "F_SETLK64", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.F_SETLKW))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "F_SETLKW", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.F_SETLKW64))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "F_SETLKW64", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.F_SETOWN))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "F_SETOWN", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.F_SETOWN_EX))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "F_SETOWN_EX", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.F_SETPIPE_SZ))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "F_SETPIPE_SZ", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.F_SETSIG))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "F_SETSIG", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.F_SHLCK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "F_SHLCK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.F_TEST))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "F_TEST", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.F_TLOCK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "F_TLOCK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.F_ULOCK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "F_ULOCK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.F_UNLCK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "F_UNLCK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.F_WRLCK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "F_WRLCK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Faccessat)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Faccessat", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Fallocate)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Fallocate", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Fchdir)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Fchdir", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Fchmod)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Fchmod", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Fchmodat)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Fchmodat", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Fchown)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Fchown", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Fchownat)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Fchownat", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.FcntlFlock)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "FcntlFlock", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.FdSet
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "FdSet", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Fdatasync)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Fdatasync", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Flock)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Flock", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.Flock_t
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Flock_t", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.ForkExec)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ForkExec", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.ForkLock)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ForkLock", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.Fsid
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Fsid", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Fstat)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Fstat", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Fstatfs)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Fstatfs", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Fsync)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Fsync", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Ftruncate)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Ftruncate", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Futimes)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Futimes", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Futimesat)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Futimesat", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Getcwd)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Getcwd", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Getdents)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Getdents", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Getegid)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Getegid", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Getenv)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Getenv", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Geteuid)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Geteuid", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Getgid)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Getgid", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Getgroups)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Getgroups", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Getpagesize)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Getpagesize", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Getpeername)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Getpeername", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Getpgid)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Getpgid", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Getpgrp)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Getpgrp", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Getpid)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Getpid", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Getppid)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Getppid", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Getpriority)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Getpriority", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Getrlimit)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Getrlimit", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Getrusage)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Getrusage", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Getsockname)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Getsockname", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.GetsockoptICMPv6Filter)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "GetsockoptICMPv6Filter", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.GetsockoptIPMreq)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "GetsockoptIPMreq", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.GetsockoptIPMreqn)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "GetsockoptIPMreqn", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.GetsockoptIPv6MTUInfo)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "GetsockoptIPv6MTUInfo", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.GetsockoptIPv6Mreq)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "GetsockoptIPv6Mreq", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.GetsockoptInet4Addr)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "GetsockoptInet4Addr", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.GetsockoptInt)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "GetsockoptInt", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.GetsockoptUcred)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "GetsockoptUcred", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Gettid)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Gettid", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Gettimeofday)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Gettimeofday", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Getuid)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Getuid", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Getwd)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Getwd", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Getxattr)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Getxattr", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.HUPCL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "HUPCL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ICANON))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ICANON", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ICMPV6_FILTER))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ICMPV6_FILTER", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.ICMPv6Filter
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "ICMPv6Filter", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.ICRNL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ICRNL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.IEXTEN))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IEXTEN", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IFA_ADDRESS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFA_ADDRESS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IFA_ANYCAST))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFA_ANYCAST", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IFA_BROADCAST))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFA_BROADCAST", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IFA_CACHEINFO))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFA_CACHEINFO", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IFA_F_DADFAILED))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFA_F_DADFAILED", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IFA_F_DEPRECATED))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFA_F_DEPRECATED", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IFA_F_HOMEADDRESS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFA_F_HOMEADDRESS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IFA_F_NODAD))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFA_F_NODAD", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IFA_F_OPTIMISTIC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFA_F_OPTIMISTIC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.IFA_F_PERMANENT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFA_F_PERMANENT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IFA_F_SECONDARY))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFA_F_SECONDARY", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IFA_F_TEMPORARY))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFA_F_TEMPORARY", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IFA_F_TENTATIVE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFA_F_TENTATIVE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IFA_LABEL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFA_LABEL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IFA_LOCAL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFA_LOCAL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IFA_MAX))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFA_MAX", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IFA_MULTICAST))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFA_MULTICAST", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IFA_UNSPEC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFA_UNSPEC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.IFF_ALLMULTI))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFF_ALLMULTI", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.IFF_AUTOMEDIA))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFF_AUTOMEDIA", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IFF_BROADCAST))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFF_BROADCAST", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IFF_DEBUG))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFF_DEBUG", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.IFF_DYNAMIC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFF_DYNAMIC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IFF_LOOPBACK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFF_LOOPBACK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.IFF_MASTER))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFF_MASTER", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.IFF_MULTICAST))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFF_MULTICAST", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.IFF_NOARP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFF_NOARP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IFF_NOTRAILERS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFF_NOTRAILERS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.IFF_NO_PI))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFF_NO_PI", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.IFF_ONE_QUEUE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFF_ONE_QUEUE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IFF_POINTOPOINT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFF_POINTOPOINT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.IFF_PORTSEL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFF_PORTSEL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.IFF_PROMISC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFF_PROMISC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IFF_RUNNING))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFF_RUNNING", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.IFF_SLAVE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFF_SLAVE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IFF_TAP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFF_TAP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IFF_TUN))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFF_TUN", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.IFF_TUN_EXCL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFF_TUN_EXCL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IFF_UP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFF_UP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.IFF_VNET_HDR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFF_VNET_HDR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IFLA_ADDRESS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFLA_ADDRESS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IFLA_BROADCAST))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFLA_BROADCAST", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IFLA_COST))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFLA_COST", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IFLA_IFALIAS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFLA_IFALIAS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IFLA_IFNAME))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFLA_IFNAME", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IFLA_LINK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFLA_LINK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IFLA_LINKINFO))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFLA_LINKINFO", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IFLA_LINKMODE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFLA_LINKMODE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IFLA_MAP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFLA_MAP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IFLA_MASTER))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFLA_MASTER", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IFLA_MAX))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFLA_MAX", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IFLA_MTU))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFLA_MTU", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IFLA_NET_NS_PID))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFLA_NET_NS_PID", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IFLA_OPERSTATE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFLA_OPERSTATE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IFLA_PRIORITY))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFLA_PRIORITY", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IFLA_PROTINFO))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFLA_PROTINFO", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IFLA_QDISC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFLA_QDISC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IFLA_STATS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFLA_STATS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IFLA_TXQLEN))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFLA_TXQLEN", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IFLA_UNSPEC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFLA_UNSPEC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IFLA_WEIGHT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFLA_WEIGHT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IFLA_WIRELESS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFLA_WIRELESS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IFNAMSIZ))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFNAMSIZ", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IGNBRK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IGNBRK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.IGNCR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IGNCR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IGNPAR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IGNPAR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.IMAXBEL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IMAXBEL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.INLCR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "INLCR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.INPCK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "INPCK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IN_ACCESS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IN_ACCESS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.IN_ALL_EVENTS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IN_ALL_EVENTS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IN_ATTRIB))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IN_ATTRIB", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.IN_CLASSA_HOST))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IN_CLASSA_HOST", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.IN_CLASSA_MAX))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IN_CLASSA_MAX", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint64(mod.IN_CLASSA_NET))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IN_CLASSA_NET", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IN_CLASSA_NSHIFT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IN_CLASSA_NSHIFT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.IN_CLASSB_HOST))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IN_CLASSB_HOST", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.IN_CLASSB_MAX))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IN_CLASSB_MAX", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint64(mod.IN_CLASSB_NET))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IN_CLASSB_NET", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IN_CLASSB_NSHIFT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IN_CLASSB_NSHIFT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.IN_CLASSC_HOST))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IN_CLASSC_HOST", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint64(mod.IN_CLASSC_NET))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IN_CLASSC_NET", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IN_CLASSC_NSHIFT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IN_CLASSC_NSHIFT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.IN_CLOEXEC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IN_CLOEXEC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IN_CLOSE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IN_CLOSE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IN_CLOSE_NOWRITE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IN_CLOSE_NOWRITE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IN_CLOSE_WRITE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IN_CLOSE_WRITE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.IN_CREATE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IN_CREATE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.IN_DELETE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IN_DELETE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.IN_DELETE_SELF))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IN_DELETE_SELF", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.IN_DONT_FOLLOW))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IN_DONT_FOLLOW", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.IN_EXCL_UNLINK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IN_EXCL_UNLINK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.IN_IGNORED))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IN_IGNORED", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.IN_ISDIR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IN_ISDIR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.IN_LOOPBACKNET))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IN_LOOPBACKNET", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.IN_MASK_ADD))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IN_MASK_ADD", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IN_MODIFY))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IN_MODIFY", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.IN_MOVE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IN_MOVE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IN_MOVED_FROM))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IN_MOVED_FROM", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.IN_MOVED_TO))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IN_MOVED_TO", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.IN_MOVE_SELF))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IN_MOVE_SELF", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.IN_NONBLOCK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IN_NONBLOCK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint64(mod.IN_ONESHOT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IN_ONESHOT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.IN_ONLYDIR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IN_ONLYDIR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IN_OPEN))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IN_OPEN", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.IN_Q_OVERFLOW))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IN_Q_OVERFLOW", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.IN_UNMOUNT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IN_UNMOUNT", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.IPMreq
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "IPMreq", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.IPMreqn
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "IPMreqn", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IPPROTO_AH))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_AH", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IPPROTO_COMP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_COMP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IPPROTO_DCCP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_DCCP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IPPROTO_DSTOPTS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_DSTOPTS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IPPROTO_EGP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_EGP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IPPROTO_ENCAP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_ENCAP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IPPROTO_ESP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_ESP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IPPROTO_FRAGMENT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_FRAGMENT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IPPROTO_GRE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_GRE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IPPROTO_HOPOPTS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_HOPOPTS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IPPROTO_ICMP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_ICMP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IPPROTO_ICMPV6))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_ICMPV6", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IPPROTO_IDP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_IDP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IPPROTO_IGMP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_IGMP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IPPROTO_IP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_IP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IPPROTO_IPIP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_IPIP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IPPROTO_IPV6))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_IPV6", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IPPROTO_MTP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_MTP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IPPROTO_NONE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_NONE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IPPROTO_PIM))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_PIM", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IPPROTO_PUP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_PUP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.IPPROTO_RAW))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_RAW", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IPPROTO_ROUTING))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_ROUTING", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IPPROTO_RSVP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_RSVP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.IPPROTO_SCTP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_SCTP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IPPROTO_TCP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_TCP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IPPROTO_TP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_TP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IPPROTO_UDP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_UDP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.IPPROTO_UDPLITE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_UDPLITE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IPV6_2292DSTOPTS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_2292DSTOPTS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IPV6_2292HOPLIMIT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_2292HOPLIMIT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IPV6_2292HOPOPTS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_2292HOPOPTS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IPV6_2292PKTINFO))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_2292PKTINFO", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IPV6_2292PKTOPTIONS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_2292PKTOPTIONS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IPV6_2292RTHDR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_2292RTHDR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IPV6_ADDRFORM))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_ADDRFORM", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IPV6_ADD_MEMBERSHIP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_ADD_MEMBERSHIP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IPV6_AUTHHDR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_AUTHHDR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IPV6_CHECKSUM))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_CHECKSUM", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IPV6_DROP_MEMBERSHIP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_DROP_MEMBERSHIP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IPV6_DSTOPTS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_DSTOPTS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IPV6_HOPLIMIT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_HOPLIMIT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IPV6_HOPOPTS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_HOPOPTS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IPV6_IPSEC_POLICY))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_IPSEC_POLICY", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IPV6_JOIN_ANYCAST))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_JOIN_ANYCAST", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IPV6_JOIN_GROUP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_JOIN_GROUP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IPV6_LEAVE_ANYCAST))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_LEAVE_ANYCAST", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IPV6_LEAVE_GROUP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_LEAVE_GROUP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IPV6_MTU))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_MTU", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IPV6_MTU_DISCOVER))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_MTU_DISCOVER", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IPV6_MULTICAST_HOPS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_MULTICAST_HOPS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IPV6_MULTICAST_IF))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_MULTICAST_IF", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IPV6_MULTICAST_LOOP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_MULTICAST_LOOP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IPV6_NEXTHOP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_NEXTHOP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IPV6_PKTINFO))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_PKTINFO", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IPV6_PMTUDISC_DO))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_PMTUDISC_DO", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IPV6_PMTUDISC_DONT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_PMTUDISC_DONT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IPV6_PMTUDISC_PROBE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_PMTUDISC_PROBE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IPV6_PMTUDISC_WANT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_PMTUDISC_WANT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IPV6_RECVDSTOPTS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_RECVDSTOPTS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IPV6_RECVERR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_RECVERR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IPV6_RECVHOPLIMIT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_RECVHOPLIMIT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IPV6_RECVHOPOPTS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_RECVHOPOPTS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IPV6_RECVPKTINFO))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_RECVPKTINFO", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IPV6_RECVRTHDR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_RECVRTHDR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IPV6_RECVTCLASS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_RECVTCLASS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IPV6_ROUTER_ALERT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_ROUTER_ALERT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IPV6_RTHDR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_RTHDR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IPV6_RTHDRDSTOPTS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_RTHDRDSTOPTS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IPV6_RTHDR_LOOSE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_RTHDR_LOOSE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IPV6_RTHDR_STRICT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_RTHDR_STRICT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IPV6_RTHDR_TYPE_0))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_RTHDR_TYPE_0", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IPV6_RXDSTOPTS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_RXDSTOPTS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IPV6_RXHOPOPTS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_RXHOPOPTS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IPV6_TCLASS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_TCLASS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IPV6_UNICAST_HOPS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_UNICAST_HOPS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IPV6_V6ONLY))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_V6ONLY", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IPV6_XFRM_POLICY))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_XFRM_POLICY", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IP_ADD_MEMBERSHIP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_ADD_MEMBERSHIP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IP_ADD_SOURCE_MEMBERSHIP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_ADD_SOURCE_MEMBERSHIP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IP_BLOCK_SOURCE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_BLOCK_SOURCE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IP_DEFAULT_MULTICAST_LOOP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_DEFAULT_MULTICAST_LOOP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IP_DEFAULT_MULTICAST_TTL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_DEFAULT_MULTICAST_TTL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.IP_DF))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_DF", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IP_DROP_MEMBERSHIP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_DROP_MEMBERSHIP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IP_DROP_SOURCE_MEMBERSHIP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_DROP_SOURCE_MEMBERSHIP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IP_FREEBIND))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_FREEBIND", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IP_HDRINCL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_HDRINCL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IP_IPSEC_POLICY))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_IPSEC_POLICY", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.IP_MAXPACKET))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_MAXPACKET", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IP_MAX_MEMBERSHIPS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_MAX_MEMBERSHIPS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.IP_MF))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_MF", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IP_MINTTL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_MINTTL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IP_MSFILTER))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_MSFILTER", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.IP_MSS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_MSS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IP_MTU))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_MTU", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IP_MTU_DISCOVER))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_MTU_DISCOVER", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IP_MULTICAST_IF))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_MULTICAST_IF", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IP_MULTICAST_LOOP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_MULTICAST_LOOP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IP_MULTICAST_TTL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_MULTICAST_TTL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.IP_OFFMASK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_OFFMASK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IP_OPTIONS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_OPTIONS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IP_ORIGDSTADDR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_ORIGDSTADDR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IP_PASSSEC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_PASSSEC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IP_PKTINFO))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_PKTINFO", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IP_PKTOPTIONS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_PKTOPTIONS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IP_PMTUDISC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_PMTUDISC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IP_PMTUDISC_DO))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_PMTUDISC_DO", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IP_PMTUDISC_DONT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_PMTUDISC_DONT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IP_PMTUDISC_PROBE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_PMTUDISC_PROBE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IP_PMTUDISC_WANT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_PMTUDISC_WANT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IP_RECVERR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_RECVERR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IP_RECVOPTS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_RECVOPTS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IP_RECVORIGDSTADDR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_RECVORIGDSTADDR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IP_RECVRETOPTS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_RECVRETOPTS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IP_RECVTOS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_RECVTOS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IP_RECVTTL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_RECVTTL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IP_RETOPTS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_RETOPTS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.IP_RF))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_RF", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IP_ROUTER_ALERT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_ROUTER_ALERT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IP_TOS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_TOS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IP_TRANSPARENT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_TRANSPARENT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IP_TTL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_TTL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IP_UNBLOCK_SOURCE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_UNBLOCK_SOURCE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.IP_XFRM_POLICY))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_XFRM_POLICY", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.IPv6MTUInfo
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "IPv6MTUInfo", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.IPv6Mreq
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "IPv6Mreq", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ISIG))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ISIG", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ISTRIP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ISTRIP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.IUCLC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IUCLC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.IUTF8))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IUTF8", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.IXANY))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IXANY", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.IXOFF))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IXOFF", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.IXON))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IXON", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.IfAddrmsg
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "IfAddrmsg", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.IfInfomsg
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "IfInfomsg", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.ImplementsGetwd)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ImplementsGetwd", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.Inet4Pktinfo
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Inet4Pktinfo", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.Inet6Pktinfo
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Inet6Pktinfo", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.InotifyAddWatch)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "InotifyAddWatch", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.InotifyEvent
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "InotifyEvent", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.InotifyInit)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "InotifyInit", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.InotifyInit1)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "InotifyInit1", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.InotifyRmWatch)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "InotifyRmWatch", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Ioperm)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Ioperm", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Iopl)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Iopl", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.Iovec
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Iovec", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Kill)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Kill", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Klogctl)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Klogctl", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.LINUX_REBOOT_CMD_CAD_OFF))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "LINUX_REBOOT_CMD_CAD_OFF", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint64(mod.LINUX_REBOOT_CMD_CAD_ON))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "LINUX_REBOOT_CMD_CAD_ON", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint64(mod.LINUX_REBOOT_CMD_HALT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "LINUX_REBOOT_CMD_HALT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.LINUX_REBOOT_CMD_KEXEC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "LINUX_REBOOT_CMD_KEXEC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.LINUX_REBOOT_CMD_POWER_OFF))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "LINUX_REBOOT_CMD_POWER_OFF", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.LINUX_REBOOT_CMD_RESTART))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "LINUX_REBOOT_CMD_RESTART", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint64(mod.LINUX_REBOOT_CMD_RESTART2))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "LINUX_REBOOT_CMD_RESTART2", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint64(mod.LINUX_REBOOT_CMD_SW_SUSPEND))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "LINUX_REBOOT_CMD_SW_SUSPEND", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint64(mod.LINUX_REBOOT_MAGIC1))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "LINUX_REBOOT_MAGIC1", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.LINUX_REBOOT_MAGIC2))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "LINUX_REBOOT_MAGIC2", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.LOCK_EX))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "LOCK_EX", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.LOCK_NB))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "LOCK_NB", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.LOCK_SH))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "LOCK_SH", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.LOCK_UN))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "LOCK_UN", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Lchown)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Lchown", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.Linger
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Linger", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Link)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Link", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Listen)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Listen", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Listxattr)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Listxattr", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.LsfJump)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "LsfJump", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.LsfSocket)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "LsfSocket", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.LsfStmt)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "LsfStmt", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Lstat)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Lstat", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.MADV_DOFORK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MADV_DOFORK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.MADV_DONTFORK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MADV_DONTFORK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.MADV_DONTNEED))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MADV_DONTNEED", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.MADV_HUGEPAGE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MADV_HUGEPAGE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.MADV_HWPOISON))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MADV_HWPOISON", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.MADV_MERGEABLE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MADV_MERGEABLE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.MADV_NOHUGEPAGE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MADV_NOHUGEPAGE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.MADV_NORMAL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MADV_NORMAL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.MADV_RANDOM))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MADV_RANDOM", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.MADV_REMOVE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MADV_REMOVE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.MADV_SEQUENTIAL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MADV_SEQUENTIAL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.MADV_UNMERGEABLE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MADV_UNMERGEABLE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.MADV_WILLNEED))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MADV_WILLNEED", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.MAP_32BIT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MAP_32BIT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.MAP_ANON))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MAP_ANON", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.MAP_ANONYMOUS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MAP_ANONYMOUS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.MAP_DENYWRITE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MAP_DENYWRITE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.MAP_EXECUTABLE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MAP_EXECUTABLE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.MAP_FILE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MAP_FILE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.MAP_FIXED))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MAP_FIXED", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.MAP_GROWSDOWN))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MAP_GROWSDOWN", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.MAP_HUGETLB))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MAP_HUGETLB", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.MAP_LOCKED))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MAP_LOCKED", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.MAP_NONBLOCK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MAP_NONBLOCK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.MAP_NORESERVE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MAP_NORESERVE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.MAP_POPULATE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MAP_POPULATE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.MAP_PRIVATE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MAP_PRIVATE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.MAP_SHARED))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MAP_SHARED", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.MAP_STACK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MAP_STACK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.MAP_TYPE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MAP_TYPE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.MCL_CURRENT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MCL_CURRENT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.MCL_FUTURE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MCL_FUTURE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.MNT_DETACH))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MNT_DETACH", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.MNT_EXPIRE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MNT_EXPIRE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.MNT_FORCE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MNT_FORCE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.MSG_CMSG_CLOEXEC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MSG_CMSG_CLOEXEC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.MSG_CONFIRM))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MSG_CONFIRM", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.MSG_CTRUNC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MSG_CTRUNC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.MSG_DONTROUTE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MSG_DONTROUTE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.MSG_DONTWAIT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MSG_DONTWAIT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.MSG_EOR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MSG_EOR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.MSG_ERRQUEUE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MSG_ERRQUEUE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.MSG_FASTOPEN))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MSG_FASTOPEN", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.MSG_FIN))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MSG_FIN", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.MSG_MORE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MSG_MORE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.MSG_NOSIGNAL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MSG_NOSIGNAL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.MSG_OOB))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MSG_OOB", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.MSG_PEEK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MSG_PEEK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.MSG_PROXY))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MSG_PROXY", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.MSG_RST))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MSG_RST", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.MSG_SYN))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MSG_SYN", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.MSG_TRUNC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MSG_TRUNC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.MSG_TRYHARD))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MSG_TRYHARD", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.MSG_WAITALL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MSG_WAITALL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.MSG_WAITFORONE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MSG_WAITFORONE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.MS_ACTIVE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MS_ACTIVE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.MS_ASYNC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MS_ASYNC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.MS_BIND))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MS_BIND", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.MS_DIRSYNC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MS_DIRSYNC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.MS_INVALIDATE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MS_INVALIDATE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.MS_I_VERSION))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MS_I_VERSION", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.MS_KERNMOUNT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MS_KERNMOUNT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.MS_MANDLOCK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MS_MANDLOCK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint64(mod.MS_MGC_MSK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MS_MGC_MSK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint64(mod.MS_MGC_VAL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MS_MGC_VAL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.MS_MOVE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MS_MOVE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.MS_NOATIME))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MS_NOATIME", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.MS_NODEV))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MS_NODEV", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.MS_NODIRATIME))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MS_NODIRATIME", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.MS_NOEXEC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MS_NOEXEC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.MS_NOSUID))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MS_NOSUID", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(int64(mod.MS_NOUSER))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MS_NOUSER", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.MS_POSIXACL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MS_POSIXACL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.MS_PRIVATE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MS_PRIVATE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.MS_RDONLY))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MS_RDONLY", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.MS_REC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MS_REC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.MS_RELATIME))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MS_RELATIME", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.MS_REMOUNT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MS_REMOUNT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.MS_RMT_MASK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MS_RMT_MASK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.MS_SHARED))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MS_SHARED", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.MS_SILENT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MS_SILENT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.MS_SLAVE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MS_SLAVE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.MS_STRICTATIME))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MS_STRICTATIME", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.MS_SYNC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MS_SYNC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.MS_SYNCHRONOUS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MS_SYNCHRONOUS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.MS_UNBINDABLE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MS_UNBINDABLE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Madvise)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Madvise", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Mkdir)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Mkdir", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Mkdirat)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Mkdirat", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Mkfifo)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Mkfifo", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Mknod)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Mknod", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Mknodat)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Mknodat", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Mlock)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Mlock", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Mlockall)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Mlockall", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Mmap)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Mmap", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Mount)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Mount", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Mprotect)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Mprotect", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.Msghdr
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Msghdr", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Munlock)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Munlock", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Munlockall)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Munlockall", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Munmap)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Munmap", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.NAME_MAX))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NAME_MAX", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.NETLINK_ADD_MEMBERSHIP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NETLINK_ADD_MEMBERSHIP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.NETLINK_AUDIT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NETLINK_AUDIT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.NETLINK_BROADCAST_ERROR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NETLINK_BROADCAST_ERROR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.NETLINK_CONNECTOR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NETLINK_CONNECTOR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.NETLINK_DNRTMSG))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NETLINK_DNRTMSG", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.NETLINK_DROP_MEMBERSHIP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NETLINK_DROP_MEMBERSHIP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.NETLINK_ECRYPTFS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NETLINK_ECRYPTFS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.NETLINK_FIB_LOOKUP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NETLINK_FIB_LOOKUP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.NETLINK_FIREWALL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NETLINK_FIREWALL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.NETLINK_GENERIC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NETLINK_GENERIC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.NETLINK_INET_DIAG))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NETLINK_INET_DIAG", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.NETLINK_IP6_FW))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NETLINK_IP6_FW", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.NETLINK_ISCSI))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NETLINK_ISCSI", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.NETLINK_KOBJECT_UEVENT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NETLINK_KOBJECT_UEVENT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.NETLINK_NETFILTER))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NETLINK_NETFILTER", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.NETLINK_NFLOG))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NETLINK_NFLOG", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.NETLINK_NO_ENOBUFS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NETLINK_NO_ENOBUFS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.NETLINK_PKTINFO))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NETLINK_PKTINFO", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.NETLINK_ROUTE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NETLINK_ROUTE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.NETLINK_SCSITRANSPORT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NETLINK_SCSITRANSPORT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.NETLINK_SELINUX))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NETLINK_SELINUX", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.NETLINK_UNUSED))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NETLINK_UNUSED", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.NETLINK_USERSOCK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NETLINK_USERSOCK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.NETLINK_XFRM))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NETLINK_XFRM", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.NLA_ALIGNTO))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NLA_ALIGNTO", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.NLA_F_NESTED))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NLA_F_NESTED", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.NLA_F_NET_BYTEORDER))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NLA_F_NET_BYTEORDER", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.NLA_HDRLEN))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NLA_HDRLEN", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.NLMSG_ALIGNTO))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NLMSG_ALIGNTO", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.NLMSG_DONE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NLMSG_DONE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.NLMSG_ERROR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NLMSG_ERROR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.NLMSG_HDRLEN))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NLMSG_HDRLEN", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.NLMSG_MIN_TYPE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NLMSG_MIN_TYPE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.NLMSG_NOOP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NLMSG_NOOP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.NLMSG_OVERRUN))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NLMSG_OVERRUN", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.NLM_F_ACK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NLM_F_ACK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.NLM_F_APPEND))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NLM_F_APPEND", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.NLM_F_ATOMIC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NLM_F_ATOMIC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.NLM_F_CREATE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NLM_F_CREATE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.NLM_F_DUMP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NLM_F_DUMP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.NLM_F_ECHO))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NLM_F_ECHO", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.NLM_F_EXCL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NLM_F_EXCL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.NLM_F_MATCH))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NLM_F_MATCH", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.NLM_F_MULTI))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NLM_F_MULTI", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.NLM_F_REPLACE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NLM_F_REPLACE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.NLM_F_REQUEST))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NLM_F_REQUEST", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.NLM_F_ROOT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NLM_F_ROOT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.NOFLSH))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NOFLSH", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Nanosleep)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Nanosleep", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.NetlinkMessage
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "NetlinkMessage", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.NetlinkRIB)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NetlinkRIB", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.NetlinkRouteAttr
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "NetlinkRouteAttr", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.NetlinkRouteRequest
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "NetlinkRouteRequest", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.NlAttr
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "NlAttr", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.NlMsgerr
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "NlMsgerr", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.NlMsghdr
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "NlMsghdr", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.NsecToTimespec)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NsecToTimespec", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.NsecToTimeval)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NsecToTimeval", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.OCRNL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "OCRNL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.OFDEL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "OFDEL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.OFILL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "OFILL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.OLCUC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "OLCUC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ONLCR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ONLCR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ONLRET))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ONLRET", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ONOCR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ONOCR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.OPOST))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "OPOST", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.O_ACCMODE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "O_ACCMODE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.O_APPEND))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "O_APPEND", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.O_ASYNC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "O_ASYNC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.O_CLOEXEC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "O_CLOEXEC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.O_CREAT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "O_CREAT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.O_DIRECT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "O_DIRECT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.O_DIRECTORY))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "O_DIRECTORY", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.O_DSYNC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "O_DSYNC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.O_EXCL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "O_EXCL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.O_FSYNC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "O_FSYNC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.O_LARGEFILE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "O_LARGEFILE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.O_NDELAY))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "O_NDELAY", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.O_NOATIME))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "O_NOATIME", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.O_NOCTTY))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "O_NOCTTY", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.O_NOFOLLOW))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "O_NOFOLLOW", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.O_NONBLOCK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "O_NONBLOCK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.O_RDONLY))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "O_RDONLY", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.O_RDWR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "O_RDWR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.O_RSYNC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "O_RSYNC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.O_SYNC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "O_SYNC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.O_TRUNC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "O_TRUNC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.O_WRONLY))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "O_WRONLY", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Open)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Open", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Openat)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Openat", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PACKET_ADD_MEMBERSHIP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PACKET_ADD_MEMBERSHIP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PACKET_BROADCAST))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PACKET_BROADCAST", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PACKET_DROP_MEMBERSHIP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PACKET_DROP_MEMBERSHIP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PACKET_FASTROUTE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PACKET_FASTROUTE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PACKET_HOST))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PACKET_HOST", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PACKET_LOOPBACK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PACKET_LOOPBACK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PACKET_MR_ALLMULTI))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PACKET_MR_ALLMULTI", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PACKET_MR_MULTICAST))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PACKET_MR_MULTICAST", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PACKET_MR_PROMISC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PACKET_MR_PROMISC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PACKET_MULTICAST))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PACKET_MULTICAST", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PACKET_OTHERHOST))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PACKET_OTHERHOST", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PACKET_OUTGOING))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PACKET_OUTGOING", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PACKET_RECV_OUTPUT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PACKET_RECV_OUTPUT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PACKET_RX_RING))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PACKET_RX_RING", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PACKET_STATISTICS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PACKET_STATISTICS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.PARENB))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PARENB", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PARMRK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PARMRK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.PARODD))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PARODD", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.PENDIN))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PENDIN", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PRIO_PGRP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PRIO_PGRP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PRIO_PROCESS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PRIO_PROCESS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PRIO_USER))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PRIO_USER", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PROT_EXEC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PROT_EXEC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.PROT_GROWSDOWN))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PROT_GROWSDOWN", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.PROT_GROWSUP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PROT_GROWSUP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PROT_NONE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PROT_NONE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PROT_READ))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PROT_READ", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PROT_WRITE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PROT_WRITE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PR_CAPBSET_DROP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PR_CAPBSET_DROP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PR_CAPBSET_READ))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PR_CAPBSET_READ", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PR_ENDIAN_BIG))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PR_ENDIAN_BIG", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PR_ENDIAN_LITTLE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PR_ENDIAN_LITTLE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PR_ENDIAN_PPC_LITTLE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PR_ENDIAN_PPC_LITTLE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PR_FPEMU_NOPRINT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PR_FPEMU_NOPRINT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PR_FPEMU_SIGFPE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PR_FPEMU_SIGFPE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PR_FP_EXC_ASYNC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PR_FP_EXC_ASYNC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PR_FP_EXC_DISABLED))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PR_FP_EXC_DISABLED", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.PR_FP_EXC_DIV))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PR_FP_EXC_DIV", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.PR_FP_EXC_INV))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PR_FP_EXC_INV", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PR_FP_EXC_NONRECOV))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PR_FP_EXC_NONRECOV", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.PR_FP_EXC_OVF))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PR_FP_EXC_OVF", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PR_FP_EXC_PRECISE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PR_FP_EXC_PRECISE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.PR_FP_EXC_RES))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PR_FP_EXC_RES", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.PR_FP_EXC_SW_ENABLE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PR_FP_EXC_SW_ENABLE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.PR_FP_EXC_UND))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PR_FP_EXC_UND", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PR_GET_DUMPABLE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PR_GET_DUMPABLE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PR_GET_ENDIAN))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PR_GET_ENDIAN", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PR_GET_FPEMU))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PR_GET_FPEMU", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PR_GET_FPEXC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PR_GET_FPEXC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PR_GET_KEEPCAPS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PR_GET_KEEPCAPS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PR_GET_NAME))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PR_GET_NAME", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PR_GET_PDEATHSIG))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PR_GET_PDEATHSIG", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PR_GET_SECCOMP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PR_GET_SECCOMP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PR_GET_SECUREBITS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PR_GET_SECUREBITS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PR_GET_TIMERSLACK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PR_GET_TIMERSLACK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PR_GET_TIMING))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PR_GET_TIMING", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PR_GET_TSC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PR_GET_TSC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PR_GET_UNALIGN))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PR_GET_UNALIGN", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PR_MCE_KILL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PR_MCE_KILL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PR_MCE_KILL_CLEAR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PR_MCE_KILL_CLEAR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PR_MCE_KILL_DEFAULT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PR_MCE_KILL_DEFAULT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PR_MCE_KILL_EARLY))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PR_MCE_KILL_EARLY", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PR_MCE_KILL_GET))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PR_MCE_KILL_GET", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PR_MCE_KILL_LATE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PR_MCE_KILL_LATE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PR_MCE_KILL_SET))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PR_MCE_KILL_SET", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PR_SET_DUMPABLE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PR_SET_DUMPABLE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PR_SET_ENDIAN))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PR_SET_ENDIAN", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PR_SET_FPEMU))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PR_SET_FPEMU", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PR_SET_FPEXC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PR_SET_FPEXC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PR_SET_KEEPCAPS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PR_SET_KEEPCAPS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PR_SET_NAME))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PR_SET_NAME", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PR_SET_PDEATHSIG))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PR_SET_PDEATHSIG", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.PR_SET_PTRACER))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PR_SET_PTRACER", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PR_SET_SECCOMP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PR_SET_SECCOMP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PR_SET_SECUREBITS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PR_SET_SECUREBITS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PR_SET_TIMERSLACK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PR_SET_TIMERSLACK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PR_SET_TIMING))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PR_SET_TIMING", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PR_SET_TSC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PR_SET_TSC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PR_SET_UNALIGN))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PR_SET_UNALIGN", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PR_TASK_PERF_EVENTS_DISABLE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PR_TASK_PERF_EVENTS_DISABLE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PR_TASK_PERF_EVENTS_ENABLE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PR_TASK_PERF_EVENTS_ENABLE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PR_TIMING_STATISTICAL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PR_TIMING_STATISTICAL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PR_TIMING_TIMESTAMP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PR_TIMING_TIMESTAMP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PR_TSC_ENABLE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PR_TSC_ENABLE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PR_TSC_SIGSEGV))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PR_TSC_SIGSEGV", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PR_UNALIGN_NOPRINT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PR_UNALIGN_NOPRINT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PR_UNALIGN_SIGBUS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PR_UNALIGN_SIGBUS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PTRACE_ARCH_PRCTL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PTRACE_ARCH_PRCTL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PTRACE_ATTACH))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PTRACE_ATTACH", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PTRACE_CONT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PTRACE_CONT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PTRACE_DETACH))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PTRACE_DETACH", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PTRACE_EVENT_CLONE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PTRACE_EVENT_CLONE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PTRACE_EVENT_EXEC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PTRACE_EVENT_EXEC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PTRACE_EVENT_EXIT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PTRACE_EVENT_EXIT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PTRACE_EVENT_FORK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PTRACE_EVENT_FORK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PTRACE_EVENT_VFORK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PTRACE_EVENT_VFORK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PTRACE_EVENT_VFORK_DONE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PTRACE_EVENT_VFORK_DONE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.PTRACE_GETEVENTMSG))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PTRACE_GETEVENTMSG", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PTRACE_GETFPREGS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PTRACE_GETFPREGS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PTRACE_GETFPXREGS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PTRACE_GETFPXREGS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PTRACE_GETREGS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PTRACE_GETREGS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.PTRACE_GETREGSET))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PTRACE_GETREGSET", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.PTRACE_GETSIGINFO))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PTRACE_GETSIGINFO", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PTRACE_GET_THREAD_AREA))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PTRACE_GET_THREAD_AREA", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PTRACE_KILL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PTRACE_KILL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PTRACE_OLDSETOPTIONS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PTRACE_OLDSETOPTIONS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.PTRACE_O_MASK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PTRACE_O_MASK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PTRACE_O_TRACECLONE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PTRACE_O_TRACECLONE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PTRACE_O_TRACEEXEC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PTRACE_O_TRACEEXEC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PTRACE_O_TRACEEXIT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PTRACE_O_TRACEEXIT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PTRACE_O_TRACEFORK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PTRACE_O_TRACEFORK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PTRACE_O_TRACESYSGOOD))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PTRACE_O_TRACESYSGOOD", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PTRACE_O_TRACEVFORK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PTRACE_O_TRACEVFORK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PTRACE_O_TRACEVFORKDONE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PTRACE_O_TRACEVFORKDONE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PTRACE_PEEKDATA))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PTRACE_PEEKDATA", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PTRACE_PEEKTEXT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PTRACE_PEEKTEXT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PTRACE_PEEKUSR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PTRACE_PEEKUSR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PTRACE_POKEDATA))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PTRACE_POKEDATA", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PTRACE_POKETEXT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PTRACE_POKETEXT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PTRACE_POKEUSR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PTRACE_POKEUSR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PTRACE_SETFPREGS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PTRACE_SETFPREGS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PTRACE_SETFPXREGS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PTRACE_SETFPXREGS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.PTRACE_SETOPTIONS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PTRACE_SETOPTIONS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PTRACE_SETREGS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PTRACE_SETREGS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.PTRACE_SETREGSET))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PTRACE_SETREGSET", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.PTRACE_SETSIGINFO))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PTRACE_SETSIGINFO", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PTRACE_SET_THREAD_AREA))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PTRACE_SET_THREAD_AREA", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PTRACE_SINGLEBLOCK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PTRACE_SINGLEBLOCK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PTRACE_SINGLESTEP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PTRACE_SINGLESTEP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PTRACE_SYSCALL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PTRACE_SYSCALL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PTRACE_SYSEMU))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PTRACE_SYSEMU", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PTRACE_SYSEMU_SINGLESTEP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PTRACE_SYSEMU_SINGLESTEP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.PTRACE_TRACEME))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PTRACE_TRACEME", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.ParseDirent)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ParseDirent", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.ParseNetlinkMessage)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ParseNetlinkMessage", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.ParseNetlinkRouteAttr)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ParseNetlinkRouteAttr", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.ParseSocketControlMessage)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ParseSocketControlMessage", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.ParseUnixCredentials)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ParseUnixCredentials", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.ParseUnixRights)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ParseUnixRights", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.PathMax))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PathMax", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Pause)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Pause", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Pipe)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Pipe", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Pipe2)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Pipe2", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.PivotRoot)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PivotRoot", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Pread)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Pread", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.ProcAttr
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "ProcAttr", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.PtraceAttach)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PtraceAttach", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.PtraceCont)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PtraceCont", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.PtraceDetach)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PtraceDetach", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.PtraceGetEventMsg)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PtraceGetEventMsg", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.PtraceGetRegs)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PtraceGetRegs", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.PtracePeekData)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PtracePeekData", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.PtracePeekText)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PtracePeekText", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.PtracePokeData)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PtracePokeData", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.PtracePokeText)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PtracePokeText", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.PtraceRegs
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "PtraceRegs", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.PtraceSetOptions)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PtraceSetOptions", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.PtraceSetRegs)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PtraceSetRegs", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.PtraceSingleStep)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PtraceSingleStep", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.PtraceSyscall)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PtraceSyscall", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Pwrite)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Pwrite", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RLIMIT_AS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RLIMIT_AS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RLIMIT_CORE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RLIMIT_CORE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RLIMIT_CPU))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RLIMIT_CPU", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RLIMIT_DATA))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RLIMIT_DATA", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RLIMIT_FSIZE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RLIMIT_FSIZE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RLIMIT_NOFILE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RLIMIT_NOFILE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RLIMIT_STACK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RLIMIT_STACK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(int(mod.RLIM_INFINITY))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RLIM_INFINITY", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTAX_ADVMSS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTAX_ADVMSS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTAX_CWND))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTAX_CWND", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTAX_FEATURES))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTAX_FEATURES", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTAX_FEATURE_ALLFRAG))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTAX_FEATURE_ALLFRAG", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTAX_FEATURE_ECN))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTAX_FEATURE_ECN", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTAX_FEATURE_SACK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTAX_FEATURE_SACK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTAX_FEATURE_TIMESTAMP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTAX_FEATURE_TIMESTAMP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTAX_HOPLIMIT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTAX_HOPLIMIT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTAX_INITCWND))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTAX_INITCWND", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTAX_INITRWND))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTAX_INITRWND", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTAX_LOCK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTAX_LOCK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTAX_MAX))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTAX_MAX", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTAX_MTU))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTAX_MTU", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTAX_REORDERING))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTAX_REORDERING", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTAX_RTO_MIN))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTAX_RTO_MIN", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTAX_RTT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTAX_RTT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTAX_RTTVAR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTAX_RTTVAR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTAX_SSTHRESH))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTAX_SSTHRESH", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTAX_UNSPEC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTAX_UNSPEC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTAX_WINDOW))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTAX_WINDOW", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTA_ALIGNTO))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTA_ALIGNTO", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTA_CACHEINFO))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTA_CACHEINFO", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTA_DST))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTA_DST", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTA_FLOW))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTA_FLOW", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTA_GATEWAY))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTA_GATEWAY", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTA_IIF))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTA_IIF", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTA_MAX))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTA_MAX", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTA_METRICS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTA_METRICS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTA_MULTIPATH))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTA_MULTIPATH", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTA_OIF))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTA_OIF", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTA_PREFSRC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTA_PREFSRC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTA_PRIORITY))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTA_PRIORITY", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTA_SRC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTA_SRC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTA_TABLE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTA_TABLE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTA_UNSPEC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTA_UNSPEC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.RTCF_DIRECTSRC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTCF_DIRECTSRC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.RTCF_DOREDIRECT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTCF_DOREDIRECT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.RTCF_LOG))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTCF_LOG", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.RTCF_MASQ))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTCF_MASQ", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.RTCF_NAT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTCF_NAT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.RTCF_VALVE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTCF_VALVE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint64(mod.RTF_ADDRCLASSMASK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTF_ADDRCLASSMASK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.RTF_ADDRCONF))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTF_ADDRCONF", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.RTF_ALLONLINK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTF_ALLONLINK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.RTF_BROADCAST))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTF_BROADCAST", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.RTF_CACHE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTF_CACHE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.RTF_DEFAULT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTF_DEFAULT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTF_DYNAMIC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTF_DYNAMIC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.RTF_FLOW))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTF_FLOW", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTF_GATEWAY))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTF_GATEWAY", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTF_HOST))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTF_HOST", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.RTF_INTERFACE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTF_INTERFACE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.RTF_IRTT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTF_IRTT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.RTF_LINKRT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTF_LINKRT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint64(mod.RTF_LOCAL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTF_LOCAL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTF_MODIFIED))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTF_MODIFIED", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTF_MSS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTF_MSS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTF_MTU))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTF_MTU", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.RTF_MULTICAST))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTF_MULTICAST", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.RTF_NAT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTF_NAT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.RTF_NOFORWARD))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTF_NOFORWARD", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.RTF_NONEXTHOP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTF_NONEXTHOP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.RTF_NOPMTUDISC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTF_NOPMTUDISC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.RTF_POLICY))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTF_POLICY", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTF_REINSTATE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTF_REINSTATE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.RTF_REJECT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTF_REJECT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.RTF_STATIC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTF_STATIC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.RTF_THROW))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTF_THROW", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTF_UP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTF_UP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.RTF_WINDOW))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTF_WINDOW", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.RTF_XRESOLVE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTF_XRESOLVE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTM_BASE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTM_BASE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTM_DELACTION))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTM_DELACTION", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTM_DELADDR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTM_DELADDR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTM_DELADDRLABEL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTM_DELADDRLABEL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTM_DELLINK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTM_DELLINK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTM_DELNEIGH))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTM_DELNEIGH", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTM_DELQDISC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTM_DELQDISC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTM_DELROUTE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTM_DELROUTE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTM_DELRULE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTM_DELRULE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTM_DELTCLASS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTM_DELTCLASS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTM_DELTFILTER))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTM_DELTFILTER", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.RTM_F_CLONED))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTM_F_CLONED", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.RTM_F_EQUALIZE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTM_F_EQUALIZE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.RTM_F_NOTIFY))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTM_F_NOTIFY", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.RTM_F_PREFIX))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTM_F_PREFIX", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTM_GETACTION))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTM_GETACTION", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTM_GETADDR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTM_GETADDR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTM_GETADDRLABEL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTM_GETADDRLABEL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTM_GETANYCAST))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTM_GETANYCAST", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTM_GETDCB))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTM_GETDCB", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTM_GETLINK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTM_GETLINK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTM_GETMULTICAST))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTM_GETMULTICAST", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTM_GETNEIGH))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTM_GETNEIGH", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTM_GETNEIGHTBL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTM_GETNEIGHTBL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTM_GETQDISC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTM_GETQDISC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTM_GETROUTE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTM_GETROUTE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTM_GETRULE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTM_GETRULE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTM_GETTCLASS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTM_GETTCLASS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTM_GETTFILTER))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTM_GETTFILTER", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTM_MAX))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTM_MAX", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTM_NEWACTION))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTM_NEWACTION", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTM_NEWADDR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTM_NEWADDR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTM_NEWADDRLABEL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTM_NEWADDRLABEL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTM_NEWLINK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTM_NEWLINK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTM_NEWNDUSEROPT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTM_NEWNDUSEROPT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTM_NEWNEIGH))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTM_NEWNEIGH", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTM_NEWNEIGHTBL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTM_NEWNEIGHTBL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTM_NEWPREFIX))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTM_NEWPREFIX", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTM_NEWQDISC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTM_NEWQDISC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTM_NEWROUTE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTM_NEWROUTE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTM_NEWRULE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTM_NEWRULE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTM_NEWTCLASS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTM_NEWTCLASS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTM_NEWTFILTER))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTM_NEWTFILTER", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTM_NR_FAMILIES))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTM_NR_FAMILIES", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTM_NR_MSGTYPES))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTM_NR_MSGTYPES", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTM_SETDCB))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTM_SETDCB", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTM_SETLINK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTM_SETLINK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTM_SETNEIGHTBL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTM_SETNEIGHTBL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTNH_ALIGNTO))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTNH_ALIGNTO", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTNH_F_DEAD))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTNH_F_DEAD", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTNH_F_ONLINK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTNH_F_ONLINK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTNH_F_PERVASIVE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTNH_F_PERVASIVE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTNLGRP_IPV4_IFADDR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTNLGRP_IPV4_IFADDR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTNLGRP_IPV4_MROUTE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTNLGRP_IPV4_MROUTE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTNLGRP_IPV4_ROUTE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTNLGRP_IPV4_ROUTE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTNLGRP_IPV4_RULE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTNLGRP_IPV4_RULE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTNLGRP_IPV6_IFADDR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTNLGRP_IPV6_IFADDR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTNLGRP_IPV6_IFINFO))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTNLGRP_IPV6_IFINFO", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTNLGRP_IPV6_MROUTE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTNLGRP_IPV6_MROUTE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTNLGRP_IPV6_PREFIX))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTNLGRP_IPV6_PREFIX", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTNLGRP_IPV6_ROUTE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTNLGRP_IPV6_ROUTE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTNLGRP_IPV6_RULE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTNLGRP_IPV6_RULE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTNLGRP_LINK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTNLGRP_LINK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTNLGRP_ND_USEROPT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTNLGRP_ND_USEROPT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTNLGRP_NEIGH))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTNLGRP_NEIGH", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTNLGRP_NONE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTNLGRP_NONE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTNLGRP_NOTIFY))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTNLGRP_NOTIFY", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTNLGRP_TC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTNLGRP_TC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTN_ANYCAST))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTN_ANYCAST", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTN_BLACKHOLE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTN_BLACKHOLE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTN_BROADCAST))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTN_BROADCAST", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTN_LOCAL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTN_LOCAL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTN_MAX))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTN_MAX", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTN_MULTICAST))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTN_MULTICAST", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTN_NAT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTN_NAT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTN_PROHIBIT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTN_PROHIBIT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTN_THROW))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTN_THROW", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTN_UNICAST))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTN_UNICAST", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTN_UNREACHABLE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTN_UNREACHABLE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTN_UNSPEC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTN_UNSPEC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTN_XRESOLVE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTN_XRESOLVE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTPROT_BIRD))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTPROT_BIRD", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTPROT_BOOT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTPROT_BOOT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTPROT_DHCP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTPROT_DHCP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTPROT_DNROUTED))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTPROT_DNROUTED", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTPROT_GATED))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTPROT_GATED", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTPROT_KERNEL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTPROT_KERNEL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTPROT_MRT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTPROT_MRT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTPROT_NTK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTPROT_NTK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTPROT_RA))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTPROT_RA", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTPROT_REDIRECT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTPROT_REDIRECT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTPROT_STATIC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTPROT_STATIC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTPROT_UNSPEC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTPROT_UNSPEC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTPROT_XORP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTPROT_XORP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RTPROT_ZEBRA))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTPROT_ZEBRA", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.RT_CLASS_DEFAULT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RT_CLASS_DEFAULT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.RT_CLASS_LOCAL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RT_CLASS_LOCAL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.RT_CLASS_MAIN))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RT_CLASS_MAIN", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.RT_CLASS_MAX))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RT_CLASS_MAX", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RT_CLASS_UNSPEC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RT_CLASS_UNSPEC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.RT_SCOPE_HOST))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RT_SCOPE_HOST", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.RT_SCOPE_LINK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RT_SCOPE_LINK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.RT_SCOPE_NOWHERE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RT_SCOPE_NOWHERE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.RT_SCOPE_SITE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RT_SCOPE_SITE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RT_SCOPE_UNIVERSE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RT_SCOPE_UNIVERSE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.RT_TABLE_COMPAT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RT_TABLE_COMPAT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.RT_TABLE_DEFAULT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RT_TABLE_DEFAULT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.RT_TABLE_LOCAL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RT_TABLE_LOCAL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.RT_TABLE_MAIN))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RT_TABLE_MAIN", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint64(mod.RT_TABLE_MAX))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RT_TABLE_MAX", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RT_TABLE_UNSPEC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RT_TABLE_UNSPEC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(int(mod.RUSAGE_CHILDREN))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RUSAGE_CHILDREN", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RUSAGE_SELF))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RUSAGE_SELF", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.RUSAGE_THREAD))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RUSAGE_THREAD", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.RawSockaddr
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "RawSockaddr", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.RawSockaddrAny
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "RawSockaddrAny", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.RawSockaddrInet4
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "RawSockaddrInet4", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.RawSockaddrInet6
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "RawSockaddrInet6", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.RawSockaddrLinklayer
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "RawSockaddrLinklayer", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.RawSockaddrNetlink
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "RawSockaddrNetlink", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.RawSockaddrUnix
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "RawSockaddrUnix", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.RawSyscall)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RawSyscall", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.RawSyscall6)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RawSyscall6", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Read)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Read", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.ReadDirent)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ReadDirent", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Readlink)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Readlink", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Reboot)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Reboot", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Recvfrom)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Recvfrom", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Recvmsg)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Recvmsg", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Removexattr)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Removexattr", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Rename)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Rename", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Renameat)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Renameat", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.Rlimit
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Rlimit", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Rmdir)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Rmdir", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.RtAttr
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "RtAttr", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.RtGenmsg
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "RtGenmsg", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.RtMsg
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "RtMsg", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.RtNexthop
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "RtNexthop", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.Rusage
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Rusage", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SCM_CREDENTIALS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SCM_CREDENTIALS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SCM_RIGHTS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SCM_RIGHTS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SCM_TIMESTAMP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SCM_TIMESTAMP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SCM_TIMESTAMPING))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SCM_TIMESTAMPING", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SCM_TIMESTAMPNS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SCM_TIMESTAMPNS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SHUT_RD))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SHUT_RD", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SHUT_RDWR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SHUT_RDWR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SHUT_WR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SHUT_WR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SIGABRT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIGABRT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SIGALRM))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIGALRM", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SIGBUS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIGBUS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SIGCHLD))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIGCHLD", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SIGCLD))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIGCLD", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SIGCONT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIGCONT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SIGFPE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIGFPE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SIGHUP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIGHUP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SIGILL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIGILL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SIGINT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIGINT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SIGIO))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIGIO", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SIGIOT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIGIOT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SIGKILL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIGKILL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SIGPIPE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIGPIPE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SIGPOLL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIGPOLL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SIGPROF))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIGPROF", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SIGPWR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIGPWR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SIGQUIT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIGQUIT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SIGSEGV))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIGSEGV", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SIGSTKFLT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIGSTKFLT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SIGSTOP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIGSTOP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SIGSYS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIGSYS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SIGTERM))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIGTERM", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SIGTRAP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIGTRAP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SIGTSTP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIGTSTP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SIGTTIN))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIGTTIN", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SIGTTOU))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIGTTOU", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SIGUNUSED))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIGUNUSED", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SIGURG))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIGURG", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SIGUSR1))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIGUSR1", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SIGUSR2))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIGUSR2", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SIGVTALRM))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIGVTALRM", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SIGWINCH))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIGWINCH", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SIGXCPU))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIGXCPU", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SIGXFSZ))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIGXFSZ", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SIOCADDDLCI))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCADDDLCI", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SIOCADDMULTI))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCADDMULTI", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SIOCADDRT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCADDRT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SIOCATMARK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCATMARK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SIOCDARP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCDARP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SIOCDELDLCI))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCDELDLCI", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SIOCDELMULTI))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCDELMULTI", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SIOCDELRT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCDELRT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SIOCDEVPRIVATE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCDEVPRIVATE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SIOCDIFADDR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCDIFADDR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SIOCDRARP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCDRARP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SIOCGARP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCGARP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SIOCGIFADDR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCGIFADDR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SIOCGIFBR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCGIFBR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SIOCGIFBRDADDR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCGIFBRDADDR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SIOCGIFCONF))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCGIFCONF", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SIOCGIFCOUNT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCGIFCOUNT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SIOCGIFDSTADDR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCGIFDSTADDR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SIOCGIFENCAP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCGIFENCAP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SIOCGIFFLAGS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCGIFFLAGS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SIOCGIFHWADDR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCGIFHWADDR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SIOCGIFINDEX))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCGIFINDEX", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SIOCGIFMAP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCGIFMAP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SIOCGIFMEM))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCGIFMEM", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SIOCGIFMETRIC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCGIFMETRIC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SIOCGIFMTU))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCGIFMTU", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SIOCGIFNAME))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCGIFNAME", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SIOCGIFNETMASK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCGIFNETMASK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SIOCGIFPFLAGS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCGIFPFLAGS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SIOCGIFSLAVE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCGIFSLAVE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SIOCGIFTXQLEN))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCGIFTXQLEN", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SIOCGPGRP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCGPGRP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SIOCGRARP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCGRARP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SIOCGSTAMP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCGSTAMP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SIOCGSTAMPNS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCGSTAMPNS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SIOCPROTOPRIVATE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCPROTOPRIVATE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SIOCRTMSG))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCRTMSG", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SIOCSARP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCSARP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SIOCSIFADDR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCSIFADDR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SIOCSIFBR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCSIFBR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SIOCSIFBRDADDR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCSIFBRDADDR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SIOCSIFDSTADDR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCSIFDSTADDR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SIOCSIFENCAP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCSIFENCAP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SIOCSIFFLAGS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCSIFFLAGS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SIOCSIFHWADDR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCSIFHWADDR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SIOCSIFHWBROADCAST))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCSIFHWBROADCAST", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SIOCSIFLINK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCSIFLINK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SIOCSIFMAP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCSIFMAP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SIOCSIFMEM))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCSIFMEM", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SIOCSIFMETRIC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCSIFMETRIC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SIOCSIFMTU))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCSIFMTU", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SIOCSIFNAME))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCSIFNAME", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SIOCSIFNETMASK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCSIFNETMASK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SIOCSIFPFLAGS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCSIFPFLAGS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SIOCSIFSLAVE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCSIFSLAVE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SIOCSIFTXQLEN))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCSIFTXQLEN", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SIOCSPGRP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCSPGRP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SIOCSRARP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCSRARP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SOCK_CLOEXEC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SOCK_CLOEXEC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SOCK_DCCP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SOCK_DCCP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SOCK_DGRAM))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SOCK_DGRAM", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SOCK_NONBLOCK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SOCK_NONBLOCK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SOCK_PACKET))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SOCK_PACKET", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SOCK_RAW))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SOCK_RAW", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SOCK_RDM))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SOCK_RDM", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SOCK_SEQPACKET))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SOCK_SEQPACKET", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SOCK_STREAM))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SOCK_STREAM", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SOL_AAL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SOL_AAL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SOL_ATM))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SOL_ATM", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SOL_DECNET))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SOL_DECNET", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SOL_ICMPV6))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SOL_ICMPV6", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SOL_IP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SOL_IP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SOL_IPV6))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SOL_IPV6", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SOL_IRDA))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SOL_IRDA", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SOL_PACKET))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SOL_PACKET", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SOL_RAW))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SOL_RAW", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SOL_SOCKET))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SOL_SOCKET", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SOL_TCP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SOL_TCP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SOL_X25))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SOL_X25", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SOMAXCONN))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SOMAXCONN", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SO_ACCEPTCONN))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SO_ACCEPTCONN", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SO_ATTACH_FILTER))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SO_ATTACH_FILTER", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SO_BINDTODEVICE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SO_BINDTODEVICE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SO_BROADCAST))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SO_BROADCAST", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SO_BSDCOMPAT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SO_BSDCOMPAT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SO_DEBUG))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SO_DEBUG", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SO_DETACH_FILTER))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SO_DETACH_FILTER", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SO_DOMAIN))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SO_DOMAIN", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SO_DONTROUTE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SO_DONTROUTE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SO_ERROR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SO_ERROR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SO_KEEPALIVE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SO_KEEPALIVE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SO_LINGER))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SO_LINGER", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SO_MARK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SO_MARK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SO_NO_CHECK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SO_NO_CHECK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SO_OOBINLINE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SO_OOBINLINE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SO_PASSCRED))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SO_PASSCRED", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SO_PASSSEC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SO_PASSSEC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SO_PEERCRED))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SO_PEERCRED", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SO_PEERNAME))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SO_PEERNAME", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SO_PEERSEC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SO_PEERSEC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SO_PRIORITY))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SO_PRIORITY", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SO_PROTOCOL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SO_PROTOCOL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SO_RCVBUF))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SO_RCVBUF", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SO_RCVBUFFORCE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SO_RCVBUFFORCE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SO_RCVLOWAT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SO_RCVLOWAT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SO_RCVTIMEO))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SO_RCVTIMEO", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SO_REUSEADDR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SO_REUSEADDR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SO_RXQ_OVFL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SO_RXQ_OVFL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SO_SECURITY_AUTHENTICATION))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SO_SECURITY_AUTHENTICATION", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SO_SECURITY_ENCRYPTION_NETWORK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SO_SECURITY_ENCRYPTION_NETWORK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SO_SECURITY_ENCRYPTION_TRANSPORT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SO_SECURITY_ENCRYPTION_TRANSPORT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SO_SNDBUF))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SO_SNDBUF", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SO_SNDBUFFORCE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SO_SNDBUFFORCE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SO_SNDLOWAT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SO_SNDLOWAT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SO_SNDTIMEO))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SO_SNDTIMEO", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SO_TIMESTAMP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SO_TIMESTAMP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SO_TIMESTAMPING))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SO_TIMESTAMPING", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SO_TIMESTAMPNS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SO_TIMESTAMPNS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SO_TYPE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SO_TYPE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_ACCEPT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_ACCEPT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_ACCEPT4))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_ACCEPT4", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_ACCESS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_ACCESS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_ACCT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_ACCT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_ADD_KEY))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_ADD_KEY", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_ADJTIMEX))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_ADJTIMEX", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_AFS_SYSCALL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_AFS_SYSCALL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_ALARM))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_ALARM", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_ARCH_PRCTL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_ARCH_PRCTL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_BIND))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_BIND", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_BRK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_BRK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_CAPGET))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_CAPGET", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_CAPSET))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_CAPSET", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_CHDIR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_CHDIR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_CHMOD))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_CHMOD", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_CHOWN))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_CHOWN", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_CHROOT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_CHROOT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_CLOCK_GETRES))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_CLOCK_GETRES", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_CLOCK_GETTIME))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_CLOCK_GETTIME", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_CLOCK_NANOSLEEP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_CLOCK_NANOSLEEP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_CLOCK_SETTIME))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_CLOCK_SETTIME", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_CLONE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_CLONE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_CLOSE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_CLOSE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_CONNECT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_CONNECT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_CREAT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_CREAT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_CREATE_MODULE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_CREATE_MODULE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_DELETE_MODULE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_DELETE_MODULE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_DUP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_DUP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_DUP2))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_DUP2", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_DUP3))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_DUP3", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_EPOLL_CREATE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_EPOLL_CREATE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_EPOLL_CREATE1))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_EPOLL_CREATE1", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_EPOLL_CTL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_EPOLL_CTL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_EPOLL_CTL_OLD))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_EPOLL_CTL_OLD", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_EPOLL_PWAIT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_EPOLL_PWAIT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_EPOLL_WAIT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_EPOLL_WAIT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_EPOLL_WAIT_OLD))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_EPOLL_WAIT_OLD", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_EVENTFD))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_EVENTFD", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_EVENTFD2))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_EVENTFD2", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_EXECVE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_EXECVE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_EXIT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_EXIT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_EXIT_GROUP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_EXIT_GROUP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_FACCESSAT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_FACCESSAT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_FADVISE64))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_FADVISE64", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_FALLOCATE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_FALLOCATE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_FANOTIFY_INIT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_FANOTIFY_INIT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_FANOTIFY_MARK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_FANOTIFY_MARK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_FCHDIR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_FCHDIR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_FCHMOD))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_FCHMOD", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_FCHMODAT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_FCHMODAT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_FCHOWN))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_FCHOWN", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_FCHOWNAT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_FCHOWNAT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_FCNTL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_FCNTL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_FDATASYNC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_FDATASYNC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_FGETXATTR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_FGETXATTR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_FLISTXATTR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_FLISTXATTR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_FLOCK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_FLOCK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_FORK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_FORK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_FREMOVEXATTR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_FREMOVEXATTR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_FSETXATTR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_FSETXATTR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_FSTAT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_FSTAT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_FSTATFS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_FSTATFS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_FSYNC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_FSYNC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_FTRUNCATE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_FTRUNCATE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_FUTEX))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_FUTEX", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_FUTIMESAT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_FUTIMESAT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_GETCWD))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_GETCWD", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_GETDENTS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_GETDENTS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_GETDENTS64))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_GETDENTS64", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_GETEGID))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_GETEGID", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_GETEUID))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_GETEUID", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_GETGID))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_GETGID", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_GETGROUPS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_GETGROUPS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_GETITIMER))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_GETITIMER", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_GETPEERNAME))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_GETPEERNAME", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_GETPGID))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_GETPGID", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_GETPGRP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_GETPGRP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_GETPID))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_GETPID", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_GETPMSG))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_GETPMSG", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_GETPPID))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_GETPPID", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_GETPRIORITY))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_GETPRIORITY", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_GETRESGID))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_GETRESGID", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_GETRESUID))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_GETRESUID", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_GETRLIMIT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_GETRLIMIT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_GETRUSAGE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_GETRUSAGE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_GETSID))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_GETSID", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_GETSOCKNAME))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_GETSOCKNAME", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_GETSOCKOPT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_GETSOCKOPT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_GETTID))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_GETTID", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_GETTIMEOFDAY))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_GETTIMEOFDAY", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_GETUID))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_GETUID", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_GETXATTR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_GETXATTR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_GET_KERNEL_SYMS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_GET_KERNEL_SYMS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_GET_MEMPOLICY))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_GET_MEMPOLICY", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_GET_ROBUST_LIST))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_GET_ROBUST_LIST", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_GET_THREAD_AREA))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_GET_THREAD_AREA", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_INIT_MODULE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_INIT_MODULE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_INOTIFY_ADD_WATCH))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_INOTIFY_ADD_WATCH", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_INOTIFY_INIT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_INOTIFY_INIT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_INOTIFY_INIT1))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_INOTIFY_INIT1", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_INOTIFY_RM_WATCH))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_INOTIFY_RM_WATCH", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_IOCTL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_IOCTL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_IOPERM))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_IOPERM", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_IOPL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_IOPL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_IOPRIO_GET))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_IOPRIO_GET", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_IOPRIO_SET))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_IOPRIO_SET", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_IO_CANCEL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_IO_CANCEL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_IO_DESTROY))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_IO_DESTROY", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_IO_GETEVENTS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_IO_GETEVENTS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_IO_SETUP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_IO_SETUP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_IO_SUBMIT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_IO_SUBMIT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_KEXEC_LOAD))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_KEXEC_LOAD", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_KEYCTL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_KEYCTL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_KILL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_KILL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_LCHOWN))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_LCHOWN", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_LGETXATTR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_LGETXATTR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_LINK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_LINK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_LINKAT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_LINKAT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_LISTEN))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_LISTEN", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_LISTXATTR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_LISTXATTR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_LLISTXATTR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_LLISTXATTR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_LOOKUP_DCOOKIE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_LOOKUP_DCOOKIE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_LREMOVEXATTR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_LREMOVEXATTR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_LSEEK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_LSEEK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_LSETXATTR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_LSETXATTR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_LSTAT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_LSTAT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_MADVISE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_MADVISE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_MBIND))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_MBIND", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_MIGRATE_PAGES))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_MIGRATE_PAGES", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_MINCORE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_MINCORE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_MKDIR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_MKDIR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_MKDIRAT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_MKDIRAT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_MKNOD))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_MKNOD", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_MKNODAT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_MKNODAT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_MLOCK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_MLOCK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_MLOCKALL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_MLOCKALL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_MMAP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_MMAP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_MODIFY_LDT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_MODIFY_LDT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_MOUNT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_MOUNT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_MOVE_PAGES))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_MOVE_PAGES", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_MPROTECT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_MPROTECT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_MQ_GETSETATTR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_MQ_GETSETATTR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_MQ_NOTIFY))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_MQ_NOTIFY", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_MQ_OPEN))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_MQ_OPEN", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_MQ_TIMEDRECEIVE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_MQ_TIMEDRECEIVE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_MQ_TIMEDSEND))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_MQ_TIMEDSEND", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_MQ_UNLINK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_MQ_UNLINK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_MREMAP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_MREMAP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_MSGCTL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_MSGCTL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_MSGGET))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_MSGGET", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_MSGRCV))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_MSGRCV", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_MSGSND))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_MSGSND", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_MSYNC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_MSYNC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_MUNLOCK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_MUNLOCK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_MUNLOCKALL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_MUNLOCKALL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_MUNMAP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_MUNMAP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_NANOSLEEP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_NANOSLEEP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_NEWFSTATAT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_NEWFSTATAT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_NFSSERVCTL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_NFSSERVCTL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_OPEN))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_OPEN", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_OPENAT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_OPENAT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_PAUSE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_PAUSE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_PERF_EVENT_OPEN))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_PERF_EVENT_OPEN", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_PERSONALITY))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_PERSONALITY", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_PIPE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_PIPE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_PIPE2))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_PIPE2", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_PIVOT_ROOT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_PIVOT_ROOT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_POLL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_POLL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_PPOLL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_PPOLL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_PRCTL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_PRCTL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_PREAD64))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_PREAD64", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_PREADV))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_PREADV", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_PRLIMIT64))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_PRLIMIT64", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_PSELECT6))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_PSELECT6", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_PTRACE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_PTRACE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_PUTPMSG))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_PUTPMSG", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_PWRITE64))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_PWRITE64", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_PWRITEV))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_PWRITEV", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_QUERY_MODULE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_QUERY_MODULE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_QUOTACTL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_QUOTACTL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_READ))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_READ", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_READAHEAD))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_READAHEAD", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_READLINK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_READLINK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_READLINKAT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_READLINKAT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_READV))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_READV", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_REBOOT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_REBOOT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_RECVFROM))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_RECVFROM", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_RECVMMSG))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_RECVMMSG", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_RECVMSG))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_RECVMSG", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_REMAP_FILE_PAGES))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_REMAP_FILE_PAGES", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_REMOVEXATTR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_REMOVEXATTR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_RENAME))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_RENAME", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_RENAMEAT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_RENAMEAT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_REQUEST_KEY))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_REQUEST_KEY", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_RESTART_SYSCALL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_RESTART_SYSCALL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_RMDIR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_RMDIR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_RT_SIGACTION))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_RT_SIGACTION", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_RT_SIGPENDING))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_RT_SIGPENDING", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_RT_SIGPROCMASK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_RT_SIGPROCMASK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_RT_SIGQUEUEINFO))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_RT_SIGQUEUEINFO", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_RT_SIGRETURN))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_RT_SIGRETURN", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_RT_SIGSUSPEND))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_RT_SIGSUSPEND", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_RT_SIGTIMEDWAIT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_RT_SIGTIMEDWAIT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_RT_TGSIGQUEUEINFO))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_RT_TGSIGQUEUEINFO", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_SCHED_GETAFFINITY))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SCHED_GETAFFINITY", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_SCHED_GETPARAM))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SCHED_GETPARAM", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_SCHED_GETSCHEDULER))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SCHED_GETSCHEDULER", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_SCHED_GET_PRIORITY_MAX))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SCHED_GET_PRIORITY_MAX", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_SCHED_GET_PRIORITY_MIN))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SCHED_GET_PRIORITY_MIN", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_SCHED_RR_GET_INTERVAL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SCHED_RR_GET_INTERVAL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_SCHED_SETAFFINITY))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SCHED_SETAFFINITY", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_SCHED_SETPARAM))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SCHED_SETPARAM", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_SCHED_SETSCHEDULER))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SCHED_SETSCHEDULER", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_SCHED_YIELD))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SCHED_YIELD", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_SECURITY))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SECURITY", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_SELECT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SELECT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_SEMCTL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SEMCTL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_SEMGET))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SEMGET", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_SEMOP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SEMOP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_SEMTIMEDOP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SEMTIMEDOP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_SENDFILE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SENDFILE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_SENDMSG))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SENDMSG", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_SENDTO))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SENDTO", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_SETDOMAINNAME))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SETDOMAINNAME", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_SETFSGID))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SETFSGID", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_SETFSUID))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SETFSUID", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_SETGID))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SETGID", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_SETGROUPS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SETGROUPS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_SETHOSTNAME))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SETHOSTNAME", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_SETITIMER))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SETITIMER", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_SETPGID))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SETPGID", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_SETPRIORITY))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SETPRIORITY", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_SETREGID))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SETREGID", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_SETRESGID))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SETRESGID", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_SETRESUID))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SETRESUID", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_SETREUID))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SETREUID", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_SETRLIMIT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SETRLIMIT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_SETSID))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SETSID", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_SETSOCKOPT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SETSOCKOPT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_SETTIMEOFDAY))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SETTIMEOFDAY", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_SETUID))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SETUID", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_SETXATTR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SETXATTR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_SET_MEMPOLICY))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SET_MEMPOLICY", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_SET_ROBUST_LIST))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SET_ROBUST_LIST", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_SET_THREAD_AREA))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SET_THREAD_AREA", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_SET_TID_ADDRESS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SET_TID_ADDRESS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_SHMAT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SHMAT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_SHMCTL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SHMCTL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_SHMDT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SHMDT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_SHMGET))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SHMGET", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_SHUTDOWN))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SHUTDOWN", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_SIGALTSTACK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SIGALTSTACK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_SIGNALFD))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SIGNALFD", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_SIGNALFD4))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SIGNALFD4", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_SOCKET))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SOCKET", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_SOCKETPAIR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SOCKETPAIR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_SPLICE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SPLICE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_STAT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_STAT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_STATFS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_STATFS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_SWAPOFF))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SWAPOFF", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_SWAPON))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SWAPON", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_SYMLINK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SYMLINK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_SYMLINKAT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SYMLINKAT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_SYNC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SYNC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_SYNC_FILE_RANGE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SYNC_FILE_RANGE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_SYSFS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SYSFS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_SYSINFO))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SYSINFO", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_SYSLOG))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SYSLOG", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_TEE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_TEE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_TGKILL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_TGKILL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_TIME))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_TIME", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_TIMERFD_CREATE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_TIMERFD_CREATE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_TIMERFD_GETTIME))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_TIMERFD_GETTIME", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_TIMERFD_SETTIME))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_TIMERFD_SETTIME", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_TIMER_CREATE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_TIMER_CREATE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_TIMER_DELETE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_TIMER_DELETE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_TIMER_GETOVERRUN))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_TIMER_GETOVERRUN", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_TIMER_GETTIME))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_TIMER_GETTIME", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_TIMER_SETTIME))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_TIMER_SETTIME", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_TIMES))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_TIMES", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_TKILL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_TKILL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_TRUNCATE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_TRUNCATE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_TUXCALL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_TUXCALL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_UMASK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_UMASK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_UMOUNT2))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_UMOUNT2", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_UNAME))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_UNAME", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_UNLINK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_UNLINK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_UNLINKAT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_UNLINKAT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_UNSHARE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_UNSHARE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_USELIB))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_USELIB", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_USTAT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_USTAT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_UTIME))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_UTIME", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_UTIMENSAT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_UTIMENSAT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_UTIMES))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_UTIMES", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_VFORK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_VFORK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_VHANGUP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_VHANGUP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_VMSPLICE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_VMSPLICE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_VSERVER))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_VSERVER", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_WAIT4))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_WAIT4", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS_WAITID))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_WAITID", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_WRITE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_WRITE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SYS_WRITEV))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_WRITEV", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.SYS__SYSCTL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS__SYSCTL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.S_BLKSIZE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "S_BLKSIZE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.S_IEXEC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "S_IEXEC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.S_IFBLK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "S_IFBLK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.S_IFCHR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "S_IFCHR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.S_IFDIR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "S_IFDIR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.S_IFIFO))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "S_IFIFO", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.S_IFLNK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "S_IFLNK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.S_IFMT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "S_IFMT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.S_IFREG))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "S_IFREG", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.S_IFSOCK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "S_IFSOCK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.S_IREAD))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "S_IREAD", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.S_IRGRP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "S_IRGRP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.S_IROTH))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "S_IROTH", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.S_IRUSR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "S_IRUSR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.S_IRWXG))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "S_IRWXG", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.S_IRWXO))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "S_IRWXO", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.S_IRWXU))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "S_IRWXU", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.S_ISGID))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "S_ISGID", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.S_ISUID))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "S_ISUID", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.S_ISVTX))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "S_ISVTX", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.S_IWGRP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "S_IWGRP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.S_IWOTH))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "S_IWOTH", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.S_IWRITE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "S_IWRITE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.S_IWUSR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "S_IWUSR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.S_IXGRP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "S_IXGRP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.S_IXOTH))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "S_IXOTH", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.S_IXUSR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "S_IXUSR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Seek)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Seek", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Select)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Select", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Sendfile)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Sendfile", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Sendmsg)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Sendmsg", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.SendmsgN)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SendmsgN", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Sendto)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Sendto", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.SetLsfPromisc)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SetLsfPromisc", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.SetNonblock)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SetNonblock", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Setdomainname)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Setdomainname", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Setenv)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Setenv", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Setfsgid)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Setfsgid", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Setfsuid)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Setfsuid", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Setgid)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Setgid", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Setgroups)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Setgroups", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Sethostname)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Sethostname", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Setpgid)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Setpgid", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Setpriority)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Setpriority", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Setregid)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Setregid", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Setresgid)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Setresgid", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Setresuid)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Setresuid", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Setreuid)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Setreuid", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Setrlimit)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Setrlimit", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Setsid)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Setsid", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.SetsockoptByte)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SetsockoptByte", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.SetsockoptICMPv6Filter)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SetsockoptICMPv6Filter", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.SetsockoptIPMreq)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SetsockoptIPMreq", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.SetsockoptIPMreqn)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SetsockoptIPMreqn", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.SetsockoptIPv6Mreq)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SetsockoptIPv6Mreq", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.SetsockoptInet4Addr)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SetsockoptInet4Addr", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.SetsockoptInt)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SetsockoptInt", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.SetsockoptLinger)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SetsockoptLinger", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.SetsockoptString)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SetsockoptString", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.SetsockoptTimeval)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SetsockoptTimeval", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Settimeofday)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Settimeofday", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Setuid)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Setuid", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Setxattr)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Setxattr", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Shutdown)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Shutdown", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.Signal
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Signal", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SizeofCmsghdr))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SizeofCmsghdr", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SizeofICMPv6Filter))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SizeofICMPv6Filter", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SizeofIPMreq))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SizeofIPMreq", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SizeofIPMreqn))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SizeofIPMreqn", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SizeofIPv6MTUInfo))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SizeofIPv6MTUInfo", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SizeofIPv6Mreq))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SizeofIPv6Mreq", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SizeofIfAddrmsg))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SizeofIfAddrmsg", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SizeofIfInfomsg))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SizeofIfInfomsg", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SizeofInet4Pktinfo))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SizeofInet4Pktinfo", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SizeofInet6Pktinfo))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SizeofInet6Pktinfo", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SizeofInotifyEvent))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SizeofInotifyEvent", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SizeofLinger))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SizeofLinger", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SizeofMsghdr))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SizeofMsghdr", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SizeofNlAttr))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SizeofNlAttr", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SizeofNlMsgerr))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SizeofNlMsgerr", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SizeofNlMsghdr))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SizeofNlMsghdr", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SizeofRtAttr))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SizeofRtAttr", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SizeofRtGenmsg))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SizeofRtGenmsg", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SizeofRtMsg))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SizeofRtMsg", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SizeofRtNexthop))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SizeofRtNexthop", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SizeofSockFilter))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SizeofSockFilter", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SizeofSockFprog))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SizeofSockFprog", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SizeofSockaddrAny))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SizeofSockaddrAny", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SizeofSockaddrInet4))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SizeofSockaddrInet4", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SizeofSockaddrInet6))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SizeofSockaddrInet6", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SizeofSockaddrLinklayer))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SizeofSockaddrLinklayer", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SizeofSockaddrNetlink))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SizeofSockaddrNetlink", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SizeofSockaddrUnix))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SizeofSockaddrUnix", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SizeofTCPInfo))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SizeofTCPInfo", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.SizeofUcred))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SizeofUcred", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.SlicePtrFromStrings)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SlicePtrFromStrings", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.SockFilter
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "SockFilter", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.SockFprog
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "SockFprog", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.SockaddrInet4
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "SockaddrInet4", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.SockaddrInet6
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "SockaddrInet6", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.SockaddrLinklayer
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "SockaddrLinklayer", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.SockaddrNetlink
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "SockaddrNetlink", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.SockaddrUnix
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "SockaddrUnix", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Socket)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Socket", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.SocketControlMessage
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "SocketControlMessage", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.SocketDisableIPv6)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SocketDisableIPv6", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Socketpair)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Socketpair", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Splice)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Splice", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.StartProcess)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "StartProcess", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Stat)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Stat", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.Stat_t
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Stat_t", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Statfs)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Statfs", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.Statfs_t
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Statfs_t", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Stderr)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Stderr", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Stdin)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Stdin", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Stdout)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Stdout", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.StringBytePtr)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "StringBytePtr", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.StringByteSlice)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "StringByteSlice", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.StringSlicePtr)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "StringSlicePtr", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Symlink)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Symlink", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Sync)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Sync", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.SyncFileRange)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SyncFileRange", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.SysProcAttr
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "SysProcAttr", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.SysProcIDMap
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "SysProcIDMap", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Syscall)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Syscall", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Syscall6)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Syscall6", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Sysinfo)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Sysinfo", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.Sysinfo_t
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Sysinfo_t", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.TCGETS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TCGETS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.TCIFLUSH))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TCIFLUSH", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.TCIOFLUSH))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TCIOFLUSH", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.TCOFLUSH))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TCOFLUSH", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.TCPInfo
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "TCPInfo", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.TCP_CONGESTION))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TCP_CONGESTION", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.TCP_CORK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TCP_CORK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.TCP_DEFER_ACCEPT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TCP_DEFER_ACCEPT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.TCP_INFO))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TCP_INFO", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.TCP_KEEPCNT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TCP_KEEPCNT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.TCP_KEEPIDLE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TCP_KEEPIDLE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.TCP_KEEPINTVL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TCP_KEEPINTVL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.TCP_LINGER2))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TCP_LINGER2", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.TCP_MAXSEG))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TCP_MAXSEG", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.TCP_MAXWIN))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TCP_MAXWIN", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.TCP_MAX_WINSHIFT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TCP_MAX_WINSHIFT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.TCP_MD5SIG))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TCP_MD5SIG", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.TCP_MD5SIG_MAXKEYLEN))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TCP_MD5SIG_MAXKEYLEN", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.TCP_MSS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TCP_MSS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.TCP_NODELAY))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TCP_NODELAY", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.TCP_QUICKACK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TCP_QUICKACK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.TCP_SYNCNT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TCP_SYNCNT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.TCP_WINDOW_CLAMP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TCP_WINDOW_CLAMP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.TCSETS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TCSETS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.TIOCCBRK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCCBRK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.TIOCCONS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCCONS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.TIOCEXCL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCEXCL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint64(mod.TIOCGDEV))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCGDEV", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.TIOCGETD))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCGETD", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.TIOCGICOUNT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCGICOUNT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.TIOCGLCKTRMIOS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCGLCKTRMIOS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.TIOCGPGRP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCGPGRP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint64(mod.TIOCGPTN))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCGPTN", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.TIOCGRS485))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCGRS485", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.TIOCGSERIAL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCGSERIAL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.TIOCGSID))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCGSID", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.TIOCGSOFTCAR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCGSOFTCAR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.TIOCGWINSZ))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCGWINSZ", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.TIOCINQ))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCINQ", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.TIOCLINUX))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCLINUX", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.TIOCMBIC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCMBIC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.TIOCMBIS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCMBIS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.TIOCMGET))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCMGET", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.TIOCMIWAIT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCMIWAIT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.TIOCMSET))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCMSET", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.TIOCM_CAR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCM_CAR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.TIOCM_CD))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCM_CD", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.TIOCM_CTS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCM_CTS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.TIOCM_DSR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCM_DSR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.TIOCM_DTR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCM_DTR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.TIOCM_LE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCM_LE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.TIOCM_RI))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCM_RI", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.TIOCM_RNG))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCM_RNG", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.TIOCM_RTS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCM_RTS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.TIOCM_SR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCM_SR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.TIOCM_ST))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCM_ST", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.TIOCNOTTY))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCNOTTY", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.TIOCNXCL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCNXCL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.TIOCOUTQ))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCOUTQ", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.TIOCPKT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCPKT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.TIOCPKT_DATA))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCPKT_DATA", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.TIOCPKT_DOSTOP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCPKT_DOSTOP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.TIOCPKT_FLUSHREAD))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCPKT_FLUSHREAD", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.TIOCPKT_FLUSHWRITE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCPKT_FLUSHWRITE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.TIOCPKT_IOCTL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCPKT_IOCTL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.TIOCPKT_NOSTOP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCPKT_NOSTOP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.TIOCPKT_START))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCPKT_START", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.TIOCPKT_STOP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCPKT_STOP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.TIOCSBRK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCSBRK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.TIOCSCTTY))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCSCTTY", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.TIOCSERCONFIG))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCSERCONFIG", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.TIOCSERGETLSR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCSERGETLSR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.TIOCSERGETMULTI))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCSERGETMULTI", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.TIOCSERGSTRUCT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCSERGSTRUCT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.TIOCSERGWILD))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCSERGWILD", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.TIOCSERSETMULTI))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCSERSETMULTI", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.TIOCSERSWILD))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCSERSWILD", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.TIOCSER_TEMT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCSER_TEMT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.TIOCSETD))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCSETD", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.TIOCSIG))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCSIG", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.TIOCSLCKTRMIOS))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCSLCKTRMIOS", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.TIOCSPGRP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCSPGRP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.TIOCSPTLCK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCSPTLCK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.TIOCSRS485))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCSRS485", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.TIOCSSERIAL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCSSERIAL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.TIOCSSOFTCAR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCSSOFTCAR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.TIOCSTI))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCSTI", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.TIOCSWINSZ))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCSWINSZ", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.TOSTOP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TOSTOP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.TUNATTACHFILTER))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TUNATTACHFILTER", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.TUNDETACHFILTER))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TUNDETACHFILTER", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint64(mod.TUNGETFEATURES))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TUNGETFEATURES", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint64(mod.TUNGETIFF))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TUNGETIFF", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint64(mod.TUNGETSNDBUF))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TUNGETSNDBUF", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint64(mod.TUNGETVNETHDRSZ))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TUNGETVNETHDRSZ", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.TUNSETDEBUG))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TUNSETDEBUG", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.TUNSETGROUP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TUNSETGROUP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.TUNSETIFF))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TUNSETIFF", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.TUNSETLINK))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TUNSETLINK", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.TUNSETNOCSUM))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TUNSETNOCSUM", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.TUNSETOFFLOAD))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TUNSETOFFLOAD", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.TUNSETOWNER))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TUNSETOWNER", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.TUNSETPERSIST))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TUNSETPERSIST", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.TUNSETSNDBUF))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TUNSETSNDBUF", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.TUNSETTXFILTER))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TUNSETTXFILTER", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.TUNSETVNETHDRSZ))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TUNSETVNETHDRSZ", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Tee)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Tee", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.Termios
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Termios", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Tgkill)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Tgkill", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Time)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Time", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.Time_t
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Time_t", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Times)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Times", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.Timespec
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Timespec", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.TimespecToNsec)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TimespecToNsec", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.Timeval
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Timeval", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.TimevalToNsec)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TimevalToNsec", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.Timex
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Timex", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.Tms
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Tms", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Truncate)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Truncate", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.Ucred
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Ucred", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Umask)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Umask", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Uname)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Uname", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.UnixCredentials)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "UnixCredentials", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.UnixRights)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "UnixRights", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Unlink)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Unlink", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Unlinkat)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Unlinkat", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Unmount)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Unmount", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Unsetenv)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Unsetenv", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Unshare)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Unshare", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Ustat)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Ustat", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.Ustat_t
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Ustat_t", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.Utimbuf
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Utimbuf", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Utime)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Utime", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Utimes)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Utimes", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.UtimesNano)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "UtimesNano", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.Utsname
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Utsname", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.VDISCARD))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "VDISCARD", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.VEOF))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "VEOF", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.VEOL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "VEOL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.VEOL2))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "VEOL2", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.VERASE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "VERASE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.VINTR))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "VINTR", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.VKILL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "VKILL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.VLNEXT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "VLNEXT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.VMIN))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "VMIN", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.VQUIT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "VQUIT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.VREPRINT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "VREPRINT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.VSTART))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "VSTART", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.VSTOP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "VSTOP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.VSUSP))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "VSUSP", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.VSWTC))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "VSWTC", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.VTIME))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "VTIME", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.VWERASE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "VWERASE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.WALL))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "WALL", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint64(mod.WCLONE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "WCLONE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.WCONTINUED))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "WCONTINUED", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.WEXITED))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "WEXITED", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.WNOHANG))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "WNOHANG", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.WNOTHREAD))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "WNOTHREAD", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint32(mod.WNOWAIT))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "WNOWAIT", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.WORDSIZE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "WORDSIZE", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.WSTOPPED))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "WSTOPPED", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.WUNTRACED))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "WUNTRACED", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Wait4)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Wait4", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.WaitStatus
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "WaitStatus", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Write)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Write", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.XCASE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "XCASE", o); raised != nil {
		return nil, raised
	}

	return nil, nil
}

var Code = pygolin_runtime.NewCode("<module>", "syscall", nil, 0, fun)

func init() {
	pygolin_runtime.RegisterModule("__go__/syscall", Code)
}
