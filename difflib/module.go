package difflib
import (
	πg "github.com/pygolin/runtime"
	// _ "github.com/pygolin/stdlib/operator"
	// _ "github.com/pygolin/stdlib/re"
	// _ "github.com/pygolin/stdlib/heapq"
	// _ "github.com/pygolin/stdlib/functools"
)
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/difflib.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ß := πg.InternStr("")
		ßDiffer := πg.InternStr("Differ")
		ßFalse := πg.InternStr("False")
		ßHtmlDiff := πg.InternStr("HtmlDiff")
		ßIS_CHARACTER_JUNK := πg.InternStr("IS_CHARACTER_JUNK")
		ßIS_LINE_JUNK := πg.InternStr("IS_LINE_JUNK")
		ßKeyError := πg.InternStr("KeyError")
		ßMatch := πg.InternStr("Match")
		ßNone := πg.InternStr("None")
		ßOrderedDict := πg.InternStr("OrderedDict")
		ßSequenceMatcher := πg.InternStr("SequenceMatcher")
		ßStopIteration := πg.InternStr("StopIteration")
		ßTrue := πg.InternStr("True")
		ßTypeError := πg.InternStr("TypeError")
		ßValueError := πg.InternStr("ValueError")
		ßX := πg.InternStr("X")
		ß__all__ := πg.InternStr("__all__")
		ß__chain_b := πg.InternStr("__chain_b")
		ß__contains__ := πg.InternStr("__contains__")
		ß__dict__ := πg.InternStr("__dict__")
		ß__doc__ := πg.InternStr("__doc__")
		ß__getnewargs__ := πg.InternStr("__getnewargs__")
		ß__getstate__ := πg.InternStr("__getstate__")
		ß__init__ := πg.InternStr("__init__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß__new__ := πg.InternStr("__new__")
		ß__repr__ := πg.InternStr("__repr__")
		ß__slots__ := πg.InternStr("__slots__")
		ß_asdict := πg.InternStr("_asdict")
		ß_calculate_ratio := πg.InternStr("_calculate_ratio")
		ß_charjunk := πg.InternStr("_charjunk")
		ß_collect_lines := πg.InternStr("_collect_lines")
		ß_convert_flags := πg.InternStr("_convert_flags")
		ß_count_leading := πg.InternStr("_count_leading")
		ß_default_prefix := πg.InternStr("_default_prefix")
		ß_dump := πg.InternStr("_dump")
		ß_fancy_helper := πg.InternStr("_fancy_helper")
		ß_fancy_replace := πg.InternStr("_fancy_replace")
		ß_fields := πg.InternStr("_fields")
		ß_file_template := πg.InternStr("_file_template")
		ß_format_line := πg.InternStr("_format_line")
		ß_format_range_context := πg.InternStr("_format_range_context")
		ß_format_range_unified := πg.InternStr("_format_range_unified")
		ß_itemgetter := πg.InternStr("_itemgetter")
		ß_legend := πg.InternStr("_legend")
		ß_line_wrapper := πg.InternStr("_line_wrapper")
		ß_linejunk := πg.InternStr("_linejunk")
		ß_make := πg.InternStr("_make")
		ß_make_prefix := πg.InternStr("_make_prefix")
		ß_mdiff := πg.InternStr("_mdiff")
		ß_plain_replace := πg.InternStr("_plain_replace")
		ß_prefix := πg.InternStr("_prefix")
		ß_property := πg.InternStr("_property")
		ß_qformat := πg.InternStr("_qformat")
		ß_replace := πg.InternStr("_replace")
		ß_split_line := πg.InternStr("_split_line")
		ß_styles := πg.InternStr("_styles")
		ß_tab_newline_replace := πg.InternStr("_tab_newline_replace")
		ß_table_template := πg.InternStr("_table_template")
		ß_tabsize := πg.InternStr("_tabsize")
		ß_tuple := πg.InternStr("_tuple")
		ß_wrapcolumn := πg.InternStr("_wrapcolumn")
		ßa := πg.InternStr("a")
		ßadd := πg.InternStr("add")
		ßany := πg.InternStr("any")
		ßappend := πg.InternStr("append")
		ßautojunk := πg.InternStr("autojunk")
		ßb := πg.InternStr("b")
		ßb2j := πg.InternStr("b2j")
		ßcharjunk := πg.InternStr("charjunk")
		ßclassmethod := πg.InternStr("classmethod")
		ßcompare := πg.InternStr("compare")
		ßcompile := πg.InternStr("compile")
		ßcontext_diff := πg.InternStr("context_diff")
		ßcount := πg.InternStr("count")
		ßdelete := πg.InternStr("delete")
		ßdict := πg.InternStr("dict")
		ßenumerate := πg.InternStr("enumerate")
		ßequal := πg.InternStr("equal")
		ßexpandtabs := πg.InternStr("expandtabs")
		ßfind_longest_match := πg.InternStr("find_longest_match")
		ßfullbcount := πg.InternStr("fullbcount")
		ßfunctools := πg.InternStr("functools")
		ßget := πg.InternStr("get")
		ßget_close_matches := πg.InternStr("get_close_matches")
		ßget_grouped_opcodes := πg.InternStr("get_grouped_opcodes")
		ßget_matching_blocks := πg.InternStr("get_matching_blocks")
		ßget_opcodes := πg.InternStr("get_opcodes")
		ßgroup := πg.InternStr("group")
		ßheapq := πg.InternStr("heapq")
		ßinsert := πg.InternStr("insert")
		ßint := πg.InternStr("int")
		ßisbjunk := πg.InternStr("isbjunk")
		ßisbpopular := πg.InternStr("isbpopular")
		ßisjunk := πg.InternStr("isjunk")
		ßitemgetter := πg.InternStr("itemgetter")
		ßitems := πg.InternStr("items")
		ßjoin := πg.InternStr("join")
		ßkeys := πg.InternStr("keys")
		ßlen := πg.InternStr("len")
		ßlinejunk := πg.InternStr("linejunk")
		ßlist := πg.InternStr("list")
		ßmake_file := πg.InternStr("make_file")
		ßmake_table := πg.InternStr("make_table")
		ßmap := πg.InternStr("map")
		ßmatch := πg.InternStr("match")
		ßmatching_blocks := πg.InternStr("matching_blocks")
		ßmax := πg.InternStr("max")
		ßmin := πg.InternStr("min")
		ßndiff := πg.InternStr("ndiff")
		ßnext := πg.InternStr("next")
		ßnlargest := πg.InternStr("nlargest")
		ßobject := πg.InternStr("object")
		ßopcodes := πg.InternStr("opcodes")
		ßoperator := πg.InternStr("operator")
		ßpop := πg.InternStr("pop")
		ßproperty := πg.InternStr("property")
		ßquick_ratio := πg.InternStr("quick_ratio")
		ßrange := πg.InternStr("range")
		ßratio := πg.InternStr("ratio")
		ßre := πg.InternStr("re")
		ßreal_quick_ratio := πg.InternStr("real_quick_ratio")
		ßreduce := πg.InternStr("reduce")
		ßreplace := πg.InternStr("replace")
		ßrestore := πg.InternStr("restore")
		ßrstrip := πg.InternStr("rstrip")
		ßset := πg.InternStr("set")
		ßset_seq1 := πg.InternStr("set_seq1")
		ßset_seq2 := πg.InternStr("set_seq2")
		ßset_seqs := πg.InternStr("set_seqs")
		ßsetdefault := πg.InternStr("setdefault")
		ßsize := πg.InternStr("size")
		ßsort := πg.InternStr("sort")
		ßspan := πg.InternStr("span")
		ßstartswith := πg.InternStr("startswith")
		ßsub := πg.InternStr("sub")
		ßtuple := πg.InternStr("tuple")
		ßunified_diff := πg.InternStr("unified_diff")
		ßxrange := πg.InternStr("xrange")
		ßzip := πg.InternStr("zip")
		var πTemp001 []*πg.Object
		_ = πTemp001
		var πTemp002 *πg.Object
		_ = πTemp002
		var πTemp003 *πg.Object
		_ = πTemp003
		var πTemp004 []πg.Param
		_ = πTemp004
		var πTemp005 *πg.Dict
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
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default: panic("unexpected function state")
			}
			// line 1: """
			πF.SetLineno(1)
			// line 1: """
			πF.SetLineno(1)
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\nModule difflib -- helpers for computing deltas between objects.\n\nFunction get_close_matches(word, possibilities, n=3, cutoff=0.6):\n    Use SequenceMatcher to return list of the best \"good enough\" matches.\n\nFunction context_diff(a, b):\n    For two lists of strings, return a delta in context diff format.\n\nFunction ndiff(a, b):\n    Return a delta: the difference between `a` and `b` (lists of strings).\n\nFunction restore(delta, which):\n    Return one of the two sequences that generated an ndiff delta.\n\nFunction unified_diff(a, b):\n    For two lists of strings, return a delta in unified diff format.\n\nClass SequenceMatcher:\n    A flexible class for comparing pairs of sequences of any type.\n\nClass Differ:\n    For producing human-readable deltas from sequences of lines of text.\n\nClass HtmlDiff:\n    For producing HTML side by side comparison with change highlights.\n").ToObject()); πE != nil {
				continue
			}
			// line 29: __all__ = ['get_close_matches', 'ndiff', 'restore', 'SequenceMatcher',
			πF.SetLineno(29)
			πTemp001 = make([]*πg.Object, 11)
			πTemp001[0] = ßget_close_matches.ToObject()
			πTemp001[1] = ßndiff.ToObject()
			πTemp001[2] = ßrestore.ToObject()
			πTemp001[3] = ßSequenceMatcher.ToObject()
			πTemp001[4] = ßDiffer.ToObject()
			πTemp001[5] = ßIS_CHARACTER_JUNK.ToObject()
			πTemp001[6] = ßIS_LINE_JUNK.ToObject()
			πTemp001[7] = ßcontext_diff.ToObject()
			πTemp001[8] = ßunified_diff.ToObject()
			πTemp001[9] = ßHtmlDiff.ToObject()
			πTemp001[10] = ßMatch.ToObject()
			πTemp002 = πg.NewList(πTemp001...).ToObject()
			if πE = πF.Globals().SetItem(πF, ß__all__.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 33: import heapq
			πF.SetLineno(33)
			if πTemp001, πE = πg.ImportModule(πF, "heapq"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ßheapq.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 36: import functools
			πF.SetLineno(36)
			if πTemp001, πE = πg.ImportModule(πF, "functools"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ßfunctools.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 37: reduce = functools.reduce
			πF.SetLineno(37)
			if πTemp002, πE = πg.ResolveGlobal(πF, ßfunctools); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßreduce, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßreduce.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 39: import operator
			πF.SetLineno(39)
			if πTemp001, πE = πg.ImportModule(πF, "operator"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ßoperator.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 40: _itemgetter = operator.itemgetter
			πF.SetLineno(40)
			if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßitemgetter, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_itemgetter.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 41: _property = property
			πF.SetLineno(41)
			if πTemp002, πE = πg.ResolveGlobal(πF, ßproperty); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_property.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 42: _tuple = tuple
			πF.SetLineno(42)
			if πTemp002, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_tuple.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 44: def setdefault(d, k, default=None):
			πF.SetLineno(44)
			πTemp004 = make([]πg.Param, 3)
			πTemp004[0] = πg.Param{Name: "d", Def: nil}
			πTemp004[1] = πg.Param{Name: "k", Def: nil}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp004[2] = πg.Param{Name: "default", Def: πTemp003}
			πTemp002 = πg.NewFunction(πg.NewCode("setdefault", "/usr/lib/python2.7/difflib.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µd *πg.Object = πArgs[0]; _ = µd
				var µk *πg.Object = πArgs[1]; _ = µk
				var µdefault *πg.Object = πArgs[2]; _ = µdefault
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 bool
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
					if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Contains(πF, µd, µk); πE != nil {
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
					// line 45: if k not in d:
					πF.SetLineno(45)
				Label1:
					// line 46: d[k] = default
					πF.SetLineno(46)
					if πE = πg.CheckLocal(πF, µdefault, "default"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µdefault); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
						continue
					}
					πTemp003 = µk
					if πE = πg.SetItem(πF, µd, πTemp003, πTemp001); πE != nil {
						continue
					}
					goto Label2
				Label2:
					// line 47: return d[k]
					πF.SetLineno(47)
					if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
						continue
					}
					πTemp001 = µk
					if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µd, πTemp001); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßsetdefault.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 50: class Match(tuple):
			πF.SetLineno(50)
			πTemp001 = make([]*πg.Object, 1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
				continue
			}
			πTemp001[0] = πTemp007
			πTemp005 = πg.NewDict()
			if πTemp003, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp005.SetItem(πF, ß__module__.ToObject(), πTemp003); πE != nil {
				continue
			}
			_, πE = πg.NewCode("Match", "/usr/lib/python2.7/difflib.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
				var πTemp012 []*πg.Object
				_ = πTemp012
				var πTemp013 πg.KWArgs
				_ = πTemp013
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 51: 'Match(a, b, size)'
					πF.SetLineno(51)
					// line 51: 'Match(a, b, size)'
					πF.SetLineno(51)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Match(a, b, size)").ToObject()); πE != nil {
						continue
					}
					// line 53: __slots__ = ()
					πF.SetLineno(53)
					πTemp001 = πg.NewTuple0().ToObject()
					if πE = πClass.SetItem(πF, ß__slots__.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 55: _fields = ('a', 'b', 'size')
					πF.SetLineno(55)
					πTemp001 = πg.NewTuple3(ßa.ToObject(), ßb.ToObject(), ßsize.ToObject()).ToObject()
					if πE = πClass.SetItem(πF, ß_fields.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 57: def __new__(_cls, a, b, size):
					πF.SetLineno(57)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "_cls", Def: nil}
					πTemp002[1] = πg.Param{Name: "a", Def: nil}
					πTemp002[2] = πg.Param{Name: "b", Def: nil}
					πTemp002[3] = πg.Param{Name: "size", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__new__", "/usr/lib/python2.7/difflib.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µ_cls *πg.Object = πArgs[0]; _ = µ_cls
						var µa *πg.Object = πArgs[1]; _ = µa
						var µb *πg.Object = πArgs[2]; _ = µb
						var µsize *πg.Object = πArgs[3]; _ = µsize
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
							// line 58: 'Create new instance of Match(a, b, size)'
							πF.SetLineno(58)
							// line 59: return _tuple.__new__(_cls, (a, b, size))
							πF.SetLineno(59)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µ_cls, "_cls"); πE != nil {
								continue
							}
							πTemp001[0] = µ_cls
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µsize, "size"); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple3(µa, µb, µsize).ToObject()
							πTemp001[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_tuple); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ß__new__, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__new__.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 58: 'Create new instance of Match(a, b, size)'
					πF.SetLineno(58)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewStr("Create new instance of Match(a, b, size)").ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ß__new__); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp004, ß__doc__, πTemp003); πE != nil {
						continue
					}
					// line 62: def _make(cls, iterable, new=tuple.__new__, len=len):
					πF.SetLineno(62)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "cls", Def: nil}
					πTemp002[1] = πg.Param{Name: "iterable", Def: nil}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßtuple); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp004, ß__new__, nil); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "new", Def: πTemp005}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßlen); πE != nil {
						continue
					}
					πTemp002[3] = πg.Param{Name: "len", Def: πTemp004}
					πTemp003 = πg.NewFunction(πg.NewCode("_make", "/usr/lib/python2.7/difflib.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µcls *πg.Object = πArgs[0]; _ = µcls
						var µiterable *πg.Object = πArgs[1]; _ = µiterable
						var µnew *πg.Object = πArgs[2]; _ = µnew
						var µlen *πg.Object = πArgs[3]; _ = µlen
						var µresult *πg.Object = πg.UnboundLocal; _ = µresult
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 63: 'Make a new Match object from a sequence or iterable'
							πF.SetLineno(63)
							// line 64: result = new(cls, iterable)
							πF.SetLineno(64)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							πTemp001[0] = µcls
							if πE = πg.CheckLocal(πF, µiterable, "iterable"); πE != nil {
								continue
							}
							πTemp001[1] = µiterable
							if πE = πg.CheckLocal(πF, µnew, "new"); πE != nil {
								continue
							}
							if πTemp002, πE = µnew.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µresult = πTemp002
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							πTemp001[0] = µresult
							if πE = πg.CheckLocal(πF, µlen, "len"); πE != nil {
								continue
							}
							if πTemp003, πE = µlen.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.NE(πF, πTemp003, πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 65: if len(result) != 3:
							πF.SetLineno(65)
						Label1:
							πTemp001 = πF.MakeArgs(1)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							πTemp005[0] = µresult
							if πE = πg.CheckLocal(πF, µlen, "len"); πE != nil {
								continue
							}
							if πTemp003, πE = µlen.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("Expected 3 arguments, got %d").ToObject(), πTemp003); πE != nil {
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
							// line 66: raise TypeError('Expected 3 arguments, got %d' % len(result))
							πF.SetLineno(66)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label2
						Label2:
							// line 67: return result
							πF.SetLineno(67)
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
					if πE = πClass.SetItem(πF, ß_make.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 63: 'Make a new Match object from a sequence or iterable'
					πF.SetLineno(63)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πg.NewStr("Make a new Match object from a sequence or iterable").ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ß_make); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp005, ß__doc__, πTemp004); πE != nil {
						continue
					}
					// line 68: _make = classmethod(_make)
					πF.SetLineno(68)
					πTemp006 = πF.MakeArgs(1)
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ß_make); πE != nil {
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
					if πE = πClass.SetItem(πF, ß_make.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 70: def __repr__(self):
					πF.SetLineno(70)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("__repr__", "/usr/lib/python2.7/difflib.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 71: 'Return a nicely formatted representation string'
							πF.SetLineno(71)
							// line 72: return 'Match(a=%r, b=%r, size=%r)' % self
							πF.SetLineno(72)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("Match(a=%r, b=%r, size=%r)").ToObject(), µself); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__repr__.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 71: 'Return a nicely formatted representation string'
					πF.SetLineno(71)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πg.NewStr("Return a nicely formatted representation string").ToObject()); πE != nil {
						continue
					}
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ß__repr__); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp007, ß__doc__, πTemp005); πE != nil {
						continue
					}
					// line 74: def _asdict(self):
					πF.SetLineno(74)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("_asdict", "/usr/lib/python2.7/difflib.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 75: 'Return a new OrderedDict which maps field names to their values'
							πF.SetLineno(75)
							// line 76: return OrderedDict(zip(self._fields, self))
							πF.SetLineno(76)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_fields, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp002[1] = µself
							if πTemp003, πE = πg.ResolveGlobal(πF, ßzip); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßOrderedDict); πE != nil {
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
					if πE = πClass.SetItem(πF, ß_asdict.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 75: 'Return a new OrderedDict which maps field names to their values'
					πF.SetLineno(75)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp007}, πg.NewStr("Return a new OrderedDict which maps field names to their values").ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ß_asdict); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp008, ß__doc__, πTemp007); πE != nil {
						continue
					}
					// line 78: def _replace(_self, **kwds):
					πF.SetLineno(78)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "_self", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("_replace", "/usr/lib/python2.7/difflib.py", πTemp002, πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µ_self *πg.Object = πArgs[0]; _ = µ_self
						var µkwds *πg.Object = πArgs[1]; _ = µkwds
						var µresult *πg.Object = πg.UnboundLocal; _ = µresult
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 79: 'Return a new Match object replacing specified fields with new values'
							πF.SetLineno(79)
							// line 80: result = _self._make(map(kwds.pop, ('a', 'b', 'size'), _self))
							πF.SetLineno(80)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µkwds, "kwds"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µkwds, ßpop, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							πTemp003 = πg.NewTuple3(ßa.ToObject(), ßb.ToObject(), ßsize.ToObject()).ToObject()
							πTemp002[1] = πTemp003
							if πE = πg.CheckLocal(πF, µ_self, "_self"); πE != nil {
								continue
							}
							πTemp002[2] = µ_self
							if πTemp003, πE = πg.ResolveGlobal(πF, ßmap); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							if πE = πg.CheckLocal(πF, µ_self, "_self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µ_self, ß_make, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µresult = πTemp004
							if πE = πg.CheckLocal(πF, µkwds, "kwds"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, µkwds); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							goto Label2
							// line 81: if kwds:
							πF.SetLineno(81)
						Label1:
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µkwds, "kwds"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µkwds, ßkeys, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Mod(πF, πg.NewStr("Got unexpected field names: %r").ToObject(), πTemp006); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 82: raise ValueError('Got unexpected field names: %r' % kwds.keys())
							πF.SetLineno(82)
							πE = πF.Raise(πTemp004, nil, nil)
							continue
							goto Label2
						Label2:
							// line 83: return result
							πF.SetLineno(83)
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
					if πE = πClass.SetItem(πF, ß_replace.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 79: 'Return a new Match object replacing specified fields with new values'
					πF.SetLineno(79)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp008}, πg.NewStr("Return a new Match object replacing specified fields with new values").ToObject()); πE != nil {
						continue
					}
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ß_replace); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp009, ß__doc__, πTemp008); πE != nil {
						continue
					}
					// line 85: def __getnewargs__(self):
					πF.SetLineno(85)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("__getnewargs__", "/usr/lib/python2.7/difflib.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 86: 'Return self as a plain tuple.  Used by copy and pickle.'
							πF.SetLineno(86)
							// line 87: return tuple(self)
							πF.SetLineno(87)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__getnewargs__.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 86: 'Return self as a plain tuple.  Used by copy and pickle.'
					πF.SetLineno(86)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp009}, πg.NewStr("Return self as a plain tuple.  Used by copy and pickle.").ToObject()); πE != nil {
						continue
					}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ß__getnewargs__); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp010, ß__doc__, πTemp009); πE != nil {
						continue
					}
					// line 89: __dict__ = _property(_asdict)
					πF.SetLineno(89)
					πTemp006 = πF.MakeArgs(1)
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ß_asdict); πE != nil {
						continue
					}
					πTemp006[0] = πTemp009
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ß_property); πE != nil {
						continue
					}
					if πTemp010, πE = πTemp009.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πE = πClass.SetItem(πF, ß__dict__.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 91: def __getstate__(self):
					πF.SetLineno(91)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("__getstate__", "/usr/lib/python2.7/difflib.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 92: 'Exclude the OrderedDict from pickling'
							πF.SetLineno(92)
							// line 93: pass
							πF.SetLineno(93)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß__getstate__.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 92: 'Exclude the OrderedDict from pickling'
					πF.SetLineno(92)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp010}, πg.NewStr("Exclude the OrderedDict from pickling").ToObject()); πE != nil {
						continue
					}
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ß__getstate__); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp011, ß__doc__, πTemp010); πE != nil {
						continue
					}
					// line 95: a = _property(_itemgetter(0), doc='Alias for field number 0')
					πF.SetLineno(95)
					πTemp006 = πF.MakeArgs(1)
					πTemp012 = πF.MakeArgs(1)
					πTemp012[0] = πg.NewInt(0).ToObject()
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ß_itemgetter); πE != nil {
						continue
					}
					if πTemp011, πE = πTemp010.Call(πF, πTemp012, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp012)
					πTemp006[0] = πTemp011
					πTemp013 = πg.KWArgs{
						{"doc", πg.NewStr("Alias for field number 0").ToObject()},
					}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ß_property); πE != nil {
						continue
					}
					if πTemp011, πE = πTemp010.Call(πF, πTemp006, πTemp013); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πE = πClass.SetItem(πF, ßa.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 97: b = _property(_itemgetter(1), doc='Alias for field number 1')
					πF.SetLineno(97)
					πTemp006 = πF.MakeArgs(1)
					πTemp012 = πF.MakeArgs(1)
					πTemp012[0] = πg.NewInt(1).ToObject()
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ß_itemgetter); πE != nil {
						continue
					}
					if πTemp011, πE = πTemp010.Call(πF, πTemp012, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp012)
					πTemp006[0] = πTemp011
					πTemp013 = πg.KWArgs{
						{"doc", πg.NewStr("Alias for field number 1").ToObject()},
					}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ß_property); πE != nil {
						continue
					}
					if πTemp011, πE = πTemp010.Call(πF, πTemp006, πTemp013); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πE = πClass.SetItem(πF, ßb.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 99: size = _property(_itemgetter(2), doc='Alias for field number 2')
					πF.SetLineno(99)
					πTemp006 = πF.MakeArgs(1)
					πTemp012 = πF.MakeArgs(1)
					πTemp012[0] = πg.NewInt(2).ToObject()
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ß_itemgetter); πE != nil {
						continue
					}
					if πTemp011, πE = πTemp010.Call(πF, πTemp012, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp012)
					πTemp006[0] = πTemp011
					πTemp013 = πg.KWArgs{
						{"doc", πg.NewStr("Alias for field number 2").ToObject()},
					}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ß_property); πE != nil {
						continue
					}
					if πTemp011, πE = πTemp010.Call(πF, πTemp006, πTemp013); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πE = πClass.SetItem(πF, ßsize.ToObject(), πTemp011); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp006, πE = πTemp005.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp006 == nil {
				πTemp006 = πg.TypeType.ToObject()
			}
			if πTemp007, πE = πTemp006.Call(πF, []*πg.Object{πg.NewStr("Match").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp005.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßMatch.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 101: def _calculate_ratio(matches, length):
			πF.SetLineno(101)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "matches", Def: nil}
			πTemp004[1] = πg.Param{Name: "length", Def: nil}
			πTemp003 = πg.NewFunction(πg.NewCode("_calculate_ratio", "/usr/lib/python2.7/difflib.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µmatches *πg.Object = πArgs[0]; _ = µmatches
				var µlength *πg.Object = πArgs[1]; _ = µlength
				var πTemp001 bool
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
					if πE = πg.CheckLocal(πF, µlength, "length"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.IsTrue(πF, µlength); πE != nil {
						continue
					}
					if πTemp001 {
						goto Label1
					}
					goto Label2
					// line 102: if length:
					πF.SetLineno(102)
				Label1:
					// line 103: return 2.0 * matches / length
					πF.SetLineno(103)
					if πE = πg.CheckLocal(πF, µmatches, "matches"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Mul(πF, πg.NewFloat(2.0).ToObject(), µmatches); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µlength, "length"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Div(πF, πTemp003, µlength); πE != nil {
						continue
					}
					πR = πTemp002
					continue
					goto Label2
				Label2:
					// line 104: return 1.0
					πF.SetLineno(104)
					πR = πg.NewFloat(1.0).ToObject()
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_calculate_ratio.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 106: class SequenceMatcher(object):
			πF.SetLineno(106)
			πTemp001 = make([]*πg.Object, 1)
			if πTemp008, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp001[0] = πTemp008
			πTemp005 = πg.NewDict()
			if πTemp006, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp005.SetItem(πF, ß__module__.ToObject(), πTemp006); πE != nil {
				continue
			}
			_, πE = πg.NewCode("SequenceMatcher", "/usr/lib/python2.7/difflib.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 108: """
					πF.SetLineno(108)
					// line 108: """
					πF.SetLineno(108)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    SequenceMatcher is a flexible class for comparing pairs of sequences of\n    any type, so long as the sequence elements are hashable.  The basic\n    algorithm predates, and is a little fancier than, an algorithm\n    published in the late 1980's by Ratcliff and Obershelp under the\n    hyperbolic name \"gestalt pattern matching\".  The basic idea is to find\n    the longest contiguous matching subsequence that contains no \"junk\"\n    elements (R-O doesn't address junk).  The same idea is then applied\n    recursively to the pieces of the sequences to the left and to the right\n    of the matching subsequence.  This does not yield minimal edit\n    sequences, but does tend to yield matches that \"look right\" to people.\n\n    SequenceMatcher tries to compute a \"human-friendly diff\" between two\n    sequences.  Unlike e.g. UNIX(tm) diff, the fundamental notion is the\n    longest *contiguous* & junk-free matching subsequence.  That's what\n    catches peoples' eyes.  The Windows(tm) windiff has another interesting\n    notion, pairing up elements that appear uniquely in each sequence.\n    That, and the method here, appear to yield more intuitive difference\n    reports than does diff.  This method appears to be the least vulnerable\n    to synching up on blocks of \"junk lines\", though (like blank lines in\n    ordinary text files, or maybe \"<P>\" lines in HTML files).  That may be\n    because this is the only method of the 3 that has a *concept* of\n    \"junk\" <wink>.\n\n    Example, comparing two strings, and considering blanks to be \"junk\":\n\n    >>> s = SequenceMatcher(lambda x: x == \" \",\n    ...                     \"private Thread currentThread;\",\n    ...                     \"private volatile Thread currentThread;\")\n    >>>\n\n    .ratio() returns a float in [0, 1], measuring the \"similarity\" of the\n    sequences.  As a rule of thumb, a .ratio() value over 0.6 means the\n    sequences are close matches:\n\n    >>> print round(s.ratio(), 3)\n    0.866\n    >>>\n\n    If you're only interested in where the sequences match,\n    .get_matching_blocks() is handy:\n\n    >>> for block in s.get_matching_blocks():\n    ...     print \"a[%d] and b[%d] match for %d elements\" % block\n    a[0] and b[0] match for 8 elements\n    a[8] and b[17] match for 21 elements\n    a[29] and b[38] match for 0 elements\n\n    Note that the last tuple returned by .get_matching_blocks() is always a\n    dummy, (len(a), len(b), 0), and this is the only case in which the last\n    tuple element (number of elements matched) is 0.\n\n    If you want to know how to change the first sequence into the second,\n    use .get_opcodes():\n\n    >>> for opcode in s.get_opcodes():\n    ...     print \"%6s a[%d:%d] b[%d:%d]\" % opcode\n     equal a[0:8] b[0:8]\n    insert a[8:8] b[8:17]\n     equal a[8:29] b[17:38]\n\n    See the Differ class for a fancy human-friendly file differencer, which\n    uses SequenceMatcher both to compare sequences of lines, and to compare\n    sequences of characters within similar (near-matching) lines.\n\n    See also function get_close_matches() in this module, which shows how\n    simple code building on SequenceMatcher can be used to do useful work.\n\n    Timing:  Basic R-O is cubic time worst case and quadratic time expected\n    case.  SequenceMatcher is quadratic time for the worst case and has\n    expected-case behavior dependent in a complicated way on how many\n    elements the sequences have in common; best case time is linear.\n\n    Methods:\n\n    __init__(isjunk=None, a='', b='')\n        Construct a SequenceMatcher.\n\n    set_seqs(a, b)\n        Set the two sequences to be compared.\n\n    set_seq1(a)\n        Set the first sequence to be compared.\n\n    set_seq2(b)\n        Set the second sequence to be compared.\n\n    find_longest_match(alo, ahi, blo, bhi)\n        Find longest matching block in a[alo:ahi] and b[blo:bhi].\n\n    get_matching_blocks()\n        Return list of triples describing matching subsequences.\n\n    get_opcodes()\n        Return list of 5-tuples describing how to turn a into b.\n\n    ratio()\n        Return a measure of the sequences' similarity (float in [0,1]).\n\n    quick_ratio()\n        Return an upper bound on .ratio() relatively quickly.\n\n    real_quick_ratio()\n        Return an upper bound on ratio() very quickly.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 214: def __init__(self, isjunk=None, a='', b='', autojunk=True):
					πF.SetLineno(214)
					πTemp002 = make([]πg.Param, 5)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[1] = πg.Param{Name: "isjunk", Def: πTemp003}
					πTemp002[2] = πg.Param{Name: "a", Def: ß.ToObject()}
					πTemp002[3] = πg.Param{Name: "b", Def: ß.ToObject()}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßTrue); πE != nil {
						continue
					}
					πTemp002[4] = πg.Param{Name: "autojunk", Def: πTemp003}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/difflib.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µisjunk *πg.Object = πArgs[1]; _ = µisjunk
						var µa *πg.Object = πArgs[2]; _ = µa
						var µb *πg.Object = πArgs[3]; _ = µb
						var µautojunk *πg.Object = πArgs[4]; _ = µautojunk
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
							// line 215: """Construct a SequenceMatcher.
							πF.SetLineno(215)
							// line 278: self.isjunk = isjunk
							πF.SetLineno(278)
							if πE = πg.CheckLocal(πF, µisjunk, "isjunk"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µisjunk); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßisjunk, πTemp001); πE != nil {
								continue
							}
							// line 279: self.a = self.b = None
							πF.SetLineno(279)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßa, πTemp002); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßb, πTemp002); πE != nil {
								continue
							}
							// line 280: self.autojunk = autojunk
							πF.SetLineno(280)
							if πE = πg.CheckLocal(πF, µautojunk, "autojunk"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µautojunk); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßautojunk, πTemp001); πE != nil {
								continue
							}
							// line 281: self.set_seqs(a, b)
							πF.SetLineno(281)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp003[0] = µa
							if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
								continue
							}
							πTemp003[1] = µb
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßset_seqs, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
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
					// line 215: """Construct a SequenceMatcher.
					πF.SetLineno(215)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewStr("Construct a SequenceMatcher.\n\n        Optional arg isjunk is None (the default), or a one-argument\n        function that takes a sequence element and returns true iff the\n        element is junk.  None is equivalent to passing \"lambda x: 0\", i.e.\n        no elements are considered to be junk.  For example, pass\n            lambda x: x in \" \\t\"\n        if you're comparing lines as sequences of characters, and don't\n        want to synch up on blanks or hard tabs.\n\n        Optional arg a is the first of two sequences to be compared.  By\n        default, an empty string.  The elements of a must be hashable.  See\n        also .set_seqs() and .set_seq1().\n\n        Optional arg b is the second of two sequences to be compared.  By\n        default, an empty string.  The elements of b must be hashable. See\n        also .set_seqs() and .set_seq2().\n\n        Optional arg autojunk should be set to False to disable the\n        \"automatic junk heuristic\" that treats popular elements as junk\n        (see module documentation for more information).\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ß__init__); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp004, ß__doc__, πTemp003); πE != nil {
						continue
					}
					// line 283: def set_seqs(self, a, b):
					πF.SetLineno(283)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "a", Def: nil}
					πTemp002[2] = πg.Param{Name: "b", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("set_seqs", "/usr/lib/python2.7/difflib.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µa *πg.Object = πArgs[1]; _ = µa
						var µb *πg.Object = πArgs[2]; _ = µb
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
							// line 284: """Set the two sequences to be compared.
							πF.SetLineno(284)
							// line 292: self.set_seq1(a)
							πF.SetLineno(292)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp001[0] = µa
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßset_seq1, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 293: self.set_seq2(b)
							πF.SetLineno(293)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
								continue
							}
							πTemp001[0] = µb
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßset_seq2, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßset_seqs.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 284: """Set the two sequences to be compared.
					πF.SetLineno(284)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πg.NewStr("Set the two sequences to be compared.\n\n        >>> s = SequenceMatcher()\n        >>> s.set_seqs(\"abcd\", \"bcde\")\n        >>> s.ratio()\n        0.75\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßset_seqs); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp005, ß__doc__, πTemp004); πE != nil {
						continue
					}
					// line 295: def set_seq1(self, a):
					πF.SetLineno(295)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "a", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("set_seq1", "/usr/lib/python2.7/difflib.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µa *πg.Object = πArgs[1]; _ = µa
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
							// line 296: """Set the first sequence to be compared.
							πF.SetLineno(296)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßa, nil); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µa == πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 316: if a is self.a:
							πF.SetLineno(316)
						Label1:
							// line 317: return
							πF.SetLineno(317)
							πR = πg.None
							continue
							goto Label2
						Label2:
							// line 318: self.a = a
							πF.SetLineno(318)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µa); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßa, πTemp001); πE != nil {
								continue
							}
							// line 319: self.matching_blocks = self.opcodes = None
							πF.SetLineno(319)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßmatching_blocks, πTemp002); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßopcodes, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ßset_seq1.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 296: """Set the first sequence to be compared.
					πF.SetLineno(296)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πg.NewStr("Set the first sequence to be compared.\n\n        The second sequence to be compared is not changed.\n\n        >>> s = SequenceMatcher(None, \"abcd\", \"bcde\")\n        >>> s.ratio()\n        0.75\n        >>> s.set_seq1(\"bcde\")\n        >>> s.ratio()\n        1.0\n        >>>\n\n        SequenceMatcher computes and caches detailed information about the\n        second sequence, so if you want to compare one sequence S against\n        many sequences, use .set_seq2(S) once and call .set_seq1(x)\n        repeatedly for each of the other sequences.\n\n        See also set_seqs() and set_seq2().\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßset_seq1); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp006, ß__doc__, πTemp005); πE != nil {
						continue
					}
					// line 321: def set_seq2(self, b):
					πF.SetLineno(321)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "b", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("set_seq2", "/usr/lib/python2.7/difflib.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µb *πg.Object = πArgs[1]; _ = µb
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
							// line 322: """Set the second sequence to be compared.
							πF.SetLineno(322)
							if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßb, nil); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µb == πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 342: if b is self.b:
							πF.SetLineno(342)
						Label1:
							// line 343: return
							πF.SetLineno(343)
							πR = πg.None
							continue
							goto Label2
						Label2:
							// line 344: self.b = b
							πF.SetLineno(344)
							if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µb); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßb, πTemp001); πE != nil {
								continue
							}
							// line 345: self.matching_blocks = self.opcodes = None
							πF.SetLineno(345)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßmatching_blocks, πTemp002); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßopcodes, πTemp002); πE != nil {
								continue
							}
							// line 346: self.fullbcount = None
							πF.SetLineno(346)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßfullbcount, πTemp002); πE != nil {
								continue
							}
							// line 347: self.__chain_b()
							πF.SetLineno(347)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß__chain_b, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßset_seq2.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 322: """Set the second sequence to be compared.
					πF.SetLineno(322)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, πg.NewStr("Set the second sequence to be compared.\n\n        The first sequence to be compared is not changed.\n\n        >>> s = SequenceMatcher(None, \"abcd\", \"bcde\")\n        >>> s.ratio()\n        0.75\n        >>> s.set_seq2(\"abcd\")\n        >>> s.ratio()\n        1.0\n        >>>\n\n        SequenceMatcher computes and caches detailed information about the\n        second sequence, so if you want to compare one sequence S against\n        many sequences, use .set_seq2(S) once and call .set_seq1(x)\n        repeatedly for each of the other sequences.\n\n        See also set_seqs() and set_seq1().\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßset_seq2); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp007, ß__doc__, πTemp006); πE != nil {
						continue
					}
					// line 366: def __chain_b(self):
					πF.SetLineno(366)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("__chain_b", "/usr/lib/python2.7/difflib.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µb *πg.Object = πg.UnboundLocal; _ = µb
						var µb2j *πg.Object = πg.UnboundLocal; _ = µb2j
						var µi *πg.Object = πg.UnboundLocal; _ = µi
						var µelt *πg.Object = πg.UnboundLocal; _ = µelt
						var µindices *πg.Object = πg.UnboundLocal; _ = µindices
						var µjunk *πg.Object = πg.UnboundLocal; _ = µjunk
						var µisjunk *πg.Object = πg.UnboundLocal; _ = µisjunk
						var µpopular *πg.Object = πg.UnboundLocal; _ = µpopular
						var µn *πg.Object = πg.UnboundLocal; _ = µn
						var µntest *πg.Object = πg.UnboundLocal; _ = µntest
						var µidxs *πg.Object = πg.UnboundLocal; _ = µidxs
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Dict
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
						var πTemp009 []*πg.Object
						_ = πTemp009
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							case 2: goto Label2
							case 6: goto Label6
							case 7: goto Label7
							case 14: goto Label14
							case 15: goto Label15
							default: panic("unexpected function state")
							}
							// line 377: b = self.b
							πF.SetLineno(377)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßb, nil); πE != nil {
								continue
							}
							µb = πTemp001
							// line 378: self.b2j = b2j = {}
							πF.SetLineno(378)
							πTemp002 = πg.NewDict()
							πTemp001 = πTemp002.ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßb2j, πTemp003); πE != nil {
								continue
							}
							µb2j = πTemp001
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
								continue
							}
							πTemp004[0] = µb
							if πTemp003, πE = πg.ResolveGlobal(πF, ßenumerate); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
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
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp008}}}, πTemp003); πE != nil {
									continue
								}
								µi = πTemp005
								µelt = πTemp008
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 381: indices = setdefault(b2j, elt, [])
							πF.SetLineno(381)
							πTemp004 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µb2j, "b2j"); πE != nil {
								continue
							}
							πTemp004[0] = µb2j
							if πE = πg.CheckLocal(πF, µelt, "elt"); πE != nil {
								continue
							}
							πTemp004[1] = µelt
							πTemp009 = make([]*πg.Object, 0)
							πTemp003 = πg.NewList(πTemp009...).ToObject()
							πTemp004[2] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßsetdefault); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µindices = πTemp005
							// line 383: indices.append(i)
							πF.SetLineno(383)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp004[0] = µi
							if πE = πg.CheckLocal(πF, µindices, "indices"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µindices, ßappend, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 386: junk = set()
							πF.SetLineno(386)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßset); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µjunk = πTemp003
							// line 387: isjunk = self.isjunk
							πF.SetLineno(387)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßisjunk, nil); πE != nil {
								continue
							}
							µisjunk = πTemp001
							if πE = πg.CheckLocal(πF, µisjunk, "isjunk"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, µisjunk); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label4
							}
							goto Label5
							// line 388: if isjunk:
							πF.SetLineno(388)
						Label4:
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µb2j, "b2j"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µb2j, ßkeys, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp005
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp001, πE = πg.Iter(πF, πTemp005); πE != nil {
								continue
							}
							πF.PushCheckpoint(7)
							πTemp006 = false
						Label6:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp006 {
								πF.PopCheckpoint()
								goto Label8
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
								µelt = πTemp003
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(6)            
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µelt, "elt"); πE != nil {
								continue
							}
							πTemp004[0] = µelt
							if πE = πg.CheckLocal(πF, µisjunk, "isjunk"); πE != nil {
								continue
							}
							if πTemp003, πE = µisjunk.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label9
							}
							goto Label10
							// line 390: if isjunk(elt):
							πF.SetLineno(390)
						Label9:
							// line 391: junk.add(elt)
							πF.SetLineno(391)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µelt, "elt"); πE != nil {
								continue
							}
							πTemp004[0] = µelt
							if πE = πg.CheckLocal(πF, µjunk, "junk"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µjunk, ßadd, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 392: del b2j[elt]
							πF.SetLineno(392)
							if πE = πg.CheckLocal(πF, µb2j, "b2j"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µelt, "elt"); πE != nil {
								continue
							}
							πTemp003 = µelt
							if πE = πg.DelItem(πF, µb2j, πTemp003); πE != nil {
								continue
							}
							goto Label10
						Label10:
							continue
						Label7:
							if πE != nil || πR != nil {
								continue
							}
						Label8:
							goto Label5
						Label5:
							// line 395: popular = set()
							πF.SetLineno(395)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßset); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µpopular = πTemp003
							// line 396: n = len(b)
							πF.SetLineno(396)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
								continue
							}
							πTemp004[0] = µb
							if πTemp001, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µn = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßautojunk, nil); πE != nil {
								continue
							}
							πTemp001 = πTemp003
							if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp006 {
								goto Label11
							}
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GE(πF, µn, πg.NewInt(200).ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp003
						Label11:
							if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label12
							}
							goto Label13
							// line 397: if self.autojunk and n >= 200:
							πF.SetLineno(397)
						Label12:
							// line 398: ntest = n // 100 + 1
							πF.SetLineno(398)
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.FloorDiv(πF, µn, πg.NewInt(100).ToObject()); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp003, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µntest = πTemp001
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µb2j, "b2j"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µb2j, ßitems, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp005
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp001, πE = πg.Iter(πF, πTemp005); πE != nil {
								continue
							}
							πF.PushCheckpoint(15)
							πTemp006 = false
						Label14:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp006 {
								πF.PopCheckpoint()
								goto Label16
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
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp008}}}, πTemp003); πE != nil {
									continue
								}
								µelt = πTemp005
								µidxs = πTemp008
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(14)            
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µidxs, "idxs"); πE != nil {
								continue
							}
							πTemp004[0] = µidxs
							if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πE = πg.CheckLocal(πF, µntest, "ntest"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GT(πF, πTemp008, µntest); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label17
							}
							goto Label18
							// line 400: if len(idxs) > ntest:
							πF.SetLineno(400)
						Label17:
							// line 401: popular.add(elt)
							πF.SetLineno(401)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µelt, "elt"); πE != nil {
								continue
							}
							πTemp004[0] = µelt
							if πE = πg.CheckLocal(πF, µpopular, "popular"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µpopular, ßadd, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 402: del b2j[elt]
							πF.SetLineno(402)
							if πE = πg.CheckLocal(πF, µb2j, "b2j"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µelt, "elt"); πE != nil {
								continue
							}
							πTemp003 = µelt
							if πE = πg.DelItem(πF, µb2j, πTemp003); πE != nil {
								continue
							}
							goto Label18
						Label18:
							continue
						Label15:
							if πE != nil || πR != nil {
								continue
							}
						Label16:
							goto Label13
						Label13:
							// line 408: self.isbjunk = junk.__contains__
							πF.SetLineno(408)
							if πE = πg.CheckLocal(πF, µjunk, "junk"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µjunk, ß__contains__, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßisbjunk, πTemp003); πE != nil {
								continue
							}
							// line 409: self.isbpopular = popular.__contains__
							πF.SetLineno(409)
							if πE = πg.CheckLocal(πF, µpopular, "popular"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µpopular, ß__contains__, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßisbpopular, πTemp003); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__chain_b.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 411: def find_longest_match(self, alo, ahi, blo, bhi):
					πF.SetLineno(411)
					πTemp002 = make([]πg.Param, 5)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "alo", Def: nil}
					πTemp002[2] = πg.Param{Name: "ahi", Def: nil}
					πTemp002[3] = πg.Param{Name: "blo", Def: nil}
					πTemp002[4] = πg.Param{Name: "bhi", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("find_longest_match", "/usr/lib/python2.7/difflib.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µalo *πg.Object = πArgs[1]; _ = µalo
						var µahi *πg.Object = πArgs[2]; _ = µahi
						var µblo *πg.Object = πArgs[3]; _ = µblo
						var µbhi *πg.Object = πArgs[4]; _ = µbhi
						var µa *πg.Object = πg.UnboundLocal; _ = µa
						var µb *πg.Object = πg.UnboundLocal; _ = µb
						var µb2j *πg.Object = πg.UnboundLocal; _ = µb2j
						var µisbjunk *πg.Object = πg.UnboundLocal; _ = µisbjunk
						var µbesti *πg.Object = πg.UnboundLocal; _ = µbesti
						var µbestj *πg.Object = πg.UnboundLocal; _ = µbestj
						var µbestsize *πg.Object = πg.UnboundLocal; _ = µbestsize
						var µj2len *πg.Object = πg.UnboundLocal; _ = µj2len
						var µnothing *πg.Object = πg.UnboundLocal; _ = µnothing
						var µi *πg.Object = πg.UnboundLocal; _ = µi
						var µj2lenget *πg.Object = πg.UnboundLocal; _ = µj2lenget
						var µnewj2len *πg.Object = πg.UnboundLocal; _ = µnewj2len
						var µj *πg.Object = πg.UnboundLocal; _ = µj
						var µk *πg.Object = πg.UnboundLocal; _ = µk
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
						var πTemp006 *πg.Dict
						_ = πTemp006
						var πTemp007 []*πg.Object
						_ = πTemp007
						var πTemp008 bool
						_ = πTemp008
						var πTemp009 bool
						_ = πTemp009
						var πTemp010 bool
						_ = πTemp010
						var πTemp011 *πg.Object
						_ = πTemp011
						var πTemp012 bool
						_ = πTemp012
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							case 2: goto Label2
							case 4: goto Label4
							case 5: goto Label5
							case 13: goto Label13
							case 14: goto Label14
							case 17: goto Label17
							case 18: goto Label18
							case 21: goto Label21
							case 22: goto Label22
							case 25: goto Label25
							case 26: goto Label26
							default: panic("unexpected function state")
							}
							// line 412: """Find longest matching block in a[alo:ahi] and b[blo:bhi].
							πF.SetLineno(412)
							// line 467: a, b, b2j, isbjunk = self.a, self.b, self.b2j, self.isbjunk
							πF.SetLineno(467)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßa, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßb, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßb2j, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßisbjunk, nil); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple4(πTemp002, πTemp003, πTemp004, πTemp005).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}}}, πTemp001); πE != nil {
								continue
							}
							µa = πTemp002
							µb = πTemp003
							µb2j = πTemp004
							µisbjunk = πTemp005
							// line 468: besti, bestj, bestsize = alo, blo, 0
							πF.SetLineno(468)
							if πE = πg.CheckLocal(πF, µalo, "alo"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µblo, "blo"); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple3(µalo, µblo, πg.NewInt(0).ToObject()).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp004}}}, πTemp001); πE != nil {
								continue
							}
							µbesti = πTemp002
							µbestj = πTemp003
							µbestsize = πTemp004
							// line 472: j2len = {}
							πF.SetLineno(472)
							πTemp006 = πg.NewDict()
							πTemp001 = πTemp006.ToObject()
							µj2len = πTemp001
							// line 473: nothing = []
							πF.SetLineno(473)
							πTemp007 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp007...).ToObject()
							µnothing = πTemp001
							πTemp007 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µalo, "alo"); πE != nil {
								continue
							}
							πTemp007[0] = µalo
							if πE = πg.CheckLocal(πF, µahi, "ahi"); πE != nil {
								continue
							}
							πTemp007[1] = µahi
							if πTemp002, πE = πg.ResolveGlobal(πF, ßxrange); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							if πTemp001, πE = πg.Iter(πF, πTemp003); πE != nil {
								continue
							}
							πF.PushCheckpoint(2)
							πTemp008 = false
						Label1:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp008 {
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
								πTemp009 = !isStop
							} else {
								πTemp009 = true
								µi = πTemp002
							}
							if πE != nil || !πTemp009 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 477: j2lenget = j2len.get
							πF.SetLineno(477)
							if πE = πg.CheckLocal(πF, µj2len, "j2len"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µj2len, ßget, nil); πE != nil {
								continue
							}
							µj2lenget = πTemp002
							// line 478: newj2len = {}
							πF.SetLineno(478)
							πTemp006 = πg.NewDict()
							πTemp002 = πTemp006.ToObject()
							µnewj2len = πTemp002
							πTemp007 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp003 = µi
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µa, πTemp003); πE != nil {
								continue
							}
							πTemp007[0] = πTemp004
							if πE = πg.CheckLocal(πF, µnothing, "nothing"); πE != nil {
								continue
							}
							πTemp007[1] = µnothing
							if πE = πg.CheckLocal(πF, µb2j, "b2j"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µb2j, ßget, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							if πTemp002, πE = πg.Iter(πF, πTemp004); πE != nil {
								continue
							}
							πF.PushCheckpoint(5)
							πTemp009 = false
						Label4:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp009 {
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
								πTemp010 = !isStop
							} else {
								πTemp010 = true
								µj = πTemp003
							}
							if πE != nil || !πTemp010 {
								continue
							}
							πF.PushCheckpoint(4)            
							if πE = πg.CheckLocal(πF, µj, "j"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µblo, "blo"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.LT(πF, µj, µblo); πE != nil {
								continue
							}
							if πTemp010, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp010 {
								goto Label7
							}
							goto Label8
							// line 481: if j < blo:
							πF.SetLineno(481)
						Label7:
							// line 482: continue
							πF.SetLineno(482)
							continue
							goto Label8
						Label8:
							if πE = πg.CheckLocal(πF, µj, "j"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µbhi, "bhi"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GE(πF, µj, µbhi); πE != nil {
								continue
							}
							if πTemp010, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp010 {
								goto Label9
							}
							goto Label10
							// line 483: if j >= bhi:
							πF.SetLineno(483)
						Label9:
							// line 484: break
							πF.SetLineno(484)
							πTemp009 = true
							continue
							goto Label10
						Label10:
							// line 485: k = newj2len[j] = j2lenget(j-1, 0) + 1
							πF.SetLineno(485)
							πTemp007 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µj, "j"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Sub(πF, µj, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp007[0] = πTemp004
							πTemp007[1] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µj2lenget, "j2lenget"); πE != nil {
								continue
							}
							if πTemp004, πE = µj2lenget.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							if πTemp003, πE = πg.Add(πF, πTemp004, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µk = πTemp003
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnewj2len, "newj2len"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µj, "j"); πE != nil {
								continue
							}
							πTemp005 = µj
							if πE = πg.SetItem(πF, µnewj2len, πTemp005, πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µbestsize, "bestsize"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GT(πF, µk, µbestsize); πE != nil {
								continue
							}
							if πTemp010, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp010 {
								goto Label11
							}
							goto Label12
							// line 486: if k > bestsize:
							πF.SetLineno(486)
						Label11:
							// line 487: besti, bestj, bestsize = i-k+1, j-k+1, k
							πF.SetLineno(487)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Sub(πF, µi, µk); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Add(πF, πTemp005, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µj, "j"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
								continue
							}
							if πTemp011, πE = πg.Sub(πF, µj, µk); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Add(πF, πTemp011, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple3(πTemp004, πTemp005, µk).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp011}}}, πTemp003); πE != nil {
								continue
							}
							µbesti = πTemp004
							µbestj = πTemp005
							µbestsize = πTemp011
							goto Label12
						Label12:
							continue
						Label5:
							if πE != nil || πR != nil {
								continue
							}
						Label6:
							// line 488: j2len = newj2len
							πF.SetLineno(488)
							if πE = πg.CheckLocal(πF, µnewj2len, "newj2len"); πE != nil {
								continue
							}
							µj2len = µnewj2len
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 494: while besti > alo and bestj > blo and \
							πF.SetLineno(494)
							πF.PushCheckpoint(14)
							πTemp008 = false
						Label13:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp008 {
								πF.PopCheckpoint()
								goto Label15
							}
							if πE = πg.CheckLocal(πF, µbesti, "besti"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µalo, "alo"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GT(πF, µbesti, µalo); πE != nil {
								continue
							}
							πTemp001 = πTemp002
							if πTemp010, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp010 {
								goto Label16
							}
							if πE = πg.CheckLocal(πF, µbestj, "bestj"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µblo, "blo"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GT(πF, µbestj, µblo); πE != nil {
								continue
							}
							πTemp001 = πTemp002
							if πTemp010, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp010 {
								goto Label16
							}
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µbestj, "bestj"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Sub(πF, µbestj, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp003 = πTemp004
							if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µb, πTemp003); πE != nil {
								continue
							}
							πTemp007[0] = πTemp004
							if πE = πg.CheckLocal(πF, µisbjunk, "isbjunk"); πE != nil {
								continue
							}
							if πTemp003, πE = µisbjunk.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							if πTemp012, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp012).ToObject()
							πTemp001 = πTemp002
							if πTemp010, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp010 {
								goto Label16
							}
							if πE = πg.CheckLocal(πF, µbesti, "besti"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Sub(πF, µbesti, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp003 = πTemp004
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µa, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µbestj, "bestj"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Sub(πF, µbestj, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp003 = πTemp005
							if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µb, πTemp003); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, πTemp004, πTemp005); πE != nil {
								continue
							}
							πTemp001 = πTemp002
						Label16:
							if πTemp009, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πE != nil || !πTemp009 {
								continue
							}
							πF.PushCheckpoint(13)            
							// line 497: besti, bestj, bestsize = besti-1, bestj-1, bestsize+1
							πF.SetLineno(497)
							if πE = πg.CheckLocal(πF, µbesti, "besti"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Sub(πF, µbesti, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µbestj, "bestj"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Sub(πF, µbestj, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µbestsize, "bestsize"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Add(πF, µbestsize, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple3(πTemp002, πTemp003, πTemp004).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp004}}}, πTemp001); πE != nil {
								continue
							}
							µbesti = πTemp002
							µbestj = πTemp003
							µbestsize = πTemp004
							continue
						Label14:
							if πE != nil || πR != nil {
								continue
							}
						Label15:
							// line 498: while besti+bestsize < ahi and bestj+bestsize < bhi and \
							πF.SetLineno(498)
							πF.PushCheckpoint(18)
							πTemp008 = false
						Label17:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp008 {
								πF.PopCheckpoint()
								goto Label19
							}
							if πE = πg.CheckLocal(πF, µbesti, "besti"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µbestsize, "bestsize"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, µbesti, µbestsize); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µahi, "ahi"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.LT(πF, πTemp003, µahi); πE != nil {
								continue
							}
							πTemp001 = πTemp002
							if πTemp010, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp010 {
								goto Label20
							}
							if πE = πg.CheckLocal(πF, µbestj, "bestj"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µbestsize, "bestsize"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, µbestj, µbestsize); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µbhi, "bhi"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.LT(πF, πTemp003, µbhi); πE != nil {
								continue
							}
							πTemp001 = πTemp002
							if πTemp010, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp010 {
								goto Label20
							}
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µbestj, "bestj"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µbestsize, "bestsize"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Add(πF, µbestj, µbestsize); πE != nil {
								continue
							}
							πTemp003 = πTemp004
							if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µb, πTemp003); πE != nil {
								continue
							}
							πTemp007[0] = πTemp004
							if πE = πg.CheckLocal(πF, µisbjunk, "isbjunk"); πE != nil {
								continue
							}
							if πTemp003, πE = µisbjunk.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							if πTemp012, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp012).ToObject()
							πTemp001 = πTemp002
							if πTemp010, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp010 {
								goto Label20
							}
							if πE = πg.CheckLocal(πF, µbesti, "besti"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µbestsize, "bestsize"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Add(πF, µbesti, µbestsize); πE != nil {
								continue
							}
							πTemp003 = πTemp004
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µa, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µbestj, "bestj"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µbestsize, "bestsize"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Add(πF, µbestj, µbestsize); πE != nil {
								continue
							}
							πTemp003 = πTemp005
							if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µb, πTemp003); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, πTemp004, πTemp005); πE != nil {
								continue
							}
							πTemp001 = πTemp002
						Label20:
							if πTemp009, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πE != nil || !πTemp009 {
								continue
							}
							πF.PushCheckpoint(17)            
							// line 501: bestsize += 1
							πF.SetLineno(501)
							if πE = πg.CheckLocal(πF, µbestsize, "bestsize"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.IAdd(πF, µbestsize, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µbestsize = πTemp001
							continue
						Label18:
							if πE != nil || πR != nil {
								continue
							}
						Label19:
							// line 510: while besti > alo and bestj > blo and \
							πF.SetLineno(510)
							πF.PushCheckpoint(22)
							πTemp008 = false
						Label21:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp008 {
								πF.PopCheckpoint()
								goto Label23
							}
							if πE = πg.CheckLocal(πF, µbesti, "besti"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µalo, "alo"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GT(πF, µbesti, µalo); πE != nil {
								continue
							}
							πTemp001 = πTemp002
							if πTemp010, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp010 {
								goto Label24
							}
							if πE = πg.CheckLocal(πF, µbestj, "bestj"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µblo, "blo"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GT(πF, µbestj, µblo); πE != nil {
								continue
							}
							πTemp001 = πTemp002
							if πTemp010, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp010 {
								goto Label24
							}
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µbestj, "bestj"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Sub(πF, µbestj, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp003
							if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µb, πTemp002); πE != nil {
								continue
							}
							πTemp007[0] = πTemp003
							if πE = πg.CheckLocal(πF, µisbjunk, "isbjunk"); πE != nil {
								continue
							}
							if πTemp002, πE = µisbjunk.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp001 = πTemp002
							if πTemp010, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp010 {
								goto Label24
							}
							if πE = πg.CheckLocal(πF, µbesti, "besti"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Sub(πF, µbesti, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp003 = πTemp004
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µa, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µbestj, "bestj"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Sub(πF, µbestj, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp003 = πTemp005
							if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µb, πTemp003); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, πTemp004, πTemp005); πE != nil {
								continue
							}
							πTemp001 = πTemp002
						Label24:
							if πTemp009, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πE != nil || !πTemp009 {
								continue
							}
							πF.PushCheckpoint(21)            
							// line 513: besti, bestj, bestsize = besti-1, bestj-1, bestsize+1
							πF.SetLineno(513)
							if πE = πg.CheckLocal(πF, µbesti, "besti"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Sub(πF, µbesti, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µbestj, "bestj"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Sub(πF, µbestj, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µbestsize, "bestsize"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Add(πF, µbestsize, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple3(πTemp002, πTemp003, πTemp004).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp004}}}, πTemp001); πE != nil {
								continue
							}
							µbesti = πTemp002
							µbestj = πTemp003
							µbestsize = πTemp004
							continue
						Label22:
							if πE != nil || πR != nil {
								continue
							}
						Label23:
							// line 514: while besti+bestsize < ahi and bestj+bestsize < bhi and \
							πF.SetLineno(514)
							πF.PushCheckpoint(26)
							πTemp008 = false
						Label25:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp008 {
								πF.PopCheckpoint()
								goto Label27
							}
							if πE = πg.CheckLocal(πF, µbesti, "besti"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µbestsize, "bestsize"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, µbesti, µbestsize); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µahi, "ahi"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.LT(πF, πTemp003, µahi); πE != nil {
								continue
							}
							πTemp001 = πTemp002
							if πTemp010, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp010 {
								goto Label28
							}
							if πE = πg.CheckLocal(πF, µbestj, "bestj"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µbestsize, "bestsize"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, µbestj, µbestsize); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µbhi, "bhi"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.LT(πF, πTemp003, µbhi); πE != nil {
								continue
							}
							πTemp001 = πTemp002
							if πTemp010, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp010 {
								goto Label28
							}
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µbestj, "bestj"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µbestsize, "bestsize"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, µbestj, µbestsize); πE != nil {
								continue
							}
							πTemp002 = πTemp003
							if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µb, πTemp002); πE != nil {
								continue
							}
							πTemp007[0] = πTemp003
							if πE = πg.CheckLocal(πF, µisbjunk, "isbjunk"); πE != nil {
								continue
							}
							if πTemp002, πE = µisbjunk.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp001 = πTemp002
							if πTemp010, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp010 {
								goto Label28
							}
							if πE = πg.CheckLocal(πF, µbesti, "besti"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µbestsize, "bestsize"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Add(πF, µbesti, µbestsize); πE != nil {
								continue
							}
							πTemp003 = πTemp004
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µa, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µbestj, "bestj"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µbestsize, "bestsize"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Add(πF, µbestj, µbestsize); πE != nil {
								continue
							}
							πTemp003 = πTemp005
							if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µb, πTemp003); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, πTemp004, πTemp005); πE != nil {
								continue
							}
							πTemp001 = πTemp002
						Label28:
							if πTemp009, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πE != nil || !πTemp009 {
								continue
							}
							πF.PushCheckpoint(25)            
							// line 517: bestsize = bestsize + 1
							πF.SetLineno(517)
							if πE = πg.CheckLocal(πF, µbestsize, "bestsize"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, µbestsize, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µbestsize = πTemp001
							continue
						Label26:
							if πE != nil || πR != nil {
								continue
							}
						Label27:
							// line 519: return Match(besti, bestj, bestsize)
							πF.SetLineno(519)
							πTemp007 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µbesti, "besti"); πE != nil {
								continue
							}
							πTemp007[0] = µbesti
							if πE = πg.CheckLocal(πF, µbestj, "bestj"); πE != nil {
								continue
							}
							πTemp007[1] = µbestj
							if πE = πg.CheckLocal(πF, µbestsize, "bestsize"); πE != nil {
								continue
							}
							πTemp007[2] = µbestsize
							if πTemp001, πE = πg.ResolveGlobal(πF, ßMatch); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
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
					if πE = πClass.SetItem(πF, ßfind_longest_match.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 412: """Find longest matching block in a[alo:ahi] and b[blo:bhi].
					πF.SetLineno(412)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp008}, πg.NewStr("Find longest matching block in a[alo:ahi] and b[blo:bhi].\n\n        If isjunk is not defined:\n\n        Return (i,j,k) such that a[i:i+k] is equal to b[j:j+k], where\n            alo <= i <= i+k <= ahi\n            blo <= j <= j+k <= bhi\n        and for all (i',j',k') meeting those conditions,\n            k >= k'\n            i <= i'\n            and if i == i', j <= j'\n\n        In other words, of all maximal matching blocks, return one that\n        starts earliest in a, and of all those maximal matching blocks that\n        start earliest in a, return the one that starts earliest in b.\n\n        >>> s = SequenceMatcher(None, \" abcd\", \"abcd abcd\")\n        >>> s.find_longest_match(0, 5, 0, 9)\n        Match(a=0, b=4, size=5)\n\n        If isjunk is defined, first the longest matching block is\n        determined as above, but with the additional restriction that no\n        junk element appears in the block.  Then that block is extended as\n        far as possible by matching (only) junk elements on both sides.  So\n        the resulting block never matches on junk except as identical junk\n        happens to be adjacent to an \"interesting\" match.\n\n        Here's the same example as before, but considering blanks to be\n        junk.  That prevents \" abcd\" from matching the \" abcd\" at the tail\n        end of the second sequence directly.  Instead only the \"abcd\" can\n        match, and matches the leftmost \"abcd\" in the second sequence:\n\n        >>> s = SequenceMatcher(lambda x: x==\" \", \" abcd\", \"abcd abcd\")\n        >>> s.find_longest_match(0, 5, 0, 9)\n        Match(a=1, b=0, size=4)\n\n        If no blocks match, return (alo, blo, 0).\n\n        >>> s = SequenceMatcher(None, \"ab\", \"c\")\n        >>> s.find_longest_match(0, 2, 0, 1)\n        Match(a=0, b=0, size=0)\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßfind_longest_match); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp009, ß__doc__, πTemp008); πE != nil {
						continue
					}
					// line 521: def get_matching_blocks(self):
					πF.SetLineno(521)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("get_matching_blocks", "/usr/lib/python2.7/difflib.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µla *πg.Object = πg.UnboundLocal; _ = µla
						var µlb *πg.Object = πg.UnboundLocal; _ = µlb
						var µqueue *πg.Object = πg.UnboundLocal; _ = µqueue
						var µmatching_blocks *πg.Object = πg.UnboundLocal; _ = µmatching_blocks
						var µalo *πg.Object = πg.UnboundLocal; _ = µalo
						var µahi *πg.Object = πg.UnboundLocal; _ = µahi
						var µblo *πg.Object = πg.UnboundLocal; _ = µblo
						var µbhi *πg.Object = πg.UnboundLocal; _ = µbhi
						var µi *πg.Object = πg.UnboundLocal; _ = µi
						var µj *πg.Object = πg.UnboundLocal; _ = µj
						var µk *πg.Object = πg.UnboundLocal; _ = µk
						var µx *πg.Object = πg.UnboundLocal; _ = µx
						var µi1 *πg.Object = πg.UnboundLocal; _ = µi1
						var µj1 *πg.Object = πg.UnboundLocal; _ = µj1
						var µk1 *πg.Object = πg.UnboundLocal; _ = µk1
						var µnon_adjacent *πg.Object = πg.UnboundLocal; _ = µnon_adjacent
						var µi2 *πg.Object = πg.UnboundLocal; _ = µi2
						var µj2 *πg.Object = πg.UnboundLocal; _ = µj2
						var µk2 *πg.Object = πg.UnboundLocal; _ = µk2
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
							case 14: goto Label14
							case 15: goto Label15
							default: panic("unexpected function state")
							}
							// line 522: """Return list of triples describing matching subsequences.
							πF.SetLineno(522)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßmatching_blocks, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp002 != πTemp003).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 540: if self.matching_blocks is not None:
							πF.SetLineno(540)
						Label1:
							// line 541: return self.matching_blocks
							πF.SetLineno(541)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßmatching_blocks, nil); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label2
						Label2:
							// line 542: la, lb = len(self.a), len(self.b)
							πF.SetLineno(542)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßa, nil); πE != nil {
								continue
							}
							πTemp005[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßb, nil); πE != nil {
								continue
							}
							πTemp005[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp001 = πg.NewTuple2(πTemp003, πTemp006).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp003}}}, πTemp001); πE != nil {
								continue
							}
							µla = πTemp002
							µlb = πTemp003
							// line 550: queue = [(0, la, 0, lb)]
							πF.SetLineno(550)
							πTemp005 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µla, "la"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µlb, "lb"); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple4(πg.NewInt(0).ToObject(), µla, πg.NewInt(0).ToObject(), µlb).ToObject()
							πTemp005[0] = πTemp001
							πTemp001 = πg.NewList(πTemp005...).ToObject()
							µqueue = πTemp001
							// line 551: matching_blocks = []
							πF.SetLineno(551)
							πTemp005 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp005...).ToObject()
							µmatching_blocks = πTemp001
							// line 552: while queue:
							πF.SetLineno(552)
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
							if πE = πg.CheckLocal(πF, µqueue, "queue"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, µqueue); πE != nil {
								continue
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(3)            
							// line 553: alo, ahi, blo, bhi = queue.pop()
							πF.SetLineno(553)
							if πE = πg.CheckLocal(πF, µqueue, "queue"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µqueue, ßpop, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp006}, πg.TieTarget{Target: &πTemp008}}}, πTemp002); πE != nil {
								continue
							}
							µalo = πTemp001
							µahi = πTemp003
							µblo = πTemp006
							µbhi = πTemp008
							// line 554: i, j, k = x = self.find_longest_match(alo, ahi, blo, bhi)
							πF.SetLineno(554)
							πTemp005 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µalo, "alo"); πE != nil {
								continue
							}
							πTemp005[0] = µalo
							if πE = πg.CheckLocal(πF, µahi, "ahi"); πE != nil {
								continue
							}
							πTemp005[1] = µahi
							if πE = πg.CheckLocal(πF, µblo, "blo"); πE != nil {
								continue
							}
							πTemp005[2] = µblo
							if πE = πg.CheckLocal(πF, µbhi, "bhi"); πE != nil {
								continue
							}
							πTemp005[3] = µbhi
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßfind_longest_match, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp006}}}, πTemp002); πE != nil {
								continue
							}
							µi = πTemp001
							µj = πTemp003
							µk = πTemp006
							µx = πTemp002
							if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, µk); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label6
							}
							goto Label7
							// line 558: if k:   # if k is 0, there was no matching block
							πF.SetLineno(558)
						Label6:
							// line 559: matching_blocks.append(x)
							πF.SetLineno(559)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp005[0] = µx
							if πE = πg.CheckLocal(πF, µmatching_blocks, "matching_blocks"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µmatching_blocks, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πE = πg.CheckLocal(πF, µalo, "alo"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.LT(πF, µalo, µi); πE != nil {
								continue
							}
							πTemp001 = πTemp002
							if πTemp007, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp007 {
								goto Label8
							}
							if πE = πg.CheckLocal(πF, µblo, "blo"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µj, "j"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.LT(πF, µblo, µj); πE != nil {
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
							goto Label10
							// line 560: if alo < i and blo < j:
							πF.SetLineno(560)
						Label9:
							// line 561: queue.append((alo, i, blo, j))
							πF.SetLineno(561)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µalo, "alo"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µblo, "blo"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µj, "j"); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple4(µalo, µi, µblo, µj).ToObject()
							πTemp005[0] = πTemp001
							if πE = πg.CheckLocal(πF, µqueue, "queue"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µqueue, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							goto Label10
						Label10:
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, µi, µk); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µahi, "ahi"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.LT(πF, πTemp003, µahi); πE != nil {
								continue
							}
							πTemp001 = πTemp002
							if πTemp007, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp007 {
								goto Label11
							}
							if πE = πg.CheckLocal(πF, µj, "j"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, µj, µk); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µbhi, "bhi"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.LT(πF, πTemp003, µbhi); πE != nil {
								continue
							}
							πTemp001 = πTemp002
						Label11:
							if πTemp007, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label12
							}
							goto Label13
							// line 562: if i+k < ahi and j+k < bhi:
							πF.SetLineno(562)
						Label12:
							// line 563: queue.append((i+k, ahi, j+k, bhi))
							πF.SetLineno(563)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, µi, µk); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µahi, "ahi"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µj, "j"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, µj, µk); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µbhi, "bhi"); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple4(πTemp002, µahi, πTemp003, µbhi).ToObject()
							πTemp005[0] = πTemp001
							if πE = πg.CheckLocal(πF, µqueue, "queue"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µqueue, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							goto Label13
						Label13:
							goto Label7
						Label7:
							continue
						Label4:
							if πE != nil || πR != nil {
								continue
							}
						Label5:
							// line 564: matching_blocks.sort()
							πF.SetLineno(564)
							if πE = πg.CheckLocal(πF, µmatching_blocks, "matching_blocks"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µmatching_blocks, ßsort, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 569: i1 = j1 = k1 = 0
							πF.SetLineno(569)
							µi1 = πg.NewInt(0).ToObject()
							µj1 = πg.NewInt(0).ToObject()
							µk1 = πg.NewInt(0).ToObject()
							// line 570: non_adjacent = []
							πF.SetLineno(570)
							πTemp005 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp005...).ToObject()
							µnon_adjacent = πTemp001
							if πE = πg.CheckLocal(πF, µmatching_blocks, "matching_blocks"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, µmatching_blocks); πE != nil {
								continue
							}
							πF.PushCheckpoint(15)
							πTemp004 = false
						Label14:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp004 {
								πF.PopCheckpoint()
								goto Label16
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
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp006}, πg.TieTarget{Target: &πTemp008}}}, πTemp002); πE != nil {
									continue
								}
								µi2 = πTemp003
								µj2 = πTemp006
								µk2 = πTemp008
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(14)            
							if πE = πg.CheckLocal(πF, µi1, "i1"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µk1, "k1"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Add(πF, µi1, µk1); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi2, "i2"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Eq(πF, πTemp006, µi2); πE != nil {
								continue
							}
							πTemp002 = πTemp003
							if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp007 {
								goto Label17
							}
							if πE = πg.CheckLocal(πF, µj1, "j1"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µk1, "k1"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Add(πF, µj1, µk1); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µj2, "j2"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Eq(πF, πTemp006, µj2); πE != nil {
								continue
							}
							πTemp002 = πTemp003
						Label17:
							if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label18
							}
							goto Label19
							// line 573: if i1 + k1 == i2 and j1 + k1 == j2:
							πF.SetLineno(573)
						Label18:
							// line 577: k1 += k2
							πF.SetLineno(577)
							if πE = πg.CheckLocal(πF, µk1, "k1"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µk2, "k2"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IAdd(πF, µk1, µk2); πE != nil {
								continue
							}
							µk1 = πTemp002
							goto Label20
						Label19:
							if πE = πg.CheckLocal(πF, µk1, "k1"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, µk1); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label21
							}
							goto Label22
							// line 582: if k1:
							πF.SetLineno(582)
						Label21:
							// line 583: non_adjacent.append((i1, j1, k1))
							πF.SetLineno(583)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µi1, "i1"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µj1, "j1"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µk1, "k1"); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple3(µi1, µj1, µk1).ToObject()
							πTemp005[0] = πTemp002
							if πE = πg.CheckLocal(πF, µnon_adjacent, "non_adjacent"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µnon_adjacent, ßappend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							goto Label22
						Label22:
							// line 584: i1, j1, k1 = i2, j2, k2
							πF.SetLineno(584)
							if πE = πg.CheckLocal(πF, µi2, "i2"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µj2, "j2"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µk2, "k2"); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple3(µi2, µj2, µk2).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp006}, πg.TieTarget{Target: &πTemp008}}}, πTemp002); πE != nil {
								continue
							}
							µi1 = πTemp003
							µj1 = πTemp006
							µk1 = πTemp008
							goto Label20
						Label20:
							continue
						Label15:
							if πE != nil || πR != nil {
								continue
							}
						Label16:
							if πE = πg.CheckLocal(πF, µk1, "k1"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µk1); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label23
							}
							goto Label24
							// line 585: if k1:
							πF.SetLineno(585)
						Label23:
							// line 586: non_adjacent.append((i1, j1, k1))
							πF.SetLineno(586)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µi1, "i1"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µj1, "j1"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µk1, "k1"); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple3(µi1, µj1, µk1).ToObject()
							πTemp005[0] = πTemp001
							if πE = πg.CheckLocal(πF, µnon_adjacent, "non_adjacent"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µnon_adjacent, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							goto Label24
						Label24:
							// line 588: non_adjacent.append( (la, lb, 0) )
							πF.SetLineno(588)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µla, "la"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µlb, "lb"); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple3(µla, µlb, πg.NewInt(0).ToObject()).ToObject()
							πTemp005[0] = πTemp001
							if πE = πg.CheckLocal(πF, µnon_adjacent, "non_adjacent"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µnon_adjacent, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 589: self.matching_blocks = map(Match._make, non_adjacent)
							πF.SetLineno(589)
							πTemp005 = πF.MakeArgs(2)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßMatch); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ß_make, nil); πE != nil {
								continue
							}
							πTemp005[0] = πTemp002
							if πE = πg.CheckLocal(πF, µnon_adjacent, "non_adjacent"); πE != nil {
								continue
							}
							πTemp005[1] = µnon_adjacent
							if πTemp001, πE = πg.ResolveGlobal(πF, ßmap); πE != nil {
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
							if πE = πg.SetAttr(πF, µself, ßmatching_blocks, πTemp001); πE != nil {
								continue
							}
							// line 590: return self.matching_blocks
							πF.SetLineno(590)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßmatching_blocks, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßget_matching_blocks.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 522: """Return list of triples describing matching subsequences.
					πF.SetLineno(522)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp009}, πg.NewStr("Return list of triples describing matching subsequences.\n\n        Each triple is of the form (i, j, n), and means that\n        a[i:i+n] == b[j:j+n].  The triples are monotonically increasing in\n        i and in j.  New in Python 2.5, it's also guaranteed that if\n        (i, j, n) and (i', j', n') are adjacent triples in the list, and\n        the second is not the last triple in the list, then i+n != i' or\n        j+n != j'.  IOW, adjacent triples never describe adjacent equal\n        blocks.\n\n        The last triple is a dummy, (len(a), len(b), 0), and is the only\n        triple with n==0.\n\n        >>> s = SequenceMatcher(None, \"abxcd\", \"abcd\")\n        >>> s.get_matching_blocks()\n        [Match(a=0, b=0, size=2), Match(a=3, b=2, size=2), Match(a=5, b=4, size=0)]\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßget_matching_blocks); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp010, ß__doc__, πTemp009); πE != nil {
						continue
					}
					// line 592: def get_opcodes(self):
					πF.SetLineno(592)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("get_opcodes", "/usr/lib/python2.7/difflib.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µi *πg.Object = πg.UnboundLocal; _ = µi
						var µj *πg.Object = πg.UnboundLocal; _ = µj
						var µanswer *πg.Object = πg.UnboundLocal; _ = µanswer
						var µai *πg.Object = πg.UnboundLocal; _ = µai
						var µbj *πg.Object = πg.UnboundLocal; _ = µbj
						var µsize *πg.Object = πg.UnboundLocal; _ = µsize
						var µtag *πg.Object = πg.UnboundLocal; _ = µtag
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
						var πTemp007 *πg.Object
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
							// line 593: """Return list of 5-tuples describing how to turn a into b.
							πF.SetLineno(593)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßopcodes, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp002 != πTemp003).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 621: if self.opcodes is not None:
							πF.SetLineno(621)
						Label1:
							// line 622: return self.opcodes
							πF.SetLineno(622)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßopcodes, nil); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label2
						Label2:
							// line 623: i = j = 0
							πF.SetLineno(623)
							µi = πg.NewInt(0).ToObject()
							µj = πg.NewInt(0).ToObject()
							// line 624: self.opcodes = answer = []
							πF.SetLineno(624)
							πTemp005 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp005...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßopcodes, πTemp002); πE != nil {
								continue
							}
							µanswer = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßget_matching_blocks, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
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
								πTemp006 = !isStop
							} else {
								πTemp006 = true
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp007}, πg.TieTarget{Target: &πTemp008}}}, πTemp002); πE != nil {
									continue
								}
								µai = πTemp003
								µbj = πTemp007
								µsize = πTemp008
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(3)            
							// line 631: tag = ''
							πF.SetLineno(631)
							µtag = ß.ToObject()
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µai, "ai"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.LT(πF, µi, µai); πE != nil {
								continue
							}
							πTemp002 = πTemp003
							if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp006 {
								goto Label6
							}
							if πE = πg.CheckLocal(πF, µj, "j"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µbj, "bj"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.LT(πF, µj, µbj); πE != nil {
								continue
							}
							πTemp002 = πTemp003
						Label6:
							if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label7
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µai, "ai"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.LT(πF, µi, µai); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label8
							}
							if πE = πg.CheckLocal(πF, µj, "j"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µbj, "bj"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.LT(πF, µj, µbj); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label9
							}
							goto Label10
							// line 632: if i < ai and j < bj:
							πF.SetLineno(632)
						Label7:
							// line 633: tag = 'replace'
							πF.SetLineno(633)
							µtag = ßreplace.ToObject()
							goto Label10
							// line 634: elif i < ai:
							πF.SetLineno(634)
						Label8:
							// line 635: tag = 'delete'
							πF.SetLineno(635)
							µtag = ßdelete.ToObject()
							goto Label10
							// line 636: elif j < bj:
							πF.SetLineno(636)
						Label9:
							// line 637: tag = 'insert'
							πF.SetLineno(637)
							µtag = ßinsert.ToObject()
							goto Label10
						Label10:
							if πE = πg.CheckLocal(πF, µtag, "tag"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, µtag); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label11
							}
							goto Label12
							// line 638: if tag:
							πF.SetLineno(638)
						Label11:
							// line 639: answer.append( (tag, i, ai, j, bj) )
							πF.SetLineno(639)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtag, "tag"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µai, "ai"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µj, "j"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µbj, "bj"); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple5(µtag, µi, µai, µj, µbj).ToObject()
							πTemp005[0] = πTemp002
							if πE = πg.CheckLocal(πF, µanswer, "answer"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µanswer, ßappend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							goto Label12
						Label12:
							// line 640: i, j = ai+size, bj+size
							πF.SetLineno(640)
							if πE = πg.CheckLocal(πF, µai, "ai"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µsize, "size"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, µai, µsize); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µbj, "bj"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µsize, "size"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Add(πF, µbj, µsize); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(πTemp003, πTemp007).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp007}}}, πTemp002); πE != nil {
								continue
							}
							µi = πTemp003
							µj = πTemp007
							if πE = πg.CheckLocal(πF, µsize, "size"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, µsize); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label13
							}
							goto Label14
							// line 643: if size:
							πF.SetLineno(643)
						Label13:
							// line 644: answer.append( ('equal', ai, i, bj, j) )
							πF.SetLineno(644)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µai, "ai"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µbj, "bj"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µj, "j"); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple5(ßequal.ToObject(), µai, µi, µbj, µj).ToObject()
							πTemp005[0] = πTemp002
							if πE = πg.CheckLocal(πF, µanswer, "answer"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µanswer, ßappend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							goto Label14
						Label14:
							continue
						Label4:
							if πE != nil || πR != nil {
								continue
							}
						Label5:
							// line 645: return answer
							πF.SetLineno(645)
							if πE = πg.CheckLocal(πF, µanswer, "answer"); πE != nil {
								continue
							}
							πR = µanswer
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßget_opcodes.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 593: """Return list of 5-tuples describing how to turn a into b.
					πF.SetLineno(593)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp010}, πg.NewStr("Return list of 5-tuples describing how to turn a into b.\n\n        Each tuple is of the form (tag, i1, i2, j1, j2).  The first tuple\n        has i1 == j1 == 0, and remaining tuples have i1 == the i2 from the\n        tuple preceding it, and likewise for j1 == the previous j2.\n\n        The tags are strings, with these meanings:\n\n        'replace':  a[i1:i2] should be replaced by b[j1:j2]\n        'delete':   a[i1:i2] should be deleted.\n                    Note that j1==j2 in this case.\n        'insert':   b[j1:j2] should be inserted at a[i1:i1].\n                    Note that i1==i2 in this case.\n        'equal':    a[i1:i2] == b[j1:j2]\n\n        >>> a = \"qabxcd\"\n        >>> b = \"abycdf\"\n        >>> s = SequenceMatcher(None, a, b)\n        >>> for tag, i1, i2, j1, j2 in s.get_opcodes():\n        ...    print (\"%7s a[%d:%d] (%s) b[%d:%d] (%s)\" %\n        ...           (tag, i1, i2, a[i1:i2], j1, j2, b[j1:j2]))\n         delete a[0:1] (q) b[0:0] ()\n          equal a[1:3] (ab) b[0:2] (ab)\n        replace a[3:4] (x) b[2:3] (y)\n          equal a[4:6] (cd) b[3:5] (cd)\n         insert a[6:6] () b[5:6] (f)\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ßget_opcodes); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp011, ß__doc__, πTemp010); πE != nil {
						continue
					}
					// line 647: def get_grouped_opcodes(self, n=3):
					πF.SetLineno(647)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "n", Def: πg.NewInt(3).ToObject()}
					πTemp010 = πg.NewFunction(πg.NewCode("get_grouped_opcodes", "/usr/lib/python2.7/difflib.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µn *πg.Object = πArgs[1]; _ = µn
						var µcodes *πg.Object = πg.UnboundLocal; _ = µcodes
						var µtag *πg.Object = πg.UnboundLocal; _ = µtag
						var µi1 *πg.Object = πg.UnboundLocal; _ = µi1
						var µi2 *πg.Object = πg.UnboundLocal; _ = µi2
						var µj1 *πg.Object = πg.UnboundLocal; _ = µj1
						var µj2 *πg.Object = πg.UnboundLocal; _ = µj2
						var µnn *πg.Object = πg.UnboundLocal; _ = µnn
						var µgroup *πg.Object = πg.UnboundLocal; _ = µgroup
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
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 bool
						_ = πTemp009
						var πTemp010 *πg.Object
						_ = πTemp010
						var πTemp011 []*πg.Object
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
								case 18: goto Label18
								case 13: goto Label13
								case 7: goto Label7
								default: panic("unexpected function state")
								}
								// line 648: """ Isolate change clusters by eliminating ranges with no changes.
								πF.SetLineno(648)
								// line 672: codes = self.get_opcodes()
								πF.SetLineno(672)
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.GetAttr(πF, µself, ßget_opcodes, nil); πE != nil {
									continue
								}
								if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
									continue
								}
								µcodes = πTemp002
								if πE = πg.CheckLocal(πF, µcodes, "codes"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.IsTrue(πF, µcodes); πE != nil {
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
								// line 673: if not codes:
								πF.SetLineno(673)
							Label1:
								// line 674: codes = [("equal", 0, 1, 0, 1)]
								πF.SetLineno(674)
								πTemp004 = make([]*πg.Object, 1)
								πTemp001 = πg.NewTuple5(ßequal.ToObject(), πg.NewInt(0).ToObject(), πg.NewInt(1).ToObject(), πg.NewInt(0).ToObject(), πg.NewInt(1).ToObject()).ToObject()
								πTemp004[0] = πTemp001
								πTemp001 = πg.NewList(πTemp004...).ToObject()
								µcodes = πTemp001
								goto Label2
							Label2:
								πTemp002 = πg.NewInt(0).ToObject()
								πTemp006 = πg.NewInt(0).ToObject()
								if πE = πg.CheckLocal(πF, µcodes, "codes"); πE != nil {
									continue
								}
								if πTemp007, πE = πg.GetItem(πF, µcodes, πTemp006); πE != nil {
									continue
								}
								if πTemp005, πE = πg.GetItem(πF, πTemp007, πTemp002); πE != nil {
									continue
								}
								if πTemp001, πE = πg.Eq(πF, πTemp005, ßequal.ToObject()); πE != nil {
									continue
								}
								if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
									continue
								}
								if πTemp003 {
									goto Label3
								}
								goto Label4
								// line 676: if codes[0][0] == 'equal':
								πF.SetLineno(676)
							Label3:
								// line 677: tag, i1, i2, j1, j2 = codes[0]
								πF.SetLineno(677)
								πTemp001 = πg.NewInt(0).ToObject()
								if πE = πg.CheckLocal(πF, µcodes, "codes"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.GetItem(πF, µcodes, πTemp001); πE != nil {
									continue
								}
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp006}, πg.TieTarget{Target: &πTemp007}, πg.TieTarget{Target: &πTemp008}}}, πTemp002); πE != nil {
									continue
								}
								µtag = πTemp001
								µi1 = πTemp005
								µi2 = πTemp006
								µj1 = πTemp007
								µj2 = πTemp008
								// line 678: codes[0] = tag, max(i1, i2-n), i2, max(j1, j2-n), j2
								πF.SetLineno(678)
								if πE = πg.CheckLocal(πF, µtag, "tag"); πE != nil {
									continue
								}
								πTemp004 = πF.MakeArgs(2)
								if πE = πg.CheckLocal(πF, µi1, "i1"); πE != nil {
									continue
								}
								πTemp004[0] = µi1
								if πE = πg.CheckLocal(πF, µi2, "i2"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.Sub(πF, µi2, µn); πE != nil {
									continue
								}
								πTemp004[1] = πTemp002
								if πTemp002, πE = πg.ResolveGlobal(πF, ßmax); πE != nil {
									continue
								}
								if πTemp005, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								if πE = πg.CheckLocal(πF, µi2, "i2"); πE != nil {
									continue
								}
								πTemp004 = πF.MakeArgs(2)
								if πE = πg.CheckLocal(πF, µj1, "j1"); πE != nil {
									continue
								}
								πTemp004[0] = µj1
								if πE = πg.CheckLocal(πF, µj2, "j2"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.Sub(πF, µj2, µn); πE != nil {
									continue
								}
								πTemp004[1] = πTemp002
								if πTemp002, πE = πg.ResolveGlobal(πF, ßmax); πE != nil {
									continue
								}
								if πTemp006, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								if πE = πg.CheckLocal(πF, µj2, "j2"); πE != nil {
									continue
								}
								πTemp001 = πg.NewTuple5(µtag, πTemp005, µi2, πTemp006, µj2).ToObject()
								if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µcodes, "codes"); πE != nil {
									continue
								}
								πTemp005 = πg.NewInt(0).ToObject()
								if πE = πg.SetItem(πF, µcodes, πTemp005, πTemp002); πE != nil {
									continue
								}
								goto Label4
							Label4:
								πTemp002 = πg.NewInt(0).ToObject()
								if πTemp007, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
									continue
								}
								πTemp006 = πTemp007
								if πE = πg.CheckLocal(πF, µcodes, "codes"); πE != nil {
									continue
								}
								if πTemp007, πE = πg.GetItem(πF, µcodes, πTemp006); πE != nil {
									continue
								}
								if πTemp005, πE = πg.GetItem(πF, πTemp007, πTemp002); πE != nil {
									continue
								}
								if πTemp001, πE = πg.Eq(πF, πTemp005, ßequal.ToObject()); πE != nil {
									continue
								}
								if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
									continue
								}
								if πTemp003 {
									goto Label5
								}
								goto Label6
								// line 679: if codes[-1][0] == 'equal':
								πF.SetLineno(679)
							Label5:
								// line 680: tag, i1, i2, j1, j2 = codes[-1]
								πF.SetLineno(680)
								if πTemp002, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
									continue
								}
								πTemp001 = πTemp002
								if πE = πg.CheckLocal(πF, µcodes, "codes"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.GetItem(πF, µcodes, πTemp001); πE != nil {
									continue
								}
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp006}, πg.TieTarget{Target: &πTemp007}, πg.TieTarget{Target: &πTemp008}}}, πTemp002); πE != nil {
									continue
								}
								µtag = πTemp001
								µi1 = πTemp005
								µi2 = πTemp006
								µj1 = πTemp007
								µj2 = πTemp008
								// line 681: codes[-1] = tag, i1, min(i2, i1+n), j1, min(j2, j1+n)
								πF.SetLineno(681)
								if πE = πg.CheckLocal(πF, µtag, "tag"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µi1, "i1"); πE != nil {
									continue
								}
								πTemp004 = πF.MakeArgs(2)
								if πE = πg.CheckLocal(πF, µi2, "i2"); πE != nil {
									continue
								}
								πTemp004[0] = µi2
								if πE = πg.CheckLocal(πF, µi1, "i1"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.Add(πF, µi1, µn); πE != nil {
									continue
								}
								πTemp004[1] = πTemp002
								if πTemp002, πE = πg.ResolveGlobal(πF, ßmin); πE != nil {
									continue
								}
								if πTemp005, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								if πE = πg.CheckLocal(πF, µj1, "j1"); πE != nil {
									continue
								}
								πTemp004 = πF.MakeArgs(2)
								if πE = πg.CheckLocal(πF, µj2, "j2"); πE != nil {
									continue
								}
								πTemp004[0] = µj2
								if πE = πg.CheckLocal(πF, µj1, "j1"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.Add(πF, µj1, µn); πE != nil {
									continue
								}
								πTemp004[1] = πTemp002
								if πTemp002, πE = πg.ResolveGlobal(πF, ßmin); πE != nil {
									continue
								}
								if πTemp006, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								πTemp001 = πg.NewTuple5(µtag, µi1, πTemp005, µj1, πTemp006).ToObject()
								if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µcodes, "codes"); πE != nil {
									continue
								}
								if πTemp006, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
									continue
								}
								πTemp005 = πTemp006
								if πE = πg.SetItem(πF, µcodes, πTemp005, πTemp002); πE != nil {
									continue
								}
								goto Label6
							Label6:
								// line 683: nn = n + n
								πF.SetLineno(683)
								if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.Add(πF, µn, µn); πE != nil {
									continue
								}
								µnn = πTemp001
								// line 684: group = []
								πF.SetLineno(684)
								πTemp004 = make([]*πg.Object, 0)
								πTemp001 = πg.NewList(πTemp004...).ToObject()
								µgroup = πTemp001
								if πE = πg.CheckLocal(πF, µcodes, "codes"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.Iter(πF, µcodes); πE != nil {
									continue
								}
								πF.PushCheckpoint(8)
								πTemp003 = false
							Label7:
								if πE != nil || πR != nil {
									continue
								}
								if πTemp003 {
									πF.PopCheckpoint()
									goto Label9
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
									if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp006}, πg.TieTarget{Target: &πTemp007}, πg.TieTarget{Target: &πTemp008}, πg.TieTarget{Target: &πTemp010}}}, πTemp002); πE != nil {
										continue
									}
									µtag = πTemp005
									µi1 = πTemp006
									µi2 = πTemp007
									µj1 = πTemp008
									µj2 = πTemp010
								}
								if πE != nil || !πTemp009 {
									continue
								}
								πF.PushCheckpoint(7)            
								if πE = πg.CheckLocal(πF, µtag, "tag"); πE != nil {
									continue
								}
								if πTemp005, πE = πg.Eq(πF, µtag, ßequal.ToObject()); πE != nil {
									continue
								}
								πTemp002 = πTemp005
								if πTemp009, πE = πg.IsTrue(πF, πTemp002); πE != nil {
									continue
								}
								if !πTemp009 {
									goto Label10
								}
								if πE = πg.CheckLocal(πF, µi2, "i2"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µi1, "i1"); πE != nil {
									continue
								}
								if πTemp006, πE = πg.Sub(πF, µi2, µi1); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µnn, "nn"); πE != nil {
									continue
								}
								if πTemp005, πE = πg.GT(πF, πTemp006, µnn); πE != nil {
									continue
								}
								πTemp002 = πTemp005
							Label10:
								if πTemp009, πE = πg.IsTrue(πF, πTemp002); πE != nil {
									continue
								}
								if πTemp009 {
									goto Label11
								}
								goto Label12
								// line 688: if tag == 'equal' and i2-i1 > nn:
								πF.SetLineno(688)
							Label11:
								// line 689: group.append((tag, i1, min(i2, i1+n), j1, min(j2, j1+n)))
								πF.SetLineno(689)
								πTemp004 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µtag, "tag"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µi1, "i1"); πE != nil {
									continue
								}
								πTemp011 = πF.MakeArgs(2)
								if πE = πg.CheckLocal(πF, µi2, "i2"); πE != nil {
									continue
								}
								πTemp011[0] = µi2
								if πE = πg.CheckLocal(πF, µi1, "i1"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
									continue
								}
								if πTemp005, πE = πg.Add(πF, µi1, µn); πE != nil {
									continue
								}
								πTemp011[1] = πTemp005
								if πTemp005, πE = πg.ResolveGlobal(πF, ßmin); πE != nil {
									continue
								}
								if πTemp006, πE = πTemp005.Call(πF, πTemp011, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp011)
								if πE = πg.CheckLocal(πF, µj1, "j1"); πE != nil {
									continue
								}
								πTemp011 = πF.MakeArgs(2)
								if πE = πg.CheckLocal(πF, µj2, "j2"); πE != nil {
									continue
								}
								πTemp011[0] = µj2
								if πE = πg.CheckLocal(πF, µj1, "j1"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
									continue
								}
								if πTemp005, πE = πg.Add(πF, µj1, µn); πE != nil {
									continue
								}
								πTemp011[1] = πTemp005
								if πTemp005, πE = πg.ResolveGlobal(πF, ßmin); πE != nil {
									continue
								}
								if πTemp007, πE = πTemp005.Call(πF, πTemp011, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp011)
								πTemp002 = πg.NewTuple5(µtag, µi1, πTemp006, µj1, πTemp007).ToObject()
								πTemp004[0] = πTemp002
								if πE = πg.CheckLocal(πF, µgroup, "group"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.GetAttr(πF, µgroup, ßappend, nil); πE != nil {
									continue
								}
								if πTemp005, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								// line 690: yield group
								πF.SetLineno(690)
								if πE = πg.CheckLocal(πF, µgroup, "group"); πE != nil {
									continue
								}
								πF.PushCheckpoint(13)
								return µgroup, nil
							Label13:
								πTemp002 = πSent
								// line 691: group = []
								πF.SetLineno(691)
								πTemp004 = make([]*πg.Object, 0)
								πTemp002 = πg.NewList(πTemp004...).ToObject()
								µgroup = πTemp002
								// line 692: i1, j1 = max(i1, i2-n), max(j1, j2-n)
								πF.SetLineno(692)
								πTemp004 = πF.MakeArgs(2)
								if πE = πg.CheckLocal(πF, µi1, "i1"); πE != nil {
									continue
								}
								πTemp004[0] = µi1
								if πE = πg.CheckLocal(πF, µi2, "i2"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
									continue
								}
								if πTemp005, πE = πg.Sub(πF, µi2, µn); πE != nil {
									continue
								}
								πTemp004[1] = πTemp005
								if πTemp005, πE = πg.ResolveGlobal(πF, ßmax); πE != nil {
									continue
								}
								if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								πTemp004 = πF.MakeArgs(2)
								if πE = πg.CheckLocal(πF, µj1, "j1"); πE != nil {
									continue
								}
								πTemp004[0] = µj1
								if πE = πg.CheckLocal(πF, µj2, "j2"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
									continue
								}
								if πTemp005, πE = πg.Sub(πF, µj2, µn); πE != nil {
									continue
								}
								πTemp004[1] = πTemp005
								if πTemp005, πE = πg.ResolveGlobal(πF, ßmax); πE != nil {
									continue
								}
								if πTemp007, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								πTemp002 = πg.NewTuple2(πTemp006, πTemp007).ToObject()
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp006}}}, πTemp002); πE != nil {
									continue
								}
								µi1 = πTemp005
								µj1 = πTemp006
								goto Label12
							Label12:
								// line 693: group.append((tag, i1, i2, j1 ,j2))
								πF.SetLineno(693)
								πTemp004 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µtag, "tag"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µi1, "i1"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µi2, "i2"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µj1, "j1"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µj2, "j2"); πE != nil {
									continue
								}
								πTemp002 = πg.NewTuple5(µtag, µi1, µi2, µj1, µj2).ToObject()
								πTemp004[0] = πTemp002
								if πE = πg.CheckLocal(πF, µgroup, "group"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.GetAttr(πF, µgroup, ßappend, nil); πE != nil {
									continue
								}
								if πTemp005, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								continue
							Label8:
								if πE != nil || πR != nil {
									continue
								}
							Label9:
								if πE = πg.CheckLocal(πF, µgroup, "group"); πE != nil {
									continue
								}
								πTemp001 = µgroup
								if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
									continue
								}
								if !πTemp003 {
									goto Label14
								}
								πTemp004 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µgroup, "group"); πE != nil {
									continue
								}
								πTemp004[0] = µgroup
								if πTemp007, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
									continue
								}
								if πTemp008, πE = πTemp007.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								if πTemp006, πE = πg.Eq(πF, πTemp008, πg.NewInt(1).ToObject()); πE != nil {
									continue
								}
								πTemp005 = πTemp006
								if πTemp009, πE = πg.IsTrue(πF, πTemp005); πE != nil {
									continue
								}
								if !πTemp009 {
									goto Label15
								}
								πTemp007 = πg.NewInt(0).ToObject()
								πTemp010 = πg.NewInt(0).ToObject()
								if πE = πg.CheckLocal(πF, µgroup, "group"); πE != nil {
									continue
								}
								if πTemp012, πE = πg.GetItem(πF, µgroup, πTemp010); πE != nil {
									continue
								}
								if πTemp008, πE = πg.GetItem(πF, πTemp012, πTemp007); πE != nil {
									continue
								}
								if πTemp006, πE = πg.Eq(πF, πTemp008, ßequal.ToObject()); πE != nil {
									continue
								}
								πTemp005 = πTemp006
							Label15:
								if πTemp009, πE = πg.IsTrue(πF, πTemp005); πE != nil {
									continue
								}
								πTemp002 = πg.GetBool(!πTemp009).ToObject()
								πTemp001 = πTemp002
							Label14:
								if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
									continue
								}
								if πTemp003 {
									goto Label16
								}
								goto Label17
								// line 694: if group and not (len(group)==1 and group[0][0] == 'equal'):
								πF.SetLineno(694)
							Label16:
								// line 695: yield group
								πF.SetLineno(695)
								if πE = πg.CheckLocal(πF, µgroup, "group"); πE != nil {
									continue
								}
								πF.PushCheckpoint(18)
								return µgroup, nil
							Label18:
								πTemp001 = πSent
								goto Label17
							Label17:
							}
							return nil, πE
						}).ToObject(), nil
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßget_grouped_opcodes.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 648: """ Isolate change clusters by eliminating ranges with no changes.
					πF.SetLineno(648)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πg.NewStr(" Isolate change clusters by eliminating ranges with no changes.\n\n        Return a generator of groups with up to n lines of context.\n        Each group is in the same format as returned by get_opcodes().\n\n        >>> from pprint import pprint\n        >>> a = map(str, range(1,40))\n        >>> b = a[:]\n        >>> b[8:8] = ['i']     # Make an insertion\n        >>> b[20] += 'x'       # Make a replacement\n        >>> b[23:28] = []      # Make a deletion\n        >>> b[30] += 'y'       # Make another replacement\n        >>> pprint(list(SequenceMatcher(None,a,b).get_grouped_opcodes()))\n        [[('equal', 5, 8, 5, 8), ('insert', 8, 8, 8, 9), ('equal', 8, 11, 9, 12)],\n         [('equal', 16, 19, 17, 20),\n          ('replace', 19, 20, 20, 21),\n          ('equal', 20, 22, 21, 23),\n          ('delete', 22, 27, 23, 23),\n          ('equal', 27, 30, 23, 26)],\n         [('equal', 31, 34, 27, 30),\n          ('replace', 34, 35, 30, 31),\n          ('equal', 35, 38, 31, 34)]]\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp012, πE = πg.ResolveClass(πF, πClass, nil, ßget_grouped_opcodes); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp012, ß__doc__, πTemp011); πE != nil {
						continue
					}
					// line 697: def ratio(self):
					πF.SetLineno(697)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp011 = πg.NewFunction(πg.NewCode("ratio", "/usr/lib/python2.7/difflib.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µmatches *πg.Object = πg.UnboundLocal; _ = µmatches
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []πg.Param
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
							// line 698: """Return a measure of the sequences' similarity (float in [0,1]).
							πF.SetLineno(698)
							// line 719: matches = reduce(lambda sum, triple: sum + triple[-1],
							πF.SetLineno(719)
							πTemp001 = πF.MakeArgs(3)
							πTemp003 = make([]πg.Param, 2)
							πTemp003[0] = πg.Param{Name: "sum", Def: nil}
							πTemp003[1] = πg.Param{Name: "triple", Def: nil}
							πTemp002 = πg.NewFunction(πg.NewCode("<lambda>", "/usr/lib/python2.7/difflib.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µsum *πg.Object = πArgs[0]; _ = µsum
								var µtriple *πg.Object = πArgs[1]; _ = µtriple
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
									// line 719: matches = reduce(lambda sum, triple: sum + triple[-1],
									πF.SetLineno(719)
									if πE = πg.CheckLocal(πF, µsum, "sum"); πE != nil {
										continue
									}
									if πTemp003, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
										continue
									}
									πTemp002 = πTemp003
									if πE = πg.CheckLocal(πF, µtriple, "triple"); πE != nil {
										continue
									}
									if πTemp003, πE = πg.GetItem(πF, µtriple, πTemp002); πE != nil {
										continue
									}
									if πTemp001, πE = πg.Add(πF, µsum, πTemp003); πE != nil {
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
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßget_matching_blocks, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp004
							πTemp001[2] = πg.NewInt(0).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßreduce); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µmatches = πTemp004
							// line 721: return _calculate_ratio(matches, len(self.a) + len(self.b))
							πF.SetLineno(721)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µmatches, "matches"); πE != nil {
								continue
							}
							πTemp001[0] = µmatches
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßa, nil); πE != nil {
								continue
							}
							πTemp005[0] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßb, nil); πE != nil {
								continue
							}
							πTemp005[0] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp002, πE = πg.Add(πF, πTemp006, πTemp007); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_calculate_ratio); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßratio.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 698: """Return a measure of the sequences' similarity (float in [0,1]).
					πF.SetLineno(698)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp012}, πg.NewStr("Return a measure of the sequences' similarity (float in [0,1]).\n\n        Where T is the total number of elements in both sequences, and\n        M is the number of matches, this is 2.0*M / T.\n        Note that this is 1 if the sequences are identical, and 0 if\n        they have nothing in common.\n\n        .ratio() is expensive to compute if you haven't already computed\n        .get_matching_blocks() or .get_opcodes(), in which case you may\n        want to try .quick_ratio() or .real_quick_ratio() first to get an\n        upper bound.\n\n        >>> s = SequenceMatcher(None, \"abcd\", \"bcde\")\n        >>> s.ratio()\n        0.75\n        >>> s.quick_ratio()\n        0.75\n        >>> s.real_quick_ratio()\n        1.0\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp013, πE = πg.ResolveClass(πF, πClass, nil, ßratio); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp013, ß__doc__, πTemp012); πE != nil {
						continue
					}
					// line 723: def quick_ratio(self):
					πF.SetLineno(723)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp012 = πg.NewFunction(πg.NewCode("quick_ratio", "/usr/lib/python2.7/difflib.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µfullbcount *πg.Object = πg.UnboundLocal; _ = µfullbcount
						var µelt *πg.Object = πg.UnboundLocal; _ = µelt
						var µavail *πg.Object = πg.UnboundLocal; _ = µavail
						var µavailhas *πg.Object = πg.UnboundLocal; _ = µavailhas
						var µmatches *πg.Object = πg.UnboundLocal; _ = µmatches
						var µnumb *πg.Object = πg.UnboundLocal; _ = µnumb
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 bool
						_ = πTemp004
						var πTemp005 *πg.Dict
						_ = πTemp005
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 []*πg.Object
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
							case 3: goto Label3
							case 4: goto Label4
							case 6: goto Label6
							case 7: goto Label7
							default: panic("unexpected function state")
							}
							// line 724: """Return an upper bound on ratio() relatively quickly.
							πF.SetLineno(724)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßfullbcount, nil); πE != nil {
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
							// line 733: if self.fullbcount is None:
							πF.SetLineno(733)
						Label1:
							// line 734: self.fullbcount = fullbcount = {}
							πF.SetLineno(734)
							πTemp005 = πg.NewDict()
							πTemp001 = πTemp005.ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßfullbcount, πTemp002); πE != nil {
								continue
							}
							µfullbcount = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßb, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, πTemp002); πE != nil {
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
								πTemp006 = !isStop
							} else {
								πTemp006 = true
								µelt = πTemp002
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(3)            
							// line 736: fullbcount[elt] = fullbcount.get(elt, 0) + 1
							πF.SetLineno(736)
							πTemp007 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µelt, "elt"); πE != nil {
								continue
							}
							πTemp007[0] = µelt
							πTemp007[1] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µfullbcount, "fullbcount"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µfullbcount, ßget, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp003.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							if πTemp002, πE = πg.Add(πF, πTemp008, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µfullbcount, "fullbcount"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µelt, "elt"); πE != nil {
								continue
							}
							πTemp008 = µelt
							if πE = πg.SetItem(πF, µfullbcount, πTemp008, πTemp003); πE != nil {
								continue
							}
							continue
						Label4:
							if πE != nil || πR != nil {
								continue
							}
						Label5:
							goto Label2
						Label2:
							// line 737: fullbcount = self.fullbcount
							πF.SetLineno(737)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßfullbcount, nil); πE != nil {
								continue
							}
							µfullbcount = πTemp001
							// line 740: avail = {}
							πF.SetLineno(740)
							πTemp005 = πg.NewDict()
							πTemp001 = πTemp005.ToObject()
							µavail = πTemp001
							// line 741: availhas, matches = avail.__contains__, 0
							πF.SetLineno(741)
							if πE = πg.CheckLocal(πF, µavail, "avail"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µavail, ß__contains__, nil); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple2(πTemp002, πg.NewInt(0).ToObject()).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp003}}}, πTemp001); πE != nil {
								continue
							}
							µavailhas = πTemp002
							µmatches = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßa, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, πTemp002); πE != nil {
								continue
							}
							πF.PushCheckpoint(7)
							πTemp004 = false
						Label6:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp004 {
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
								πTemp006 = !isStop
							} else {
								πTemp006 = true
								µelt = πTemp002
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(6)            
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µelt, "elt"); πE != nil {
								continue
							}
							πTemp007[0] = µelt
							if πE = πg.CheckLocal(πF, µavailhas, "availhas"); πE != nil {
								continue
							}
							if πTemp002, πE = µavailhas.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label9
							}
							goto Label10
							// line 743: if availhas(elt):
							πF.SetLineno(743)
						Label9:
							// line 744: numb = avail[elt]
							πF.SetLineno(744)
							if πE = πg.CheckLocal(πF, µelt, "elt"); πE != nil {
								continue
							}
							πTemp002 = µelt
							if πE = πg.CheckLocal(πF, µavail, "avail"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µavail, πTemp002); πE != nil {
								continue
							}
							µnumb = πTemp003
							goto Label11
						Label10:
							// line 746: numb = fullbcount.get(elt, 0)
							πF.SetLineno(746)
							πTemp007 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µelt, "elt"); πE != nil {
								continue
							}
							πTemp007[0] = µelt
							πTemp007[1] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µfullbcount, "fullbcount"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µfullbcount, ßget, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							µnumb = πTemp003
							goto Label11
						Label11:
							// line 747: avail[elt] = numb - 1
							πF.SetLineno(747)
							if πE = πg.CheckLocal(πF, µnumb, "numb"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Sub(πF, µnumb, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µavail, "avail"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µelt, "elt"); πE != nil {
								continue
							}
							πTemp008 = µelt
							if πE = πg.SetItem(πF, µavail, πTemp008, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnumb, "numb"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GT(πF, µnumb, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label12
							}
							goto Label13
							// line 748: if numb > 0:
							πF.SetLineno(748)
						Label12:
							// line 749: matches = matches + 1
							πF.SetLineno(749)
							if πE = πg.CheckLocal(πF, µmatches, "matches"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, µmatches, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µmatches = πTemp002
							goto Label13
						Label13:
							continue
						Label7:
							if πE != nil || πR != nil {
								continue
							}
						Label8:
							// line 750: return _calculate_ratio(matches, len(self.a) + len(self.b))
							πF.SetLineno(750)
							πTemp007 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µmatches, "matches"); πE != nil {
								continue
							}
							πTemp007[0] = µmatches
							πTemp009 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßa, nil); πE != nil {
								continue
							}
							πTemp009[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							πTemp009 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßb, nil); πE != nil {
								continue
							}
							πTemp009[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp002.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							if πTemp001, πE = πg.Add(πF, πTemp003, πTemp008); πE != nil {
								continue
							}
							πTemp007[1] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_calculate_ratio); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
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
					if πE = πClass.SetItem(πF, ßquick_ratio.ToObject(), πTemp012); πE != nil {
						continue
					}
					// line 724: """Return an upper bound on ratio() relatively quickly.
					πF.SetLineno(724)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp013}, πg.NewStr("Return an upper bound on ratio() relatively quickly.\n\n        This isn't defined beyond that it is an upper bound on .ratio(), and\n        is faster to compute.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp014, πE = πg.ResolveClass(πF, πClass, nil, ßquick_ratio); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp014, ß__doc__, πTemp013); πE != nil {
						continue
					}
					// line 752: def real_quick_ratio(self):
					πF.SetLineno(752)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp013 = πg.NewFunction(πg.NewCode("real_quick_ratio", "/usr/lib/python2.7/difflib.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µla *πg.Object = πg.UnboundLocal; _ = µla
						var µlb *πg.Object = πg.UnboundLocal; _ = µlb
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
						var πTemp006 []*πg.Object
						_ = πTemp006
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 753: """Return an upper bound on ratio() very quickly.
							πF.SetLineno(753)
							// line 759: la, lb = len(self.a), len(self.b)
							πF.SetLineno(759)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßa, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßb, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
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
							µla = πTemp003
							µlb = πTemp004
							// line 762: return _calculate_ratio(min(la, lb), la + lb)
							πF.SetLineno(762)
							πTemp002 = πF.MakeArgs(2)
							πTemp006 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µla, "la"); πE != nil {
								continue
							}
							πTemp006[0] = µla
							if πE = πg.CheckLocal(πF, µlb, "lb"); πE != nil {
								continue
							}
							πTemp006[1] = µlb
							if πTemp001, πE = πg.ResolveGlobal(πF, ßmin); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µla, "la"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µlb, "lb"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, µla, µlb); πE != nil {
								continue
							}
							πTemp002[1] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_calculate_ratio); πE != nil {
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
					if πE = πClass.SetItem(πF, ßreal_quick_ratio.ToObject(), πTemp013); πE != nil {
						continue
					}
					// line 753: """Return an upper bound on ratio() very quickly.
					πF.SetLineno(753)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp014}, πg.NewStr("Return an upper bound on ratio() very quickly.\n\n        This isn't defined beyond that it is an upper bound on .ratio(), and\n        is faster to compute than either .ratio() or .quick_ratio().\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp015, πE = πg.ResolveClass(πF, πClass, nil, ßreal_quick_ratio); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp015, ß__doc__, πTemp014); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp007, πE = πTemp005.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp007 == nil {
				πTemp007 = πg.TypeType.ToObject()
			}
			if πTemp008, πE = πTemp007.Call(πF, []*πg.Object{πg.NewStr("SequenceMatcher").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp005.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßSequenceMatcher.ToObject(), πTemp008); πE != nil {
				continue
			}
			// line 764: def get_close_matches(word, possibilities, n=3, cutoff=0.6):
			πF.SetLineno(764)
			πTemp004 = make([]πg.Param, 4)
			πTemp004[0] = πg.Param{Name: "word", Def: nil}
			πTemp004[1] = πg.Param{Name: "possibilities", Def: nil}
			πTemp004[2] = πg.Param{Name: "n", Def: πg.NewInt(3).ToObject()}
			πTemp004[3] = πg.Param{Name: "cutoff", Def: πg.NewFloat(0.6).ToObject()}
			πTemp006 = πg.NewFunction(πg.NewCode("get_close_matches", "/usr/lib/python2.7/difflib.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µword *πg.Object = πArgs[0]; _ = µword
				var µpossibilities *πg.Object = πArgs[1]; _ = µpossibilities
				var µn *πg.Object = πArgs[2]; _ = µn
				var µcutoff *πg.Object = πArgs[3]; _ = µcutoff
				var µresult *πg.Object = πg.UnboundLocal; _ = µresult
				var µs *πg.Object = πg.UnboundLocal; _ = µs
				var µx *πg.Object = πg.UnboundLocal; _ = µx
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
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 *πg.Object
				_ = πTemp008
				var πTemp009 []πg.Param
				_ = πTemp009
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 6: goto Label6
					case 7: goto Label7
					default: panic("unexpected function state")
					}
					// line 765: """Use SequenceMatcher to return list of the best "good enough" matches.
					πF.SetLineno(765)
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GT(πF, µn, πg.NewInt(0).ToObject()); πE != nil {
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
					// line 793: if not n >  0:
					πF.SetLineno(793)
				Label1:
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple1(µn).ToObject()
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("n must be > 0: %r").ToObject(), πTemp002); πE != nil {
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
					// line 794: raise ValueError("n must be > 0: %r" % (n,))
					πF.SetLineno(794)
					πE = πF.Raise(πTemp002, nil, nil)
					continue
					goto Label2
				Label2:
					if πE = πg.CheckLocal(πF, µcutoff, "cutoff"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.LE(πF, πg.NewFloat(0.0).ToObject(), µcutoff); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp003 {
						goto Label3
					}
					if πTemp002, πE = πg.LE(πF, µcutoff, πg.NewFloat(1.0).ToObject()); πE != nil {
						continue
					}
				Label3:
					if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(!πTemp003).ToObject()
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label4
					}
					goto Label5
					// line 795: if not 0.0 <= cutoff <= 1.0:
					πF.SetLineno(795)
				Label4:
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µcutoff, "cutoff"); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple1(µcutoff).ToObject()
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("cutoff must be in [0.0, 1.0]: %r").ToObject(), πTemp002); πE != nil {
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
					// line 796: raise ValueError("cutoff must be in [0.0, 1.0]: %r" % (cutoff,))
					πF.SetLineno(796)
					πE = πF.Raise(πTemp002, nil, nil)
					continue
					goto Label5
				Label5:
					// line 797: result = []
					πF.SetLineno(797)
					πTemp004 = make([]*πg.Object, 0)
					πTemp001 = πg.NewList(πTemp004...).ToObject()
					µresult = πTemp001
					// line 798: s = SequenceMatcher()
					πF.SetLineno(798)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßSequenceMatcher); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					µs = πTemp002
					// line 799: s.set_seq2(word)
					πF.SetLineno(799)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µword, "word"); πE != nil {
						continue
					}
					πTemp004[0] = µword
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µs, ßset_seq2, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πE = πg.CheckLocal(πF, µpossibilities, "possibilities"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Iter(πF, µpossibilities); πE != nil {
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
						πTemp005 = !isStop
					} else {
						πTemp005 = true
						µx = πTemp002
					}
					if πE != nil || !πTemp005 {
						continue
					}
					πF.PushCheckpoint(6)            
					// line 801: s.set_seq1(x)
					πF.SetLineno(801)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp004[0] = µx
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µs, ßset_seq1, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, µs, ßreal_quick_ratio, nil); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp007.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcutoff, "cutoff"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GE(πF, πTemp008, µcutoff); πE != nil {
						continue
					}
					πTemp002 = πTemp006
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp005 {
						goto Label9
					}
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, µs, ßquick_ratio, nil); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp007.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcutoff, "cutoff"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GE(πF, πTemp008, µcutoff); πE != nil {
						continue
					}
					πTemp002 = πTemp006
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp005 {
						goto Label9
					}
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, µs, ßratio, nil); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp007.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcutoff, "cutoff"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GE(πF, πTemp008, µcutoff); πE != nil {
						continue
					}
					πTemp002 = πTemp006
				Label9:
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label10
					}
					goto Label11
					// line 802: if s.real_quick_ratio() >= cutoff and \
					πF.SetLineno(802)
				Label10:
					// line 805: result.append((s.ratio(), x))
					πF.SetLineno(805)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, µs, ßratio, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(πTemp007, µx).ToObject()
					πTemp004[0] = πTemp002
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µresult, ßappend, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					goto Label11
				Label11:
					continue
				Label7:
					if πE != nil || πR != nil {
						continue
					}
				Label8:
					// line 808: result = heapq.nlargest(n, result)
					πF.SetLineno(808)
					πTemp004 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					πTemp004[0] = µn
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					πTemp004[1] = µresult
					if πTemp001, πE = πg.ResolveGlobal(πF, ßheapq); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßnlargest, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µresult = πTemp001
					// line 810: return [x for score, x in result]
					πF.SetLineno(810)
					πTemp009 = make([]πg.Param, 0)
					πTemp002 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/difflib.py", πTemp009, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µscore *πg.Object = πg.UnboundLocal; _ = µscore
						var µx *πg.Object = πg.UnboundLocal; _ = µx
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
								if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.Iter(πF, µresult); πE != nil {
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
									if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp006}}}, πTemp004); πE != nil {
										continue
									}
									µscore = πTemp005
									µx = πTemp006
								}
								if πE != nil || !πTemp003 {
									continue
								}
								πF.PushCheckpoint(1)            
								// line 810: return [x for score, x in result]
								πF.SetLineno(810)
								if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
									continue
								}
								πF.PushCheckpoint(4)
								return µx, nil
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
					if πTemp006, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.ListType.Call(πF, πg.Args{πTemp006}, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßget_close_matches.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 765: """Use SequenceMatcher to return list of the best "good enough" matches.
			πF.SetLineno(765)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp007}, πg.NewStr("Use SequenceMatcher to return list of the best \"good enough\" matches.\n\n    word is a sequence for which close matches are desired (typically a\n    string).\n\n    possibilities is a list of sequences against which to match word\n    (typically a list of strings).\n\n    Optional arg n (default 3) is the maximum number of close matches to\n    return.  n must be > 0.\n\n    Optional arg cutoff (default 0.6) is a float in [0, 1].  Possibilities\n    that don't score at least that similar to word are ignored.\n\n    The best (no more than n) matches among the possibilities are returned\n    in a list, sorted by similarity score, most similar first.\n\n    >>> get_close_matches(\"appel\", [\"ape\", \"apple\", \"peach\", \"puppy\"])\n    ['apple', 'ape']\n    >>> import keyword as _keyword\n    >>> get_close_matches(\"wheel\", _keyword.kwlist)\n    ['while']\n    >>> get_close_matches(\"apple\", _keyword.kwlist)\n    []\n    >>> get_close_matches(\"accept\", _keyword.kwlist)\n    ['except']\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp008, πE = πg.ResolveGlobal(πF, ßget_close_matches); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp008, ß__doc__, πTemp007); πE != nil {
				continue
			}
			// line 812: def _count_leading(line, ch):
			πF.SetLineno(812)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "line", Def: nil}
			πTemp004[1] = πg.Param{Name: "ch", Def: nil}
			πTemp007 = πg.NewFunction(πg.NewCode("_count_leading", "/usr/lib/python2.7/difflib.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µline *πg.Object = πArgs[0]; _ = µline
				var µch *πg.Object = πArgs[1]; _ = µch
				var µi *πg.Object = πg.UnboundLocal; _ = µi
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
				var πTemp007 bool
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
					// line 813: """
					πF.SetLineno(813)
					// line 822: i, n = 0, len(line)
					πF.SetLineno(822)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					πTemp002[0] = µline
					if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001 = πg.NewTuple2(πg.NewInt(0).ToObject(), πTemp004).ToObject()
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp004}}}, πTemp001); πE != nil {
						continue
					}
					µi = πTemp003
					µn = πTemp004
					// line 823: while i < n and line[i] == ch:
					πF.SetLineno(823)
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
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.LT(πF, µi, µn); πE != nil {
						continue
					}
					πTemp001 = πTemp003
					if πTemp007, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp007 {
						goto Label4
					}
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					πTemp004 = µi
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetItem(πF, µline, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µch, "ch"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Eq(πF, πTemp008, µch); πE != nil {
						continue
					}
					πTemp001 = πTemp003
				Label4:
					if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πE != nil || !πTemp006 {
						continue
					}
					πF.PushCheckpoint(1)            
					// line 824: i += 1
					πF.SetLineno(824)
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.IAdd(πF, µi, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µi = πTemp001
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
					// line 825: return i
					πF.SetLineno(825)
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
			if πE = πF.Globals().SetItem(πF, ß_count_leading.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 813: """
			πF.SetLineno(813)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp008}, πg.NewStr("\n    Return number of `ch` characters at the start of `line`.\n\n    Example:\n\n    >>> _count_leading('   abc', ' ')\n    3\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp009, πE = πg.ResolveGlobal(πF, ß_count_leading); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp009, ß__doc__, πTemp008); πE != nil {
				continue
			}
			// line 827: class Differ(object):
			πF.SetLineno(827)
			πTemp001 = make([]*πg.Object, 1)
			if πTemp010, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp001[0] = πTemp010
			πTemp005 = πg.NewDict()
			if πTemp008, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp005.SetItem(πF, ß__module__.ToObject(), πTemp008); πE != nil {
				continue
			}
			_, πE = πg.NewCode("Differ", "/usr/lib/python2.7/difflib.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 828: r"""
					πF.SetLineno(828)
					// line 828: r"""
					πF.SetLineno(828)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Differ is a class for comparing sequences of lines of text, and\n    producing human-readable differences or deltas.  Differ uses\n    SequenceMatcher both to compare sequences of lines, and to compare\n    sequences of characters within similar (near-matching) lines.\n\n    Each line of a Differ delta begins with a two-letter code:\n\n        '- '    line unique to sequence 1\n        '+ '    line unique to sequence 2\n        '  '    line common to both sequences\n        '? '    line not present in either input sequence\n\n    Lines beginning with '? ' attempt to guide the eye to intraline\n    differences, and were not present in either input sequence.  These lines\n    can be confusing if the sequences contain tab characters.\n\n    Note that Differ makes no claim to produce a *minimal* diff.  To the\n    contrary, minimal diffs are often counter-intuitive, because they synch\n    up anywhere possible, sometimes accidental matches 100 pages apart.\n    Restricting synch points to contiguous matches preserves some notion of\n    locality, at the occasional cost of producing a longer diff.\n\n    Example: Comparing two texts.\n\n    First we set up the texts, sequences of individual single-line strings\n    ending with newlines (such sequences can also be obtained from the\n    `readlines()` method of file-like objects):\n\n    >>> text1 = '''  1. Beautiful is better than ugly.\n    ...   2. Explicit is better than implicit.\n    ...   3. Simple is better than complex.\n    ...   4. Complex is better than complicated.\n    ... '''.splitlines(1)\n    >>> len(text1)\n    4\n    >>> text1[0][-1]\n    '\\n'\n    >>> text2 = '''  1. Beautiful is better than ugly.\n    ...   3.   Simple is better than complex.\n    ...   4. Complicated is better than complex.\n    ...   5. Flat is better than nested.\n    ... '''.splitlines(1)\n\n    Next we instantiate a Differ object:\n\n    >>> d = Differ()\n\n    Note that when instantiating a Differ object we may pass functions to\n    filter out line and character 'junk'.  See Differ.__init__ for details.\n\n    Finally, we compare the two:\n\n    >>> result = list(d.compare(text1, text2))\n\n    'result' is a list of strings, so let's pretty-print it:\n\n    >>> from pprint import pprint as _pprint\n    >>> _pprint(result)\n    ['    1. Beautiful is better than ugly.\\n',\n     '-   2. Explicit is better than implicit.\\n',\n     '-   3. Simple is better than complex.\\n',\n     '+   3.   Simple is better than complex.\\n',\n     '?     ++\\n',\n     '-   4. Complex is better than complicated.\\n',\n     '?            ^                     ---- ^\\n',\n     '+   4. Complicated is better than complex.\\n',\n     '?           ++++ ^                      ^\\n',\n     '+   5. Flat is better than nested.\\n']\n\n    As a single multi-line string it looks like this:\n\n    >>> print ''.join(result),\n        1. Beautiful is better than ugly.\n    -   2. Explicit is better than implicit.\n    -   3. Simple is better than complex.\n    +   3.   Simple is better than complex.\n    ?     ++\n    -   4. Complex is better than complicated.\n    ?            ^                     ---- ^\n    +   4. Complicated is better than complex.\n    ?           ++++ ^                      ^\n    +   5. Flat is better than nested.\n\n    Methods:\n\n    __init__(linejunk=None, charjunk=None)\n        Construct a text differencer, with optional filters.\n\n    compare(a, b)\n        Compare two sequences of lines; generate the resulting delta.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 921: def __init__(self, linejunk=None, charjunk=None):
					πF.SetLineno(921)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[1] = πg.Param{Name: "linejunk", Def: πTemp003}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "charjunk", Def: πTemp003}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/difflib.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µlinejunk *πg.Object = πArgs[1]; _ = µlinejunk
						var µcharjunk *πg.Object = πArgs[2]; _ = µcharjunk
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 922: """
							πF.SetLineno(922)
							// line 942: self.linejunk = linejunk
							πF.SetLineno(942)
							if πE = πg.CheckLocal(πF, µlinejunk, "linejunk"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µlinejunk); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßlinejunk, πTemp001); πE != nil {
								continue
							}
							// line 943: self.charjunk = charjunk
							πF.SetLineno(943)
							if πE = πg.CheckLocal(πF, µcharjunk, "charjunk"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µcharjunk); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßcharjunk, πTemp001); πE != nil {
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
					// line 922: """
					πF.SetLineno(922)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewStr("\n        Construct a text differencer, with optional filters.\n\n        The two optional keyword parameters are for filter functions:\n\n        - `linejunk`: A function that should accept a single string argument,\n          and return true iff the string is junk. The module-level function\n          `IS_LINE_JUNK` may be used to filter out lines without visible\n          characters, except for at most one splat ('#').  It is recommended\n          to leave linejunk None; as of Python 2.3, the underlying\n          SequenceMatcher class has grown an adaptive notion of \"noise\" lines\n          that's better than any static definition the author has ever been\n          able to craft.\n\n        - `charjunk`: A function that should accept a string of length 1. The\n          module-level function `IS_CHARACTER_JUNK` may be used to filter out\n          whitespace characters (a blank or tab; **note**: bad idea to include\n          newline in this!).  Use of IS_CHARACTER_JUNK is recommended.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ß__init__); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp004, ß__doc__, πTemp003); πE != nil {
						continue
					}
					// line 945: def compare(self, a, b):
					πF.SetLineno(945)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "a", Def: nil}
					πTemp002[2] = πg.Param{Name: "b", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("compare", "/usr/lib/python2.7/difflib.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µa *πg.Object = πArgs[1]; _ = µa
						var µb *πg.Object = πArgs[2]; _ = µb
						var µcruncher *πg.Object = πg.UnboundLocal; _ = µcruncher
						var µtag *πg.Object = πg.UnboundLocal; _ = µtag
						var µalo *πg.Object = πg.UnboundLocal; _ = µalo
						var µahi *πg.Object = πg.UnboundLocal; _ = µahi
						var µblo *πg.Object = πg.UnboundLocal; _ = µblo
						var µbhi *πg.Object = πg.UnboundLocal; _ = µbhi
						var µg *πg.Object = πg.UnboundLocal; _ = µg
						var µline *πg.Object = πg.UnboundLocal; _ = µline
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
						var πTemp011 bool
						_ = πTemp011
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
							for ; πF.State() >= 0; πF.PopCheckpoint() {
								switch πF.State() {
								case 0:
								case 1: goto Label1
								case 2: goto Label2
								case 11: goto Label11
								case 10: goto Label10
								case 13: goto Label13
								default: panic("unexpected function state")
								}
								// line 946: r"""
								πF.SetLineno(946)
								// line 970: cruncher = SequenceMatcher(self.linejunk, a, b)
								πF.SetLineno(970)
								πTemp001 = πF.MakeArgs(3)
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.GetAttr(πF, µself, ßlinejunk, nil); πE != nil {
									continue
								}
								πTemp001[0] = πTemp002
								if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
									continue
								}
								πTemp001[1] = µa
								if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
									continue
								}
								πTemp001[2] = µb
								if πTemp002, πE = πg.ResolveGlobal(πF, ßSequenceMatcher); πE != nil {
									continue
								}
								if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								µcruncher = πTemp003
								if πE = πg.CheckLocal(πF, µcruncher, "cruncher"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.GetAttr(πF, µcruncher, ßget_opcodes, nil); πE != nil {
									continue
								}
								if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
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
									if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}, πg.TieTarget{Target: &πTemp008}, πg.TieTarget{Target: &πTemp009}, πg.TieTarget{Target: &πTemp010}}}, πTemp003); πE != nil {
										continue
									}
									µtag = πTemp004
									µalo = πTemp007
									µahi = πTemp008
									µblo = πTemp009
									µbhi = πTemp010
								}
								if πE != nil || !πTemp006 {
									continue
								}
								πF.PushCheckpoint(1)            
								if πE = πg.CheckLocal(πF, µtag, "tag"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.Eq(πF, µtag, ßreplace.ToObject()); πE != nil {
									continue
								}
								if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
									continue
								}
								if πTemp006 {
									goto Label4
								}
								if πE = πg.CheckLocal(πF, µtag, "tag"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.Eq(πF, µtag, ßdelete.ToObject()); πE != nil {
									continue
								}
								if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
									continue
								}
								if πTemp006 {
									goto Label5
								}
								if πE = πg.CheckLocal(πF, µtag, "tag"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.Eq(πF, µtag, ßinsert.ToObject()); πE != nil {
									continue
								}
								if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
									continue
								}
								if πTemp006 {
									goto Label6
								}
								if πE = πg.CheckLocal(πF, µtag, "tag"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.Eq(πF, µtag, ßequal.ToObject()); πE != nil {
									continue
								}
								if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
									continue
								}
								if πTemp006 {
									goto Label7
								}
								goto Label8
								// line 972: if tag == 'replace':
								πF.SetLineno(972)
							Label4:
								// line 973: g = self._fancy_replace(a, alo, ahi, b, blo, bhi)
								πF.SetLineno(973)
								πTemp001 = πF.MakeArgs(6)
								if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
									continue
								}
								πTemp001[0] = µa
								if πE = πg.CheckLocal(πF, µalo, "alo"); πE != nil {
									continue
								}
								πTemp001[1] = µalo
								if πE = πg.CheckLocal(πF, µahi, "ahi"); πE != nil {
									continue
								}
								πTemp001[2] = µahi
								if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
									continue
								}
								πTemp001[3] = µb
								if πE = πg.CheckLocal(πF, µblo, "blo"); πE != nil {
									continue
								}
								πTemp001[4] = µblo
								if πE = πg.CheckLocal(πF, µbhi, "bhi"); πE != nil {
									continue
								}
								πTemp001[5] = µbhi
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.GetAttr(πF, µself, ß_fancy_replace, nil); πE != nil {
									continue
								}
								if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								µg = πTemp004
								goto Label9
								// line 974: elif tag == 'delete':
								πF.SetLineno(974)
							Label5:
								// line 975: g = self._dump('-', a, alo, ahi)
								πF.SetLineno(975)
								πTemp001 = πF.MakeArgs(4)
								πTemp001[0] = πg.NewStr("-").ToObject()
								if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
									continue
								}
								πTemp001[1] = µa
								if πE = πg.CheckLocal(πF, µalo, "alo"); πE != nil {
									continue
								}
								πTemp001[2] = µalo
								if πE = πg.CheckLocal(πF, µahi, "ahi"); πE != nil {
									continue
								}
								πTemp001[3] = µahi
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.GetAttr(πF, µself, ß_dump, nil); πE != nil {
									continue
								}
								if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								µg = πTemp004
								goto Label9
								// line 976: elif tag == 'insert':
								πF.SetLineno(976)
							Label6:
								// line 977: g = self._dump('+', b, blo, bhi)
								πF.SetLineno(977)
								πTemp001 = πF.MakeArgs(4)
								πTemp001[0] = πg.NewStr("+").ToObject()
								if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
									continue
								}
								πTemp001[1] = µb
								if πE = πg.CheckLocal(πF, µblo, "blo"); πE != nil {
									continue
								}
								πTemp001[2] = µblo
								if πE = πg.CheckLocal(πF, µbhi, "bhi"); πE != nil {
									continue
								}
								πTemp001[3] = µbhi
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.GetAttr(πF, µself, ß_dump, nil); πE != nil {
									continue
								}
								if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								µg = πTemp004
								goto Label9
								// line 978: elif tag == 'equal':
								πF.SetLineno(978)
							Label7:
								// line 979: g = self._dump(' ', a, alo, ahi)
								πF.SetLineno(979)
								πTemp001 = πF.MakeArgs(4)
								πTemp001[0] = πg.NewStr(" ").ToObject()
								if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
									continue
								}
								πTemp001[1] = µa
								if πE = πg.CheckLocal(πF, µalo, "alo"); πE != nil {
									continue
								}
								πTemp001[2] = µalo
								if πE = πg.CheckLocal(πF, µahi, "ahi"); πE != nil {
									continue
								}
								πTemp001[3] = µahi
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.GetAttr(πF, µself, ß_dump, nil); πE != nil {
									continue
								}
								if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								µg = πTemp004
								goto Label9
							Label8:
								if πTemp003, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µtag, "tag"); πE != nil {
									continue
								}
								πTemp007 = πg.NewTuple1(µtag).ToObject()
								if πTemp004, πE = πg.Mod(πF, πg.NewStr("unknown tag %r").ToObject(), πTemp007); πE != nil {
									continue
								}
								// line 981: raise ValueError, 'unknown tag %r' % (tag,)
								πF.SetLineno(981)
								πE = πF.Raise(πTemp003, πTemp004, nil)
								continue
								goto Label9
							Label9:
								if πE = πg.CheckLocal(πF, µg, "g"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.Iter(πF, µg); πE != nil {
									continue
								}
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
								if πTemp004, πE = πg.Next(πF, πTemp003); πE != nil {
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
									µline = πTemp004
								}
								if πE != nil || !πTemp011 {
									continue
								}
								πF.PushCheckpoint(10)            
								// line 984: yield line
								πF.SetLineno(984)
								if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
									continue
								}
								πF.PushCheckpoint(13)
								return µline, nil
							Label13:
								πTemp004 = πSent
								continue
							Label11:
								if πE != nil || πR != nil {
									continue
								}
							Label12:
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
					if πE = πClass.SetItem(πF, ßcompare.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 946: r"""
					πF.SetLineno(946)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πg.NewStr("\n        Compare two sequences of lines; generate the resulting delta.\n\n        Each sequence must contain individual single-line strings ending with\n        newlines. Such sequences can be obtained from the `readlines()` method\n        of file-like objects.  The delta generated also consists of newline-\n        terminated strings, ready to be printed as-is via the writeline()\n        method of a file-like object.\n\n        Example:\n\n        >>> print ''.join(Differ().compare('one\\ntwo\\nthree\\n'.splitlines(1),\n        ...                                'ore\\ntree\\nemu\\n'.splitlines(1))),\n        - one\n        ?  ^\n        + ore\n        ?  ^\n        - two\n        - three\n        ?  -\n        + tree\n        + emu\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßcompare); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp005, ß__doc__, πTemp004); πE != nil {
						continue
					}
					// line 986: def _dump(self, tag, x, lo, hi):
					πF.SetLineno(986)
					πTemp002 = make([]πg.Param, 5)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "tag", Def: nil}
					πTemp002[2] = πg.Param{Name: "x", Def: nil}
					πTemp002[3] = πg.Param{Name: "lo", Def: nil}
					πTemp002[4] = πg.Param{Name: "hi", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("_dump", "/usr/lib/python2.7/difflib.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µtag *πg.Object = πArgs[1]; _ = µtag
						var µx *πg.Object = πArgs[2]; _ = µx
						var µlo *πg.Object = πArgs[3]; _ = µlo
						var µhi *πg.Object = πArgs[4]; _ = µhi
						var µi *πg.Object = πg.UnboundLocal; _ = µi
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
								// line 987: """Generate comparison results for a same-tagged range."""
								πF.SetLineno(987)
								πTemp002 = πF.MakeArgs(2)
								if πE = πg.CheckLocal(πF, µlo, "lo"); πE != nil {
									continue
								}
								πTemp002[0] = µlo
								if πE = πg.CheckLocal(πF, µhi, "hi"); πE != nil {
									continue
								}
								πTemp002[1] = µhi
								if πTemp003, πE = πg.ResolveGlobal(πF, ßxrange); πE != nil {
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
									µi = πTemp003
								}
								if πE != nil || !πTemp006 {
									continue
								}
								πF.PushCheckpoint(1)            
								// line 989: yield '%s %s' % (tag, x[i])
								πF.SetLineno(989)
								if πE = πg.CheckLocal(πF, µtag, "tag"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
									continue
								}
								πTemp007 = µi
								if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
									continue
								}
								if πTemp008, πE = πg.GetItem(πF, µx, πTemp007); πE != nil {
									continue
								}
								πTemp004 = πg.NewTuple2(µtag, πTemp008).ToObject()
								if πTemp003, πE = πg.Mod(πF, πg.NewStr("%s %s").ToObject(), πTemp004); πE != nil {
									continue
								}
								πF.PushCheckpoint(4)
								return πTemp003, nil
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
					if πE = πClass.SetItem(πF, ß_dump.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 987: """Generate comparison results for a same-tagged range."""
					πF.SetLineno(987)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πg.NewStr("Generate comparison results for a same-tagged range.").ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ß_dump); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp006, ß__doc__, πTemp005); πE != nil {
						continue
					}
					// line 991: def _plain_replace(self, a, alo, ahi, b, blo, bhi):
					πF.SetLineno(991)
					πTemp002 = make([]πg.Param, 7)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "a", Def: nil}
					πTemp002[2] = πg.Param{Name: "alo", Def: nil}
					πTemp002[3] = πg.Param{Name: "ahi", Def: nil}
					πTemp002[4] = πg.Param{Name: "b", Def: nil}
					πTemp002[5] = πg.Param{Name: "blo", Def: nil}
					πTemp002[6] = πg.Param{Name: "bhi", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("_plain_replace", "/usr/lib/python2.7/difflib.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µa *πg.Object = πArgs[1]; _ = µa
						var µalo *πg.Object = πArgs[2]; _ = µalo
						var µahi *πg.Object = πArgs[3]; _ = µahi
						var µb *πg.Object = πArgs[4]; _ = µb
						var µblo *πg.Object = πArgs[5]; _ = µblo
						var µbhi *πg.Object = πArgs[6]; _ = µbhi
						var µfirst *πg.Object = πg.UnboundLocal; _ = µfirst
						var µsecond *πg.Object = πg.UnboundLocal; _ = µsecond
						var µg *πg.Object = πg.UnboundLocal; _ = µg
						var µline *πg.Object = πg.UnboundLocal; _ = µline
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
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 bool
						_ = πTemp007
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
							for ; πF.State() >= 0; πF.PopCheckpoint() {
								switch πF.State() {
								case 0:
								case 8: goto Label8
								case 9: goto Label9
								case 11: goto Label11
								case 5: goto Label5
								case 6: goto Label6
								default: panic("unexpected function state")
								}
								// line 992: assert alo < ahi and blo < bhi
								πF.SetLineno(992)
								if πE = πg.CheckLocal(πF, µalo, "alo"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µahi, "ahi"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.LT(πF, µalo, µahi); πE != nil {
									continue
								}
								πTemp001 = πTemp003
								if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
									continue
								}
								if !πTemp002 {
									goto Label1
								}
								if πE = πg.CheckLocal(πF, µblo, "blo"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µbhi, "bhi"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.LT(πF, µblo, µbhi); πE != nil {
									continue
								}
								πTemp001 = πTemp003
							Label1:
								if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µbhi, "bhi"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µblo, "blo"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.Sub(πF, µbhi, µblo); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µahi, "ahi"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µalo, "alo"); πE != nil {
									continue
								}
								if πTemp004, πE = πg.Sub(πF, µahi, µalo); πE != nil {
									continue
								}
								if πTemp001, πE = πg.LT(πF, πTemp003, πTemp004); πE != nil {
									continue
								}
								if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
									continue
								}
								if πTemp002 {
									goto Label2
								}
								goto Label3
								// line 995: if bhi - blo < ahi - alo:
								πF.SetLineno(995)
							Label2:
								// line 996: first  = self._dump('+', b, blo, bhi)
								πF.SetLineno(996)
								πTemp005 = πF.MakeArgs(4)
								πTemp005[0] = πg.NewStr("+").ToObject()
								if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
									continue
								}
								πTemp005[1] = µb
								if πE = πg.CheckLocal(πF, µblo, "blo"); πE != nil {
									continue
								}
								πTemp005[2] = µblo
								if πE = πg.CheckLocal(πF, µbhi, "bhi"); πE != nil {
									continue
								}
								πTemp005[3] = µbhi
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.GetAttr(πF, µself, ß_dump, nil); πE != nil {
									continue
								}
								if πTemp003, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp005)
								µfirst = πTemp003
								// line 997: second = self._dump('-', a, alo, ahi)
								πF.SetLineno(997)
								πTemp005 = πF.MakeArgs(4)
								πTemp005[0] = πg.NewStr("-").ToObject()
								if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
									continue
								}
								πTemp005[1] = µa
								if πE = πg.CheckLocal(πF, µalo, "alo"); πE != nil {
									continue
								}
								πTemp005[2] = µalo
								if πE = πg.CheckLocal(πF, µahi, "ahi"); πE != nil {
									continue
								}
								πTemp005[3] = µahi
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.GetAttr(πF, µself, ß_dump, nil); πE != nil {
									continue
								}
								if πTemp003, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp005)
								µsecond = πTemp003
								goto Label4
							Label3:
								// line 999: first  = self._dump('-', a, alo, ahi)
								πF.SetLineno(999)
								πTemp005 = πF.MakeArgs(4)
								πTemp005[0] = πg.NewStr("-").ToObject()
								if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
									continue
								}
								πTemp005[1] = µa
								if πE = πg.CheckLocal(πF, µalo, "alo"); πE != nil {
									continue
								}
								πTemp005[2] = µalo
								if πE = πg.CheckLocal(πF, µahi, "ahi"); πE != nil {
									continue
								}
								πTemp005[3] = µahi
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.GetAttr(πF, µself, ß_dump, nil); πE != nil {
									continue
								}
								if πTemp003, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp005)
								µfirst = πTemp003
								// line 1000: second = self._dump('+', b, blo, bhi)
								πF.SetLineno(1000)
								πTemp005 = πF.MakeArgs(4)
								πTemp005[0] = πg.NewStr("+").ToObject()
								if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
									continue
								}
								πTemp005[1] = µb
								if πE = πg.CheckLocal(πF, µblo, "blo"); πE != nil {
									continue
								}
								πTemp005[2] = µblo
								if πE = πg.CheckLocal(πF, µbhi, "bhi"); πE != nil {
									continue
								}
								πTemp005[3] = µbhi
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.GetAttr(πF, µself, ß_dump, nil); πE != nil {
									continue
								}
								if πTemp003, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp005)
								µsecond = πTemp003
								goto Label4
							Label4:
								if πE = πg.CheckLocal(πF, µfirst, "first"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µsecond, "second"); πE != nil {
									continue
								}
								πTemp003 = πg.NewTuple2(µfirst, µsecond).ToObject()
								if πTemp001, πE = πg.Iter(πF, πTemp003); πE != nil {
									continue
								}
								πF.PushCheckpoint(6)
								πTemp002 = false
							Label5:
								if πE != nil || πR != nil {
									continue
								}
								if πTemp002 {
									πF.PopCheckpoint()
									goto Label7
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
									µg = πTemp003
								}
								if πE != nil || !πTemp006 {
									continue
								}
								πF.PushCheckpoint(5)            
								if πE = πg.CheckLocal(πF, µg, "g"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.Iter(πF, µg); πE != nil {
									continue
								}
								πF.PushCheckpoint(9)
								πTemp006 = false
							Label8:
								if πE != nil || πR != nil {
									continue
								}
								if πTemp006 {
									πF.PopCheckpoint()
									goto Label10
								}
								if πTemp004, πE = πg.Next(πF, πTemp003); πE != nil {
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
									µline = πTemp004
								}
								if πE != nil || !πTemp007 {
									continue
								}
								πF.PushCheckpoint(8)            
								// line 1004: yield line
								πF.SetLineno(1004)
								if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
									continue
								}
								πF.PushCheckpoint(11)
								return µline, nil
							Label11:
								πTemp004 = πSent
								continue
							Label9:
								if πE != nil || πR != nil {
									continue
								}
							Label10:
								continue
							Label6:
								if πE != nil || πR != nil {
									continue
								}
							Label7:
							}
							return nil, πE
						}).ToObject(), nil
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß_plain_replace.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 1006: def _fancy_replace(self, a, alo, ahi, b, blo, bhi):
					πF.SetLineno(1006)
					πTemp002 = make([]πg.Param, 7)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "a", Def: nil}
					πTemp002[2] = πg.Param{Name: "alo", Def: nil}
					πTemp002[3] = πg.Param{Name: "ahi", Def: nil}
					πTemp002[4] = πg.Param{Name: "b", Def: nil}
					πTemp002[5] = πg.Param{Name: "blo", Def: nil}
					πTemp002[6] = πg.Param{Name: "bhi", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("_fancy_replace", "/usr/lib/python2.7/difflib.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µa *πg.Object = πArgs[1]; _ = µa
						var µalo *πg.Object = πArgs[2]; _ = µalo
						var µahi *πg.Object = πArgs[3]; _ = µahi
						var µb *πg.Object = πArgs[4]; _ = µb
						var µblo *πg.Object = πArgs[5]; _ = µblo
						var µbhi *πg.Object = πArgs[6]; _ = µbhi
						var µbest_ratio *πg.Object = πg.UnboundLocal; _ = µbest_ratio
						var µcutoff *πg.Object = πg.UnboundLocal; _ = µcutoff
						var µcruncher *πg.Object = πg.UnboundLocal; _ = µcruncher
						var µeqi *πg.Object = πg.UnboundLocal; _ = µeqi
						var µeqj *πg.Object = πg.UnboundLocal; _ = µeqj
						var µj *πg.Object = πg.UnboundLocal; _ = µj
						var µbj *πg.Object = πg.UnboundLocal; _ = µbj
						var µi *πg.Object = πg.UnboundLocal; _ = µi
						var µai *πg.Object = πg.UnboundLocal; _ = µai
						var µbest_i *πg.Object = πg.UnboundLocal; _ = µbest_i
						var µbest_j *πg.Object = πg.UnboundLocal; _ = µbest_j
						var µline *πg.Object = πg.UnboundLocal; _ = µline
						var µaelt *πg.Object = πg.UnboundLocal; _ = µaelt
						var µbelt *πg.Object = πg.UnboundLocal; _ = µbelt
						var µatags *πg.Object = πg.UnboundLocal; _ = µatags
						var µbtags *πg.Object = πg.UnboundLocal; _ = µbtags
						var µtag *πg.Object = πg.UnboundLocal; _ = µtag
						var µai1 *πg.Object = πg.UnboundLocal; _ = µai1
						var µai2 *πg.Object = πg.UnboundLocal; _ = µai2
						var µbj1 *πg.Object = πg.UnboundLocal; _ = µbj1
						var µbj2 *πg.Object = πg.UnboundLocal; _ = µbj2
						var µla *πg.Object = πg.UnboundLocal; _ = µla
						var µlb *πg.Object = πg.UnboundLocal; _ = µlb
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
						var πTemp008 bool
						_ = πTemp008
						var πTemp009 *πg.Object
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
								case 5: goto Label5
								case 39: goto Label39
								case 40: goto Label40
								case 42: goto Label42
								case 43: goto Label43
								case 44: goto Label44
								case 45: goto Label45
								case 47: goto Label47
								case 19: goto Label19
								case 20: goto Label20
								case 22: goto Label22
								case 23: goto Label23
								case 24: goto Label24
								case 26: goto Label26
								case 30: goto Label30
								case 31: goto Label31
								default: panic("unexpected function state")
								}
								// line 1007: r"""
								πF.SetLineno(1007)
								// line 1027: best_ratio, cutoff = 0.74, 0.75
								πF.SetLineno(1027)
								πTemp001 = πg.NewTuple2(πg.NewFloat(0.74).ToObject(), πg.NewFloat(0.75).ToObject()).ToObject()
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp003}}}, πTemp001); πE != nil {
									continue
								}
								µbest_ratio = πTemp002
								µcutoff = πTemp003
								// line 1028: cruncher = SequenceMatcher(self.charjunk)
								πF.SetLineno(1028)
								πTemp004 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.GetAttr(πF, µself, ßcharjunk, nil); πE != nil {
									continue
								}
								πTemp004[0] = πTemp001
								if πTemp001, πE = πg.ResolveGlobal(πF, ßSequenceMatcher); πE != nil {
									continue
								}
								if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								µcruncher = πTemp002
								// line 1029: eqi, eqj = None, None   # 1st indices of equal lines (if any)
								πF.SetLineno(1029)
								if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
									continue
								}
								if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
									continue
								}
								πTemp001 = πg.NewTuple2(πTemp002, πTemp003).ToObject()
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp003}}}, πTemp001); πE != nil {
									continue
								}
								µeqi = πTemp002
								µeqj = πTemp003
								πTemp004 = πF.MakeArgs(2)
								if πE = πg.CheckLocal(πF, µblo, "blo"); πE != nil {
									continue
								}
								πTemp004[0] = µblo
								if πE = πg.CheckLocal(πF, µbhi, "bhi"); πE != nil {
									continue
								}
								πTemp004[1] = µbhi
								if πTemp002, πE = πg.ResolveGlobal(πF, ßxrange); πE != nil {
									continue
								}
								if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
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
									µj = πTemp002
								}
								if πE != nil || !πTemp006 {
									continue
								}
								πF.PushCheckpoint(1)            
								// line 1035: bj = b[j]
								πF.SetLineno(1035)
								if πE = πg.CheckLocal(πF, µj, "j"); πE != nil {
									continue
								}
								πTemp002 = µj
								if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.GetItem(πF, µb, πTemp002); πE != nil {
									continue
								}
								µbj = πTemp003
								// line 1036: cruncher.set_seq2(bj)
								πF.SetLineno(1036)
								πTemp004 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µbj, "bj"); πE != nil {
									continue
								}
								πTemp004[0] = µbj
								if πE = πg.CheckLocal(πF, µcruncher, "cruncher"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.GetAttr(πF, µcruncher, ßset_seq2, nil); πE != nil {
									continue
								}
								if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								πTemp004 = πF.MakeArgs(2)
								if πE = πg.CheckLocal(πF, µalo, "alo"); πE != nil {
									continue
								}
								πTemp004[0] = µalo
								if πE = πg.CheckLocal(πF, µahi, "ahi"); πE != nil {
									continue
								}
								πTemp004[1] = µahi
								if πTemp003, πE = πg.ResolveGlobal(πF, ßxrange); πE != nil {
									continue
								}
								if πTemp007, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								if πTemp002, πE = πg.Iter(πF, πTemp007); πE != nil {
									continue
								}
								πF.PushCheckpoint(5)
								πTemp006 = false
							Label4:
								if πE != nil || πR != nil {
									continue
								}
								if πTemp006 {
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
									πTemp008 = !isStop
								} else {
									πTemp008 = true
									µi = πTemp003
								}
								if πE != nil || !πTemp008 {
									continue
								}
								πF.PushCheckpoint(4)            
								// line 1038: ai = a[i]
								πF.SetLineno(1038)
								if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
									continue
								}
								πTemp003 = µi
								if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
									continue
								}
								if πTemp007, πE = πg.GetItem(πF, µa, πTemp003); πE != nil {
									continue
								}
								µai = πTemp007
								if πE = πg.CheckLocal(πF, µai, "ai"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µbj, "bj"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.Eq(πF, µai, µbj); πE != nil {
									continue
								}
								if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
									continue
								}
								if πTemp008 {
									goto Label7
								}
								goto Label8
								// line 1039: if ai == bj:
								πF.SetLineno(1039)
							Label7:
								if πE = πg.CheckLocal(πF, µeqi, "eqi"); πE != nil {
									continue
								}
								if πTemp007, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
									continue
								}
								πTemp003 = πg.GetBool(µeqi == πTemp007).ToObject()
								if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
									continue
								}
								if πTemp008 {
									goto Label9
								}
								goto Label10
								// line 1040: if eqi is None:
								πF.SetLineno(1040)
							Label9:
								// line 1041: eqi, eqj = i, j
								πF.SetLineno(1041)
								if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µj, "j"); πE != nil {
									continue
								}
								πTemp003 = πg.NewTuple2(µi, µj).ToObject()
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp007}, πg.TieTarget{Target: &πTemp009}}}, πTemp003); πE != nil {
									continue
								}
								µeqi = πTemp007
								µeqj = πTemp009
								goto Label10
							Label10:
								// line 1042: continue
								πF.SetLineno(1042)
								continue
								goto Label8
							Label8:
								// line 1043: cruncher.set_seq1(ai)
								πF.SetLineno(1043)
								πTemp004 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µai, "ai"); πE != nil {
									continue
								}
								πTemp004[0] = µai
								if πE = πg.CheckLocal(πF, µcruncher, "cruncher"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.GetAttr(πF, µcruncher, ßset_seq1, nil); πE != nil {
									continue
								}
								if πTemp007, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								if πE = πg.CheckLocal(πF, µcruncher, "cruncher"); πE != nil {
									continue
								}
								if πTemp009, πE = πg.GetAttr(πF, µcruncher, ßreal_quick_ratio, nil); πE != nil {
									continue
								}
								if πTemp010, πE = πTemp009.Call(πF, nil, nil); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µbest_ratio, "best_ratio"); πE != nil {
									continue
								}
								if πTemp007, πE = πg.GT(πF, πTemp010, µbest_ratio); πE != nil {
									continue
								}
								πTemp003 = πTemp007
								if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
									continue
								}
								if !πTemp008 {
									goto Label11
								}
								if πE = πg.CheckLocal(πF, µcruncher, "cruncher"); πE != nil {
									continue
								}
								if πTemp009, πE = πg.GetAttr(πF, µcruncher, ßquick_ratio, nil); πE != nil {
									continue
								}
								if πTemp010, πE = πTemp009.Call(πF, nil, nil); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µbest_ratio, "best_ratio"); πE != nil {
									continue
								}
								if πTemp007, πE = πg.GT(πF, πTemp010, µbest_ratio); πE != nil {
									continue
								}
								πTemp003 = πTemp007
								if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
									continue
								}
								if !πTemp008 {
									goto Label11
								}
								if πE = πg.CheckLocal(πF, µcruncher, "cruncher"); πE != nil {
									continue
								}
								if πTemp009, πE = πg.GetAttr(πF, µcruncher, ßratio, nil); πE != nil {
									continue
								}
								if πTemp010, πE = πTemp009.Call(πF, nil, nil); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µbest_ratio, "best_ratio"); πE != nil {
									continue
								}
								if πTemp007, πE = πg.GT(πF, πTemp010, µbest_ratio); πE != nil {
									continue
								}
								πTemp003 = πTemp007
							Label11:
								if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
									continue
								}
								if πTemp008 {
									goto Label12
								}
								goto Label13
								// line 1050: if cruncher.real_quick_ratio() > best_ratio and \
								πF.SetLineno(1050)
							Label12:
								// line 1053: best_ratio, best_i, best_j = cruncher.ratio(), i, j
								πF.SetLineno(1053)
								if πE = πg.CheckLocal(πF, µcruncher, "cruncher"); πE != nil {
									continue
								}
								if πTemp007, πE = πg.GetAttr(πF, µcruncher, ßratio, nil); πE != nil {
									continue
								}
								if πTemp009, πE = πTemp007.Call(πF, nil, nil); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µj, "j"); πE != nil {
									continue
								}
								πTemp003 = πg.NewTuple3(πTemp009, µi, µj).ToObject()
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp007}, πg.TieTarget{Target: &πTemp009}, πg.TieTarget{Target: &πTemp010}}}, πTemp003); πE != nil {
									continue
								}
								µbest_ratio = πTemp007
								µbest_i = πTemp009
								µbest_j = πTemp010
								goto Label13
							Label13:
								continue
							Label5:
								if πE != nil || πR != nil {
									continue
								}
							Label6:
								continue
							Label2:
								if πE != nil || πR != nil {
									continue
								}
							Label3:
								if πE = πg.CheckLocal(πF, µbest_ratio, "best_ratio"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µcutoff, "cutoff"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.LT(πF, µbest_ratio, µcutoff); πE != nil {
									continue
								}
								if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
									continue
								}
								if πTemp005 {
									goto Label14
								}
								goto Label15
								// line 1054: if best_ratio < cutoff:
								πF.SetLineno(1054)
							Label14:
								if πE = πg.CheckLocal(πF, µeqi, "eqi"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
									continue
								}
								πTemp001 = πg.GetBool(µeqi == πTemp002).ToObject()
								if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
									continue
								}
								if πTemp005 {
									goto Label17
								}
								goto Label18
								// line 1056: if eqi is None:
								πF.SetLineno(1056)
							Label17:
								πTemp004 = πF.MakeArgs(6)
								if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
									continue
								}
								πTemp004[0] = µa
								if πE = πg.CheckLocal(πF, µalo, "alo"); πE != nil {
									continue
								}
								πTemp004[1] = µalo
								if πE = πg.CheckLocal(πF, µahi, "ahi"); πE != nil {
									continue
								}
								πTemp004[2] = µahi
								if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
									continue
								}
								πTemp004[3] = µb
								if πE = πg.CheckLocal(πF, µblo, "blo"); πE != nil {
									continue
								}
								πTemp004[4] = µblo
								if πE = πg.CheckLocal(πF, µbhi, "bhi"); πE != nil {
									continue
								}
								πTemp004[5] = µbhi
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.GetAttr(πF, µself, ß_plain_replace, nil); πE != nil {
									continue
								}
								if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								if πTemp001, πE = πg.Iter(πF, πTemp003); πE != nil {
									continue
								}
								πF.PushCheckpoint(20)
								πTemp005 = false
							Label19:
								if πE != nil || πR != nil {
									continue
								}
								if πTemp005 {
									πF.PopCheckpoint()
									goto Label21
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
								πF.PushCheckpoint(19)            
								// line 1059: yield line
								πF.SetLineno(1059)
								if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
									continue
								}
								πF.PushCheckpoint(22)
								return µline, nil
							Label22:
								πTemp002 = πSent
								continue
							Label20:
								if πE != nil || πR != nil {
									continue
								}
							Label21:
								// line 1060: return
								πF.SetLineno(1060)
								πR = πg.None
								continue
								goto Label18
							Label18:
								// line 1062: best_i, best_j, best_ratio = eqi, eqj, 1.0
								πF.SetLineno(1062)
								if πE = πg.CheckLocal(πF, µeqi, "eqi"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µeqj, "eqj"); πE != nil {
									continue
								}
								πTemp001 = πg.NewTuple3(µeqi, µeqj, πg.NewFloat(1.0).ToObject()).ToObject()
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp007}}}, πTemp001); πE != nil {
									continue
								}
								µbest_i = πTemp002
								µbest_j = πTemp003
								µbest_ratio = πTemp007
								goto Label16
							Label15:
								// line 1065: eqi = None
								πF.SetLineno(1065)
								if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
									continue
								}
								µeqi = πTemp001
								goto Label16
							Label16:
								πTemp004 = πF.MakeArgs(6)
								if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
									continue
								}
								πTemp004[0] = µa
								if πE = πg.CheckLocal(πF, µalo, "alo"); πE != nil {
									continue
								}
								πTemp004[1] = µalo
								if πE = πg.CheckLocal(πF, µbest_i, "best_i"); πE != nil {
									continue
								}
								πTemp004[2] = µbest_i
								if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
									continue
								}
								πTemp004[3] = µb
								if πE = πg.CheckLocal(πF, µblo, "blo"); πE != nil {
									continue
								}
								πTemp004[4] = µblo
								if πE = πg.CheckLocal(πF, µbest_j, "best_j"); πE != nil {
									continue
								}
								πTemp004[5] = µbest_j
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.GetAttr(πF, µself, ß_fancy_helper, nil); πE != nil {
									continue
								}
								if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								if πTemp001, πE = πg.Iter(πF, πTemp003); πE != nil {
									continue
								}
								πF.PushCheckpoint(24)
								πTemp005 = false
							Label23:
								if πE != nil || πR != nil {
									continue
								}
								if πTemp005 {
									πF.PopCheckpoint()
									goto Label25
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
								πF.PushCheckpoint(23)            
								// line 1072: yield line
								πF.SetLineno(1072)
								if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
									continue
								}
								πF.PushCheckpoint(26)
								return µline, nil
							Label26:
								πTemp002 = πSent
								continue
							Label24:
								if πE != nil || πR != nil {
									continue
								}
							Label25:
								// line 1075: aelt, belt = a[best_i], b[best_j]
								πF.SetLineno(1075)
								if πE = πg.CheckLocal(πF, µbest_i, "best_i"); πE != nil {
									continue
								}
								πTemp002 = µbest_i
								if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.GetItem(πF, µa, πTemp002); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µbest_j, "best_j"); πE != nil {
									continue
								}
								πTemp002 = µbest_j
								if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
									continue
								}
								if πTemp007, πE = πg.GetItem(πF, µb, πTemp002); πE != nil {
									continue
								}
								πTemp001 = πg.NewTuple2(πTemp003, πTemp007).ToObject()
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp003}}}, πTemp001); πE != nil {
									continue
								}
								µaelt = πTemp002
								µbelt = πTemp003
								if πE = πg.CheckLocal(πF, µeqi, "eqi"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
									continue
								}
								πTemp001 = πg.GetBool(µeqi == πTemp002).ToObject()
								if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
									continue
								}
								if πTemp005 {
									goto Label27
								}
								goto Label28
								// line 1076: if eqi is None:
								πF.SetLineno(1076)
							Label27:
								// line 1078: atags = btags = ""
								πF.SetLineno(1078)
								µatags = ß.ToObject()
								µbtags = ß.ToObject()
								// line 1079: cruncher.set_seqs(aelt, belt)
								πF.SetLineno(1079)
								πTemp004 = πF.MakeArgs(2)
								if πE = πg.CheckLocal(πF, µaelt, "aelt"); πE != nil {
									continue
								}
								πTemp004[0] = µaelt
								if πE = πg.CheckLocal(πF, µbelt, "belt"); πE != nil {
									continue
								}
								πTemp004[1] = µbelt
								if πE = πg.CheckLocal(πF, µcruncher, "cruncher"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.GetAttr(πF, µcruncher, ßset_seqs, nil); πE != nil {
									continue
								}
								if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								if πE = πg.CheckLocal(πF, µcruncher, "cruncher"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.GetAttr(πF, µcruncher, ßget_opcodes, nil); πE != nil {
									continue
								}
								if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
									continue
								}
								if πTemp001, πE = πg.Iter(πF, πTemp003); πE != nil {
									continue
								}
								πF.PushCheckpoint(31)
								πTemp005 = false
							Label30:
								if πE != nil || πR != nil {
									continue
								}
								if πTemp005 {
									πF.PopCheckpoint()
									goto Label32
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
									if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp007}, πg.TieTarget{Target: &πTemp009}, πg.TieTarget{Target: &πTemp010}, πg.TieTarget{Target: &πTemp011}}}, πTemp002); πE != nil {
										continue
									}
									µtag = πTemp003
									µai1 = πTemp007
									µai2 = πTemp009
									µbj1 = πTemp010
									µbj2 = πTemp011
								}
								if πE != nil || !πTemp006 {
									continue
								}
								πF.PushCheckpoint(30)            
								// line 1081: la, lb = ai2 - ai1, bj2 - bj1
								πF.SetLineno(1081)
								if πE = πg.CheckLocal(πF, µai2, "ai2"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µai1, "ai1"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.Sub(πF, µai2, µai1); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µbj2, "bj2"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µbj1, "bj1"); πE != nil {
									continue
								}
								if πTemp007, πE = πg.Sub(πF, µbj2, µbj1); πE != nil {
									continue
								}
								πTemp002 = πg.NewTuple2(πTemp003, πTemp007).ToObject()
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp007}}}, πTemp002); πE != nil {
									continue
								}
								µla = πTemp003
								µlb = πTemp007
								if πE = πg.CheckLocal(πF, µtag, "tag"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.Eq(πF, µtag, ßreplace.ToObject()); πE != nil {
									continue
								}
								if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
									continue
								}
								if πTemp006 {
									goto Label33
								}
								if πE = πg.CheckLocal(πF, µtag, "tag"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.Eq(πF, µtag, ßdelete.ToObject()); πE != nil {
									continue
								}
								if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
									continue
								}
								if πTemp006 {
									goto Label34
								}
								if πE = πg.CheckLocal(πF, µtag, "tag"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.Eq(πF, µtag, ßinsert.ToObject()); πE != nil {
									continue
								}
								if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
									continue
								}
								if πTemp006 {
									goto Label35
								}
								if πE = πg.CheckLocal(πF, µtag, "tag"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.Eq(πF, µtag, ßequal.ToObject()); πE != nil {
									continue
								}
								if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
									continue
								}
								if πTemp006 {
									goto Label36
								}
								goto Label37
								// line 1082: if tag == 'replace':
								πF.SetLineno(1082)
							Label33:
								// line 1083: atags += '^' * la
								πF.SetLineno(1083)
								if πE = πg.CheckLocal(πF, µatags, "atags"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µla, "la"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.Mul(πF, πg.NewStr("^").ToObject(), µla); πE != nil {
									continue
								}
								if πTemp003, πE = πg.IAdd(πF, µatags, πTemp002); πE != nil {
									continue
								}
								µatags = πTemp003
								// line 1084: btags += '^' * lb
								πF.SetLineno(1084)
								if πE = πg.CheckLocal(πF, µbtags, "btags"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µlb, "lb"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.Mul(πF, πg.NewStr("^").ToObject(), µlb); πE != nil {
									continue
								}
								if πTemp003, πE = πg.IAdd(πF, µbtags, πTemp002); πE != nil {
									continue
								}
								µbtags = πTemp003
								goto Label38
								// line 1085: elif tag == 'delete':
								πF.SetLineno(1085)
							Label34:
								// line 1086: atags += '-' * la
								πF.SetLineno(1086)
								if πE = πg.CheckLocal(πF, µatags, "atags"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µla, "la"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.Mul(πF, πg.NewStr("-").ToObject(), µla); πE != nil {
									continue
								}
								if πTemp003, πE = πg.IAdd(πF, µatags, πTemp002); πE != nil {
									continue
								}
								µatags = πTemp003
								goto Label38
								// line 1087: elif tag == 'insert':
								πF.SetLineno(1087)
							Label35:
								// line 1088: btags += '+' * lb
								πF.SetLineno(1088)
								if πE = πg.CheckLocal(πF, µbtags, "btags"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µlb, "lb"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.Mul(πF, πg.NewStr("+").ToObject(), µlb); πE != nil {
									continue
								}
								if πTemp003, πE = πg.IAdd(πF, µbtags, πTemp002); πE != nil {
									continue
								}
								µbtags = πTemp003
								goto Label38
								// line 1089: elif tag == 'equal':
								πF.SetLineno(1089)
							Label36:
								// line 1090: atags += ' ' * la
								πF.SetLineno(1090)
								if πE = πg.CheckLocal(πF, µatags, "atags"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µla, "la"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.Mul(πF, πg.NewStr(" ").ToObject(), µla); πE != nil {
									continue
								}
								if πTemp003, πE = πg.IAdd(πF, µatags, πTemp002); πE != nil {
									continue
								}
								µatags = πTemp003
								// line 1091: btags += ' ' * lb
								πF.SetLineno(1091)
								if πE = πg.CheckLocal(πF, µbtags, "btags"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µlb, "lb"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.Mul(πF, πg.NewStr(" ").ToObject(), µlb); πE != nil {
									continue
								}
								if πTemp003, πE = πg.IAdd(πF, µbtags, πTemp002); πE != nil {
									continue
								}
								µbtags = πTemp003
								goto Label38
							Label37:
								if πTemp002, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µtag, "tag"); πE != nil {
									continue
								}
								πTemp007 = πg.NewTuple1(µtag).ToObject()
								if πTemp003, πE = πg.Mod(πF, πg.NewStr("unknown tag %r").ToObject(), πTemp007); πE != nil {
									continue
								}
								// line 1093: raise ValueError, 'unknown tag %r' % (tag,)
								πF.SetLineno(1093)
								πE = πF.Raise(πTemp002, πTemp003, nil)
								continue
								goto Label38
							Label38:
								continue
							Label31:
								if πE != nil || πR != nil {
									continue
								}
							Label32:
								πTemp004 = πF.MakeArgs(4)
								if πE = πg.CheckLocal(πF, µaelt, "aelt"); πE != nil {
									continue
								}
								πTemp004[0] = µaelt
								if πE = πg.CheckLocal(πF, µbelt, "belt"); πE != nil {
									continue
								}
								πTemp004[1] = µbelt
								if πE = πg.CheckLocal(πF, µatags, "atags"); πE != nil {
									continue
								}
								πTemp004[2] = µatags
								if πE = πg.CheckLocal(πF, µbtags, "btags"); πE != nil {
									continue
								}
								πTemp004[3] = µbtags
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.GetAttr(πF, µself, ß_qformat, nil); πE != nil {
									continue
								}
								if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								if πTemp001, πE = πg.Iter(πF, πTemp003); πE != nil {
									continue
								}
								πF.PushCheckpoint(40)
								πTemp005 = false
							Label39:
								if πE != nil || πR != nil {
									continue
								}
								if πTemp005 {
									πF.PopCheckpoint()
									goto Label41
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
								πF.PushCheckpoint(39)            
								// line 1095: yield line
								πF.SetLineno(1095)
								if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
									continue
								}
								πF.PushCheckpoint(42)
								return µline, nil
							Label42:
								πTemp002 = πSent
								continue
							Label40:
								if πE != nil || πR != nil {
									continue
								}
							Label41:
								goto Label29
							Label28:
								// line 1098: yield '  ' + aelt
								πF.SetLineno(1098)
								if πE = πg.CheckLocal(πF, µaelt, "aelt"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.Add(πF, πg.NewStr("  ").ToObject(), µaelt); πE != nil {
									continue
								}
								πF.PushCheckpoint(43)
								return πTemp001, nil
							Label43:
								πTemp002 = πSent
								goto Label29
							Label29:
								πTemp004 = πF.MakeArgs(6)
								if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
									continue
								}
								πTemp004[0] = µa
								if πE = πg.CheckLocal(πF, µbest_i, "best_i"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.Add(πF, µbest_i, πg.NewInt(1).ToObject()); πE != nil {
									continue
								}
								πTemp004[1] = πTemp003
								if πE = πg.CheckLocal(πF, µahi, "ahi"); πE != nil {
									continue
								}
								πTemp004[2] = µahi
								if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
									continue
								}
								πTemp004[3] = µb
								if πE = πg.CheckLocal(πF, µbest_j, "best_j"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.Add(πF, µbest_j, πg.NewInt(1).ToObject()); πE != nil {
									continue
								}
								πTemp004[4] = πTemp003
								if πE = πg.CheckLocal(πF, µbhi, "bhi"); πE != nil {
									continue
								}
								πTemp004[5] = µbhi
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.GetAttr(πF, µself, ß_fancy_helper, nil); πE != nil {
									continue
								}
								if πTemp007, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								if πTemp002, πE = πg.Iter(πF, πTemp007); πE != nil {
									continue
								}
								πF.PushCheckpoint(45)
								πTemp005 = false
							Label44:
								if πE != nil || πR != nil {
									continue
								}
								if πTemp005 {
									πF.PopCheckpoint()
									goto Label46
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
									µline = πTemp003
								}
								if πE != nil || !πTemp006 {
									continue
								}
								πF.PushCheckpoint(44)            
								// line 1102: yield line
								πF.SetLineno(1102)
								if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
									continue
								}
								πF.PushCheckpoint(47)
								return µline, nil
							Label47:
								πTemp003 = πSent
								continue
							Label45:
								if πE != nil || πR != nil {
									continue
								}
							Label46:
							}
							return nil, πE
						}).ToObject(), nil
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß_fancy_replace.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 1007: r"""
					πF.SetLineno(1007)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp007}, πg.NewStr("\n        When replacing one block of lines with another, search the blocks\n        for *similar* lines; the best-matching pair (if any) is used as a\n        synch point, and intraline difference marking is done on the\n        similar pair. Lots of work, but often worth it.\n\n        Example:\n\n        >>> d = Differ()\n        >>> results = d._fancy_replace(['abcDefghiJkl\\n'], 0, 1,\n        ...                            ['abcdefGhijkl\\n'], 0, 1)\n        >>> print ''.join(results),\n        - abcDefghiJkl\n        ?    ^  ^  ^\n        + abcdefGhijkl\n        ?    ^  ^  ^\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ß_fancy_replace); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp008, ß__doc__, πTemp007); πE != nil {
						continue
					}
					// line 1104: def _fancy_helper(self, a, alo, ahi, b, blo, bhi):
					πF.SetLineno(1104)
					πTemp002 = make([]πg.Param, 7)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "a", Def: nil}
					πTemp002[2] = πg.Param{Name: "alo", Def: nil}
					πTemp002[3] = πg.Param{Name: "ahi", Def: nil}
					πTemp002[4] = πg.Param{Name: "b", Def: nil}
					πTemp002[5] = πg.Param{Name: "blo", Def: nil}
					πTemp002[6] = πg.Param{Name: "bhi", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("_fancy_helper", "/usr/lib/python2.7/difflib.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µa *πg.Object = πArgs[1]; _ = µa
						var µalo *πg.Object = πArgs[2]; _ = µalo
						var µahi *πg.Object = πArgs[3]; _ = µahi
						var µb *πg.Object = πArgs[4]; _ = µb
						var µblo *πg.Object = πArgs[5]; _ = µblo
						var µbhi *πg.Object = πArgs[6]; _ = µbhi
						var µg *πg.Object = πg.UnboundLocal; _ = µg
						var µline *πg.Object = πg.UnboundLocal; _ = µline
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 bool
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 bool
						_ = πTemp005
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
							for ; πF.State() >= 0; πF.PopCheckpoint() {
								switch πF.State() {
								case 0:
								case 8: goto Label8
								case 10: goto Label10
								case 7: goto Label7
								default: panic("unexpected function state")
								}
								// line 1105: g = []
								πF.SetLineno(1105)
								πTemp001 = make([]*πg.Object, 0)
								πTemp002 = πg.NewList(πTemp001...).ToObject()
								µg = πTemp002
								if πE = πg.CheckLocal(πF, µalo, "alo"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µahi, "ahi"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.LT(πF, µalo, µahi); πE != nil {
									continue
								}
								if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
									continue
								}
								if πTemp003 {
									goto Label1
								}
								if πE = πg.CheckLocal(πF, µblo, "blo"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µbhi, "bhi"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.LT(πF, µblo, µbhi); πE != nil {
									continue
								}
								if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
									continue
								}
								if πTemp003 {
									goto Label2
								}
								goto Label3
								// line 1106: if alo < ahi:
								πF.SetLineno(1106)
							Label1:
								if πE = πg.CheckLocal(πF, µblo, "blo"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µbhi, "bhi"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.LT(πF, µblo, µbhi); πE != nil {
									continue
								}
								if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
									continue
								}
								if πTemp003 {
									goto Label4
								}
								goto Label5
								// line 1107: if blo < bhi:
								πF.SetLineno(1107)
							Label4:
								// line 1108: g = self._fancy_replace(a, alo, ahi, b, blo, bhi)
								πF.SetLineno(1108)
								πTemp001 = πF.MakeArgs(6)
								if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
									continue
								}
								πTemp001[0] = µa
								if πE = πg.CheckLocal(πF, µalo, "alo"); πE != nil {
									continue
								}
								πTemp001[1] = µalo
								if πE = πg.CheckLocal(πF, µahi, "ahi"); πE != nil {
									continue
								}
								πTemp001[2] = µahi
								if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
									continue
								}
								πTemp001[3] = µb
								if πE = πg.CheckLocal(πF, µblo, "blo"); πE != nil {
									continue
								}
								πTemp001[4] = µblo
								if πE = πg.CheckLocal(πF, µbhi, "bhi"); πE != nil {
									continue
								}
								πTemp001[5] = µbhi
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.GetAttr(πF, µself, ß_fancy_replace, nil); πE != nil {
									continue
								}
								if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								µg = πTemp004
								goto Label6
							Label5:
								// line 1110: g = self._dump('-', a, alo, ahi)
								πF.SetLineno(1110)
								πTemp001 = πF.MakeArgs(4)
								πTemp001[0] = πg.NewStr("-").ToObject()
								if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
									continue
								}
								πTemp001[1] = µa
								if πE = πg.CheckLocal(πF, µalo, "alo"); πE != nil {
									continue
								}
								πTemp001[2] = µalo
								if πE = πg.CheckLocal(πF, µahi, "ahi"); πE != nil {
									continue
								}
								πTemp001[3] = µahi
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.GetAttr(πF, µself, ß_dump, nil); πE != nil {
									continue
								}
								if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								µg = πTemp004
								goto Label6
							Label6:
								goto Label3
								// line 1111: elif blo < bhi:
								πF.SetLineno(1111)
							Label2:
								// line 1112: g = self._dump('+', b, blo, bhi)
								πF.SetLineno(1112)
								πTemp001 = πF.MakeArgs(4)
								πTemp001[0] = πg.NewStr("+").ToObject()
								if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
									continue
								}
								πTemp001[1] = µb
								if πE = πg.CheckLocal(πF, µblo, "blo"); πE != nil {
									continue
								}
								πTemp001[2] = µblo
								if πE = πg.CheckLocal(πF, µbhi, "bhi"); πE != nil {
									continue
								}
								πTemp001[3] = µbhi
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.GetAttr(πF, µself, ß_dump, nil); πE != nil {
									continue
								}
								if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								µg = πTemp004
								goto Label3
							Label3:
								if πE = πg.CheckLocal(πF, µg, "g"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.Iter(πF, µg); πE != nil {
									continue
								}
								πF.PushCheckpoint(8)
								πTemp003 = false
							Label7:
								if πE != nil || πR != nil {
									continue
								}
								if πTemp003 {
									πF.PopCheckpoint()
									goto Label9
								}
								if πTemp004, πE = πg.Next(πF, πTemp002); πE != nil {
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
									µline = πTemp004
								}
								if πE != nil || !πTemp005 {
									continue
								}
								πF.PushCheckpoint(7)            
								// line 1115: yield line
								πF.SetLineno(1115)
								if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
									continue
								}
								πF.PushCheckpoint(10)
								return µline, nil
							Label10:
								πTemp004 = πSent
								continue
							Label8:
								if πE != nil || πR != nil {
									continue
								}
							Label9:
							}
							return nil, πE
						}).ToObject(), nil
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß_fancy_helper.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 1117: def _qformat(self, aline, bline, atags, btags):
					πF.SetLineno(1117)
					πTemp002 = make([]πg.Param, 5)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "aline", Def: nil}
					πTemp002[2] = πg.Param{Name: "bline", Def: nil}
					πTemp002[3] = πg.Param{Name: "atags", Def: nil}
					πTemp002[4] = πg.Param{Name: "btags", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("_qformat", "/usr/lib/python2.7/difflib.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µaline *πg.Object = πArgs[1]; _ = µaline
						var µbline *πg.Object = πArgs[2]; _ = µbline
						var µatags *πg.Object = πArgs[3]; _ = µatags
						var µbtags *πg.Object = πArgs[4]; _ = µbtags
						var µcommon *πg.Object = πg.UnboundLocal; _ = µcommon
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
							for ; πF.State() >= 0; πF.PopCheckpoint() {
								switch πF.State() {
								case 0:
								case 8: goto Label8
								case 1: goto Label1
								case 4: goto Label4
								case 5: goto Label5
								default: panic("unexpected function state")
								}
								// line 1118: r"""
								πF.SetLineno(1118)
								// line 1135: common = min(_count_leading(aline, "\t"),
								πF.SetLineno(1135)
								πTemp001 = πF.MakeArgs(2)
								πTemp002 = πF.MakeArgs(2)
								if πE = πg.CheckLocal(πF, µaline, "aline"); πE != nil {
									continue
								}
								πTemp002[0] = µaline
								πTemp002[1] = πg.NewStr("\t").ToObject()
								if πTemp003, πE = πg.ResolveGlobal(πF, ß_count_leading); πE != nil {
									continue
								}
								if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp002)
								πTemp001[0] = πTemp004
								πTemp002 = πF.MakeArgs(2)
								if πE = πg.CheckLocal(πF, µbline, "bline"); πE != nil {
									continue
								}
								πTemp002[0] = µbline
								πTemp002[1] = πg.NewStr("\t").ToObject()
								if πTemp003, πE = πg.ResolveGlobal(πF, ß_count_leading); πE != nil {
									continue
								}
								if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp002)
								πTemp001[1] = πTemp004
								if πTemp003, πE = πg.ResolveGlobal(πF, ßmin); πE != nil {
									continue
								}
								if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								µcommon = πTemp004
								// line 1137: common = min(common, _count_leading(atags[:common], " "))
								πF.SetLineno(1137)
								πTemp001 = πF.MakeArgs(2)
								if πE = πg.CheckLocal(πF, µcommon, "common"); πE != nil {
									continue
								}
								πTemp001[0] = µcommon
								πTemp002 = πF.MakeArgs(2)
								if πE = πg.CheckLocal(πF, µcommon, "common"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.None, µcommon, πg.None}, nil); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µatags, "atags"); πE != nil {
									continue
								}
								if πTemp004, πE = πg.GetItem(πF, µatags, πTemp003); πE != nil {
									continue
								}
								πTemp002[0] = πTemp004
								πTemp002[1] = πg.NewStr(" ").ToObject()
								if πTemp003, πE = πg.ResolveGlobal(πF, ß_count_leading); πE != nil {
									continue
								}
								if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp002)
								πTemp001[1] = πTemp004
								if πTemp003, πE = πg.ResolveGlobal(πF, ßmin); πE != nil {
									continue
								}
								if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								µcommon = πTemp004
								// line 1138: common = min(common, _count_leading(btags[:common], " "))
								πF.SetLineno(1138)
								πTemp001 = πF.MakeArgs(2)
								if πE = πg.CheckLocal(πF, µcommon, "common"); πE != nil {
									continue
								}
								πTemp001[0] = µcommon
								πTemp002 = πF.MakeArgs(2)
								if πE = πg.CheckLocal(πF, µcommon, "common"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.None, µcommon, πg.None}, nil); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µbtags, "btags"); πE != nil {
									continue
								}
								if πTemp004, πE = πg.GetItem(πF, µbtags, πTemp003); πE != nil {
									continue
								}
								πTemp002[0] = πTemp004
								πTemp002[1] = πg.NewStr(" ").ToObject()
								if πTemp003, πE = πg.ResolveGlobal(πF, ß_count_leading); πE != nil {
									continue
								}
								if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp002)
								πTemp001[1] = πTemp004
								if πTemp003, πE = πg.ResolveGlobal(πF, ßmin); πE != nil {
									continue
								}
								if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								µcommon = πTemp004
								// line 1139: atags = atags[common:].rstrip()
								πF.SetLineno(1139)
								if πE = πg.CheckLocal(πF, µcommon, "common"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{µcommon, πg.None, πg.None}, nil); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µatags, "atags"); πE != nil {
									continue
								}
								if πTemp004, πE = πg.GetItem(πF, µatags, πTemp003); πE != nil {
									continue
								}
								if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßrstrip, nil); πE != nil {
									continue
								}
								if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
									continue
								}
								µatags = πTemp004
								// line 1140: btags = btags[common:].rstrip()
								πF.SetLineno(1140)
								if πE = πg.CheckLocal(πF, µcommon, "common"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{µcommon, πg.None, πg.None}, nil); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µbtags, "btags"); πE != nil {
									continue
								}
								if πTemp004, πE = πg.GetItem(πF, µbtags, πTemp003); πE != nil {
									continue
								}
								if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßrstrip, nil); πE != nil {
									continue
								}
								if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
									continue
								}
								µbtags = πTemp004
								// line 1142: yield "- " + aline
								πF.SetLineno(1142)
								if πE = πg.CheckLocal(πF, µaline, "aline"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.Add(πF, πg.NewStr("- ").ToObject(), µaline); πE != nil {
									continue
								}
								πF.PushCheckpoint(1)
								return πTemp003, nil
							Label1:
								πTemp004 = πSent
								if πE = πg.CheckLocal(πF, µatags, "atags"); πE != nil {
									continue
								}
								if πTemp005, πE = πg.IsTrue(πF, µatags); πE != nil {
									continue
								}
								if πTemp005 {
									goto Label2
								}
								goto Label3
								// line 1143: if atags:
								πF.SetLineno(1143)
							Label2:
								// line 1144: yield "? %s%s\n" % ("\t" * common, atags)
								πF.SetLineno(1144)
								if πE = πg.CheckLocal(πF, µcommon, "common"); πE != nil {
									continue
								}
								if πTemp007, πE = πg.Mul(πF, πg.NewStr("\t").ToObject(), µcommon); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µatags, "atags"); πE != nil {
									continue
								}
								πTemp006 = πg.NewTuple2(πTemp007, µatags).ToObject()
								if πTemp004, πE = πg.Mod(πF, πg.NewStr("? %s%s\n").ToObject(), πTemp006); πE != nil {
									continue
								}
								πF.PushCheckpoint(4)
								return πTemp004, nil
							Label4:
								πTemp006 = πSent
								goto Label3
							Label3:
								// line 1146: yield "+ " + bline
								πF.SetLineno(1146)
								if πE = πg.CheckLocal(πF, µbline, "bline"); πE != nil {
									continue
								}
								if πTemp006, πE = πg.Add(πF, πg.NewStr("+ ").ToObject(), µbline); πE != nil {
									continue
								}
								πF.PushCheckpoint(5)
								return πTemp006, nil
							Label5:
								πTemp007 = πSent
								if πE = πg.CheckLocal(πF, µbtags, "btags"); πE != nil {
									continue
								}
								if πTemp005, πE = πg.IsTrue(πF, µbtags); πE != nil {
									continue
								}
								if πTemp005 {
									goto Label6
								}
								goto Label7
								// line 1147: if btags:
								πF.SetLineno(1147)
							Label6:
								// line 1148: yield "? %s%s\n" % ("\t" * common, btags)
								πF.SetLineno(1148)
								if πE = πg.CheckLocal(πF, µcommon, "common"); πE != nil {
									continue
								}
								if πTemp009, πE = πg.Mul(πF, πg.NewStr("\t").ToObject(), µcommon); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µbtags, "btags"); πE != nil {
									continue
								}
								πTemp008 = πg.NewTuple2(πTemp009, µbtags).ToObject()
								if πTemp007, πE = πg.Mod(πF, πg.NewStr("? %s%s\n").ToObject(), πTemp008); πE != nil {
									continue
								}
								πF.PushCheckpoint(8)
								return πTemp007, nil
							Label8:
								πTemp008 = πSent
								goto Label7
							Label7:
							}
							return nil, πE
						}).ToObject(), nil
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß_qformat.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 1118: r"""
					πF.SetLineno(1118)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp009}, πg.NewStr("\n        Format \"?\" output and deal with leading tabs.\n\n        Example:\n\n        >>> d = Differ()\n        >>> results = d._qformat('\\tabcDefghiJkl\\n', '\\tabcdefGhijkl\\n',\n        ...                      '  ^ ^  ^      ', '  ^ ^  ^      ')\n        >>> for line in results: print repr(line)\n        ...\n        '- \\tabcDefghiJkl\\n'\n        '? \\t ^ ^  ^\\n'\n        '+ \\tabcdefGhijkl\\n'\n        '? \\t ^ ^  ^\\n'\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ß_qformat); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp010, ß__doc__, πTemp009); πE != nil {
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
			if πTemp010, πE = πTemp009.Call(πF, []*πg.Object{πg.NewStr("Differ").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp005.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßDiffer.ToObject(), πTemp010); πE != nil {
				continue
			}
			// line 1167: import re
			πF.SetLineno(1167)
			if πTemp001, πE = πg.ImportModule(πF, "re"); πE != nil {
				continue
			}
			πTemp008 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ßre.ToObject(), πTemp008); πE != nil {
				continue
			}
			// line 1169: def IS_LINE_JUNK(line, pat=re.compile(r"\s*#?\s*$").match):
			πF.SetLineno(1169)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "line", Def: nil}
			πTemp001 = πF.MakeArgs(1)
			πTemp001[0] = πg.NewStr("\\s*#?\\s*$").ToObject()
			if πTemp009, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
				continue
			}
			if πTemp010, πE = πg.GetAttr(πF, πTemp009, ßcompile, nil); πE != nil {
				continue
			}
			if πTemp009, πE = πTemp010.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πTemp010, πE = πg.GetAttr(πF, πTemp009, ßmatch, nil); πE != nil {
				continue
			}
			πTemp004[1] = πg.Param{Name: "pat", Def: πTemp010}
			πTemp008 = πg.NewFunction(πg.NewCode("IS_LINE_JUNK", "/usr/lib/python2.7/difflib.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µline *πg.Object = πArgs[0]; _ = µline
				var µpat *πg.Object = πArgs[1]; _ = µpat
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
					// line 1170: r"""
					πF.SetLineno(1170)
					// line 1183: return pat(line) is not None
					πF.SetLineno(1183)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					πTemp002[0] = µline
					if πE = πg.CheckLocal(πF, µpat, "pat"); πE != nil {
						continue
					}
					if πTemp003, πE = µpat.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(πTemp003 != πTemp004).ToObject()
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
			if πE = πF.Globals().SetItem(πF, ßIS_LINE_JUNK.ToObject(), πTemp008); πE != nil {
				continue
			}
			// line 1170: r"""
			πF.SetLineno(1170)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp009}, πg.NewStr("\n    Return 1 for ignorable line: iff `line` is blank or contains a single '#'.\n\n    Examples:\n\n    >>> IS_LINE_JUNK('\\n')\n    True\n    >>> IS_LINE_JUNK('  #   \\n')\n    True\n    >>> IS_LINE_JUNK('hello\\n')\n    False\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp010, πE = πg.ResolveGlobal(πF, ßIS_LINE_JUNK); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp010, ß__doc__, πTemp009); πE != nil {
				continue
			}
			// line 1185: def IS_CHARACTER_JUNK(ch, ws=" \t"):
			πF.SetLineno(1185)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "ch", Def: nil}
			πTemp004[1] = πg.Param{Name: "ws", Def: πg.NewStr(" \t").ToObject()}
			πTemp009 = πg.NewFunction(πg.NewCode("IS_CHARACTER_JUNK", "/usr/lib/python2.7/difflib.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µch *πg.Object = πArgs[0]; _ = µch
				var µws *πg.Object = πArgs[1]; _ = µws
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 bool
				_ = πTemp002
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 1186: r"""
					πF.SetLineno(1186)
					// line 1201: return ch in ws
					πF.SetLineno(1201)
					if πE = πg.CheckLocal(πF, µch, "ch"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µws, "ws"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Contains(πF, µws, µch); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(πTemp002).ToObject()
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
			if πE = πF.Globals().SetItem(πF, ßIS_CHARACTER_JUNK.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1186: r"""
			πF.SetLineno(1186)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp010}, πg.NewStr("\n    Return 1 for ignorable character: iff `ch` is a space or tab.\n\n    Examples:\n\n    >>> IS_CHARACTER_JUNK(' ')\n    True\n    >>> IS_CHARACTER_JUNK('\\t')\n    True\n    >>> IS_CHARACTER_JUNK('\\n')\n    False\n    >>> IS_CHARACTER_JUNK('x')\n    False\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp011, πE = πg.ResolveGlobal(πF, ßIS_CHARACTER_JUNK); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp011, ß__doc__, πTemp010); πE != nil {
				continue
			}
			// line 1208: def _format_range_unified(start, stop):
			πF.SetLineno(1208)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "start", Def: nil}
			πTemp004[1] = πg.Param{Name: "stop", Def: nil}
			πTemp010 = πg.NewFunction(πg.NewCode("_format_range_unified", "/usr/lib/python2.7/difflib.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µstart *πg.Object = πArgs[0]; _ = µstart
				var µstop *πg.Object = πArgs[1]; _ = µstop
				var µbeginning *πg.Object = πg.UnboundLocal; _ = µbeginning
				var µlength *πg.Object = πg.UnboundLocal; _ = µlength
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 bool
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
					// line 1209: 'Convert range to the "ed" format'
					πF.SetLineno(1209)
					// line 1211: beginning = start + 1     # lines start numbering with one
					πF.SetLineno(1211)
					if πE = πg.CheckLocal(πF, µstart, "start"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, µstart, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µbeginning = πTemp001
					// line 1212: length = stop - start
					πF.SetLineno(1212)
					if πE = πg.CheckLocal(πF, µstop, "stop"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µstart, "start"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Sub(πF, µstop, µstart); πE != nil {
						continue
					}
					µlength = πTemp001
					if πE = πg.CheckLocal(πF, µlength, "length"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, µlength, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label1
					}
					goto Label2
					// line 1213: if length == 1:
					πF.SetLineno(1213)
				Label1:
					// line 1215: return '%s' % (beginning)
					πF.SetLineno(1215)
					if πE = πg.CheckLocal(πF, µbeginning, "beginning"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("%s").ToObject(), µbeginning); πE != nil {
						continue
					}
					πR = πTemp001
					continue
					goto Label2
				Label2:
					if πE = πg.CheckLocal(πF, µlength, "length"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, µlength); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(!πTemp002).ToObject()
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label3
					}
					goto Label4
					// line 1216: if not length:
					πF.SetLineno(1216)
				Label3:
					// line 1217: beginning -= 1        # empty ranges begin at line just before the range
					πF.SetLineno(1217)
					if πE = πg.CheckLocal(πF, µbeginning, "beginning"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.ISub(πF, µbeginning, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µbeginning = πTemp001
					goto Label4
				Label4:
					// line 1218: return '%s,%s' % (beginning, length)
					πF.SetLineno(1218)
					if πE = πg.CheckLocal(πF, µbeginning, "beginning"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µlength, "length"); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple2(µbeginning, µlength).ToObject()
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("%s,%s").ToObject(), πTemp003); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ß_format_range_unified.ToObject(), πTemp010); πE != nil {
				continue
			}
			// line 1209: 'Convert range to the "ed" format'
			πF.SetLineno(1209)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πg.NewStr("Convert range to the \"ed\" format").ToObject()); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ß_format_range_unified); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp012, ß__doc__, πTemp011); πE != nil {
				continue
			}
			// line 1220: def unified_diff(a, b, fromfile='', tofile='', fromfiledate='',
			πF.SetLineno(1220)
			πTemp004 = make([]πg.Param, 8)
			πTemp004[0] = πg.Param{Name: "a", Def: nil}
			πTemp004[1] = πg.Param{Name: "b", Def: nil}
			πTemp004[2] = πg.Param{Name: "fromfile", Def: ß.ToObject()}
			πTemp004[3] = πg.Param{Name: "tofile", Def: ß.ToObject()}
			πTemp004[4] = πg.Param{Name: "fromfiledate", Def: ß.ToObject()}
			πTemp004[5] = πg.Param{Name: "tofiledate", Def: ß.ToObject()}
			πTemp004[6] = πg.Param{Name: "n", Def: πg.NewInt(3).ToObject()}
			πTemp004[7] = πg.Param{Name: "lineterm", Def: πg.NewStr("\n").ToObject()}
			πTemp011 = πg.NewFunction(πg.NewCode("unified_diff", "/usr/lib/python2.7/difflib.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µa *πg.Object = πArgs[0]; _ = µa
				var µb *πg.Object = πArgs[1]; _ = µb
				var µfromfile *πg.Object = πArgs[2]; _ = µfromfile
				var µtofile *πg.Object = πArgs[3]; _ = µtofile
				var µfromfiledate *πg.Object = πArgs[4]; _ = µfromfiledate
				var µtofiledate *πg.Object = πArgs[5]; _ = µtofiledate
				var µn *πg.Object = πArgs[6]; _ = µn
				var µlineterm *πg.Object = πArgs[7]; _ = µlineterm
				var µstarted *πg.Object = πg.UnboundLocal; _ = µstarted
				var µgroup *πg.Object = πg.UnboundLocal; _ = µgroup
				var µfromdate *πg.Object = πg.UnboundLocal; _ = µfromdate
				var µtodate *πg.Object = πg.UnboundLocal; _ = µtodate
				var µfirst *πg.Object = πg.UnboundLocal; _ = µfirst
				var µlast *πg.Object = πg.UnboundLocal; _ = µlast
				var µfile1_range *πg.Object = πg.UnboundLocal; _ = µfile1_range
				var µfile2_range *πg.Object = πg.UnboundLocal; _ = µfile2_range
				var µtag *πg.Object = πg.UnboundLocal; _ = µtag
				var µi1 *πg.Object = πg.UnboundLocal; _ = µi1
				var µi2 *πg.Object = πg.UnboundLocal; _ = µi2
				var µj1 *πg.Object = πg.UnboundLocal; _ = µj1
				var µj2 *πg.Object = πg.UnboundLocal; _ = µj2
				var µline *πg.Object = πg.UnboundLocal; _ = µline
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []*πg.Object
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
				var πTemp009 *πg.Object
				_ = πTemp009
				var πTemp010 *πg.Object
				_ = πTemp010
				var πTemp011 *πg.Object
				_ = πTemp011
				var πTemp012 bool
				_ = πTemp012
				var πTemp013 *πg.Object
				_ = πTemp013
				var πTemp014 *πg.Object
				_ = πTemp014
				var πTemp015 *πg.Object
				_ = πTemp015
				var πTemp016 *πg.Object
				_ = πTemp016
				var πTemp017 bool
				_ = πTemp017
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
					for ; πF.State() >= 0; πF.PopCheckpoint() {
						switch πF.State() {
						case 0:
						case 1: goto Label1
						case 2: goto Label2
						case 33: goto Label33
						case 10: goto Label10
						case 11: goto Label11
						case 12: goto Label12
						case 13: goto Label13
						case 14: goto Label14
						case 18: goto Label18
						case 19: goto Label19
						case 21: goto Label21
						case 24: goto Label24
						case 25: goto Label25
						case 27: goto Label27
						case 30: goto Label30
						case 31: goto Label31
						default: panic("unexpected function state")
						}
						// line 1222: r"""
						πF.SetLineno(1222)
						// line 1261: started = False
						πF.SetLineno(1261)
						if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
							continue
						}
						µstarted = πTemp001
						πTemp002 = πF.MakeArgs(1)
						if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
							continue
						}
						πTemp002[0] = µn
						πTemp003 = πF.MakeArgs(3)
						if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
							continue
						}
						πTemp003[0] = πTemp004
						if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
							continue
						}
						πTemp003[1] = µa
						if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
							continue
						}
						πTemp003[2] = µb
						if πTemp004, πE = πg.ResolveGlobal(πF, ßSequenceMatcher); πE != nil {
							continue
						}
						if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp003)
						if πTemp004, πE = πg.GetAttr(πF, πTemp005, ßget_grouped_opcodes, nil); πE != nil {
							continue
						}
						if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
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
						if πTemp004, πE = πg.Next(πF, πTemp001); πE != nil {
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
							µgroup = πTemp004
						}
						if πE != nil || !πTemp007 {
							continue
						}
						πF.PushCheckpoint(1)            
						if πE = πg.CheckLocal(πF, µstarted, "started"); πE != nil {
							continue
						}
						if πTemp007, πE = πg.IsTrue(πF, µstarted); πE != nil {
							continue
						}
						πTemp004 = πg.GetBool(!πTemp007).ToObject()
						if πTemp007, πE = πg.IsTrue(πF, πTemp004); πE != nil {
							continue
						}
						if πTemp007 {
							goto Label4
						}
						goto Label5
						// line 1263: if not started:
						πF.SetLineno(1263)
					Label4:
						// line 1264: started = True
						πF.SetLineno(1264)
						if πTemp004, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
							continue
						}
						µstarted = πTemp004
						// line 1266: fromdate = '\t%s' % (fromfiledate) if fromfiledate else ''
						πF.SetLineno(1266)
						if πE = πg.CheckLocal(πF, µfromfiledate, "fromfiledate"); πE != nil {
							continue
						}
						if πTemp007, πE = πg.IsTrue(πF, µfromfiledate); πE != nil {
							continue
						}
						if !πTemp007 {
							goto Label6
						}
						if πE = πg.CheckLocal(πF, µfromfiledate, "fromfiledate"); πE != nil {
							continue
						}
						if πTemp005, πE = πg.Mod(πF, πg.NewStr("\t%s").ToObject(), µfromfiledate); πE != nil {
							continue
						}
						πTemp004 = πTemp005
						goto Label7
					Label6:
						πTemp004 = ß.ToObject()
					Label7:
						µfromdate = πTemp004
						// line 1268: todate = '\t%s' % (tofiledate) if tofiledate else ''
						πF.SetLineno(1268)
						if πE = πg.CheckLocal(πF, µtofiledate, "tofiledate"); πE != nil {
							continue
						}
						if πTemp007, πE = πg.IsTrue(πF, µtofiledate); πE != nil {
							continue
						}
						if !πTemp007 {
							goto Label8
						}
						if πE = πg.CheckLocal(πF, µtofiledate, "tofiledate"); πE != nil {
							continue
						}
						if πTemp005, πE = πg.Mod(πF, πg.NewStr("\t%s").ToObject(), µtofiledate); πE != nil {
							continue
						}
						πTemp004 = πTemp005
						goto Label9
					Label8:
						πTemp004 = ß.ToObject()
					Label9:
						µtodate = πTemp004
						// line 1270: yield '--- %s%s%s' % (fromfile, fromdate, lineterm)
						πF.SetLineno(1270)
						if πE = πg.CheckLocal(πF, µfromfile, "fromfile"); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µfromdate, "fromdate"); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µlineterm, "lineterm"); πE != nil {
							continue
						}
						πTemp005 = πg.NewTuple3(µfromfile, µfromdate, µlineterm).ToObject()
						if πTemp004, πE = πg.Mod(πF, πg.NewStr("--- %s%s%s").ToObject(), πTemp005); πE != nil {
							continue
						}
						πF.PushCheckpoint(10)
						return πTemp004, nil
					Label10:
						πTemp005 = πSent
						// line 1272: yield '+++ %s%s%s' % (tofile, todate, lineterm)
						πF.SetLineno(1272)
						if πE = πg.CheckLocal(πF, µtofile, "tofile"); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µtodate, "todate"); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µlineterm, "lineterm"); πE != nil {
							continue
						}
						πTemp008 = πg.NewTuple3(µtofile, µtodate, µlineterm).ToObject()
						if πTemp005, πE = πg.Mod(πF, πg.NewStr("+++ %s%s%s").ToObject(), πTemp008); πE != nil {
							continue
						}
						πF.PushCheckpoint(11)
						return πTemp005, nil
					Label11:
						πTemp008 = πSent
						goto Label5
					Label5:
						// line 1274: first, last = group[0], group[-1]
						πF.SetLineno(1274)
						πTemp009 = πg.NewInt(0).ToObject()
						if πE = πg.CheckLocal(πF, µgroup, "group"); πE != nil {
							continue
						}
						if πTemp010, πE = πg.GetItem(πF, µgroup, πTemp009); πE != nil {
							continue
						}
						if πTemp011, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
							continue
						}
						πTemp009 = πTemp011
						if πE = πg.CheckLocal(πF, µgroup, "group"); πE != nil {
							continue
						}
						if πTemp011, πE = πg.GetItem(πF, µgroup, πTemp009); πE != nil {
							continue
						}
						πTemp008 = πg.NewTuple2(πTemp010, πTemp011).ToObject()
						if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp009}, πg.TieTarget{Target: &πTemp010}}}, πTemp008); πE != nil {
							continue
						}
						µfirst = πTemp009
						µlast = πTemp010
						// line 1275: file1_range = _format_range_unified(first[1], last[2])
						πF.SetLineno(1275)
						πTemp002 = πF.MakeArgs(2)
						πTemp008 = πg.NewInt(1).ToObject()
						if πE = πg.CheckLocal(πF, µfirst, "first"); πE != nil {
							continue
						}
						if πTemp009, πE = πg.GetItem(πF, µfirst, πTemp008); πE != nil {
							continue
						}
						πTemp002[0] = πTemp009
						πTemp008 = πg.NewInt(2).ToObject()
						if πE = πg.CheckLocal(πF, µlast, "last"); πE != nil {
							continue
						}
						if πTemp009, πE = πg.GetItem(πF, µlast, πTemp008); πE != nil {
							continue
						}
						πTemp002[1] = πTemp009
						if πTemp008, πE = πg.ResolveGlobal(πF, ß_format_range_unified); πE != nil {
							continue
						}
						if πTemp009, πE = πTemp008.Call(πF, πTemp002, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp002)
						µfile1_range = πTemp009
						// line 1276: file2_range = _format_range_unified(first[3], last[4])
						πF.SetLineno(1276)
						πTemp002 = πF.MakeArgs(2)
						πTemp008 = πg.NewInt(3).ToObject()
						if πE = πg.CheckLocal(πF, µfirst, "first"); πE != nil {
							continue
						}
						if πTemp009, πE = πg.GetItem(πF, µfirst, πTemp008); πE != nil {
							continue
						}
						πTemp002[0] = πTemp009
						πTemp008 = πg.NewInt(4).ToObject()
						if πE = πg.CheckLocal(πF, µlast, "last"); πE != nil {
							continue
						}
						if πTemp009, πE = πg.GetItem(πF, µlast, πTemp008); πE != nil {
							continue
						}
						πTemp002[1] = πTemp009
						if πTemp008, πE = πg.ResolveGlobal(πF, ß_format_range_unified); πE != nil {
							continue
						}
						if πTemp009, πE = πTemp008.Call(πF, πTemp002, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp002)
						µfile2_range = πTemp009
						// line 1278: yield '@@ -%s +%s @@%s' % (file1_range, file2_range, lineterm)
						πF.SetLineno(1278)
						if πE = πg.CheckLocal(πF, µfile1_range, "file1_range"); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µfile2_range, "file2_range"); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µlineterm, "lineterm"); πE != nil {
							continue
						}
						πTemp009 = πg.NewTuple3(µfile1_range, µfile2_range, µlineterm).ToObject()
						if πTemp008, πE = πg.Mod(πF, πg.NewStr("@@ -%s +%s @@%s").ToObject(), πTemp009); πE != nil {
							continue
						}
						πF.PushCheckpoint(12)
						return πTemp008, nil
					Label12:
						πTemp009 = πSent
						if πE = πg.CheckLocal(πF, µgroup, "group"); πE != nil {
							continue
						}
						if πTemp009, πE = πg.Iter(πF, µgroup); πE != nil {
							continue
						}
						πF.PushCheckpoint(14)
						πTemp007 = false
					Label13:
						if πE != nil || πR != nil {
							continue
						}
						if πTemp007 {
							πF.PopCheckpoint()
							goto Label15
						}
						if πTemp010, πE = πg.Next(πF, πTemp009); πE != nil {
							isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
							if exc != nil {
								πE = exc
							} else if isStop {
								πE = nil
								πF.RestoreExc(nil, nil)
							}
							πTemp012 = !isStop
						} else {
							πTemp012 = true
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp011}, πg.TieTarget{Target: &πTemp013}, πg.TieTarget{Target: &πTemp014}, πg.TieTarget{Target: &πTemp015}, πg.TieTarget{Target: &πTemp016}}}, πTemp010); πE != nil {
								continue
							}
							µtag = πTemp011
							µi1 = πTemp013
							µi2 = πTemp014
							µj1 = πTemp015
							µj2 = πTemp016
						}
						if πE != nil || !πTemp012 {
							continue
						}
						πF.PushCheckpoint(13)            
						if πE = πg.CheckLocal(πF, µtag, "tag"); πE != nil {
							continue
						}
						if πTemp010, πE = πg.Eq(πF, µtag, ßequal.ToObject()); πE != nil {
							continue
						}
						if πTemp012, πE = πg.IsTrue(πF, πTemp010); πE != nil {
							continue
						}
						if πTemp012 {
							goto Label16
						}
						goto Label17
						// line 1281: if tag == 'equal':
						πF.SetLineno(1281)
					Label16:
						if πE = πg.CheckLocal(πF, µi1, "i1"); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µi2, "i2"); πE != nil {
							continue
						}
						if πTemp011, πE = πg.SliceType.Call(πF, πg.Args{µi1, µi2, πg.None}, nil); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
							continue
						}
						if πTemp013, πE = πg.GetItem(πF, µa, πTemp011); πE != nil {
							continue
						}
						if πTemp010, πE = πg.Iter(πF, πTemp013); πE != nil {
							continue
						}
						πF.PushCheckpoint(19)
						πTemp012 = false
					Label18:
						if πE != nil || πR != nil {
							continue
						}
						if πTemp012 {
							πF.PopCheckpoint()
							goto Label20
						}
						if πTemp011, πE = πg.Next(πF, πTemp010); πE != nil {
							isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
							if exc != nil {
								πE = exc
							} else if isStop {
								πE = nil
								πF.RestoreExc(nil, nil)
							}
							πTemp017 = !isStop
						} else {
							πTemp017 = true
							µline = πTemp011
						}
						if πE != nil || !πTemp017 {
							continue
						}
						πF.PushCheckpoint(18)            
						// line 1283: yield ' ' + line
						πF.SetLineno(1283)
						if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
							continue
						}
						if πTemp011, πE = πg.Add(πF, πg.NewStr(" ").ToObject(), µline); πE != nil {
							continue
						}
						πF.PushCheckpoint(21)
						return πTemp011, nil
					Label21:
						πTemp013 = πSent
						continue
					Label19:
						if πE != nil || πR != nil {
							continue
						}
					Label20:
						// line 1284: continue
						πF.SetLineno(1284)
						continue
						goto Label17
					Label17:
						if πE = πg.CheckLocal(πF, µtag, "tag"); πE != nil {
							continue
						}
						πTemp013 = πg.NewTuple2(ßreplace.ToObject(), ßdelete.ToObject()).ToObject()
						if πTemp012, πE = πg.Contains(πF, πTemp013, µtag); πE != nil {
							continue
						}
						πTemp010 = πg.GetBool(πTemp012).ToObject()
						if πTemp012, πE = πg.IsTrue(πF, πTemp010); πE != nil {
							continue
						}
						if πTemp012 {
							goto Label22
						}
						goto Label23
						// line 1285: if tag in ('replace', 'delete'):
						πF.SetLineno(1285)
					Label22:
						if πE = πg.CheckLocal(πF, µi1, "i1"); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µi2, "i2"); πE != nil {
							continue
						}
						if πTemp013, πE = πg.SliceType.Call(πF, πg.Args{µi1, µi2, πg.None}, nil); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
							continue
						}
						if πTemp014, πE = πg.GetItem(πF, µa, πTemp013); πE != nil {
							continue
						}
						if πTemp010, πE = πg.Iter(πF, πTemp014); πE != nil {
							continue
						}
						πF.PushCheckpoint(25)
						πTemp012 = false
					Label24:
						if πE != nil || πR != nil {
							continue
						}
						if πTemp012 {
							πF.PopCheckpoint()
							goto Label26
						}
						if πTemp013, πE = πg.Next(πF, πTemp010); πE != nil {
							isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
							if exc != nil {
								πE = exc
							} else if isStop {
								πE = nil
								πF.RestoreExc(nil, nil)
							}
							πTemp017 = !isStop
						} else {
							πTemp017 = true
							µline = πTemp013
						}
						if πE != nil || !πTemp017 {
							continue
						}
						πF.PushCheckpoint(24)            
						// line 1287: yield '-' + line
						πF.SetLineno(1287)
						if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
							continue
						}
						if πTemp013, πE = πg.Add(πF, πg.NewStr("-").ToObject(), µline); πE != nil {
							continue
						}
						πF.PushCheckpoint(27)
						return πTemp013, nil
					Label27:
						πTemp014 = πSent
						continue
					Label25:
						if πE != nil || πR != nil {
							continue
						}
					Label26:
						goto Label23
					Label23:
						if πE = πg.CheckLocal(πF, µtag, "tag"); πE != nil {
							continue
						}
						πTemp014 = πg.NewTuple2(ßreplace.ToObject(), ßinsert.ToObject()).ToObject()
						if πTemp012, πE = πg.Contains(πF, πTemp014, µtag); πE != nil {
							continue
						}
						πTemp010 = πg.GetBool(πTemp012).ToObject()
						if πTemp012, πE = πg.IsTrue(πF, πTemp010); πE != nil {
							continue
						}
						if πTemp012 {
							goto Label28
						}
						goto Label29
						// line 1288: if tag in ('replace', 'insert'):
						πF.SetLineno(1288)
					Label28:
						if πE = πg.CheckLocal(πF, µj1, "j1"); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µj2, "j2"); πE != nil {
							continue
						}
						if πTemp014, πE = πg.SliceType.Call(πF, πg.Args{µj1, µj2, πg.None}, nil); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
							continue
						}
						if πTemp015, πE = πg.GetItem(πF, µb, πTemp014); πE != nil {
							continue
						}
						if πTemp010, πE = πg.Iter(πF, πTemp015); πE != nil {
							continue
						}
						πF.PushCheckpoint(31)
						πTemp012 = false
					Label30:
						if πE != nil || πR != nil {
							continue
						}
						if πTemp012 {
							πF.PopCheckpoint()
							goto Label32
						}
						if πTemp014, πE = πg.Next(πF, πTemp010); πE != nil {
							isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
							if exc != nil {
								πE = exc
							} else if isStop {
								πE = nil
								πF.RestoreExc(nil, nil)
							}
							πTemp017 = !isStop
						} else {
							πTemp017 = true
							µline = πTemp014
						}
						if πE != nil || !πTemp017 {
							continue
						}
						πF.PushCheckpoint(30)            
						// line 1290: yield '+' + line
						πF.SetLineno(1290)
						if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
							continue
						}
						if πTemp014, πE = πg.Add(πF, πg.NewStr("+").ToObject(), µline); πE != nil {
							continue
						}
						πF.PushCheckpoint(33)
						return πTemp014, nil
					Label33:
						πTemp015 = πSent
						continue
					Label31:
						if πE != nil || πR != nil {
							continue
						}
					Label32:
						goto Label29
					Label29:
						continue
					Label14:
						if πE != nil || πR != nil {
							continue
						}
					Label15:
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
			if πE = πF.Globals().SetItem(πF, ßunified_diff.ToObject(), πTemp011); πE != nil {
				continue
			}
			// line 1222: r"""
			πF.SetLineno(1222)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp012}, πg.NewStr("\n    Compare two sequences of lines; generate the delta as a unified diff.\n\n    Unified diffs are a compact way of showing line changes and a few\n    lines of context.  The number of context lines is set by 'n' which\n    defaults to three.\n\n    By default, the diff control lines (those with ---, +++, or @@) are\n    created with a trailing newline.  This is helpful so that inputs\n    created from file.readlines() result in diffs that are suitable for\n    file.writelines() since both the inputs and outputs have trailing\n    newlines.\n\n    For inputs that do not have trailing newlines, set the lineterm\n    argument to \"\" so that the output will be uniformly newline free.\n\n    The unidiff format normally has a header for filenames and modification\n    times.  Any or all of these may be specified using strings for\n    'fromfile', 'tofile', 'fromfiledate', and 'tofiledate'.\n    The modification times are normally expressed in the ISO 8601 format.\n\n    Example:\n\n    >>> for line in unified_diff('one two three four'.split(),\n    ...             'zero one tree four'.split(), 'Original', 'Current',\n    ...             '2005-01-26 23:30:50', '2010-04-02 10:20:52',\n    ...             lineterm=''):\n    ...     print line                  # doctest: +NORMALIZE_WHITESPACE\n    --- Original        2005-01-26 23:30:50\n    +++ Current         2010-04-02 10:20:52\n    @@ -1,4 +1,4 @@\n    +zero\n     one\n    -two\n    -three\n    +tree\n     four\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp013, πE = πg.ResolveGlobal(πF, ßunified_diff); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp013, ß__doc__, πTemp012); πE != nil {
				continue
			}
			// line 1297: def _format_range_context(start, stop):
			πF.SetLineno(1297)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "start", Def: nil}
			πTemp004[1] = πg.Param{Name: "stop", Def: nil}
			πTemp012 = πg.NewFunction(πg.NewCode("_format_range_context", "/usr/lib/python2.7/difflib.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µstart *πg.Object = πArgs[0]; _ = µstart
				var µstop *πg.Object = πArgs[1]; _ = µstop
				var µbeginning *πg.Object = πg.UnboundLocal; _ = µbeginning
				var µlength *πg.Object = πg.UnboundLocal; _ = µlength
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
					// line 1298: 'Convert range to the "ed" format'
					πF.SetLineno(1298)
					// line 1300: beginning = start + 1     # lines start numbering with one
					πF.SetLineno(1300)
					if πE = πg.CheckLocal(πF, µstart, "start"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, µstart, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µbeginning = πTemp001
					// line 1301: length = stop - start
					πF.SetLineno(1301)
					if πE = πg.CheckLocal(πF, µstop, "stop"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µstart, "start"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Sub(πF, µstop, µstart); πE != nil {
						continue
					}
					µlength = πTemp001
					if πE = πg.CheckLocal(πF, µlength, "length"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, µlength); πE != nil {
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
					// line 1302: if not length:
					πF.SetLineno(1302)
				Label1:
					// line 1303: beginning -= 1        # empty ranges begin at line just before the range
					πF.SetLineno(1303)
					if πE = πg.CheckLocal(πF, µbeginning, "beginning"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.ISub(πF, µbeginning, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µbeginning = πTemp001
					goto Label2
				Label2:
					if πE = πg.CheckLocal(πF, µlength, "length"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.LE(πF, µlength, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label3
					}
					goto Label4
					// line 1304: if length <= 1:
					πF.SetLineno(1304)
				Label3:
					// line 1306: return '%s' % (beginning)
					πF.SetLineno(1306)
					if πE = πg.CheckLocal(πF, µbeginning, "beginning"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("%s").ToObject(), µbeginning); πE != nil {
						continue
					}
					πR = πTemp001
					continue
					goto Label4
				Label4:
					// line 1308: return '%s,%s' % (beginning, beginning + length - 1)
					πF.SetLineno(1308)
					if πE = πg.CheckLocal(πF, µbeginning, "beginning"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µbeginning, "beginning"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µlength, "length"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Add(πF, µbeginning, µlength); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Sub(πF, πTemp005, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple2(µbeginning, πTemp004).ToObject()
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("%s,%s").ToObject(), πTemp003); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ß_format_range_context.ToObject(), πTemp012); πE != nil {
				continue
			}
			// line 1298: 'Convert range to the "ed" format'
			πF.SetLineno(1298)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp013}, πg.NewStr("Convert range to the \"ed\" format").ToObject()); πE != nil {
				continue
			}
			if πTemp014, πE = πg.ResolveGlobal(πF, ß_format_range_context); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp014, ß__doc__, πTemp013); πE != nil {
				continue
			}
			// line 1311: def context_diff(a, b, fromfile='', tofile='',
			πF.SetLineno(1311)
			πTemp004 = make([]πg.Param, 8)
			πTemp004[0] = πg.Param{Name: "a", Def: nil}
			πTemp004[1] = πg.Param{Name: "b", Def: nil}
			πTemp004[2] = πg.Param{Name: "fromfile", Def: ß.ToObject()}
			πTemp004[3] = πg.Param{Name: "tofile", Def: ß.ToObject()}
			πTemp004[4] = πg.Param{Name: "fromfiledate", Def: ß.ToObject()}
			πTemp004[5] = πg.Param{Name: "tofiledate", Def: ß.ToObject()}
			πTemp004[6] = πg.Param{Name: "n", Def: πg.NewInt(3).ToObject()}
			πTemp004[7] = πg.Param{Name: "lineterm", Def: πg.NewStr("\n").ToObject()}
			πTemp013 = πg.NewFunction(πg.NewCode("context_diff", "/usr/lib/python2.7/difflib.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µa *πg.Object = πArgs[0]; _ = µa
				var µb *πg.Object = πArgs[1]; _ = µb
				var µfromfile *πg.Object = πArgs[2]; _ = µfromfile
				var µtofile *πg.Object = πArgs[3]; _ = µtofile
				var µfromfiledate *πg.Object = πArgs[4]; _ = µfromfiledate
				var µtofiledate *πg.Object = πArgs[5]; _ = µtofiledate
				var µn *πg.Object = πArgs[6]; _ = µn
				var µlineterm *πg.Object = πArgs[7]; _ = µlineterm
				var µprefix *πg.Object = πg.UnboundLocal; _ = µprefix
				var µstarted *πg.Object = πg.UnboundLocal; _ = µstarted
				var µgroup *πg.Object = πg.UnboundLocal; _ = µgroup
				var µfromdate *πg.Object = πg.UnboundLocal; _ = µfromdate
				var µtodate *πg.Object = πg.UnboundLocal; _ = µtodate
				var µfirst *πg.Object = πg.UnboundLocal; _ = µfirst
				var µlast *πg.Object = πg.UnboundLocal; _ = µlast
				var µfile1_range *πg.Object = πg.UnboundLocal; _ = µfile1_range
				var µtag *πg.Object = πg.UnboundLocal; _ = µtag
				var µi1 *πg.Object = πg.UnboundLocal; _ = µi1
				var µi2 *πg.Object = πg.UnboundLocal; _ = µi2
				var µ_ *πg.Object = πg.UnboundLocal; _ = µ_
				var µline *πg.Object = πg.UnboundLocal; _ = µline
				var µfile2_range *πg.Object = πg.UnboundLocal; _ = µfile2_range
				var µj1 *πg.Object = πg.UnboundLocal; _ = µj1
				var µj2 *πg.Object = πg.UnboundLocal; _ = µj2
				var πTemp001 πg.KWArgs
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 []*πg.Object
				_ = πTemp004
				var πTemp005 []*πg.Object
				_ = πTemp005
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 bool
				_ = πTemp007
				var πTemp008 bool
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
				var πTemp014 *πg.Object
				_ = πTemp014
				var πTemp015 bool
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
				var πTemp021 bool
				_ = πTemp021
				var πTemp022 *πg.Object
				_ = πTemp022
				var πTemp023 *πg.Object
				_ = πTemp023
				var πTemp024 *πg.Object
				_ = πTemp024
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
					for ; πF.State() >= 0; πF.PopCheckpoint() {
						switch πF.State() {
						case 0:
						case 1: goto Label1
						case 2: goto Label2
						case 34: goto Label34
						case 36: goto Label36
						case 33: goto Label33
						case 10: goto Label10
						case 11: goto Label11
						case 12: goto Label12
						case 13: goto Label13
						case 16: goto Label16
						case 17: goto Label17
						case 21: goto Label21
						case 22: goto Label22
						case 24: goto Label24
						case 25: goto Label25
						case 28: goto Label28
						case 29: goto Label29
						default: panic("unexpected function state")
						}
						// line 1313: r"""
						πF.SetLineno(1313)
						// line 1354: prefix = dict(insert='+ ', delete='- ', replace='! ', equal='  ')
						πF.SetLineno(1354)
						πTemp001 = πg.KWArgs{
							{"insert", πg.NewStr("+ ").ToObject()},
							{"delete", πg.NewStr("- ").ToObject()},
							{"replace", πg.NewStr("! ").ToObject()},
							{"equal", πg.NewStr("  ").ToObject()},
						}
						if πTemp002, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
							continue
						}
						if πTemp003, πE = πTemp002.Call(πF, nil, πTemp001); πE != nil {
							continue
						}
						µprefix = πTemp003
						// line 1355: started = False
						πF.SetLineno(1355)
						if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
							continue
						}
						µstarted = πTemp002
						πTemp004 = πF.MakeArgs(1)
						if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
							continue
						}
						πTemp004[0] = µn
						πTemp005 = πF.MakeArgs(3)
						if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
							continue
						}
						πTemp005[0] = πTemp003
						if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
							continue
						}
						πTemp005[1] = µa
						if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
							continue
						}
						πTemp005[2] = µb
						if πTemp003, πE = πg.ResolveGlobal(πF, ßSequenceMatcher); πE != nil {
							continue
						}
						if πTemp006, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp005)
						if πTemp003, πE = πg.GetAttr(πF, πTemp006, ßget_grouped_opcodes, nil); πE != nil {
							continue
						}
						if πTemp006, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp004)
						if πTemp002, πE = πg.Iter(πF, πTemp006); πE != nil {
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
						if πTemp003, πE = πg.Next(πF, πTemp002); πE != nil {
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
							µgroup = πTemp003
						}
						if πE != nil || !πTemp008 {
							continue
						}
						πF.PushCheckpoint(1)            
						if πE = πg.CheckLocal(πF, µstarted, "started"); πE != nil {
							continue
						}
						if πTemp008, πE = πg.IsTrue(πF, µstarted); πE != nil {
							continue
						}
						πTemp003 = πg.GetBool(!πTemp008).ToObject()
						if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
							continue
						}
						if πTemp008 {
							goto Label4
						}
						goto Label5
						// line 1357: if not started:
						πF.SetLineno(1357)
					Label4:
						// line 1358: started = True
						πF.SetLineno(1358)
						if πTemp003, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
							continue
						}
						µstarted = πTemp003
						// line 1360: fromdate = '\t%s' % (fromfiledate) if fromfiledate else ''
						πF.SetLineno(1360)
						if πE = πg.CheckLocal(πF, µfromfiledate, "fromfiledate"); πE != nil {
							continue
						}
						if πTemp008, πE = πg.IsTrue(πF, µfromfiledate); πE != nil {
							continue
						}
						if !πTemp008 {
							goto Label6
						}
						if πE = πg.CheckLocal(πF, µfromfiledate, "fromfiledate"); πE != nil {
							continue
						}
						if πTemp006, πE = πg.Mod(πF, πg.NewStr("\t%s").ToObject(), µfromfiledate); πE != nil {
							continue
						}
						πTemp003 = πTemp006
						goto Label7
					Label6:
						πTemp003 = ß.ToObject()
					Label7:
						µfromdate = πTemp003
						// line 1362: todate = '\t%s' % (tofiledate) if tofiledate else ''
						πF.SetLineno(1362)
						if πE = πg.CheckLocal(πF, µtofiledate, "tofiledate"); πE != nil {
							continue
						}
						if πTemp008, πE = πg.IsTrue(πF, µtofiledate); πE != nil {
							continue
						}
						if !πTemp008 {
							goto Label8
						}
						if πE = πg.CheckLocal(πF, µtofiledate, "tofiledate"); πE != nil {
							continue
						}
						if πTemp006, πE = πg.Mod(πF, πg.NewStr("\t%s").ToObject(), µtofiledate); πE != nil {
							continue
						}
						πTemp003 = πTemp006
						goto Label9
					Label8:
						πTemp003 = ß.ToObject()
					Label9:
						µtodate = πTemp003
						// line 1364: yield '*** %s%s%s' % (fromfile, fromdate, lineterm)
						πF.SetLineno(1364)
						if πE = πg.CheckLocal(πF, µfromfile, "fromfile"); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µfromdate, "fromdate"); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µlineterm, "lineterm"); πE != nil {
							continue
						}
						πTemp006 = πg.NewTuple3(µfromfile, µfromdate, µlineterm).ToObject()
						if πTemp003, πE = πg.Mod(πF, πg.NewStr("*** %s%s%s").ToObject(), πTemp006); πE != nil {
							continue
						}
						πF.PushCheckpoint(10)
						return πTemp003, nil
					Label10:
						πTemp006 = πSent
						// line 1366: yield '--- %s%s%s' % (tofile, todate, lineterm)
						πF.SetLineno(1366)
						if πE = πg.CheckLocal(πF, µtofile, "tofile"); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µtodate, "todate"); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µlineterm, "lineterm"); πE != nil {
							continue
						}
						πTemp009 = πg.NewTuple3(µtofile, µtodate, µlineterm).ToObject()
						if πTemp006, πE = πg.Mod(πF, πg.NewStr("--- %s%s%s").ToObject(), πTemp009); πE != nil {
							continue
						}
						πF.PushCheckpoint(11)
						return πTemp006, nil
					Label11:
						πTemp009 = πSent
						goto Label5
					Label5:
						// line 1368: first, last = group[0], group[-1]
						πF.SetLineno(1368)
						πTemp010 = πg.NewInt(0).ToObject()
						if πE = πg.CheckLocal(πF, µgroup, "group"); πE != nil {
							continue
						}
						if πTemp011, πE = πg.GetItem(πF, µgroup, πTemp010); πE != nil {
							continue
						}
						if πTemp012, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
							continue
						}
						πTemp010 = πTemp012
						if πE = πg.CheckLocal(πF, µgroup, "group"); πE != nil {
							continue
						}
						if πTemp012, πE = πg.GetItem(πF, µgroup, πTemp010); πE != nil {
							continue
						}
						πTemp009 = πg.NewTuple2(πTemp011, πTemp012).ToObject()
						if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp010}, πg.TieTarget{Target: &πTemp011}}}, πTemp009); πE != nil {
							continue
						}
						µfirst = πTemp010
						µlast = πTemp011
						// line 1369: yield '***************' + lineterm
						πF.SetLineno(1369)
						if πE = πg.CheckLocal(πF, µlineterm, "lineterm"); πE != nil {
							continue
						}
						if πTemp009, πE = πg.Add(πF, πg.NewStr("***************").ToObject(), µlineterm); πE != nil {
							continue
						}
						πF.PushCheckpoint(12)
						return πTemp009, nil
					Label12:
						πTemp010 = πSent
						// line 1371: file1_range = _format_range_context(first[1], last[2])
						πF.SetLineno(1371)
						πTemp004 = πF.MakeArgs(2)
						πTemp010 = πg.NewInt(1).ToObject()
						if πE = πg.CheckLocal(πF, µfirst, "first"); πE != nil {
							continue
						}
						if πTemp011, πE = πg.GetItem(πF, µfirst, πTemp010); πE != nil {
							continue
						}
						πTemp004[0] = πTemp011
						πTemp010 = πg.NewInt(2).ToObject()
						if πE = πg.CheckLocal(πF, µlast, "last"); πE != nil {
							continue
						}
						if πTemp011, πE = πg.GetItem(πF, µlast, πTemp010); πE != nil {
							continue
						}
						πTemp004[1] = πTemp011
						if πTemp010, πE = πg.ResolveGlobal(πF, ß_format_range_context); πE != nil {
							continue
						}
						if πTemp011, πE = πTemp010.Call(πF, πTemp004, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp004)
						µfile1_range = πTemp011
						// line 1373: yield '*** %s ****%s' % (file1_range, lineterm)
						πF.SetLineno(1373)
						if πE = πg.CheckLocal(πF, µfile1_range, "file1_range"); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µlineterm, "lineterm"); πE != nil {
							continue
						}
						πTemp011 = πg.NewTuple2(µfile1_range, µlineterm).ToObject()
						if πTemp010, πE = πg.Mod(πF, πg.NewStr("*** %s ****%s").ToObject(), πTemp011); πE != nil {
							continue
						}
						πF.PushCheckpoint(13)
						return πTemp010, nil
					Label13:
						πTemp011 = πSent
						πTemp004 = πF.MakeArgs(1)
						πTemp013 = make([]πg.Param, 0)
						πTemp011 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/difflib.py", πTemp013, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
							var µtag *πg.Object = πg.UnboundLocal; _ = µtag
							var µ_ *πg.Object = πg.UnboundLocal; _ = µ_
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
									if πE = πg.CheckLocal(πF, µgroup, "group"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.Iter(πF, µgroup); πE != nil {
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
										if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp006}, πg.TieTarget{Target: &πTemp007}, πg.TieTarget{Target: &πTemp008}, πg.TieTarget{Target: &πTemp009}}}, πTemp004); πE != nil {
											continue
										}
										µtag = πTemp005
										µ_ = πTemp006
										µ_ = πTemp007
										µ_ = πTemp008
										µ_ = πTemp009
									}
									if πE != nil || !πTemp003 {
										continue
									}
									πF.PushCheckpoint(1)            
									// line 1375: if any(tag in ('replace', 'delete') for tag, _, _, _, _ in group):
									πF.SetLineno(1375)
									if πE = πg.CheckLocal(πF, µtag, "tag"); πE != nil {
										continue
									}
									πTemp005 = πg.NewTuple2(ßreplace.ToObject(), ßdelete.ToObject()).ToObject()
									if πTemp003, πE = πg.Contains(πF, πTemp005, µtag); πE != nil {
										continue
									}
									πTemp004 = πg.GetBool(πTemp003).ToObject()
									πF.PushCheckpoint(4)
									return πTemp004, nil
								Label4:
									πTemp005 = πSent
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
						if πTemp012, πE = πTemp011.Call(πF, nil, nil); πE != nil {
							continue
						}
						πTemp004[0] = πTemp012
						if πTemp012, πE = πg.ResolveGlobal(πF, ßany); πE != nil {
							continue
						}
						if πTemp014, πE = πTemp012.Call(πF, πTemp004, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp004)
						if πTemp008, πE = πg.IsTrue(πF, πTemp014); πE != nil {
							continue
						}
						if πTemp008 {
							goto Label14
						}
						goto Label15
						// line 1375: if any(tag in ('replace', 'delete') for tag, _, _, _, _ in group):
						πF.SetLineno(1375)
					Label14:
						if πE = πg.CheckLocal(πF, µgroup, "group"); πE != nil {
							continue
						}
						if πTemp012, πE = πg.Iter(πF, µgroup); πE != nil {
							continue
						}
						πF.PushCheckpoint(17)
						πTemp008 = false
					Label16:
						if πE != nil || πR != nil {
							continue
						}
						if πTemp008 {
							πF.PopCheckpoint()
							goto Label18
						}
						if πTemp014, πE = πg.Next(πF, πTemp012); πE != nil {
							isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
							if exc != nil {
								πE = exc
							} else if isStop {
								πE = nil
								πF.RestoreExc(nil, nil)
							}
							πTemp015 = !isStop
						} else {
							πTemp015 = true
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp016}, πg.TieTarget{Target: &πTemp017}, πg.TieTarget{Target: &πTemp018}, πg.TieTarget{Target: &πTemp019}, πg.TieTarget{Target: &πTemp020}}}, πTemp014); πE != nil {
								continue
							}
							µtag = πTemp016
							µi1 = πTemp017
							µi2 = πTemp018
							µ_ = πTemp019
							µ_ = πTemp020
						}
						if πE != nil || !πTemp015 {
							continue
						}
						πF.PushCheckpoint(16)            
						if πE = πg.CheckLocal(πF, µtag, "tag"); πE != nil {
							continue
						}
						if πTemp014, πE = πg.NE(πF, µtag, ßinsert.ToObject()); πE != nil {
							continue
						}
						if πTemp015, πE = πg.IsTrue(πF, πTemp014); πE != nil {
							continue
						}
						if πTemp015 {
							goto Label19
						}
						goto Label20
						// line 1377: if tag != 'insert':
						πF.SetLineno(1377)
					Label19:
						if πE = πg.CheckLocal(πF, µi1, "i1"); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µi2, "i2"); πE != nil {
							continue
						}
						if πTemp016, πE = πg.SliceType.Call(πF, πg.Args{µi1, µi2, πg.None}, nil); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
							continue
						}
						if πTemp017, πE = πg.GetItem(πF, µa, πTemp016); πE != nil {
							continue
						}
						if πTemp014, πE = πg.Iter(πF, πTemp017); πE != nil {
							continue
						}
						πF.PushCheckpoint(22)
						πTemp015 = false
					Label21:
						if πE != nil || πR != nil {
							continue
						}
						if πTemp015 {
							πF.PopCheckpoint()
							goto Label23
						}
						if πTemp016, πE = πg.Next(πF, πTemp014); πE != nil {
							isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
							if exc != nil {
								πE = exc
							} else if isStop {
								πE = nil
								πF.RestoreExc(nil, nil)
							}
							πTemp021 = !isStop
						} else {
							πTemp021 = true
							µline = πTemp016
						}
						if πE != nil || !πTemp021 {
							continue
						}
						πF.PushCheckpoint(21)            
						// line 1379: yield prefix[tag] + line
						πF.SetLineno(1379)
						if πE = πg.CheckLocal(πF, µtag, "tag"); πE != nil {
							continue
						}
						πTemp017 = µtag
						if πE = πg.CheckLocal(πF, µprefix, "prefix"); πE != nil {
							continue
						}
						if πTemp018, πE = πg.GetItem(πF, µprefix, πTemp017); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
							continue
						}
						if πTemp016, πE = πg.Add(πF, πTemp018, µline); πE != nil {
							continue
						}
						πF.PushCheckpoint(24)
						return πTemp016, nil
					Label24:
						πTemp017 = πSent
						continue
					Label22:
						if πE != nil || πR != nil {
							continue
						}
					Label23:
						goto Label20
					Label20:
						continue
					Label17:
						if πE != nil || πR != nil {
							continue
						}
					Label18:
						goto Label15
					Label15:
						// line 1381: file2_range = _format_range_context(first[3], last[4])
						πF.SetLineno(1381)
						πTemp004 = πF.MakeArgs(2)
						πTemp012 = πg.NewInt(3).ToObject()
						if πE = πg.CheckLocal(πF, µfirst, "first"); πE != nil {
							continue
						}
						if πTemp014, πE = πg.GetItem(πF, µfirst, πTemp012); πE != nil {
							continue
						}
						πTemp004[0] = πTemp014
						πTemp012 = πg.NewInt(4).ToObject()
						if πE = πg.CheckLocal(πF, µlast, "last"); πE != nil {
							continue
						}
						if πTemp014, πE = πg.GetItem(πF, µlast, πTemp012); πE != nil {
							continue
						}
						πTemp004[1] = πTemp014
						if πTemp012, πE = πg.ResolveGlobal(πF, ß_format_range_context); πE != nil {
							continue
						}
						if πTemp014, πE = πTemp012.Call(πF, πTemp004, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp004)
						µfile2_range = πTemp014
						// line 1383: yield '--- %s ----%s' % (file2_range, lineterm)
						πF.SetLineno(1383)
						if πE = πg.CheckLocal(πF, µfile2_range, "file2_range"); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µlineterm, "lineterm"); πE != nil {
							continue
						}
						πTemp014 = πg.NewTuple2(µfile2_range, µlineterm).ToObject()
						if πTemp012, πE = πg.Mod(πF, πg.NewStr("--- %s ----%s").ToObject(), πTemp014); πE != nil {
							continue
						}
						πF.PushCheckpoint(25)
						return πTemp012, nil
					Label25:
						πTemp014 = πSent
						πTemp004 = πF.MakeArgs(1)
						πTemp013 = make([]πg.Param, 0)
						πTemp014 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/difflib.py", πTemp013, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
							var µtag *πg.Object = πg.UnboundLocal; _ = µtag
							var µ_ *πg.Object = πg.UnboundLocal; _ = µ_
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
									if πE = πg.CheckLocal(πF, µgroup, "group"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.Iter(πF, µgroup); πE != nil {
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
										if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp006}, πg.TieTarget{Target: &πTemp007}, πg.TieTarget{Target: &πTemp008}, πg.TieTarget{Target: &πTemp009}}}, πTemp004); πE != nil {
											continue
										}
										µtag = πTemp005
										µ_ = πTemp006
										µ_ = πTemp007
										µ_ = πTemp008
										µ_ = πTemp009
									}
									if πE != nil || !πTemp003 {
										continue
									}
									πF.PushCheckpoint(1)            
									// line 1385: if any(tag in ('replace', 'insert') for tag, _, _, _, _ in group):
									πF.SetLineno(1385)
									if πE = πg.CheckLocal(πF, µtag, "tag"); πE != nil {
										continue
									}
									πTemp005 = πg.NewTuple2(ßreplace.ToObject(), ßinsert.ToObject()).ToObject()
									if πTemp003, πE = πg.Contains(πF, πTemp005, µtag); πE != nil {
										continue
									}
									πTemp004 = πg.GetBool(πTemp003).ToObject()
									πF.PushCheckpoint(4)
									return πTemp004, nil
								Label4:
									πTemp005 = πSent
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
						if πTemp017, πE = πTemp014.Call(πF, nil, nil); πE != nil {
							continue
						}
						πTemp004[0] = πTemp017
						if πTemp017, πE = πg.ResolveGlobal(πF, ßany); πE != nil {
							continue
						}
						if πTemp018, πE = πTemp017.Call(πF, πTemp004, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp004)
						if πTemp008, πE = πg.IsTrue(πF, πTemp018); πE != nil {
							continue
						}
						if πTemp008 {
							goto Label26
						}
						goto Label27
						// line 1385: if any(tag in ('replace', 'insert') for tag, _, _, _, _ in group):
						πF.SetLineno(1385)
					Label26:
						if πE = πg.CheckLocal(πF, µgroup, "group"); πE != nil {
							continue
						}
						if πTemp017, πE = πg.Iter(πF, µgroup); πE != nil {
							continue
						}
						πF.PushCheckpoint(29)
						πTemp008 = false
					Label28:
						if πE != nil || πR != nil {
							continue
						}
						if πTemp008 {
							πF.PopCheckpoint()
							goto Label30
						}
						if πTemp018, πE = πg.Next(πF, πTemp017); πE != nil {
							isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
							if exc != nil {
								πE = exc
							} else if isStop {
								πE = nil
								πF.RestoreExc(nil, nil)
							}
							πTemp015 = !isStop
						} else {
							πTemp015 = true
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp019}, πg.TieTarget{Target: &πTemp020}, πg.TieTarget{Target: &πTemp022}, πg.TieTarget{Target: &πTemp023}, πg.TieTarget{Target: &πTemp024}}}, πTemp018); πE != nil {
								continue
							}
							µtag = πTemp019
							µ_ = πTemp020
							µ_ = πTemp022
							µj1 = πTemp023
							µj2 = πTemp024
						}
						if πE != nil || !πTemp015 {
							continue
						}
						πF.PushCheckpoint(28)            
						if πE = πg.CheckLocal(πF, µtag, "tag"); πE != nil {
							continue
						}
						if πTemp018, πE = πg.NE(πF, µtag, ßdelete.ToObject()); πE != nil {
							continue
						}
						if πTemp015, πE = πg.IsTrue(πF, πTemp018); πE != nil {
							continue
						}
						if πTemp015 {
							goto Label31
						}
						goto Label32
						// line 1387: if tag != 'delete':
						πF.SetLineno(1387)
					Label31:
						if πE = πg.CheckLocal(πF, µj1, "j1"); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µj2, "j2"); πE != nil {
							continue
						}
						if πTemp019, πE = πg.SliceType.Call(πF, πg.Args{µj1, µj2, πg.None}, nil); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
							continue
						}
						if πTemp020, πE = πg.GetItem(πF, µb, πTemp019); πE != nil {
							continue
						}
						if πTemp018, πE = πg.Iter(πF, πTemp020); πE != nil {
							continue
						}
						πF.PushCheckpoint(34)
						πTemp015 = false
					Label33:
						if πE != nil || πR != nil {
							continue
						}
						if πTemp015 {
							πF.PopCheckpoint()
							goto Label35
						}
						if πTemp019, πE = πg.Next(πF, πTemp018); πE != nil {
							isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
							if exc != nil {
								πE = exc
							} else if isStop {
								πE = nil
								πF.RestoreExc(nil, nil)
							}
							πTemp021 = !isStop
						} else {
							πTemp021 = true
							µline = πTemp019
						}
						if πE != nil || !πTemp021 {
							continue
						}
						πF.PushCheckpoint(33)            
						// line 1389: yield prefix[tag] + line
						πF.SetLineno(1389)
						if πE = πg.CheckLocal(πF, µtag, "tag"); πE != nil {
							continue
						}
						πTemp020 = µtag
						if πE = πg.CheckLocal(πF, µprefix, "prefix"); πE != nil {
							continue
						}
						if πTemp022, πE = πg.GetItem(πF, µprefix, πTemp020); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
							continue
						}
						if πTemp019, πE = πg.Add(πF, πTemp022, µline); πE != nil {
							continue
						}
						πF.PushCheckpoint(36)
						return πTemp019, nil
					Label36:
						πTemp020 = πSent
						continue
					Label34:
						if πE != nil || πR != nil {
							continue
						}
					Label35:
						goto Label32
					Label32:
						continue
					Label29:
						if πE != nil || πR != nil {
							continue
						}
					Label30:
						goto Label27
					Label27:
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
			if πE = πF.Globals().SetItem(πF, ßcontext_diff.ToObject(), πTemp013); πE != nil {
				continue
			}
			// line 1313: r"""
			πF.SetLineno(1313)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp014}, πg.NewStr("\n    Compare two sequences of lines; generate the delta as a context diff.\n\n    Context diffs are a compact way of showing line changes and a few\n    lines of context.  The number of context lines is set by 'n' which\n    defaults to three.\n\n    By default, the diff control lines (those with *** or ---) are\n    created with a trailing newline.  This is helpful so that inputs\n    created from file.readlines() result in diffs that are suitable for\n    file.writelines() since both the inputs and outputs have trailing\n    newlines.\n\n    For inputs that do not have trailing newlines, set the lineterm\n    argument to \"\" so that the output will be uniformly newline free.\n\n    The context diff format normally has a header for filenames and\n    modification times.  Any or all of these may be specified using\n    strings for 'fromfile', 'tofile', 'fromfiledate', and 'tofiledate'.\n    The modification times are normally expressed in the ISO 8601 format.\n    If not specified, the strings default to blanks.\n\n    Example:\n\n    >>> print ''.join(context_diff('one\\ntwo\\nthree\\nfour\\n'.splitlines(1),\n    ...       'zero\\none\\ntree\\nfour\\n'.splitlines(1), 'Original', 'Current')),\n    *** Original\n    --- Current\n    ***************\n    *** 1,4 ****\n      one\n    ! two\n    ! three\n      four\n    --- 1,4 ----\n    + zero\n      one\n    ! tree\n      four\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp015, πE = πg.ResolveGlobal(πF, ßcontext_diff); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp015, ß__doc__, πTemp014); πE != nil {
				continue
			}
			// line 1391: def ndiff(a, b, linejunk=None, charjunk=IS_CHARACTER_JUNK):
			πF.SetLineno(1391)
			πTemp004 = make([]πg.Param, 4)
			πTemp004[0] = πg.Param{Name: "a", Def: nil}
			πTemp004[1] = πg.Param{Name: "b", Def: nil}
			if πTemp015, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp004[2] = πg.Param{Name: "linejunk", Def: πTemp015}
			if πTemp015, πE = πg.ResolveGlobal(πF, ßIS_CHARACTER_JUNK); πE != nil {
				continue
			}
			πTemp004[3] = πg.Param{Name: "charjunk", Def: πTemp015}
			πTemp014 = πg.NewFunction(πg.NewCode("ndiff", "/usr/lib/python2.7/difflib.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µa *πg.Object = πArgs[0]; _ = µa
				var µb *πg.Object = πArgs[1]; _ = µb
				var µlinejunk *πg.Object = πArgs[2]; _ = µlinejunk
				var µcharjunk *πg.Object = πArgs[3]; _ = µcharjunk
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
					// line 1392: r"""
					πF.SetLineno(1392)
					// line 1425: return Differ(linejunk, charjunk).compare(a, b)
					πF.SetLineno(1425)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					πTemp001[0] = µa
					if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
						continue
					}
					πTemp001[1] = µb
					πTemp002 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µlinejunk, "linejunk"); πE != nil {
						continue
					}
					πTemp002[0] = µlinejunk
					if πE = πg.CheckLocal(πF, µcharjunk, "charjunk"); πE != nil {
						continue
					}
					πTemp002[1] = µcharjunk
					if πTemp003, πE = πg.ResolveGlobal(πF, ßDiffer); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßcompare, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßndiff.ToObject(), πTemp014); πE != nil {
				continue
			}
			// line 1392: r"""
			πF.SetLineno(1392)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp015}, πg.NewStr("\n    Compare `a` and `b` (lists of strings); return a `Differ`-style delta.\n\n    Optional keyword parameters `linejunk` and `charjunk` are for filter\n    functions (or None):\n\n    - linejunk: A function that should accept a single string argument, and\n      return true iff the string is junk.  The default is None, and is\n      recommended; as of Python 2.3, an adaptive notion of \"noise\" lines is\n      used that does a good job on its own.\n\n    - charjunk: A function that should accept a string of length 1. The\n      default is module-level function IS_CHARACTER_JUNK, which filters out\n      whitespace characters (a blank or tab; note: bad idea to include newline\n      in this!).\n\n    Tools/scripts/ndiff.py is a command-line front-end to this function.\n\n    Example:\n\n    >>> diff = ndiff('one\\ntwo\\nthree\\n'.splitlines(1),\n    ...              'ore\\ntree\\nemu\\n'.splitlines(1))\n    >>> print ''.join(diff),\n    - one\n    ?  ^\n    + ore\n    ?  ^\n    - two\n    - three\n    ?  -\n    + tree\n    + emu\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp016, πE = πg.ResolveGlobal(πF, ßndiff); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp016, ß__doc__, πTemp015); πE != nil {
				continue
			}
			// line 1427: def _mdiff(fromlines, tolines, context=None, linejunk=None,
			πF.SetLineno(1427)
			πTemp004 = make([]πg.Param, 5)
			πTemp004[0] = πg.Param{Name: "fromlines", Def: nil}
			πTemp004[1] = πg.Param{Name: "tolines", Def: nil}
			if πTemp016, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp004[2] = πg.Param{Name: "context", Def: πTemp016}
			if πTemp016, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp004[3] = πg.Param{Name: "linejunk", Def: πTemp016}
			if πTemp016, πE = πg.ResolveGlobal(πF, ßIS_CHARACTER_JUNK); πE != nil {
				continue
			}
			πTemp004[4] = πg.Param{Name: "charjunk", Def: πTemp016}
			πTemp015 = πg.NewFunction(πg.NewCode("_mdiff", "/usr/lib/python2.7/difflib.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µfromlines *πg.Object = πArgs[0]; _ = µfromlines
				var µtolines *πg.Object = πArgs[1]; _ = µtolines
				var µcontext *πg.Object = πArgs[2]; _ = µcontext
				var µlinejunk *πg.Object = πArgs[3]; _ = µlinejunk
				var µcharjunk *πg.Object = πArgs[4]; _ = µcharjunk
				var µre *πg.Object = πg.UnboundLocal; _ = µre
				var µchange_re *πg.Object = πg.UnboundLocal; _ = µchange_re
				var µdiff_lines_iterator *πg.Object = πg.UnboundLocal; _ = µdiff_lines_iterator
				var µ_make_line *πg.Object = πg.UnboundLocal; _ = µ_make_line
				var µ_line_iterator *πg.Object = πg.UnboundLocal; _ = µ_line_iterator
				var µ_line_pair_iterator *πg.Object = πg.UnboundLocal; _ = µ_line_pair_iterator
				var µline_pair_iterator *πg.Object = πg.UnboundLocal; _ = µline_pair_iterator
				var µlines_to_write *πg.Object = πg.UnboundLocal; _ = µlines_to_write
				var µindex *πg.Object = πg.UnboundLocal; _ = µindex
				var µcontextLines *πg.Object = πg.UnboundLocal; _ = µcontextLines
				var µfound_diff *πg.Object = πg.UnboundLocal; _ = µfound_diff
				var µfrom_line *πg.Object = πg.UnboundLocal; _ = µfrom_line
				var µto_line *πg.Object = πg.UnboundLocal; _ = µto_line
				var µi *πg.Object = πg.UnboundLocal; _ = µi
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []*πg.Object
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
				var πTemp008 bool
				_ = πTemp008
				var πTemp009 bool
				_ = πTemp009
				var πTemp010 *πg.Object
				_ = πTemp010
				var πTemp011 *πg.Object
				_ = πTemp011
				var πTemp012 bool
				_ = πTemp012
				var πTemp013 *πg.Object
				_ = πTemp013
				var πTemp014 *πg.Object
				_ = πTemp014
				var πTemp015 *πg.Object
				_ = πTemp015
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
					for ; πF.State() >= 0; πF.PopCheckpoint() {
						switch πF.State() {
						case 0:
						case 4: goto Label4
						case 5: goto Label5
						case 7: goto Label7
						case 8: goto Label8
						case 9: goto Label9
						case 11: goto Label11
						case 12: goto Label12
						case 17: goto Label17
						case 18: goto Label18
						case 19: goto Label19
						case 21: goto Label21
						case 22: goto Label22
						case 23: goto Label23
						case 28: goto Label28
						default: panic("unexpected function state")
						}
						// line 1429: r"""Returns generator yielding marked up from/to side by side differences.
						πF.SetLineno(1429)
						// line 1461: import re
						πF.SetLineno(1461)
						if πTemp002, πE = πg.ImportModule(πF, "re"); πE != nil {
							continue
						}
						πTemp001 = πTemp002[0]
						µre = πTemp001
						// line 1464: change_re = re.compile('(\++|\-+|\^+)')
						πF.SetLineno(1464)
						πTemp002 = πF.MakeArgs(1)
						πTemp002[0] = πg.NewStr("(\\++|\\-+|\\^+)").ToObject()
						if πE = πg.CheckLocal(πF, µre, "re"); πE != nil {
							continue
						}
						if πTemp001, πE = πg.GetAttr(πF, µre, ßcompile, nil); πE != nil {
							continue
						}
						if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp002)
						µchange_re = πTemp003
						// line 1467: diff_lines_iterator = ndiff(fromlines,tolines,linejunk,charjunk)
						πF.SetLineno(1467)
						πTemp002 = πF.MakeArgs(4)
						if πE = πg.CheckLocal(πF, µfromlines, "fromlines"); πE != nil {
							continue
						}
						πTemp002[0] = µfromlines
						if πE = πg.CheckLocal(πF, µtolines, "tolines"); πE != nil {
							continue
						}
						πTemp002[1] = µtolines
						if πE = πg.CheckLocal(πF, µlinejunk, "linejunk"); πE != nil {
							continue
						}
						πTemp002[2] = µlinejunk
						if πE = πg.CheckLocal(πF, µcharjunk, "charjunk"); πE != nil {
							continue
						}
						πTemp002[3] = µcharjunk
						if πTemp001, πE = πg.ResolveGlobal(πF, ßndiff); πE != nil {
							continue
						}
						if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp002)
						µdiff_lines_iterator = πTemp003
						// line 1469: def _make_line(lines, format_key, side, num_lines=[0,0]):
						πF.SetLineno(1469)
						πTemp004 = make([]πg.Param, 4)
						πTemp004[0] = πg.Param{Name: "lines", Def: nil}
						πTemp004[1] = πg.Param{Name: "format_key", Def: nil}
						πTemp004[2] = πg.Param{Name: "side", Def: nil}
						πTemp002 = make([]*πg.Object, 2)
						πTemp002[0] = πg.NewInt(0).ToObject()
						πTemp002[1] = πg.NewInt(0).ToObject()
						πTemp003 = πg.NewList(πTemp002...).ToObject()
						πTemp004[3] = πg.Param{Name: "num_lines", Def: πTemp003}
						πTemp001 = πg.NewFunction(πg.NewCode("_make_line", "/usr/lib/python2.7/difflib.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
							var µlines *πg.Object = πArgs[0]; _ = µlines
							var µformat_key *πg.Object = πArgs[1]; _ = µformat_key
							var µside *πg.Object = πArgs[2]; _ = µside
							var µnum_lines *πg.Object = πArgs[3]; _ = µnum_lines
							var µtext *πg.Object = πg.UnboundLocal; _ = µtext
							var µmarkers *πg.Object = πg.UnboundLocal; _ = µmarkers
							var µsub_info *πg.Object = πg.UnboundLocal; _ = µsub_info
							var µrecord_sub_info *πg.Object = πg.UnboundLocal; _ = µrecord_sub_info
							var µkey *πg.Object = πg.UnboundLocal; _ = µkey
							var µbegin *πg.Object = πg.UnboundLocal; _ = µbegin
							var µend *πg.Object = πg.UnboundLocal; _ = µend
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
							var πTemp006 []*πg.Object
							_ = πTemp006
							var πTemp007 *πg.Object
							_ = πTemp007
							var πTemp008 *πg.Object
							_ = πTemp008
							var πTemp009 []πg.Param
							_ = πTemp009
							var πTemp010 bool
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
								case 6: goto Label6
								case 7: goto Label7
								default: panic("unexpected function state")
								}
								// line 1470: """Returns line of text with user's change markup and line formatting.
								πF.SetLineno(1470)
								// line 1492: num_lines[side] += 1
								πF.SetLineno(1492)
								if πE = πg.CheckLocal(πF, µside, "side"); πE != nil {
									continue
								}
								πTemp001 = µside
								if πE = πg.CheckLocal(πF, µnum_lines, "num_lines"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.GetItem(πF, µnum_lines, πTemp001); πE != nil {
									continue
								}
								if πTemp001, πE = πg.IAdd(πF, πTemp002, πg.NewInt(1).ToObject()); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µnum_lines, "num_lines"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µside, "side"); πE != nil {
									continue
								}
								πTemp003 = µside
								if πE = πg.SetItem(πF, µnum_lines, πTemp003, πTemp001); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µformat_key, "format_key"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
									continue
								}
								πTemp001 = πg.GetBool(µformat_key == πTemp002).ToObject()
								if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
									continue
								}
								if πTemp004 {
									goto Label1
								}
								goto Label2
								// line 1495: if format_key is None:
								πF.SetLineno(1495)
							Label1:
								// line 1496: return (num_lines[side],lines.pop(0)[2:])
								πF.SetLineno(1496)
								if πE = πg.CheckLocal(πF, µside, "side"); πE != nil {
									continue
								}
								πTemp002 = µside
								if πE = πg.CheckLocal(πF, µnum_lines, "num_lines"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.GetItem(πF, µnum_lines, πTemp002); πE != nil {
									continue
								}
								if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(2).ToObject(), πg.None, πg.None}, nil); πE != nil {
									continue
								}
								πTemp006 = πF.MakeArgs(1)
								πTemp006[0] = πg.NewInt(0).ToObject()
								if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
									continue
								}
								if πTemp007, πE = πg.GetAttr(πF, µlines, ßpop, nil); πE != nil {
									continue
								}
								if πTemp008, πE = πTemp007.Call(πF, πTemp006, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp006)
								if πTemp005, πE = πg.GetItem(πF, πTemp008, πTemp002); πE != nil {
									continue
								}
								πTemp001 = πg.NewTuple2(πTemp003, πTemp005).ToObject()
								πR = πTemp001
								continue
								goto Label2
							Label2:
								if πE = πg.CheckLocal(πF, µformat_key, "format_key"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.Eq(πF, µformat_key, πg.NewStr("?").ToObject()); πE != nil {
									continue
								}
								if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
									continue
								}
								if πTemp004 {
									goto Label3
								}
								goto Label4
								// line 1498: if format_key == '?':
								πF.SetLineno(1498)
							Label3:
								// line 1499: text, markers = lines.pop(0), lines.pop(0)
								πF.SetLineno(1499)
								πTemp006 = πF.MakeArgs(1)
								πTemp006[0] = πg.NewInt(0).ToObject()
								if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.GetAttr(πF, µlines, ßpop, nil); πE != nil {
									continue
								}
								if πTemp003, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp006)
								πTemp006 = πF.MakeArgs(1)
								πTemp006[0] = πg.NewInt(0).ToObject()
								if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.GetAttr(πF, µlines, ßpop, nil); πE != nil {
									continue
								}
								if πTemp005, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp006)
								πTemp001 = πg.NewTuple2(πTemp003, πTemp005).ToObject()
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp003}}}, πTemp001); πE != nil {
									continue
								}
								µtext = πTemp002
								µmarkers = πTemp003
								// line 1501: sub_info = []
								πF.SetLineno(1501)
								πTemp006 = make([]*πg.Object, 0)
								πTemp001 = πg.NewList(πTemp006...).ToObject()
								µsub_info = πTemp001
								// line 1502: def record_sub_info(match_object,sub_info=sub_info):
								πF.SetLineno(1502)
								πTemp009 = make([]πg.Param, 2)
								πTemp009[0] = πg.Param{Name: "match_object", Def: nil}
								if πE = πg.CheckLocal(πF, µsub_info, "sub_info"); πE != nil {
									continue
								}
								πTemp009[1] = πg.Param{Name: "sub_info", Def: µsub_info}
								πTemp001 = πg.NewFunction(πg.NewCode("record_sub_info", "/usr/lib/python2.7/difflib.py", πTemp009, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
									var µmatch_object *πg.Object = πArgs[0]; _ = µmatch_object
									var µsub_info *πg.Object = πArgs[1]; _ = µsub_info
									var πTemp001 []*πg.Object
									_ = πTemp001
									var πTemp002 []*πg.Object
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
										// line 1503: sub_info.append([match_object.group(1)[0],match_object.span()])
										πF.SetLineno(1503)
										πTemp001 = πF.MakeArgs(1)
										πTemp002 = make([]*πg.Object, 2)
										πTemp003 = πg.NewInt(0).ToObject()
										πTemp005 = πF.MakeArgs(1)
										πTemp005[0] = πg.NewInt(1).ToObject()
										if πE = πg.CheckLocal(πF, µmatch_object, "match_object"); πE != nil {
											continue
										}
										if πTemp006, πE = πg.GetAttr(πF, µmatch_object, ßgroup, nil); πE != nil {
											continue
										}
										if πTemp007, πE = πTemp006.Call(πF, πTemp005, nil); πE != nil {
											continue
										}
										πF.FreeArgs(πTemp005)
										if πTemp004, πE = πg.GetItem(πF, πTemp007, πTemp003); πE != nil {
											continue
										}
										πTemp002[0] = πTemp004
										if πE = πg.CheckLocal(πF, µmatch_object, "match_object"); πE != nil {
											continue
										}
										if πTemp003, πE = πg.GetAttr(πF, µmatch_object, ßspan, nil); πE != nil {
											continue
										}
										if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
											continue
										}
										πTemp002[1] = πTemp004
										πTemp003 = πg.NewList(πTemp002...).ToObject()
										πTemp001[0] = πTemp003
										if πE = πg.CheckLocal(πF, µsub_info, "sub_info"); πE != nil {
											continue
										}
										if πTemp003, πE = πg.GetAttr(πF, µsub_info, ßappend, nil); πE != nil {
											continue
										}
										if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
											continue
										}
										πF.FreeArgs(πTemp001)
										// line 1504: return match_object.group(1)
										πF.SetLineno(1504)
										πTemp001 = πF.MakeArgs(1)
										πTemp001[0] = πg.NewInt(1).ToObject()
										if πE = πg.CheckLocal(πF, µmatch_object, "match_object"); πE != nil {
											continue
										}
										if πTemp003, πE = πg.GetAttr(πF, µmatch_object, ßgroup, nil); πE != nil {
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
								µrecord_sub_info = πTemp001
								// line 1505: change_re.sub(record_sub_info,markers)
								πF.SetLineno(1505)
								πTemp006 = πF.MakeArgs(2)
								if πE = πg.CheckLocal(πF, µrecord_sub_info, "record_sub_info"); πE != nil {
									continue
								}
								πTemp006[0] = µrecord_sub_info
								if πE = πg.CheckLocal(πF, µmarkers, "markers"); πE != nil {
									continue
								}
								πTemp006[1] = µmarkers
								if πE = πg.CheckLocal(πF, µchange_re, "change_re"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.GetAttr(πF, µchange_re, ßsub, nil); πE != nil {
									continue
								}
								if πTemp003, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp006)
								if πTemp005, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
									continue
								}
								if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.None, πTemp005}, nil); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µsub_info, "sub_info"); πE != nil {
									continue
								}
								if πTemp005, πE = πg.GetItem(πF, µsub_info, πTemp003); πE != nil {
									continue
								}
								if πTemp002, πE = πg.Iter(πF, πTemp005); πE != nil {
									continue
								}
								πF.PushCheckpoint(7)
								πTemp004 = false
							Label6:
								if πE != nil || πR != nil {
									continue
								}
								if πTemp004 {
									πF.PopCheckpoint()
									goto Label8
								}
								if πTemp003, πE = πg.Next(πF, πTemp002); πE != nil {
									isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
									if exc != nil {
										πE = exc
									} else if isStop {
										πE = nil
										πF.RestoreExc(nil, nil)
									}
									πTemp010 = !isStop
								} else {
									πTemp010 = true
									if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp007}, πg.TieTarget{Target: &πTemp008}}}}}, πTemp003); πE != nil {
										continue
									}
									µkey = πTemp005
									µbegin = πTemp007
									µend = πTemp008
								}
								if πE != nil || !πTemp010 {
									continue
								}
								πF.PushCheckpoint(6)            
								// line 1509: text = text[0:begin]+'\0'+key+text[begin:end]+'\1'+text[end:]
								πF.SetLineno(1509)
								if πE = πg.CheckLocal(πF, µbegin, "begin"); πE != nil {
									continue
								}
								if πTemp012, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(0).ToObject(), µbegin, πg.None}, nil); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
									continue
								}
								if πTemp013, πE = πg.GetItem(πF, µtext, πTemp012); πE != nil {
									continue
								}
								if πTemp011, πE = πg.Add(πF, πTemp013, πg.NewStr("\x00").ToObject()); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
									continue
								}
								if πTemp008, πE = πg.Add(πF, πTemp011, µkey); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µbegin, "begin"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
									continue
								}
								if πTemp011, πE = πg.SliceType.Call(πF, πg.Args{µbegin, µend, πg.None}, nil); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
									continue
								}
								if πTemp012, πE = πg.GetItem(πF, µtext, πTemp011); πE != nil {
									continue
								}
								if πTemp007, πE = πg.Add(πF, πTemp008, πTemp012); πE != nil {
									continue
								}
								if πTemp005, πE = πg.Add(πF, πTemp007, πg.NewStr("\x01").ToObject()); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
									continue
								}
								if πTemp007, πE = πg.SliceType.Call(πF, πg.Args{µend, πg.None, πg.None}, nil); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
									continue
								}
								if πTemp008, πE = πg.GetItem(πF, µtext, πTemp007); πE != nil {
									continue
								}
								if πTemp003, πE = πg.Add(πF, πTemp005, πTemp008); πE != nil {
									continue
								}
								µtext = πTemp003
								continue
							Label7:
								if πE != nil || πR != nil {
									continue
								}
							Label8:
								// line 1510: text = text[2:]
								πF.SetLineno(1510)
								if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(2).ToObject(), πg.None, πg.None}, nil); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.GetItem(πF, µtext, πTemp002); πE != nil {
									continue
								}
								µtext = πTemp003
								goto Label5
							Label4:
								// line 1513: text = lines.pop(0)[2:]
								πF.SetLineno(1513)
								if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(2).ToObject(), πg.None, πg.None}, nil); πE != nil {
									continue
								}
								πTemp006 = πF.MakeArgs(1)
								πTemp006[0] = πg.NewInt(0).ToObject()
								if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
									continue
								}
								if πTemp005, πE = πg.GetAttr(πF, µlines, ßpop, nil); πE != nil {
									continue
								}
								if πTemp007, πE = πTemp005.Call(πF, πTemp006, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp006)
								if πTemp003, πE = πg.GetItem(πF, πTemp007, πTemp002); πE != nil {
									continue
								}
								µtext = πTemp003
								if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
									continue
								}
								if πTemp004, πE = πg.IsTrue(πF, µtext); πE != nil {
									continue
								}
								πTemp002 = πg.GetBool(!πTemp004).ToObject()
								if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
									continue
								}
								if πTemp004 {
									goto Label9
								}
								goto Label10
								// line 1516: if not text:
								πF.SetLineno(1516)
							Label9:
								// line 1517: text = ' '
								πF.SetLineno(1517)
								µtext = πg.NewStr(" ").ToObject()
								goto Label10
							Label10:
								// line 1519: text = '\0' + format_key + text + '\1'
								πF.SetLineno(1519)
								if πE = πg.CheckLocal(πF, µformat_key, "format_key"); πE != nil {
									continue
								}
								if πTemp005, πE = πg.Add(πF, πg.NewStr("\x00").ToObject(), µformat_key); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.Add(πF, πTemp005, µtext); πE != nil {
									continue
								}
								if πTemp002, πE = πg.Add(πF, πTemp003, πg.NewStr("\x01").ToObject()); πE != nil {
									continue
								}
								µtext = πTemp002
								goto Label5
							Label5:
								// line 1523: return (num_lines[side],text)
								πF.SetLineno(1523)
								if πE = πg.CheckLocal(πF, µside, "side"); πE != nil {
									continue
								}
								πTemp003 = µside
								if πE = πg.CheckLocal(πF, µnum_lines, "num_lines"); πE != nil {
									continue
								}
								if πTemp005, πE = πg.GetItem(πF, µnum_lines, πTemp003); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
									continue
								}
								πTemp002 = πg.NewTuple2(πTemp005, µtext).ToObject()
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
						µ_make_line = πTemp001
						// line 1470: """Returns line of text with user's change markup and line formatting.
						πF.SetLineno(1470)
						if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewStr("Returns line of text with user's change markup and line formatting.\n\n        lines -- list of lines from the ndiff generator to produce a line of\n                 text from.  When producing the line of text to return, the\n                 lines used are removed from this list.\n        format_key -- '+' return first line in list with \"add\" markup around\n                          the entire line.\n                      '-' return first line in list with \"delete\" markup around\n                          the entire line.\n                      '?' return first line in list with add/delete/change\n                          intraline markup (indices obtained from second line)\n                      None return first line in list with no markup\n        side -- indice into the num_lines list (0=from,1=to)\n        num_lines -- from/to current line number.  This is NOT intended to be a\n                     passed parameter.  It is present as a keyword argument to\n                     maintain memory of the current line numbers between calls\n                     of this function.\n\n        Note, this function is purposefully not defined at the module scope so\n        that data it needs from its parent function (within whose context it\n        is defined) does not need to be of module scope.\n        ").ToObject()); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µ_make_line, "_make_line"); πE != nil {
							continue
						}
						if πE = πg.SetAttr(πF, µ_make_line, ß__doc__, πTemp003); πE != nil {
							continue
						}
						// line 1525: def _line_iterator():
						πF.SetLineno(1525)
						πTemp004 = make([]πg.Param, 0)
						πTemp003 = πg.NewFunction(πg.NewCode("_line_iterator", "/usr/lib/python2.7/difflib.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
							var µlines *πg.Object = πg.UnboundLocal; _ = µlines
							var µnum_blanks_pending *πg.Object = πg.UnboundLocal; _ = µnum_blanks_pending
							var µnum_blanks_to_yield *πg.Object = πg.UnboundLocal; _ = µnum_blanks_to_yield
							var µs *πg.Object = πg.UnboundLocal; _ = µs
							var µfrom_line *πg.Object = πg.UnboundLocal; _ = µfrom_line
							var µto_line *πg.Object = πg.UnboundLocal; _ = µto_line
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
							var πTemp007 bool
							_ = πTemp007
							var πTemp008 *πg.BaseException
							_ = πTemp008
							var πTemp009 *πg.Traceback
							_ = πTemp009
							var πTemp010 []πg.Param
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
							var πR *πg.Object; _ = πR
							var πE *πg.BaseException; _ = πE
							return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
								for ; πF.State() >= 0; πF.PopCheckpoint() {
									switch πF.State() {
									case 0:
									case 1: goto Label1
									case 2: goto Label2
									case 35: goto Label35
									case 4: goto Label4
									case 5: goto Label5
									case 33: goto Label33
									case 8: goto Label8
									case 41: goto Label41
									case 34: goto Label34
									case 37: goto Label37
									case 22: goto Label22
									case 23: goto Label23
									case 24: goto Label24
									case 25: goto Label25
									case 26: goto Label26
									case 27: goto Label27
									case 28: goto Label28
									case 29: goto Label29
									case 30: goto Label30
									case 31: goto Label31
									default: panic("unexpected function state")
									}
									// line 1526: """Yields from/to lines of text with a change indication.
									πF.SetLineno(1526)
									// line 1539: lines = []
									πF.SetLineno(1539)
									πTemp001 = make([]*πg.Object, 0)
									πTemp002 = πg.NewList(πTemp001...).ToObject()
									µlines = πTemp002
									// line 1540: num_blanks_pending, num_blanks_to_yield = 0, 0
									πF.SetLineno(1540)
									πTemp002 = πg.NewTuple2(πg.NewInt(0).ToObject(), πg.NewInt(0).ToObject()).ToObject()
									if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp004}}}, πTemp002); πE != nil {
										continue
									}
									µnum_blanks_pending = πTemp003
									µnum_blanks_to_yield = πTemp004
									// line 1541: while True:
									πF.SetLineno(1541)
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
									if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
										continue
									}
									if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
										continue
									}
									if πE != nil || !πTemp006 {
										continue
									}
									πF.PushCheckpoint(1)            
									// line 1545: while len(lines) < 4:
									πF.SetLineno(1545)
									πF.PushCheckpoint(5)
									πTemp006 = false
								Label4:
									if πE != nil || πR != nil {
										continue
									}
									if πTemp006 {
										πF.PopCheckpoint()
										goto Label6
									}
									πTemp001 = πF.MakeArgs(1)
									if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
										continue
									}
									πTemp001[0] = µlines
									if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
										continue
									}
									if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp001)
									if πTemp002, πE = πg.LT(πF, πTemp004, πg.NewInt(4).ToObject()); πE != nil {
										continue
									}
									if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
										continue
									}
									if πE != nil || !πTemp007 {
										continue
									}
									πF.PushCheckpoint(4)            
									// line 1546: try:
									πF.SetLineno(1546)
									πF.PushCheckpoint(8)
									// line 1547: lines.append(diff_lines_iterator.next())
									πF.SetLineno(1547)
									πTemp001 = πF.MakeArgs(1)
									if πE = πg.CheckLocal(πF, µdiff_lines_iterator, "diff_lines_iterator"); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, µdiff_lines_iterator, ßnext, nil); πE != nil {
										continue
									}
									if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
										continue
									}
									πTemp001[0] = πTemp003
									if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, µlines, ßappend, nil); πE != nil {
										continue
									}
									if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp001)
									πF.PopCheckpoint()
									goto Label7
								Label8:
									if πE == nil {
									  continue
									}
									πE = nil
									πTemp008, πTemp009 = πF.ExcInfo()
									if πTemp002, πE = πg.ResolveGlobal(πF, ßStopIteration); πE != nil {
										continue
									}
									if πTemp007, πE = πg.IsInstance(πF, πTemp008.ToObject(), πTemp002); πE != nil {
										continue
									}
									if πTemp007 {
										goto Label9
									}
									πE = πF.Raise(πTemp008.ToObject(), nil, πTemp009.ToObject())
									continue
									// line 1548: except StopIteration:
									πF.SetLineno(1548)
								Label9:
									// line 1549: lines.append('X')
									πF.SetLineno(1549)
									πTemp001 = πF.MakeArgs(1)
									πTemp001[0] = ßX.ToObject()
									if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, µlines, ßappend, nil); πE != nil {
										continue
									}
									if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp001)
									πF.RestoreExc(nil, nil)
									goto Label7
								Label7:
									continue
								Label5:
									if πE != nil || πR != nil {
										continue
									}
								Label6:
									// line 1550: s = ''.join([line[0] for line in lines])
									πF.SetLineno(1550)
									πTemp001 = πF.MakeArgs(1)
									πTemp010 = make([]πg.Param, 0)
									πTemp003 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/difflib.py", πTemp010, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
										var µline *πg.Object = πg.UnboundLocal; _ = µline
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
												if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
													continue
												}
												if πTemp001, πE = πg.Iter(πF, µlines); πE != nil {
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
													µline = πTemp004
												}
												if πE != nil || !πTemp003 {
													continue
												}
												πF.PushCheckpoint(1)            
												// line 1550: s = ''.join([line[0] for line in lines])
												πF.SetLineno(1550)
												πTemp004 = πg.NewInt(0).ToObject()
												if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
													continue
												}
												if πTemp005, πE = πg.GetItem(πF, µline, πTemp004); πE != nil {
													continue
												}
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
									if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
										continue
									}
									if πTemp002, πE = πg.ListType.Call(πF, πg.Args{πTemp004}, nil); πE != nil {
										continue
									}
									πTemp001[0] = πTemp002
									if πTemp002, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
										continue
									}
									if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp001)
									µs = πTemp004
									πTemp001 = πF.MakeArgs(1)
									πTemp001[0] = ßX.ToObject()
									if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, µs, ßstartswith, nil); πE != nil {
										continue
									}
									if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp001)
									if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
										continue
									}
									if πTemp006 {
										goto Label10
									}
									πTemp001 = πF.MakeArgs(1)
									πTemp001[0] = πg.NewStr("-?+?").ToObject()
									if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, µs, ßstartswith, nil); πE != nil {
										continue
									}
									if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp001)
									if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
										continue
									}
									if πTemp006 {
										goto Label11
									}
									πTemp001 = πF.MakeArgs(1)
									πTemp001[0] = πg.NewStr("--++").ToObject()
									if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, µs, ßstartswith, nil); πE != nil {
										continue
									}
									if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp001)
									if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
										continue
									}
									if πTemp006 {
										goto Label12
									}
									πTemp001 = πF.MakeArgs(1)
									πTemp002 = πg.NewTuple3(πg.NewStr("--?+").ToObject(), πg.NewStr("--+").ToObject(), πg.NewStr("- ").ToObject()).ToObject()
									πTemp001[0] = πTemp002
									if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, µs, ßstartswith, nil); πE != nil {
										continue
									}
									if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp001)
									if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
										continue
									}
									if πTemp006 {
										goto Label13
									}
									πTemp001 = πF.MakeArgs(1)
									πTemp001[0] = πg.NewStr("-+?").ToObject()
									if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, µs, ßstartswith, nil); πE != nil {
										continue
									}
									if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp001)
									if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
										continue
									}
									if πTemp006 {
										goto Label14
									}
									πTemp001 = πF.MakeArgs(1)
									πTemp001[0] = πg.NewStr("-?+").ToObject()
									if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, µs, ßstartswith, nil); πE != nil {
										continue
									}
									if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp001)
									if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
										continue
									}
									if πTemp006 {
										goto Label15
									}
									πTemp001 = πF.MakeArgs(1)
									πTemp001[0] = πg.NewStr("-").ToObject()
									if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, µs, ßstartswith, nil); πE != nil {
										continue
									}
									if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp001)
									if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
										continue
									}
									if πTemp006 {
										goto Label16
									}
									πTemp001 = πF.MakeArgs(1)
									πTemp001[0] = πg.NewStr("+--").ToObject()
									if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, µs, ßstartswith, nil); πE != nil {
										continue
									}
									if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp001)
									if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
										continue
									}
									if πTemp006 {
										goto Label17
									}
									πTemp001 = πF.MakeArgs(1)
									πTemp002 = πg.NewTuple2(πg.NewStr("+ ").ToObject(), πg.NewStr("+-").ToObject()).ToObject()
									πTemp001[0] = πTemp002
									if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, µs, ßstartswith, nil); πE != nil {
										continue
									}
									if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp001)
									if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
										continue
									}
									if πTemp006 {
										goto Label18
									}
									πTemp001 = πF.MakeArgs(1)
									πTemp001[0] = πg.NewStr("+").ToObject()
									if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, µs, ßstartswith, nil); πE != nil {
										continue
									}
									if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp001)
									if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
										continue
									}
									if πTemp006 {
										goto Label19
									}
									πTemp001 = πF.MakeArgs(1)
									πTemp001[0] = πg.NewStr(" ").ToObject()
									if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, µs, ßstartswith, nil); πE != nil {
										continue
									}
									if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp001)
									if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
										continue
									}
									if πTemp006 {
										goto Label20
									}
									goto Label21
									// line 1551: if s.startswith('X'):
									πF.SetLineno(1551)
								Label10:
									// line 1555: num_blanks_to_yield = num_blanks_pending
									πF.SetLineno(1555)
									if πE = πg.CheckLocal(πF, µnum_blanks_pending, "num_blanks_pending"); πE != nil {
										continue
									}
									µnum_blanks_to_yield = µnum_blanks_pending
									goto Label21
									// line 1556: elif s.startswith('-?+?'):
									πF.SetLineno(1556)
								Label11:
									// line 1558: yield _make_line(lines,'?',0), _make_line(lines,'?',1), True
									πF.SetLineno(1558)
									πTemp001 = πF.MakeArgs(3)
									if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
										continue
									}
									πTemp001[0] = µlines
									πTemp001[1] = πg.NewStr("?").ToObject()
									πTemp001[2] = πg.NewInt(0).ToObject()
									if πE = πg.CheckLocal(πF, µ_make_line, "_make_line"); πE != nil {
										continue
									}
									if πTemp004, πE = µ_make_line.Call(πF, πTemp001, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp001)
									πTemp001 = πF.MakeArgs(3)
									if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
										continue
									}
									πTemp001[0] = µlines
									πTemp001[1] = πg.NewStr("?").ToObject()
									πTemp001[2] = πg.NewInt(1).ToObject()
									if πE = πg.CheckLocal(πF, µ_make_line, "_make_line"); πE != nil {
										continue
									}
									if πTemp011, πE = µ_make_line.Call(πF, πTemp001, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp001)
									if πTemp012, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
										continue
									}
									πTemp002 = πg.NewTuple3(πTemp004, πTemp011, πTemp012).ToObject()
									πF.PushCheckpoint(22)
									return πTemp002, nil
								Label22:
									πTemp004 = πSent
									// line 1559: continue
									πF.SetLineno(1559)
									continue
									goto Label21
									// line 1560: elif s.startswith('--++'):
									πF.SetLineno(1560)
								Label12:
									// line 1563: num_blanks_pending -= 1
									πF.SetLineno(1563)
									if πE = πg.CheckLocal(πF, µnum_blanks_pending, "num_blanks_pending"); πE != nil {
										continue
									}
									if πTemp004, πE = πg.ISub(πF, µnum_blanks_pending, πg.NewInt(1).ToObject()); πE != nil {
										continue
									}
									µnum_blanks_pending = πTemp004
									// line 1564: yield _make_line(lines,'-',0), None, True
									πF.SetLineno(1564)
									πTemp001 = πF.MakeArgs(3)
									if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
										continue
									}
									πTemp001[0] = µlines
									πTemp001[1] = πg.NewStr("-").ToObject()
									πTemp001[2] = πg.NewInt(0).ToObject()
									if πE = πg.CheckLocal(πF, µ_make_line, "_make_line"); πE != nil {
										continue
									}
									if πTemp011, πE = µ_make_line.Call(πF, πTemp001, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp001)
									if πTemp012, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
										continue
									}
									if πTemp013, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
										continue
									}
									πTemp004 = πg.NewTuple3(πTemp011, πTemp012, πTemp013).ToObject()
									πF.PushCheckpoint(23)
									return πTemp004, nil
								Label23:
									πTemp011 = πSent
									// line 1565: continue
									πF.SetLineno(1565)
									continue
									goto Label21
									// line 1566: elif s.startswith(('--?+', '--+', '- ')):
									πF.SetLineno(1566)
								Label13:
									// line 1569: from_line,to_line = _make_line(lines,'-',0), None
									πF.SetLineno(1569)
									πTemp001 = πF.MakeArgs(3)
									if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
										continue
									}
									πTemp001[0] = µlines
									πTemp001[1] = πg.NewStr("-").ToObject()
									πTemp001[2] = πg.NewInt(0).ToObject()
									if πE = πg.CheckLocal(πF, µ_make_line, "_make_line"); πE != nil {
										continue
									}
									if πTemp012, πE = µ_make_line.Call(πF, πTemp001, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp001)
									if πTemp013, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
										continue
									}
									πTemp011 = πg.NewTuple2(πTemp012, πTemp013).ToObject()
									if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp012}, πg.TieTarget{Target: &πTemp013}}}, πTemp011); πE != nil {
										continue
									}
									µfrom_line = πTemp012
									µto_line = πTemp013
									// line 1570: num_blanks_to_yield,num_blanks_pending = num_blanks_pending-1,0
									πF.SetLineno(1570)
									if πE = πg.CheckLocal(πF, µnum_blanks_pending, "num_blanks_pending"); πE != nil {
										continue
									}
									if πTemp012, πE = πg.Sub(πF, µnum_blanks_pending, πg.NewInt(1).ToObject()); πE != nil {
										continue
									}
									πTemp011 = πg.NewTuple2(πTemp012, πg.NewInt(0).ToObject()).ToObject()
									if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp012}, πg.TieTarget{Target: &πTemp013}}}, πTemp011); πE != nil {
										continue
									}
									µnum_blanks_to_yield = πTemp012
									µnum_blanks_pending = πTemp013
									goto Label21
									// line 1571: elif s.startswith('-+?'):
									πF.SetLineno(1571)
								Label14:
									// line 1573: yield _make_line(lines,None,0), _make_line(lines,'?',1), True
									πF.SetLineno(1573)
									πTemp001 = πF.MakeArgs(3)
									if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
										continue
									}
									πTemp001[0] = µlines
									if πTemp012, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
										continue
									}
									πTemp001[1] = πTemp012
									πTemp001[2] = πg.NewInt(0).ToObject()
									if πE = πg.CheckLocal(πF, µ_make_line, "_make_line"); πE != nil {
										continue
									}
									if πTemp012, πE = µ_make_line.Call(πF, πTemp001, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp001)
									πTemp001 = πF.MakeArgs(3)
									if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
										continue
									}
									πTemp001[0] = µlines
									πTemp001[1] = πg.NewStr("?").ToObject()
									πTemp001[2] = πg.NewInt(1).ToObject()
									if πE = πg.CheckLocal(πF, µ_make_line, "_make_line"); πE != nil {
										continue
									}
									if πTemp013, πE = µ_make_line.Call(πF, πTemp001, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp001)
									if πTemp014, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
										continue
									}
									πTemp011 = πg.NewTuple3(πTemp012, πTemp013, πTemp014).ToObject()
									πF.PushCheckpoint(24)
									return πTemp011, nil
								Label24:
									πTemp012 = πSent
									// line 1574: continue
									πF.SetLineno(1574)
									continue
									goto Label21
									// line 1575: elif s.startswith('-?+'):
									πF.SetLineno(1575)
								Label15:
									// line 1577: yield _make_line(lines,'?',0), _make_line(lines,None,1), True
									πF.SetLineno(1577)
									πTemp001 = πF.MakeArgs(3)
									if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
										continue
									}
									πTemp001[0] = µlines
									πTemp001[1] = πg.NewStr("?").ToObject()
									πTemp001[2] = πg.NewInt(0).ToObject()
									if πE = πg.CheckLocal(πF, µ_make_line, "_make_line"); πE != nil {
										continue
									}
									if πTemp013, πE = µ_make_line.Call(πF, πTemp001, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp001)
									πTemp001 = πF.MakeArgs(3)
									if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
										continue
									}
									πTemp001[0] = µlines
									if πTemp014, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
										continue
									}
									πTemp001[1] = πTemp014
									πTemp001[2] = πg.NewInt(1).ToObject()
									if πE = πg.CheckLocal(πF, µ_make_line, "_make_line"); πE != nil {
										continue
									}
									if πTemp014, πE = µ_make_line.Call(πF, πTemp001, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp001)
									if πTemp015, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
										continue
									}
									πTemp012 = πg.NewTuple3(πTemp013, πTemp014, πTemp015).ToObject()
									πF.PushCheckpoint(25)
									return πTemp012, nil
								Label25:
									πTemp013 = πSent
									// line 1578: continue
									πF.SetLineno(1578)
									continue
									goto Label21
									// line 1579: elif s.startswith('-'):
									πF.SetLineno(1579)
								Label16:
									// line 1581: num_blanks_pending -= 1
									πF.SetLineno(1581)
									if πE = πg.CheckLocal(πF, µnum_blanks_pending, "num_blanks_pending"); πE != nil {
										continue
									}
									if πTemp013, πE = πg.ISub(πF, µnum_blanks_pending, πg.NewInt(1).ToObject()); πE != nil {
										continue
									}
									µnum_blanks_pending = πTemp013
									// line 1582: yield _make_line(lines,'-',0), None, True
									πF.SetLineno(1582)
									πTemp001 = πF.MakeArgs(3)
									if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
										continue
									}
									πTemp001[0] = µlines
									πTemp001[1] = πg.NewStr("-").ToObject()
									πTemp001[2] = πg.NewInt(0).ToObject()
									if πE = πg.CheckLocal(πF, µ_make_line, "_make_line"); πE != nil {
										continue
									}
									if πTemp014, πE = µ_make_line.Call(πF, πTemp001, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp001)
									if πTemp015, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
										continue
									}
									if πTemp016, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
										continue
									}
									πTemp013 = πg.NewTuple3(πTemp014, πTemp015, πTemp016).ToObject()
									πF.PushCheckpoint(26)
									return πTemp013, nil
								Label26:
									πTemp014 = πSent
									// line 1583: continue
									πF.SetLineno(1583)
									continue
									goto Label21
									// line 1584: elif s.startswith('+--'):
									πF.SetLineno(1584)
								Label17:
									// line 1587: num_blanks_pending += 1
									πF.SetLineno(1587)
									if πE = πg.CheckLocal(πF, µnum_blanks_pending, "num_blanks_pending"); πE != nil {
										continue
									}
									if πTemp014, πE = πg.IAdd(πF, µnum_blanks_pending, πg.NewInt(1).ToObject()); πE != nil {
										continue
									}
									µnum_blanks_pending = πTemp014
									// line 1588: yield None, _make_line(lines,'+',1), True
									πF.SetLineno(1588)
									if πTemp015, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
										continue
									}
									πTemp001 = πF.MakeArgs(3)
									if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
										continue
									}
									πTemp001[0] = µlines
									πTemp001[1] = πg.NewStr("+").ToObject()
									πTemp001[2] = πg.NewInt(1).ToObject()
									if πE = πg.CheckLocal(πF, µ_make_line, "_make_line"); πE != nil {
										continue
									}
									if πTemp016, πE = µ_make_line.Call(πF, πTemp001, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp001)
									if πTemp017, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
										continue
									}
									πTemp014 = πg.NewTuple3(πTemp015, πTemp016, πTemp017).ToObject()
									πF.PushCheckpoint(27)
									return πTemp014, nil
								Label27:
									πTemp015 = πSent
									// line 1589: continue
									πF.SetLineno(1589)
									continue
									goto Label21
									// line 1590: elif s.startswith(('+ ', '+-')):
									πF.SetLineno(1590)
								Label18:
									// line 1592: from_line, to_line = None, _make_line(lines,'+',1)
									πF.SetLineno(1592)
									if πTemp016, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
										continue
									}
									πTemp001 = πF.MakeArgs(3)
									if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
										continue
									}
									πTemp001[0] = µlines
									πTemp001[1] = πg.NewStr("+").ToObject()
									πTemp001[2] = πg.NewInt(1).ToObject()
									if πE = πg.CheckLocal(πF, µ_make_line, "_make_line"); πE != nil {
										continue
									}
									if πTemp017, πE = µ_make_line.Call(πF, πTemp001, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp001)
									πTemp015 = πg.NewTuple2(πTemp016, πTemp017).ToObject()
									if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp016}, πg.TieTarget{Target: &πTemp017}}}, πTemp015); πE != nil {
										continue
									}
									µfrom_line = πTemp016
									µto_line = πTemp017
									// line 1593: num_blanks_to_yield,num_blanks_pending = num_blanks_pending+1,0
									πF.SetLineno(1593)
									if πE = πg.CheckLocal(πF, µnum_blanks_pending, "num_blanks_pending"); πE != nil {
										continue
									}
									if πTemp016, πE = πg.Add(πF, µnum_blanks_pending, πg.NewInt(1).ToObject()); πE != nil {
										continue
									}
									πTemp015 = πg.NewTuple2(πTemp016, πg.NewInt(0).ToObject()).ToObject()
									if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp016}, πg.TieTarget{Target: &πTemp017}}}, πTemp015); πE != nil {
										continue
									}
									µnum_blanks_to_yield = πTemp016
									µnum_blanks_pending = πTemp017
									goto Label21
									// line 1594: elif s.startswith('+'):
									πF.SetLineno(1594)
								Label19:
									// line 1596: num_blanks_pending += 1
									πF.SetLineno(1596)
									if πE = πg.CheckLocal(πF, µnum_blanks_pending, "num_blanks_pending"); πE != nil {
										continue
									}
									if πTemp015, πE = πg.IAdd(πF, µnum_blanks_pending, πg.NewInt(1).ToObject()); πE != nil {
										continue
									}
									µnum_blanks_pending = πTemp015
									// line 1597: yield None, _make_line(lines,'+',1), True
									πF.SetLineno(1597)
									if πTemp016, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
										continue
									}
									πTemp001 = πF.MakeArgs(3)
									if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
										continue
									}
									πTemp001[0] = µlines
									πTemp001[1] = πg.NewStr("+").ToObject()
									πTemp001[2] = πg.NewInt(1).ToObject()
									if πE = πg.CheckLocal(πF, µ_make_line, "_make_line"); πE != nil {
										continue
									}
									if πTemp017, πE = µ_make_line.Call(πF, πTemp001, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp001)
									if πTemp018, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
										continue
									}
									πTemp015 = πg.NewTuple3(πTemp016, πTemp017, πTemp018).ToObject()
									πF.PushCheckpoint(28)
									return πTemp015, nil
								Label28:
									πTemp016 = πSent
									// line 1598: continue
									πF.SetLineno(1598)
									continue
									goto Label21
									// line 1599: elif s.startswith(' '):
									πF.SetLineno(1599)
								Label20:
									// line 1601: yield _make_line(lines[:],None,0),_make_line(lines,None,1),False
									πF.SetLineno(1601)
									πTemp001 = πF.MakeArgs(3)
									if πTemp017, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.None, πg.None}, nil); πE != nil {
										continue
									}
									if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
										continue
									}
									if πTemp018, πE = πg.GetItem(πF, µlines, πTemp017); πE != nil {
										continue
									}
									πTemp001[0] = πTemp018
									if πTemp017, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
										continue
									}
									πTemp001[1] = πTemp017
									πTemp001[2] = πg.NewInt(0).ToObject()
									if πE = πg.CheckLocal(πF, µ_make_line, "_make_line"); πE != nil {
										continue
									}
									if πTemp017, πE = µ_make_line.Call(πF, πTemp001, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp001)
									πTemp001 = πF.MakeArgs(3)
									if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
										continue
									}
									πTemp001[0] = µlines
									if πTemp018, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
										continue
									}
									πTemp001[1] = πTemp018
									πTemp001[2] = πg.NewInt(1).ToObject()
									if πE = πg.CheckLocal(πF, µ_make_line, "_make_line"); πE != nil {
										continue
									}
									if πTemp018, πE = µ_make_line.Call(πF, πTemp001, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp001)
									if πTemp019, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
										continue
									}
									πTemp016 = πg.NewTuple3(πTemp017, πTemp018, πTemp019).ToObject()
									πF.PushCheckpoint(29)
									return πTemp016, nil
								Label29:
									πTemp017 = πSent
									// line 1602: continue
									πF.SetLineno(1602)
									continue
									goto Label21
								Label21:
									// line 1605: while(num_blanks_to_yield < 0):
									πF.SetLineno(1605)
									πF.PushCheckpoint(31)
									πTemp006 = false
								Label30:
									if πE != nil || πR != nil {
										continue
									}
									if πTemp006 {
										πF.PopCheckpoint()
										goto Label32
									}
									if πE = πg.CheckLocal(πF, µnum_blanks_to_yield, "num_blanks_to_yield"); πE != nil {
										continue
									}
									if πTemp017, πE = πg.LT(πF, µnum_blanks_to_yield, πg.NewInt(0).ToObject()); πE != nil {
										continue
									}
									if πTemp007, πE = πg.IsTrue(πF, πTemp017); πE != nil {
										continue
									}
									if πE != nil || !πTemp007 {
										continue
									}
									πF.PushCheckpoint(30)            
									// line 1606: num_blanks_to_yield += 1
									πF.SetLineno(1606)
									if πE = πg.CheckLocal(πF, µnum_blanks_to_yield, "num_blanks_to_yield"); πE != nil {
										continue
									}
									if πTemp017, πE = πg.IAdd(πF, µnum_blanks_to_yield, πg.NewInt(1).ToObject()); πE != nil {
										continue
									}
									µnum_blanks_to_yield = πTemp017
									// line 1607: yield None,('','\n'),True
									πF.SetLineno(1607)
									if πTemp018, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
										continue
									}
									πTemp019 = πg.NewTuple2(ß.ToObject(), πg.NewStr("\n").ToObject()).ToObject()
									if πTemp020, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
										continue
									}
									πTemp017 = πg.NewTuple3(πTemp018, πTemp019, πTemp020).ToObject()
									πF.PushCheckpoint(33)
									return πTemp017, nil
								Label33:
									πTemp018 = πSent
									continue
								Label31:
									if πE != nil || πR != nil {
										continue
									}
								Label32:
									// line 1608: while(num_blanks_to_yield > 0):
									πF.SetLineno(1608)
									πF.PushCheckpoint(35)
									πTemp006 = false
								Label34:
									if πE != nil || πR != nil {
										continue
									}
									if πTemp006 {
										πF.PopCheckpoint()
										goto Label36
									}
									if πE = πg.CheckLocal(πF, µnum_blanks_to_yield, "num_blanks_to_yield"); πE != nil {
										continue
									}
									if πTemp018, πE = πg.GT(πF, µnum_blanks_to_yield, πg.NewInt(0).ToObject()); πE != nil {
										continue
									}
									if πTemp007, πE = πg.IsTrue(πF, πTemp018); πE != nil {
										continue
									}
									if πE != nil || !πTemp007 {
										continue
									}
									πF.PushCheckpoint(34)            
									// line 1609: num_blanks_to_yield -= 1
									πF.SetLineno(1609)
									if πE = πg.CheckLocal(πF, µnum_blanks_to_yield, "num_blanks_to_yield"); πE != nil {
										continue
									}
									if πTemp018, πE = πg.ISub(πF, µnum_blanks_to_yield, πg.NewInt(1).ToObject()); πE != nil {
										continue
									}
									µnum_blanks_to_yield = πTemp018
									// line 1610: yield ('','\n'),None,True
									πF.SetLineno(1610)
									πTemp019 = πg.NewTuple2(ß.ToObject(), πg.NewStr("\n").ToObject()).ToObject()
									if πTemp020, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
										continue
									}
									if πTemp021, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
										continue
									}
									πTemp018 = πg.NewTuple3(πTemp019, πTemp020, πTemp021).ToObject()
									πF.PushCheckpoint(37)
									return πTemp018, nil
								Label37:
									πTemp019 = πSent
									continue
								Label35:
									if πE != nil || πR != nil {
										continue
									}
								Label36:
									πTemp001 = πF.MakeArgs(1)
									πTemp001[0] = ßX.ToObject()
									if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
										continue
									}
									if πTemp019, πE = πg.GetAttr(πF, µs, ßstartswith, nil); πE != nil {
										continue
									}
									if πTemp020, πE = πTemp019.Call(πF, πTemp001, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp001)
									if πTemp006, πE = πg.IsTrue(πF, πTemp020); πE != nil {
										continue
									}
									if πTemp006 {
										goto Label38
									}
									goto Label39
									// line 1611: if s.startswith('X'):
									πF.SetLineno(1611)
								Label38:
									if πTemp019, πE = πg.ResolveGlobal(πF, ßStopIteration); πE != nil {
										continue
									}
									// line 1612: raise StopIteration
									πF.SetLineno(1612)
									πE = πF.Raise(πTemp019, nil, nil)
									continue
									goto Label40
								Label39:
									// line 1614: yield from_line,to_line,True
									πF.SetLineno(1614)
									if πE = πg.CheckLocal(πF, µfrom_line, "from_line"); πE != nil {
										continue
									}
									if πE = πg.CheckLocal(πF, µto_line, "to_line"); πE != nil {
										continue
									}
									if πTemp020, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
										continue
									}
									πTemp019 = πg.NewTuple3(µfrom_line, µto_line, πTemp020).ToObject()
									πF.PushCheckpoint(41)
									return πTemp019, nil
								Label41:
									πTemp020 = πSent
									goto Label40
								Label40:
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
						µ_line_iterator = πTemp003
						// line 1526: """Yields from/to lines of text with a change indication.
						πF.SetLineno(1526)
						if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πg.NewStr("Yields from/to lines of text with a change indication.\n\n        This function is an iterator.  It itself pulls lines from a\n        differencing iterator, processes them and yields them.  When it can\n        it yields both a \"from\" and a \"to\" line, otherwise it will yield one\n        or the other.  In addition to yielding the lines of from/to text, a\n        boolean flag is yielded to indicate if the text line(s) have\n        differences in them.\n\n        Note, this function is purposefully not defined at the module scope so\n        that data it needs from its parent function (within whose context it\n        is defined) does not need to be of module scope.\n        ").ToObject()); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µ_line_iterator, "_line_iterator"); πE != nil {
							continue
						}
						if πE = πg.SetAttr(πF, µ_line_iterator, ß__doc__, πTemp005); πE != nil {
							continue
						}
						// line 1616: def _line_pair_iterator():
						πF.SetLineno(1616)
						πTemp004 = make([]πg.Param, 0)
						πTemp005 = πg.NewFunction(πg.NewCode("_line_pair_iterator", "/usr/lib/python2.7/difflib.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
							var µline_iterator *πg.Object = πg.UnboundLocal; _ = µline_iterator
							var µfromlines *πg.Object = πg.UnboundLocal; _ = µfromlines
							var µtolines *πg.Object = πg.UnboundLocal; _ = µtolines
							var µfrom_line *πg.Object = πg.UnboundLocal; _ = µfrom_line
							var µto_line *πg.Object = πg.UnboundLocal; _ = µto_line
							var µfound_diff *πg.Object = πg.UnboundLocal; _ = µfound_diff
							var µfromDiff *πg.Object = πg.UnboundLocal; _ = µfromDiff
							var µto_diff *πg.Object = πg.UnboundLocal; _ = µto_diff
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
							var πTemp007 bool
							_ = πTemp007
							var πTemp008 bool
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
									case 5: goto Label5
									case 13: goto Label13
									default: panic("unexpected function state")
									}
									// line 1617: """Yields from/to lines of text with a change indication.
									πF.SetLineno(1617)
									// line 1629: line_iterator = _line_iterator()
									πF.SetLineno(1629)
									if πE = πg.CheckLocal(πF, µ_line_iterator, "_line_iterator"); πE != nil {
										continue
									}
									if πTemp001, πE = µ_line_iterator.Call(πF, nil, nil); πE != nil {
										continue
									}
									µline_iterator = πTemp001
									// line 1630: fromlines,tolines=[],[]
									πF.SetLineno(1630)
									πTemp002 = make([]*πg.Object, 0)
									πTemp003 = πg.NewList(πTemp002...).ToObject()
									πTemp002 = make([]*πg.Object, 0)
									πTemp004 = πg.NewList(πTemp002...).ToObject()
									πTemp001 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
									if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp004}}}, πTemp001); πE != nil {
										continue
									}
									µfromlines = πTemp003
									µtolines = πTemp004
									// line 1631: while True:
									πF.SetLineno(1631)
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
									if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
										continue
									}
									if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
										continue
									}
									if πE != nil || !πTemp006 {
										continue
									}
									πF.PushCheckpoint(1)            
									// line 1633: while (len(fromlines)==0 or len(tolines)==0):
									πF.SetLineno(1633)
									πF.PushCheckpoint(5)
									πTemp006 = false
								Label4:
									if πE != nil || πR != nil {
										continue
									}
									if πTemp006 {
										πF.PopCheckpoint()
										goto Label6
									}
									πTemp002 = πF.MakeArgs(1)
									if πE = πg.CheckLocal(πF, µfromlines, "fromlines"); πE != nil {
										continue
									}
									πTemp002[0] = µfromlines
									if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
										continue
									}
									if πTemp009, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp002)
									if πTemp003, πE = πg.Eq(πF, πTemp009, πg.NewInt(0).ToObject()); πE != nil {
										continue
									}
									πTemp001 = πTemp003
									if πTemp008, πE = πg.IsTrue(πF, πTemp001); πE != nil {
										continue
									}
									if πTemp008 {
										goto Label7
									}
									πTemp002 = πF.MakeArgs(1)
									if πE = πg.CheckLocal(πF, µtolines, "tolines"); πE != nil {
										continue
									}
									πTemp002[0] = µtolines
									if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
										continue
									}
									if πTemp009, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp002)
									if πTemp003, πE = πg.Eq(πF, πTemp009, πg.NewInt(0).ToObject()); πE != nil {
										continue
									}
									πTemp001 = πTemp003
								Label7:
									if πTemp007, πE = πg.IsTrue(πF, πTemp001); πE != nil {
										continue
									}
									if πE != nil || !πTemp007 {
										continue
									}
									πF.PushCheckpoint(4)            
									// line 1634: from_line, to_line, found_diff =line_iterator.next()
									πF.SetLineno(1634)
									if πE = πg.CheckLocal(πF, µline_iterator, "line_iterator"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.GetAttr(πF, µline_iterator, ßnext, nil); πE != nil {
										continue
									}
									if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
										continue
									}
									if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp009}}}, πTemp003); πE != nil {
										continue
									}
									µfrom_line = πTemp001
									µto_line = πTemp004
									µfound_diff = πTemp009
									if πE = πg.CheckLocal(πF, µfrom_line, "from_line"); πE != nil {
										continue
									}
									if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
										continue
									}
									πTemp001 = πg.GetBool(µfrom_line != πTemp003).ToObject()
									if πTemp007, πE = πg.IsTrue(πF, πTemp001); πE != nil {
										continue
									}
									if πTemp007 {
										goto Label8
									}
									goto Label9
									// line 1635: if from_line is not None:
									πF.SetLineno(1635)
								Label8:
									// line 1636: fromlines.append((from_line,found_diff))
									πF.SetLineno(1636)
									πTemp002 = πF.MakeArgs(1)
									if πE = πg.CheckLocal(πF, µfrom_line, "from_line"); πE != nil {
										continue
									}
									if πE = πg.CheckLocal(πF, µfound_diff, "found_diff"); πE != nil {
										continue
									}
									πTemp001 = πg.NewTuple2(µfrom_line, µfound_diff).ToObject()
									πTemp002[0] = πTemp001
									if πE = πg.CheckLocal(πF, µfromlines, "fromlines"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.GetAttr(πF, µfromlines, ßappend, nil); πE != nil {
										continue
									}
									if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp002)
									goto Label9
								Label9:
									if πE = πg.CheckLocal(πF, µto_line, "to_line"); πE != nil {
										continue
									}
									if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
										continue
									}
									πTemp001 = πg.GetBool(µto_line != πTemp003).ToObject()
									if πTemp007, πE = πg.IsTrue(πF, πTemp001); πE != nil {
										continue
									}
									if πTemp007 {
										goto Label10
									}
									goto Label11
									// line 1637: if to_line is not None:
									πF.SetLineno(1637)
								Label10:
									// line 1638: tolines.append((to_line,found_diff))
									πF.SetLineno(1638)
									πTemp002 = πF.MakeArgs(1)
									if πE = πg.CheckLocal(πF, µto_line, "to_line"); πE != nil {
										continue
									}
									if πE = πg.CheckLocal(πF, µfound_diff, "found_diff"); πE != nil {
										continue
									}
									πTemp001 = πg.NewTuple2(µto_line, µfound_diff).ToObject()
									πTemp002[0] = πTemp001
									if πE = πg.CheckLocal(πF, µtolines, "tolines"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.GetAttr(πF, µtolines, ßappend, nil); πE != nil {
										continue
									}
									if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp002)
									goto Label11
								Label11:
									continue
								Label5:
									if πE != nil || πR != nil {
										continue
									}
								Label6:
									// line 1640: from_line, fromDiff = fromlines.pop(0)
									πF.SetLineno(1640)
									πTemp002 = πF.MakeArgs(1)
									πTemp002[0] = πg.NewInt(0).ToObject()
									if πE = πg.CheckLocal(πF, µfromlines, "fromlines"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.GetAttr(πF, µfromlines, ßpop, nil); πE != nil {
										continue
									}
									if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp002)
									if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp004}}}, πTemp003); πE != nil {
										continue
									}
									µfrom_line = πTemp001
									µfromDiff = πTemp004
									// line 1641: to_line, to_diff = tolines.pop(0)
									πF.SetLineno(1641)
									πTemp002 = πF.MakeArgs(1)
									πTemp002[0] = πg.NewInt(0).ToObject()
									if πE = πg.CheckLocal(πF, µtolines, "tolines"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.GetAttr(πF, µtolines, ßpop, nil); πE != nil {
										continue
									}
									if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp002)
									if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp004}}}, πTemp003); πE != nil {
										continue
									}
									µto_line = πTemp001
									µto_diff = πTemp004
									// line 1642: yield (from_line,to_line,fromDiff or to_diff)
									πF.SetLineno(1642)
									if πE = πg.CheckLocal(πF, µfrom_line, "from_line"); πE != nil {
										continue
									}
									if πE = πg.CheckLocal(πF, µto_line, "to_line"); πE != nil {
										continue
									}
									if πE = πg.CheckLocal(πF, µfromDiff, "fromDiff"); πE != nil {
										continue
									}
									πTemp003 = µfromDiff
									if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
										continue
									}
									if πTemp006 {
										goto Label12
									}
									if πE = πg.CheckLocal(πF, µto_diff, "to_diff"); πE != nil {
										continue
									}
									πTemp003 = µto_diff
								Label12:
									πTemp001 = πg.NewTuple3(µfrom_line, µto_line, πTemp003).ToObject()
									πF.PushCheckpoint(13)
									return πTemp001, nil
								Label13:
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
						µ_line_pair_iterator = πTemp005
						// line 1617: """Yields from/to lines of text with a change indication.
						πF.SetLineno(1617)
						if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, πg.NewStr("Yields from/to lines of text with a change indication.\n\n        This function is an iterator.  It itself pulls lines from the line\n        iterator.  Its difference from that iterator is that this function\n        always yields a pair of from/to text lines (with the change\n        indication).  If necessary it will collect single from/to lines\n        until it has a matching pair from/to pair to yield.\n\n        Note, this function is purposefully not defined at the module scope so\n        that data it needs from its parent function (within whose context it\n        is defined) does not need to be of module scope.\n        ").ToObject()); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µ_line_pair_iterator, "_line_pair_iterator"); πE != nil {
							continue
						}
						if πE = πg.SetAttr(πF, µ_line_pair_iterator, ß__doc__, πTemp006); πE != nil {
							continue
						}
						// line 1646: line_pair_iterator = _line_pair_iterator()
						πF.SetLineno(1646)
						if πE = πg.CheckLocal(πF, µ_line_pair_iterator, "_line_pair_iterator"); πE != nil {
							continue
						}
						if πTemp006, πE = µ_line_pair_iterator.Call(πF, nil, nil); πE != nil {
							continue
						}
						µline_pair_iterator = πTemp006
						if πE = πg.CheckLocal(πF, µcontext, "context"); πE != nil {
							continue
						}
						if πTemp007, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
							continue
						}
						πTemp006 = πg.GetBool(µcontext == πTemp007).ToObject()
						if πTemp008, πE = πg.IsTrue(πF, πTemp006); πE != nil {
							continue
						}
						if πTemp008 {
							goto Label1
						}
						goto Label2
						// line 1647: if context is None:
						πF.SetLineno(1647)
					Label1:
						// line 1648: while True:
						πF.SetLineno(1648)
						πF.PushCheckpoint(5)
						πTemp008 = false
					Label4:
						if πE != nil || πR != nil {
							continue
						}
						if πTemp008 {
							πF.PopCheckpoint()
							goto Label6
						}
						if πTemp006, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
							continue
						}
						if πTemp009, πE = πg.IsTrue(πF, πTemp006); πE != nil {
							continue
						}
						if πE != nil || !πTemp009 {
							continue
						}
						πF.PushCheckpoint(4)            
						// line 1649: yield line_pair_iterator.next()
						πF.SetLineno(1649)
						if πE = πg.CheckLocal(πF, µline_pair_iterator, "line_pair_iterator"); πE != nil {
							continue
						}
						if πTemp006, πE = πg.GetAttr(πF, µline_pair_iterator, ßnext, nil); πE != nil {
							continue
						}
						if πTemp007, πE = πTemp006.Call(πF, nil, nil); πE != nil {
							continue
						}
						πF.PushCheckpoint(7)
						return πTemp007, nil
					Label7:
						πTemp006 = πSent
						continue
					Label5:
						if πE != nil || πR != nil {
							continue
						}
					Label6:
						goto Label3
					Label2:
						// line 1653: context += 1
						πF.SetLineno(1653)
						if πE = πg.CheckLocal(πF, µcontext, "context"); πE != nil {
							continue
						}
						if πTemp006, πE = πg.IAdd(πF, µcontext, πg.NewInt(1).ToObject()); πE != nil {
							continue
						}
						µcontext = πTemp006
						// line 1654: lines_to_write = 0
						πF.SetLineno(1654)
						µlines_to_write = πg.NewInt(0).ToObject()
						// line 1655: while True:
						πF.SetLineno(1655)
						πF.PushCheckpoint(9)
						πTemp008 = false
					Label8:
						if πE != nil || πR != nil {
							continue
						}
						if πTemp008 {
							πF.PopCheckpoint()
							goto Label10
						}
						if πTemp006, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
							continue
						}
						if πTemp009, πE = πg.IsTrue(πF, πTemp006); πE != nil {
							continue
						}
						if πE != nil || !πTemp009 {
							continue
						}
						πF.PushCheckpoint(8)            
						// line 1659: index, contextLines = 0, [None]*(context)
						πF.SetLineno(1659)
						πTemp002 = make([]*πg.Object, 1)
						if πTemp011, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
							continue
						}
						πTemp002[0] = πTemp011
						πTemp011 = πg.NewList(πTemp002...).ToObject()
						if πE = πg.CheckLocal(πF, µcontext, "context"); πE != nil {
							continue
						}
						if πTemp010, πE = πg.Mul(πF, πTemp011, µcontext); πE != nil {
							continue
						}
						πTemp006 = πg.NewTuple2(πg.NewInt(0).ToObject(), πTemp010).ToObject()
						if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp010}, πg.TieTarget{Target: &πTemp011}}}, πTemp006); πE != nil {
							continue
						}
						µindex = πTemp010
						µcontextLines = πTemp011
						// line 1660: found_diff = False
						πF.SetLineno(1660)
						if πTemp006, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
							continue
						}
						µfound_diff = πTemp006
						// line 1661: while(found_diff is False):
						πF.SetLineno(1661)
						πF.PushCheckpoint(12)
						πTemp009 = false
					Label11:
						if πE != nil || πR != nil {
							continue
						}
						if πTemp009 {
							πF.PopCheckpoint()
							goto Label13
						}
						if πE = πg.CheckLocal(πF, µfound_diff, "found_diff"); πE != nil {
							continue
						}
						if πTemp010, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
							continue
						}
						πTemp006 = πg.GetBool(µfound_diff == πTemp010).ToObject()
						if πTemp012, πE = πg.IsTrue(πF, πTemp006); πE != nil {
							continue
						}
						if πE != nil || !πTemp012 {
							continue
						}
						πF.PushCheckpoint(11)            
						// line 1662: from_line, to_line, found_diff = line_pair_iterator.next()
						πF.SetLineno(1662)
						if πE = πg.CheckLocal(πF, µline_pair_iterator, "line_pair_iterator"); πE != nil {
							continue
						}
						if πTemp006, πE = πg.GetAttr(πF, µline_pair_iterator, ßnext, nil); πE != nil {
							continue
						}
						if πTemp010, πE = πTemp006.Call(πF, nil, nil); πE != nil {
							continue
						}
						if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp006}, πg.TieTarget{Target: &πTemp011}, πg.TieTarget{Target: &πTemp013}}}, πTemp010); πE != nil {
							continue
						}
						µfrom_line = πTemp006
						µto_line = πTemp011
						µfound_diff = πTemp013
						// line 1663: i = index % context
						πF.SetLineno(1663)
						if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µcontext, "context"); πE != nil {
							continue
						}
						if πTemp006, πE = πg.Mod(πF, µindex, µcontext); πE != nil {
							continue
						}
						µi = πTemp006
						// line 1664: contextLines[i] = (from_line, to_line, found_diff)
						πF.SetLineno(1664)
						if πE = πg.CheckLocal(πF, µfrom_line, "from_line"); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µto_line, "to_line"); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µfound_diff, "found_diff"); πE != nil {
							continue
						}
						πTemp006 = πg.NewTuple3(µfrom_line, µto_line, µfound_diff).ToObject()
						if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp010}, πTemp006); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µcontextLines, "contextLines"); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
							continue
						}
						πTemp011 = µi
						if πE = πg.SetItem(πF, µcontextLines, πTemp011, πTemp010); πE != nil {
							continue
						}
						// line 1665: index += 1
						πF.SetLineno(1665)
						if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
							continue
						}
						if πTemp006, πE = πg.IAdd(πF, µindex, πg.NewInt(1).ToObject()); πE != nil {
							continue
						}
						µindex = πTemp006
						continue
					Label12:
						if πE != nil || πR != nil {
							continue
						}
					Label13:
						if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µcontext, "context"); πE != nil {
							continue
						}
						if πTemp006, πE = πg.GT(πF, µindex, µcontext); πE != nil {
							continue
						}
						if πTemp009, πE = πg.IsTrue(πF, πTemp006); πE != nil {
							continue
						}
						if πTemp009 {
							goto Label14
						}
						goto Label15
						// line 1668: if index > context:
						πF.SetLineno(1668)
					Label14:
						// line 1669: yield None, None, None
						πF.SetLineno(1669)
						if πTemp010, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
							continue
						}
						if πTemp011, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
							continue
						}
						if πTemp013, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
							continue
						}
						πTemp006 = πg.NewTuple3(πTemp010, πTemp011, πTemp013).ToObject()
						πF.PushCheckpoint(17)
						return πTemp006, nil
					Label17:
						πTemp010 = πSent
						// line 1670: lines_to_write = context
						πF.SetLineno(1670)
						if πE = πg.CheckLocal(πF, µcontext, "context"); πE != nil {
							continue
						}
						µlines_to_write = µcontext
						goto Label16
					Label15:
						// line 1672: lines_to_write = index
						πF.SetLineno(1672)
						if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
							continue
						}
						µlines_to_write = µindex
						// line 1673: index = 0
						πF.SetLineno(1673)
						µindex = πg.NewInt(0).ToObject()
						goto Label16
					Label16:
						// line 1674: while(lines_to_write):
						πF.SetLineno(1674)
						πF.PushCheckpoint(19)
						πTemp009 = false
					Label18:
						if πE != nil || πR != nil {
							continue
						}
						if πTemp009 {
							πF.PopCheckpoint()
							goto Label20
						}
						if πE = πg.CheckLocal(πF, µlines_to_write, "lines_to_write"); πE != nil {
							continue
						}
						if πTemp012, πE = πg.IsTrue(πF, µlines_to_write); πE != nil {
							continue
						}
						if πE != nil || !πTemp012 {
							continue
						}
						πF.PushCheckpoint(18)            
						// line 1675: i = index % context
						πF.SetLineno(1675)
						if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µcontext, "context"); πE != nil {
							continue
						}
						if πTemp010, πE = πg.Mod(πF, µindex, µcontext); πE != nil {
							continue
						}
						µi = πTemp010
						// line 1676: index += 1
						πF.SetLineno(1676)
						if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
							continue
						}
						if πTemp010, πE = πg.IAdd(πF, µindex, πg.NewInt(1).ToObject()); πE != nil {
							continue
						}
						µindex = πTemp010
						// line 1677: yield contextLines[i]
						πF.SetLineno(1677)
						if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
							continue
						}
						πTemp010 = µi
						if πE = πg.CheckLocal(πF, µcontextLines, "contextLines"); πE != nil {
							continue
						}
						if πTemp011, πE = πg.GetItem(πF, µcontextLines, πTemp010); πE != nil {
							continue
						}
						πF.PushCheckpoint(21)
						return πTemp011, nil
					Label21:
						πTemp010 = πSent
						// line 1678: lines_to_write -= 1
						πF.SetLineno(1678)
						if πE = πg.CheckLocal(πF, µlines_to_write, "lines_to_write"); πE != nil {
							continue
						}
						if πTemp010, πE = πg.ISub(πF, µlines_to_write, πg.NewInt(1).ToObject()); πE != nil {
							continue
						}
						µlines_to_write = πTemp010
						continue
					Label19:
						if πE != nil || πR != nil {
							continue
						}
					Label20:
						// line 1680: lines_to_write = context-1
						πF.SetLineno(1680)
						if πE = πg.CheckLocal(πF, µcontext, "context"); πE != nil {
							continue
						}
						if πTemp010, πE = πg.Sub(πF, µcontext, πg.NewInt(1).ToObject()); πE != nil {
							continue
						}
						µlines_to_write = πTemp010
						// line 1681: while(lines_to_write):
						πF.SetLineno(1681)
						πF.PushCheckpoint(23)
						πTemp009 = false
					Label22:
						if πE != nil || πR != nil {
							continue
						}
						if πTemp009 {
							πF.PopCheckpoint()
							goto Label24
						}
						if πE = πg.CheckLocal(πF, µlines_to_write, "lines_to_write"); πE != nil {
							continue
						}
						if πTemp012, πE = πg.IsTrue(πF, µlines_to_write); πE != nil {
							continue
						}
						if πE != nil || !πTemp012 {
							continue
						}
						πF.PushCheckpoint(22)            
						// line 1682: from_line, to_line, found_diff = line_pair_iterator.next()
						πF.SetLineno(1682)
						if πE = πg.CheckLocal(πF, µline_pair_iterator, "line_pair_iterator"); πE != nil {
							continue
						}
						if πTemp010, πE = πg.GetAttr(πF, µline_pair_iterator, ßnext, nil); πE != nil {
							continue
						}
						if πTemp013, πE = πTemp010.Call(πF, nil, nil); πE != nil {
							continue
						}
						if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp010}, πg.TieTarget{Target: &πTemp014}, πg.TieTarget{Target: &πTemp015}}}, πTemp013); πE != nil {
							continue
						}
						µfrom_line = πTemp010
						µto_line = πTemp014
						µfound_diff = πTemp015
						if πE = πg.CheckLocal(πF, µfound_diff, "found_diff"); πE != nil {
							continue
						}
						if πTemp012, πE = πg.IsTrue(πF, µfound_diff); πE != nil {
							continue
						}
						if πTemp012 {
							goto Label25
						}
						goto Label26
						// line 1684: if found_diff:
						πF.SetLineno(1684)
					Label25:
						// line 1685: lines_to_write = context-1
						πF.SetLineno(1685)
						if πE = πg.CheckLocal(πF, µcontext, "context"); πE != nil {
							continue
						}
						if πTemp010, πE = πg.Sub(πF, µcontext, πg.NewInt(1).ToObject()); πE != nil {
							continue
						}
						µlines_to_write = πTemp010
						goto Label27
					Label26:
						// line 1687: lines_to_write -= 1
						πF.SetLineno(1687)
						if πE = πg.CheckLocal(πF, µlines_to_write, "lines_to_write"); πE != nil {
							continue
						}
						if πTemp010, πE = πg.ISub(πF, µlines_to_write, πg.NewInt(1).ToObject()); πE != nil {
							continue
						}
						µlines_to_write = πTemp010
						goto Label27
					Label27:
						// line 1688: yield from_line, to_line, found_diff
						πF.SetLineno(1688)
						if πE = πg.CheckLocal(πF, µfrom_line, "from_line"); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µto_line, "to_line"); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µfound_diff, "found_diff"); πE != nil {
							continue
						}
						πTemp010 = πg.NewTuple3(µfrom_line, µto_line, µfound_diff).ToObject()
						πF.PushCheckpoint(28)
						return πTemp010, nil
					Label28:
						πTemp013 = πSent
						continue
					Label23:
						if πE != nil || πR != nil {
							continue
						}
					Label24:
						continue
					Label9:
						if πE != nil || πR != nil {
							continue
						}
					Label10:
						goto Label3
					Label3:
					}
					return nil, πE
				}).ToObject(), nil
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_mdiff.ToObject(), πTemp015); πE != nil {
				continue
			}
			// line 1429: r"""Returns generator yielding marked up from/to side by side differences.
			πF.SetLineno(1429)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp016}, πg.NewStr("Returns generator yielding marked up from/to side by side differences.\n\n    Arguments:\n    fromlines -- list of text lines to compared to tolines\n    tolines -- list of text lines to be compared to fromlines\n    context -- number of context lines to display on each side of difference,\n               if None, all from/to text lines will be generated.\n    linejunk -- passed on to ndiff (see ndiff documentation)\n    charjunk -- passed on to ndiff (see ndiff documentation)\n\n    This function returns an iterator which returns a tuple:\n    (from line tuple, to line tuple, boolean flag)\n\n    from/to line tuple -- (line num, line text)\n        line num -- integer or None (to indicate a context separation)\n        line text -- original line text with following markers inserted:\n            '\\0+' -- marks start of added text\n            '\\0-' -- marks start of deleted text\n            '\\0^' -- marks start of changed text\n            '\\1' -- marks end of added/deleted/changed text\n\n    boolean flag -- None indicates context separation, True indicates\n        either \"from\" or \"to\" line contains a change, otherwise False.\n\n    This function/iterator was originally developed to generate side by side\n    file difference for making HTML pages (see HtmlDiff class for example\n    usage).\n\n    Note, this function utilizes the ndiff function to generate the side by\n    side difference markup.  Optional ndiff arguments may be passed to this\n    function and they in turn will be passed to ndiff.\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp017, πE = πg.ResolveGlobal(πF, ß_mdiff); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp017, ß__doc__, πTemp016); πE != nil {
				continue
			}
			// line 1691: _file_template = """
			πF.SetLineno(1691)
			if πE = πF.Globals().SetItem(πF, ß_file_template.ToObject(), πg.NewStr("\n<!DOCTYPE html PUBLIC \"-//W3C//DTD XHTML 1.0 Transitional//EN\"\n          \"http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd\">\n\n<html>\n\n<head>\n    <meta http-equiv=\"Content-Type\"\n          content=\"text/html; charset=ISO-8859-1\" />\n    <title></title>\n    <style type=\"text/css\">%(styles)s\n    </style>\n</head>\n\n<body>\n    %(table)s%(legend)s\n</body>\n\n</html>").ToObject()); πE != nil {
				continue
			}
			// line 1711: _styles = """
			πF.SetLineno(1711)
			if πE = πF.Globals().SetItem(πF, ß_styles.ToObject(), πg.NewStr("\n        table.diff {font-family:Courier; border:medium;}\n        .diff_header {background-color:#e0e0e0}\n        td.diff_header {text-align:right}\n        .diff_next {background-color:#c0c0c0}\n        .diff_add {background-color:#aaffaa}\n        .diff_chg {background-color:#ffff77}\n        .diff_sub {background-color:#ffaaaa}").ToObject()); πE != nil {
				continue
			}
			// line 1720: _table_template = """
			πF.SetLineno(1720)
			if πE = πF.Globals().SetItem(πF, ß_table_template.ToObject(), πg.NewStr("\n    <table class=\"diff\" id=\"difflib_chg_%(prefix)s_top\"\n           cellspacing=\"0\" cellpadding=\"0\" rules=\"groups\" >\n        <colgroup></colgroup> <colgroup></colgroup> <colgroup></colgroup>\n        <colgroup></colgroup> <colgroup></colgroup> <colgroup></colgroup>\n        %(header_row)s\n        <tbody>\n%(data_rows)s        </tbody>\n    </table>").ToObject()); πE != nil {
				continue
			}
			// line 1730: _legend = """
			πF.SetLineno(1730)
			if πE = πF.Globals().SetItem(πF, ß_legend.ToObject(), πg.NewStr("\n    <table class=\"diff\" summary=\"Legends\">\n        <tr> <th colspan=\"2\"> Legends </th> </tr>\n        <tr> <td> <table border=\"\" summary=\"Colors\">\n                      <tr><th> Colors </th> </tr>\n                      <tr><td class=\"diff_add\">&nbsp;Added&nbsp;</td></tr>\n                      <tr><td class=\"diff_chg\">Changed</td> </tr>\n                      <tr><td class=\"diff_sub\">Deleted</td> </tr>\n                  </table></td>\n             <td> <table border=\"\" summary=\"Links\">\n                      <tr><th colspan=\"2\"> Links </th> </tr>\n                      <tr><td>(f)irst change</td> </tr>\n                      <tr><td>(n)ext change</td> </tr>\n                      <tr><td>(t)op</td> </tr>\n                  </table></td> </tr>\n    </table>").ToObject()); πE != nil {
				continue
			}
			// line 1747: class HtmlDiff(object):
			πF.SetLineno(1747)
			πTemp001 = make([]*πg.Object, 1)
			if πTemp018, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp001[0] = πTemp018
			πTemp005 = πg.NewDict()
			if πTemp016, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp005.SetItem(πF, ß__module__.ToObject(), πTemp016); πE != nil {
				continue
			}
			_, πE = πg.NewCode("HtmlDiff", "/usr/lib/python2.7/difflib.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 1748: """For producing HTML side by side comparison with change highlights.
					πF.SetLineno(1748)
					// line 1748: """For producing HTML side by side comparison with change highlights.
					πF.SetLineno(1748)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("For producing HTML side by side comparison with change highlights.\n\n    This class can be used to create an HTML table (or a complete HTML file\n    containing the table) showing a side by side, line by line comparison\n    of text with inter-line and intra-line change highlights.  The table can\n    be generated in either full or contextual difference mode.\n\n    The following methods are provided for HTML generation:\n\n    make_table -- generates HTML for a single side by side table\n    make_file -- generates complete HTML file with a single side by side table\n\n    See tools/scripts/diff.py for an example usage of this class.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 1763: _file_template = _file_template
					πF.SetLineno(1763)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ß_file_template); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ß_file_template.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 1764: _styles = _styles
					πF.SetLineno(1764)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ß_styles); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ß_styles.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 1765: _table_template = _table_template
					πF.SetLineno(1765)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ß_table_template); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ß_table_template.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 1766: _legend = _legend
					πF.SetLineno(1766)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ß_legend); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ß_legend.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 1767: _default_prefix = 0
					πF.SetLineno(1767)
					if πE = πClass.SetItem(πF, ß_default_prefix.ToObject(), πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					// line 1769: def __init__(self,tabsize=8,wrapcolumn=None,linejunk=None,
					πF.SetLineno(1769)
					πTemp002 = make([]πg.Param, 5)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "tabsize", Def: πg.NewInt(8).ToObject()}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "wrapcolumn", Def: πTemp003}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[3] = πg.Param{Name: "linejunk", Def: πTemp003}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßIS_CHARACTER_JUNK); πE != nil {
						continue
					}
					πTemp002[4] = πg.Param{Name: "charjunk", Def: πTemp003}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/difflib.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µtabsize *πg.Object = πArgs[1]; _ = µtabsize
						var µwrapcolumn *πg.Object = πArgs[2]; _ = µwrapcolumn
						var µlinejunk *πg.Object = πArgs[3]; _ = µlinejunk
						var µcharjunk *πg.Object = πArgs[4]; _ = µcharjunk
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 1771: """HtmlDiff instance initializer
							πF.SetLineno(1771)
							// line 1781: self._tabsize = tabsize
							πF.SetLineno(1781)
							if πE = πg.CheckLocal(πF, µtabsize, "tabsize"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µtabsize); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_tabsize, πTemp001); πE != nil {
								continue
							}
							// line 1782: self._wrapcolumn = wrapcolumn
							πF.SetLineno(1782)
							if πE = πg.CheckLocal(πF, µwrapcolumn, "wrapcolumn"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µwrapcolumn); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_wrapcolumn, πTemp001); πE != nil {
								continue
							}
							// line 1783: self._linejunk = linejunk
							πF.SetLineno(1783)
							if πE = πg.CheckLocal(πF, µlinejunk, "linejunk"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µlinejunk); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_linejunk, πTemp001); πE != nil {
								continue
							}
							// line 1784: self._charjunk = charjunk
							πF.SetLineno(1784)
							if πE = πg.CheckLocal(πF, µcharjunk, "charjunk"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µcharjunk); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_charjunk, πTemp001); πE != nil {
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
					// line 1771: """HtmlDiff instance initializer
					πF.SetLineno(1771)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewStr("HtmlDiff instance initializer\n\n        Arguments:\n        tabsize -- tab stop spacing, defaults to 8.\n        wrapcolumn -- column number where lines are broken and wrapped,\n            defaults to None where lines are not wrapped.\n        linejunk,charjunk -- keyword arguments passed into ndiff() (used to by\n            HtmlDiff() to generate the side by side HTML differences).  See\n            ndiff() documentation for argument default values and descriptions.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ß__init__); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp004, ß__doc__, πTemp003); πE != nil {
						continue
					}
					// line 1786: def make_file(self,fromlines,tolines,fromdesc='',todesc='',context=False,
					πF.SetLineno(1786)
					πTemp002 = make([]πg.Param, 7)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "fromlines", Def: nil}
					πTemp002[2] = πg.Param{Name: "tolines", Def: nil}
					πTemp002[3] = πg.Param{Name: "fromdesc", Def: ß.ToObject()}
					πTemp002[4] = πg.Param{Name: "todesc", Def: ß.ToObject()}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßFalse); πE != nil {
						continue
					}
					πTemp002[5] = πg.Param{Name: "context", Def: πTemp004}
					πTemp002[6] = πg.Param{Name: "numlines", Def: πg.NewInt(5).ToObject()}
					πTemp003 = πg.NewFunction(πg.NewCode("make_file", "/usr/lib/python2.7/difflib.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µfromlines *πg.Object = πArgs[1]; _ = µfromlines
						var µtolines *πg.Object = πArgs[2]; _ = µtolines
						var µfromdesc *πg.Object = πArgs[3]; _ = µfromdesc
						var µtodesc *πg.Object = πArgs[4]; _ = µtodesc
						var µcontext *πg.Object = πArgs[5]; _ = µcontext
						var µnumlines *πg.Object = πArgs[6]; _ = µnumlines
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
						var πTemp006 πg.KWArgs
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
							// line 1788: """Returns HTML file of side by side comparison with change highlights
							πF.SetLineno(1788)
							// line 1804: return self._file_template % dict(
							πF.SetLineno(1804)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_file_template, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_styles, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß_legend, nil); πE != nil {
								continue
							}
							πTemp005 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µfromlines, "fromlines"); πE != nil {
								continue
							}
							πTemp005[0] = µfromlines
							if πE = πg.CheckLocal(πF, µtolines, "tolines"); πE != nil {
								continue
							}
							πTemp005[1] = µtolines
							if πE = πg.CheckLocal(πF, µfromdesc, "fromdesc"); πE != nil {
								continue
							}
							πTemp005[2] = µfromdesc
							if πE = πg.CheckLocal(πF, µtodesc, "todesc"); πE != nil {
								continue
							}
							πTemp005[3] = µtodesc
							if πE = πg.CheckLocal(πF, µcontext, "context"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnumlines, "numlines"); πE != nil {
								continue
							}
							πTemp006 = πg.KWArgs{
								{"context", µcontext},
								{"numlines", µnumlines},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßmake_table, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp007.Call(πF, πTemp005, πTemp006); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp006 = πg.KWArgs{
								{"styles", πTemp003},
								{"legend", πTemp004},
								{"table", πTemp008},
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, πTemp006); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πTemp002, πTemp004); πE != nil {
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
					if πE = πClass.SetItem(πF, ßmake_file.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 1788: """Returns HTML file of side by side comparison with change highlights
					πF.SetLineno(1788)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πg.NewStr("Returns HTML file of side by side comparison with change highlights\n\n        Arguments:\n        fromlines -- list of \"from\" lines\n        tolines -- list of \"to\" lines\n        fromdesc -- \"from\" file column header string\n        todesc -- \"to\" file column header string\n        context -- set to True for contextual differences (defaults to False\n            which shows full differences).\n        numlines -- number of context lines.  When context is set True,\n            controls number of lines displayed before and after the change.\n            When context is False, controls the number of lines to place\n            the \"next\" link anchors before the next change (so click of\n            \"next\" link jumps to just before the change).\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßmake_file); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp005, ß__doc__, πTemp004); πE != nil {
						continue
					}
					// line 1810: def _tab_newline_replace(self,fromlines,tolines):
					πF.SetLineno(1810)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "fromlines", Def: nil}
					πTemp002[2] = πg.Param{Name: "tolines", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("_tab_newline_replace", "/usr/lib/python2.7/difflib.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µfromlines *πg.Object = πArgs[1]; _ = µfromlines
						var µtolines *πg.Object = πArgs[2]; _ = µtolines
						var µexpand_tabs *πg.Object = πg.UnboundLocal; _ = µexpand_tabs
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 1811: """Returns from/to line lists with tabs expanded and newlines removed.
							πF.SetLineno(1811)
							// line 1820: def expand_tabs(line):
							πF.SetLineno(1820)
							πTemp002 = make([]πg.Param, 1)
							πTemp002[0] = πg.Param{Name: "line", Def: nil}
							πTemp001 = πg.NewFunction(πg.NewCode("expand_tabs", "/usr/lib/python2.7/difflib.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µline *πg.Object = πArgs[0]; _ = µline
								var πTemp001 []*πg.Object
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
									// line 1822: line = line.replace(' ','\0')
									πF.SetLineno(1822)
									πTemp001 = πF.MakeArgs(2)
									πTemp001[0] = πg.NewStr(" ").ToObject()
									πTemp001[1] = πg.NewStr("\x00").ToObject()
									if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, µline, ßreplace, nil); πE != nil {
										continue
									}
									if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp001)
									µline = πTemp003
									// line 1824: line = line.expandtabs(self._tabsize)
									πF.SetLineno(1824)
									πTemp001 = πF.MakeArgs(1)
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, µself, ß_tabsize, nil); πE != nil {
										continue
									}
									πTemp001[0] = πTemp002
									if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, µline, ßexpandtabs, nil); πE != nil {
										continue
									}
									if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp001)
									µline = πTemp003
									// line 1827: line = line.replace(' ','\t')
									πF.SetLineno(1827)
									πTemp001 = πF.MakeArgs(2)
									πTemp001[0] = πg.NewStr(" ").ToObject()
									πTemp001[1] = πg.NewStr("\t").ToObject()
									if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, µline, ßreplace, nil); πE != nil {
										continue
									}
									if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp001)
									µline = πTemp003
									// line 1828: return line.replace('\0',' ').rstrip('\n')
									πF.SetLineno(1828)
									πTemp001 = πF.MakeArgs(1)
									πTemp001[0] = πg.NewStr("\n").ToObject()
									πTemp004 = πF.MakeArgs(2)
									πTemp004[0] = πg.NewStr("\x00").ToObject()
									πTemp004[1] = πg.NewStr(" ").ToObject()
									if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, µline, ßreplace, nil); πE != nil {
										continue
									}
									if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp004)
									if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßrstrip, nil); πE != nil {
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
							µexpand_tabs = πTemp001
							// line 1829: fromlines = [expand_tabs(line) for line in fromlines]
							πF.SetLineno(1829)
							πTemp002 = make([]πg.Param, 0)
							πTemp004 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/difflib.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µline *πg.Object = πg.UnboundLocal; _ = µline
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
										if πE = πg.CheckLocal(πF, µfromlines, "fromlines"); πE != nil {
											continue
										}
										if πTemp001, πE = πg.Iter(πF, µfromlines); πE != nil {
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
											µline = πTemp004
										}
										if πE != nil || !πTemp003 {
											continue
										}
										πF.PushCheckpoint(1)            
										// line 1829: fromlines = [expand_tabs(line) for line in fromlines]
										πF.SetLineno(1829)
										πTemp005 = πF.MakeArgs(1)
										if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
											continue
										}
										πTemp005[0] = µline
										if πE = πg.CheckLocal(πF, µexpand_tabs, "expand_tabs"); πE != nil {
											continue
										}
										if πTemp004, πE = µexpand_tabs.Call(πF, πTemp005, nil); πE != nil {
											continue
										}
										πF.FreeArgs(πTemp005)
										πF.PushCheckpoint(4)
										return πTemp004, nil
									Label4:
										πTemp006 = πSent
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
							µfromlines = πTemp003
							// line 1830: tolines = [expand_tabs(line) for line in tolines]
							πF.SetLineno(1830)
							πTemp002 = make([]πg.Param, 0)
							πTemp005 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/difflib.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µline *πg.Object = πg.UnboundLocal; _ = µline
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
										if πE = πg.CheckLocal(πF, µtolines, "tolines"); πE != nil {
											continue
										}
										if πTemp001, πE = πg.Iter(πF, µtolines); πE != nil {
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
											µline = πTemp004
										}
										if πE != nil || !πTemp003 {
											continue
										}
										πF.PushCheckpoint(1)            
										// line 1830: tolines = [expand_tabs(line) for line in tolines]
										πF.SetLineno(1830)
										πTemp005 = πF.MakeArgs(1)
										if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
											continue
										}
										πTemp005[0] = µline
										if πE = πg.CheckLocal(πF, µexpand_tabs, "expand_tabs"); πE != nil {
											continue
										}
										if πTemp004, πE = µexpand_tabs.Call(πF, πTemp005, nil); πE != nil {
											continue
										}
										πF.FreeArgs(πTemp005)
										πF.PushCheckpoint(4)
										return πTemp004, nil
									Label4:
										πTemp006 = πSent
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
							µtolines = πTemp003
							// line 1831: return fromlines,tolines
							πF.SetLineno(1831)
							if πE = πg.CheckLocal(πF, µfromlines, "fromlines"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtolines, "tolines"); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple2(µfromlines, µtolines).ToObject()
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
					if πE = πClass.SetItem(πF, ß_tab_newline_replace.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 1811: """Returns from/to line lists with tabs expanded and newlines removed.
					πF.SetLineno(1811)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πg.NewStr("Returns from/to line lists with tabs expanded and newlines removed.\n\n        Instead of tab characters being replaced by the number of spaces\n        needed to fill in to the next tab stop, this function will fill\n        the space with tab characters.  This is done so that the difference\n        algorithms can identify changes in a file when tabs are replaced by\n        spaces and vice versa.  At the end of the HTML generation, the tab\n        characters will be replaced with a nonbreakable space.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ß_tab_newline_replace); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp006, ß__doc__, πTemp005); πE != nil {
						continue
					}
					// line 1833: def _split_line(self,data_list,line_num,text):
					πF.SetLineno(1833)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "data_list", Def: nil}
					πTemp002[2] = πg.Param{Name: "line_num", Def: nil}
					πTemp002[3] = πg.Param{Name: "text", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("_split_line", "/usr/lib/python2.7/difflib.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µdata_list *πg.Object = πArgs[1]; _ = µdata_list
						var µline_num *πg.Object = πArgs[2]; _ = µline_num
						var µtext *πg.Object = πArgs[3]; _ = µtext
						var µsize *πg.Object = πg.UnboundLocal; _ = µsize
						var µmax *πg.Object = πg.UnboundLocal; _ = µmax
						var µi *πg.Object = πg.UnboundLocal; _ = µi
						var µn *πg.Object = πg.UnboundLocal; _ = µn
						var µmark *πg.Object = πg.UnboundLocal; _ = µmark
						var µline1 *πg.Object = πg.UnboundLocal; _ = µline1
						var µline2 *πg.Object = πg.UnboundLocal; _ = µline2
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
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 bool
						_ = πTemp009
						var πTemp010 bool
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
							// line 1834: """Builds list of text lines by splitting text lines at wrap point
							πF.SetLineno(1834)
							if πE = πg.CheckLocal(πF, µline_num, "line_num"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, µline_num); πE != nil {
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
							// line 1843: if not line_num:
							πF.SetLineno(1843)
						Label1:
							// line 1844: data_list.append((line_num,text))
							πF.SetLineno(1844)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µline_num, "line_num"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple2(µline_num, µtext).ToObject()
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µdata_list, "data_list"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µdata_list, ßappend, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 1845: return
							πF.SetLineno(1845)
							πR = πg.None
							continue
							goto Label2
						Label2:
							// line 1848: size = len(text)
							πF.SetLineno(1848)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							πTemp003[0] = µtext
							if πTemp001, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µsize = πTemp004
							// line 1849: max = self._wrapcolumn
							πF.SetLineno(1849)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_wrapcolumn, nil); πE != nil {
								continue
							}
							µmax = πTemp001
							if πE = πg.CheckLocal(πF, µsize, "size"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µmax, "max"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.LE(πF, µsize, µmax); πE != nil {
								continue
							}
							πTemp001 = πTemp004
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label3
							}
							if πE = πg.CheckLocal(πF, µsize, "size"); πE != nil {
								continue
							}
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewStr("\x00").ToObject()
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µtext, ßcount, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp007.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp006, πE = πg.Mul(πF, πTemp008, πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Sub(πF, µsize, πTemp006); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µmax, "max"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.LE(πF, πTemp005, µmax); πE != nil {
								continue
							}
							πTemp001 = πTemp004
						Label3:
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label4
							}
							goto Label5
							// line 1850: if (size <= max) or ((size -(text.count('\0')*3)) <= max):
							πF.SetLineno(1850)
						Label4:
							// line 1851: data_list.append((line_num,text))
							πF.SetLineno(1851)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µline_num, "line_num"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple2(µline_num, µtext).ToObject()
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µdata_list, "data_list"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µdata_list, ßappend, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 1852: return
							πF.SetLineno(1852)
							πR = πg.None
							continue
							goto Label5
						Label5:
							// line 1856: i = 0
							πF.SetLineno(1856)
							µi = πg.NewInt(0).ToObject()
							// line 1857: n = 0
							πF.SetLineno(1857)
							µn = πg.NewInt(0).ToObject()
							// line 1858: mark = ''
							πF.SetLineno(1858)
							µmark = ß.ToObject()
							// line 1859: while n < max and i < size:
							πF.SetLineno(1859)
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
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µmax, "max"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.LT(πF, µn, µmax); πE != nil {
								continue
							}
							πTemp001 = πTemp004
							if πTemp010, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp010 {
								goto Label9
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µsize, "size"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.LT(πF, µi, µsize); πE != nil {
								continue
							}
							πTemp001 = πTemp004
						Label9:
							if πTemp009, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πE != nil || !πTemp009 {
								continue
							}
							πF.PushCheckpoint(6)            
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp004 = µi
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µtext, πTemp004); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, πTemp005, πg.NewStr("\x00").ToObject()); πE != nil {
								continue
							}
							if πTemp009, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp009 {
								goto Label10
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp004 = µi
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µtext, πTemp004); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, πTemp005, πg.NewStr("\x01").ToObject()); πE != nil {
								continue
							}
							if πTemp009, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp009 {
								goto Label11
							}
							goto Label12
							// line 1860: if text[i] == '\0':
							πF.SetLineno(1860)
						Label10:
							// line 1861: i += 1
							πF.SetLineno(1861)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.IAdd(πF, µi, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µi = πTemp001
							// line 1862: mark = text[i]
							πF.SetLineno(1862)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp001 = µi
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µtext, πTemp001); πE != nil {
								continue
							}
							µmark = πTemp004
							// line 1863: i += 1
							πF.SetLineno(1863)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.IAdd(πF, µi, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µi = πTemp001
							goto Label13
							// line 1864: elif text[i] == '\1':
							πF.SetLineno(1864)
						Label11:
							// line 1865: i += 1
							πF.SetLineno(1865)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.IAdd(πF, µi, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µi = πTemp001
							// line 1866: mark = ''
							πF.SetLineno(1866)
							µmark = ß.ToObject()
							goto Label13
						Label12:
							// line 1868: i += 1
							πF.SetLineno(1868)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.IAdd(πF, µi, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µi = πTemp001
							// line 1869: n += 1
							πF.SetLineno(1869)
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.IAdd(πF, µn, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µn = πTemp001
							goto Label13
						Label13:
							continue
						Label7:
							if πE != nil || πR != nil {
								continue
							}
						Label8:
							// line 1872: line1 = text[:i]
							πF.SetLineno(1872)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.None, µi, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µtext, πTemp001); πE != nil {
								continue
							}
							µline1 = πTemp004
							// line 1873: line2 = text[i:]
							πF.SetLineno(1873)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{µi, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µtext, πTemp001); πE != nil {
								continue
							}
							µline2 = πTemp004
							if πE = πg.CheckLocal(πF, µmark, "mark"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, µmark); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label14
							}
							goto Label15
							// line 1878: if mark:
							πF.SetLineno(1878)
						Label14:
							// line 1879: line1 = line1 + '\1'
							πF.SetLineno(1879)
							if πE = πg.CheckLocal(πF, µline1, "line1"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, µline1, πg.NewStr("\x01").ToObject()); πE != nil {
								continue
							}
							µline1 = πTemp001
							// line 1880: line2 = '\0' + mark + line2
							πF.SetLineno(1880)
							if πE = πg.CheckLocal(πF, µmark, "mark"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Add(πF, πg.NewStr("\x00").ToObject(), µmark); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µline2, "line2"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp004, µline2); πE != nil {
								continue
							}
							µline2 = πTemp001
							goto Label15
						Label15:
							// line 1883: data_list.append((line_num,line1))
							πF.SetLineno(1883)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µline_num, "line_num"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µline1, "line1"); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple2(µline_num, µline1).ToObject()
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µdata_list, "data_list"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µdata_list, ßappend, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 1886: self._split_line(data_list,'>',line2)
							πF.SetLineno(1886)
							πTemp003 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µdata_list, "data_list"); πE != nil {
								continue
							}
							πTemp003[0] = µdata_list
							πTemp003[1] = πg.NewStr(">").ToObject()
							if πE = πg.CheckLocal(πF, µline2, "line2"); πE != nil {
								continue
							}
							πTemp003[2] = µline2
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_split_line, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß_split_line.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 1834: """Builds list of text lines by splitting text lines at wrap point
					πF.SetLineno(1834)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, πg.NewStr("Builds list of text lines by splitting text lines at wrap point\n\n        This function will determine if the input text line needs to be\n        wrapped (split) into separate lines.  If so, the first wrap point\n        will be determined and the first line appended to the output\n        text line list.  This function is used recursively to handle\n        the second part of the split line to further split it.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ß_split_line); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp007, ß__doc__, πTemp006); πE != nil {
						continue
					}
					// line 1888: def _line_wrapper(self,diffs):
					πF.SetLineno(1888)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "diffs", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("_line_wrapper", "/usr/lib/python2.7/difflib.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µdiffs *πg.Object = πArgs[1]; _ = µdiffs
						var µfromdata *πg.Object = πg.UnboundLocal; _ = µfromdata
						var µtodata *πg.Object = πg.UnboundLocal; _ = µtodata
						var µflag *πg.Object = πg.UnboundLocal; _ = µflag
						var µfromline *πg.Object = πg.UnboundLocal; _ = µfromline
						var µfromtext *πg.Object = πg.UnboundLocal; _ = µfromtext
						var µtoline *πg.Object = πg.UnboundLocal; _ = µtoline
						var µtotext *πg.Object = πg.UnboundLocal; _ = µtotext
						var µfromlist *πg.Object = πg.UnboundLocal; _ = µfromlist
						var µtolist *πg.Object = πg.UnboundLocal; _ = µtolist
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
						var πTemp011 bool
						_ = πTemp011
						var πTemp012 bool
						_ = πTemp012
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
							for ; πF.State() >= 0; πF.PopCheckpoint() {
								switch πF.State() {
								case 0:
								case 1: goto Label1
								case 2: goto Label2
								case 6: goto Label6
								case 7: goto Label7
								case 8: goto Label8
								case 17: goto Label17
								default: panic("unexpected function state")
								}
								// line 1889: """Returns iterator that splits (wraps) mdiff text lines"""
								πF.SetLineno(1889)
								if πE = πg.CheckLocal(πF, µdiffs, "diffs"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.Iter(πF, µdiffs); πE != nil {
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
									µfromdata = πTemp005
									µtodata = πTemp006
									µflag = πTemp007
								}
								if πE != nil || !πTemp003 {
									continue
								}
								πF.PushCheckpoint(1)            
								if πE = πg.CheckLocal(πF, µflag, "flag"); πE != nil {
									continue
								}
								if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
									continue
								}
								πTemp004 = πg.GetBool(µflag == πTemp005).ToObject()
								if πTemp003, πE = πg.IsTrue(πF, πTemp004); πE != nil {
									continue
								}
								if πTemp003 {
									goto Label4
								}
								goto Label5
								// line 1894: if flag is None:
								πF.SetLineno(1894)
							Label4:
								// line 1895: yield fromdata,todata,flag
								πF.SetLineno(1895)
								if πE = πg.CheckLocal(πF, µfromdata, "fromdata"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µtodata, "todata"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µflag, "flag"); πE != nil {
									continue
								}
								πTemp004 = πg.NewTuple3(µfromdata, µtodata, µflag).ToObject()
								πF.PushCheckpoint(6)
								return πTemp004, nil
							Label6:
								πTemp005 = πSent
								// line 1896: continue
								πF.SetLineno(1896)
								continue
								goto Label5
							Label5:
								// line 1897: (fromline,fromtext),(toline,totext) = fromdata,todata
								πF.SetLineno(1897)
								if πE = πg.CheckLocal(πF, µfromdata, "fromdata"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µtodata, "todata"); πE != nil {
									continue
								}
								πTemp005 = πg.NewTuple2(µfromdata, µtodata).ToObject()
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp006}, πg.TieTarget{Target: &πTemp007}}}, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp008}, πg.TieTarget{Target: &πTemp009}}}}}, πTemp005); πE != nil {
									continue
								}
								µfromline = πTemp006
								µfromtext = πTemp007
								µtoline = πTemp008
								µtotext = πTemp009
								// line 1900: fromlist,tolist = [],[]
								πF.SetLineno(1900)
								πTemp010 = make([]*πg.Object, 0)
								πTemp006 = πg.NewList(πTemp010...).ToObject()
								πTemp010 = make([]*πg.Object, 0)
								πTemp007 = πg.NewList(πTemp010...).ToObject()
								πTemp005 = πg.NewTuple2(πTemp006, πTemp007).ToObject()
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp006}, πg.TieTarget{Target: &πTemp007}}}, πTemp005); πE != nil {
									continue
								}
								µfromlist = πTemp006
								µtolist = πTemp007
								// line 1901: self._split_line(fromlist,fromline,fromtext)
								πF.SetLineno(1901)
								πTemp010 = πF.MakeArgs(3)
								if πE = πg.CheckLocal(πF, µfromlist, "fromlist"); πE != nil {
									continue
								}
								πTemp010[0] = µfromlist
								if πE = πg.CheckLocal(πF, µfromline, "fromline"); πE != nil {
									continue
								}
								πTemp010[1] = µfromline
								if πE = πg.CheckLocal(πF, µfromtext, "fromtext"); πE != nil {
									continue
								}
								πTemp010[2] = µfromtext
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp005, πE = πg.GetAttr(πF, µself, ß_split_line, nil); πE != nil {
									continue
								}
								if πTemp006, πE = πTemp005.Call(πF, πTemp010, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp010)
								// line 1902: self._split_line(tolist,toline,totext)
								πF.SetLineno(1902)
								πTemp010 = πF.MakeArgs(3)
								if πE = πg.CheckLocal(πF, µtolist, "tolist"); πE != nil {
									continue
								}
								πTemp010[0] = µtolist
								if πE = πg.CheckLocal(πF, µtoline, "toline"); πE != nil {
									continue
								}
								πTemp010[1] = µtoline
								if πE = πg.CheckLocal(πF, µtotext, "totext"); πE != nil {
									continue
								}
								πTemp010[2] = µtotext
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp005, πE = πg.GetAttr(πF, µself, ß_split_line, nil); πE != nil {
									continue
								}
								if πTemp006, πE = πTemp005.Call(πF, πTemp010, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp010)
								// line 1905: while fromlist or tolist:
								πF.SetLineno(1905)
								πF.PushCheckpoint(8)
								πTemp003 = false
							Label7:
								if πE != nil || πR != nil {
									continue
								}
								if πTemp003 {
									πF.PopCheckpoint()
									goto Label9
								}
								if πE = πg.CheckLocal(πF, µfromlist, "fromlist"); πE != nil {
									continue
								}
								πTemp005 = µfromlist
								if πTemp012, πE = πg.IsTrue(πF, πTemp005); πE != nil {
									continue
								}
								if πTemp012 {
									goto Label10
								}
								if πE = πg.CheckLocal(πF, µtolist, "tolist"); πE != nil {
									continue
								}
								πTemp005 = µtolist
							Label10:
								if πTemp011, πE = πg.IsTrue(πF, πTemp005); πE != nil {
									continue
								}
								if πE != nil || !πTemp011 {
									continue
								}
								πF.PushCheckpoint(7)            
								if πE = πg.CheckLocal(πF, µfromlist, "fromlist"); πE != nil {
									continue
								}
								if πTemp011, πE = πg.IsTrue(πF, µfromlist); πE != nil {
									continue
								}
								if πTemp011 {
									goto Label11
								}
								goto Label12
								// line 1906: if fromlist:
								πF.SetLineno(1906)
							Label11:
								// line 1907: fromdata = fromlist.pop(0)
								πF.SetLineno(1907)
								πTemp010 = πF.MakeArgs(1)
								πTemp010[0] = πg.NewInt(0).ToObject()
								if πE = πg.CheckLocal(πF, µfromlist, "fromlist"); πE != nil {
									continue
								}
								if πTemp005, πE = πg.GetAttr(πF, µfromlist, ßpop, nil); πE != nil {
									continue
								}
								if πTemp006, πE = πTemp005.Call(πF, πTemp010, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp010)
								µfromdata = πTemp006
								goto Label13
							Label12:
								// line 1909: fromdata = ('',' ')
								πF.SetLineno(1909)
								πTemp005 = πg.NewTuple2(ß.ToObject(), πg.NewStr(" ").ToObject()).ToObject()
								µfromdata = πTemp005
								goto Label13
							Label13:
								if πE = πg.CheckLocal(πF, µtolist, "tolist"); πE != nil {
									continue
								}
								if πTemp011, πE = πg.IsTrue(πF, µtolist); πE != nil {
									continue
								}
								if πTemp011 {
									goto Label14
								}
								goto Label15
								// line 1910: if tolist:
								πF.SetLineno(1910)
							Label14:
								// line 1911: todata = tolist.pop(0)
								πF.SetLineno(1911)
								πTemp010 = πF.MakeArgs(1)
								πTemp010[0] = πg.NewInt(0).ToObject()
								if πE = πg.CheckLocal(πF, µtolist, "tolist"); πE != nil {
									continue
								}
								if πTemp005, πE = πg.GetAttr(πF, µtolist, ßpop, nil); πE != nil {
									continue
								}
								if πTemp006, πE = πTemp005.Call(πF, πTemp010, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp010)
								µtodata = πTemp006
								goto Label16
							Label15:
								// line 1913: todata = ('',' ')
								πF.SetLineno(1913)
								πTemp005 = πg.NewTuple2(ß.ToObject(), πg.NewStr(" ").ToObject()).ToObject()
								µtodata = πTemp005
								goto Label16
							Label16:
								// line 1914: yield fromdata,todata,flag
								πF.SetLineno(1914)
								if πE = πg.CheckLocal(πF, µfromdata, "fromdata"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µtodata, "todata"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µflag, "flag"); πE != nil {
									continue
								}
								πTemp005 = πg.NewTuple3(µfromdata, µtodata, µflag).ToObject()
								πF.PushCheckpoint(17)
								return πTemp005, nil
							Label17:
								πTemp006 = πSent
								continue
							Label8:
								if πE != nil || πR != nil {
									continue
								}
							Label9:
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
					if πE = πClass.SetItem(πF, ß_line_wrapper.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 1889: """Returns iterator that splits (wraps) mdiff text lines"""
					πF.SetLineno(1889)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp007}, πg.NewStr("Returns iterator that splits (wraps) mdiff text lines").ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ß_line_wrapper); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp008, ß__doc__, πTemp007); πE != nil {
						continue
					}
					// line 1916: def _collect_lines(self,diffs):
					πF.SetLineno(1916)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "diffs", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("_collect_lines", "/usr/lib/python2.7/difflib.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µdiffs *πg.Object = πArgs[1]; _ = µdiffs
						var µfromlist *πg.Object = πg.UnboundLocal; _ = µfromlist
						var µtolist *πg.Object = πg.UnboundLocal; _ = µtolist
						var µflaglist *πg.Object = πg.UnboundLocal; _ = µflaglist
						var µfromdata *πg.Object = πg.UnboundLocal; _ = µfromdata
						var µtodata *πg.Object = πg.UnboundLocal; _ = µtodata
						var µflag *πg.Object = πg.UnboundLocal; _ = µflag
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
						var πTemp008 *πg.Object
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
							case 1: goto Label1
							case 2: goto Label2
							case 5: goto Label5
							default: panic("unexpected function state")
							}
							// line 1917: """Collects mdiff output into separate lists
							πF.SetLineno(1917)
							// line 1923: fromlist,tolist,flaglist = [],[],[]
							πF.SetLineno(1923)
							πTemp002 = make([]*πg.Object, 0)
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp002 = make([]*πg.Object, 0)
							πTemp004 = πg.NewList(πTemp002...).ToObject()
							πTemp002 = make([]*πg.Object, 0)
							πTemp005 = πg.NewList(πTemp002...).ToObject()
							πTemp001 = πg.NewTuple3(πTemp003, πTemp004, πTemp005).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}}}, πTemp001); πE != nil {
								continue
							}
							µfromlist = πTemp003
							µtolist = πTemp004
							µflaglist = πTemp005
							if πE = πg.CheckLocal(πF, µdiffs, "diffs"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, µdiffs); πE != nil {
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
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp008}}}, πTemp003); πE != nil {
									continue
								}
								µfromdata = πTemp004
								µtodata = πTemp005
								µflag = πTemp008
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 1926: try:
							πF.SetLineno(1926)
							πF.PushCheckpoint(5)
							// line 1928: fromlist.append(self._format_line(0,flag,*fromdata))
							πF.SetLineno(1928)
							πTemp002 = πF.MakeArgs(1)
							πTemp009 = πF.MakeArgs(2)
							πTemp009[0] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µflag, "flag"); πE != nil {
								continue
							}
							πTemp009[1] = µflag
							if πE = πg.CheckLocal(πF, µfromdata, "fromdata"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_format_line, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Invoke(πF, πTemp003, πTemp009, µfromdata, nil, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							πTemp002[0] = πTemp004
							if πE = πg.CheckLocal(πF, µfromlist, "fromlist"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µfromlist, ßappend, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 1929: tolist.append(self._format_line(1,flag,*todata))
							πF.SetLineno(1929)
							πTemp002 = πF.MakeArgs(1)
							πTemp009 = πF.MakeArgs(2)
							πTemp009[0] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µflag, "flag"); πE != nil {
								continue
							}
							πTemp009[1] = µflag
							if πE = πg.CheckLocal(πF, µtodata, "todata"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_format_line, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Invoke(πF, πTemp003, πTemp009, µtodata, nil, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							πTemp002[0] = πTemp004
							if πE = πg.CheckLocal(πF, µtolist, "tolist"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µtolist, ßappend, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πF.PopCheckpoint()
							goto Label4
						Label5:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp010, πTemp011 = πF.ExcInfo()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsInstance(πF, πTemp010.ToObject(), πTemp003); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label6
							}
							πE = πF.Raise(πTemp010.ToObject(), nil, πTemp011.ToObject())
							continue
							// line 1930: except TypeError:
							πF.SetLineno(1930)
						Label6:
							// line 1932: fromlist.append(None)
							πF.SetLineno(1932)
							πTemp002 = πF.MakeArgs(1)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µfromlist, "fromlist"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µfromlist, ßappend, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 1933: tolist.append(None)
							πF.SetLineno(1933)
							πTemp002 = πF.MakeArgs(1)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µtolist, "tolist"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µtolist, ßappend, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πF.RestoreExc(nil, nil)
							goto Label4
						Label4:
							// line 1934: flaglist.append(flag)
							πF.SetLineno(1934)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µflag, "flag"); πE != nil {
								continue
							}
							πTemp002[0] = µflag
							if πE = πg.CheckLocal(πF, µflaglist, "flaglist"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µflaglist, ßappend, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 1935: return fromlist,tolist,flaglist
							πF.SetLineno(1935)
							if πE = πg.CheckLocal(πF, µfromlist, "fromlist"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtolist, "tolist"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µflaglist, "flaglist"); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple3(µfromlist, µtolist, µflaglist).ToObject()
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
					if πE = πClass.SetItem(πF, ß_collect_lines.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 1917: """Collects mdiff output into separate lists
					πF.SetLineno(1917)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp008}, πg.NewStr("Collects mdiff output into separate lists\n\n        Before storing the mdiff from/to data into a list, it is converted\n        into a single line of text with HTML markup.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ß_collect_lines); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp009, ß__doc__, πTemp008); πE != nil {
						continue
					}
					// line 1937: def _format_line(self,side,flag,linenum,text):
					πF.SetLineno(1937)
					πTemp002 = make([]πg.Param, 5)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "side", Def: nil}
					πTemp002[2] = πg.Param{Name: "flag", Def: nil}
					πTemp002[3] = πg.Param{Name: "linenum", Def: nil}
					πTemp002[4] = πg.Param{Name: "text", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("_format_line", "/usr/lib/python2.7/difflib.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µside *πg.Object = πArgs[1]; _ = µside
						var µflag *πg.Object = πArgs[2]; _ = µflag
						var µlinenum *πg.Object = πArgs[3]; _ = µlinenum
						var µtext *πg.Object = πArgs[4]; _ = µtext
						var µid *πg.Object = πg.UnboundLocal; _ = µid
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
						var πTemp006 *πg.BaseException
						_ = πTemp006
						var πTemp007 *πg.Traceback
						_ = πTemp007
						var πTemp008 bool
						_ = πTemp008
						var πTemp009 []*πg.Object
						_ = πTemp009
						var πTemp010 []*πg.Object
						_ = πTemp010
						var πTemp011 []*πg.Object
						_ = πTemp011
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							// line 1938: """Returns HTML markup of "from" / "to" text lines
							πF.SetLineno(1938)
							// line 1945: try:
							πF.SetLineno(1945)
							πF.PushCheckpoint(2)
							// line 1946: linenum = '%d' % linenum
							πF.SetLineno(1946)
							if πE = πg.CheckLocal(πF, µlinenum, "linenum"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("%d").ToObject(), µlinenum); πE != nil {
								continue
							}
							µlinenum = πTemp001
							// line 1947: id = ' id="%s%s"' % (self._prefix[side],linenum)
							πF.SetLineno(1947)
							if πE = πg.CheckLocal(πF, µside, "side"); πE != nil {
								continue
							}
							πTemp003 = µside
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ß_prefix, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, πTemp005, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µlinenum, "linenum"); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(πTemp004, µlinenum).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr(" id=\"%s%s\"").ToObject(), πTemp002); πE != nil {
								continue
							}
							µid = πTemp001
							πF.PopCheckpoint()
							goto Label1
						Label2:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp006, πTemp007 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp008, πE = πg.IsInstance(πF, πTemp006.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label3
							}
							πE = πF.Raise(πTemp006.ToObject(), nil, πTemp007.ToObject())
							continue
							// line 1948: except TypeError:
							πF.SetLineno(1948)
						Label3:
							// line 1950: id = ''
							πF.SetLineno(1950)
							µid = ß.ToObject()
							πF.RestoreExc(nil, nil)
							goto Label1
						Label1:
							// line 1952: text=text.replace("&","&amp;").replace(">","&gt;").replace("<","&lt;")
							πF.SetLineno(1952)
							πTemp009 = πF.MakeArgs(2)
							πTemp009[0] = πg.NewStr("<").ToObject()
							πTemp009[1] = πg.NewStr("&lt;").ToObject()
							πTemp010 = πF.MakeArgs(2)
							πTemp010[0] = πg.NewStr(">").ToObject()
							πTemp010[1] = πg.NewStr("&gt;").ToObject()
							πTemp011 = πF.MakeArgs(2)
							πTemp011[0] = πg.NewStr("&").ToObject()
							πTemp011[1] = πg.NewStr("&amp;").ToObject()
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µtext, ßreplace, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp011, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp011)
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßreplace, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp010, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp010)
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßreplace, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							µtext = πTemp002
							// line 1955: text = text.replace(' ','&nbsp;').rstrip()
							πF.SetLineno(1955)
							πTemp009 = πF.MakeArgs(2)
							πTemp009[0] = πg.NewStr(" ").ToObject()
							πTemp009[1] = πg.NewStr("&nbsp;").ToObject()
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µtext, ßreplace, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßrstrip, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µtext = πTemp002
							// line 1957: return '<td class="diff_header"%s>%s</td><td nowrap="nowrap">%s</td>' \
							πF.SetLineno(1957)
							if πE = πg.CheckLocal(πF, µid, "id"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µlinenum, "linenum"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple3(µid, µlinenum, µtext).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("<td class=\"diff_header\"%s>%s</td><td nowrap=\"nowrap\">%s</td>").ToObject(), πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ß_format_line.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 1938: """Returns HTML markup of "from" / "to" text lines
					πF.SetLineno(1938)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp009}, πg.NewStr("Returns HTML markup of \"from\" / \"to\" text lines\n\n        side -- 0 or 1 indicating \"from\" or \"to\" text\n        flag -- indicates if difference on line\n        linenum -- line number (used for line number column)\n        text -- line text to be marked up\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ß_format_line); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp010, ß__doc__, πTemp009); πE != nil {
						continue
					}
					// line 1960: def _make_prefix(self):
					πF.SetLineno(1960)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("_make_prefix", "/usr/lib/python2.7/difflib.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µfromprefix *πg.Object = πg.UnboundLocal; _ = µfromprefix
						var µtoprefix *πg.Object = πg.UnboundLocal; _ = µtoprefix
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
							// line 1961: """Create unique anchor prefixes"""
							πF.SetLineno(1961)
							// line 1965: fromprefix = "from%d_" % HtmlDiff._default_prefix
							πF.SetLineno(1965)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßHtmlDiff); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ß_default_prefix, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("from%d_").ToObject(), πTemp003); πE != nil {
								continue
							}
							µfromprefix = πTemp001
							// line 1966: toprefix = "to%d_" % HtmlDiff._default_prefix
							πF.SetLineno(1966)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßHtmlDiff); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ß_default_prefix, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("to%d_").ToObject(), πTemp003); πE != nil {
								continue
							}
							µtoprefix = πTemp001
							// line 1967: HtmlDiff._default_prefix += 1
							πF.SetLineno(1967)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßHtmlDiff); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ß_default_prefix, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.IAdd(πF, πTemp002, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßHtmlDiff); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, πTemp003, ß_default_prefix, πTemp001); πE != nil {
								continue
							}
							// line 1969: self._prefix = [fromprefix,toprefix]
							πF.SetLineno(1969)
							πTemp004 = make([]*πg.Object, 2)
							if πE = πg.CheckLocal(πF, µfromprefix, "fromprefix"); πE != nil {
								continue
							}
							πTemp004[0] = µfromprefix
							if πE = πg.CheckLocal(πF, µtoprefix, "toprefix"); πE != nil {
								continue
							}
							πTemp004[1] = µtoprefix
							πTemp001 = πg.NewList(πTemp004...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_prefix, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ß_make_prefix.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 1961: """Create unique anchor prefixes"""
					πF.SetLineno(1961)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp010}, πg.NewStr("Create unique anchor prefixes").ToObject()); πE != nil {
						continue
					}
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ß_make_prefix); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp011, ß__doc__, πTemp010); πE != nil {
						continue
					}
					// line 1971: def _convert_flags(self,fromlist,tolist,flaglist,context,numlines):
					πF.SetLineno(1971)
					πTemp002 = make([]πg.Param, 6)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "fromlist", Def: nil}
					πTemp002[2] = πg.Param{Name: "tolist", Def: nil}
					πTemp002[3] = πg.Param{Name: "flaglist", Def: nil}
					πTemp002[4] = πg.Param{Name: "context", Def: nil}
					πTemp002[5] = πg.Param{Name: "numlines", Def: nil}
					πTemp010 = πg.NewFunction(πg.NewCode("_convert_flags", "/usr/lib/python2.7/difflib.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µfromlist *πg.Object = πArgs[1]; _ = µfromlist
						var µtolist *πg.Object = πArgs[2]; _ = µtolist
						var µflaglist *πg.Object = πArgs[3]; _ = µflaglist
						var µcontext *πg.Object = πArgs[4]; _ = µcontext
						var µnumlines *πg.Object = πArgs[5]; _ = µnumlines
						var µtoprefix *πg.Object = πg.UnboundLocal; _ = µtoprefix
						var µnext_id *πg.Object = πg.UnboundLocal; _ = µnext_id
						var µnext_href *πg.Object = πg.UnboundLocal; _ = µnext_href
						var µnum_chg *πg.Object = πg.UnboundLocal; _ = µnum_chg
						var µin_change *πg.Object = πg.UnboundLocal; _ = µin_change
						var µlast *πg.Object = πg.UnboundLocal; _ = µlast
						var µi *πg.Object = πg.UnboundLocal; _ = µi
						var µflag *πg.Object = πg.UnboundLocal; _ = µflag
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
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
							// line 1972: """Makes list of "next" links"""
							πF.SetLineno(1972)
							// line 1975: toprefix = self._prefix[1]
							πF.SetLineno(1975)
							πTemp001 = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_prefix, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp003, πTemp001); πE != nil {
								continue
							}
							µtoprefix = πTemp002
							// line 1978: next_id = ['']*len(flaglist)
							πF.SetLineno(1978)
							πTemp004 = make([]*πg.Object, 1)
							πTemp004[0] = ß.ToObject()
							πTemp002 = πg.NewList(πTemp004...).ToObject()
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µflaglist, "flaglist"); πE != nil {
								continue
							}
							πTemp004[0] = µflaglist
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp001, πE = πg.Mul(πF, πTemp002, πTemp005); πE != nil {
								continue
							}
							µnext_id = πTemp001
							// line 1979: next_href = ['']*len(flaglist)
							πF.SetLineno(1979)
							πTemp004 = make([]*πg.Object, 1)
							πTemp004[0] = ß.ToObject()
							πTemp002 = πg.NewList(πTemp004...).ToObject()
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µflaglist, "flaglist"); πE != nil {
								continue
							}
							πTemp004[0] = µflaglist
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp001, πE = πg.Mul(πF, πTemp002, πTemp005); πE != nil {
								continue
							}
							µnext_href = πTemp001
							// line 1980: num_chg, in_change = 0, False
							πF.SetLineno(1980)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple2(πg.NewInt(0).ToObject(), πTemp002).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp003}}}, πTemp001); πE != nil {
								continue
							}
							µnum_chg = πTemp002
							µin_change = πTemp003
							// line 1981: last = 0
							πF.SetLineno(1981)
							µlast = πg.NewInt(0).ToObject()
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µflaglist, "flaglist"); πE != nil {
								continue
							}
							πTemp004[0] = µflaglist
							if πTemp002, πE = πg.ResolveGlobal(πF, ßenumerate); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp001, πE = πg.Iter(πF, πTemp003); πE != nil {
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
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp005}}}, πTemp002); πE != nil {
									continue
								}
								µi = πTemp003
								µflag = πTemp005
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(1)            
							if πE = πg.CheckLocal(πF, µflag, "flag"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, µflag); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label4
							}
							goto Label5
							// line 1983: if flag:
							πF.SetLineno(1983)
						Label4:
							if πE = πg.CheckLocal(πF, µin_change, "in_change"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, µin_change); πE != nil {
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
							// line 1984: if not in_change:
							πF.SetLineno(1984)
						Label7:
							// line 1985: in_change = True
							πF.SetLineno(1985)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							µin_change = πTemp002
							// line 1986: last = i
							πF.SetLineno(1986)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							µlast = µi
							// line 1990: i = max([0,i-numlines])
							πF.SetLineno(1990)
							πTemp004 = πF.MakeArgs(1)
							πTemp008 = make([]*πg.Object, 2)
							πTemp008[0] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnumlines, "numlines"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Sub(πF, µi, µnumlines); πE != nil {
								continue
							}
							πTemp008[1] = πTemp002
							πTemp002 = πg.NewList(πTemp008...).ToObject()
							πTemp004[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßmax); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µi = πTemp003
							// line 1991: next_id[i] = ' id="difflib_chg_%s_%d"' % (toprefix,num_chg)
							πF.SetLineno(1991)
							if πE = πg.CheckLocal(πF, µtoprefix, "toprefix"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnum_chg, "num_chg"); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple2(µtoprefix, µnum_chg).ToObject()
							if πTemp002, πE = πg.Mod(πF, πg.NewStr(" id=\"difflib_chg_%s_%d\"").ToObject(), πTemp003); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnext_id, "next_id"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp005 = µi
							if πE = πg.SetItem(πF, µnext_id, πTemp005, πTemp003); πE != nil {
								continue
							}
							// line 1994: num_chg += 1
							πF.SetLineno(1994)
							if πE = πg.CheckLocal(πF, µnum_chg, "num_chg"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IAdd(πF, µnum_chg, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µnum_chg = πTemp002
							// line 1995: next_href[last] = '<a href="#difflib_chg_%s_%d">n</a>' % (
							πF.SetLineno(1995)
							if πE = πg.CheckLocal(πF, µtoprefix, "toprefix"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnum_chg, "num_chg"); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple2(µtoprefix, µnum_chg).ToObject()
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("<a href=\"#difflib_chg_%s_%d\">n</a>").ToObject(), πTemp003); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnext_href, "next_href"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µlast, "last"); πE != nil {
								continue
							}
							πTemp005 = µlast
							if πE = πg.SetItem(πF, µnext_href, πTemp005, πTemp003); πE != nil {
								continue
							}
							goto Label8
						Label8:
							goto Label6
						Label5:
							// line 1998: in_change = False
							πF.SetLineno(1998)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							µin_change = πTemp002
							goto Label6
						Label6:
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							if πE = πg.CheckLocal(πF, µflaglist, "flaglist"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, µflaglist); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp006).ToObject()
							if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label9
							}
							goto Label10
							// line 2000: if not flaglist:
							πF.SetLineno(2000)
						Label9:
							// line 2001: flaglist = [False]
							πF.SetLineno(2001)
							πTemp004 = make([]*πg.Object, 1)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
							πTemp001 = πg.NewList(πTemp004...).ToObject()
							µflaglist = πTemp001
							// line 2002: next_id = ['']
							πF.SetLineno(2002)
							πTemp004 = make([]*πg.Object, 1)
							πTemp004[0] = ß.ToObject()
							πTemp001 = πg.NewList(πTemp004...).ToObject()
							µnext_id = πTemp001
							// line 2003: next_href = ['']
							πF.SetLineno(2003)
							πTemp004 = make([]*πg.Object, 1)
							πTemp004[0] = ß.ToObject()
							πTemp001 = πg.NewList(πTemp004...).ToObject()
							µnext_href = πTemp001
							// line 2004: last = 0
							πF.SetLineno(2004)
							µlast = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µcontext, "context"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, µcontext); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label11
							}
							goto Label12
							// line 2005: if context:
							πF.SetLineno(2005)
						Label11:
							// line 2006: fromlist = ['<td></td><td>&nbsp;No Differences Found&nbsp;</td>']
							πF.SetLineno(2006)
							πTemp004 = make([]*πg.Object, 1)
							πTemp004[0] = πg.NewStr("<td></td><td>&nbsp;No Differences Found&nbsp;</td>").ToObject()
							πTemp001 = πg.NewList(πTemp004...).ToObject()
							µfromlist = πTemp001
							// line 2007: tolist = fromlist
							πF.SetLineno(2007)
							if πE = πg.CheckLocal(πF, µfromlist, "fromlist"); πE != nil {
								continue
							}
							µtolist = µfromlist
							goto Label13
						Label12:
							// line 2009: fromlist = tolist = ['<td></td><td>&nbsp;Empty File&nbsp;</td>']
							πF.SetLineno(2009)
							πTemp004 = make([]*πg.Object, 1)
							πTemp004[0] = πg.NewStr("<td></td><td>&nbsp;Empty File&nbsp;</td>").ToObject()
							πTemp001 = πg.NewList(πTemp004...).ToObject()
							µfromlist = πTemp001
							µtolist = πTemp001
							goto Label13
						Label13:
							goto Label10
						Label10:
							πTemp002 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µflaglist, "flaglist"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µflaglist, πTemp002); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp006).ToObject()
							if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label14
							}
							goto Label15
							// line 2011: if not flaglist[0]:
							πF.SetLineno(2011)
						Label14:
							// line 2012: next_href[0] = '<a href="#difflib_chg_%s_0">f</a>' % toprefix
							πF.SetLineno(2012)
							if πE = πg.CheckLocal(πF, µtoprefix, "toprefix"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("<a href=\"#difflib_chg_%s_0\">f</a>").ToObject(), µtoprefix); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnext_href, "next_href"); πE != nil {
								continue
							}
							πTemp003 = πg.NewInt(0).ToObject()
							if πE = πg.SetItem(πF, µnext_href, πTemp003, πTemp002); πE != nil {
								continue
							}
							goto Label15
						Label15:
							// line 2014: next_href[last] = '<a href="#difflib_chg_%s_top">t</a>' % (toprefix)
							πF.SetLineno(2014)
							if πE = πg.CheckLocal(πF, µtoprefix, "toprefix"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("<a href=\"#difflib_chg_%s_top\">t</a>").ToObject(), µtoprefix); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnext_href, "next_href"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µlast, "last"); πE != nil {
								continue
							}
							πTemp003 = µlast
							if πE = πg.SetItem(πF, µnext_href, πTemp003, πTemp002); πE != nil {
								continue
							}
							// line 2016: return fromlist,tolist,flaglist,next_href,next_id
							πF.SetLineno(2016)
							if πE = πg.CheckLocal(πF, µfromlist, "fromlist"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtolist, "tolist"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µflaglist, "flaglist"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnext_href, "next_href"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnext_id, "next_id"); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple5(µfromlist, µtolist, µflaglist, µnext_href, µnext_id).ToObject()
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
					if πE = πClass.SetItem(πF, ß_convert_flags.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 1972: """Makes list of "next" links"""
					πF.SetLineno(1972)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πg.NewStr("Makes list of \"next\" links").ToObject()); πE != nil {
						continue
					}
					if πTemp012, πE = πg.ResolveClass(πF, πClass, nil, ß_convert_flags); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp012, ß__doc__, πTemp011); πE != nil {
						continue
					}
					// line 2018: def make_table(self,fromlines,tolines,fromdesc='',todesc='',context=False,
					πF.SetLineno(2018)
					πTemp002 = make([]πg.Param, 7)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "fromlines", Def: nil}
					πTemp002[2] = πg.Param{Name: "tolines", Def: nil}
					πTemp002[3] = πg.Param{Name: "fromdesc", Def: ß.ToObject()}
					πTemp002[4] = πg.Param{Name: "todesc", Def: ß.ToObject()}
					if πTemp012, πE = πg.ResolveClass(πF, πClass, nil, ßFalse); πE != nil {
						continue
					}
					πTemp002[5] = πg.Param{Name: "context", Def: πTemp012}
					πTemp002[6] = πg.Param{Name: "numlines", Def: πg.NewInt(5).ToObject()}
					πTemp011 = πg.NewFunction(πg.NewCode("make_table", "/usr/lib/python2.7/difflib.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µfromlines *πg.Object = πArgs[1]; _ = µfromlines
						var µtolines *πg.Object = πArgs[2]; _ = µtolines
						var µfromdesc *πg.Object = πArgs[3]; _ = µfromdesc
						var µtodesc *πg.Object = πArgs[4]; _ = µtodesc
						var µcontext *πg.Object = πArgs[5]; _ = µcontext
						var µnumlines *πg.Object = πArgs[6]; _ = µnumlines
						var µcontext_lines *πg.Object = πg.UnboundLocal; _ = µcontext_lines
						var µdiffs *πg.Object = πg.UnboundLocal; _ = µdiffs
						var µfromlist *πg.Object = πg.UnboundLocal; _ = µfromlist
						var µtolist *πg.Object = πg.UnboundLocal; _ = µtolist
						var µflaglist *πg.Object = πg.UnboundLocal; _ = µflaglist
						var µnext_href *πg.Object = πg.UnboundLocal; _ = µnext_href
						var µnext_id *πg.Object = πg.UnboundLocal; _ = µnext_id
						var µs *πg.Object = πg.UnboundLocal; _ = µs
						var µfmt *πg.Object = πg.UnboundLocal; _ = µfmt
						var µi *πg.Object = πg.UnboundLocal; _ = µi
						var µheader_row *πg.Object = πg.UnboundLocal; _ = µheader_row
						var µtable *πg.Object = πg.UnboundLocal; _ = µtable
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
						var πTemp006 πg.KWArgs
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 *πg.Object
						_ = πTemp009
						var πTemp010 []*πg.Object
						_ = πTemp010
						var πTemp011 bool
						_ = πTemp011
						var πTemp012 *πg.Object
						_ = πTemp012
						var πTemp013 *πg.Object
						_ = πTemp013
						var πTemp014 *πg.Object
						_ = πTemp014
						var πTemp015 []*πg.Object
						_ = πTemp015
						var πTemp016 []*πg.Object
						_ = πTemp016
						var πTemp017 []*πg.Object
						_ = πTemp017
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 6: goto Label6
							case 7: goto Label7
							default: panic("unexpected function state")
							}
							// line 2020: """Returns HTML table of side by side comparison with change highlights
							πF.SetLineno(2020)
							// line 2038: self._make_prefix()
							πF.SetLineno(2038)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_make_prefix, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 2042: fromlines,tolines = self._tab_newline_replace(fromlines,tolines)
							πF.SetLineno(2042)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µfromlines, "fromlines"); πE != nil {
								continue
							}
							πTemp003[0] = µfromlines
							if πE = πg.CheckLocal(πF, µtolines, "tolines"); πE != nil {
								continue
							}
							πTemp003[1] = µtolines
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_tab_newline_replace, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp004}}}, πTemp002); πE != nil {
								continue
							}
							µfromlines = πTemp001
							µtolines = πTemp004
							if πE = πg.CheckLocal(πF, µcontext, "context"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, µcontext); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							goto Label2
							// line 2045: if context:
							πF.SetLineno(2045)
						Label1:
							// line 2046: context_lines = numlines
							πF.SetLineno(2046)
							if πE = πg.CheckLocal(πF, µnumlines, "numlines"); πE != nil {
								continue
							}
							µcontext_lines = µnumlines
							goto Label3
						Label2:
							// line 2048: context_lines = None
							πF.SetLineno(2048)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							µcontext_lines = πTemp001
							goto Label3
						Label3:
							// line 2049: diffs = _mdiff(fromlines,tolines,context_lines,linejunk=self._linejunk,
							πF.SetLineno(2049)
							πTemp003 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µfromlines, "fromlines"); πE != nil {
								continue
							}
							πTemp003[0] = µfromlines
							if πE = πg.CheckLocal(πF, µtolines, "tolines"); πE != nil {
								continue
							}
							πTemp003[1] = µtolines
							if πE = πg.CheckLocal(πF, µcontext_lines, "context_lines"); πE != nil {
								continue
							}
							πTemp003[2] = µcontext_lines
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_linejunk, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_charjunk, nil); πE != nil {
								continue
							}
							πTemp006 = πg.KWArgs{
								{"linejunk", πTemp001},
								{"charjunk", πTemp002},
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_mdiff); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, πTemp006); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µdiffs = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_wrapcolumn, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label4
							}
							goto Label5
							// line 2053: if self._wrapcolumn:
							πF.SetLineno(2053)
						Label4:
							// line 2054: diffs = self._line_wrapper(diffs)
							πF.SetLineno(2054)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdiffs, "diffs"); πE != nil {
								continue
							}
							πTemp003[0] = µdiffs
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_line_wrapper, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µdiffs = πTemp002
							goto Label5
						Label5:
							// line 2057: fromlist,tolist,flaglist = self._collect_lines(diffs)
							πF.SetLineno(2057)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdiffs, "diffs"); πE != nil {
								continue
							}
							πTemp003[0] = µdiffs
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_collect_lines, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp002); πE != nil {
								continue
							}
							µfromlist = πTemp001
							µtolist = πTemp004
							µflaglist = πTemp007
							// line 2060: fromlist,tolist,flaglist,next_href,next_id = self._convert_flags(
							πF.SetLineno(2060)
							πTemp003 = πF.MakeArgs(5)
							if πE = πg.CheckLocal(πF, µfromlist, "fromlist"); πE != nil {
								continue
							}
							πTemp003[0] = µfromlist
							if πE = πg.CheckLocal(πF, µtolist, "tolist"); πE != nil {
								continue
							}
							πTemp003[1] = µtolist
							if πE = πg.CheckLocal(πF, µflaglist, "flaglist"); πE != nil {
								continue
							}
							πTemp003[2] = µflaglist
							if πE = πg.CheckLocal(πF, µcontext, "context"); πE != nil {
								continue
							}
							πTemp003[3] = µcontext
							if πE = πg.CheckLocal(πF, µnumlines, "numlines"); πE != nil {
								continue
							}
							πTemp003[4] = µnumlines
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_convert_flags, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}, πg.TieTarget{Target: &πTemp008}, πg.TieTarget{Target: &πTemp009}}}, πTemp002); πE != nil {
								continue
							}
							µfromlist = πTemp001
							µtolist = πTemp004
							µflaglist = πTemp007
							µnext_href = πTemp008
							µnext_id = πTemp009
							// line 2063: s = []
							πF.SetLineno(2063)
							πTemp003 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp003...).ToObject()
							µs = πTemp001
							// line 2064: fmt = '            <tr><td class="diff_next"%s>%s</td>%s' + \
							πF.SetLineno(2064)
							if πTemp001, πE = πg.Add(πF, πg.NewStr("            <tr><td class=\"diff_next\"%s>%s</td>%s").ToObject(), πg.NewStr("<td class=\"diff_next\">%s</td>%s</tr>\n").ToObject()); πE != nil {
								continue
							}
							µfmt = πTemp001
							πTemp003 = πF.MakeArgs(1)
							πTemp010 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µflaglist, "flaglist"); πE != nil {
								continue
							}
							πTemp010[0] = µflaglist
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp010, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp010)
							πTemp003[0] = πTemp004
							if πTemp002, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp001, πE = πg.Iter(πF, πTemp004); πE != nil {
								continue
							}
							πF.PushCheckpoint(7)
							πTemp005 = false
						Label6:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp005 {
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
								πTemp011 = !isStop
							} else {
								πTemp011 = true
								µi = πTemp002
							}
							if πE != nil || !πTemp011 {
								continue
							}
							πF.PushCheckpoint(6)            
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp004 = µi
							if πE = πg.CheckLocal(πF, µflaglist, "flaglist"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, µflaglist, πTemp004); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp007 == πTemp004).ToObject()
							if πTemp011, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp011 {
								goto Label9
							}
							goto Label10
							// line 2067: if flaglist[i] is None:
							πF.SetLineno(2067)
						Label9:
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GT(πF, µi, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp011, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp011 {
								goto Label12
							}
							goto Label13
							// line 2070: if i > 0:
							πF.SetLineno(2070)
						Label12:
							// line 2071: s.append('        </tbody>        \n        <tbody>\n')
							πF.SetLineno(2071)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewStr("        </tbody>        \n        <tbody>\n").ToObject()
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µs, ßappend, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							goto Label13
						Label13:
							goto Label11
						Label10:
							// line 2073: s.append( fmt % (next_id[i],next_href[i],fromlist[i],
							πF.SetLineno(2073)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfmt, "fmt"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp007 = µi
							if πE = πg.CheckLocal(πF, µnext_id, "next_id"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, µnext_id, πTemp007); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp007 = µi
							if πE = πg.CheckLocal(πF, µnext_href, "next_href"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetItem(πF, µnext_href, πTemp007); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp007 = µi
							if πE = πg.CheckLocal(πF, µfromlist, "fromlist"); πE != nil {
								continue
							}
							if πTemp012, πE = πg.GetItem(πF, µfromlist, πTemp007); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp007 = µi
							if πE = πg.CheckLocal(πF, µnext_href, "next_href"); πE != nil {
								continue
							}
							if πTemp013, πE = πg.GetItem(πF, µnext_href, πTemp007); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp007 = µi
							if πE = πg.CheckLocal(πF, µtolist, "tolist"); πE != nil {
								continue
							}
							if πTemp014, πE = πg.GetItem(πF, µtolist, πTemp007); πE != nil {
								continue
							}
							πTemp004 = πg.NewTuple5(πTemp008, πTemp009, πTemp012, πTemp013, πTemp014).ToObject()
							if πTemp002, πE = πg.Mod(πF, µfmt, πTemp004); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µs, ßappend, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							goto Label11
						Label11:
							continue
						Label7:
							if πE != nil || πR != nil {
								continue
							}
						Label8:
							if πE = πg.CheckLocal(πF, µfromdesc, "fromdesc"); πE != nil {
								continue
							}
							πTemp001 = µfromdesc
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label14
							}
							if πE = πg.CheckLocal(πF, µtodesc, "todesc"); πE != nil {
								continue
							}
							πTemp001 = µtodesc
						Label14:
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label15
							}
							goto Label16
							// line 2075: if fromdesc or todesc:
							πF.SetLineno(2075)
						Label15:
							// line 2076: header_row = '<thead><tr>%s%s%s%s</tr></thead>' % (
							πF.SetLineno(2076)
							if πE = πg.CheckLocal(πF, µfromdesc, "fromdesc"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Mod(πF, πg.NewStr("<th colspan=\"2\" class=\"diff_header\">%s</th>").ToObject(), µfromdesc); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtodesc, "todesc"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Mod(πF, πg.NewStr("<th colspan=\"2\" class=\"diff_header\">%s</th>").ToObject(), µtodesc); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple4(πg.NewStr("<th class=\"diff_next\"><br /></th>").ToObject(), πTemp004, πg.NewStr("<th class=\"diff_next\"><br /></th>").ToObject(), πTemp007).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("<thead><tr>%s%s%s%s</tr></thead>").ToObject(), πTemp002); πE != nil {
								continue
							}
							µheader_row = πTemp001
							goto Label17
						Label16:
							// line 2082: header_row = ''
							πF.SetLineno(2082)
							µheader_row = ß.ToObject()
							goto Label17
						Label17:
							// line 2084: table = self._table_template % dict(
							πF.SetLineno(2084)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_table_template, nil); πE != nil {
								continue
							}
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp003[0] = µs
							if πTemp004, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πE = πg.CheckLocal(πF, µheader_row, "header_row"); πE != nil {
								continue
							}
							πTemp004 = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, µself, ß_prefix, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, πTemp009, πTemp004); πE != nil {
								continue
							}
							πTemp006 = πg.KWArgs{
								{"data_rows", πTemp007},
								{"header_row", µheader_row},
								{"prefix", πTemp008},
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp004.Call(πF, nil, πTemp006); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πTemp002, πTemp007); πE != nil {
								continue
							}
							µtable = πTemp001
							// line 2089: return table.replace('\0+','<span class="diff_add">'). \
							πF.SetLineno(2089)
							πTemp003 = πF.MakeArgs(2)
							πTemp003[0] = πg.NewStr("\t").ToObject()
							πTemp003[1] = πg.NewStr("&nbsp;").ToObject()
							πTemp010 = πF.MakeArgs(2)
							πTemp010[0] = πg.NewStr("\x01").ToObject()
							πTemp010[1] = πg.NewStr("</span>").ToObject()
							πTemp015 = πF.MakeArgs(2)
							πTemp015[0] = πg.NewStr("\x00^").ToObject()
							πTemp015[1] = πg.NewStr("<span class=\"diff_chg\">").ToObject()
							πTemp016 = πF.MakeArgs(2)
							πTemp016[0] = πg.NewStr("\x00-").ToObject()
							πTemp016[1] = πg.NewStr("<span class=\"diff_sub\">").ToObject()
							πTemp017 = πF.MakeArgs(2)
							πTemp017[0] = πg.NewStr("\x00+").ToObject()
							πTemp017[1] = πg.NewStr("<span class=\"diff_add\">").ToObject()
							if πE = πg.CheckLocal(πF, µtable, "table"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µtable, ßreplace, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp017, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp017)
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßreplace, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp016, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp016)
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßreplace, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp015, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp015)
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßreplace, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp010, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp010)
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßreplace, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßmake_table.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 2020: """Returns HTML table of side by side comparison with change highlights
					πF.SetLineno(2020)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp012}, πg.NewStr("Returns HTML table of side by side comparison with change highlights\n\n        Arguments:\n        fromlines -- list of \"from\" lines\n        tolines -- list of \"to\" lines\n        fromdesc -- \"from\" file column header string\n        todesc -- \"to\" file column header string\n        context -- set to True for contextual differences (defaults to False\n            which shows full differences).\n        numlines -- number of context lines.  When context is set True,\n            controls number of lines displayed before and after the change.\n            When context is False, controls the number of lines to place\n            the \"next\" link anchors before the next change (so click of\n            \"next\" link jumps to just before the change).\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp013, πE = πg.ResolveClass(πF, πClass, nil, ßmake_table); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp013, ß__doc__, πTemp012); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp017, πE = πTemp005.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp017 == nil {
				πTemp017 = πg.TypeType.ToObject()
			}
			if πTemp018, πE = πTemp017.Call(πF, []*πg.Object{πg.NewStr("HtmlDiff").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp005.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßHtmlDiff.ToObject(), πTemp018); πE != nil {
				continue
			}
			// line 2095: del re
			πF.SetLineno(2095)
			if πE = πg.DelVar(πF, πF.Globals(), ßre); πE != nil {
				continue
			}
			// line 2097: def restore(delta, which):
			πF.SetLineno(2097)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "delta", Def: nil}
			πTemp004[1] = πg.Param{Name: "which", Def: nil}
			πTemp016 = πg.NewFunction(πg.NewCode("restore", "/usr/lib/python2.7/difflib.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µdelta *πg.Object = πArgs[0]; _ = µdelta
				var µwhich *πg.Object = πArgs[1]; _ = µwhich
				var µtag *πg.Object = πg.UnboundLocal; _ = µtag
				var µprefixes *πg.Object = πg.UnboundLocal; _ = µprefixes
				var µline *πg.Object = πg.UnboundLocal; _ = µline
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []*πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 *πg.Dict
				_ = πTemp005
				var πTemp006 *πg.BaseException
				_ = πTemp006
				var πTemp007 *πg.Traceback
				_ = πTemp007
				var πTemp008 bool
				_ = πTemp008
				var πTemp009 bool
				_ = πTemp009
				var πTemp010 *πg.Object
				_ = πTemp010
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
					for ; πF.State() >= 0; πF.PopCheckpoint() {
						switch πF.State() {
						case 0:
						case 9: goto Label9
						case 2: goto Label2
						case 4: goto Label4
						case 5: goto Label5
						default: panic("unexpected function state")
						}
						// line 2098: r"""
						πF.SetLineno(2098)
						// line 2119: try:
						πF.SetLineno(2119)
						πF.PushCheckpoint(2)
						// line 2120: tag = {1: "- ", 2: "+ "}[int(which)]
						πF.SetLineno(2120)
						πTemp002 = πF.MakeArgs(1)
						if πE = πg.CheckLocal(πF, µwhich, "which"); πE != nil {
							continue
						}
						πTemp002[0] = µwhich
						if πTemp003, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
							continue
						}
						if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp002)
						πTemp001 = πTemp004
						πTemp005 = πg.NewDict()
						if πE = πTemp005.SetItem(πF, πg.NewInt(1).ToObject(), πg.NewStr("- ").ToObject()); πE != nil {
							continue
						}
						if πE = πTemp005.SetItem(πF, πg.NewInt(2).ToObject(), πg.NewStr("+ ").ToObject()); πE != nil {
							continue
						}
						πTemp004 = πTemp005.ToObject()
						if πTemp003, πE = πg.GetItem(πF, πTemp004, πTemp001); πE != nil {
							continue
						}
						µtag = πTemp003
						πF.PopCheckpoint()
						goto Label1
					Label2:
						if πE == nil {
						  continue
						}
						πE = nil
						πTemp006, πTemp007 = πF.ExcInfo()
						if πTemp001, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
							continue
						}
						if πTemp008, πE = πg.IsInstance(πF, πTemp006.ToObject(), πTemp001); πE != nil {
							continue
						}
						if πTemp008 {
							goto Label3
						}
						πE = πF.Raise(πTemp006.ToObject(), nil, πTemp007.ToObject())
						continue
						// line 2121: except KeyError:
						πF.SetLineno(2121)
					Label3:
						if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µwhich, "which"); πE != nil {
							continue
						}
						if πTemp003, πE = πg.Mod(πF, πg.NewStr("unknown delta choice (must be 1 or 2): %r").ToObject(), µwhich); πE != nil {
							continue
						}
						// line 2122: raise ValueError, ('unknown delta choice (must be 1 or 2): %r'
						πF.SetLineno(2122)
						πE = πF.Raise(πTemp001, πTemp003, nil)
						continue
						πF.RestoreExc(nil, nil)
						goto Label1
					Label1:
						// line 2124: prefixes = ("  ", tag)
						πF.SetLineno(2124)
						if πE = πg.CheckLocal(πF, µtag, "tag"); πE != nil {
							continue
						}
						πTemp001 = πg.NewTuple2(πg.NewStr("  ").ToObject(), µtag).ToObject()
						µprefixes = πTemp001
						if πE = πg.CheckLocal(πF, µdelta, "delta"); πE != nil {
							continue
						}
						if πTemp001, πE = πg.Iter(πF, µdelta); πE != nil {
							continue
						}
						πF.PushCheckpoint(5)
						πTemp008 = false
					Label4:
						if πE != nil || πR != nil {
							continue
						}
						if πTemp008 {
							πF.PopCheckpoint()
							goto Label6
						}
						if πTemp003, πE = πg.Next(πF, πTemp001); πE != nil {
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
							µline = πTemp003
						}
						if πE != nil || !πTemp009 {
							continue
						}
						πF.PushCheckpoint(4)            
						if πTemp004, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.NewInt(2).ToObject(), πg.None}, nil); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
							continue
						}
						if πTemp010, πE = πg.GetItem(πF, µline, πTemp004); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µprefixes, "prefixes"); πE != nil {
							continue
						}
						if πTemp009, πE = πg.Contains(πF, µprefixes, πTemp010); πE != nil {
							continue
						}
						πTemp003 = πg.GetBool(πTemp009).ToObject()
						if πTemp009, πE = πg.IsTrue(πF, πTemp003); πE != nil {
							continue
						}
						if πTemp009 {
							goto Label7
						}
						goto Label8
						// line 2126: if line[:2] in prefixes:
						πF.SetLineno(2126)
					Label7:
						// line 2127: yield line[2:]
						πF.SetLineno(2127)
						if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(2).ToObject(), πg.None, πg.None}, nil); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
							continue
						}
						if πTemp004, πE = πg.GetItem(πF, µline, πTemp003); πE != nil {
							continue
						}
						πF.PushCheckpoint(9)
						return πTemp004, nil
					Label9:
						πTemp003 = πSent
						goto Label8
					Label8:
						continue
					Label5:
						if πE != nil || πR != nil {
							continue
						}
					Label6:
					}
					return nil, πE
				}).ToObject(), nil
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßrestore.ToObject(), πTemp016); πE != nil {
				continue
			}
			// line 2098: r"""
			πF.SetLineno(2098)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp017}, πg.NewStr("\n    Generate one of the two sequences that generated a delta.\n\n    Given a `delta` produced by `Differ.compare()` or `ndiff()`, extract\n    lines originating from file 1 or 2 (parameter `which`), stripping off line\n    prefixes.\n\n    Examples:\n\n    >>> diff = ndiff('one\\ntwo\\nthree\\n'.splitlines(1),\n    ...              'ore\\ntree\\nemu\\n'.splitlines(1))\n    >>> diff = list(diff)\n    >>> print ''.join(restore(diff, 1)),\n    one\n    two\n    three\n    >>> print ''.join(restore(diff, 2)),\n    ore\n    tree\n    emu\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp018, πE = πg.ResolveGlobal(πF, ßrestore); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp018, ß__doc__, πTemp017); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("difflib", Code)
}

