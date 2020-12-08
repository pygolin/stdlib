package binascii
import (
	πg "github.com/pygolin/runtime"
)
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/binascii.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ß := πg.InternStr("")
		ß0 := πg.InternStr("0")
		ß0123456789ABCDEF := πg.InternStr("0123456789ABCDEF")
		ß1 := πg.InternStr("1")
		ß2 := πg.InternStr("2")
		ß3 := πg.InternStr("3")
		ß4 := πg.InternStr("4")
		ß5 := πg.InternStr("5")
		ß6 := πg.InternStr("6")
		ß7 := πg.InternStr("7")
		ß8 := πg.InternStr("8")
		ß9 := πg.InternStr("9")
		ßA := πg.InternStr("A")
		ßB := πg.InternStr("B")
		ßC := πg.InternStr("C")
		ßD := πg.InternStr("D")
		ßDONE := πg.InternStr("DONE")
		ßDone := πg.InternStr("Done")
		ßE := πg.InternStr("E")
		ßError := πg.InternStr("Error")
		ßException := πg.InternStr("Exception")
		ßF := πg.InternStr("F")
		ßFAIL := πg.InternStr("FAIL")
		ßFalse := πg.InternStr("False")
		ßG := πg.InternStr("G")
		ßH := πg.InternStr("H")
		ßI := πg.InternStr("I")
		ßIncomplete := πg.InternStr("Incomplete")
		ßIndexError := πg.InternStr("IndexError")
		ßJ := πg.InternStr("J")
		ßK := πg.InternStr("K")
		ßKeyError := πg.InternStr("KeyError")
		ßL := πg.InternStr("L")
		ßM := πg.InternStr("M")
		ßN := πg.InternStr("N")
		ßNone := πg.InternStr("None")
		ßO := πg.InternStr("O")
		ßP := πg.InternStr("P")
		ßQ := πg.InternStr("Q")
		ßR := πg.InternStr("R")
		ßS := πg.InternStr("S")
		ßSKIP := πg.InternStr("SKIP")
		ßT := πg.InternStr("T")
		ßTrue := πg.InternStr("True")
		ßTypeError := πg.InternStr("TypeError")
		ßU := πg.InternStr("U")
		ßV := πg.InternStr("V")
		ßValueError := πg.InternStr("ValueError")
		ßW := πg.InternStr("W")
		ßX := πg.InternStr("X")
		ßY := πg.InternStr("Y")
		ßZ := πg.InternStr("Z")
		ß_ := πg.InternStr("_")
		ß__doc__ := πg.InternStr("__doc__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ßa := πg.InternStr("a")
		ßa2b_base64 := πg.InternStr("a2b_base64")
		ßa2b_hex := πg.InternStr("a2b_hex")
		ßa2b_hqx := πg.InternStr("a2b_hqx")
		ßa2b_qp := πg.InternStr("a2b_qp")
		ßa2b_uu := πg.InternStr("a2b_uu")
		ßappend := πg.InternStr("append")
		ßb := πg.InternStr("b")
		ßb2a_base64 := πg.InternStr("b2a_base64")
		ßb2a_hex := πg.InternStr("b2a_hex")
		ßb2a_hqx := πg.InternStr("b2a_hqx")
		ßb2a_qp := πg.InternStr("b2a_qp")
		ßb2a_uu := πg.InternStr("b2a_uu")
		ßc := πg.InternStr("c")
		ßchr := πg.InternStr("chr")
		ßcrc32 := πg.InternStr("crc32")
		ßcrc_32_tab := πg.InternStr("crc_32_tab")
		ßcrc_hqx := πg.InternStr("crc_hqx")
		ßcrctab_hqx := πg.InternStr("crctab_hqx")
		ßd := πg.InternStr("d")
		ße := πg.InternStr("e")
		ßenumerate := πg.InternStr("enumerate")
		ßf := πg.InternStr("f")
		ßfind := πg.InternStr("find")
		ßg := πg.InternStr("g")
		ßh := πg.InternStr("h")
		ßhex := πg.InternStr("hex")
		ßhex_numbers := πg.InternStr("hex_numbers")
		ßhexlify := πg.InternStr("hexlify")
		ßhqx_encoding := πg.InternStr("hqx_encoding")
		ßi := πg.InternStr("i")
		ßindex := πg.InternStr("index")
		ßint := πg.InternStr("int")
		ßisinstance := πg.InternStr("isinstance")
		ßj := πg.InternStr("j")
		ßjoin := πg.InternStr("join")
		ßk := πg.InternStr("k")
		ßl := πg.InternStr("l")
		ßlen := πg.InternStr("len")
		ßlong := πg.InternStr("long")
		ßm := πg.InternStr("m")
		ßn := πg.InternStr("n")
		ßo := πg.InternStr("o")
		ßord := πg.InternStr("ord")
		ßp := πg.InternStr("p")
		ßq := πg.InternStr("q")
		ßr := πg.InternStr("r")
		ßrange := πg.InternStr("range")
		ßrlecode_hqx := πg.InternStr("rlecode_hqx")
		ßrledecode_hqx := πg.InternStr("rledecode_hqx")
		ßrstrip := πg.InternStr("rstrip")
		ßs := πg.InternStr("s")
		ßsplit := πg.InternStr("split")
		ßstr := πg.InternStr("str")
		ßstrhex_to_int := πg.InternStr("strhex_to_int")
		ßt := πg.InternStr("t")
		ßtable_a2b_base64 := πg.InternStr("table_a2b_base64")
		ßtable_a2b_hqx := πg.InternStr("table_a2b_hqx")
		ßtable_b2a_base64 := πg.InternStr("table_b2a_base64")
		ßtable_hex := πg.InternStr("table_hex")
		ßtuple := πg.InternStr("tuple")
		ßtwo_hex_digits := πg.InternStr("two_hex_digits")
		ßu := πg.InternStr("u")
		ßunhexlify := πg.InternStr("unhexlify")
		ßunicode := πg.InternStr("unicode")
		ßv := πg.InternStr("v")
		ßw := πg.InternStr("w")
		ßx := πg.InternStr("x")
		ßy := πg.InternStr("y")
		ßz := πg.InternStr("z")
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
			// line 1: """A pure Python implementation of binascii.
			πF.SetLineno(1)
			// line 1: """A pure Python implementation of binascii.
			πF.SetLineno(1)
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("A pure Python implementation of binascii.\n\nRather slow and buggy in corner cases.\nPyPy provides an RPython version too.\n").ToObject()); πE != nil {
				continue
			}
			// line 7: class Error(Exception):
			πF.SetLineno(7)
			πTemp003 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßException); πE != nil {
				continue
			}
			πTemp003[0] = πTemp005
			πTemp001 = πg.NewDict()
			if πTemp002, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ß__module__.ToObject(), πTemp002); πE != nil {
				continue
			}
			_, πE = πg.NewCode("Error", "/usr/lib/python2.7/binascii.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp001
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 8: pass
					πF.SetLineno(8)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp004, πE = πTemp001.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp004 == nil {
				πTemp004 = πg.TypeType.ToObject()
			}
			if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("Error").ToObject(), πg.NewTuple(πTemp003...).ToObject(), πTemp001.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßError.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 10: class Done(Exception):
			πF.SetLineno(10)
			πTemp003 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßException); πE != nil {
				continue
			}
			πTemp003[0] = πTemp005
			πTemp001 = πg.NewDict()
			if πTemp002, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ß__module__.ToObject(), πTemp002); πE != nil {
				continue
			}
			_, πE = πg.NewCode("Done", "/usr/lib/python2.7/binascii.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp001
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 11: pass
					πF.SetLineno(11)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp004, πE = πTemp001.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp004 == nil {
				πTemp004 = πg.TypeType.ToObject()
			}
			if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("Done").ToObject(), πg.NewTuple(πTemp003...).ToObject(), πTemp001.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßDone.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 13: class Incomplete(Exception):
			πF.SetLineno(13)
			πTemp003 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßException); πE != nil {
				continue
			}
			πTemp003[0] = πTemp005
			πTemp001 = πg.NewDict()
			if πTemp002, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ß__module__.ToObject(), πTemp002); πE != nil {
				continue
			}
			_, πE = πg.NewCode("Incomplete", "/usr/lib/python2.7/binascii.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp001
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 14: pass
					πF.SetLineno(14)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp004, πE = πTemp001.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp004 == nil {
				πTemp004 = πg.TypeType.ToObject()
			}
			if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("Incomplete").ToObject(), πg.NewTuple(πTemp003...).ToObject(), πTemp001.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßIncomplete.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 16: def a2b_uu(s):
			πF.SetLineno(16)
			πTemp006 = make([]πg.Param, 1)
			πTemp006[0] = πg.Param{Name: "s", Def: nil}
			πTemp002 = πg.NewFunction(πg.NewCode("a2b_uu", "/usr/lib/python2.7/binascii.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µs *πg.Object = πArgs[0]; _ = µs
				var µlength *πg.Object = πg.UnboundLocal; _ = µlength
				var µquadruplets_gen *πg.Object = πg.UnboundLocal; _ = µquadruplets_gen
				var µresult *πg.Object = πg.UnboundLocal; _ = µresult
				var µtrailingdata *πg.Object = πg.UnboundLocal; _ = µtrailingdata
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
				var πTemp007 []πg.Param
				_ = πTemp007
				var πTemp008 *πg.BaseException
				_ = πTemp008
				var πTemp009 *πg.Traceback
				_ = πTemp009
				var πTemp010 *πg.Object
				_ = πTemp010
				var πTemp011 *πg.Object
				_ = πTemp011
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 4: goto Label4
					default: panic("unexpected function state")
					}
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, µs); πE != nil {
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
					// line 17: if not s:
					πF.SetLineno(17)
				Label1:
					// line 18: return ''
					πF.SetLineno(18)
					πR = ß.ToObject()
					continue
					goto Label2
				Label2:
					// line 20: length = (ord(s[0]) - 0x20) % 64
					πF.SetLineno(20)
					πTemp004 = πF.MakeArgs(1)
					πTemp005 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, µs, πTemp005); πE != nil {
						continue
					}
					πTemp004[0] = πTemp006
					if πTemp005, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp003, πE = πg.Sub(πF, πTemp006, πg.NewInt(32).ToObject()); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Mod(πF, πTemp003, πg.NewInt(64).ToObject()); πE != nil {
						continue
					}
					µlength = πTemp001
					// line 22: def quadruplets_gen(s):
					πF.SetLineno(22)
					πTemp007 = make([]πg.Param, 1)
					πTemp007[0] = πg.Param{Name: "s", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("quadruplets_gen", "/usr/lib/python2.7/binascii.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µs *πg.Object = πArgs[0]; _ = µs
						var πTemp001 bool
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
						var πTemp012 *πg.Object
						_ = πTemp012
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
							for ; πF.State() >= 0; πF.PopCheckpoint() {
								switch πF.State() {
								case 0:
								case 8: goto Label8
								case 1: goto Label1
								case 2: goto Label2
								case 5: goto Label5
								case 6: goto Label6
								default: panic("unexpected function state")
								}
								// line 23: while s:
								πF.SetLineno(23)
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
								if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.IsTrue(πF, µs); πE != nil {
									continue
								}
								if πE != nil || !πTemp002 {
									continue
								}
								πF.PushCheckpoint(1)            
								// line 24: try:
								πF.SetLineno(24)
								πF.PushCheckpoint(5)
								// line 25: yield ord(s[0]), ord(s[1]), ord(s[2]), ord(s[3])
								πF.SetLineno(25)
								πTemp004 = πF.MakeArgs(1)
								πTemp005 = πg.NewInt(0).ToObject()
								if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
									continue
								}
								if πTemp006, πE = πg.GetItem(πF, µs, πTemp005); πE != nil {
									continue
								}
								πTemp004[0] = πTemp006
								if πTemp005, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
									continue
								}
								if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								πTemp004 = πF.MakeArgs(1)
								πTemp005 = πg.NewInt(1).ToObject()
								if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
									continue
								}
								if πTemp007, πE = πg.GetItem(πF, µs, πTemp005); πE != nil {
									continue
								}
								πTemp004[0] = πTemp007
								if πTemp005, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
									continue
								}
								if πTemp007, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								πTemp004 = πF.MakeArgs(1)
								πTemp005 = πg.NewInt(2).ToObject()
								if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
									continue
								}
								if πTemp008, πE = πg.GetItem(πF, µs, πTemp005); πE != nil {
									continue
								}
								πTemp004[0] = πTemp008
								if πTemp005, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
									continue
								}
								if πTemp008, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								πTemp004 = πF.MakeArgs(1)
								πTemp005 = πg.NewInt(3).ToObject()
								if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
									continue
								}
								if πTemp009, πE = πg.GetItem(πF, µs, πTemp005); πE != nil {
									continue
								}
								πTemp004[0] = πTemp009
								if πTemp005, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
									continue
								}
								if πTemp009, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								πTemp003 = πg.NewTuple4(πTemp006, πTemp007, πTemp008, πTemp009).ToObject()
								πF.PushCheckpoint(6)
								return πTemp003, nil
							Label6:
								πTemp005 = πSent
								πF.PopCheckpoint()
								goto Label4
							Label5:
								if πE == nil {
								  continue
								}
								πE = nil
								πTemp010, πTemp011 = πF.ExcInfo()
								if πTemp005, πE = πg.ResolveGlobal(πF, ßIndexError); πE != nil {
									continue
								}
								if πTemp002, πE = πg.IsInstance(πF, πTemp010.ToObject(), πTemp005); πE != nil {
									continue
								}
								if πTemp002 {
									goto Label7
								}
								πE = πF.Raise(πTemp010.ToObject(), nil, πTemp011.ToObject())
								continue
								// line 26: except IndexError:
								πF.SetLineno(26)
							Label7:
								// line 27: s += '   '
								πF.SetLineno(27)
								if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
									continue
								}
								if πTemp005, πE = πg.IAdd(πF, µs, πg.NewStr("   ").ToObject()); πE != nil {
									continue
								}
								µs = πTemp005
								// line 28: yield ord(s[0]), ord(s[1]), ord(s[2]), ord(s[3])
								πF.SetLineno(28)
								πTemp004 = πF.MakeArgs(1)
								πTemp006 = πg.NewInt(0).ToObject()
								if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
									continue
								}
								if πTemp007, πE = πg.GetItem(πF, µs, πTemp006); πE != nil {
									continue
								}
								πTemp004[0] = πTemp007
								if πTemp006, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
									continue
								}
								if πTemp007, πE = πTemp006.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								πTemp004 = πF.MakeArgs(1)
								πTemp006 = πg.NewInt(1).ToObject()
								if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
									continue
								}
								if πTemp008, πE = πg.GetItem(πF, µs, πTemp006); πE != nil {
									continue
								}
								πTemp004[0] = πTemp008
								if πTemp006, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
									continue
								}
								if πTemp008, πE = πTemp006.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								πTemp004 = πF.MakeArgs(1)
								πTemp006 = πg.NewInt(2).ToObject()
								if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
									continue
								}
								if πTemp009, πE = πg.GetItem(πF, µs, πTemp006); πE != nil {
									continue
								}
								πTemp004[0] = πTemp009
								if πTemp006, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
									continue
								}
								if πTemp009, πE = πTemp006.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								πTemp004 = πF.MakeArgs(1)
								πTemp006 = πg.NewInt(3).ToObject()
								if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
									continue
								}
								if πTemp012, πE = πg.GetItem(πF, µs, πTemp006); πE != nil {
									continue
								}
								πTemp004[0] = πTemp012
								if πTemp006, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
									continue
								}
								if πTemp012, πE = πTemp006.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								πTemp005 = πg.NewTuple4(πTemp007, πTemp008, πTemp009, πTemp012).ToObject()
								πF.PushCheckpoint(8)
								return πTemp005, nil
							Label8:
								πTemp006 = πSent
								// line 29: return
								πF.SetLineno(29)
								πR = πg.None
								continue
								πF.RestoreExc(nil, nil)
								goto Label4
							Label4:
								// line 30: s = s[4:]
								πF.SetLineno(30)
								if πTemp006, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(4).ToObject(), πg.None, πg.None}, nil); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
									continue
								}
								if πTemp007, πE = πg.GetItem(πF, µs, πTemp006); πE != nil {
									continue
								}
								µs = πTemp007
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
					µquadruplets_gen = πTemp001
					// line 32: try:
					πF.SetLineno(32)
					πF.PushCheckpoint(4)
					// line 33: result = [''.join(
					πF.SetLineno(33)
					πTemp007 = make([]πg.Param, 0)
					πTemp005 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/binascii.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µA *πg.Object = πg.UnboundLocal; _ = µA
						var µB *πg.Object = πg.UnboundLocal; _ = µB
						var µC *πg.Object = πg.UnboundLocal; _ = µC
						var µD *πg.Object = πg.UnboundLocal; _ = µD
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
						var πTemp010 []*πg.Object
						_ = πTemp010
						var πTemp011 []*πg.Object
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
								default: panic("unexpected function state")
								}
								πTemp002 = πF.MakeArgs(1)
								if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
									continue
								}
								if πTemp004, πE = πg.GetItem(πF, µs, πTemp003); πE != nil {
									continue
								}
								if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßrstrip, nil); πE != nil {
									continue
								}
								if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
									continue
								}
								πTemp002[0] = πTemp004
								if πE = πg.CheckLocal(πF, µquadruplets_gen, "quadruplets_gen"); πE != nil {
									continue
								}
								if πTemp003, πE = µquadruplets_gen.Call(πF, πTemp002, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp002)
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
									if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}, πg.TieTarget{Target: &πTemp008}, πg.TieTarget{Target: &πTemp009}}}, πTemp003); πE != nil {
										continue
									}
									µA = πTemp004
									µB = πTemp007
									µC = πTemp008
									µD = πTemp009
								}
								if πE != nil || !πTemp006 {
									continue
								}
								πF.PushCheckpoint(1)            
								// line 33: result = [''.join(
								πF.SetLineno(33)
								πTemp002 = πF.MakeArgs(1)
								πTemp010 = make([]*πg.Object, 3)
								πTemp011 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µA, "A"); πE != nil {
									continue
								}
								if πTemp007, πE = πg.Sub(πF, µA, πg.NewInt(32).ToObject()); πE != nil {
									continue
								}
								if πTemp004, πE = πg.LShift(πF, πTemp007, πg.NewInt(2).ToObject()); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µB, "B"); πE != nil {
									continue
								}
								if πTemp009, πE = πg.Sub(πF, µB, πg.NewInt(32).ToObject()); πE != nil {
									continue
								}
								if πTemp008, πE = πg.RShift(πF, πTemp009, πg.NewInt(4).ToObject()); πE != nil {
									continue
								}
								if πTemp007, πE = πg.And(πF, πTemp008, πg.NewInt(3).ToObject()); πE != nil {
									continue
								}
								if πTemp003, πE = πg.Or(πF, πTemp004, πTemp007); πE != nil {
									continue
								}
								πTemp011[0] = πTemp003
								if πTemp003, πE = πg.ResolveGlobal(πF, ßchr); πE != nil {
									continue
								}
								if πTemp004, πE = πTemp003.Call(πF, πTemp011, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp011)
								πTemp010[0] = πTemp004
								πTemp011 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µB, "B"); πE != nil {
									continue
								}
								if πTemp008, πE = πg.Sub(πF, µB, πg.NewInt(32).ToObject()); πE != nil {
									continue
								}
								if πTemp007, πE = πg.And(πF, πTemp008, πg.NewInt(15).ToObject()); πE != nil {
									continue
								}
								if πTemp004, πE = πg.LShift(πF, πTemp007, πg.NewInt(4).ToObject()); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µC, "C"); πE != nil {
									continue
								}
								if πTemp009, πE = πg.Sub(πF, µC, πg.NewInt(32).ToObject()); πE != nil {
									continue
								}
								if πTemp008, πE = πg.RShift(πF, πTemp009, πg.NewInt(2).ToObject()); πE != nil {
									continue
								}
								if πTemp007, πE = πg.And(πF, πTemp008, πg.NewInt(15).ToObject()); πE != nil {
									continue
								}
								if πTemp003, πE = πg.Or(πF, πTemp004, πTemp007); πE != nil {
									continue
								}
								πTemp011[0] = πTemp003
								if πTemp003, πE = πg.ResolveGlobal(πF, ßchr); πE != nil {
									continue
								}
								if πTemp004, πE = πTemp003.Call(πF, πTemp011, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp011)
								πTemp010[1] = πTemp004
								πTemp011 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µC, "C"); πE != nil {
									continue
								}
								if πTemp008, πE = πg.Sub(πF, µC, πg.NewInt(32).ToObject()); πE != nil {
									continue
								}
								if πTemp007, πE = πg.And(πF, πTemp008, πg.NewInt(3).ToObject()); πE != nil {
									continue
								}
								if πTemp004, πE = πg.LShift(πF, πTemp007, πg.NewInt(6).ToObject()); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µD, "D"); πE != nil {
									continue
								}
								if πTemp008, πE = πg.Sub(πF, µD, πg.NewInt(32).ToObject()); πE != nil {
									continue
								}
								if πTemp007, πE = πg.And(πF, πTemp008, πg.NewInt(63).ToObject()); πE != nil {
									continue
								}
								if πTemp003, πE = πg.Or(πF, πTemp004, πTemp007); πE != nil {
									continue
								}
								πTemp011[0] = πTemp003
								if πTemp003, πE = πg.ResolveGlobal(πF, ßchr); πE != nil {
									continue
								}
								if πTemp004, πE = πTemp003.Call(πF, πTemp011, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp011)
								πTemp010[2] = πTemp004
								πTemp003 = πg.NewList(πTemp010...).ToObject()
								πTemp002[0] = πTemp003
								if πTemp003, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
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
					if πTemp006, πE = πTemp005.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ListType.Call(πF, πg.Args{πTemp006}, nil); πE != nil {
						continue
					}
					µresult = πTemp003
					πF.PopCheckpoint()
					goto Label3
				Label4:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp008, πTemp009 = πF.ExcInfo()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsInstance(πF, πTemp008.ToObject(), πTemp003); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label5
					}
					πE = πF.Raise(πTemp008.ToObject(), nil, πTemp009.ToObject())
					continue
					// line 38: except ValueError:
					πF.SetLineno(38)
				Label5:
					πTemp004 = πF.MakeArgs(1)
					πTemp004[0] = πg.NewStr("Illegal char").ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßError); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 39: raise Error('Illegal char')
					πF.SetLineno(39)
					πE = πF.Raise(πTemp006, nil, nil)
					continue
					πF.RestoreExc(nil, nil)
					goto Label3
				Label3:
					// line 40: result = ''.join(result)
					πF.SetLineno(40)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					πTemp004[0] = µresult
					if πTemp003, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µresult = πTemp006
					// line 41: trailingdata = result[length:]
					πF.SetLineno(41)
					if πE = πg.CheckLocal(πF, µlength, "length"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{µlength, πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, µresult, πTemp003); πE != nil {
						continue
					}
					µtrailingdata = πTemp006
					// line 44: result = result[:length]
					πF.SetLineno(44)
					if πE = πg.CheckLocal(πF, µlength, "length"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.None, µlength, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, µresult, πTemp003); πE != nil {
						continue
					}
					µresult = πTemp006
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					πTemp004[0] = µresult
					if πTemp006, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp010, πE = πTemp006.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πE = πg.CheckLocal(πF, µlength, "length"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.LT(πF, πTemp010, µlength); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label6
					}
					goto Label7
					// line 45: if len(result) < length:
					πF.SetLineno(45)
				Label6:
					// line 46: result += ((length - len(result)) * '\x00')
					πF.SetLineno(46)
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µlength, "length"); πE != nil {
						continue
					}
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					πTemp004[0] = µresult
					if πTemp010, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp011, πE = πTemp010.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp006, πE = πg.Sub(πF, µlength, πTemp011); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Mul(πF, πTemp006, πg.NewStr("\x00").ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IAdd(πF, µresult, πTemp003); πE != nil {
						continue
					}
					µresult = πTemp006
					goto Label7
				Label7:
					// line 47: return result
					πF.SetLineno(47)
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
			if πE = πF.Globals().SetItem(πF, ßa2b_uu.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 50: def b2a_uu(s):
			πF.SetLineno(50)
			πTemp006 = make([]πg.Param, 1)
			πTemp006[0] = πg.Param{Name: "s", Def: nil}
			πTemp004 = πg.NewFunction(πg.NewCode("b2a_uu", "/usr/lib/python2.7/binascii.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µs *πg.Object = πArgs[0]; _ = µs
				var µlength *πg.Object = πg.UnboundLocal; _ = µlength
				var µtriples_gen *πg.Object = πg.UnboundLocal; _ = µtriples_gen
				var µresult *πg.Object = πg.UnboundLocal; _ = µresult
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 bool
				_ = πTemp004
				var πTemp005 []πg.Param
				_ = πTemp005
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 *πg.Object
				_ = πTemp008
				var πTemp009 []*πg.Object
				_ = πTemp009
				var πTemp010 *πg.Object
				_ = πTemp010
				var πTemp011 *πg.Object
				_ = πTemp011
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 51: length = len(s)
					πF.SetLineno(51)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp001[0] = µs
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µlength = πTemp003
					if πE = πg.CheckLocal(πF, µlength, "length"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GT(πF, µlength, πg.NewInt(45).ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label1
					}
					goto Label2
					// line 52: if length > 45:
					πF.SetLineno(52)
				Label1:
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("At most 45 bytes at once").ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 53: raise Error('At most 45 bytes at once')
					πF.SetLineno(53)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					goto Label2
				Label2:
					// line 55: def triples_gen(s):
					πF.SetLineno(55)
					πTemp005 = make([]πg.Param, 1)
					πTemp005[0] = πg.Param{Name: "s", Def: nil}
					πTemp002 = πg.NewFunction(πg.NewCode("triples_gen", "/usr/lib/python2.7/binascii.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µs *πg.Object = πArgs[0]; _ = µs
						var πTemp001 bool
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
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 *πg.BaseException
						_ = πTemp009
						var πTemp010 *πg.Traceback
						_ = πTemp010
						var πTemp011 *πg.Object
						_ = πTemp011
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
							for ; πF.State() >= 0; πF.PopCheckpoint() {
								switch πF.State() {
								case 0:
								case 8: goto Label8
								case 1: goto Label1
								case 2: goto Label2
								case 5: goto Label5
								case 6: goto Label6
								default: panic("unexpected function state")
								}
								// line 56: while s:
								πF.SetLineno(56)
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
								if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.IsTrue(πF, µs); πE != nil {
									continue
								}
								if πE != nil || !πTemp002 {
									continue
								}
								πF.PushCheckpoint(1)            
								// line 57: try:
								πF.SetLineno(57)
								πF.PushCheckpoint(5)
								// line 58: yield ord(s[0]), ord(s[1]), ord(s[2])
								πF.SetLineno(58)
								πTemp004 = πF.MakeArgs(1)
								πTemp005 = πg.NewInt(0).ToObject()
								if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
									continue
								}
								if πTemp006, πE = πg.GetItem(πF, µs, πTemp005); πE != nil {
									continue
								}
								πTemp004[0] = πTemp006
								if πTemp005, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
									continue
								}
								if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								πTemp004 = πF.MakeArgs(1)
								πTemp005 = πg.NewInt(1).ToObject()
								if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
									continue
								}
								if πTemp007, πE = πg.GetItem(πF, µs, πTemp005); πE != nil {
									continue
								}
								πTemp004[0] = πTemp007
								if πTemp005, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
									continue
								}
								if πTemp007, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								πTemp004 = πF.MakeArgs(1)
								πTemp005 = πg.NewInt(2).ToObject()
								if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
									continue
								}
								if πTemp008, πE = πg.GetItem(πF, µs, πTemp005); πE != nil {
									continue
								}
								πTemp004[0] = πTemp008
								if πTemp005, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
									continue
								}
								if πTemp008, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								πTemp003 = πg.NewTuple3(πTemp006, πTemp007, πTemp008).ToObject()
								πF.PushCheckpoint(6)
								return πTemp003, nil
							Label6:
								πTemp005 = πSent
								πF.PopCheckpoint()
								goto Label4
							Label5:
								if πE == nil {
								  continue
								}
								πE = nil
								πTemp009, πTemp010 = πF.ExcInfo()
								if πTemp005, πE = πg.ResolveGlobal(πF, ßIndexError); πE != nil {
									continue
								}
								if πTemp002, πE = πg.IsInstance(πF, πTemp009.ToObject(), πTemp005); πE != nil {
									continue
								}
								if πTemp002 {
									goto Label7
								}
								πE = πF.Raise(πTemp009.ToObject(), nil, πTemp010.ToObject())
								continue
								// line 59: except IndexError:
								πF.SetLineno(59)
							Label7:
								// line 60: s += '\0\0'
								πF.SetLineno(60)
								if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
									continue
								}
								if πTemp005, πE = πg.IAdd(πF, µs, πg.NewStr("\x00\x00").ToObject()); πE != nil {
									continue
								}
								µs = πTemp005
								// line 61: yield ord(s[0]), ord(s[1]), ord(s[2])
								πF.SetLineno(61)
								πTemp004 = πF.MakeArgs(1)
								πTemp006 = πg.NewInt(0).ToObject()
								if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
									continue
								}
								if πTemp007, πE = πg.GetItem(πF, µs, πTemp006); πE != nil {
									continue
								}
								πTemp004[0] = πTemp007
								if πTemp006, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
									continue
								}
								if πTemp007, πE = πTemp006.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								πTemp004 = πF.MakeArgs(1)
								πTemp006 = πg.NewInt(1).ToObject()
								if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
									continue
								}
								if πTemp008, πE = πg.GetItem(πF, µs, πTemp006); πE != nil {
									continue
								}
								πTemp004[0] = πTemp008
								if πTemp006, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
									continue
								}
								if πTemp008, πE = πTemp006.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								πTemp004 = πF.MakeArgs(1)
								πTemp006 = πg.NewInt(2).ToObject()
								if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
									continue
								}
								if πTemp011, πE = πg.GetItem(πF, µs, πTemp006); πE != nil {
									continue
								}
								πTemp004[0] = πTemp011
								if πTemp006, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
									continue
								}
								if πTemp011, πE = πTemp006.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								πTemp005 = πg.NewTuple3(πTemp007, πTemp008, πTemp011).ToObject()
								πF.PushCheckpoint(8)
								return πTemp005, nil
							Label8:
								πTemp006 = πSent
								// line 62: return
								πF.SetLineno(62)
								πR = πg.None
								continue
								πF.RestoreExc(nil, nil)
								goto Label4
							Label4:
								// line 63: s = s[3:]
								πF.SetLineno(63)
								if πTemp006, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(3).ToObject(), πg.None, πg.None}, nil); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
									continue
								}
								if πTemp007, πE = πg.GetItem(πF, µs, πTemp006); πE != nil {
									continue
								}
								µs = πTemp007
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
					µtriples_gen = πTemp002
					// line 65: result = [''.join(
					πF.SetLineno(65)
					πTemp005 = make([]πg.Param, 0)
					πTemp006 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/binascii.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µA *πg.Object = πg.UnboundLocal; _ = µA
						var µB *πg.Object = πg.UnboundLocal; _ = µB
						var µC *πg.Object = πg.UnboundLocal; _ = µC
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
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
						var πTemp009 []*πg.Object
						_ = πTemp009
						var πTemp010 []*πg.Object
						_ = πTemp010
						var πTemp011 *πg.Object
						_ = πTemp011
						var πTemp012 *πg.Object
						_ = πTemp012
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
								if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
									continue
								}
								πTemp002[0] = µs
								if πE = πg.CheckLocal(πF, µtriples_gen, "triples_gen"); πE != nil {
									continue
								}
								if πTemp003, πE = µtriples_gen.Call(πF, πTemp002, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp002)
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
								if πTemp003, πE = πg.Next(πF, πTemp001); πE != nil {
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
									if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp006}, πg.TieTarget{Target: &πTemp007}, πg.TieTarget{Target: &πTemp008}}}, πTemp003); πE != nil {
										continue
									}
									µA = πTemp006
									µB = πTemp007
									µC = πTemp008
								}
								if πE != nil || !πTemp005 {
									continue
								}
								πF.PushCheckpoint(1)            
								// line 65: result = [''.join(
								πF.SetLineno(65)
								πTemp002 = πF.MakeArgs(1)
								πTemp009 = make([]*πg.Object, 4)
								πTemp010 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µA, "A"); πE != nil {
									continue
								}
								if πTemp007, πE = πg.RShift(πF, µA, πg.NewInt(2).ToObject()); πE != nil {
									continue
								}
								if πTemp006, πE = πg.And(πF, πTemp007, πg.NewInt(63).ToObject()); πE != nil {
									continue
								}
								if πTemp003, πE = πg.Add(πF, πg.NewInt(32).ToObject(), πTemp006); πE != nil {
									continue
								}
								πTemp010[0] = πTemp003
								if πTemp003, πE = πg.ResolveGlobal(πF, ßchr); πE != nil {
									continue
								}
								if πTemp006, πE = πTemp003.Call(πF, πTemp010, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp010)
								πTemp009[0] = πTemp006
								πTemp010 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µA, "A"); πE != nil {
									continue
								}
								if πTemp008, πE = πg.LShift(πF, µA, πg.NewInt(4).ToObject()); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µB, "B"); πE != nil {
									continue
								}
								if πTemp012, πE = πg.RShift(πF, µB, πg.NewInt(4).ToObject()); πE != nil {
									continue
								}
								if πTemp011, πE = πg.And(πF, πTemp012, πg.NewInt(15).ToObject()); πE != nil {
									continue
								}
								if πTemp007, πE = πg.Or(πF, πTemp008, πTemp011); πE != nil {
									continue
								}
								if πTemp006, πE = πg.And(πF, πTemp007, πg.NewInt(63).ToObject()); πE != nil {
									continue
								}
								if πTemp003, πE = πg.Add(πF, πg.NewInt(32).ToObject(), πTemp006); πE != nil {
									continue
								}
								πTemp010[0] = πTemp003
								if πTemp003, πE = πg.ResolveGlobal(πF, ßchr); πE != nil {
									continue
								}
								if πTemp006, πE = πTemp003.Call(πF, πTemp010, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp010)
								πTemp009[1] = πTemp006
								πTemp010 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µB, "B"); πE != nil {
									continue
								}
								if πTemp008, πE = πg.LShift(πF, µB, πg.NewInt(2).ToObject()); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µC, "C"); πE != nil {
									continue
								}
								if πTemp012, πE = πg.RShift(πF, µC, πg.NewInt(6).ToObject()); πE != nil {
									continue
								}
								if πTemp011, πE = πg.And(πF, πTemp012, πg.NewInt(3).ToObject()); πE != nil {
									continue
								}
								if πTemp007, πE = πg.Or(πF, πTemp008, πTemp011); πE != nil {
									continue
								}
								if πTemp006, πE = πg.And(πF, πTemp007, πg.NewInt(63).ToObject()); πE != nil {
									continue
								}
								if πTemp003, πE = πg.Add(πF, πg.NewInt(32).ToObject(), πTemp006); πE != nil {
									continue
								}
								πTemp010[0] = πTemp003
								if πTemp003, πE = πg.ResolveGlobal(πF, ßchr); πE != nil {
									continue
								}
								if πTemp006, πE = πTemp003.Call(πF, πTemp010, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp010)
								πTemp009[2] = πTemp006
								πTemp010 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µC, "C"); πE != nil {
									continue
								}
								if πTemp006, πE = πg.And(πF, µC, πg.NewInt(63).ToObject()); πE != nil {
									continue
								}
								if πTemp003, πE = πg.Add(πF, πg.NewInt(32).ToObject(), πTemp006); πE != nil {
									continue
								}
								πTemp010[0] = πTemp003
								if πTemp003, πE = πg.ResolveGlobal(πF, ßchr); πE != nil {
									continue
								}
								if πTemp006, πE = πTemp003.Call(πF, πTemp010, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp010)
								πTemp009[3] = πTemp006
								πTemp003 = πg.NewList(πTemp009...).ToObject()
								πTemp002[0] = πTemp003
								if πTemp003, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
									continue
								}
								if πTemp006, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp002)
								πF.PushCheckpoint(4)
								return πTemp006, nil
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
					if πTemp007, πE = πTemp006.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ListType.Call(πF, πg.Args{πTemp007}, nil); πE != nil {
						continue
					}
					µresult = πTemp003
					// line 71: return chr(ord(' ') + (length & 077)) + ''.join(result) + '\n'
					πF.SetLineno(71)
					πTemp001 = πF.MakeArgs(1)
					πTemp009 = πF.MakeArgs(1)
					πTemp009[0] = πg.NewStr(" ").ToObject()
					if πTemp010, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
						continue
					}
					if πTemp011, πE = πTemp010.Call(πF, πTemp009, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp009)
					if πE = πg.CheckLocal(πF, µlength, "length"); πE != nil {
						continue
					}
					if πTemp010, πE = πg.And(πF, µlength, πg.NewInt(63).ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.Add(πF, πTemp011, πTemp010); πE != nil {
						continue
					}
					πTemp001[0] = πTemp008
					if πTemp008, πE = πg.ResolveGlobal(πF, ßchr); πE != nil {
						continue
					}
					if πTemp010, πE = πTemp008.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					πTemp001[0] = µresult
					if πTemp008, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
						continue
					}
					if πTemp011, πE = πTemp008.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp007, πE = πg.Add(πF, πTemp010, πTemp011); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Add(πF, πTemp007, πg.NewStr("\n").ToObject()); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßb2a_uu.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 74: table_a2b_base64 = {
			πF.SetLineno(74)
			πTemp001 = πg.NewDict()
			if πE = πTemp001.SetItem(πF, ßA.ToObject(), πg.NewInt(0).ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßB.ToObject(), πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßC.ToObject(), πg.NewInt(2).ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßD.ToObject(), πg.NewInt(3).ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßE.ToObject(), πg.NewInt(4).ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßF.ToObject(), πg.NewInt(5).ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßG.ToObject(), πg.NewInt(6).ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßH.ToObject(), πg.NewInt(7).ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßI.ToObject(), πg.NewInt(8).ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßJ.ToObject(), πg.NewInt(9).ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßK.ToObject(), πg.NewInt(10).ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßL.ToObject(), πg.NewInt(11).ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßM.ToObject(), πg.NewInt(12).ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßN.ToObject(), πg.NewInt(13).ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßO.ToObject(), πg.NewInt(14).ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßP.ToObject(), πg.NewInt(15).ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßQ.ToObject(), πg.NewInt(16).ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßR.ToObject(), πg.NewInt(17).ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßS.ToObject(), πg.NewInt(18).ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßT.ToObject(), πg.NewInt(19).ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßU.ToObject(), πg.NewInt(20).ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßV.ToObject(), πg.NewInt(21).ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßW.ToObject(), πg.NewInt(22).ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßX.ToObject(), πg.NewInt(23).ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßY.ToObject(), πg.NewInt(24).ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßZ.ToObject(), πg.NewInt(25).ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßa.ToObject(), πg.NewInt(26).ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßb.ToObject(), πg.NewInt(27).ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßc.ToObject(), πg.NewInt(28).ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßd.ToObject(), πg.NewInt(29).ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ße.ToObject(), πg.NewInt(30).ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßf.ToObject(), πg.NewInt(31).ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßg.ToObject(), πg.NewInt(32).ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßh.ToObject(), πg.NewInt(33).ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßi.ToObject(), πg.NewInt(34).ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßj.ToObject(), πg.NewInt(35).ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßk.ToObject(), πg.NewInt(36).ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßl.ToObject(), πg.NewInt(37).ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßm.ToObject(), πg.NewInt(38).ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßn.ToObject(), πg.NewInt(39).ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßo.ToObject(), πg.NewInt(40).ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßp.ToObject(), πg.NewInt(41).ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßq.ToObject(), πg.NewInt(42).ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßr.ToObject(), πg.NewInt(43).ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßs.ToObject(), πg.NewInt(44).ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßt.ToObject(), πg.NewInt(45).ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßu.ToObject(), πg.NewInt(46).ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßv.ToObject(), πg.NewInt(47).ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßw.ToObject(), πg.NewInt(48).ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßx.ToObject(), πg.NewInt(49).ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßy.ToObject(), πg.NewInt(50).ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßz.ToObject(), πg.NewInt(51).ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ß0.ToObject(), πg.NewInt(52).ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ß1.ToObject(), πg.NewInt(53).ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ß2.ToObject(), πg.NewInt(54).ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ß3.ToObject(), πg.NewInt(55).ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ß4.ToObject(), πg.NewInt(56).ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ß5.ToObject(), πg.NewInt(57).ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ß6.ToObject(), πg.NewInt(58).ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ß7.ToObject(), πg.NewInt(59).ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ß8.ToObject(), πg.NewInt(60).ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ß9.ToObject(), πg.NewInt(61).ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewStr("+").ToObject(), πg.NewInt(62).ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewStr("/").ToObject(), πg.NewInt(63).ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewStr("=").ToObject(), πg.NewInt(0).ToObject()); πE != nil {
				continue
			}
			πTemp005 = πTemp001.ToObject()
			if πE = πF.Globals().SetItem(πF, ßtable_a2b_base64.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 143: def a2b_base64(s):
			πF.SetLineno(143)
			πTemp006 = make([]πg.Param, 1)
			πTemp006[0] = πg.Param{Name: "s", Def: nil}
			πTemp005 = πg.NewFunction(πg.NewCode("a2b_base64", "/usr/lib/python2.7/binascii.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µs *πg.Object = πArgs[0]; _ = µs
				var µnext_valid_char *πg.Object = πg.UnboundLocal; _ = µnext_valid_char
				var µquad_pos *πg.Object = πg.UnboundLocal; _ = µquad_pos
				var µleftbits *πg.Object = πg.UnboundLocal; _ = µleftbits
				var µleftchar *πg.Object = πg.UnboundLocal; _ = µleftchar
				var µres *πg.Object = πg.UnboundLocal; _ = µres
				var µi *πg.Object = πg.UnboundLocal; _ = µi
				var µc *πg.Object = πg.UnboundLocal; _ = µc
				var µnext_c *πg.Object = πg.UnboundLocal; _ = µnext_c
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
				var πTemp006 bool
				_ = πTemp006
				var πTemp007 []πg.Param
				_ = πTemp007
				var πTemp008 bool
				_ = πTemp008
				var πTemp009 *πg.Object
				_ = πTemp009
				var πTemp010 bool
				_ = πTemp010
				var πTemp011 *πg.Object
				_ = πTemp011
				var πTemp012 *πg.BaseException
				_ = πTemp012
				var πTemp013 *πg.Traceback
				_ = πTemp013
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 17: goto Label17
					case 3: goto Label3
					case 4: goto Label4
					default: panic("unexpected function state")
					}
					πTemp002 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp002[0] = µs
					if πTemp004, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveGlobal(πF, ßunicode); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple2(πTemp004, πTemp005).ToObject()
					πTemp002[1] = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(!πTemp006).ToObject()
					if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label1
					}
					goto Label2
					// line 144: if not isinstance(s, (str, unicode)):
					πF.SetLineno(144)
				Label1:
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple1(µs).ToObject()
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("expected string or unicode, got %r").ToObject(), πTemp003); πE != nil {
						continue
					}
					πTemp002[0] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					// line 145: raise TypeError("expected string or unicode, got %r" % (s,))
					πF.SetLineno(145)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					goto Label2
				Label2:
					// line 146: s = s.rstrip()
					πF.SetLineno(146)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µs, ßrstrip, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					µs = πTemp003
					// line 150: def next_valid_char(s, pos):
					πF.SetLineno(150)
					πTemp007 = make([]πg.Param, 2)
					πTemp007[0] = πg.Param{Name: "s", Def: nil}
					πTemp007[1] = πg.Param{Name: "pos", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("next_valid_char", "/usr/lib/python2.7/binascii.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µs *πg.Object = πArgs[0]; _ = µs
						var µpos *πg.Object = πArgs[1]; _ = µpos
						var µi *πg.Object = πg.UnboundLocal; _ = µi
						var µc *πg.Object = πg.UnboundLocal; _ = µc
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 []*πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 bool
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
							case 1: goto Label1
							case 2: goto Label2
							case 7: goto Label7
							default: panic("unexpected function state")
							}
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µpos, "pos"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, µpos, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp004[0] = µs
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp002[1] = πTemp005
							if πTemp003, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp001, πE = πg.Iter(πF, πTemp005); πE != nil {
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
							if πTemp003, πE = πg.Next(πF, πTemp001); πE != nil {
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
							πF.PushCheckpoint(1)            
							// line 152: c = s[i]
							πF.SetLineno(152)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp003 = µi
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µs, πTemp003); πE != nil {
								continue
							}
							µc = πTemp005
							if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.LT(πF, µc, πg.NewStr("\x7f").ToObject()); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label4
							}
							goto Label5
							// line 153: if c < '\x7f':
							πF.SetLineno(153)
						Label4:
							// line 154: try:
							πF.SetLineno(154)
							πF.PushCheckpoint(7)
							// line 155: table_a2b_base64[c]
							πF.SetLineno(155)
							if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
								continue
							}
							πTemp003 = µc
							if πTemp008, πE = πg.ResolveGlobal(πF, ßtable_a2b_base64); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, πTemp008, πTemp003); πE != nil {
								continue
							}
							// line 156: return c
							πF.SetLineno(156)
							if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
								continue
							}
							πR = µc
							continue
							πF.PopCheckpoint()
							goto Label6
						Label7:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp009, πTemp010 = πF.ExcInfo()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsInstance(πF, πTemp009.ToObject(), πTemp003); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label8
							}
							πE = πF.Raise(πTemp009.ToObject(), nil, πTemp010.ToObject())
							continue
							// line 157: except KeyError:
							πF.SetLineno(157)
						Label8:
							// line 158: pass
							πF.SetLineno(158)
							πF.RestoreExc(nil, nil)
							goto Label6
						Label6:
							goto Label5
						Label5:
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 159: return None
							πF.SetLineno(159)
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
					µnext_valid_char = πTemp001
					// line 161: quad_pos = 0
					πF.SetLineno(161)
					µquad_pos = πg.NewInt(0).ToObject()
					// line 162: leftbits = 0
					πF.SetLineno(162)
					µleftbits = πg.NewInt(0).ToObject()
					// line 163: leftchar = 0
					πF.SetLineno(163)
					µleftchar = πg.NewInt(0).ToObject()
					// line 164: res = []
					πF.SetLineno(164)
					πTemp002 = make([]*πg.Object, 0)
					πTemp003 = πg.NewList(πTemp002...).ToObject()
					µres = πTemp003
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp002[0] = µs
					if πTemp004, πE = πg.ResolveGlobal(πF, ßenumerate); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp003, πE = πg.Iter(πF, πTemp005); πE != nil {
						continue
					}
					πF.PushCheckpoint(4)
					πTemp006 = false
				Label3:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp006 {
						πF.PopCheckpoint()
						goto Label5
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
						if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp009}}}, πTemp004); πE != nil {
							continue
						}
						µi = πTemp005
						µc = πTemp009
					}
					if πE != nil || !πTemp008 {
						continue
					}
					πF.PushCheckpoint(3)            
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GT(πF, µc, πg.NewStr("\x7f").ToObject()); πE != nil {
						continue
					}
					πTemp004 = πTemp005
					if πTemp008, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label6
					}
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Eq(πF, µc, πg.NewStr("\n").ToObject()); πE != nil {
						continue
					}
					πTemp004 = πTemp005
					if πTemp008, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label6
					}
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Eq(πF, µc, πg.NewStr("\r").ToObject()); πE != nil {
						continue
					}
					πTemp004 = πTemp005
					if πTemp008, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label6
					}
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Eq(πF, µc, πg.NewStr(" ").ToObject()); πE != nil {
						continue
					}
					πTemp004 = πTemp005
				Label6:
					if πTemp008, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label7
					}
					goto Label8
					// line 166: if c > '\x7f' or c == '\n' or c == '\r' or c == ' ':
					πF.SetLineno(166)
				Label7:
					// line 167: continue
					πF.SetLineno(167)
					continue
					goto Label8
				Label8:
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Eq(πF, µc, πg.NewStr("=").ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label9
					}
					goto Label10
					// line 168: if c == '=':
					πF.SetLineno(168)
				Label9:
					if πE = πg.CheckLocal(πF, µquad_pos, "quad_pos"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.LT(πF, µquad_pos, πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
					πTemp004 = πTemp005
					if πTemp008, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label11
					}
					if πE = πg.CheckLocal(πF, µquad_pos, "quad_pos"); πE != nil {
						continue
					}
					if πTemp009, πE = πg.Eq(πF, µquad_pos, πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
					πTemp005 = πTemp009
					if πTemp010, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					if !πTemp010 {
						goto Label12
					}
					πTemp002 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp002[0] = µs
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					πTemp002[1] = µi
					if πE = πg.CheckLocal(πF, µnext_valid_char, "next_valid_char"); πE != nil {
						continue
					}
					if πTemp011, πE = µnext_valid_char.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp009, πE = πg.NE(πF, πTemp011, πg.NewStr("=").ToObject()); πE != nil {
						continue
					}
					πTemp005 = πTemp009
				Label12:
					πTemp004 = πTemp005
				Label11:
					if πTemp008, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label13
					}
					goto Label14
					// line 169: if quad_pos < 2 or (quad_pos == 2 and next_valid_char(s, i) != '='):
					πF.SetLineno(169)
				Label13:
					// line 170: continue
					πF.SetLineno(170)
					continue
					goto Label15
				Label14:
					// line 172: leftbits = 0
					πF.SetLineno(172)
					µleftbits = πg.NewInt(0).ToObject()
					// line 173: break
					πF.SetLineno(173)
					πTemp006 = true
					continue
					goto Label15
				Label15:
					goto Label10
				Label10:
					// line 174: try:
					πF.SetLineno(174)
					πF.PushCheckpoint(17)
					// line 175: next_c = table_a2b_base64[c]
					πF.SetLineno(175)
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					πTemp004 = µc
					if πTemp009, πE = πg.ResolveGlobal(πF, ßtable_a2b_base64); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, πTemp009, πTemp004); πE != nil {
						continue
					}
					µnext_c = πTemp005
					πF.PopCheckpoint()
					goto Label16
				Label17:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp012, πTemp013 = πF.ExcInfo()
					if πTemp004, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsInstance(πF, πTemp012.ToObject(), πTemp004); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label18
					}
					πE = πF.Raise(πTemp012.ToObject(), nil, πTemp013.ToObject())
					continue
					// line 176: except KeyError:
					πF.SetLineno(176)
				Label18:
					// line 177: continue
					πF.SetLineno(177)
					continue
					πF.RestoreExc(nil, nil)
					goto Label16
				Label16:
					// line 178: quad_pos = (quad_pos + 1) & 0x03
					πF.SetLineno(178)
					if πE = πg.CheckLocal(πF, µquad_pos, "quad_pos"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Add(πF, µquad_pos, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.And(πF, πTemp005, πg.NewInt(3).ToObject()); πE != nil {
						continue
					}
					µquad_pos = πTemp004
					// line 179: leftchar = (leftchar << 6) | next_c
					πF.SetLineno(179)
					if πE = πg.CheckLocal(πF, µleftchar, "leftchar"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.LShift(πF, µleftchar, πg.NewInt(6).ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µnext_c, "next_c"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Or(πF, πTemp005, µnext_c); πE != nil {
						continue
					}
					µleftchar = πTemp004
					// line 180: leftbits += 6
					πF.SetLineno(180)
					if πE = πg.CheckLocal(πF, µleftbits, "leftbits"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IAdd(πF, µleftbits, πg.NewInt(6).ToObject()); πE != nil {
						continue
					}
					µleftbits = πTemp004
					if πE = πg.CheckLocal(πF, µleftbits, "leftbits"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GE(πF, µleftbits, πg.NewInt(8).ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label19
					}
					goto Label20
					// line 181: if leftbits >= 8:
					πF.SetLineno(181)
				Label19:
					// line 182: leftbits -= 8
					πF.SetLineno(182)
					if πE = πg.CheckLocal(πF, µleftbits, "leftbits"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ISub(πF, µleftbits, πg.NewInt(8).ToObject()); πE != nil {
						continue
					}
					µleftbits = πTemp004
					// line 183: res.append((leftchar >> leftbits & 0xff))
					πF.SetLineno(183)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µleftchar, "leftchar"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µleftbits, "leftbits"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.RShift(πF, µleftchar, µleftbits); πE != nil {
						continue
					}
					if πTemp004, πE = πg.And(πF, πTemp005, πg.NewInt(255).ToObject()); πE != nil {
						continue
					}
					πTemp002[0] = πTemp004
					if πE = πg.CheckLocal(πF, µres, "res"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µres, ßappend, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					// line 184: leftchar &= ((1 << leftbits) - 1)
					πF.SetLineno(184)
					if πE = πg.CheckLocal(πF, µleftchar, "leftchar"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µleftbits, "leftbits"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.LShift(πF, πg.NewInt(1).ToObject(), µleftbits); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Sub(πF, πTemp005, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IAnd(πF, µleftchar, πTemp004); πE != nil {
						continue
					}
					µleftchar = πTemp005
					goto Label20
				Label20:
					continue
				Label4:
					if πE != nil || πR != nil {
						continue
					}
				Label5:
					if πE = πg.CheckLocal(πF, µleftbits, "leftbits"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.NE(πF, µleftbits, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label21
					}
					goto Label22
					// line 185: if leftbits != 0:
					πF.SetLineno(185)
				Label21:
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewStr("Incorrect padding").ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßError); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					// line 186: raise Error('Incorrect padding')
					πF.SetLineno(186)
					πE = πF.Raise(πTemp004, nil, nil)
					continue
					goto Label22
				Label22:
					// line 188: return ''.join([chr(i) for i in res])
					πF.SetLineno(188)
					πTemp002 = πF.MakeArgs(1)
					πTemp007 = make([]πg.Param, 0)
					πTemp004 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/binascii.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
								if πE = πg.CheckLocal(πF, µres, "res"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.Iter(πF, µres); πE != nil {
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
								// line 188: return ''.join([chr(i) for i in res])
								πF.SetLineno(188)
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
					if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ListType.Call(πF, πg.Args{πTemp005}, nil); πE != nil {
						continue
					}
					πTemp002[0] = πTemp003
					if πTemp003, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
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
			if πE = πF.Globals().SetItem(πF, ßa2b_base64.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 190: table_b2a_base64 = \
			πF.SetLineno(190)
			if πE = πF.Globals().SetItem(πF, ßtable_b2a_base64.ToObject(), πg.NewStr("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/").ToObject()); πE != nil {
				continue
			}
			// line 193: def b2a_base64(s):
			πF.SetLineno(193)
			πTemp006 = make([]πg.Param, 1)
			πTemp006[0] = πg.Param{Name: "s", Def: nil}
			πTemp007 = πg.NewFunction(πg.NewCode("b2a_base64", "/usr/lib/python2.7/binascii.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µs *πg.Object = πArgs[0]; _ = µs
				var µlength *πg.Object = πg.UnboundLocal; _ = µlength
				var µfinal_length *πg.Object = πg.UnboundLocal; _ = µfinal_length
				var µtriples_gen *πg.Object = πg.UnboundLocal; _ = µtriples_gen
				var µa *πg.Object = πg.UnboundLocal; _ = µa
				var µresult *πg.Object = πg.UnboundLocal; _ = µresult
				var µfinal *πg.Object = πg.UnboundLocal; _ = µfinal
				var µsnippet *πg.Object = πg.UnboundLocal; _ = µsnippet
				var µb *πg.Object = πg.UnboundLocal; _ = µb
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
				var πTemp006 *πg.Object
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
				var πTemp013 *πg.Object
				_ = πTemp013
				var πTemp014 *πg.Object
				_ = πTemp014
				var πTemp015 *πg.Object
				_ = πTemp015
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 194: length = len(s)
					πF.SetLineno(194)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp001[0] = µs
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µlength = πTemp003
					// line 195: final_length = length % 3
					πF.SetLineno(195)
					if πE = πg.CheckLocal(πF, µlength, "length"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mod(πF, µlength, πg.NewInt(3).ToObject()); πE != nil {
						continue
					}
					µfinal_length = πTemp002
					// line 197: def triples_gen(s):
					πF.SetLineno(197)
					πTemp004 = make([]πg.Param, 1)
					πTemp004[0] = πg.Param{Name: "s", Def: nil}
					πTemp002 = πg.NewFunction(πg.NewCode("triples_gen", "/usr/lib/python2.7/binascii.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µs *πg.Object = πArgs[0]; _ = µs
						var πTemp001 bool
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
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 *πg.BaseException
						_ = πTemp009
						var πTemp010 *πg.Traceback
						_ = πTemp010
						var πTemp011 *πg.Object
						_ = πTemp011
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
							for ; πF.State() >= 0; πF.PopCheckpoint() {
								switch πF.State() {
								case 0:
								case 8: goto Label8
								case 1: goto Label1
								case 2: goto Label2
								case 5: goto Label5
								case 6: goto Label6
								default: panic("unexpected function state")
								}
								// line 198: while s:
								πF.SetLineno(198)
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
								if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.IsTrue(πF, µs); πE != nil {
									continue
								}
								if πE != nil || !πTemp002 {
									continue
								}
								πF.PushCheckpoint(1)            
								// line 199: try:
								πF.SetLineno(199)
								πF.PushCheckpoint(5)
								// line 200: yield ord(s[0]), ord(s[1]), ord(s[2])
								πF.SetLineno(200)
								πTemp004 = πF.MakeArgs(1)
								πTemp005 = πg.NewInt(0).ToObject()
								if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
									continue
								}
								if πTemp006, πE = πg.GetItem(πF, µs, πTemp005); πE != nil {
									continue
								}
								πTemp004[0] = πTemp006
								if πTemp005, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
									continue
								}
								if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								πTemp004 = πF.MakeArgs(1)
								πTemp005 = πg.NewInt(1).ToObject()
								if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
									continue
								}
								if πTemp007, πE = πg.GetItem(πF, µs, πTemp005); πE != nil {
									continue
								}
								πTemp004[0] = πTemp007
								if πTemp005, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
									continue
								}
								if πTemp007, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								πTemp004 = πF.MakeArgs(1)
								πTemp005 = πg.NewInt(2).ToObject()
								if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
									continue
								}
								if πTemp008, πE = πg.GetItem(πF, µs, πTemp005); πE != nil {
									continue
								}
								πTemp004[0] = πTemp008
								if πTemp005, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
									continue
								}
								if πTemp008, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								πTemp003 = πg.NewTuple3(πTemp006, πTemp007, πTemp008).ToObject()
								πF.PushCheckpoint(6)
								return πTemp003, nil
							Label6:
								πTemp005 = πSent
								πF.PopCheckpoint()
								goto Label4
							Label5:
								if πE == nil {
								  continue
								}
								πE = nil
								πTemp009, πTemp010 = πF.ExcInfo()
								if πTemp005, πE = πg.ResolveGlobal(πF, ßIndexError); πE != nil {
									continue
								}
								if πTemp002, πE = πg.IsInstance(πF, πTemp009.ToObject(), πTemp005); πE != nil {
									continue
								}
								if πTemp002 {
									goto Label7
								}
								πE = πF.Raise(πTemp009.ToObject(), nil, πTemp010.ToObject())
								continue
								// line 201: except IndexError:
								πF.SetLineno(201)
							Label7:
								// line 202: s += '\0\0'
								πF.SetLineno(202)
								if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
									continue
								}
								if πTemp005, πE = πg.IAdd(πF, µs, πg.NewStr("\x00\x00").ToObject()); πE != nil {
									continue
								}
								µs = πTemp005
								// line 203: yield ord(s[0]), ord(s[1]), ord(s[2])
								πF.SetLineno(203)
								πTemp004 = πF.MakeArgs(1)
								πTemp006 = πg.NewInt(0).ToObject()
								if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
									continue
								}
								if πTemp007, πE = πg.GetItem(πF, µs, πTemp006); πE != nil {
									continue
								}
								πTemp004[0] = πTemp007
								if πTemp006, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
									continue
								}
								if πTemp007, πE = πTemp006.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								πTemp004 = πF.MakeArgs(1)
								πTemp006 = πg.NewInt(1).ToObject()
								if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
									continue
								}
								if πTemp008, πE = πg.GetItem(πF, µs, πTemp006); πE != nil {
									continue
								}
								πTemp004[0] = πTemp008
								if πTemp006, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
									continue
								}
								if πTemp008, πE = πTemp006.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								πTemp004 = πF.MakeArgs(1)
								πTemp006 = πg.NewInt(2).ToObject()
								if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
									continue
								}
								if πTemp011, πE = πg.GetItem(πF, µs, πTemp006); πE != nil {
									continue
								}
								πTemp004[0] = πTemp011
								if πTemp006, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
									continue
								}
								if πTemp011, πE = πTemp006.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								πTemp005 = πg.NewTuple3(πTemp007, πTemp008, πTemp011).ToObject()
								πF.PushCheckpoint(8)
								return πTemp005, nil
							Label8:
								πTemp006 = πSent
								// line 204: return
								πF.SetLineno(204)
								πR = πg.None
								continue
								πF.RestoreExc(nil, nil)
								goto Label4
							Label4:
								// line 205: s = s[3:]
								πF.SetLineno(205)
								if πTemp006, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(3).ToObject(), πg.None, πg.None}, nil); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
									continue
								}
								if πTemp007, πE = πg.GetItem(πF, µs, πTemp006); πE != nil {
									continue
								}
								µs = πTemp007
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
					µtriples_gen = πTemp002
					// line 208: a = triples_gen(s[ :length - final_length])
					πF.SetLineno(208)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µlength, "length"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µfinal_length, "final_length"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Sub(πF, µlength, µfinal_length); πE != nil {
						continue
					}
					if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πTemp005, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µs, πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp005
					if πE = πg.CheckLocal(πF, µtriples_gen, "triples_gen"); πE != nil {
						continue
					}
					if πTemp003, πE = µtriples_gen.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µa = πTemp003
					// line 210: result = [''.join(
					πF.SetLineno(210)
					πTemp004 = make([]πg.Param, 0)
					πTemp005 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/binascii.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µA *πg.Object = πg.UnboundLocal; _ = µA
						var µB *πg.Object = πg.UnboundLocal; _ = µB
						var µC *πg.Object = πg.UnboundLocal; _ = µC
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
						var πTemp006 *πg.Object
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
								default: panic("unexpected function state")
								}
								if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.Iter(πF, µa); πE != nil {
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
									if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp006}, πg.TieTarget{Target: &πTemp007}}}, πTemp004); πE != nil {
										continue
									}
									µA = πTemp005
									µB = πTemp006
									µC = πTemp007
								}
								if πE != nil || !πTemp003 {
									continue
								}
								πF.PushCheckpoint(1)            
								// line 210: result = [''.join(
								πF.SetLineno(210)
								πTemp008 = πF.MakeArgs(1)
								πTemp009 = make([]*πg.Object, 4)
								if πE = πg.CheckLocal(πF, µA, "A"); πE != nil {
									continue
								}
								if πTemp006, πE = πg.RShift(πF, µA, πg.NewInt(2).ToObject()); πE != nil {
									continue
								}
								if πTemp005, πE = πg.And(πF, πTemp006, πg.NewInt(63).ToObject()); πE != nil {
									continue
								}
								πTemp004 = πTemp005
								if πTemp006, πE = πg.ResolveGlobal(πF, ßtable_b2a_base64); πE != nil {
									continue
								}
								if πTemp005, πE = πg.GetItem(πF, πTemp006, πTemp004); πE != nil {
									continue
								}
								πTemp009[0] = πTemp005
								if πE = πg.CheckLocal(πF, µA, "A"); πE != nil {
									continue
								}
								if πTemp007, πE = πg.LShift(πF, µA, πg.NewInt(4).ToObject()); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µB, "B"); πE != nil {
									continue
								}
								if πTemp011, πE = πg.RShift(πF, µB, πg.NewInt(4).ToObject()); πE != nil {
									continue
								}
								if πTemp010, πE = πg.And(πF, πTemp011, πg.NewInt(15).ToObject()); πE != nil {
									continue
								}
								if πTemp006, πE = πg.Or(πF, πTemp007, πTemp010); πE != nil {
									continue
								}
								if πTemp005, πE = πg.And(πF, πTemp006, πg.NewInt(63).ToObject()); πE != nil {
									continue
								}
								πTemp004 = πTemp005
								if πTemp006, πE = πg.ResolveGlobal(πF, ßtable_b2a_base64); πE != nil {
									continue
								}
								if πTemp005, πE = πg.GetItem(πF, πTemp006, πTemp004); πE != nil {
									continue
								}
								πTemp009[1] = πTemp005
								if πE = πg.CheckLocal(πF, µB, "B"); πE != nil {
									continue
								}
								if πTemp007, πE = πg.LShift(πF, µB, πg.NewInt(2).ToObject()); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µC, "C"); πE != nil {
									continue
								}
								if πTemp011, πE = πg.RShift(πF, µC, πg.NewInt(6).ToObject()); πE != nil {
									continue
								}
								if πTemp010, πE = πg.And(πF, πTemp011, πg.NewInt(3).ToObject()); πE != nil {
									continue
								}
								if πTemp006, πE = πg.Or(πF, πTemp007, πTemp010); πE != nil {
									continue
								}
								if πTemp005, πE = πg.And(πF, πTemp006, πg.NewInt(63).ToObject()); πE != nil {
									continue
								}
								πTemp004 = πTemp005
								if πTemp006, πE = πg.ResolveGlobal(πF, ßtable_b2a_base64); πE != nil {
									continue
								}
								if πTemp005, πE = πg.GetItem(πF, πTemp006, πTemp004); πE != nil {
									continue
								}
								πTemp009[2] = πTemp005
								if πE = πg.CheckLocal(πF, µC, "C"); πE != nil {
									continue
								}
								if πTemp005, πE = πg.And(πF, µC, πg.NewInt(63).ToObject()); πE != nil {
									continue
								}
								πTemp004 = πTemp005
								if πTemp006, πE = πg.ResolveGlobal(πF, ßtable_b2a_base64); πE != nil {
									continue
								}
								if πTemp005, πE = πg.GetItem(πF, πTemp006, πTemp004); πE != nil {
									continue
								}
								πTemp009[3] = πTemp005
								πTemp004 = πg.NewList(πTemp009...).ToObject()
								πTemp008[0] = πTemp004
								if πTemp004, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
									continue
								}
								if πTemp005, πE = πTemp004.Call(πF, πTemp008, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp008)
								πF.PushCheckpoint(4)
								return πTemp005, nil
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
					if πTemp006, πE = πTemp005.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ListType.Call(πF, πg.Args{πTemp006}, nil); πE != nil {
						continue
					}
					µresult = πTemp003
					// line 217: final = s[length - final_length:]
					πF.SetLineno(217)
					if πE = πg.CheckLocal(πF, µlength, "length"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µfinal_length, "final_length"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.Sub(πF, µlength, µfinal_length); πE != nil {
						continue
					}
					if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πTemp006, πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, µs, πTemp003); πE != nil {
						continue
					}
					µfinal = πTemp006
					if πE = πg.CheckLocal(πF, µfinal_length, "final_length"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Eq(πF, µfinal_length, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label1
					}
					if πE = πg.CheckLocal(πF, µfinal_length, "final_length"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Eq(πF, µfinal_length, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label2
					}
					goto Label3
					// line 218: if final_length == 0:
					πF.SetLineno(218)
				Label1:
					// line 219: snippet = ''
					πF.SetLineno(219)
					µsnippet = ß.ToObject()
					goto Label4
					// line 220: elif final_length == 1:
					πF.SetLineno(220)
				Label2:
					// line 221: a = ord(final[0])
					πF.SetLineno(221)
					πTemp001 = πF.MakeArgs(1)
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µfinal, "final"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, µfinal, πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp006
					if πTemp003, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µa = πTemp006
					// line 222: snippet = table_b2a_base64[(a >> 2 ) & 0x3F] + \
					πF.SetLineno(222)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					if πTemp010, πE = πg.RShift(πF, µa, πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
					if πTemp009, πE = πg.And(πF, πTemp010, πg.NewInt(63).ToObject()); πE != nil {
						continue
					}
					πTemp008 = πTemp009
					if πTemp010, πE = πg.ResolveGlobal(πF, ßtable_b2a_base64); πE != nil {
						continue
					}
					if πTemp009, πE = πg.GetItem(πF, πTemp010, πTemp008); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					if πTemp011, πE = πg.LShift(πF, µa, πg.NewInt(4).ToObject()); πE != nil {
						continue
					}
					if πTemp010, πE = πg.And(πF, πTemp011, πg.NewInt(63).ToObject()); πE != nil {
						continue
					}
					πTemp008 = πTemp010
					if πTemp011, πE = πg.ResolveGlobal(πF, ßtable_b2a_base64); πE != nil {
						continue
					}
					if πTemp010, πE = πg.GetItem(πF, πTemp011, πTemp008); πE != nil {
						continue
					}
					if πTemp006, πE = πg.Add(πF, πTemp009, πTemp010); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Add(πF, πTemp006, πg.NewStr("==").ToObject()); πE != nil {
						continue
					}
					µsnippet = πTemp003
					goto Label4
				Label3:
					// line 225: a = ord(final[0])
					πF.SetLineno(225)
					πTemp001 = πF.MakeArgs(1)
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µfinal, "final"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, µfinal, πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp006
					if πTemp003, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µa = πTemp006
					// line 226: b = ord(final[1])
					πF.SetLineno(226)
					πTemp001 = πF.MakeArgs(1)
					πTemp003 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µfinal, "final"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, µfinal, πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp006
					if πTemp003, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µb = πTemp006
					// line 227: snippet = table_b2a_base64[(a >> 2) & 0x3F] + \
					πF.SetLineno(227)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					if πTemp011, πE = πg.RShift(πF, µa, πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
					if πTemp010, πE = πg.And(πF, πTemp011, πg.NewInt(63).ToObject()); πE != nil {
						continue
					}
					πTemp009 = πTemp010
					if πTemp011, πE = πg.ResolveGlobal(πF, ßtable_b2a_base64); πE != nil {
						continue
					}
					if πTemp010, πE = πg.GetItem(πF, πTemp011, πTemp009); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					if πTemp013, πE = πg.LShift(πF, µa, πg.NewInt(4).ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
						continue
					}
					if πTemp015, πE = πg.RShift(πF, µb, πg.NewInt(4).ToObject()); πE != nil {
						continue
					}
					if πTemp014, πE = πg.And(πF, πTemp015, πg.NewInt(15).ToObject()); πE != nil {
						continue
					}
					if πTemp012, πE = πg.Or(πF, πTemp013, πTemp014); πE != nil {
						continue
					}
					if πTemp011, πE = πg.And(πF, πTemp012, πg.NewInt(63).ToObject()); πE != nil {
						continue
					}
					πTemp009 = πTemp011
					if πTemp012, πE = πg.ResolveGlobal(πF, ßtable_b2a_base64); πE != nil {
						continue
					}
					if πTemp011, πE = πg.GetItem(πF, πTemp012, πTemp009); πE != nil {
						continue
					}
					if πTemp008, πE = πg.Add(πF, πTemp010, πTemp011); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
						continue
					}
					if πTemp011, πE = πg.LShift(πF, µb, πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
					if πTemp010, πE = πg.And(πF, πTemp011, πg.NewInt(63).ToObject()); πE != nil {
						continue
					}
					πTemp009 = πTemp010
					if πTemp011, πE = πg.ResolveGlobal(πF, ßtable_b2a_base64); πE != nil {
						continue
					}
					if πTemp010, πE = πg.GetItem(πF, πTemp011, πTemp009); πE != nil {
						continue
					}
					if πTemp006, πE = πg.Add(πF, πTemp008, πTemp010); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Add(πF, πTemp006, πg.NewStr("=").ToObject()); πE != nil {
						continue
					}
					µsnippet = πTemp003
					goto Label4
				Label4:
					// line 230: return ''.join(result) + snippet + '\n'
					πF.SetLineno(230)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					πTemp001[0] = µresult
					if πTemp008, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
						continue
					}
					if πTemp009, πE = πTemp008.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.CheckLocal(πF, µsnippet, "snippet"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.Add(πF, πTemp009, µsnippet); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Add(πF, πTemp006, πg.NewStr("\n").ToObject()); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßb2a_base64.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 232: def a2b_qp(s, header=False):
			πF.SetLineno(232)
			πTemp006 = make([]πg.Param, 2)
			πTemp006[0] = πg.Param{Name: "s", Def: nil}
			if πTemp009, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			πTemp006[1] = πg.Param{Name: "header", Def: πTemp009}
			πTemp008 = πg.NewFunction(πg.NewCode("a2b_qp", "/usr/lib/python2.7/binascii.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µs *πg.Object = πArgs[0]; _ = µs
				var µheader *πg.Object = πArgs[1]; _ = µheader
				var µinp *πg.Object = πg.UnboundLocal; _ = µinp
				var µodata *πg.Object = πg.UnboundLocal; _ = µodata
				var µch *πg.Object = πg.UnboundLocal; _ = µch
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
				var πTemp008 bool
				_ = πTemp008
				var πTemp009 bool
				_ = πTemp009
				var πTemp010 []*πg.Object
				_ = πTemp010
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 1: goto Label1
					case 2: goto Label2
					case 20: goto Label20
					case 21: goto Label21
					default: panic("unexpected function state")
					}
					// line 233: inp = 0
					πF.SetLineno(233)
					µinp = πg.NewInt(0).ToObject()
					// line 234: odata = []
					πF.SetLineno(234)
					πTemp001 = make([]*πg.Object, 0)
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					µodata = πTemp002
					// line 235: while inp < len(s):
					πF.SetLineno(235)
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
					if πE = πg.CheckLocal(πF, µinp, "inp"); πE != nil {
						continue
					}
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp001[0] = µs
					if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp002, πE = πg.LT(πF, µinp, πTemp006); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πE != nil || !πTemp004 {
						continue
					}
					πF.PushCheckpoint(1)            
					if πE = πg.CheckLocal(πF, µinp, "inp"); πE != nil {
						continue
					}
					πTemp005 = µinp
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, µs, πTemp005); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, πTemp006, πg.NewStr("=").ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label4
					}
					if πE = πg.CheckLocal(πF, µheader, "header"); πE != nil {
						continue
					}
					πTemp002 = µheader
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp004 {
						goto Label5
					}
					if πE = πg.CheckLocal(πF, µinp, "inp"); πE != nil {
						continue
					}
					πTemp006 = µinp
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetItem(πF, µs, πTemp006); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Eq(πF, πTemp007, ß_.ToObject()); πE != nil {
						continue
					}
					πTemp002 = πTemp005
				Label5:
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label6
					}
					goto Label7
					// line 236: if s[inp] == '=':
					πF.SetLineno(236)
				Label4:
					// line 237: inp += 1
					πF.SetLineno(237)
					if πE = πg.CheckLocal(πF, µinp, "inp"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IAdd(πF, µinp, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µinp = πTemp002
					if πE = πg.CheckLocal(πF, µinp, "inp"); πE != nil {
						continue
					}
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp001[0] = µs
					if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp002, πE = πg.GE(πF, µinp, πTemp006); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label9
					}
					goto Label10
					// line 238: if inp >= len(s):
					πF.SetLineno(238)
				Label9:
					// line 239: break
					πF.SetLineno(239)
					πTemp003 = true
					continue
					goto Label10
				Label10:
					if πE = πg.CheckLocal(πF, µinp, "inp"); πE != nil {
						continue
					}
					πTemp006 = µinp
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetItem(πF, µs, πTemp006); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Eq(πF, πTemp007, πg.NewStr("\n").ToObject()); πE != nil {
						continue
					}
					πTemp002 = πTemp005
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label11
					}
					if πE = πg.CheckLocal(πF, µinp, "inp"); πE != nil {
						continue
					}
					πTemp006 = µinp
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetItem(πF, µs, πTemp006); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Eq(πF, πTemp007, πg.NewStr("\r").ToObject()); πE != nil {
						continue
					}
					πTemp002 = πTemp005
				Label11:
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label12
					}
					if πE = πg.CheckLocal(πF, µinp, "inp"); πE != nil {
						continue
					}
					πTemp005 = µinp
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, µs, πTemp005); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, πTemp006, πg.NewStr("=").ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label13
					}
					if πE = πg.CheckLocal(πF, µinp, "inp"); πE != nil {
						continue
					}
					πTemp006 = µinp
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetItem(πF, µs, πTemp006); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveGlobal(πF, ßhex_numbers); πE != nil {
						continue
					}
					if πTemp008, πE = πg.Contains(πF, πTemp006, πTemp007); πE != nil {
						continue
					}
					πTemp005 = πg.GetBool(πTemp008).ToObject()
					πTemp002 = πTemp005
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp004 {
						goto Label14
					}
					if πE = πg.CheckLocal(πF, µinp, "inp"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.Add(πF, µinp, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp006 = πTemp007
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetItem(πF, µs, πTemp006); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveGlobal(πF, ßhex_numbers); πE != nil {
						continue
					}
					if πTemp008, πE = πg.Contains(πF, πTemp006, πTemp007); πE != nil {
						continue
					}
					πTemp005 = πg.GetBool(πTemp008).ToObject()
					πTemp002 = πTemp005
				Label14:
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label15
					}
					goto Label16
					// line 241: if (s[inp] == '\n') or (s[inp] == '\r'):
					πF.SetLineno(241)
				Label12:
					if πE = πg.CheckLocal(πF, µinp, "inp"); πE != nil {
						continue
					}
					πTemp005 = µinp
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, µs, πTemp005); πE != nil {
						continue
					}
					if πTemp002, πE = πg.NE(πF, πTemp006, πg.NewStr("\n").ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label18
					}
					goto Label19
					// line 242: if s[inp] != '\n':
					πF.SetLineno(242)
				Label18:
					// line 243: while inp < len(s) and s[inp] != '\n':
					πF.SetLineno(243)
					πF.PushCheckpoint(21)
					πTemp004 = false
				Label20:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp004 {
						πF.PopCheckpoint()
						goto Label22
					}
					if πE = πg.CheckLocal(πF, µinp, "inp"); πE != nil {
						continue
					}
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp001[0] = µs
					if πTemp006, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp005, πE = πg.LT(πF, µinp, πTemp007); πE != nil {
						continue
					}
					πTemp002 = πTemp005
					if πTemp009, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp009 {
						goto Label23
					}
					if πE = πg.CheckLocal(πF, µinp, "inp"); πE != nil {
						continue
					}
					πTemp006 = µinp
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetItem(πF, µs, πTemp006); πE != nil {
						continue
					}
					if πTemp005, πE = πg.NE(πF, πTemp007, πg.NewStr("\n").ToObject()); πE != nil {
						continue
					}
					πTemp002 = πTemp005
				Label23:
					if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πE != nil || !πTemp008 {
						continue
					}
					πF.PushCheckpoint(20)            
					// line 244: inp += 1
					πF.SetLineno(244)
					if πE = πg.CheckLocal(πF, µinp, "inp"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IAdd(πF, µinp, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µinp = πTemp002
					continue
				Label21:
					if πE != nil || πR != nil {
						continue
					}
				Label22:
					goto Label19
				Label19:
					if πE = πg.CheckLocal(πF, µinp, "inp"); πE != nil {
						continue
					}
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp001[0] = µs
					if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp002, πE = πg.LT(πF, µinp, πTemp006); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label24
					}
					goto Label25
					// line 245: if inp < len(s):
					πF.SetLineno(245)
				Label24:
					// line 246: inp += 1
					πF.SetLineno(246)
					if πE = πg.CheckLocal(πF, µinp, "inp"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IAdd(πF, µinp, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µinp = πTemp002
					goto Label25
				Label25:
					goto Label17
					// line 247: elif s[inp] == '=':
					πF.SetLineno(247)
				Label13:
					// line 249: odata.append('=')
					πF.SetLineno(249)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("=").ToObject()
					if πE = πg.CheckLocal(πF, µodata, "odata"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µodata, ßappend, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 250: inp += 1
					πF.SetLineno(250)
					if πE = πg.CheckLocal(πF, µinp, "inp"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IAdd(πF, µinp, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µinp = πTemp002
					goto Label17
					// line 251: elif s[inp] in hex_numbers and s[inp + 1] in hex_numbers:
					πF.SetLineno(251)
				Label15:
					// line 252: ch = chr(int(s[inp:inp+2], 16))
					πF.SetLineno(252)
					πTemp001 = πF.MakeArgs(1)
					πTemp010 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µinp, "inp"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µinp, "inp"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Add(πF, µinp, πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{µinp, πTemp005, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µs, πTemp002); πE != nil {
						continue
					}
					πTemp010[0] = πTemp005
					πTemp010[1] = πg.NewInt(16).ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, πTemp010, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp010)
					πTemp001[0] = πTemp005
					if πTemp002, πE = πg.ResolveGlobal(πF, ßchr); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µch = πTemp005
					// line 253: inp += 2
					πF.SetLineno(253)
					if πE = πg.CheckLocal(πF, µinp, "inp"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IAdd(πF, µinp, πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
					µinp = πTemp002
					// line 254: odata.append(ch)
					πF.SetLineno(254)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µch, "ch"); πE != nil {
						continue
					}
					πTemp001[0] = µch
					if πE = πg.CheckLocal(πF, µodata, "odata"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µodata, ßappend, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					goto Label17
				Label16:
					// line 256: odata.append('=')
					πF.SetLineno(256)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("=").ToObject()
					if πE = πg.CheckLocal(πF, µodata, "odata"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µodata, ßappend, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					goto Label17
				Label17:
					goto Label8
					// line 257: elif header and s[inp] == '_':
					πF.SetLineno(257)
				Label6:
					// line 258: odata.append(' ')
					πF.SetLineno(258)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr(" ").ToObject()
					if πE = πg.CheckLocal(πF, µodata, "odata"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µodata, ßappend, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 259: inp += 1
					πF.SetLineno(259)
					if πE = πg.CheckLocal(πF, µinp, "inp"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IAdd(πF, µinp, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µinp = πTemp002
					goto Label8
				Label7:
					// line 261: odata.append(s[inp])
					πF.SetLineno(261)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µinp, "inp"); πE != nil {
						continue
					}
					πTemp002 = µinp
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µs, πTemp002); πE != nil {
						continue
					}
					πTemp001[0] = πTemp005
					if πE = πg.CheckLocal(πF, µodata, "odata"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µodata, ßappend, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 262: inp += 1
					πF.SetLineno(262)
					if πE = πg.CheckLocal(πF, µinp, "inp"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IAdd(πF, µinp, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µinp = πTemp002
					goto Label8
				Label8:
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
					// line 263: return ''.join(odata)
					πF.SetLineno(263)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µodata, "odata"); πE != nil {
						continue
					}
					πTemp001[0] = µodata
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
			if πE = πF.Globals().SetItem(πF, ßa2b_qp.ToObject(), πTemp008); πE != nil {
				continue
			}
			// line 265: def b2a_qp(data, quotetabs=False, istext=True, header=False):
			πF.SetLineno(265)
			πTemp006 = make([]πg.Param, 4)
			πTemp006[0] = πg.Param{Name: "data", Def: nil}
			if πTemp010, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			πTemp006[1] = πg.Param{Name: "quotetabs", Def: πTemp010}
			if πTemp010, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
				continue
			}
			πTemp006[2] = πg.Param{Name: "istext", Def: πTemp010}
			if πTemp010, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			πTemp006[3] = πg.Param{Name: "header", Def: πTemp010}
			πTemp009 = πg.NewFunction(πg.NewCode("b2a_qp", "/usr/lib/python2.7/binascii.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µdata *πg.Object = πArgs[0]; _ = µdata
				var µquotetabs *πg.Object = πArgs[1]; _ = µquotetabs
				var µistext *πg.Object = πArgs[2]; _ = µistext
				var µheader *πg.Object = πArgs[3]; _ = µheader
				var µMAXLINESIZE *πg.Object = πg.UnboundLocal; _ = µMAXLINESIZE
				var µlf *πg.Object = πg.UnboundLocal; _ = µlf
				var µcrlf *πg.Object = πg.UnboundLocal; _ = µcrlf
				var µinp *πg.Object = πg.UnboundLocal; _ = µinp
				var µlinelen *πg.Object = πg.UnboundLocal; _ = µlinelen
				var µodata *πg.Object = πg.UnboundLocal; _ = µodata
				var µc *πg.Object = πg.UnboundLocal; _ = µc
				var µch *πg.Object = πg.UnboundLocal; _ = µch
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
				var πTemp007 bool
				_ = πTemp007
				var πTemp008 bool
				_ = πTemp008
				var πTemp009 bool
				_ = πTemp009
				var πTemp010 *πg.Object
				_ = πTemp010
				var πTemp011 *πg.Object
				_ = πTemp011
				var πTemp012 *πg.Object
				_ = πTemp012
				var πTemp013 bool
				_ = πTemp013
				var πTemp014 bool
				_ = πTemp014
				var πTemp015 []*πg.Object
				_ = πTemp015
				var πTemp016 []*πg.Object
				_ = πTemp016
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 2: goto Label2
					case 3: goto Label3
					default: panic("unexpected function state")
					}
					// line 266: """quotetabs=True means that tab and space characters are always
					πF.SetLineno(266)
					// line 272: MAXLINESIZE = 76
					πF.SetLineno(272)
					µMAXLINESIZE = πg.NewInt(76).ToObject()
					// line 275: lf = data.find('\n')
					πF.SetLineno(275)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("\n").ToObject()
					if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µdata, ßfind, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µlf = πTemp003
					// line 276: crlf = lf > 0 and data[lf-1] == '\r'
					πF.SetLineno(276)
					if πE = πg.CheckLocal(πF, µlf, "lf"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GT(πF, µlf, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					πTemp002 = πTemp003
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp004 {
						goto Label1
					}
					if πE = πg.CheckLocal(πF, µlf, "lf"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.Sub(πF, µlf, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp005 = πTemp006
					if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, µdata, πTemp005); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Eq(πF, πTemp006, πg.NewStr("\r").ToObject()); πE != nil {
						continue
					}
					πTemp002 = πTemp003
				Label1:
					µcrlf = πTemp002
					// line 278: inp = 0
					πF.SetLineno(278)
					µinp = πg.NewInt(0).ToObject()
					// line 279: linelen = 0
					πF.SetLineno(279)
					µlinelen = πg.NewInt(0).ToObject()
					// line 280: odata = []
					πF.SetLineno(280)
					πTemp001 = make([]*πg.Object, 0)
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					µodata = πTemp002
					// line 281: while inp < len(data):
					πF.SetLineno(281)
					πF.PushCheckpoint(3)
					πTemp004 = false
				Label2:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp004 {
						πF.PopCheckpoint()
						goto Label4
					}
					if πE = πg.CheckLocal(πF, µinp, "inp"); πE != nil {
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
					if πTemp002, πE = πg.LT(πF, µinp, πTemp005); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πE != nil || !πTemp007 {
						continue
					}
					πF.PushCheckpoint(2)            
					// line 282: c = data[inp]
					πF.SetLineno(282)
					if πE = πg.CheckLocal(πF, µinp, "inp"); πE != nil {
						continue
					}
					πTemp002 = µinp
					if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µdata, πTemp002); πE != nil {
						continue
					}
					µc = πTemp003
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GT(πF, µc, πg.NewStr("~").ToObject()); πE != nil {
						continue
					}
					πTemp002 = πTemp003
					if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label5
					}
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Eq(πF, µc, πg.NewStr("=").ToObject()); πE != nil {
						continue
					}
					πTemp002 = πTemp003
					if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label5
					}
					if πE = πg.CheckLocal(πF, µheader, "header"); πE != nil {
						continue
					}
					πTemp003 = µheader
					if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if !πTemp008 {
						goto Label6
					}
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Eq(πF, µc, ß_.ToObject()); πE != nil {
						continue
					}
					πTemp003 = πTemp005
				Label6:
					πTemp002 = πTemp003
					if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label5
					}
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Eq(πF, µc, πg.NewStr(".").ToObject()); πE != nil {
						continue
					}
					πTemp003 = πTemp005
					if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if !πTemp008 {
						goto Label7
					}
					if πE = πg.CheckLocal(πF, µlinelen, "linelen"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Eq(πF, µlinelen, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					πTemp003 = πTemp005
					if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if !πTemp008 {
						goto Label7
					}
					if πE = πg.CheckLocal(πF, µinp, "inp"); πE != nil {
						continue
					}
					if πTemp010, πE = πg.Add(πF, µinp, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
						continue
					}
					πTemp001[0] = µdata
					if πTemp011, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp012, πE = πTemp011.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp006, πE = πg.Eq(πF, πTemp010, πTemp012); πE != nil {
						continue
					}
					πTemp005 = πTemp006
					if πTemp009, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					if πTemp009 {
						goto Label8
					}
					if πE = πg.CheckLocal(πF, µinp, "inp"); πE != nil {
						continue
					}
					if πTemp011, πE = πg.Add(πF, µinp, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp010 = πTemp011
					if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
						continue
					}
					if πTemp011, πE = πg.GetItem(πF, µdata, πTemp010); πE != nil {
						continue
					}
					if πTemp006, πE = πg.Eq(πF, πTemp011, πg.NewStr("\n").ToObject()); πE != nil {
						continue
					}
					πTemp005 = πTemp006
					if πTemp009, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					if πTemp009 {
						goto Label8
					}
					if πE = πg.CheckLocal(πF, µinp, "inp"); πE != nil {
						continue
					}
					if πTemp011, πE = πg.Add(πF, µinp, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp010 = πTemp011
					if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
						continue
					}
					if πTemp011, πE = πg.GetItem(πF, µdata, πTemp010); πE != nil {
						continue
					}
					if πTemp006, πE = πg.Eq(πF, πTemp011, πg.NewStr("\r").ToObject()); πE != nil {
						continue
					}
					πTemp005 = πTemp006
				Label8:
					πTemp003 = πTemp005
				Label7:
					πTemp002 = πTemp003
					if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label5
					}
					if πE = πg.CheckLocal(πF, µistext, "istext"); πE != nil {
						continue
					}
					if πTemp009, πE = πg.IsTrue(πF, µistext); πE != nil {
						continue
					}
					πTemp005 = πg.GetBool(!πTemp009).ToObject()
					πTemp003 = πTemp005
					if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if !πTemp008 {
						goto Label9
					}
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.Eq(πF, µc, πg.NewStr("\r").ToObject()); πE != nil {
						continue
					}
					πTemp005 = πTemp006
					if πTemp009, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					if πTemp009 {
						goto Label10
					}
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.Eq(πF, µc, πg.NewStr("\n").ToObject()); πE != nil {
						continue
					}
					πTemp005 = πTemp006
				Label10:
					πTemp003 = πTemp005
				Label9:
					πTemp002 = πTemp003
					if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label5
					}
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.Eq(πF, µc, πg.NewStr("\t").ToObject()); πE != nil {
						continue
					}
					πTemp005 = πTemp006
					if πTemp009, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					if πTemp009 {
						goto Label12
					}
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.Eq(πF, µc, πg.NewStr(" ").ToObject()); πE != nil {
						continue
					}
					πTemp005 = πTemp006
				Label12:
					πTemp003 = πTemp005
					if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if !πTemp008 {
						goto Label11
					}
					if πE = πg.CheckLocal(πF, µinp, "inp"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.Add(πF, µinp, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
						continue
					}
					πTemp001[0] = µdata
					if πTemp010, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp011, πE = πTemp010.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp005, πE = πg.Eq(πF, πTemp006, πTemp011); πE != nil {
						continue
					}
					πTemp003 = πTemp005
				Label11:
					πTemp002 = πTemp003
					if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label5
					}
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.LE(πF, µc, πg.NewStr(" ").ToObject()); πE != nil {
						continue
					}
					πTemp003 = πTemp005
					if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if !πTemp008 {
						goto Label13
					}
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.NE(πF, µc, πg.NewStr("\r").ToObject()); πE != nil {
						continue
					}
					πTemp003 = πTemp005
					if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if !πTemp008 {
						goto Label13
					}
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.NE(πF, µc, πg.NewStr("\n").ToObject()); πE != nil {
						continue
					}
					πTemp003 = πTemp005
					if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if !πTemp008 {
						goto Label13
					}
					if πE = πg.CheckLocal(πF, µquotetabs, "quotetabs"); πE != nil {
						continue
					}
					πTemp005 = µquotetabs
					if πTemp009, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					if πTemp009 {
						goto Label14
					}
					if πE = πg.CheckLocal(πF, µquotetabs, "quotetabs"); πE != nil {
						continue
					}
					if πTemp014, πE = πg.IsTrue(πF, µquotetabs); πE != nil {
						continue
					}
					πTemp010 = πg.GetBool(!πTemp014).ToObject()
					πTemp006 = πTemp010
					if πTemp013, πE = πg.IsTrue(πF, πTemp006); πE != nil {
						continue
					}
					if !πTemp013 {
						goto Label15
					}
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					if πTemp011, πE = πg.NE(πF, µc, πg.NewStr("\t").ToObject()); πE != nil {
						continue
					}
					πTemp010 = πTemp011
					if πTemp014, πE = πg.IsTrue(πF, πTemp010); πE != nil {
						continue
					}
					if !πTemp014 {
						goto Label16
					}
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					if πTemp011, πE = πg.NE(πF, µc, πg.NewStr(" ").ToObject()); πE != nil {
						continue
					}
					πTemp010 = πTemp011
				Label16:
					πTemp006 = πTemp010
				Label15:
					πTemp005 = πTemp006
				Label14:
					πTemp003 = πTemp005
				Label13:
					πTemp002 = πTemp003
				Label5:
					if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label17
					}
					if πE = πg.CheckLocal(πF, µistext, "istext"); πE != nil {
						continue
					}
					πTemp002 = µistext
					if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp007 {
						goto Label18
					}
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Eq(πF, µc, πg.NewStr("\n").ToObject()); πE != nil {
						continue
					}
					πTemp003 = πTemp005
					if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label19
					}
					if πE = πg.CheckLocal(πF, µinp, "inp"); πE != nil {
						continue
					}
					if πTemp010, πE = πg.Add(πF, µinp, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
						continue
					}
					πTemp001[0] = µdata
					if πTemp011, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp012, πE = πTemp011.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp006, πE = πg.LT(πF, πTemp010, πTemp012); πE != nil {
						continue
					}
					πTemp005 = πTemp006
					if πTemp009, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					if !πTemp009 {
						goto Label20
					}
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.Eq(πF, µc, πg.NewStr("\r").ToObject()); πE != nil {
						continue
					}
					πTemp005 = πTemp006
					if πTemp009, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					if !πTemp009 {
						goto Label20
					}
					if πE = πg.CheckLocal(πF, µinp, "inp"); πE != nil {
						continue
					}
					if πTemp011, πE = πg.Add(πF, µinp, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp010 = πTemp011
					if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
						continue
					}
					if πTemp011, πE = πg.GetItem(πF, µdata, πTemp010); πE != nil {
						continue
					}
					if πTemp006, πE = πg.Eq(πF, πTemp011, πg.NewStr("\n").ToObject()); πE != nil {
						continue
					}
					πTemp005 = πTemp006
				Label20:
					πTemp003 = πTemp005
				Label19:
					πTemp002 = πTemp003
				Label18:
					if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label21
					}
					goto Label22
					// line 283: if (c > '~' or
					πF.SetLineno(283)
				Label17:
					// line 293: linelen += 3
					πF.SetLineno(293)
					if πE = πg.CheckLocal(πF, µlinelen, "linelen"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IAdd(πF, µlinelen, πg.NewInt(3).ToObject()); πE != nil {
						continue
					}
					µlinelen = πTemp002
					if πE = πg.CheckLocal(πF, µlinelen, "linelen"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µMAXLINESIZE, "MAXLINESIZE"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GE(πF, µlinelen, µMAXLINESIZE); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label24
					}
					goto Label25
					// line 294: if linelen >= MAXLINESIZE:
					πF.SetLineno(294)
				Label24:
					// line 295: odata.append('=')
					πF.SetLineno(295)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("=").ToObject()
					if πE = πg.CheckLocal(πF, µodata, "odata"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µodata, ßappend, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.CheckLocal(πF, µcrlf, "crlf"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsTrue(πF, µcrlf); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label26
					}
					goto Label27
					// line 296: if crlf: odata.append('\r')
					πF.SetLineno(296)
				Label26:
					// line 296: if crlf: odata.append('\r')
					πF.SetLineno(296)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("\r").ToObject()
					if πE = πg.CheckLocal(πF, µodata, "odata"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µodata, ßappend, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					goto Label27
				Label27:
					// line 297: odata.append('\n')
					πF.SetLineno(297)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("\n").ToObject()
					if πE = πg.CheckLocal(πF, µodata, "odata"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µodata, ßappend, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 298: linelen = 3
					πF.SetLineno(298)
					µlinelen = πg.NewInt(3).ToObject()
					goto Label25
				Label25:
					// line 299: odata.append('=' + two_hex_digits(ord(c)))
					πF.SetLineno(299)
					πTemp001 = πF.MakeArgs(1)
					πTemp015 = πF.MakeArgs(1)
					πTemp016 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					πTemp016[0] = µc
					if πTemp003, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp016, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp016)
					πTemp015[0] = πTemp005
					if πTemp003, πE = πg.ResolveGlobal(πF, ßtwo_hex_digits); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp015, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp015)
					if πTemp002, πE = πg.Add(πF, πg.NewStr("=").ToObject(), πTemp005); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πE = πg.CheckLocal(πF, µodata, "odata"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µodata, ßappend, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 300: inp += 1
					πF.SetLineno(300)
					if πE = πg.CheckLocal(πF, µinp, "inp"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IAdd(πF, µinp, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µinp = πTemp002
					goto Label23
					// line 302: if (istext and
					πF.SetLineno(302)
				Label21:
					// line 305: linelen = 0
					πF.SetLineno(305)
					µlinelen = πg.NewInt(0).ToObject()
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µodata, "odata"); πE != nil {
						continue
					}
					πTemp001[0] = µodata
					if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp003, πE = πg.GT(πF, πTemp006, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					πTemp002 = πTemp003
					if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp007 {
						goto Label28
					}
					if πTemp010, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp006 = πTemp010
					if πE = πg.CheckLocal(πF, µodata, "odata"); πE != nil {
						continue
					}
					if πTemp010, πE = πg.GetItem(πF, µodata, πTemp006); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Eq(πF, πTemp010, πg.NewStr(" ").ToObject()); πE != nil {
						continue
					}
					πTemp003 = πTemp005
					if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label29
					}
					if πTemp010, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp006 = πTemp010
					if πE = πg.CheckLocal(πF, µodata, "odata"); πE != nil {
						continue
					}
					if πTemp010, πE = πg.GetItem(πF, µodata, πTemp006); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Eq(πF, πTemp010, πg.NewStr("\t").ToObject()); πE != nil {
						continue
					}
					πTemp003 = πTemp005
				Label29:
					πTemp002 = πTemp003
				Label28:
					if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label30
					}
					goto Label31
					// line 307: if (len(odata) > 0 and
					πF.SetLineno(307)
				Label30:
					// line 309: ch = ord(odata[-1])
					πF.SetLineno(309)
					πTemp001 = πF.MakeArgs(1)
					if πTemp003, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp002 = πTemp003
					if πE = πg.CheckLocal(πF, µodata, "odata"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µodata, πTemp002); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					if πTemp002, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µch = πTemp003
					// line 310: odata[-1] = '='
					πF.SetLineno(310)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewStr("=").ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µodata, "odata"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp003 = πTemp005
					if πE = πg.SetItem(πF, µodata, πTemp003, πTemp002); πE != nil {
						continue
					}
					// line 311: odata.append(two_hex_digits(ch))
					πF.SetLineno(311)
					πTemp001 = πF.MakeArgs(1)
					πTemp015 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µch, "ch"); πE != nil {
						continue
					}
					πTemp015[0] = µch
					if πTemp002, πE = πg.ResolveGlobal(πF, ßtwo_hex_digits); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp015, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp015)
					πTemp001[0] = πTemp003
					if πE = πg.CheckLocal(πF, µodata, "odata"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µodata, ßappend, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					goto Label31
				Label31:
					if πE = πg.CheckLocal(πF, µcrlf, "crlf"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsTrue(πF, µcrlf); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label32
					}
					goto Label33
					// line 313: if crlf: odata.append('\r')
					πF.SetLineno(313)
				Label32:
					// line 313: if crlf: odata.append('\r')
					πF.SetLineno(313)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("\r").ToObject()
					if πE = πg.CheckLocal(πF, µodata, "odata"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µodata, ßappend, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					goto Label33
				Label33:
					// line 314: odata.append('\n')
					πF.SetLineno(314)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("\n").ToObject()
					if πE = πg.CheckLocal(πF, µodata, "odata"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µodata, ßappend, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µc, πg.NewStr("\r").ToObject()); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label34
					}
					goto Label35
					// line 315: if c == '\r':
					πF.SetLineno(315)
				Label34:
					// line 316: inp += 2
					πF.SetLineno(316)
					if πE = πg.CheckLocal(πF, µinp, "inp"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IAdd(πF, µinp, πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
					µinp = πTemp002
					goto Label36
				Label35:
					// line 318: inp += 1
					πF.SetLineno(318)
					if πE = πg.CheckLocal(πF, µinp, "inp"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IAdd(πF, µinp, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µinp = πTemp002
					goto Label36
				Label36:
					goto Label23
				Label22:
					if πE = πg.CheckLocal(πF, µinp, "inp"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Add(πF, µinp, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
						continue
					}
					πTemp001[0] = µdata
					if πTemp006, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp010, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp003, πE = πg.LT(πF, πTemp005, πTemp010); πE != nil {
						continue
					}
					πTemp002 = πTemp003
					if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp007 {
						goto Label37
					}
					if πE = πg.CheckLocal(πF, µinp, "inp"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.Add(πF, µinp, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp005 = πTemp006
					if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, µdata, πTemp005); πE != nil {
						continue
					}
					if πTemp003, πE = πg.NE(πF, πTemp006, πg.NewStr("\n").ToObject()); πE != nil {
						continue
					}
					πTemp002 = πTemp003
					if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp007 {
						goto Label37
					}
					if πE = πg.CheckLocal(πF, µlinelen, "linelen"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Add(πF, µlinelen, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µMAXLINESIZE, "MAXLINESIZE"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GE(πF, πTemp005, µMAXLINESIZE); πE != nil {
						continue
					}
					πTemp002 = πTemp003
				Label37:
					if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label38
					}
					goto Label39
					// line 320: if (inp + 1 < len(data) and
					πF.SetLineno(320)
				Label38:
					// line 323: odata.append('=')
					πF.SetLineno(323)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("=").ToObject()
					if πE = πg.CheckLocal(πF, µodata, "odata"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µodata, ßappend, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.CheckLocal(πF, µcrlf, "crlf"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsTrue(πF, µcrlf); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label40
					}
					goto Label41
					// line 324: if crlf: odata.append('\r')
					πF.SetLineno(324)
				Label40:
					// line 324: if crlf: odata.append('\r')
					πF.SetLineno(324)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("\r").ToObject()
					if πE = πg.CheckLocal(πF, µodata, "odata"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µodata, ßappend, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					goto Label41
				Label41:
					// line 325: odata.append('\n')
					πF.SetLineno(325)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("\n").ToObject()
					if πE = πg.CheckLocal(πF, µodata, "odata"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µodata, ßappend, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 326: linelen = 0
					πF.SetLineno(326)
					µlinelen = πg.NewInt(0).ToObject()
					goto Label39
				Label39:
					// line 328: linelen += 1
					πF.SetLineno(328)
					if πE = πg.CheckLocal(πF, µlinelen, "linelen"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IAdd(πF, µlinelen, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µlinelen = πTemp002
					if πE = πg.CheckLocal(πF, µheader, "header"); πE != nil {
						continue
					}
					πTemp002 = µheader
					if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp007 {
						goto Label42
					}
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Eq(πF, µc, πg.NewStr(" ").ToObject()); πE != nil {
						continue
					}
					πTemp002 = πTemp003
				Label42:
					if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label43
					}
					goto Label44
					// line 329: if header and c == ' ':
					πF.SetLineno(329)
				Label43:
					// line 330: c = '_'
					πF.SetLineno(330)
					µc = ß_.ToObject()
					goto Label44
				Label44:
					// line 331: odata.append(c)
					πF.SetLineno(331)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					πTemp001[0] = µc
					if πE = πg.CheckLocal(πF, µodata, "odata"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µodata, ßappend, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 332: inp += 1
					πF.SetLineno(332)
					if πE = πg.CheckLocal(πF, µinp, "inp"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IAdd(πF, µinp, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µinp = πTemp002
					goto Label23
				Label23:
					continue
				Label3:
					if πE != nil || πR != nil {
						continue
					}
				Label4:
					// line 333: return ''.join(odata)
					πF.SetLineno(333)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µodata, "odata"); πE != nil {
						continue
					}
					πTemp001[0] = µodata
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
			if πE = πF.Globals().SetItem(πF, ßb2a_qp.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 266: """quotetabs=True means that tab and space characters are always
			πF.SetLineno(266)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp010}, πg.NewStr("quotetabs=True means that tab and space characters are always\n       quoted.\n       istext=False means that \r and \n are treated as regular characters\n       header=True encodes space characters with '_' and requires\n       real '_' characters to be quoted.\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp011, πE = πg.ResolveGlobal(πF, ßb2a_qp); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp011, ß__doc__, πTemp010); πE != nil {
				continue
			}
			// line 335: hex_numbers = '0123456789ABCDEF'
			πF.SetLineno(335)
			if πE = πF.Globals().SetItem(πF, ßhex_numbers.ToObject(), ß0123456789ABCDEF.ToObject()); πE != nil {
				continue
			}
			// line 336: def hex(n):
			πF.SetLineno(336)
			πTemp006 = make([]πg.Param, 1)
			πTemp006[0] = πg.Param{Name: "n", Def: nil}
			πTemp010 = πg.NewFunction(πg.NewCode("hex", "/usr/lib/python2.7/binascii.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µn *πg.Object = πArgs[0]; _ = µn
				var µsign *πg.Object = πg.UnboundLocal; _ = µsign
				var µarr *πg.Object = πg.UnboundLocal; _ = µarr
				var µhex_gen *πg.Object = πg.UnboundLocal; _ = µhex_gen
				var µnibble *πg.Object = πg.UnboundLocal; _ = µnibble
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 bool
				_ = πTemp002
				var πTemp003 []*πg.Object
				_ = πTemp003
				var πTemp004 []πg.Param
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 bool
				_ = πTemp007
				var πTemp008 *πg.Object
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
					case 6: goto Label6
					case 7: goto Label7
					default: panic("unexpected function state")
					}
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, µn, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label1
					}
					goto Label2
					// line 337: if n == 0:
					πF.SetLineno(337)
				Label1:
					// line 338: return '0'
					πF.SetLineno(338)
					πR = ß0.ToObject()
					continue
					goto Label2
				Label2:
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.LT(πF, µn, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label3
					}
					goto Label4
					// line 340: if n < 0:
					πF.SetLineno(340)
				Label3:
					// line 341: n = -n
					πF.SetLineno(341)
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Neg(πF, µn); πE != nil {
						continue
					}
					µn = πTemp001
					// line 342: sign = '-'
					πF.SetLineno(342)
					µsign = πg.NewStr("-").ToObject()
					goto Label5
				Label4:
					// line 344: sign = ''
					πF.SetLineno(344)
					µsign = ß.ToObject()
					goto Label5
				Label5:
					// line 345: arr = []
					πF.SetLineno(345)
					πTemp003 = make([]*πg.Object, 0)
					πTemp001 = πg.NewList(πTemp003...).ToObject()
					µarr = πTemp001
					// line 347: def hex_gen(n):
					πF.SetLineno(347)
					πTemp004 = make([]πg.Param, 1)
					πTemp004[0] = πg.Param{Name: "n", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("hex_gen", "/usr/lib/python2.7/binascii.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µn *πg.Object = πArgs[0]; _ = µn
						var πTemp001 bool
						_ = πTemp001
						var πTemp002 bool
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
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
								// line 348: """ Yield a nibble at a time. """
								πF.SetLineno(348)
								// line 349: while n:
								πF.SetLineno(349)
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
								if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.IsTrue(πF, µn); πE != nil {
									continue
								}
								if πE != nil || !πTemp002 {
									continue
								}
								πF.PushCheckpoint(1)            
								// line 350: yield n % 0x10
								πF.SetLineno(350)
								if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.Mod(πF, µn, πg.NewInt(16).ToObject()); πE != nil {
									continue
								}
								πF.PushCheckpoint(4)
								return πTemp003, nil
							Label4:
								πTemp004 = πSent
								// line 351: n = n / 0x10
								πF.SetLineno(351)
								if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
									continue
								}
								if πTemp004, πE = πg.Div(πF, µn, πg.NewInt(16).ToObject()); πE != nil {
									continue
								}
								µn = πTemp004
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
					µhex_gen = πTemp001
					// line 348: """ Yield a nibble at a time. """
					πF.SetLineno(348)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πg.NewStr(" Yield a nibble at a time. ").ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µhex_gen, "hex_gen"); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, µhex_gen, ß__doc__, πTemp005); πE != nil {
						continue
					}
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					πTemp003[0] = µn
					if πE = πg.CheckLocal(πF, µhex_gen, "hex_gen"); πE != nil {
						continue
					}
					if πTemp006, πE = µhex_gen.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp005, πE = πg.Iter(πF, πTemp006); πE != nil {
						continue
					}
					πF.PushCheckpoint(7)
					πTemp002 = false
				Label6:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp002 {
						πF.PopCheckpoint()
						goto Label8
					}
					if πTemp006, πE = πg.Next(πF, πTemp005); πE != nil {
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
						µnibble = πTemp006
					}
					if πE != nil || !πTemp007 {
						continue
					}
					πF.PushCheckpoint(6)            
					// line 354: arr = [hex_numbers[nibble]] + arr
					πF.SetLineno(354)
					πTemp003 = make([]*πg.Object, 1)
					if πE = πg.CheckLocal(πF, µnibble, "nibble"); πE != nil {
						continue
					}
					πTemp008 = µnibble
					if πTemp010, πE = πg.ResolveGlobal(πF, ßhex_numbers); πE != nil {
						continue
					}
					if πTemp009, πE = πg.GetItem(πF, πTemp010, πTemp008); πE != nil {
						continue
					}
					πTemp003[0] = πTemp009
					πTemp008 = πg.NewList(πTemp003...).ToObject()
					if πE = πg.CheckLocal(πF, µarr, "arr"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.Add(πF, πTemp008, µarr); πE != nil {
						continue
					}
					µarr = πTemp006
					continue
				Label7:
					if πE != nil || πR != nil {
						continue
					}
				Label8:
					// line 355: return sign + ''.join(arr)
					πF.SetLineno(355)
					if πE = πg.CheckLocal(πF, µsign, "sign"); πE != nil {
						continue
					}
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µarr, "arr"); πE != nil {
						continue
					}
					πTemp003[0] = µarr
					if πTemp006, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp006.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp005, πE = πg.Add(πF, µsign, πTemp008); πE != nil {
						continue
					}
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
			if πE = πF.Globals().SetItem(πF, ßhex.ToObject(), πTemp010); πE != nil {
				continue
			}
			// line 357: def two_hex_digits(n):
			πF.SetLineno(357)
			πTemp006 = make([]πg.Param, 1)
			πTemp006[0] = πg.Param{Name: "n", Def: nil}
			πTemp011 = πg.NewFunction(πg.NewCode("two_hex_digits", "/usr/lib/python2.7/binascii.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µn *πg.Object = πArgs[0]; _ = µn
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
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 358: return hex_numbers[n / 0x10] + hex_numbers[n % 0x10]
					πF.SetLineno(358)
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Div(πF, µn, πg.NewInt(16).ToObject()); πE != nil {
						continue
					}
					πTemp002 = πTemp003
					if πTemp004, πE = πg.ResolveGlobal(πF, ßhex_numbers); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, πTemp004, πTemp002); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Mod(πF, µn, πg.NewInt(16).ToObject()); πE != nil {
						continue
					}
					πTemp002 = πTemp004
					if πTemp005, πE = πg.ResolveGlobal(πF, ßhex_numbers); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, πTemp005, πTemp002); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, πTemp003, πTemp004); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßtwo_hex_digits.ToObject(), πTemp011); πE != nil {
				continue
			}
			// line 361: def strhex_to_int(s):
			πF.SetLineno(361)
			πTemp006 = make([]πg.Param, 1)
			πTemp006[0] = πg.Param{Name: "s", Def: nil}
			πTemp012 = πg.NewFunction(πg.NewCode("strhex_to_int", "/usr/lib/python2.7/binascii.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µs *πg.Object = πArgs[0]; _ = µs
				var µi *πg.Object = πg.UnboundLocal; _ = µi
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
					case 1: goto Label1
					case 2: goto Label2
					default: panic("unexpected function state")
					}
					// line 362: i = 0
					πF.SetLineno(362)
					µi = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Iter(πF, µs); πE != nil {
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
					// line 364: i = i * 0x10 + hex_numbers.index(c)
					πF.SetLineno(364)
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Mul(πF, µi, πg.NewInt(16).ToObject()); πE != nil {
						continue
					}
					πTemp006 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					πTemp006[0] = µc
					if πTemp007, πE = πg.ResolveGlobal(πF, ßhex_numbers); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetAttr(πF, πTemp007, ßindex, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp008.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πTemp004, πE = πg.Add(πF, πTemp005, πTemp007); πE != nil {
						continue
					}
					µi = πTemp004
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
					// line 365: return i
					πF.SetLineno(365)
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					πR = µi
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßstrhex_to_int.ToObject(), πTemp012); πE != nil {
				continue
			}
			// line 367: hqx_encoding = '!"#$%&\'()*+,-012345689@ABCDEFGHIJKLMNPQRSTUVXYZ[`abcdefhijklmpqr'
			πF.SetLineno(367)
			if πE = πF.Globals().SetItem(πF, ßhqx_encoding.ToObject(), πg.NewStr("!\"#$%&'()*+,-012345689@ABCDEFGHIJKLMNPQRSTUVXYZ[`abcdefhijklmpqr").ToObject()); πE != nil {
				continue
			}
			// line 369: DONE = 0x7f
			πF.SetLineno(369)
			if πE = πF.Globals().SetItem(πF, ßDONE.ToObject(), πg.NewInt(127).ToObject()); πE != nil {
				continue
			}
			// line 370: SKIP = 0x7e
			πF.SetLineno(370)
			if πE = πF.Globals().SetItem(πF, ßSKIP.ToObject(), πg.NewInt(126).ToObject()); πE != nil {
				continue
			}
			// line 371: FAIL = 0x7d
			πF.SetLineno(371)
			if πE = πF.Globals().SetItem(πF, ßFAIL.ToObject(), πg.NewInt(125).ToObject()); πE != nil {
				continue
			}
			// line 373: table_a2b_hqx = [
			πF.SetLineno(373)
			πTemp003 = make([]*πg.Object, 256)
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[0] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[1] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[2] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[3] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[4] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[5] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[6] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[7] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[8] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[9] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßSKIP); πE != nil {
				continue
			}
			πTemp003[10] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[11] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[12] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßSKIP); πE != nil {
				continue
			}
			πTemp003[13] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[14] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[15] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[16] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[17] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[18] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[19] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[20] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[21] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[22] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[23] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[24] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[25] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[26] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[27] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[28] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[29] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[30] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[31] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[32] = πTemp013
			πTemp003[33] = πg.NewInt(0).ToObject()
			πTemp003[34] = πg.NewInt(1).ToObject()
			πTemp003[35] = πg.NewInt(2).ToObject()
			πTemp003[36] = πg.NewInt(3).ToObject()
			πTemp003[37] = πg.NewInt(4).ToObject()
			πTemp003[38] = πg.NewInt(5).ToObject()
			πTemp003[39] = πg.NewInt(6).ToObject()
			πTemp003[40] = πg.NewInt(7).ToObject()
			πTemp003[41] = πg.NewInt(8).ToObject()
			πTemp003[42] = πg.NewInt(9).ToObject()
			πTemp003[43] = πg.NewInt(10).ToObject()
			πTemp003[44] = πg.NewInt(11).ToObject()
			πTemp003[45] = πg.NewInt(12).ToObject()
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[46] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[47] = πTemp013
			πTemp003[48] = πg.NewInt(13).ToObject()
			πTemp003[49] = πg.NewInt(14).ToObject()
			πTemp003[50] = πg.NewInt(15).ToObject()
			πTemp003[51] = πg.NewInt(16).ToObject()
			πTemp003[52] = πg.NewInt(17).ToObject()
			πTemp003[53] = πg.NewInt(18).ToObject()
			πTemp003[54] = πg.NewInt(19).ToObject()
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[55] = πTemp013
			πTemp003[56] = πg.NewInt(20).ToObject()
			πTemp003[57] = πg.NewInt(21).ToObject()
			if πTemp013, πE = πg.ResolveGlobal(πF, ßDONE); πE != nil {
				continue
			}
			πTemp003[58] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[59] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[60] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[61] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[62] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[63] = πTemp013
			πTemp003[64] = πg.NewInt(22).ToObject()
			πTemp003[65] = πg.NewInt(23).ToObject()
			πTemp003[66] = πg.NewInt(24).ToObject()
			πTemp003[67] = πg.NewInt(25).ToObject()
			πTemp003[68] = πg.NewInt(26).ToObject()
			πTemp003[69] = πg.NewInt(27).ToObject()
			πTemp003[70] = πg.NewInt(28).ToObject()
			πTemp003[71] = πg.NewInt(29).ToObject()
			πTemp003[72] = πg.NewInt(30).ToObject()
			πTemp003[73] = πg.NewInt(31).ToObject()
			πTemp003[74] = πg.NewInt(32).ToObject()
			πTemp003[75] = πg.NewInt(33).ToObject()
			πTemp003[76] = πg.NewInt(34).ToObject()
			πTemp003[77] = πg.NewInt(35).ToObject()
			πTemp003[78] = πg.NewInt(36).ToObject()
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[79] = πTemp013
			πTemp003[80] = πg.NewInt(37).ToObject()
			πTemp003[81] = πg.NewInt(38).ToObject()
			πTemp003[82] = πg.NewInt(39).ToObject()
			πTemp003[83] = πg.NewInt(40).ToObject()
			πTemp003[84] = πg.NewInt(41).ToObject()
			πTemp003[85] = πg.NewInt(42).ToObject()
			πTemp003[86] = πg.NewInt(43).ToObject()
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[87] = πTemp013
			πTemp003[88] = πg.NewInt(44).ToObject()
			πTemp003[89] = πg.NewInt(45).ToObject()
			πTemp003[90] = πg.NewInt(46).ToObject()
			πTemp003[91] = πg.NewInt(47).ToObject()
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[92] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[93] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[94] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[95] = πTemp013
			πTemp003[96] = πg.NewInt(48).ToObject()
			πTemp003[97] = πg.NewInt(49).ToObject()
			πTemp003[98] = πg.NewInt(50).ToObject()
			πTemp003[99] = πg.NewInt(51).ToObject()
			πTemp003[100] = πg.NewInt(52).ToObject()
			πTemp003[101] = πg.NewInt(53).ToObject()
			πTemp003[102] = πg.NewInt(54).ToObject()
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[103] = πTemp013
			πTemp003[104] = πg.NewInt(55).ToObject()
			πTemp003[105] = πg.NewInt(56).ToObject()
			πTemp003[106] = πg.NewInt(57).ToObject()
			πTemp003[107] = πg.NewInt(58).ToObject()
			πTemp003[108] = πg.NewInt(59).ToObject()
			πTemp003[109] = πg.NewInt(60).ToObject()
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[110] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[111] = πTemp013
			πTemp003[112] = πg.NewInt(61).ToObject()
			πTemp003[113] = πg.NewInt(62).ToObject()
			πTemp003[114] = πg.NewInt(63).ToObject()
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[115] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[116] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[117] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[118] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[119] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[120] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[121] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[122] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[123] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[124] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[125] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[126] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[127] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[128] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[129] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[130] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[131] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[132] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[133] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[134] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[135] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[136] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[137] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[138] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[139] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[140] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[141] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[142] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[143] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[144] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[145] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[146] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[147] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[148] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[149] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[150] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[151] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[152] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[153] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[154] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[155] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[156] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[157] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[158] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[159] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[160] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[161] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[162] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[163] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[164] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[165] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[166] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[167] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[168] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[169] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[170] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[171] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[172] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[173] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[174] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[175] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[176] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[177] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[178] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[179] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[180] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[181] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[182] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[183] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[184] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[185] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[186] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[187] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[188] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[189] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[190] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[191] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[192] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[193] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[194] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[195] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[196] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[197] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[198] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[199] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[200] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[201] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[202] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[203] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[204] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[205] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[206] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[207] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[208] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[209] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[210] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[211] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[212] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[213] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[214] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[215] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[216] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[217] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[218] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[219] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[220] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[221] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[222] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[223] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[224] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[225] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[226] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[227] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[228] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[229] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[230] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[231] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[232] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[233] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[234] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[235] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[236] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[237] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[238] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[239] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[240] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[241] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[242] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[243] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[244] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[245] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[246] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[247] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[248] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[249] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[250] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[251] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[252] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[253] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[254] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp003[255] = πTemp013
			πTemp013 = πg.NewList(πTemp003...).ToObject()
			if πE = πF.Globals().SetItem(πF, ßtable_a2b_hqx.ToObject(), πTemp013); πE != nil {
				continue
			}
			// line 424: def a2b_hqx(s):
			πF.SetLineno(424)
			πTemp006 = make([]πg.Param, 1)
			πTemp006[0] = πg.Param{Name: "s", Def: nil}
			πTemp013 = πg.NewFunction(πg.NewCode("a2b_hqx", "/usr/lib/python2.7/binascii.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µs *πg.Object = πArgs[0]; _ = µs
				var µresult *πg.Object = πg.UnboundLocal; _ = µresult
				var µquadruples_gen *πg.Object = πg.UnboundLocal; _ = µquadruples_gen
				var µdone *πg.Object = πg.UnboundLocal; _ = µdone
				var µsnippet *πg.Object = πg.UnboundLocal; _ = µsnippet
				var µlength *πg.Object = πg.UnboundLocal; _ = µlength
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
				var πTemp006 bool
				_ = πTemp006
				var πTemp007 bool
				_ = πTemp007
				var πTemp008 *πg.Object
				_ = πTemp008
				var πTemp009 []*πg.Object
				_ = πTemp009
				var πTemp010 *πg.Object
				_ = πTemp010
				var πTemp011 *πg.Object
				_ = πTemp011
				var πTemp012 *πg.Object
				_ = πTemp012
				var πTemp013 *πg.BaseException
				_ = πTemp013
				var πTemp014 *πg.Traceback
				_ = πTemp014
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 2: goto Label2
					case 3: goto Label3
					case 4: goto Label4
					default: panic("unexpected function state")
					}
					// line 425: result = []
					πF.SetLineno(425)
					πTemp001 = make([]*πg.Object, 0)
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					µresult = πTemp002
					// line 427: def quadruples_gen(s):
					πF.SetLineno(427)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "s", Def: nil}
					πTemp002 = πg.NewFunction(πg.NewCode("quadruples_gen", "/usr/lib/python2.7/binascii.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µs *πg.Object = πArgs[0]; _ = µs
						var µt *πg.Object = πg.UnboundLocal; _ = µt
						var µc *πg.Object = πg.UnboundLocal; _ = µc
						var µres *πg.Object = πg.UnboundLocal; _ = µres
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
						return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
							for ; πF.State() >= 0; πF.PopCheckpoint() {
								switch πF.State() {
								case 0:
								case 1: goto Label1
								case 2: goto Label2
								case 12: goto Label12
								case 13: goto Label13
								case 9: goto Label9
								default: panic("unexpected function state")
								}
								// line 428: t = []
								πF.SetLineno(428)
								πTemp001 = make([]*πg.Object, 0)
								πTemp002 = πg.NewList(πTemp001...).ToObject()
								µt = πTemp002
								if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.Iter(πF, µs); πE != nil {
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
								if πTemp005, πE = πg.Next(πF, πTemp002); πE != nil {
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
									µc = πTemp005
								}
								if πE != nil || !πTemp004 {
									continue
								}
								πF.PushCheckpoint(1)            
								// line 430: res = table_a2b_hqx[ord(c)]
								πF.SetLineno(430)
								πTemp001 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
									continue
								}
								πTemp001[0] = µc
								if πTemp006, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
									continue
								}
								if πTemp007, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								πTemp005 = πTemp007
								if πTemp007, πE = πg.ResolveGlobal(πF, ßtable_a2b_hqx); πE != nil {
									continue
								}
								if πTemp006, πE = πg.GetItem(πF, πTemp007, πTemp005); πE != nil {
									continue
								}
								µres = πTemp006
								if πE = πg.CheckLocal(πF, µres, "res"); πE != nil {
									continue
								}
								if πTemp006, πE = πg.ResolveGlobal(πF, ßSKIP); πE != nil {
									continue
								}
								if πTemp005, πE = πg.Eq(πF, µres, πTemp006); πE != nil {
									continue
								}
								if πTemp004, πE = πg.IsTrue(πF, πTemp005); πE != nil {
									continue
								}
								if πTemp004 {
									goto Label4
								}
								if πE = πg.CheckLocal(πF, µres, "res"); πE != nil {
									continue
								}
								if πTemp006, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
									continue
								}
								if πTemp005, πE = πg.Eq(πF, µres, πTemp006); πE != nil {
									continue
								}
								if πTemp004, πE = πg.IsTrue(πF, πTemp005); πE != nil {
									continue
								}
								if πTemp004 {
									goto Label5
								}
								if πE = πg.CheckLocal(πF, µres, "res"); πE != nil {
									continue
								}
								if πTemp006, πE = πg.ResolveGlobal(πF, ßDONE); πE != nil {
									continue
								}
								if πTemp005, πE = πg.Eq(πF, µres, πTemp006); πE != nil {
									continue
								}
								if πTemp004, πE = πg.IsTrue(πF, πTemp005); πE != nil {
									continue
								}
								if πTemp004 {
									goto Label6
								}
								goto Label7
								// line 431: if res == SKIP:
								πF.SetLineno(431)
							Label4:
								// line 432: continue
								πF.SetLineno(432)
								continue
								goto Label8
								// line 433: elif res == FAIL:
								πF.SetLineno(433)
							Label5:
								πTemp001 = πF.MakeArgs(1)
								πTemp001[0] = πg.NewStr("Illegal character").ToObject()
								if πTemp005, πE = πg.ResolveGlobal(πF, ßError); πE != nil {
									continue
								}
								if πTemp006, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								// line 434: raise Error('Illegal character')
								πF.SetLineno(434)
								πE = πF.Raise(πTemp006, nil, nil)
								continue
								goto Label8
								// line 435: elif res == DONE:
								πF.SetLineno(435)
							Label6:
								// line 436: yield t
								πF.SetLineno(436)
								if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
									continue
								}
								πF.PushCheckpoint(9)
								return µt, nil
							Label9:
								πTemp005 = πSent
								if πTemp005, πE = πg.ResolveGlobal(πF, ßDone); πE != nil {
									continue
								}
								// line 437: raise Done
								πF.SetLineno(437)
								πE = πF.Raise(πTemp005, nil, nil)
								continue
								goto Label8
							Label7:
								// line 439: t.append(res)
								πF.SetLineno(439)
								πTemp001 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µres, "res"); πE != nil {
									continue
								}
								πTemp001[0] = µres
								if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
									continue
								}
								if πTemp005, πE = πg.GetAttr(πF, µt, ßappend, nil); πE != nil {
									continue
								}
								if πTemp006, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								goto Label8
							Label8:
								πTemp001 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
									continue
								}
								πTemp001[0] = µt
								if πTemp006, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
									continue
								}
								if πTemp007, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								if πTemp005, πE = πg.Eq(πF, πTemp007, πg.NewInt(4).ToObject()); πE != nil {
									continue
								}
								if πTemp004, πE = πg.IsTrue(πF, πTemp005); πE != nil {
									continue
								}
								if πTemp004 {
									goto Label10
								}
								goto Label11
								// line 440: if len(t) == 4:
								πF.SetLineno(440)
							Label10:
								// line 441: yield t
								πF.SetLineno(441)
								if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
									continue
								}
								πF.PushCheckpoint(12)
								return µt, nil
							Label12:
								πTemp005 = πSent
								// line 442: t = []
								πF.SetLineno(442)
								πTemp001 = make([]*πg.Object, 0)
								πTemp005 = πg.NewList(πTemp001...).ToObject()
								µt = πTemp005
								goto Label11
							Label11:
								continue
							Label2:
								if πE != nil || πR != nil {
									continue
								}
							Label3:
								// line 443: yield t
								πF.SetLineno(443)
								if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
									continue
								}
								πF.PushCheckpoint(13)
								return µt, nil
							Label13:
								πTemp002 = πSent
							}
							return nil, πE
						}).ToObject(), nil
					}), πF.Globals()).ToObject()
					µquadruples_gen = πTemp002
					// line 445: done = 0
					πF.SetLineno(445)
					µdone = πg.NewInt(0).ToObject()
					// line 446: try:
					πF.SetLineno(446)
					πF.PushCheckpoint(2)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp001[0] = µs
					if πE = πg.CheckLocal(πF, µquadruples_gen, "quadruples_gen"); πE != nil {
						continue
					}
					if πTemp005, πE = µquadruples_gen.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp004, πE = πg.Iter(πF, πTemp005); πE != nil {
						continue
					}
					πF.PushCheckpoint(4)
					πTemp006 = false
				Label3:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp006 {
						πF.PopCheckpoint()
						goto Label5
					}
					if πTemp005, πE = πg.Next(πF, πTemp004); πE != nil {
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
						µsnippet = πTemp005
					}
					if πE != nil || !πTemp007 {
						continue
					}
					πF.PushCheckpoint(3)            
					// line 448: length = len(snippet)
					πF.SetLineno(448)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µsnippet, "snippet"); πE != nil {
						continue
					}
					πTemp001[0] = µsnippet
					if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µlength = πTemp008
					if πE = πg.CheckLocal(πF, µlength, "length"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Eq(πF, µlength, πg.NewInt(4).ToObject()); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label6
					}
					if πE = πg.CheckLocal(πF, µlength, "length"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Eq(πF, µlength, πg.NewInt(3).ToObject()); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label7
					}
					if πE = πg.CheckLocal(πF, µlength, "length"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Eq(πF, µlength, πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label8
					}
					goto Label9
					// line 449: if length == 4:
					πF.SetLineno(449)
				Label6:
					// line 450: result.append(chr(((snippet[0] & 0x3f) << 2) | (snippet[1] >> 4)))
					πF.SetLineno(450)
					πTemp001 = πF.MakeArgs(1)
					πTemp009 = πF.MakeArgs(1)
					πTemp011 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µsnippet, "snippet"); πE != nil {
						continue
					}
					if πTemp012, πE = πg.GetItem(πF, µsnippet, πTemp011); πE != nil {
						continue
					}
					if πTemp010, πE = πg.And(πF, πTemp012, πg.NewInt(63).ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.LShift(πF, πTemp010, πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
					πTemp011 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µsnippet, "snippet"); πE != nil {
						continue
					}
					if πTemp012, πE = πg.GetItem(πF, µsnippet, πTemp011); πE != nil {
						continue
					}
					if πTemp010, πE = πg.RShift(πF, πTemp012, πg.NewInt(4).ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Or(πF, πTemp008, πTemp010); πE != nil {
						continue
					}
					πTemp009[0] = πTemp005
					if πTemp005, πE = πg.ResolveGlobal(πF, ßchr); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp005.Call(πF, πTemp009, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp009)
					πTemp001[0] = πTemp008
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µresult, ßappend, nil); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 451: result.append(chr(((snippet[1] & 0x0f) << 4) | (snippet[2] >> 2)))
					πF.SetLineno(451)
					πTemp001 = πF.MakeArgs(1)
					πTemp009 = πF.MakeArgs(1)
					πTemp011 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µsnippet, "snippet"); πE != nil {
						continue
					}
					if πTemp012, πE = πg.GetItem(πF, µsnippet, πTemp011); πE != nil {
						continue
					}
					if πTemp010, πE = πg.And(πF, πTemp012, πg.NewInt(15).ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.LShift(πF, πTemp010, πg.NewInt(4).ToObject()); πE != nil {
						continue
					}
					πTemp011 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µsnippet, "snippet"); πE != nil {
						continue
					}
					if πTemp012, πE = πg.GetItem(πF, µsnippet, πTemp011); πE != nil {
						continue
					}
					if πTemp010, πE = πg.RShift(πF, πTemp012, πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Or(πF, πTemp008, πTemp010); πE != nil {
						continue
					}
					πTemp009[0] = πTemp005
					if πTemp005, πE = πg.ResolveGlobal(πF, ßchr); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp005.Call(πF, πTemp009, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp009)
					πTemp001[0] = πTemp008
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µresult, ßappend, nil); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 452: result.append(chr(((snippet[2] & 0x03) << 6) | (snippet[3])))
					πF.SetLineno(452)
					πTemp001 = πF.MakeArgs(1)
					πTemp009 = πF.MakeArgs(1)
					πTemp011 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µsnippet, "snippet"); πE != nil {
						continue
					}
					if πTemp012, πE = πg.GetItem(πF, µsnippet, πTemp011); πE != nil {
						continue
					}
					if πTemp010, πE = πg.And(πF, πTemp012, πg.NewInt(3).ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.LShift(πF, πTemp010, πg.NewInt(6).ToObject()); πE != nil {
						continue
					}
					πTemp010 = πg.NewInt(3).ToObject()
					if πE = πg.CheckLocal(πF, µsnippet, "snippet"); πE != nil {
						continue
					}
					if πTemp011, πE = πg.GetItem(πF, µsnippet, πTemp010); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Or(πF, πTemp008, πTemp011); πE != nil {
						continue
					}
					πTemp009[0] = πTemp005
					if πTemp005, πE = πg.ResolveGlobal(πF, ßchr); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp005.Call(πF, πTemp009, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp009)
					πTemp001[0] = πTemp008
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µresult, ßappend, nil); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					goto Label9
					// line 453: elif length == 3:
					πF.SetLineno(453)
				Label7:
					// line 454: result.append(chr(((snippet[0] & 0x3f) << 2) | (snippet[1] >> 4)))
					πF.SetLineno(454)
					πTemp001 = πF.MakeArgs(1)
					πTemp009 = πF.MakeArgs(1)
					πTemp011 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µsnippet, "snippet"); πE != nil {
						continue
					}
					if πTemp012, πE = πg.GetItem(πF, µsnippet, πTemp011); πE != nil {
						continue
					}
					if πTemp010, πE = πg.And(πF, πTemp012, πg.NewInt(63).ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.LShift(πF, πTemp010, πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
					πTemp011 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µsnippet, "snippet"); πE != nil {
						continue
					}
					if πTemp012, πE = πg.GetItem(πF, µsnippet, πTemp011); πE != nil {
						continue
					}
					if πTemp010, πE = πg.RShift(πF, πTemp012, πg.NewInt(4).ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Or(πF, πTemp008, πTemp010); πE != nil {
						continue
					}
					πTemp009[0] = πTemp005
					if πTemp005, πE = πg.ResolveGlobal(πF, ßchr); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp005.Call(πF, πTemp009, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp009)
					πTemp001[0] = πTemp008
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µresult, ßappend, nil); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 455: result.append(chr(((snippet[1] & 0x0f) << 4) | (snippet[2] >> 2)))
					πF.SetLineno(455)
					πTemp001 = πF.MakeArgs(1)
					πTemp009 = πF.MakeArgs(1)
					πTemp011 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µsnippet, "snippet"); πE != nil {
						continue
					}
					if πTemp012, πE = πg.GetItem(πF, µsnippet, πTemp011); πE != nil {
						continue
					}
					if πTemp010, πE = πg.And(πF, πTemp012, πg.NewInt(15).ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.LShift(πF, πTemp010, πg.NewInt(4).ToObject()); πE != nil {
						continue
					}
					πTemp011 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µsnippet, "snippet"); πE != nil {
						continue
					}
					if πTemp012, πE = πg.GetItem(πF, µsnippet, πTemp011); πE != nil {
						continue
					}
					if πTemp010, πE = πg.RShift(πF, πTemp012, πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Or(πF, πTemp008, πTemp010); πE != nil {
						continue
					}
					πTemp009[0] = πTemp005
					if πTemp005, πE = πg.ResolveGlobal(πF, ßchr); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp005.Call(πF, πTemp009, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp009)
					πTemp001[0] = πTemp008
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µresult, ßappend, nil); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					goto Label9
					// line 456: elif length == 2:
					πF.SetLineno(456)
				Label8:
					// line 457: result.append(chr(((snippet[0] & 0x3f) << 2) | (snippet[1] >> 4)))
					πF.SetLineno(457)
					πTemp001 = πF.MakeArgs(1)
					πTemp009 = πF.MakeArgs(1)
					πTemp011 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µsnippet, "snippet"); πE != nil {
						continue
					}
					if πTemp012, πE = πg.GetItem(πF, µsnippet, πTemp011); πE != nil {
						continue
					}
					if πTemp010, πE = πg.And(πF, πTemp012, πg.NewInt(63).ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.LShift(πF, πTemp010, πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
					πTemp011 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µsnippet, "snippet"); πE != nil {
						continue
					}
					if πTemp012, πE = πg.GetItem(πF, µsnippet, πTemp011); πE != nil {
						continue
					}
					if πTemp010, πE = πg.RShift(πF, πTemp012, πg.NewInt(4).ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Or(πF, πTemp008, πTemp010); πE != nil {
						continue
					}
					πTemp009[0] = πTemp005
					if πTemp005, πE = πg.ResolveGlobal(πF, ßchr); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp005.Call(πF, πTemp009, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp009)
					πTemp001[0] = πTemp008
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µresult, ßappend, nil); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					goto Label9
				Label9:
					continue
				Label4:
					if πE != nil || πR != nil {
						continue
					}
				Label5:
					πF.PopCheckpoint()
					goto Label1
				Label2:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp013, πTemp014 = πF.ExcInfo()
					if πTemp004, πE = πg.ResolveGlobal(πF, ßDone); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsInstance(πF, πTemp013.ToObject(), πTemp004); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label10
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßError); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsInstance(πF, πTemp013.ToObject(), πTemp004); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label11
					}
					πE = πF.Raise(πTemp013.ToObject(), nil, πTemp014.ToObject())
					continue
					// line 458: except Done:
					πF.SetLineno(458)
				Label10:
					// line 459: done = 1
					πF.SetLineno(459)
					µdone = πg.NewInt(1).ToObject()
					πF.RestoreExc(nil, nil)
					goto Label1
					// line 460: except Error:
					πF.SetLineno(460)
				Label11:
					// line 461: raise
					πF.SetLineno(461)
					πE = πF.Raise(nil, nil, nil)
					continue
					πF.RestoreExc(nil, nil)
					goto Label1
				Label1:
					// line 462: return (''.join(result), done)
					πF.SetLineno(462)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					πTemp001[0] = µresult
					if πTemp005, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.CheckLocal(πF, µdone, "done"); πE != nil {
						continue
					}
					πTemp004 = πg.NewTuple2(πTemp008, µdone).ToObject()
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
			if πE = πF.Globals().SetItem(πF, ßa2b_hqx.ToObject(), πTemp013); πE != nil {
				continue
			}
			// line 464: def b2a_hqx(s):
			πF.SetLineno(464)
			πTemp006 = make([]πg.Param, 1)
			πTemp006[0] = πg.Param{Name: "s", Def: nil}
			πTemp014 = πg.NewFunction(πg.NewCode("b2a_hqx", "/usr/lib/python2.7/binascii.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µs *πg.Object = πArgs[0]; _ = µs
				var µresult *πg.Object = πg.UnboundLocal; _ = µresult
				var µtriples_gen *πg.Object = πg.UnboundLocal; _ = µtriples_gen
				var µsnippet *πg.Object = πg.UnboundLocal; _ = µsnippet
				var µlength *πg.Object = πg.UnboundLocal; _ = µlength
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
				var πTemp013 *πg.Object
				_ = πTemp013
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 1: goto Label1
					case 2: goto Label2
					default: panic("unexpected function state")
					}
					// line 465: result =[]
					πF.SetLineno(465)
					πTemp001 = make([]*πg.Object, 0)
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					µresult = πTemp002
					// line 467: def triples_gen(s):
					πF.SetLineno(467)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "s", Def: nil}
					πTemp002 = πg.NewFunction(πg.NewCode("triples_gen", "/usr/lib/python2.7/binascii.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µs *πg.Object = πArgs[0]; _ = µs
						var πTemp001 bool
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
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 *πg.BaseException
						_ = πTemp009
						var πTemp010 *πg.Traceback
						_ = πTemp010
						var πTemp011 []πg.Param
						_ = πTemp011
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
							for ; πF.State() >= 0; πF.PopCheckpoint() {
								switch πF.State() {
								case 0:
								case 8: goto Label8
								case 1: goto Label1
								case 2: goto Label2
								case 5: goto Label5
								case 6: goto Label6
								default: panic("unexpected function state")
								}
								// line 468: while s:
								πF.SetLineno(468)
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
								if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.IsTrue(πF, µs); πE != nil {
									continue
								}
								if πE != nil || !πTemp002 {
									continue
								}
								πF.PushCheckpoint(1)            
								// line 469: try:
								πF.SetLineno(469)
								πF.PushCheckpoint(5)
								// line 470: yield ord(s[0]), ord(s[1]), ord(s[2])
								πF.SetLineno(470)
								πTemp004 = πF.MakeArgs(1)
								πTemp005 = πg.NewInt(0).ToObject()
								if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
									continue
								}
								if πTemp006, πE = πg.GetItem(πF, µs, πTemp005); πE != nil {
									continue
								}
								πTemp004[0] = πTemp006
								if πTemp005, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
									continue
								}
								if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								πTemp004 = πF.MakeArgs(1)
								πTemp005 = πg.NewInt(1).ToObject()
								if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
									continue
								}
								if πTemp007, πE = πg.GetItem(πF, µs, πTemp005); πE != nil {
									continue
								}
								πTemp004[0] = πTemp007
								if πTemp005, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
									continue
								}
								if πTemp007, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								πTemp004 = πF.MakeArgs(1)
								πTemp005 = πg.NewInt(2).ToObject()
								if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
									continue
								}
								if πTemp008, πE = πg.GetItem(πF, µs, πTemp005); πE != nil {
									continue
								}
								πTemp004[0] = πTemp008
								if πTemp005, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
									continue
								}
								if πTemp008, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								πTemp003 = πg.NewTuple3(πTemp006, πTemp007, πTemp008).ToObject()
								πF.PushCheckpoint(6)
								return πTemp003, nil
							Label6:
								πTemp005 = πSent
								πF.PopCheckpoint()
								goto Label4
							Label5:
								if πE == nil {
								  continue
								}
								πE = nil
								πTemp009, πTemp010 = πF.ExcInfo()
								if πTemp005, πE = πg.ResolveGlobal(πF, ßIndexError); πE != nil {
									continue
								}
								if πTemp002, πE = πg.IsInstance(πF, πTemp009.ToObject(), πTemp005); πE != nil {
									continue
								}
								if πTemp002 {
									goto Label7
								}
								πE = πF.Raise(πTemp009.ToObject(), nil, πTemp010.ToObject())
								continue
								// line 471: except IndexError:
								πF.SetLineno(471)
							Label7:
								// line 472: yield tuple([ord(c) for c in s])
								πF.SetLineno(472)
								πTemp004 = πF.MakeArgs(1)
								πTemp011 = make([]πg.Param, 0)
								πTemp006 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/binascii.py", πTemp011, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
											if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
												continue
											}
											if πTemp001, πE = πg.Iter(πF, µs); πE != nil {
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
											// line 472: yield tuple([ord(c) for c in s])
											πF.SetLineno(472)
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
								if πTemp007, πE = πTemp006.Call(πF, nil, nil); πE != nil {
									continue
								}
								if πTemp005, πE = πg.ListType.Call(πF, πg.Args{πTemp007}, nil); πE != nil {
									continue
								}
								πTemp004[0] = πTemp005
								if πTemp005, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
									continue
								}
								if πTemp007, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								πF.PushCheckpoint(8)
								return πTemp007, nil
							Label8:
								πTemp005 = πSent
								πF.RestoreExc(nil, nil)
								goto Label4
							Label4:
								// line 473: s = s[3:]
								πF.SetLineno(473)
								if πTemp005, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(3).ToObject(), πg.None, πg.None}, nil); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
									continue
								}
								if πTemp008, πE = πg.GetItem(πF, µs, πTemp005); πE != nil {
									continue
								}
								µs = πTemp008
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
					µtriples_gen = πTemp002
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp001[0] = µs
					if πE = πg.CheckLocal(πF, µtriples_gen, "triples_gen"); πE != nil {
						continue
					}
					if πTemp005, πE = µtriples_gen.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp004, πE = πg.Iter(πF, πTemp005); πE != nil {
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
					if πTemp005, πE = πg.Next(πF, πTemp004); πE != nil {
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
						µsnippet = πTemp005
					}
					if πE != nil || !πTemp007 {
						continue
					}
					πF.PushCheckpoint(1)            
					// line 476: length = len(snippet)
					πF.SetLineno(476)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µsnippet, "snippet"); πE != nil {
						continue
					}
					πTemp001[0] = µsnippet
					if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µlength = πTemp008
					if πE = πg.CheckLocal(πF, µlength, "length"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Eq(πF, µlength, πg.NewInt(3).ToObject()); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label4
					}
					if πE = πg.CheckLocal(πF, µlength, "length"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Eq(πF, µlength, πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label5
					}
					if πE = πg.CheckLocal(πF, µlength, "length"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Eq(πF, µlength, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label6
					}
					goto Label7
					// line 477: if length == 3:
					πF.SetLineno(477)
				Label4:
					// line 478: result.append(
					πF.SetLineno(478)
					πTemp001 = πF.MakeArgs(1)
					πTemp010 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µsnippet, "snippet"); πE != nil {
						continue
					}
					if πTemp011, πE = πg.GetItem(πF, µsnippet, πTemp010); πE != nil {
						continue
					}
					if πTemp009, πE = πg.And(πF, πTemp011, πg.NewInt(252).ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.RShift(πF, πTemp009, πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
					πTemp005 = πTemp008
					if πTemp009, πE = πg.ResolveGlobal(πF, ßhqx_encoding); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetItem(πF, πTemp009, πTemp005); πE != nil {
						continue
					}
					πTemp001[0] = πTemp008
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µresult, ßappend, nil); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 480: result.append(hqx_encoding[
					πF.SetLineno(480)
					πTemp001 = πF.MakeArgs(1)
					πTemp011 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µsnippet, "snippet"); πE != nil {
						continue
					}
					if πTemp012, πE = πg.GetItem(πF, µsnippet, πTemp011); πE != nil {
						continue
					}
					if πTemp010, πE = πg.And(πF, πTemp012, πg.NewInt(3).ToObject()); πE != nil {
						continue
					}
					if πTemp009, πE = πg.LShift(πF, πTemp010, πg.NewInt(4).ToObject()); πE != nil {
						continue
					}
					πTemp012 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µsnippet, "snippet"); πE != nil {
						continue
					}
					if πTemp013, πE = πg.GetItem(πF, µsnippet, πTemp012); πE != nil {
						continue
					}
					if πTemp011, πE = πg.And(πF, πTemp013, πg.NewInt(240).ToObject()); πE != nil {
						continue
					}
					if πTemp010, πE = πg.RShift(πF, πTemp011, πg.NewInt(4).ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.Or(πF, πTemp009, πTemp010); πE != nil {
						continue
					}
					πTemp005 = πTemp008
					if πTemp009, πE = πg.ResolveGlobal(πF, ßhqx_encoding); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetItem(πF, πTemp009, πTemp005); πE != nil {
						continue
					}
					πTemp001[0] = πTemp008
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µresult, ßappend, nil); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 482: result.append(hqx_encoding[
					πF.SetLineno(482)
					πTemp001 = πF.MakeArgs(1)
					πTemp011 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µsnippet, "snippet"); πE != nil {
						continue
					}
					if πTemp012, πE = πg.GetItem(πF, µsnippet, πTemp011); πE != nil {
						continue
					}
					if πTemp010, πE = πg.And(πF, πTemp012, πg.NewInt(15).ToObject()); πE != nil {
						continue
					}
					if πTemp009, πE = πg.LShift(πF, πTemp010, πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
					πTemp012 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µsnippet, "snippet"); πE != nil {
						continue
					}
					if πTemp013, πE = πg.GetItem(πF, µsnippet, πTemp012); πE != nil {
						continue
					}
					if πTemp011, πE = πg.And(πF, πTemp013, πg.NewInt(192).ToObject()); πE != nil {
						continue
					}
					if πTemp010, πE = πg.RShift(πF, πTemp011, πg.NewInt(6).ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.Or(πF, πTemp009, πTemp010); πE != nil {
						continue
					}
					πTemp005 = πTemp008
					if πTemp009, πE = πg.ResolveGlobal(πF, ßhqx_encoding); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetItem(πF, πTemp009, πTemp005); πE != nil {
						continue
					}
					πTemp001[0] = πTemp008
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µresult, ßappend, nil); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 484: result.append(hqx_encoding[snippet[2] & 0x3f])
					πF.SetLineno(484)
					πTemp001 = πF.MakeArgs(1)
					πTemp009 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µsnippet, "snippet"); πE != nil {
						continue
					}
					if πTemp010, πE = πg.GetItem(πF, µsnippet, πTemp009); πE != nil {
						continue
					}
					if πTemp008, πE = πg.And(πF, πTemp010, πg.NewInt(63).ToObject()); πE != nil {
						continue
					}
					πTemp005 = πTemp008
					if πTemp009, πE = πg.ResolveGlobal(πF, ßhqx_encoding); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetItem(πF, πTemp009, πTemp005); πE != nil {
						continue
					}
					πTemp001[0] = πTemp008
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µresult, ßappend, nil); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					goto Label7
					// line 485: elif length == 2:
					πF.SetLineno(485)
				Label5:
					// line 486: result.append(
					πF.SetLineno(486)
					πTemp001 = πF.MakeArgs(1)
					πTemp010 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µsnippet, "snippet"); πE != nil {
						continue
					}
					if πTemp011, πE = πg.GetItem(πF, µsnippet, πTemp010); πE != nil {
						continue
					}
					if πTemp009, πE = πg.And(πF, πTemp011, πg.NewInt(252).ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.RShift(πF, πTemp009, πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
					πTemp005 = πTemp008
					if πTemp009, πE = πg.ResolveGlobal(πF, ßhqx_encoding); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetItem(πF, πTemp009, πTemp005); πE != nil {
						continue
					}
					πTemp001[0] = πTemp008
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µresult, ßappend, nil); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 488: result.append(hqx_encoding[
					πF.SetLineno(488)
					πTemp001 = πF.MakeArgs(1)
					πTemp011 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µsnippet, "snippet"); πE != nil {
						continue
					}
					if πTemp012, πE = πg.GetItem(πF, µsnippet, πTemp011); πE != nil {
						continue
					}
					if πTemp010, πE = πg.And(πF, πTemp012, πg.NewInt(3).ToObject()); πE != nil {
						continue
					}
					if πTemp009, πE = πg.LShift(πF, πTemp010, πg.NewInt(4).ToObject()); πE != nil {
						continue
					}
					πTemp012 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µsnippet, "snippet"); πE != nil {
						continue
					}
					if πTemp013, πE = πg.GetItem(πF, µsnippet, πTemp012); πE != nil {
						continue
					}
					if πTemp011, πE = πg.And(πF, πTemp013, πg.NewInt(240).ToObject()); πE != nil {
						continue
					}
					if πTemp010, πE = πg.RShift(πF, πTemp011, πg.NewInt(4).ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.Or(πF, πTemp009, πTemp010); πE != nil {
						continue
					}
					πTemp005 = πTemp008
					if πTemp009, πE = πg.ResolveGlobal(πF, ßhqx_encoding); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetItem(πF, πTemp009, πTemp005); πE != nil {
						continue
					}
					πTemp001[0] = πTemp008
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µresult, ßappend, nil); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 490: result.append(hqx_encoding[
					πF.SetLineno(490)
					πTemp001 = πF.MakeArgs(1)
					πTemp010 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µsnippet, "snippet"); πE != nil {
						continue
					}
					if πTemp011, πE = πg.GetItem(πF, µsnippet, πTemp010); πE != nil {
						continue
					}
					if πTemp009, πE = πg.And(πF, πTemp011, πg.NewInt(15).ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.LShift(πF, πTemp009, πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
					πTemp005 = πTemp008
					if πTemp009, πE = πg.ResolveGlobal(πF, ßhqx_encoding); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetItem(πF, πTemp009, πTemp005); πE != nil {
						continue
					}
					πTemp001[0] = πTemp008
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µresult, ßappend, nil); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					goto Label7
					// line 492: elif length == 1:
					πF.SetLineno(492)
				Label6:
					// line 493: result.append(
					πF.SetLineno(493)
					πTemp001 = πF.MakeArgs(1)
					πTemp010 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µsnippet, "snippet"); πE != nil {
						continue
					}
					if πTemp011, πE = πg.GetItem(πF, µsnippet, πTemp010); πE != nil {
						continue
					}
					if πTemp009, πE = πg.And(πF, πTemp011, πg.NewInt(252).ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.RShift(πF, πTemp009, πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
					πTemp005 = πTemp008
					if πTemp009, πE = πg.ResolveGlobal(πF, ßhqx_encoding); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetItem(πF, πTemp009, πTemp005); πE != nil {
						continue
					}
					πTemp001[0] = πTemp008
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µresult, ßappend, nil); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 495: result.append(hqx_encoding[
					πF.SetLineno(495)
					πTemp001 = πF.MakeArgs(1)
					πTemp010 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µsnippet, "snippet"); πE != nil {
						continue
					}
					if πTemp011, πE = πg.GetItem(πF, µsnippet, πTemp010); πE != nil {
						continue
					}
					if πTemp009, πE = πg.And(πF, πTemp011, πg.NewInt(3).ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.LShift(πF, πTemp009, πg.NewInt(4).ToObject()); πE != nil {
						continue
					}
					πTemp005 = πTemp008
					if πTemp009, πE = πg.ResolveGlobal(πF, ßhqx_encoding); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetItem(πF, πTemp009, πTemp005); πE != nil {
						continue
					}
					πTemp001[0] = πTemp008
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µresult, ßappend, nil); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					goto Label7
				Label7:
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
					// line 497: return ''.join(result)
					πF.SetLineno(497)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					πTemp001[0] = µresult
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
			if πE = πF.Globals().SetItem(πF, ßb2a_hqx.ToObject(), πTemp014); πE != nil {
				continue
			}
			// line 499: crctab_hqx = [
			πF.SetLineno(499)
			πTemp003 = make([]*πg.Object, 256)
			πTemp003[0] = πg.NewInt(0).ToObject()
			πTemp003[1] = πg.NewInt(4129).ToObject()
			πTemp003[2] = πg.NewInt(8258).ToObject()
			πTemp003[3] = πg.NewInt(12387).ToObject()
			πTemp003[4] = πg.NewInt(16516).ToObject()
			πTemp003[5] = πg.NewInt(20645).ToObject()
			πTemp003[6] = πg.NewInt(24774).ToObject()
			πTemp003[7] = πg.NewInt(28903).ToObject()
			πTemp003[8] = πg.NewInt(33032).ToObject()
			πTemp003[9] = πg.NewInt(37161).ToObject()
			πTemp003[10] = πg.NewInt(41290).ToObject()
			πTemp003[11] = πg.NewInt(45419).ToObject()
			πTemp003[12] = πg.NewInt(49548).ToObject()
			πTemp003[13] = πg.NewInt(53677).ToObject()
			πTemp003[14] = πg.NewInt(57806).ToObject()
			πTemp003[15] = πg.NewInt(61935).ToObject()
			πTemp003[16] = πg.NewInt(4657).ToObject()
			πTemp003[17] = πg.NewInt(528).ToObject()
			πTemp003[18] = πg.NewInt(12915).ToObject()
			πTemp003[19] = πg.NewInt(8786).ToObject()
			πTemp003[20] = πg.NewInt(21173).ToObject()
			πTemp003[21] = πg.NewInt(17044).ToObject()
			πTemp003[22] = πg.NewInt(29431).ToObject()
			πTemp003[23] = πg.NewInt(25302).ToObject()
			πTemp003[24] = πg.NewInt(37689).ToObject()
			πTemp003[25] = πg.NewInt(33560).ToObject()
			πTemp003[26] = πg.NewInt(45947).ToObject()
			πTemp003[27] = πg.NewInt(41818).ToObject()
			πTemp003[28] = πg.NewInt(54205).ToObject()
			πTemp003[29] = πg.NewInt(50076).ToObject()
			πTemp003[30] = πg.NewInt(62463).ToObject()
			πTemp003[31] = πg.NewInt(58334).ToObject()
			πTemp003[32] = πg.NewInt(9314).ToObject()
			πTemp003[33] = πg.NewInt(13379).ToObject()
			πTemp003[34] = πg.NewInt(1056).ToObject()
			πTemp003[35] = πg.NewInt(5121).ToObject()
			πTemp003[36] = πg.NewInt(25830).ToObject()
			πTemp003[37] = πg.NewInt(29895).ToObject()
			πTemp003[38] = πg.NewInt(17572).ToObject()
			πTemp003[39] = πg.NewInt(21637).ToObject()
			πTemp003[40] = πg.NewInt(42346).ToObject()
			πTemp003[41] = πg.NewInt(46411).ToObject()
			πTemp003[42] = πg.NewInt(34088).ToObject()
			πTemp003[43] = πg.NewInt(38153).ToObject()
			πTemp003[44] = πg.NewInt(58862).ToObject()
			πTemp003[45] = πg.NewInt(62927).ToObject()
			πTemp003[46] = πg.NewInt(50604).ToObject()
			πTemp003[47] = πg.NewInt(54669).ToObject()
			πTemp003[48] = πg.NewInt(13907).ToObject()
			πTemp003[49] = πg.NewInt(9842).ToObject()
			πTemp003[50] = πg.NewInt(5649).ToObject()
			πTemp003[51] = πg.NewInt(1584).ToObject()
			πTemp003[52] = πg.NewInt(30423).ToObject()
			πTemp003[53] = πg.NewInt(26358).ToObject()
			πTemp003[54] = πg.NewInt(22165).ToObject()
			πTemp003[55] = πg.NewInt(18100).ToObject()
			πTemp003[56] = πg.NewInt(46939).ToObject()
			πTemp003[57] = πg.NewInt(42874).ToObject()
			πTemp003[58] = πg.NewInt(38681).ToObject()
			πTemp003[59] = πg.NewInt(34616).ToObject()
			πTemp003[60] = πg.NewInt(63455).ToObject()
			πTemp003[61] = πg.NewInt(59390).ToObject()
			πTemp003[62] = πg.NewInt(55197).ToObject()
			πTemp003[63] = πg.NewInt(51132).ToObject()
			πTemp003[64] = πg.NewInt(18628).ToObject()
			πTemp003[65] = πg.NewInt(22757).ToObject()
			πTemp003[66] = πg.NewInt(26758).ToObject()
			πTemp003[67] = πg.NewInt(30887).ToObject()
			πTemp003[68] = πg.NewInt(2112).ToObject()
			πTemp003[69] = πg.NewInt(6241).ToObject()
			πTemp003[70] = πg.NewInt(10242).ToObject()
			πTemp003[71] = πg.NewInt(14371).ToObject()
			πTemp003[72] = πg.NewInt(51660).ToObject()
			πTemp003[73] = πg.NewInt(55789).ToObject()
			πTemp003[74] = πg.NewInt(59790).ToObject()
			πTemp003[75] = πg.NewInt(63919).ToObject()
			πTemp003[76] = πg.NewInt(35144).ToObject()
			πTemp003[77] = πg.NewInt(39273).ToObject()
			πTemp003[78] = πg.NewInt(43274).ToObject()
			πTemp003[79] = πg.NewInt(47403).ToObject()
			πTemp003[80] = πg.NewInt(23285).ToObject()
			πTemp003[81] = πg.NewInt(19156).ToObject()
			πTemp003[82] = πg.NewInt(31415).ToObject()
			πTemp003[83] = πg.NewInt(27286).ToObject()
			πTemp003[84] = πg.NewInt(6769).ToObject()
			πTemp003[85] = πg.NewInt(2640).ToObject()
			πTemp003[86] = πg.NewInt(14899).ToObject()
			πTemp003[87] = πg.NewInt(10770).ToObject()
			πTemp003[88] = πg.NewInt(56317).ToObject()
			πTemp003[89] = πg.NewInt(52188).ToObject()
			πTemp003[90] = πg.NewInt(64447).ToObject()
			πTemp003[91] = πg.NewInt(60318).ToObject()
			πTemp003[92] = πg.NewInt(39801).ToObject()
			πTemp003[93] = πg.NewInt(35672).ToObject()
			πTemp003[94] = πg.NewInt(47931).ToObject()
			πTemp003[95] = πg.NewInt(43802).ToObject()
			πTemp003[96] = πg.NewInt(27814).ToObject()
			πTemp003[97] = πg.NewInt(31879).ToObject()
			πTemp003[98] = πg.NewInt(19684).ToObject()
			πTemp003[99] = πg.NewInt(23749).ToObject()
			πTemp003[100] = πg.NewInt(11298).ToObject()
			πTemp003[101] = πg.NewInt(15363).ToObject()
			πTemp003[102] = πg.NewInt(3168).ToObject()
			πTemp003[103] = πg.NewInt(7233).ToObject()
			πTemp003[104] = πg.NewInt(60846).ToObject()
			πTemp003[105] = πg.NewInt(64911).ToObject()
			πTemp003[106] = πg.NewInt(52716).ToObject()
			πTemp003[107] = πg.NewInt(56781).ToObject()
			πTemp003[108] = πg.NewInt(44330).ToObject()
			πTemp003[109] = πg.NewInt(48395).ToObject()
			πTemp003[110] = πg.NewInt(36200).ToObject()
			πTemp003[111] = πg.NewInt(40265).ToObject()
			πTemp003[112] = πg.NewInt(32407).ToObject()
			πTemp003[113] = πg.NewInt(28342).ToObject()
			πTemp003[114] = πg.NewInt(24277).ToObject()
			πTemp003[115] = πg.NewInt(20212).ToObject()
			πTemp003[116] = πg.NewInt(15891).ToObject()
			πTemp003[117] = πg.NewInt(11826).ToObject()
			πTemp003[118] = πg.NewInt(7761).ToObject()
			πTemp003[119] = πg.NewInt(3696).ToObject()
			πTemp003[120] = πg.NewInt(65439).ToObject()
			πTemp003[121] = πg.NewInt(61374).ToObject()
			πTemp003[122] = πg.NewInt(57309).ToObject()
			πTemp003[123] = πg.NewInt(53244).ToObject()
			πTemp003[124] = πg.NewInt(48923).ToObject()
			πTemp003[125] = πg.NewInt(44858).ToObject()
			πTemp003[126] = πg.NewInt(40793).ToObject()
			πTemp003[127] = πg.NewInt(36728).ToObject()
			πTemp003[128] = πg.NewInt(37256).ToObject()
			πTemp003[129] = πg.NewInt(33193).ToObject()
			πTemp003[130] = πg.NewInt(45514).ToObject()
			πTemp003[131] = πg.NewInt(41451).ToObject()
			πTemp003[132] = πg.NewInt(53516).ToObject()
			πTemp003[133] = πg.NewInt(49453).ToObject()
			πTemp003[134] = πg.NewInt(61774).ToObject()
			πTemp003[135] = πg.NewInt(57711).ToObject()
			πTemp003[136] = πg.NewInt(4224).ToObject()
			πTemp003[137] = πg.NewInt(161).ToObject()
			πTemp003[138] = πg.NewInt(12482).ToObject()
			πTemp003[139] = πg.NewInt(8419).ToObject()
			πTemp003[140] = πg.NewInt(20484).ToObject()
			πTemp003[141] = πg.NewInt(16421).ToObject()
			πTemp003[142] = πg.NewInt(28742).ToObject()
			πTemp003[143] = πg.NewInt(24679).ToObject()
			πTemp003[144] = πg.NewInt(33721).ToObject()
			πTemp003[145] = πg.NewInt(37784).ToObject()
			πTemp003[146] = πg.NewInt(41979).ToObject()
			πTemp003[147] = πg.NewInt(46042).ToObject()
			πTemp003[148] = πg.NewInt(49981).ToObject()
			πTemp003[149] = πg.NewInt(54044).ToObject()
			πTemp003[150] = πg.NewInt(58239).ToObject()
			πTemp003[151] = πg.NewInt(62302).ToObject()
			πTemp003[152] = πg.NewInt(689).ToObject()
			πTemp003[153] = πg.NewInt(4752).ToObject()
			πTemp003[154] = πg.NewInt(8947).ToObject()
			πTemp003[155] = πg.NewInt(13010).ToObject()
			πTemp003[156] = πg.NewInt(16949).ToObject()
			πTemp003[157] = πg.NewInt(21012).ToObject()
			πTemp003[158] = πg.NewInt(25207).ToObject()
			πTemp003[159] = πg.NewInt(29270).ToObject()
			πTemp003[160] = πg.NewInt(46570).ToObject()
			πTemp003[161] = πg.NewInt(42443).ToObject()
			πTemp003[162] = πg.NewInt(38312).ToObject()
			πTemp003[163] = πg.NewInt(34185).ToObject()
			πTemp003[164] = πg.NewInt(62830).ToObject()
			πTemp003[165] = πg.NewInt(58703).ToObject()
			πTemp003[166] = πg.NewInt(54572).ToObject()
			πTemp003[167] = πg.NewInt(50445).ToObject()
			πTemp003[168] = πg.NewInt(13538).ToObject()
			πTemp003[169] = πg.NewInt(9411).ToObject()
			πTemp003[170] = πg.NewInt(5280).ToObject()
			πTemp003[171] = πg.NewInt(1153).ToObject()
			πTemp003[172] = πg.NewInt(29798).ToObject()
			πTemp003[173] = πg.NewInt(25671).ToObject()
			πTemp003[174] = πg.NewInt(21540).ToObject()
			πTemp003[175] = πg.NewInt(17413).ToObject()
			πTemp003[176] = πg.NewInt(42971).ToObject()
			πTemp003[177] = πg.NewInt(47098).ToObject()
			πTemp003[178] = πg.NewInt(34713).ToObject()
			πTemp003[179] = πg.NewInt(38840).ToObject()
			πTemp003[180] = πg.NewInt(59231).ToObject()
			πTemp003[181] = πg.NewInt(63358).ToObject()
			πTemp003[182] = πg.NewInt(50973).ToObject()
			πTemp003[183] = πg.NewInt(55100).ToObject()
			πTemp003[184] = πg.NewInt(9939).ToObject()
			πTemp003[185] = πg.NewInt(14066).ToObject()
			πTemp003[186] = πg.NewInt(1681).ToObject()
			πTemp003[187] = πg.NewInt(5808).ToObject()
			πTemp003[188] = πg.NewInt(26199).ToObject()
			πTemp003[189] = πg.NewInt(30326).ToObject()
			πTemp003[190] = πg.NewInt(17941).ToObject()
			πTemp003[191] = πg.NewInt(22068).ToObject()
			πTemp003[192] = πg.NewInt(55628).ToObject()
			πTemp003[193] = πg.NewInt(51565).ToObject()
			πTemp003[194] = πg.NewInt(63758).ToObject()
			πTemp003[195] = πg.NewInt(59695).ToObject()
			πTemp003[196] = πg.NewInt(39368).ToObject()
			πTemp003[197] = πg.NewInt(35305).ToObject()
			πTemp003[198] = πg.NewInt(47498).ToObject()
			πTemp003[199] = πg.NewInt(43435).ToObject()
			πTemp003[200] = πg.NewInt(22596).ToObject()
			πTemp003[201] = πg.NewInt(18533).ToObject()
			πTemp003[202] = πg.NewInt(30726).ToObject()
			πTemp003[203] = πg.NewInt(26663).ToObject()
			πTemp003[204] = πg.NewInt(6336).ToObject()
			πTemp003[205] = πg.NewInt(2273).ToObject()
			πTemp003[206] = πg.NewInt(14466).ToObject()
			πTemp003[207] = πg.NewInt(10403).ToObject()
			πTemp003[208] = πg.NewInt(52093).ToObject()
			πTemp003[209] = πg.NewInt(56156).ToObject()
			πTemp003[210] = πg.NewInt(60223).ToObject()
			πTemp003[211] = πg.NewInt(64286).ToObject()
			πTemp003[212] = πg.NewInt(35833).ToObject()
			πTemp003[213] = πg.NewInt(39896).ToObject()
			πTemp003[214] = πg.NewInt(43963).ToObject()
			πTemp003[215] = πg.NewInt(48026).ToObject()
			πTemp003[216] = πg.NewInt(19061).ToObject()
			πTemp003[217] = πg.NewInt(23124).ToObject()
			πTemp003[218] = πg.NewInt(27191).ToObject()
			πTemp003[219] = πg.NewInt(31254).ToObject()
			πTemp003[220] = πg.NewInt(2801).ToObject()
			πTemp003[221] = πg.NewInt(6864).ToObject()
			πTemp003[222] = πg.NewInt(10931).ToObject()
			πTemp003[223] = πg.NewInt(14994).ToObject()
			πTemp003[224] = πg.NewInt(64814).ToObject()
			πTemp003[225] = πg.NewInt(60687).ToObject()
			πTemp003[226] = πg.NewInt(56684).ToObject()
			πTemp003[227] = πg.NewInt(52557).ToObject()
			πTemp003[228] = πg.NewInt(48554).ToObject()
			πTemp003[229] = πg.NewInt(44427).ToObject()
			πTemp003[230] = πg.NewInt(40424).ToObject()
			πTemp003[231] = πg.NewInt(36297).ToObject()
			πTemp003[232] = πg.NewInt(31782).ToObject()
			πTemp003[233] = πg.NewInt(27655).ToObject()
			πTemp003[234] = πg.NewInt(23652).ToObject()
			πTemp003[235] = πg.NewInt(19525).ToObject()
			πTemp003[236] = πg.NewInt(15522).ToObject()
			πTemp003[237] = πg.NewInt(11395).ToObject()
			πTemp003[238] = πg.NewInt(7392).ToObject()
			πTemp003[239] = πg.NewInt(3265).ToObject()
			πTemp003[240] = πg.NewInt(61215).ToObject()
			πTemp003[241] = πg.NewInt(65342).ToObject()
			πTemp003[242] = πg.NewInt(53085).ToObject()
			πTemp003[243] = πg.NewInt(57212).ToObject()
			πTemp003[244] = πg.NewInt(44955).ToObject()
			πTemp003[245] = πg.NewInt(49082).ToObject()
			πTemp003[246] = πg.NewInt(36825).ToObject()
			πTemp003[247] = πg.NewInt(40952).ToObject()
			πTemp003[248] = πg.NewInt(28183).ToObject()
			πTemp003[249] = πg.NewInt(32310).ToObject()
			πTemp003[250] = πg.NewInt(20053).ToObject()
			πTemp003[251] = πg.NewInt(24180).ToObject()
			πTemp003[252] = πg.NewInt(11923).ToObject()
			πTemp003[253] = πg.NewInt(16050).ToObject()
			πTemp003[254] = πg.NewInt(3793).ToObject()
			πTemp003[255] = πg.NewInt(7920).ToObject()
			πTemp015 = πg.NewList(πTemp003...).ToObject()
			if πE = πF.Globals().SetItem(πF, ßcrctab_hqx.ToObject(), πTemp015); πE != nil {
				continue
			}
			// line 534: def crc_hqx(s, crc):
			πF.SetLineno(534)
			πTemp006 = make([]πg.Param, 2)
			πTemp006[0] = πg.Param{Name: "s", Def: nil}
			πTemp006[1] = πg.Param{Name: "crc", Def: nil}
			πTemp015 = πg.NewFunction(πg.NewCode("crc_hqx", "/usr/lib/python2.7/binascii.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µs *πg.Object = πArgs[0]; _ = µs
				var µcrc *πg.Object = πArgs[1]; _ = µcrc
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
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 *πg.Object
				_ = πTemp008
				var πTemp009 *πg.Object
				_ = πTemp009
				var πTemp010 []*πg.Object
				_ = πTemp010
				var πTemp011 *πg.Object
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
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Iter(πF, µs); πE != nil {
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
					// line 536: crc = ((crc << 8) & 0xff00) ^ crctab_hqx[((crc >> 8) & 0xff) ^ ord(c)]
					πF.SetLineno(536)
					if πE = πg.CheckLocal(πF, µcrc, "crc"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.LShift(πF, µcrc, πg.NewInt(8).ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.And(πF, πTemp006, πg.NewInt(65280).ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcrc, "crc"); πE != nil {
						continue
					}
					if πTemp009, πE = πg.RShift(πF, µcrc, πg.NewInt(8).ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.And(πF, πTemp009, πg.NewInt(255).ToObject()); πE != nil {
						continue
					}
					πTemp010 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					πTemp010[0] = µc
					if πTemp009, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
						continue
					}
					if πTemp011, πE = πTemp009.Call(πF, πTemp010, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp010)
					if πTemp007, πE = πg.Xor(πF, πTemp008, πTemp011); πE != nil {
						continue
					}
					πTemp006 = πTemp007
					if πTemp008, πE = πg.ResolveGlobal(πF, ßcrctab_hqx); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetItem(πF, πTemp008, πTemp006); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Xor(πF, πTemp005, πTemp007); πE != nil {
						continue
					}
					µcrc = πTemp004
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
					// line 538: return crc
					πF.SetLineno(538)
					if πE = πg.CheckLocal(πF, µcrc, "crc"); πE != nil {
						continue
					}
					πR = µcrc
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßcrc_hqx.ToObject(), πTemp015); πE != nil {
				continue
			}
			// line 540: def rlecode_hqx(s):
			πF.SetLineno(540)
			πTemp006 = make([]πg.Param, 1)
			πTemp006[0] = πg.Param{Name: "s", Def: nil}
			πTemp016 = πg.NewFunction(πg.NewCode("rlecode_hqx", "/usr/lib/python2.7/binascii.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µs *πg.Object = πArgs[0]; _ = µs
				var µresult *πg.Object = πg.UnboundLocal; _ = µresult
				var µprev *πg.Object = πg.UnboundLocal; _ = µprev
				var µcount *πg.Object = πg.UnboundLocal; _ = µcount
				var µc *πg.Object = πg.UnboundLocal; _ = µc
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
				var πTemp006 bool
				_ = πTemp006
				var πTemp007 []*πg.Object
				_ = πTemp007
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 6: goto Label6
					case 7: goto Label7
					default: panic("unexpected function state")
					}
					// line 541: """
					πF.SetLineno(541)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, µs); πE != nil {
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
					// line 546: if not s:
					πF.SetLineno(546)
				Label1:
					// line 547: return ''
					πF.SetLineno(547)
					πR = ß.ToObject()
					continue
					goto Label2
				Label2:
					// line 548: result = []
					πF.SetLineno(548)
					πTemp003 = make([]*πg.Object, 0)
					πTemp001 = πg.NewList(πTemp003...).ToObject()
					µresult = πTemp001
					// line 549: prev = s[0]
					πF.SetLineno(549)
					πTemp001 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µs, πTemp001); πE != nil {
						continue
					}
					µprev = πTemp004
					// line 550: count = 1
					πF.SetLineno(550)
					µcount = πg.NewInt(1).ToObject()
					if πTemp005, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp004 = πTemp005
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µs, πTemp004); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, πTemp005, πg.NewStr("!").ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label3
					}
					goto Label4
					// line 555: if s[-1] == '!':
					πF.SetLineno(555)
				Label3:
					// line 556: s = s[1:] + '?'
					πF.SetLineno(556)
					if πTemp004, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µs, πTemp004); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, πTemp005, πg.NewStr("?").ToObject()); πE != nil {
						continue
					}
					µs = πTemp001
					goto Label5
				Label4:
					// line 558: s = s[1:] + '!'
					πF.SetLineno(558)
					if πTemp004, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µs, πTemp004); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, πTemp005, πg.NewStr("!").ToObject()); πE != nil {
						continue
					}
					µs = πTemp001
					goto Label5
				Label5:
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Iter(πF, µs); πE != nil {
						continue
					}
					πF.PushCheckpoint(7)
					πTemp002 = false
				Label6:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp002 {
						πF.PopCheckpoint()
						goto Label8
					}
					if πTemp004, πE = πg.Next(πF, πTemp001); πE != nil {
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
						µc = πTemp004
					}
					if πE != nil || !πTemp006 {
						continue
					}
					πF.PushCheckpoint(6)            
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µprev, "prev"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Eq(πF, µc, µprev); πE != nil {
						continue
					}
					πTemp004 = πTemp005
					if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if !πTemp006 {
						goto Label9
					}
					if πE = πg.CheckLocal(πF, µcount, "count"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.LT(πF, µcount, πg.NewInt(255).ToObject()); πE != nil {
						continue
					}
					πTemp004 = πTemp005
				Label9:
					if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label10
					}
					goto Label11
					// line 561: if c == prev and count < 255:
					πF.SetLineno(561)
				Label10:
					// line 562: count += 1
					πF.SetLineno(562)
					if πE = πg.CheckLocal(πF, µcount, "count"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IAdd(πF, µcount, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µcount = πTemp004
					goto Label12
				Label11:
					if πE = πg.CheckLocal(πF, µcount, "count"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Eq(πF, µcount, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label13
					}
					if πE = πg.CheckLocal(πF, µcount, "count"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.LT(πF, µcount, πg.NewInt(4).ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label14
					}
					if πE = πg.CheckLocal(πF, µprev, "prev"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.NE(πF, µprev, πg.NewStr("\x90").ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label15
					}
					goto Label16
					// line 564: if count == 1:
					πF.SetLineno(564)
				Label13:
					if πE = πg.CheckLocal(πF, µprev, "prev"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.NE(πF, µprev, πg.NewStr("\x90").ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label18
					}
					goto Label19
					// line 565: if prev != '\x90':
					πF.SetLineno(565)
				Label18:
					// line 566: result.append(prev)
					πF.SetLineno(566)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µprev, "prev"); πE != nil {
						continue
					}
					πTemp003[0] = µprev
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µresult, ßappend, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					goto Label20
				Label19:
					// line 568: result += ['\x90', '\x00']
					πF.SetLineno(568)
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					πTemp003 = make([]*πg.Object, 2)
					πTemp003[0] = πg.NewStr("\x90").ToObject()
					πTemp003[1] = πg.NewStr("\x00").ToObject()
					πTemp004 = πg.NewList(πTemp003...).ToObject()
					if πTemp005, πE = πg.IAdd(πF, µresult, πTemp004); πE != nil {
						continue
					}
					µresult = πTemp005
					goto Label20
				Label20:
					goto Label17
					// line 569: elif count < 4:
					πF.SetLineno(569)
				Label14:
					if πE = πg.CheckLocal(πF, µprev, "prev"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.NE(πF, µprev, πg.NewStr("\x90").ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label21
					}
					goto Label22
					// line 570: if prev != '\x90':
					πF.SetLineno(570)
				Label21:
					// line 571: result += [prev] * count
					πF.SetLineno(571)
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					πTemp003 = make([]*πg.Object, 1)
					if πE = πg.CheckLocal(πF, µprev, "prev"); πE != nil {
						continue
					}
					πTemp003[0] = µprev
					πTemp005 = πg.NewList(πTemp003...).ToObject()
					if πE = πg.CheckLocal(πF, µcount, "count"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Mul(πF, πTemp005, µcount); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IAdd(πF, µresult, πTemp004); πE != nil {
						continue
					}
					µresult = πTemp005
					goto Label23
				Label22:
					// line 573: result += ['\x90', '\x00'] * count
					πF.SetLineno(573)
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					πTemp003 = make([]*πg.Object, 2)
					πTemp003[0] = πg.NewStr("\x90").ToObject()
					πTemp003[1] = πg.NewStr("\x00").ToObject()
					πTemp005 = πg.NewList(πTemp003...).ToObject()
					if πE = πg.CheckLocal(πF, µcount, "count"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Mul(πF, πTemp005, µcount); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IAdd(πF, µresult, πTemp004); πE != nil {
						continue
					}
					µresult = πTemp005
					goto Label23
				Label23:
					goto Label17
					// line 575: if prev != '\x90':
					πF.SetLineno(575)
				Label15:
					// line 576: result += [prev, '\x90', chr(count)]
					πF.SetLineno(576)
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					πTemp003 = make([]*πg.Object, 3)
					if πE = πg.CheckLocal(πF, µprev, "prev"); πE != nil {
						continue
					}
					πTemp003[0] = µprev
					πTemp003[1] = πg.NewStr("\x90").ToObject()
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µcount, "count"); πE != nil {
						continue
					}
					πTemp007[0] = µcount
					if πTemp004, πE = πg.ResolveGlobal(πF, ßchr); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					πTemp003[2] = πTemp005
					πTemp004 = πg.NewList(πTemp003...).ToObject()
					if πTemp005, πE = πg.IAdd(πF, µresult, πTemp004); πE != nil {
						continue
					}
					µresult = πTemp005
					goto Label17
				Label16:
					// line 578: result += ['\x90', '\x00', '\x90', chr(count)]
					πF.SetLineno(578)
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					πTemp003 = make([]*πg.Object, 4)
					πTemp003[0] = πg.NewStr("\x90").ToObject()
					πTemp003[1] = πg.NewStr("\x00").ToObject()
					πTemp003[2] = πg.NewStr("\x90").ToObject()
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µcount, "count"); πE != nil {
						continue
					}
					πTemp007[0] = µcount
					if πTemp004, πE = πg.ResolveGlobal(πF, ßchr); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					πTemp003[3] = πTemp005
					πTemp004 = πg.NewList(πTemp003...).ToObject()
					if πTemp005, πE = πg.IAdd(πF, µresult, πTemp004); πE != nil {
						continue
					}
					µresult = πTemp005
					goto Label17
				Label17:
					// line 579: count = 1
					πF.SetLineno(579)
					µcount = πg.NewInt(1).ToObject()
					// line 580: prev = c
					πF.SetLineno(580)
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					µprev = µc
					goto Label12
				Label12:
					continue
				Label7:
					if πE != nil || πR != nil {
						continue
					}
				Label8:
					// line 582: return ''.join(result)
					πF.SetLineno(582)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					πTemp003[0] = µresult
					if πTemp001, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßrlecode_hqx.ToObject(), πTemp016); πE != nil {
				continue
			}
			// line 541: """
			πF.SetLineno(541)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp017}, πg.NewStr("\n    Run length encoding for binhex4.\n    The CPython implementation does not do run length encoding\n    of \x90 characters. This implementation does.\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp018, πE = πg.ResolveGlobal(πF, ßrlecode_hqx); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp018, ß__doc__, πTemp017); πE != nil {
				continue
			}
			// line 584: def rledecode_hqx(s):
			πF.SetLineno(584)
			πTemp006 = make([]πg.Param, 1)
			πTemp006[0] = πg.Param{Name: "s", Def: nil}
			πTemp017 = πg.NewFunction(πg.NewCode("rledecode_hqx", "/usr/lib/python2.7/binascii.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µs *πg.Object = πArgs[0]; _ = µs
				var µresult *πg.Object = πg.UnboundLocal; _ = µresult
				var µprev *πg.Object = πg.UnboundLocal; _ = µprev
				var µsnippet *πg.Object = πg.UnboundLocal; _ = µsnippet
				var µcount *πg.Object = πg.UnboundLocal; _ = µcount
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
					// line 585: s = s.split('\x90')
					πF.SetLineno(585)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("\x90").ToObject()
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
					µs = πTemp003
					// line 586: result = [s[0]]
					πF.SetLineno(586)
					πTemp001 = make([]*πg.Object, 1)
					πTemp002 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µs, πTemp002); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					µresult = πTemp002
					// line 587: prev = s[0]
					πF.SetLineno(587)
					πTemp002 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µs, πTemp002); πE != nil {
						continue
					}
					µprev = πTemp003
					if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µs, πTemp003); πE != nil {
						continue
					}
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
						µsnippet = πTemp003
					}
					if πE != nil || !πTemp006 {
						continue
					}
					πF.PushCheckpoint(1)            
					// line 589: count = ord(snippet[0])
					πF.SetLineno(589)
					πTemp001 = πF.MakeArgs(1)
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µsnippet, "snippet"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µsnippet, πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp004
					if πTemp003, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µcount = πTemp004
					if πE = πg.CheckLocal(πF, µcount, "count"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GT(πF, µcount, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label4
					}
					goto Label5
					// line 590: if count > 0:
					πF.SetLineno(590)
				Label4:
					// line 591: result.append(prev[-1] * (count-1))
					πF.SetLineno(591)
					πTemp001 = πF.MakeArgs(1)
					if πTemp007, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp004 = πTemp007
					if πE = πg.CheckLocal(πF, µprev, "prev"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetItem(πF, µprev, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcount, "count"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Sub(πF, µcount, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Mul(πF, πTemp007, πTemp004); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
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
					// line 592: prev = snippet
					πF.SetLineno(592)
					if πE = πg.CheckLocal(πF, µsnippet, "snippet"); πE != nil {
						continue
					}
					µprev = µsnippet
					goto Label6
				Label5:
					// line 594: result. append('\x90')
					πF.SetLineno(594)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("\x90").ToObject()
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
					// line 595: prev = '\x90'
					πF.SetLineno(595)
					µprev = πg.NewStr("\x90").ToObject()
					goto Label6
				Label6:
					// line 596: result.append(snippet[1:])
					πF.SetLineno(596)
					πTemp001 = πF.MakeArgs(1)
					if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsnippet, "snippet"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µsnippet, πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp004
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
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
					// line 598: return ''.join(result)
					πF.SetLineno(598)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					πTemp001[0] = µresult
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
			if πE = πF.Globals().SetItem(πF, ßrledecode_hqx.ToObject(), πTemp017); πE != nil {
				continue
			}
			// line 600: crc_32_tab = [
			πF.SetLineno(600)
			πTemp003 = make([]*πg.Object, 256)
			πTemp003[0] = πg.NewLongFromBytes([]byte{}).ToObject()
			πTemp003[1] = πg.NewLongFromBytes([]byte{0x77,0x7,0x30,0x96,}).ToObject()
			πTemp003[2] = πg.NewLongFromBytes([]byte{0xee,0xe,0x61,0x2c,}).ToObject()
			πTemp003[3] = πg.NewLongFromBytes([]byte{0x99,0x9,0x51,0xba,}).ToObject()
			πTemp003[4] = πg.NewLongFromBytes([]byte{0x7,0x6d,0xc4,0x19,}).ToObject()
			πTemp003[5] = πg.NewLongFromBytes([]byte{0x70,0x6a,0xf4,0x8f,}).ToObject()
			πTemp003[6] = πg.NewLongFromBytes([]byte{0xe9,0x63,0xa5,0x35,}).ToObject()
			πTemp003[7] = πg.NewLongFromBytes([]byte{0x9e,0x64,0x95,0xa3,}).ToObject()
			πTemp003[8] = πg.NewLongFromBytes([]byte{0xe,0xdb,0x88,0x32,}).ToObject()
			πTemp003[9] = πg.NewLongFromBytes([]byte{0x79,0xdc,0xb8,0xa4,}).ToObject()
			πTemp003[10] = πg.NewLongFromBytes([]byte{0xe0,0xd5,0xe9,0x1e,}).ToObject()
			πTemp003[11] = πg.NewLongFromBytes([]byte{0x97,0xd2,0xd9,0x88,}).ToObject()
			πTemp003[12] = πg.NewLongFromBytes([]byte{0x9,0xb6,0x4c,0x2b,}).ToObject()
			πTemp003[13] = πg.NewLongFromBytes([]byte{0x7e,0xb1,0x7c,0xbd,}).ToObject()
			πTemp003[14] = πg.NewLongFromBytes([]byte{0xe7,0xb8,0x2d,0x7,}).ToObject()
			πTemp003[15] = πg.NewLongFromBytes([]byte{0x90,0xbf,0x1d,0x91,}).ToObject()
			πTemp003[16] = πg.NewLongFromBytes([]byte{0x1d,0xb7,0x10,0x64,}).ToObject()
			πTemp003[17] = πg.NewLongFromBytes([]byte{0x6a,0xb0,0x20,0xf2,}).ToObject()
			πTemp003[18] = πg.NewLongFromBytes([]byte{0xf3,0xb9,0x71,0x48,}).ToObject()
			πTemp003[19] = πg.NewLongFromBytes([]byte{0x84,0xbe,0x41,0xde,}).ToObject()
			πTemp003[20] = πg.NewLongFromBytes([]byte{0x1a,0xda,0xd4,0x7d,}).ToObject()
			πTemp003[21] = πg.NewLongFromBytes([]byte{0x6d,0xdd,0xe4,0xeb,}).ToObject()
			πTemp003[22] = πg.NewLongFromBytes([]byte{0xf4,0xd4,0xb5,0x51,}).ToObject()
			πTemp003[23] = πg.NewLongFromBytes([]byte{0x83,0xd3,0x85,0xc7,}).ToObject()
			πTemp003[24] = πg.NewLongFromBytes([]byte{0x13,0x6c,0x98,0x56,}).ToObject()
			πTemp003[25] = πg.NewLongFromBytes([]byte{0x64,0x6b,0xa8,0xc0,}).ToObject()
			πTemp003[26] = πg.NewLongFromBytes([]byte{0xfd,0x62,0xf9,0x7a,}).ToObject()
			πTemp003[27] = πg.NewLongFromBytes([]byte{0x8a,0x65,0xc9,0xec,}).ToObject()
			πTemp003[28] = πg.NewLongFromBytes([]byte{0x14,0x1,0x5c,0x4f,}).ToObject()
			πTemp003[29] = πg.NewLongFromBytes([]byte{0x63,0x6,0x6c,0xd9,}).ToObject()
			πTemp003[30] = πg.NewLongFromBytes([]byte{0xfa,0xf,0x3d,0x63,}).ToObject()
			πTemp003[31] = πg.NewLongFromBytes([]byte{0x8d,0x8,0xd,0xf5,}).ToObject()
			πTemp003[32] = πg.NewLongFromBytes([]byte{0x3b,0x6e,0x20,0xc8,}).ToObject()
			πTemp003[33] = πg.NewLongFromBytes([]byte{0x4c,0x69,0x10,0x5e,}).ToObject()
			πTemp003[34] = πg.NewLongFromBytes([]byte{0xd5,0x60,0x41,0xe4,}).ToObject()
			πTemp003[35] = πg.NewLongFromBytes([]byte{0xa2,0x67,0x71,0x72,}).ToObject()
			πTemp003[36] = πg.NewLongFromBytes([]byte{0x3c,0x3,0xe4,0xd1,}).ToObject()
			πTemp003[37] = πg.NewLongFromBytes([]byte{0x4b,0x4,0xd4,0x47,}).ToObject()
			πTemp003[38] = πg.NewLongFromBytes([]byte{0xd2,0xd,0x85,0xfd,}).ToObject()
			πTemp003[39] = πg.NewLongFromBytes([]byte{0xa5,0xa,0xb5,0x6b,}).ToObject()
			πTemp003[40] = πg.NewLongFromBytes([]byte{0x35,0xb5,0xa8,0xfa,}).ToObject()
			πTemp003[41] = πg.NewLongFromBytes([]byte{0x42,0xb2,0x98,0x6c,}).ToObject()
			πTemp003[42] = πg.NewLongFromBytes([]byte{0xdb,0xbb,0xc9,0xd6,}).ToObject()
			πTemp003[43] = πg.NewLongFromBytes([]byte{0xac,0xbc,0xf9,0x40,}).ToObject()
			πTemp003[44] = πg.NewLongFromBytes([]byte{0x32,0xd8,0x6c,0xe3,}).ToObject()
			πTemp003[45] = πg.NewLongFromBytes([]byte{0x45,0xdf,0x5c,0x75,}).ToObject()
			πTemp003[46] = πg.NewLongFromBytes([]byte{0xdc,0xd6,0xd,0xcf,}).ToObject()
			πTemp003[47] = πg.NewLongFromBytes([]byte{0xab,0xd1,0x3d,0x59,}).ToObject()
			πTemp003[48] = πg.NewLongFromBytes([]byte{0x26,0xd9,0x30,0xac,}).ToObject()
			πTemp003[49] = πg.NewLongFromBytes([]byte{0x51,0xde,0x0,0x3a,}).ToObject()
			πTemp003[50] = πg.NewLongFromBytes([]byte{0xc8,0xd7,0x51,0x80,}).ToObject()
			πTemp003[51] = πg.NewLongFromBytes([]byte{0xbf,0xd0,0x61,0x16,}).ToObject()
			πTemp003[52] = πg.NewLongFromBytes([]byte{0x21,0xb4,0xf4,0xb5,}).ToObject()
			πTemp003[53] = πg.NewLongFromBytes([]byte{0x56,0xb3,0xc4,0x23,}).ToObject()
			πTemp003[54] = πg.NewLongFromBytes([]byte{0xcf,0xba,0x95,0x99,}).ToObject()
			πTemp003[55] = πg.NewLongFromBytes([]byte{0xb8,0xbd,0xa5,0xf,}).ToObject()
			πTemp003[56] = πg.NewLongFromBytes([]byte{0x28,0x2,0xb8,0x9e,}).ToObject()
			πTemp003[57] = πg.NewLongFromBytes([]byte{0x5f,0x5,0x88,0x8,}).ToObject()
			πTemp003[58] = πg.NewLongFromBytes([]byte{0xc6,0xc,0xd9,0xb2,}).ToObject()
			πTemp003[59] = πg.NewLongFromBytes([]byte{0xb1,0xb,0xe9,0x24,}).ToObject()
			πTemp003[60] = πg.NewLongFromBytes([]byte{0x2f,0x6f,0x7c,0x87,}).ToObject()
			πTemp003[61] = πg.NewLongFromBytes([]byte{0x58,0x68,0x4c,0x11,}).ToObject()
			πTemp003[62] = πg.NewLongFromBytes([]byte{0xc1,0x61,0x1d,0xab,}).ToObject()
			πTemp003[63] = πg.NewLongFromBytes([]byte{0xb6,0x66,0x2d,0x3d,}).ToObject()
			πTemp003[64] = πg.NewLongFromBytes([]byte{0x76,0xdc,0x41,0x90,}).ToObject()
			πTemp003[65] = πg.NewLongFromBytes([]byte{0x1,0xdb,0x71,0x6,}).ToObject()
			πTemp003[66] = πg.NewLongFromBytes([]byte{0x98,0xd2,0x20,0xbc,}).ToObject()
			πTemp003[67] = πg.NewLongFromBytes([]byte{0xef,0xd5,0x10,0x2a,}).ToObject()
			πTemp003[68] = πg.NewLongFromBytes([]byte{0x71,0xb1,0x85,0x89,}).ToObject()
			πTemp003[69] = πg.NewLongFromBytes([]byte{0x6,0xb6,0xb5,0x1f,}).ToObject()
			πTemp003[70] = πg.NewLongFromBytes([]byte{0x9f,0xbf,0xe4,0xa5,}).ToObject()
			πTemp003[71] = πg.NewLongFromBytes([]byte{0xe8,0xb8,0xd4,0x33,}).ToObject()
			πTemp003[72] = πg.NewLongFromBytes([]byte{0x78,0x7,0xc9,0xa2,}).ToObject()
			πTemp003[73] = πg.NewLongFromBytes([]byte{0xf,0x0,0xf9,0x34,}).ToObject()
			πTemp003[74] = πg.NewLongFromBytes([]byte{0x96,0x9,0xa8,0x8e,}).ToObject()
			πTemp003[75] = πg.NewLongFromBytes([]byte{0xe1,0xe,0x98,0x18,}).ToObject()
			πTemp003[76] = πg.NewLongFromBytes([]byte{0x7f,0x6a,0xd,0xbb,}).ToObject()
			πTemp003[77] = πg.NewLongFromBytes([]byte{0x8,0x6d,0x3d,0x2d,}).ToObject()
			πTemp003[78] = πg.NewLongFromBytes([]byte{0x91,0x64,0x6c,0x97,}).ToObject()
			πTemp003[79] = πg.NewLongFromBytes([]byte{0xe6,0x63,0x5c,0x1,}).ToObject()
			πTemp003[80] = πg.NewLongFromBytes([]byte{0x6b,0x6b,0x51,0xf4,}).ToObject()
			πTemp003[81] = πg.NewLongFromBytes([]byte{0x1c,0x6c,0x61,0x62,}).ToObject()
			πTemp003[82] = πg.NewLongFromBytes([]byte{0x85,0x65,0x30,0xd8,}).ToObject()
			πTemp003[83] = πg.NewLongFromBytes([]byte{0xf2,0x62,0x0,0x4e,}).ToObject()
			πTemp003[84] = πg.NewLongFromBytes([]byte{0x6c,0x6,0x95,0xed,}).ToObject()
			πTemp003[85] = πg.NewLongFromBytes([]byte{0x1b,0x1,0xa5,0x7b,}).ToObject()
			πTemp003[86] = πg.NewLongFromBytes([]byte{0x82,0x8,0xf4,0xc1,}).ToObject()
			πTemp003[87] = πg.NewLongFromBytes([]byte{0xf5,0xf,0xc4,0x57,}).ToObject()
			πTemp003[88] = πg.NewLongFromBytes([]byte{0x65,0xb0,0xd9,0xc6,}).ToObject()
			πTemp003[89] = πg.NewLongFromBytes([]byte{0x12,0xb7,0xe9,0x50,}).ToObject()
			πTemp003[90] = πg.NewLongFromBytes([]byte{0x8b,0xbe,0xb8,0xea,}).ToObject()
			πTemp003[91] = πg.NewLongFromBytes([]byte{0xfc,0xb9,0x88,0x7c,}).ToObject()
			πTemp003[92] = πg.NewLongFromBytes([]byte{0x62,0xdd,0x1d,0xdf,}).ToObject()
			πTemp003[93] = πg.NewLongFromBytes([]byte{0x15,0xda,0x2d,0x49,}).ToObject()
			πTemp003[94] = πg.NewLongFromBytes([]byte{0x8c,0xd3,0x7c,0xf3,}).ToObject()
			πTemp003[95] = πg.NewLongFromBytes([]byte{0xfb,0xd4,0x4c,0x65,}).ToObject()
			πTemp003[96] = πg.NewLongFromBytes([]byte{0x4d,0xb2,0x61,0x58,}).ToObject()
			πTemp003[97] = πg.NewLongFromBytes([]byte{0x3a,0xb5,0x51,0xce,}).ToObject()
			πTemp003[98] = πg.NewLongFromBytes([]byte{0xa3,0xbc,0x0,0x74,}).ToObject()
			πTemp003[99] = πg.NewLongFromBytes([]byte{0xd4,0xbb,0x30,0xe2,}).ToObject()
			πTemp003[100] = πg.NewLongFromBytes([]byte{0x4a,0xdf,0xa5,0x41,}).ToObject()
			πTemp003[101] = πg.NewLongFromBytes([]byte{0x3d,0xd8,0x95,0xd7,}).ToObject()
			πTemp003[102] = πg.NewLongFromBytes([]byte{0xa4,0xd1,0xc4,0x6d,}).ToObject()
			πTemp003[103] = πg.NewLongFromBytes([]byte{0xd3,0xd6,0xf4,0xfb,}).ToObject()
			πTemp003[104] = πg.NewLongFromBytes([]byte{0x43,0x69,0xe9,0x6a,}).ToObject()
			πTemp003[105] = πg.NewLongFromBytes([]byte{0x34,0x6e,0xd9,0xfc,}).ToObject()
			πTemp003[106] = πg.NewLongFromBytes([]byte{0xad,0x67,0x88,0x46,}).ToObject()
			πTemp003[107] = πg.NewLongFromBytes([]byte{0xda,0x60,0xb8,0xd0,}).ToObject()
			πTemp003[108] = πg.NewLongFromBytes([]byte{0x44,0x4,0x2d,0x73,}).ToObject()
			πTemp003[109] = πg.NewLongFromBytes([]byte{0x33,0x3,0x1d,0xe5,}).ToObject()
			πTemp003[110] = πg.NewLongFromBytes([]byte{0xaa,0xa,0x4c,0x5f,}).ToObject()
			πTemp003[111] = πg.NewLongFromBytes([]byte{0xdd,0xd,0x7c,0xc9,}).ToObject()
			πTemp003[112] = πg.NewLongFromBytes([]byte{0x50,0x5,0x71,0x3c,}).ToObject()
			πTemp003[113] = πg.NewLongFromBytes([]byte{0x27,0x2,0x41,0xaa,}).ToObject()
			πTemp003[114] = πg.NewLongFromBytes([]byte{0xbe,0xb,0x10,0x10,}).ToObject()
			πTemp003[115] = πg.NewLongFromBytes([]byte{0xc9,0xc,0x20,0x86,}).ToObject()
			πTemp003[116] = πg.NewLongFromBytes([]byte{0x57,0x68,0xb5,0x25,}).ToObject()
			πTemp003[117] = πg.NewLongFromBytes([]byte{0x20,0x6f,0x85,0xb3,}).ToObject()
			πTemp003[118] = πg.NewLongFromBytes([]byte{0xb9,0x66,0xd4,0x9,}).ToObject()
			πTemp003[119] = πg.NewLongFromBytes([]byte{0xce,0x61,0xe4,0x9f,}).ToObject()
			πTemp003[120] = πg.NewLongFromBytes([]byte{0x5e,0xde,0xf9,0xe,}).ToObject()
			πTemp003[121] = πg.NewLongFromBytes([]byte{0x29,0xd9,0xc9,0x98,}).ToObject()
			πTemp003[122] = πg.NewLongFromBytes([]byte{0xb0,0xd0,0x98,0x22,}).ToObject()
			πTemp003[123] = πg.NewLongFromBytes([]byte{0xc7,0xd7,0xa8,0xb4,}).ToObject()
			πTemp003[124] = πg.NewLongFromBytes([]byte{0x59,0xb3,0x3d,0x17,}).ToObject()
			πTemp003[125] = πg.NewLongFromBytes([]byte{0x2e,0xb4,0xd,0x81,}).ToObject()
			πTemp003[126] = πg.NewLongFromBytes([]byte{0xb7,0xbd,0x5c,0x3b,}).ToObject()
			πTemp003[127] = πg.NewLongFromBytes([]byte{0xc0,0xba,0x6c,0xad,}).ToObject()
			πTemp003[128] = πg.NewLongFromBytes([]byte{0xed,0xb8,0x83,0x20,}).ToObject()
			πTemp003[129] = πg.NewLongFromBytes([]byte{0x9a,0xbf,0xb3,0xb6,}).ToObject()
			πTemp003[130] = πg.NewLongFromBytes([]byte{0x3,0xb6,0xe2,0xc,}).ToObject()
			πTemp003[131] = πg.NewLongFromBytes([]byte{0x74,0xb1,0xd2,0x9a,}).ToObject()
			πTemp003[132] = πg.NewLongFromBytes([]byte{0xea,0xd5,0x47,0x39,}).ToObject()
			πTemp003[133] = πg.NewLongFromBytes([]byte{0x9d,0xd2,0x77,0xaf,}).ToObject()
			πTemp003[134] = πg.NewLongFromBytes([]byte{0x4,0xdb,0x26,0x15,}).ToObject()
			πTemp003[135] = πg.NewLongFromBytes([]byte{0x73,0xdc,0x16,0x83,}).ToObject()
			πTemp003[136] = πg.NewLongFromBytes([]byte{0xe3,0x63,0xb,0x12,}).ToObject()
			πTemp003[137] = πg.NewLongFromBytes([]byte{0x94,0x64,0x3b,0x84,}).ToObject()
			πTemp003[138] = πg.NewLongFromBytes([]byte{0xd,0x6d,0x6a,0x3e,}).ToObject()
			πTemp003[139] = πg.NewLongFromBytes([]byte{0x7a,0x6a,0x5a,0xa8,}).ToObject()
			πTemp003[140] = πg.NewLongFromBytes([]byte{0xe4,0xe,0xcf,0xb,}).ToObject()
			πTemp003[141] = πg.NewLongFromBytes([]byte{0x93,0x9,0xff,0x9d,}).ToObject()
			πTemp003[142] = πg.NewLongFromBytes([]byte{0xa,0x0,0xae,0x27,}).ToObject()
			πTemp003[143] = πg.NewLongFromBytes([]byte{0x7d,0x7,0x9e,0xb1,}).ToObject()
			πTemp003[144] = πg.NewLongFromBytes([]byte{0xf0,0xf,0x93,0x44,}).ToObject()
			πTemp003[145] = πg.NewLongFromBytes([]byte{0x87,0x8,0xa3,0xd2,}).ToObject()
			πTemp003[146] = πg.NewLongFromBytes([]byte{0x1e,0x1,0xf2,0x68,}).ToObject()
			πTemp003[147] = πg.NewLongFromBytes([]byte{0x69,0x6,0xc2,0xfe,}).ToObject()
			πTemp003[148] = πg.NewLongFromBytes([]byte{0xf7,0x62,0x57,0x5d,}).ToObject()
			πTemp003[149] = πg.NewLongFromBytes([]byte{0x80,0x65,0x67,0xcb,}).ToObject()
			πTemp003[150] = πg.NewLongFromBytes([]byte{0x19,0x6c,0x36,0x71,}).ToObject()
			πTemp003[151] = πg.NewLongFromBytes([]byte{0x6e,0x6b,0x6,0xe7,}).ToObject()
			πTemp003[152] = πg.NewLongFromBytes([]byte{0xfe,0xd4,0x1b,0x76,}).ToObject()
			πTemp003[153] = πg.NewLongFromBytes([]byte{0x89,0xd3,0x2b,0xe0,}).ToObject()
			πTemp003[154] = πg.NewLongFromBytes([]byte{0x10,0xda,0x7a,0x5a,}).ToObject()
			πTemp003[155] = πg.NewLongFromBytes([]byte{0x67,0xdd,0x4a,0xcc,}).ToObject()
			πTemp003[156] = πg.NewLongFromBytes([]byte{0xf9,0xb9,0xdf,0x6f,}).ToObject()
			πTemp003[157] = πg.NewLongFromBytes([]byte{0x8e,0xbe,0xef,0xf9,}).ToObject()
			πTemp003[158] = πg.NewLongFromBytes([]byte{0x17,0xb7,0xbe,0x43,}).ToObject()
			πTemp003[159] = πg.NewLongFromBytes([]byte{0x60,0xb0,0x8e,0xd5,}).ToObject()
			πTemp003[160] = πg.NewLongFromBytes([]byte{0xd6,0xd6,0xa3,0xe8,}).ToObject()
			πTemp003[161] = πg.NewLongFromBytes([]byte{0xa1,0xd1,0x93,0x7e,}).ToObject()
			πTemp003[162] = πg.NewLongFromBytes([]byte{0x38,0xd8,0xc2,0xc4,}).ToObject()
			πTemp003[163] = πg.NewLongFromBytes([]byte{0x4f,0xdf,0xf2,0x52,}).ToObject()
			πTemp003[164] = πg.NewLongFromBytes([]byte{0xd1,0xbb,0x67,0xf1,}).ToObject()
			πTemp003[165] = πg.NewLongFromBytes([]byte{0xa6,0xbc,0x57,0x67,}).ToObject()
			πTemp003[166] = πg.NewLongFromBytes([]byte{0x3f,0xb5,0x6,0xdd,}).ToObject()
			πTemp003[167] = πg.NewLongFromBytes([]byte{0x48,0xb2,0x36,0x4b,}).ToObject()
			πTemp003[168] = πg.NewLongFromBytes([]byte{0xd8,0xd,0x2b,0xda,}).ToObject()
			πTemp003[169] = πg.NewLongFromBytes([]byte{0xaf,0xa,0x1b,0x4c,}).ToObject()
			πTemp003[170] = πg.NewLongFromBytes([]byte{0x36,0x3,0x4a,0xf6,}).ToObject()
			πTemp003[171] = πg.NewLongFromBytes([]byte{0x41,0x4,0x7a,0x60,}).ToObject()
			πTemp003[172] = πg.NewLongFromBytes([]byte{0xdf,0x60,0xef,0xc3,}).ToObject()
			πTemp003[173] = πg.NewLongFromBytes([]byte{0xa8,0x67,0xdf,0x55,}).ToObject()
			πTemp003[174] = πg.NewLongFromBytes([]byte{0x31,0x6e,0x8e,0xef,}).ToObject()
			πTemp003[175] = πg.NewLongFromBytes([]byte{0x46,0x69,0xbe,0x79,}).ToObject()
			πTemp003[176] = πg.NewLongFromBytes([]byte{0xcb,0x61,0xb3,0x8c,}).ToObject()
			πTemp003[177] = πg.NewLongFromBytes([]byte{0xbc,0x66,0x83,0x1a,}).ToObject()
			πTemp003[178] = πg.NewLongFromBytes([]byte{0x25,0x6f,0xd2,0xa0,}).ToObject()
			πTemp003[179] = πg.NewLongFromBytes([]byte{0x52,0x68,0xe2,0x36,}).ToObject()
			πTemp003[180] = πg.NewLongFromBytes([]byte{0xcc,0xc,0x77,0x95,}).ToObject()
			πTemp003[181] = πg.NewLongFromBytes([]byte{0xbb,0xb,0x47,0x3,}).ToObject()
			πTemp003[182] = πg.NewLongFromBytes([]byte{0x22,0x2,0x16,0xb9,}).ToObject()
			πTemp003[183] = πg.NewLongFromBytes([]byte{0x55,0x5,0x26,0x2f,}).ToObject()
			πTemp003[184] = πg.NewLongFromBytes([]byte{0xc5,0xba,0x3b,0xbe,}).ToObject()
			πTemp003[185] = πg.NewLongFromBytes([]byte{0xb2,0xbd,0xb,0x28,}).ToObject()
			πTemp003[186] = πg.NewLongFromBytes([]byte{0x2b,0xb4,0x5a,0x92,}).ToObject()
			πTemp003[187] = πg.NewLongFromBytes([]byte{0x5c,0xb3,0x6a,0x4,}).ToObject()
			πTemp003[188] = πg.NewLongFromBytes([]byte{0xc2,0xd7,0xff,0xa7,}).ToObject()
			πTemp003[189] = πg.NewLongFromBytes([]byte{0xb5,0xd0,0xcf,0x31,}).ToObject()
			πTemp003[190] = πg.NewLongFromBytes([]byte{0x2c,0xd9,0x9e,0x8b,}).ToObject()
			πTemp003[191] = πg.NewLongFromBytes([]byte{0x5b,0xde,0xae,0x1d,}).ToObject()
			πTemp003[192] = πg.NewLongFromBytes([]byte{0x9b,0x64,0xc2,0xb0,}).ToObject()
			πTemp003[193] = πg.NewLongFromBytes([]byte{0xec,0x63,0xf2,0x26,}).ToObject()
			πTemp003[194] = πg.NewLongFromBytes([]byte{0x75,0x6a,0xa3,0x9c,}).ToObject()
			πTemp003[195] = πg.NewLongFromBytes([]byte{0x2,0x6d,0x93,0xa,}).ToObject()
			πTemp003[196] = πg.NewLongFromBytes([]byte{0x9c,0x9,0x6,0xa9,}).ToObject()
			πTemp003[197] = πg.NewLongFromBytes([]byte{0xeb,0xe,0x36,0x3f,}).ToObject()
			πTemp003[198] = πg.NewLongFromBytes([]byte{0x72,0x7,0x67,0x85,}).ToObject()
			πTemp003[199] = πg.NewLongFromBytes([]byte{0x5,0x0,0x57,0x13,}).ToObject()
			πTemp003[200] = πg.NewLongFromBytes([]byte{0x95,0xbf,0x4a,0x82,}).ToObject()
			πTemp003[201] = πg.NewLongFromBytes([]byte{0xe2,0xb8,0x7a,0x14,}).ToObject()
			πTemp003[202] = πg.NewLongFromBytes([]byte{0x7b,0xb1,0x2b,0xae,}).ToObject()
			πTemp003[203] = πg.NewLongFromBytes([]byte{0xc,0xb6,0x1b,0x38,}).ToObject()
			πTemp003[204] = πg.NewLongFromBytes([]byte{0x92,0xd2,0x8e,0x9b,}).ToObject()
			πTemp003[205] = πg.NewLongFromBytes([]byte{0xe5,0xd5,0xbe,0xd,}).ToObject()
			πTemp003[206] = πg.NewLongFromBytes([]byte{0x7c,0xdc,0xef,0xb7,}).ToObject()
			πTemp003[207] = πg.NewLongFromBytes([]byte{0xb,0xdb,0xdf,0x21,}).ToObject()
			πTemp003[208] = πg.NewLongFromBytes([]byte{0x86,0xd3,0xd2,0xd4,}).ToObject()
			πTemp003[209] = πg.NewLongFromBytes([]byte{0xf1,0xd4,0xe2,0x42,}).ToObject()
			πTemp003[210] = πg.NewLongFromBytes([]byte{0x68,0xdd,0xb3,0xf8,}).ToObject()
			πTemp003[211] = πg.NewLongFromBytes([]byte{0x1f,0xda,0x83,0x6e,}).ToObject()
			πTemp003[212] = πg.NewLongFromBytes([]byte{0x81,0xbe,0x16,0xcd,}).ToObject()
			πTemp003[213] = πg.NewLongFromBytes([]byte{0xf6,0xb9,0x26,0x5b,}).ToObject()
			πTemp003[214] = πg.NewLongFromBytes([]byte{0x6f,0xb0,0x77,0xe1,}).ToObject()
			πTemp003[215] = πg.NewLongFromBytes([]byte{0x18,0xb7,0x47,0x77,}).ToObject()
			πTemp003[216] = πg.NewLongFromBytes([]byte{0x88,0x8,0x5a,0xe6,}).ToObject()
			πTemp003[217] = πg.NewLongFromBytes([]byte{0xff,0xf,0x6a,0x70,}).ToObject()
			πTemp003[218] = πg.NewLongFromBytes([]byte{0x66,0x6,0x3b,0xca,}).ToObject()
			πTemp003[219] = πg.NewLongFromBytes([]byte{0x11,0x1,0xb,0x5c,}).ToObject()
			πTemp003[220] = πg.NewLongFromBytes([]byte{0x8f,0x65,0x9e,0xff,}).ToObject()
			πTemp003[221] = πg.NewLongFromBytes([]byte{0xf8,0x62,0xae,0x69,}).ToObject()
			πTemp003[222] = πg.NewLongFromBytes([]byte{0x61,0x6b,0xff,0xd3,}).ToObject()
			πTemp003[223] = πg.NewLongFromBytes([]byte{0x16,0x6c,0xcf,0x45,}).ToObject()
			πTemp003[224] = πg.NewLongFromBytes([]byte{0xa0,0xa,0xe2,0x78,}).ToObject()
			πTemp003[225] = πg.NewLongFromBytes([]byte{0xd7,0xd,0xd2,0xee,}).ToObject()
			πTemp003[226] = πg.NewLongFromBytes([]byte{0x4e,0x4,0x83,0x54,}).ToObject()
			πTemp003[227] = πg.NewLongFromBytes([]byte{0x39,0x3,0xb3,0xc2,}).ToObject()
			πTemp003[228] = πg.NewLongFromBytes([]byte{0xa7,0x67,0x26,0x61,}).ToObject()
			πTemp003[229] = πg.NewLongFromBytes([]byte{0xd0,0x60,0x16,0xf7,}).ToObject()
			πTemp003[230] = πg.NewLongFromBytes([]byte{0x49,0x69,0x47,0x4d,}).ToObject()
			πTemp003[231] = πg.NewLongFromBytes([]byte{0x3e,0x6e,0x77,0xdb,}).ToObject()
			πTemp003[232] = πg.NewLongFromBytes([]byte{0xae,0xd1,0x6a,0x4a,}).ToObject()
			πTemp003[233] = πg.NewLongFromBytes([]byte{0xd9,0xd6,0x5a,0xdc,}).ToObject()
			πTemp003[234] = πg.NewLongFromBytes([]byte{0x40,0xdf,0xb,0x66,}).ToObject()
			πTemp003[235] = πg.NewLongFromBytes([]byte{0x37,0xd8,0x3b,0xf0,}).ToObject()
			πTemp003[236] = πg.NewLongFromBytes([]byte{0xa9,0xbc,0xae,0x53,}).ToObject()
			πTemp003[237] = πg.NewLongFromBytes([]byte{0xde,0xbb,0x9e,0xc5,}).ToObject()
			πTemp003[238] = πg.NewLongFromBytes([]byte{0x47,0xb2,0xcf,0x7f,}).ToObject()
			πTemp003[239] = πg.NewLongFromBytes([]byte{0x30,0xb5,0xff,0xe9,}).ToObject()
			πTemp003[240] = πg.NewLongFromBytes([]byte{0xbd,0xbd,0xf2,0x1c,}).ToObject()
			πTemp003[241] = πg.NewLongFromBytes([]byte{0xca,0xba,0xc2,0x8a,}).ToObject()
			πTemp003[242] = πg.NewLongFromBytes([]byte{0x53,0xb3,0x93,0x30,}).ToObject()
			πTemp003[243] = πg.NewLongFromBytes([]byte{0x24,0xb4,0xa3,0xa6,}).ToObject()
			πTemp003[244] = πg.NewLongFromBytes([]byte{0xba,0xd0,0x36,0x5,}).ToObject()
			πTemp003[245] = πg.NewLongFromBytes([]byte{0xcd,0xd7,0x6,0x93,}).ToObject()
			πTemp003[246] = πg.NewLongFromBytes([]byte{0x54,0xde,0x57,0x29,}).ToObject()
			πTemp003[247] = πg.NewLongFromBytes([]byte{0x23,0xd9,0x67,0xbf,}).ToObject()
			πTemp003[248] = πg.NewLongFromBytes([]byte{0xb3,0x66,0x7a,0x2e,}).ToObject()
			πTemp003[249] = πg.NewLongFromBytes([]byte{0xc4,0x61,0x4a,0xb8,}).ToObject()
			πTemp003[250] = πg.NewLongFromBytes([]byte{0x5d,0x68,0x1b,0x2,}).ToObject()
			πTemp003[251] = πg.NewLongFromBytes([]byte{0x2a,0x6f,0x2b,0x94,}).ToObject()
			πTemp003[252] = πg.NewLongFromBytes([]byte{0xb4,0xb,0xbe,0x37,}).ToObject()
			πTemp003[253] = πg.NewLongFromBytes([]byte{0xc3,0xc,0x8e,0xa1,}).ToObject()
			πTemp003[254] = πg.NewLongFromBytes([]byte{0x5a,0x5,0xdf,0x1b,}).ToObject()
			πTemp003[255] = πg.NewLongFromBytes([]byte{0x2d,0x2,0xef,0x8d,}).ToObject()
			πTemp018 = πg.NewList(πTemp003...).ToObject()
			if πE = πF.Globals().SetItem(πF, ßcrc_32_tab.ToObject(), πTemp018); πE != nil {
				continue
			}
			// line 655: def crc32(s, crc=0):
			πF.SetLineno(655)
			πTemp006 = make([]πg.Param, 2)
			πTemp006[0] = πg.Param{Name: "s", Def: nil}
			πTemp006[1] = πg.Param{Name: "crc", Def: πg.NewInt(0).ToObject()}
			πTemp018 = πg.NewFunction(πg.NewCode("crc32", "/usr/lib/python2.7/binascii.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µs *πg.Object = πArgs[0]; _ = µs
				var µcrc *πg.Object = πArgs[1]; _ = µcrc
				var µresult *πg.Object = πg.UnboundLocal; _ = µresult
				var µc *πg.Object = πg.UnboundLocal; _ = µc
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
				var πTemp009 []*πg.Object
				_ = πTemp009
				var πTemp010 *πg.Object
				_ = πTemp010
				var πTemp011 *πg.Object
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
					// line 656: result = 0
					πF.SetLineno(656)
					µresult = πg.NewInt(0).ToObject()
					// line 657: crc = ~long(crc) & 0xffffffffL
					πF.SetLineno(657)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µcrc, "crc"); πE != nil {
						continue
					}
					πTemp003[0] = µcrc
					if πTemp004, πE = πg.ResolveGlobal(πF, ßlong); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp002, πE = πg.Invert(πF, πTemp005); πE != nil {
						continue
					}
					if πTemp001, πE = πg.And(πF, πTemp002, πg.NewLongFromBytes([]byte{0xff,0xff,0xff,0xff,}).ToObject()); πE != nil {
						continue
					}
					µcrc = πTemp001
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Iter(πF, µs); πE != nil {
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
						µc = πTemp002
					}
					if πE != nil || !πTemp007 {
						continue
					}
					πF.PushCheckpoint(1)            
					// line 659: crc = crc_32_tab[(crc ^ long(ord(c))) & 0xffL] ^ (crc >> 8)
					πF.SetLineno(659)
					if πE = πg.CheckLocal(πF, µcrc, "crc"); πE != nil {
						continue
					}
					πTemp003 = πF.MakeArgs(1)
					πTemp009 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					πTemp009[0] = µc
					if πTemp010, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
						continue
					}
					if πTemp011, πE = πTemp010.Call(πF, πTemp009, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp009)
					πTemp003[0] = πTemp011
					if πTemp010, πE = πg.ResolveGlobal(πF, ßlong); πE != nil {
						continue
					}
					if πTemp011, πE = πTemp010.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp008, πE = πg.Xor(πF, µcrc, πTemp011); πE != nil {
						continue
					}
					if πTemp005, πE = πg.And(πF, πTemp008, πg.NewLongFromBytes([]byte{0xff,}).ToObject()); πE != nil {
						continue
					}
					πTemp004 = πTemp005
					if πTemp008, πE = πg.ResolveGlobal(πF, ßcrc_32_tab); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, πTemp008, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcrc, "crc"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.RShift(πF, µcrc, πg.NewInt(8).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Xor(πF, πTemp005, πTemp004); πE != nil {
						continue
					}
					µcrc = πTemp002
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
					// line 662: result = crc ^ 0xffffffffL
					πF.SetLineno(662)
					if πE = πg.CheckLocal(πF, µcrc, "crc"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Xor(πF, µcrc, πg.NewLongFromBytes([]byte{0xff,0xff,0xff,0xff,}).ToObject()); πE != nil {
						continue
					}
					µresult = πTemp001
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.LShift(πF, πg.NewInt(1).ToObject(), πg.NewInt(31).ToObject()); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GT(πF, µresult, πTemp002); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label4
					}
					goto Label5
					// line 664: if result > (1 << 31):
					πF.SetLineno(664)
				Label4:
					// line 665: result = ((result + (1<<31)) % (1<<32)) - (1<<31)
					πF.SetLineno(665)
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.LShift(πF, πg.NewInt(1).ToObject(), πg.NewInt(31).ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Add(πF, µresult, πTemp005); πE != nil {
						continue
					}
					if πTemp005, πE = πg.LShift(πF, πg.NewInt(1).ToObject(), πg.NewInt(32).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mod(πF, πTemp004, πTemp005); πE != nil {
						continue
					}
					if πTemp004, πE = πg.LShift(πF, πg.NewInt(1).ToObject(), πg.NewInt(31).ToObject()); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Sub(πF, πTemp002, πTemp004); πE != nil {
						continue
					}
					µresult = πTemp001
					goto Label5
				Label5:
					// line 667: return result
					πF.SetLineno(667)
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
			if πE = πF.Globals().SetItem(πF, ßcrc32.ToObject(), πTemp018); πE != nil {
				continue
			}
			// line 669: def b2a_hex(s):
			πF.SetLineno(669)
			πTemp006 = make([]πg.Param, 1)
			πTemp006[0] = πg.Param{Name: "s", Def: nil}
			πTemp019 = πg.NewFunction(πg.NewCode("b2a_hex", "/usr/lib/python2.7/binascii.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µs *πg.Object = πArgs[0]; _ = µs
				var µresult *πg.Object = πg.UnboundLocal; _ = µresult
				var µchar *πg.Object = πg.UnboundLocal; _ = µchar
				var µc *πg.Object = πg.UnboundLocal; _ = µc
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
				var πTemp008 *πg.Object
				_ = πTemp008
				var πTemp009 []*πg.Object
				_ = πTemp009
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 1: goto Label1
					case 2: goto Label2
					default: panic("unexpected function state")
					}
					// line 670: result = []
					πF.SetLineno(670)
					πTemp001 = make([]*πg.Object, 0)
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					µresult = πTemp002
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Iter(πF, µs); πE != nil {
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
					if πTemp005, πE = πg.Next(πF, πTemp002); πE != nil {
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
						µchar = πTemp005
					}
					if πE != nil || !πTemp004 {
						continue
					}
					πF.PushCheckpoint(1)            
					// line 672: c = (ord(char) >> 4) & 0xf
					πF.SetLineno(672)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µchar, "char"); πE != nil {
						continue
					}
					πTemp001[0] = µchar
					if πTemp007, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp007.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp006, πE = πg.RShift(πF, πTemp008, πg.NewInt(4).ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.And(πF, πTemp006, πg.NewInt(15).ToObject()); πE != nil {
						continue
					}
					µc = πTemp005
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GT(πF, µc, πg.NewInt(9).ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label4
					}
					goto Label5
					// line 673: if c > 9:
					πF.SetLineno(673)
				Label4:
					// line 674: c = c + ord('a') - 10
					πF.SetLineno(674)
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = ßa.ToObject()
					if πTemp007, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp007.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp006, πE = πg.Add(πF, µc, πTemp008); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Sub(πF, πTemp006, πg.NewInt(10).ToObject()); πE != nil {
						continue
					}
					µc = πTemp005
					goto Label6
				Label5:
					// line 676: c = c + ord('0')
					πF.SetLineno(676)
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = ß0.ToObject()
					if πTemp006, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp005, πE = πg.Add(πF, µc, πTemp007); πE != nil {
						continue
					}
					µc = πTemp005
					goto Label6
				Label6:
					// line 677: result.append(chr(c))
					πF.SetLineno(677)
					πTemp001 = πF.MakeArgs(1)
					πTemp009 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					πTemp009[0] = µc
					if πTemp005, πE = πg.ResolveGlobal(πF, ßchr); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp009, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp009)
					πTemp001[0] = πTemp006
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µresult, ßappend, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 678: c = ord(char) & 0xf
					πF.SetLineno(678)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µchar, "char"); πE != nil {
						continue
					}
					πTemp001[0] = µchar
					if πTemp006, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp005, πE = πg.And(πF, πTemp007, πg.NewInt(15).ToObject()); πE != nil {
						continue
					}
					µc = πTemp005
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GT(πF, µc, πg.NewInt(9).ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label7
					}
					goto Label8
					// line 679: if c > 9:
					πF.SetLineno(679)
				Label7:
					// line 680: c = c + ord('a') - 10
					πF.SetLineno(680)
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = ßa.ToObject()
					if πTemp007, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp007.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp006, πE = πg.Add(πF, µc, πTemp008); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Sub(πF, πTemp006, πg.NewInt(10).ToObject()); πE != nil {
						continue
					}
					µc = πTemp005
					goto Label9
				Label8:
					// line 682: c = c + ord('0')
					πF.SetLineno(682)
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = ß0.ToObject()
					if πTemp006, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp005, πE = πg.Add(πF, µc, πTemp007); πE != nil {
						continue
					}
					µc = πTemp005
					goto Label9
				Label9:
					// line 683: result.append(chr(c))
					πF.SetLineno(683)
					πTemp001 = πF.MakeArgs(1)
					πTemp009 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					πTemp009[0] = µc
					if πTemp005, πE = πg.ResolveGlobal(πF, ßchr); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp009, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp009)
					πTemp001[0] = πTemp006
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µresult, ßappend, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
					// line 684: return ''.join(result)
					πF.SetLineno(684)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					πTemp001[0] = µresult
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
			if πE = πF.Globals().SetItem(πF, ßb2a_hex.ToObject(), πTemp019); πE != nil {
				continue
			}
			// line 686: hexlify = b2a_hex
			πF.SetLineno(686)
			if πTemp020, πE = πg.ResolveGlobal(πF, ßb2a_hex); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßhexlify.ToObject(), πTemp020); πE != nil {
				continue
			}
			// line 688: table_hex = [
			πF.SetLineno(688)
			πTemp003 = make([]*πg.Object, 128)
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[0] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[1] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[2] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[3] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[4] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[5] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[6] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[7] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[8] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[9] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[10] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[11] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[12] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[13] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[14] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[15] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[16] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[17] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[18] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[19] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[20] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[21] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[22] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[23] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[24] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[25] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[26] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[27] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[28] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[29] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[30] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[31] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[32] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[33] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[34] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[35] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[36] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[37] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[38] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[39] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[40] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[41] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[42] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[43] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[44] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[45] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[46] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[47] = πTemp020
			πTemp003[48] = πg.NewInt(0).ToObject()
			πTemp003[49] = πg.NewInt(1).ToObject()
			πTemp003[50] = πg.NewInt(2).ToObject()
			πTemp003[51] = πg.NewInt(3).ToObject()
			πTemp003[52] = πg.NewInt(4).ToObject()
			πTemp003[53] = πg.NewInt(5).ToObject()
			πTemp003[54] = πg.NewInt(6).ToObject()
			πTemp003[55] = πg.NewInt(7).ToObject()
			πTemp003[56] = πg.NewInt(8).ToObject()
			πTemp003[57] = πg.NewInt(9).ToObject()
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[58] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[59] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[60] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[61] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[62] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[63] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[64] = πTemp020
			πTemp003[65] = πg.NewInt(10).ToObject()
			πTemp003[66] = πg.NewInt(11).ToObject()
			πTemp003[67] = πg.NewInt(12).ToObject()
			πTemp003[68] = πg.NewInt(13).ToObject()
			πTemp003[69] = πg.NewInt(14).ToObject()
			πTemp003[70] = πg.NewInt(15).ToObject()
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[71] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[72] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[73] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[74] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[75] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[76] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[77] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[78] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[79] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[80] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[81] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[82] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[83] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[84] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[85] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[86] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[87] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[88] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[89] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[90] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[91] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[92] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[93] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[94] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[95] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[96] = πTemp020
			πTemp003[97] = πg.NewInt(10).ToObject()
			πTemp003[98] = πg.NewInt(11).ToObject()
			πTemp003[99] = πg.NewInt(12).ToObject()
			πTemp003[100] = πg.NewInt(13).ToObject()
			πTemp003[101] = πg.NewInt(14).ToObject()
			πTemp003[102] = πg.NewInt(15).ToObject()
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[103] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[104] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[105] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[106] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[107] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[108] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[109] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[110] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[111] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[112] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[113] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[114] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[115] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[116] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[117] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[118] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[119] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[120] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[121] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[122] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[123] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[124] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[125] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[126] = πTemp020
			if πTemp020, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp003[127] = πTemp020
			πTemp020 = πg.NewList(πTemp003...).ToObject()
			if πE = πF.Globals().SetItem(πF, ßtable_hex.ToObject(), πTemp020); πE != nil {
				continue
			}
			// line 700: def a2b_hex(t):
			πF.SetLineno(700)
			πTemp006 = make([]πg.Param, 1)
			πTemp006[0] = πg.Param{Name: "t", Def: nil}
			πTemp020 = πg.NewFunction(πg.NewCode("a2b_hex", "/usr/lib/python2.7/binascii.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µt *πg.Object = πArgs[0]; _ = µt
				var µresult *πg.Object = πg.UnboundLocal; _ = µresult
				var µpairs_gen *πg.Object = πg.UnboundLocal; _ = µpairs_gen
				var µa *πg.Object = πg.UnboundLocal; _ = µa
				var µb *πg.Object = πg.UnboundLocal; _ = µb
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
				var πTemp006 bool
				_ = πTemp006
				var πTemp007 bool
				_ = πTemp007
				var πTemp008 *πg.Object
				_ = πTemp008
				var πTemp009 *πg.Object
				_ = πTemp009
				var πTemp010 []*πg.Object
				_ = πTemp010
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 1: goto Label1
					case 2: goto Label2
					default: panic("unexpected function state")
					}
					// line 701: result = []
					πF.SetLineno(701)
					πTemp001 = make([]*πg.Object, 0)
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					µresult = πTemp002
					// line 703: def pairs_gen(s):
					πF.SetLineno(703)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "s", Def: nil}
					πTemp002 = πg.NewFunction(πg.NewCode("pairs_gen", "/usr/lib/python2.7/binascii.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µs *πg.Object = πArgs[0]; _ = µs
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
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 *πg.BaseException
						_ = πTemp009
						var πTemp010 *πg.Traceback
						_ = πTemp010
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
							for ; πF.State() >= 0; πF.PopCheckpoint() {
								switch πF.State() {
								case 0:
								case 1: goto Label1
								case 2: goto Label2
								case 5: goto Label5
								case 6: goto Label6
								default: panic("unexpected function state")
								}
								// line 704: while s:
								πF.SetLineno(704)
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
								if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.IsTrue(πF, µs); πE != nil {
									continue
								}
								if πE != nil || !πTemp002 {
									continue
								}
								πF.PushCheckpoint(1)            
								// line 705: try:
								πF.SetLineno(705)
								πF.PushCheckpoint(5)
								// line 706: yield table_hex[ord(s[0])], table_hex[ord(s[1])]
								πF.SetLineno(706)
								πTemp005 = πF.MakeArgs(1)
								πTemp006 = πg.NewInt(0).ToObject()
								if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
									continue
								}
								if πTemp007, πE = πg.GetItem(πF, µs, πTemp006); πE != nil {
									continue
								}
								πTemp005[0] = πTemp007
								if πTemp006, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
									continue
								}
								if πTemp007, πE = πTemp006.Call(πF, πTemp005, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp005)
								πTemp004 = πTemp007
								if πTemp007, πE = πg.ResolveGlobal(πF, ßtable_hex); πE != nil {
									continue
								}
								if πTemp006, πE = πg.GetItem(πF, πTemp007, πTemp004); πE != nil {
									continue
								}
								πTemp005 = πF.MakeArgs(1)
								πTemp007 = πg.NewInt(1).ToObject()
								if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
									continue
								}
								if πTemp008, πE = πg.GetItem(πF, µs, πTemp007); πE != nil {
									continue
								}
								πTemp005[0] = πTemp008
								if πTemp007, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
									continue
								}
								if πTemp008, πE = πTemp007.Call(πF, πTemp005, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp005)
								πTemp004 = πTemp008
								if πTemp008, πE = πg.ResolveGlobal(πF, ßtable_hex); πE != nil {
									continue
								}
								if πTemp007, πE = πg.GetItem(πF, πTemp008, πTemp004); πE != nil {
									continue
								}
								πTemp003 = πg.NewTuple2(πTemp006, πTemp007).ToObject()
								πF.PushCheckpoint(6)
								return πTemp003, nil
							Label6:
								πTemp004 = πSent
								πF.PopCheckpoint()
								goto Label4
							Label5:
								if πE == nil {
								  continue
								}
								πE = nil
								πTemp009, πTemp010 = πF.ExcInfo()
								if πTemp004, πE = πg.ResolveGlobal(πF, ßIndexError); πE != nil {
									continue
								}
								if πTemp002, πE = πg.IsInstance(πF, πTemp009.ToObject(), πTemp004); πE != nil {
									continue
								}
								if πTemp002 {
									goto Label7
								}
								πE = πF.Raise(πTemp009.ToObject(), nil, πTemp010.ToObject())
								continue
								// line 707: except IndexError:
								πF.SetLineno(707)
							Label7:
								πTemp005 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
									continue
								}
								πTemp005[0] = µs
								if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
									continue
								}
								if πTemp006, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp005)
								if πTemp002, πE = πg.IsTrue(πF, πTemp006); πE != nil {
									continue
								}
								if πTemp002 {
									goto Label8
								}
								goto Label9
								// line 708: if len(s):
								πF.SetLineno(708)
							Label8:
								πTemp005 = πF.MakeArgs(1)
								πTemp005[0] = πg.NewStr("Odd-length string").ToObject()
								if πTemp004, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
									continue
								}
								if πTemp006, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp005)
								// line 709: raise TypeError('Odd-length string')
								πF.SetLineno(709)
								πE = πF.Raise(πTemp006, nil, nil)
								continue
								goto Label9
							Label9:
								// line 710: return
								πF.SetLineno(710)
								πR = πg.None
								continue
								πF.RestoreExc(nil, nil)
								goto Label4
							Label4:
								// line 711: s = s[2:]
								πF.SetLineno(711)
								if πTemp004, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(2).ToObject(), πg.None, πg.None}, nil); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
									continue
								}
								if πTemp006, πE = πg.GetItem(πF, µs, πTemp004); πE != nil {
									continue
								}
								µs = πTemp006
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
					µpairs_gen = πTemp002
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
						continue
					}
					πTemp001[0] = µt
					if πE = πg.CheckLocal(πF, µpairs_gen, "pairs_gen"); πE != nil {
						continue
					}
					if πTemp005, πE = µpairs_gen.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp004, πE = πg.Iter(πF, πTemp005); πE != nil {
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
					if πTemp005, πE = πg.Next(πF, πTemp004); πE != nil {
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
						if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp008}, πg.TieTarget{Target: &πTemp009}}}, πTemp005); πE != nil {
							continue
						}
						µa = πTemp008
						µb = πTemp009
					}
					if πE != nil || !πTemp007 {
						continue
					}
					πF.PushCheckpoint(1)            
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.LT(πF, µa, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					πTemp005 = πTemp008
					if πTemp007, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label4
					}
					if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.LT(πF, µb, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					πTemp005 = πTemp008
				Label4:
					if πTemp007, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label5
					}
					goto Label6
					// line 714: if a < 0 or b < 0:
					πF.SetLineno(714)
				Label5:
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("Non-hexadecimal digit found").ToObject()
					if πTemp005, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 715: raise TypeError('Non-hexadecimal digit found')
					πF.SetLineno(715)
					πE = πF.Raise(πTemp008, nil, nil)
					continue
					goto Label6
				Label6:
					// line 716: result.append(chr((a << 4) + b))
					πF.SetLineno(716)
					πTemp001 = πF.MakeArgs(1)
					πTemp010 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.LShift(πF, µa, πg.NewInt(4).ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Add(πF, πTemp008, µb); πE != nil {
						continue
					}
					πTemp010[0] = πTemp005
					if πTemp005, πE = πg.ResolveGlobal(πF, ßchr); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp005.Call(πF, πTemp010, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp010)
					πTemp001[0] = πTemp008
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µresult, ßappend, nil); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
					// line 717: return ''.join(result)
					πF.SetLineno(717)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					πTemp001[0] = µresult
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
			if πE = πF.Globals().SetItem(πF, ßa2b_hex.ToObject(), πTemp020); πE != nil {
				continue
			}
			// line 720: unhexlify = a2b_hex
			πF.SetLineno(720)
			if πTemp021, πE = πg.ResolveGlobal(πF, ßa2b_hex); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßunhexlify.ToObject(), πTemp021); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("binascii", Code)
}

