package _sha256
import (
	πg "github.com/pygolin/runtime"
	// _ "github.com/pygolin/stdlib/_struct"
)
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/_sha256.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ß := πg.InternStr("")
		ß03d9963e05a094593190b6fc794cb1a3e1ac7d7883f0b5855268afeccc70d461 := πg.InternStr("03d9963e05a094593190b6fc794cb1a3e1ac7d7883f0b5855268afeccc70d461")
		ß8113ebf33c97daa9998762aacafe750c7cefc2b2f173c90c59663a57fe626f21 := πg.InternStr("8113ebf33c97daa9998762aacafe750c7cefc2b2f173c90c59663a57fe626f21")
		ßB := πg.InternStr("B")
		ßCh := πg.InternStr("Ch")
		ßGamma0 := πg.InternStr("Gamma0")
		ßGamma1 := πg.InternStr("Gamma1")
		ßMaj := πg.InternStr("Maj")
		ßNone := πg.InternStr("None")
		ßR := πg.InternStr("R")
		ßROR := πg.InternStr("ROR")
		ßS := πg.InternStr("S")
		ßSHA_BLOCKSIZE := πg.InternStr("SHA_BLOCKSIZE")
		ßSHA_DIGESTSIZE := πg.InternStr("SHA_DIGESTSIZE")
		ßSigma0 := πg.InternStr("Sigma0")
		ßSigma1 := πg.InternStr("Sigma1")
		ß__init__ := πg.InternStr("__init__")
		ß__main__ := πg.InternStr("__main__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß__new__ := πg.InternStr("__new__")
		ß_sha := πg.InternStr("_sha")
		ßappend := πg.InternStr("append")
		ßblock_size := πg.InternStr("block_size")
		ßbuffer := πg.InternStr("buffer")
		ßchr := πg.InternStr("chr")
		ßcopy := πg.InternStr("copy")
		ßcount_hi := πg.InternStr("count_hi")
		ßcount_lo := πg.InternStr("count_lo")
		ßd7b553c6f09ac85d142415f857c5310f3bbbe7cdd787cce4b985acedd585266f := πg.InternStr("d7b553c6f09ac85d142415f857c5310f3bbbe7cdd787cce4b985acedd585266f")
		ßdata := πg.InternStr("data")
		ßdigest := πg.InternStr("digest")
		ßdigest_size := πg.InternStr("digest_size")
		ßdigestsize := πg.InternStr("digestsize")
		ße3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855 := πg.InternStr("e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855")
		ßenumerate := πg.InternStr("enumerate")
		ßextend := πg.InternStr("extend")
		ßgetbuf := πg.InternStr("getbuf")
		ßhexdigest := πg.InternStr("hexdigest")
		ßisinstance := πg.InternStr("isinstance")
		ßjoin := πg.InternStr("join")
		ßlen := πg.InternStr("len")
		ßlocal := πg.InternStr("local")
		ßnew_shaobject := πg.InternStr("new_shaobject")
		ßobject := πg.InternStr("object")
		ßord := πg.InternStr("ord")
		ßsha224 := πg.InternStr("sha224")
		ßsha224_init := πg.InternStr("sha224_init")
		ßsha256 := πg.InternStr("sha256")
		ßsha_final := πg.InternStr("sha_final")
		ßsha_init := πg.InternStr("sha_init")
		ßsha_transform := πg.InternStr("sha_transform")
		ßsha_update := πg.InternStr("sha_update")
		ßstr := πg.InternStr("str")
		ßstruct := πg.InternStr("struct")
		ßtest := πg.InternStr("test")
		ßunicode := πg.InternStr("unicode")
		ßunpack := πg.InternStr("unpack")
		ßupdate := πg.InternStr("update")
		ßxrange := πg.InternStr("xrange")
		var πTemp001 *πg.Object
		_ = πTemp001
		var πTemp002 []*πg.Object
		_ = πTemp002
		var πTemp003 []πg.Param
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
		var πTemp010 *πg.Dict
		_ = πTemp010
		var πTemp011 *πg.Object
		_ = πTemp011
		var πTemp012 *πg.Object
		_ = πTemp012
		var πTemp013 *πg.Object
		_ = πTemp013
		var πTemp014 bool
		_ = πTemp014
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default: panic("unexpected function state")
			}
			// line 1: import _struct as struct
			πF.SetLineno(1)
			if πTemp002, πE = πg.ImportModule(πF, "_struct"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßstruct.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 3: SHA_BLOCKSIZE = 64
			πF.SetLineno(3)
			if πE = πF.Globals().SetItem(πF, ßSHA_BLOCKSIZE.ToObject(), πg.NewInt(64).ToObject()); πE != nil {
				continue
			}
			// line 4: SHA_DIGESTSIZE = 32
			πF.SetLineno(4)
			if πE = πF.Globals().SetItem(πF, ßSHA_DIGESTSIZE.ToObject(), πg.NewInt(32).ToObject()); πE != nil {
				continue
			}
			// line 7: def new_shaobject():
			πF.SetLineno(7)
			πTemp003 = make([]πg.Param, 0)
			πTemp001 = πg.NewFunction(πg.NewCode("new_shaobject", "/usr/lib/python2.7/_sha256.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 8: return {
					πF.SetLineno(8)
					πTemp001 = πg.NewDict()
					πTemp003 = make([]*πg.Object, 1)
					πTemp003[0] = πg.NewInt(0).ToObject()
					πTemp004 = πg.NewList(πTemp003...).ToObject()
					if πTemp002, πE = πg.Mul(πF, πTemp004, πg.NewInt(8).ToObject()); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßdigest.ToObject(), πTemp002); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßcount_lo.ToObject(), πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßcount_hi.ToObject(), πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					πTemp003 = make([]*πg.Object, 1)
					πTemp003[0] = πg.NewInt(0).ToObject()
					πTemp004 = πg.NewList(πTemp003...).ToObject()
					if πTemp005, πE = πg.ResolveGlobal(πF, ßSHA_BLOCKSIZE); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mul(πF, πTemp004, πTemp005); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßdata.ToObject(), πTemp002); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßlocal.ToObject(), πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßdigestsize.ToObject(), πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					πTemp002 = πTemp001.ToObject()
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
			if πE = πF.Globals().SetItem(πF, ßnew_shaobject.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 17: ROR = lambda x, y: (((x & 0xffffffff) >> (y & 31)) | (x << (32 - (y & 31)))) & 0xffffffff
			πF.SetLineno(17)
			πTemp003 = make([]πg.Param, 2)
			πTemp003[0] = πg.Param{Name: "x", Def: nil}
			πTemp003[1] = πg.Param{Name: "y", Def: nil}
			πTemp004 = πg.NewFunction(πg.NewCode("<lambda>", "/usr/lib/python2.7/_sha256.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πArgs[0]; _ = µx
				var µy *πg.Object = πArgs[1]; _ = µy
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
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 17: ROR = lambda x, y: (((x & 0xffffffff) >> (y & 31)) | (x << (32 - (y & 31)))) & 0xffffffff
					πF.SetLineno(17)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.And(πF, µx, πg.NewInt(4294967295).ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.And(πF, µy, πg.NewInt(31).ToObject()); πE != nil {
						continue
					}
					if πTemp003, πE = πg.RShift(πF, πTemp004, πTemp005); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.And(πF, µy, πg.NewInt(31).ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Sub(πF, πg.NewInt(32).ToObject(), πTemp006); πE != nil {
						continue
					}
					if πTemp004, πE = πg.LShift(πF, µx, πTemp005); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Or(πF, πTemp003, πTemp004); πE != nil {
						continue
					}
					if πTemp001, πE = πg.And(πF, πTemp002, πg.NewInt(4294967295).ToObject()); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßROR.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 18: Ch = lambda x, y, z: (z ^ (x & (y ^ z)))
			πF.SetLineno(18)
			πTemp003 = make([]πg.Param, 3)
			πTemp003[0] = πg.Param{Name: "x", Def: nil}
			πTemp003[1] = πg.Param{Name: "y", Def: nil}
			πTemp003[2] = πg.Param{Name: "z", Def: nil}
			πTemp004 = πg.NewFunction(πg.NewCode("<lambda>", "/usr/lib/python2.7/_sha256.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πArgs[0]; _ = µx
				var µy *πg.Object = πArgs[1]; _ = µy
				var µz *πg.Object = πArgs[2]; _ = µz
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
					// line 18: Ch = lambda x, y, z: (z ^ (x & (y ^ z)))
					πF.SetLineno(18)
					if πE = πg.CheckLocal(πF, µz, "z"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µz, "z"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Xor(πF, µy, µz); πE != nil {
						continue
					}
					if πTemp002, πE = πg.And(πF, µx, πTemp003); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Xor(πF, µz, πTemp002); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßCh.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 19: Maj = lambda x, y, z: (((x | y) & z) | (x & y))
			πF.SetLineno(19)
			πTemp003 = make([]πg.Param, 3)
			πTemp003[0] = πg.Param{Name: "x", Def: nil}
			πTemp003[1] = πg.Param{Name: "y", Def: nil}
			πTemp003[2] = πg.Param{Name: "z", Def: nil}
			πTemp004 = πg.NewFunction(πg.NewCode("<lambda>", "/usr/lib/python2.7/_sha256.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πArgs[0]; _ = µx
				var µy *πg.Object = πArgs[1]; _ = µy
				var µz *πg.Object = πArgs[2]; _ = µz
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
					// line 19: Maj = lambda x, y, z: (((x | y) & z) | (x & y))
					πF.SetLineno(19)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Or(πF, µx, µy); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µz, "z"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.And(πF, πTemp003, µz); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.And(πF, µx, µy); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Or(πF, πTemp002, πTemp003); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßMaj.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 20: S = lambda x, n: ROR(x, n)
			πF.SetLineno(20)
			πTemp003 = make([]πg.Param, 2)
			πTemp003[0] = πg.Param{Name: "x", Def: nil}
			πTemp003[1] = πg.Param{Name: "n", Def: nil}
			πTemp004 = πg.NewFunction(πg.NewCode("<lambda>", "/usr/lib/python2.7/_sha256.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πArgs[0]; _ = µx
				var µn *πg.Object = πArgs[1]; _ = µn
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
					// line 20: S = lambda x, n: ROR(x, n)
					πF.SetLineno(20)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp001[0] = µx
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					πTemp001[1] = µn
					if πTemp002, πE = πg.ResolveGlobal(πF, ßROR); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßS.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 21: R = lambda x, n: (x & 0xffffffff) >> n
			πF.SetLineno(21)
			πTemp003 = make([]πg.Param, 2)
			πTemp003[0] = πg.Param{Name: "x", Def: nil}
			πTemp003[1] = πg.Param{Name: "n", Def: nil}
			πTemp004 = πg.NewFunction(πg.NewCode("<lambda>", "/usr/lib/python2.7/_sha256.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πArgs[0]; _ = µx
				var µn *πg.Object = πArgs[1]; _ = µn
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
					// line 21: R = lambda x, n: (x & 0xffffffff) >> n
					πF.SetLineno(21)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.And(πF, µx, πg.NewInt(4294967295).ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.RShift(πF, πTemp002, µn); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßR.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 22: Sigma0 = lambda x: (S(x, 2) ^ S(x, 13) ^ S(x, 22))
			πF.SetLineno(22)
			πTemp003 = make([]πg.Param, 1)
			πTemp003[0] = πg.Param{Name: "x", Def: nil}
			πTemp004 = πg.NewFunction(πg.NewCode("<lambda>", "/usr/lib/python2.7/_sha256.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πArgs[0]; _ = µx
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
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 22: Sigma0 = lambda x: (S(x, 2) ^ S(x, 13) ^ S(x, 22))
					πF.SetLineno(22)
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp003[0] = µx
					πTemp003[1] = πg.NewInt(2).ToObject()
					if πTemp004, πE = πg.ResolveGlobal(πF, ßS); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp003[0] = µx
					πTemp003[1] = πg.NewInt(13).ToObject()
					if πTemp004, πE = πg.ResolveGlobal(πF, ßS); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp002, πE = πg.Xor(πF, πTemp005, πTemp006); πE != nil {
						continue
					}
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp003[0] = µx
					πTemp003[1] = πg.NewInt(22).ToObject()
					if πTemp004, πE = πg.ResolveGlobal(πF, ßS); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp001, πE = πg.Xor(πF, πTemp002, πTemp005); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßSigma0.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 23: Sigma1 = lambda x: (S(x, 6) ^ S(x, 11) ^ S(x, 25))
			πF.SetLineno(23)
			πTemp003 = make([]πg.Param, 1)
			πTemp003[0] = πg.Param{Name: "x", Def: nil}
			πTemp004 = πg.NewFunction(πg.NewCode("<lambda>", "/usr/lib/python2.7/_sha256.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πArgs[0]; _ = µx
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
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 23: Sigma1 = lambda x: (S(x, 6) ^ S(x, 11) ^ S(x, 25))
					πF.SetLineno(23)
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp003[0] = µx
					πTemp003[1] = πg.NewInt(6).ToObject()
					if πTemp004, πE = πg.ResolveGlobal(πF, ßS); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp003[0] = µx
					πTemp003[1] = πg.NewInt(11).ToObject()
					if πTemp004, πE = πg.ResolveGlobal(πF, ßS); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp002, πE = πg.Xor(πF, πTemp005, πTemp006); πE != nil {
						continue
					}
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp003[0] = µx
					πTemp003[1] = πg.NewInt(25).ToObject()
					if πTemp004, πE = πg.ResolveGlobal(πF, ßS); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp001, πE = πg.Xor(πF, πTemp002, πTemp005); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßSigma1.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 24: Gamma0 = lambda x: (S(x, 7) ^ S(x, 18) ^ R(x, 3))
			πF.SetLineno(24)
			πTemp003 = make([]πg.Param, 1)
			πTemp003[0] = πg.Param{Name: "x", Def: nil}
			πTemp004 = πg.NewFunction(πg.NewCode("<lambda>", "/usr/lib/python2.7/_sha256.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πArgs[0]; _ = µx
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
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 24: Gamma0 = lambda x: (S(x, 7) ^ S(x, 18) ^ R(x, 3))
					πF.SetLineno(24)
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp003[0] = µx
					πTemp003[1] = πg.NewInt(7).ToObject()
					if πTemp004, πE = πg.ResolveGlobal(πF, ßS); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp003[0] = µx
					πTemp003[1] = πg.NewInt(18).ToObject()
					if πTemp004, πE = πg.ResolveGlobal(πF, ßS); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp002, πE = πg.Xor(πF, πTemp005, πTemp006); πE != nil {
						continue
					}
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp003[0] = µx
					πTemp003[1] = πg.NewInt(3).ToObject()
					if πTemp004, πE = πg.ResolveGlobal(πF, ßR); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp001, πE = πg.Xor(πF, πTemp002, πTemp005); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßGamma0.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 25: Gamma1 = lambda x: (S(x, 17) ^ S(x, 19) ^ R(x, 10))
			πF.SetLineno(25)
			πTemp003 = make([]πg.Param, 1)
			πTemp003[0] = πg.Param{Name: "x", Def: nil}
			πTemp004 = πg.NewFunction(πg.NewCode("<lambda>", "/usr/lib/python2.7/_sha256.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πArgs[0]; _ = µx
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
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 25: Gamma1 = lambda x: (S(x, 17) ^ S(x, 19) ^ R(x, 10))
					πF.SetLineno(25)
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp003[0] = µx
					πTemp003[1] = πg.NewInt(17).ToObject()
					if πTemp004, πE = πg.ResolveGlobal(πF, ßS); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp003[0] = µx
					πTemp003[1] = πg.NewInt(19).ToObject()
					if πTemp004, πE = πg.ResolveGlobal(πF, ßS); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp002, πE = πg.Xor(πF, πTemp005, πTemp006); πE != nil {
						continue
					}
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp003[0] = µx
					πTemp003[1] = πg.NewInt(10).ToObject()
					if πTemp004, πE = πg.ResolveGlobal(πF, ßR); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp001, πE = πg.Xor(πF, πTemp002, πTemp005); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßGamma1.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 27: def sha_transform(sha_info):
			πF.SetLineno(27)
			πTemp003 = make([]πg.Param, 1)
			πTemp003[0] = πg.Param{Name: "sha_info", Def: nil}
			πTemp004 = πg.NewFunction(πg.NewCode("sha_transform", "/usr/lib/python2.7/_sha256.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µsha_info *πg.Object = πArgs[0]; _ = µsha_info
				var µW *πg.Object = πg.UnboundLocal; _ = µW
				var µd *πg.Object = πg.UnboundLocal; _ = µd
				var µi *πg.Object = πg.UnboundLocal; _ = µi
				var µss *πg.Object = πg.UnboundLocal; _ = µss
				var µRND *πg.Object = πg.UnboundLocal; _ = µRND
				var µdig *πg.Object = πg.UnboundLocal; _ = µdig
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
				var πTemp013 []*πg.Object
				_ = πTemp013
				var πTemp014 []πg.Param
				_ = πTemp014
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
					default: panic("unexpected function state")
					}
					// line 28: W = []
					πF.SetLineno(28)
					πTemp001 = make([]*πg.Object, 0)
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					µW = πTemp002
					// line 30: d = sha_info['data']
					πF.SetLineno(30)
					πTemp002 = ßdata.ToObject()
					if πE = πg.CheckLocal(πF, µsha_info, "sha_info"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µsha_info, πTemp002); πE != nil {
						continue
					}
					µd = πTemp003
					πTemp001 = πF.MakeArgs(2)
					πTemp001[0] = πg.NewInt(0).ToObject()
					πTemp001[1] = πg.NewInt(16).ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßxrange); πE != nil {
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
						µi = πTemp003
					}
					if πE != nil || !πTemp006 {
						continue
					}
					πF.PushCheckpoint(1)            
					// line 32: W.append( (d[4*i]<<24) + (d[4*i+1]<<16) + (d[4*i+2]<<8) + d[4*i+3])
					πF.SetLineno(32)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp010, πE = πg.Mul(πF, πg.NewInt(4).ToObject(), µi); πE != nil {
						continue
					}
					πTemp009 = πTemp010
					if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
						continue
					}
					if πTemp010, πE = πg.GetItem(πF, µd, πTemp009); πE != nil {
						continue
					}
					if πTemp008, πE = πg.LShift(πF, πTemp010, πg.NewInt(24).ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp012, πE = πg.Mul(πF, πg.NewInt(4).ToObject(), µi); πE != nil {
						continue
					}
					if πTemp011, πE = πg.Add(πF, πTemp012, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp010 = πTemp011
					if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
						continue
					}
					if πTemp011, πE = πg.GetItem(πF, µd, πTemp010); πE != nil {
						continue
					}
					if πTemp009, πE = πg.LShift(πF, πTemp011, πg.NewInt(16).ToObject()); πE != nil {
						continue
					}
					if πTemp007, πE = πg.Add(πF, πTemp008, πTemp009); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp011, πE = πg.Mul(πF, πg.NewInt(4).ToObject(), µi); πE != nil {
						continue
					}
					if πTemp010, πE = πg.Add(πF, πTemp011, πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
					πTemp009 = πTemp010
					if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
						continue
					}
					if πTemp010, πE = πg.GetItem(πF, µd, πTemp009); πE != nil {
						continue
					}
					if πTemp008, πE = πg.LShift(πF, πTemp010, πg.NewInt(8).ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Add(πF, πTemp007, πTemp008); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp009, πE = πg.Mul(πF, πg.NewInt(4).ToObject(), µi); πE != nil {
						continue
					}
					if πTemp008, πE = πg.Add(πF, πTemp009, πg.NewInt(3).ToObject()); πE != nil {
						continue
					}
					πTemp007 = πTemp008
					if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetItem(πF, µd, πTemp007); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Add(πF, πTemp004, πTemp008); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					if πE = πg.CheckLocal(πF, µW, "W"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µW, ßappend, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
					πTemp001 = πF.MakeArgs(2)
					πTemp001[0] = πg.NewInt(16).ToObject()
					πTemp001[1] = πg.NewInt(64).ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßxrange); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp002, πE = πg.Iter(πF, πTemp004); πE != nil {
						continue
					}
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
						µi = πTemp003
					}
					if πE != nil || !πTemp006 {
						continue
					}
					πF.PushCheckpoint(4)            
					// line 35: W.append( (Gamma1(W[i - 2]) + W[i - 7] + Gamma0(W[i - 15]) + W[i - 16]) & 0xffffffff )
					πF.SetLineno(35)
					πTemp001 = πF.MakeArgs(1)
					πTemp013 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp010, πE = πg.Sub(πF, µi, πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
					πTemp009 = πTemp010
					if πE = πg.CheckLocal(πF, µW, "W"); πE != nil {
						continue
					}
					if πTemp010, πE = πg.GetItem(πF, µW, πTemp009); πE != nil {
						continue
					}
					πTemp013[0] = πTemp010
					if πTemp009, πE = πg.ResolveGlobal(πF, ßGamma1); πE != nil {
						continue
					}
					if πTemp010, πE = πTemp009.Call(πF, πTemp013, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp013)
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp011, πE = πg.Sub(πF, µi, πg.NewInt(7).ToObject()); πE != nil {
						continue
					}
					πTemp009 = πTemp011
					if πE = πg.CheckLocal(πF, µW, "W"); πE != nil {
						continue
					}
					if πTemp011, πE = πg.GetItem(πF, µW, πTemp009); πE != nil {
						continue
					}
					if πTemp008, πE = πg.Add(πF, πTemp010, πTemp011); πE != nil {
						continue
					}
					πTemp013 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp010, πE = πg.Sub(πF, µi, πg.NewInt(15).ToObject()); πE != nil {
						continue
					}
					πTemp009 = πTemp010
					if πE = πg.CheckLocal(πF, µW, "W"); πE != nil {
						continue
					}
					if πTemp010, πE = πg.GetItem(πF, µW, πTemp009); πE != nil {
						continue
					}
					πTemp013[0] = πTemp010
					if πTemp009, πE = πg.ResolveGlobal(πF, ßGamma0); πE != nil {
						continue
					}
					if πTemp010, πE = πTemp009.Call(πF, πTemp013, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp013)
					if πTemp007, πE = πg.Add(πF, πTemp008, πTemp010); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp009, πE = πg.Sub(πF, µi, πg.NewInt(16).ToObject()); πE != nil {
						continue
					}
					πTemp008 = πTemp009
					if πE = πg.CheckLocal(πF, µW, "W"); πE != nil {
						continue
					}
					if πTemp009, πE = πg.GetItem(πF, µW, πTemp008); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Add(πF, πTemp007, πTemp009); πE != nil {
						continue
					}
					if πTemp003, πE = πg.And(πF, πTemp004, πg.NewInt(4294967295).ToObject()); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					if πE = πg.CheckLocal(πF, µW, "W"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µW, ßappend, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					continue
				Label5:
					if πE != nil || πR != nil {
						continue
					}
				Label6:
					// line 37: ss = sha_info['digest'][:]
					πF.SetLineno(37)
					if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.None, πg.None}, nil); πE != nil {
						continue
					}
					πTemp004 = ßdigest.ToObject()
					if πE = πg.CheckLocal(πF, µsha_info, "sha_info"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetItem(πF, µsha_info, πTemp004); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, πTemp007, πTemp002); πE != nil {
						continue
					}
					µss = πTemp003
					// line 39: def RND(a,b,c,d,e,f,g,h,i,ki):
					πF.SetLineno(39)
					πTemp014 = make([]πg.Param, 10)
					πTemp014[0] = πg.Param{Name: "a", Def: nil}
					πTemp014[1] = πg.Param{Name: "b", Def: nil}
					πTemp014[2] = πg.Param{Name: "c", Def: nil}
					πTemp014[3] = πg.Param{Name: "d", Def: nil}
					πTemp014[4] = πg.Param{Name: "e", Def: nil}
					πTemp014[5] = πg.Param{Name: "f", Def: nil}
					πTemp014[6] = πg.Param{Name: "g", Def: nil}
					πTemp014[7] = πg.Param{Name: "h", Def: nil}
					πTemp014[8] = πg.Param{Name: "i", Def: nil}
					πTemp014[9] = πg.Param{Name: "ki", Def: nil}
					πTemp002 = πg.NewFunction(πg.NewCode("RND", "/usr/lib/python2.7/_sha256.py", πTemp014, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µa *πg.Object = πArgs[0]; _ = µa
						var µb *πg.Object = πArgs[1]; _ = µb
						var µc *πg.Object = πArgs[2]; _ = µc
						var µd *πg.Object = πArgs[3]; _ = µd
						var µe *πg.Object = πArgs[4]; _ = µe
						var µf *πg.Object = πArgs[5]; _ = µf
						var µg *πg.Object = πArgs[6]; _ = µg
						var µh *πg.Object = πArgs[7]; _ = µh
						var µi *πg.Object = πArgs[8]; _ = µi
						var µki *πg.Object = πArgs[9]; _ = µki
						var µt0 *πg.Object = πg.UnboundLocal; _ = µt0
						var µt1 *πg.Object = πg.UnboundLocal; _ = µt1
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
							// line 40: t0 = h + Sigma1(e) + Ch(e, f, g) + ki + W[i];
							πF.SetLineno(40)
							if πE = πg.CheckLocal(πF, µh, "h"); πE != nil {
								continue
							}
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µe, "e"); πE != nil {
								continue
							}
							πTemp005[0] = µe
							if πTemp006, πE = πg.ResolveGlobal(πF, ßSigma1); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp004, πE = πg.Add(πF, µh, πTemp007); πE != nil {
								continue
							}
							πTemp005 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µe, "e"); πE != nil {
								continue
							}
							πTemp005[0] = µe
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							πTemp005[1] = µf
							if πE = πg.CheckLocal(πF, µg, "g"); πE != nil {
								continue
							}
							πTemp005[2] = µg
							if πTemp006, πE = πg.ResolveGlobal(πF, ßCh); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp003, πE = πg.Add(πF, πTemp004, πTemp007); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µki, "ki"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp003, µki); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp003 = µi
							if πE = πg.CheckLocal(πF, µW, "W"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µW, πTemp003); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp002, πTemp004); πE != nil {
								continue
							}
							µt0 = πTemp001
							// line 41: t1 = Sigma0(a) + Maj(a, b, c);
							πF.SetLineno(41)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp005[0] = µa
							if πTemp002, πE = πg.ResolveGlobal(πF, ßSigma0); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp005 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp005[0] = µa
							if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
								continue
							}
							πTemp005[1] = µb
							if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
								continue
							}
							πTemp005[2] = µc
							if πTemp002, πE = πg.ResolveGlobal(πF, ßMaj); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp001, πE = πg.Add(πF, πTemp003, πTemp004); πE != nil {
								continue
							}
							µt1 = πTemp001
							// line 42: d += t0;
							πF.SetLineno(42)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µt0, "t0"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.IAdd(πF, µd, µt0); πE != nil {
								continue
							}
							µd = πTemp001
							// line 43: h  = t0 + t1;
							πF.SetLineno(43)
							if πE = πg.CheckLocal(πF, µt0, "t0"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µt1, "t1"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, µt0, µt1); πE != nil {
								continue
							}
							µh = πTemp001
							// line 44: return d & 0xffffffff, h & 0xffffffff
							πF.SetLineno(44)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.And(πF, µd, πg.NewInt(4294967295).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µh, "h"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.And(πF, µh, πg.NewInt(4294967295).ToObject()); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple2(πTemp002, πTemp003).ToObject()
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
					µRND = πTemp002
					// line 46: ss[3], ss[7] = RND(ss[0],ss[1],ss[2],ss[3],ss[4],ss[5],ss[6],ss[7],0,0x428a2f98);
					πF.SetLineno(46)
					πTemp001 = πF.MakeArgs(10)
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp004
					πTemp003 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[1] = πTemp004
					πTemp003 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[2] = πTemp004
					πTemp003 = πg.NewInt(3).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[3] = πTemp004
					πTemp003 = πg.NewInt(4).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[4] = πTemp004
					πTemp003 = πg.NewInt(5).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[5] = πTemp004
					πTemp003 = πg.NewInt(6).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[6] = πTemp004
					πTemp003 = πg.NewInt(7).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[7] = πTemp004
					πTemp001[8] = πg.NewInt(0).ToObject()
					πTemp001[9] = πg.NewInt(1116352408).ToObject()
					if πE = πg.CheckLocal(πF, µRND, "RND"); πE != nil {
						continue
					}
					if πTemp003, πE = µRND.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp008 = πg.NewInt(3).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp008, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp004 = πg.NewInt(7).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp004, πTemp007); πE != nil {
						continue
					}
					// line 47: ss[2], ss[6] = RND(ss[7],ss[0],ss[1],ss[2],ss[3],ss[4],ss[5],ss[6],1,0x71374491);
					πF.SetLineno(47)
					πTemp001 = πF.MakeArgs(10)
					πTemp003 = πg.NewInt(7).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp004
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[1] = πTemp004
					πTemp003 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[2] = πTemp004
					πTemp003 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[3] = πTemp004
					πTemp003 = πg.NewInt(3).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[4] = πTemp004
					πTemp003 = πg.NewInt(4).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[5] = πTemp004
					πTemp003 = πg.NewInt(5).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[6] = πTemp004
					πTemp003 = πg.NewInt(6).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[7] = πTemp004
					πTemp001[8] = πg.NewInt(1).ToObject()
					πTemp001[9] = πg.NewInt(1899447441).ToObject()
					if πE = πg.CheckLocal(πF, µRND, "RND"); πE != nil {
						continue
					}
					if πTemp003, πE = µRND.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp008 = πg.NewInt(2).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp008, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp004 = πg.NewInt(6).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp004, πTemp007); πE != nil {
						continue
					}
					// line 48: ss[1], ss[5] = RND(ss[6],ss[7],ss[0],ss[1],ss[2],ss[3],ss[4],ss[5],2,0xb5c0fbcf);
					πF.SetLineno(48)
					πTemp001 = πF.MakeArgs(10)
					πTemp003 = πg.NewInt(6).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp004
					πTemp003 = πg.NewInt(7).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[1] = πTemp004
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[2] = πTemp004
					πTemp003 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[3] = πTemp004
					πTemp003 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[4] = πTemp004
					πTemp003 = πg.NewInt(3).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[5] = πTemp004
					πTemp003 = πg.NewInt(4).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[6] = πTemp004
					πTemp003 = πg.NewInt(5).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[7] = πTemp004
					πTemp001[8] = πg.NewInt(2).ToObject()
					πTemp001[9] = πg.NewInt(3049323471).ToObject()
					if πE = πg.CheckLocal(πF, µRND, "RND"); πE != nil {
						continue
					}
					if πTemp003, πE = µRND.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp008 = πg.NewInt(1).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp008, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp004 = πg.NewInt(5).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp004, πTemp007); πE != nil {
						continue
					}
					// line 49: ss[0], ss[4] = RND(ss[5],ss[6],ss[7],ss[0],ss[1],ss[2],ss[3],ss[4],3,0xe9b5dba5);
					πF.SetLineno(49)
					πTemp001 = πF.MakeArgs(10)
					πTemp003 = πg.NewInt(5).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp004
					πTemp003 = πg.NewInt(6).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[1] = πTemp004
					πTemp003 = πg.NewInt(7).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[2] = πTemp004
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[3] = πTemp004
					πTemp003 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[4] = πTemp004
					πTemp003 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[5] = πTemp004
					πTemp003 = πg.NewInt(3).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[6] = πTemp004
					πTemp003 = πg.NewInt(4).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[7] = πTemp004
					πTemp001[8] = πg.NewInt(3).ToObject()
					πTemp001[9] = πg.NewInt(3921009573).ToObject()
					if πE = πg.CheckLocal(πF, µRND, "RND"); πE != nil {
						continue
					}
					if πTemp003, πE = µRND.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp008 = πg.NewInt(0).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp008, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp004 = πg.NewInt(4).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp004, πTemp007); πE != nil {
						continue
					}
					// line 50: ss[7], ss[3] = RND(ss[4],ss[5],ss[6],ss[7],ss[0],ss[1],ss[2],ss[3],4,0x3956c25b);
					πF.SetLineno(50)
					πTemp001 = πF.MakeArgs(10)
					πTemp003 = πg.NewInt(4).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp004
					πTemp003 = πg.NewInt(5).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[1] = πTemp004
					πTemp003 = πg.NewInt(6).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[2] = πTemp004
					πTemp003 = πg.NewInt(7).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[3] = πTemp004
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[4] = πTemp004
					πTemp003 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[5] = πTemp004
					πTemp003 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[6] = πTemp004
					πTemp003 = πg.NewInt(3).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[7] = πTemp004
					πTemp001[8] = πg.NewInt(4).ToObject()
					πTemp001[9] = πg.NewInt(961987163).ToObject()
					if πE = πg.CheckLocal(πF, µRND, "RND"); πE != nil {
						continue
					}
					if πTemp003, πE = µRND.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp008 = πg.NewInt(7).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp008, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp004 = πg.NewInt(3).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp004, πTemp007); πE != nil {
						continue
					}
					// line 51: ss[6], ss[2] = RND(ss[3],ss[4],ss[5],ss[6],ss[7],ss[0],ss[1],ss[2],5,0x59f111f1);
					πF.SetLineno(51)
					πTemp001 = πF.MakeArgs(10)
					πTemp003 = πg.NewInt(3).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp004
					πTemp003 = πg.NewInt(4).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[1] = πTemp004
					πTemp003 = πg.NewInt(5).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[2] = πTemp004
					πTemp003 = πg.NewInt(6).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[3] = πTemp004
					πTemp003 = πg.NewInt(7).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[4] = πTemp004
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[5] = πTemp004
					πTemp003 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[6] = πTemp004
					πTemp003 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[7] = πTemp004
					πTemp001[8] = πg.NewInt(5).ToObject()
					πTemp001[9] = πg.NewInt(1508970993).ToObject()
					if πE = πg.CheckLocal(πF, µRND, "RND"); πE != nil {
						continue
					}
					if πTemp003, πE = µRND.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp008 = πg.NewInt(6).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp008, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp004 = πg.NewInt(2).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp004, πTemp007); πE != nil {
						continue
					}
					// line 52: ss[5], ss[1] = RND(ss[2],ss[3],ss[4],ss[5],ss[6],ss[7],ss[0],ss[1],6,0x923f82a4);
					πF.SetLineno(52)
					πTemp001 = πF.MakeArgs(10)
					πTemp003 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp004
					πTemp003 = πg.NewInt(3).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[1] = πTemp004
					πTemp003 = πg.NewInt(4).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[2] = πTemp004
					πTemp003 = πg.NewInt(5).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[3] = πTemp004
					πTemp003 = πg.NewInt(6).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[4] = πTemp004
					πTemp003 = πg.NewInt(7).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[5] = πTemp004
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[6] = πTemp004
					πTemp003 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[7] = πTemp004
					πTemp001[8] = πg.NewInt(6).ToObject()
					πTemp001[9] = πg.NewInt(2453635748).ToObject()
					if πE = πg.CheckLocal(πF, µRND, "RND"); πE != nil {
						continue
					}
					if πTemp003, πE = µRND.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp008 = πg.NewInt(5).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp008, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp004 = πg.NewInt(1).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp004, πTemp007); πE != nil {
						continue
					}
					// line 53: ss[4], ss[0] = RND(ss[1],ss[2],ss[3],ss[4],ss[5],ss[6],ss[7],ss[0],7,0xab1c5ed5);
					πF.SetLineno(53)
					πTemp001 = πF.MakeArgs(10)
					πTemp003 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp004
					πTemp003 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[1] = πTemp004
					πTemp003 = πg.NewInt(3).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[2] = πTemp004
					πTemp003 = πg.NewInt(4).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[3] = πTemp004
					πTemp003 = πg.NewInt(5).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[4] = πTemp004
					πTemp003 = πg.NewInt(6).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[5] = πTemp004
					πTemp003 = πg.NewInt(7).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[6] = πTemp004
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[7] = πTemp004
					πTemp001[8] = πg.NewInt(7).ToObject()
					πTemp001[9] = πg.NewInt(2870763221).ToObject()
					if πE = πg.CheckLocal(πF, µRND, "RND"); πE != nil {
						continue
					}
					if πTemp003, πE = µRND.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp008 = πg.NewInt(4).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp008, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp004 = πg.NewInt(0).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp004, πTemp007); πE != nil {
						continue
					}
					// line 54: ss[3], ss[7] = RND(ss[0],ss[1],ss[2],ss[3],ss[4],ss[5],ss[6],ss[7],8,0xd807aa98);
					πF.SetLineno(54)
					πTemp001 = πF.MakeArgs(10)
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp004
					πTemp003 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[1] = πTemp004
					πTemp003 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[2] = πTemp004
					πTemp003 = πg.NewInt(3).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[3] = πTemp004
					πTemp003 = πg.NewInt(4).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[4] = πTemp004
					πTemp003 = πg.NewInt(5).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[5] = πTemp004
					πTemp003 = πg.NewInt(6).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[6] = πTemp004
					πTemp003 = πg.NewInt(7).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[7] = πTemp004
					πTemp001[8] = πg.NewInt(8).ToObject()
					πTemp001[9] = πg.NewInt(3624381080).ToObject()
					if πE = πg.CheckLocal(πF, µRND, "RND"); πE != nil {
						continue
					}
					if πTemp003, πE = µRND.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp008 = πg.NewInt(3).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp008, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp004 = πg.NewInt(7).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp004, πTemp007); πE != nil {
						continue
					}
					// line 55: ss[2], ss[6] = RND(ss[7],ss[0],ss[1],ss[2],ss[3],ss[4],ss[5],ss[6],9,0x12835b01);
					πF.SetLineno(55)
					πTemp001 = πF.MakeArgs(10)
					πTemp003 = πg.NewInt(7).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp004
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[1] = πTemp004
					πTemp003 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[2] = πTemp004
					πTemp003 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[3] = πTemp004
					πTemp003 = πg.NewInt(3).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[4] = πTemp004
					πTemp003 = πg.NewInt(4).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[5] = πTemp004
					πTemp003 = πg.NewInt(5).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[6] = πTemp004
					πTemp003 = πg.NewInt(6).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[7] = πTemp004
					πTemp001[8] = πg.NewInt(9).ToObject()
					πTemp001[9] = πg.NewInt(310598401).ToObject()
					if πE = πg.CheckLocal(πF, µRND, "RND"); πE != nil {
						continue
					}
					if πTemp003, πE = µRND.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp008 = πg.NewInt(2).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp008, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp004 = πg.NewInt(6).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp004, πTemp007); πE != nil {
						continue
					}
					// line 56: ss[1], ss[5] = RND(ss[6],ss[7],ss[0],ss[1],ss[2],ss[3],ss[4],ss[5],10,0x243185be);
					πF.SetLineno(56)
					πTemp001 = πF.MakeArgs(10)
					πTemp003 = πg.NewInt(6).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp004
					πTemp003 = πg.NewInt(7).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[1] = πTemp004
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[2] = πTemp004
					πTemp003 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[3] = πTemp004
					πTemp003 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[4] = πTemp004
					πTemp003 = πg.NewInt(3).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[5] = πTemp004
					πTemp003 = πg.NewInt(4).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[6] = πTemp004
					πTemp003 = πg.NewInt(5).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[7] = πTemp004
					πTemp001[8] = πg.NewInt(10).ToObject()
					πTemp001[9] = πg.NewInt(607225278).ToObject()
					if πE = πg.CheckLocal(πF, µRND, "RND"); πE != nil {
						continue
					}
					if πTemp003, πE = µRND.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp008 = πg.NewInt(1).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp008, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp004 = πg.NewInt(5).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp004, πTemp007); πE != nil {
						continue
					}
					// line 57: ss[0], ss[4] = RND(ss[5],ss[6],ss[7],ss[0],ss[1],ss[2],ss[3],ss[4],11,0x550c7dc3);
					πF.SetLineno(57)
					πTemp001 = πF.MakeArgs(10)
					πTemp003 = πg.NewInt(5).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp004
					πTemp003 = πg.NewInt(6).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[1] = πTemp004
					πTemp003 = πg.NewInt(7).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[2] = πTemp004
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[3] = πTemp004
					πTemp003 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[4] = πTemp004
					πTemp003 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[5] = πTemp004
					πTemp003 = πg.NewInt(3).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[6] = πTemp004
					πTemp003 = πg.NewInt(4).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[7] = πTemp004
					πTemp001[8] = πg.NewInt(11).ToObject()
					πTemp001[9] = πg.NewInt(1426881987).ToObject()
					if πE = πg.CheckLocal(πF, µRND, "RND"); πE != nil {
						continue
					}
					if πTemp003, πE = µRND.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp008 = πg.NewInt(0).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp008, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp004 = πg.NewInt(4).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp004, πTemp007); πE != nil {
						continue
					}
					// line 58: ss[7], ss[3] = RND(ss[4],ss[5],ss[6],ss[7],ss[0],ss[1],ss[2],ss[3],12,0x72be5d74);
					πF.SetLineno(58)
					πTemp001 = πF.MakeArgs(10)
					πTemp003 = πg.NewInt(4).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp004
					πTemp003 = πg.NewInt(5).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[1] = πTemp004
					πTemp003 = πg.NewInt(6).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[2] = πTemp004
					πTemp003 = πg.NewInt(7).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[3] = πTemp004
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[4] = πTemp004
					πTemp003 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[5] = πTemp004
					πTemp003 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[6] = πTemp004
					πTemp003 = πg.NewInt(3).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[7] = πTemp004
					πTemp001[8] = πg.NewInt(12).ToObject()
					πTemp001[9] = πg.NewInt(1925078388).ToObject()
					if πE = πg.CheckLocal(πF, µRND, "RND"); πE != nil {
						continue
					}
					if πTemp003, πE = µRND.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp008 = πg.NewInt(7).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp008, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp004 = πg.NewInt(3).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp004, πTemp007); πE != nil {
						continue
					}
					// line 59: ss[6], ss[2] = RND(ss[3],ss[4],ss[5],ss[6],ss[7],ss[0],ss[1],ss[2],13,0x80deb1fe);
					πF.SetLineno(59)
					πTemp001 = πF.MakeArgs(10)
					πTemp003 = πg.NewInt(3).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp004
					πTemp003 = πg.NewInt(4).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[1] = πTemp004
					πTemp003 = πg.NewInt(5).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[2] = πTemp004
					πTemp003 = πg.NewInt(6).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[3] = πTemp004
					πTemp003 = πg.NewInt(7).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[4] = πTemp004
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[5] = πTemp004
					πTemp003 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[6] = πTemp004
					πTemp003 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[7] = πTemp004
					πTemp001[8] = πg.NewInt(13).ToObject()
					πTemp001[9] = πg.NewInt(2162078206).ToObject()
					if πE = πg.CheckLocal(πF, µRND, "RND"); πE != nil {
						continue
					}
					if πTemp003, πE = µRND.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp008 = πg.NewInt(6).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp008, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp004 = πg.NewInt(2).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp004, πTemp007); πE != nil {
						continue
					}
					// line 60: ss[5], ss[1] = RND(ss[2],ss[3],ss[4],ss[5],ss[6],ss[7],ss[0],ss[1],14,0x9bdc06a7);
					πF.SetLineno(60)
					πTemp001 = πF.MakeArgs(10)
					πTemp003 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp004
					πTemp003 = πg.NewInt(3).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[1] = πTemp004
					πTemp003 = πg.NewInt(4).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[2] = πTemp004
					πTemp003 = πg.NewInt(5).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[3] = πTemp004
					πTemp003 = πg.NewInt(6).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[4] = πTemp004
					πTemp003 = πg.NewInt(7).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[5] = πTemp004
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[6] = πTemp004
					πTemp003 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[7] = πTemp004
					πTemp001[8] = πg.NewInt(14).ToObject()
					πTemp001[9] = πg.NewInt(2614888103).ToObject()
					if πE = πg.CheckLocal(πF, µRND, "RND"); πE != nil {
						continue
					}
					if πTemp003, πE = µRND.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp008 = πg.NewInt(5).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp008, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp004 = πg.NewInt(1).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp004, πTemp007); πE != nil {
						continue
					}
					// line 61: ss[4], ss[0] = RND(ss[1],ss[2],ss[3],ss[4],ss[5],ss[6],ss[7],ss[0],15,0xc19bf174);
					πF.SetLineno(61)
					πTemp001 = πF.MakeArgs(10)
					πTemp003 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp004
					πTemp003 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[1] = πTemp004
					πTemp003 = πg.NewInt(3).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[2] = πTemp004
					πTemp003 = πg.NewInt(4).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[3] = πTemp004
					πTemp003 = πg.NewInt(5).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[4] = πTemp004
					πTemp003 = πg.NewInt(6).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[5] = πTemp004
					πTemp003 = πg.NewInt(7).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[6] = πTemp004
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[7] = πTemp004
					πTemp001[8] = πg.NewInt(15).ToObject()
					πTemp001[9] = πg.NewInt(3248222580).ToObject()
					if πE = πg.CheckLocal(πF, µRND, "RND"); πE != nil {
						continue
					}
					if πTemp003, πE = µRND.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp008 = πg.NewInt(4).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp008, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp004 = πg.NewInt(0).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp004, πTemp007); πE != nil {
						continue
					}
					// line 62: ss[3], ss[7] = RND(ss[0],ss[1],ss[2],ss[3],ss[4],ss[5],ss[6],ss[7],16,0xe49b69c1);
					πF.SetLineno(62)
					πTemp001 = πF.MakeArgs(10)
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp004
					πTemp003 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[1] = πTemp004
					πTemp003 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[2] = πTemp004
					πTemp003 = πg.NewInt(3).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[3] = πTemp004
					πTemp003 = πg.NewInt(4).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[4] = πTemp004
					πTemp003 = πg.NewInt(5).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[5] = πTemp004
					πTemp003 = πg.NewInt(6).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[6] = πTemp004
					πTemp003 = πg.NewInt(7).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[7] = πTemp004
					πTemp001[8] = πg.NewInt(16).ToObject()
					πTemp001[9] = πg.NewInt(3835390401).ToObject()
					if πE = πg.CheckLocal(πF, µRND, "RND"); πE != nil {
						continue
					}
					if πTemp003, πE = µRND.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp008 = πg.NewInt(3).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp008, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp004 = πg.NewInt(7).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp004, πTemp007); πE != nil {
						continue
					}
					// line 63: ss[2], ss[6] = RND(ss[7],ss[0],ss[1],ss[2],ss[3],ss[4],ss[5],ss[6],17,0xefbe4786);
					πF.SetLineno(63)
					πTemp001 = πF.MakeArgs(10)
					πTemp003 = πg.NewInt(7).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp004
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[1] = πTemp004
					πTemp003 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[2] = πTemp004
					πTemp003 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[3] = πTemp004
					πTemp003 = πg.NewInt(3).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[4] = πTemp004
					πTemp003 = πg.NewInt(4).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[5] = πTemp004
					πTemp003 = πg.NewInt(5).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[6] = πTemp004
					πTemp003 = πg.NewInt(6).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[7] = πTemp004
					πTemp001[8] = πg.NewInt(17).ToObject()
					πTemp001[9] = πg.NewInt(4022224774).ToObject()
					if πE = πg.CheckLocal(πF, µRND, "RND"); πE != nil {
						continue
					}
					if πTemp003, πE = µRND.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp008 = πg.NewInt(2).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp008, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp004 = πg.NewInt(6).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp004, πTemp007); πE != nil {
						continue
					}
					// line 64: ss[1], ss[5] = RND(ss[6],ss[7],ss[0],ss[1],ss[2],ss[3],ss[4],ss[5],18,0x0fc19dc6);
					πF.SetLineno(64)
					πTemp001 = πF.MakeArgs(10)
					πTemp003 = πg.NewInt(6).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp004
					πTemp003 = πg.NewInt(7).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[1] = πTemp004
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[2] = πTemp004
					πTemp003 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[3] = πTemp004
					πTemp003 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[4] = πTemp004
					πTemp003 = πg.NewInt(3).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[5] = πTemp004
					πTemp003 = πg.NewInt(4).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[6] = πTemp004
					πTemp003 = πg.NewInt(5).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[7] = πTemp004
					πTemp001[8] = πg.NewInt(18).ToObject()
					πTemp001[9] = πg.NewInt(264347078).ToObject()
					if πE = πg.CheckLocal(πF, µRND, "RND"); πE != nil {
						continue
					}
					if πTemp003, πE = µRND.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp008 = πg.NewInt(1).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp008, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp004 = πg.NewInt(5).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp004, πTemp007); πE != nil {
						continue
					}
					// line 65: ss[0], ss[4] = RND(ss[5],ss[6],ss[7],ss[0],ss[1],ss[2],ss[3],ss[4],19,0x240ca1cc);
					πF.SetLineno(65)
					πTemp001 = πF.MakeArgs(10)
					πTemp003 = πg.NewInt(5).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp004
					πTemp003 = πg.NewInt(6).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[1] = πTemp004
					πTemp003 = πg.NewInt(7).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[2] = πTemp004
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[3] = πTemp004
					πTemp003 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[4] = πTemp004
					πTemp003 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[5] = πTemp004
					πTemp003 = πg.NewInt(3).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[6] = πTemp004
					πTemp003 = πg.NewInt(4).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[7] = πTemp004
					πTemp001[8] = πg.NewInt(19).ToObject()
					πTemp001[9] = πg.NewInt(604807628).ToObject()
					if πE = πg.CheckLocal(πF, µRND, "RND"); πE != nil {
						continue
					}
					if πTemp003, πE = µRND.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp008 = πg.NewInt(0).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp008, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp004 = πg.NewInt(4).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp004, πTemp007); πE != nil {
						continue
					}
					// line 66: ss[7], ss[3] = RND(ss[4],ss[5],ss[6],ss[7],ss[0],ss[1],ss[2],ss[3],20,0x2de92c6f);
					πF.SetLineno(66)
					πTemp001 = πF.MakeArgs(10)
					πTemp003 = πg.NewInt(4).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp004
					πTemp003 = πg.NewInt(5).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[1] = πTemp004
					πTemp003 = πg.NewInt(6).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[2] = πTemp004
					πTemp003 = πg.NewInt(7).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[3] = πTemp004
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[4] = πTemp004
					πTemp003 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[5] = πTemp004
					πTemp003 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[6] = πTemp004
					πTemp003 = πg.NewInt(3).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[7] = πTemp004
					πTemp001[8] = πg.NewInt(20).ToObject()
					πTemp001[9] = πg.NewInt(770255983).ToObject()
					if πE = πg.CheckLocal(πF, µRND, "RND"); πE != nil {
						continue
					}
					if πTemp003, πE = µRND.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp008 = πg.NewInt(7).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp008, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp004 = πg.NewInt(3).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp004, πTemp007); πE != nil {
						continue
					}
					// line 67: ss[6], ss[2] = RND(ss[3],ss[4],ss[5],ss[6],ss[7],ss[0],ss[1],ss[2],21,0x4a7484aa);
					πF.SetLineno(67)
					πTemp001 = πF.MakeArgs(10)
					πTemp003 = πg.NewInt(3).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp004
					πTemp003 = πg.NewInt(4).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[1] = πTemp004
					πTemp003 = πg.NewInt(5).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[2] = πTemp004
					πTemp003 = πg.NewInt(6).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[3] = πTemp004
					πTemp003 = πg.NewInt(7).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[4] = πTemp004
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[5] = πTemp004
					πTemp003 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[6] = πTemp004
					πTemp003 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[7] = πTemp004
					πTemp001[8] = πg.NewInt(21).ToObject()
					πTemp001[9] = πg.NewInt(1249150122).ToObject()
					if πE = πg.CheckLocal(πF, µRND, "RND"); πE != nil {
						continue
					}
					if πTemp003, πE = µRND.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp008 = πg.NewInt(6).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp008, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp004 = πg.NewInt(2).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp004, πTemp007); πE != nil {
						continue
					}
					// line 68: ss[5], ss[1] = RND(ss[2],ss[3],ss[4],ss[5],ss[6],ss[7],ss[0],ss[1],22,0x5cb0a9dc);
					πF.SetLineno(68)
					πTemp001 = πF.MakeArgs(10)
					πTemp003 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp004
					πTemp003 = πg.NewInt(3).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[1] = πTemp004
					πTemp003 = πg.NewInt(4).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[2] = πTemp004
					πTemp003 = πg.NewInt(5).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[3] = πTemp004
					πTemp003 = πg.NewInt(6).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[4] = πTemp004
					πTemp003 = πg.NewInt(7).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[5] = πTemp004
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[6] = πTemp004
					πTemp003 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[7] = πTemp004
					πTemp001[8] = πg.NewInt(22).ToObject()
					πTemp001[9] = πg.NewInt(1555081692).ToObject()
					if πE = πg.CheckLocal(πF, µRND, "RND"); πE != nil {
						continue
					}
					if πTemp003, πE = µRND.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp008 = πg.NewInt(5).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp008, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp004 = πg.NewInt(1).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp004, πTemp007); πE != nil {
						continue
					}
					// line 69: ss[4], ss[0] = RND(ss[1],ss[2],ss[3],ss[4],ss[5],ss[6],ss[7],ss[0],23,0x76f988da);
					πF.SetLineno(69)
					πTemp001 = πF.MakeArgs(10)
					πTemp003 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp004
					πTemp003 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[1] = πTemp004
					πTemp003 = πg.NewInt(3).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[2] = πTemp004
					πTemp003 = πg.NewInt(4).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[3] = πTemp004
					πTemp003 = πg.NewInt(5).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[4] = πTemp004
					πTemp003 = πg.NewInt(6).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[5] = πTemp004
					πTemp003 = πg.NewInt(7).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[6] = πTemp004
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[7] = πTemp004
					πTemp001[8] = πg.NewInt(23).ToObject()
					πTemp001[9] = πg.NewInt(1996064986).ToObject()
					if πE = πg.CheckLocal(πF, µRND, "RND"); πE != nil {
						continue
					}
					if πTemp003, πE = µRND.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp008 = πg.NewInt(4).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp008, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp004 = πg.NewInt(0).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp004, πTemp007); πE != nil {
						continue
					}
					// line 70: ss[3], ss[7] = RND(ss[0],ss[1],ss[2],ss[3],ss[4],ss[5],ss[6],ss[7],24,0x983e5152);
					πF.SetLineno(70)
					πTemp001 = πF.MakeArgs(10)
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp004
					πTemp003 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[1] = πTemp004
					πTemp003 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[2] = πTemp004
					πTemp003 = πg.NewInt(3).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[3] = πTemp004
					πTemp003 = πg.NewInt(4).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[4] = πTemp004
					πTemp003 = πg.NewInt(5).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[5] = πTemp004
					πTemp003 = πg.NewInt(6).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[6] = πTemp004
					πTemp003 = πg.NewInt(7).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[7] = πTemp004
					πTemp001[8] = πg.NewInt(24).ToObject()
					πTemp001[9] = πg.NewInt(2554220882).ToObject()
					if πE = πg.CheckLocal(πF, µRND, "RND"); πE != nil {
						continue
					}
					if πTemp003, πE = µRND.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp008 = πg.NewInt(3).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp008, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp004 = πg.NewInt(7).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp004, πTemp007); πE != nil {
						continue
					}
					// line 71: ss[2], ss[6] = RND(ss[7],ss[0],ss[1],ss[2],ss[3],ss[4],ss[5],ss[6],25,0xa831c66d);
					πF.SetLineno(71)
					πTemp001 = πF.MakeArgs(10)
					πTemp003 = πg.NewInt(7).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp004
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[1] = πTemp004
					πTemp003 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[2] = πTemp004
					πTemp003 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[3] = πTemp004
					πTemp003 = πg.NewInt(3).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[4] = πTemp004
					πTemp003 = πg.NewInt(4).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[5] = πTemp004
					πTemp003 = πg.NewInt(5).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[6] = πTemp004
					πTemp003 = πg.NewInt(6).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[7] = πTemp004
					πTemp001[8] = πg.NewInt(25).ToObject()
					πTemp001[9] = πg.NewInt(2821834349).ToObject()
					if πE = πg.CheckLocal(πF, µRND, "RND"); πE != nil {
						continue
					}
					if πTemp003, πE = µRND.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp008 = πg.NewInt(2).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp008, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp004 = πg.NewInt(6).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp004, πTemp007); πE != nil {
						continue
					}
					// line 72: ss[1], ss[5] = RND(ss[6],ss[7],ss[0],ss[1],ss[2],ss[3],ss[4],ss[5],26,0xb00327c8);
					πF.SetLineno(72)
					πTemp001 = πF.MakeArgs(10)
					πTemp003 = πg.NewInt(6).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp004
					πTemp003 = πg.NewInt(7).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[1] = πTemp004
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[2] = πTemp004
					πTemp003 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[3] = πTemp004
					πTemp003 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[4] = πTemp004
					πTemp003 = πg.NewInt(3).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[5] = πTemp004
					πTemp003 = πg.NewInt(4).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[6] = πTemp004
					πTemp003 = πg.NewInt(5).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[7] = πTemp004
					πTemp001[8] = πg.NewInt(26).ToObject()
					πTemp001[9] = πg.NewInt(2952996808).ToObject()
					if πE = πg.CheckLocal(πF, µRND, "RND"); πE != nil {
						continue
					}
					if πTemp003, πE = µRND.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp008 = πg.NewInt(1).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp008, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp004 = πg.NewInt(5).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp004, πTemp007); πE != nil {
						continue
					}
					// line 73: ss[0], ss[4] = RND(ss[5],ss[6],ss[7],ss[0],ss[1],ss[2],ss[3],ss[4],27,0xbf597fc7);
					πF.SetLineno(73)
					πTemp001 = πF.MakeArgs(10)
					πTemp003 = πg.NewInt(5).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp004
					πTemp003 = πg.NewInt(6).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[1] = πTemp004
					πTemp003 = πg.NewInt(7).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[2] = πTemp004
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[3] = πTemp004
					πTemp003 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[4] = πTemp004
					πTemp003 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[5] = πTemp004
					πTemp003 = πg.NewInt(3).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[6] = πTemp004
					πTemp003 = πg.NewInt(4).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[7] = πTemp004
					πTemp001[8] = πg.NewInt(27).ToObject()
					πTemp001[9] = πg.NewInt(3210313671).ToObject()
					if πE = πg.CheckLocal(πF, µRND, "RND"); πE != nil {
						continue
					}
					if πTemp003, πE = µRND.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp008 = πg.NewInt(0).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp008, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp004 = πg.NewInt(4).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp004, πTemp007); πE != nil {
						continue
					}
					// line 74: ss[7], ss[3] = RND(ss[4],ss[5],ss[6],ss[7],ss[0],ss[1],ss[2],ss[3],28,0xc6e00bf3);
					πF.SetLineno(74)
					πTemp001 = πF.MakeArgs(10)
					πTemp003 = πg.NewInt(4).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp004
					πTemp003 = πg.NewInt(5).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[1] = πTemp004
					πTemp003 = πg.NewInt(6).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[2] = πTemp004
					πTemp003 = πg.NewInt(7).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[3] = πTemp004
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[4] = πTemp004
					πTemp003 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[5] = πTemp004
					πTemp003 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[6] = πTemp004
					πTemp003 = πg.NewInt(3).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[7] = πTemp004
					πTemp001[8] = πg.NewInt(28).ToObject()
					πTemp001[9] = πg.NewInt(3336571891).ToObject()
					if πE = πg.CheckLocal(πF, µRND, "RND"); πE != nil {
						continue
					}
					if πTemp003, πE = µRND.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp008 = πg.NewInt(7).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp008, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp004 = πg.NewInt(3).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp004, πTemp007); πE != nil {
						continue
					}
					// line 75: ss[6], ss[2] = RND(ss[3],ss[4],ss[5],ss[6],ss[7],ss[0],ss[1],ss[2],29,0xd5a79147);
					πF.SetLineno(75)
					πTemp001 = πF.MakeArgs(10)
					πTemp003 = πg.NewInt(3).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp004
					πTemp003 = πg.NewInt(4).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[1] = πTemp004
					πTemp003 = πg.NewInt(5).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[2] = πTemp004
					πTemp003 = πg.NewInt(6).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[3] = πTemp004
					πTemp003 = πg.NewInt(7).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[4] = πTemp004
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[5] = πTemp004
					πTemp003 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[6] = πTemp004
					πTemp003 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[7] = πTemp004
					πTemp001[8] = πg.NewInt(29).ToObject()
					πTemp001[9] = πg.NewInt(3584528711).ToObject()
					if πE = πg.CheckLocal(πF, µRND, "RND"); πE != nil {
						continue
					}
					if πTemp003, πE = µRND.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp008 = πg.NewInt(6).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp008, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp004 = πg.NewInt(2).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp004, πTemp007); πE != nil {
						continue
					}
					// line 76: ss[5], ss[1] = RND(ss[2],ss[3],ss[4],ss[5],ss[6],ss[7],ss[0],ss[1],30,0x06ca6351);
					πF.SetLineno(76)
					πTemp001 = πF.MakeArgs(10)
					πTemp003 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp004
					πTemp003 = πg.NewInt(3).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[1] = πTemp004
					πTemp003 = πg.NewInt(4).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[2] = πTemp004
					πTemp003 = πg.NewInt(5).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[3] = πTemp004
					πTemp003 = πg.NewInt(6).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[4] = πTemp004
					πTemp003 = πg.NewInt(7).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[5] = πTemp004
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[6] = πTemp004
					πTemp003 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[7] = πTemp004
					πTemp001[8] = πg.NewInt(30).ToObject()
					πTemp001[9] = πg.NewInt(113926993).ToObject()
					if πE = πg.CheckLocal(πF, µRND, "RND"); πE != nil {
						continue
					}
					if πTemp003, πE = µRND.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp008 = πg.NewInt(5).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp008, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp004 = πg.NewInt(1).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp004, πTemp007); πE != nil {
						continue
					}
					// line 77: ss[4], ss[0] = RND(ss[1],ss[2],ss[3],ss[4],ss[5],ss[6],ss[7],ss[0],31,0x14292967);
					πF.SetLineno(77)
					πTemp001 = πF.MakeArgs(10)
					πTemp003 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp004
					πTemp003 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[1] = πTemp004
					πTemp003 = πg.NewInt(3).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[2] = πTemp004
					πTemp003 = πg.NewInt(4).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[3] = πTemp004
					πTemp003 = πg.NewInt(5).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[4] = πTemp004
					πTemp003 = πg.NewInt(6).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[5] = πTemp004
					πTemp003 = πg.NewInt(7).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[6] = πTemp004
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[7] = πTemp004
					πTemp001[8] = πg.NewInt(31).ToObject()
					πTemp001[9] = πg.NewInt(338241895).ToObject()
					if πE = πg.CheckLocal(πF, µRND, "RND"); πE != nil {
						continue
					}
					if πTemp003, πE = µRND.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp008 = πg.NewInt(4).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp008, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp004 = πg.NewInt(0).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp004, πTemp007); πE != nil {
						continue
					}
					// line 78: ss[3], ss[7] = RND(ss[0],ss[1],ss[2],ss[3],ss[4],ss[5],ss[6],ss[7],32,0x27b70a85);
					πF.SetLineno(78)
					πTemp001 = πF.MakeArgs(10)
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp004
					πTemp003 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[1] = πTemp004
					πTemp003 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[2] = πTemp004
					πTemp003 = πg.NewInt(3).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[3] = πTemp004
					πTemp003 = πg.NewInt(4).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[4] = πTemp004
					πTemp003 = πg.NewInt(5).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[5] = πTemp004
					πTemp003 = πg.NewInt(6).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[6] = πTemp004
					πTemp003 = πg.NewInt(7).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[7] = πTemp004
					πTemp001[8] = πg.NewInt(32).ToObject()
					πTemp001[9] = πg.NewInt(666307205).ToObject()
					if πE = πg.CheckLocal(πF, µRND, "RND"); πE != nil {
						continue
					}
					if πTemp003, πE = µRND.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp008 = πg.NewInt(3).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp008, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp004 = πg.NewInt(7).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp004, πTemp007); πE != nil {
						continue
					}
					// line 79: ss[2], ss[6] = RND(ss[7],ss[0],ss[1],ss[2],ss[3],ss[4],ss[5],ss[6],33,0x2e1b2138);
					πF.SetLineno(79)
					πTemp001 = πF.MakeArgs(10)
					πTemp003 = πg.NewInt(7).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp004
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[1] = πTemp004
					πTemp003 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[2] = πTemp004
					πTemp003 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[3] = πTemp004
					πTemp003 = πg.NewInt(3).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[4] = πTemp004
					πTemp003 = πg.NewInt(4).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[5] = πTemp004
					πTemp003 = πg.NewInt(5).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[6] = πTemp004
					πTemp003 = πg.NewInt(6).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[7] = πTemp004
					πTemp001[8] = πg.NewInt(33).ToObject()
					πTemp001[9] = πg.NewInt(773529912).ToObject()
					if πE = πg.CheckLocal(πF, µRND, "RND"); πE != nil {
						continue
					}
					if πTemp003, πE = µRND.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp008 = πg.NewInt(2).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp008, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp004 = πg.NewInt(6).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp004, πTemp007); πE != nil {
						continue
					}
					// line 80: ss[1], ss[5] = RND(ss[6],ss[7],ss[0],ss[1],ss[2],ss[3],ss[4],ss[5],34,0x4d2c6dfc);
					πF.SetLineno(80)
					πTemp001 = πF.MakeArgs(10)
					πTemp003 = πg.NewInt(6).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp004
					πTemp003 = πg.NewInt(7).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[1] = πTemp004
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[2] = πTemp004
					πTemp003 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[3] = πTemp004
					πTemp003 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[4] = πTemp004
					πTemp003 = πg.NewInt(3).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[5] = πTemp004
					πTemp003 = πg.NewInt(4).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[6] = πTemp004
					πTemp003 = πg.NewInt(5).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[7] = πTemp004
					πTemp001[8] = πg.NewInt(34).ToObject()
					πTemp001[9] = πg.NewInt(1294757372).ToObject()
					if πE = πg.CheckLocal(πF, µRND, "RND"); πE != nil {
						continue
					}
					if πTemp003, πE = µRND.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp008 = πg.NewInt(1).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp008, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp004 = πg.NewInt(5).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp004, πTemp007); πE != nil {
						continue
					}
					// line 81: ss[0], ss[4] = RND(ss[5],ss[6],ss[7],ss[0],ss[1],ss[2],ss[3],ss[4],35,0x53380d13);
					πF.SetLineno(81)
					πTemp001 = πF.MakeArgs(10)
					πTemp003 = πg.NewInt(5).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp004
					πTemp003 = πg.NewInt(6).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[1] = πTemp004
					πTemp003 = πg.NewInt(7).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[2] = πTemp004
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[3] = πTemp004
					πTemp003 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[4] = πTemp004
					πTemp003 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[5] = πTemp004
					πTemp003 = πg.NewInt(3).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[6] = πTemp004
					πTemp003 = πg.NewInt(4).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[7] = πTemp004
					πTemp001[8] = πg.NewInt(35).ToObject()
					πTemp001[9] = πg.NewInt(1396182291).ToObject()
					if πE = πg.CheckLocal(πF, µRND, "RND"); πE != nil {
						continue
					}
					if πTemp003, πE = µRND.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp008 = πg.NewInt(0).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp008, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp004 = πg.NewInt(4).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp004, πTemp007); πE != nil {
						continue
					}
					// line 82: ss[7], ss[3] = RND(ss[4],ss[5],ss[6],ss[7],ss[0],ss[1],ss[2],ss[3],36,0x650a7354);
					πF.SetLineno(82)
					πTemp001 = πF.MakeArgs(10)
					πTemp003 = πg.NewInt(4).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp004
					πTemp003 = πg.NewInt(5).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[1] = πTemp004
					πTemp003 = πg.NewInt(6).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[2] = πTemp004
					πTemp003 = πg.NewInt(7).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[3] = πTemp004
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[4] = πTemp004
					πTemp003 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[5] = πTemp004
					πTemp003 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[6] = πTemp004
					πTemp003 = πg.NewInt(3).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[7] = πTemp004
					πTemp001[8] = πg.NewInt(36).ToObject()
					πTemp001[9] = πg.NewInt(1695183700).ToObject()
					if πE = πg.CheckLocal(πF, µRND, "RND"); πE != nil {
						continue
					}
					if πTemp003, πE = µRND.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp008 = πg.NewInt(7).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp008, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp004 = πg.NewInt(3).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp004, πTemp007); πE != nil {
						continue
					}
					// line 83: ss[6], ss[2] = RND(ss[3],ss[4],ss[5],ss[6],ss[7],ss[0],ss[1],ss[2],37,0x766a0abb);
					πF.SetLineno(83)
					πTemp001 = πF.MakeArgs(10)
					πTemp003 = πg.NewInt(3).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp004
					πTemp003 = πg.NewInt(4).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[1] = πTemp004
					πTemp003 = πg.NewInt(5).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[2] = πTemp004
					πTemp003 = πg.NewInt(6).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[3] = πTemp004
					πTemp003 = πg.NewInt(7).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[4] = πTemp004
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[5] = πTemp004
					πTemp003 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[6] = πTemp004
					πTemp003 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[7] = πTemp004
					πTemp001[8] = πg.NewInt(37).ToObject()
					πTemp001[9] = πg.NewInt(1986661051).ToObject()
					if πE = πg.CheckLocal(πF, µRND, "RND"); πE != nil {
						continue
					}
					if πTemp003, πE = µRND.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp008 = πg.NewInt(6).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp008, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp004 = πg.NewInt(2).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp004, πTemp007); πE != nil {
						continue
					}
					// line 84: ss[5], ss[1] = RND(ss[2],ss[3],ss[4],ss[5],ss[6],ss[7],ss[0],ss[1],38,0x81c2c92e);
					πF.SetLineno(84)
					πTemp001 = πF.MakeArgs(10)
					πTemp003 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp004
					πTemp003 = πg.NewInt(3).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[1] = πTemp004
					πTemp003 = πg.NewInt(4).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[2] = πTemp004
					πTemp003 = πg.NewInt(5).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[3] = πTemp004
					πTemp003 = πg.NewInt(6).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[4] = πTemp004
					πTemp003 = πg.NewInt(7).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[5] = πTemp004
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[6] = πTemp004
					πTemp003 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[7] = πTemp004
					πTemp001[8] = πg.NewInt(38).ToObject()
					πTemp001[9] = πg.NewInt(2177026350).ToObject()
					if πE = πg.CheckLocal(πF, µRND, "RND"); πE != nil {
						continue
					}
					if πTemp003, πE = µRND.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp008 = πg.NewInt(5).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp008, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp004 = πg.NewInt(1).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp004, πTemp007); πE != nil {
						continue
					}
					// line 85: ss[4], ss[0] = RND(ss[1],ss[2],ss[3],ss[4],ss[5],ss[6],ss[7],ss[0],39,0x92722c85);
					πF.SetLineno(85)
					πTemp001 = πF.MakeArgs(10)
					πTemp003 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp004
					πTemp003 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[1] = πTemp004
					πTemp003 = πg.NewInt(3).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[2] = πTemp004
					πTemp003 = πg.NewInt(4).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[3] = πTemp004
					πTemp003 = πg.NewInt(5).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[4] = πTemp004
					πTemp003 = πg.NewInt(6).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[5] = πTemp004
					πTemp003 = πg.NewInt(7).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[6] = πTemp004
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[7] = πTemp004
					πTemp001[8] = πg.NewInt(39).ToObject()
					πTemp001[9] = πg.NewInt(2456956037).ToObject()
					if πE = πg.CheckLocal(πF, µRND, "RND"); πE != nil {
						continue
					}
					if πTemp003, πE = µRND.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp008 = πg.NewInt(4).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp008, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp004 = πg.NewInt(0).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp004, πTemp007); πE != nil {
						continue
					}
					// line 86: ss[3], ss[7] = RND(ss[0],ss[1],ss[2],ss[3],ss[4],ss[5],ss[6],ss[7],40,0xa2bfe8a1);
					πF.SetLineno(86)
					πTemp001 = πF.MakeArgs(10)
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp004
					πTemp003 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[1] = πTemp004
					πTemp003 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[2] = πTemp004
					πTemp003 = πg.NewInt(3).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[3] = πTemp004
					πTemp003 = πg.NewInt(4).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[4] = πTemp004
					πTemp003 = πg.NewInt(5).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[5] = πTemp004
					πTemp003 = πg.NewInt(6).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[6] = πTemp004
					πTemp003 = πg.NewInt(7).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[7] = πTemp004
					πTemp001[8] = πg.NewInt(40).ToObject()
					πTemp001[9] = πg.NewInt(2730485921).ToObject()
					if πE = πg.CheckLocal(πF, µRND, "RND"); πE != nil {
						continue
					}
					if πTemp003, πE = µRND.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp008 = πg.NewInt(3).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp008, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp004 = πg.NewInt(7).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp004, πTemp007); πE != nil {
						continue
					}
					// line 87: ss[2], ss[6] = RND(ss[7],ss[0],ss[1],ss[2],ss[3],ss[4],ss[5],ss[6],41,0xa81a664b);
					πF.SetLineno(87)
					πTemp001 = πF.MakeArgs(10)
					πTemp003 = πg.NewInt(7).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp004
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[1] = πTemp004
					πTemp003 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[2] = πTemp004
					πTemp003 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[3] = πTemp004
					πTemp003 = πg.NewInt(3).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[4] = πTemp004
					πTemp003 = πg.NewInt(4).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[5] = πTemp004
					πTemp003 = πg.NewInt(5).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[6] = πTemp004
					πTemp003 = πg.NewInt(6).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[7] = πTemp004
					πTemp001[8] = πg.NewInt(41).ToObject()
					πTemp001[9] = πg.NewInt(2820302411).ToObject()
					if πE = πg.CheckLocal(πF, µRND, "RND"); πE != nil {
						continue
					}
					if πTemp003, πE = µRND.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp008 = πg.NewInt(2).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp008, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp004 = πg.NewInt(6).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp004, πTemp007); πE != nil {
						continue
					}
					// line 88: ss[1], ss[5] = RND(ss[6],ss[7],ss[0],ss[1],ss[2],ss[3],ss[4],ss[5],42,0xc24b8b70);
					πF.SetLineno(88)
					πTemp001 = πF.MakeArgs(10)
					πTemp003 = πg.NewInt(6).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp004
					πTemp003 = πg.NewInt(7).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[1] = πTemp004
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[2] = πTemp004
					πTemp003 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[3] = πTemp004
					πTemp003 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[4] = πTemp004
					πTemp003 = πg.NewInt(3).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[5] = πTemp004
					πTemp003 = πg.NewInt(4).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[6] = πTemp004
					πTemp003 = πg.NewInt(5).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[7] = πTemp004
					πTemp001[8] = πg.NewInt(42).ToObject()
					πTemp001[9] = πg.NewInt(3259730800).ToObject()
					if πE = πg.CheckLocal(πF, µRND, "RND"); πE != nil {
						continue
					}
					if πTemp003, πE = µRND.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp008 = πg.NewInt(1).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp008, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp004 = πg.NewInt(5).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp004, πTemp007); πE != nil {
						continue
					}
					// line 89: ss[0], ss[4] = RND(ss[5],ss[6],ss[7],ss[0],ss[1],ss[2],ss[3],ss[4],43,0xc76c51a3);
					πF.SetLineno(89)
					πTemp001 = πF.MakeArgs(10)
					πTemp003 = πg.NewInt(5).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp004
					πTemp003 = πg.NewInt(6).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[1] = πTemp004
					πTemp003 = πg.NewInt(7).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[2] = πTemp004
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[3] = πTemp004
					πTemp003 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[4] = πTemp004
					πTemp003 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[5] = πTemp004
					πTemp003 = πg.NewInt(3).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[6] = πTemp004
					πTemp003 = πg.NewInt(4).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[7] = πTemp004
					πTemp001[8] = πg.NewInt(43).ToObject()
					πTemp001[9] = πg.NewInt(3345764771).ToObject()
					if πE = πg.CheckLocal(πF, µRND, "RND"); πE != nil {
						continue
					}
					if πTemp003, πE = µRND.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp008 = πg.NewInt(0).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp008, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp004 = πg.NewInt(4).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp004, πTemp007); πE != nil {
						continue
					}
					// line 90: ss[7], ss[3] = RND(ss[4],ss[5],ss[6],ss[7],ss[0],ss[1],ss[2],ss[3],44,0xd192e819);
					πF.SetLineno(90)
					πTemp001 = πF.MakeArgs(10)
					πTemp003 = πg.NewInt(4).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp004
					πTemp003 = πg.NewInt(5).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[1] = πTemp004
					πTemp003 = πg.NewInt(6).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[2] = πTemp004
					πTemp003 = πg.NewInt(7).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[3] = πTemp004
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[4] = πTemp004
					πTemp003 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[5] = πTemp004
					πTemp003 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[6] = πTemp004
					πTemp003 = πg.NewInt(3).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[7] = πTemp004
					πTemp001[8] = πg.NewInt(44).ToObject()
					πTemp001[9] = πg.NewInt(3516065817).ToObject()
					if πE = πg.CheckLocal(πF, µRND, "RND"); πE != nil {
						continue
					}
					if πTemp003, πE = µRND.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp008 = πg.NewInt(7).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp008, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp004 = πg.NewInt(3).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp004, πTemp007); πE != nil {
						continue
					}
					// line 91: ss[6], ss[2] = RND(ss[3],ss[4],ss[5],ss[6],ss[7],ss[0],ss[1],ss[2],45,0xd6990624);
					πF.SetLineno(91)
					πTemp001 = πF.MakeArgs(10)
					πTemp003 = πg.NewInt(3).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp004
					πTemp003 = πg.NewInt(4).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[1] = πTemp004
					πTemp003 = πg.NewInt(5).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[2] = πTemp004
					πTemp003 = πg.NewInt(6).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[3] = πTemp004
					πTemp003 = πg.NewInt(7).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[4] = πTemp004
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[5] = πTemp004
					πTemp003 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[6] = πTemp004
					πTemp003 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[7] = πTemp004
					πTemp001[8] = πg.NewInt(45).ToObject()
					πTemp001[9] = πg.NewInt(3600352804).ToObject()
					if πE = πg.CheckLocal(πF, µRND, "RND"); πE != nil {
						continue
					}
					if πTemp003, πE = µRND.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp008 = πg.NewInt(6).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp008, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp004 = πg.NewInt(2).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp004, πTemp007); πE != nil {
						continue
					}
					// line 92: ss[5], ss[1] = RND(ss[2],ss[3],ss[4],ss[5],ss[6],ss[7],ss[0],ss[1],46,0xf40e3585);
					πF.SetLineno(92)
					πTemp001 = πF.MakeArgs(10)
					πTemp003 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp004
					πTemp003 = πg.NewInt(3).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[1] = πTemp004
					πTemp003 = πg.NewInt(4).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[2] = πTemp004
					πTemp003 = πg.NewInt(5).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[3] = πTemp004
					πTemp003 = πg.NewInt(6).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[4] = πTemp004
					πTemp003 = πg.NewInt(7).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[5] = πTemp004
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[6] = πTemp004
					πTemp003 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[7] = πTemp004
					πTemp001[8] = πg.NewInt(46).ToObject()
					πTemp001[9] = πg.NewInt(4094571909).ToObject()
					if πE = πg.CheckLocal(πF, µRND, "RND"); πE != nil {
						continue
					}
					if πTemp003, πE = µRND.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp008 = πg.NewInt(5).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp008, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp004 = πg.NewInt(1).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp004, πTemp007); πE != nil {
						continue
					}
					// line 93: ss[4], ss[0] = RND(ss[1],ss[2],ss[3],ss[4],ss[5],ss[6],ss[7],ss[0],47,0x106aa070);
					πF.SetLineno(93)
					πTemp001 = πF.MakeArgs(10)
					πTemp003 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp004
					πTemp003 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[1] = πTemp004
					πTemp003 = πg.NewInt(3).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[2] = πTemp004
					πTemp003 = πg.NewInt(4).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[3] = πTemp004
					πTemp003 = πg.NewInt(5).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[4] = πTemp004
					πTemp003 = πg.NewInt(6).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[5] = πTemp004
					πTemp003 = πg.NewInt(7).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[6] = πTemp004
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[7] = πTemp004
					πTemp001[8] = πg.NewInt(47).ToObject()
					πTemp001[9] = πg.NewInt(275423344).ToObject()
					if πE = πg.CheckLocal(πF, µRND, "RND"); πE != nil {
						continue
					}
					if πTemp003, πE = µRND.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp008 = πg.NewInt(4).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp008, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp004 = πg.NewInt(0).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp004, πTemp007); πE != nil {
						continue
					}
					// line 94: ss[3], ss[7] = RND(ss[0],ss[1],ss[2],ss[3],ss[4],ss[5],ss[6],ss[7],48,0x19a4c116);
					πF.SetLineno(94)
					πTemp001 = πF.MakeArgs(10)
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp004
					πTemp003 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[1] = πTemp004
					πTemp003 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[2] = πTemp004
					πTemp003 = πg.NewInt(3).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[3] = πTemp004
					πTemp003 = πg.NewInt(4).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[4] = πTemp004
					πTemp003 = πg.NewInt(5).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[5] = πTemp004
					πTemp003 = πg.NewInt(6).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[6] = πTemp004
					πTemp003 = πg.NewInt(7).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[7] = πTemp004
					πTemp001[8] = πg.NewInt(48).ToObject()
					πTemp001[9] = πg.NewInt(430227734).ToObject()
					if πE = πg.CheckLocal(πF, µRND, "RND"); πE != nil {
						continue
					}
					if πTemp003, πE = µRND.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp008 = πg.NewInt(3).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp008, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp004 = πg.NewInt(7).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp004, πTemp007); πE != nil {
						continue
					}
					// line 95: ss[2], ss[6] = RND(ss[7],ss[0],ss[1],ss[2],ss[3],ss[4],ss[5],ss[6],49,0x1e376c08);
					πF.SetLineno(95)
					πTemp001 = πF.MakeArgs(10)
					πTemp003 = πg.NewInt(7).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp004
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[1] = πTemp004
					πTemp003 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[2] = πTemp004
					πTemp003 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[3] = πTemp004
					πTemp003 = πg.NewInt(3).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[4] = πTemp004
					πTemp003 = πg.NewInt(4).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[5] = πTemp004
					πTemp003 = πg.NewInt(5).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[6] = πTemp004
					πTemp003 = πg.NewInt(6).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[7] = πTemp004
					πTemp001[8] = πg.NewInt(49).ToObject()
					πTemp001[9] = πg.NewInt(506948616).ToObject()
					if πE = πg.CheckLocal(πF, µRND, "RND"); πE != nil {
						continue
					}
					if πTemp003, πE = µRND.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp008 = πg.NewInt(2).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp008, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp004 = πg.NewInt(6).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp004, πTemp007); πE != nil {
						continue
					}
					// line 96: ss[1], ss[5] = RND(ss[6],ss[7],ss[0],ss[1],ss[2],ss[3],ss[4],ss[5],50,0x2748774c);
					πF.SetLineno(96)
					πTemp001 = πF.MakeArgs(10)
					πTemp003 = πg.NewInt(6).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp004
					πTemp003 = πg.NewInt(7).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[1] = πTemp004
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[2] = πTemp004
					πTemp003 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[3] = πTemp004
					πTemp003 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[4] = πTemp004
					πTemp003 = πg.NewInt(3).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[5] = πTemp004
					πTemp003 = πg.NewInt(4).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[6] = πTemp004
					πTemp003 = πg.NewInt(5).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[7] = πTemp004
					πTemp001[8] = πg.NewInt(50).ToObject()
					πTemp001[9] = πg.NewInt(659060556).ToObject()
					if πE = πg.CheckLocal(πF, µRND, "RND"); πE != nil {
						continue
					}
					if πTemp003, πE = µRND.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp008 = πg.NewInt(1).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp008, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp004 = πg.NewInt(5).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp004, πTemp007); πE != nil {
						continue
					}
					// line 97: ss[0], ss[4] = RND(ss[5],ss[6],ss[7],ss[0],ss[1],ss[2],ss[3],ss[4],51,0x34b0bcb5);
					πF.SetLineno(97)
					πTemp001 = πF.MakeArgs(10)
					πTemp003 = πg.NewInt(5).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp004
					πTemp003 = πg.NewInt(6).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[1] = πTemp004
					πTemp003 = πg.NewInt(7).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[2] = πTemp004
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[3] = πTemp004
					πTemp003 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[4] = πTemp004
					πTemp003 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[5] = πTemp004
					πTemp003 = πg.NewInt(3).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[6] = πTemp004
					πTemp003 = πg.NewInt(4).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[7] = πTemp004
					πTemp001[8] = πg.NewInt(51).ToObject()
					πTemp001[9] = πg.NewInt(883997877).ToObject()
					if πE = πg.CheckLocal(πF, µRND, "RND"); πE != nil {
						continue
					}
					if πTemp003, πE = µRND.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp008 = πg.NewInt(0).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp008, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp004 = πg.NewInt(4).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp004, πTemp007); πE != nil {
						continue
					}
					// line 98: ss[7], ss[3] = RND(ss[4],ss[5],ss[6],ss[7],ss[0],ss[1],ss[2],ss[3],52,0x391c0cb3);
					πF.SetLineno(98)
					πTemp001 = πF.MakeArgs(10)
					πTemp003 = πg.NewInt(4).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp004
					πTemp003 = πg.NewInt(5).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[1] = πTemp004
					πTemp003 = πg.NewInt(6).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[2] = πTemp004
					πTemp003 = πg.NewInt(7).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[3] = πTemp004
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[4] = πTemp004
					πTemp003 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[5] = πTemp004
					πTemp003 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[6] = πTemp004
					πTemp003 = πg.NewInt(3).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[7] = πTemp004
					πTemp001[8] = πg.NewInt(52).ToObject()
					πTemp001[9] = πg.NewInt(958139571).ToObject()
					if πE = πg.CheckLocal(πF, µRND, "RND"); πE != nil {
						continue
					}
					if πTemp003, πE = µRND.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp008 = πg.NewInt(7).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp008, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp004 = πg.NewInt(3).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp004, πTemp007); πE != nil {
						continue
					}
					// line 99: ss[6], ss[2] = RND(ss[3],ss[4],ss[5],ss[6],ss[7],ss[0],ss[1],ss[2],53,0x4ed8aa4a);
					πF.SetLineno(99)
					πTemp001 = πF.MakeArgs(10)
					πTemp003 = πg.NewInt(3).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp004
					πTemp003 = πg.NewInt(4).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[1] = πTemp004
					πTemp003 = πg.NewInt(5).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[2] = πTemp004
					πTemp003 = πg.NewInt(6).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[3] = πTemp004
					πTemp003 = πg.NewInt(7).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[4] = πTemp004
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[5] = πTemp004
					πTemp003 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[6] = πTemp004
					πTemp003 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[7] = πTemp004
					πTemp001[8] = πg.NewInt(53).ToObject()
					πTemp001[9] = πg.NewInt(1322822218).ToObject()
					if πE = πg.CheckLocal(πF, µRND, "RND"); πE != nil {
						continue
					}
					if πTemp003, πE = µRND.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp008 = πg.NewInt(6).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp008, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp004 = πg.NewInt(2).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp004, πTemp007); πE != nil {
						continue
					}
					// line 100: ss[5], ss[1] = RND(ss[2],ss[3],ss[4],ss[5],ss[6],ss[7],ss[0],ss[1],54,0x5b9cca4f);
					πF.SetLineno(100)
					πTemp001 = πF.MakeArgs(10)
					πTemp003 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp004
					πTemp003 = πg.NewInt(3).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[1] = πTemp004
					πTemp003 = πg.NewInt(4).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[2] = πTemp004
					πTemp003 = πg.NewInt(5).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[3] = πTemp004
					πTemp003 = πg.NewInt(6).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[4] = πTemp004
					πTemp003 = πg.NewInt(7).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[5] = πTemp004
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[6] = πTemp004
					πTemp003 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[7] = πTemp004
					πTemp001[8] = πg.NewInt(54).ToObject()
					πTemp001[9] = πg.NewInt(1537002063).ToObject()
					if πE = πg.CheckLocal(πF, µRND, "RND"); πE != nil {
						continue
					}
					if πTemp003, πE = µRND.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp008 = πg.NewInt(5).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp008, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp004 = πg.NewInt(1).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp004, πTemp007); πE != nil {
						continue
					}
					// line 101: ss[4], ss[0] = RND(ss[1],ss[2],ss[3],ss[4],ss[5],ss[6],ss[7],ss[0],55,0x682e6ff3);
					πF.SetLineno(101)
					πTemp001 = πF.MakeArgs(10)
					πTemp003 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp004
					πTemp003 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[1] = πTemp004
					πTemp003 = πg.NewInt(3).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[2] = πTemp004
					πTemp003 = πg.NewInt(4).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[3] = πTemp004
					πTemp003 = πg.NewInt(5).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[4] = πTemp004
					πTemp003 = πg.NewInt(6).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[5] = πTemp004
					πTemp003 = πg.NewInt(7).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[6] = πTemp004
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[7] = πTemp004
					πTemp001[8] = πg.NewInt(55).ToObject()
					πTemp001[9] = πg.NewInt(1747873779).ToObject()
					if πE = πg.CheckLocal(πF, µRND, "RND"); πE != nil {
						continue
					}
					if πTemp003, πE = µRND.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp008 = πg.NewInt(4).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp008, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp004 = πg.NewInt(0).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp004, πTemp007); πE != nil {
						continue
					}
					// line 102: ss[3], ss[7] = RND(ss[0],ss[1],ss[2],ss[3],ss[4],ss[5],ss[6],ss[7],56,0x748f82ee);
					πF.SetLineno(102)
					πTemp001 = πF.MakeArgs(10)
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp004
					πTemp003 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[1] = πTemp004
					πTemp003 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[2] = πTemp004
					πTemp003 = πg.NewInt(3).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[3] = πTemp004
					πTemp003 = πg.NewInt(4).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[4] = πTemp004
					πTemp003 = πg.NewInt(5).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[5] = πTemp004
					πTemp003 = πg.NewInt(6).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[6] = πTemp004
					πTemp003 = πg.NewInt(7).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[7] = πTemp004
					πTemp001[8] = πg.NewInt(56).ToObject()
					πTemp001[9] = πg.NewInt(1955562222).ToObject()
					if πE = πg.CheckLocal(πF, µRND, "RND"); πE != nil {
						continue
					}
					if πTemp003, πE = µRND.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp008 = πg.NewInt(3).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp008, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp004 = πg.NewInt(7).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp004, πTemp007); πE != nil {
						continue
					}
					// line 103: ss[2], ss[6] = RND(ss[7],ss[0],ss[1],ss[2],ss[3],ss[4],ss[5],ss[6],57,0x78a5636f);
					πF.SetLineno(103)
					πTemp001 = πF.MakeArgs(10)
					πTemp003 = πg.NewInt(7).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp004
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[1] = πTemp004
					πTemp003 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[2] = πTemp004
					πTemp003 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[3] = πTemp004
					πTemp003 = πg.NewInt(3).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[4] = πTemp004
					πTemp003 = πg.NewInt(4).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[5] = πTemp004
					πTemp003 = πg.NewInt(5).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[6] = πTemp004
					πTemp003 = πg.NewInt(6).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[7] = πTemp004
					πTemp001[8] = πg.NewInt(57).ToObject()
					πTemp001[9] = πg.NewInt(2024104815).ToObject()
					if πE = πg.CheckLocal(πF, µRND, "RND"); πE != nil {
						continue
					}
					if πTemp003, πE = µRND.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp008 = πg.NewInt(2).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp008, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp004 = πg.NewInt(6).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp004, πTemp007); πE != nil {
						continue
					}
					// line 104: ss[1], ss[5] = RND(ss[6],ss[7],ss[0],ss[1],ss[2],ss[3],ss[4],ss[5],58,0x84c87814);
					πF.SetLineno(104)
					πTemp001 = πF.MakeArgs(10)
					πTemp003 = πg.NewInt(6).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp004
					πTemp003 = πg.NewInt(7).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[1] = πTemp004
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[2] = πTemp004
					πTemp003 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[3] = πTemp004
					πTemp003 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[4] = πTemp004
					πTemp003 = πg.NewInt(3).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[5] = πTemp004
					πTemp003 = πg.NewInt(4).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[6] = πTemp004
					πTemp003 = πg.NewInt(5).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[7] = πTemp004
					πTemp001[8] = πg.NewInt(58).ToObject()
					πTemp001[9] = πg.NewInt(2227730452).ToObject()
					if πE = πg.CheckLocal(πF, µRND, "RND"); πE != nil {
						continue
					}
					if πTemp003, πE = µRND.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp008 = πg.NewInt(1).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp008, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp004 = πg.NewInt(5).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp004, πTemp007); πE != nil {
						continue
					}
					// line 105: ss[0], ss[4] = RND(ss[5],ss[6],ss[7],ss[0],ss[1],ss[2],ss[3],ss[4],59,0x8cc70208);
					πF.SetLineno(105)
					πTemp001 = πF.MakeArgs(10)
					πTemp003 = πg.NewInt(5).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp004
					πTemp003 = πg.NewInt(6).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[1] = πTemp004
					πTemp003 = πg.NewInt(7).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[2] = πTemp004
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[3] = πTemp004
					πTemp003 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[4] = πTemp004
					πTemp003 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[5] = πTemp004
					πTemp003 = πg.NewInt(3).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[6] = πTemp004
					πTemp003 = πg.NewInt(4).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[7] = πTemp004
					πTemp001[8] = πg.NewInt(59).ToObject()
					πTemp001[9] = πg.NewInt(2361852424).ToObject()
					if πE = πg.CheckLocal(πF, µRND, "RND"); πE != nil {
						continue
					}
					if πTemp003, πE = µRND.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp008 = πg.NewInt(0).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp008, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp004 = πg.NewInt(4).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp004, πTemp007); πE != nil {
						continue
					}
					// line 106: ss[7], ss[3] = RND(ss[4],ss[5],ss[6],ss[7],ss[0],ss[1],ss[2],ss[3],60,0x90befffa);
					πF.SetLineno(106)
					πTemp001 = πF.MakeArgs(10)
					πTemp003 = πg.NewInt(4).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp004
					πTemp003 = πg.NewInt(5).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[1] = πTemp004
					πTemp003 = πg.NewInt(6).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[2] = πTemp004
					πTemp003 = πg.NewInt(7).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[3] = πTemp004
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[4] = πTemp004
					πTemp003 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[5] = πTemp004
					πTemp003 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[6] = πTemp004
					πTemp003 = πg.NewInt(3).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[7] = πTemp004
					πTemp001[8] = πg.NewInt(60).ToObject()
					πTemp001[9] = πg.NewInt(2428436474).ToObject()
					if πE = πg.CheckLocal(πF, µRND, "RND"); πE != nil {
						continue
					}
					if πTemp003, πE = µRND.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp008 = πg.NewInt(7).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp008, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp004 = πg.NewInt(3).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp004, πTemp007); πE != nil {
						continue
					}
					// line 107: ss[6], ss[2] = RND(ss[3],ss[4],ss[5],ss[6],ss[7],ss[0],ss[1],ss[2],61,0xa4506ceb);
					πF.SetLineno(107)
					πTemp001 = πF.MakeArgs(10)
					πTemp003 = πg.NewInt(3).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp004
					πTemp003 = πg.NewInt(4).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[1] = πTemp004
					πTemp003 = πg.NewInt(5).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[2] = πTemp004
					πTemp003 = πg.NewInt(6).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[3] = πTemp004
					πTemp003 = πg.NewInt(7).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[4] = πTemp004
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[5] = πTemp004
					πTemp003 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[6] = πTemp004
					πTemp003 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[7] = πTemp004
					πTemp001[8] = πg.NewInt(61).ToObject()
					πTemp001[9] = πg.NewInt(2756734187).ToObject()
					if πE = πg.CheckLocal(πF, µRND, "RND"); πE != nil {
						continue
					}
					if πTemp003, πE = µRND.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp008 = πg.NewInt(6).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp008, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp004 = πg.NewInt(2).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp004, πTemp007); πE != nil {
						continue
					}
					// line 108: ss[5], ss[1] = RND(ss[2],ss[3],ss[4],ss[5],ss[6],ss[7],ss[0],ss[1],62,0xbef9a3f7);
					πF.SetLineno(108)
					πTemp001 = πF.MakeArgs(10)
					πTemp003 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp004
					πTemp003 = πg.NewInt(3).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[1] = πTemp004
					πTemp003 = πg.NewInt(4).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[2] = πTemp004
					πTemp003 = πg.NewInt(5).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[3] = πTemp004
					πTemp003 = πg.NewInt(6).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[4] = πTemp004
					πTemp003 = πg.NewInt(7).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[5] = πTemp004
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[6] = πTemp004
					πTemp003 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[7] = πTemp004
					πTemp001[8] = πg.NewInt(62).ToObject()
					πTemp001[9] = πg.NewInt(3204031479).ToObject()
					if πE = πg.CheckLocal(πF, µRND, "RND"); πE != nil {
						continue
					}
					if πTemp003, πE = µRND.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp008 = πg.NewInt(5).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp008, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp004 = πg.NewInt(1).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp004, πTemp007); πE != nil {
						continue
					}
					// line 109: ss[4], ss[0] = RND(ss[1],ss[2],ss[3],ss[4],ss[5],ss[6],ss[7],ss[0],63,0xc67178f2);
					πF.SetLineno(109)
					πTemp001 = πF.MakeArgs(10)
					πTemp003 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp004
					πTemp003 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[1] = πTemp004
					πTemp003 = πg.NewInt(3).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[2] = πTemp004
					πTemp003 = πg.NewInt(4).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[3] = πTemp004
					πTemp003 = πg.NewInt(5).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[4] = πTemp004
					πTemp003 = πg.NewInt(6).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[5] = πTemp004
					πTemp003 = πg.NewInt(7).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[6] = πTemp004
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µss, πTemp003); πE != nil {
						continue
					}
					πTemp001[7] = πTemp004
					πTemp001[8] = πg.NewInt(63).ToObject()
					πTemp001[9] = πg.NewInt(3329325298).ToObject()
					if πE = πg.CheckLocal(πF, µRND, "RND"); πE != nil {
						continue
					}
					if πTemp003, πE = µRND.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp008 = πg.NewInt(4).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp008, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp004 = πg.NewInt(0).ToObject()
					if πE = πg.SetItem(πF, µss, πTemp004, πTemp007); πE != nil {
						continue
					}
					// line 111: dig = []
					πF.SetLineno(111)
					πTemp001 = make([]*πg.Object, 0)
					πTemp003 = πg.NewList(πTemp001...).ToObject()
					µdig = πTemp003
					πTemp001 = πF.MakeArgs(1)
					πTemp004 = ßdigest.ToObject()
					if πE = πg.CheckLocal(πF, µsha_info, "sha_info"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetItem(πF, µsha_info, πTemp004); πE != nil {
						continue
					}
					πTemp001[0] = πTemp007
					if πTemp004, πE = πg.ResolveGlobal(πF, ßenumerate); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp003, πE = πg.Iter(πF, πTemp007); πE != nil {
						continue
					}
					πF.PushCheckpoint(8)
					πTemp005 = false
				Label7:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp005 {
						πF.PopCheckpoint()
						goto Label9
					}
					if πTemp004, πE = πg.Next(πF, πTemp003); πE != nil {
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
						if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp007}, πg.TieTarget{Target: &πTemp008}}}, πTemp004); πE != nil {
							continue
						}
						µi = πTemp007
						µx = πTemp008
					}
					if πE != nil || !πTemp006 {
						continue
					}
					πF.PushCheckpoint(7)            
					// line 113: dig.append( (x + ss[i]) & 0xffffffff )
					πF.SetLineno(113)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					πTemp008 = µi
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					if πTemp009, πE = πg.GetItem(πF, µss, πTemp008); πE != nil {
						continue
					}
					if πTemp007, πE = πg.Add(πF, µx, πTemp009); πE != nil {
						continue
					}
					if πTemp004, πE = πg.And(πF, πTemp007, πg.NewInt(4294967295).ToObject()); πE != nil {
						continue
					}
					πTemp001[0] = πTemp004
					if πE = πg.CheckLocal(πF, µdig, "dig"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µdig, ßappend, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					continue
				Label8:
					if πE != nil || πR != nil {
						continue
					}
				Label9:
					// line 114: sha_info['digest'] = dig
					πF.SetLineno(114)
					if πE = πg.CheckLocal(πF, µdig, "dig"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, µdig); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsha_info, "sha_info"); πE != nil {
						continue
					}
					πTemp004 = ßdigest.ToObject()
					if πE = πg.SetItem(πF, µsha_info, πTemp004, πTemp003); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßsha_transform.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 116: def sha_init():
			πF.SetLineno(116)
			πTemp003 = make([]πg.Param, 0)
			πTemp005 = πg.NewFunction(πg.NewCode("sha_init", "/usr/lib/python2.7/_sha256.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µsha_info *πg.Object = πg.UnboundLocal; _ = µsha_info
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 []*πg.Object
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
					// line 117: sha_info = new_shaobject()
					πF.SetLineno(117)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßnew_shaobject); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					µsha_info = πTemp002
					// line 118: sha_info['digest'] = [0x6A09E667, 0xBB67AE85, 0x3C6EF372, 0xA54FF53A, 0x510E527F, 0x9B05688C, 0x1F83D9AB, 0x5BE0CD19]
					πF.SetLineno(118)
					πTemp003 = make([]*πg.Object, 8)
					πTemp003[0] = πg.NewInt(1779033703).ToObject()
					πTemp003[1] = πg.NewInt(3144134277).ToObject()
					πTemp003[2] = πg.NewInt(1013904242).ToObject()
					πTemp003[3] = πg.NewInt(2773480762).ToObject()
					πTemp003[4] = πg.NewInt(1359893119).ToObject()
					πTemp003[5] = πg.NewInt(2600822924).ToObject()
					πTemp003[6] = πg.NewInt(528734635).ToObject()
					πTemp003[7] = πg.NewInt(1541459225).ToObject()
					πTemp001 = πg.NewList(πTemp003...).ToObject()
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsha_info, "sha_info"); πE != nil {
						continue
					}
					πTemp004 = ßdigest.ToObject()
					if πE = πg.SetItem(πF, µsha_info, πTemp004, πTemp002); πE != nil {
						continue
					}
					// line 119: sha_info['count_lo'] = 0
					πF.SetLineno(119)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsha_info, "sha_info"); πE != nil {
						continue
					}
					πTemp002 = ßcount_lo.ToObject()
					if πE = πg.SetItem(πF, µsha_info, πTemp002, πTemp001); πE != nil {
						continue
					}
					// line 120: sha_info['count_hi'] = 0
					πF.SetLineno(120)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsha_info, "sha_info"); πE != nil {
						continue
					}
					πTemp002 = ßcount_hi.ToObject()
					if πE = πg.SetItem(πF, µsha_info, πTemp002, πTemp001); πE != nil {
						continue
					}
					// line 121: sha_info['local'] = 0
					πF.SetLineno(121)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsha_info, "sha_info"); πE != nil {
						continue
					}
					πTemp002 = ßlocal.ToObject()
					if πE = πg.SetItem(πF, µsha_info, πTemp002, πTemp001); πE != nil {
						continue
					}
					// line 122: sha_info['digestsize'] = 32
					πF.SetLineno(122)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(32).ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsha_info, "sha_info"); πE != nil {
						continue
					}
					πTemp002 = ßdigestsize.ToObject()
					if πE = πg.SetItem(πF, µsha_info, πTemp002, πTemp001); πE != nil {
						continue
					}
					// line 123: return sha_info
					πF.SetLineno(123)
					if πE = πg.CheckLocal(πF, µsha_info, "sha_info"); πE != nil {
						continue
					}
					πR = µsha_info
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßsha_init.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 125: def sha224_init():
			πF.SetLineno(125)
			πTemp003 = make([]πg.Param, 0)
			πTemp006 = πg.NewFunction(πg.NewCode("sha224_init", "/usr/lib/python2.7/_sha256.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µsha_info *πg.Object = πg.UnboundLocal; _ = µsha_info
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 []*πg.Object
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
					// line 126: sha_info = new_shaobject()
					πF.SetLineno(126)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßnew_shaobject); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					µsha_info = πTemp002
					// line 127: sha_info['digest'] = [0xc1059ed8, 0x367cd507, 0x3070dd17, 0xf70e5939, 0xffc00b31, 0x68581511, 0x64f98fa7, 0xbefa4fa4]
					πF.SetLineno(127)
					πTemp003 = make([]*πg.Object, 8)
					πTemp003[0] = πg.NewInt(3238371032).ToObject()
					πTemp003[1] = πg.NewInt(914150663).ToObject()
					πTemp003[2] = πg.NewInt(812702999).ToObject()
					πTemp003[3] = πg.NewInt(4144912697).ToObject()
					πTemp003[4] = πg.NewInt(4290775857).ToObject()
					πTemp003[5] = πg.NewInt(1750603025).ToObject()
					πTemp003[6] = πg.NewInt(1694076839).ToObject()
					πTemp003[7] = πg.NewInt(3204075428).ToObject()
					πTemp001 = πg.NewList(πTemp003...).ToObject()
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsha_info, "sha_info"); πE != nil {
						continue
					}
					πTemp004 = ßdigest.ToObject()
					if πE = πg.SetItem(πF, µsha_info, πTemp004, πTemp002); πE != nil {
						continue
					}
					// line 128: sha_info['count_lo'] = 0
					πF.SetLineno(128)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsha_info, "sha_info"); πE != nil {
						continue
					}
					πTemp002 = ßcount_lo.ToObject()
					if πE = πg.SetItem(πF, µsha_info, πTemp002, πTemp001); πE != nil {
						continue
					}
					// line 129: sha_info['count_hi'] = 0
					πF.SetLineno(129)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsha_info, "sha_info"); πE != nil {
						continue
					}
					πTemp002 = ßcount_hi.ToObject()
					if πE = πg.SetItem(πF, µsha_info, πTemp002, πTemp001); πE != nil {
						continue
					}
					// line 130: sha_info['local'] = 0
					πF.SetLineno(130)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsha_info, "sha_info"); πE != nil {
						continue
					}
					πTemp002 = ßlocal.ToObject()
					if πE = πg.SetItem(πF, µsha_info, πTemp002, πTemp001); πE != nil {
						continue
					}
					// line 131: sha_info['digestsize'] = 28
					πF.SetLineno(131)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(28).ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsha_info, "sha_info"); πE != nil {
						continue
					}
					πTemp002 = ßdigestsize.ToObject()
					if πE = πg.SetItem(πF, µsha_info, πTemp002, πTemp001); πE != nil {
						continue
					}
					// line 132: return sha_info
					πF.SetLineno(132)
					if πE = πg.CheckLocal(πF, µsha_info, "sha_info"); πE != nil {
						continue
					}
					πR = µsha_info
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßsha224_init.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 134: def getbuf(s):
			πF.SetLineno(134)
			πTemp003 = make([]πg.Param, 1)
			πTemp003[0] = πg.Param{Name: "s", Def: nil}
			πTemp007 = πg.NewFunction(πg.NewCode("getbuf", "/usr/lib/python2.7/_sha256.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µs *πg.Object = πArgs[0]; _ = µs
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
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp001[0] = µs
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
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp001[0] = µs
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
					if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label2
					}
					goto Label3
					// line 135: if isinstance(s, str):
					πF.SetLineno(135)
				Label1:
					// line 136: return s
					πF.SetLineno(136)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πR = µs
					continue
					goto Label4
					// line 137: elif isinstance(s, unicode):
					πF.SetLineno(137)
				Label2:
					// line 138: return str(s)
					πF.SetLineno(138)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp001[0] = µs
					if πTemp002, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πR = πTemp003
					continue
					goto Label4
				Label3:
					// line 140: return buffer(s)
					πF.SetLineno(140)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp001[0] = µs
					if πTemp002, πE = πg.ResolveGlobal(πF, ßbuffer); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πR = πTemp003
					continue
					goto Label4
				Label4:
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßgetbuf.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 142: def sha_update(sha_info, buffer):
			πF.SetLineno(142)
			πTemp003 = make([]πg.Param, 2)
			πTemp003[0] = πg.Param{Name: "sha_info", Def: nil}
			πTemp003[1] = πg.Param{Name: "buffer", Def: nil}
			πTemp008 = πg.NewFunction(πg.NewCode("sha_update", "/usr/lib/python2.7/_sha256.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µsha_info *πg.Object = πArgs[0]; _ = µsha_info
				var µbuffer *πg.Object = πArgs[1]; _ = µbuffer
				var µcount *πg.Object = πg.UnboundLocal; _ = µcount
				var µbuffer_idx *πg.Object = πg.UnboundLocal; _ = µbuffer_idx
				var µclo *πg.Object = πg.UnboundLocal; _ = µclo
				var µi *πg.Object = πg.UnboundLocal; _ = µi
				var µx *πg.Object = πg.UnboundLocal; _ = µx
				var µpos *πg.Object = πg.UnboundLocal; _ = µpos
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
				var πTemp013 []πg.Param
				_ = πTemp013
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 8: goto Label8
					case 13: goto Label13
					case 14: goto Label14
					case 7: goto Label7
					default: panic("unexpected function state")
					}
					// line 143: count = len(buffer)
					πF.SetLineno(143)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µbuffer, "buffer"); πE != nil {
						continue
					}
					πTemp001[0] = µbuffer
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µcount = πTemp003
					// line 144: buffer_idx = 0
					πF.SetLineno(144)
					µbuffer_idx = πg.NewInt(0).ToObject()
					// line 145: clo = (sha_info['count_lo'] + (count << 3)) & 0xffffffff
					πF.SetLineno(145)
					πTemp004 = ßcount_lo.ToObject()
					if πE = πg.CheckLocal(πF, µsha_info, "sha_info"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µsha_info, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcount, "count"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.LShift(πF, µcount, πg.NewInt(3).ToObject()); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Add(πF, πTemp005, πTemp004); πE != nil {
						continue
					}
					if πTemp002, πE = πg.And(πF, πTemp003, πg.NewInt(4294967295).ToObject()); πE != nil {
						continue
					}
					µclo = πTemp002
					if πE = πg.CheckLocal(πF, µclo, "clo"); πE != nil {
						continue
					}
					πTemp003 = ßcount_lo.ToObject()
					if πE = πg.CheckLocal(πF, µsha_info, "sha_info"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µsha_info, πTemp003); πE != nil {
						continue
					}
					if πTemp002, πE = πg.LT(πF, µclo, πTemp004); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label1
					}
					goto Label2
					// line 146: if clo < sha_info['count_lo']:
					πF.SetLineno(146)
				Label1:
					// line 147: sha_info['count_hi'] += 1
					πF.SetLineno(147)
					πTemp002 = ßcount_hi.ToObject()
					if πE = πg.CheckLocal(πF, µsha_info, "sha_info"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µsha_info, πTemp002); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IAdd(πF, πTemp003, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsha_info, "sha_info"); πE != nil {
						continue
					}
					πTemp004 = ßcount_hi.ToObject()
					if πE = πg.SetItem(πF, µsha_info, πTemp004, πTemp002); πE != nil {
						continue
					}
					goto Label2
				Label2:
					// line 148: sha_info['count_lo'] = clo
					πF.SetLineno(148)
					if πE = πg.CheckLocal(πF, µclo, "clo"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µclo); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsha_info, "sha_info"); πE != nil {
						continue
					}
					πTemp003 = ßcount_lo.ToObject()
					if πE = πg.SetItem(πF, µsha_info, πTemp003, πTemp002); πE != nil {
						continue
					}
					// line 150: sha_info['count_hi'] += (count >> 29)
					πF.SetLineno(150)
					πTemp002 = ßcount_hi.ToObject()
					if πE = πg.CheckLocal(πF, µsha_info, "sha_info"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µsha_info, πTemp002); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcount, "count"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.RShift(πF, µcount, πg.NewInt(29).ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IAdd(πF, πTemp003, πTemp002); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsha_info, "sha_info"); πE != nil {
						continue
					}
					πTemp005 = ßcount_hi.ToObject()
					if πE = πg.SetItem(πF, µsha_info, πTemp005, πTemp004); πE != nil {
						continue
					}
					πTemp002 = ßlocal.ToObject()
					if πE = πg.CheckLocal(πF, µsha_info, "sha_info"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µsha_info, πTemp002); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label3
					}
					goto Label4
					// line 152: if sha_info['local']:
					πF.SetLineno(152)
				Label3:
					// line 153: i = SHA_BLOCKSIZE - sha_info['local']
					πF.SetLineno(153)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßSHA_BLOCKSIZE); πE != nil {
						continue
					}
					πTemp004 = ßlocal.ToObject()
					if πE = πg.CheckLocal(πF, µsha_info, "sha_info"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µsha_info, πTemp004); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Sub(πF, πTemp003, πTemp005); πE != nil {
						continue
					}
					µi = πTemp002
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcount, "count"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GT(πF, µi, µcount); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label5
					}
					goto Label6
					// line 154: if i > count:
					πF.SetLineno(154)
				Label5:
					// line 155: i = count
					πF.SetLineno(155)
					if πE = πg.CheckLocal(πF, µcount, "count"); πE != nil {
						continue
					}
					µi = µcount
					goto Label6
				Label6:
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µbuffer_idx, "buffer_idx"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µbuffer_idx, "buffer_idx"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Add(πF, µbuffer_idx, µi); πE != nil {
						continue
					}
					if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{µbuffer_idx, πTemp004, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µbuffer, "buffer"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µbuffer, πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp004
					if πTemp003, πE = πg.ResolveGlobal(πF, ßenumerate); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp002, πE = πg.Iter(πF, πTemp004); πE != nil {
						continue
					}
					πF.PushCheckpoint(8)
					πTemp006 = false
				Label7:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp006 {
						πF.PopCheckpoint()
						goto Label9
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
						µx = πTemp003
					}
					if πE != nil || !πTemp007 {
						continue
					}
					πF.PushCheckpoint(7)            
					// line 159: sha_info['data'][sha_info['local']+x[0]] = struct.unpack('B', x[1])[0]
					πF.SetLineno(159)
					πTemp003 = πg.NewInt(0).ToObject()
					πTemp001 = πF.MakeArgs(2)
					πTemp001[0] = ßB.ToObject()
					πTemp005 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetItem(πF, µx, πTemp005); πE != nil {
						continue
					}
					πTemp001[1] = πTemp008
					if πTemp005, πE = πg.ResolveGlobal(πF, ßstruct); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetAttr(πF, πTemp005, ßunpack, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp008.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp004, πE = πg.GetItem(πF, πTemp005, πTemp003); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp004); πE != nil {
						continue
					}
					πTemp005 = ßdata.ToObject()
					if πE = πg.CheckLocal(πF, µsha_info, "sha_info"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetItem(πF, µsha_info, πTemp005); πE != nil {
						continue
					}
					πTemp010 = ßlocal.ToObject()
					if πE = πg.CheckLocal(πF, µsha_info, "sha_info"); πE != nil {
						continue
					}
					if πTemp011, πE = πg.GetItem(πF, µsha_info, πTemp010); πE != nil {
						continue
					}
					πTemp010 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					if πTemp012, πE = πg.GetItem(πF, µx, πTemp010); πE != nil {
						continue
					}
					if πTemp009, πE = πg.Add(πF, πTemp011, πTemp012); πE != nil {
						continue
					}
					πTemp005 = πTemp009
					if πE = πg.SetItem(πF, πTemp008, πTemp005, πTemp003); πE != nil {
						continue
					}
					continue
				Label8:
					if πE != nil || πR != nil {
						continue
					}
				Label9:
					// line 161: count -= i
					πF.SetLineno(161)
					if πE = πg.CheckLocal(πF, µcount, "count"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ISub(πF, µcount, µi); πE != nil {
						continue
					}
					µcount = πTemp002
					// line 162: buffer_idx += i
					πF.SetLineno(162)
					if πE = πg.CheckLocal(πF, µbuffer_idx, "buffer_idx"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IAdd(πF, µbuffer_idx, µi); πE != nil {
						continue
					}
					µbuffer_idx = πTemp002
					// line 164: sha_info['local'] += i
					πF.SetLineno(164)
					πTemp002 = ßlocal.ToObject()
					if πE = πg.CheckLocal(πF, µsha_info, "sha_info"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µsha_info, πTemp002); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IAdd(πF, πTemp003, µi); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsha_info, "sha_info"); πE != nil {
						continue
					}
					πTemp004 = ßlocal.ToObject()
					if πE = πg.SetItem(πF, µsha_info, πTemp004, πTemp002); πE != nil {
						continue
					}
					πTemp003 = ßlocal.ToObject()
					if πE = πg.CheckLocal(πF, µsha_info, "sha_info"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µsha_info, πTemp003); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßSHA_BLOCKSIZE); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, πTemp004, πTemp003); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label10
					}
					goto Label11
					// line 165: if sha_info['local'] == SHA_BLOCKSIZE:
					πF.SetLineno(165)
				Label10:
					// line 166: sha_transform(sha_info)
					πF.SetLineno(166)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µsha_info, "sha_info"); πE != nil {
						continue
					}
					πTemp001[0] = µsha_info
					if πTemp002, πE = πg.ResolveGlobal(πF, ßsha_transform); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 167: sha_info['local'] = 0
					πF.SetLineno(167)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsha_info, "sha_info"); πE != nil {
						continue
					}
					πTemp003 = ßlocal.ToObject()
					if πE = πg.SetItem(πF, µsha_info, πTemp003, πTemp002); πE != nil {
						continue
					}
					goto Label12
				Label11:
					// line 169: return
					πF.SetLineno(169)
					πR = πg.None
					continue
					goto Label12
				Label12:
					goto Label4
				Label4:
					// line 171: while count >= SHA_BLOCKSIZE:
					πF.SetLineno(171)
					πF.PushCheckpoint(14)
					πTemp006 = false
				Label13:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp006 {
						πF.PopCheckpoint()
						goto Label15
					}
					if πE = πg.CheckLocal(πF, µcount, "count"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßSHA_BLOCKSIZE); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GE(πF, µcount, πTemp003); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πE != nil || !πTemp007 {
						continue
					}
					πF.PushCheckpoint(13)            
					// line 173: sha_info['data'] = [struct.unpack('B',c)[0] for c in buffer[buffer_idx:buffer_idx + SHA_BLOCKSIZE]]
					πF.SetLineno(173)
					πTemp013 = make([]πg.Param, 0)
					πTemp003 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/_sha256.py", πTemp013, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µc *πg.Object = πg.UnboundLocal; _ = µc
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
						var πTemp008 *πg.Object
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
								default: panic("unexpected function state")
								}
								if πE = πg.CheckLocal(πF, µbuffer_idx, "buffer_idx"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µbuffer_idx, "buffer_idx"); πE != nil {
									continue
								}
								if πTemp004, πE = πg.ResolveGlobal(πF, ßSHA_BLOCKSIZE); πE != nil {
									continue
								}
								if πTemp003, πE = πg.Add(πF, µbuffer_idx, πTemp004); πE != nil {
									continue
								}
								if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{µbuffer_idx, πTemp003, πg.None}, nil); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µbuffer, "buffer"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.GetItem(πF, µbuffer, πTemp002); πE != nil {
									continue
								}
								if πTemp001, πE = πg.Iter(πF, πTemp003); πE != nil {
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
									µc = πTemp002
								}
								if πE != nil || !πTemp006 {
									continue
								}
								πF.PushCheckpoint(1)            
								// line 173: sha_info['data'] = [struct.unpack('B',c)[0] for c in buffer[buffer_idx:buffer_idx + SHA_BLOCKSIZE]]
								πF.SetLineno(173)
								πTemp002 = πg.NewInt(0).ToObject()
								πTemp007 = πF.MakeArgs(2)
								πTemp007[0] = ßB.ToObject()
								if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
									continue
								}
								πTemp007[1] = µc
								if πTemp004, πE = πg.ResolveGlobal(πF, ßstruct); πE != nil {
									continue
								}
								if πTemp008, πE = πg.GetAttr(πF, πTemp004, ßunpack, nil); πE != nil {
									continue
								}
								if πTemp004, πE = πTemp008.Call(πF, πTemp007, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp007)
								if πTemp003, πE = πg.GetItem(πF, πTemp004, πTemp002); πE != nil {
									continue
								}
								πF.PushCheckpoint(4)
								return πTemp003, nil
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
					if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ListType.Call(πF, πg.Args{πTemp004}, nil); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp002); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsha_info, "sha_info"); πE != nil {
						continue
					}
					πTemp005 = ßdata.ToObject()
					if πE = πg.SetItem(πF, µsha_info, πTemp005, πTemp004); πE != nil {
						continue
					}
					// line 174: count -= SHA_BLOCKSIZE
					πF.SetLineno(174)
					if πE = πg.CheckLocal(πF, µcount, "count"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßSHA_BLOCKSIZE); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ISub(πF, µcount, πTemp002); πE != nil {
						continue
					}
					µcount = πTemp004
					// line 175: buffer_idx += SHA_BLOCKSIZE
					πF.SetLineno(175)
					if πE = πg.CheckLocal(πF, µbuffer_idx, "buffer_idx"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßSHA_BLOCKSIZE); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IAdd(πF, µbuffer_idx, πTemp002); πE != nil {
						continue
					}
					µbuffer_idx = πTemp004
					// line 176: sha_transform(sha_info)
					πF.SetLineno(176)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µsha_info, "sha_info"); πE != nil {
						continue
					}
					πTemp001[0] = µsha_info
					if πTemp002, πE = πg.ResolveGlobal(πF, ßsha_transform); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					continue
				Label14:
					if πE != nil || πR != nil {
						continue
					}
				Label15:
					// line 180: pos = sha_info['local']
					πF.SetLineno(180)
					πTemp002 = ßlocal.ToObject()
					if πE = πg.CheckLocal(πF, µsha_info, "sha_info"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µsha_info, πTemp002); πE != nil {
						continue
					}
					µpos = πTemp004
					// line 181: sha_info['data'][pos:pos+count] = [struct.unpack('B',c)[0] for c in buffer[buffer_idx:buffer_idx + count]]
					πF.SetLineno(181)
					πTemp013 = make([]πg.Param, 0)
					πTemp004 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/_sha256.py", πTemp013, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µc *πg.Object = πg.UnboundLocal; _ = µc
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
						return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
							for ; πF.State() >= 0; πF.PopCheckpoint() {
								switch πF.State() {
								case 0:
								case 1: goto Label1
								case 2: goto Label2
								case 4: goto Label4
								default: panic("unexpected function state")
								}
								if πE = πg.CheckLocal(πF, µbuffer_idx, "buffer_idx"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µbuffer_idx, "buffer_idx"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µcount, "count"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.Add(πF, µbuffer_idx, µcount); πE != nil {
									continue
								}
								if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{µbuffer_idx, πTemp003, πg.None}, nil); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µbuffer, "buffer"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.GetItem(πF, µbuffer, πTemp002); πE != nil {
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
									µc = πTemp002
								}
								if πE != nil || !πTemp005 {
									continue
								}
								πF.PushCheckpoint(1)            
								// line 181: sha_info['data'][pos:pos+count] = [struct.unpack('B',c)[0] for c in buffer[buffer_idx:buffer_idx + count]]
								πF.SetLineno(181)
								πTemp002 = πg.NewInt(0).ToObject()
								πTemp006 = πF.MakeArgs(2)
								πTemp006[0] = ßB.ToObject()
								if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
									continue
								}
								πTemp006[1] = µc
								if πTemp007, πE = πg.ResolveGlobal(πF, ßstruct); πE != nil {
									continue
								}
								if πTemp008, πE = πg.GetAttr(πF, πTemp007, ßunpack, nil); πE != nil {
									continue
								}
								if πTemp007, πE = πTemp008.Call(πF, πTemp006, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp006)
								if πTemp003, πE = πg.GetItem(πF, πTemp007, πTemp002); πE != nil {
									continue
								}
								πF.PushCheckpoint(4)
								return πTemp003, nil
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
					if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ListType.Call(πF, πg.Args{πTemp005}, nil); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp002); πE != nil {
						continue
					}
					πTemp008 = ßdata.ToObject()
					if πE = πg.CheckLocal(πF, µsha_info, "sha_info"); πE != nil {
						continue
					}
					if πTemp009, πE = πg.GetItem(πF, µsha_info, πTemp008); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µpos, "pos"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µpos, "pos"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcount, "count"); πE != nil {
						continue
					}
					if πTemp010, πE = πg.Add(πF, µpos, µcount); πE != nil {
						continue
					}
					if πTemp008, πE = πg.SliceType.Call(πF, πg.Args{µpos, πTemp010, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.SetItem(πF, πTemp009, πTemp008, πTemp005); πE != nil {
						continue
					}
					// line 182: sha_info['local'] = count
					πF.SetLineno(182)
					if πE = πg.CheckLocal(πF, µcount, "count"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µcount); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsha_info, "sha_info"); πE != nil {
						continue
					}
					πTemp005 = ßlocal.ToObject()
					if πE = πg.SetItem(πF, µsha_info, πTemp005, πTemp002); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßsha_update.ToObject(), πTemp008); πE != nil {
				continue
			}
			// line 184: def sha_final(sha_info):
			πF.SetLineno(184)
			πTemp003 = make([]πg.Param, 1)
			πTemp003[0] = πg.Param{Name: "sha_info", Def: nil}
			πTemp009 = πg.NewFunction(πg.NewCode("sha_final", "/usr/lib/python2.7/_sha256.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µsha_info *πg.Object = πArgs[0]; _ = µsha_info
				var µlo_bit_count *πg.Object = πg.UnboundLocal; _ = µlo_bit_count
				var µhi_bit_count *πg.Object = πg.UnboundLocal; _ = µhi_bit_count
				var µcount *πg.Object = πg.UnboundLocal; _ = µcount
				var µdig *πg.Object = πg.UnboundLocal; _ = µdig
				var µi *πg.Object = πg.UnboundLocal; _ = µi
				var πTemp001 *πg.Object
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
				var πTemp007 []*πg.Object
				_ = πTemp007
				var πTemp008 *πg.Object
				_ = πTemp008
				var πTemp009 bool
				_ = πTemp009
				var πTemp010 []*πg.Object
				_ = πTemp010
				var πTemp011 []πg.Param
				_ = πTemp011
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 4: goto Label4
					case 5: goto Label5
					default: panic("unexpected function state")
					}
					// line 185: lo_bit_count = sha_info['count_lo']
					πF.SetLineno(185)
					πTemp001 = ßcount_lo.ToObject()
					if πE = πg.CheckLocal(πF, µsha_info, "sha_info"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetItem(πF, µsha_info, πTemp001); πE != nil {
						continue
					}
					µlo_bit_count = πTemp002
					// line 186: hi_bit_count = sha_info['count_hi']
					πF.SetLineno(186)
					πTemp001 = ßcount_hi.ToObject()
					if πE = πg.CheckLocal(πF, µsha_info, "sha_info"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetItem(πF, µsha_info, πTemp001); πE != nil {
						continue
					}
					µhi_bit_count = πTemp002
					// line 187: count = (lo_bit_count >> 3) & 0x3f
					πF.SetLineno(187)
					if πE = πg.CheckLocal(πF, µlo_bit_count, "lo_bit_count"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.RShift(πF, µlo_bit_count, πg.NewInt(3).ToObject()); πE != nil {
						continue
					}
					if πTemp001, πE = πg.And(πF, πTemp002, πg.NewInt(63).ToObject()); πE != nil {
						continue
					}
					µcount = πTemp001
					// line 188: sha_info['data'][count] = 0x80;
					πF.SetLineno(188)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(128).ToObject()); πE != nil {
						continue
					}
					πTemp002 = ßdata.ToObject()
					if πE = πg.CheckLocal(πF, µsha_info, "sha_info"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µsha_info, πTemp002); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcount, "count"); πE != nil {
						continue
					}
					πTemp002 = µcount
					if πE = πg.SetItem(πF, πTemp003, πTemp002, πTemp001); πE != nil {
						continue
					}
					// line 189: count += 1
					πF.SetLineno(189)
					if πE = πg.CheckLocal(πF, µcount, "count"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.IAdd(πF, µcount, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µcount = πTemp001
					if πE = πg.CheckLocal(πF, µcount, "count"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßSHA_BLOCKSIZE); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Sub(πF, πTemp003, πg.NewInt(8).ToObject()); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GT(πF, µcount, πTemp002); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label1
					}
					goto Label2
					// line 190: if count > SHA_BLOCKSIZE - 8:
					πF.SetLineno(190)
				Label1:
					// line 192: sha_info['data'] = sha_info['data'][:count] + ([0] * (SHA_BLOCKSIZE - count))
					πF.SetLineno(192)
					if πE = πg.CheckLocal(πF, µcount, "count"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.None, µcount, πg.None}, nil); πE != nil {
						continue
					}
					πTemp005 = ßdata.ToObject()
					if πE = πg.CheckLocal(πF, µsha_info, "sha_info"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, µsha_info, πTemp005); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, πTemp006, πTemp002); πE != nil {
						continue
					}
					πTemp007 = make([]*πg.Object, 1)
					πTemp007[0] = πg.NewInt(0).ToObject()
					πTemp005 = πg.NewList(πTemp007...).ToObject()
					if πTemp008, πE = πg.ResolveGlobal(πF, ßSHA_BLOCKSIZE); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcount, "count"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.Sub(πF, πTemp008, µcount); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mul(πF, πTemp005, πTemp006); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, πTemp003, πTemp002); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsha_info, "sha_info"); πE != nil {
						continue
					}
					πTemp003 = ßdata.ToObject()
					if πE = πg.SetItem(πF, µsha_info, πTemp003, πTemp002); πE != nil {
						continue
					}
					// line 193: sha_transform(sha_info)
					πF.SetLineno(193)
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µsha_info, "sha_info"); πE != nil {
						continue
					}
					πTemp007[0] = µsha_info
					if πTemp001, πE = πg.ResolveGlobal(πF, ßsha_transform); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					// line 195: sha_info['data'] = [0] * SHA_BLOCKSIZE
					πF.SetLineno(195)
					πTemp007 = make([]*πg.Object, 1)
					πTemp007[0] = πg.NewInt(0).ToObject()
					πTemp002 = πg.NewList(πTemp007...).ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßSHA_BLOCKSIZE); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Mul(πF, πTemp002, πTemp003); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsha_info, "sha_info"); πE != nil {
						continue
					}
					πTemp003 = ßdata.ToObject()
					if πE = πg.SetItem(πF, µsha_info, πTemp003, πTemp002); πE != nil {
						continue
					}
					goto Label3
				Label2:
					// line 197: sha_info['data'] = sha_info['data'][:count] + ([0] * (SHA_BLOCKSIZE - count))
					πF.SetLineno(197)
					if πE = πg.CheckLocal(πF, µcount, "count"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.None, µcount, πg.None}, nil); πE != nil {
						continue
					}
					πTemp005 = ßdata.ToObject()
					if πE = πg.CheckLocal(πF, µsha_info, "sha_info"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, µsha_info, πTemp005); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, πTemp006, πTemp002); πE != nil {
						continue
					}
					πTemp007 = make([]*πg.Object, 1)
					πTemp007[0] = πg.NewInt(0).ToObject()
					πTemp005 = πg.NewList(πTemp007...).ToObject()
					if πTemp008, πE = πg.ResolveGlobal(πF, ßSHA_BLOCKSIZE); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcount, "count"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.Sub(πF, πTemp008, µcount); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mul(πF, πTemp005, πTemp006); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, πTemp003, πTemp002); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsha_info, "sha_info"); πE != nil {
						continue
					}
					πTemp003 = ßdata.ToObject()
					if πE = πg.SetItem(πF, µsha_info, πTemp003, πTemp002); πE != nil {
						continue
					}
					goto Label3
				Label3:
					// line 199: sha_info['data'][56] = (hi_bit_count >> 24) & 0xff
					πF.SetLineno(199)
					if πE = πg.CheckLocal(πF, µhi_bit_count, "hi_bit_count"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.RShift(πF, µhi_bit_count, πg.NewInt(24).ToObject()); πE != nil {
						continue
					}
					if πTemp001, πE = πg.And(πF, πTemp002, πg.NewInt(255).ToObject()); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
						continue
					}
					πTemp003 = ßdata.ToObject()
					if πE = πg.CheckLocal(πF, µsha_info, "sha_info"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µsha_info, πTemp003); πE != nil {
						continue
					}
					πTemp003 = πg.NewInt(56).ToObject()
					if πE = πg.SetItem(πF, πTemp005, πTemp003, πTemp002); πE != nil {
						continue
					}
					// line 200: sha_info['data'][57] = (hi_bit_count >> 16) & 0xff
					πF.SetLineno(200)
					if πE = πg.CheckLocal(πF, µhi_bit_count, "hi_bit_count"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.RShift(πF, µhi_bit_count, πg.NewInt(16).ToObject()); πE != nil {
						continue
					}
					if πTemp001, πE = πg.And(πF, πTemp002, πg.NewInt(255).ToObject()); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
						continue
					}
					πTemp003 = ßdata.ToObject()
					if πE = πg.CheckLocal(πF, µsha_info, "sha_info"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µsha_info, πTemp003); πE != nil {
						continue
					}
					πTemp003 = πg.NewInt(57).ToObject()
					if πE = πg.SetItem(πF, πTemp005, πTemp003, πTemp002); πE != nil {
						continue
					}
					// line 201: sha_info['data'][58] = (hi_bit_count >>  8) & 0xff
					πF.SetLineno(201)
					if πE = πg.CheckLocal(πF, µhi_bit_count, "hi_bit_count"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.RShift(πF, µhi_bit_count, πg.NewInt(8).ToObject()); πE != nil {
						continue
					}
					if πTemp001, πE = πg.And(πF, πTemp002, πg.NewInt(255).ToObject()); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
						continue
					}
					πTemp003 = ßdata.ToObject()
					if πE = πg.CheckLocal(πF, µsha_info, "sha_info"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µsha_info, πTemp003); πE != nil {
						continue
					}
					πTemp003 = πg.NewInt(58).ToObject()
					if πE = πg.SetItem(πF, πTemp005, πTemp003, πTemp002); πE != nil {
						continue
					}
					// line 202: sha_info['data'][59] = (hi_bit_count >>  0) & 0xff
					πF.SetLineno(202)
					if πE = πg.CheckLocal(πF, µhi_bit_count, "hi_bit_count"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.RShift(πF, µhi_bit_count, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					if πTemp001, πE = πg.And(πF, πTemp002, πg.NewInt(255).ToObject()); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
						continue
					}
					πTemp003 = ßdata.ToObject()
					if πE = πg.CheckLocal(πF, µsha_info, "sha_info"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µsha_info, πTemp003); πE != nil {
						continue
					}
					πTemp003 = πg.NewInt(59).ToObject()
					if πE = πg.SetItem(πF, πTemp005, πTemp003, πTemp002); πE != nil {
						continue
					}
					// line 203: sha_info['data'][60] = (lo_bit_count >> 24) & 0xff
					πF.SetLineno(203)
					if πE = πg.CheckLocal(πF, µlo_bit_count, "lo_bit_count"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.RShift(πF, µlo_bit_count, πg.NewInt(24).ToObject()); πE != nil {
						continue
					}
					if πTemp001, πE = πg.And(πF, πTemp002, πg.NewInt(255).ToObject()); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
						continue
					}
					πTemp003 = ßdata.ToObject()
					if πE = πg.CheckLocal(πF, µsha_info, "sha_info"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µsha_info, πTemp003); πE != nil {
						continue
					}
					πTemp003 = πg.NewInt(60).ToObject()
					if πE = πg.SetItem(πF, πTemp005, πTemp003, πTemp002); πE != nil {
						continue
					}
					// line 204: sha_info['data'][61] = (lo_bit_count >> 16) & 0xff
					πF.SetLineno(204)
					if πE = πg.CheckLocal(πF, µlo_bit_count, "lo_bit_count"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.RShift(πF, µlo_bit_count, πg.NewInt(16).ToObject()); πE != nil {
						continue
					}
					if πTemp001, πE = πg.And(πF, πTemp002, πg.NewInt(255).ToObject()); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
						continue
					}
					πTemp003 = ßdata.ToObject()
					if πE = πg.CheckLocal(πF, µsha_info, "sha_info"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µsha_info, πTemp003); πE != nil {
						continue
					}
					πTemp003 = πg.NewInt(61).ToObject()
					if πE = πg.SetItem(πF, πTemp005, πTemp003, πTemp002); πE != nil {
						continue
					}
					// line 205: sha_info['data'][62] = (lo_bit_count >>  8) & 0xff
					πF.SetLineno(205)
					if πE = πg.CheckLocal(πF, µlo_bit_count, "lo_bit_count"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.RShift(πF, µlo_bit_count, πg.NewInt(8).ToObject()); πE != nil {
						continue
					}
					if πTemp001, πE = πg.And(πF, πTemp002, πg.NewInt(255).ToObject()); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
						continue
					}
					πTemp003 = ßdata.ToObject()
					if πE = πg.CheckLocal(πF, µsha_info, "sha_info"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µsha_info, πTemp003); πE != nil {
						continue
					}
					πTemp003 = πg.NewInt(62).ToObject()
					if πE = πg.SetItem(πF, πTemp005, πTemp003, πTemp002); πE != nil {
						continue
					}
					// line 206: sha_info['data'][63] = (lo_bit_count >>  0) & 0xff
					πF.SetLineno(206)
					if πE = πg.CheckLocal(πF, µlo_bit_count, "lo_bit_count"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.RShift(πF, µlo_bit_count, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					if πTemp001, πE = πg.And(πF, πTemp002, πg.NewInt(255).ToObject()); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
						continue
					}
					πTemp003 = ßdata.ToObject()
					if πE = πg.CheckLocal(πF, µsha_info, "sha_info"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µsha_info, πTemp003); πE != nil {
						continue
					}
					πTemp003 = πg.NewInt(63).ToObject()
					if πE = πg.SetItem(πF, πTemp005, πTemp003, πTemp002); πE != nil {
						continue
					}
					// line 208: sha_transform(sha_info)
					πF.SetLineno(208)
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µsha_info, "sha_info"); πE != nil {
						continue
					}
					πTemp007[0] = µsha_info
					if πTemp001, πE = πg.ResolveGlobal(πF, ßsha_transform); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					// line 210: dig = []
					πF.SetLineno(210)
					πTemp007 = make([]*πg.Object, 0)
					πTemp001 = πg.NewList(πTemp007...).ToObject()
					µdig = πTemp001
					πTemp002 = ßdigest.ToObject()
					if πE = πg.CheckLocal(πF, µsha_info, "sha_info"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µsha_info, πTemp002); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Iter(πF, πTemp003); πE != nil {
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
						µi = πTemp002
					}
					if πE != nil || !πTemp009 {
						continue
					}
					πF.PushCheckpoint(4)            
					// line 212: dig.extend([ ((i>>24) & 0xff), ((i>>16) & 0xff), ((i>>8) & 0xff), (i & 0xff) ])
					πF.SetLineno(212)
					πTemp007 = πF.MakeArgs(1)
					πTemp010 = make([]*πg.Object, 4)
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.RShift(πF, µi, πg.NewInt(24).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.And(πF, πTemp003, πg.NewInt(255).ToObject()); πE != nil {
						continue
					}
					πTemp010[0] = πTemp002
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.RShift(πF, µi, πg.NewInt(16).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.And(πF, πTemp003, πg.NewInt(255).ToObject()); πE != nil {
						continue
					}
					πTemp010[1] = πTemp002
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.RShift(πF, µi, πg.NewInt(8).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.And(πF, πTemp003, πg.NewInt(255).ToObject()); πE != nil {
						continue
					}
					πTemp010[2] = πTemp002
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.And(πF, µi, πg.NewInt(255).ToObject()); πE != nil {
						continue
					}
					πTemp010[3] = πTemp002
					πTemp002 = πg.NewList(πTemp010...).ToObject()
					πTemp007[0] = πTemp002
					if πE = πg.CheckLocal(πF, µdig, "dig"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µdig, ßextend, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					continue
				Label5:
					if πE != nil || πR != nil {
						continue
					}
				Label6:
					// line 213: return ''.join([chr(i) for i in dig])
					πF.SetLineno(213)
					πTemp007 = πF.MakeArgs(1)
					πTemp011 = make([]πg.Param, 0)
					πTemp002 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/_sha256.py", πTemp011, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µi *πg.Object = πg.UnboundLocal; _ = µi
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
								if πE = πg.CheckLocal(πF, µdig, "dig"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.Iter(πF, µdig); πE != nil {
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
									µi = πTemp004
								}
								if πE != nil || !πTemp003 {
									continue
								}
								πF.PushCheckpoint(1)            
								// line 213: return ''.join([chr(i) for i in dig])
								πF.SetLineno(213)
								πTemp005 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
									continue
								}
								πTemp005[0] = µi
								if πTemp004, πE = πg.ResolveGlobal(πF, ßchr); πE != nil {
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
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.ListType.Call(πF, πg.Args{πTemp003}, nil); πE != nil {
						continue
					}
					πTemp007[0] = πTemp001
					if πTemp001, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
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
			if πE = πF.Globals().SetItem(πF, ßsha_final.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 215: class sha256(object):
			πF.SetLineno(215)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp013, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp013
			πTemp010 = πg.NewDict()
			if πTemp011, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp010.SetItem(πF, ß__module__.ToObject(), πTemp011); πE != nil {
				continue
			}
			_, πE = πg.NewCode("sha256", "/usr/lib/python2.7/_sha256.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 216: digest_size = digestsize = SHA_DIGESTSIZE
					πF.SetLineno(216)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßSHA_DIGESTSIZE); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßdigest_size.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßdigestsize.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 217: block_size = SHA_BLOCKSIZE
					πF.SetLineno(217)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßSHA_BLOCKSIZE); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßblock_size.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 219: def __init__(self, s=None):
					πF.SetLineno(219)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[1] = πg.Param{Name: "s", Def: πTemp003}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/_sha256.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µs *πg.Object = πArgs[1]; _ = µs
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 bool
						_ = πTemp003
						var πTemp004 []*πg.Object
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
							// line 220: self._sha = sha_init()
							πF.SetLineno(220)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßsha_init); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_sha, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, µs); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 221: if s:
							πF.SetLineno(221)
						Label1:
							// line 222: sha_update(self._sha, getbuf(s))
							πF.SetLineno(222)
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_sha, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp005[0] = µs
							if πTemp001, πE = πg.ResolveGlobal(πF, ßgetbuf); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp004[1] = πTemp002
							if πTemp001, πE = πg.ResolveGlobal(πF, ßsha_update); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
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
					if πE = πClass.SetItem(πF, ß__init__.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 224: def update(self, s):
					πF.SetLineno(224)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "s", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("update", "/usr/lib/python2.7/_sha256.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µs *πg.Object = πArgs[1]; _ = µs
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
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
							// line 225: sha_update(self._sha, getbuf(s))
							πF.SetLineno(225)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_sha, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp003[0] = µs
							if πTemp002, πE = πg.ResolveGlobal(πF, ßgetbuf); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πTemp001[1] = πTemp004
							if πTemp002, πE = πg.ResolveGlobal(πF, ßsha_update); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßupdate.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 227: def digest(self):
					πF.SetLineno(227)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("digest", "/usr/lib/python2.7/_sha256.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
							// line 228: return sha_final(self._sha.copy())[:self._sha['digestsize']]
							πF.SetLineno(228)
							πTemp002 = ßdigestsize.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß_sha, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp004, πTemp002); πE != nil {
								continue
							}
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πTemp003, πg.None}, nil); πE != nil {
								continue
							}
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_sha, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßcopy, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp005[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßsha_final); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp002, πE = πg.GetItem(πF, πTemp004, πTemp001); πE != nil {
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
					if πE = πClass.SetItem(πF, ßdigest.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 230: def hexdigest(self):
					πF.SetLineno(230)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("hexdigest", "/usr/lib/python2.7/_sha256.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 []πg.Param
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
							// line 231: return ''.join([('0%x' % ord(i))[-2:] for i in self.digest()])
							πF.SetLineno(231)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = make([]πg.Param, 0)
							πTemp003 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/_sha256.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µi *πg.Object = πg.UnboundLocal; _ = µi
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
								var πTemp007 []*πg.Object
								_ = πTemp007
								var πTemp008 *πg.Object
								_ = πTemp008
								var πTemp009 *πg.Object
								_ = πTemp009
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
										if πTemp002, πE = πg.GetAttr(πF, µself, ßdigest, nil); πE != nil {
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
											µi = πTemp002
										}
										if πE != nil || !πTemp005 {
											continue
										}
										πF.PushCheckpoint(1)            
										// line 231: return ''.join([('0%x' % ord(i))[-2:] for i in self.digest()])
										πF.SetLineno(231)
										if πTemp003, πE = πg.Neg(πF, πg.NewInt(2).ToObject()); πE != nil {
											continue
										}
										if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πTemp003, πg.None, πg.None}, nil); πE != nil {
											continue
										}
										πTemp007 = πF.MakeArgs(1)
										if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
											continue
										}
										πTemp007[0] = µi
										if πTemp008, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
											continue
										}
										if πTemp009, πE = πTemp008.Call(πF, πTemp007, nil); πE != nil {
											continue
										}
										πF.FreeArgs(πTemp007)
										if πTemp006, πE = πg.Mod(πF, πg.NewStr("0%x").ToObject(), πTemp009); πE != nil {
											continue
										}
										if πTemp003, πE = πg.GetItem(πF, πTemp006, πTemp002); πE != nil {
											continue
										}
										πF.PushCheckpoint(4)
										return πTemp003, nil
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
							if πTemp005, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ListType.Call(πF, πg.Args{πTemp005}, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
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
					if πE = πClass.SetItem(πF, ßhexdigest.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 233: def copy(self):
					πF.SetLineno(233)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("copy", "/usr/lib/python2.7/_sha256.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µnew *πg.Object = πg.UnboundLocal; _ = µnew
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
							// line 234: new = sha256.__new__(sha256)
							πF.SetLineno(234)
							πTemp001 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßsha256); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßsha256); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ß__new__, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µnew = πTemp002
							// line 235: new._sha = self._sha.copy()
							πF.SetLineno(235)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_sha, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßcopy, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnew, "new"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µnew, ß_sha, πTemp003); πE != nil {
								continue
							}
							// line 236: return new
							πF.SetLineno(236)
							if πE = πg.CheckLocal(πF, µnew, "new"); πE != nil {
								continue
							}
							πR = µnew
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßcopy.ToObject(), πTemp006); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp012, πE = πTemp010.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp012 == nil {
				πTemp012 = πg.TypeType.ToObject()
			}
			if πTemp013, πE = πTemp012.Call(πF, []*πg.Object{πg.NewStr("sha256").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp010.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßsha256.ToObject(), πTemp013); πE != nil {
				continue
			}
			// line 238: class sha224(sha256):
			πF.SetLineno(238)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp013, πE = πg.ResolveGlobal(πF, ßsha256); πE != nil {
				continue
			}
			πTemp002[0] = πTemp013
			πTemp010 = πg.NewDict()
			if πTemp011, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp010.SetItem(πF, ß__module__.ToObject(), πTemp011); πE != nil {
				continue
			}
			_, πE = πg.NewCode("sha224", "/usr/lib/python2.7/_sha256.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 239: digest_size = digestsize = 28
					πF.SetLineno(239)
					if πE = πClass.SetItem(πF, ßdigest_size.ToObject(), πg.NewInt(28).ToObject()); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßdigestsize.ToObject(), πg.NewInt(28).ToObject()); πE != nil {
						continue
					}
					// line 241: def __init__(self, s=None):
					πF.SetLineno(241)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[1] = πg.Param{Name: "s", Def: πTemp003}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/_sha256.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µs *πg.Object = πArgs[1]; _ = µs
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 bool
						_ = πTemp003
						var πTemp004 []*πg.Object
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
							// line 242: self._sha = sha224_init()
							πF.SetLineno(242)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßsha224_init); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_sha, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, µs); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 243: if s:
							πF.SetLineno(243)
						Label1:
							// line 244: sha_update(self._sha, getbuf(s))
							πF.SetLineno(244)
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_sha, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp005[0] = µs
							if πTemp001, πE = πg.ResolveGlobal(πF, ßgetbuf); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp004[1] = πTemp002
							if πTemp001, πE = πg.ResolveGlobal(πF, ßsha_update); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
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
					if πE = πClass.SetItem(πF, ß__init__.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 246: def copy(self):
					πF.SetLineno(246)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("copy", "/usr/lib/python2.7/_sha256.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µnew *πg.Object = πg.UnboundLocal; _ = µnew
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
							// line 247: new = sha224.__new__(sha224)
							πF.SetLineno(247)
							πTemp001 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßsha224); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßsha224); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ß__new__, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µnew = πTemp002
							// line 248: new._sha = self._sha.copy()
							πF.SetLineno(248)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_sha, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßcopy, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnew, "new"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µnew, ß_sha, πTemp003); πE != nil {
								continue
							}
							// line 249: return new
							πF.SetLineno(249)
							if πE = πg.CheckLocal(πF, µnew, "new"); πE != nil {
								continue
							}
							πR = µnew
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßcopy.ToObject(), πTemp003); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp012, πE = πTemp010.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp012 == nil {
				πTemp012 = πg.TypeType.ToObject()
			}
			if πTemp013, πE = πTemp012.Call(πF, []*πg.Object{πg.NewStr("sha224").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp010.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßsha224.ToObject(), πTemp013); πE != nil {
				continue
			}
			// line 251: def test():
			πF.SetLineno(251)
			πTemp003 = make([]πg.Param, 0)
			πTemp011 = πg.NewFunction(πg.NewCode("test", "/usr/lib/python2.7/_sha256.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µa_str *πg.Object = πg.UnboundLocal; _ = µa_str
				var µs *πg.Object = πg.UnboundLocal; _ = µs
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
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
					// line 252: a_str = "just a test string"
					πF.SetLineno(252)
					µa_str = πg.NewStr("just a test string").ToObject()
					// line 254: assert 'e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855' == sha256().hexdigest()
					πF.SetLineno(254)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßsha256); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßhexdigest, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, ße3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855.ToObject(), πTemp003); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
						continue
					}
					// line 255: assert 'd7b553c6f09ac85d142415f857c5310f3bbbe7cdd787cce4b985acedd585266f' == sha256(a_str).hexdigest()
					πF.SetLineno(255)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µa_str, "a_str"); πE != nil {
						continue
					}
					πTemp004[0] = µa_str
					if πTemp002, πE = πg.ResolveGlobal(πF, ßsha256); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßhexdigest, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, ßd7b553c6f09ac85d142415f857c5310f3bbbe7cdd787cce4b985acedd585266f.ToObject(), πTemp003); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
						continue
					}
					// line 256: assert '8113ebf33c97daa9998762aacafe750c7cefc2b2f173c90c59663a57fe626f21' == sha256(a_str*7).hexdigest()
					πF.SetLineno(256)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µa_str, "a_str"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mul(πF, µa_str, πg.NewInt(7).ToObject()); πE != nil {
						continue
					}
					πTemp004[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßsha256); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßhexdigest, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, ß8113ebf33c97daa9998762aacafe750c7cefc2b2f173c90c59663a57fe626f21.ToObject(), πTemp003); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
						continue
					}
					// line 258: s = sha256(a_str)
					πF.SetLineno(258)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µa_str, "a_str"); πE != nil {
						continue
					}
					πTemp004[0] = µa_str
					if πTemp001, πE = πg.ResolveGlobal(πF, ßsha256); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µs = πTemp002
					// line 259: s.update(a_str)
					πF.SetLineno(259)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µa_str, "a_str"); πE != nil {
						continue
					}
					πTemp004[0] = µa_str
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µs, ßupdate, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 260: assert '03d9963e05a094593190b6fc794cb1a3e1ac7d7883f0b5855268afeccc70d461' == s.hexdigest()
					πF.SetLineno(260)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µs, ßhexdigest, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, ß03d9963e05a094593190b6fc794cb1a3e1ac7d7883f0b5855268afeccc70d461.ToObject(), πTemp003); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßtest.ToObject(), πTemp011); πE != nil {
				continue
			}
			if πTemp013, πE = πg.ResolveGlobal(πF, ß__name__); πE != nil {
				continue
			}
			if πTemp012, πE = πg.Eq(πF, πTemp013, ß__main__.ToObject()); πE != nil {
				continue
			}
			if πTemp014, πE = πg.IsTrue(πF, πTemp012); πE != nil {
				continue
			}
			if πTemp014 {
				goto Label1
			}
			goto Label2
			// line 262: if __name__ == "__main__":
			πF.SetLineno(262)
		Label1:
			// line 263: test()
			πF.SetLineno(263)
			if πTemp012, πE = πg.ResolveGlobal(πF, ßtest); πE != nil {
				continue
			}
			if πTemp013, πE = πTemp012.Call(πF, nil, nil); πE != nil {
				continue
			}
			goto Label2
		Label2:
		}
		return nil, πE
	})
	πg.RegisterModule("_sha256", Code)
}

