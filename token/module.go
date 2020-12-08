package token
import (
	πg "github.com/pygolin/runtime"
	// _ "github.com/pygolin/stdlib/sys"
	// _ "github.com/pygolin/stdlib/re"
)
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/token.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ßAMPER := πg.InternStr("AMPER")
		ßAMPEREQUAL := πg.InternStr("AMPEREQUAL")
		ßASYNC := πg.InternStr("ASYNC")
		ßAT := πg.InternStr("AT")
		ßATEQUAL := πg.InternStr("ATEQUAL")
		ßAWAIT := πg.InternStr("AWAIT")
		ßCIRCUMFLEX := πg.InternStr("CIRCUMFLEX")
		ßCIRCUMFLEXEQUAL := πg.InternStr("CIRCUMFLEXEQUAL")
		ßCOLON := πg.InternStr("COLON")
		ßCOMMA := πg.InternStr("COMMA")
		ßDEDENT := πg.InternStr("DEDENT")
		ßDOT := πg.InternStr("DOT")
		ßDOUBLESLASH := πg.InternStr("DOUBLESLASH")
		ßDOUBLESLASHEQUAL := πg.InternStr("DOUBLESLASHEQUAL")
		ßDOUBLESTAR := πg.InternStr("DOUBLESTAR")
		ßDOUBLESTAREQUAL := πg.InternStr("DOUBLESTAREQUAL")
		ßELLIPSIS := πg.InternStr("ELLIPSIS")
		ßENDMARKER := πg.InternStr("ENDMARKER")
		ßEQEQUAL := πg.InternStr("EQEQUAL")
		ßEQUAL := πg.InternStr("EQUAL")
		ßERRORTOKEN := πg.InternStr("ERRORTOKEN")
		ßGREATER := πg.InternStr("GREATER")
		ßGREATEREQUAL := πg.InternStr("GREATEREQUAL")
		ßIGNORECASE := πg.InternStr("IGNORECASE")
		ßINDENT := πg.InternStr("INDENT")
		ßISEOF := πg.InternStr("ISEOF")
		ßISNONTERMINAL := πg.InternStr("ISNONTERMINAL")
		ßISTERMINAL := πg.InternStr("ISTERMINAL")
		ßLBRACE := πg.InternStr("LBRACE")
		ßLEFTSHIFT := πg.InternStr("LEFTSHIFT")
		ßLEFTSHIFTEQUAL := πg.InternStr("LEFTSHIFTEQUAL")
		ßLESS := πg.InternStr("LESS")
		ßLESSEQUAL := πg.InternStr("LESSEQUAL")
		ßLPAR := πg.InternStr("LPAR")
		ßLSQB := πg.InternStr("LSQB")
		ßMINEQUAL := πg.InternStr("MINEQUAL")
		ßMINUS := πg.InternStr("MINUS")
		ßNAME := πg.InternStr("NAME")
		ßNEWLINE := πg.InternStr("NEWLINE")
		ßNOTEQUAL := πg.InternStr("NOTEQUAL")
		ßNT_OFFSET := πg.InternStr("NT_OFFSET")
		ßNUMBER := πg.InternStr("NUMBER")
		ßN_TOKENS := πg.InternStr("N_TOKENS")
		ßOP := πg.InternStr("OP")
		ßOSError := πg.InternStr("OSError")
		ßPERCENT := πg.InternStr("PERCENT")
		ßPERCENTEQUAL := πg.InternStr("PERCENTEQUAL")
		ßPLUS := πg.InternStr("PLUS")
		ßPLUSEQUAL := πg.InternStr("PLUSEQUAL")
		ßRARROW := πg.InternStr("RARROW")
		ßRBRACE := πg.InternStr("RBRACE")
		ßRIGHTSHIFT := πg.InternStr("RIGHTSHIFT")
		ßRIGHTSHIFTEQUAL := πg.InternStr("RIGHTSHIFTEQUAL")
		ßRPAR := πg.InternStr("RPAR")
		ßRSQB := πg.InternStr("RSQB")
		ßSEMI := πg.InternStr("SEMI")
		ßSLASH := πg.InternStr("SLASH")
		ßSLASHEQUAL := πg.InternStr("SLASHEQUAL")
		ßSTAR := πg.InternStr("STAR")
		ßSTAREQUAL := πg.InternStr("STAREQUAL")
		ßSTRING := πg.InternStr("STRING")
		ßTILDE := πg.InternStr("TILDE")
		ßVBAR := πg.InternStr("VBAR")
		ßVBAREQUAL := πg.InternStr("VBAREQUAL")
		ßValueError := πg.InternStr("ValueError")
		ß_ := πg.InternStr("_")
		ß__all__ := πg.InternStr("__all__")
		ß__doc__ := πg.InternStr("__doc__")
		ß__enter__ := πg.InternStr("__enter__")
		ß__exit__ := πg.InternStr("__exit__")
		ß__main := πg.InternStr("__main")
		ß__main__ := πg.InternStr("__main__")
		ß__name__ := πg.InternStr("__name__")
		ßappend := πg.InternStr("append")
		ßargv := πg.InternStr("argv")
		ßcompile := πg.InternStr("compile")
		ßexit := πg.InternStr("exit")
		ßextend := πg.InternStr("extend")
		ßglobals := πg.InternStr("globals")
		ßgroup := πg.InternStr("group")
		ßindex := πg.InternStr("index")
		ßint := πg.InternStr("int")
		ßisinstance := πg.InternStr("isinstance")
		ßitems := πg.InternStr("items")
		ßjoin := πg.InternStr("join")
		ßkeys := πg.InternStr("keys")
		ßlen := πg.InternStr("len")
		ßmatch := πg.InternStr("match")
		ßopen := πg.InternStr("open")
		ßread := πg.InternStr("read")
		ßsorted := πg.InternStr("sorted")
		ßsplit := πg.InternStr("split")
		ßstartswith := πg.InternStr("startswith")
		ßstderr := πg.InternStr("stderr")
		ßstdout := πg.InternStr("stdout")
		ßstr := πg.InternStr("str")
		ßtok_name := πg.InternStr("tok_name")
		ßvalues := πg.InternStr("values")
		ßw := πg.InternStr("w")
		ßwrite := πg.InternStr("write")
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
		var πTemp007 *πg.Object
		_ = πTemp007
		var πTemp008 *πg.Object
		_ = πTemp008
		var πTemp009 *πg.Object
		_ = πTemp009
		var πTemp010 bool
		_ = πTemp010
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default: panic("unexpected function state")
			}
			// line 1: """Token constants (from "token.h")."""
			πF.SetLineno(1)
			// line 1: """Token constants (from "token.h")."""
			πF.SetLineno(1)
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Token constants (from \"token.h\").").ToObject()); πE != nil {
				continue
			}
			// line 3: __all__ = ['tok_name', 'ISTERMINAL', 'ISNONTERMINAL', 'ISEOF']
			πF.SetLineno(3)
			πTemp001 = make([]*πg.Object, 4)
			πTemp001[0] = ßtok_name.ToObject()
			πTemp001[1] = ßISTERMINAL.ToObject()
			πTemp001[2] = ßISNONTERMINAL.ToObject()
			πTemp001[3] = ßISEOF.ToObject()
			πTemp002 = πg.NewList(πTemp001...).ToObject()
			if πE = πF.Globals().SetItem(πF, ß__all__.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 13: ENDMARKER = 0
			πF.SetLineno(13)
			if πE = πF.Globals().SetItem(πF, ßENDMARKER.ToObject(), πg.NewInt(0).ToObject()); πE != nil {
				continue
			}
			// line 14: NAME = 1
			πF.SetLineno(14)
			if πE = πF.Globals().SetItem(πF, ßNAME.ToObject(), πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			// line 15: NUMBER = 2
			πF.SetLineno(15)
			if πE = πF.Globals().SetItem(πF, ßNUMBER.ToObject(), πg.NewInt(2).ToObject()); πE != nil {
				continue
			}
			// line 16: STRING = 3
			πF.SetLineno(16)
			if πE = πF.Globals().SetItem(πF, ßSTRING.ToObject(), πg.NewInt(3).ToObject()); πE != nil {
				continue
			}
			// line 17: NEWLINE = 4
			πF.SetLineno(17)
			if πE = πF.Globals().SetItem(πF, ßNEWLINE.ToObject(), πg.NewInt(4).ToObject()); πE != nil {
				continue
			}
			// line 18: INDENT = 5
			πF.SetLineno(18)
			if πE = πF.Globals().SetItem(πF, ßINDENT.ToObject(), πg.NewInt(5).ToObject()); πE != nil {
				continue
			}
			// line 19: DEDENT = 6
			πF.SetLineno(19)
			if πE = πF.Globals().SetItem(πF, ßDEDENT.ToObject(), πg.NewInt(6).ToObject()); πE != nil {
				continue
			}
			// line 20: LPAR = 7
			πF.SetLineno(20)
			if πE = πF.Globals().SetItem(πF, ßLPAR.ToObject(), πg.NewInt(7).ToObject()); πE != nil {
				continue
			}
			// line 21: RPAR = 8
			πF.SetLineno(21)
			if πE = πF.Globals().SetItem(πF, ßRPAR.ToObject(), πg.NewInt(8).ToObject()); πE != nil {
				continue
			}
			// line 22: LSQB = 9
			πF.SetLineno(22)
			if πE = πF.Globals().SetItem(πF, ßLSQB.ToObject(), πg.NewInt(9).ToObject()); πE != nil {
				continue
			}
			// line 23: RSQB = 10
			πF.SetLineno(23)
			if πE = πF.Globals().SetItem(πF, ßRSQB.ToObject(), πg.NewInt(10).ToObject()); πE != nil {
				continue
			}
			// line 24: COLON = 11
			πF.SetLineno(24)
			if πE = πF.Globals().SetItem(πF, ßCOLON.ToObject(), πg.NewInt(11).ToObject()); πE != nil {
				continue
			}
			// line 25: COMMA = 12
			πF.SetLineno(25)
			if πE = πF.Globals().SetItem(πF, ßCOMMA.ToObject(), πg.NewInt(12).ToObject()); πE != nil {
				continue
			}
			// line 26: SEMI = 13
			πF.SetLineno(26)
			if πE = πF.Globals().SetItem(πF, ßSEMI.ToObject(), πg.NewInt(13).ToObject()); πE != nil {
				continue
			}
			// line 27: PLUS = 14
			πF.SetLineno(27)
			if πE = πF.Globals().SetItem(πF, ßPLUS.ToObject(), πg.NewInt(14).ToObject()); πE != nil {
				continue
			}
			// line 28: MINUS = 15
			πF.SetLineno(28)
			if πE = πF.Globals().SetItem(πF, ßMINUS.ToObject(), πg.NewInt(15).ToObject()); πE != nil {
				continue
			}
			// line 29: STAR = 16
			πF.SetLineno(29)
			if πE = πF.Globals().SetItem(πF, ßSTAR.ToObject(), πg.NewInt(16).ToObject()); πE != nil {
				continue
			}
			// line 30: SLASH = 17
			πF.SetLineno(30)
			if πE = πF.Globals().SetItem(πF, ßSLASH.ToObject(), πg.NewInt(17).ToObject()); πE != nil {
				continue
			}
			// line 31: VBAR = 18
			πF.SetLineno(31)
			if πE = πF.Globals().SetItem(πF, ßVBAR.ToObject(), πg.NewInt(18).ToObject()); πE != nil {
				continue
			}
			// line 32: AMPER = 19
			πF.SetLineno(32)
			if πE = πF.Globals().SetItem(πF, ßAMPER.ToObject(), πg.NewInt(19).ToObject()); πE != nil {
				continue
			}
			// line 33: LESS = 20
			πF.SetLineno(33)
			if πE = πF.Globals().SetItem(πF, ßLESS.ToObject(), πg.NewInt(20).ToObject()); πE != nil {
				continue
			}
			// line 34: GREATER = 21
			πF.SetLineno(34)
			if πE = πF.Globals().SetItem(πF, ßGREATER.ToObject(), πg.NewInt(21).ToObject()); πE != nil {
				continue
			}
			// line 35: EQUAL = 22
			πF.SetLineno(35)
			if πE = πF.Globals().SetItem(πF, ßEQUAL.ToObject(), πg.NewInt(22).ToObject()); πE != nil {
				continue
			}
			// line 36: DOT = 23
			πF.SetLineno(36)
			if πE = πF.Globals().SetItem(πF, ßDOT.ToObject(), πg.NewInt(23).ToObject()); πE != nil {
				continue
			}
			// line 37: PERCENT = 24
			πF.SetLineno(37)
			if πE = πF.Globals().SetItem(πF, ßPERCENT.ToObject(), πg.NewInt(24).ToObject()); πE != nil {
				continue
			}
			// line 38: LBRACE = 25
			πF.SetLineno(38)
			if πE = πF.Globals().SetItem(πF, ßLBRACE.ToObject(), πg.NewInt(25).ToObject()); πE != nil {
				continue
			}
			// line 39: RBRACE = 26
			πF.SetLineno(39)
			if πE = πF.Globals().SetItem(πF, ßRBRACE.ToObject(), πg.NewInt(26).ToObject()); πE != nil {
				continue
			}
			// line 40: EQEQUAL = 27
			πF.SetLineno(40)
			if πE = πF.Globals().SetItem(πF, ßEQEQUAL.ToObject(), πg.NewInt(27).ToObject()); πE != nil {
				continue
			}
			// line 41: NOTEQUAL = 28
			πF.SetLineno(41)
			if πE = πF.Globals().SetItem(πF, ßNOTEQUAL.ToObject(), πg.NewInt(28).ToObject()); πE != nil {
				continue
			}
			// line 42: LESSEQUAL = 29
			πF.SetLineno(42)
			if πE = πF.Globals().SetItem(πF, ßLESSEQUAL.ToObject(), πg.NewInt(29).ToObject()); πE != nil {
				continue
			}
			// line 43: GREATEREQUAL = 30
			πF.SetLineno(43)
			if πE = πF.Globals().SetItem(πF, ßGREATEREQUAL.ToObject(), πg.NewInt(30).ToObject()); πE != nil {
				continue
			}
			// line 44: TILDE = 31
			πF.SetLineno(44)
			if πE = πF.Globals().SetItem(πF, ßTILDE.ToObject(), πg.NewInt(31).ToObject()); πE != nil {
				continue
			}
			// line 45: CIRCUMFLEX = 32
			πF.SetLineno(45)
			if πE = πF.Globals().SetItem(πF, ßCIRCUMFLEX.ToObject(), πg.NewInt(32).ToObject()); πE != nil {
				continue
			}
			// line 46: LEFTSHIFT = 33
			πF.SetLineno(46)
			if πE = πF.Globals().SetItem(πF, ßLEFTSHIFT.ToObject(), πg.NewInt(33).ToObject()); πE != nil {
				continue
			}
			// line 47: RIGHTSHIFT = 34
			πF.SetLineno(47)
			if πE = πF.Globals().SetItem(πF, ßRIGHTSHIFT.ToObject(), πg.NewInt(34).ToObject()); πE != nil {
				continue
			}
			// line 48: DOUBLESTAR = 35
			πF.SetLineno(48)
			if πE = πF.Globals().SetItem(πF, ßDOUBLESTAR.ToObject(), πg.NewInt(35).ToObject()); πE != nil {
				continue
			}
			// line 49: PLUSEQUAL = 36
			πF.SetLineno(49)
			if πE = πF.Globals().SetItem(πF, ßPLUSEQUAL.ToObject(), πg.NewInt(36).ToObject()); πE != nil {
				continue
			}
			// line 50: MINEQUAL = 37
			πF.SetLineno(50)
			if πE = πF.Globals().SetItem(πF, ßMINEQUAL.ToObject(), πg.NewInt(37).ToObject()); πE != nil {
				continue
			}
			// line 51: STAREQUAL = 38
			πF.SetLineno(51)
			if πE = πF.Globals().SetItem(πF, ßSTAREQUAL.ToObject(), πg.NewInt(38).ToObject()); πE != nil {
				continue
			}
			// line 52: SLASHEQUAL = 39
			πF.SetLineno(52)
			if πE = πF.Globals().SetItem(πF, ßSLASHEQUAL.ToObject(), πg.NewInt(39).ToObject()); πE != nil {
				continue
			}
			// line 53: PERCENTEQUAL = 40
			πF.SetLineno(53)
			if πE = πF.Globals().SetItem(πF, ßPERCENTEQUAL.ToObject(), πg.NewInt(40).ToObject()); πE != nil {
				continue
			}
			// line 54: AMPEREQUAL = 41
			πF.SetLineno(54)
			if πE = πF.Globals().SetItem(πF, ßAMPEREQUAL.ToObject(), πg.NewInt(41).ToObject()); πE != nil {
				continue
			}
			// line 55: VBAREQUAL = 42
			πF.SetLineno(55)
			if πE = πF.Globals().SetItem(πF, ßVBAREQUAL.ToObject(), πg.NewInt(42).ToObject()); πE != nil {
				continue
			}
			// line 56: CIRCUMFLEXEQUAL = 43
			πF.SetLineno(56)
			if πE = πF.Globals().SetItem(πF, ßCIRCUMFLEXEQUAL.ToObject(), πg.NewInt(43).ToObject()); πE != nil {
				continue
			}
			// line 57: LEFTSHIFTEQUAL = 44
			πF.SetLineno(57)
			if πE = πF.Globals().SetItem(πF, ßLEFTSHIFTEQUAL.ToObject(), πg.NewInt(44).ToObject()); πE != nil {
				continue
			}
			// line 58: RIGHTSHIFTEQUAL = 45
			πF.SetLineno(58)
			if πE = πF.Globals().SetItem(πF, ßRIGHTSHIFTEQUAL.ToObject(), πg.NewInt(45).ToObject()); πE != nil {
				continue
			}
			// line 59: DOUBLESTAREQUAL = 46
			πF.SetLineno(59)
			if πE = πF.Globals().SetItem(πF, ßDOUBLESTAREQUAL.ToObject(), πg.NewInt(46).ToObject()); πE != nil {
				continue
			}
			// line 60: DOUBLESLASH = 47
			πF.SetLineno(60)
			if πE = πF.Globals().SetItem(πF, ßDOUBLESLASH.ToObject(), πg.NewInt(47).ToObject()); πE != nil {
				continue
			}
			// line 61: DOUBLESLASHEQUAL = 48
			πF.SetLineno(61)
			if πE = πF.Globals().SetItem(πF, ßDOUBLESLASHEQUAL.ToObject(), πg.NewInt(48).ToObject()); πE != nil {
				continue
			}
			// line 62: AT = 49
			πF.SetLineno(62)
			if πE = πF.Globals().SetItem(πF, ßAT.ToObject(), πg.NewInt(49).ToObject()); πE != nil {
				continue
			}
			// line 63: ATEQUAL = 50
			πF.SetLineno(63)
			if πE = πF.Globals().SetItem(πF, ßATEQUAL.ToObject(), πg.NewInt(50).ToObject()); πE != nil {
				continue
			}
			// line 64: RARROW = 51
			πF.SetLineno(64)
			if πE = πF.Globals().SetItem(πF, ßRARROW.ToObject(), πg.NewInt(51).ToObject()); πE != nil {
				continue
			}
			// line 65: ELLIPSIS = 52
			πF.SetLineno(65)
			if πE = πF.Globals().SetItem(πF, ßELLIPSIS.ToObject(), πg.NewInt(52).ToObject()); πE != nil {
				continue
			}
			// line 66: OP = 53
			πF.SetLineno(66)
			if πE = πF.Globals().SetItem(πF, ßOP.ToObject(), πg.NewInt(53).ToObject()); πE != nil {
				continue
			}
			// line 67: AWAIT = 54
			πF.SetLineno(67)
			if πE = πF.Globals().SetItem(πF, ßAWAIT.ToObject(), πg.NewInt(54).ToObject()); πE != nil {
				continue
			}
			// line 68: ASYNC = 55
			πF.SetLineno(68)
			if πE = πF.Globals().SetItem(πF, ßASYNC.ToObject(), πg.NewInt(55).ToObject()); πE != nil {
				continue
			}
			// line 69: ERRORTOKEN = 56
			πF.SetLineno(69)
			if πE = πF.Globals().SetItem(πF, ßERRORTOKEN.ToObject(), πg.NewInt(56).ToObject()); πE != nil {
				continue
			}
			// line 70: N_TOKENS = 57
			πF.SetLineno(70)
			if πE = πF.Globals().SetItem(πF, ßN_TOKENS.ToObject(), πg.NewInt(57).ToObject()); πE != nil {
				continue
			}
			// line 71: NT_OFFSET = 256
			πF.SetLineno(71)
			if πE = πF.Globals().SetItem(πF, ßNT_OFFSET.ToObject(), πg.NewInt(256).ToObject()); πE != nil {
				continue
			}
			// line 74: tok_name = {value: name
			πF.SetLineno(74)
			πTemp004 = make([]πg.Param, 0)
			πTemp003 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/token.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µname *πg.Object = πg.UnboundLocal; _ = µname
				var µvalue *πg.Object = πg.UnboundLocal; _ = µvalue
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
				var πTemp009 bool
				_ = πTemp009
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
					for ; πF.State() >= 0; πF.PopCheckpoint() {
						switch πF.State() {
						case 0:
						case 1: goto Label1
						case 2: goto Label2
						case 7: goto Label7
						default: panic("unexpected function state")
						}
						if πTemp002, πE = πg.ResolveGlobal(πF, ßglobals); πE != nil {
							continue
						}
						if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
							continue
						}
						if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßitems, nil); πE != nil {
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
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp006}}}, πTemp002); πE != nil {
								continue
							}
							µname = πTemp003
							µvalue = πTemp006
						}
						if πE != nil || !πTemp005 {
							continue
						}
						πF.PushCheckpoint(1)            
						πTemp007 = πF.MakeArgs(2)
						if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
							continue
						}
						πTemp007[0] = µvalue
						if πTemp003, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
							continue
						}
						πTemp007[1] = πTemp003
						if πTemp003, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
							continue
						}
						if πTemp006, πE = πTemp003.Call(πF, πTemp007, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp007)
						πTemp002 = πTemp006
						if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
							continue
						}
						if !πTemp005 {
							goto Label4
						}
						πTemp007 = πF.MakeArgs(1)
						πTemp007[0] = ß_.ToObject()
						if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
							continue
						}
						if πTemp006, πE = πg.GetAttr(πF, µname, ßstartswith, nil); πE != nil {
							continue
						}
						if πTemp008, πE = πTemp006.Call(πF, πTemp007, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp007)
						if πTemp009, πE = πg.IsTrue(πF, πTemp008); πE != nil {
							continue
						}
						πTemp003 = πg.GetBool(!πTemp009).ToObject()
						πTemp002 = πTemp003
					Label4:
						if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
							continue
						}
						if πTemp005 {
							goto Label5
						}
						goto Label6
						// line 74: tok_name = {value: name
						πF.SetLineno(74)
					Label5:
						// line 74: tok_name = {value: name
						πF.SetLineno(74)
						if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
							continue
						}
						πTemp002 = πg.NewTuple2(µvalue, µname).ToObject()
						πF.PushCheckpoint(7)
						return πTemp002, nil
					Label7:
						πTemp003 = πSent
						goto Label6
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
			if πTemp005, πE = πTemp003.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πTemp002, πE = πg.DictType.Call(πF, πg.Args{πTemp005}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßtok_name.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 77: __all__.extend(tok_name.values())
			πF.SetLineno(77)
			πTemp001 = πF.MakeArgs(1)
			if πTemp002, πE = πg.ResolveGlobal(πF, ßtok_name); πE != nil {
				continue
			}
			if πTemp005, πE = πg.GetAttr(πF, πTemp002, ßvalues, nil); πE != nil {
				continue
			}
			if πTemp002, πE = πTemp005.Call(πF, nil, nil); πE != nil {
				continue
			}
			πTemp001[0] = πTemp002
			if πTemp002, πE = πg.ResolveGlobal(πF, ß__all__); πE != nil {
				continue
			}
			if πTemp005, πE = πg.GetAttr(πF, πTemp002, ßextend, nil); πE != nil {
				continue
			}
			if πTemp002, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			// line 79: def ISTERMINAL(x):
			πF.SetLineno(79)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "x", Def: nil}
			πTemp002 = πg.NewFunction(πg.NewCode("ISTERMINAL", "/usr/lib/python2.7/token.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πArgs[0]; _ = µx
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
					// line 80: return x < NT_OFFSET
					πF.SetLineno(80)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNT_OFFSET); πE != nil {
						continue
					}
					if πTemp001, πE = πg.LT(πF, µx, πTemp002); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßISTERMINAL.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 82: def ISNONTERMINAL(x):
			πF.SetLineno(82)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "x", Def: nil}
			πTemp005 = πg.NewFunction(πg.NewCode("ISNONTERMINAL", "/usr/lib/python2.7/token.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πArgs[0]; _ = µx
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
					// line 83: return x >= NT_OFFSET
					πF.SetLineno(83)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNT_OFFSET); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GE(πF, µx, πTemp002); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßISNONTERMINAL.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 85: def ISEOF(x):
			πF.SetLineno(85)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "x", Def: nil}
			πTemp006 = πg.NewFunction(πg.NewCode("ISEOF", "/usr/lib/python2.7/token.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πArgs[0]; _ = µx
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
					// line 86: return x == ENDMARKER
					πF.SetLineno(86)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßENDMARKER); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, µx, πTemp002); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßISEOF.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 89: def __main():
			πF.SetLineno(89)
			πTemp004 = make([]πg.Param, 0)
			πTemp007 = πg.NewFunction(πg.NewCode("__main", "/usr/lib/python2.7/token.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µre *πg.Object = πg.UnboundLocal; _ = µre
				var µsys *πg.Object = πg.UnboundLocal; _ = µsys
				var µargs *πg.Object = πg.UnboundLocal; _ = µargs
				var µinFileName *πg.Object = πg.UnboundLocal; _ = µinFileName
				var µoutFileName *πg.Object = πg.UnboundLocal; _ = µoutFileName
				var µfp *πg.Object = πg.UnboundLocal; _ = µfp
				var µerr *πg.Object = πg.UnboundLocal; _ = µerr
				var µlines *πg.Object = πg.UnboundLocal; _ = µlines
				var µprog *πg.Object = πg.UnboundLocal; _ = µprog
				var µtokens *πg.Object = πg.UnboundLocal; _ = µtokens
				var µline *πg.Object = πg.UnboundLocal; _ = µline
				var µmatch *πg.Object = πg.UnboundLocal; _ = µmatch
				var µname *πg.Object = πg.UnboundLocal; _ = µname
				var µval *πg.Object = πg.UnboundLocal; _ = µval
				var µkeys *πg.Object = πg.UnboundLocal; _ = µkeys
				var µformat *πg.Object = πg.UnboundLocal; _ = µformat
				var µstart *πg.Object = πg.UnboundLocal; _ = µstart
				var µend *πg.Object = πg.UnboundLocal; _ = µend
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
				var πTemp008 *πg.BaseException
				_ = πTemp008
				var πTemp009 *πg.Traceback
				_ = πTemp009
				var πTemp010 []*πg.Object
				_ = πTemp010
				var πTemp011 *πg.Type
				_ = πTemp011
				var πTemp012 *πg.Dict
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
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 6: goto Label6
					case 8: goto Label8
					case 9: goto Label9
					case 10: goto Label10
					case 15: goto Label15
					case 17: goto Label17
					case 19: goto Label19
					case 21: goto Label21
					case 22: goto Label22
					case 25: goto Label25
					case 27: goto Label27
					default: panic("unexpected function state")
					}
					// line 90: import re
					πF.SetLineno(90)
					if πTemp002, πE = πg.ImportModule(πF, "re"); πE != nil {
						continue
					}
					πTemp001 = πTemp002[0]
					µre = πTemp001
					// line 91: import sys
					πF.SetLineno(91)
					if πTemp002, πE = πg.ImportModule(πF, "sys"); πE != nil {
						continue
					}
					πTemp001 = πTemp002[0]
					µsys = πTemp001
					// line 92: args = sys.argv[1:]
					πF.SetLineno(92)
					if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsys, "sys"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µsys, ßargv, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, πTemp004, πTemp001); πE != nil {
						continue
					}
					µargs = πTemp003
					// line 93: inFileName = args and args[0] or "Include/token.h"
					πF.SetLineno(93)
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					πTemp003 = µargs
					if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if !πTemp006 {
						goto Label2
					}
					πTemp004 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetItem(πF, µargs, πTemp004); πE != nil {
						continue
					}
					πTemp003 = πTemp007
				Label2:
					πTemp001 = πTemp003
					if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label1
					}
					πTemp001 = πg.NewStr("Include/token.h").ToObject()
				Label1:
					µinFileName = πTemp001
					// line 94: outFileName = "Lib/token.py"
					πF.SetLineno(94)
					µoutFileName = πg.NewStr("Lib/token.py").ToObject()
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					πTemp002[0] = µargs
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
						goto Label3
					}
					goto Label4
					// line 95: if len(args) > 1:
					πF.SetLineno(95)
				Label3:
					// line 96: outFileName = args[1]
					πF.SetLineno(96)
					πTemp001 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µargs, πTemp001); πE != nil {
						continue
					}
					µoutFileName = πTemp003
					goto Label4
				Label4:
					// line 97: try:
					πF.SetLineno(97)
					πF.PushCheckpoint(6)
					// line 98: fp = open(inFileName)
					πF.SetLineno(98)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µinFileName, "inFileName"); πE != nil {
						continue
					}
					πTemp002[0] = µinFileName
					if πTemp001, πE = πg.ResolveGlobal(πF, ßopen); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µfp = πTemp003
					πF.PopCheckpoint()
					goto Label5
				Label6:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp008, πTemp009 = πF.ExcInfo()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßOSError); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsInstance(πF, πTemp008.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label7
					}
					πE = πF.Raise(πTemp008.ToObject(), nil, πTemp009.ToObject())
					continue
					// line 99: except OSError as err:
					πF.SetLineno(99)
				Label7:
					µerr = πTemp008.ToObject()
					// line 100: sys.stdout.write("I/O error: %s\n" % str(err))
					πF.SetLineno(100)
					πTemp002 = πF.MakeArgs(1)
					πTemp010 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
						continue
					}
					πTemp010[0] = µerr
					if πTemp003, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp010, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp010)
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("I/O error: %s\n").ToObject(), πTemp004); πE != nil {
						continue
					}
					πTemp002[0] = πTemp001
					if πE = πg.CheckLocal(πF, µsys, "sys"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µsys, ßstdout, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßwrite, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					// line 101: sys.exit(1)
					πF.SetLineno(101)
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µsys, "sys"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µsys, ßexit, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πF.RestoreExc(nil, nil)
					goto Label5
				Label5:
					// line 102: with fp:
					πF.SetLineno(102)
					if πE = πg.CheckLocal(πF, µfp, "fp"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µfp.Type().ToObject(), ß__exit__, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µfp.Type().ToObject(), ß__enter__, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp003.Call(πF, πg.Args{µfp}, nil); πE != nil {
						continue
					}
					πF.PushCheckpoint(8)
					// line 103: lines = fp.read().split("\n")
					πF.SetLineno(103)
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewStr("\n").ToObject()
					if πE = πg.CheckLocal(πF, µfp, "fp"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µfp, ßread, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp007, ßsplit, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µlines = πTemp007
					πF.PopCheckpoint()
				Label8:
					πTemp008, πTemp009 = nil, nil
					if πE != nil {
						πTemp008, πTemp009 = πF.ExcInfo()
					}
					if πTemp008 != nil {
						πTemp011 = πTemp008.Type()
						if πTemp004, πE = πTemp001.Call(πF, πg.Args{µfp, πTemp011.ToObject(), πTemp008.ToObject(), πTemp009.ToObject()}, nil); πE != nil {
							continue
						}
					} else {
						if πTemp004, πE = πTemp001.Call(πF, πg.Args{µfp, πg.None, πg.None, πg.None}, nil); πE != nil {
							continue
						}
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp008 != nil && πTemp005 != true {
						πE = πF.Raise(nil, nil, nil)
						continue
					}
					if πR != nil {
						continue
					}
					// line 104: prog = re.compile(
					πF.SetLineno(104)
					πTemp002 = πF.MakeArgs(2)
					πTemp002[0] = πg.NewStr("#define[ \t][ \t]*([A-Z0-9][A-Z0-9_]*)[ \t][ \t]*([0-9][0-9]*)").ToObject()
					if πE = πg.CheckLocal(πF, µre, "re"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µre, ßIGNORECASE, nil); πE != nil {
						continue
					}
					πTemp002[1] = πTemp004
					if πE = πg.CheckLocal(πF, µre, "re"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µre, ßcompile, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µprog = πTemp007
					// line 107: tokens = {}
					πF.SetLineno(107)
					πTemp012 = πg.NewDict()
					πTemp004 = πTemp012.ToObject()
					µtokens = πTemp004
					if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Iter(πF, µlines); πE != nil {
						continue
					}
					πF.PushCheckpoint(10)
					πTemp005 = false
				Label9:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp005 {
						πF.PopCheckpoint()
						goto Label11
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
						µline = πTemp007
					}
					if πE != nil || !πTemp006 {
						continue
					}
					πF.PushCheckpoint(9)            
					// line 109: match = prog.match(line)
					πF.SetLineno(109)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					πTemp002[0] = µline
					if πE = πg.CheckLocal(πF, µprog, "prog"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, µprog, ßmatch, nil); πE != nil {
						continue
					}
					if πTemp013, πE = πTemp007.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µmatch = πTemp013
					if πE = πg.CheckLocal(πF, µmatch, "match"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, µmatch); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label12
					}
					goto Label13
					// line 110: if match:
					πF.SetLineno(110)
				Label12:
					// line 111: name, val = match.group(1, 2)
					πF.SetLineno(111)
					πTemp002 = πF.MakeArgs(2)
					πTemp002[0] = πg.NewInt(1).ToObject()
					πTemp002[1] = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µmatch, "match"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, µmatch, ßgroup, nil); πE != nil {
						continue
					}
					if πTemp013, πE = πTemp007.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp007}, πg.TieTarget{Target: &πTemp014}}}, πTemp013); πE != nil {
						continue
					}
					µname = πTemp007
					µval = πTemp014
					// line 112: val = int(val)
					πF.SetLineno(112)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µval, "val"); πE != nil {
						continue
					}
					πTemp002[0] = µval
					if πTemp007, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
						continue
					}
					if πTemp013, πE = πTemp007.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µval = πTemp013
					// line 113: tokens[val] = name          # reverse so we can sort them...
					πF.SetLineno(113)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp007}, µname); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µtokens, "tokens"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µval, "val"); πE != nil {
						continue
					}
					πTemp013 = µval
					if πE = πg.SetItem(πF, µtokens, πTemp013, πTemp007); πE != nil {
						continue
					}
					goto Label13
				Label13:
					continue
				Label10:
					if πE != nil || πR != nil {
						continue
					}
				Label11:
					// line 114: keys = sorted(tokens.keys())
					πF.SetLineno(114)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µtokens, "tokens"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µtokens, ßkeys, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp002[0] = πTemp007
					if πTemp004, πE = πg.ResolveGlobal(πF, ßsorted); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µkeys = πTemp007
					// line 116: try:
					πF.SetLineno(116)
					πF.PushCheckpoint(15)
					// line 117: fp = open(outFileName)
					πF.SetLineno(117)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µoutFileName, "outFileName"); πE != nil {
						continue
					}
					πTemp002[0] = µoutFileName
					if πTemp004, πE = πg.ResolveGlobal(πF, ßopen); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µfp = πTemp007
					πF.PopCheckpoint()
					goto Label14
				Label15:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp008, πTemp009 = πF.ExcInfo()
					if πTemp004, πE = πg.ResolveGlobal(πF, ßOSError); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsInstance(πF, πTemp008.ToObject(), πTemp004); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label16
					}
					πE = πF.Raise(πTemp008.ToObject(), nil, πTemp009.ToObject())
					continue
					// line 118: except OSError as err:
					πF.SetLineno(118)
				Label16:
					µerr = πTemp008.ToObject()
					// line 119: sys.stderr.write("I/O error: %s\n" % str(err))
					πF.SetLineno(119)
					πTemp002 = πF.MakeArgs(1)
					πTemp010 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
						continue
					}
					πTemp010[0] = µerr
					if πTemp007, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
						continue
					}
					if πTemp013, πE = πTemp007.Call(πF, πTemp010, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp010)
					if πTemp004, πE = πg.Mod(πF, πg.NewStr("I/O error: %s\n").ToObject(), πTemp013); πE != nil {
						continue
					}
					πTemp002[0] = πTemp004
					if πE = πg.CheckLocal(πF, µsys, "sys"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µsys, ßstderr, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, πTemp004, ßwrite, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp007.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					// line 120: sys.exit(2)
					πF.SetLineno(120)
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µsys, "sys"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µsys, ßexit, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πF.RestoreExc(nil, nil)
					goto Label14
				Label14:
					// line 121: with fp:
					πF.SetLineno(121)
					if πE = πg.CheckLocal(πF, µfp, "fp"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µfp.Type().ToObject(), ß__exit__, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, µfp.Type().ToObject(), ß__enter__, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp007.Call(πF, πg.Args{µfp}, nil); πE != nil {
						continue
					}
					πF.PushCheckpoint(17)
					// line 122: format = fp.read().split("\n")
					πF.SetLineno(122)
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewStr("\n").ToObject()
					if πE = πg.CheckLocal(πF, µfp, "fp"); πE != nil {
						continue
					}
					if πTemp013, πE = πg.GetAttr(πF, µfp, ßread, nil); πE != nil {
						continue
					}
					if πTemp014, πE = πTemp013.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp013, πE = πg.GetAttr(πF, πTemp014, ßsplit, nil); πE != nil {
						continue
					}
					if πTemp014, πE = πTemp013.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µformat = πTemp014
					πF.PopCheckpoint()
				Label17:
					πTemp008, πTemp009 = nil, nil
					if πE != nil {
						πTemp008, πTemp009 = πF.ExcInfo()
					}
					if πTemp008 != nil {
						πTemp011 = πTemp008.Type()
						if πTemp013, πE = πTemp004.Call(πF, πg.Args{µfp, πTemp011.ToObject(), πTemp008.ToObject(), πTemp009.ToObject()}, nil); πE != nil {
							continue
						}
					} else {
						if πTemp013, πE = πTemp004.Call(πF, πg.Args{µfp, πg.None, πg.None, πg.None}, nil); πE != nil {
							continue
						}
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp013); πE != nil {
						continue
					}
					if πTemp008 != nil && πTemp005 != true {
						πE = πF.Raise(nil, nil, nil)
						continue
					}
					if πR != nil {
						continue
					}
					// line 123: try:
					πF.SetLineno(123)
					πF.PushCheckpoint(19)
					// line 124: start = format.index("#--start constants--") + 1
					πF.SetLineno(124)
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewStr("#--start constants--").ToObject()
					if πE = πg.CheckLocal(πF, µformat, "format"); πE != nil {
						continue
					}
					if πTemp014, πE = πg.GetAttr(πF, µformat, ßindex, nil); πE != nil {
						continue
					}
					if πTemp015, πE = πTemp014.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp013, πE = πg.Add(πF, πTemp015, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µstart = πTemp013
					// line 125: end = format.index("#--end constants--")
					πF.SetLineno(125)
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewStr("#--end constants--").ToObject()
					if πE = πg.CheckLocal(πF, µformat, "format"); πE != nil {
						continue
					}
					if πTemp013, πE = πg.GetAttr(πF, µformat, ßindex, nil); πE != nil {
						continue
					}
					if πTemp014, πE = πTemp013.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µend = πTemp014
					πF.PopCheckpoint()
					goto Label18
				Label19:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp008, πTemp009 = πF.ExcInfo()
					if πTemp013, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsInstance(πF, πTemp008.ToObject(), πTemp013); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label20
					}
					πE = πF.Raise(πTemp008.ToObject(), nil, πTemp009.ToObject())
					continue
					// line 126: except ValueError:
					πF.SetLineno(126)
				Label20:
					// line 127: sys.stderr.write("target does not contain format markers")
					πF.SetLineno(127)
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewStr("target does not contain format markers").ToObject()
					if πE = πg.CheckLocal(πF, µsys, "sys"); πE != nil {
						continue
					}
					if πTemp013, πE = πg.GetAttr(πF, µsys, ßstderr, nil); πE != nil {
						continue
					}
					if πTemp014, πE = πg.GetAttr(πF, πTemp013, ßwrite, nil); πE != nil {
						continue
					}
					if πTemp013, πE = πTemp014.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					// line 128: sys.exit(3)
					πF.SetLineno(128)
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewInt(3).ToObject()
					if πE = πg.CheckLocal(πF, µsys, "sys"); πE != nil {
						continue
					}
					if πTemp013, πE = πg.GetAttr(πF, µsys, ßexit, nil); πE != nil {
						continue
					}
					if πTemp014, πE = πTemp013.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πF.RestoreExc(nil, nil)
					goto Label18
				Label18:
					// line 129: lines = []
					πF.SetLineno(129)
					πTemp002 = make([]*πg.Object, 0)
					πTemp013 = πg.NewList(πTemp002...).ToObject()
					µlines = πTemp013
					if πE = πg.CheckLocal(πF, µkeys, "keys"); πE != nil {
						continue
					}
					if πTemp013, πE = πg.Iter(πF, µkeys); πE != nil {
						continue
					}
					πF.PushCheckpoint(22)
					πTemp005 = false
				Label21:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp005 {
						πF.PopCheckpoint()
						goto Label23
					}
					if πTemp014, πE = πg.Next(πF, πTemp013); πE != nil {
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
						µval = πTemp014
					}
					if πE != nil || !πTemp006 {
						continue
					}
					πF.PushCheckpoint(21)            
					// line 131: lines.append("%s = %d" % (tokens[val], val))
					πF.SetLineno(131)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µval, "val"); πE != nil {
						continue
					}
					πTemp016 = µval
					if πE = πg.CheckLocal(πF, µtokens, "tokens"); πE != nil {
						continue
					}
					if πTemp017, πE = πg.GetItem(πF, µtokens, πTemp016); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µval, "val"); πE != nil {
						continue
					}
					πTemp015 = πg.NewTuple2(πTemp017, µval).ToObject()
					if πTemp014, πE = πg.Mod(πF, πg.NewStr("%s = %d").ToObject(), πTemp015); πE != nil {
						continue
					}
					πTemp002[0] = πTemp014
					if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
						continue
					}
					if πTemp014, πE = πg.GetAttr(πF, µlines, ßappend, nil); πE != nil {
						continue
					}
					if πTemp015, πE = πTemp014.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					continue
				Label22:
					if πE != nil || πR != nil {
						continue
					}
				Label23:
					// line 132: format[start:end] = lines
					πF.SetLineno(132)
					if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp013}, µlines); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µformat, "format"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µstart, "start"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
						continue
					}
					if πTemp014, πE = πg.SliceType.Call(πF, πg.Args{µstart, µend, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.SetItem(πF, µformat, πTemp014, πTemp013); πE != nil {
						continue
					}
					// line 133: try:
					πF.SetLineno(133)
					πF.PushCheckpoint(25)
					// line 134: fp = open(outFileName, 'w')
					πF.SetLineno(134)
					πTemp002 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µoutFileName, "outFileName"); πE != nil {
						continue
					}
					πTemp002[0] = µoutFileName
					πTemp002[1] = ßw.ToObject()
					if πTemp013, πE = πg.ResolveGlobal(πF, ßopen); πE != nil {
						continue
					}
					if πTemp014, πE = πTemp013.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µfp = πTemp014
					πF.PopCheckpoint()
					goto Label24
				Label25:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp008, πTemp009 = πF.ExcInfo()
					if πTemp013, πE = πg.ResolveGlobal(πF, ßOSError); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsInstance(πF, πTemp008.ToObject(), πTemp013); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label26
					}
					πE = πF.Raise(πTemp008.ToObject(), nil, πTemp009.ToObject())
					continue
					// line 135: except OSError as err:
					πF.SetLineno(135)
				Label26:
					µerr = πTemp008.ToObject()
					// line 136: sys.stderr.write("I/O error: %s\n" % str(err))
					πF.SetLineno(136)
					πTemp002 = πF.MakeArgs(1)
					πTemp010 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
						continue
					}
					πTemp010[0] = µerr
					if πTemp014, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
						continue
					}
					if πTemp015, πE = πTemp014.Call(πF, πTemp010, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp010)
					if πTemp013, πE = πg.Mod(πF, πg.NewStr("I/O error: %s\n").ToObject(), πTemp015); πE != nil {
						continue
					}
					πTemp002[0] = πTemp013
					if πE = πg.CheckLocal(πF, µsys, "sys"); πE != nil {
						continue
					}
					if πTemp013, πE = πg.GetAttr(πF, µsys, ßstderr, nil); πE != nil {
						continue
					}
					if πTemp014, πE = πg.GetAttr(πF, πTemp013, ßwrite, nil); πE != nil {
						continue
					}
					if πTemp013, πE = πTemp014.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					// line 137: sys.exit(4)
					πF.SetLineno(137)
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewInt(4).ToObject()
					if πE = πg.CheckLocal(πF, µsys, "sys"); πE != nil {
						continue
					}
					if πTemp013, πE = πg.GetAttr(πF, µsys, ßexit, nil); πE != nil {
						continue
					}
					if πTemp014, πE = πTemp013.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πF.RestoreExc(nil, nil)
					goto Label24
				Label24:
					// line 138: with fp:
					πF.SetLineno(138)
					if πE = πg.CheckLocal(πF, µfp, "fp"); πE != nil {
						continue
					}
					if πTemp013, πE = πg.GetAttr(πF, µfp.Type().ToObject(), ß__exit__, nil); πE != nil {
						continue
					}
					if πTemp014, πE = πg.GetAttr(πF, µfp.Type().ToObject(), ß__enter__, nil); πE != nil {
						continue
					}
					if πTemp014, πE = πTemp014.Call(πF, πg.Args{µfp}, nil); πE != nil {
						continue
					}
					πF.PushCheckpoint(27)
					// line 139: fp.write("\n".join(format))
					πF.SetLineno(139)
					πTemp002 = πF.MakeArgs(1)
					πTemp010 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µformat, "format"); πE != nil {
						continue
					}
					πTemp010[0] = µformat
					if πTemp015, πE = πg.GetAttr(πF, πg.NewStr("\n").ToObject(), ßjoin, nil); πE != nil {
						continue
					}
					if πTemp016, πE = πTemp015.Call(πF, πTemp010, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp010)
					πTemp002[0] = πTemp016
					if πE = πg.CheckLocal(πF, µfp, "fp"); πE != nil {
						continue
					}
					if πTemp015, πE = πg.GetAttr(πF, µfp, ßwrite, nil); πE != nil {
						continue
					}
					if πTemp016, πE = πTemp015.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πF.PopCheckpoint()
				Label27:
					πTemp008, πTemp009 = nil, nil
					if πE != nil {
						πTemp008, πTemp009 = πF.ExcInfo()
					}
					if πTemp008 != nil {
						πTemp011 = πTemp008.Type()
						if πTemp015, πE = πTemp013.Call(πF, πg.Args{µfp, πTemp011.ToObject(), πTemp008.ToObject(), πTemp009.ToObject()}, nil); πE != nil {
							continue
						}
					} else {
						if πTemp015, πE = πTemp013.Call(πF, πg.Args{µfp, πg.None, πg.None, πg.None}, nil); πE != nil {
							continue
						}
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp015); πE != nil {
						continue
					}
					if πTemp008 != nil && πTemp005 != true {
						πE = πF.Raise(nil, nil, nil)
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
			if πE = πF.Globals().SetItem(πF, ß__main.ToObject(), πTemp007); πE != nil {
				continue
			}
			if πTemp009, πE = πg.ResolveGlobal(πF, ß__name__); πE != nil {
				continue
			}
			if πTemp008, πE = πg.Eq(πF, πTemp009, ß__main__.ToObject()); πE != nil {
				continue
			}
			if πTemp010, πE = πg.IsTrue(πF, πTemp008); πE != nil {
				continue
			}
			if πTemp010 {
				goto Label1
			}
			goto Label2
			// line 142: if __name__ == "__main__":
			πF.SetLineno(142)
		Label1:
			// line 143: __main()
			πF.SetLineno(143)
			if πTemp008, πE = πg.ResolveGlobal(πF, ß__main); πE != nil {
				continue
			}
			if πTemp009, πE = πTemp008.Call(πF, nil, nil); πE != nil {
				continue
			}
			goto Label2
		Label2:
		}
		return nil, πE
	})
	πg.RegisterModule("token", Code)
}

