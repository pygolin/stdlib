package datetime
import (
	πg "github.com/pygolin/runtime"
	// _ "github.com/pygolin/stdlib/_struct"
	// _ "github.com/pygolin/stdlib/math"
	// _ "github.com/pygolin/stdlib/time"
)
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/datetime.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ß := πg.InternStr("")
		ß10B := πg.InternStr("10B")
		ß4B := πg.InternStr("4B")
		ß6B := πg.InternStr("6B")
		ßApr := πg.InternStr("Apr")
		ßAttributeError := πg.InternStr("AttributeError")
		ßAug := πg.InternStr("Aug")
		ßDec := πg.InternStr("Dec")
		ßFalse := πg.InternStr("False")
		ßFeb := πg.InternStr("Feb")
		ßFri := πg.InternStr("Fri")
		ßJan := πg.InternStr("Jan")
		ßJul := πg.InternStr("Jul")
		ßJun := πg.InternStr("Jun")
		ßMAXYEAR := πg.InternStr("MAXYEAR")
		ßMINYEAR := πg.InternStr("MINYEAR")
		ßMar := πg.InternStr("Mar")
		ßMay := πg.InternStr("May")
		ßMon := πg.InternStr("Mon")
		ßNone := πg.InternStr("None")
		ßNotImplemented := πg.InternStr("NotImplemented")
		ßNotImplementedError := πg.InternStr("NotImplementedError")
		ßNov := πg.InternStr("Nov")
		ßOct := πg.InternStr("Oct")
		ßOverflowError := πg.InternStr("OverflowError")
		ßSat := πg.InternStr("Sat")
		ßSep := πg.InternStr("Sep")
		ßSun := πg.InternStr("Sun")
		ßT := πg.InternStr("T")
		ßThu := πg.InternStr("Thu")
		ßTrue := πg.InternStr("True")
		ßTue := πg.InternStr("Tue")
		ßTypeError := πg.InternStr("TypeError")
		ßValueError := πg.InternStr("ValueError")
		ßWed := πg.InternStr("Wed")
		ß_DAYNAMES := πg.InternStr("_DAYNAMES")
		ß_DAYS_BEFORE_MONTH := πg.InternStr("_DAYS_BEFORE_MONTH")
		ß_DAYS_IN_MONTH := πg.InternStr("_DAYS_IN_MONTH")
		ß_DI100Y := πg.InternStr("_DI100Y")
		ß_DI400Y := πg.InternStr("_DI400Y")
		ß_DI4Y := πg.InternStr("_DI4Y")
		ß_MAX_DELTA_DAYS := πg.InternStr("_MAX_DELTA_DAYS")
		ß_MINYEARFMT := πg.InternStr("_MINYEARFMT")
		ß_MONTHNAMES := πg.InternStr("_MONTHNAMES")
		ß_SECONDS_PER_DAY := πg.InternStr("_SECONDS_PER_DAY")
		ß_SENTINEL := πg.InternStr("_SENTINEL")
		ß_US_PER_DAY := πg.InternStr("_US_PER_DAY")
		ß_US_PER_HOUR := πg.InternStr("_US_PER_HOUR")
		ß_US_PER_MINUTE := πg.InternStr("_US_PER_MINUTE")
		ß_US_PER_MS := πg.InternStr("_US_PER_MS")
		ß_US_PER_SECOND := πg.InternStr("_US_PER_SECOND")
		ß_US_PER_US := πg.InternStr("_US_PER_US")
		ß_US_PER_WEEK := πg.InternStr("_US_PER_WEEK")
		ß__abs__ := πg.InternStr("__abs__")
		ß__add__ := πg.InternStr("__add__")
		ß__class__ := πg.InternStr("__class__")
		ß__dict__ := πg.InternStr("__dict__")
		ß__div__ := πg.InternStr("__div__")
		ß__doc__ := πg.InternStr("__doc__")
		ß__eq__ := πg.InternStr("__eq__")
		ß__floordiv__ := πg.InternStr("__floordiv__")
		ß__format__ := πg.InternStr("__format__")
		ß__ge__ := πg.InternStr("__ge__")
		ß__getinitargs__ := πg.InternStr("__getinitargs__")
		ß__getstate__ := πg.InternStr("__getstate__")
		ß__gt__ := πg.InternStr("__gt__")
		ß__hash__ := πg.InternStr("__hash__")
		ß__int__ := πg.InternStr("__int__")
		ß__le__ := πg.InternStr("__le__")
		ß__lt__ := πg.InternStr("__lt__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__mul__ := πg.InternStr("__mul__")
		ß__name__ := πg.InternStr("__name__")
		ß__ne__ := πg.InternStr("__ne__")
		ß__neg__ := πg.InternStr("__neg__")
		ß__new__ := πg.InternStr("__new__")
		ß__nonzero__ := πg.InternStr("__nonzero__")
		ß__pos__ := πg.InternStr("__pos__")
		ß__radd__ := πg.InternStr("__radd__")
		ß__reduce__ := πg.InternStr("__reduce__")
		ß__repr__ := πg.InternStr("__repr__")
		ß__rmul__ := πg.InternStr("__rmul__")
		ß__setstate := πg.InternStr("__setstate")
		ß__slots__ := πg.InternStr("__slots__")
		ß__str__ := πg.InternStr("__str__")
		ß__sub__ := πg.InternStr("__sub__")
		ß_accum := πg.InternStr("_accum")
		ß_add_timedelta := πg.InternStr("_add_timedelta")
		ß_build_struct_time := πg.InternStr("_build_struct_time")
		ß_check_date_fields := πg.InternStr("_check_date_fields")
		ß_check_int_field := πg.InternStr("_check_int_field")
		ß_check_time_fields := πg.InternStr("_check_time_fields")
		ß_check_tzinfo_arg := πg.InternStr("_check_tzinfo_arg")
		ß_check_tzname := πg.InternStr("_check_tzname")
		ß_check_utc_offset := πg.InternStr("_check_utc_offset")
		ß_cmp := πg.InternStr("_cmp")
		ß_cmperror := πg.InternStr("_cmperror")
		ß_create := πg.InternStr("_create")
		ß_date_class := πg.InternStr("_date_class")
		ß_day := πg.InternStr("_day")
		ß_days := πg.InternStr("_days")
		ß_days_before_month := πg.InternStr("_days_before_month")
		ß_days_before_year := πg.InternStr("_days_before_year")
		ß_days_in_month := πg.InternStr("_days_in_month")
		ß_dst := πg.InternStr("_dst")
		ß_format_time := πg.InternStr("_format_time")
		ß_from_microseconds := πg.InternStr("_from_microseconds")
		ß_from_timestamp := πg.InternStr("_from_timestamp")
		ß_getstate := πg.InternStr("_getstate")
		ß_hashcode := πg.InternStr("_hashcode")
		ß_hour := πg.InternStr("_hour")
		ß_is_leap := πg.InternStr("_is_leap")
		ß_isoweek1monday := πg.InternStr("_isoweek1monday")
		ß_math := πg.InternStr("_math")
		ß_microsecond := πg.InternStr("_microsecond")
		ß_microseconds := πg.InternStr("_microseconds")
		ß_minute := πg.InternStr("_minute")
		ß_month := πg.InternStr("_month")
		ß_normalize_date := πg.InternStr("_normalize_date")
		ß_normalize_datetime := πg.InternStr("_normalize_datetime")
		ß_normalize_pair := πg.InternStr("_normalize_pair")
		ß_ord2ymd := πg.InternStr("_ord2ymd")
		ß_round := πg.InternStr("_round")
		ß_second := πg.InternStr("_second")
		ß_seconds := πg.InternStr("_seconds")
		ß_struct := πg.InternStr("_struct")
		ß_time := πg.InternStr("_time")
		ß_time_class := πg.InternStr("_time_class")
		ß_to_microseconds := πg.InternStr("_to_microseconds")
		ß_tzinfo := πg.InternStr("_tzinfo")
		ß_tzinfo_class := πg.InternStr("_tzinfo_class")
		ß_tzstr := πg.InternStr("_tzstr")
		ß_utcoffset := πg.InternStr("_utcoffset")
		ß_year := πg.InternStr("_year")
		ß_ymd2ord := πg.InternStr("_ymd2ord")
		ßabs := πg.InternStr("abs")
		ßappend := πg.InternStr("append")
		ßastimezone := πg.InternStr("astimezone")
		ßceil := πg.InternStr("ceil")
		ßclassmethod := πg.InternStr("classmethod")
		ßcombine := πg.InternStr("combine")
		ßctime := πg.InternStr("ctime")
		ßdate := πg.InternStr("date")
		ßdatetime := πg.InternStr("datetime")
		ßday := πg.InternStr("day")
		ßdays := πg.InternStr("days")
		ßdbm := πg.InternStr("dbm")
		ßdim := πg.InternStr("dim")
		ßdivmod := πg.InternStr("divmod")
		ßdst := πg.InternStr("dst")
		ßfloat := πg.InternStr("float")
		ßfloor := πg.InternStr("floor")
		ßfromordinal := πg.InternStr("fromordinal")
		ßfromtimestamp := πg.InternStr("fromtimestamp")
		ßfromutc := πg.InternStr("fromutc")
		ßgetattr := πg.InternStr("getattr")
		ßgmtime := πg.InternStr("gmtime")
		ßhasattr := πg.InternStr("hasattr")
		ßhash := πg.InternStr("hash")
		ßhour := πg.InternStr("hour")
		ßhours := πg.InternStr("hours")
		ßint := πg.InternStr("int")
		ßisinstance := πg.InternStr("isinstance")
		ßisocalendar := πg.InternStr("isocalendar")
		ßisoformat := πg.InternStr("isoformat")
		ßisoweekday := πg.InternStr("isoweekday")
		ßjoin := πg.InternStr("join")
		ßlen := πg.InternStr("len")
		ßlocaltime := πg.InternStr("localtime")
		ßlong := πg.InternStr("long")
		ßmap := πg.InternStr("map")
		ßmax := πg.InternStr("max")
		ßmicrosecond := πg.InternStr("microsecond")
		ßmicroseconds := πg.InternStr("microseconds")
		ßmilliseconds := πg.InternStr("milliseconds")
		ßmin := πg.InternStr("min")
		ßminute := πg.InternStr("minute")
		ßminutes := πg.InternStr("minutes")
		ßmodf := πg.InternStr("modf")
		ßmonth := πg.InternStr("month")
		ßnow := πg.InternStr("now")
		ßobject := πg.InternStr("object")
		ßord := πg.InternStr("ord")
		ßpack := πg.InternStr("pack")
		ßproperty := πg.InternStr("property")
		ßreplace := πg.InternStr("replace")
		ßresolution := πg.InternStr("resolution")
		ßs := πg.InternStr("s")
		ßsecond := πg.InternStr("second")
		ßseconds := πg.InternStr("seconds")
		ßstr := πg.InternStr("str")
		ßstrftime := πg.InternStr("strftime")
		ßstruct_time := πg.InternStr("struct_time")
		ßtime := πg.InternStr("time")
		ßtimedelta := πg.InternStr("timedelta")
		ßtimetuple := πg.InternStr("timetuple")
		ßtimetz := πg.InternStr("timetz")
		ßtoday := πg.InternStr("today")
		ßtoordinal := πg.InternStr("toordinal")
		ßtotal_seconds := πg.InternStr("total_seconds")
		ßtype := πg.InternStr("type")
		ßtzinfo := πg.InternStr("tzinfo")
		ßtzname := πg.InternStr("tzname")
		ßunicode := πg.InternStr("unicode")
		ßutcfromtimestamp := πg.InternStr("utcfromtimestamp")
		ßutcnow := πg.InternStr("utcnow")
		ßutcoffset := πg.InternStr("utcoffset")
		ßutctimetuple := πg.InternStr("utctimetuple")
		ßweekday := πg.InternStr("weekday")
		ßweeks := πg.InternStr("weeks")
		ßyear := πg.InternStr("year")
		ßzfill := πg.InternStr("zfill")
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
		var πTemp010 bool
		_ = πTemp010
		var πTemp011 bool
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
		var πTemp027 *πg.Dict
		_ = πTemp027
		var πTemp028 *πg.Object
		_ = πTemp028
		var πTemp029 *πg.Object
		_ = πTemp029
		var πTemp030 *πg.Object
		_ = πTemp030
		var πTemp031 πg.KWArgs
		_ = πTemp031
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			case 1: goto Label1
			case 2: goto Label2
			default: panic("unexpected function state")
			}
			// line 1: """Concrete date/time and related types -- prototype implemented in Python.
			πF.SetLineno(1)
			// line 1: """Concrete date/time and related types -- prototype implemented in Python.
			πF.SetLineno(1)
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Concrete date/time and related types -- prototype implemented in Python.\n\nSee http://www.zope.org/Members/fdrake/DateTimeWiki/FrontPage\n\nSee also http://dir.yahoo.com/Reference/calendars/\n\nFor a primer on DST, including many current DST rules, see\nhttp://webexhibits.org/daylightsaving/\n\nFor more about DST than you ever wanted to know, see\nftp://elsie.nci.nih.gov/pub/\n\nSources for time zone and DST data: http://www.twinsun.com/tz/tz-link.htm\n\nThis was originally copied from the sandbox of the CPython CVS repository.\nThanks to Tim Peters for suggesting using it.\n").ToObject()); πE != nil {
				continue
			}
			// line 20: import time as _time
			πF.SetLineno(20)
			if πTemp002, πE = πg.ImportModule(πF, "time"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ß_time.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 21: import math as _math
			πF.SetLineno(21)
			if πTemp002, πE = πg.ImportModule(πF, "math"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ß_math.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 23: import _struct
			πF.SetLineno(23)
			if πTemp002, πE = πg.ImportModule(πF, "_struct"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ß_struct.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 25: def divmod(x, y):
			πF.SetLineno(25)
			πTemp003 = make([]πg.Param, 2)
			πTemp003[0] = πg.Param{Name: "x", Def: nil}
			πTemp003[1] = πg.Param{Name: "y", Def: nil}
			πTemp001 = πg.NewFunction(πg.NewCode("divmod", "/usr/lib/python2.7/datetime.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πArgs[0]; _ = µx
				var µy *πg.Object = πArgs[1]; _ = µy
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
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 26: x, y = int(x), int(y)
					πF.SetLineno(26)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp002[0] = µx
					if πTemp003, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
						continue
					}
					πTemp002[0] = µy
					if πTemp003, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001 = πg.NewTuple2(πTemp004, πTemp005).ToObject()
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp004}}}, πTemp001); πE != nil {
						continue
					}
					µx = πTemp003
					µy = πTemp004
					// line 27: return x / y, x % y
					πF.SetLineno(27)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Div(πF, µx, µy); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Mod(πF, µx, µy); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
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
			if πE = πF.Globals().SetItem(πF, ßdivmod.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 29: _SENTINEL = object()
			πF.SetLineno(29)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_SENTINEL.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 31: def _cmp(x, y):
			πF.SetLineno(31)
			πTemp003 = make([]πg.Param, 2)
			πTemp003[0] = πg.Param{Name: "x", Def: nil}
			πTemp003[1] = πg.Param{Name: "y", Def: nil}
			πTemp004 = πg.NewFunction(πg.NewCode("_cmp", "/usr/lib/python2.7/datetime.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πArgs[0]; _ = µx
				var µy *πg.Object = πArgs[1]; _ = µy
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 bool
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
					// line 32: return 0 if x == y else 1 if x > y else -1
					πF.SetLineno(32)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µx, µy); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp003 {
						goto Label1
					}
					πTemp001 = πg.NewInt(0).ToObject()
					goto Label2
				Label1:
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GT(πF, µx, µy); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if !πTemp003 {
						goto Label3
					}
					πTemp002 = πg.NewInt(1).ToObject()
					goto Label4
				Label3:
					if πTemp004, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp002 = πTemp004
				Label4:
					πTemp001 = πTemp002
				Label2:
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
			if πE = πF.Globals().SetItem(πF, ß_cmp.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 34: def _round(x):
			πF.SetLineno(34)
			πTemp003 = make([]πg.Param, 1)
			πTemp003[0] = πg.Param{Name: "x", Def: nil}
			πTemp005 = πg.NewFunction(πg.NewCode("_round", "/usr/lib/python2.7/datetime.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πArgs[0]; _ = µx
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
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 35: return int(_math.floor(x + 0.5) if x >= 0.0 else _math.ceil(x - 0.5))
					πF.SetLineno(35)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GE(πF, µx, πg.NewFloat(0.0).ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if !πTemp004 {
						goto Label1
					}
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Add(πF, µx, πg.NewFloat(0.5).ToObject()); πE != nil {
						continue
					}
					πTemp005[0] = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_math); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp003, ßfloor, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp006.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πTemp002 = πTemp003
					goto Label2
				Label1:
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Sub(πF, µx, πg.NewFloat(0.5).ToObject()); πE != nil {
						continue
					}
					πTemp005[0] = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_math); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp003, ßceil, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp006.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πTemp002 = πTemp003
				Label2:
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ß_round.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 37: MINYEAR = 1
			πF.SetLineno(37)
			if πE = πF.Globals().SetItem(πF, ßMINYEAR.ToObject(), πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			// line 38: MAXYEAR = 9999
			πF.SetLineno(38)
			if πE = πF.Globals().SetItem(πF, ßMAXYEAR.ToObject(), πg.NewInt(9999).ToObject()); πE != nil {
				continue
			}
			// line 39: _MINYEARFMT = 1900
			πF.SetLineno(39)
			if πE = πF.Globals().SetItem(πF, ß_MINYEARFMT.ToObject(), πg.NewInt(1900).ToObject()); πE != nil {
				continue
			}
			// line 41: _MAX_DELTA_DAYS = 999999999
			πF.SetLineno(41)
			if πE = πF.Globals().SetItem(πF, ß_MAX_DELTA_DAYS.ToObject(), πg.NewInt(999999999).ToObject()); πE != nil {
				continue
			}
			// line 52: _DAYS_IN_MONTH = [-1, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31]
			πF.SetLineno(52)
			πTemp002 = make([]*πg.Object, 13)
			if πTemp006, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp002[0] = πTemp006
			πTemp002[1] = πg.NewInt(31).ToObject()
			πTemp002[2] = πg.NewInt(28).ToObject()
			πTemp002[3] = πg.NewInt(31).ToObject()
			πTemp002[4] = πg.NewInt(30).ToObject()
			πTemp002[5] = πg.NewInt(31).ToObject()
			πTemp002[6] = πg.NewInt(30).ToObject()
			πTemp002[7] = πg.NewInt(31).ToObject()
			πTemp002[8] = πg.NewInt(31).ToObject()
			πTemp002[9] = πg.NewInt(30).ToObject()
			πTemp002[10] = πg.NewInt(31).ToObject()
			πTemp002[11] = πg.NewInt(30).ToObject()
			πTemp002[12] = πg.NewInt(31).ToObject()
			πTemp006 = πg.NewList(πTemp002...).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_DAYS_IN_MONTH.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 54: _DAYS_BEFORE_MONTH = [-1]
			πF.SetLineno(54)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp006, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp002[0] = πTemp006
			πTemp006 = πg.NewList(πTemp002...).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_DAYS_BEFORE_MONTH.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 55: dbm = 0
			πF.SetLineno(55)
			if πE = πF.Globals().SetItem(πF, ßdbm.ToObject(), πg.NewInt(0).ToObject()); πE != nil {
				continue
			}
			if πTemp007, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
				continue
			}
			if πTemp009, πE = πg.ResolveGlobal(πF, ß_DAYS_IN_MONTH); πE != nil {
				continue
			}
			if πTemp008, πE = πg.GetItem(πF, πTemp009, πTemp007); πE != nil {
				continue
			}
			if πTemp006, πE = πg.Iter(πF, πTemp008); πE != nil {
				continue
			}
			πF.PushCheckpoint(2)
			πTemp010 = false
		Label1:
			if πE != nil || πR != nil {
				continue
			}
			if πTemp010 {
				πF.PopCheckpoint()
				goto Label3
			}
			if πTemp007, πE = πg.Next(πF, πTemp006); πE != nil {
				isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
				if exc != nil {
					πE = exc
				} else if isStop {
					πE = nil
					πF.RestoreExc(nil, nil)
				}
				πTemp011 = !isStop
			} else {
				πTemp011 = true
				if πE = πF.Globals().SetItem(πF, ßdim.ToObject(), πTemp007); πE != nil {
					continue
				}
			}
			if πE != nil || !πTemp011 {
				continue
			}
			πF.PushCheckpoint(1)            
			// line 57: _DAYS_BEFORE_MONTH.append(dbm)
			πF.SetLineno(57)
			πTemp002 = πF.MakeArgs(1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßdbm); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			if πTemp007, πE = πg.ResolveGlobal(πF, ß_DAYS_BEFORE_MONTH); πE != nil {
				continue
			}
			if πTemp008, πE = πg.GetAttr(πF, πTemp007, ßappend, nil); πE != nil {
				continue
			}
			if πTemp007, πE = πTemp008.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			// line 58: dbm += dim
			πF.SetLineno(58)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßdbm); πE != nil {
				continue
			}
			if πTemp008, πE = πg.ResolveGlobal(πF, ßdim); πE != nil {
				continue
			}
			if πTemp009, πE = πg.IAdd(πF, πTemp007, πTemp008); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßdbm.ToObject(), πTemp009); πE != nil {
				continue
			}
			continue
		Label2:
			if πE != nil || πR != nil {
				continue
			}
		Label3:
			// line 59: del dbm, dim
			πF.SetLineno(59)
			if πE = πg.DelVar(πF, πF.Globals(), ßdbm); πE != nil {
				continue
			}
			if πE = πg.DelVar(πF, πF.Globals(), ßdim); πE != nil {
				continue
			}
			// line 61: def _is_leap(year):
			πF.SetLineno(61)
			πTemp003 = make([]πg.Param, 1)
			πTemp003[0] = πg.Param{Name: "year", Def: nil}
			πTemp006 = πg.NewFunction(πg.NewCode("_is_leap", "/usr/lib/python2.7/datetime.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µyear *πg.Object = πArgs[0]; _ = µyear
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 bool
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
					// line 62: "year -> 1 if leap year, else 0."
					πF.SetLineno(62)
					// line 63: return year % 4 == 0 and (year % 100 != 0 or year % 400 == 0)
					πF.SetLineno(63)
					if πE = πg.CheckLocal(πF, µyear, "year"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Mod(πF, µyear, πg.NewInt(4).ToObject()); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Eq(πF, πTemp004, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					πTemp001 = πTemp003
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp002 {
						goto Label1
					}
					if πE = πg.CheckLocal(πF, µyear, "year"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.Mod(πF, µyear, πg.NewInt(100).ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.NE(πF, πTemp006, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					πTemp003 = πTemp004
					if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label2
					}
					if πE = πg.CheckLocal(πF, µyear, "year"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.Mod(πF, µyear, πg.NewInt(400).ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Eq(πF, πTemp006, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					πTemp003 = πTemp004
				Label2:
					πTemp001 = πTemp003
				Label1:
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
			if πE = πF.Globals().SetItem(πF, ß_is_leap.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 62: "year -> 1 if leap year, else 0."
			πF.SetLineno(62)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp007}, πg.NewStr("year -> 1 if leap year, else 0.").ToObject()); πE != nil {
				continue
			}
			if πTemp008, πE = πg.ResolveGlobal(πF, ß_is_leap); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp008, ß__doc__, πTemp007); πE != nil {
				continue
			}
			// line 65: def _days_before_year(year):
			πF.SetLineno(65)
			πTemp003 = make([]πg.Param, 1)
			πTemp003[0] = πg.Param{Name: "year", Def: nil}
			πTemp007 = πg.NewFunction(πg.NewCode("_days_before_year", "/usr/lib/python2.7/datetime.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µyear *πg.Object = πArgs[0]; _ = µyear
				var µy *πg.Object = πg.UnboundLocal; _ = µy
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
					// line 66: "year -> number of days before January 1st of year."
					πF.SetLineno(66)
					// line 67: y = year - 1
					πF.SetLineno(67)
					if πE = πg.CheckLocal(πF, µyear, "year"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Sub(πF, µyear, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µy = πTemp001
					// line 68: return y*365 + y//4 - y//100 + y//400
					πF.SetLineno(68)
					if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Mul(πF, µy, πg.NewInt(365).ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.FloorDiv(πF, µy, πg.NewInt(4).ToObject()); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Add(πF, πTemp004, πTemp005); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.FloorDiv(πF, µy, πg.NewInt(100).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Sub(πF, πTemp003, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.FloorDiv(πF, µy, πg.NewInt(400).ToObject()); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, πTemp002, πTemp003); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ß_days_before_year.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 66: "year -> number of days before January 1st of year."
			πF.SetLineno(66)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp008}, πg.NewStr("year -> number of days before January 1st of year.").ToObject()); πE != nil {
				continue
			}
			if πTemp009, πE = πg.ResolveGlobal(πF, ß_days_before_year); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp009, ß__doc__, πTemp008); πE != nil {
				continue
			}
			// line 70: def _days_in_month(year, month):
			πF.SetLineno(70)
			πTemp003 = make([]πg.Param, 2)
			πTemp003[0] = πg.Param{Name: "year", Def: nil}
			πTemp003[1] = πg.Param{Name: "month", Def: nil}
			πTemp008 = πg.NewFunction(πg.NewCode("_days_in_month", "/usr/lib/python2.7/datetime.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µyear *πg.Object = πArgs[0]; _ = µyear
				var µmonth *πg.Object = πArgs[1]; _ = µmonth
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
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 71: "year, month -> number of days in that month in that year."
					πF.SetLineno(71)
					// line 72: assert 1 <= month <= 12, month
					πF.SetLineno(72)
					if πE = πg.CheckLocal(πF, µmonth, "month"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µmonth, "month"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.LE(πF, πg.NewInt(1).ToObject(), µmonth); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp002 {
						goto Label1
					}
					if πTemp001, πE = πg.LE(πF, µmonth, πg.NewInt(12).ToObject()); πE != nil {
						continue
					}
				Label1:
					if πE = πg.Assert(πF, πTemp001, µmonth); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µmonth, "month"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Eq(πF, µmonth, πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
					πTemp001 = πTemp003
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp002 {
						goto Label2
					}
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µyear, "year"); πE != nil {
						continue
					}
					πTemp004[0] = µyear
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_is_leap); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					πTemp001 = πTemp005
				Label2:
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label3
					}
					goto Label4
					// line 73: if month == 2 and _is_leap(year):
					πF.SetLineno(73)
				Label3:
					// line 74: return 29
					πF.SetLineno(74)
					πR = πg.NewInt(29).ToObject()
					continue
					goto Label4
				Label4:
					// line 75: return _DAYS_IN_MONTH[month]
					πF.SetLineno(75)
					if πE = πg.CheckLocal(πF, µmonth, "month"); πE != nil {
						continue
					}
					πTemp001 = µmonth
					if πTemp005, πE = πg.ResolveGlobal(πF, ß_DAYS_IN_MONTH); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, πTemp005, πTemp001); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ß_days_in_month.ToObject(), πTemp008); πE != nil {
				continue
			}
			// line 71: "year, month -> number of days in that month in that year."
			πF.SetLineno(71)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp009}, πg.NewStr("year, month -> number of days in that month in that year.").ToObject()); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ß_days_in_month); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp012, ß__doc__, πTemp009); πE != nil {
				continue
			}
			// line 77: def _days_before_month(year, month):
			πF.SetLineno(77)
			πTemp003 = make([]πg.Param, 2)
			πTemp003[0] = πg.Param{Name: "year", Def: nil}
			πTemp003[1] = πg.Param{Name: "month", Def: nil}
			πTemp009 = πg.NewFunction(πg.NewCode("_days_before_month", "/usr/lib/python2.7/datetime.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µyear *πg.Object = πArgs[0]; _ = µyear
				var µmonth *πg.Object = πArgs[1]; _ = µmonth
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
				var πTemp006 []*πg.Object
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
					// line 78: "year, month -> number of days in year preceding first day of month."
					πF.SetLineno(78)
					// line 79: assert 1 <= month <= 12, 'month must be in 1..12'
					πF.SetLineno(79)
					if πE = πg.CheckLocal(πF, µmonth, "month"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.LE(πF, πg.NewInt(1).ToObject(), µmonth); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp002 {
						goto Label1
					}
					if πTemp001, πE = πg.LE(πF, µmonth, πg.NewInt(12).ToObject()); πE != nil {
						continue
					}
				Label1:
					if πE = πg.Assert(πF, πTemp001, πg.NewStr("month must be in 1..12").ToObject()); πE != nil {
						continue
					}
					// line 80: return _DAYS_BEFORE_MONTH[month] + (month > 2 and _is_leap(year))
					πF.SetLineno(80)
					if πE = πg.CheckLocal(πF, µmonth, "month"); πE != nil {
						continue
					}
					πTemp003 = µmonth
					if πTemp005, πE = πg.ResolveGlobal(πF, ß_DAYS_BEFORE_MONTH); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, πTemp005, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µmonth, "month"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GT(πF, µmonth, πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
					πTemp003 = πTemp005
					if πTemp002, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if !πTemp002 {
						goto Label2
					}
					πTemp006 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µyear, "year"); πE != nil {
						continue
					}
					πTemp006[0] = µyear
					if πTemp005, πE = πg.ResolveGlobal(πF, ß_is_leap); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp005.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					πTemp003 = πTemp007
				Label2:
					if πTemp001, πE = πg.Add(πF, πTemp004, πTemp003); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ß_days_before_month.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 78: "year, month -> number of days in year preceding first day of month."
			πF.SetLineno(78)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp012}, πg.NewStr("year, month -> number of days in year preceding first day of month.").ToObject()); πE != nil {
				continue
			}
			if πTemp013, πE = πg.ResolveGlobal(πF, ß_days_before_month); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp013, ß__doc__, πTemp012); πE != nil {
				continue
			}
			// line 82: def _ymd2ord(year, month, day):
			πF.SetLineno(82)
			πTemp003 = make([]πg.Param, 3)
			πTemp003[0] = πg.Param{Name: "year", Def: nil}
			πTemp003[1] = πg.Param{Name: "month", Def: nil}
			πTemp003[2] = πg.Param{Name: "day", Def: nil}
			πTemp012 = πg.NewFunction(πg.NewCode("_ymd2ord", "/usr/lib/python2.7/datetime.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µyear *πg.Object = πArgs[0]; _ = µyear
				var µmonth *πg.Object = πArgs[1]; _ = µmonth
				var µday *πg.Object = πArgs[2]; _ = µday
				var µdim *πg.Object = πg.UnboundLocal; _ = µdim
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
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 83: "year, month, day -> ordinal, considering 01-Jan-0001 as day 1."
					πF.SetLineno(83)
					// line 84: assert 1 <= month <= 12, 'month must be in 1..12'
					πF.SetLineno(84)
					if πE = πg.CheckLocal(πF, µmonth, "month"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.LE(πF, πg.NewInt(1).ToObject(), µmonth); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp002 {
						goto Label1
					}
					if πTemp001, πE = πg.LE(πF, µmonth, πg.NewInt(12).ToObject()); πE != nil {
						continue
					}
				Label1:
					if πE = πg.Assert(πF, πTemp001, πg.NewStr("month must be in 1..12").ToObject()); πE != nil {
						continue
					}
					// line 85: dim = _days_in_month(year, month)
					πF.SetLineno(85)
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µyear, "year"); πE != nil {
						continue
					}
					πTemp003[0] = µyear
					if πE = πg.CheckLocal(πF, µmonth, "month"); πE != nil {
						continue
					}
					πTemp003[1] = µmonth
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_days_in_month); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µdim = πTemp004
					// line 86: assert 1 <= day <= dim, ('day must be in 1..%d' % dim)
					πF.SetLineno(86)
					if πE = πg.CheckLocal(πF, µdim, "dim"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("day must be in 1..%d").ToObject(), µdim); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µday, "day"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.LE(πF, πg.NewInt(1).ToObject(), µday); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if !πTemp002 {
						goto Label2
					}
					if πE = πg.CheckLocal(πF, µdim, "dim"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.LE(πF, µday, µdim); πE != nil {
						continue
					}
				Label2:
					if πE = πg.Assert(πF, πTemp004, πTemp001); πE != nil {
						continue
					}
					// line 87: return (_days_before_year(year) +
					πF.SetLineno(87)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µyear, "year"); πE != nil {
						continue
					}
					πTemp003[0] = µyear
					if πTemp005, πE = πg.ResolveGlobal(πF, ß_days_before_year); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µyear, "year"); πE != nil {
						continue
					}
					πTemp003[0] = µyear
					if πE = πg.CheckLocal(πF, µmonth, "month"); πE != nil {
						continue
					}
					πTemp003[1] = µmonth
					if πTemp005, πE = πg.ResolveGlobal(πF, ß_days_before_month); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp004, πE = πg.Add(πF, πTemp006, πTemp007); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µday, "day"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, πTemp004, µday); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ß_ymd2ord.ToObject(), πTemp012); πE != nil {
				continue
			}
			// line 83: "year, month, day -> ordinal, considering 01-Jan-0001 as day 1."
			πF.SetLineno(83)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp013}, πg.NewStr("year, month, day -> ordinal, considering 01-Jan-0001 as day 1.").ToObject()); πE != nil {
				continue
			}
			if πTemp014, πE = πg.ResolveGlobal(πF, ß_ymd2ord); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp014, ß__doc__, πTemp013); πE != nil {
				continue
			}
			// line 91: _DI400Y = _days_before_year(401)    # number of days in 400 years
			πF.SetLineno(91)
			πTemp002 = πF.MakeArgs(1)
			πTemp002[0] = πg.NewInt(401).ToObject()
			if πTemp013, πE = πg.ResolveGlobal(πF, ß_days_before_year); πE != nil {
				continue
			}
			if πTemp014, πE = πTemp013.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πF.Globals().SetItem(πF, ß_DI400Y.ToObject(), πTemp014); πE != nil {
				continue
			}
			// line 92: _DI100Y = _days_before_year(101)    #    "    "   "   " 100   "
			πF.SetLineno(92)
			πTemp002 = πF.MakeArgs(1)
			πTemp002[0] = πg.NewInt(101).ToObject()
			if πTemp013, πE = πg.ResolveGlobal(πF, ß_days_before_year); πE != nil {
				continue
			}
			if πTemp014, πE = πTemp013.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πF.Globals().SetItem(πF, ß_DI100Y.ToObject(), πTemp014); πE != nil {
				continue
			}
			// line 93: _DI4Y   = _days_before_year(5)      #    "    "   "   "   4   "
			πF.SetLineno(93)
			πTemp002 = πF.MakeArgs(1)
			πTemp002[0] = πg.NewInt(5).ToObject()
			if πTemp013, πE = πg.ResolveGlobal(πF, ß_days_before_year); πE != nil {
				continue
			}
			if πTemp014, πE = πTemp013.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πF.Globals().SetItem(πF, ß_DI4Y.ToObject(), πTemp014); πE != nil {
				continue
			}
			// line 97: assert _DI4Y == 4 * 365 + 1
			πF.SetLineno(97)
			if πTemp014, πE = πg.ResolveGlobal(πF, ß_DI4Y); πE != nil {
				continue
			}
			if πTemp016, πE = πg.Mul(πF, πg.NewInt(4).ToObject(), πg.NewInt(365).ToObject()); πE != nil {
				continue
			}
			if πTemp015, πE = πg.Add(πF, πTemp016, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			if πTemp013, πE = πg.Eq(πF, πTemp014, πTemp015); πE != nil {
				continue
			}
			if πE = πg.Assert(πF, πTemp013, nil); πE != nil {
				continue
			}
			// line 101: assert _DI400Y == 4 * _DI100Y + 1
			πF.SetLineno(101)
			if πTemp014, πE = πg.ResolveGlobal(πF, ß_DI400Y); πE != nil {
				continue
			}
			if πTemp017, πE = πg.ResolveGlobal(πF, ß_DI100Y); πE != nil {
				continue
			}
			if πTemp016, πE = πg.Mul(πF, πg.NewInt(4).ToObject(), πTemp017); πE != nil {
				continue
			}
			if πTemp015, πE = πg.Add(πF, πTemp016, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			if πTemp013, πE = πg.Eq(πF, πTemp014, πTemp015); πE != nil {
				continue
			}
			if πE = πg.Assert(πF, πTemp013, nil); πE != nil {
				continue
			}
			// line 105: assert _DI100Y == 25 * _DI4Y - 1
			πF.SetLineno(105)
			if πTemp014, πE = πg.ResolveGlobal(πF, ß_DI100Y); πE != nil {
				continue
			}
			if πTemp017, πE = πg.ResolveGlobal(πF, ß_DI4Y); πE != nil {
				continue
			}
			if πTemp016, πE = πg.Mul(πF, πg.NewInt(25).ToObject(), πTemp017); πE != nil {
				continue
			}
			if πTemp015, πE = πg.Sub(πF, πTemp016, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			if πTemp013, πE = πg.Eq(πF, πTemp014, πTemp015); πE != nil {
				continue
			}
			if πE = πg.Assert(πF, πTemp013, nil); πE != nil {
				continue
			}
			// line 107: _US_PER_US = 1
			πF.SetLineno(107)
			if πE = πF.Globals().SetItem(πF, ß_US_PER_US.ToObject(), πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			// line 108: _US_PER_MS = 1000
			πF.SetLineno(108)
			if πE = πF.Globals().SetItem(πF, ß_US_PER_MS.ToObject(), πg.NewInt(1000).ToObject()); πE != nil {
				continue
			}
			// line 109: _US_PER_SECOND = 1000000
			πF.SetLineno(109)
			if πE = πF.Globals().SetItem(πF, ß_US_PER_SECOND.ToObject(), πg.NewInt(1000000).ToObject()); πE != nil {
				continue
			}
			// line 110: _US_PER_MINUTE = 60000000
			πF.SetLineno(110)
			if πE = πF.Globals().SetItem(πF, ß_US_PER_MINUTE.ToObject(), πg.NewInt(60000000).ToObject()); πE != nil {
				continue
			}
			// line 111: _SECONDS_PER_DAY = 24 * 3600
			πF.SetLineno(111)
			if πTemp013, πE = πg.Mul(πF, πg.NewInt(24).ToObject(), πg.NewInt(3600).ToObject()); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_SECONDS_PER_DAY.ToObject(), πTemp013); πE != nil {
				continue
			}
			// line 112: _US_PER_HOUR = 3600000000
			πF.SetLineno(112)
			if πE = πF.Globals().SetItem(πF, ß_US_PER_HOUR.ToObject(), πg.NewInt(3600000000).ToObject()); πE != nil {
				continue
			}
			// line 113: _US_PER_DAY = 86400000000
			πF.SetLineno(113)
			if πE = πF.Globals().SetItem(πF, ß_US_PER_DAY.ToObject(), πg.NewInt(86400000000).ToObject()); πE != nil {
				continue
			}
			// line 114: _US_PER_WEEK = 604800000000
			πF.SetLineno(114)
			if πE = πF.Globals().SetItem(πF, ß_US_PER_WEEK.ToObject(), πg.NewInt(604800000000).ToObject()); πE != nil {
				continue
			}
			// line 116: def _ord2ymd(n):
			πF.SetLineno(116)
			πTemp003 = make([]πg.Param, 1)
			πTemp003[0] = πg.Param{Name: "n", Def: nil}
			πTemp013 = πg.NewFunction(πg.NewCode("_ord2ymd", "/usr/lib/python2.7/datetime.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µn *πg.Object = πArgs[0]; _ = µn
				var µn400 *πg.Object = πg.UnboundLocal; _ = µn400
				var µyear *πg.Object = πg.UnboundLocal; _ = µyear
				var µn100 *πg.Object = πg.UnboundLocal; _ = µn100
				var µn4 *πg.Object = πg.UnboundLocal; _ = µn4
				var µn1 *πg.Object = πg.UnboundLocal; _ = µn1
				var µleapyear *πg.Object = πg.UnboundLocal; _ = µleapyear
				var µmonth *πg.Object = πg.UnboundLocal; _ = µmonth
				var µpreceding *πg.Object = πg.UnboundLocal; _ = µpreceding
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
				var πTemp007 bool
				_ = πTemp007
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 117: "ordinal -> (year, month, day), considering 01-Jan-0001 as day 1."
					πF.SetLineno(117)
					// line 139: n -= 1
					πF.SetLineno(139)
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.ISub(πF, µn, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µn = πTemp001
					// line 140: n400, n = divmod(n, _DI400Y)
					πF.SetLineno(140)
					πTemp002 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					πTemp002[0] = µn
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_DI400Y); πE != nil {
						continue
					}
					πTemp002[1] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßdivmod); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp004}}}, πTemp003); πE != nil {
						continue
					}
					µn400 = πTemp001
					µn = πTemp004
					// line 141: year = n400 * 400 + 1   # ..., -399, 1, 401, ...
					πF.SetLineno(141)
					if πE = πg.CheckLocal(πF, µn400, "n400"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Mul(πF, µn400, πg.NewInt(400).ToObject()); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, πTemp003, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µyear = πTemp001
					// line 148: n100, n = divmod(n, _DI100Y)
					πF.SetLineno(148)
					πTemp002 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					πTemp002[0] = µn
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_DI100Y); πE != nil {
						continue
					}
					πTemp002[1] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßdivmod); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp004}}}, πTemp003); πE != nil {
						continue
					}
					µn100 = πTemp001
					µn = πTemp004
					// line 151: n4, n = divmod(n, _DI4Y)
					πF.SetLineno(151)
					πTemp002 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					πTemp002[0] = µn
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_DI4Y); πE != nil {
						continue
					}
					πTemp002[1] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßdivmod); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp004}}}, πTemp003); πE != nil {
						continue
					}
					µn4 = πTemp001
					µn = πTemp004
					// line 155: n1, n = divmod(n, 365)
					πF.SetLineno(155)
					πTemp002 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					πTemp002[0] = µn
					πTemp002[1] = πg.NewInt(365).ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßdivmod); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp004}}}, πTemp003); πE != nil {
						continue
					}
					µn1 = πTemp001
					µn = πTemp004
					// line 157: year += n100 * 100 + n4 * 4 + n1
					πF.SetLineno(157)
					if πE = πg.CheckLocal(πF, µyear, "year"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µn100, "n100"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Mul(πF, µn100, πg.NewInt(100).ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µn4, "n4"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Mul(πF, µn4, πg.NewInt(4).ToObject()); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Add(πF, πTemp004, πTemp005); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µn1, "n1"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, πTemp003, µn1); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IAdd(πF, µyear, πTemp001); πE != nil {
						continue
					}
					µyear = πTemp003
					if πE = πg.CheckLocal(πF, µn1, "n1"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Eq(πF, µn1, πg.NewInt(4).ToObject()); πE != nil {
						continue
					}
					πTemp001 = πTemp003
					if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label1
					}
					if πE = πg.CheckLocal(πF, µn100, "n100"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Eq(πF, µn100, πg.NewInt(4).ToObject()); πE != nil {
						continue
					}
					πTemp001 = πTemp003
				Label1:
					if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label2
					}
					goto Label3
					// line 158: if n1 == 4 or n100 == 4:
					πF.SetLineno(158)
				Label2:
					// line 159: assert n == 0
					πF.SetLineno(159)
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, µn, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
						continue
					}
					// line 160: return year-1, 12, 31
					πF.SetLineno(160)
					if πE = πg.CheckLocal(πF, µyear, "year"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Sub(πF, µyear, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple3(πTemp003, πg.NewInt(12).ToObject(), πg.NewInt(31).ToObject()).ToObject()
					πR = πTemp001
					continue
					goto Label3
				Label3:
					// line 164: leapyear = n1 == 3 and (n4 != 24 or n100 == 3)
					πF.SetLineno(164)
					if πE = πg.CheckLocal(πF, µn1, "n1"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Eq(πF, µn1, πg.NewInt(3).ToObject()); πE != nil {
						continue
					}
					πTemp001 = πTemp003
					if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp006 {
						goto Label4
					}
					if πE = πg.CheckLocal(πF, µn4, "n4"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.NE(πF, µn4, πg.NewInt(24).ToObject()); πE != nil {
						continue
					}
					πTemp003 = πTemp004
					if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label5
					}
					if πE = πg.CheckLocal(πF, µn100, "n100"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Eq(πF, µn100, πg.NewInt(3).ToObject()); πE != nil {
						continue
					}
					πTemp003 = πTemp004
				Label5:
					πTemp001 = πTemp003
				Label4:
					µleapyear = πTemp001
					// line 165: assert leapyear == _is_leap(year)
					πF.SetLineno(165)
					if πE = πg.CheckLocal(πF, µleapyear, "leapyear"); πE != nil {
						continue
					}
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µyear, "year"); πE != nil {
						continue
					}
					πTemp002[0] = µyear
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_is_leap); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.Eq(πF, µleapyear, πTemp004); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
						continue
					}
					// line 166: month = (n + 50) >> 5
					πF.SetLineno(166)
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Add(πF, µn, πg.NewInt(50).ToObject()); πE != nil {
						continue
					}
					if πTemp001, πE = πg.RShift(πF, πTemp003, πg.NewInt(5).ToObject()); πE != nil {
						continue
					}
					µmonth = πTemp001
					// line 167: preceding = _DAYS_BEFORE_MONTH[month] + (month > 2 and leapyear)
					πF.SetLineno(167)
					if πE = πg.CheckLocal(πF, µmonth, "month"); πE != nil {
						continue
					}
					πTemp003 = µmonth
					if πTemp005, πE = πg.ResolveGlobal(πF, ß_DAYS_BEFORE_MONTH); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, πTemp005, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µmonth, "month"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GT(πF, µmonth, πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
					πTemp003 = πTemp005
					if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if !πTemp006 {
						goto Label6
					}
					if πE = πg.CheckLocal(πF, µleapyear, "leapyear"); πE != nil {
						continue
					}
					πTemp003 = µleapyear
				Label6:
					if πTemp001, πE = πg.Add(πF, πTemp004, πTemp003); πE != nil {
						continue
					}
					µpreceding = πTemp001
					if πE = πg.CheckLocal(πF, µpreceding, "preceding"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GT(πF, µpreceding, µn); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label7
					}
					goto Label8
					// line 168: if preceding > n:  # estimate is too large
					πF.SetLineno(168)
				Label7:
					// line 169: month -= 1
					πF.SetLineno(169)
					if πE = πg.CheckLocal(πF, µmonth, "month"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.ISub(πF, µmonth, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µmonth = πTemp001
					// line 170: preceding -= _DAYS_IN_MONTH[month] + (month == 2 and leapyear)
					πF.SetLineno(170)
					if πE = πg.CheckLocal(πF, µpreceding, "preceding"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µmonth, "month"); πE != nil {
						continue
					}
					πTemp003 = µmonth
					if πTemp005, πE = πg.ResolveGlobal(πF, ß_DAYS_IN_MONTH); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, πTemp005, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µmonth, "month"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Eq(πF, µmonth, πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
					πTemp003 = πTemp005
					if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if !πTemp006 {
						goto Label9
					}
					if πE = πg.CheckLocal(πF, µleapyear, "leapyear"); πE != nil {
						continue
					}
					πTemp003 = µleapyear
				Label9:
					if πTemp001, πE = πg.Add(πF, πTemp004, πTemp003); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ISub(πF, µpreceding, πTemp001); πE != nil {
						continue
					}
					µpreceding = πTemp003
					goto Label8
				Label8:
					// line 171: n -= preceding
					πF.SetLineno(171)
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µpreceding, "preceding"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.ISub(πF, µn, µpreceding); πE != nil {
						continue
					}
					µn = πTemp001
					// line 172: assert 0 <= n < _days_in_month(year, month)
					πF.SetLineno(172)
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.LE(πF, πg.NewInt(0).ToObject(), µn); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp006 {
						goto Label10
					}
					πTemp002 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µyear, "year"); πE != nil {
						continue
					}
					πTemp002[0] = µyear
					if πE = πg.CheckLocal(πF, µmonth, "month"); πE != nil {
						continue
					}
					πTemp002[1] = µmonth
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_days_in_month); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.LT(πF, µn, πTemp004); πE != nil {
						continue
					}
				Label10:
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
						continue
					}
					// line 176: return year, month, n+1
					πF.SetLineno(176)
					if πE = πg.CheckLocal(πF, µyear, "year"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µmonth, "month"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Add(πF, µn, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple3(µyear, µmonth, πTemp003).ToObject()
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
			if πE = πF.Globals().SetItem(πF, ß_ord2ymd.ToObject(), πTemp013); πE != nil {
				continue
			}
			// line 117: "ordinal -> (year, month, day), considering 01-Jan-0001 as day 1."
			πF.SetLineno(117)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp014}, πg.NewStr("ordinal -> (year, month, day), considering 01-Jan-0001 as day 1.").ToObject()); πE != nil {
				continue
			}
			if πTemp015, πE = πg.ResolveGlobal(πF, ß_ord2ymd); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp015, ß__doc__, πTemp014); πE != nil {
				continue
			}
			// line 179: _MONTHNAMES = [None, "Jan", "Feb", "Mar", "Apr", "May", "Jun",
			πF.SetLineno(179)
			πTemp002 = make([]*πg.Object, 13)
			if πTemp014, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp002[0] = πTemp014
			πTemp002[1] = ßJan.ToObject()
			πTemp002[2] = ßFeb.ToObject()
			πTemp002[3] = ßMar.ToObject()
			πTemp002[4] = ßApr.ToObject()
			πTemp002[5] = ßMay.ToObject()
			πTemp002[6] = ßJun.ToObject()
			πTemp002[7] = ßJul.ToObject()
			πTemp002[8] = ßAug.ToObject()
			πTemp002[9] = ßSep.ToObject()
			πTemp002[10] = ßOct.ToObject()
			πTemp002[11] = ßNov.ToObject()
			πTemp002[12] = ßDec.ToObject()
			πTemp014 = πg.NewList(πTemp002...).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_MONTHNAMES.ToObject(), πTemp014); πE != nil {
				continue
			}
			// line 181: _DAYNAMES = [None, "Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"]
			πF.SetLineno(181)
			πTemp002 = make([]*πg.Object, 8)
			if πTemp014, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp002[0] = πTemp014
			πTemp002[1] = ßMon.ToObject()
			πTemp002[2] = ßTue.ToObject()
			πTemp002[3] = ßWed.ToObject()
			πTemp002[4] = ßThu.ToObject()
			πTemp002[5] = ßFri.ToObject()
			πTemp002[6] = ßSat.ToObject()
			πTemp002[7] = ßSun.ToObject()
			πTemp014 = πg.NewList(πTemp002...).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_DAYNAMES.ToObject(), πTemp014); πE != nil {
				continue
			}
			// line 184: def _build_struct_time(y, m, d, hh, mm, ss, dstflag):
			πF.SetLineno(184)
			πTemp003 = make([]πg.Param, 7)
			πTemp003[0] = πg.Param{Name: "y", Def: nil}
			πTemp003[1] = πg.Param{Name: "m", Def: nil}
			πTemp003[2] = πg.Param{Name: "d", Def: nil}
			πTemp003[3] = πg.Param{Name: "hh", Def: nil}
			πTemp003[4] = πg.Param{Name: "mm", Def: nil}
			πTemp003[5] = πg.Param{Name: "ss", Def: nil}
			πTemp003[6] = πg.Param{Name: "dstflag", Def: nil}
			πTemp014 = πg.NewFunction(πg.NewCode("_build_struct_time", "/usr/lib/python2.7/datetime.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µy *πg.Object = πArgs[0]; _ = µy
				var µm *πg.Object = πArgs[1]; _ = µm
				var µd *πg.Object = πArgs[2]; _ = µd
				var µhh *πg.Object = πArgs[3]; _ = µhh
				var µmm *πg.Object = πArgs[4]; _ = µmm
				var µss *πg.Object = πArgs[5]; _ = µss
				var µdstflag *πg.Object = πArgs[6]; _ = µdstflag
				var µwday *πg.Object = πg.UnboundLocal; _ = µwday
				var µdnum *πg.Object = πg.UnboundLocal; _ = µdnum
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
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 185: wday = (_ymd2ord(y, m, d) + 6) % 7
					πF.SetLineno(185)
					πTemp003 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
						continue
					}
					πTemp003[0] = µy
					if πE = πg.CheckLocal(πF, µm, "m"); πE != nil {
						continue
					}
					πTemp003[1] = µm
					if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
						continue
					}
					πTemp003[2] = µd
					if πTemp004, πE = πg.ResolveGlobal(πF, ß_ymd2ord); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp002, πE = πg.Add(πF, πTemp005, πg.NewInt(6).ToObject()); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Mod(πF, πTemp002, πg.NewInt(7).ToObject()); πE != nil {
						continue
					}
					µwday = πTemp001
					// line 186: dnum = _days_before_month(y, m) + d
					πF.SetLineno(186)
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
						continue
					}
					πTemp003[0] = µy
					if πE = πg.CheckLocal(πF, µm, "m"); πE != nil {
						continue
					}
					πTemp003[1] = µm
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_days_before_month); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, πTemp004, µd); πE != nil {
						continue
					}
					µdnum = πTemp001
					// line 187: return _time.struct_time((y, m, d, hh, mm, ss, wday, dnum, dstflag))
					πF.SetLineno(187)
					πTemp003 = πF.MakeArgs(1)
					πTemp006 = make([]*πg.Object, 9)
					if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
						continue
					}
					πTemp006[0] = µy
					if πE = πg.CheckLocal(πF, µm, "m"); πE != nil {
						continue
					}
					πTemp006[1] = µm
					if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
						continue
					}
					πTemp006[2] = µd
					if πE = πg.CheckLocal(πF, µhh, "hh"); πE != nil {
						continue
					}
					πTemp006[3] = µhh
					if πE = πg.CheckLocal(πF, µmm, "mm"); πE != nil {
						continue
					}
					πTemp006[4] = µmm
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp006[5] = µss
					if πE = πg.CheckLocal(πF, µwday, "wday"); πE != nil {
						continue
					}
					πTemp006[6] = µwday
					if πE = πg.CheckLocal(πF, µdnum, "dnum"); πE != nil {
						continue
					}
					πTemp006[7] = µdnum
					if πE = πg.CheckLocal(πF, µdstflag, "dstflag"); πE != nil {
						continue
					}
					πTemp006[8] = µdstflag
					πTemp001 = πg.NewTuple(πTemp006...).ToObject()
					πTemp003[0] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_time); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßstruct_time, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
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
			if πE = πF.Globals().SetItem(πF, ß_build_struct_time.ToObject(), πTemp014); πE != nil {
				continue
			}
			// line 189: def _format_time(hh, mm, ss, us):
			πF.SetLineno(189)
			πTemp003 = make([]πg.Param, 4)
			πTemp003[0] = πg.Param{Name: "hh", Def: nil}
			πTemp003[1] = πg.Param{Name: "mm", Def: nil}
			πTemp003[2] = πg.Param{Name: "ss", Def: nil}
			πTemp003[3] = πg.Param{Name: "us", Def: nil}
			πTemp015 = πg.NewFunction(πg.NewCode("_format_time", "/usr/lib/python2.7/datetime.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µhh *πg.Object = πArgs[0]; _ = µhh
				var µmm *πg.Object = πArgs[1]; _ = µmm
				var µss *πg.Object = πArgs[2]; _ = µss
				var µus *πg.Object = πArgs[3]; _ = µus
				var µresult *πg.Object = πg.UnboundLocal; _ = µresult
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
					// line 191: result = "%02d:%02d:%02d" % (hh, mm, ss)
					πF.SetLineno(191)
					if πE = πg.CheckLocal(πF, µhh, "hh"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µmm, "mm"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple3(µhh, µmm, µss).ToObject()
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("%02d:%02d:%02d").ToObject(), πTemp002); πE != nil {
						continue
					}
					µresult = πTemp001
					if πE = πg.CheckLocal(πF, µus, "us"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsTrue(πF, µus); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label1
					}
					goto Label2
					// line 192: if us:
					πF.SetLineno(192)
				Label1:
					// line 193: result += ".%06d" % us
					πF.SetLineno(193)
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µus, "us"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Mod(πF, πg.NewStr(".%06d").ToObject(), µus); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IAdd(πF, µresult, πTemp001); πE != nil {
						continue
					}
					µresult = πTemp002
					goto Label2
				Label2:
					// line 194: return result
					πF.SetLineno(194)
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
			if πE = πF.Globals().SetItem(πF, ß_format_time.ToObject(), πTemp015); πE != nil {
				continue
			}
			// line 258: def _check_tzname(name):
			πF.SetLineno(258)
			πTemp003 = make([]πg.Param, 1)
			πTemp003[0] = πg.Param{Name: "name", Def: nil}
			πTemp016 = πg.NewFunction(πg.NewCode("_check_tzname", "/usr/lib/python2.7/datetime.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µname *πg.Object = πArgs[0]; _ = µname
				var πTemp001 *πg.Object
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
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(µname != πTemp004).ToObject()
					πTemp001 = πTemp003
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp002 {
						goto Label1
					}
					πTemp005 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp005[0] = µname
					if πTemp004, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
						continue
					}
					πTemp005[1] = πTemp004
					if πTemp004, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πTemp007, πE = πg.IsTrue(πF, πTemp006); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(!πTemp007).ToObject()
					πTemp001 = πTemp003
				Label1:
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label2
					}
					goto Label3
					// line 259: if name is not None and not isinstance(name, str):
					πF.SetLineno(259)
				Label2:
					πTemp005 = πF.MakeArgs(1)
					πTemp008 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp008[0] = µname
					if πTemp003, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp008, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp008)
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("tzinfo.tzname() must return None or string, not '%s'").ToObject(), πTemp004); πE != nil {
						continue
					}
					πTemp005[0] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					// line 260: raise TypeError("tzinfo.tzname() must return None or string, "
					πF.SetLineno(260)
					πE = πF.Raise(πTemp003, nil, nil)
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
			if πE = πF.Globals().SetItem(πF, ß_check_tzname.ToObject(), πTemp016); πE != nil {
				continue
			}
			// line 269: def _check_utc_offset(name, offset):
			πF.SetLineno(269)
			πTemp003 = make([]πg.Param, 2)
			πTemp003[0] = πg.Param{Name: "name", Def: nil}
			πTemp003[1] = πg.Param{Name: "offset", Def: nil}
			πTemp017 = πg.NewFunction(πg.NewCode("_check_utc_offset", "/usr/lib/python2.7/datetime.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µname *πg.Object = πArgs[0]; _ = µname
				var µoffset *πg.Object = πArgs[1]; _ = µoffset
				var µdays *πg.Object = πg.UnboundLocal; _ = µdays
				var µseconds *πg.Object = πg.UnboundLocal; _ = µseconds
				var µminutes *πg.Object = πg.UnboundLocal; _ = µminutes
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
				var πTemp006 []*πg.Object
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
					// line 270: assert name in ("utcoffset", "dst")
					πF.SetLineno(270)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(ßutcoffset.ToObject(), ßdst.ToObject()).ToObject()
					if πTemp003, πE = πg.Contains(πF, πTemp002, µname); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(πTemp003).ToObject()
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µoffset, "offset"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µoffset == πTemp002).ToObject()
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label1
					}
					goto Label2
					// line 271: if offset is None:
					πF.SetLineno(271)
				Label1:
					// line 272: return
					πF.SetLineno(272)
					πR = πg.None
					continue
					goto Label2
				Label2:
					πTemp004 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µoffset, "offset"); πE != nil {
						continue
					}
					πTemp004[0] = µoffset
					if πTemp002, πE = πg.ResolveGlobal(πF, ßtimedelta); πE != nil {
						continue
					}
					πTemp004[1] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp003, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(!πTemp003).ToObject()
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label3
					}
					goto Label4
					// line 273: if not isinstance(offset, timedelta):
					πF.SetLineno(273)
				Label3:
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp006 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µoffset, "offset"); πE != nil {
						continue
					}
					πTemp006[0] = µoffset
					if πTemp005, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp005.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					πTemp002 = πg.NewTuple2(µname, πTemp007).ToObject()
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("tzinfo.%s() must return None or timedelta, not '%s'").ToObject(), πTemp002); πE != nil {
						continue
					}
					πTemp004[0] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 274: raise TypeError("tzinfo.%s() must return None "
					πF.SetLineno(274)
					πE = πF.Raise(πTemp002, nil, nil)
					continue
					goto Label4
				Label4:
					// line 276: days = offset.days
					πF.SetLineno(276)
					if πE = πg.CheckLocal(πF, µoffset, "offset"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µoffset, ßdays, nil); πE != nil {
						continue
					}
					µdays = πTemp001
					if πE = πg.CheckLocal(πF, µdays, "days"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.LT(πF, µdays, πTemp005); πE != nil {
						continue
					}
					πTemp001 = πTemp002
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label5
					}
					if πE = πg.CheckLocal(πF, µdays, "days"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GT(πF, µdays, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					πTemp001 = πTemp002
				Label5:
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label6
					}
					goto Label7
					// line 277: if days < -1 or days > 0:
					πF.SetLineno(277)
				Label6:
					// line 278: offset = 1440  # trigger out-of-range
					πF.SetLineno(278)
					µoffset = πg.NewInt(1440).ToObject()
					goto Label8
				Label7:
					// line 280: seconds = days * 86400 + offset.seconds
					πF.SetLineno(280)
					if πE = πg.CheckLocal(πF, µdays, "days"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mul(πF, µdays, πg.NewInt(86400).ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µoffset, "offset"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µoffset, ßseconds, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, πTemp002, πTemp005); πE != nil {
						continue
					}
					µseconds = πTemp001
					// line 281: minutes, seconds = divmod(seconds, 60)
					πF.SetLineno(281)
					πTemp004 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µseconds, "seconds"); πE != nil {
						continue
					}
					πTemp004[0] = µseconds
					πTemp004[1] = πg.NewInt(60).ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßdivmod); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp005}}}, πTemp002); πE != nil {
						continue
					}
					µminutes = πTemp001
					µseconds = πTemp005
					if πE = πg.CheckLocal(πF, µseconds, "seconds"); πE != nil {
						continue
					}
					πTemp001 = µseconds
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label9
					}
					if πE = πg.CheckLocal(πF, µoffset, "offset"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µoffset, ßmicroseconds, nil); πE != nil {
						continue
					}
					πTemp001 = πTemp002
				Label9:
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label10
					}
					goto Label11
					// line 282: if seconds or offset.microseconds:
					πF.SetLineno(282)
				Label10:
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("tzinfo.%s() must return a whole number of minutes").ToObject(), µname); πE != nil {
						continue
					}
					πTemp004[0] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 283: raise ValueError("tzinfo.%s() must return a whole number "
					πF.SetLineno(283)
					πE = πF.Raise(πTemp002, nil, nil)
					continue
					goto Label11
				Label11:
					// line 285: offset = minutes
					πF.SetLineno(285)
					if πE = πg.CheckLocal(πF, µminutes, "minutes"); πE != nil {
						continue
					}
					µoffset = µminutes
					goto Label8
				Label8:
					if πTemp005, πE = πg.Neg(πF, πg.NewInt(1440).ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µoffset, "offset"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.LT(πF, πTemp005, µoffset); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp003 {
						goto Label12
					}
					if πTemp002, πE = πg.LT(πF, µoffset, πg.NewInt(1440).ToObject()); πE != nil {
						continue
					}
				Label12:
					if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(!πTemp003).ToObject()
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label13
					}
					goto Label14
					// line 286: if not -1440 < offset < 1440:
					πF.SetLineno(286)
				Label13:
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µoffset, "offset"); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(µname, µoffset).ToObject()
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("%s()=%d, must be in -1439..1439").ToObject(), πTemp002); πE != nil {
						continue
					}
					πTemp004[0] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 287: raise ValueError("%s()=%d, must be in -1439..1439" % (name, offset))
					πF.SetLineno(287)
					πE = πF.Raise(πTemp002, nil, nil)
					continue
					goto Label14
				Label14:
					// line 288: return offset
					πF.SetLineno(288)
					if πE = πg.CheckLocal(πF, µoffset, "offset"); πE != nil {
						continue
					}
					πR = µoffset
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_check_utc_offset.ToObject(), πTemp017); πE != nil {
				continue
			}
			// line 290: def _check_int_field(value):
			πF.SetLineno(290)
			πTemp003 = make([]πg.Param, 1)
			πTemp003[0] = πg.Param{Name: "value", Def: nil}
			πTemp018 = πg.NewFunction(πg.NewCode("_check_int_field", "/usr/lib/python2.7/datetime.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µvalue *πg.Object = πArgs[0]; _ = µvalue
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
				var πTemp006 []*πg.Object
				_ = πTemp006
				var πTemp007 *πg.BaseException
				_ = πTemp007
				var πTemp008 *πg.Traceback
				_ = πTemp008
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 6: goto Label6
					default: panic("unexpected function state")
					}
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					πTemp001[0] = µvalue
					if πTemp002, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
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
					goto Label2
					// line 291: if isinstance(value, int):
					πF.SetLineno(291)
				Label1:
					// line 292: return int(value)
					πF.SetLineno(292)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					πTemp001[0] = µvalue
					if πTemp002, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
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
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					πTemp001[0] = µvalue
					if πTemp003, πE = πg.ResolveGlobal(πF, ßfloat); πE != nil {
						continue
					}
					πTemp001[1] = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp004, πE = πg.IsTrue(πF, πTemp005); πE != nil {
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
					// line 293: if not isinstance(value, float):
					πF.SetLineno(293)
				Label3:
					// line 294: try:
					πF.SetLineno(294)
					πF.PushCheckpoint(6)
					// line 295: value = value.__int__()
					πF.SetLineno(295)
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µvalue, ß__int__, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					µvalue = πTemp003
					πF.PopCheckpoint()
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					πTemp001[0] = µvalue
					if πTemp002, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
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
						goto Label7
					}
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					πTemp001[0] = µvalue
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlong); πE != nil {
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
						goto Label8
					}
					goto Label9
					// line 299: if isinstance(value, int):
					πF.SetLineno(299)
				Label7:
					// line 300: return int(value)
					πF.SetLineno(300)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					πTemp001[0] = µvalue
					if πTemp002, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πR = πTemp003
					continue
					goto Label9
					// line 301: elif isinstance(value, long):
					πF.SetLineno(301)
				Label8:
					// line 302: return int(long(value))
					πF.SetLineno(302)
					πTemp001 = πF.MakeArgs(1)
					πTemp006 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					πTemp006[0] = µvalue
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlong); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					πTemp001[0] = πTemp003
					if πTemp002, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πR = πTemp003
					continue
					goto Label9
				Label9:
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("__int__ method should return an integer").ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 303: raise TypeError('__int__ method should return an integer')
					πF.SetLineno(303)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					goto Label5
				Label6:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp007, πTemp008 = πF.ExcInfo()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßAttributeError); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsInstance(πF, πTemp007.ToObject(), πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label10
					}
					πE = πF.Raise(πTemp007.ToObject(), nil, πTemp008.ToObject())
					continue
					// line 296: except AttributeError:
					πF.SetLineno(296)
				Label10:
					// line 297: pass
					πF.SetLineno(297)
					πF.RestoreExc(nil, nil)
					goto Label5
				Label5:
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("an integer is required").ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 304: raise TypeError('an integer is required')
					πF.SetLineno(304)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					goto Label4
				Label4:
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("integer argument expected, got float").ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 305: raise TypeError('integer argument expected, got float')
					πF.SetLineno(305)
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
			if πE = πF.Globals().SetItem(πF, ß_check_int_field.ToObject(), πTemp018); πE != nil {
				continue
			}
			// line 307: def _check_date_fields(year, month, day):
			πF.SetLineno(307)
			πTemp003 = make([]πg.Param, 3)
			πTemp003[0] = πg.Param{Name: "year", Def: nil}
			πTemp003[1] = πg.Param{Name: "month", Def: nil}
			πTemp003[2] = πg.Param{Name: "day", Def: nil}
			πTemp019 = πg.NewFunction(πg.NewCode("_check_date_fields", "/usr/lib/python2.7/datetime.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µyear *πg.Object = πArgs[0]; _ = µyear
				var µmonth *πg.Object = πArgs[1]; _ = µmonth
				var µday *πg.Object = πArgs[2]; _ = µday
				var µdim *πg.Object = πg.UnboundLocal; _ = µdim
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
					// line 308: year = _check_int_field(year)
					πF.SetLineno(308)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µyear, "year"); πE != nil {
						continue
					}
					πTemp001[0] = µyear
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_check_int_field); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µyear = πTemp003
					// line 309: month = _check_int_field(month)
					πF.SetLineno(309)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µmonth, "month"); πE != nil {
						continue
					}
					πTemp001[0] = µmonth
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_check_int_field); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µmonth = πTemp003
					// line 310: day = _check_int_field(day)
					πF.SetLineno(310)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µday, "day"); πE != nil {
						continue
					}
					πTemp001[0] = µday
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_check_int_field); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µday = πTemp003
					if πTemp004, πE = πg.ResolveGlobal(πF, ßMINYEAR); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µyear, "year"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.LE(πF, πTemp004, µyear); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if !πTemp005 {
						goto Label1
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßMAXYEAR); πE != nil {
						continue
					}
					if πTemp003, πE = πg.LE(πF, µyear, πTemp004); πE != nil {
						continue
					}
				Label1:
					if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(!πTemp005).ToObject()
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label2
					}
					goto Label3
					// line 311: if not MINYEAR <= year <= MAXYEAR:
					πF.SetLineno(311)
				Label2:
					πTemp001 = πF.MakeArgs(2)
					if πTemp004, πE = πg.ResolveGlobal(πF, ßMINYEAR); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveGlobal(πF, ßMAXYEAR); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple2(πTemp004, πTemp006).ToObject()
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("year must be in %d..%d").ToObject(), πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πE = πg.CheckLocal(πF, µyear, "year"); πE != nil {
						continue
					}
					πTemp001[1] = µyear
					if πTemp002, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 312: raise ValueError('year must be in %d..%d' % (MINYEAR, MAXYEAR), year)
					πF.SetLineno(312)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					goto Label3
				Label3:
					if πE = πg.CheckLocal(πF, µmonth, "month"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.LE(πF, πg.NewInt(1).ToObject(), µmonth); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if !πTemp005 {
						goto Label4
					}
					if πTemp003, πE = πg.LE(πF, µmonth, πg.NewInt(12).ToObject()); πE != nil {
						continue
					}
				Label4:
					if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(!πTemp005).ToObject()
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label5
					}
					goto Label6
					// line 313: if not 1 <= month <= 12:
					πF.SetLineno(313)
				Label5:
					πTemp001 = πF.MakeArgs(2)
					πTemp001[0] = πg.NewStr("month must be in 1..12").ToObject()
					if πE = πg.CheckLocal(πF, µmonth, "month"); πE != nil {
						continue
					}
					πTemp001[1] = µmonth
					if πTemp002, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 314: raise ValueError('month must be in 1..12', month)
					πF.SetLineno(314)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					goto Label6
				Label6:
					// line 315: dim = _days_in_month(year, month)
					πF.SetLineno(315)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µyear, "year"); πE != nil {
						continue
					}
					πTemp001[0] = µyear
					if πE = πg.CheckLocal(πF, µmonth, "month"); πE != nil {
						continue
					}
					πTemp001[1] = µmonth
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_days_in_month); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µdim = πTemp003
					if πE = πg.CheckLocal(πF, µday, "day"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.LE(πF, πg.NewInt(1).ToObject(), µday); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if !πTemp005 {
						goto Label7
					}
					if πE = πg.CheckLocal(πF, µdim, "dim"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.LE(πF, µday, µdim); πE != nil {
						continue
					}
				Label7:
					if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(!πTemp005).ToObject()
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label8
					}
					goto Label9
					// line 316: if not 1 <= day <= dim:
					πF.SetLineno(316)
				Label8:
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µdim, "dim"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("day must be in 1..%d").ToObject(), µdim); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πE = πg.CheckLocal(πF, µday, "day"); πE != nil {
						continue
					}
					πTemp001[1] = µday
					if πTemp002, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 317: raise ValueError('day must be in 1..%d' % dim, day)
					πF.SetLineno(317)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					goto Label9
				Label9:
					// line 318: return year, month, day
					πF.SetLineno(318)
					if πE = πg.CheckLocal(πF, µyear, "year"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µmonth, "month"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µday, "day"); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple3(µyear, µmonth, µday).ToObject()
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
			if πE = πF.Globals().SetItem(πF, ß_check_date_fields.ToObject(), πTemp019); πE != nil {
				continue
			}
			// line 320: def _check_time_fields(hour, minute, second, microsecond):
			πF.SetLineno(320)
			πTemp003 = make([]πg.Param, 4)
			πTemp003[0] = πg.Param{Name: "hour", Def: nil}
			πTemp003[1] = πg.Param{Name: "minute", Def: nil}
			πTemp003[2] = πg.Param{Name: "second", Def: nil}
			πTemp003[3] = πg.Param{Name: "microsecond", Def: nil}
			πTemp020 = πg.NewFunction(πg.NewCode("_check_time_fields", "/usr/lib/python2.7/datetime.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µhour *πg.Object = πArgs[0]; _ = µhour
				var µminute *πg.Object = πArgs[1]; _ = µminute
				var µsecond *πg.Object = πArgs[2]; _ = µsecond
				var µmicrosecond *πg.Object = πArgs[3]; _ = µmicrosecond
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
					// line 321: hour = _check_int_field(hour)
					πF.SetLineno(321)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µhour, "hour"); πE != nil {
						continue
					}
					πTemp001[0] = µhour
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_check_int_field); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µhour = πTemp003
					// line 322: minute = _check_int_field(minute)
					πF.SetLineno(322)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µminute, "minute"); πE != nil {
						continue
					}
					πTemp001[0] = µminute
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_check_int_field); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µminute = πTemp003
					// line 323: second = _check_int_field(second)
					πF.SetLineno(323)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µsecond, "second"); πE != nil {
						continue
					}
					πTemp001[0] = µsecond
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_check_int_field); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µsecond = πTemp003
					// line 324: microsecond = _check_int_field(microsecond)
					πF.SetLineno(324)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µmicrosecond, "microsecond"); πE != nil {
						continue
					}
					πTemp001[0] = µmicrosecond
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_check_int_field); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µmicrosecond = πTemp003
					if πE = πg.CheckLocal(πF, µhour, "hour"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.LE(πF, πg.NewInt(0).ToObject(), µhour); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if !πTemp004 {
						goto Label1
					}
					if πTemp003, πE = πg.LE(πF, µhour, πg.NewInt(23).ToObject()); πE != nil {
						continue
					}
				Label1:
					if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(!πTemp004).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label2
					}
					goto Label3
					// line 325: if not 0 <= hour <= 23:
					πF.SetLineno(325)
				Label2:
					πTemp001 = πF.MakeArgs(2)
					πTemp001[0] = πg.NewStr("hour must be in 0..23").ToObject()
					if πE = πg.CheckLocal(πF, µhour, "hour"); πE != nil {
						continue
					}
					πTemp001[1] = µhour
					if πTemp002, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 326: raise ValueError('hour must be in 0..23', hour)
					πF.SetLineno(326)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					goto Label3
				Label3:
					if πE = πg.CheckLocal(πF, µminute, "minute"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.LE(πF, πg.NewInt(0).ToObject(), µminute); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if !πTemp004 {
						goto Label4
					}
					if πTemp003, πE = πg.LE(πF, µminute, πg.NewInt(59).ToObject()); πE != nil {
						continue
					}
				Label4:
					if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(!πTemp004).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label5
					}
					goto Label6
					// line 327: if not 0 <= minute <= 59:
					πF.SetLineno(327)
				Label5:
					πTemp001 = πF.MakeArgs(2)
					πTemp001[0] = πg.NewStr("minute must be in 0..59").ToObject()
					if πE = πg.CheckLocal(πF, µminute, "minute"); πE != nil {
						continue
					}
					πTemp001[1] = µminute
					if πTemp002, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 328: raise ValueError('minute must be in 0..59', minute)
					πF.SetLineno(328)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					goto Label6
				Label6:
					if πE = πg.CheckLocal(πF, µsecond, "second"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.LE(πF, πg.NewInt(0).ToObject(), µsecond); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if !πTemp004 {
						goto Label7
					}
					if πTemp003, πE = πg.LE(πF, µsecond, πg.NewInt(59).ToObject()); πE != nil {
						continue
					}
				Label7:
					if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(!πTemp004).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label8
					}
					goto Label9
					// line 329: if not 0 <= second <= 59:
					πF.SetLineno(329)
				Label8:
					πTemp001 = πF.MakeArgs(2)
					πTemp001[0] = πg.NewStr("second must be in 0..59").ToObject()
					if πE = πg.CheckLocal(πF, µsecond, "second"); πE != nil {
						continue
					}
					πTemp001[1] = µsecond
					if πTemp002, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 330: raise ValueError('second must be in 0..59', second)
					πF.SetLineno(330)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					goto Label9
				Label9:
					if πE = πg.CheckLocal(πF, µmicrosecond, "microsecond"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.LE(πF, πg.NewInt(0).ToObject(), µmicrosecond); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if !πTemp004 {
						goto Label10
					}
					if πTemp003, πE = πg.LE(πF, µmicrosecond, πg.NewInt(999999).ToObject()); πE != nil {
						continue
					}
				Label10:
					if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(!πTemp004).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label11
					}
					goto Label12
					// line 331: if not 0 <= microsecond <= 999999:
					πF.SetLineno(331)
				Label11:
					πTemp001 = πF.MakeArgs(2)
					πTemp001[0] = πg.NewStr("microsecond must be in 0..999999").ToObject()
					if πE = πg.CheckLocal(πF, µmicrosecond, "microsecond"); πE != nil {
						continue
					}
					πTemp001[1] = µmicrosecond
					if πTemp002, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 332: raise ValueError('microsecond must be in 0..999999', microsecond)
					πF.SetLineno(332)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					goto Label12
				Label12:
					// line 333: return hour, minute, second, microsecond
					πF.SetLineno(333)
					if πE = πg.CheckLocal(πF, µhour, "hour"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µminute, "minute"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsecond, "second"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µmicrosecond, "microsecond"); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple4(µhour, µminute, µsecond, µmicrosecond).ToObject()
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
			if πE = πF.Globals().SetItem(πF, ß_check_time_fields.ToObject(), πTemp020); πE != nil {
				continue
			}
			// line 335: def _check_tzinfo_arg(tz):
			πF.SetLineno(335)
			πTemp003 = make([]πg.Param, 1)
			πTemp003[0] = πg.Param{Name: "tz", Def: nil}
			πTemp021 = πg.NewFunction(πg.NewCode("_check_tzinfo_arg", "/usr/lib/python2.7/datetime.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µtz *πg.Object = πArgs[0]; _ = µtz
				var πTemp001 *πg.Object
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
				var πTemp007 bool
				_ = πTemp007
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					if πE = πg.CheckLocal(πF, µtz, "tz"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(µtz != πTemp004).ToObject()
					πTemp001 = πTemp003
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp002 {
						goto Label1
					}
					πTemp005 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µtz, "tz"); πE != nil {
						continue
					}
					πTemp005[0] = µtz
					if πTemp004, πE = πg.ResolveGlobal(πF, ßtzinfo); πE != nil {
						continue
					}
					πTemp005[1] = πTemp004
					if πTemp004, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πTemp007, πE = πg.IsTrue(πF, πTemp006); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(!πTemp007).ToObject()
					πTemp001 = πTemp003
				Label1:
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label2
					}
					goto Label3
					// line 336: if tz is not None and not isinstance(tz, tzinfo):
					πF.SetLineno(336)
				Label2:
					πTemp005 = πF.MakeArgs(1)
					πTemp005[0] = πg.NewStr("tzinfo argument must be None or of a tzinfo subclass").ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					// line 337: raise TypeError("tzinfo argument must be None or of a tzinfo subclass")
					πF.SetLineno(337)
					πE = πF.Raise(πTemp003, nil, nil)
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
			if πE = πF.Globals().SetItem(πF, ß_check_tzinfo_arg.ToObject(), πTemp021); πE != nil {
				continue
			}
			// line 363: def _cmperror(x, y):
			πF.SetLineno(363)
			πTemp003 = make([]πg.Param, 2)
			πTemp003[0] = πg.Param{Name: "x", Def: nil}
			πTemp003[1] = πg.Param{Name: "y", Def: nil}
			πTemp022 = πg.NewFunction(πg.NewCode("_cmperror", "/usr/lib/python2.7/datetime.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πArgs[0]; _ = µx
				var µy *πg.Object = πArgs[1]; _ = µy
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
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
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					πTemp001 = πF.MakeArgs(1)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp004[0] = µx
					if πTemp005, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp005, πE = πg.GetAttr(πF, πTemp006, ß__name__, nil); πE != nil {
						continue
					}
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
						continue
					}
					πTemp004[0] = µy
					if πTemp006, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp006, πE = πg.GetAttr(πF, πTemp007, ß__name__, nil); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple2(πTemp005, πTemp006).ToObject()
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("can't compare '%s' to '%s'").ToObject(), πTemp003); πE != nil {
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
					// line 364: raise TypeError("can't compare '%s' to '%s'" % (
					πF.SetLineno(364)
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
			if πE = πF.Globals().SetItem(πF, ß_cmperror.ToObject(), πTemp022); πE != nil {
				continue
			}
			// line 367: def _normalize_pair(hi, lo, factor):
			πF.SetLineno(367)
			πTemp003 = make([]πg.Param, 3)
			πTemp003[0] = πg.Param{Name: "hi", Def: nil}
			πTemp003[1] = πg.Param{Name: "lo", Def: nil}
			πTemp003[2] = πg.Param{Name: "factor", Def: nil}
			πTemp023 = πg.NewFunction(πg.NewCode("_normalize_pair", "/usr/lib/python2.7/datetime.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µhi *πg.Object = πArgs[0]; _ = µhi
				var µlo *πg.Object = πArgs[1]; _ = µlo
				var µfactor *πg.Object = πArgs[2]; _ = µfactor
				var µinc *πg.Object = πg.UnboundLocal; _ = µinc
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
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					if πE = πg.CheckLocal(πF, µlo, "lo"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.LE(πF, πg.NewInt(0).ToObject(), µlo); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp003 {
						goto Label1
					}
					if πE = πg.CheckLocal(πF, µfactor, "factor"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Sub(πF, µfactor, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.LE(πF, µlo, πTemp004); πE != nil {
						continue
					}
				Label1:
					if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(!πTemp003).ToObject()
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label2
					}
					goto Label3
					// line 368: if not 0 <= lo <= factor-1:
					πF.SetLineno(368)
				Label2:
					// line 369: inc, lo = divmod(lo, factor)
					πF.SetLineno(369)
					πTemp005 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µlo, "lo"); πE != nil {
						continue
					}
					πTemp005[0] = µlo
					if πE = πg.CheckLocal(πF, µfactor, "factor"); πE != nil {
						continue
					}
					πTemp005[1] = µfactor
					if πTemp001, πE = πg.ResolveGlobal(πF, ßdivmod); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp004}}}, πTemp002); πE != nil {
						continue
					}
					µinc = πTemp001
					µlo = πTemp004
					// line 370: hi += inc
					πF.SetLineno(370)
					if πE = πg.CheckLocal(πF, µhi, "hi"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µinc, "inc"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.IAdd(πF, µhi, µinc); πE != nil {
						continue
					}
					µhi = πTemp001
					goto Label3
				Label3:
					// line 371: return hi, lo
					πF.SetLineno(371)
					if πE = πg.CheckLocal(πF, µhi, "hi"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µlo, "lo"); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple2(µhi, µlo).ToObject()
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
			if πE = πF.Globals().SetItem(πF, ß_normalize_pair.ToObject(), πTemp023); πE != nil {
				continue
			}
			// line 373: def _normalize_datetime(y, m, d, hh, mm, ss, us, ignore_overflow=False):
			πF.SetLineno(373)
			πTemp003 = make([]πg.Param, 8)
			πTemp003[0] = πg.Param{Name: "y", Def: nil}
			πTemp003[1] = πg.Param{Name: "m", Def: nil}
			πTemp003[2] = πg.Param{Name: "d", Def: nil}
			πTemp003[3] = πg.Param{Name: "hh", Def: nil}
			πTemp003[4] = πg.Param{Name: "mm", Def: nil}
			πTemp003[5] = πg.Param{Name: "ss", Def: nil}
			πTemp003[6] = πg.Param{Name: "us", Def: nil}
			if πTemp025, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			πTemp003[7] = πg.Param{Name: "ignore_overflow", Def: πTemp025}
			πTemp024 = πg.NewFunction(πg.NewCode("_normalize_datetime", "/usr/lib/python2.7/datetime.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µy *πg.Object = πArgs[0]; _ = µy
				var µm *πg.Object = πArgs[1]; _ = µm
				var µd *πg.Object = πArgs[2]; _ = µd
				var µhh *πg.Object = πArgs[3]; _ = µhh
				var µmm *πg.Object = πArgs[4]; _ = µmm
				var µss *πg.Object = πArgs[5]; _ = µss
				var µus *πg.Object = πArgs[6]; _ = µus
				var µignore_overflow *πg.Object = πArgs[7]; _ = µignore_overflow
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
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 375: ss, us = _normalize_pair(ss, us, 1000000)
					πF.SetLineno(375)
					πTemp001 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp001[0] = µss
					if πE = πg.CheckLocal(πF, µus, "us"); πE != nil {
						continue
					}
					πTemp001[1] = µus
					πTemp001[2] = πg.NewInt(1000000).ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_normalize_pair); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp004}}}, πTemp003); πE != nil {
						continue
					}
					µss = πTemp002
					µus = πTemp004
					// line 376: mm, ss = _normalize_pair(mm, ss, 60)
					πF.SetLineno(376)
					πTemp001 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µmm, "mm"); πE != nil {
						continue
					}
					πTemp001[0] = µmm
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp001[1] = µss
					πTemp001[2] = πg.NewInt(60).ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_normalize_pair); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp004}}}, πTemp003); πE != nil {
						continue
					}
					µmm = πTemp002
					µss = πTemp004
					// line 377: hh, mm = _normalize_pair(hh, mm, 60)
					πF.SetLineno(377)
					πTemp001 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µhh, "hh"); πE != nil {
						continue
					}
					πTemp001[0] = µhh
					if πE = πg.CheckLocal(πF, µmm, "mm"); πE != nil {
						continue
					}
					πTemp001[1] = µmm
					πTemp001[2] = πg.NewInt(60).ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_normalize_pair); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp004}}}, πTemp003); πE != nil {
						continue
					}
					µhh = πTemp002
					µmm = πTemp004
					// line 378: d, hh = _normalize_pair(d, hh, 24)
					πF.SetLineno(378)
					πTemp001 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
						continue
					}
					πTemp001[0] = µd
					if πE = πg.CheckLocal(πF, µhh, "hh"); πE != nil {
						continue
					}
					πTemp001[1] = µhh
					πTemp001[2] = πg.NewInt(24).ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_normalize_pair); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp004}}}, πTemp003); πE != nil {
						continue
					}
					µd = πTemp002
					µhh = πTemp004
					// line 379: y, m, d = _normalize_date(y, m, d, ignore_overflow)
					πF.SetLineno(379)
					πTemp001 = πF.MakeArgs(4)
					if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
						continue
					}
					πTemp001[0] = µy
					if πE = πg.CheckLocal(πF, µm, "m"); πE != nil {
						continue
					}
					πTemp001[1] = µm
					if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
						continue
					}
					πTemp001[2] = µd
					if πE = πg.CheckLocal(πF, µignore_overflow, "ignore_overflow"); πE != nil {
						continue
					}
					πTemp001[3] = µignore_overflow
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_normalize_date); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}}}, πTemp003); πE != nil {
						continue
					}
					µy = πTemp002
					µm = πTemp004
					µd = πTemp005
					// line 380: return y, m, d, hh, mm, ss, us
					πF.SetLineno(380)
					πTemp001 = make([]*πg.Object, 7)
					if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
						continue
					}
					πTemp001[0] = µy
					if πE = πg.CheckLocal(πF, µm, "m"); πE != nil {
						continue
					}
					πTemp001[1] = µm
					if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
						continue
					}
					πTemp001[2] = µd
					if πE = πg.CheckLocal(πF, µhh, "hh"); πE != nil {
						continue
					}
					πTemp001[3] = µhh
					if πE = πg.CheckLocal(πF, µmm, "mm"); πE != nil {
						continue
					}
					πTemp001[4] = µmm
					if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
						continue
					}
					πTemp001[5] = µss
					if πE = πg.CheckLocal(πF, µus, "us"); πE != nil {
						continue
					}
					πTemp001[6] = µus
					πTemp002 = πg.NewTuple(πTemp001...).ToObject()
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
			if πE = πF.Globals().SetItem(πF, ß_normalize_datetime.ToObject(), πTemp024); πE != nil {
				continue
			}
			// line 382: def _normalize_date(year, month, day, ignore_overflow=False):
			πF.SetLineno(382)
			πTemp003 = make([]πg.Param, 4)
			πTemp003[0] = πg.Param{Name: "year", Def: nil}
			πTemp003[1] = πg.Param{Name: "month", Def: nil}
			πTemp003[2] = πg.Param{Name: "day", Def: nil}
			if πTemp026, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			πTemp003[3] = πg.Param{Name: "ignore_overflow", Def: πTemp026}
			πTemp025 = πg.NewFunction(πg.NewCode("_normalize_date", "/usr/lib/python2.7/datetime.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µyear *πg.Object = πArgs[0]; _ = µyear
				var µmonth *πg.Object = πArgs[1]; _ = µmonth
				var µday *πg.Object = πArgs[2]; _ = µday
				var µignore_overflow *πg.Object = πArgs[3]; _ = µignore_overflow
				var µdim *πg.Object = πg.UnboundLocal; _ = µdim
				var µordinal *πg.Object = πg.UnboundLocal; _ = µordinal
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
				var πTemp007 bool
				_ = πTemp007
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					if πE = πg.CheckLocal(πF, µmonth, "month"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.LE(πF, πg.NewInt(1).ToObject(), µmonth); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp003 {
						goto Label1
					}
					if πTemp002, πE = πg.LE(πF, µmonth, πg.NewInt(12).ToObject()); πE != nil {
						continue
					}
				Label1:
					if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(!πTemp003).ToObject()
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label2
					}
					goto Label3
					// line 389: if not 1 <= month <= 12:
					πF.SetLineno(389)
				Label2:
					// line 390: year, month = _normalize_pair(year, month-1, 12)
					πF.SetLineno(390)
					πTemp004 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µyear, "year"); πE != nil {
						continue
					}
					πTemp004[0] = µyear
					if πE = πg.CheckLocal(πF, µmonth, "month"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Sub(πF, µmonth, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp004[1] = πTemp001
					πTemp004[2] = πg.NewInt(12).ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_normalize_pair); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp005}}}, πTemp002); πE != nil {
						continue
					}
					µyear = πTemp001
					µmonth = πTemp005
					// line 391: month += 1
					πF.SetLineno(391)
					if πE = πg.CheckLocal(πF, µmonth, "month"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.IAdd(πF, µmonth, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µmonth = πTemp001
					// line 392: assert 1 <= month <= 12
					πF.SetLineno(392)
					if πE = πg.CheckLocal(πF, µmonth, "month"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.LE(πF, πg.NewInt(1).ToObject(), µmonth); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp003 {
						goto Label4
					}
					if πTemp001, πE = πg.LE(πF, µmonth, πg.NewInt(12).ToObject()); πE != nil {
						continue
					}
				Label4:
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
						continue
					}
					goto Label3
				Label3:
					// line 398: dim = _days_in_month(year, month)
					πF.SetLineno(398)
					πTemp004 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µyear, "year"); πE != nil {
						continue
					}
					πTemp004[0] = µyear
					if πE = πg.CheckLocal(πF, µmonth, "month"); πE != nil {
						continue
					}
					πTemp004[1] = µmonth
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_days_in_month); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µdim = πTemp002
					if πE = πg.CheckLocal(πF, µday, "day"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.LE(πF, πg.NewInt(1).ToObject(), µday); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp003 {
						goto Label5
					}
					if πE = πg.CheckLocal(πF, µdim, "dim"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.LE(πF, µday, µdim); πE != nil {
						continue
					}
				Label5:
					if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(!πTemp003).ToObject()
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label6
					}
					goto Label7
					// line 399: if not 1 <= day <= dim:
					πF.SetLineno(399)
				Label6:
					if πE = πg.CheckLocal(πF, µday, "day"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, µday, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label8
					}
					if πE = πg.CheckLocal(πF, µday, "day"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µdim, "dim"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Add(πF, µdim, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, µday, πTemp002); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label9
					}
					goto Label10
					// line 403: if day == 0:    # move back a day
					πF.SetLineno(403)
				Label8:
					// line 404: month -= 1
					πF.SetLineno(404)
					if πE = πg.CheckLocal(πF, µmonth, "month"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.ISub(πF, µmonth, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µmonth = πTemp001
					if πE = πg.CheckLocal(πF, µmonth, "month"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GT(πF, µmonth, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label12
					}
					goto Label13
					// line 405: if month > 0:
					πF.SetLineno(405)
				Label12:
					// line 406: day = _days_in_month(year, month)
					πF.SetLineno(406)
					πTemp004 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µyear, "year"); πE != nil {
						continue
					}
					πTemp004[0] = µyear
					if πE = πg.CheckLocal(πF, µmonth, "month"); πE != nil {
						continue
					}
					πTemp004[1] = µmonth
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_days_in_month); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µday = πTemp002
					goto Label14
				Label13:
					// line 408: year, month, day = year-1, 12, 31
					πF.SetLineno(408)
					if πE = πg.CheckLocal(πF, µyear, "year"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Sub(πF, µyear, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple3(πTemp002, πg.NewInt(12).ToObject(), πg.NewInt(31).ToObject()).ToObject()
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp006}}}, πTemp001); πE != nil {
						continue
					}
					µyear = πTemp002
					µmonth = πTemp005
					µday = πTemp006
					goto Label14
				Label14:
					goto Label11
					// line 409: elif day == dim + 1:    # move forward a day
					πF.SetLineno(409)
				Label9:
					// line 410: month += 1
					πF.SetLineno(410)
					if πE = πg.CheckLocal(πF, µmonth, "month"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.IAdd(πF, µmonth, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µmonth = πTemp001
					// line 411: day = 1
					πF.SetLineno(411)
					µday = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µmonth, "month"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GT(πF, µmonth, πg.NewInt(12).ToObject()); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label15
					}
					goto Label16
					// line 412: if month > 12:
					πF.SetLineno(412)
				Label15:
					// line 413: month = 1
					πF.SetLineno(413)
					µmonth = πg.NewInt(1).ToObject()
					// line 414: year += 1
					πF.SetLineno(414)
					if πE = πg.CheckLocal(πF, µyear, "year"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.IAdd(πF, µyear, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µyear = πTemp001
					goto Label16
				Label16:
					goto Label11
				Label10:
					// line 416: ordinal = _ymd2ord(year, month, 1) + (day - 1)
					πF.SetLineno(416)
					πTemp004 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µyear, "year"); πE != nil {
						continue
					}
					πTemp004[0] = µyear
					if πE = πg.CheckLocal(πF, µmonth, "month"); πE != nil {
						continue
					}
					πTemp004[1] = µmonth
					πTemp004[2] = πg.NewInt(1).ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_ymd2ord); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πE = πg.CheckLocal(πF, µday, "day"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Sub(πF, µday, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, πTemp005, πTemp002); πE != nil {
						continue
					}
					µordinal = πTemp001
					// line 417: year, month, day = _ord2ymd(ordinal)
					πF.SetLineno(417)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µordinal, "ordinal"); πE != nil {
						continue
					}
					πTemp004[0] = µordinal
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_ord2ymd); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp006}}}, πTemp002); πE != nil {
						continue
					}
					µyear = πTemp001
					µmonth = πTemp005
					µday = πTemp006
					goto Label11
				Label11:
					goto Label7
				Label7:
					if πE = πg.CheckLocal(πF, µignore_overflow, "ignore_overflow"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsTrue(πF, µignore_overflow); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(!πTemp007).ToObject()
					πTemp001 = πTemp002
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp003 {
						goto Label17
					}
					if πTemp006, πE = πg.ResolveGlobal(πF, ßMINYEAR); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µyear, "year"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.LE(πF, πTemp006, µyear); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					if !πTemp007 {
						goto Label18
					}
					if πTemp006, πE = πg.ResolveGlobal(πF, ßMAXYEAR); πE != nil {
						continue
					}
					if πTemp005, πE = πg.LE(πF, µyear, πTemp006); πE != nil {
						continue
					}
				Label18:
					if πTemp007, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(!πTemp007).ToObject()
					πTemp001 = πTemp002
				Label17:
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label19
					}
					goto Label20
					// line 419: if not ignore_overflow and not MINYEAR <= year <= MAXYEAR:
					πF.SetLineno(419)
				Label19:
					πTemp004 = πF.MakeArgs(1)
					πTemp004[0] = πg.NewStr("date value out of range").ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßOverflowError); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 420: raise OverflowError("date value out of range")
					πF.SetLineno(420)
					πE = πF.Raise(πTemp002, nil, nil)
					continue
					goto Label20
				Label20:
					// line 421: return year, month, day
					πF.SetLineno(421)
					if πE = πg.CheckLocal(πF, µyear, "year"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µmonth, "month"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µday, "day"); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple3(µyear, µmonth, µday).ToObject()
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
			if πE = πF.Globals().SetItem(πF, ß_normalize_date.ToObject(), πTemp025); πE != nil {
				continue
			}
			// line 423: def _accum(tag, sofar, num, factor, leftover):
			πF.SetLineno(423)
			πTemp003 = make([]πg.Param, 5)
			πTemp003[0] = πg.Param{Name: "tag", Def: nil}
			πTemp003[1] = πg.Param{Name: "sofar", Def: nil}
			πTemp003[2] = πg.Param{Name: "num", Def: nil}
			πTemp003[3] = πg.Param{Name: "factor", Def: nil}
			πTemp003[4] = πg.Param{Name: "leftover", Def: nil}
			πTemp026 = πg.NewFunction(πg.NewCode("_accum", "/usr/lib/python2.7/datetime.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µtag *πg.Object = πArgs[0]; _ = µtag
				var µsofar *πg.Object = πArgs[1]; _ = µsofar
				var µnum *πg.Object = πArgs[2]; _ = µnum
				var µfactor *πg.Object = πArgs[3]; _ = µfactor
				var µleftover *πg.Object = πArgs[4]; _ = µleftover
				var µprod *πg.Object = πg.UnboundLocal; _ = µprod
				var µrsum *πg.Object = πg.UnboundLocal; _ = µrsum
				var µfracpart *πg.Object = πg.UnboundLocal; _ = µfracpart
				var µintpart *πg.Object = πg.UnboundLocal; _ = µintpart
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
				var πTemp006 []*πg.Object
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
					if πE = πg.CheckLocal(πF, µnum, "num"); πE != nil {
						continue
					}
					πTemp001[0] = µnum
					if πTemp003, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßlong); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
					πTemp001[1] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label1
					}
					goto Label2
					// line 424: if isinstance(num, (int, long)):
					πF.SetLineno(424)
				Label1:
					// line 425: prod = num * factor
					πF.SetLineno(425)
					if πE = πg.CheckLocal(πF, µnum, "num"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µfactor, "factor"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mul(πF, µnum, µfactor); πE != nil {
						continue
					}
					µprod = πTemp002
					// line 426: rsum = sofar + prod
					πF.SetLineno(426)
					if πE = πg.CheckLocal(πF, µsofar, "sofar"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µprod, "prod"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Add(πF, µsofar, µprod); πE != nil {
						continue
					}
					µrsum = πTemp002
					// line 427: return rsum, leftover
					πF.SetLineno(427)
					if πE = πg.CheckLocal(πF, µrsum, "rsum"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µleftover, "leftover"); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(µrsum, µleftover).ToObject()
					πR = πTemp002
					continue
					goto Label2
				Label2:
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µnum, "num"); πE != nil {
						continue
					}
					πTemp001[0] = µnum
					if πTemp002, πE = πg.ResolveGlobal(πF, ßfloat); πE != nil {
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
					if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label3
					}
					goto Label4
					// line 428: if isinstance(num, float):
					πF.SetLineno(428)
				Label3:
					// line 429: fracpart, intpart = _math.modf(num)
					πF.SetLineno(429)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µnum, "num"); πE != nil {
						continue
					}
					πTemp001[0] = µnum
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_math); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßmodf, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp004}}}, πTemp002); πE != nil {
						continue
					}
					µfracpart = πTemp003
					µintpart = πTemp004
					// line 430: prod = int(intpart) * factor
					πF.SetLineno(430)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µintpart, "intpart"); πE != nil {
						continue
					}
					πTemp001[0] = µintpart
					if πTemp003, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.CheckLocal(πF, µfactor, "factor"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mul(πF, πTemp004, µfactor); πE != nil {
						continue
					}
					µprod = πTemp002
					// line 431: rsum = sofar + prod
					πF.SetLineno(431)
					if πE = πg.CheckLocal(πF, µsofar, "sofar"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µprod, "prod"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Add(πF, µsofar, µprod); πE != nil {
						continue
					}
					µrsum = πTemp002
					if πE = πg.CheckLocal(πF, µfracpart, "fracpart"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µfracpart, πg.NewFloat(0.0).ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label5
					}
					goto Label6
					// line 432: if fracpart == 0.0:
					πF.SetLineno(432)
				Label5:
					// line 433: return rsum, leftover
					πF.SetLineno(433)
					if πE = πg.CheckLocal(πF, µrsum, "rsum"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µleftover, "leftover"); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(µrsum, µleftover).ToObject()
					πR = πTemp002
					continue
					goto Label6
				Label6:
					// line 434: assert isinstance(factor, (int, long))
					πF.SetLineno(434)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µfactor, "factor"); πE != nil {
						continue
					}
					πTemp001[0] = µfactor
					if πTemp003, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßlong); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
					πTemp001[1] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Assert(πF, πTemp003, nil); πE != nil {
						continue
					}
					// line 435: fracpart, intpart = _math.modf(factor * fracpart)
					πF.SetLineno(435)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfactor, "factor"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µfracpart, "fracpart"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mul(πF, µfactor, µfracpart); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_math); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßmodf, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp004}}}, πTemp002); πE != nil {
						continue
					}
					µfracpart = πTemp003
					µintpart = πTemp004
					// line 436: rsum += int(intpart)
					πF.SetLineno(436)
					if πE = πg.CheckLocal(πF, µrsum, "rsum"); πE != nil {
						continue
					}
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µintpart, "intpart"); πE != nil {
						continue
					}
					πTemp001[0] = µintpart
					if πTemp002, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp002, πE = πg.IAdd(πF, µrsum, πTemp003); πE != nil {
						continue
					}
					µrsum = πTemp002
					// line 437: return rsum, leftover + fracpart
					πF.SetLineno(437)
					if πE = πg.CheckLocal(πF, µrsum, "rsum"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µleftover, "leftover"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µfracpart, "fracpart"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Add(πF, µleftover, µfracpart); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(µrsum, πTemp003).ToObject()
					πR = πTemp002
					continue
					goto Label4
				Label4:
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µtag, "tag"); πE != nil {
						continue
					}
					πTemp006 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µnum, "num"); πE != nil {
						continue
					}
					πTemp006[0] = µnum
					if πTemp004, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					πTemp003 = πg.NewTuple2(µtag, πTemp007).ToObject()
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("unsupported type for timedelta %s component: %s").ToObject(), πTemp003); πE != nil {
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
					// line 438: raise TypeError("unsupported type for timedelta %s component: %s" %
					πF.SetLineno(438)
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
			if πE = πF.Globals().SetItem(πF, ß_accum.ToObject(), πTemp026); πE != nil {
				continue
			}
			// line 441: class timedelta(object):
			πF.SetLineno(441)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp030, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp030
			πTemp027 = πg.NewDict()
			if πTemp028, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp027.SetItem(πF, ß__module__.ToObject(), πTemp028); πE != nil {
				continue
			}
			_, πE = πg.NewCode("timedelta", "/usr/lib/python2.7/datetime.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp027
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []πg.Param
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
				var πTemp030 *πg.Object
				_ = πTemp030
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 442: """Represent the difference between two datetime objects.
					πF.SetLineno(442)
					// line 442: """Represent the difference between two datetime objects.
					πF.SetLineno(442)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Represent the difference between two datetime objects.\n\n    Supported operators:\n\n    - add, subtract timedelta\n    - unary plus, minus, abs\n    - compare to timedelta\n    - multiply, divide by int/long\n\n    In addition, datetime supports subtraction of two datetime objects\n    returning a timedelta, and addition or subtraction of a datetime\n    and a timedelta giving a datetime.\n\n    Representation: (days, seconds, microseconds).  Why?  Because I\n    felt like it.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 458: __slots__ = '_days', '_seconds', '_microseconds', '_hashcode'
					πF.SetLineno(458)
					πTemp001 = πg.NewTuple4(ß_days.ToObject(), ß_seconds.ToObject(), ß_microseconds.ToObject(), ß_hashcode.ToObject()).ToObject()
					if πE = πClass.SetItem(πF, ß__slots__.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 460: def __new__(cls, days=_SENTINEL, seconds=_SENTINEL, microseconds=_SENTINEL,
					πF.SetLineno(460)
					πTemp002 = make([]πg.Param, 8)
					πTemp002[0] = πg.Param{Name: "cls", Def: nil}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ß_SENTINEL); πE != nil {
						continue
					}
					πTemp002[1] = πg.Param{Name: "days", Def: πTemp003}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ß_SENTINEL); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "seconds", Def: πTemp003}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ß_SENTINEL); πE != nil {
						continue
					}
					πTemp002[3] = πg.Param{Name: "microseconds", Def: πTemp003}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ß_SENTINEL); πE != nil {
						continue
					}
					πTemp002[4] = πg.Param{Name: "milliseconds", Def: πTemp003}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ß_SENTINEL); πE != nil {
						continue
					}
					πTemp002[5] = πg.Param{Name: "minutes", Def: πTemp003}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ß_SENTINEL); πE != nil {
						continue
					}
					πTemp002[6] = πg.Param{Name: "hours", Def: πTemp003}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ß_SENTINEL); πE != nil {
						continue
					}
					πTemp002[7] = πg.Param{Name: "weeks", Def: πTemp003}
					πTemp001 = πg.NewFunction(πg.NewCode("__new__", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µcls *πg.Object = πArgs[0]; _ = µcls
						var µdays *πg.Object = πArgs[1]; _ = µdays
						var µseconds *πg.Object = πArgs[2]; _ = µseconds
						var µmicroseconds *πg.Object = πArgs[3]; _ = µmicroseconds
						var µmilliseconds *πg.Object = πArgs[4]; _ = µmilliseconds
						var µminutes *πg.Object = πArgs[5]; _ = µminutes
						var µhours *πg.Object = πArgs[6]; _ = µhours
						var µweeks *πg.Object = πArgs[7]; _ = µweeks
						var µx *πg.Object = πg.UnboundLocal; _ = µx
						var µleftover *πg.Object = πg.UnboundLocal; _ = µleftover
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
							// line 462: x = 0
							πF.SetLineno(462)
							µx = πg.NewInt(0).ToObject()
							// line 463: leftover = 0.0
							πF.SetLineno(463)
							µleftover = πg.NewFloat(0.0).ToObject()
							if πE = πg.CheckLocal(πF, µmicroseconds, "microseconds"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_SENTINEL); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µmicroseconds != πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 464: if microseconds is not _SENTINEL:
							πF.SetLineno(464)
						Label1:
							// line 465: x, leftover = _accum("microseconds", x, microseconds, _US_PER_US, leftover)
							πF.SetLineno(465)
							πTemp004 = πF.MakeArgs(5)
							πTemp004[0] = ßmicroseconds.ToObject()
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp004[1] = µx
							if πE = πg.CheckLocal(πF, µmicroseconds, "microseconds"); πE != nil {
								continue
							}
							πTemp004[2] = µmicroseconds
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_US_PER_US); πE != nil {
								continue
							}
							πTemp004[3] = πTemp001
							if πE = πg.CheckLocal(πF, µleftover, "leftover"); πE != nil {
								continue
							}
							πTemp004[4] = µleftover
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_accum); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp005}}}, πTemp002); πE != nil {
								continue
							}
							µx = πTemp001
							µleftover = πTemp005
							goto Label2
						Label2:
							if πE = πg.CheckLocal(πF, µmilliseconds, "milliseconds"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_SENTINEL); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µmilliseconds != πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label3
							}
							goto Label4
							// line 466: if milliseconds is not _SENTINEL:
							πF.SetLineno(466)
						Label3:
							// line 467: x, leftover = _accum("milliseconds", x, milliseconds, _US_PER_MS, leftover)
							πF.SetLineno(467)
							πTemp004 = πF.MakeArgs(5)
							πTemp004[0] = ßmilliseconds.ToObject()
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp004[1] = µx
							if πE = πg.CheckLocal(πF, µmilliseconds, "milliseconds"); πE != nil {
								continue
							}
							πTemp004[2] = µmilliseconds
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_US_PER_MS); πE != nil {
								continue
							}
							πTemp004[3] = πTemp001
							if πE = πg.CheckLocal(πF, µleftover, "leftover"); πE != nil {
								continue
							}
							πTemp004[4] = µleftover
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_accum); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp005}}}, πTemp002); πE != nil {
								continue
							}
							µx = πTemp001
							µleftover = πTemp005
							goto Label4
						Label4:
							if πE = πg.CheckLocal(πF, µseconds, "seconds"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_SENTINEL); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µseconds != πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label5
							}
							goto Label6
							// line 468: if seconds is not _SENTINEL:
							πF.SetLineno(468)
						Label5:
							// line 469: x, leftover = _accum("seconds", x, seconds, _US_PER_SECOND, leftover)
							πF.SetLineno(469)
							πTemp004 = πF.MakeArgs(5)
							πTemp004[0] = ßseconds.ToObject()
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp004[1] = µx
							if πE = πg.CheckLocal(πF, µseconds, "seconds"); πE != nil {
								continue
							}
							πTemp004[2] = µseconds
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_US_PER_SECOND); πE != nil {
								continue
							}
							πTemp004[3] = πTemp001
							if πE = πg.CheckLocal(πF, µleftover, "leftover"); πE != nil {
								continue
							}
							πTemp004[4] = µleftover
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_accum); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp005}}}, πTemp002); πE != nil {
								continue
							}
							µx = πTemp001
							µleftover = πTemp005
							goto Label6
						Label6:
							if πE = πg.CheckLocal(πF, µminutes, "minutes"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_SENTINEL); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µminutes != πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label7
							}
							goto Label8
							// line 470: if minutes is not _SENTINEL:
							πF.SetLineno(470)
						Label7:
							// line 471: x, leftover = _accum("minutes", x, minutes, _US_PER_MINUTE, leftover)
							πF.SetLineno(471)
							πTemp004 = πF.MakeArgs(5)
							πTemp004[0] = ßminutes.ToObject()
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp004[1] = µx
							if πE = πg.CheckLocal(πF, µminutes, "minutes"); πE != nil {
								continue
							}
							πTemp004[2] = µminutes
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_US_PER_MINUTE); πE != nil {
								continue
							}
							πTemp004[3] = πTemp001
							if πE = πg.CheckLocal(πF, µleftover, "leftover"); πE != nil {
								continue
							}
							πTemp004[4] = µleftover
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_accum); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp005}}}, πTemp002); πE != nil {
								continue
							}
							µx = πTemp001
							µleftover = πTemp005
							goto Label8
						Label8:
							if πE = πg.CheckLocal(πF, µhours, "hours"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_SENTINEL); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µhours != πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label9
							}
							goto Label10
							// line 472: if hours is not _SENTINEL:
							πF.SetLineno(472)
						Label9:
							// line 473: x, leftover = _accum("hours", x, hours, _US_PER_HOUR, leftover)
							πF.SetLineno(473)
							πTemp004 = πF.MakeArgs(5)
							πTemp004[0] = ßhours.ToObject()
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp004[1] = µx
							if πE = πg.CheckLocal(πF, µhours, "hours"); πE != nil {
								continue
							}
							πTemp004[2] = µhours
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_US_PER_HOUR); πE != nil {
								continue
							}
							πTemp004[3] = πTemp001
							if πE = πg.CheckLocal(πF, µleftover, "leftover"); πE != nil {
								continue
							}
							πTemp004[4] = µleftover
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_accum); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp005}}}, πTemp002); πE != nil {
								continue
							}
							µx = πTemp001
							µleftover = πTemp005
							goto Label10
						Label10:
							if πE = πg.CheckLocal(πF, µdays, "days"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_SENTINEL); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µdays != πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label11
							}
							goto Label12
							// line 474: if days is not _SENTINEL:
							πF.SetLineno(474)
						Label11:
							// line 475: x, leftover = _accum("days", x, days, _US_PER_DAY, leftover)
							πF.SetLineno(475)
							πTemp004 = πF.MakeArgs(5)
							πTemp004[0] = ßdays.ToObject()
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp004[1] = µx
							if πE = πg.CheckLocal(πF, µdays, "days"); πE != nil {
								continue
							}
							πTemp004[2] = µdays
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_US_PER_DAY); πE != nil {
								continue
							}
							πTemp004[3] = πTemp001
							if πE = πg.CheckLocal(πF, µleftover, "leftover"); πE != nil {
								continue
							}
							πTemp004[4] = µleftover
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_accum); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp005}}}, πTemp002); πE != nil {
								continue
							}
							µx = πTemp001
							µleftover = πTemp005
							goto Label12
						Label12:
							if πE = πg.CheckLocal(πF, µweeks, "weeks"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_SENTINEL); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µweeks != πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label13
							}
							goto Label14
							// line 476: if weeks is not _SENTINEL:
							πF.SetLineno(476)
						Label13:
							// line 477: x, leftover = _accum("weeks", x, weeks, _US_PER_WEEK, leftover)
							πF.SetLineno(477)
							πTemp004 = πF.MakeArgs(5)
							πTemp004[0] = ßweeks.ToObject()
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp004[1] = µx
							if πE = πg.CheckLocal(πF, µweeks, "weeks"); πE != nil {
								continue
							}
							πTemp004[2] = µweeks
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_US_PER_WEEK); πE != nil {
								continue
							}
							πTemp004[3] = πTemp001
							if πE = πg.CheckLocal(πF, µleftover, "leftover"); πE != nil {
								continue
							}
							πTemp004[4] = µleftover
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_accum); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp005}}}, πTemp002); πE != nil {
								continue
							}
							µx = πTemp001
							µleftover = πTemp005
							goto Label14
						Label14:
							if πE = πg.CheckLocal(πF, µleftover, "leftover"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.NE(πF, µleftover, πg.NewFloat(0.0).ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label15
							}
							goto Label16
							// line 478: if leftover != 0.0:
							πF.SetLineno(478)
						Label15:
							// line 479: x += _round(leftover)
							πF.SetLineno(479)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µleftover, "leftover"); πE != nil {
								continue
							}
							πTemp004[0] = µleftover
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_round); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp001, πE = πg.IAdd(πF, µx, πTemp002); πE != nil {
								continue
							}
							µx = πTemp001
							goto Label16
						Label16:
							// line 480: return cls._from_microseconds(x)
							πF.SetLineno(480)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp004[0] = µx
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µcls, ß_from_microseconds, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
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
					// line 483: def _from_microseconds(cls, us):
					πF.SetLineno(483)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "cls", Def: nil}
					πTemp002[1] = πg.Param{Name: "us", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("_from_microseconds", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µcls *πg.Object = πArgs[0]; _ = µcls
						var µus *πg.Object = πArgs[1]; _ = µus
						var µs *πg.Object = πg.UnboundLocal; _ = µs
						var µd *πg.Object = πg.UnboundLocal; _ = µd
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
							// line 484: s, us = divmod(us, _US_PER_SECOND)
							πF.SetLineno(484)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µus, "us"); πE != nil {
								continue
							}
							πTemp001[0] = µus
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_US_PER_SECOND); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßdivmod); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp004}}}, πTemp003); πE != nil {
								continue
							}
							µs = πTemp002
							µus = πTemp004
							// line 485: d, s = divmod(s, _SECONDS_PER_DAY)
							πF.SetLineno(485)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp001[0] = µs
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_SECONDS_PER_DAY); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßdivmod); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp004}}}, πTemp003); πE != nil {
								continue
							}
							µd = πTemp002
							µs = πTemp004
							// line 486: return cls._create(d, s, us, False)
							πF.SetLineno(486)
							πTemp001 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp001[0] = µd
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp001[1] = µs
							if πE = πg.CheckLocal(πF, µus, "us"); πE != nil {
								continue
							}
							πTemp001[2] = µus
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp001[3] = πTemp002
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µcls, ß_create, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ß_from_microseconds.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 482: @classmethod
					πF.SetLineno(482)
					πTemp004 = πF.MakeArgs(1)
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ß_from_microseconds); πE != nil {
						continue
					}
					πTemp004[0] = πTemp005
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßclassmethod); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πE = πClass.SetItem(πF, ß_from_microseconds.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 489: def _create(cls, d, s, us, normalize):
					πF.SetLineno(489)
					πTemp002 = make([]πg.Param, 5)
					πTemp002[0] = πg.Param{Name: "cls", Def: nil}
					πTemp002[1] = πg.Param{Name: "d", Def: nil}
					πTemp002[2] = πg.Param{Name: "s", Def: nil}
					πTemp002[3] = πg.Param{Name: "us", Def: nil}
					πTemp002[4] = πg.Param{Name: "normalize", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("_create", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µcls *πg.Object = πArgs[0]; _ = µcls
						var µd *πg.Object = πArgs[1]; _ = µd
						var µs *πg.Object = πArgs[2]; _ = µs
						var µus *πg.Object = πArgs[3]; _ = µus
						var µnormalize *πg.Object = πArgs[4]; _ = µnormalize
						var µself *πg.Object = πg.UnboundLocal; _ = µself
						var πTemp001 bool
						_ = πTemp001
						var πTemp002 []*πg.Object
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
							if πE = πg.CheckLocal(πF, µnormalize, "normalize"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.IsTrue(πF, µnormalize); πE != nil {
								continue
							}
							if πTemp001 {
								goto Label1
							}
							goto Label2
							// line 490: if normalize:
							πF.SetLineno(490)
						Label1:
							// line 491: s, us = _normalize_pair(s, us, 1000000)
							πF.SetLineno(491)
							πTemp002 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp002[0] = µs
							if πE = πg.CheckLocal(πF, µus, "us"); πE != nil {
								continue
							}
							πTemp002[1] = µus
							πTemp002[2] = πg.NewInt(1000000).ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ß_normalize_pair); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp005}}}, πTemp004); πE != nil {
								continue
							}
							µs = πTemp003
							µus = πTemp005
							// line 492: d, s = _normalize_pair(d, s, 24*3600)
							πF.SetLineno(492)
							πTemp002 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp002[0] = µd
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp002[1] = µs
							if πTemp003, πE = πg.Mul(πF, πg.NewInt(24).ToObject(), πg.NewInt(3600).ToObject()); πE != nil {
								continue
							}
							πTemp002[2] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ß_normalize_pair); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp005}}}, πTemp004); πE != nil {
								continue
							}
							µd = πTemp003
							µs = πTemp005
							goto Label2
						Label2:
							if πTemp006, πE = πg.ResolveGlobal(πF, ß_MAX_DELTA_DAYS); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Neg(πF, πTemp006); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.LE(πF, πTemp005, µd); πE != nil {
								continue
							}
							if πTemp001, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if !πTemp001 {
								goto Label3
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ß_MAX_DELTA_DAYS); πE != nil {
								continue
							}
							if πTemp004, πE = πg.LE(πF, µd, πTemp005); πE != nil {
								continue
							}
						Label3:
							if πTemp001, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp001).ToObject()
							if πTemp001, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp001 {
								goto Label4
							}
							goto Label5
							// line 494: if not -_MAX_DELTA_DAYS <= d <= _MAX_DELTA_DAYS:
							πF.SetLineno(494)
						Label4:
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ß_MAX_DELTA_DAYS); πE != nil {
								continue
							}
							πTemp004 = πg.NewTuple2(µd, πTemp005).ToObject()
							if πTemp003, πE = πg.Mod(πF, πg.NewStr("days=%d; must have magnitude <= %d").ToObject(), πTemp004); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßOverflowError); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 495: raise OverflowError("days=%d; must have magnitude <= %d" % (d, _MAX_DELTA_DAYS))
							πF.SetLineno(495)
							πE = πF.Raise(πTemp004, nil, nil)
							continue
							goto Label5
						Label5:
							// line 497: self = object.__new__(cls)
							πF.SetLineno(497)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							πTemp002[0] = µcls
							if πTemp003, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ß__new__, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µself = πTemp003
							// line 498: self._days = d
							πF.SetLineno(498)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, µd); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_days, πTemp003); πE != nil {
								continue
							}
							// line 499: self._seconds = s
							πF.SetLineno(499)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, µs); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_seconds, πTemp003); πE != nil {
								continue
							}
							// line 500: self._microseconds = us
							πF.SetLineno(500)
							if πE = πg.CheckLocal(πF, µus, "us"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, µus); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_microseconds, πTemp003); πE != nil {
								continue
							}
							// line 501: self._hashcode = -1
							πF.SetLineno(501)
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_hashcode, πTemp004); πE != nil {
								continue
							}
							// line 502: return self
							πF.SetLineno(502)
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
					if πE = πClass.SetItem(πF, ß_create.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 488: @classmethod
					πF.SetLineno(488)
					πTemp004 = πF.MakeArgs(1)
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ß_create); πE != nil {
						continue
					}
					πTemp004[0] = πTemp006
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßclassmethod); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πE = πClass.SetItem(πF, ß_create.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 504: def _to_microseconds(self):
					πF.SetLineno(504)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("_to_microseconds", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
							// line 505: return ((self._days * _SECONDS_PER_DAY + self._seconds) * _US_PER_SECOND +
							πF.SetLineno(505)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ß_days, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.ResolveGlobal(πF, ß_SECONDS_PER_DAY); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Mul(πF, πTemp005, πTemp006); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ß_seconds, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, πTemp004, πTemp005); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ß_US_PER_SECOND); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mul(πF, πTemp003, πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_microseconds, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp002, πTemp003); πE != nil {
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
					if πE = πClass.SetItem(πF, ß_to_microseconds.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 508: def __repr__(self):
					πF.SetLineno(508)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("__repr__", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µmodule *πg.Object = πg.UnboundLocal; _ = µmodule
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
							// line 509: module = "datetime." if self.__class__ is timedelta else ""
							πF.SetLineno(509)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß__class__, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßtimedelta); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp003 == πTemp004).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp005 {
								goto Label1
							}
							πTemp001 = πg.NewStr("datetime.").ToObject()
							goto Label2
						Label1:
							πTemp001 = ß.ToObject()
						Label2:
							µmodule = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_microseconds, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label3
							}
							goto Label4
							// line 510: if self._microseconds:
							πF.SetLineno(510)
						Label3:
							// line 511: return "%s(%d, %d, %d)" % (module + self.__class__.__name__,
							πF.SetLineno(511)
							if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß__class__, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp004, ß__name__, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, µmodule, πTemp006); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß_days, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ß_seconds, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ß_microseconds, nil); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple4(πTemp003, πTemp004, πTemp006, πTemp007).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("%s(%d, %d, %d)").ToObject(), πTemp002); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label4
						Label4:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_seconds, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label5
							}
							goto Label6
							// line 515: if self._seconds:
							πF.SetLineno(515)
						Label5:
							// line 516: return "%s(%d, %d)" % (module + self.__class__.__name__,
							πF.SetLineno(516)
							if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß__class__, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp004, ß__name__, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, µmodule, πTemp006); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß_days, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ß_seconds, nil); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple3(πTemp003, πTemp004, πTemp006).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("%s(%d, %d)").ToObject(), πTemp002); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label6
						Label6:
							// line 519: return "%s(%d)" % (module + self.__class__.__name__, self._days)
							πF.SetLineno(519)
							if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß__class__, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp004, ß__name__, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, µmodule, πTemp006); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß_days, nil); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("%s(%d)").ToObject(), πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__repr__.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 521: def __str__(self):
					πF.SetLineno(521)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("__str__", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µmm *πg.Object = πg.UnboundLocal; _ = µmm
						var µss *πg.Object = πg.UnboundLocal; _ = µss
						var µhh *πg.Object = πg.UnboundLocal; _ = µhh
						var µs *πg.Object = πg.UnboundLocal; _ = µs
						var µplural *πg.Object = πg.UnboundLocal; _ = µplural
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
						var πTemp006 []πg.Param
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
							// line 522: mm, ss = divmod(self._seconds, 60)
							πF.SetLineno(522)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_seconds, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = πg.NewInt(60).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßdivmod); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp004}}}, πTemp003); πE != nil {
								continue
							}
							µmm = πTemp002
							µss = πTemp004
							// line 523: hh, mm = divmod(mm, 60)
							πF.SetLineno(523)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µmm, "mm"); πE != nil {
								continue
							}
							πTemp001[0] = µmm
							πTemp001[1] = πg.NewInt(60).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßdivmod); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp004}}}, πTemp003); πE != nil {
								continue
							}
							µhh = πTemp002
							µmm = πTemp004
							// line 524: s = "%d:%02d:%02d" % (hh, mm, ss)
							πF.SetLineno(524)
							if πE = πg.CheckLocal(πF, µhh, "hh"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µmm, "mm"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple3(µhh, µmm, µss).ToObject()
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("%d:%02d:%02d").ToObject(), πTemp003); πE != nil {
								continue
							}
							µs = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_days, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							goto Label2
							// line 525: if self._days:
							πF.SetLineno(525)
						Label1:
							// line 526: def plural(n):
							πF.SetLineno(526)
							πTemp006 = make([]πg.Param, 1)
							πTemp006[0] = πg.Param{Name: "n", Def: nil}
							πTemp002 = πg.NewFunction(πg.NewCode("plural", "/usr/lib/python2.7/datetime.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µn *πg.Object = πArgs[0]; _ = µn
								var πTemp001 *πg.Object
								_ = πTemp001
								var πTemp002 *πg.Object
								_ = πTemp002
								var πTemp003 bool
								_ = πTemp003
								var πTemp004 *πg.Object
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
								for ; πF.State() >= 0; πF.PopCheckpoint() {
									switch πF.State() {
									case 0:
									default: panic("unexpected function state")
									}
									// line 527: return n, abs(n) != 1 and "s" or ""
									πF.SetLineno(527)
									if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
										continue
									}
									πTemp007 = πF.MakeArgs(1)
									if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
										continue
									}
									πTemp007[0] = µn
									if πTemp008, πE = πg.ResolveGlobal(πF, ßabs); πE != nil {
										continue
									}
									if πTemp009, πE = πTemp008.Call(πF, πTemp007, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp007)
									if πTemp006, πE = πg.NE(πF, πTemp009, πg.NewInt(1).ToObject()); πE != nil {
										continue
									}
									πTemp004 = πTemp006
									if πTemp005, πE = πg.IsTrue(πF, πTemp004); πE != nil {
										continue
									}
									if !πTemp005 {
										goto Label2
									}
									πTemp004 = ßs.ToObject()
								Label2:
									πTemp002 = πTemp004
									if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
										continue
									}
									if πTemp003 {
										goto Label1
									}
									πTemp002 = ß.ToObject()
								Label1:
									πTemp001 = πg.NewTuple2(µn, πTemp002).ToObject()
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
							µplural = πTemp002
							// line 528: s = ("%d day%s, " % plural(self._days)) + s
							πF.SetLineno(528)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ß_days, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp007
							if πE = πg.CheckLocal(πF, µplural, "plural"); πE != nil {
								continue
							}
							if πTemp007, πE = µplural.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp004, πE = πg.Mod(πF, πg.NewStr("%d day%s, ").ToObject(), πTemp007); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, πTemp004, µs); πE != nil {
								continue
							}
							µs = πTemp003
							goto Label2
						Label2:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_microseconds, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label3
							}
							goto Label4
							// line 529: if self._microseconds:
							πF.SetLineno(529)
						Label3:
							// line 530: s = s + ".%06d" % self._microseconds
							πF.SetLineno(530)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ß_microseconds, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Mod(πF, πg.NewStr(".%06d").ToObject(), πTemp007); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, µs, πTemp004); πE != nil {
								continue
							}
							µs = πTemp003
							goto Label4
						Label4:
							// line 531: return s
							πF.SetLineno(531)
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
					if πE = πClass.SetItem(πF, ß__str__.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 533: def total_seconds(self):
					πF.SetLineno(533)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("total_seconds", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 534: """Total seconds in the duration."""
							πF.SetLineno(534)
							// line 536: return float(self._to_microseconds()) / float(10**6)
							πF.SetLineno(536)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_to_microseconds, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßfloat); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp002 = πF.MakeArgs(1)
							if πTemp003, πE = πg.Pow(πF, πg.NewInt(10).ToObject(), πg.NewInt(6).ToObject()); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßfloat); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp001, πE = πg.Div(πF, πTemp004, πTemp005); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtotal_seconds.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 534: """Total seconds in the duration."""
					πF.SetLineno(534)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp010}, πg.NewStr("Total seconds in the duration.").ToObject()); πE != nil {
						continue
					}
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ßtotal_seconds); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp011, ß__doc__, πTemp010); πE != nil {
						continue
					}
					// line 540: def days(self):
					πF.SetLineno(540)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp010 = πg.NewFunction(πg.NewCode("days", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 541: """days"""
							πF.SetLineno(541)
							// line 542: return self._days
							πF.SetLineno(542)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_days, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßdays.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 541: """days"""
					πF.SetLineno(541)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, ßdays.ToObject()); πE != nil {
						continue
					}
					if πTemp012, πE = πg.ResolveClass(πF, πClass, nil, ßdays); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp012, ß__doc__, πTemp011); πE != nil {
						continue
					}
					// line 539: @property
					πF.SetLineno(539)
					πTemp004 = πF.MakeArgs(1)
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ßdays); πE != nil {
						continue
					}
					πTemp004[0] = πTemp011
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ßproperty); πE != nil {
						continue
					}
					if πTemp012, πE = πTemp011.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πE = πClass.SetItem(πF, ßdays.ToObject(), πTemp012); πE != nil {
						continue
					}
					// line 545: def seconds(self):
					πF.SetLineno(545)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp011 = πg.NewFunction(πg.NewCode("seconds", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 546: """seconds"""
							πF.SetLineno(546)
							// line 547: return self._seconds
							πF.SetLineno(547)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_seconds, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßseconds.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 546: """seconds"""
					πF.SetLineno(546)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp012}, ßseconds.ToObject()); πE != nil {
						continue
					}
					if πTemp013, πE = πg.ResolveClass(πF, πClass, nil, ßseconds); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp013, ß__doc__, πTemp012); πE != nil {
						continue
					}
					// line 544: @property
					πF.SetLineno(544)
					πTemp004 = πF.MakeArgs(1)
					if πTemp012, πE = πg.ResolveClass(πF, πClass, nil, ßseconds); πE != nil {
						continue
					}
					πTemp004[0] = πTemp012
					if πTemp012, πE = πg.ResolveClass(πF, πClass, nil, ßproperty); πE != nil {
						continue
					}
					if πTemp013, πE = πTemp012.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πE = πClass.SetItem(πF, ßseconds.ToObject(), πTemp013); πE != nil {
						continue
					}
					// line 550: def microseconds(self):
					πF.SetLineno(550)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp012 = πg.NewFunction(πg.NewCode("microseconds", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 551: """microseconds"""
							πF.SetLineno(551)
							// line 552: return self._microseconds
							πF.SetLineno(552)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_microseconds, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßmicroseconds.ToObject(), πTemp012); πE != nil {
						continue
					}
					// line 551: """microseconds"""
					πF.SetLineno(551)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp013}, ßmicroseconds.ToObject()); πE != nil {
						continue
					}
					if πTemp014, πE = πg.ResolveClass(πF, πClass, nil, ßmicroseconds); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp014, ß__doc__, πTemp013); πE != nil {
						continue
					}
					// line 549: @property
					πF.SetLineno(549)
					πTemp004 = πF.MakeArgs(1)
					if πTemp013, πE = πg.ResolveClass(πF, πClass, nil, ßmicroseconds); πE != nil {
						continue
					}
					πTemp004[0] = πTemp013
					if πTemp013, πE = πg.ResolveClass(πF, πClass, nil, ßproperty); πE != nil {
						continue
					}
					if πTemp014, πE = πTemp013.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πE = πClass.SetItem(πF, ßmicroseconds.ToObject(), πTemp014); πE != nil {
						continue
					}
					// line 554: def __add__(self, other):
					πF.SetLineno(554)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "other", Def: nil}
					πTemp013 = πg.NewFunction(πg.NewCode("__add__", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
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
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtimedelta); πE != nil {
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
							goto Label2
							// line 555: if isinstance(other, timedelta):
							πF.SetLineno(555)
						Label1:
							// line 558: return timedelta._create(self._days + other._days,
							πF.SetLineno(558)
							πTemp001 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_days, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µother, ß_days, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp003, πTemp005); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_seconds, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µother, ß_seconds, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp003, πTemp005); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_microseconds, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µother, ß_microseconds, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp003, πTemp005); πE != nil {
								continue
							}
							πTemp001[2] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[3] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtimedelta); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ß_create, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πR = πTemp002
							continue
							goto Label2
						Label2:
							// line 562: return NotImplemented
							πF.SetLineno(562)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNotImplemented); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__add__.ToObject(), πTemp013); πE != nil {
						continue
					}
					// line 564: def __sub__(self, other):
					πF.SetLineno(564)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "other", Def: nil}
					πTemp014 = πg.NewFunction(πg.NewCode("__sub__", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
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
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtimedelta); πE != nil {
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
							goto Label2
							// line 565: if isinstance(other, timedelta):
							πF.SetLineno(565)
						Label1:
							// line 568: return timedelta._create(self._days - other._days,
							πF.SetLineno(568)
							πTemp001 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_days, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µother, ß_days, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Sub(πF, πTemp003, πTemp005); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_seconds, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µother, ß_seconds, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Sub(πF, πTemp003, πTemp005); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_microseconds, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µother, ß_microseconds, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Sub(πF, πTemp003, πTemp005); πE != nil {
								continue
							}
							πTemp001[2] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[3] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtimedelta); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ß_create, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πR = πTemp002
							continue
							goto Label2
						Label2:
							// line 572: return NotImplemented
							πF.SetLineno(572)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNotImplemented); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__sub__.ToObject(), πTemp014); πE != nil {
						continue
					}
					// line 574: def __neg__(self):
					πF.SetLineno(574)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp015 = πg.NewFunction(πg.NewCode("__neg__", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 577: return timedelta._create(-self._days,
							πF.SetLineno(577)
							πTemp001 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_days, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Neg(πF, πTemp003); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_seconds, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Neg(πF, πTemp003); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_microseconds, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Neg(πF, πTemp003); πE != nil {
								continue
							}
							πTemp001[2] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[3] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtimedelta); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ß_create, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__neg__.ToObject(), πTemp015); πE != nil {
						continue
					}
					// line 582: def __pos__(self):
					πF.SetLineno(582)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp016 = πg.NewFunction(πg.NewCode("__pos__", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 585: return timedelta._create(self._days,
							πF.SetLineno(585)
							πTemp001 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_days, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_seconds, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_microseconds, nil); πE != nil {
								continue
							}
							πTemp001[2] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp001[3] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtimedelta); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ß_create, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__pos__.ToObject(), πTemp016); πE != nil {
						continue
					}
					// line 590: def __abs__(self):
					πF.SetLineno(590)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp017 = πg.NewFunction(πg.NewCode("__abs__", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_days, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.LT(πF, πTemp002, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 591: if self._days < 0:
							πF.SetLineno(591)
						Label1:
							// line 592: return -self
							πF.SetLineno(592)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Neg(πF, µself); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label3
						Label2:
							// line 594: return self
							πF.SetLineno(594)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πR = µself
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
					if πE = πClass.SetItem(πF, ß__abs__.ToObject(), πTemp017); πE != nil {
						continue
					}
					// line 596: def __mul__(self, other):
					πF.SetLineno(596)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "other", Def: nil}
					πTemp018 = πg.NewFunction(πg.NewCode("__mul__", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
						var µusec *πg.Object = πg.UnboundLocal; _ = µusec
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp002[0] = µother
							if πTemp004, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßlong); πE != nil {
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
							// line 597: if not isinstance(other, (int, long)):
							πF.SetLineno(597)
						Label1:
							// line 598: return NotImplemented
							πF.SetLineno(598)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNotImplemented); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label2
						Label2:
							// line 599: usec = self._to_microseconds()
							πF.SetLineno(599)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_to_microseconds, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µusec = πTemp003
							// line 600: return timedelta._from_microseconds(usec * other)
							πF.SetLineno(600)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µusec, "usec"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mul(πF, µusec, µother); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßtimedelta); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ß_from_microseconds, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
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
					if πE = πClass.SetItem(πF, ß__mul__.ToObject(), πTemp018); πE != nil {
						continue
					}
					// line 602: __rmul__ = __mul__
					πF.SetLineno(602)
					if πTemp019, πE = πg.ResolveClass(πF, πClass, nil, ß__mul__); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ß__rmul__.ToObject(), πTemp019); πE != nil {
						continue
					}
					// line 604: def __div__(self, other):
					πF.SetLineno(604)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "other", Def: nil}
					πTemp019 = πg.NewFunction(πg.NewCode("__div__", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
						var µusec *πg.Object = πg.UnboundLocal; _ = µusec
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
						var πTemp007 []*πg.Object
						_ = πTemp007
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp002[0] = µother
							if πTemp004, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßlong); πE != nil {
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
							// line 605: if not isinstance(other, (int, long)):
							πF.SetLineno(605)
						Label1:
							// line 606: return NotImplemented
							πF.SetLineno(606)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNotImplemented); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label2
						Label2:
							// line 607: usec = self._to_microseconds()
							πF.SetLineno(607)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_to_microseconds, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µusec = πTemp003
							// line 609: return timedelta._from_microseconds(int(usec) / int(other))
							πF.SetLineno(609)
							πTemp002 = πF.MakeArgs(1)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µusec, "usec"); πE != nil {
								continue
							}
							πTemp007[0] = µusec
							if πTemp003, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp007[0] = µother
							if πTemp003, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							if πTemp001, πE = πg.Div(πF, πTemp004, πTemp005); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßtimedelta); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ß_from_microseconds, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
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
					if πE = πClass.SetItem(πF, ß__div__.ToObject(), πTemp019); πE != nil {
						continue
					}
					// line 611: __floordiv__ = __div__
					πF.SetLineno(611)
					if πTemp020, πE = πg.ResolveClass(πF, πClass, nil, ß__div__); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ß__floordiv__.ToObject(), πTemp020); πE != nil {
						continue
					}
					// line 615: def __eq__(self, other):
					πF.SetLineno(615)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "other", Def: nil}
					πTemp020 = πg.NewFunction(πg.NewCode("__eq__", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
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
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtimedelta); πE != nil {
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
							goto Label2
							// line 616: if isinstance(other, timedelta):
							πF.SetLineno(616)
						Label1:
							// line 617: return self._cmp(other) == 0
							πF.SetLineno(617)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_cmp, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.Eq(πF, πTemp005, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							πR = πTemp002
							continue
							goto Label3
						Label2:
							// line 619: return False
							πF.SetLineno(619)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__eq__.ToObject(), πTemp020); πE != nil {
						continue
					}
					// line 621: def __ne__(self, other):
					πF.SetLineno(621)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "other", Def: nil}
					πTemp021 = πg.NewFunction(πg.NewCode("__ne__", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
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
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtimedelta); πE != nil {
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
							goto Label2
							// line 622: if isinstance(other, timedelta):
							πF.SetLineno(622)
						Label1:
							// line 623: return self._cmp(other) != 0
							πF.SetLineno(623)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_cmp, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.NE(πF, πTemp005, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							πR = πTemp002
							continue
							goto Label3
						Label2:
							// line 625: return True
							πF.SetLineno(625)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__ne__.ToObject(), πTemp021); πE != nil {
						continue
					}
					// line 627: def __le__(self, other):
					πF.SetLineno(627)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "other", Def: nil}
					πTemp022 = πg.NewFunction(πg.NewCode("__le__", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
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
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtimedelta); πE != nil {
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
							goto Label2
							// line 628: if isinstance(other, timedelta):
							πF.SetLineno(628)
						Label1:
							// line 629: return self._cmp(other) <= 0
							πF.SetLineno(629)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_cmp, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.LE(πF, πTemp005, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							πR = πTemp002
							continue
							goto Label3
						Label2:
							// line 631: _cmperror(self, other)
							πF.SetLineno(631)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[1] = µother
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_cmperror); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
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
					if πE = πClass.SetItem(πF, ß__le__.ToObject(), πTemp022); πE != nil {
						continue
					}
					// line 633: def __lt__(self, other):
					πF.SetLineno(633)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "other", Def: nil}
					πTemp023 = πg.NewFunction(πg.NewCode("__lt__", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
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
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtimedelta); πE != nil {
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
							goto Label2
							// line 634: if isinstance(other, timedelta):
							πF.SetLineno(634)
						Label1:
							// line 635: return self._cmp(other) < 0
							πF.SetLineno(635)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_cmp, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.LT(πF, πTemp005, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							πR = πTemp002
							continue
							goto Label3
						Label2:
							// line 637: _cmperror(self, other)
							πF.SetLineno(637)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[1] = µother
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_cmperror); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
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
					if πE = πClass.SetItem(πF, ß__lt__.ToObject(), πTemp023); πE != nil {
						continue
					}
					// line 639: def __ge__(self, other):
					πF.SetLineno(639)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "other", Def: nil}
					πTemp024 = πg.NewFunction(πg.NewCode("__ge__", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
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
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtimedelta); πE != nil {
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
							goto Label2
							// line 640: if isinstance(other, timedelta):
							πF.SetLineno(640)
						Label1:
							// line 641: return self._cmp(other) >= 0
							πF.SetLineno(641)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_cmp, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.GE(πF, πTemp005, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							πR = πTemp002
							continue
							goto Label3
						Label2:
							// line 643: _cmperror(self, other)
							πF.SetLineno(643)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[1] = µother
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_cmperror); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
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
					if πE = πClass.SetItem(πF, ß__ge__.ToObject(), πTemp024); πE != nil {
						continue
					}
					// line 645: def __gt__(self, other):
					πF.SetLineno(645)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "other", Def: nil}
					πTemp025 = πg.NewFunction(πg.NewCode("__gt__", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
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
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtimedelta); πE != nil {
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
							goto Label2
							// line 646: if isinstance(other, timedelta):
							πF.SetLineno(646)
						Label1:
							// line 647: return self._cmp(other) > 0
							πF.SetLineno(647)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_cmp, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.GT(πF, πTemp005, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							πR = πTemp002
							continue
							goto Label3
						Label2:
							// line 649: _cmperror(self, other)
							πF.SetLineno(649)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[1] = µother
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_cmperror); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
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
					if πE = πClass.SetItem(πF, ß__gt__.ToObject(), πTemp025); πE != nil {
						continue
					}
					// line 651: def _cmp(self, other):
					πF.SetLineno(651)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "other", Def: nil}
					πTemp026 = πg.NewFunction(πg.NewCode("_cmp", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
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
							// line 652: assert isinstance(other, timedelta)
							πF.SetLineno(652)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtimedelta); πE != nil {
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
							if πE = πg.Assert(πF, πTemp003, nil); πE != nil {
								continue
							}
							// line 653: return _cmp(self._getstate(), other._getstate())
							πF.SetLineno(653)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_getstate, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µother, ß_getstate, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_cmp); πE != nil {
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
					if πE = πClass.SetItem(πF, ß_cmp.ToObject(), πTemp026); πE != nil {
						continue
					}
					// line 655: def __hash__(self):
					πF.SetLineno(655)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp027 = πg.NewFunction(πg.NewCode("__hash__", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_hashcode, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
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
							// line 656: if self._hashcode == -1:
							πF.SetLineno(656)
						Label1:
							// line 657: self._hashcode = hash(self._getstate())
							πF.SetLineno(657)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_getstate, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp005[0] = πTemp002
							if πTemp001, πE = πg.ResolveGlobal(πF, ßhash); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_hashcode, πTemp001); πE != nil {
								continue
							}
							goto Label2
						Label2:
							// line 658: return self._hashcode
							πF.SetLineno(658)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_hashcode, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__hash__.ToObject(), πTemp027); πE != nil {
						continue
					}
					// line 660: def __nonzero__(self):
					πF.SetLineno(660)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp028 = πg.NewFunction(πg.NewCode("__nonzero__", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
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
							// line 661: return (self._days != 0 or
							πF.SetLineno(661)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß_days, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.NE(πF, πTemp004, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp003
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß_seconds, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.NE(πF, πTemp004, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp003
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß_microseconds, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.NE(πF, πTemp004, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp003
						Label1:
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
					if πE = πClass.SetItem(πF, ß__nonzero__.ToObject(), πTemp028); πE != nil {
						continue
					}
					// line 667: def _getstate(self):
					πF.SetLineno(667)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp029 = πg.NewFunction(πg.NewCode("_getstate", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πTemp001 *πg.Object
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
							// line 668: return (self._days, self._seconds, self._microseconds)
							πF.SetLineno(668)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_days, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_seconds, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß_microseconds, nil); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple3(πTemp002, πTemp003, πTemp004).ToObject()
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
					if πE = πClass.SetItem(πF, ß_getstate.ToObject(), πTemp029); πE != nil {
						continue
					}
					// line 670: def __reduce__(self):
					πF.SetLineno(670)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp030 = πg.NewFunction(πg.NewCode("__reduce__", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πTemp001 *πg.Object
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
							// line 671: return (self.__class__, self._getstate())
							πF.SetLineno(671)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß__class__, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_getstate, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple2(πTemp002, πTemp004).ToObject()
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
					if πE = πClass.SetItem(πF, ß__reduce__.ToObject(), πTemp030); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp029, πE = πTemp027.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp029 == nil {
				πTemp029 = πg.TypeType.ToObject()
			}
			if πTemp030, πE = πTemp029.Call(πF, []*πg.Object{πg.NewStr("timedelta").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp027.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßtimedelta.ToObject(), πTemp030); πE != nil {
				continue
			}
			// line 673: timedelta.min = timedelta(-_MAX_DELTA_DAYS)
			πF.SetLineno(673)
			πTemp002 = πF.MakeArgs(1)
			if πTemp029, πE = πg.ResolveGlobal(πF, ß_MAX_DELTA_DAYS); πE != nil {
				continue
			}
			if πTemp028, πE = πg.Neg(πF, πTemp029); πE != nil {
				continue
			}
			πTemp002[0] = πTemp028
			if πTemp028, πE = πg.ResolveGlobal(πF, ßtimedelta); πE != nil {
				continue
			}
			if πTemp029, πE = πTemp028.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp028}, πTemp029); πE != nil {
				continue
			}
			if πTemp030, πE = πg.ResolveGlobal(πF, ßtimedelta); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp030, ßmin, πTemp028); πE != nil {
				continue
			}
			// line 674: timedelta.max = timedelta(_MAX_DELTA_DAYS, 24*3600-1, 1000000-1)
			πF.SetLineno(674)
			πTemp002 = πF.MakeArgs(3)
			if πTemp028, πE = πg.ResolveGlobal(πF, ß_MAX_DELTA_DAYS); πE != nil {
				continue
			}
			πTemp002[0] = πTemp028
			if πTemp029, πE = πg.Mul(πF, πg.NewInt(24).ToObject(), πg.NewInt(3600).ToObject()); πE != nil {
				continue
			}
			if πTemp028, πE = πg.Sub(πF, πTemp029, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp002[1] = πTemp028
			if πTemp028, πE = πg.Sub(πF, πg.NewInt(1000000).ToObject(), πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp002[2] = πTemp028
			if πTemp028, πE = πg.ResolveGlobal(πF, ßtimedelta); πE != nil {
				continue
			}
			if πTemp029, πE = πTemp028.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp028}, πTemp029); πE != nil {
				continue
			}
			if πTemp030, πE = πg.ResolveGlobal(πF, ßtimedelta); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp030, ßmax, πTemp028); πE != nil {
				continue
			}
			// line 675: timedelta.resolution = timedelta(microseconds=1)
			πF.SetLineno(675)
			πTemp031 = πg.KWArgs{
				{"microseconds", πg.NewInt(1).ToObject()},
			}
			if πTemp028, πE = πg.ResolveGlobal(πF, ßtimedelta); πE != nil {
				continue
			}
			if πTemp029, πE = πTemp028.Call(πF, nil, πTemp031); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp028}, πTemp029); πE != nil {
				continue
			}
			if πTemp030, πE = πg.ResolveGlobal(πF, ßtimedelta); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp030, ßresolution, πTemp028); πE != nil {
				continue
			}
			// line 677: class date(object):
			πF.SetLineno(677)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp030, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp030
			πTemp027 = πg.NewDict()
			if πTemp028, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp027.SetItem(πF, ß__module__.ToObject(), πTemp028); πE != nil {
				continue
			}
			_, πE = πg.NewCode("date", "/usr/lib/python2.7/datetime.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp027
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
				var πTemp030 *πg.Object
				_ = πTemp030
				var πTemp031 *πg.Object
				_ = πTemp031
				var πTemp032 *πg.Object
				_ = πTemp032
				var πTemp033 *πg.Object
				_ = πTemp033
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 678: """Concrete date type.
					πF.SetLineno(678)
					// line 678: """Concrete date type.
					πF.SetLineno(678)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Concrete date type.\n\n    Constructors:\n\n    __new__()\n    fromtimestamp()\n    today()\n    fromordinal()\n\n    Operators:\n\n    __repr__, __str__\n    __cmp__, __hash__\n    __add__, __radd__, __sub__ (add/radd only with timedelta arg)\n\n    Methods:\n\n    timetuple()\n    toordinal()\n    weekday()\n    isoweekday(), isocalendar(), isoformat()\n    ctime()\n    strftime()\n\n    Properties (readonly):\n    year, month, day\n    ").ToObject()); πE != nil {
						continue
					}
					// line 705: __slots__ = '_year', '_month', '_day', '_hashcode'
					πF.SetLineno(705)
					πTemp001 = πg.NewTuple4(ß_year.ToObject(), ß_month.ToObject(), ß_day.ToObject(), ß_hashcode.ToObject()).ToObject()
					if πE = πClass.SetItem(πF, ß__slots__.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 707: def __new__(cls, year, month=None, day=None):
					πF.SetLineno(707)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "cls", Def: nil}
					πTemp002[1] = πg.Param{Name: "year", Def: nil}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "month", Def: πTemp003}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[3] = πg.Param{Name: "day", Def: πTemp003}
					πTemp001 = πg.NewFunction(πg.NewCode("__new__", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µcls *πg.Object = πArgs[0]; _ = µcls
						var µyear *πg.Object = πArgs[1]; _ = µyear
						var µmonth *πg.Object = πArgs[2]; _ = µmonth
						var µday *πg.Object = πArgs[3]; _ = µday
						var µself *πg.Object = πg.UnboundLocal; _ = µself
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 708: """Constructor.
							πF.SetLineno(708)
							// line 721: year, month, day = _check_date_fields(year, month, day)
							πF.SetLineno(721)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µyear, "year"); πE != nil {
								continue
							}
							πTemp001[0] = µyear
							if πE = πg.CheckLocal(πF, µmonth, "month"); πE != nil {
								continue
							}
							πTemp001[1] = µmonth
							if πE = πg.CheckLocal(πF, µday, "day"); πE != nil {
								continue
							}
							πTemp001[2] = µday
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_check_date_fields); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}}}, πTemp003); πE != nil {
								continue
							}
							µyear = πTemp002
							µmonth = πTemp004
							µday = πTemp005
							// line 722: self = object.__new__(cls)
							πF.SetLineno(722)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							πTemp001[0] = µcls
							if πTemp002, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ß__new__, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µself = πTemp002
							// line 723: self._year = year
							πF.SetLineno(723)
							if πE = πg.CheckLocal(πF, µyear, "year"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µyear); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_year, πTemp002); πE != nil {
								continue
							}
							// line 724: self._month = month
							πF.SetLineno(724)
							if πE = πg.CheckLocal(πF, µmonth, "month"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µmonth); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_month, πTemp002); πE != nil {
								continue
							}
							// line 725: self._day = day
							πF.SetLineno(725)
							if πE = πg.CheckLocal(πF, µday, "day"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µday); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_day, πTemp002); πE != nil {
								continue
							}
							// line 726: self._hashcode = -1
							πF.SetLineno(726)
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_hashcode, πTemp003); πE != nil {
								continue
							}
							// line 727: return self
							πF.SetLineno(727)
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
					if πE = πClass.SetItem(πF, ß__new__.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 708: """Constructor.
					πF.SetLineno(708)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewStr("Constructor.\n\n        Arguments:\n\n        year, month, day (required, base 1)\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ß__new__); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp004, ß__doc__, πTemp003); πE != nil {
						continue
					}
					// line 732: def fromtimestamp(cls, t):
					πF.SetLineno(732)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "cls", Def: nil}
					πTemp002[1] = πg.Param{Name: "t", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("fromtimestamp", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µcls *πg.Object = πArgs[0]; _ = µcls
						var µt *πg.Object = πArgs[1]; _ = µt
						var µy *πg.Object = πg.UnboundLocal; _ = µy
						var µm *πg.Object = πg.UnboundLocal; _ = µm
						var µd *πg.Object = πg.UnboundLocal; _ = µd
						var µhh *πg.Object = πg.UnboundLocal; _ = µhh
						var µmm *πg.Object = πg.UnboundLocal; _ = µmm
						var µss *πg.Object = πg.UnboundLocal; _ = µss
						var µweekday *πg.Object = πg.UnboundLocal; _ = µweekday
						var µjday *πg.Object = πg.UnboundLocal; _ = µjday
						var µdst *πg.Object = πg.UnboundLocal; _ = µdst
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 733: "Construct a date from a POSIX timestamp (like time.time())."
							πF.SetLineno(733)
							// line 734: y, m, d, hh, mm, ss, weekday, jday, dst = _time.localtime(t)
							πF.SetLineno(734)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
								continue
							}
							πTemp001[0] = µt
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_time); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßlocaltime, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp006}, πg.TieTarget{Target: &πTemp007}, πg.TieTarget{Target: &πTemp008}, πg.TieTarget{Target: &πTemp009}, πg.TieTarget{Target: &πTemp010}, πg.TieTarget{Target: &πTemp011}}}, πTemp002); πE != nil {
								continue
							}
							µy = πTemp003
							µm = πTemp004
							µd = πTemp005
							µhh = πTemp006
							µmm = πTemp007
							µss = πTemp008
							µweekday = πTemp009
							µjday = πTemp010
							µdst = πTemp011
							// line 735: return cls(y, m, d)
							πF.SetLineno(735)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
								continue
							}
							πTemp001[0] = µy
							if πE = πg.CheckLocal(πF, µm, "m"); πE != nil {
								continue
							}
							πTemp001[1] = µm
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp001[2] = µd
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							if πTemp002, πE = µcls.Call(πF, πTemp001, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßfromtimestamp.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 733: "Construct a date from a POSIX timestamp (like time.time())."
					πF.SetLineno(733)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πg.NewStr("Construct a date from a POSIX timestamp (like time.time()).").ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßfromtimestamp); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp005, ß__doc__, πTemp004); πE != nil {
						continue
					}
					// line 731: @classmethod
					πF.SetLineno(731)
					πTemp006 = πF.MakeArgs(1)
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßfromtimestamp); πE != nil {
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
					if πE = πClass.SetItem(πF, ßfromtimestamp.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 738: def today(cls):
					πF.SetLineno(738)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "cls", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("today", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µcls *πg.Object = πArgs[0]; _ = µcls
						var µt *πg.Object = πg.UnboundLocal; _ = µt
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 739: "Construct a date from time.time()."
							πF.SetLineno(739)
							// line 740: t = _time.time()
							πF.SetLineno(740)
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_time); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßtime, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µt = πTemp001
							// line 741: return cls.fromtimestamp(t)
							πF.SetLineno(741)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
								continue
							}
							πTemp003[0] = µt
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µcls, ßfromtimestamp, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
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
					if πE = πClass.SetItem(πF, ßtoday.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 739: "Construct a date from time.time()."
					πF.SetLineno(739)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πg.NewStr("Construct a date from time.time().").ToObject()); πE != nil {
						continue
					}
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßtoday); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp007, ß__doc__, πTemp005); πE != nil {
						continue
					}
					// line 737: @classmethod
					πF.SetLineno(737)
					πTemp006 = πF.MakeArgs(1)
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßtoday); πE != nil {
						continue
					}
					πTemp006[0] = πTemp005
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßclassmethod); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp005.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πE = πClass.SetItem(πF, ßtoday.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 744: def fromordinal(cls, n):
					πF.SetLineno(744)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "cls", Def: nil}
					πTemp002[1] = πg.Param{Name: "n", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("fromordinal", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µcls *πg.Object = πArgs[0]; _ = µcls
						var µn *πg.Object = πArgs[1]; _ = µn
						var µy *πg.Object = πg.UnboundLocal; _ = µy
						var µm *πg.Object = πg.UnboundLocal; _ = µm
						var µd *πg.Object = πg.UnboundLocal; _ = µd
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 745: """Contruct a date from a proleptic Gregorian ordinal.
							πF.SetLineno(745)
							// line 750: y, m, d = _ord2ymd(n)
							πF.SetLineno(750)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							πTemp001[0] = µn
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_ord2ymd); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}}}, πTemp003); πE != nil {
								continue
							}
							µy = πTemp002
							µm = πTemp004
							µd = πTemp005
							// line 751: return cls(y, m, d)
							πF.SetLineno(751)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
								continue
							}
							πTemp001[0] = µy
							if πE = πg.CheckLocal(πF, µm, "m"); πE != nil {
								continue
							}
							πTemp001[1] = µm
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp001[2] = µd
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							if πTemp002, πE = µcls.Call(πF, πTemp001, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßfromordinal.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 745: """Contruct a date from a proleptic Gregorian ordinal.
					πF.SetLineno(745)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp007}, πg.NewStr("Contruct a date from a proleptic Gregorian ordinal.\n\n        January 1 of year 1 is day 1.  Only the year, month and day are\n        non-zero in the result.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ßfromordinal); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp008, ß__doc__, πTemp007); πE != nil {
						continue
					}
					// line 743: @classmethod
					πF.SetLineno(743)
					πTemp006 = πF.MakeArgs(1)
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßfromordinal); πE != nil {
						continue
					}
					πTemp006[0] = πTemp007
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßclassmethod); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp007.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πE = πClass.SetItem(πF, ßfromordinal.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 755: def __repr__(self):
					πF.SetLineno(755)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("__repr__", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µmodule *πg.Object = πg.UnboundLocal; _ = µmodule
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
							// line 756: """Convert to formal string, for repr().
							πF.SetLineno(756)
							// line 766: module = "datetime." if self.__class__ is date else ""
							πF.SetLineno(766)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß__class__, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßdate); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp003 == πTemp004).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp005 {
								goto Label1
							}
							πTemp001 = πg.NewStr("datetime.").ToObject()
							goto Label2
						Label1:
							πTemp001 = ß.ToObject()
						Label2:
							µmodule = πTemp001
							// line 767: return "%s(%d, %d, %d)" % (module + self.__class__.__name__,
							πF.SetLineno(767)
							if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß__class__, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp004, ß__name__, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, µmodule, πTemp006); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß_year, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ß_month, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ß_day, nil); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple4(πTemp003, πTemp004, πTemp006, πTemp007).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("%s(%d, %d, %d)").ToObject(), πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__repr__.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 756: """Convert to formal string, for repr().
					πF.SetLineno(756)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp008}, πg.NewStr("Convert to formal string, for repr().\n\n        >>> dt = datetime(2010, 1, 1)\n        >>> repr(dt)\n        'datetime.datetime(2010, 1, 1, 0, 0)'\n\n        >>> dt = datetime(2010, 1, 1, tzinfo=timezone.utc)\n        >>> repr(dt)\n        'datetime.datetime(2010, 1, 1, 0, 0, tzinfo=datetime.timezone.utc)'\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ß__repr__); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp009, ß__doc__, πTemp008); πE != nil {
						continue
					}
					// line 777: def ctime(self):
					πF.SetLineno(777)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("ctime", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µweekday *πg.Object = πg.UnboundLocal; _ = µweekday
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 778: "Return ctime() style string."
							πF.SetLineno(778)
							// line 779: weekday = self.toordinal() % 7 or 7
							πF.SetLineno(779)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßtoordinal, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Mod(πF, πTemp005, πg.NewInt(7).ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp003
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							πTemp001 = πg.NewInt(7).ToObject()
						Label1:
							µweekday = πTemp001
							// line 780: return "%s %s %2d 00:00:00 %04d" % (
							πF.SetLineno(780)
							if πE = πg.CheckLocal(πF, µweekday, "weekday"); πE != nil {
								continue
							}
							πTemp004 = µweekday
							if πTemp006, πE = πg.ResolveGlobal(πF, ß_DAYNAMES); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, πTemp006, πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ß_month, nil); πE != nil {
								continue
							}
							πTemp004 = πTemp006
							if πTemp007, πE = πg.ResolveGlobal(πF, ß_MONTHNAMES); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, πTemp007, πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß_day, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ß_year, nil); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple4(πTemp005, πTemp006, πTemp004, πTemp007).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("%s %s %2d 00:00:00 %04d").ToObject(), πTemp003); πE != nil {
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
					if πE = πClass.SetItem(πF, ßctime.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 778: "Return ctime() style string."
					πF.SetLineno(778)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp009}, πg.NewStr("Return ctime() style string.").ToObject()); πE != nil {
						continue
					}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßctime); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp010, ß__doc__, πTemp009); πE != nil {
						continue
					}
					// line 789: def __format__(self, fmt):
					πF.SetLineno(789)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "fmt", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("__format__", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µfmt *πg.Object = πArgs[1]; _ = µfmt
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µfmt, "fmt"); πE != nil {
								continue
							}
							πTemp002[0] = µfmt
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
							// line 790: if not isinstance(fmt, (str, unicode)):
							πF.SetLineno(790)
						Label1:
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfmt, "fmt"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µfmt, ß__class__, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ß__name__, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("__format__ expects str or unicode, not %s").ToObject(), πTemp004); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 791: raise ValueError("__format__ expects str or unicode, not %s" %
							πF.SetLineno(791)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label2
						Label2:
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfmt, "fmt"); πE != nil {
								continue
							}
							πTemp002[0] = µfmt
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp001, πE = πg.NE(πF, πTemp004, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label3
							}
							goto Label4
							// line 793: if len(fmt) != 0:
							πF.SetLineno(793)
						Label3:
							// line 794: return self.strftime(fmt)
							πF.SetLineno(794)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfmt, "fmt"); πE != nil {
								continue
							}
							πTemp002[0] = µfmt
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstrftime, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πR = πTemp003
							continue
							goto Label4
						Label4:
							// line 795: return str(self)
							πF.SetLineno(795)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp002[0] = µself
							if πTemp001, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
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
					if πE = πClass.SetItem(πF, ß__format__.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 797: def isoformat(self):
					πF.SetLineno(797)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp010 = πg.NewFunction(πg.NewCode("isoformat", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
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
						var πTemp008 *πg.Object
						_ = πTemp008
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 798: """Return the date formatted according to ISO.
							πF.SetLineno(798)
							// line 807: return "%s-%s-%s" % (str(self._year).zfill(4), str(self._month).zfill(2), str(self._day).zfill(2))
							πF.SetLineno(807)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewInt(4).ToObject()
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ß_year, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp005
							if πTemp005, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp005, πE = πg.GetAttr(πF, πTemp006, ßzfill, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewInt(2).ToObject()
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ß_month, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp005
							if πTemp005, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp005, πE = πg.GetAttr(πF, πTemp007, ßzfill, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewInt(2).ToObject()
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ß_day, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp005
							if πTemp005, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp005, πE = πg.GetAttr(πF, πTemp008, ßzfill, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πTemp002 = πg.NewTuple3(πTemp006, πTemp007, πTemp008).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("%s-%s-%s").ToObject(), πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ßisoformat.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 798: """Return the date formatted according to ISO.
					πF.SetLineno(798)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πg.NewStr("Return the date formatted according to ISO.\n\n        This is 'YYYY-MM-DD'.\n\n        References:\n        - http://www.w3.org/TR/NOTE-datetime\n        - http://www.cl.cam.ac.uk/~mgk25/iso-time.html\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp012, πE = πg.ResolveClass(πF, πClass, nil, ßisoformat); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp012, ß__doc__, πTemp011); πE != nil {
						continue
					}
					// line 809: __str__ = isoformat
					πF.SetLineno(809)
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ßisoformat); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ß__str__.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 813: def year(self):
					πF.SetLineno(813)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp011 = πg.NewFunction(πg.NewCode("year", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 814: """year (1-9999)"""
							πF.SetLineno(814)
							// line 815: return self._year
							πF.SetLineno(815)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_year, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßyear.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 814: """year (1-9999)"""
					πF.SetLineno(814)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp012}, πg.NewStr("year (1-9999)").ToObject()); πE != nil {
						continue
					}
					if πTemp013, πE = πg.ResolveClass(πF, πClass, nil, ßyear); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp013, ß__doc__, πTemp012); πE != nil {
						continue
					}
					// line 812: @property
					πF.SetLineno(812)
					πTemp006 = πF.MakeArgs(1)
					if πTemp012, πE = πg.ResolveClass(πF, πClass, nil, ßyear); πE != nil {
						continue
					}
					πTemp006[0] = πTemp012
					if πTemp012, πE = πg.ResolveClass(πF, πClass, nil, ßproperty); πE != nil {
						continue
					}
					if πTemp013, πE = πTemp012.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πE = πClass.SetItem(πF, ßyear.ToObject(), πTemp013); πE != nil {
						continue
					}
					// line 818: def month(self):
					πF.SetLineno(818)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp012 = πg.NewFunction(πg.NewCode("month", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 819: """month (1-12)"""
							πF.SetLineno(819)
							// line 820: return self._month
							πF.SetLineno(820)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_month, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßmonth.ToObject(), πTemp012); πE != nil {
						continue
					}
					// line 819: """month (1-12)"""
					πF.SetLineno(819)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp013}, πg.NewStr("month (1-12)").ToObject()); πE != nil {
						continue
					}
					if πTemp014, πE = πg.ResolveClass(πF, πClass, nil, ßmonth); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp014, ß__doc__, πTemp013); πE != nil {
						continue
					}
					// line 817: @property
					πF.SetLineno(817)
					πTemp006 = πF.MakeArgs(1)
					if πTemp013, πE = πg.ResolveClass(πF, πClass, nil, ßmonth); πE != nil {
						continue
					}
					πTemp006[0] = πTemp013
					if πTemp013, πE = πg.ResolveClass(πF, πClass, nil, ßproperty); πE != nil {
						continue
					}
					if πTemp014, πE = πTemp013.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πE = πClass.SetItem(πF, ßmonth.ToObject(), πTemp014); πE != nil {
						continue
					}
					// line 823: def day(self):
					πF.SetLineno(823)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp013 = πg.NewFunction(πg.NewCode("day", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 824: """day (1-31)"""
							πF.SetLineno(824)
							// line 825: return self._day
							πF.SetLineno(825)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_day, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßday.ToObject(), πTemp013); πE != nil {
						continue
					}
					// line 824: """day (1-31)"""
					πF.SetLineno(824)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp014}, πg.NewStr("day (1-31)").ToObject()); πE != nil {
						continue
					}
					if πTemp015, πE = πg.ResolveClass(πF, πClass, nil, ßday); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp015, ß__doc__, πTemp014); πE != nil {
						continue
					}
					// line 822: @property
					πF.SetLineno(822)
					πTemp006 = πF.MakeArgs(1)
					if πTemp014, πE = πg.ResolveClass(πF, πClass, nil, ßday); πE != nil {
						continue
					}
					πTemp006[0] = πTemp014
					if πTemp014, πE = πg.ResolveClass(πF, πClass, nil, ßproperty); πE != nil {
						continue
					}
					if πTemp015, πE = πTemp014.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πE = πClass.SetItem(πF, ßday.ToObject(), πTemp015); πE != nil {
						continue
					}
					// line 829: def timetuple(self):
					πF.SetLineno(829)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp014 = πg.NewFunction(πg.NewCode("timetuple", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 830: "Return local time tuple compatible with time.localtime()."
							πF.SetLineno(830)
							// line 831: return _build_struct_time(self._year, self._month, self._day,
							πF.SetLineno(831)
							πTemp001 = πF.MakeArgs(7)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_year, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_month, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_day, nil); πE != nil {
								continue
							}
							πTemp001[2] = πTemp002
							πTemp001[3] = πg.NewInt(0).ToObject()
							πTemp001[4] = πg.NewInt(0).ToObject()
							πTemp001[5] = πg.NewInt(0).ToObject()
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001[6] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_build_struct_time); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtimetuple.ToObject(), πTemp014); πE != nil {
						continue
					}
					// line 830: "Return local time tuple compatible with time.localtime()."
					πF.SetLineno(830)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp015}, πg.NewStr("Return local time tuple compatible with time.localtime().").ToObject()); πE != nil {
						continue
					}
					if πTemp016, πE = πg.ResolveClass(πF, πClass, nil, ßtimetuple); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp016, ß__doc__, πTemp015); πE != nil {
						continue
					}
					// line 834: def toordinal(self):
					πF.SetLineno(834)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp015 = πg.NewFunction(πg.NewCode("toordinal", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 835: """Return proleptic Gregorian ordinal for the year, month and day.
							πF.SetLineno(835)
							// line 840: return _ymd2ord(self._year, self._month, self._day)
							πF.SetLineno(840)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_year, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_month, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_day, nil); πE != nil {
								continue
							}
							πTemp001[2] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_ymd2ord); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtoordinal.ToObject(), πTemp015); πE != nil {
						continue
					}
					// line 835: """Return proleptic Gregorian ordinal for the year, month and day.
					πF.SetLineno(835)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp016}, πg.NewStr("Return proleptic Gregorian ordinal for the year, month and day.\n\n        January 1 of year 1 is day 1.  Only the year, month and day values\n        contribute to the result.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp017, πE = πg.ResolveClass(πF, πClass, nil, ßtoordinal); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp017, ß__doc__, πTemp016); πE != nil {
						continue
					}
					// line 842: def replace(self, year=None, month=None, day=None):
					πF.SetLineno(842)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					if πTemp017, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[1] = πg.Param{Name: "year", Def: πTemp017}
					if πTemp017, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "month", Def: πTemp017}
					if πTemp017, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[3] = πg.Param{Name: "day", Def: πTemp017}
					πTemp016 = πg.NewFunction(πg.NewCode("replace", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µyear *πg.Object = πArgs[1]; _ = µyear
						var µmonth *πg.Object = πArgs[2]; _ = µmonth
						var µday *πg.Object = πArgs[3]; _ = µday
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
							// line 843: """Return a new date with new values for the specified fields."""
							πF.SetLineno(843)
							if πE = πg.CheckLocal(πF, µyear, "year"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µyear == πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 844: if year is None:
							πF.SetLineno(844)
						Label1:
							// line 845: year = self._year
							πF.SetLineno(845)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_year, nil); πE != nil {
								continue
							}
							µyear = πTemp001
							goto Label2
						Label2:
							if πE = πg.CheckLocal(πF, µmonth, "month"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µmonth == πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label3
							}
							goto Label4
							// line 846: if month is None:
							πF.SetLineno(846)
						Label3:
							// line 847: month = self._month
							πF.SetLineno(847)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_month, nil); πE != nil {
								continue
							}
							µmonth = πTemp001
							goto Label4
						Label4:
							if πE = πg.CheckLocal(πF, µday, "day"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µday == πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label5
							}
							goto Label6
							// line 848: if day is None:
							πF.SetLineno(848)
						Label5:
							// line 849: day = self._day
							πF.SetLineno(849)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_day, nil); πE != nil {
								continue
							}
							µday = πTemp001
							goto Label6
						Label6:
							// line 850: return date.__new__(type(self), year, month, day)
							πF.SetLineno(850)
							πTemp004 = πF.MakeArgs(4)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp005[0] = µself
							if πTemp001, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp004[0] = πTemp002
							if πE = πg.CheckLocal(πF, µyear, "year"); πE != nil {
								continue
							}
							πTemp004[1] = µyear
							if πE = πg.CheckLocal(πF, µmonth, "month"); πE != nil {
								continue
							}
							πTemp004[2] = µmonth
							if πE = πg.CheckLocal(πF, µday, "day"); πE != nil {
								continue
							}
							πTemp004[3] = µday
							if πTemp001, πE = πg.ResolveGlobal(πF, ßdate); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ß__new__, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
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
					if πE = πClass.SetItem(πF, ßreplace.ToObject(), πTemp016); πE != nil {
						continue
					}
					// line 843: """Return a new date with new values for the specified fields."""
					πF.SetLineno(843)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp017}, πg.NewStr("Return a new date with new values for the specified fields.").ToObject()); πE != nil {
						continue
					}
					if πTemp018, πE = πg.ResolveClass(πF, πClass, nil, ßreplace); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp018, ß__doc__, πTemp017); πE != nil {
						continue
					}
					// line 854: def __eq__(self, other):
					πF.SetLineno(854)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "other", Def: nil}
					πTemp017 = πg.NewFunction(πg.NewCode("__eq__", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
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
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							if πTemp002, πE = πg.ResolveGlobal(πF, ßdate); πE != nil {
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
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							πTemp001[1] = ßtimetuple.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
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
							// line 855: if isinstance(other, date):
							πF.SetLineno(855)
						Label1:
							// line 856: return self._cmp(other) == 0
							πF.SetLineno(856)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_cmp, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.Eq(πF, πTemp005, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							πR = πTemp002
							continue
							goto Label4
							// line 857: elif hasattr(other, "timetuple"):
							πF.SetLineno(857)
						Label2:
							// line 858: return NotImplemented
							πF.SetLineno(858)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNotImplemented); πE != nil {
								continue
							}
							πR = πTemp002
							continue
							goto Label4
						Label3:
							// line 860: return False
							πF.SetLineno(860)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πR = πTemp002
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
					if πE = πClass.SetItem(πF, ß__eq__.ToObject(), πTemp017); πE != nil {
						continue
					}
					// line 862: def __ne__(self, other):
					πF.SetLineno(862)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "other", Def: nil}
					πTemp018 = πg.NewFunction(πg.NewCode("__ne__", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
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
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							if πTemp002, πE = πg.ResolveGlobal(πF, ßdate); πE != nil {
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
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							πTemp001[1] = ßtimetuple.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
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
							// line 863: if isinstance(other, date):
							πF.SetLineno(863)
						Label1:
							// line 864: return self._cmp(other) != 0
							πF.SetLineno(864)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_cmp, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.NE(πF, πTemp005, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							πR = πTemp002
							continue
							goto Label4
							// line 865: elif hasattr(other, "timetuple"):
							πF.SetLineno(865)
						Label2:
							// line 866: return NotImplemented
							πF.SetLineno(866)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNotImplemented); πE != nil {
								continue
							}
							πR = πTemp002
							continue
							goto Label4
						Label3:
							// line 868: return True
							πF.SetLineno(868)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πR = πTemp002
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
					if πE = πClass.SetItem(πF, ß__ne__.ToObject(), πTemp018); πE != nil {
						continue
					}
					// line 870: def __le__(self, other):
					πF.SetLineno(870)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "other", Def: nil}
					πTemp019 = πg.NewFunction(πg.NewCode("__le__", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
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
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							if πTemp002, πE = πg.ResolveGlobal(πF, ßdate); πE != nil {
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
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							πTemp001[1] = ßtimetuple.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
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
							// line 871: if isinstance(other, date):
							πF.SetLineno(871)
						Label1:
							// line 872: return self._cmp(other) <= 0
							πF.SetLineno(872)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_cmp, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.LE(πF, πTemp005, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							πR = πTemp002
							continue
							goto Label4
							// line 873: elif hasattr(other, "timetuple"):
							πF.SetLineno(873)
						Label2:
							// line 874: return NotImplemented
							πF.SetLineno(874)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNotImplemented); πE != nil {
								continue
							}
							πR = πTemp002
							continue
							goto Label4
						Label3:
							// line 876: _cmperror(self, other)
							πF.SetLineno(876)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[1] = µother
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_cmperror); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
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
					if πE = πClass.SetItem(πF, ß__le__.ToObject(), πTemp019); πE != nil {
						continue
					}
					// line 878: def __lt__(self, other):
					πF.SetLineno(878)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "other", Def: nil}
					πTemp020 = πg.NewFunction(πg.NewCode("__lt__", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
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
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							if πTemp002, πE = πg.ResolveGlobal(πF, ßdate); πE != nil {
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
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							πTemp001[1] = ßtimetuple.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
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
							// line 879: if isinstance(other, date):
							πF.SetLineno(879)
						Label1:
							// line 880: return self._cmp(other) < 0
							πF.SetLineno(880)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_cmp, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.LT(πF, πTemp005, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							πR = πTemp002
							continue
							goto Label4
							// line 881: elif hasattr(other, "timetuple"):
							πF.SetLineno(881)
						Label2:
							// line 882: return NotImplemented
							πF.SetLineno(882)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNotImplemented); πE != nil {
								continue
							}
							πR = πTemp002
							continue
							goto Label4
						Label3:
							// line 884: _cmperror(self, other)
							πF.SetLineno(884)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[1] = µother
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_cmperror); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
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
					if πE = πClass.SetItem(πF, ß__lt__.ToObject(), πTemp020); πE != nil {
						continue
					}
					// line 886: def __ge__(self, other):
					πF.SetLineno(886)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "other", Def: nil}
					πTemp021 = πg.NewFunction(πg.NewCode("__ge__", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
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
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							if πTemp002, πE = πg.ResolveGlobal(πF, ßdate); πE != nil {
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
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							πTemp001[1] = ßtimetuple.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
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
							// line 887: if isinstance(other, date):
							πF.SetLineno(887)
						Label1:
							// line 888: return self._cmp(other) >= 0
							πF.SetLineno(888)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_cmp, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.GE(πF, πTemp005, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							πR = πTemp002
							continue
							goto Label4
							// line 889: elif hasattr(other, "timetuple"):
							πF.SetLineno(889)
						Label2:
							// line 890: return NotImplemented
							πF.SetLineno(890)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNotImplemented); πE != nil {
								continue
							}
							πR = πTemp002
							continue
							goto Label4
						Label3:
							// line 892: _cmperror(self, other)
							πF.SetLineno(892)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[1] = µother
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_cmperror); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
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
					if πE = πClass.SetItem(πF, ß__ge__.ToObject(), πTemp021); πE != nil {
						continue
					}
					// line 894: def __gt__(self, other):
					πF.SetLineno(894)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "other", Def: nil}
					πTemp022 = πg.NewFunction(πg.NewCode("__gt__", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
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
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							if πTemp002, πE = πg.ResolveGlobal(πF, ßdate); πE != nil {
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
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							πTemp001[1] = ßtimetuple.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
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
							// line 895: if isinstance(other, date):
							πF.SetLineno(895)
						Label1:
							// line 896: return self._cmp(other) > 0
							πF.SetLineno(896)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_cmp, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.GT(πF, πTemp005, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							πR = πTemp002
							continue
							goto Label4
							// line 897: elif hasattr(other, "timetuple"):
							πF.SetLineno(897)
						Label2:
							// line 898: return NotImplemented
							πF.SetLineno(898)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNotImplemented); πE != nil {
								continue
							}
							πR = πTemp002
							continue
							goto Label4
						Label3:
							// line 900: _cmperror(self, other)
							πF.SetLineno(900)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[1] = µother
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_cmperror); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
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
					if πE = πClass.SetItem(πF, ß__gt__.ToObject(), πTemp022); πE != nil {
						continue
					}
					// line 902: def _cmp(self, other):
					πF.SetLineno(902)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "other", Def: nil}
					πTemp023 = πg.NewFunction(πg.NewCode("_cmp", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
						var µy *πg.Object = πg.UnboundLocal; _ = µy
						var µm *πg.Object = πg.UnboundLocal; _ = µm
						var µd *πg.Object = πg.UnboundLocal; _ = µd
						var µy2 *πg.Object = πg.UnboundLocal; _ = µy2
						var µm2 *πg.Object = πg.UnboundLocal; _ = µm2
						var µd2 *πg.Object = πg.UnboundLocal; _ = µd2
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 903: assert isinstance(other, date)
							πF.SetLineno(903)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							if πTemp002, πE = πg.ResolveGlobal(πF, ßdate); πE != nil {
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
							if πE = πg.Assert(πF, πTemp003, nil); πE != nil {
								continue
							}
							// line 904: y, m, d = self._year, self._month, self._day
							πF.SetLineno(904)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_year, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß_month, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ß_day, nil); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple3(πTemp003, πTemp004, πTemp005).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}}}, πTemp002); πE != nil {
								continue
							}
							µy = πTemp003
							µm = πTemp004
							µd = πTemp005
							// line 905: y2, m2, d2 = other._year, other._month, other._day
							πF.SetLineno(905)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µother, ß_year, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µother, ß_month, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µother, ß_day, nil); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple3(πTemp003, πTemp004, πTemp005).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}}}, πTemp002); πE != nil {
								continue
							}
							µy2 = πTemp003
							µm2 = πTemp004
							µd2 = πTemp005
							// line 906: return _cmp((y, m, d), (y2, m2, d2))
							πF.SetLineno(906)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µm, "m"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple3(µy, µm, µd).ToObject()
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µy2, "y2"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µm2, "m2"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µd2, "d2"); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple3(µy2, µm2, µd2).ToObject()
							πTemp001[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_cmp); πE != nil {
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
					if πE = πClass.SetItem(πF, ß_cmp.ToObject(), πTemp023); πE != nil {
						continue
					}
					// line 908: def __hash__(self):
					πF.SetLineno(908)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp024 = πg.NewFunction(πg.NewCode("__hash__", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 909: "Hash."
							πF.SetLineno(909)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_hashcode, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
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
							// line 910: if self._hashcode == -1:
							πF.SetLineno(910)
						Label1:
							// line 911: self._hashcode = hash(self._getstate())
							πF.SetLineno(911)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_getstate, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp005[0] = πTemp002
							if πTemp001, πE = πg.ResolveGlobal(πF, ßhash); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_hashcode, πTemp001); πE != nil {
								continue
							}
							goto Label2
						Label2:
							// line 912: return self._hashcode
							πF.SetLineno(912)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_hashcode, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__hash__.ToObject(), πTemp024); πE != nil {
						continue
					}
					// line 909: "Hash."
					πF.SetLineno(909)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp025}, πg.NewStr("Hash.").ToObject()); πE != nil {
						continue
					}
					if πTemp026, πE = πg.ResolveClass(πF, πClass, nil, ß__hash__); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp026, ß__doc__, πTemp025); πE != nil {
						continue
					}
					// line 916: def _add_timedelta(self, other, factor):
					πF.SetLineno(916)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "other", Def: nil}
					πTemp002[2] = πg.Param{Name: "factor", Def: nil}
					πTemp025 = πg.NewFunction(πg.NewCode("_add_timedelta", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
						var µfactor *πg.Object = πArgs[2]; _ = µfactor
						var µy *πg.Object = πg.UnboundLocal; _ = µy
						var µm *πg.Object = πg.UnboundLocal; _ = µm
						var µd *πg.Object = πg.UnboundLocal; _ = µd
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 917: y, m, d = _normalize_date(
							πF.SetLineno(917)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_year, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_month, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_day, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µother, ßdays, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µfactor, "factor"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Mul(πF, πTemp005, µfactor); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp003, πTemp004); πE != nil {
								continue
							}
							πTemp001[2] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_normalize_date); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}}}, πTemp003); πE != nil {
								continue
							}
							µy = πTemp002
							µm = πTemp004
							µd = πTemp005
							// line 921: return date(y, m, d)
							πF.SetLineno(921)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
								continue
							}
							πTemp001[0] = µy
							if πE = πg.CheckLocal(πF, µm, "m"); πE != nil {
								continue
							}
							πTemp001[1] = µm
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp001[2] = µd
							if πTemp002, πE = πg.ResolveGlobal(πF, ßdate); πE != nil {
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
					if πE = πClass.SetItem(πF, ß_add_timedelta.ToObject(), πTemp025); πE != nil {
						continue
					}
					// line 923: def __add__(self, other):
					πF.SetLineno(923)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "other", Def: nil}
					πTemp026 = πg.NewFunction(πg.NewCode("__add__", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
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
							// line 924: "Add a date to a timedelta."
							πF.SetLineno(924)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtimedelta); πE != nil {
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
							goto Label2
							// line 925: if isinstance(other, timedelta):
							πF.SetLineno(925)
						Label1:
							// line 926: return self._add_timedelta(other, 1)
							πF.SetLineno(926)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							πTemp001[1] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_add_timedelta, nil); πE != nil {
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
							// line 927: return NotImplemented
							πF.SetLineno(927)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNotImplemented); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__add__.ToObject(), πTemp026); πE != nil {
						continue
					}
					// line 924: "Add a date to a timedelta."
					πF.SetLineno(924)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp027}, πg.NewStr("Add a date to a timedelta.").ToObject()); πE != nil {
						continue
					}
					if πTemp028, πE = πg.ResolveClass(πF, πClass, nil, ß__add__); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp028, ß__doc__, πTemp027); πE != nil {
						continue
					}
					// line 929: __radd__ = __add__
					πF.SetLineno(929)
					if πTemp027, πE = πg.ResolveClass(πF, πClass, nil, ß__add__); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ß__radd__.ToObject(), πTemp027); πE != nil {
						continue
					}
					// line 931: def __sub__(self, other):
					πF.SetLineno(931)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "other", Def: nil}
					πTemp027 = πg.NewFunction(πg.NewCode("__sub__", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
						var µdays1 *πg.Object = πg.UnboundLocal; _ = µdays1
						var µdays2 *πg.Object = πg.UnboundLocal; _ = µdays2
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
							// line 932: """Subtract two dates, or a date and a timedelta."""
							πF.SetLineno(932)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							if πTemp002, πE = πg.ResolveGlobal(πF, ßdate); πE != nil {
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
							goto Label2
							// line 933: if isinstance(other, date):
							πF.SetLineno(933)
						Label1:
							// line 934: days1 = self.toordinal()
							πF.SetLineno(934)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtoordinal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µdays1 = πTemp003
							// line 935: days2 = other.toordinal()
							πF.SetLineno(935)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µother, ßtoordinal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µdays2 = πTemp003
							// line 936: return timedelta._create(days1 - days2, 0, 0, False)
							πF.SetLineno(936)
							πTemp001 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µdays1, "days1"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µdays2, "days2"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Sub(πF, µdays1, µdays2); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = πg.NewInt(0).ToObject()
							πTemp001[2] = πg.NewInt(0).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp001[3] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtimedelta); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ß_create, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πR = πTemp002
							continue
							goto Label2
						Label2:
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtimedelta); πE != nil {
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
								goto Label3
							}
							goto Label4
							// line 937: if isinstance(other, timedelta):
							πF.SetLineno(937)
						Label3:
							// line 938: return self._add_timedelta(other, -1)
							πF.SetLineno(938)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_add_timedelta, nil); πE != nil {
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
							// line 939: return NotImplemented
							πF.SetLineno(939)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNotImplemented); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__sub__.ToObject(), πTemp027); πE != nil {
						continue
					}
					// line 932: """Subtract two dates, or a date and a timedelta."""
					πF.SetLineno(932)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp028}, πg.NewStr("Subtract two dates, or a date and a timedelta.").ToObject()); πE != nil {
						continue
					}
					if πTemp029, πE = πg.ResolveClass(πF, πClass, nil, ß__sub__); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp029, ß__doc__, πTemp028); πE != nil {
						continue
					}
					// line 941: def weekday(self):
					πF.SetLineno(941)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp028 = πg.NewFunction(πg.NewCode("weekday", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πTemp001 *πg.Object
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
							// line 942: "Return day of the week, where Monday == 0 ... Sunday == 6."
							πF.SetLineno(942)
							// line 943: return (self.toordinal() + 6) % 7
							πF.SetLineno(943)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtoordinal, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp004, πg.NewInt(6).ToObject()); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πTemp002, πg.NewInt(7).ToObject()); πE != nil {
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
					if πE = πClass.SetItem(πF, ßweekday.ToObject(), πTemp028); πE != nil {
						continue
					}
					// line 942: "Return day of the week, where Monday == 0 ... Sunday == 6."
					πF.SetLineno(942)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp029}, πg.NewStr("Return day of the week, where Monday == 0 ... Sunday == 6.").ToObject()); πE != nil {
						continue
					}
					if πTemp030, πE = πg.ResolveClass(πF, πClass, nil, ßweekday); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp030, ß__doc__, πTemp029); πE != nil {
						continue
					}
					// line 947: def isoweekday(self):
					πF.SetLineno(947)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp029 = πg.NewFunction(πg.NewCode("isoweekday", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 948: "Return day of the week, where Monday == 1 ... Sunday == 7."
							πF.SetLineno(948)
							// line 950: return self.toordinal() % 7 or 7
							πF.SetLineno(950)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßtoordinal, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Mod(πF, πTemp005, πg.NewInt(7).ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp003
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							πTemp001 = πg.NewInt(7).ToObject()
						Label1:
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
					if πE = πClass.SetItem(πF, ßisoweekday.ToObject(), πTemp029); πE != nil {
						continue
					}
					// line 948: "Return day of the week, where Monday == 1 ... Sunday == 7."
					πF.SetLineno(948)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp030}, πg.NewStr("Return day of the week, where Monday == 1 ... Sunday == 7.").ToObject()); πE != nil {
						continue
					}
					if πTemp031, πE = πg.ResolveClass(πF, πClass, nil, ßisoweekday); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp031, ß__doc__, πTemp030); πE != nil {
						continue
					}
					// line 952: def isocalendar(self):
					πF.SetLineno(952)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp030 = πg.NewFunction(πg.NewCode("isocalendar", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µyear *πg.Object = πg.UnboundLocal; _ = µyear
						var µweek1monday *πg.Object = πg.UnboundLocal; _ = µweek1monday
						var µtoday *πg.Object = πg.UnboundLocal; _ = µtoday
						var µweek *πg.Object = πg.UnboundLocal; _ = µweek
						var µday *πg.Object = πg.UnboundLocal; _ = µday
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 953: """Return a 3-tuple containing ISO year, week number, and weekday.
							πF.SetLineno(953)
							// line 964: year = self._year
							πF.SetLineno(964)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_year, nil); πE != nil {
								continue
							}
							µyear = πTemp001
							// line 965: week1monday = _isoweek1monday(year)
							πF.SetLineno(965)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µyear, "year"); πE != nil {
								continue
							}
							πTemp002[0] = µyear
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_isoweek1monday); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µweek1monday = πTemp003
							// line 966: today = _ymd2ord(self._year, self._month, self._day)
							πF.SetLineno(966)
							πTemp002 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_year, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_month, nil); πE != nil {
								continue
							}
							πTemp002[1] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_day, nil); πE != nil {
								continue
							}
							πTemp002[2] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_ymd2ord); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µtoday = πTemp003
							// line 968: week, day = divmod(today - week1monday, 7)
							πF.SetLineno(968)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µtoday, "today"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µweek1monday, "week1monday"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Sub(πF, µtoday, µweek1monday); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							πTemp002[1] = πg.NewInt(7).ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßdivmod); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp004}}}, πTemp003); πE != nil {
								continue
							}
							µweek = πTemp001
							µday = πTemp004
							if πE = πg.CheckLocal(πF, µweek, "week"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.LT(πF, µweek, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µweek, "week"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GE(πF, µweek, πg.NewInt(52).ToObject()); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label2
							}
							goto Label3
							// line 969: if week < 0:
							πF.SetLineno(969)
						Label1:
							// line 970: year -= 1
							πF.SetLineno(970)
							if πE = πg.CheckLocal(πF, µyear, "year"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.ISub(πF, µyear, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µyear = πTemp001
							// line 971: week1monday = _isoweek1monday(year)
							πF.SetLineno(971)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µyear, "year"); πE != nil {
								continue
							}
							πTemp002[0] = µyear
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_isoweek1monday); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µweek1monday = πTemp003
							// line 972: week, day = divmod(today - week1monday, 7)
							πF.SetLineno(972)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µtoday, "today"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µweek1monday, "week1monday"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Sub(πF, µtoday, µweek1monday); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							πTemp002[1] = πg.NewInt(7).ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßdivmod); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp004}}}, πTemp003); πE != nil {
								continue
							}
							µweek = πTemp001
							µday = πTemp004
							goto Label3
							// line 973: elif week >= 52:
							πF.SetLineno(973)
						Label2:
							if πE = πg.CheckLocal(πF, µtoday, "today"); πE != nil {
								continue
							}
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µyear, "year"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, µyear, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ß_isoweek1monday); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp001, πE = πg.GE(πF, µtoday, πTemp004); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label4
							}
							goto Label5
							// line 974: if today >= _isoweek1monday(year+1):
							πF.SetLineno(974)
						Label4:
							// line 975: year += 1
							πF.SetLineno(975)
							if πE = πg.CheckLocal(πF, µyear, "year"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.IAdd(πF, µyear, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µyear = πTemp001
							// line 976: week = 0
							πF.SetLineno(976)
							µweek = πg.NewInt(0).ToObject()
							goto Label5
						Label5:
							goto Label3
						Label3:
							// line 977: return year, week+1, day+1
							πF.SetLineno(977)
							if πE = πg.CheckLocal(πF, µyear, "year"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µweek, "week"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, µweek, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µday, "day"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Add(πF, µday, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple3(µyear, πTemp003, πTemp004).ToObject()
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
					if πE = πClass.SetItem(πF, ßisocalendar.ToObject(), πTemp030); πE != nil {
						continue
					}
					// line 953: """Return a 3-tuple containing ISO year, week number, and weekday.
					πF.SetLineno(953)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp031}, πg.NewStr("Return a 3-tuple containing ISO year, week number, and weekday.\n\n        The first ISO week of the year is the (Mon-Sun) week\n        containing the year's first Thursday; everything else derives\n        from that.\n\n        The first week is 1; Monday is 1 ... Sunday is 7.\n\n        ISO calendar algorithm taken from\n        http://www.phys.uu.nl/~vgent/calendar/isocalendar.htm\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp032, πE = πg.ResolveClass(πF, πClass, nil, ßisocalendar); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp032, ß__doc__, πTemp031); πE != nil {
						continue
					}
					// line 981: def _getstate(self):
					πF.SetLineno(981)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp031 = πg.NewFunction(πg.NewCode("_getstate", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µyhi *πg.Object = πg.UnboundLocal; _ = µyhi
						var µylo *πg.Object = πg.UnboundLocal; _ = µylo
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
							// line 982: yhi, ylo = divmod(self._year, 256)
							πF.SetLineno(982)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_year, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = πg.NewInt(256).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßdivmod); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp004}}}, πTemp003); πE != nil {
								continue
							}
							µyhi = πTemp002
							µylo = πTemp004
							// line 983: return (_struct.pack('4B', yhi, ylo, self._month, self._day),)
							πF.SetLineno(983)
							πTemp001 = πF.MakeArgs(5)
							πTemp001[0] = ß4B.ToObject()
							if πE = πg.CheckLocal(πF, µyhi, "yhi"); πE != nil {
								continue
							}
							πTemp001[1] = µyhi
							if πE = πg.CheckLocal(πF, µylo, "ylo"); πE != nil {
								continue
							}
							πTemp001[2] = µylo
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_month, nil); πE != nil {
								continue
							}
							πTemp001[3] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_day, nil); πE != nil {
								continue
							}
							πTemp001[4] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ß_struct); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßpack, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp002 = πg.NewTuple1(πTemp003).ToObject()
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
					if πE = πClass.SetItem(πF, ß_getstate.ToObject(), πTemp031); πE != nil {
						continue
					}
					// line 985: def __setstate(self, string):
					πF.SetLineno(985)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "string", Def: nil}
					πTemp032 = πg.NewFunction(πg.NewCode("__setstate", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µstring *πg.Object = πArgs[1]; _ = µstring
						var µyhi *πg.Object = πg.UnboundLocal; _ = µyhi
						var µylo *πg.Object = πg.UnboundLocal; _ = µylo
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
							// line 986: yhi, ylo, self._month, self._day = (ord(string[0]), ord(string[1]),
							πF.SetLineno(986)
							πTemp002 = πF.MakeArgs(1)
							πTemp003 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µstring, πTemp003); πE != nil {
								continue
							}
							πTemp002[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp002 = πF.MakeArgs(1)
							πTemp003 = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µstring, πTemp003); πE != nil {
								continue
							}
							πTemp002[0] = πTemp005
							if πTemp003, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp002 = πF.MakeArgs(1)
							πTemp003 = πg.NewInt(2).ToObject()
							if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µstring, πTemp003); πE != nil {
								continue
							}
							πTemp002[0] = πTemp006
							if πTemp003, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp002 = πF.MakeArgs(1)
							πTemp003 = πg.NewInt(3).ToObject()
							if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, µstring, πTemp003); πE != nil {
								continue
							}
							πTemp002[0] = πTemp007
							if πTemp003, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001 = πg.NewTuple4(πTemp004, πTemp005, πTemp006, πTemp007).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp006}}}, πTemp001); πE != nil {
								continue
							}
							µyhi = πTemp003
							µylo = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_month, πTemp005); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_day, πTemp006); πE != nil {
								continue
							}
							// line 988: self._year = yhi * 256 + ylo
							πF.SetLineno(988)
							if πE = πg.CheckLocal(πF, µyhi, "yhi"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Mul(πF, µyhi, πg.NewInt(256).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µylo, "ylo"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp003, µylo); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_year, πTemp003); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__setstate.ToObject(), πTemp032); πE != nil {
						continue
					}
					// line 990: def __reduce__(self):
					πF.SetLineno(990)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp033 = πg.NewFunction(πg.NewCode("__reduce__", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πTemp001 *πg.Object
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
							// line 991: return (self.__class__, self._getstate())
							πF.SetLineno(991)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß__class__, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_getstate, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple2(πTemp002, πTemp004).ToObject()
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
					if πE = πClass.SetItem(πF, ß__reduce__.ToObject(), πTemp033); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp029, πE = πTemp027.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp029 == nil {
				πTemp029 = πg.TypeType.ToObject()
			}
			if πTemp030, πE = πTemp029.Call(πF, []*πg.Object{πg.NewStr("date").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp027.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßdate.ToObject(), πTemp030); πE != nil {
				continue
			}
			// line 993: _date_class = date  # so functions w/ args named "date" can get at the class
			πF.SetLineno(993)
			if πTemp028, πE = πg.ResolveGlobal(πF, ßdate); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_date_class.ToObject(), πTemp028); πE != nil {
				continue
			}
			// line 995: date.min = date(1, 1, 1)
			πF.SetLineno(995)
			πTemp002 = πF.MakeArgs(3)
			πTemp002[0] = πg.NewInt(1).ToObject()
			πTemp002[1] = πg.NewInt(1).ToObject()
			πTemp002[2] = πg.NewInt(1).ToObject()
			if πTemp028, πE = πg.ResolveGlobal(πF, ßdate); πE != nil {
				continue
			}
			if πTemp029, πE = πTemp028.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp028}, πTemp029); πE != nil {
				continue
			}
			if πTemp030, πE = πg.ResolveGlobal(πF, ßdate); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp030, ßmin, πTemp028); πE != nil {
				continue
			}
			// line 996: date.max = date(9999, 12, 31)
			πF.SetLineno(996)
			πTemp002 = πF.MakeArgs(3)
			πTemp002[0] = πg.NewInt(9999).ToObject()
			πTemp002[1] = πg.NewInt(12).ToObject()
			πTemp002[2] = πg.NewInt(31).ToObject()
			if πTemp028, πE = πg.ResolveGlobal(πF, ßdate); πE != nil {
				continue
			}
			if πTemp029, πE = πTemp028.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp028}, πTemp029); πE != nil {
				continue
			}
			if πTemp030, πE = πg.ResolveGlobal(πF, ßdate); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp030, ßmax, πTemp028); πE != nil {
				continue
			}
			// line 997: date.resolution = timedelta(days=1)
			πF.SetLineno(997)
			πTemp031 = πg.KWArgs{
				{"days", πg.NewInt(1).ToObject()},
			}
			if πTemp028, πE = πg.ResolveGlobal(πF, ßtimedelta); πE != nil {
				continue
			}
			if πTemp029, πE = πTemp028.Call(πF, nil, πTemp031); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp028}, πTemp029); πE != nil {
				continue
			}
			if πTemp030, πE = πg.ResolveGlobal(πF, ßdate); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp030, ßresolution, πTemp028); πE != nil {
				continue
			}
			// line 999: class tzinfo(object):
			πF.SetLineno(999)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp030, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp030
			πTemp027 = πg.NewDict()
			if πTemp028, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp027.SetItem(πF, ß__module__.ToObject(), πTemp028); πE != nil {
				continue
			}
			_, πE = πg.NewCode("tzinfo", "/usr/lib/python2.7/datetime.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp027
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 1000: """Abstract base class for time zone info classes.
					πF.SetLineno(1000)
					// line 1000: """Abstract base class for time zone info classes.
					πF.SetLineno(1000)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Abstract base class for time zone info classes.\n\n    Subclasses must override the name(), utcoffset() and dst() methods.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 1004: __slots__ = ()
					πF.SetLineno(1004)
					πTemp001 = πg.NewTuple0().ToObject()
					if πE = πClass.SetItem(πF, ß__slots__.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 1006: def tzname(self, dt):
					πF.SetLineno(1006)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "dt", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("tzname", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µdt *πg.Object = πArgs[1]; _ = µdt
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
							// line 1007: "datetime -> string name of time zone."
							πF.SetLineno(1007)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("tzinfo subclass must override tzname()").ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNotImplementedError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1008: raise NotImplementedError("tzinfo subclass must override tzname()")
							πF.SetLineno(1008)
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
					if πE = πClass.SetItem(πF, ßtzname.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 1007: "datetime -> string name of time zone."
					πF.SetLineno(1007)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewStr("datetime -> string name of time zone.").ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßtzname); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp004, ß__doc__, πTemp003); πE != nil {
						continue
					}
					// line 1010: def utcoffset(self, dt):
					πF.SetLineno(1010)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "dt", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("utcoffset", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µdt *πg.Object = πArgs[1]; _ = µdt
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
							// line 1011: "datetime -> minutes east of UTC (negative for west of UTC)"
							πF.SetLineno(1011)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("tzinfo subclass must override utcoffset()").ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNotImplementedError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1012: raise NotImplementedError("tzinfo subclass must override utcoffset()")
							πF.SetLineno(1012)
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
					if πE = πClass.SetItem(πF, ßutcoffset.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 1011: "datetime -> minutes east of UTC (negative for west of UTC)"
					πF.SetLineno(1011)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πg.NewStr("datetime -> minutes east of UTC (negative for west of UTC)").ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßutcoffset); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp005, ß__doc__, πTemp004); πE != nil {
						continue
					}
					// line 1014: def dst(self, dt):
					πF.SetLineno(1014)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "dt", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("dst", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µdt *πg.Object = πArgs[1]; _ = µdt
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
							// line 1015: """datetime -> DST offset in minutes east of UTC.
							πF.SetLineno(1015)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("tzinfo subclass must override dst()").ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNotImplementedError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1020: raise NotImplementedError("tzinfo subclass must override dst()")
							πF.SetLineno(1020)
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
					if πE = πClass.SetItem(πF, ßdst.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 1015: """datetime -> DST offset in minutes east of UTC.
					πF.SetLineno(1015)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πg.NewStr("datetime -> DST offset in minutes east of UTC.\n\n        Return 0 if DST not in effect.  utcoffset() must include the DST\n        offset.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßdst); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp006, ß__doc__, πTemp005); πE != nil {
						continue
					}
					// line 1022: def fromutc(self, dt):
					πF.SetLineno(1022)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "dt", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("fromutc", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µdt *πg.Object = πArgs[1]; _ = µdt
						var µdtoff *πg.Object = πg.UnboundLocal; _ = µdtoff
						var µdtdst *πg.Object = πg.UnboundLocal; _ = µdtdst
						var µdelta *πg.Object = πg.UnboundLocal; _ = µdelta
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 1023: "datetime in UTC -> datetime in local time."
							πF.SetLineno(1023)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µdt, "dt"); πE != nil {
								continue
							}
							πTemp002[0] = µdt
							if πTemp003, πE = πg.ResolveGlobal(πF, ßdatetime); πE != nil {
								continue
							}
							πTemp002[1] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp005, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp005).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							goto Label2
							// line 1025: if not isinstance(dt, datetime):
							πF.SetLineno(1025)
						Label1:
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("fromutc() requires a datetime argument").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 1026: raise TypeError("fromutc() requires a datetime argument")
							πF.SetLineno(1026)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label2
						Label2:
							if πE = πg.CheckLocal(πF, µdt, "dt"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µdt, ßtzinfo, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp003 != µself).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label3
							}
							goto Label4
							// line 1027: if dt.tzinfo is not self:
							πF.SetLineno(1027)
						Label3:
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("dt.tzinfo is not self").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 1028: raise ValueError("dt.tzinfo is not self")
							πF.SetLineno(1028)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label4
						Label4:
							// line 1030: dtoff = dt.utcoffset()
							πF.SetLineno(1030)
							if πE = πg.CheckLocal(πF, µdt, "dt"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µdt, ßutcoffset, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µdtoff = πTemp003
							if πE = πg.CheckLocal(πF, µdtoff, "dtoff"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µdtoff == πTemp003).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label5
							}
							goto Label6
							// line 1031: if dtoff is None:
							πF.SetLineno(1031)
						Label5:
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("fromutc() requires a non-None utcoffset() result").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 1032: raise ValueError("fromutc() requires a non-None utcoffset() "
							πF.SetLineno(1032)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label6
						Label6:
							// line 1037: dtdst = dt.dst()
							πF.SetLineno(1037)
							if πE = πg.CheckLocal(πF, µdt, "dt"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µdt, ßdst, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µdtdst = πTemp003
							if πE = πg.CheckLocal(πF, µdtdst, "dtdst"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µdtdst == πTemp003).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label7
							}
							goto Label8
							// line 1038: if dtdst is None:
							πF.SetLineno(1038)
						Label7:
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("fromutc() requires a non-None dst() result").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 1039: raise ValueError("fromutc() requires a non-None dst() result")
							πF.SetLineno(1039)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label8
						Label8:
							// line 1040: delta = dtoff - dtdst
							πF.SetLineno(1040)
							if πE = πg.CheckLocal(πF, µdtoff, "dtoff"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µdtdst, "dtdst"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Sub(πF, µdtoff, µdtdst); πE != nil {
								continue
							}
							µdelta = πTemp001
							if πE = πg.CheckLocal(πF, µdelta, "delta"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, µdelta); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label9
							}
							goto Label10
							// line 1041: if delta:
							πF.SetLineno(1041)
						Label9:
							// line 1042: dt += delta
							πF.SetLineno(1042)
							if πE = πg.CheckLocal(πF, µdt, "dt"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µdelta, "delta"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.IAdd(πF, µdt, µdelta); πE != nil {
								continue
							}
							µdt = πTemp001
							// line 1043: dtdst = dt.dst()
							πF.SetLineno(1043)
							if πE = πg.CheckLocal(πF, µdt, "dt"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µdt, ßdst, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µdtdst = πTemp003
							if πE = πg.CheckLocal(πF, µdtdst, "dtdst"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µdtdst == πTemp003).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label11
							}
							goto Label12
							// line 1044: if dtdst is None:
							πF.SetLineno(1044)
						Label11:
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("fromutc(): dt.dst gave inconsistent results; cannot convert").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 1045: raise ValueError("fromutc(): dt.dst gave inconsistent "
							πF.SetLineno(1045)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label12
						Label12:
							goto Label10
						Label10:
							if πE = πg.CheckLocal(πF, µdtdst, "dtdst"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, µdtdst); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label13
							}
							goto Label14
							// line 1047: if dtdst:
							πF.SetLineno(1047)
						Label13:
							// line 1048: return dt + dtdst
							πF.SetLineno(1048)
							if πE = πg.CheckLocal(πF, µdt, "dt"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µdtdst, "dtdst"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, µdt, µdtdst); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label15
						Label14:
							// line 1050: return dt
							πF.SetLineno(1050)
							if πE = πg.CheckLocal(πF, µdt, "dt"); πE != nil {
								continue
							}
							πR = µdt
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
					if πE = πClass.SetItem(πF, ßfromutc.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 1023: "datetime in UTC -> datetime in local time."
					πF.SetLineno(1023)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, πg.NewStr("datetime in UTC -> datetime in local time.").ToObject()); πE != nil {
						continue
					}
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßfromutc); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp007, ß__doc__, πTemp006); πE != nil {
						continue
					}
					// line 1054: def __reduce__(self):
					πF.SetLineno(1054)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("__reduce__", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µgetinitargs *πg.Object = πg.UnboundLocal; _ = µgetinitargs
						var µargs *πg.Object = πg.UnboundLocal; _ = µargs
						var µgetstate *πg.Object = πg.UnboundLocal; _ = µgetstate
						var µstate *πg.Object = πg.UnboundLocal; _ = µstate
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
							// line 1055: getinitargs = getattr(self, "__getinitargs__", None)
							πF.SetLineno(1055)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							πTemp001[1] = ß__getinitargs__.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[2] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µgetinitargs = πTemp003
							if πE = πg.CheckLocal(πF, µgetinitargs, "getinitargs"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µgetinitargs); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 1056: if getinitargs:
							πF.SetLineno(1056)
						Label1:
							// line 1057: args = getinitargs()
							πF.SetLineno(1057)
							if πE = πg.CheckLocal(πF, µgetinitargs, "getinitargs"); πE != nil {
								continue
							}
							if πTemp002, πE = µgetinitargs.Call(πF, nil, nil); πE != nil {
								continue
							}
							µargs = πTemp002
							goto Label3
						Label2:
							// line 1059: args = ()
							πF.SetLineno(1059)
							πTemp002 = πg.NewTuple0().ToObject()
							µargs = πTemp002
							goto Label3
						Label3:
							// line 1060: getstate = getattr(self, "__getstate__", None)
							πF.SetLineno(1060)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							πTemp001[1] = ß__getstate__.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[2] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µgetstate = πTemp003
							if πE = πg.CheckLocal(πF, µgetstate, "getstate"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µgetstate); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label4
							}
							goto Label5
							// line 1061: if getstate:
							πF.SetLineno(1061)
						Label4:
							// line 1062: state = getstate()
							πF.SetLineno(1062)
							if πE = πg.CheckLocal(πF, µgetstate, "getstate"); πE != nil {
								continue
							}
							if πTemp002, πE = µgetstate.Call(πF, nil, nil); πE != nil {
								continue
							}
							µstate = πTemp002
							goto Label6
						Label5:
							// line 1064: state = getattr(self, "__dict__", None) or None
							πF.SetLineno(1064)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							πTemp001[1] = ß__dict__.ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[2] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp002 = πTemp005
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label7
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πTemp003
						Label7:
							µstate = πTemp002
							goto Label6
						Label6:
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(µstate == πTemp003).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label8
							}
							goto Label9
							// line 1065: if state is None:
							πF.SetLineno(1065)
						Label8:
							// line 1066: return (self.__class__, args)
							πF.SetLineno(1066)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß__class__, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(πTemp003, µargs).ToObject()
							πR = πTemp002
							continue
							goto Label10
						Label9:
							// line 1068: return (self.__class__, args, state)
							πF.SetLineno(1068)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß__class__, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple3(πTemp003, µargs, µstate).ToObject()
							πR = πTemp002
							continue
							goto Label10
						Label10:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß__reduce__.ToObject(), πTemp006); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp029, πE = πTemp027.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp029 == nil {
				πTemp029 = πg.TypeType.ToObject()
			}
			if πTemp030, πE = πTemp029.Call(πF, []*πg.Object{πg.NewStr("tzinfo").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp027.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßtzinfo.ToObject(), πTemp030); πE != nil {
				continue
			}
			// line 1070: _tzinfo_class = tzinfo
			πF.SetLineno(1070)
			if πTemp028, πE = πg.ResolveGlobal(πF, ßtzinfo); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_tzinfo_class.ToObject(), πTemp028); πE != nil {
				continue
			}
			// line 1072: class time(object):
			πF.SetLineno(1072)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp030, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp030
			πTemp027 = πg.NewDict()
			if πTemp028, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp027.SetItem(πF, ß__module__.ToObject(), πTemp028); πE != nil {
				continue
			}
			_, πE = πg.NewCode("time", "/usr/lib/python2.7/datetime.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp027
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
				var πTemp030 *πg.Object
				_ = πTemp030
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 1073: """Time with time zone.
					πF.SetLineno(1073)
					// line 1073: """Time with time zone.
					πF.SetLineno(1073)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Time with time zone.\n\n    Constructors:\n\n    __new__()\n\n    Operators:\n\n    __repr__, __str__\n    __cmp__, __hash__\n\n    Methods:\n\n    strftime()\n    isoformat()\n    utcoffset()\n    tzname()\n    dst()\n\n    Properties (readonly):\n    hour, minute, second, microsecond, tzinfo\n    ").ToObject()); πE != nil {
						continue
					}
					// line 1095: __slots__ = '_hour', '_minute', '_second', '_microsecond', '_tzinfo', '_hashcode'
					πF.SetLineno(1095)
					πTemp001 = πg.NewTuple6(ß_hour.ToObject(), ß_minute.ToObject(), ß_second.ToObject(), ß_microsecond.ToObject(), ß_tzinfo.ToObject(), ß_hashcode.ToObject()).ToObject()
					if πE = πClass.SetItem(πF, ß__slots__.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 1097: def __new__(cls, hour=0, minute=0, second=0, microsecond=0, tzinfo=None):
					πF.SetLineno(1097)
					πTemp002 = make([]πg.Param, 6)
					πTemp002[0] = πg.Param{Name: "cls", Def: nil}
					πTemp002[1] = πg.Param{Name: "hour", Def: πg.NewInt(0).ToObject()}
					πTemp002[2] = πg.Param{Name: "minute", Def: πg.NewInt(0).ToObject()}
					πTemp002[3] = πg.Param{Name: "second", Def: πg.NewInt(0).ToObject()}
					πTemp002[4] = πg.Param{Name: "microsecond", Def: πg.NewInt(0).ToObject()}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[5] = πg.Param{Name: "tzinfo", Def: πTemp003}
					πTemp001 = πg.NewFunction(πg.NewCode("__new__", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µcls *πg.Object = πArgs[0]; _ = µcls
						var µhour *πg.Object = πArgs[1]; _ = µhour
						var µminute *πg.Object = πArgs[2]; _ = µminute
						var µsecond *πg.Object = πArgs[3]; _ = µsecond
						var µmicrosecond *πg.Object = πArgs[4]; _ = µmicrosecond
						var µtzinfo *πg.Object = πArgs[5]; _ = µtzinfo
						var µself *πg.Object = πg.UnboundLocal; _ = µself
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 1098: """Constructor.
							πF.SetLineno(1098)
							// line 1112: hour, minute, second, microsecond = _check_time_fields(
							πF.SetLineno(1112)
							πTemp001 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µhour, "hour"); πE != nil {
								continue
							}
							πTemp001[0] = µhour
							if πE = πg.CheckLocal(πF, µminute, "minute"); πE != nil {
								continue
							}
							πTemp001[1] = µminute
							if πE = πg.CheckLocal(πF, µsecond, "second"); πE != nil {
								continue
							}
							πTemp001[2] = µsecond
							if πE = πg.CheckLocal(πF, µmicrosecond, "microsecond"); πE != nil {
								continue
							}
							πTemp001[3] = µmicrosecond
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_check_time_fields); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp006}}}, πTemp003); πE != nil {
								continue
							}
							µhour = πTemp002
							µminute = πTemp004
							µsecond = πTemp005
							µmicrosecond = πTemp006
							// line 1114: _check_tzinfo_arg(tzinfo)
							πF.SetLineno(1114)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtzinfo, "tzinfo"); πE != nil {
								continue
							}
							πTemp001[0] = µtzinfo
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_check_tzinfo_arg); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1115: self = object.__new__(cls)
							πF.SetLineno(1115)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							πTemp001[0] = µcls
							if πTemp002, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ß__new__, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µself = πTemp002
							// line 1116: self._hour = hour
							πF.SetLineno(1116)
							if πE = πg.CheckLocal(πF, µhour, "hour"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µhour); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_hour, πTemp002); πE != nil {
								continue
							}
							// line 1117: self._minute = minute
							πF.SetLineno(1117)
							if πE = πg.CheckLocal(πF, µminute, "minute"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µminute); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_minute, πTemp002); πE != nil {
								continue
							}
							// line 1118: self._second = second
							πF.SetLineno(1118)
							if πE = πg.CheckLocal(πF, µsecond, "second"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µsecond); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_second, πTemp002); πE != nil {
								continue
							}
							// line 1119: self._microsecond = microsecond
							πF.SetLineno(1119)
							if πE = πg.CheckLocal(πF, µmicrosecond, "microsecond"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µmicrosecond); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_microsecond, πTemp002); πE != nil {
								continue
							}
							// line 1120: self._tzinfo = tzinfo
							πF.SetLineno(1120)
							if πE = πg.CheckLocal(πF, µtzinfo, "tzinfo"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µtzinfo); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_tzinfo, πTemp002); πE != nil {
								continue
							}
							// line 1121: self._hashcode = -1
							πF.SetLineno(1121)
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_hashcode, πTemp003); πE != nil {
								continue
							}
							// line 1122: return self
							πF.SetLineno(1122)
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
					if πE = πClass.SetItem(πF, ß__new__.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 1098: """Constructor.
					πF.SetLineno(1098)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewStr("Constructor.\n\n        Arguments:\n\n        hour, minute (required)\n        second, microsecond (default to zero)\n        tzinfo (default to None)\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ß__new__); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp004, ß__doc__, πTemp003); πE != nil {
						continue
					}
					// line 1126: def hour(self):
					πF.SetLineno(1126)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("hour", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1127: """hour (0-23)"""
							πF.SetLineno(1127)
							// line 1128: return self._hour
							πF.SetLineno(1128)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_hour, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßhour.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 1127: """hour (0-23)"""
					πF.SetLineno(1127)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πg.NewStr("hour (0-23)").ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßhour); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp005, ß__doc__, πTemp004); πE != nil {
						continue
					}
					// line 1125: @property
					πF.SetLineno(1125)
					πTemp006 = πF.MakeArgs(1)
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßhour); πE != nil {
						continue
					}
					πTemp006[0] = πTemp004
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßproperty); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πE = πClass.SetItem(πF, ßhour.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 1131: def minute(self):
					πF.SetLineno(1131)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("minute", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1132: """minute (0-59)"""
							πF.SetLineno(1132)
							// line 1133: return self._minute
							πF.SetLineno(1133)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_minute, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßminute.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 1132: """minute (0-59)"""
					πF.SetLineno(1132)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πg.NewStr("minute (0-59)").ToObject()); πE != nil {
						continue
					}
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßminute); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp007, ß__doc__, πTemp005); πE != nil {
						continue
					}
					// line 1130: @property
					πF.SetLineno(1130)
					πTemp006 = πF.MakeArgs(1)
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßminute); πE != nil {
						continue
					}
					πTemp006[0] = πTemp005
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßproperty); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp005.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πE = πClass.SetItem(πF, ßminute.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 1136: def second(self):
					πF.SetLineno(1136)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("second", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1137: """second (0-59)"""
							πF.SetLineno(1137)
							// line 1138: return self._second
							πF.SetLineno(1138)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_second, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßsecond.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 1137: """second (0-59)"""
					πF.SetLineno(1137)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp007}, πg.NewStr("second (0-59)").ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ßsecond); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp008, ß__doc__, πTemp007); πE != nil {
						continue
					}
					// line 1135: @property
					πF.SetLineno(1135)
					πTemp006 = πF.MakeArgs(1)
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßsecond); πE != nil {
						continue
					}
					πTemp006[0] = πTemp007
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßproperty); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp007.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πE = πClass.SetItem(πF, ßsecond.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 1141: def microsecond(self):
					πF.SetLineno(1141)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("microsecond", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1142: """microsecond (0-999999)"""
							πF.SetLineno(1142)
							// line 1143: return self._microsecond
							πF.SetLineno(1143)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_microsecond, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßmicrosecond.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 1142: """microsecond (0-999999)"""
					πF.SetLineno(1142)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp008}, πg.NewStr("microsecond (0-999999)").ToObject()); πE != nil {
						continue
					}
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßmicrosecond); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp009, ß__doc__, πTemp008); πE != nil {
						continue
					}
					// line 1140: @property
					πF.SetLineno(1140)
					πTemp006 = πF.MakeArgs(1)
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ßmicrosecond); πE != nil {
						continue
					}
					πTemp006[0] = πTemp008
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ßproperty); πE != nil {
						continue
					}
					if πTemp009, πE = πTemp008.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πE = πClass.SetItem(πF, ßmicrosecond.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 1146: def tzinfo(self):
					πF.SetLineno(1146)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("tzinfo", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1147: """timezone info object"""
							πF.SetLineno(1147)
							// line 1148: return self._tzinfo
							πF.SetLineno(1148)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_tzinfo, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtzinfo.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 1147: """timezone info object"""
					πF.SetLineno(1147)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp009}, πg.NewStr("timezone info object").ToObject()); πE != nil {
						continue
					}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßtzinfo); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp010, ß__doc__, πTemp009); πE != nil {
						continue
					}
					// line 1145: @property
					πF.SetLineno(1145)
					πTemp006 = πF.MakeArgs(1)
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßtzinfo); πE != nil {
						continue
					}
					πTemp006[0] = πTemp009
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßproperty); πE != nil {
						continue
					}
					if πTemp010, πE = πTemp009.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πE = πClass.SetItem(πF, ßtzinfo.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 1154: def __eq__(self, other):
					πF.SetLineno(1154)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "other", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("__eq__", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
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
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtime); πE != nil {
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
							goto Label2
							// line 1155: if isinstance(other, time):
							πF.SetLineno(1155)
						Label1:
							// line 1156: return self._cmp(other) == 0
							πF.SetLineno(1156)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_cmp, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.Eq(πF, πTemp005, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							πR = πTemp002
							continue
							goto Label3
						Label2:
							// line 1158: return False
							πF.SetLineno(1158)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__eq__.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 1160: def __ne__(self, other):
					πF.SetLineno(1160)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "other", Def: nil}
					πTemp010 = πg.NewFunction(πg.NewCode("__ne__", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
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
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtime); πE != nil {
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
							goto Label2
							// line 1161: if isinstance(other, time):
							πF.SetLineno(1161)
						Label1:
							// line 1162: return self._cmp(other) != 0
							πF.SetLineno(1162)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_cmp, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.NE(πF, πTemp005, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							πR = πTemp002
							continue
							goto Label3
						Label2:
							// line 1164: return True
							πF.SetLineno(1164)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__ne__.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 1166: def __le__(self, other):
					πF.SetLineno(1166)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "other", Def: nil}
					πTemp011 = πg.NewFunction(πg.NewCode("__le__", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
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
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtime); πE != nil {
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
							goto Label2
							// line 1167: if isinstance(other, time):
							πF.SetLineno(1167)
						Label1:
							// line 1168: return self._cmp(other) <= 0
							πF.SetLineno(1168)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_cmp, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.LE(πF, πTemp005, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							πR = πTemp002
							continue
							goto Label3
						Label2:
							// line 1170: _cmperror(self, other)
							πF.SetLineno(1170)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[1] = µother
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_cmperror); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
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
					if πE = πClass.SetItem(πF, ß__le__.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 1172: def __lt__(self, other):
					πF.SetLineno(1172)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "other", Def: nil}
					πTemp012 = πg.NewFunction(πg.NewCode("__lt__", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
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
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtime); πE != nil {
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
							goto Label2
							// line 1173: if isinstance(other, time):
							πF.SetLineno(1173)
						Label1:
							// line 1174: return self._cmp(other) < 0
							πF.SetLineno(1174)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_cmp, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.LT(πF, πTemp005, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							πR = πTemp002
							continue
							goto Label3
						Label2:
							// line 1176: _cmperror(self, other)
							πF.SetLineno(1176)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[1] = µother
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_cmperror); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
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
					if πE = πClass.SetItem(πF, ß__lt__.ToObject(), πTemp012); πE != nil {
						continue
					}
					// line 1178: def __ge__(self, other):
					πF.SetLineno(1178)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "other", Def: nil}
					πTemp013 = πg.NewFunction(πg.NewCode("__ge__", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
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
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtime); πE != nil {
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
							goto Label2
							// line 1179: if isinstance(other, time):
							πF.SetLineno(1179)
						Label1:
							// line 1180: return self._cmp(other) >= 0
							πF.SetLineno(1180)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_cmp, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.GE(πF, πTemp005, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							πR = πTemp002
							continue
							goto Label3
						Label2:
							// line 1182: _cmperror(self, other)
							πF.SetLineno(1182)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[1] = µother
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_cmperror); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
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
					if πE = πClass.SetItem(πF, ß__ge__.ToObject(), πTemp013); πE != nil {
						continue
					}
					// line 1184: def __gt__(self, other):
					πF.SetLineno(1184)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "other", Def: nil}
					πTemp014 = πg.NewFunction(πg.NewCode("__gt__", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
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
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtime); πE != nil {
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
							goto Label2
							// line 1185: if isinstance(other, time):
							πF.SetLineno(1185)
						Label1:
							// line 1186: return self._cmp(other) > 0
							πF.SetLineno(1186)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_cmp, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.GT(πF, πTemp005, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							πR = πTemp002
							continue
							goto Label3
						Label2:
							// line 1188: _cmperror(self, other)
							πF.SetLineno(1188)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[1] = µother
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_cmperror); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
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
					if πE = πClass.SetItem(πF, ß__gt__.ToObject(), πTemp014); πE != nil {
						continue
					}
					// line 1190: def _cmp(self, other):
					πF.SetLineno(1190)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "other", Def: nil}
					πTemp015 = πg.NewFunction(πg.NewCode("_cmp", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
						var µmytz *πg.Object = πg.UnboundLocal; _ = µmytz
						var µottz *πg.Object = πg.UnboundLocal; _ = µottz
						var µmyoff *πg.Object = πg.UnboundLocal; _ = µmyoff
						var µotoff *πg.Object = πg.UnboundLocal; _ = µotoff
						var µbase_compare *πg.Object = πg.UnboundLocal; _ = µbase_compare
						var µmyhhmm *πg.Object = πg.UnboundLocal; _ = µmyhhmm
						var µothhmm *πg.Object = πg.UnboundLocal; _ = µothhmm
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
						var πTemp007 *πg.Object
						_ = πTemp007
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 1191: assert isinstance(other, time)
							πF.SetLineno(1191)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtime); πE != nil {
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
							if πE = πg.Assert(πF, πTemp003, nil); πE != nil {
								continue
							}
							// line 1192: mytz = self._tzinfo
							πF.SetLineno(1192)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_tzinfo, nil); πE != nil {
								continue
							}
							µmytz = πTemp002
							// line 1193: ottz = other._tzinfo
							πF.SetLineno(1193)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µother, ß_tzinfo, nil); πE != nil {
								continue
							}
							µottz = πTemp002
							// line 1194: myoff = otoff = None
							πF.SetLineno(1194)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							µmyoff = πTemp002
							µotoff = πTemp002
							if πE = πg.CheckLocal(πF, µmytz, "mytz"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µottz, "ottz"); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(µmytz == µottz).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 1196: if mytz is ottz:
							πF.SetLineno(1196)
						Label1:
							// line 1197: base_compare = True
							πF.SetLineno(1197)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							µbase_compare = πTemp002
							goto Label3
						Label2:
							// line 1199: myoff = self._utcoffset()
							πF.SetLineno(1199)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_utcoffset, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µmyoff = πTemp003
							// line 1200: otoff = other._utcoffset()
							πF.SetLineno(1200)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µother, ß_utcoffset, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µotoff = πTemp003
							// line 1201: base_compare = myoff == otoff
							πF.SetLineno(1201)
							if πE = πg.CheckLocal(πF, µmyoff, "myoff"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µotoff, "otoff"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, µmyoff, µotoff); πE != nil {
								continue
							}
							µbase_compare = πTemp002
							goto Label3
						Label3:
							if πE = πg.CheckLocal(πF, µbase_compare, "base_compare"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µbase_compare); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label4
							}
							goto Label5
							// line 1203: if base_compare:
							πF.SetLineno(1203)
						Label4:
							// line 1204: return _cmp((self._hour, self._minute, self._second,
							πF.SetLineno(1204)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_hour, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ß_minute, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ß_second, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ß_microsecond, nil); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple4(πTemp003, πTemp005, πTemp006, πTemp007).ToObject()
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µother, ß_hour, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µother, ß_minute, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µother, ß_second, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µother, ß_microsecond, nil); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple4(πTemp003, πTemp005, πTemp006, πTemp007).ToObject()
							πTemp001[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_cmp); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πR = πTemp003
							continue
							goto Label5
						Label5:
							if πE = πg.CheckLocal(πF, µmyoff, "myoff"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(µmyoff == πTemp005).ToObject()
							πTemp002 = πTemp003
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label6
							}
							if πE = πg.CheckLocal(πF, µotoff, "otoff"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(µotoff == πTemp005).ToObject()
							πTemp002 = πTemp003
						Label6:
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label7
							}
							goto Label8
							// line 1208: if myoff is None or otoff is None:
							πF.SetLineno(1208)
						Label7:
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("can't compare offset-naive and offset-aware times").ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1209: raise TypeError("can't compare offset-naive and offset-aware times")
							πF.SetLineno(1209)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label8
						Label8:
							// line 1210: myhhmm = self._hour * 60 + self._minute - myoff
							πF.SetLineno(1210)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ß_hour, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Mul(πF, πTemp006, πg.NewInt(60).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ß_minute, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, πTemp005, πTemp006); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µmyoff, "myoff"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Sub(πF, πTemp003, µmyoff); πE != nil {
								continue
							}
							µmyhhmm = πTemp002
							// line 1211: othhmm = other._hour * 60 + other._minute - otoff
							πF.SetLineno(1211)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µother, ß_hour, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Mul(πF, πTemp006, πg.NewInt(60).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µother, ß_minute, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, πTemp005, πTemp006); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µotoff, "otoff"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Sub(πF, πTemp003, µotoff); πE != nil {
								continue
							}
							µothhmm = πTemp002
							// line 1212: return _cmp((myhhmm, self._second, self._microsecond),
							πF.SetLineno(1212)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µmyhhmm, "myhhmm"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_second, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ß_microsecond, nil); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple3(µmyhhmm, πTemp003, πTemp005).ToObject()
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µothhmm, "othhmm"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µother, ß_second, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µother, ß_microsecond, nil); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple3(µothhmm, πTemp003, πTemp005).ToObject()
							πTemp001[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_cmp); πE != nil {
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
					if πE = πClass.SetItem(πF, ß_cmp.ToObject(), πTemp015); πE != nil {
						continue
					}
					// line 1215: def __hash__(self):
					πF.SetLineno(1215)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp016 = πg.NewFunction(πg.NewCode("__hash__", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µtzoff *πg.Object = πg.UnboundLocal; _ = µtzoff
						var µh *πg.Object = πg.UnboundLocal; _ = µh
						var µm *πg.Object = πg.UnboundLocal; _ = µm
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
						var πTemp007 []*πg.Object
						_ = πTemp007
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 1216: """Hash."""
							πF.SetLineno(1216)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_hashcode, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
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
							// line 1217: if self._hashcode == -1:
							πF.SetLineno(1217)
						Label1:
							// line 1218: tzoff = self._utcoffset()
							πF.SetLineno(1218)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_utcoffset, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µtzoff = πTemp002
							if πE = πg.CheckLocal(πF, µtzoff, "tzoff"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µtzoff); πE != nil {
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
							// line 1219: if not tzoff:  # zero or None
							πF.SetLineno(1219)
						Label3:
							// line 1220: self._hashcode = hash(self._getstate()[0])
							πF.SetLineno(1220)
							πTemp005 = πF.MakeArgs(1)
							πTemp001 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_getstate, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp006, πTemp001); πE != nil {
								continue
							}
							πTemp005[0] = πTemp002
							if πTemp001, πE = πg.ResolveGlobal(πF, ßhash); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_hashcode, πTemp001); πE != nil {
								continue
							}
							goto Label5
						Label4:
							// line 1222: h, m = divmod(self.hour * 60 + self.minute - tzoff, 60)
							πF.SetLineno(1222)
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßhour, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Mul(πF, πTemp006, πg.NewInt(60).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßminute, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp003, πTemp006); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtzoff, "tzoff"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Sub(πF, πTemp002, µtzoff); πE != nil {
								continue
							}
							πTemp005[0] = πTemp001
							πTemp005[1] = πg.NewInt(60).ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßdivmod); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp003}}}, πTemp002); πE != nil {
								continue
							}
							µh = πTemp001
							µm = πTemp003
							if πE = πg.CheckLocal(πF, µh, "h"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.LE(πF, πg.NewInt(0).ToObject(), µh); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp004 {
								goto Label6
							}
							if πTemp001, πE = πg.LT(πF, µh, πg.NewInt(24).ToObject()); πE != nil {
								continue
							}
						Label6:
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label7
							}
							goto Label8
							// line 1223: if 0 <= h < 24:
							πF.SetLineno(1223)
						Label7:
							// line 1224: self._hashcode = hash(time(h, m, self.second, self.microsecond))
							πF.SetLineno(1224)
							πTemp005 = πF.MakeArgs(1)
							πTemp007 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µh, "h"); πE != nil {
								continue
							}
							πTemp007[0] = µh
							if πE = πg.CheckLocal(πF, µm, "m"); πE != nil {
								continue
							}
							πTemp007[1] = µm
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßsecond, nil); πE != nil {
								continue
							}
							πTemp007[2] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßmicrosecond, nil); πE != nil {
								continue
							}
							πTemp007[3] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßtime); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp005[0] = πTemp002
							if πTemp001, πE = πg.ResolveGlobal(πF, ßhash); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_hashcode, πTemp001); πE != nil {
								continue
							}
							goto Label9
						Label8:
							// line 1226: self._hashcode = hash((h, m, self.second, self.microsecond))
							πF.SetLineno(1226)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µh, "h"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µm, "m"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßsecond, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmicrosecond, nil); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple4(µh, µm, πTemp002, πTemp003).ToObject()
							πTemp005[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßhash); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_hashcode, πTemp001); πE != nil {
								continue
							}
							goto Label9
						Label9:
							goto Label5
						Label5:
							goto Label2
						Label2:
							// line 1227: return self._hashcode
							πF.SetLineno(1227)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_hashcode, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__hash__.ToObject(), πTemp016); πE != nil {
						continue
					}
					// line 1216: """Hash."""
					πF.SetLineno(1216)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp017}, πg.NewStr("Hash.").ToObject()); πE != nil {
						continue
					}
					if πTemp018, πE = πg.ResolveClass(πF, πClass, nil, ß__hash__); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp018, ß__doc__, πTemp017); πE != nil {
						continue
					}
					// line 1231: def _tzstr(self, sep=":"):
					πF.SetLineno(1231)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "sep", Def: πg.NewStr(":").ToObject()}
					πTemp017 = πg.NewFunction(πg.NewCode("_tzstr", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µsep *πg.Object = πArgs[1]; _ = µsep
						var µoff *πg.Object = πg.UnboundLocal; _ = µoff
						var µsign *πg.Object = πg.UnboundLocal; _ = µsign
						var µhh *πg.Object = πg.UnboundLocal; _ = µhh
						var µmm *πg.Object = πg.UnboundLocal; _ = µmm
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
							// line 1232: """Return formatted timezone offset (+xx:xx) or None."""
							πF.SetLineno(1232)
							// line 1233: off = self._utcoffset()
							πF.SetLineno(1233)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_utcoffset, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µoff = πTemp002
							if πE = πg.CheckLocal(πF, µoff, "off"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µoff != πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 1234: if off is not None:
							πF.SetLineno(1234)
						Label1:
							if πE = πg.CheckLocal(πF, µoff, "off"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.LT(πF, µoff, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label3
							}
							goto Label4
							// line 1235: if off < 0:
							πF.SetLineno(1235)
						Label3:
							// line 1236: sign = "-"
							πF.SetLineno(1236)
							µsign = πg.NewStr("-").ToObject()
							// line 1237: off = -off
							πF.SetLineno(1237)
							if πE = πg.CheckLocal(πF, µoff, "off"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Neg(πF, µoff); πE != nil {
								continue
							}
							µoff = πTemp001
							goto Label5
						Label4:
							// line 1239: sign = "+"
							πF.SetLineno(1239)
							µsign = πg.NewStr("+").ToObject()
							goto Label5
						Label5:
							// line 1240: hh, mm = divmod(off, 60)
							πF.SetLineno(1240)
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µoff, "off"); πE != nil {
								continue
							}
							πTemp004[0] = µoff
							πTemp004[1] = πg.NewInt(60).ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßdivmod); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp005}}}, πTemp002); πE != nil {
								continue
							}
							µhh = πTemp001
							µmm = πTemp005
							// line 1241: assert 0 <= hh < 24
							πF.SetLineno(1241)
							if πE = πg.CheckLocal(πF, µhh, "hh"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.LE(πF, πg.NewInt(0).ToObject(), µhh); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp003 {
								goto Label6
							}
							if πTemp001, πE = πg.LT(πF, µhh, πg.NewInt(24).ToObject()); πE != nil {
								continue
							}
						Label6:
							if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
								continue
							}
							// line 1242: off = "%s%02d%s%02d" % (sign, hh, sep, mm)
							πF.SetLineno(1242)
							if πE = πg.CheckLocal(πF, µsign, "sign"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µhh, "hh"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µsep, "sep"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µmm, "mm"); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple4(µsign, µhh, µsep, µmm).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("%s%02d%s%02d").ToObject(), πTemp002); πE != nil {
								continue
							}
							µoff = πTemp001
							goto Label2
						Label2:
							// line 1243: return off
							πF.SetLineno(1243)
							if πE = πg.CheckLocal(πF, µoff, "off"); πE != nil {
								continue
							}
							πR = µoff
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß_tzstr.ToObject(), πTemp017); πE != nil {
						continue
					}
					// line 1232: """Return formatted timezone offset (+xx:xx) or None."""
					πF.SetLineno(1232)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp018}, πg.NewStr("Return formatted timezone offset (+xx:xx) or None.").ToObject()); πE != nil {
						continue
					}
					if πTemp019, πE = πg.ResolveClass(πF, πClass, nil, ß_tzstr); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp019, ß__doc__, πTemp018); πE != nil {
						continue
					}
					// line 1245: def __repr__(self):
					πF.SetLineno(1245)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp018 = πg.NewFunction(πg.NewCode("__repr__", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µs *πg.Object = πg.UnboundLocal; _ = µs
						var µmodule *πg.Object = πg.UnboundLocal; _ = µmodule
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 1246: """Convert to formal string, for repr()."""
							πF.SetLineno(1246)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_microsecond, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.NE(πF, πTemp002, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_second, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.NE(πF, πTemp002, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label2
							}
							goto Label3
							// line 1247: if self._microsecond != 0:
							πF.SetLineno(1247)
						Label1:
							// line 1248: s = ", %d, %d" % (self._second, self._microsecond)
							πF.SetLineno(1248)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß_second, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ß_microsecond, nil); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(πTemp004, πTemp005).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr(", %d, %d").ToObject(), πTemp002); πE != nil {
								continue
							}
							µs = πTemp001
							goto Label4
							// line 1249: elif self._second != 0:
							πF.SetLineno(1249)
						Label2:
							// line 1250: s = ", %d" % self._second
							πF.SetLineno(1250)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_second, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr(", %d").ToObject(), πTemp002); πE != nil {
								continue
							}
							µs = πTemp001
							goto Label4
						Label3:
							// line 1252: s = ""
							πF.SetLineno(1252)
							µs = ß.ToObject()
							goto Label4
						Label4:
							// line 1253: module = "datetime." if self.__class__ is time else ""
							πF.SetLineno(1253)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß__class__, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßtime); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp004 == πTemp005).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp003 {
								goto Label5
							}
							πTemp001 = πg.NewStr("datetime.").ToObject()
							goto Label6
						Label5:
							πTemp001 = ß.ToObject()
						Label6:
							µmodule = πTemp001
							// line 1254: s= "%s(%d, %d%s)" % (module + self.__class__.__name__,
							πF.SetLineno(1254)
							if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ß__class__, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp005, ß__name__, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Add(πF, µmodule, πTemp006); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ß_hour, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ß_minute, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple4(πTemp004, πTemp005, πTemp006, µs).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("%s(%d, %d%s)").ToObject(), πTemp002); πE != nil {
								continue
							}
							µs = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_tzinfo, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp002 != πTemp004).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label7
							}
							goto Label8
							// line 1256: if self._tzinfo is not None:
							πF.SetLineno(1256)
						Label7:
							// line 1257: assert s[-1:] == ")"
							πF.SetLineno(1257)
							if πTemp004, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πTemp004, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µs, πTemp002); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, πTemp004, πg.NewStr(")").ToObject()); πE != nil {
								continue
							}
							if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
								continue
							}
							// line 1258: s = s[:-1] + ", tzinfo=%r" % self._tzinfo + ")"
							πF.SetLineno(1258)
							if πTemp005, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πTemp005, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µs, πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ß_tzinfo, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Mod(πF, πg.NewStr(", tzinfo=%r").ToObject(), πTemp006); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp005, πTemp004); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp002, πg.NewStr(")").ToObject()); πE != nil {
								continue
							}
							µs = πTemp001
							goto Label8
						Label8:
							// line 1259: return s
							πF.SetLineno(1259)
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
					if πE = πClass.SetItem(πF, ß__repr__.ToObject(), πTemp018); πE != nil {
						continue
					}
					// line 1246: """Convert to formal string, for repr()."""
					πF.SetLineno(1246)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp019}, πg.NewStr("Convert to formal string, for repr().").ToObject()); πE != nil {
						continue
					}
					if πTemp020, πE = πg.ResolveClass(πF, πClass, nil, ß__repr__); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp020, ß__doc__, πTemp019); πE != nil {
						continue
					}
					// line 1261: def isoformat(self):
					πF.SetLineno(1261)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp019 = πg.NewFunction(πg.NewCode("isoformat", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µs *πg.Object = πg.UnboundLocal; _ = µs
						var µtz *πg.Object = πg.UnboundLocal; _ = µtz
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
							// line 1262: """Return the time formatted according to ISO.
							πF.SetLineno(1262)
							// line 1267: s = _format_time(self._hour, self._minute, self._second,
							πF.SetLineno(1267)
							πTemp001 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_hour, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_minute, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_second, nil); πE != nil {
								continue
							}
							πTemp001[2] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_microsecond, nil); πE != nil {
								continue
							}
							πTemp001[3] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_format_time); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µs = πTemp003
							// line 1269: tz = self._tzstr()
							πF.SetLineno(1269)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_tzstr, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µtz = πTemp003
							if πE = πg.CheckLocal(πF, µtz, "tz"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µtz); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 1270: if tz:
							πF.SetLineno(1270)
						Label1:
							// line 1271: s += tz
							πF.SetLineno(1271)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtz, "tz"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IAdd(πF, µs, µtz); πE != nil {
								continue
							}
							µs = πTemp002
							goto Label2
						Label2:
							// line 1272: return s
							πF.SetLineno(1272)
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
					if πE = πClass.SetItem(πF, ßisoformat.ToObject(), πTemp019); πE != nil {
						continue
					}
					// line 1262: """Return the time formatted according to ISO.
					πF.SetLineno(1262)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp020}, πg.NewStr("Return the time formatted according to ISO.\n\n        This is 'HH:MM:SS.mmmmmm+zz:zz', or 'HH:MM:SS+zz:zz' if\n        self.microsecond == 0.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp021, πE = πg.ResolveClass(πF, πClass, nil, ßisoformat); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp021, ß__doc__, πTemp020); πE != nil {
						continue
					}
					// line 1274: __str__ = isoformat
					πF.SetLineno(1274)
					if πTemp020, πE = πg.ResolveClass(πF, πClass, nil, ßisoformat); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ß__str__.ToObject(), πTemp020); πE != nil {
						continue
					}
					// line 1287: def __format__(self, fmt):
					πF.SetLineno(1287)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "fmt", Def: nil}
					πTemp020 = πg.NewFunction(πg.NewCode("__format__", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µfmt *πg.Object = πArgs[1]; _ = µfmt
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µfmt, "fmt"); πE != nil {
								continue
							}
							πTemp002[0] = µfmt
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
							// line 1288: if not isinstance(fmt, (str, unicode)):
							πF.SetLineno(1288)
						Label1:
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfmt, "fmt"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µfmt, ß__class__, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ß__name__, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("__format__ expects str or unicode, not %s").ToObject(), πTemp004); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 1289: raise ValueError("__format__ expects str or unicode, not %s" %
							πF.SetLineno(1289)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label2
						Label2:
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfmt, "fmt"); πE != nil {
								continue
							}
							πTemp002[0] = µfmt
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp001, πE = πg.NE(πF, πTemp004, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label3
							}
							goto Label4
							// line 1291: if len(fmt) != 0:
							πF.SetLineno(1291)
						Label3:
							// line 1292: return self.strftime(fmt)
							πF.SetLineno(1292)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfmt, "fmt"); πE != nil {
								continue
							}
							πTemp002[0] = µfmt
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstrftime, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πR = πTemp003
							continue
							goto Label4
						Label4:
							// line 1293: return str(self)
							πF.SetLineno(1293)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp002[0] = µself
							if πTemp001, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
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
					if πE = πClass.SetItem(πF, ß__format__.ToObject(), πTemp020); πE != nil {
						continue
					}
					// line 1297: def utcoffset(self):
					πF.SetLineno(1297)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp021 = πg.NewFunction(πg.NewCode("utcoffset", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µoffset *πg.Object = πg.UnboundLocal; _ = µoffset
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 1298: """Return the timezone offset in minutes east of UTC (negative west of
							πF.SetLineno(1298)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_tzinfo, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp002 == πTemp003).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 1300: if self._tzinfo is None:
							πF.SetLineno(1300)
						Label1:
							// line 1301: return None
							πF.SetLineno(1301)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label2
						Label2:
							// line 1302: offset = self._tzinfo.utcoffset(None)
							πF.SetLineno(1302)
							πTemp005 = πF.MakeArgs(1)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp005[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_tzinfo, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßutcoffset, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µoffset = πTemp001
							// line 1303: offset = _check_utc_offset("utcoffset", offset)
							πF.SetLineno(1303)
							πTemp005 = πF.MakeArgs(2)
							πTemp005[0] = ßutcoffset.ToObject()
							if πE = πg.CheckLocal(πF, µoffset, "offset"); πE != nil {
								continue
							}
							πTemp005[1] = µoffset
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_check_utc_offset); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µoffset = πTemp002
							if πE = πg.CheckLocal(πF, µoffset, "offset"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µoffset != πTemp002).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label3
							}
							goto Label4
							// line 1304: if offset is not None:
							πF.SetLineno(1304)
						Label3:
							// line 1305: offset = timedelta._create(0, offset * 60, 0, True)
							πF.SetLineno(1305)
							πTemp005 = πF.MakeArgs(4)
							πTemp005[0] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µoffset, "offset"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mul(πF, µoffset, πg.NewInt(60).ToObject()); πE != nil {
								continue
							}
							πTemp005[1] = πTemp001
							πTemp005[2] = πg.NewInt(0).ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp005[3] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßtimedelta); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ß_create, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µoffset = πTemp001
							goto Label4
						Label4:
							// line 1306: return offset
							πF.SetLineno(1306)
							if πE = πg.CheckLocal(πF, µoffset, "offset"); πE != nil {
								continue
							}
							πR = µoffset
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßutcoffset.ToObject(), πTemp021); πE != nil {
						continue
					}
					// line 1298: """Return the timezone offset in minutes east of UTC (negative west of
					πF.SetLineno(1298)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp022}, πg.NewStr("Return the timezone offset in minutes east of UTC (negative west of\n        UTC).").ToObject()); πE != nil {
						continue
					}
					if πTemp023, πE = πg.ResolveClass(πF, πClass, nil, ßutcoffset); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp023, ß__doc__, πTemp022); πE != nil {
						continue
					}
					// line 1309: def _utcoffset(self):
					πF.SetLineno(1309)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp022 = πg.NewFunction(πg.NewCode("_utcoffset", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µoffset *πg.Object = πg.UnboundLocal; _ = µoffset
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_tzinfo, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp002 == πTemp003).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 1310: if self._tzinfo is None:
							πF.SetLineno(1310)
						Label1:
							// line 1311: return None
							πF.SetLineno(1311)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label2
						Label2:
							// line 1312: offset = self._tzinfo.utcoffset(None)
							πF.SetLineno(1312)
							πTemp005 = πF.MakeArgs(1)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp005[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_tzinfo, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßutcoffset, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µoffset = πTemp001
							// line 1313: offset = _check_utc_offset("utcoffset", offset)
							πF.SetLineno(1313)
							πTemp005 = πF.MakeArgs(2)
							πTemp005[0] = ßutcoffset.ToObject()
							if πE = πg.CheckLocal(πF, µoffset, "offset"); πE != nil {
								continue
							}
							πTemp005[1] = µoffset
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_check_utc_offset); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µoffset = πTemp002
							// line 1314: return offset
							πF.SetLineno(1314)
							if πE = πg.CheckLocal(πF, µoffset, "offset"); πE != nil {
								continue
							}
							πR = µoffset
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß_utcoffset.ToObject(), πTemp022); πE != nil {
						continue
					}
					// line 1316: def tzname(self):
					πF.SetLineno(1316)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp023 = πg.NewFunction(πg.NewCode("tzname", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µname *πg.Object = πg.UnboundLocal; _ = µname
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 1317: """Return the timezone name.
							πF.SetLineno(1317)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_tzinfo, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp002 == πTemp003).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 1323: if self._tzinfo is None:
							πF.SetLineno(1323)
						Label1:
							// line 1324: return None
							πF.SetLineno(1324)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label2
						Label2:
							// line 1325: name = self._tzinfo.tzname(None)
							πF.SetLineno(1325)
							πTemp005 = πF.MakeArgs(1)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp005[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_tzinfo, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßtzname, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µname = πTemp001
							// line 1326: _check_tzname(name)
							πF.SetLineno(1326)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp005[0] = µname
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_check_tzname); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 1327: return name
							πF.SetLineno(1327)
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
					if πE = πClass.SetItem(πF, ßtzname.ToObject(), πTemp023); πE != nil {
						continue
					}
					// line 1317: """Return the timezone name.
					πF.SetLineno(1317)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp024}, πg.NewStr("Return the timezone name.\n\n        Note that the name is 100% informational -- there's no requirement that\n        it mean anything in particular. For example, \"GMT\", \"UTC\", \"-500\",\n        \"-5:00\", \"EDT\", \"US/Eastern\", \"America/New York\" are all valid replies.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp025, πE = πg.ResolveClass(πF, πClass, nil, ßtzname); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp025, ß__doc__, πTemp024); πE != nil {
						continue
					}
					// line 1329: def dst(self):
					πF.SetLineno(1329)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp024 = πg.NewFunction(πg.NewCode("dst", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µoffset *πg.Object = πg.UnboundLocal; _ = µoffset
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 1330: """Return 0 if DST is not in effect, or the DST offset (in minutes
							πF.SetLineno(1330)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_tzinfo, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp002 == πTemp003).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 1338: if self._tzinfo is None:
							πF.SetLineno(1338)
						Label1:
							// line 1339: return None
							πF.SetLineno(1339)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label2
						Label2:
							// line 1340: offset = self._tzinfo.dst(None)
							πF.SetLineno(1340)
							πTemp005 = πF.MakeArgs(1)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp005[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_tzinfo, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßdst, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µoffset = πTemp001
							// line 1341: offset = _check_utc_offset("dst", offset)
							πF.SetLineno(1341)
							πTemp005 = πF.MakeArgs(2)
							πTemp005[0] = ßdst.ToObject()
							if πE = πg.CheckLocal(πF, µoffset, "offset"); πE != nil {
								continue
							}
							πTemp005[1] = µoffset
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_check_utc_offset); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µoffset = πTemp002
							if πE = πg.CheckLocal(πF, µoffset, "offset"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µoffset != πTemp002).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label3
							}
							goto Label4
							// line 1342: if offset is not None:
							πF.SetLineno(1342)
						Label3:
							// line 1343: offset = timedelta._create(0, offset * 60, 0, True)
							πF.SetLineno(1343)
							πTemp005 = πF.MakeArgs(4)
							πTemp005[0] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µoffset, "offset"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mul(πF, µoffset, πg.NewInt(60).ToObject()); πE != nil {
								continue
							}
							πTemp005[1] = πTemp001
							πTemp005[2] = πg.NewInt(0).ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp005[3] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßtimedelta); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ß_create, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µoffset = πTemp001
							goto Label4
						Label4:
							// line 1344: return offset
							πF.SetLineno(1344)
							if πE = πg.CheckLocal(πF, µoffset, "offset"); πE != nil {
								continue
							}
							πR = µoffset
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdst.ToObject(), πTemp024); πE != nil {
						continue
					}
					// line 1330: """Return 0 if DST is not in effect, or the DST offset (in minutes
					πF.SetLineno(1330)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp025}, πg.NewStr("Return 0 if DST is not in effect, or the DST offset (in minutes\n        eastward) if DST is in effect.\n\n        This is purely informational; the DST offset has already been added to\n        the UTC offset returned by utcoffset() if applicable, so there's no\n        need to consult dst() unless you're interested in displaying the DST\n        info.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp026, πE = πg.ResolveClass(πF, πClass, nil, ßdst); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp026, ß__doc__, πTemp025); πE != nil {
						continue
					}
					// line 1347: def _dst(self):
					πF.SetLineno(1347)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp025 = πg.NewFunction(πg.NewCode("_dst", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µoffset *πg.Object = πg.UnboundLocal; _ = µoffset
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_tzinfo, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp002 == πTemp003).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 1348: if self._tzinfo is None:
							πF.SetLineno(1348)
						Label1:
							// line 1349: return None
							πF.SetLineno(1349)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label2
						Label2:
							// line 1350: offset = self._tzinfo.dst(None)
							πF.SetLineno(1350)
							πTemp005 = πF.MakeArgs(1)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp005[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_tzinfo, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßdst, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µoffset = πTemp001
							// line 1351: offset = _check_utc_offset("dst", offset)
							πF.SetLineno(1351)
							πTemp005 = πF.MakeArgs(2)
							πTemp005[0] = ßdst.ToObject()
							if πE = πg.CheckLocal(πF, µoffset, "offset"); πE != nil {
								continue
							}
							πTemp005[1] = µoffset
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_check_utc_offset); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µoffset = πTemp002
							// line 1352: return offset
							πF.SetLineno(1352)
							if πE = πg.CheckLocal(πF, µoffset, "offset"); πE != nil {
								continue
							}
							πR = µoffset
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß_dst.ToObject(), πTemp025); πE != nil {
						continue
					}
					// line 1354: def replace(self, hour=None, minute=None, second=None, microsecond=None,
					πF.SetLineno(1354)
					πTemp002 = make([]πg.Param, 6)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					if πTemp027, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[1] = πg.Param{Name: "hour", Def: πTemp027}
					if πTemp027, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "minute", Def: πTemp027}
					if πTemp027, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[3] = πg.Param{Name: "second", Def: πTemp027}
					if πTemp027, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[4] = πg.Param{Name: "microsecond", Def: πTemp027}
					if πTemp027, πE = πg.ResolveClass(πF, πClass, nil, ßTrue); πE != nil {
						continue
					}
					πTemp002[5] = πg.Param{Name: "tzinfo", Def: πTemp027}
					πTemp026 = πg.NewFunction(πg.NewCode("replace", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µhour *πg.Object = πArgs[1]; _ = µhour
						var µminute *πg.Object = πArgs[2]; _ = µminute
						var µsecond *πg.Object = πArgs[3]; _ = µsecond
						var µmicrosecond *πg.Object = πArgs[4]; _ = µmicrosecond
						var µtzinfo *πg.Object = πArgs[5]; _ = µtzinfo
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
							// line 1356: """Return a new time with new values for the specified fields."""
							πF.SetLineno(1356)
							if πE = πg.CheckLocal(πF, µhour, "hour"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µhour == πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 1357: if hour is None:
							πF.SetLineno(1357)
						Label1:
							// line 1358: hour = self.hour
							πF.SetLineno(1358)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßhour, nil); πE != nil {
								continue
							}
							µhour = πTemp001
							goto Label2
						Label2:
							if πE = πg.CheckLocal(πF, µminute, "minute"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µminute == πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label3
							}
							goto Label4
							// line 1359: if minute is None:
							πF.SetLineno(1359)
						Label3:
							// line 1360: minute = self.minute
							πF.SetLineno(1360)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßminute, nil); πE != nil {
								continue
							}
							µminute = πTemp001
							goto Label4
						Label4:
							if πE = πg.CheckLocal(πF, µsecond, "second"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µsecond == πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label5
							}
							goto Label6
							// line 1361: if second is None:
							πF.SetLineno(1361)
						Label5:
							// line 1362: second = self.second
							πF.SetLineno(1362)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßsecond, nil); πE != nil {
								continue
							}
							µsecond = πTemp001
							goto Label6
						Label6:
							if πE = πg.CheckLocal(πF, µmicrosecond, "microsecond"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µmicrosecond == πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label7
							}
							goto Label8
							// line 1363: if microsecond is None:
							πF.SetLineno(1363)
						Label7:
							// line 1364: microsecond = self.microsecond
							πF.SetLineno(1364)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßmicrosecond, nil); πE != nil {
								continue
							}
							µmicrosecond = πTemp001
							goto Label8
						Label8:
							if πE = πg.CheckLocal(πF, µtzinfo, "tzinfo"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µtzinfo == πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label9
							}
							goto Label10
							// line 1365: if tzinfo is True:
							πF.SetLineno(1365)
						Label9:
							// line 1366: tzinfo = self.tzinfo
							πF.SetLineno(1366)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßtzinfo, nil); πE != nil {
								continue
							}
							µtzinfo = πTemp001
							goto Label10
						Label10:
							// line 1367: return time.__new__(type(self),
							πF.SetLineno(1367)
							πTemp004 = πF.MakeArgs(6)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp005[0] = µself
							if πTemp001, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp004[0] = πTemp002
							if πE = πg.CheckLocal(πF, µhour, "hour"); πE != nil {
								continue
							}
							πTemp004[1] = µhour
							if πE = πg.CheckLocal(πF, µminute, "minute"); πE != nil {
								continue
							}
							πTemp004[2] = µminute
							if πE = πg.CheckLocal(πF, µsecond, "second"); πE != nil {
								continue
							}
							πTemp004[3] = µsecond
							if πE = πg.CheckLocal(πF, µmicrosecond, "microsecond"); πE != nil {
								continue
							}
							πTemp004[4] = µmicrosecond
							if πE = πg.CheckLocal(πF, µtzinfo, "tzinfo"); πE != nil {
								continue
							}
							πTemp004[5] = µtzinfo
							if πTemp001, πE = πg.ResolveGlobal(πF, ßtime); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ß__new__, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
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
					if πE = πClass.SetItem(πF, ßreplace.ToObject(), πTemp026); πE != nil {
						continue
					}
					// line 1356: """Return a new time with new values for the specified fields."""
					πF.SetLineno(1356)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp027}, πg.NewStr("Return a new time with new values for the specified fields.").ToObject()); πE != nil {
						continue
					}
					if πTemp028, πE = πg.ResolveClass(πF, πClass, nil, ßreplace); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp028, ß__doc__, πTemp027); πE != nil {
						continue
					}
					// line 1370: def __nonzero__(self):
					πF.SetLineno(1370)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp027 = πg.NewFunction(πg.NewCode("__nonzero__", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µoffset *πg.Object = πg.UnboundLocal; _ = µoffset
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßsecond, nil); πE != nil {
								continue
							}
							πTemp001 = πTemp003
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmicrosecond, nil); πE != nil {
								continue
							}
							πTemp001 = πTemp003
						Label1:
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label2
							}
							goto Label3
							// line 1371: if self.second or self.microsecond:
							πF.SetLineno(1371)
						Label2:
							// line 1372: return True
							πF.SetLineno(1372)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label3
						Label3:
							// line 1373: offset = self._utcoffset() or 0
							πF.SetLineno(1373)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_utcoffset, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001 = πTemp004
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label4
							}
							πTemp001 = πg.NewInt(0).ToObject()
						Label4:
							µoffset = πTemp001
							// line 1374: return self.hour * 60 + self.minute != offset
							πF.SetLineno(1374)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßhour, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Mul(πF, πTemp005, πg.NewInt(60).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßminute, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, πTemp004, πTemp005); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µoffset, "offset"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.NE(πF, πTemp003, µoffset); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__nonzero__.ToObject(), πTemp027); πE != nil {
						continue
					}
					// line 1378: def _getstate(self):
					πF.SetLineno(1378)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp028 = πg.NewFunction(πg.NewCode("_getstate", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µus2 *πg.Object = πg.UnboundLocal; _ = µus2
						var µus3 *πg.Object = πg.UnboundLocal; _ = µus3
						var µus1 *πg.Object = πg.UnboundLocal; _ = µus1
						var µbasestate *πg.Object = πg.UnboundLocal; _ = µbasestate
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
							// line 1379: us2, us3 = divmod(self._microsecond, 256)
							πF.SetLineno(1379)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_microsecond, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = πg.NewInt(256).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßdivmod); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp004}}}, πTemp003); πE != nil {
								continue
							}
							µus2 = πTemp002
							µus3 = πTemp004
							// line 1380: us1, us2 = divmod(us2, 256)
							πF.SetLineno(1380)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µus2, "us2"); πE != nil {
								continue
							}
							πTemp001[0] = µus2
							πTemp001[1] = πg.NewInt(256).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßdivmod); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp004}}}, πTemp003); πE != nil {
								continue
							}
							µus1 = πTemp002
							µus2 = πTemp004
							// line 1381: basestate = _struct.pack('6B', self._hour, self._minute, self._second,
							πF.SetLineno(1381)
							πTemp001 = πF.MakeArgs(7)
							πTemp001[0] = ß6B.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_hour, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_minute, nil); πE != nil {
								continue
							}
							πTemp001[2] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_second, nil); πE != nil {
								continue
							}
							πTemp001[3] = πTemp002
							if πE = πg.CheckLocal(πF, µus1, "us1"); πE != nil {
								continue
							}
							πTemp001[4] = µus1
							if πE = πg.CheckLocal(πF, µus2, "us2"); πE != nil {
								continue
							}
							πTemp001[5] = µus2
							if πE = πg.CheckLocal(πF, µus3, "us3"); πE != nil {
								continue
							}
							πTemp001[6] = µus3
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_struct); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßpack, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µbasestate = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_tzinfo, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp003 == πTemp004).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							goto Label2
							// line 1383: if self._tzinfo is None:
							πF.SetLineno(1383)
						Label1:
							// line 1384: return (basestate,)
							πF.SetLineno(1384)
							if πE = πg.CheckLocal(πF, µbasestate, "basestate"); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple1(µbasestate).ToObject()
							πR = πTemp002
							continue
							goto Label3
						Label2:
							// line 1386: return (basestate, self._tzinfo)
							πF.SetLineno(1386)
							if πE = πg.CheckLocal(πF, µbasestate, "basestate"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_tzinfo, nil); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(µbasestate, πTemp003).ToObject()
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
					if πE = πClass.SetItem(πF, ß_getstate.ToObject(), πTemp028); πE != nil {
						continue
					}
					// line 1388: def __setstate(self, string, tzinfo):
					πF.SetLineno(1388)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "string", Def: nil}
					πTemp002[2] = πg.Param{Name: "tzinfo", Def: nil}
					πTemp029 = πg.NewFunction(πg.NewCode("__setstate", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µstring *πg.Object = πArgs[1]; _ = µstring
						var µtzinfo *πg.Object = πArgs[2]; _ = µtzinfo
						var µus1 *πg.Object = πg.UnboundLocal; _ = µus1
						var µus2 *πg.Object = πg.UnboundLocal; _ = µus2
						var µus3 *πg.Object = πg.UnboundLocal; _ = µus3
						var πTemp001 *πg.Object
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							if πE = πg.CheckLocal(πF, µtzinfo, "tzinfo"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(µtzinfo != πTemp004).ToObject()
							πTemp001 = πTemp003
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp002 {
								goto Label1
							}
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µtzinfo, "tzinfo"); πE != nil {
								continue
							}
							πTemp005[0] = µtzinfo
							if πTemp004, πE = πg.ResolveGlobal(πF, ß_tzinfo_class); πE != nil {
								continue
							}
							πTemp005[1] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp007, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp007).ToObject()
							πTemp001 = πTemp003
						Label1:
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label2
							}
							goto Label3
							// line 1389: if tzinfo is not None and not isinstance(tzinfo, _tzinfo_class):
							πF.SetLineno(1389)
						Label2:
							πTemp005 = πF.MakeArgs(1)
							πTemp005[0] = πg.NewStr("bad tzinfo state arg").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 1390: raise TypeError("bad tzinfo state arg")
							πF.SetLineno(1390)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label3
						Label3:
							// line 1391: self._hour, self._minute, self._second, us1, us2, us3 = (
							πF.SetLineno(1391)
							πTemp005 = πF.MakeArgs(1)
							πTemp003 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µstring, πTemp003); πE != nil {
								continue
							}
							πTemp005[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp005 = πF.MakeArgs(1)
							πTemp003 = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µstring, πTemp003); πE != nil {
								continue
							}
							πTemp005[0] = πTemp006
							if πTemp003, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp005 = πF.MakeArgs(1)
							πTemp003 = πg.NewInt(2).ToObject()
							if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, µstring, πTemp003); πE != nil {
								continue
							}
							πTemp005[0] = πTemp008
							if πTemp003, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp005 = πF.MakeArgs(1)
							πTemp003 = πg.NewInt(3).ToObject()
							if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetItem(πF, µstring, πTemp003); πE != nil {
								continue
							}
							πTemp005[0] = πTemp009
							if πTemp003, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp005 = πF.MakeArgs(1)
							πTemp003 = πg.NewInt(4).ToObject()
							if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetItem(πF, µstring, πTemp003); πE != nil {
								continue
							}
							πTemp005[0] = πTemp010
							if πTemp003, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
								continue
							}
							if πTemp010, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp005 = πF.MakeArgs(1)
							πTemp003 = πg.NewInt(5).ToObject()
							if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetItem(πF, µstring, πTemp003); πE != nil {
								continue
							}
							πTemp005[0] = πTemp011
							if πTemp003, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
								continue
							}
							if πTemp011, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp001 = πg.NewTuple6(πTemp004, πTemp006, πTemp008, πTemp009, πTemp010, πTemp011).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp006}, πg.TieTarget{Target: &πTemp008}, πg.TieTarget{Target: &πTemp009}, πg.TieTarget{Target: &πTemp010}}}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_hour, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_minute, πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_second, πTemp006); πE != nil {
								continue
							}
							µus1 = πTemp008
							µus2 = πTemp009
							µus3 = πTemp010
							// line 1394: self._microsecond = (((us1 << 8) | us2) << 8) | us3
							πF.SetLineno(1394)
							if πE = πg.CheckLocal(πF, µus1, "us1"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.LShift(πF, µus1, πg.NewInt(8).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µus2, "us2"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Or(πF, πTemp006, µus2); πE != nil {
								continue
							}
							if πTemp003, πE = πg.LShift(πF, πTemp004, πg.NewInt(8).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µus3, "us3"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Or(πF, πTemp003, µus3); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_microsecond, πTemp003); πE != nil {
								continue
							}
							// line 1395: self._tzinfo = tzinfo
							πF.SetLineno(1395)
							if πE = πg.CheckLocal(πF, µtzinfo, "tzinfo"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µtzinfo); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_tzinfo, πTemp001); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__setstate.ToObject(), πTemp029); πE != nil {
						continue
					}
					// line 1397: def __reduce__(self):
					πF.SetLineno(1397)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp030 = πg.NewFunction(πg.NewCode("__reduce__", "/usr/lib/python2.7/datetime.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πTemp001 *πg.Object
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
							// line 1398: return (time, self._getstate())
							πF.SetLineno(1398)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtime); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_getstate, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple2(πTemp002, πTemp004).ToObject()
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
					if πE = πClass.SetItem(πF, ß__reduce__.ToObject(), πTemp030); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp029, πE = πTemp027.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp029 == nil {
				πTemp029 = πg.TypeType.ToObject()
			}
			if πTemp030, πE = πTemp029.Call(πF, []*πg.Object{πg.NewStr("time").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp027.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßtime.ToObject(), πTemp030); πE != nil {
				continue
			}
			// line 1400: _time_class = time  # so functions w/ args named "time" can get at the class
			πF.SetLineno(1400)
			if πTemp028, πE = πg.ResolveGlobal(πF, ßtime); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_time_class.ToObject(), πTemp028); πE != nil {
				continue
			}
			// line 1402: time.min = time(0, 0, 0)
			πF.SetLineno(1402)
			πTemp002 = πF.MakeArgs(3)
			πTemp002[0] = πg.NewInt(0).ToObject()
			πTemp002[1] = πg.NewInt(0).ToObject()
			πTemp002[2] = πg.NewInt(0).ToObject()
			if πTemp028, πE = πg.ResolveGlobal(πF, ßtime); πE != nil {
				continue
			}
			if πTemp029, πE = πTemp028.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp028}, πTemp029); πE != nil {
				continue
			}
			if πTemp030, πE = πg.ResolveGlobal(πF, ßtime); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp030, ßmin, πTemp028); πE != nil {
				continue
			}
			// line 1403: time.max = time(23, 59, 59, 999999)
			πF.SetLineno(1403)
			πTemp002 = πF.MakeArgs(4)
			πTemp002[0] = πg.NewInt(23).ToObject()
			πTemp002[1] = πg.NewInt(59).ToObject()
			πTemp002[2] = πg.NewInt(59).ToObject()
			πTemp002[3] = πg.NewInt(999999).ToObject()
			if πTemp028, πE = πg.ResolveGlobal(πF, ßtime); πE != nil {
				continue
			}
			if πTemp029, πE = πTemp028.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp028}, πTemp029); πE != nil {
				continue
			}
			if πTemp030, πE = πg.ResolveGlobal(πF, ßtime); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp030, ßmax, πTemp028); πE != nil {
				continue
			}
			// line 1404: time.resolution = timedelta(microseconds=1)
			πF.SetLineno(1404)
			πTemp031 = πg.KWArgs{
				{"microseconds", πg.NewInt(1).ToObject()},
			}
			if πTemp028, πE = πg.ResolveGlobal(πF, ßtimedelta); πE != nil {
				continue
			}
			if πTemp029, πE = πTemp028.Call(πF, nil, πTemp031); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp028}, πTemp029); πE != nil {
				continue
			}
			if πTemp030, πE = πg.ResolveGlobal(πF, ßtime); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp030, ßresolution, πTemp028); πE != nil {
				continue
			}
			// line 1406: class datetime(date):
			πF.SetLineno(1406)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp030, πE = πg.ResolveGlobal(πF, ßdate); πE != nil {
				continue
			}
			πTemp002[0] = πTemp030
			πTemp027 = πg.NewDict()
			if πTemp028, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp027.SetItem(πF, ß__module__.ToObject(), πTemp028); πE != nil {
				continue
			}
			_, πE = πg.NewCode("datetime", "/usr/lib/python2.7/datetime.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp027
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 []πg.Param
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
				var πTemp030 *πg.Object
				_ = πTemp030
				var πTemp031 *πg.Object
				_ = πTemp031
				var πTemp032 *πg.Object
				_ = πTemp032
				var πTemp033 *πg.Object
				_ = πTemp033
				var πTemp034 *πg.Object
				_ = πTemp034
				var πTemp035 *πg.Object
				_ = πTemp035
				var πTemp036 *πg.Object
				_ = πTemp036
				var πTemp037 *πg.Object
				_ = πTemp037
				var πTemp038 *πg.Object
				_ = πTemp038
				var πTemp039 *πg.Object
				_ = πTemp039
				var πTemp040 *πg.Object
				_ = πTemp040
				var πTemp041 *πg.Object
				_ = πTemp041
				var πTemp042 *πg.Object
				_ = πTemp042
				var πTemp043 *πg.Object
				_ = πTemp043
				var πTemp044 *πg.Object
				_ = πTemp044
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 1407: """datetime(year, month, day[, hour[, minute[, second[, microsecond[,tzinfo]]]]])
					πF.SetLineno(1407)
					// line 1407: """datetime(year, month, day[, hour[, minute[, second[, microsecond[,tzinfo]]]]])
					πF.SetLineno(1407)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("datetime(year, month, day[, hour[, minute[, second[, microsecond[,tzinfo]]]]])\n\n    The year, month and day arguments are required. tzinfo may be None, or an\n    instance of a tzinfo subclass. The remaining arguments may be ints or longs.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 1412: __slots__ = date.__slots__ + time.__slots__
					πF.SetLineno(1412)
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßdate); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ß__slots__, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßtime); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp002, ß__slots__, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, πTemp003, πTemp004); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ß__slots__.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 1414: def __new__(cls, year, month=None, day=None, hour=0, minute=0, second=0,
					πF.SetLineno(1414)
					πTemp005 = make([]πg.Param, 9)
					πTemp005[0] = πg.Param{Name: "cls", Def: nil}
					πTemp005[1] = πg.Param{Name: "year", Def: nil}
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp005[2] = πg.Param{Name: "month", Def: πTemp002}
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp005[3] = πg.Param{Name: "day", Def: πTemp002}
					πTemp005[4] = πg.Param{Name: "hour", Def: πg.NewInt(0).ToObject()}
					πTemp005[5] = πg.Param{Name: "minute", Def: πg.NewInt(0).ToObject()}
					πTemp005[6] = πg.Param{Name: "second", Def: πg.NewInt(0).ToObject()}
					πTemp005[7] = πg.Param{Name: "microsecond", Def: πg.NewInt(0).ToObject()}
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp005[8] = πg.Param{Name: "tzinfo", Def: πTemp002}
					πTemp001 = πg.NewFunction(πg.NewCode("__new__", "/usr/lib/python2.7/datetime.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µcls *πg.Object = πArgs[0]; _ = µcls
						var µyear *πg.Object = πArgs[1]; _ = µyear
						var µmonth *πg.Object = πArgs[2]; _ = µmonth
						var µday *πg.Object = πArgs[3]; _ = µday
						var µhour *πg.Object = πArgs[4]; _ = µhour
						var µminute *πg.Object = πArgs[5]; _ = µminute
						var µsecond *πg.Object = πArgs[6]; _ = µsecond
						var µmicrosecond *πg.Object = πArgs[7]; _ = µmicrosecond
						var µtzinfo *πg.Object = πArgs[8]; _ = µtzinfo
						var µself *πg.Object = πg.UnboundLocal; _ = µself
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 1423: year, month, day = _check_date_fields(year, month, day)
							πF.SetLineno(1423)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µyear, "year"); πE != nil {
								continue
							}
							πTemp001[0] = µyear
							if πE = πg.CheckLocal(πF, µmonth, "month"); πE != nil {
								continue
							}
							πTemp001[1] = µmonth
							if πE = πg.CheckLocal(πF, µday, "day"); πE != nil {
								continue
							}
							πTemp001[2] = µday
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_check_date_fields); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}}}, πTemp003); πE != nil {
								continue
							}
							µyear = πTemp002
							µmonth = πTemp004
							µday = πTemp005
							// line 1424: hour, minute, second, microsecond = _check_time_fields(
							πF.SetLineno(1424)
							πTemp001 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µhour, "hour"); πE != nil {
								continue
							}
							πTemp001[0] = µhour
							if πE = πg.CheckLocal(πF, µminute, "minute"); πE != nil {
								continue
							}
							πTemp001[1] = µminute
							if πE = πg.CheckLocal(πF, µsecond, "second"); πE != nil {
								continue
							}
							πTemp001[2] = µsecond
							if πE = πg.CheckLocal(πF, µmicrosecond, "microsecond"); πE != nil {
								continue
							}
							πTemp001[3] = µmicrosecond
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_check_time_fields); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp006}}}, πTemp003); πE != nil {
								continue
							}
							µhour = πTemp002
							µminute = πTemp004
							µsecond = πTemp005
							µmicrosecond = πTemp006
							// line 1426: _check_tzinfo_arg(tzinfo)
							πF.SetLineno(1426)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtzinfo, "tzinfo"); πE != nil {
								continue
							}
							πTemp001[0] = µtzinfo
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_check_tzinfo_arg); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1427: self = object.__new__(cls)
							πF.SetLineno(1427)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							πTemp001[0] = µcls
							if πTemp002, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ß__new__, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µself = πTemp002
							// line 1428: self._year = year
							πF.SetLineno(1428)
							if πE = πg.CheckLocal(πF, µyear, "year"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µyear); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_year, πTemp002); πE != nil {
								continue
							}
							// line 1429: self._month = month
							πF.SetLineno(1429)
							if πE = πg.CheckLocal(πF, µmonth, "month"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µmonth); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_month, πTemp002); πE != nil {
								continue
							}
							// line 1430: self._day = day
							πF.SetLineno(1430)
							if πE = πg.CheckLocal(πF, µday, "day"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µday); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_day, πTemp002); πE != nil {
								continue
							}
							// line 1431: self._hour = hour
							πF.SetLineno(1431)
							if πE = πg.CheckLocal(πF, µhour, "hour"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µhour); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_hour, πTemp002); πE != nil {
								continue
							}
							// line 1432: self._minute = minute
							πF.SetLineno(1432)
							if πE = πg.CheckLocal(πF, µminute, "minute"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µminute); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_minute, πTemp002); πE != nil {
								continue
							}
							// line 1433: self._second = second
							πF.SetLineno(1433)
							if πE = πg.CheckLocal(πF, µsecond, "second"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µsecond); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_second, πTemp002); πE != nil {
								continue
							}
							// line 1434: self._microsecond = microsecond
							πF.SetLineno(1434)
							if πE = πg.CheckLocal(πF, µmicrosecond, "microsecond"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µmicrosecond); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_microsecond, πTemp002); πE != nil {
								continue
							}
							// line 1435: self._tzinfo = tzinfo
							πF.SetLineno(1435)
							if πE = πg.CheckLocal(πF, µtzinfo, "tzinfo"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µtzinfo); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_tzinfo, πTemp002); πE != nil {
								continue
							}
							// line 1436: self._hashcode = -1
							πF.SetLineno(1436)
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_hashcode, πTemp003); πE != nil {
								continue
							}
							// line 1437: return self
							πF.SetLineno(1437)
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
					if πE = πClass.SetItem(πF, ß__new__.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 1441: def hour(self):
					πF.SetLineno(1441)
					πTemp005 = make([]πg.Param, 1)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp002 = πg.NewFunction(πg.NewCode("hour", "/usr/lib/python2.7/datetime.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1442: """hour (0-23)"""
							πF.SetLineno(1442)
							// line 1443: return self._hour
							πF.SetLineno(1443)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_hour, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßhour.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 1442: """hour (0-23)"""
					πF.SetLineno(1442)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewStr("hour (0-23)").ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßhour); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp004, ß__doc__, πTemp003); πE != nil {
						continue
					}
					// line 1440: @property
					πF.SetLineno(1440)
					πTemp006 = πF.MakeArgs(1)
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßhour); πE != nil {
						continue
					}
					πTemp006[0] = πTemp003
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßproperty); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πE = πClass.SetItem(πF, ßhour.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 1446: def minute(self):
					πF.SetLineno(1446)
					πTemp005 = make([]πg.Param, 1)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("minute", "/usr/lib/python2.7/datetime.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1447: """minute (0-59)"""
							πF.SetLineno(1447)
							// line 1448: return self._minute
							πF.SetLineno(1448)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_minute, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßminute.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 1447: """minute (0-59)"""
					πF.SetLineno(1447)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πg.NewStr("minute (0-59)").ToObject()); πE != nil {
						continue
					}
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßminute); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp007, ß__doc__, πTemp004); πE != nil {
						continue
					}
					// line 1445: @property
					πF.SetLineno(1445)
					πTemp006 = πF.MakeArgs(1)
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßminute); πE != nil {
						continue
					}
					πTemp006[0] = πTemp004
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßproperty); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πE = πClass.SetItem(πF, ßminute.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 1451: def second(self):
					πF.SetLineno(1451)
					πTemp005 = make([]πg.Param, 1)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("second", "/usr/lib/python2.7/datetime.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1452: """second (0-59)"""
							πF.SetLineno(1452)
							// line 1453: return self._second
							πF.SetLineno(1453)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_second, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßsecond.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 1452: """second (0-59)"""
					πF.SetLineno(1452)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp007}, πg.NewStr("second (0-59)").ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ßsecond); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp008, ß__doc__, πTemp007); πE != nil {
						continue
					}
					// line 1450: @property
					πF.SetLineno(1450)
					πTemp006 = πF.MakeArgs(1)
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßsecond); πE != nil {
						continue
					}
					πTemp006[0] = πTemp007
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßproperty); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp007.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πE = πClass.SetItem(πF, ßsecond.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 1456: def microsecond(self):
					πF.SetLineno(1456)
					πTemp005 = make([]πg.Param, 1)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("microsecond", "/usr/lib/python2.7/datetime.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1457: """microsecond (0-999999)"""
							πF.SetLineno(1457)
							// line 1458: return self._microsecond
							πF.SetLineno(1458)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_microsecond, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßmicrosecond.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 1457: """microsecond (0-999999)"""
					πF.SetLineno(1457)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp008}, πg.NewStr("microsecond (0-999999)").ToObject()); πE != nil {
						continue
					}
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßmicrosecond); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp009, ß__doc__, πTemp008); πE != nil {
						continue
					}
					// line 1455: @property
					πF.SetLineno(1455)
					πTemp006 = πF.MakeArgs(1)
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ßmicrosecond); πE != nil {
						continue
					}
					πTemp006[0] = πTemp008
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ßproperty); πE != nil {
						continue
					}
					if πTemp009, πE = πTemp008.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πE = πClass.SetItem(πF, ßmicrosecond.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 1461: def tzinfo(self):
					πF.SetLineno(1461)
					πTemp005 = make([]πg.Param, 1)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("tzinfo", "/usr/lib/python2.7/datetime.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1462: """timezone info object"""
							πF.SetLineno(1462)
							// line 1463: return self._tzinfo
							πF.SetLineno(1463)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_tzinfo, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtzinfo.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 1462: """timezone info object"""
					πF.SetLineno(1462)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp009}, πg.NewStr("timezone info object").ToObject()); πE != nil {
						continue
					}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßtzinfo); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp010, ß__doc__, πTemp009); πE != nil {
						continue
					}
					// line 1460: @property
					πF.SetLineno(1460)
					πTemp006 = πF.MakeArgs(1)
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßtzinfo); πE != nil {
						continue
					}
					πTemp006[0] = πTemp009
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßproperty); πE != nil {
						continue
					}
					if πTemp010, πE = πTemp009.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πE = πClass.SetItem(πF, ßtzinfo.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 1466: def fromtimestamp(cls, timestamp, tz=None):
					πF.SetLineno(1466)
					πTemp005 = make([]πg.Param, 3)
					πTemp005[0] = πg.Param{Name: "cls", Def: nil}
					πTemp005[1] = πg.Param{Name: "timestamp", Def: nil}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp005[2] = πg.Param{Name: "tz", Def: πTemp010}
					πTemp009 = πg.NewFunction(πg.NewCode("fromtimestamp", "/usr/lib/python2.7/datetime.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µcls *πg.Object = πArgs[0]; _ = µcls
						var µtimestamp *πg.Object = πArgs[1]; _ = µtimestamp
						var µtz *πg.Object = πArgs[2]; _ = µtz
						var µconverter *πg.Object = πg.UnboundLocal; _ = µconverter
						var µself *πg.Object = πg.UnboundLocal; _ = µself
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
							// line 1467: """Construct a datetime from a POSIX timestamp (like time.time()).
							πF.SetLineno(1467)
							// line 1471: _check_tzinfo_arg(tz)
							πF.SetLineno(1471)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtz, "tz"); πE != nil {
								continue
							}
							πTemp001[0] = µtz
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_check_tzinfo_arg); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1472: converter = _time.localtime if tz is None else _time.gmtime
							πF.SetLineno(1472)
							if πE = πg.CheckLocal(πF, µtz, "tz"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(µtz == πTemp004).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if !πTemp005 {
								goto Label1
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ß_time); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßlocaltime, nil); πE != nil {
								continue
							}
							πTemp002 = πTemp004
							goto Label2
						Label1:
							if πTemp003, πE = πg.ResolveGlobal(πF, ß_time); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßgmtime, nil); πE != nil {
								continue
							}
							πTemp002 = πTemp004
						Label2:
							µconverter = πTemp002
							// line 1473: self = cls._from_timestamp(converter, timestamp, tz)
							πF.SetLineno(1473)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µconverter, "converter"); πE != nil {
								continue
							}
							πTemp001[0] = µconverter
							if πE = πg.CheckLocal(πF, µtimestamp, "timestamp"); πE != nil {
								continue
							}
							πTemp001[1] = µtimestamp
							if πE = πg.CheckLocal(πF, µtz, "tz"); πE != nil {
								continue
							}
							πTemp001[2] = µtz
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µcls, ß_from_timestamp, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µself = πTemp003
							if πE = πg.CheckLocal(πF, µtz, "tz"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(µtz != πTemp003).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label3
							}
							goto Label4
							// line 1474: if tz is not None:
							πF.SetLineno(1474)
						Label3:
							// line 1475: self = tz.fromutc(self)
							πF.SetLineno(1475)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µtz, "tz"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µtz, ßfromutc, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µself = πTemp003
							goto Label4
						Label4:
							// line 1476: return self
							πF.SetLineno(1476)
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
					if πE = πClass.SetItem(πF, ßfromtimestamp.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 1467: """Construct a datetime from a POSIX timestamp (like time.time()).
					πF.SetLineno(1467)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp010}, πg.NewStr("Construct a datetime from a POSIX timestamp (like time.time()).\n\n        A timezone info object may be passed in as well.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ßfromtimestamp); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp011, ß__doc__, πTemp010); πE != nil {
						continue
					}
					// line 1465: @classmethod
					πF.SetLineno(1465)
					πTemp006 = πF.MakeArgs(1)
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßfromtimestamp); πE != nil {
						continue
					}
					πTemp006[0] = πTemp010
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßclassmethod); πE != nil {
						continue
					}
					if πTemp011, πE = πTemp010.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πE = πClass.SetItem(πF, ßfromtimestamp.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 1479: def utcfromtimestamp(cls, t):
					πF.SetLineno(1479)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "cls", Def: nil}
					πTemp005[1] = πg.Param{Name: "t", Def: nil}
					πTemp010 = πg.NewFunction(πg.NewCode("utcfromtimestamp", "/usr/lib/python2.7/datetime.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µcls *πg.Object = πArgs[0]; _ = µcls
						var µt *πg.Object = πArgs[1]; _ = µt
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
							// line 1480: "Construct a UTC datetime from a POSIX timestamp (like time.time())."
							πF.SetLineno(1480)
							// line 1481: return cls._from_timestamp(_time.gmtime, t, None)
							πF.SetLineno(1481)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_time); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßgmtime, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
								continue
							}
							πTemp001[1] = µt
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[2] = πTemp002
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µcls, ß_from_timestamp, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßutcfromtimestamp.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 1480: "Construct a UTC datetime from a POSIX timestamp (like time.time())."
					πF.SetLineno(1480)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πg.NewStr("Construct a UTC datetime from a POSIX timestamp (like time.time()).").ToObject()); πE != nil {
						continue
					}
					if πTemp012, πE = πg.ResolveClass(πF, πClass, nil, ßutcfromtimestamp); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp012, ß__doc__, πTemp011); πE != nil {
						continue
					}
					// line 1478: @classmethod
					πF.SetLineno(1478)
					πTemp006 = πF.MakeArgs(1)
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ßutcfromtimestamp); πE != nil {
						continue
					}
					πTemp006[0] = πTemp011
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ßclassmethod); πE != nil {
						continue
					}
					if πTemp012, πE = πTemp011.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πE = πClass.SetItem(πF, ßutcfromtimestamp.ToObject(), πTemp012); πE != nil {
						continue
					}
					// line 1484: def _from_timestamp(cls, converter, timestamp, tzinfo):
					πF.SetLineno(1484)
					πTemp005 = make([]πg.Param, 4)
					πTemp005[0] = πg.Param{Name: "cls", Def: nil}
					πTemp005[1] = πg.Param{Name: "converter", Def: nil}
					πTemp005[2] = πg.Param{Name: "timestamp", Def: nil}
					πTemp005[3] = πg.Param{Name: "tzinfo", Def: nil}
					πTemp011 = πg.NewFunction(πg.NewCode("_from_timestamp", "/usr/lib/python2.7/datetime.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µcls *πg.Object = πArgs[0]; _ = µcls
						var µconverter *πg.Object = πArgs[1]; _ = µconverter
						var µtimestamp *πg.Object = πArgs[2]; _ = µtimestamp
						var µtzinfo *πg.Object = πArgs[3]; _ = µtzinfo
						var µt_full *πg.Object = πg.UnboundLocal; _ = µt_full
						var µfrac *πg.Object = πg.UnboundLocal; _ = µfrac
						var µus *πg.Object = πg.UnboundLocal; _ = µus
						var µy *πg.Object = πg.UnboundLocal; _ = µy
						var µm *πg.Object = πg.UnboundLocal; _ = µm
						var µd *πg.Object = πg.UnboundLocal; _ = µd
						var µhh *πg.Object = πg.UnboundLocal; _ = µhh
						var µmm *πg.Object = πg.UnboundLocal; _ = µmm
						var µss *πg.Object = πg.UnboundLocal; _ = µss
						var µweekday *πg.Object = πg.UnboundLocal; _ = µweekday
						var µjday *πg.Object = πg.UnboundLocal; _ = µjday
						var µdst *πg.Object = πg.UnboundLocal; _ = µdst
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
							default: panic("unexpected function state")
							}
							// line 1485: t_full = timestamp
							πF.SetLineno(1485)
							if πE = πg.CheckLocal(πF, µtimestamp, "timestamp"); πE != nil {
								continue
							}
							µt_full = µtimestamp
							// line 1486: timestamp = int(_math.floor(timestamp))
							πF.SetLineno(1486)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtimestamp, "timestamp"); πE != nil {
								continue
							}
							πTemp002[0] = µtimestamp
							if πTemp003, πE = πg.ResolveGlobal(πF, ß_math); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßfloor, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µtimestamp = πTemp004
							// line 1487: frac = t_full - timestamp
							πF.SetLineno(1487)
							if πE = πg.CheckLocal(πF, µt_full, "t_full"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtimestamp, "timestamp"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Sub(πF, µt_full, µtimestamp); πE != nil {
								continue
							}
							µfrac = πTemp003
							// line 1488: us = _round(frac * 1e6)
							πF.SetLineno(1488)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfrac, "frac"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Mul(πF, µfrac, πg.NewFloat(1000000.0).ToObject()); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ß_round); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µus = πTemp004
							if πE = πg.CheckLocal(πF, µus, "us"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Eq(πF, µus, πg.NewInt(1000000).ToObject()); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							goto Label2
							// line 1494: if us == 1000000:
							πF.SetLineno(1494)
						Label1:
							// line 1495: timestamp += 1
							πF.SetLineno(1495)
							if πE = πg.CheckLocal(πF, µtimestamp, "timestamp"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IAdd(πF, µtimestamp, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µtimestamp = πTemp003
							// line 1496: us = 0
							πF.SetLineno(1496)
							µus = πg.NewInt(0).ToObject()
							goto Label2
						Label2:
							// line 1497: y, m, d, hh, mm, ss, weekday, jday, dst = converter(timestamp)
							πF.SetLineno(1497)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtimestamp, "timestamp"); πE != nil {
								continue
							}
							πTemp001[0] = µtimestamp
							if πE = πg.CheckLocal(πF, µconverter, "converter"); πE != nil {
								continue
							}
							if πTemp003, πE = µconverter.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp006}, πg.TieTarget{Target: &πTemp007}, πg.TieTarget{Target: &πTemp008}, πg.TieTarget{Target: &πTemp009}, πg.TieTarget{Target: &πTemp010}, πg.TieTarget{Target: &πTemp011}, πg.TieTarget{Target: &πTemp012}, πg.TieTarget{Target: &πTemp013}}}, πTemp003); πE != nil {
								continue
							}
							µy = πTemp004
							µm = πTemp006
							µd = πTemp007
							µhh = πTemp008
							µmm = πTemp009
							µss = πTemp010
							µweekday = πTemp011
							µjday = πTemp012
							µdst = πTemp013
							// line 1498: ss = min(ss, 59)    # clamp out leap seconds if the platform has them
							πF.SetLineno(1498)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
								continue
							}
							πTemp001[0] = µss
							πTemp001[1] = πg.NewInt(59).ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßmin); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µss = πTemp004
							// line 1499: return cls(y, m, d, hh, mm, ss, us, tzinfo)
							πF.SetLineno(1499)
							πTemp001 = πF.MakeArgs(8)
							if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
								continue
							}
							πTemp001[0] = µy
							if πE = πg.CheckLocal(πF, µm, "m"); πE != nil {
								continue
							}
							πTemp001[1] = µm
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp001[2] = µd
							if πE = πg.CheckLocal(πF, µhh, "hh"); πE != nil {
								continue
							}
							πTemp001[3] = µhh
							if πE = πg.CheckLocal(πF, µmm, "mm"); πE != nil {
								continue
							}
							πTemp001[4] = µmm
							if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
								continue
							}
							πTemp001[5] = µss
							if πE = πg.CheckLocal(πF, µus, "us"); πE != nil {
								continue
							}
							πTemp001[6] = µus
							if πE = πg.CheckLocal(πF, µtzinfo, "tzinfo"); πE != nil {
								continue
							}
							πTemp001[7] = µtzinfo
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							if πTemp003, πE = µcls.Call(πF, πTemp001, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ß_from_timestamp.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 1483: @classmethod
					πF.SetLineno(1483)
					πTemp006 = πF.MakeArgs(1)
					if πTemp012, πE = πg.ResolveClass(πF, πClass, nil, ß_from_timestamp); πE != nil {
						continue
					}
					πTemp006[0] = πTemp012
					if πTemp012, πE = πg.ResolveClass(πF, πClass, nil, ßclassmethod); πE != nil {
						continue
					}
					if πTemp013, πE = πTemp012.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πE = πClass.SetItem(πF, ß_from_timestamp.ToObject(), πTemp013); πE != nil {
						continue
					}
					// line 1502: def now(cls, tz=None):
					πF.SetLineno(1502)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "cls", Def: nil}
					if πTemp013, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp005[1] = πg.Param{Name: "tz", Def: πTemp013}
					πTemp012 = πg.NewFunction(πg.NewCode("now", "/usr/lib/python2.7/datetime.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µcls *πg.Object = πArgs[0]; _ = µcls
						var µtz *πg.Object = πArgs[1]; _ = µtz
						var µt *πg.Object = πg.UnboundLocal; _ = µt
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 1503: "Construct a datetime from time.time() and optional time zone info."
							πF.SetLineno(1503)
							// line 1504: t = _time.time()
							πF.SetLineno(1504)
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_time); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßtime, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µt = πTemp001
							// line 1505: return cls.fromtimestamp(t, tz)
							πF.SetLineno(1505)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
								continue
							}
							πTemp003[0] = µt
							if πE = πg.CheckLocal(πF, µtz, "tz"); πE != nil {
								continue
							}
							πTemp003[1] = µtz
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µcls, ßfromtimestamp, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
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
					if πE = πClass.SetItem(πF, ßnow.ToObject(), πTemp012); πE != nil {
						continue
					}
					// line 1503: "Construct a datetime from time.time() and optional time zone info."
					πF.SetLineno(1503)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp013}, πg.NewStr("Construct a datetime from time.time() and optional time zone info.").ToObject()); πE != nil {
						continue
					}
					if πTemp014, πE = πg.ResolveClass(πF, πClass, nil, ßnow); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp014, ß__doc__, πTemp013); πE != nil {
						continue
					}
					// line 1501: @classmethod
					πF.SetLineno(1501)
					πTemp006 = πF.MakeArgs(1)
					if πTemp013, πE = πg.ResolveClass(πF, πClass, nil, ßnow); πE != nil {
						continue
					}
					πTemp006[0] = πTemp013
					if πTemp013, πE = πg.ResolveClass(πF, πClass, nil, ßclassmethod); πE != nil {
						continue
					}
					if πTemp014, πE = πTemp013.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πE = πClass.SetItem(πF, ßnow.ToObject(), πTemp014); πE != nil {
						continue
					}
					// line 1508: def utcnow(cls):
					πF.SetLineno(1508)
					πTemp005 = make([]πg.Param, 1)
					πTemp005[0] = πg.Param{Name: "cls", Def: nil}
					πTemp013 = πg.NewFunction(πg.NewCode("utcnow", "/usr/lib/python2.7/datetime.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µcls *πg.Object = πArgs[0]; _ = µcls
						var µt *πg.Object = πg.UnboundLocal; _ = µt
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 1509: "Construct a UTC datetime from time.time()."
							πF.SetLineno(1509)
							// line 1510: t = _time.time()
							πF.SetLineno(1510)
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_time); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßtime, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µt = πTemp001
							// line 1511: return cls.utcfromtimestamp(t)
							πF.SetLineno(1511)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
								continue
							}
							πTemp003[0] = µt
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µcls, ßutcfromtimestamp, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
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
					if πE = πClass.SetItem(πF, ßutcnow.ToObject(), πTemp013); πE != nil {
						continue
					}
					// line 1509: "Construct a UTC datetime from time.time()."
					πF.SetLineno(1509)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp014}, πg.NewStr("Construct a UTC datetime from time.time().").ToObject()); πE != nil {
						continue
					}
					if πTemp015, πE = πg.ResolveClass(πF, πClass, nil, ßutcnow); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp015, ß__doc__, πTemp014); πE != nil {
						continue
					}
					// line 1507: @classmethod
					πF.SetLineno(1507)
					πTemp006 = πF.MakeArgs(1)
					if πTemp014, πE = πg.ResolveClass(πF, πClass, nil, ßutcnow); πE != nil {
						continue
					}
					πTemp006[0] = πTemp014
					if πTemp014, πE = πg.ResolveClass(πF, πClass, nil, ßclassmethod); πE != nil {
						continue
					}
					if πTemp015, πE = πTemp014.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πE = πClass.SetItem(πF, ßutcnow.ToObject(), πTemp015); πE != nil {
						continue
					}
					// line 1514: def combine(cls, date, time):
					πF.SetLineno(1514)
					πTemp005 = make([]πg.Param, 3)
					πTemp005[0] = πg.Param{Name: "cls", Def: nil}
					πTemp005[1] = πg.Param{Name: "date", Def: nil}
					πTemp005[2] = πg.Param{Name: "time", Def: nil}
					πTemp014 = πg.NewFunction(πg.NewCode("combine", "/usr/lib/python2.7/datetime.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µcls *πg.Object = πArgs[0]; _ = µcls
						var µdate *πg.Object = πArgs[1]; _ = µdate
						var µtime *πg.Object = πArgs[2]; _ = µtime
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 1515: "Construct a datetime from a given date and a given time."
							πF.SetLineno(1515)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µdate, "date"); πE != nil {
								continue
							}
							πTemp002[0] = µdate
							if πTemp003, πE = πg.ResolveGlobal(πF, ß_date_class); πE != nil {
								continue
							}
							πTemp002[1] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp005, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp005).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							goto Label2
							// line 1516: if not isinstance(date, _date_class):
							πF.SetLineno(1516)
						Label1:
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("date argument must be a date instance").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 1517: raise TypeError("date argument must be a date instance")
							πF.SetLineno(1517)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label2
						Label2:
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µtime, "time"); πE != nil {
								continue
							}
							πTemp002[0] = µtime
							if πTemp003, πE = πg.ResolveGlobal(πF, ß_time_class); πE != nil {
								continue
							}
							πTemp002[1] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp005, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp005).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label3
							}
							goto Label4
							// line 1518: if not isinstance(time, _time_class):
							πF.SetLineno(1518)
						Label3:
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("time argument must be a time instance").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 1519: raise TypeError("time argument must be a time instance")
							πF.SetLineno(1519)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label4
						Label4:
							// line 1520: return cls(date.year, date.month, date.day,
							πF.SetLineno(1520)
							πTemp002 = πF.MakeArgs(8)
							if πE = πg.CheckLocal(πF, µdate, "date"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µdate, ßyear, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							if πE = πg.CheckLocal(πF, µdate, "date"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µdate, ßmonth, nil); πE != nil {
								continue
							}
							πTemp002[1] = πTemp001
							if πE = πg.CheckLocal(πF, µdate, "date"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µdate, ßday, nil); πE != nil {
								continue
							}
							πTemp002[2] = πTemp001
							if πE = πg.CheckLocal(πF, µtime, "time"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µtime, ßhour, nil); πE != nil {
								continue
							}
							πTemp002[3] = πTemp001
							if πE = πg.CheckLocal(πF, µtime, "time"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µtime, ßminute, nil); πE != nil {
								continue
							}
							πTemp002[4] = πTemp001
							if πE = πg.CheckLocal(πF, µtime, "time"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µtime, ßsecond, nil); πE != nil {
								continue
							}
							πTemp002[5] = πTemp001
							if πE = πg.CheckLocal(πF, µtime, "time"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µtime, ßmicrosecond, nil); πE != nil {
								continue
							}
							πTemp002[6] = πTemp001
							if πE = πg.CheckLocal(πF, µtime, "time"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µtime, ßtzinfo, nil); πE != nil {
								continue
							}
							πTemp002[7] = πTemp001
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							if πTemp001, πE = µcls.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
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
					if πE = πClass.SetItem(πF, ßcombine.ToObject(), πTemp014); πE != nil {
						continue
					}
					// line 1515: "Construct a datetime from a given date and a given time."
					πF.SetLineno(1515)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp015}, πg.NewStr("Construct a datetime from a given date and a given time.").ToObject()); πE != nil {
						continue
					}
					if πTemp016, πE = πg.ResolveClass(πF, πClass, nil, ßcombine); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp016, ß__doc__, πTemp015); πE != nil {
						continue
					}
					// line 1513: @classmethod
					πF.SetLineno(1513)
					πTemp006 = πF.MakeArgs(1)
					if πTemp015, πE = πg.ResolveClass(πF, πClass, nil, ßcombine); πE != nil {
						continue
					}
					πTemp006[0] = πTemp015
					if πTemp015, πE = πg.ResolveClass(πF, πClass, nil, ßclassmethod); πE != nil {
						continue
					}
					if πTemp016, πE = πTemp015.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πE = πClass.SetItem(πF, ßcombine.ToObject(), πTemp016); πE != nil {
						continue
					}
					// line 1524: def timetuple(self):
					πF.SetLineno(1524)
					πTemp005 = make([]πg.Param, 1)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp015 = πg.NewFunction(πg.NewCode("timetuple", "/usr/lib/python2.7/datetime.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µdst *πg.Object = πg.UnboundLocal; _ = µdst
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
							// line 1525: "Return local time tuple compatible with time.localtime()."
							πF.SetLineno(1525)
							// line 1526: dst = self._dst()
							πF.SetLineno(1526)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_dst, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µdst = πTemp002
							if πE = πg.CheckLocal(πF, µdst, "dst"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µdst == πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µdst, "dst"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, µdst); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label2
							}
							goto Label3
							// line 1527: if dst is None:
							πF.SetLineno(1527)
						Label1:
							// line 1528: dst = -1
							πF.SetLineno(1528)
							if πTemp001, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µdst = πTemp001
							goto Label3
							// line 1529: elif dst:
							πF.SetLineno(1529)
						Label2:
							// line 1530: dst = 1
							πF.SetLineno(1530)
							µdst = πg.NewInt(1).ToObject()
							goto Label3
						Label3:
							// line 1531: return _build_struct_time(self.year, self.month, self.day,
							πF.SetLineno(1531)
							πTemp004 = πF.MakeArgs(7)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßyear, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßmonth, nil); πE != nil {
								continue
							}
							πTemp004[1] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßday, nil); πE != nil {
								continue
							}
							πTemp004[2] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßhour, nil); πE != nil {
								continue
							}
							πTemp004[3] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßminute, nil); πE != nil {
								continue
							}
							πTemp004[4] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßsecond, nil); πE != nil {
								continue
							}
							πTemp004[5] = πTemp001
							if πE = πg.CheckLocal(πF, µdst, "dst"); πE != nil {
								continue
							}
							πTemp004[6] = µdst
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_build_struct_time); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
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
					if πE = πClass.SetItem(πF, ßtimetuple.ToObject(), πTemp015); πE != nil {
						continue
					}
					// line 1525: "Return local time tuple compatible with time.localtime()."
					πF.SetLineno(1525)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp016}, πg.NewStr("Return local time tuple compatible with time.localtime().").ToObject()); πE != nil {
						continue
					}
					if πTemp017, πE = πg.ResolveClass(πF, πClass, nil, ßtimetuple); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp017, ß__doc__, πTemp016); πE != nil {
						continue
					}
					// line 1535: def utctimetuple(self):
					πF.SetLineno(1535)
					πTemp005 = make([]πg.Param, 1)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp016 = πg.NewFunction(πg.NewCode("utctimetuple", "/usr/lib/python2.7/datetime.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µy *πg.Object = πg.UnboundLocal; _ = µy
						var µm *πg.Object = πg.UnboundLocal; _ = µm
						var µd *πg.Object = πg.UnboundLocal; _ = µd
						var µhh *πg.Object = πg.UnboundLocal; _ = µhh
						var µmm *πg.Object = πg.UnboundLocal; _ = µmm
						var µss *πg.Object = πg.UnboundLocal; _ = µss
						var µoffset *πg.Object = πg.UnboundLocal; _ = µoffset
						var µ_ *πg.Object = πg.UnboundLocal; _ = µ_
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
						var πTemp007 πg.KWArgs
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 *πg.Object
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
							// line 1536: "Return UTC time tuple compatible with time.gmtime()."
							πF.SetLineno(1536)
							// line 1537: y, m, d = self.year, self.month, self.day
							πF.SetLineno(1537)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßyear, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmonth, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßday, nil); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple3(πTemp002, πTemp003, πTemp004).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp004}}}, πTemp001); πE != nil {
								continue
							}
							µy = πTemp002
							µm = πTemp003
							µd = πTemp004
							// line 1538: hh, mm, ss = self.hour, self.minute, self.second
							πF.SetLineno(1538)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßhour, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßminute, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßsecond, nil); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple3(πTemp002, πTemp003, πTemp004).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp004}}}, πTemp001); πE != nil {
								continue
							}
							µhh = πTemp002
							µmm = πTemp003
							µss = πTemp004
							// line 1539: offset = self._utcoffset()
							πF.SetLineno(1539)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_utcoffset, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µoffset = πTemp002
							if πE = πg.CheckLocal(πF, µoffset, "offset"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, µoffset); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							goto Label2
							// line 1540: if offset:  # neither None nor 0
							πF.SetLineno(1540)
						Label1:
							// line 1541: mm -= offset
							πF.SetLineno(1541)
							if πE = πg.CheckLocal(πF, µmm, "mm"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µoffset, "offset"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.ISub(πF, µmm, µoffset); πE != nil {
								continue
							}
							µmm = πTemp001
							// line 1542: y, m, d, hh, mm, ss, _ = _normalize_datetime(
							πF.SetLineno(1542)
							πTemp006 = πF.MakeArgs(7)
							if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
								continue
							}
							πTemp006[0] = µy
							if πE = πg.CheckLocal(πF, µm, "m"); πE != nil {
								continue
							}
							πTemp006[1] = µm
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp006[2] = µd
							if πE = πg.CheckLocal(πF, µhh, "hh"); πE != nil {
								continue
							}
							πTemp006[3] = µhh
							if πE = πg.CheckLocal(πF, µmm, "mm"); πE != nil {
								continue
							}
							πTemp006[4] = µmm
							if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
								continue
							}
							πTemp006[5] = µss
							πTemp006[6] = πg.NewInt(0).ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp007 = πg.KWArgs{
								{"ignore_overflow", πTemp001},
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_normalize_datetime); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp006, πTemp007); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp008}, πg.TieTarget{Target: &πTemp009}, πg.TieTarget{Target: &πTemp010}, πg.TieTarget{Target: &πTemp011}}}, πTemp002); πE != nil {
								continue
							}
							µy = πTemp001
							µm = πTemp003
							µd = πTemp004
							µhh = πTemp008
							µmm = πTemp009
							µss = πTemp010
							µ_ = πTemp011
							goto Label2
						Label2:
							// line 1544: return _build_struct_time(y, m, d, hh, mm, ss, 0)
							πF.SetLineno(1544)
							πTemp006 = πF.MakeArgs(7)
							if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
								continue
							}
							πTemp006[0] = µy
							if πE = πg.CheckLocal(πF, µm, "m"); πE != nil {
								continue
							}
							πTemp006[1] = µm
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp006[2] = µd
							if πE = πg.CheckLocal(πF, µhh, "hh"); πE != nil {
								continue
							}
							πTemp006[3] = µhh
							if πE = πg.CheckLocal(πF, µmm, "mm"); πE != nil {
								continue
							}
							πTemp006[4] = µmm
							if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
								continue
							}
							πTemp006[5] = µss
							πTemp006[6] = πg.NewInt(0).ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_build_struct_time); πE != nil {
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
					if πE = πClass.SetItem(πF, ßutctimetuple.ToObject(), πTemp016); πE != nil {
						continue
					}
					// line 1536: "Return UTC time tuple compatible with time.gmtime()."
					πF.SetLineno(1536)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp017}, πg.NewStr("Return UTC time tuple compatible with time.gmtime().").ToObject()); πE != nil {
						continue
					}
					if πTemp018, πE = πg.ResolveClass(πF, πClass, nil, ßutctimetuple); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp018, ß__doc__, πTemp017); πE != nil {
						continue
					}
					// line 1546: def date(self):
					πF.SetLineno(1546)
					πTemp005 = make([]πg.Param, 1)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp017 = πg.NewFunction(πg.NewCode("date", "/usr/lib/python2.7/datetime.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1547: "Return the date part."
							πF.SetLineno(1547)
							// line 1548: return date(self._year, self._month, self._day)
							πF.SetLineno(1548)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_year, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_month, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_day, nil); πE != nil {
								continue
							}
							πTemp001[2] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßdate); πE != nil {
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
					if πE = πClass.SetItem(πF, ßdate.ToObject(), πTemp017); πE != nil {
						continue
					}
					// line 1547: "Return the date part."
					πF.SetLineno(1547)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp018}, πg.NewStr("Return the date part.").ToObject()); πE != nil {
						continue
					}
					if πTemp019, πE = πg.ResolveClass(πF, πClass, nil, ßdate); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp019, ß__doc__, πTemp018); πE != nil {
						continue
					}
					// line 1550: def time(self):
					πF.SetLineno(1550)
					πTemp005 = make([]πg.Param, 1)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp018 = πg.NewFunction(πg.NewCode("time", "/usr/lib/python2.7/datetime.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1551: "Return the time part, with tzinfo None."
							πF.SetLineno(1551)
							// line 1552: return time(self.hour, self.minute, self.second, self.microsecond)
							πF.SetLineno(1552)
							πTemp001 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßhour, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßminute, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßsecond, nil); πE != nil {
								continue
							}
							πTemp001[2] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßmicrosecond, nil); πE != nil {
								continue
							}
							πTemp001[3] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtime); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtime.ToObject(), πTemp018); πE != nil {
						continue
					}
					// line 1551: "Return the time part, with tzinfo None."
					πF.SetLineno(1551)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp019}, πg.NewStr("Return the time part, with tzinfo None.").ToObject()); πE != nil {
						continue
					}
					if πTemp020, πE = πg.ResolveClass(πF, πClass, nil, ßtime); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp020, ß__doc__, πTemp019); πE != nil {
						continue
					}
					// line 1554: def timetz(self):
					πF.SetLineno(1554)
					πTemp005 = make([]πg.Param, 1)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp019 = πg.NewFunction(πg.NewCode("timetz", "/usr/lib/python2.7/datetime.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1555: "Return the time part, with same tzinfo."
							πF.SetLineno(1555)
							// line 1556: return time(self.hour, self.minute, self.second, self.microsecond,
							πF.SetLineno(1556)
							πTemp001 = πF.MakeArgs(5)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßhour, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßminute, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßsecond, nil); πE != nil {
								continue
							}
							πTemp001[2] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßmicrosecond, nil); πE != nil {
								continue
							}
							πTemp001[3] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_tzinfo, nil); πE != nil {
								continue
							}
							πTemp001[4] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtime); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtimetz.ToObject(), πTemp019); πE != nil {
						continue
					}
					// line 1555: "Return the time part, with same tzinfo."
					πF.SetLineno(1555)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp020}, πg.NewStr("Return the time part, with same tzinfo.").ToObject()); πE != nil {
						continue
					}
					if πTemp021, πE = πg.ResolveClass(πF, πClass, nil, ßtimetz); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp021, ß__doc__, πTemp020); πE != nil {
						continue
					}
					// line 1559: def replace(self, year=None, month=None, day=None, hour=None,
					πF.SetLineno(1559)
					πTemp005 = make([]πg.Param, 9)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					if πTemp021, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp005[1] = πg.Param{Name: "year", Def: πTemp021}
					if πTemp021, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp005[2] = πg.Param{Name: "month", Def: πTemp021}
					if πTemp021, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp005[3] = πg.Param{Name: "day", Def: πTemp021}
					if πTemp021, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp005[4] = πg.Param{Name: "hour", Def: πTemp021}
					if πTemp021, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp005[5] = πg.Param{Name: "minute", Def: πTemp021}
					if πTemp021, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp005[6] = πg.Param{Name: "second", Def: πTemp021}
					if πTemp021, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp005[7] = πg.Param{Name: "microsecond", Def: πTemp021}
					if πTemp021, πE = πg.ResolveClass(πF, πClass, nil, ßTrue); πE != nil {
						continue
					}
					πTemp005[8] = πg.Param{Name: "tzinfo", Def: πTemp021}
					πTemp020 = πg.NewFunction(πg.NewCode("replace", "/usr/lib/python2.7/datetime.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µyear *πg.Object = πArgs[1]; _ = µyear
						var µmonth *πg.Object = πArgs[2]; _ = µmonth
						var µday *πg.Object = πArgs[3]; _ = µday
						var µhour *πg.Object = πArgs[4]; _ = µhour
						var µminute *πg.Object = πArgs[5]; _ = µminute
						var µsecond *πg.Object = πArgs[6]; _ = µsecond
						var µmicrosecond *πg.Object = πArgs[7]; _ = µmicrosecond
						var µtzinfo *πg.Object = πArgs[8]; _ = µtzinfo
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
							// line 1561: """Return a new datetime with new values for the specified fields."""
							πF.SetLineno(1561)
							if πE = πg.CheckLocal(πF, µyear, "year"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µyear == πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 1562: if year is None:
							πF.SetLineno(1562)
						Label1:
							// line 1563: year = self.year
							πF.SetLineno(1563)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßyear, nil); πE != nil {
								continue
							}
							µyear = πTemp001
							goto Label2
						Label2:
							if πE = πg.CheckLocal(πF, µmonth, "month"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µmonth == πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label3
							}
							goto Label4
							// line 1564: if month is None:
							πF.SetLineno(1564)
						Label3:
							// line 1565: month = self.month
							πF.SetLineno(1565)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßmonth, nil); πE != nil {
								continue
							}
							µmonth = πTemp001
							goto Label4
						Label4:
							if πE = πg.CheckLocal(πF, µday, "day"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µday == πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label5
							}
							goto Label6
							// line 1566: if day is None:
							πF.SetLineno(1566)
						Label5:
							// line 1567: day = self.day
							πF.SetLineno(1567)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßday, nil); πE != nil {
								continue
							}
							µday = πTemp001
							goto Label6
						Label6:
							if πE = πg.CheckLocal(πF, µhour, "hour"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µhour == πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label7
							}
							goto Label8
							// line 1568: if hour is None:
							πF.SetLineno(1568)
						Label7:
							// line 1569: hour = self.hour
							πF.SetLineno(1569)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßhour, nil); πE != nil {
								continue
							}
							µhour = πTemp001
							goto Label8
						Label8:
							if πE = πg.CheckLocal(πF, µminute, "minute"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µminute == πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label9
							}
							goto Label10
							// line 1570: if minute is None:
							πF.SetLineno(1570)
						Label9:
							// line 1571: minute = self.minute
							πF.SetLineno(1571)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßminute, nil); πE != nil {
								continue
							}
							µminute = πTemp001
							goto Label10
						Label10:
							if πE = πg.CheckLocal(πF, µsecond, "second"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µsecond == πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label11
							}
							goto Label12
							// line 1572: if second is None:
							πF.SetLineno(1572)
						Label11:
							// line 1573: second = self.second
							πF.SetLineno(1573)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßsecond, nil); πE != nil {
								continue
							}
							µsecond = πTemp001
							goto Label12
						Label12:
							if πE = πg.CheckLocal(πF, µmicrosecond, "microsecond"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µmicrosecond == πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label13
							}
							goto Label14
							// line 1574: if microsecond is None:
							πF.SetLineno(1574)
						Label13:
							// line 1575: microsecond = self.microsecond
							πF.SetLineno(1575)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßmicrosecond, nil); πE != nil {
								continue
							}
							µmicrosecond = πTemp001
							goto Label14
						Label14:
							if πE = πg.CheckLocal(πF, µtzinfo, "tzinfo"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µtzinfo == πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label15
							}
							goto Label16
							// line 1576: if tzinfo is True:
							πF.SetLineno(1576)
						Label15:
							// line 1577: tzinfo = self.tzinfo
							πF.SetLineno(1577)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßtzinfo, nil); πE != nil {
								continue
							}
							µtzinfo = πTemp001
							goto Label16
						Label16:
							// line 1578: return datetime.__new__(type(self),
							πF.SetLineno(1578)
							πTemp004 = πF.MakeArgs(9)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp005[0] = µself
							if πTemp001, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp004[0] = πTemp002
							if πE = πg.CheckLocal(πF, µyear, "year"); πE != nil {
								continue
							}
							πTemp004[1] = µyear
							if πE = πg.CheckLocal(πF, µmonth, "month"); πE != nil {
								continue
							}
							πTemp004[2] = µmonth
							if πE = πg.CheckLocal(πF, µday, "day"); πE != nil {
								continue
							}
							πTemp004[3] = µday
							if πE = πg.CheckLocal(πF, µhour, "hour"); πE != nil {
								continue
							}
							πTemp004[4] = µhour
							if πE = πg.CheckLocal(πF, µminute, "minute"); πE != nil {
								continue
							}
							πTemp004[5] = µminute
							if πE = πg.CheckLocal(πF, µsecond, "second"); πE != nil {
								continue
							}
							πTemp004[6] = µsecond
							if πE = πg.CheckLocal(πF, µmicrosecond, "microsecond"); πE != nil {
								continue
							}
							πTemp004[7] = µmicrosecond
							if πE = πg.CheckLocal(πF, µtzinfo, "tzinfo"); πE != nil {
								continue
							}
							πTemp004[8] = µtzinfo
							if πTemp001, πE = πg.ResolveGlobal(πF, ßdatetime); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ß__new__, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
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
					if πE = πClass.SetItem(πF, ßreplace.ToObject(), πTemp020); πE != nil {
						continue
					}
					// line 1561: """Return a new datetime with new values for the specified fields."""
					πF.SetLineno(1561)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp021}, πg.NewStr("Return a new datetime with new values for the specified fields.").ToObject()); πE != nil {
						continue
					}
					if πTemp022, πE = πg.ResolveClass(πF, πClass, nil, ßreplace); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp022, ß__doc__, πTemp021); πE != nil {
						continue
					}
					// line 1582: def astimezone(self, tz):
					πF.SetLineno(1582)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "tz", Def: nil}
					πTemp021 = πg.NewFunction(πg.NewCode("astimezone", "/usr/lib/python2.7/datetime.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µtz *πg.Object = πArgs[1]; _ = µtz
						var µmytz *πg.Object = πg.UnboundLocal; _ = µmytz
						var µmyoffset *πg.Object = πg.UnboundLocal; _ = µmyoffset
						var µutc *πg.Object = πg.UnboundLocal; _ = µutc
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
						var πTemp006 πg.KWArgs
						_ = πTemp006
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µtz, "tz"); πE != nil {
								continue
							}
							πTemp002[0] = µtz
							if πTemp003, πE = πg.ResolveGlobal(πF, ßtzinfo); πE != nil {
								continue
							}
							πTemp002[1] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp005, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp005).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							goto Label2
							// line 1583: if not isinstance(tz, tzinfo):
							πF.SetLineno(1583)
						Label1:
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("tz argument must be an instance of tzinfo").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 1584: raise TypeError("tz argument must be an instance of tzinfo")
							πF.SetLineno(1584)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label2
						Label2:
							// line 1586: mytz = self.tzinfo
							πF.SetLineno(1586)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßtzinfo, nil); πE != nil {
								continue
							}
							µmytz = πTemp001
							if πE = πg.CheckLocal(πF, µmytz, "mytz"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µmytz == πTemp003).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label3
							}
							goto Label4
							// line 1587: if mytz is None:
							πF.SetLineno(1587)
						Label3:
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("astimezone() requires an aware datetime").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 1588: raise ValueError("astimezone() requires an aware datetime")
							πF.SetLineno(1588)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label4
						Label4:
							if πE = πg.CheckLocal(πF, µtz, "tz"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µmytz, "mytz"); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µtz == µmytz).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label5
							}
							goto Label6
							// line 1590: if tz is mytz:
							πF.SetLineno(1590)
						Label5:
							// line 1591: return self
							πF.SetLineno(1591)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πR = µself
							continue
							goto Label6
						Label6:
							// line 1594: myoffset = self.utcoffset()
							πF.SetLineno(1594)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßutcoffset, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µmyoffset = πTemp003
							if πE = πg.CheckLocal(πF, µmyoffset, "myoffset"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µmyoffset == πTemp003).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label7
							}
							goto Label8
							// line 1595: if myoffset is None:
							πF.SetLineno(1595)
						Label7:
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("astimezone() requires an aware datetime").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 1596: raise ValueError("astimezone() requires an aware datetime")
							πF.SetLineno(1596)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label8
						Label8:
							// line 1597: utc = (self - myoffset).replace(tzinfo=tz)
							πF.SetLineno(1597)
							if πE = πg.CheckLocal(πF, µtz, "tz"); πE != nil {
								continue
							}
							πTemp006 = πg.KWArgs{
								{"tzinfo", µtz},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µmyoffset, "myoffset"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Sub(πF, µself, µmyoffset); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßreplace, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, nil, πTemp006); πE != nil {
								continue
							}
							µutc = πTemp001
							// line 1600: return tz.fromutc(utc)
							πF.SetLineno(1600)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µutc, "utc"); πE != nil {
								continue
							}
							πTemp002[0] = µutc
							if πE = πg.CheckLocal(πF, µtz, "tz"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µtz, ßfromutc, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
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
					if πE = πClass.SetItem(πF, ßastimezone.ToObject(), πTemp021); πE != nil {
						continue
					}
					// line 1604: def ctime(self):
					πF.SetLineno(1604)
					πTemp005 = make([]πg.Param, 1)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp022 = πg.NewFunction(πg.NewCode("ctime", "/usr/lib/python2.7/datetime.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µweekday *πg.Object = πg.UnboundLocal; _ = µweekday
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
						var πTemp006 []*πg.Object
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
							// line 1605: "Return ctime() style string."
							πF.SetLineno(1605)
							// line 1606: weekday = self.toordinal() % 7 or 7
							πF.SetLineno(1606)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßtoordinal, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Mod(πF, πTemp005, πg.NewInt(7).ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp003
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							πTemp001 = πg.NewInt(7).ToObject()
						Label1:
							µweekday = πTemp001
							// line 1607: return "%s %s %2d %02d:%02d:%02d %04d" % (
							πF.SetLineno(1607)
							πTemp006 = make([]*πg.Object, 7)
							if πE = πg.CheckLocal(πF, µweekday, "weekday"); πE != nil {
								continue
							}
							πTemp004 = µweekday
							if πTemp007, πE = πg.ResolveGlobal(πF, ß_DAYNAMES); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, πTemp007, πTemp004); πE != nil {
								continue
							}
							πTemp006[0] = πTemp005
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ß_month, nil); πE != nil {
								continue
							}
							πTemp004 = πTemp005
							if πTemp007, πE = πg.ResolveGlobal(πF, ß_MONTHNAMES); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, πTemp007, πTemp004); πE != nil {
								continue
							}
							πTemp006[1] = πTemp005
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß_day, nil); πE != nil {
								continue
							}
							πTemp006[2] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß_hour, nil); πE != nil {
								continue
							}
							πTemp006[3] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß_minute, nil); πE != nil {
								continue
							}
							πTemp006[4] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß_second, nil); πE != nil {
								continue
							}
							πTemp006[5] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß_year, nil); πE != nil {
								continue
							}
							πTemp006[6] = πTemp004
							πTemp003 = πg.NewTuple(πTemp006...).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("%s %s %2d %02d:%02d:%02d %04d").ToObject(), πTemp003); πE != nil {
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
					if πE = πClass.SetItem(πF, ßctime.ToObject(), πTemp022); πE != nil {
						continue
					}
					// line 1605: "Return ctime() style string."
					πF.SetLineno(1605)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp023}, πg.NewStr("Return ctime() style string.").ToObject()); πE != nil {
						continue
					}
					if πTemp024, πE = πg.ResolveClass(πF, πClass, nil, ßctime); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp024, ß__doc__, πTemp023); πE != nil {
						continue
					}
					// line 1614: def isoformat(self, sep='T'):
					πF.SetLineno(1614)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "sep", Def: ßT.ToObject()}
					πTemp023 = πg.NewFunction(πg.NewCode("isoformat", "/usr/lib/python2.7/datetime.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µsep *πg.Object = πArgs[1]; _ = µsep
						var µs *πg.Object = πg.UnboundLocal; _ = µs
						var µoff *πg.Object = πg.UnboundLocal; _ = µoff
						var µsign *πg.Object = πg.UnboundLocal; _ = µsign
						var µhh *πg.Object = πg.UnboundLocal; _ = µhh
						var µmm *πg.Object = πg.UnboundLocal; _ = µmm
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
						var πTemp007 []*πg.Object
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
							// line 1615: """Return the time formatted according to ISO.
							πF.SetLineno(1615)
							// line 1626: s = ("%04d-%02d-%02d%c" % (self._year, self._month, self._day, sep) +
							πF.SetLineno(1626)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß_year, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ß_month, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ß_day, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µsep, "sep"); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple4(πTemp004, πTemp005, πTemp006, µsep).ToObject()
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("%04d-%02d-%02d%c").ToObject(), πTemp003); πE != nil {
								continue
							}
							πTemp007 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_hour, nil); πE != nil {
								continue
							}
							πTemp007[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_minute, nil); πE != nil {
								continue
							}
							πTemp007[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_second, nil); πE != nil {
								continue
							}
							πTemp007[2] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_microsecond, nil); πE != nil {
								continue
							}
							πTemp007[3] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ß_format_time); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							if πTemp001, πE = πg.Add(πF, πTemp002, πTemp004); πE != nil {
								continue
							}
							µs = πTemp001
							// line 1629: off = self._utcoffset()
							πF.SetLineno(1629)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_utcoffset, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µoff = πTemp002
							if πE = πg.CheckLocal(πF, µoff, "off"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µoff != πTemp002).ToObject()
							if πTemp008, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label1
							}
							goto Label2
							// line 1630: if off is not None:
							πF.SetLineno(1630)
						Label1:
							if πE = πg.CheckLocal(πF, µoff, "off"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.LT(πF, µoff, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp008, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label3
							}
							goto Label4
							// line 1631: if off < 0:
							πF.SetLineno(1631)
						Label3:
							// line 1632: sign = "-"
							πF.SetLineno(1632)
							µsign = πg.NewStr("-").ToObject()
							// line 1633: off = -off
							πF.SetLineno(1633)
							if πE = πg.CheckLocal(πF, µoff, "off"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Neg(πF, µoff); πE != nil {
								continue
							}
							µoff = πTemp001
							goto Label5
						Label4:
							// line 1635: sign = "+"
							πF.SetLineno(1635)
							µsign = πg.NewStr("+").ToObject()
							goto Label5
						Label5:
							// line 1636: hh, mm = divmod(off, 60)
							πF.SetLineno(1636)
							πTemp007 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µoff, "off"); πE != nil {
								continue
							}
							πTemp007[0] = µoff
							πTemp007[1] = πg.NewInt(60).ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßdivmod); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp003}}}, πTemp002); πE != nil {
								continue
							}
							µhh = πTemp001
							µmm = πTemp003
							// line 1637: s += "%s%02d:%02d" % (sign, hh, mm)
							πF.SetLineno(1637)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µsign, "sign"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µhh, "hh"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µmm, "mm"); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple3(µsign, µhh, µmm).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("%s%02d:%02d").ToObject(), πTemp002); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IAdd(πF, µs, πTemp001); πE != nil {
								continue
							}
							µs = πTemp002
							goto Label2
						Label2:
							// line 1638: return s
							πF.SetLineno(1638)
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
					if πE = πClass.SetItem(πF, ßisoformat.ToObject(), πTemp023); πE != nil {
						continue
					}
					// line 1615: """Return the time formatted according to ISO.
					πF.SetLineno(1615)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp024}, πg.NewStr("Return the time formatted according to ISO.\n\n        This is 'YYYY-MM-DD HH:MM:SS.mmmmmm', or 'YYYY-MM-DD HH:MM:SS' if\n        self.microsecond == 0.\n\n        If self.tzinfo is not None, the UTC offset is also attached, giving\n        'YYYY-MM-DD HH:MM:SS.mmmmmm+HH:MM' or 'YYYY-MM-DD HH:MM:SS+HH:MM'.\n\n        Optional argument sep specifies the separator between date and\n        time, default 'T'.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp025, πE = πg.ResolveClass(πF, πClass, nil, ßisoformat); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp025, ß__doc__, πTemp024); πE != nil {
						continue
					}
					// line 1640: def __repr__(self):
					πF.SetLineno(1640)
					πTemp005 = make([]πg.Param, 1)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp024 = πg.NewFunction(πg.NewCode("__repr__", "/usr/lib/python2.7/datetime.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µL *πg.Object = πg.UnboundLocal; _ = µL
						var µs *πg.Object = πg.UnboundLocal; _ = µs
						var µmodule *πg.Object = πg.UnboundLocal; _ = µmodule
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
							// line 1641: """Convert to formal string, for repr()."""
							πF.SetLineno(1641)
							// line 1642: L = [self._year, self._month, self._day,  # These are never zero
							πF.SetLineno(1642)
							πTemp001 = make([]*πg.Object, 7)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_year, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_month, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_day, nil); πE != nil {
								continue
							}
							πTemp001[2] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_hour, nil); πE != nil {
								continue
							}
							πTemp001[3] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_minute, nil); πE != nil {
								continue
							}
							πTemp001[4] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_second, nil); πE != nil {
								continue
							}
							πTemp001[5] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_microsecond, nil); πE != nil {
								continue
							}
							πTemp001[6] = πTemp002
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µL = πTemp002
							if πTemp004, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp003 = πTemp004
							if πE = πg.CheckLocal(πF, µL, "L"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µL, πTemp003); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, πTemp004, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							goto Label2
							// line 1644: if L[-1] == 0:
							πF.SetLineno(1644)
						Label1:
							// line 1645: del L[-1]
							πF.SetLineno(1645)
							if πE = πg.CheckLocal(πF, µL, "L"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp003
							if πE = πg.DelItem(πF, µL, πTemp002); πE != nil {
								continue
							}
							goto Label2
						Label2:
							if πTemp004, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp003 = πTemp004
							if πE = πg.CheckLocal(πF, µL, "L"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µL, πTemp003); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, πTemp004, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label3
							}
							goto Label4
							// line 1646: if L[-1] == 0:
							πF.SetLineno(1646)
						Label3:
							// line 1647: del L[-1]
							πF.SetLineno(1647)
							if πE = πg.CheckLocal(πF, µL, "L"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp003
							if πE = πg.DelItem(πF, µL, πTemp002); πE != nil {
								continue
							}
							goto Label4
						Label4:
							// line 1648: s = ", ".join(map(str, L))
							πF.SetLineno(1648)
							πTemp001 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
								continue
							}
							πTemp006[0] = πTemp002
							if πE = πg.CheckLocal(πF, µL, "L"); πE != nil {
								continue
							}
							πTemp006[1] = µL
							if πTemp002, πE = πg.ResolveGlobal(πF, ßmap); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp001[0] = πTemp003
							if πTemp002, πE = πg.GetAttr(πF, πg.NewStr(", ").ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µs = πTemp003
							// line 1649: module = "datetime." if self.__class__ is datetime else ""
							πF.SetLineno(1649)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß__class__, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.ResolveGlobal(πF, ßdatetime); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(πTemp004 == πTemp007).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if !πTemp005 {
								goto Label5
							}
							πTemp002 = πg.NewStr("datetime.").ToObject()
							goto Label6
						Label5:
							πTemp002 = ß.ToObject()
						Label6:
							µmodule = πTemp002
							// line 1650: s = "%s(%s)" % (module + self.__class__.__name__, s)
							πF.SetLineno(1650)
							if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ß__class__, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp007, ß__name__, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Add(πF, µmodule, πTemp008); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple2(πTemp004, µs).ToObject()
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("%s(%s)").ToObject(), πTemp003); πE != nil {
								continue
							}
							µs = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_tzinfo, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp003 != πTemp004).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label7
							}
							goto Label8
							// line 1651: if self._tzinfo is not None:
							πF.SetLineno(1651)
						Label7:
							// line 1652: assert s[-1:] == ")"
							πF.SetLineno(1652)
							if πTemp004, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
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
							if πTemp002, πE = πg.Eq(πF, πTemp004, πg.NewStr(")").ToObject()); πE != nil {
								continue
							}
							if πE = πg.Assert(πF, πTemp002, nil); πE != nil {
								continue
							}
							// line 1653: s = s[:-1] + ", tzinfo=%r" % self._tzinfo + ")"
							πF.SetLineno(1653)
							if πTemp007, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πTemp007, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, µs, πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µself, ß_tzinfo, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Mod(πF, πg.NewStr(", tzinfo=%r").ToObject(), πTemp008); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, πTemp007, πTemp004); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp003, πg.NewStr(")").ToObject()); πE != nil {
								continue
							}
							µs = πTemp002
							goto Label8
						Label8:
							// line 1654: return s
							πF.SetLineno(1654)
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
					if πE = πClass.SetItem(πF, ß__repr__.ToObject(), πTemp024); πE != nil {
						continue
					}
					// line 1641: """Convert to formal string, for repr()."""
					πF.SetLineno(1641)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp025}, πg.NewStr("Convert to formal string, for repr().").ToObject()); πE != nil {
						continue
					}
					if πTemp026, πE = πg.ResolveClass(πF, πClass, nil, ß__repr__); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp026, ß__doc__, πTemp025); πE != nil {
						continue
					}
					// line 1656: def __str__(self):
					πF.SetLineno(1656)
					πTemp005 = make([]πg.Param, 1)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp025 = πg.NewFunction(πg.NewCode("__str__", "/usr/lib/python2.7/datetime.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πTemp001 πg.KWArgs
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
							// line 1657: "Convert to string, for str()."
							πF.SetLineno(1657)
							// line 1658: return self.isoformat(sep=' ')
							πF.SetLineno(1658)
							πTemp001 = πg.KWArgs{
								{"sep", πg.NewStr(" ").ToObject()},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßisoformat, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, πTemp001); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__str__.ToObject(), πTemp025); πE != nil {
						continue
					}
					// line 1657: "Convert to string, for str()."
					πF.SetLineno(1657)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp026}, πg.NewStr("Convert to string, for str().").ToObject()); πE != nil {
						continue
					}
					if πTemp027, πE = πg.ResolveClass(πF, πClass, nil, ß__str__); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp027, ß__doc__, πTemp026); πE != nil {
						continue
					}
					// line 1670: def utcoffset(self):
					πF.SetLineno(1670)
					πTemp005 = make([]πg.Param, 1)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp026 = πg.NewFunction(πg.NewCode("utcoffset", "/usr/lib/python2.7/datetime.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µoffset *πg.Object = πg.UnboundLocal; _ = µoffset
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 1671: """Return the timezone offset in minutes east of UTC (negative west of
							πF.SetLineno(1671)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_tzinfo, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp002 == πTemp003).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 1673: if self._tzinfo is None:
							πF.SetLineno(1673)
						Label1:
							// line 1674: return None
							πF.SetLineno(1674)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label2
						Label2:
							// line 1675: offset = self._tzinfo.utcoffset(self)
							πF.SetLineno(1675)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp005[0] = µself
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_tzinfo, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßutcoffset, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µoffset = πTemp001
							// line 1676: offset = _check_utc_offset("utcoffset", offset)
							πF.SetLineno(1676)
							πTemp005 = πF.MakeArgs(2)
							πTemp005[0] = ßutcoffset.ToObject()
							if πE = πg.CheckLocal(πF, µoffset, "offset"); πE != nil {
								continue
							}
							πTemp005[1] = µoffset
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_check_utc_offset); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µoffset = πTemp002
							if πE = πg.CheckLocal(πF, µoffset, "offset"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µoffset != πTemp002).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label3
							}
							goto Label4
							// line 1677: if offset is not None:
							πF.SetLineno(1677)
						Label3:
							// line 1678: offset = timedelta._create(0, offset * 60, 0, True)
							πF.SetLineno(1678)
							πTemp005 = πF.MakeArgs(4)
							πTemp005[0] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µoffset, "offset"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mul(πF, µoffset, πg.NewInt(60).ToObject()); πE != nil {
								continue
							}
							πTemp005[1] = πTemp001
							πTemp005[2] = πg.NewInt(0).ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp005[3] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßtimedelta); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ß_create, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µoffset = πTemp001
							goto Label4
						Label4:
							// line 1679: return offset
							πF.SetLineno(1679)
							if πE = πg.CheckLocal(πF, µoffset, "offset"); πE != nil {
								continue
							}
							πR = µoffset
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßutcoffset.ToObject(), πTemp026); πE != nil {
						continue
					}
					// line 1671: """Return the timezone offset in minutes east of UTC (negative west of
					πF.SetLineno(1671)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp027}, πg.NewStr("Return the timezone offset in minutes east of UTC (negative west of\n        UTC).").ToObject()); πE != nil {
						continue
					}
					if πTemp028, πE = πg.ResolveClass(πF, πClass, nil, ßutcoffset); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp028, ß__doc__, πTemp027); πE != nil {
						continue
					}
					// line 1682: def _utcoffset(self):
					πF.SetLineno(1682)
					πTemp005 = make([]πg.Param, 1)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp027 = πg.NewFunction(πg.NewCode("_utcoffset", "/usr/lib/python2.7/datetime.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µoffset *πg.Object = πg.UnboundLocal; _ = µoffset
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_tzinfo, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp002 == πTemp003).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 1683: if self._tzinfo is None:
							πF.SetLineno(1683)
						Label1:
							// line 1684: return None
							πF.SetLineno(1684)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label2
						Label2:
							// line 1685: offset = self._tzinfo.utcoffset(self)
							πF.SetLineno(1685)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp005[0] = µself
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_tzinfo, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßutcoffset, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µoffset = πTemp001
							// line 1686: offset = _check_utc_offset("utcoffset", offset)
							πF.SetLineno(1686)
							πTemp005 = πF.MakeArgs(2)
							πTemp005[0] = ßutcoffset.ToObject()
							if πE = πg.CheckLocal(πF, µoffset, "offset"); πE != nil {
								continue
							}
							πTemp005[1] = µoffset
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_check_utc_offset); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µoffset = πTemp002
							// line 1687: return offset
							πF.SetLineno(1687)
							if πE = πg.CheckLocal(πF, µoffset, "offset"); πE != nil {
								continue
							}
							πR = µoffset
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß_utcoffset.ToObject(), πTemp027); πE != nil {
						continue
					}
					// line 1689: def tzname(self):
					πF.SetLineno(1689)
					πTemp005 = make([]πg.Param, 1)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp028 = πg.NewFunction(πg.NewCode("tzname", "/usr/lib/python2.7/datetime.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µname *πg.Object = πg.UnboundLocal; _ = µname
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 1690: """Return the timezone name.
							πF.SetLineno(1690)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_tzinfo, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp002 == πTemp003).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 1696: if self._tzinfo is None:
							πF.SetLineno(1696)
						Label1:
							// line 1697: return None
							πF.SetLineno(1697)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label2
						Label2:
							// line 1698: name = self._tzinfo.tzname(self)
							πF.SetLineno(1698)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp005[0] = µself
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_tzinfo, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßtzname, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µname = πTemp001
							// line 1699: _check_tzname(name)
							πF.SetLineno(1699)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp005[0] = µname
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_check_tzname); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 1700: return name
							πF.SetLineno(1700)
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
					if πE = πClass.SetItem(πF, ßtzname.ToObject(), πTemp028); πE != nil {
						continue
					}
					// line 1690: """Return the timezone name.
					πF.SetLineno(1690)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp029}, πg.NewStr("Return the timezone name.\n\n        Note that the name is 100% informational -- there's no requirement that\n        it mean anything in particular. For example, \"GMT\", \"UTC\", \"-500\",\n        \"-5:00\", \"EDT\", \"US/Eastern\", \"America/New York\" are all valid replies.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp030, πE = πg.ResolveClass(πF, πClass, nil, ßtzname); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp030, ß__doc__, πTemp029); πE != nil {
						continue
					}
					// line 1702: def dst(self):
					πF.SetLineno(1702)
					πTemp005 = make([]πg.Param, 1)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp029 = πg.NewFunction(πg.NewCode("dst", "/usr/lib/python2.7/datetime.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µoffset *πg.Object = πg.UnboundLocal; _ = µoffset
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 1703: """Return 0 if DST is not in effect, or the DST offset (in minutes
							πF.SetLineno(1703)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_tzinfo, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp002 == πTemp003).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 1711: if self._tzinfo is None:
							πF.SetLineno(1711)
						Label1:
							// line 1712: return None
							πF.SetLineno(1712)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label2
						Label2:
							// line 1713: offset = self._tzinfo.dst(self)
							πF.SetLineno(1713)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp005[0] = µself
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_tzinfo, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßdst, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µoffset = πTemp001
							// line 1714: offset = _check_utc_offset("dst", offset)
							πF.SetLineno(1714)
							πTemp005 = πF.MakeArgs(2)
							πTemp005[0] = ßdst.ToObject()
							if πE = πg.CheckLocal(πF, µoffset, "offset"); πE != nil {
								continue
							}
							πTemp005[1] = µoffset
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_check_utc_offset); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µoffset = πTemp002
							if πE = πg.CheckLocal(πF, µoffset, "offset"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µoffset != πTemp002).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label3
							}
							goto Label4
							// line 1715: if offset is not None:
							πF.SetLineno(1715)
						Label3:
							// line 1716: offset = timedelta._create(0, offset * 60, 0, True)
							πF.SetLineno(1716)
							πTemp005 = πF.MakeArgs(4)
							πTemp005[0] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µoffset, "offset"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mul(πF, µoffset, πg.NewInt(60).ToObject()); πE != nil {
								continue
							}
							πTemp005[1] = πTemp001
							πTemp005[2] = πg.NewInt(0).ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp005[3] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßtimedelta); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ß_create, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µoffset = πTemp001
							goto Label4
						Label4:
							// line 1717: return offset
							πF.SetLineno(1717)
							if πE = πg.CheckLocal(πF, µoffset, "offset"); πE != nil {
								continue
							}
							πR = µoffset
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdst.ToObject(), πTemp029); πE != nil {
						continue
					}
					// line 1703: """Return 0 if DST is not in effect, or the DST offset (in minutes
					πF.SetLineno(1703)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp030}, πg.NewStr("Return 0 if DST is not in effect, or the DST offset (in minutes\n        eastward) if DST is in effect.\n\n        This is purely informational; the DST offset has already been added to\n        the UTC offset returned by utcoffset() if applicable, so there's no\n        need to consult dst() unless you're interested in displaying the DST\n        info.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp031, πE = πg.ResolveClass(πF, πClass, nil, ßdst); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp031, ß__doc__, πTemp030); πE != nil {
						continue
					}
					// line 1720: def _dst(self):
					πF.SetLineno(1720)
					πTemp005 = make([]πg.Param, 1)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp030 = πg.NewFunction(πg.NewCode("_dst", "/usr/lib/python2.7/datetime.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µoffset *πg.Object = πg.UnboundLocal; _ = µoffset
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_tzinfo, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp002 == πTemp003).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 1721: if self._tzinfo is None:
							πF.SetLineno(1721)
						Label1:
							// line 1722: return None
							πF.SetLineno(1722)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label2
						Label2:
							// line 1723: offset = self._tzinfo.dst(self)
							πF.SetLineno(1723)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp005[0] = µself
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_tzinfo, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßdst, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µoffset = πTemp001
							// line 1724: offset = _check_utc_offset("dst", offset)
							πF.SetLineno(1724)
							πTemp005 = πF.MakeArgs(2)
							πTemp005[0] = ßdst.ToObject()
							if πE = πg.CheckLocal(πF, µoffset, "offset"); πE != nil {
								continue
							}
							πTemp005[1] = µoffset
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_check_utc_offset); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µoffset = πTemp002
							// line 1725: return offset
							πF.SetLineno(1725)
							if πE = πg.CheckLocal(πF, µoffset, "offset"); πE != nil {
								continue
							}
							πR = µoffset
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß_dst.ToObject(), πTemp030); πE != nil {
						continue
					}
					// line 1729: def __eq__(self, other):
					πF.SetLineno(1729)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "other", Def: nil}
					πTemp031 = πg.NewFunction(πg.NewCode("__eq__", "/usr/lib/python2.7/datetime.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							if πTemp002, πE = πg.ResolveGlobal(πF, ßdatetime); πE != nil {
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
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							πTemp001[1] = ßtimetuple.ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp002 = πTemp005
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp004 {
								goto Label2
							}
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							if πTemp005, πE = πg.ResolveGlobal(πF, ßdate); πE != nil {
								continue
							}
							πTemp001[1] = πTemp005
							if πTemp005, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp007, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp007).ToObject()
							πTemp002 = πTemp003
						Label2:
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label3
							}
							goto Label4
							// line 1730: if isinstance(other, datetime):
							πF.SetLineno(1730)
						Label1:
							// line 1731: return self._cmp(other) == 0
							πF.SetLineno(1731)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_cmp, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.Eq(πF, πTemp005, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							πR = πTemp002
							continue
							goto Label5
							// line 1732: elif hasattr(other, "timetuple") and not isinstance(other, date):
							πF.SetLineno(1732)
						Label3:
							// line 1733: return NotImplemented
							πF.SetLineno(1733)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNotImplemented); πE != nil {
								continue
							}
							πR = πTemp002
							continue
							goto Label5
						Label4:
							// line 1735: return False
							πF.SetLineno(1735)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πR = πTemp002
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
					if πE = πClass.SetItem(πF, ß__eq__.ToObject(), πTemp031); πE != nil {
						continue
					}
					// line 1737: def __ne__(self, other):
					πF.SetLineno(1737)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "other", Def: nil}
					πTemp032 = πg.NewFunction(πg.NewCode("__ne__", "/usr/lib/python2.7/datetime.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							if πTemp002, πE = πg.ResolveGlobal(πF, ßdatetime); πE != nil {
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
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							πTemp001[1] = ßtimetuple.ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp002 = πTemp005
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp004 {
								goto Label2
							}
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							if πTemp005, πE = πg.ResolveGlobal(πF, ßdate); πE != nil {
								continue
							}
							πTemp001[1] = πTemp005
							if πTemp005, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp007, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp007).ToObject()
							πTemp002 = πTemp003
						Label2:
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label3
							}
							goto Label4
							// line 1738: if isinstance(other, datetime):
							πF.SetLineno(1738)
						Label1:
							// line 1739: return self._cmp(other) != 0
							πF.SetLineno(1739)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_cmp, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.NE(πF, πTemp005, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							πR = πTemp002
							continue
							goto Label5
							// line 1740: elif hasattr(other, "timetuple") and not isinstance(other, date):
							πF.SetLineno(1740)
						Label3:
							// line 1741: return NotImplemented
							πF.SetLineno(1741)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNotImplemented); πE != nil {
								continue
							}
							πR = πTemp002
							continue
							goto Label5
						Label4:
							// line 1743: return True
							πF.SetLineno(1743)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πR = πTemp002
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
					if πE = πClass.SetItem(πF, ß__ne__.ToObject(), πTemp032); πE != nil {
						continue
					}
					// line 1745: def __le__(self, other):
					πF.SetLineno(1745)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "other", Def: nil}
					πTemp033 = πg.NewFunction(πg.NewCode("__le__", "/usr/lib/python2.7/datetime.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							if πTemp002, πE = πg.ResolveGlobal(πF, ßdatetime); πE != nil {
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
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							πTemp001[1] = ßtimetuple.ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp002 = πTemp005
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp004 {
								goto Label2
							}
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							if πTemp005, πE = πg.ResolveGlobal(πF, ßdate); πE != nil {
								continue
							}
							πTemp001[1] = πTemp005
							if πTemp005, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp007, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp007).ToObject()
							πTemp002 = πTemp003
						Label2:
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label3
							}
							goto Label4
							// line 1746: if isinstance(other, datetime):
							πF.SetLineno(1746)
						Label1:
							// line 1747: return self._cmp(other) <= 0
							πF.SetLineno(1747)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_cmp, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.LE(πF, πTemp005, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							πR = πTemp002
							continue
							goto Label5
							// line 1748: elif hasattr(other, "timetuple") and not isinstance(other, date):
							πF.SetLineno(1748)
						Label3:
							// line 1749: return NotImplemented
							πF.SetLineno(1749)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNotImplemented); πE != nil {
								continue
							}
							πR = πTemp002
							continue
							goto Label5
						Label4:
							// line 1751: _cmperror(self, other)
							πF.SetLineno(1751)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[1] = µother
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_cmperror); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
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
					if πE = πClass.SetItem(πF, ß__le__.ToObject(), πTemp033); πE != nil {
						continue
					}
					// line 1753: def __lt__(self, other):
					πF.SetLineno(1753)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "other", Def: nil}
					πTemp034 = πg.NewFunction(πg.NewCode("__lt__", "/usr/lib/python2.7/datetime.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							if πTemp002, πE = πg.ResolveGlobal(πF, ßdatetime); πE != nil {
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
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							πTemp001[1] = ßtimetuple.ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp002 = πTemp005
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp004 {
								goto Label2
							}
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							if πTemp005, πE = πg.ResolveGlobal(πF, ßdate); πE != nil {
								continue
							}
							πTemp001[1] = πTemp005
							if πTemp005, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp007, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp007).ToObject()
							πTemp002 = πTemp003
						Label2:
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label3
							}
							goto Label4
							// line 1754: if isinstance(other, datetime):
							πF.SetLineno(1754)
						Label1:
							// line 1755: return self._cmp(other) < 0
							πF.SetLineno(1755)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_cmp, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.LT(πF, πTemp005, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							πR = πTemp002
							continue
							goto Label5
							// line 1756: elif hasattr(other, "timetuple") and not isinstance(other, date):
							πF.SetLineno(1756)
						Label3:
							// line 1757: return NotImplemented
							πF.SetLineno(1757)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNotImplemented); πE != nil {
								continue
							}
							πR = πTemp002
							continue
							goto Label5
						Label4:
							// line 1759: _cmperror(self, other)
							πF.SetLineno(1759)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[1] = µother
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_cmperror); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
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
					if πE = πClass.SetItem(πF, ß__lt__.ToObject(), πTemp034); πE != nil {
						continue
					}
					// line 1761: def __ge__(self, other):
					πF.SetLineno(1761)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "other", Def: nil}
					πTemp035 = πg.NewFunction(πg.NewCode("__ge__", "/usr/lib/python2.7/datetime.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							if πTemp002, πE = πg.ResolveGlobal(πF, ßdatetime); πE != nil {
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
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							πTemp001[1] = ßtimetuple.ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp002 = πTemp005
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp004 {
								goto Label2
							}
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							if πTemp005, πE = πg.ResolveGlobal(πF, ßdate); πE != nil {
								continue
							}
							πTemp001[1] = πTemp005
							if πTemp005, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp007, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp007).ToObject()
							πTemp002 = πTemp003
						Label2:
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label3
							}
							goto Label4
							// line 1762: if isinstance(other, datetime):
							πF.SetLineno(1762)
						Label1:
							// line 1763: return self._cmp(other) >= 0
							πF.SetLineno(1763)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_cmp, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.GE(πF, πTemp005, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							πR = πTemp002
							continue
							goto Label5
							// line 1764: elif hasattr(other, "timetuple") and not isinstance(other, date):
							πF.SetLineno(1764)
						Label3:
							// line 1765: return NotImplemented
							πF.SetLineno(1765)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNotImplemented); πE != nil {
								continue
							}
							πR = πTemp002
							continue
							goto Label5
						Label4:
							// line 1767: _cmperror(self, other)
							πF.SetLineno(1767)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[1] = µother
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_cmperror); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
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
					if πE = πClass.SetItem(πF, ß__ge__.ToObject(), πTemp035); πE != nil {
						continue
					}
					// line 1769: def __gt__(self, other):
					πF.SetLineno(1769)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "other", Def: nil}
					πTemp036 = πg.NewFunction(πg.NewCode("__gt__", "/usr/lib/python2.7/datetime.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							if πTemp002, πE = πg.ResolveGlobal(πF, ßdatetime); πE != nil {
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
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							πTemp001[1] = ßtimetuple.ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp002 = πTemp005
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp004 {
								goto Label2
							}
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							if πTemp005, πE = πg.ResolveGlobal(πF, ßdate); πE != nil {
								continue
							}
							πTemp001[1] = πTemp005
							if πTemp005, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp007, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp007).ToObject()
							πTemp002 = πTemp003
						Label2:
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label3
							}
							goto Label4
							// line 1770: if isinstance(other, datetime):
							πF.SetLineno(1770)
						Label1:
							// line 1771: return self._cmp(other) > 0
							πF.SetLineno(1771)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_cmp, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.GT(πF, πTemp005, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							πR = πTemp002
							continue
							goto Label5
							// line 1772: elif hasattr(other, "timetuple") and not isinstance(other, date):
							πF.SetLineno(1772)
						Label3:
							// line 1773: return NotImplemented
							πF.SetLineno(1773)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNotImplemented); πE != nil {
								continue
							}
							πR = πTemp002
							continue
							goto Label5
						Label4:
							// line 1775: _cmperror(self, other)
							πF.SetLineno(1775)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[1] = µother
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_cmperror); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
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
					if πE = πClass.SetItem(πF, ß__gt__.ToObject(), πTemp036); πE != nil {
						continue
					}
					// line 1777: def _cmp(self, other):
					πF.SetLineno(1777)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "other", Def: nil}
					πTemp037 = πg.NewFunction(πg.NewCode("_cmp", "/usr/lib/python2.7/datetime.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
						var µmytz *πg.Object = πg.UnboundLocal; _ = µmytz
						var µottz *πg.Object = πg.UnboundLocal; _ = µottz
						var µmyoff *πg.Object = πg.UnboundLocal; _ = µmyoff
						var µotoff *πg.Object = πg.UnboundLocal; _ = µotoff
						var µbase_compare *πg.Object = πg.UnboundLocal; _ = µbase_compare
						var µdiff *πg.Object = πg.UnboundLocal; _ = µdiff
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 1778: assert isinstance(other, datetime)
							πF.SetLineno(1778)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							if πTemp002, πE = πg.ResolveGlobal(πF, ßdatetime); πE != nil {
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
							if πE = πg.Assert(πF, πTemp003, nil); πE != nil {
								continue
							}
							// line 1779: mytz = self._tzinfo
							πF.SetLineno(1779)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_tzinfo, nil); πE != nil {
								continue
							}
							µmytz = πTemp002
							// line 1780: ottz = other._tzinfo
							πF.SetLineno(1780)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µother, ß_tzinfo, nil); πE != nil {
								continue
							}
							µottz = πTemp002
							// line 1781: myoff = otoff = None
							πF.SetLineno(1781)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							µmyoff = πTemp002
							µotoff = πTemp002
							if πE = πg.CheckLocal(πF, µmytz, "mytz"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µottz, "ottz"); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(µmytz == µottz).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 1783: if mytz is ottz:
							πF.SetLineno(1783)
						Label1:
							// line 1784: base_compare = True
							πF.SetLineno(1784)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							µbase_compare = πTemp002
							goto Label3
						Label2:
							if πE = πg.CheckLocal(πF, µmytz, "mytz"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(µmytz != πTemp003).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label4
							}
							goto Label5
							// line 1786: if mytz is not None:
							πF.SetLineno(1786)
						Label4:
							// line 1787: myoff = self._utcoffset()
							πF.SetLineno(1787)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_utcoffset, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µmyoff = πTemp003
							goto Label5
						Label5:
							if πE = πg.CheckLocal(πF, µottz, "ottz"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(µottz != πTemp003).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label6
							}
							goto Label7
							// line 1788: if ottz is not None:
							πF.SetLineno(1788)
						Label6:
							// line 1789: otoff = other._utcoffset()
							πF.SetLineno(1789)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µother, ß_utcoffset, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µotoff = πTemp003
							goto Label7
						Label7:
							// line 1790: base_compare = myoff == otoff
							πF.SetLineno(1790)
							if πE = πg.CheckLocal(πF, µmyoff, "myoff"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µotoff, "otoff"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, µmyoff, µotoff); πE != nil {
								continue
							}
							µbase_compare = πTemp002
							goto Label3
						Label3:
							if πE = πg.CheckLocal(πF, µbase_compare, "base_compare"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µbase_compare); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label8
							}
							goto Label9
							// line 1792: if base_compare:
							πF.SetLineno(1792)
						Label8:
							// line 1793: return _cmp((self._year, self._month, self._day,
							πF.SetLineno(1793)
							πTemp001 = πF.MakeArgs(2)
							πTemp005 = make([]*πg.Object, 7)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_year, nil); πE != nil {
								continue
							}
							πTemp005[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_month, nil); πE != nil {
								continue
							}
							πTemp005[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_day, nil); πE != nil {
								continue
							}
							πTemp005[2] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_hour, nil); πE != nil {
								continue
							}
							πTemp005[3] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_minute, nil); πE != nil {
								continue
							}
							πTemp005[4] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_second, nil); πE != nil {
								continue
							}
							πTemp005[5] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_microsecond, nil); πE != nil {
								continue
							}
							πTemp005[6] = πTemp003
							πTemp002 = πg.NewTuple(πTemp005...).ToObject()
							πTemp001[0] = πTemp002
							πTemp005 = make([]*πg.Object, 7)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µother, ß_year, nil); πE != nil {
								continue
							}
							πTemp005[0] = πTemp003
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µother, ß_month, nil); πE != nil {
								continue
							}
							πTemp005[1] = πTemp003
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µother, ß_day, nil); πE != nil {
								continue
							}
							πTemp005[2] = πTemp003
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µother, ß_hour, nil); πE != nil {
								continue
							}
							πTemp005[3] = πTemp003
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µother, ß_minute, nil); πE != nil {
								continue
							}
							πTemp005[4] = πTemp003
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µother, ß_second, nil); πE != nil {
								continue
							}
							πTemp005[5] = πTemp003
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µother, ß_microsecond, nil); πE != nil {
								continue
							}
							πTemp005[6] = πTemp003
							πTemp002 = πg.NewTuple(πTemp005...).ToObject()
							πTemp001[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_cmp); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πR = πTemp003
							continue
							goto Label9
						Label9:
							if πE = πg.CheckLocal(πF, µmyoff, "myoff"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(µmyoff == πTemp006).ToObject()
							πTemp002 = πTemp003
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label10
							}
							if πE = πg.CheckLocal(πF, µotoff, "otoff"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(µotoff == πTemp006).ToObject()
							πTemp002 = πTemp003
						Label10:
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label11
							}
							goto Label12
							// line 1799: if myoff is None or otoff is None:
							πF.SetLineno(1799)
						Label11:
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("can't compare offset-naive and offset-aware datetimes").ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1800: raise TypeError("can't compare offset-naive and offset-aware datetimes")
							πF.SetLineno(1800)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label12
						Label12:
							// line 1802: diff = self - other     # this will take offsets into account
							πF.SetLineno(1802)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Sub(πF, µself, µother); πE != nil {
								continue
							}
							µdiff = πTemp002
							if πE = πg.CheckLocal(πF, µdiff, "diff"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µdiff, ßdays, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.LT(πF, πTemp003, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label13
							}
							goto Label14
							// line 1803: if diff.days < 0:
							πF.SetLineno(1803)
						Label13:
							// line 1804: return -1
							πF.SetLineno(1804)
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πR = πTemp002
							continue
							goto Label14
						Label14:
							// line 1805: return diff and 1 or 0
							πF.SetLineno(1805)
							if πE = πg.CheckLocal(πF, µdiff, "diff"); πE != nil {
								continue
							}
							πTemp003 = µdiff
							if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if !πTemp007 {
								goto Label16
							}
							πTemp003 = πg.NewInt(1).ToObject()
						Label16:
							πTemp002 = πTemp003
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label15
							}
							πTemp002 = πg.NewInt(0).ToObject()
						Label15:
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
					if πE = πClass.SetItem(πF, ß_cmp.ToObject(), πTemp037); πE != nil {
						continue
					}
					// line 1807: def _add_timedelta(self, other, factor):
					πF.SetLineno(1807)
					πTemp005 = make([]πg.Param, 3)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "other", Def: nil}
					πTemp005[2] = πg.Param{Name: "factor", Def: nil}
					πTemp038 = πg.NewFunction(πg.NewCode("_add_timedelta", "/usr/lib/python2.7/datetime.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
						var µfactor *πg.Object = πArgs[2]; _ = µfactor
						var µy *πg.Object = πg.UnboundLocal; _ = µy
						var µm *πg.Object = πg.UnboundLocal; _ = µm
						var µd *πg.Object = πg.UnboundLocal; _ = µd
						var µhh *πg.Object = πg.UnboundLocal; _ = µhh
						var µmm *πg.Object = πg.UnboundLocal; _ = µmm
						var µss *πg.Object = πg.UnboundLocal; _ = µss
						var µus *πg.Object = πg.UnboundLocal; _ = µus
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
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 *πg.Object
						_ = πTemp009
						var πTemp010 πg.KWArgs
						_ = πTemp010
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 1808: y, m, d, hh, mm, ss, us = _normalize_datetime(
							πF.SetLineno(1808)
							πTemp001 = πF.MakeArgs(7)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_year, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_month, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_day, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µother, ßdays, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µfactor, "factor"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Mul(πF, πTemp005, µfactor); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp003, πTemp004); πE != nil {
								continue
							}
							πTemp001[2] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_hour, nil); πE != nil {
								continue
							}
							πTemp001[3] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_minute, nil); πE != nil {
								continue
							}
							πTemp001[4] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_second, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µother, ßseconds, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µfactor, "factor"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Mul(πF, πTemp005, µfactor); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp003, πTemp004); πE != nil {
								continue
							}
							πTemp001[5] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_microsecond, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µother, ßmicroseconds, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µfactor, "factor"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Mul(πF, πTemp005, µfactor); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp003, πTemp004); πE != nil {
								continue
							}
							πTemp001[6] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_normalize_datetime); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp006}, πg.TieTarget{Target: &πTemp007}, πg.TieTarget{Target: &πTemp008}, πg.TieTarget{Target: &πTemp009}}}, πTemp003); πE != nil {
								continue
							}
							µy = πTemp002
							µm = πTemp004
							µd = πTemp005
							µhh = πTemp006
							µmm = πTemp007
							µss = πTemp008
							µus = πTemp009
							// line 1816: return datetime(y, m, d, hh, mm, ss, us, tzinfo=self._tzinfo)
							πF.SetLineno(1816)
							πTemp001 = πF.MakeArgs(7)
							if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
								continue
							}
							πTemp001[0] = µy
							if πE = πg.CheckLocal(πF, µm, "m"); πE != nil {
								continue
							}
							πTemp001[1] = µm
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp001[2] = µd
							if πE = πg.CheckLocal(πF, µhh, "hh"); πE != nil {
								continue
							}
							πTemp001[3] = µhh
							if πE = πg.CheckLocal(πF, µmm, "mm"); πE != nil {
								continue
							}
							πTemp001[4] = µmm
							if πE = πg.CheckLocal(πF, µss, "ss"); πE != nil {
								continue
							}
							πTemp001[5] = µss
							if πE = πg.CheckLocal(πF, µus, "us"); πE != nil {
								continue
							}
							πTemp001[6] = µus
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_tzinfo, nil); πE != nil {
								continue
							}
							πTemp010 = πg.KWArgs{
								{"tzinfo", πTemp002},
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßdatetime); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, πTemp010); πE != nil {
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
					if πE = πClass.SetItem(πF, ß_add_timedelta.ToObject(), πTemp038); πE != nil {
						continue
					}
					// line 1818: def __add__(self, other):
					πF.SetLineno(1818)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "other", Def: nil}
					πTemp039 = πg.NewFunction(πg.NewCode("__add__", "/usr/lib/python2.7/datetime.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 1819: "Add a datetime and a timedelta."
							πF.SetLineno(1819)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp002[0] = µother
							if πTemp003, πE = πg.ResolveGlobal(πF, ßtimedelta); πE != nil {
								continue
							}
							πTemp002[1] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp005, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp005).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							goto Label2
							// line 1820: if not isinstance(other, timedelta):
							πF.SetLineno(1820)
						Label1:
							// line 1821: return NotImplemented
							πF.SetLineno(1821)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNotImplemented); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label2
						Label2:
							// line 1822: return self._add_timedelta(other, 1)
							πF.SetLineno(1822)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp002[0] = µother
							πTemp002[1] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_add_timedelta, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
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
					if πE = πClass.SetItem(πF, ß__add__.ToObject(), πTemp039); πE != nil {
						continue
					}
					// line 1819: "Add a datetime and a timedelta."
					πF.SetLineno(1819)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp040}, πg.NewStr("Add a datetime and a timedelta.").ToObject()); πE != nil {
						continue
					}
					if πTemp041, πE = πg.ResolveClass(πF, πClass, nil, ß__add__); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp041, ß__doc__, πTemp040); πE != nil {
						continue
					}
					// line 1824: __radd__ = __add__
					πF.SetLineno(1824)
					if πTemp040, πE = πg.ResolveClass(πF, πClass, nil, ß__add__); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ß__radd__.ToObject(), πTemp040); πE != nil {
						continue
					}
					// line 1826: def __sub__(self, other):
					πF.SetLineno(1826)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "other", Def: nil}
					πTemp040 = πg.NewFunction(πg.NewCode("__sub__", "/usr/lib/python2.7/datetime.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
						var µdelta_d *πg.Object = πg.UnboundLocal; _ = µdelta_d
						var µdelta_s *πg.Object = πg.UnboundLocal; _ = µdelta_s
						var µdelta_us *πg.Object = πg.UnboundLocal; _ = µdelta_us
						var µbase *πg.Object = πg.UnboundLocal; _ = µbase
						var µmyoff *πg.Object = πg.UnboundLocal; _ = µmyoff
						var µotoff *πg.Object = πg.UnboundLocal; _ = µotoff
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
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 *πg.Object
						_ = πTemp009
						var πTemp010 πg.KWArgs
						_ = πTemp010
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 1827: "Subtract two datetimes, or a datetime and a timedelta."
							πF.SetLineno(1827)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp002[0] = µother
							if πTemp003, πE = πg.ResolveGlobal(πF, ßdatetime); πE != nil {
								continue
							}
							πTemp002[1] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp005, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp005).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							goto Label2
							// line 1828: if not isinstance(other, datetime):
							πF.SetLineno(1828)
						Label1:
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp002[0] = µother
							if πTemp001, πE = πg.ResolveGlobal(πF, ßtimedelta); πE != nil {
								continue
							}
							πTemp002[1] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label3
							}
							goto Label4
							// line 1829: if isinstance(other, timedelta):
							πF.SetLineno(1829)
						Label3:
							// line 1830: return self._add_timedelta(other, -1)
							πF.SetLineno(1830)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp002[0] = µother
							if πTemp001, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp002[1] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_add_timedelta, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πR = πTemp003
							continue
							goto Label4
						Label4:
							// line 1831: return NotImplemented
							πF.SetLineno(1831)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNotImplemented); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label2
						Label2:
							// line 1833: delta_d = self.toordinal() - other.toordinal()
							πF.SetLineno(1833)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtoordinal, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µother, ßtoordinal, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Sub(πF, πTemp004, πTemp006); πE != nil {
								continue
							}
							µdelta_d = πTemp001
							// line 1834: delta_s = (self._hour - other._hour) * 3600 + \
							πF.SetLineno(1834)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ß_hour, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µother, ß_hour, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Sub(πF, πTemp007, πTemp008); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Mul(πF, πTemp006, πg.NewInt(3600).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µself, ß_minute, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, µother, ß_minute, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Sub(πF, πTemp008, πTemp009); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Mul(πF, πTemp007, πg.NewInt(60).ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, πTemp004, πTemp006); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ß_second, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µother, ß_second, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Sub(πF, πTemp006, πTemp007); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp003, πTemp004); πE != nil {
								continue
							}
							µdelta_s = πTemp001
							// line 1837: delta_us = self._microsecond - other._microsecond
							πF.SetLineno(1837)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_microsecond, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µother, ß_microsecond, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Sub(πF, πTemp003, πTemp004); πE != nil {
								continue
							}
							µdelta_us = πTemp001
							// line 1838: base = timedelta._create(delta_d, delta_s, delta_us, True)
							πF.SetLineno(1838)
							πTemp002 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µdelta_d, "delta_d"); πE != nil {
								continue
							}
							πTemp002[0] = µdelta_d
							if πE = πg.CheckLocal(πF, µdelta_s, "delta_s"); πE != nil {
								continue
							}
							πTemp002[1] = µdelta_s
							if πE = πg.CheckLocal(πF, µdelta_us, "delta_us"); πE != nil {
								continue
							}
							πTemp002[2] = µdelta_us
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp002[3] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßtimedelta); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ß_create, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µbase = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_tzinfo, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µother, ß_tzinfo, nil); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp003 == πTemp004).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label5
							}
							goto Label6
							// line 1839: if self._tzinfo is other._tzinfo:
							πF.SetLineno(1839)
						Label5:
							// line 1840: return base
							πF.SetLineno(1840)
							if πE = πg.CheckLocal(πF, µbase, "base"); πE != nil {
								continue
							}
							πR = µbase
							continue
							goto Label6
						Label6:
							// line 1841: myoff = self._utcoffset()
							πF.SetLineno(1841)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_utcoffset, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µmyoff = πTemp003
							// line 1842: otoff = other._utcoffset()
							πF.SetLineno(1842)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µother, ß_utcoffset, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µotoff = πTemp003
							if πE = πg.CheckLocal(πF, µmyoff, "myoff"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µotoff, "otoff"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, µmyoff, µotoff); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label7
							}
							goto Label8
							// line 1843: if myoff == otoff:
							πF.SetLineno(1843)
						Label7:
							// line 1844: return base
							πF.SetLineno(1844)
							if πE = πg.CheckLocal(πF, µbase, "base"); πE != nil {
								continue
							}
							πR = µbase
							continue
							goto Label8
						Label8:
							if πE = πg.CheckLocal(πF, µmyoff, "myoff"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(µmyoff == πTemp004).ToObject()
							πTemp001 = πTemp003
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label9
							}
							if πE = πg.CheckLocal(πF, µotoff, "otoff"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(µotoff == πTemp004).ToObject()
							πTemp001 = πTemp003
						Label9:
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label10
							}
							goto Label11
							// line 1845: if myoff is None or otoff is None:
							πF.SetLineno(1845)
						Label10:
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("can't subtract offset-naive and offset-aware datetimes").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 1846: raise TypeError("can't subtract offset-naive and offset-aware datetimes")
							πF.SetLineno(1846)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label11
						Label11:
							// line 1847: return base + timedelta(minutes = otoff-myoff)
							πF.SetLineno(1847)
							if πE = πg.CheckLocal(πF, µbase, "base"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µotoff, "otoff"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µmyoff, "myoff"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Sub(πF, µotoff, µmyoff); πE != nil {
								continue
							}
							πTemp010 = πg.KWArgs{
								{"minutes", πTemp003},
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßtimedelta); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, πTemp010); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, µbase, πTemp004); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__sub__.ToObject(), πTemp040); πE != nil {
						continue
					}
					// line 1827: "Subtract two datetimes, or a datetime and a timedelta."
					πF.SetLineno(1827)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp041}, πg.NewStr("Subtract two datetimes, or a datetime and a timedelta.").ToObject()); πE != nil {
						continue
					}
					if πTemp042, πE = πg.ResolveClass(πF, πClass, nil, ß__sub__); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp042, ß__doc__, πTemp041); πE != nil {
						continue
					}
					// line 1849: def __hash__(self):
					πF.SetLineno(1849)
					πTemp005 = make([]πg.Param, 1)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp041 = πg.NewFunction(πg.NewCode("__hash__", "/usr/lib/python2.7/datetime.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µtzoff *πg.Object = πg.UnboundLocal; _ = µtzoff
						var µdays *πg.Object = πg.UnboundLocal; _ = µdays
						var µseconds *πg.Object = πg.UnboundLocal; _ = µseconds
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
							default: panic("unexpected function state")
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_hashcode, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
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
							// line 1850: if self._hashcode == -1:
							πF.SetLineno(1850)
						Label1:
							// line 1851: tzoff = self._utcoffset()
							πF.SetLineno(1851)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_utcoffset, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µtzoff = πTemp002
							if πE = πg.CheckLocal(πF, µtzoff, "tzoff"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µtzoff == πTemp002).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label3
							}
							goto Label4
							// line 1852: if tzoff is None:
							πF.SetLineno(1852)
						Label3:
							// line 1853: self._hashcode = hash(self._getstate()[0])
							πF.SetLineno(1853)
							πTemp005 = πF.MakeArgs(1)
							πTemp001 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_getstate, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp006, πTemp001); πE != nil {
								continue
							}
							πTemp005[0] = πTemp002
							if πTemp001, πE = πg.ResolveGlobal(πF, ßhash); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_hashcode, πTemp001); πE != nil {
								continue
							}
							goto Label5
						Label4:
							// line 1855: days = _ymd2ord(self.year, self.month, self.day)
							πF.SetLineno(1855)
							πTemp005 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßyear, nil); πE != nil {
								continue
							}
							πTemp005[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßmonth, nil); πE != nil {
								continue
							}
							πTemp005[1] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßday, nil); πE != nil {
								continue
							}
							πTemp005[2] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_ymd2ord); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µdays = πTemp002
							// line 1856: seconds = self.hour * 3600 + (self.minute - tzoff) * 60 + self.second
							πF.SetLineno(1856)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßhour, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Mul(πF, πTemp006, πg.NewInt(3600).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µself, ßminute, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtzoff, "tzoff"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Sub(πF, πTemp008, µtzoff); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Mul(πF, πTemp007, πg.NewInt(60).ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp003, πTemp006); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßsecond, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp002, πTemp003); πE != nil {
								continue
							}
							µseconds = πTemp001
							// line 1857: self._hashcode = hash(timedelta(days, seconds, self.microsecond))
							πF.SetLineno(1857)
							πTemp005 = πF.MakeArgs(1)
							πTemp009 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µdays, "days"); πE != nil {
								continue
							}
							πTemp009[0] = µdays
							if πE = πg.CheckLocal(πF, µseconds, "seconds"); πE != nil {
								continue
							}
							πTemp009[1] = µseconds
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßmicrosecond, nil); πE != nil {
								continue
							}
							πTemp009[2] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßtimedelta); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							πTemp005[0] = πTemp002
							if πTemp001, πE = πg.ResolveGlobal(πF, ßhash); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_hashcode, πTemp001); πE != nil {
								continue
							}
							goto Label5
						Label5:
							goto Label2
						Label2:
							// line 1858: return self._hashcode
							πF.SetLineno(1858)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_hashcode, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__hash__.ToObject(), πTemp041); πE != nil {
						continue
					}
					// line 1862: def _getstate(self):
					πF.SetLineno(1862)
					πTemp005 = make([]πg.Param, 1)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp042 = πg.NewFunction(πg.NewCode("_getstate", "/usr/lib/python2.7/datetime.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µyhi *πg.Object = πg.UnboundLocal; _ = µyhi
						var µylo *πg.Object = πg.UnboundLocal; _ = µylo
						var µus2 *πg.Object = πg.UnboundLocal; _ = µus2
						var µus3 *πg.Object = πg.UnboundLocal; _ = µus3
						var µus1 *πg.Object = πg.UnboundLocal; _ = µus1
						var µbasestate *πg.Object = πg.UnboundLocal; _ = µbasestate
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
							// line 1863: yhi, ylo = divmod(self._year, 256)
							πF.SetLineno(1863)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_year, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = πg.NewInt(256).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßdivmod); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp004}}}, πTemp003); πE != nil {
								continue
							}
							µyhi = πTemp002
							µylo = πTemp004
							// line 1864: us2, us3 = divmod(self._microsecond, 256)
							πF.SetLineno(1864)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_microsecond, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = πg.NewInt(256).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßdivmod); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp004}}}, πTemp003); πE != nil {
								continue
							}
							µus2 = πTemp002
							µus3 = πTemp004
							// line 1865: us1, us2 = divmod(us2, 256)
							πF.SetLineno(1865)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µus2, "us2"); πE != nil {
								continue
							}
							πTemp001[0] = µus2
							πTemp001[1] = πg.NewInt(256).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßdivmod); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp004}}}, πTemp003); πE != nil {
								continue
							}
							µus1 = πTemp002
							µus2 = πTemp004
							// line 1866: basestate = _struct.pack('10B', yhi, ylo, self._month, self._day,
							πF.SetLineno(1866)
							πTemp001 = πF.MakeArgs(11)
							πTemp001[0] = ß10B.ToObject()
							if πE = πg.CheckLocal(πF, µyhi, "yhi"); πE != nil {
								continue
							}
							πTemp001[1] = µyhi
							if πE = πg.CheckLocal(πF, µylo, "ylo"); πE != nil {
								continue
							}
							πTemp001[2] = µylo
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_month, nil); πE != nil {
								continue
							}
							πTemp001[3] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_day, nil); πE != nil {
								continue
							}
							πTemp001[4] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_hour, nil); πE != nil {
								continue
							}
							πTemp001[5] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_minute, nil); πE != nil {
								continue
							}
							πTemp001[6] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_second, nil); πE != nil {
								continue
							}
							πTemp001[7] = πTemp002
							if πE = πg.CheckLocal(πF, µus1, "us1"); πE != nil {
								continue
							}
							πTemp001[8] = µus1
							if πE = πg.CheckLocal(πF, µus2, "us2"); πE != nil {
								continue
							}
							πTemp001[9] = µus2
							if πE = πg.CheckLocal(πF, µus3, "us3"); πE != nil {
								continue
							}
							πTemp001[10] = µus3
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_struct); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßpack, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µbasestate = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_tzinfo, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp003 == πTemp004).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							goto Label2
							// line 1869: if self._tzinfo is None:
							πF.SetLineno(1869)
						Label1:
							// line 1870: return (basestate,)
							πF.SetLineno(1870)
							if πE = πg.CheckLocal(πF, µbasestate, "basestate"); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple1(µbasestate).ToObject()
							πR = πTemp002
							continue
							goto Label3
						Label2:
							// line 1872: return (basestate, self._tzinfo)
							πF.SetLineno(1872)
							if πE = πg.CheckLocal(πF, µbasestate, "basestate"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_tzinfo, nil); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(µbasestate, πTemp003).ToObject()
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
					if πE = πClass.SetItem(πF, ß_getstate.ToObject(), πTemp042); πE != nil {
						continue
					}
					// line 1874: def __setstate(self, string, tzinfo):
					πF.SetLineno(1874)
					πTemp005 = make([]πg.Param, 3)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "string", Def: nil}
					πTemp005[2] = πg.Param{Name: "tzinfo", Def: nil}
					πTemp043 = πg.NewFunction(πg.NewCode("__setstate", "/usr/lib/python2.7/datetime.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µstring *πg.Object = πArgs[1]; _ = µstring
						var µtzinfo *πg.Object = πArgs[2]; _ = µtzinfo
						var µyhi *πg.Object = πg.UnboundLocal; _ = µyhi
						var µylo *πg.Object = πg.UnboundLocal; _ = µylo
						var µus1 *πg.Object = πg.UnboundLocal; _ = µus1
						var µus2 *πg.Object = πg.UnboundLocal; _ = µus2
						var µus3 *πg.Object = πg.UnboundLocal; _ = µus3
						var πTemp001 *πg.Object
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
						var πTemp007 bool
						_ = πTemp007
						var πTemp008 []*πg.Object
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
							if πE = πg.CheckLocal(πF, µtzinfo, "tzinfo"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(µtzinfo != πTemp004).ToObject()
							πTemp001 = πTemp003
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp002 {
								goto Label1
							}
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µtzinfo, "tzinfo"); πE != nil {
								continue
							}
							πTemp005[0] = µtzinfo
							if πTemp004, πE = πg.ResolveGlobal(πF, ß_tzinfo_class); πE != nil {
								continue
							}
							πTemp005[1] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp007, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp007).ToObject()
							πTemp001 = πTemp003
						Label1:
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label2
							}
							goto Label3
							// line 1875: if tzinfo is not None and not isinstance(tzinfo, _tzinfo_class):
							πF.SetLineno(1875)
						Label2:
							πTemp005 = πF.MakeArgs(1)
							πTemp005[0] = πg.NewStr("bad tzinfo state arg").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 1876: raise TypeError("bad tzinfo state arg")
							πF.SetLineno(1876)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label3
						Label3:
							// line 1877: (yhi, ylo, self._month, self._day, self._hour,
							πF.SetLineno(1877)
							πTemp005 = make([]*πg.Object, 10)
							πTemp008 = πF.MakeArgs(1)
							πTemp003 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µstring, πTemp003); πE != nil {
								continue
							}
							πTemp008[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp005[0] = πTemp004
							πTemp008 = πF.MakeArgs(1)
							πTemp003 = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µstring, πTemp003); πE != nil {
								continue
							}
							πTemp008[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp005[1] = πTemp004
							πTemp008 = πF.MakeArgs(1)
							πTemp003 = πg.NewInt(2).ToObject()
							if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µstring, πTemp003); πE != nil {
								continue
							}
							πTemp008[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp005[2] = πTemp004
							πTemp008 = πF.MakeArgs(1)
							πTemp003 = πg.NewInt(3).ToObject()
							if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µstring, πTemp003); πE != nil {
								continue
							}
							πTemp008[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp005[3] = πTemp004
							πTemp008 = πF.MakeArgs(1)
							πTemp003 = πg.NewInt(4).ToObject()
							if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µstring, πTemp003); πE != nil {
								continue
							}
							πTemp008[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp005[4] = πTemp004
							πTemp008 = πF.MakeArgs(1)
							πTemp003 = πg.NewInt(5).ToObject()
							if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µstring, πTemp003); πE != nil {
								continue
							}
							πTemp008[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp005[5] = πTemp004
							πTemp008 = πF.MakeArgs(1)
							πTemp003 = πg.NewInt(6).ToObject()
							if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µstring, πTemp003); πE != nil {
								continue
							}
							πTemp008[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp005[6] = πTemp004
							πTemp008 = πF.MakeArgs(1)
							πTemp003 = πg.NewInt(7).ToObject()
							if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µstring, πTemp003); πE != nil {
								continue
							}
							πTemp008[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp005[7] = πTemp004
							πTemp008 = πF.MakeArgs(1)
							πTemp003 = πg.NewInt(8).ToObject()
							if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µstring, πTemp003); πE != nil {
								continue
							}
							πTemp008[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp005[8] = πTemp004
							πTemp008 = πF.MakeArgs(1)
							πTemp003 = πg.NewInt(9).ToObject()
							if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µstring, πTemp003); πE != nil {
								continue
							}
							πTemp008[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp005[9] = πTemp004
							πTemp001 = πg.NewTuple(πTemp005...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp006}, πg.TieTarget{Target: &πTemp009}, πg.TieTarget{Target: &πTemp010}, πg.TieTarget{Target: &πTemp011}, πg.TieTarget{Target: &πTemp012}, πg.TieTarget{Target: &πTemp013}, πg.TieTarget{Target: &πTemp014}, πg.TieTarget{Target: &πTemp015}}}, πTemp001); πE != nil {
								continue
							}
							µyhi = πTemp003
							µylo = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_month, πTemp006); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_day, πTemp009); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_hour, πTemp010); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_minute, πTemp011); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_second, πTemp012); πE != nil {
								continue
							}
							µus1 = πTemp013
							µus2 = πTemp014
							µus3 = πTemp015
							// line 1882: self._year = yhi * 256 + ylo
							πF.SetLineno(1882)
							if πE = πg.CheckLocal(πF, µyhi, "yhi"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Mul(πF, µyhi, πg.NewInt(256).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µylo, "ylo"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp003, µylo); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_year, πTemp003); πE != nil {
								continue
							}
							// line 1883: self._microsecond = (((us1 << 8) | us2) << 8) | us3
							πF.SetLineno(1883)
							if πE = πg.CheckLocal(πF, µus1, "us1"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.LShift(πF, µus1, πg.NewInt(8).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µus2, "us2"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Or(πF, πTemp006, µus2); πE != nil {
								continue
							}
							if πTemp003, πE = πg.LShift(πF, πTemp004, πg.NewInt(8).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µus3, "us3"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Or(πF, πTemp003, µus3); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_microsecond, πTemp003); πE != nil {
								continue
							}
							// line 1884: self._tzinfo = tzinfo
							πF.SetLineno(1884)
							if πE = πg.CheckLocal(πF, µtzinfo, "tzinfo"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µtzinfo); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_tzinfo, πTemp001); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__setstate.ToObject(), πTemp043); πE != nil {
						continue
					}
					// line 1886: def __reduce__(self):
					πF.SetLineno(1886)
					πTemp005 = make([]πg.Param, 1)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp044 = πg.NewFunction(πg.NewCode("__reduce__", "/usr/lib/python2.7/datetime.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πTemp001 *πg.Object
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
							// line 1887: return (self.__class__, self._getstate())
							πF.SetLineno(1887)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß__class__, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_getstate, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple2(πTemp002, πTemp004).ToObject()
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
					if πE = πClass.SetItem(πF, ß__reduce__.ToObject(), πTemp044); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp029, πE = πTemp027.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp029 == nil {
				πTemp029 = πg.TypeType.ToObject()
			}
			if πTemp030, πE = πTemp029.Call(πF, []*πg.Object{πg.NewStr("datetime").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp027.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßdatetime.ToObject(), πTemp030); πE != nil {
				continue
			}
			// line 1890: datetime.min = datetime(1, 1, 1)
			πF.SetLineno(1890)
			πTemp002 = πF.MakeArgs(3)
			πTemp002[0] = πg.NewInt(1).ToObject()
			πTemp002[1] = πg.NewInt(1).ToObject()
			πTemp002[2] = πg.NewInt(1).ToObject()
			if πTemp028, πE = πg.ResolveGlobal(πF, ßdatetime); πE != nil {
				continue
			}
			if πTemp029, πE = πTemp028.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp028}, πTemp029); πE != nil {
				continue
			}
			if πTemp030, πE = πg.ResolveGlobal(πF, ßdatetime); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp030, ßmin, πTemp028); πE != nil {
				continue
			}
			// line 1891: datetime.max = datetime(9999, 12, 31, 23, 59, 59, 999999)
			πF.SetLineno(1891)
			πTemp002 = πF.MakeArgs(7)
			πTemp002[0] = πg.NewInt(9999).ToObject()
			πTemp002[1] = πg.NewInt(12).ToObject()
			πTemp002[2] = πg.NewInt(31).ToObject()
			πTemp002[3] = πg.NewInt(23).ToObject()
			πTemp002[4] = πg.NewInt(59).ToObject()
			πTemp002[5] = πg.NewInt(59).ToObject()
			πTemp002[6] = πg.NewInt(999999).ToObject()
			if πTemp028, πE = πg.ResolveGlobal(πF, ßdatetime); πE != nil {
				continue
			}
			if πTemp029, πE = πTemp028.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp028}, πTemp029); πE != nil {
				continue
			}
			if πTemp030, πE = πg.ResolveGlobal(πF, ßdatetime); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp030, ßmax, πTemp028); πE != nil {
				continue
			}
			// line 1892: datetime.resolution = timedelta(microseconds=1)
			πF.SetLineno(1892)
			πTemp031 = πg.KWArgs{
				{"microseconds", πg.NewInt(1).ToObject()},
			}
			if πTemp028, πE = πg.ResolveGlobal(πF, ßtimedelta); πE != nil {
				continue
			}
			if πTemp029, πE = πTemp028.Call(πF, nil, πTemp031); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp028}, πTemp029); πE != nil {
				continue
			}
			if πTemp030, πE = πg.ResolveGlobal(πF, ßdatetime); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp030, ßresolution, πTemp028); πE != nil {
				continue
			}
			// line 1895: def _isoweek1monday(year):
			πF.SetLineno(1895)
			πTemp003 = make([]πg.Param, 1)
			πTemp003[0] = πg.Param{Name: "year", Def: nil}
			πTemp028 = πg.NewFunction(πg.NewCode("_isoweek1monday", "/usr/lib/python2.7/datetime.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µyear *πg.Object = πArgs[0]; _ = µyear
				var µTHURSDAY *πg.Object = πg.UnboundLocal; _ = µTHURSDAY
				var µfirstday *πg.Object = πg.UnboundLocal; _ = µfirstday
				var µfirstweekday *πg.Object = πg.UnboundLocal; _ = µfirstweekday
				var µweek1monday *πg.Object = πg.UnboundLocal; _ = µweek1monday
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
					// line 1898: THURSDAY = 3
					πF.SetLineno(1898)
					µTHURSDAY = πg.NewInt(3).ToObject()
					// line 1899: firstday = _ymd2ord(year, 1, 1)
					πF.SetLineno(1899)
					πTemp001 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µyear, "year"); πE != nil {
						continue
					}
					πTemp001[0] = µyear
					πTemp001[1] = πg.NewInt(1).ToObject()
					πTemp001[2] = πg.NewInt(1).ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_ymd2ord); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µfirstday = πTemp003
					// line 1900: firstweekday = (firstday + 6) % 7  # See weekday() above
					πF.SetLineno(1900)
					if πE = πg.CheckLocal(πF, µfirstday, "firstday"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Add(πF, µfirstday, πg.NewInt(6).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mod(πF, πTemp003, πg.NewInt(7).ToObject()); πE != nil {
						continue
					}
					µfirstweekday = πTemp002
					// line 1901: week1monday = firstday - firstweekday
					πF.SetLineno(1901)
					if πE = πg.CheckLocal(πF, µfirstday, "firstday"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µfirstweekday, "firstweekday"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Sub(πF, µfirstday, µfirstweekday); πE != nil {
						continue
					}
					µweek1monday = πTemp002
					if πE = πg.CheckLocal(πF, µfirstweekday, "firstweekday"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µTHURSDAY, "THURSDAY"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GT(πF, µfirstweekday, µTHURSDAY); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label1
					}
					goto Label2
					// line 1902: if firstweekday > THURSDAY:
					πF.SetLineno(1902)
				Label1:
					// line 1903: week1monday += 7
					πF.SetLineno(1903)
					if πE = πg.CheckLocal(πF, µweek1monday, "week1monday"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IAdd(πF, µweek1monday, πg.NewInt(7).ToObject()); πE != nil {
						continue
					}
					µweek1monday = πTemp002
					goto Label2
				Label2:
					// line 1904: return week1monday
					πF.SetLineno(1904)
					if πE = πg.CheckLocal(πF, µweek1monday, "week1monday"); πE != nil {
						continue
					}
					πR = µweek1monday
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_isoweek1monday.ToObject(), πTemp028); πE != nil {
				continue
			}
			// line 1906: """
			πF.SetLineno(1906)
		}
		return nil, πE
	})
	πg.RegisterModule("datetime", Code)
}

