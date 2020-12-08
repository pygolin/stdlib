package _sha512
import (
	πg "github.com/pygolin/runtime"
	// _ "github.com/pygolin/stdlib/_struct"
)
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/_sha512.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ß := πg.InternStr("")
		ßB := πg.InternStr("B")
		ßCh := πg.InternStr("Ch")
		ßGamma0 := πg.InternStr("Gamma0")
		ßGamma1 := πg.InternStr("Gamma1")
		ßMaj := πg.InternStr("Maj")
		ßNone := πg.InternStr("None")
		ßR := πg.InternStr("R")
		ßROR64 := πg.InternStr("ROR64")
		ßS := πg.InternStr("S")
		ßSHA_BLOCKSIZE := πg.InternStr("SHA_BLOCKSIZE")
		ßSHA_DIGESTSIZE := πg.InternStr("SHA_DIGESTSIZE")
		ßSigma0 := πg.InternStr("Sigma0")
		ßSigma1 := πg.InternStr("Sigma1")
		ß__doc__ := πg.InternStr("__doc__")
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
		ßdata := πg.InternStr("data")
		ßdigest := πg.InternStr("digest")
		ßdigest_size := πg.InternStr("digest_size")
		ßdigestsize := πg.InternStr("digestsize")
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
		ßsha384 := πg.InternStr("sha384")
		ßsha384_init := πg.InternStr("sha384_init")
		ßsha512 := πg.InternStr("sha512")
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
			// line 1: """
			πF.SetLineno(1)
			// line 1: """
			πF.SetLineno(1)
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\nThis code was Ported from CPython's sha512module.c\n").ToObject()); πE != nil {
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
			// line 7: SHA_BLOCKSIZE = 128
			πF.SetLineno(7)
			if πE = πF.Globals().SetItem(πF, ßSHA_BLOCKSIZE.ToObject(), πg.NewInt(128).ToObject()); πE != nil {
				continue
			}
			// line 8: SHA_DIGESTSIZE = 64
			πF.SetLineno(8)
			if πE = πF.Globals().SetItem(πF, ßSHA_DIGESTSIZE.ToObject(), πg.NewInt(64).ToObject()); πE != nil {
				continue
			}
			// line 11: def new_shaobject():
			πF.SetLineno(11)
			πTemp003 = make([]πg.Param, 0)
			πTemp001 = πg.NewFunction(πg.NewCode("new_shaobject", "/usr/lib/python2.7/_sha512.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 12: return {
					πF.SetLineno(12)
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
			// line 21: ROR64 = lambda x, y: (((x & 0xffffffffffffffff) >> (y & 63)) | (x << (64 - (y & 63)))) & 0xffffffffffffffff
			πF.SetLineno(21)
			πTemp003 = make([]πg.Param, 2)
			πTemp003[0] = πg.Param{Name: "x", Def: nil}
			πTemp003[1] = πg.Param{Name: "y", Def: nil}
			πTemp004 = πg.NewFunction(πg.NewCode("<lambda>", "/usr/lib/python2.7/_sha512.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 21: ROR64 = lambda x, y: (((x & 0xffffffffffffffff) >> (y & 63)) | (x << (64 - (y & 63)))) & 0xffffffffffffffff
					πF.SetLineno(21)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.And(πF, µx, πg.NewLongFromBytes([]byte{0xff,0xff,0xff,0xff,0xff,0xff,0xff,0xff,}).ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.And(πF, µy, πg.NewInt(63).ToObject()); πE != nil {
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
					if πTemp006, πE = πg.And(πF, µy, πg.NewInt(63).ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Sub(πF, πg.NewInt(64).ToObject(), πTemp006); πE != nil {
						continue
					}
					if πTemp004, πE = πg.LShift(πF, µx, πTemp005); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Or(πF, πTemp003, πTemp004); πE != nil {
						continue
					}
					if πTemp001, πE = πg.And(πF, πTemp002, πg.NewLongFromBytes([]byte{0xff,0xff,0xff,0xff,0xff,0xff,0xff,0xff,}).ToObject()); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßROR64.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 22: Ch = lambda x, y, z: (z ^ (x & (y ^ z)))
			πF.SetLineno(22)
			πTemp003 = make([]πg.Param, 3)
			πTemp003[0] = πg.Param{Name: "x", Def: nil}
			πTemp003[1] = πg.Param{Name: "y", Def: nil}
			πTemp003[2] = πg.Param{Name: "z", Def: nil}
			πTemp004 = πg.NewFunction(πg.NewCode("<lambda>", "/usr/lib/python2.7/_sha512.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 22: Ch = lambda x, y, z: (z ^ (x & (y ^ z)))
					πF.SetLineno(22)
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
			// line 23: Maj = lambda x, y, z: (((x | y) & z) | (x & y))
			πF.SetLineno(23)
			πTemp003 = make([]πg.Param, 3)
			πTemp003[0] = πg.Param{Name: "x", Def: nil}
			πTemp003[1] = πg.Param{Name: "y", Def: nil}
			πTemp003[2] = πg.Param{Name: "z", Def: nil}
			πTemp004 = πg.NewFunction(πg.NewCode("<lambda>", "/usr/lib/python2.7/_sha512.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 23: Maj = lambda x, y, z: (((x | y) & z) | (x & y))
					πF.SetLineno(23)
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
			// line 24: S = lambda x, n: ROR64(x, n)
			πF.SetLineno(24)
			πTemp003 = make([]πg.Param, 2)
			πTemp003[0] = πg.Param{Name: "x", Def: nil}
			πTemp003[1] = πg.Param{Name: "n", Def: nil}
			πTemp004 = πg.NewFunction(πg.NewCode("<lambda>", "/usr/lib/python2.7/_sha512.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 24: S = lambda x, n: ROR64(x, n)
					πF.SetLineno(24)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp001[0] = µx
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					πTemp001[1] = µn
					if πTemp002, πE = πg.ResolveGlobal(πF, ßROR64); πE != nil {
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
			// line 25: R = lambda x, n: (x & 0xffffffffffffffff) >> n
			πF.SetLineno(25)
			πTemp003 = make([]πg.Param, 2)
			πTemp003[0] = πg.Param{Name: "x", Def: nil}
			πTemp003[1] = πg.Param{Name: "n", Def: nil}
			πTemp004 = πg.NewFunction(πg.NewCode("<lambda>", "/usr/lib/python2.7/_sha512.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 25: R = lambda x, n: (x & 0xffffffffffffffff) >> n
					πF.SetLineno(25)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.And(πF, µx, πg.NewLongFromBytes([]byte{0xff,0xff,0xff,0xff,0xff,0xff,0xff,0xff,}).ToObject()); πE != nil {
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
			// line 26: Sigma0 = lambda x: (S(x, 28) ^ S(x, 34) ^ S(x, 39))
			πF.SetLineno(26)
			πTemp003 = make([]πg.Param, 1)
			πTemp003[0] = πg.Param{Name: "x", Def: nil}
			πTemp004 = πg.NewFunction(πg.NewCode("<lambda>", "/usr/lib/python2.7/_sha512.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 26: Sigma0 = lambda x: (S(x, 28) ^ S(x, 34) ^ S(x, 39))
					πF.SetLineno(26)
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp003[0] = µx
					πTemp003[1] = πg.NewInt(28).ToObject()
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
					πTemp003[1] = πg.NewInt(34).ToObject()
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
					πTemp003[1] = πg.NewInt(39).ToObject()
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
			// line 27: Sigma1 = lambda x: (S(x, 14) ^ S(x, 18) ^ S(x, 41))
			πF.SetLineno(27)
			πTemp003 = make([]πg.Param, 1)
			πTemp003[0] = πg.Param{Name: "x", Def: nil}
			πTemp004 = πg.NewFunction(πg.NewCode("<lambda>", "/usr/lib/python2.7/_sha512.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 27: Sigma1 = lambda x: (S(x, 14) ^ S(x, 18) ^ S(x, 41))
					πF.SetLineno(27)
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp003[0] = µx
					πTemp003[1] = πg.NewInt(14).ToObject()
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
					πTemp003[1] = πg.NewInt(41).ToObject()
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
			// line 28: Gamma0 = lambda x: (S(x, 1) ^ S(x, 8) ^ R(x, 7))
			πF.SetLineno(28)
			πTemp003 = make([]πg.Param, 1)
			πTemp003[0] = πg.Param{Name: "x", Def: nil}
			πTemp004 = πg.NewFunction(πg.NewCode("<lambda>", "/usr/lib/python2.7/_sha512.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 28: Gamma0 = lambda x: (S(x, 1) ^ S(x, 8) ^ R(x, 7))
					πF.SetLineno(28)
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp003[0] = µx
					πTemp003[1] = πg.NewInt(1).ToObject()
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
					πTemp003[1] = πg.NewInt(8).ToObject()
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
					πTemp003[1] = πg.NewInt(7).ToObject()
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
			// line 29: Gamma1 = lambda x: (S(x, 19) ^ S(x, 61) ^ R(x, 6))
			πF.SetLineno(29)
			πTemp003 = make([]πg.Param, 1)
			πTemp003[0] = πg.Param{Name: "x", Def: nil}
			πTemp004 = πg.NewFunction(πg.NewCode("<lambda>", "/usr/lib/python2.7/_sha512.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 29: Gamma1 = lambda x: (S(x, 19) ^ S(x, 61) ^ R(x, 6))
					πF.SetLineno(29)
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp003[0] = µx
					πTemp003[1] = πg.NewInt(19).ToObject()
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
					πTemp003[1] = πg.NewInt(61).ToObject()
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
					πTemp003[1] = πg.NewInt(6).ToObject()
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
			// line 31: def sha_transform(sha_info):
			πF.SetLineno(31)
			πTemp003 = make([]πg.Param, 1)
			πTemp003[0] = πg.Param{Name: "sha_info", Def: nil}
			πTemp004 = πg.NewFunction(πg.NewCode("sha_transform", "/usr/lib/python2.7/_sha512.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
				var πTemp013 *πg.Object
				_ = πTemp013
				var πTemp014 *πg.Object
				_ = πTemp014
				var πTemp015 *πg.Object
				_ = πTemp015
				var πTemp016 *πg.Object
				_ = πTemp016
				var πTemp017 []*πg.Object
				_ = πTemp017
				var πTemp018 []πg.Param
				_ = πTemp018
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
					// line 32: W = []
					πF.SetLineno(32)
					πTemp001 = make([]*πg.Object, 0)
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					µW = πTemp002
					// line 34: d = sha_info['data']
					πF.SetLineno(34)
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
					// line 36: W.append( (d[8*i]<<56) + (d[8*i+1]<<48) + (d[8*i+2]<<40) + (d[8*i+3]<<32) + (d[8*i+4]<<24) + (d[8*i+5]<<16) + (d[8*i+6]<<8) + d[8*i+7])
					πF.SetLineno(36)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp014, πE = πg.Mul(πF, πg.NewInt(8).ToObject(), µi); πE != nil {
						continue
					}
					πTemp013 = πTemp014
					if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
						continue
					}
					if πTemp014, πE = πg.GetItem(πF, µd, πTemp013); πE != nil {
						continue
					}
					if πTemp012, πE = πg.LShift(πF, πTemp014, πg.NewInt(56).ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp016, πE = πg.Mul(πF, πg.NewInt(8).ToObject(), µi); πE != nil {
						continue
					}
					if πTemp015, πE = πg.Add(πF, πTemp016, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp014 = πTemp015
					if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
						continue
					}
					if πTemp015, πE = πg.GetItem(πF, µd, πTemp014); πE != nil {
						continue
					}
					if πTemp013, πE = πg.LShift(πF, πTemp015, πg.NewInt(48).ToObject()); πE != nil {
						continue
					}
					if πTemp011, πE = πg.Add(πF, πTemp012, πTemp013); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp015, πE = πg.Mul(πF, πg.NewInt(8).ToObject(), µi); πE != nil {
						continue
					}
					if πTemp014, πE = πg.Add(πF, πTemp015, πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
					πTemp013 = πTemp014
					if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
						continue
					}
					if πTemp014, πE = πg.GetItem(πF, µd, πTemp013); πE != nil {
						continue
					}
					if πTemp012, πE = πg.LShift(πF, πTemp014, πg.NewInt(40).ToObject()); πE != nil {
						continue
					}
					if πTemp010, πE = πg.Add(πF, πTemp011, πTemp012); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp014, πE = πg.Mul(πF, πg.NewInt(8).ToObject(), µi); πE != nil {
						continue
					}
					if πTemp013, πE = πg.Add(πF, πTemp014, πg.NewInt(3).ToObject()); πE != nil {
						continue
					}
					πTemp012 = πTemp013
					if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
						continue
					}
					if πTemp013, πE = πg.GetItem(πF, µd, πTemp012); πE != nil {
						continue
					}
					if πTemp011, πE = πg.LShift(πF, πTemp013, πg.NewInt(32).ToObject()); πE != nil {
						continue
					}
					if πTemp009, πE = πg.Add(πF, πTemp010, πTemp011); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp013, πE = πg.Mul(πF, πg.NewInt(8).ToObject(), µi); πE != nil {
						continue
					}
					if πTemp012, πE = πg.Add(πF, πTemp013, πg.NewInt(4).ToObject()); πE != nil {
						continue
					}
					πTemp011 = πTemp012
					if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
						continue
					}
					if πTemp012, πE = πg.GetItem(πF, µd, πTemp011); πE != nil {
						continue
					}
					if πTemp010, πE = πg.LShift(πF, πTemp012, πg.NewInt(24).ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.Add(πF, πTemp009, πTemp010); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp012, πE = πg.Mul(πF, πg.NewInt(8).ToObject(), µi); πE != nil {
						continue
					}
					if πTemp011, πE = πg.Add(πF, πTemp012, πg.NewInt(5).ToObject()); πE != nil {
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
					if πTemp011, πE = πg.Mul(πF, πg.NewInt(8).ToObject(), µi); πE != nil {
						continue
					}
					if πTemp010, πE = πg.Add(πF, πTemp011, πg.NewInt(6).ToObject()); πE != nil {
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
					if πTemp009, πE = πg.Mul(πF, πg.NewInt(8).ToObject(), µi); πE != nil {
						continue
					}
					if πTemp008, πE = πg.Add(πF, πTemp009, πg.NewInt(7).ToObject()); πE != nil {
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
					πTemp001[1] = πg.NewInt(80).ToObject()
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
					// line 39: W.append( (Gamma1(W[i - 2]) + W[i - 7] + Gamma0(W[i - 15]) + W[i - 16]) & 0xffffffffffffffff )
					πF.SetLineno(39)
					πTemp001 = πF.MakeArgs(1)
					πTemp017 = πF.MakeArgs(1)
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
					πTemp017[0] = πTemp010
					if πTemp009, πE = πg.ResolveGlobal(πF, ßGamma1); πE != nil {
						continue
					}
					if πTemp010, πE = πTemp009.Call(πF, πTemp017, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp017)
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
					πTemp017 = πF.MakeArgs(1)
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
					πTemp017[0] = πTemp010
					if πTemp009, πE = πg.ResolveGlobal(πF, ßGamma0); πE != nil {
						continue
					}
					if πTemp010, πE = πTemp009.Call(πF, πTemp017, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp017)
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
					if πTemp003, πE = πg.And(πF, πTemp004, πg.NewLongFromBytes([]byte{0xff,0xff,0xff,0xff,0xff,0xff,0xff,0xff,}).ToObject()); πE != nil {
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
					// line 41: ss = sha_info['digest'][:]
					πF.SetLineno(41)
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
					// line 43: def RND(a,b,c,d,e,f,g,h,i,ki):
					πF.SetLineno(43)
					πTemp018 = make([]πg.Param, 10)
					πTemp018[0] = πg.Param{Name: "a", Def: nil}
					πTemp018[1] = πg.Param{Name: "b", Def: nil}
					πTemp018[2] = πg.Param{Name: "c", Def: nil}
					πTemp018[3] = πg.Param{Name: "d", Def: nil}
					πTemp018[4] = πg.Param{Name: "e", Def: nil}
					πTemp018[5] = πg.Param{Name: "f", Def: nil}
					πTemp018[6] = πg.Param{Name: "g", Def: nil}
					πTemp018[7] = πg.Param{Name: "h", Def: nil}
					πTemp018[8] = πg.Param{Name: "i", Def: nil}
					πTemp018[9] = πg.Param{Name: "ki", Def: nil}
					πTemp002 = πg.NewFunction(πg.NewCode("RND", "/usr/lib/python2.7/_sha512.py", πTemp018, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
						var πTemp005 *πg.Object
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
							default: panic("unexpected function state")
							}
							// line 44: t0 = (h + Sigma1(e) + Ch(e, f, g) + ki + W[i]) & 0xffffffffffffffff
							πF.SetLineno(44)
							if πE = πg.CheckLocal(πF, µh, "h"); πE != nil {
								continue
							}
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µe, "e"); πE != nil {
								continue
							}
							πTemp006[0] = µe
							if πTemp007, πE = πg.ResolveGlobal(πF, ßSigma1); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp007.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πTemp005, πE = πg.Add(πF, µh, πTemp008); πE != nil {
								continue
							}
							πTemp006 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µe, "e"); πE != nil {
								continue
							}
							πTemp006[0] = µe
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							πTemp006[1] = µf
							if πE = πg.CheckLocal(πF, µg, "g"); πE != nil {
								continue
							}
							πTemp006[2] = µg
							if πTemp007, πE = πg.ResolveGlobal(πF, ßCh); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp007.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πTemp004, πE = πg.Add(πF, πTemp005, πTemp008); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µki, "ki"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, πTemp004, µki); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp004 = µi
							if πE = πg.CheckLocal(πF, µW, "W"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µW, πTemp004); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp003, πTemp005); πE != nil {
								continue
							}
							if πTemp001, πE = πg.And(πF, πTemp002, πg.NewLongFromBytes([]byte{0xff,0xff,0xff,0xff,0xff,0xff,0xff,0xff,}).ToObject()); πE != nil {
								continue
							}
							µt0 = πTemp001
							// line 45: t1 = (Sigma0(a) + Maj(a, b, c)) & 0xffffffffffffffff
							πF.SetLineno(45)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp006[0] = µa
							if πTemp003, πE = πg.ResolveGlobal(πF, ßSigma0); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp006 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp006[0] = µa
							if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
								continue
							}
							πTemp006[1] = µb
							if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
								continue
							}
							πTemp006[2] = µc
							if πTemp003, πE = πg.ResolveGlobal(πF, ßMaj); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πTemp002, πE = πg.Add(πF, πTemp004, πTemp005); πE != nil {
								continue
							}
							if πTemp001, πE = πg.And(πF, πTemp002, πg.NewLongFromBytes([]byte{0xff,0xff,0xff,0xff,0xff,0xff,0xff,0xff,}).ToObject()); πE != nil {
								continue
							}
							µt1 = πTemp001
							// line 46: d = (d + t0) & 0xffffffffffffffff
							πF.SetLineno(46)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µt0, "t0"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, µd, µt0); πE != nil {
								continue
							}
							if πTemp001, πE = πg.And(πF, πTemp002, πg.NewLongFromBytes([]byte{0xff,0xff,0xff,0xff,0xff,0xff,0xff,0xff,}).ToObject()); πE != nil {
								continue
							}
							µd = πTemp001
							// line 47: h = (t0 + t1) & 0xffffffffffffffff
							πF.SetLineno(47)
							if πE = πg.CheckLocal(πF, µt0, "t0"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µt1, "t1"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, µt0, µt1); πE != nil {
								continue
							}
							if πTemp001, πE = πg.And(πF, πTemp002, πg.NewLongFromBytes([]byte{0xff,0xff,0xff,0xff,0xff,0xff,0xff,0xff,}).ToObject()); πE != nil {
								continue
							}
							µh = πTemp001
							// line 48: return d & 0xffffffffffffffff, h & 0xffffffffffffffff
							πF.SetLineno(48)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.And(πF, µd, πg.NewLongFromBytes([]byte{0xff,0xff,0xff,0xff,0xff,0xff,0xff,0xff,}).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µh, "h"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.And(πF, µh, πg.NewLongFromBytes([]byte{0xff,0xff,0xff,0xff,0xff,0xff,0xff,0xff,}).ToObject()); πE != nil {
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
					// line 50: ss[3], ss[7] = RND(ss[0],ss[1],ss[2],ss[3],ss[4],ss[5],ss[6],ss[7],0,0x428a2f98d728ae22)
					πF.SetLineno(50)
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
					πTemp001[9] = πg.NewInt(4794697086780616226).ToObject()
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
					// line 51: ss[2], ss[6] = RND(ss[7],ss[0],ss[1],ss[2],ss[3],ss[4],ss[5],ss[6],1,0x7137449123ef65cd)
					πF.SetLineno(51)
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
					πTemp001[9] = πg.NewInt(8158064640168781261).ToObject()
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
					// line 52: ss[1], ss[5] = RND(ss[6],ss[7],ss[0],ss[1],ss[2],ss[3],ss[4],ss[5],2,0xb5c0fbcfec4d3b2f)
					πF.SetLineno(52)
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
					πTemp001[9] = πg.NewLongFromBytes([]byte{0xb5,0xc0,0xfb,0xcf,0xec,0x4d,0x3b,0x2f,}).ToObject()
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
					// line 53: ss[0], ss[4] = RND(ss[5],ss[6],ss[7],ss[0],ss[1],ss[2],ss[3],ss[4],3,0xe9b5dba58189dbbc)
					πF.SetLineno(53)
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
					πTemp001[9] = πg.NewLongFromBytes([]byte{0xe9,0xb5,0xdb,0xa5,0x81,0x89,0xdb,0xbc,}).ToObject()
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
					// line 54: ss[7], ss[3] = RND(ss[4],ss[5],ss[6],ss[7],ss[0],ss[1],ss[2],ss[3],4,0x3956c25bf348b538)
					πF.SetLineno(54)
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
					πTemp001[9] = πg.NewInt(4131703408338449720).ToObject()
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
					// line 55: ss[6], ss[2] = RND(ss[3],ss[4],ss[5],ss[6],ss[7],ss[0],ss[1],ss[2],5,0x59f111f1b605d019)
					πF.SetLineno(55)
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
					πTemp001[9] = πg.NewInt(6480981068601479193).ToObject()
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
					// line 56: ss[5], ss[1] = RND(ss[2],ss[3],ss[4],ss[5],ss[6],ss[7],ss[0],ss[1],6,0x923f82a4af194f9b)
					πF.SetLineno(56)
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
					πTemp001[9] = πg.NewLongFromBytes([]byte{0x92,0x3f,0x82,0xa4,0xaf,0x19,0x4f,0x9b,}).ToObject()
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
					// line 57: ss[4], ss[0] = RND(ss[1],ss[2],ss[3],ss[4],ss[5],ss[6],ss[7],ss[0],7,0xab1c5ed5da6d8118)
					πF.SetLineno(57)
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
					πTemp001[9] = πg.NewLongFromBytes([]byte{0xab,0x1c,0x5e,0xd5,0xda,0x6d,0x81,0x18,}).ToObject()
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
					// line 58: ss[3], ss[7] = RND(ss[0],ss[1],ss[2],ss[3],ss[4],ss[5],ss[6],ss[7],8,0xd807aa98a3030242)
					πF.SetLineno(58)
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
					πTemp001[9] = πg.NewLongFromBytes([]byte{0xd8,0x7,0xaa,0x98,0xa3,0x3,0x2,0x42,}).ToObject()
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
					// line 59: ss[2], ss[6] = RND(ss[7],ss[0],ss[1],ss[2],ss[3],ss[4],ss[5],ss[6],9,0x12835b0145706fbe)
					πF.SetLineno(59)
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
					πTemp001[9] = πg.NewInt(1334009975649890238).ToObject()
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
					// line 60: ss[1], ss[5] = RND(ss[6],ss[7],ss[0],ss[1],ss[2],ss[3],ss[4],ss[5],10,0x243185be4ee4b28c)
					πF.SetLineno(60)
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
					πTemp001[9] = πg.NewInt(2608012711638119052).ToObject()
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
					// line 61: ss[0], ss[4] = RND(ss[5],ss[6],ss[7],ss[0],ss[1],ss[2],ss[3],ss[4],11,0x550c7dc3d5ffb4e2)
					πF.SetLineno(61)
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
					πTemp001[9] = πg.NewInt(6128411473006802146).ToObject()
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
					// line 62: ss[7], ss[3] = RND(ss[4],ss[5],ss[6],ss[7],ss[0],ss[1],ss[2],ss[3],12,0x72be5d74f27b896f)
					πF.SetLineno(62)
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
					πTemp001[9] = πg.NewInt(8268148722764581231).ToObject()
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
					// line 63: ss[6], ss[2] = RND(ss[3],ss[4],ss[5],ss[6],ss[7],ss[0],ss[1],ss[2],13,0x80deb1fe3b1696b1)
					πF.SetLineno(63)
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
					πTemp001[9] = πg.NewLongFromBytes([]byte{0x80,0xde,0xb1,0xfe,0x3b,0x16,0x96,0xb1,}).ToObject()
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
					// line 64: ss[5], ss[1] = RND(ss[2],ss[3],ss[4],ss[5],ss[6],ss[7],ss[0],ss[1],14,0x9bdc06a725c71235)
					πF.SetLineno(64)
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
					πTemp001[9] = πg.NewLongFromBytes([]byte{0x9b,0xdc,0x6,0xa7,0x25,0xc7,0x12,0x35,}).ToObject()
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
					// line 65: ss[4], ss[0] = RND(ss[1],ss[2],ss[3],ss[4],ss[5],ss[6],ss[7],ss[0],15,0xc19bf174cf692694)
					πF.SetLineno(65)
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
					πTemp001[9] = πg.NewLongFromBytes([]byte{0xc1,0x9b,0xf1,0x74,0xcf,0x69,0x26,0x94,}).ToObject()
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
					// line 66: ss[3], ss[7] = RND(ss[0],ss[1],ss[2],ss[3],ss[4],ss[5],ss[6],ss[7],16,0xe49b69c19ef14ad2)
					πF.SetLineno(66)
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
					πTemp001[9] = πg.NewLongFromBytes([]byte{0xe4,0x9b,0x69,0xc1,0x9e,0xf1,0x4a,0xd2,}).ToObject()
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
					// line 67: ss[2], ss[6] = RND(ss[7],ss[0],ss[1],ss[2],ss[3],ss[4],ss[5],ss[6],17,0xefbe4786384f25e3)
					πF.SetLineno(67)
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
					πTemp001[9] = πg.NewLongFromBytes([]byte{0xef,0xbe,0x47,0x86,0x38,0x4f,0x25,0xe3,}).ToObject()
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
					// line 68: ss[1], ss[5] = RND(ss[6],ss[7],ss[0],ss[1],ss[2],ss[3],ss[4],ss[5],18,0x0fc19dc68b8cd5b5)
					πF.SetLineno(68)
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
					πTemp001[9] = πg.NewInt(1135362057144423861).ToObject()
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
					// line 69: ss[0], ss[4] = RND(ss[5],ss[6],ss[7],ss[0],ss[1],ss[2],ss[3],ss[4],19,0x240ca1cc77ac9c65)
					πF.SetLineno(69)
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
					πTemp001[9] = πg.NewInt(2597628984639134821).ToObject()
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
					// line 70: ss[7], ss[3] = RND(ss[4],ss[5],ss[6],ss[7],ss[0],ss[1],ss[2],ss[3],20,0x2de92c6f592b0275)
					πF.SetLineno(70)
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
					πTemp001[9] = πg.NewInt(3308224258029322869).ToObject()
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
					// line 71: ss[6], ss[2] = RND(ss[3],ss[4],ss[5],ss[6],ss[7],ss[0],ss[1],ss[2],21,0x4a7484aa6ea6e483)
					πF.SetLineno(71)
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
					πTemp001[9] = πg.NewInt(5365058923640841347).ToObject()
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
					// line 72: ss[5], ss[1] = RND(ss[2],ss[3],ss[4],ss[5],ss[6],ss[7],ss[0],ss[1],22,0x5cb0a9dcbd41fbd4)
					πF.SetLineno(72)
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
					πTemp001[9] = πg.NewInt(6679025012923562964).ToObject()
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
					// line 73: ss[4], ss[0] = RND(ss[1],ss[2],ss[3],ss[4],ss[5],ss[6],ss[7],ss[0],23,0x76f988da831153b5)
					πF.SetLineno(73)
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
					πTemp001[9] = πg.NewInt(8573033837759648693).ToObject()
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
					// line 74: ss[3], ss[7] = RND(ss[0],ss[1],ss[2],ss[3],ss[4],ss[5],ss[6],ss[7],24,0x983e5152ee66dfab)
					πF.SetLineno(74)
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
					πTemp001[9] = πg.NewLongFromBytes([]byte{0x98,0x3e,0x51,0x52,0xee,0x66,0xdf,0xab,}).ToObject()
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
					// line 75: ss[2], ss[6] = RND(ss[7],ss[0],ss[1],ss[2],ss[3],ss[4],ss[5],ss[6],25,0xa831c66d2db43210)
					πF.SetLineno(75)
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
					πTemp001[9] = πg.NewLongFromBytes([]byte{0xa8,0x31,0xc6,0x6d,0x2d,0xb4,0x32,0x10,}).ToObject()
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
					// line 76: ss[1], ss[5] = RND(ss[6],ss[7],ss[0],ss[1],ss[2],ss[3],ss[4],ss[5],26,0xb00327c898fb213f)
					πF.SetLineno(76)
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
					πTemp001[9] = πg.NewLongFromBytes([]byte{0xb0,0x3,0x27,0xc8,0x98,0xfb,0x21,0x3f,}).ToObject()
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
					// line 77: ss[0], ss[4] = RND(ss[5],ss[6],ss[7],ss[0],ss[1],ss[2],ss[3],ss[4],27,0xbf597fc7beef0ee4)
					πF.SetLineno(77)
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
					πTemp001[9] = πg.NewLongFromBytes([]byte{0xbf,0x59,0x7f,0xc7,0xbe,0xef,0xe,0xe4,}).ToObject()
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
					// line 78: ss[7], ss[3] = RND(ss[4],ss[5],ss[6],ss[7],ss[0],ss[1],ss[2],ss[3],28,0xc6e00bf33da88fc2)
					πF.SetLineno(78)
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
					πTemp001[9] = πg.NewLongFromBytes([]byte{0xc6,0xe0,0xb,0xf3,0x3d,0xa8,0x8f,0xc2,}).ToObject()
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
					// line 79: ss[6], ss[2] = RND(ss[3],ss[4],ss[5],ss[6],ss[7],ss[0],ss[1],ss[2],29,0xd5a79147930aa725)
					πF.SetLineno(79)
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
					πTemp001[9] = πg.NewLongFromBytes([]byte{0xd5,0xa7,0x91,0x47,0x93,0xa,0xa7,0x25,}).ToObject()
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
					// line 80: ss[5], ss[1] = RND(ss[2],ss[3],ss[4],ss[5],ss[6],ss[7],ss[0],ss[1],30,0x06ca6351e003826f)
					πF.SetLineno(80)
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
					πTemp001[9] = πg.NewInt(489312712824947311).ToObject()
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
					// line 81: ss[4], ss[0] = RND(ss[1],ss[2],ss[3],ss[4],ss[5],ss[6],ss[7],ss[0],31,0x142929670a0e6e70)
					πF.SetLineno(81)
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
					πTemp001[9] = πg.NewInt(1452737877330783856).ToObject()
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
					// line 82: ss[3], ss[7] = RND(ss[0],ss[1],ss[2],ss[3],ss[4],ss[5],ss[6],ss[7],32,0x27b70a8546d22ffc)
					πF.SetLineno(82)
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
					πTemp001[9] = πg.NewInt(2861767655752347644).ToObject()
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
					// line 83: ss[2], ss[6] = RND(ss[7],ss[0],ss[1],ss[2],ss[3],ss[4],ss[5],ss[6],33,0x2e1b21385c26c926)
					πF.SetLineno(83)
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
					πTemp001[9] = πg.NewInt(3322285676063803686).ToObject()
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
					// line 84: ss[1], ss[5] = RND(ss[6],ss[7],ss[0],ss[1],ss[2],ss[3],ss[4],ss[5],34,0x4d2c6dfc5ac42aed)
					πF.SetLineno(84)
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
					πTemp001[9] = πg.NewInt(5560940570517711597).ToObject()
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
					// line 85: ss[0], ss[4] = RND(ss[5],ss[6],ss[7],ss[0],ss[1],ss[2],ss[3],ss[4],35,0x53380d139d95b3df)
					πF.SetLineno(85)
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
					πTemp001[9] = πg.NewInt(5996557281743188959).ToObject()
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
					// line 86: ss[7], ss[3] = RND(ss[4],ss[5],ss[6],ss[7],ss[0],ss[1],ss[2],ss[3],36,0x650a73548baf63de)
					πF.SetLineno(86)
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
					πTemp001[9] = πg.NewInt(7280758554555802590).ToObject()
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
					// line 87: ss[6], ss[2] = RND(ss[3],ss[4],ss[5],ss[6],ss[7],ss[0],ss[1],ss[2],37,0x766a0abb3c77b2a8)
					πF.SetLineno(87)
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
					πTemp001[9] = πg.NewInt(8532644243296465576).ToObject()
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
					// line 88: ss[5], ss[1] = RND(ss[2],ss[3],ss[4],ss[5],ss[6],ss[7],ss[0],ss[1],38,0x81c2c92e47edaee6)
					πF.SetLineno(88)
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
					πTemp001[9] = πg.NewLongFromBytes([]byte{0x81,0xc2,0xc9,0x2e,0x47,0xed,0xae,0xe6,}).ToObject()
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
					// line 89: ss[4], ss[0] = RND(ss[1],ss[2],ss[3],ss[4],ss[5],ss[6],ss[7],ss[0],39,0x92722c851482353b)
					πF.SetLineno(89)
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
					πTemp001[9] = πg.NewLongFromBytes([]byte{0x92,0x72,0x2c,0x85,0x14,0x82,0x35,0x3b,}).ToObject()
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
					// line 90: ss[3], ss[7] = RND(ss[0],ss[1],ss[2],ss[3],ss[4],ss[5],ss[6],ss[7],40,0xa2bfe8a14cf10364)
					πF.SetLineno(90)
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
					πTemp001[9] = πg.NewLongFromBytes([]byte{0xa2,0xbf,0xe8,0xa1,0x4c,0xf1,0x3,0x64,}).ToObject()
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
					// line 91: ss[2], ss[6] = RND(ss[7],ss[0],ss[1],ss[2],ss[3],ss[4],ss[5],ss[6],41,0xa81a664bbc423001)
					πF.SetLineno(91)
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
					πTemp001[9] = πg.NewLongFromBytes([]byte{0xa8,0x1a,0x66,0x4b,0xbc,0x42,0x30,0x1,}).ToObject()
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
					// line 92: ss[1], ss[5] = RND(ss[6],ss[7],ss[0],ss[1],ss[2],ss[3],ss[4],ss[5],42,0xc24b8b70d0f89791)
					πF.SetLineno(92)
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
					πTemp001[9] = πg.NewLongFromBytes([]byte{0xc2,0x4b,0x8b,0x70,0xd0,0xf8,0x97,0x91,}).ToObject()
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
					// line 93: ss[0], ss[4] = RND(ss[5],ss[6],ss[7],ss[0],ss[1],ss[2],ss[3],ss[4],43,0xc76c51a30654be30)
					πF.SetLineno(93)
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
					πTemp001[9] = πg.NewLongFromBytes([]byte{0xc7,0x6c,0x51,0xa3,0x6,0x54,0xbe,0x30,}).ToObject()
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
					// line 94: ss[7], ss[3] = RND(ss[4],ss[5],ss[6],ss[7],ss[0],ss[1],ss[2],ss[3],44,0xd192e819d6ef5218)
					πF.SetLineno(94)
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
					πTemp001[9] = πg.NewLongFromBytes([]byte{0xd1,0x92,0xe8,0x19,0xd6,0xef,0x52,0x18,}).ToObject()
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
					// line 95: ss[6], ss[2] = RND(ss[3],ss[4],ss[5],ss[6],ss[7],ss[0],ss[1],ss[2],45,0xd69906245565a910)
					πF.SetLineno(95)
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
					πTemp001[9] = πg.NewLongFromBytes([]byte{0xd6,0x99,0x6,0x24,0x55,0x65,0xa9,0x10,}).ToObject()
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
					// line 96: ss[5], ss[1] = RND(ss[2],ss[3],ss[4],ss[5],ss[6],ss[7],ss[0],ss[1],46,0xf40e35855771202a)
					πF.SetLineno(96)
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
					πTemp001[9] = πg.NewLongFromBytes([]byte{0xf4,0xe,0x35,0x85,0x57,0x71,0x20,0x2a,}).ToObject()
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
					// line 97: ss[4], ss[0] = RND(ss[1],ss[2],ss[3],ss[4],ss[5],ss[6],ss[7],ss[0],47,0x106aa07032bbd1b8)
					πF.SetLineno(97)
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
					πTemp001[9] = πg.NewInt(1182934255886127544).ToObject()
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
					// line 98: ss[3], ss[7] = RND(ss[0],ss[1],ss[2],ss[3],ss[4],ss[5],ss[6],ss[7],48,0x19a4c116b8d2d0c8)
					πF.SetLineno(98)
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
					πTemp001[9] = πg.NewInt(1847814050463011016).ToObject()
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
					// line 99: ss[2], ss[6] = RND(ss[7],ss[0],ss[1],ss[2],ss[3],ss[4],ss[5],ss[6],49,0x1e376c085141ab53)
					πF.SetLineno(99)
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
					πTemp001[9] = πg.NewInt(2177327727835720531).ToObject()
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
					// line 100: ss[1], ss[5] = RND(ss[6],ss[7],ss[0],ss[1],ss[2],ss[3],ss[4],ss[5],50,0x2748774cdf8eeb99)
					πF.SetLineno(100)
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
					πTemp001[9] = πg.NewInt(2830643537854262169).ToObject()
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
					// line 101: ss[0], ss[4] = RND(ss[5],ss[6],ss[7],ss[0],ss[1],ss[2],ss[3],ss[4],51,0x34b0bcb5e19b48a8)
					πF.SetLineno(101)
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
					πTemp001[9] = πg.NewInt(3796741975233480872).ToObject()
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
					// line 102: ss[7], ss[3] = RND(ss[4],ss[5],ss[6],ss[7],ss[0],ss[1],ss[2],ss[3],52,0x391c0cb3c5c95a63)
					πF.SetLineno(102)
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
					πTemp001[9] = πg.NewInt(4115178125766777443).ToObject()
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
					// line 103: ss[6], ss[2] = RND(ss[3],ss[4],ss[5],ss[6],ss[7],ss[0],ss[1],ss[2],53,0x4ed8aa4ae3418acb)
					πF.SetLineno(103)
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
					πTemp001[9] = πg.NewInt(5681478168544905931).ToObject()
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
					// line 104: ss[5], ss[1] = RND(ss[2],ss[3],ss[4],ss[5],ss[6],ss[7],ss[0],ss[1],54,0x5b9cca4f7763e373)
					πF.SetLineno(104)
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
					πTemp001[9] = πg.NewInt(6601373596472566643).ToObject()
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
					// line 105: ss[4], ss[0] = RND(ss[1],ss[2],ss[3],ss[4],ss[5],ss[6],ss[7],ss[0],55,0x682e6ff3d6b2b8a3)
					πF.SetLineno(105)
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
					πTemp001[9] = πg.NewInt(7507060721942968483).ToObject()
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
					// line 106: ss[3], ss[7] = RND(ss[0],ss[1],ss[2],ss[3],ss[4],ss[5],ss[6],ss[7],56,0x748f82ee5defb2fc)
					πF.SetLineno(106)
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
					πTemp001[9] = πg.NewInt(8399075790359081724).ToObject()
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
					// line 107: ss[2], ss[6] = RND(ss[7],ss[0],ss[1],ss[2],ss[3],ss[4],ss[5],ss[6],57,0x78a5636f43172f60)
					πF.SetLineno(107)
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
					πTemp001[9] = πg.NewInt(8693463985226723168).ToObject()
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
					// line 108: ss[1], ss[5] = RND(ss[6],ss[7],ss[0],ss[1],ss[2],ss[3],ss[4],ss[5],58,0x84c87814a1f0ab72)
					πF.SetLineno(108)
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
					πTemp001[9] = πg.NewLongFromBytes([]byte{0x84,0xc8,0x78,0x14,0xa1,0xf0,0xab,0x72,}).ToObject()
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
					// line 109: ss[0], ss[4] = RND(ss[5],ss[6],ss[7],ss[0],ss[1],ss[2],ss[3],ss[4],59,0x8cc702081a6439ec)
					πF.SetLineno(109)
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
					πTemp001[9] = πg.NewLongFromBytes([]byte{0x8c,0xc7,0x2,0x8,0x1a,0x64,0x39,0xec,}).ToObject()
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
					// line 110: ss[7], ss[3] = RND(ss[4],ss[5],ss[6],ss[7],ss[0],ss[1],ss[2],ss[3],60,0x90befffa23631e28)
					πF.SetLineno(110)
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
					πTemp001[9] = πg.NewLongFromBytes([]byte{0x90,0xbe,0xff,0xfa,0x23,0x63,0x1e,0x28,}).ToObject()
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
					// line 111: ss[6], ss[2] = RND(ss[3],ss[4],ss[5],ss[6],ss[7],ss[0],ss[1],ss[2],61,0xa4506cebde82bde9)
					πF.SetLineno(111)
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
					πTemp001[9] = πg.NewLongFromBytes([]byte{0xa4,0x50,0x6c,0xeb,0xde,0x82,0xbd,0xe9,}).ToObject()
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
					// line 112: ss[5], ss[1] = RND(ss[2],ss[3],ss[4],ss[5],ss[6],ss[7],ss[0],ss[1],62,0xbef9a3f7b2c67915)
					πF.SetLineno(112)
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
					πTemp001[9] = πg.NewLongFromBytes([]byte{0xbe,0xf9,0xa3,0xf7,0xb2,0xc6,0x79,0x15,}).ToObject()
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
					// line 113: ss[4], ss[0] = RND(ss[1],ss[2],ss[3],ss[4],ss[5],ss[6],ss[7],ss[0],63,0xc67178f2e372532b)
					πF.SetLineno(113)
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
					πTemp001[9] = πg.NewLongFromBytes([]byte{0xc6,0x71,0x78,0xf2,0xe3,0x72,0x53,0x2b,}).ToObject()
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
					// line 114: ss[3], ss[7] = RND(ss[0],ss[1],ss[2],ss[3],ss[4],ss[5],ss[6],ss[7],64,0xca273eceea26619c)
					πF.SetLineno(114)
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
					πTemp001[8] = πg.NewInt(64).ToObject()
					πTemp001[9] = πg.NewLongFromBytes([]byte{0xca,0x27,0x3e,0xce,0xea,0x26,0x61,0x9c,}).ToObject()
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
					// line 115: ss[2], ss[6] = RND(ss[7],ss[0],ss[1],ss[2],ss[3],ss[4],ss[5],ss[6],65,0xd186b8c721c0c207)
					πF.SetLineno(115)
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
					πTemp001[8] = πg.NewInt(65).ToObject()
					πTemp001[9] = πg.NewLongFromBytes([]byte{0xd1,0x86,0xb8,0xc7,0x21,0xc0,0xc2,0x7,}).ToObject()
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
					// line 116: ss[1], ss[5] = RND(ss[6],ss[7],ss[0],ss[1],ss[2],ss[3],ss[4],ss[5],66,0xeada7dd6cde0eb1e)
					πF.SetLineno(116)
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
					πTemp001[8] = πg.NewInt(66).ToObject()
					πTemp001[9] = πg.NewLongFromBytes([]byte{0xea,0xda,0x7d,0xd6,0xcd,0xe0,0xeb,0x1e,}).ToObject()
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
					// line 117: ss[0], ss[4] = RND(ss[5],ss[6],ss[7],ss[0],ss[1],ss[2],ss[3],ss[4],67,0xf57d4f7fee6ed178)
					πF.SetLineno(117)
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
					πTemp001[8] = πg.NewInt(67).ToObject()
					πTemp001[9] = πg.NewLongFromBytes([]byte{0xf5,0x7d,0x4f,0x7f,0xee,0x6e,0xd1,0x78,}).ToObject()
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
					// line 118: ss[7], ss[3] = RND(ss[4],ss[5],ss[6],ss[7],ss[0],ss[1],ss[2],ss[3],68,0x06f067aa72176fba)
					πF.SetLineno(118)
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
					πTemp001[8] = πg.NewInt(68).ToObject()
					πTemp001[9] = πg.NewInt(500013540394364858).ToObject()
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
					// line 119: ss[6], ss[2] = RND(ss[3],ss[4],ss[5],ss[6],ss[7],ss[0],ss[1],ss[2],69,0x0a637dc5a2c898a6)
					πF.SetLineno(119)
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
					πTemp001[8] = πg.NewInt(69).ToObject()
					πTemp001[9] = πg.NewInt(748580250866718886).ToObject()
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
					// line 120: ss[5], ss[1] = RND(ss[2],ss[3],ss[4],ss[5],ss[6],ss[7],ss[0],ss[1],70,0x113f9804bef90dae)
					πF.SetLineno(120)
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
					πTemp001[8] = πg.NewInt(70).ToObject()
					πTemp001[9] = πg.NewInt(1242879168328830382).ToObject()
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
					// line 121: ss[4], ss[0] = RND(ss[1],ss[2],ss[3],ss[4],ss[5],ss[6],ss[7],ss[0],71,0x1b710b35131c471b)
					πF.SetLineno(121)
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
					πTemp001[8] = πg.NewInt(71).ToObject()
					πTemp001[9] = πg.NewInt(1977374033974150939).ToObject()
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
					// line 122: ss[3], ss[7] = RND(ss[0],ss[1],ss[2],ss[3],ss[4],ss[5],ss[6],ss[7],72,0x28db77f523047d84)
					πF.SetLineno(122)
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
					πTemp001[8] = πg.NewInt(72).ToObject()
					πTemp001[9] = πg.NewInt(2944078676154940804).ToObject()
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
					// line 123: ss[2], ss[6] = RND(ss[7],ss[0],ss[1],ss[2],ss[3],ss[4],ss[5],ss[6],73,0x32caab7b40c72493)
					πF.SetLineno(123)
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
					πTemp001[8] = πg.NewInt(73).ToObject()
					πTemp001[9] = πg.NewInt(3659926193048069267).ToObject()
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
					// line 124: ss[1], ss[5] = RND(ss[6],ss[7],ss[0],ss[1],ss[2],ss[3],ss[4],ss[5],74,0x3c9ebe0a15c9bebc)
					πF.SetLineno(124)
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
					πTemp001[8] = πg.NewInt(74).ToObject()
					πTemp001[9] = πg.NewInt(4368137639120453308).ToObject()
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
					// line 125: ss[0], ss[4] = RND(ss[5],ss[6],ss[7],ss[0],ss[1],ss[2],ss[3],ss[4],75,0x431d67c49c100d4c)
					πF.SetLineno(125)
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
					πTemp001[8] = πg.NewInt(75).ToObject()
					πTemp001[9] = πg.NewInt(4836135668995329356).ToObject()
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
					// line 126: ss[7], ss[3] = RND(ss[4],ss[5],ss[6],ss[7],ss[0],ss[1],ss[2],ss[3],76,0x4cc5d4becb3e42b6)
					πF.SetLineno(126)
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
					πTemp001[8] = πg.NewInt(76).ToObject()
					πTemp001[9] = πg.NewInt(5532061633213252278).ToObject()
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
					// line 127: ss[6], ss[2] = RND(ss[3],ss[4],ss[5],ss[6],ss[7],ss[0],ss[1],ss[2],77,0x597f299cfc657e2a)
					πF.SetLineno(127)
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
					πTemp001[8] = πg.NewInt(77).ToObject()
					πTemp001[9] = πg.NewInt(6448918945643986474).ToObject()
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
					// line 128: ss[5], ss[1] = RND(ss[2],ss[3],ss[4],ss[5],ss[6],ss[7],ss[0],ss[1],78,0x5fcb6fab3ad6faec)
					πF.SetLineno(128)
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
					πTemp001[8] = πg.NewInt(78).ToObject()
					πTemp001[9] = πg.NewInt(6902733635092675308).ToObject()
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
					// line 129: ss[4], ss[0] = RND(ss[1],ss[2],ss[3],ss[4],ss[5],ss[6],ss[7],ss[0],79,0x6c44198c4a475817)
					πF.SetLineno(129)
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
					πTemp001[8] = πg.NewInt(79).ToObject()
					πTemp001[9] = πg.NewInt(7801388544844847127).ToObject()
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
					// line 131: dig = []
					πF.SetLineno(131)
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
					// line 133: dig.append( (x + ss[i]) & 0xffffffffffffffff )
					πF.SetLineno(133)
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
					if πTemp004, πE = πg.And(πF, πTemp007, πg.NewLongFromBytes([]byte{0xff,0xff,0xff,0xff,0xff,0xff,0xff,0xff,}).ToObject()); πE != nil {
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
					// line 134: sha_info['digest'] = dig
					πF.SetLineno(134)
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
			// line 136: def sha_init():
			πF.SetLineno(136)
			πTemp003 = make([]πg.Param, 0)
			πTemp005 = πg.NewFunction(πg.NewCode("sha_init", "/usr/lib/python2.7/_sha512.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 137: sha_info = new_shaobject()
					πF.SetLineno(137)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßnew_shaobject); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					µsha_info = πTemp002
					// line 138: sha_info['digest'] = [ 0x6a09e667f3bcc908, 0xbb67ae8584caa73b, 0x3c6ef372fe94f82b, 0xa54ff53a5f1d36f1, 0x510e527fade682d1, 0x9b05688c2b3e6c1f, 0x1f83d9abfb41bd6b, 0x5be0cd19137e2179]
					πF.SetLineno(138)
					πTemp003 = make([]*πg.Object, 8)
					πTemp003[0] = πg.NewInt(7640891576956012808).ToObject()
					πTemp003[1] = πg.NewLongFromBytes([]byte{0xbb,0x67,0xae,0x85,0x84,0xca,0xa7,0x3b,}).ToObject()
					πTemp003[2] = πg.NewInt(4354685564936845355).ToObject()
					πTemp003[3] = πg.NewLongFromBytes([]byte{0xa5,0x4f,0xf5,0x3a,0x5f,0x1d,0x36,0xf1,}).ToObject()
					πTemp003[4] = πg.NewInt(5840696475078001361).ToObject()
					πTemp003[5] = πg.NewLongFromBytes([]byte{0x9b,0x5,0x68,0x8c,0x2b,0x3e,0x6c,0x1f,}).ToObject()
					πTemp003[6] = πg.NewInt(2270897969802886507).ToObject()
					πTemp003[7] = πg.NewInt(6620516959819538809).ToObject()
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
					// line 139: sha_info['count_lo'] = 0
					πF.SetLineno(139)
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
					// line 140: sha_info['count_hi'] = 0
					πF.SetLineno(140)
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
					// line 141: sha_info['local'] = 0
					πF.SetLineno(141)
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
					// line 142: sha_info['digestsize'] = 64
					πF.SetLineno(142)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(64).ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsha_info, "sha_info"); πE != nil {
						continue
					}
					πTemp002 = ßdigestsize.ToObject()
					if πE = πg.SetItem(πF, µsha_info, πTemp002, πTemp001); πE != nil {
						continue
					}
					// line 143: return sha_info
					πF.SetLineno(143)
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
			// line 145: def sha384_init():
			πF.SetLineno(145)
			πTemp003 = make([]πg.Param, 0)
			πTemp006 = πg.NewFunction(πg.NewCode("sha384_init", "/usr/lib/python2.7/_sha512.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 146: sha_info = new_shaobject()
					πF.SetLineno(146)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßnew_shaobject); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					µsha_info = πTemp002
					// line 147: sha_info['digest'] = [ 0xcbbb9d5dc1059ed8, 0x629a292a367cd507, 0x9159015a3070dd17, 0x152fecd8f70e5939, 0x67332667ffc00b31, 0x8eb44a8768581511, 0xdb0c2e0d64f98fa7, 0x47b5481dbefa4fa4]
					πF.SetLineno(147)
					πTemp003 = make([]*πg.Object, 8)
					πTemp003[0] = πg.NewLongFromBytes([]byte{0xcb,0xbb,0x9d,0x5d,0xc1,0x5,0x9e,0xd8,}).ToObject()
					πTemp003[1] = πg.NewInt(7105036623409894663).ToObject()
					πTemp003[2] = πg.NewLongFromBytes([]byte{0x91,0x59,0x1,0x5a,0x30,0x70,0xdd,0x17,}).ToObject()
					πTemp003[3] = πg.NewInt(1526699215303891257).ToObject()
					πTemp003[4] = πg.NewInt(7436329637833083697).ToObject()
					πTemp003[5] = πg.NewLongFromBytes([]byte{0x8e,0xb4,0x4a,0x87,0x68,0x58,0x15,0x11,}).ToObject()
					πTemp003[6] = πg.NewLongFromBytes([]byte{0xdb,0xc,0x2e,0xd,0x64,0xf9,0x8f,0xa7,}).ToObject()
					πTemp003[7] = πg.NewInt(5167115440072839076).ToObject()
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
					// line 148: sha_info['count_lo'] = 0
					πF.SetLineno(148)
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
					// line 149: sha_info['count_hi'] = 0
					πF.SetLineno(149)
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
					// line 150: sha_info['local'] = 0
					πF.SetLineno(150)
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
					// line 151: sha_info['digestsize'] = 48
					πF.SetLineno(151)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(48).ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsha_info, "sha_info"); πE != nil {
						continue
					}
					πTemp002 = ßdigestsize.ToObject()
					if πE = πg.SetItem(πF, µsha_info, πTemp002, πTemp001); πE != nil {
						continue
					}
					// line 152: return sha_info
					πF.SetLineno(152)
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
			if πE = πF.Globals().SetItem(πF, ßsha384_init.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 154: def getbuf(s):
			πF.SetLineno(154)
			πTemp003 = make([]πg.Param, 1)
			πTemp003[0] = πg.Param{Name: "s", Def: nil}
			πTemp007 = πg.NewFunction(πg.NewCode("getbuf", "/usr/lib/python2.7/_sha512.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 155: if isinstance(s, str):
					πF.SetLineno(155)
				Label1:
					// line 156: return s
					πF.SetLineno(156)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πR = µs
					continue
					goto Label4
					// line 157: elif isinstance(s, unicode):
					πF.SetLineno(157)
				Label2:
					// line 158: return str(s)
					πF.SetLineno(158)
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
					// line 160: return buffer(s)
					πF.SetLineno(160)
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
			// line 162: def sha_update(sha_info, buffer):
			πF.SetLineno(162)
			πTemp003 = make([]πg.Param, 2)
			πTemp003[0] = πg.Param{Name: "sha_info", Def: nil}
			πTemp003[1] = πg.Param{Name: "buffer", Def: nil}
			πTemp008 = πg.NewFunction(πg.NewCode("sha_update", "/usr/lib/python2.7/_sha512.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 163: count = len(buffer)
					πF.SetLineno(163)
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
					// line 164: buffer_idx = 0
					πF.SetLineno(164)
					µbuffer_idx = πg.NewInt(0).ToObject()
					// line 165: clo = (sha_info['count_lo'] + (count << 3)) & 0xffffffff
					πF.SetLineno(165)
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
					// line 166: if clo < sha_info['count_lo']:
					πF.SetLineno(166)
				Label1:
					// line 167: sha_info['count_hi'] += 1
					πF.SetLineno(167)
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
					// line 168: sha_info['count_lo'] = clo
					πF.SetLineno(168)
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
					// line 170: sha_info['count_hi'] += (count >> 29)
					πF.SetLineno(170)
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
					// line 172: if sha_info['local']:
					πF.SetLineno(172)
				Label3:
					// line 173: i = SHA_BLOCKSIZE - sha_info['local']
					πF.SetLineno(173)
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
					// line 174: if i > count:
					πF.SetLineno(174)
				Label5:
					// line 175: i = count
					πF.SetLineno(175)
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
					// line 179: sha_info['data'][sha_info['local']+x[0]] = struct.unpack('B', x[1])[0]
					πF.SetLineno(179)
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
					// line 181: count -= i
					πF.SetLineno(181)
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
					// line 182: buffer_idx += i
					πF.SetLineno(182)
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
					// line 184: sha_info['local'] += i
					πF.SetLineno(184)
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
					// line 185: if sha_info['local'] == SHA_BLOCKSIZE:
					πF.SetLineno(185)
				Label10:
					// line 186: sha_transform(sha_info)
					πF.SetLineno(186)
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
					// line 187: sha_info['local'] = 0
					πF.SetLineno(187)
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
					// line 189: return
					πF.SetLineno(189)
					πR = πg.None
					continue
					goto Label12
				Label12:
					goto Label4
				Label4:
					// line 191: while count >= SHA_BLOCKSIZE:
					πF.SetLineno(191)
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
					// line 193: sha_info['data'] = [struct.unpack('B',c)[0] for c in buffer[buffer_idx:buffer_idx + SHA_BLOCKSIZE]]
					πF.SetLineno(193)
					πTemp013 = make([]πg.Param, 0)
					πTemp003 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/_sha512.py", πTemp013, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
								// line 193: sha_info['data'] = [struct.unpack('B',c)[0] for c in buffer[buffer_idx:buffer_idx + SHA_BLOCKSIZE]]
								πF.SetLineno(193)
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
					// line 194: count -= SHA_BLOCKSIZE
					πF.SetLineno(194)
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
					// line 195: buffer_idx += SHA_BLOCKSIZE
					πF.SetLineno(195)
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
					// line 196: sha_transform(sha_info)
					πF.SetLineno(196)
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
					// line 199: pos = sha_info['local']
					πF.SetLineno(199)
					πTemp002 = ßlocal.ToObject()
					if πE = πg.CheckLocal(πF, µsha_info, "sha_info"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µsha_info, πTemp002); πE != nil {
						continue
					}
					µpos = πTemp004
					// line 200: sha_info['data'][pos:pos+count] = [struct.unpack('B',c)[0] for c in buffer[buffer_idx:buffer_idx + count]]
					πF.SetLineno(200)
					πTemp013 = make([]πg.Param, 0)
					πTemp004 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/_sha512.py", πTemp013, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
								// line 200: sha_info['data'][pos:pos+count] = [struct.unpack('B',c)[0] for c in buffer[buffer_idx:buffer_idx + count]]
								πF.SetLineno(200)
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
					// line 201: sha_info['local'] = count
					πF.SetLineno(201)
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
			// line 203: def sha_final(sha_info):
			πF.SetLineno(203)
			πTemp003 = make([]πg.Param, 1)
			πTemp003[0] = πg.Param{Name: "sha_info", Def: nil}
			πTemp009 = πg.NewFunction(πg.NewCode("sha_final", "/usr/lib/python2.7/_sha512.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 204: lo_bit_count = sha_info['count_lo']
					πF.SetLineno(204)
					πTemp001 = ßcount_lo.ToObject()
					if πE = πg.CheckLocal(πF, µsha_info, "sha_info"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetItem(πF, µsha_info, πTemp001); πE != nil {
						continue
					}
					µlo_bit_count = πTemp002
					// line 205: hi_bit_count = sha_info['count_hi']
					πF.SetLineno(205)
					πTemp001 = ßcount_hi.ToObject()
					if πE = πg.CheckLocal(πF, µsha_info, "sha_info"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetItem(πF, µsha_info, πTemp001); πE != nil {
						continue
					}
					µhi_bit_count = πTemp002
					// line 206: count = (lo_bit_count >> 3) & 0x7f
					πF.SetLineno(206)
					if πE = πg.CheckLocal(πF, µlo_bit_count, "lo_bit_count"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.RShift(πF, µlo_bit_count, πg.NewInt(3).ToObject()); πE != nil {
						continue
					}
					if πTemp001, πE = πg.And(πF, πTemp002, πg.NewInt(127).ToObject()); πE != nil {
						continue
					}
					µcount = πTemp001
					// line 207: sha_info['data'][count] = 0x80;
					πF.SetLineno(207)
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
					// line 208: count += 1
					πF.SetLineno(208)
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
					if πTemp002, πE = πg.Sub(πF, πTemp003, πg.NewInt(16).ToObject()); πE != nil {
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
					// line 209: if count > SHA_BLOCKSIZE - 16:
					πF.SetLineno(209)
				Label1:
					// line 211: sha_info['data'] = sha_info['data'][:count] + ([0] * (SHA_BLOCKSIZE - count))
					πF.SetLineno(211)
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
					// line 212: sha_transform(sha_info)
					πF.SetLineno(212)
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
					// line 214: sha_info['data'] = [0] * SHA_BLOCKSIZE
					πF.SetLineno(214)
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
					// line 216: sha_info['data'] = sha_info['data'][:count] + ([0] * (SHA_BLOCKSIZE - count))
					πF.SetLineno(216)
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
					// line 218: sha_info['data'][112] = 0;
					πF.SetLineno(218)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					πTemp002 = ßdata.ToObject()
					if πE = πg.CheckLocal(πF, µsha_info, "sha_info"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µsha_info, πTemp002); πE != nil {
						continue
					}
					πTemp002 = πg.NewInt(112).ToObject()
					if πE = πg.SetItem(πF, πTemp003, πTemp002, πTemp001); πE != nil {
						continue
					}
					// line 219: sha_info['data'][113] = 0;
					πF.SetLineno(219)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					πTemp002 = ßdata.ToObject()
					if πE = πg.CheckLocal(πF, µsha_info, "sha_info"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µsha_info, πTemp002); πE != nil {
						continue
					}
					πTemp002 = πg.NewInt(113).ToObject()
					if πE = πg.SetItem(πF, πTemp003, πTemp002, πTemp001); πE != nil {
						continue
					}
					// line 220: sha_info['data'][114] = 0;
					πF.SetLineno(220)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					πTemp002 = ßdata.ToObject()
					if πE = πg.CheckLocal(πF, µsha_info, "sha_info"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µsha_info, πTemp002); πE != nil {
						continue
					}
					πTemp002 = πg.NewInt(114).ToObject()
					if πE = πg.SetItem(πF, πTemp003, πTemp002, πTemp001); πE != nil {
						continue
					}
					// line 221: sha_info['data'][115] = 0;
					πF.SetLineno(221)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					πTemp002 = ßdata.ToObject()
					if πE = πg.CheckLocal(πF, µsha_info, "sha_info"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µsha_info, πTemp002); πE != nil {
						continue
					}
					πTemp002 = πg.NewInt(115).ToObject()
					if πE = πg.SetItem(πF, πTemp003, πTemp002, πTemp001); πE != nil {
						continue
					}
					// line 222: sha_info['data'][116] = 0;
					πF.SetLineno(222)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					πTemp002 = ßdata.ToObject()
					if πE = πg.CheckLocal(πF, µsha_info, "sha_info"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µsha_info, πTemp002); πE != nil {
						continue
					}
					πTemp002 = πg.NewInt(116).ToObject()
					if πE = πg.SetItem(πF, πTemp003, πTemp002, πTemp001); πE != nil {
						continue
					}
					// line 223: sha_info['data'][117] = 0;
					πF.SetLineno(223)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					πTemp002 = ßdata.ToObject()
					if πE = πg.CheckLocal(πF, µsha_info, "sha_info"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µsha_info, πTemp002); πE != nil {
						continue
					}
					πTemp002 = πg.NewInt(117).ToObject()
					if πE = πg.SetItem(πF, πTemp003, πTemp002, πTemp001); πE != nil {
						continue
					}
					// line 224: sha_info['data'][118] = 0;
					πF.SetLineno(224)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					πTemp002 = ßdata.ToObject()
					if πE = πg.CheckLocal(πF, µsha_info, "sha_info"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µsha_info, πTemp002); πE != nil {
						continue
					}
					πTemp002 = πg.NewInt(118).ToObject()
					if πE = πg.SetItem(πF, πTemp003, πTemp002, πTemp001); πE != nil {
						continue
					}
					// line 225: sha_info['data'][119] = 0;
					πF.SetLineno(225)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					πTemp002 = ßdata.ToObject()
					if πE = πg.CheckLocal(πF, µsha_info, "sha_info"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µsha_info, πTemp002); πE != nil {
						continue
					}
					πTemp002 = πg.NewInt(119).ToObject()
					if πE = πg.SetItem(πF, πTemp003, πTemp002, πTemp001); πE != nil {
						continue
					}
					// line 227: sha_info['data'][120] = (hi_bit_count >> 24) & 0xff
					πF.SetLineno(227)
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
					πTemp003 = πg.NewInt(120).ToObject()
					if πE = πg.SetItem(πF, πTemp005, πTemp003, πTemp002); πE != nil {
						continue
					}
					// line 228: sha_info['data'][121] = (hi_bit_count >> 16) & 0xff
					πF.SetLineno(228)
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
					πTemp003 = πg.NewInt(121).ToObject()
					if πE = πg.SetItem(πF, πTemp005, πTemp003, πTemp002); πE != nil {
						continue
					}
					// line 229: sha_info['data'][122] = (hi_bit_count >>  8) & 0xff
					πF.SetLineno(229)
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
					πTemp003 = πg.NewInt(122).ToObject()
					if πE = πg.SetItem(πF, πTemp005, πTemp003, πTemp002); πE != nil {
						continue
					}
					// line 230: sha_info['data'][123] = (hi_bit_count >>  0) & 0xff
					πF.SetLineno(230)
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
					πTemp003 = πg.NewInt(123).ToObject()
					if πE = πg.SetItem(πF, πTemp005, πTemp003, πTemp002); πE != nil {
						continue
					}
					// line 231: sha_info['data'][124] = (lo_bit_count >> 24) & 0xff
					πF.SetLineno(231)
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
					πTemp003 = πg.NewInt(124).ToObject()
					if πE = πg.SetItem(πF, πTemp005, πTemp003, πTemp002); πE != nil {
						continue
					}
					// line 232: sha_info['data'][125] = (lo_bit_count >> 16) & 0xff
					πF.SetLineno(232)
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
					πTemp003 = πg.NewInt(125).ToObject()
					if πE = πg.SetItem(πF, πTemp005, πTemp003, πTemp002); πE != nil {
						continue
					}
					// line 233: sha_info['data'][126] = (lo_bit_count >>  8) & 0xff
					πF.SetLineno(233)
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
					πTemp003 = πg.NewInt(126).ToObject()
					if πE = πg.SetItem(πF, πTemp005, πTemp003, πTemp002); πE != nil {
						continue
					}
					// line 234: sha_info['data'][127] = (lo_bit_count >>  0) & 0xff
					πF.SetLineno(234)
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
					πTemp003 = πg.NewInt(127).ToObject()
					if πE = πg.SetItem(πF, πTemp005, πTemp003, πTemp002); πE != nil {
						continue
					}
					// line 236: sha_transform(sha_info)
					πF.SetLineno(236)
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
					// line 238: dig = []
					πF.SetLineno(238)
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
					// line 240: dig.extend([ ((i>>56) & 0xff), ((i>>48) & 0xff), ((i>>40) & 0xff), ((i>>32) & 0xff), ((i>>24) & 0xff), ((i>>16) & 0xff), ((i>>8) & 0xff), (i & 0xff) ])
					πF.SetLineno(240)
					πTemp007 = πF.MakeArgs(1)
					πTemp010 = make([]*πg.Object, 8)
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.RShift(πF, µi, πg.NewInt(56).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.And(πF, πTemp003, πg.NewInt(255).ToObject()); πE != nil {
						continue
					}
					πTemp010[0] = πTemp002
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.RShift(πF, µi, πg.NewInt(48).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.And(πF, πTemp003, πg.NewInt(255).ToObject()); πE != nil {
						continue
					}
					πTemp010[1] = πTemp002
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.RShift(πF, µi, πg.NewInt(40).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.And(πF, πTemp003, πg.NewInt(255).ToObject()); πE != nil {
						continue
					}
					πTemp010[2] = πTemp002
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.RShift(πF, µi, πg.NewInt(32).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.And(πF, πTemp003, πg.NewInt(255).ToObject()); πE != nil {
						continue
					}
					πTemp010[3] = πTemp002
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.RShift(πF, µi, πg.NewInt(24).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.And(πF, πTemp003, πg.NewInt(255).ToObject()); πE != nil {
						continue
					}
					πTemp010[4] = πTemp002
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.RShift(πF, µi, πg.NewInt(16).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.And(πF, πTemp003, πg.NewInt(255).ToObject()); πE != nil {
						continue
					}
					πTemp010[5] = πTemp002
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.RShift(πF, µi, πg.NewInt(8).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.And(πF, πTemp003, πg.NewInt(255).ToObject()); πE != nil {
						continue
					}
					πTemp010[6] = πTemp002
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.And(πF, µi, πg.NewInt(255).ToObject()); πE != nil {
						continue
					}
					πTemp010[7] = πTemp002
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
					// line 241: return ''.join([chr(i) for i in dig])
					πF.SetLineno(241)
					πTemp007 = πF.MakeArgs(1)
					πTemp011 = make([]πg.Param, 0)
					πTemp002 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/_sha512.py", πTemp011, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
								// line 241: return ''.join([chr(i) for i in dig])
								πF.SetLineno(241)
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
			// line 243: class sha512(object):
			πF.SetLineno(243)
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
			_, πE = πg.NewCode("sha512", "/usr/lib/python2.7/_sha512.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 244: digest_size = digestsize = SHA_DIGESTSIZE
					πF.SetLineno(244)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßSHA_DIGESTSIZE); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßdigest_size.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßdigestsize.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 245: block_size = SHA_BLOCKSIZE
					πF.SetLineno(245)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßSHA_BLOCKSIZE); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßblock_size.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 247: def __init__(self, s=None):
					πF.SetLineno(247)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[1] = πg.Param{Name: "s", Def: πTemp003}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/_sha512.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 248: self._sha = sha_init()
							πF.SetLineno(248)
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
							// line 249: if s:
							πF.SetLineno(249)
						Label1:
							// line 250: sha_update(self._sha, getbuf(s))
							πF.SetLineno(250)
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
					// line 252: def update(self, s):
					πF.SetLineno(252)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "s", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("update", "/usr/lib/python2.7/_sha512.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 253: sha_update(self._sha, getbuf(s))
							πF.SetLineno(253)
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
					// line 255: def digest(self):
					πF.SetLineno(255)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("digest", "/usr/lib/python2.7/_sha512.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 256: return sha_final(self._sha.copy())[:self._sha['digestsize']]
							πF.SetLineno(256)
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
					// line 258: def hexdigest(self):
					πF.SetLineno(258)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("hexdigest", "/usr/lib/python2.7/_sha512.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 259: return ''.join([('0%x' % ord(i))[-2:] for i in self.digest()])
							πF.SetLineno(259)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = make([]πg.Param, 0)
							πTemp003 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/_sha512.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
										// line 259: return ''.join([('0%x' % ord(i))[-2:] for i in self.digest()])
										πF.SetLineno(259)
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
					// line 261: def copy(self):
					πF.SetLineno(261)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("copy", "/usr/lib/python2.7/_sha512.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 262: new = sha512.__new__(sha512)
							πF.SetLineno(262)
							πTemp001 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßsha512); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßsha512); πE != nil {
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
							// line 263: new._sha = self._sha.copy()
							πF.SetLineno(263)
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
							// line 264: return new
							πF.SetLineno(264)
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
			if πTemp013, πE = πTemp012.Call(πF, []*πg.Object{πg.NewStr("sha512").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp010.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßsha512.ToObject(), πTemp013); πE != nil {
				continue
			}
			// line 266: class sha384(sha512):
			πF.SetLineno(266)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp013, πE = πg.ResolveGlobal(πF, ßsha512); πE != nil {
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
			_, πE = πg.NewCode("sha384", "/usr/lib/python2.7/_sha512.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 267: digest_size = digestsize = 48
					πF.SetLineno(267)
					if πE = πClass.SetItem(πF, ßdigest_size.ToObject(), πg.NewInt(48).ToObject()); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßdigestsize.ToObject(), πg.NewInt(48).ToObject()); πE != nil {
						continue
					}
					// line 269: def __init__(self, s=None):
					πF.SetLineno(269)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[1] = πg.Param{Name: "s", Def: πTemp003}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/_sha512.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 270: self._sha = sha384_init()
							πF.SetLineno(270)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßsha384_init); πE != nil {
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
							// line 271: if s:
							πF.SetLineno(271)
						Label1:
							// line 272: sha_update(self._sha, getbuf(s))
							πF.SetLineno(272)
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
					// line 274: def copy(self):
					πF.SetLineno(274)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("copy", "/usr/lib/python2.7/_sha512.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 275: new = sha384.__new__(sha384)
							πF.SetLineno(275)
							πTemp001 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßsha384); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßsha384); πE != nil {
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
							// line 276: new._sha = self._sha.copy()
							πF.SetLineno(276)
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
							// line 277: return new
							πF.SetLineno(277)
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
			if πTemp013, πE = πTemp012.Call(πF, []*πg.Object{πg.NewStr("sha384").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp010.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßsha384.ToObject(), πTemp013); πE != nil {
				continue
			}
			// line 279: def test():
			πF.SetLineno(279)
			πTemp003 = make([]πg.Param, 0)
			πTemp011 = πg.NewFunction(πg.NewCode("test", "/usr/lib/python2.7/_sha512.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µa_str *πg.Object = πg.UnboundLocal; _ = µa_str
				var µs *πg.Object = πg.UnboundLocal; _ = µs
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
					// line 282: a_str = "just a test string"
					πF.SetLineno(282)
					µa_str = πg.NewStr("just a test string").ToObject()
					// line 284: assert sha512().hexdigest() == sha512().hexdigest()
					πF.SetLineno(284)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßsha512); πE != nil {
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
					if πTemp002, πE = πg.ResolveGlobal(πF, ßsha512); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp004, ßhexdigest, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, πTemp003, πTemp004); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
						continue
					}
					// line 285: assert sha512(a_str).hexdigest() == sha512(a_str).hexdigest()
					πF.SetLineno(285)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µa_str, "a_str"); πE != nil {
						continue
					}
					πTemp005[0] = µa_str
					if πTemp002, πE = πg.ResolveGlobal(πF, ßsha512); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßhexdigest, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µa_str, "a_str"); πE != nil {
						continue
					}
					πTemp005[0] = µa_str
					if πTemp002, πE = πg.ResolveGlobal(πF, ßsha512); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πTemp002, πE = πg.GetAttr(πF, πTemp004, ßhexdigest, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, πTemp003, πTemp004); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
						continue
					}
					// line 286: assert sha512(a_str*7).hexdigest() == sha512(a_str*7).hexdigest()
					πF.SetLineno(286)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µa_str, "a_str"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mul(πF, µa_str, πg.NewInt(7).ToObject()); πE != nil {
						continue
					}
					πTemp005[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßsha512); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßhexdigest, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µa_str, "a_str"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mul(πF, µa_str, πg.NewInt(7).ToObject()); πE != nil {
						continue
					}
					πTemp005[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßsha512); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πTemp002, πE = πg.GetAttr(πF, πTemp004, ßhexdigest, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, πTemp003, πTemp004); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
						continue
					}
					// line 288: s = sha512(a_str)
					πF.SetLineno(288)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µa_str, "a_str"); πE != nil {
						continue
					}
					πTemp005[0] = µa_str
					if πTemp001, πE = πg.ResolveGlobal(πF, ßsha512); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					µs = πTemp002
					// line 289: s.update(a_str)
					πF.SetLineno(289)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µa_str, "a_str"); πE != nil {
						continue
					}
					πTemp005[0] = µa_str
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µs, ßupdate, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					// line 290: assert sha512(a_str+a_str).hexdigest() == s.hexdigest()
					πF.SetLineno(290)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µa_str, "a_str"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µa_str, "a_str"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Add(πF, µa_str, µa_str); πE != nil {
						continue
					}
					πTemp005[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßsha512); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßhexdigest, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µs, ßhexdigest, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, πTemp003, πTemp004); πE != nil {
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
			// line 292: if __name__ == "__main__":
			πF.SetLineno(292)
		Label1:
			// line 293: test()
			πF.SetLineno(293)
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
	πg.RegisterModule("_sha512", Code)
}

