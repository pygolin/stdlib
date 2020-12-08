package collections
import (
	πg "github.com/pygolin/runtime"
	// _ "github.com/pygolin/stdlib/heapq"
	// _ "github.com/pygolin/stdlib/_collections"
	// _ "github.com/pygolin/stdlib/keyword"
	// _ "github.com/pygolin/stdlib/sys"
	// _ "github.com/pygolin/stdlib/thread"
	// _ "github.com/pygolin/stdlib/itertools"
	// _ "github.com/pygolin/stdlib/_abcoll"
	// _ "github.com/pygolin/stdlib/operator"
)
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/collections.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ßAttributeError := πg.InternStr("AttributeError")
		ßCounter := πg.InternStr("Counter")
		ßItemsView := πg.InternStr("ItemsView")
		ßKeyError := πg.InternStr("KeyError")
		ßKeysView := πg.InternStr("KeysView")
		ßMapping := πg.InternStr("Mapping")
		ßMutableMapping := πg.InternStr("MutableMapping")
		ßNone := πg.InternStr("None")
		ßNotImplemented := πg.InternStr("NotImplemented")
		ßNotImplementedError := πg.InternStr("NotImplementedError")
		ßOrderedDict := πg.InternStr("OrderedDict")
		ßTrue := πg.InternStr("True")
		ßTypeError := πg.InternStr("TypeError")
		ßValuesView := πg.InternStr("ValuesView")
		ß__add__ := πg.InternStr("__add__")
		ß__all__ := πg.InternStr("__all__")
		ß__and__ := πg.InternStr("__and__")
		ß__class__ := πg.InternStr("__class__")
		ß__delitem__ := πg.InternStr("__delitem__")
		ß__doc__ := πg.InternStr("__doc__")
		ß__eq__ := πg.InternStr("__eq__")
		ß__init__ := πg.InternStr("__init__")
		ß__iter__ := πg.InternStr("__iter__")
		ß__main__ := πg.InternStr("__main__")
		ß__map := πg.InternStr("__map")
		ß__marker := πg.InternStr("__marker")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__missing__ := πg.InternStr("__missing__")
		ß__mod__ := πg.InternStr("__mod__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß__ne__ := πg.InternStr("__ne__")
		ß__or__ := πg.InternStr("__or__")
		ß__reduce__ := πg.InternStr("__reduce__")
		ß__repr__ := πg.InternStr("__repr__")
		ß__reversed__ := πg.InternStr("__reversed__")
		ß__root := πg.InternStr("__root")
		ß__setitem__ := πg.InternStr("__setitem__")
		ß__sub__ := πg.InternStr("__sub__")
		ß__update := πg.InternStr("__update")
		ß_abcoll := πg.InternStr("_abcoll")
		ß_chain := πg.InternStr("_chain")
		ß_class_template := πg.InternStr("_class_template")
		ß_collections := πg.InternStr("_collections")
		ß_eq := πg.InternStr("_eq")
		ß_field_template := πg.InternStr("_field_template")
		ß_get_ident := πg.InternStr("_get_ident")
		ß_heapq := πg.InternStr("_heapq")
		ß_imap := πg.InternStr("_imap")
		ß_iskeyword := πg.InternStr("_iskeyword")
		ß_itemgetter := πg.InternStr("_itemgetter")
		ß_repeat := πg.InternStr("_repeat")
		ß_repr_template := πg.InternStr("_repr_template")
		ß_starmap := πg.InternStr("_starmap")
		ß_sys := πg.InternStr("_sys")
		ßall := πg.InternStr("all")
		ßchain := πg.InternStr("chain")
		ßclassmethod := πg.InternStr("classmethod")
		ßclear := πg.InternStr("clear")
		ßcopy := πg.InternStr("copy")
		ßdeque := πg.InternStr("deque")
		ßdict := πg.InternStr("dict")
		ßelements := πg.InternStr("elements")
		ßeq := πg.InternStr("eq")
		ßfrom_iterable := πg.InternStr("from_iterable")
		ßfromkeys := πg.InternStr("fromkeys")
		ßget := πg.InternStr("get")
		ßget_ident := πg.InternStr("get_ident")
		ßgetattr := πg.InternStr("getattr")
		ßglobals := πg.InternStr("globals")
		ßid := πg.InternStr("id")
		ßimap := πg.InternStr("imap")
		ßisinstance := πg.InternStr("isinstance")
		ßiskeyword := πg.InternStr("iskeyword")
		ßitemgetter := πg.InternStr("itemgetter")
		ßitems := πg.InternStr("items")
		ßiter := πg.InternStr("iter")
		ßiteritems := πg.InternStr("iteritems")
		ßiterkeys := πg.InternStr("iterkeys")
		ßitertools := πg.InternStr("itertools")
		ßitervalues := πg.InternStr("itervalues")
		ßjoin := πg.InternStr("join")
		ßkeys := πg.InternStr("keys")
		ßkeyword := πg.InternStr("keyword")
		ßlen := πg.InternStr("len")
		ßlist := πg.InternStr("list")
		ßmap := πg.InternStr("map")
		ßmost_common := πg.InternStr("most_common")
		ßname := πg.InternStr("name")
		ßnamedtuple := πg.InternStr("namedtuple")
		ßnext := πg.InternStr("next")
		ßnlargest := πg.InternStr("nlargest")
		ßobject := πg.InternStr("object")
		ßoperator := πg.InternStr("operator")
		ßpop := πg.InternStr("pop")
		ßpopitem := πg.InternStr("popitem")
		ßrepeat := πg.InternStr("repeat")
		ßreversed := πg.InternStr("reversed")
		ßsetdefault := πg.InternStr("setdefault")
		ßsorted := πg.InternStr("sorted")
		ßstarmap := πg.InternStr("starmap")
		ßsubtract := πg.InternStr("subtract")
		ßsuper := πg.InternStr("super")
		ßthread := πg.InternStr("thread")
		ßupdate := πg.InternStr("update")
		ßvalues := πg.InternStr("values")
		ßvars := πg.InternStr("vars")
		ßviewitems := πg.InternStr("viewitems")
		ßviewkeys := πg.InternStr("viewkeys")
		ßviewvalues := πg.InternStr("viewvalues")
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
		var πTemp010 *πg.Dict
		_ = πTemp010
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			case 1: goto Label1
			case 2: goto Label2
			default: panic("unexpected function state")
			}
			// line 1: '''This module implements specialized container datatypes providing
			πF.SetLineno(1)
			// line 1: '''This module implements specialized container datatypes providing
			πF.SetLineno(1)
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("This module implements specialized container datatypes providing\nalternatives to Python's general purpose built-in containers, dict,\nlist, set, and tuple.\n\n* namedtuple   factory function for creating tuple subclasses with named fields\n* deque        list-like container with fast appends and pops on either end\n* Counter      dict subclass for counting hashable objects\n* OrderedDict  dict subclass that remembers the order entries were added\n* defaultdict  dict subclass that calls a factory function to supply missing values\n\n").ToObject()); πE != nil {
				continue
			}
			// line 13: __all__ = ['Counter', 'namedtuple', 'OrderedDict', 'deque'] # 'deque', 'defaultdict',
			πF.SetLineno(13)
			πTemp001 = make([]*πg.Object, 4)
			πTemp001[0] = ßCounter.ToObject()
			πTemp001[1] = ßnamedtuple.ToObject()
			πTemp001[2] = ßOrderedDict.ToObject()
			πTemp001[3] = ßdeque.ToObject()
			πTemp002 = πg.NewList(πTemp001...).ToObject()
			if πE = πF.Globals().SetItem(πF, ß__all__.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 16: import _abcoll
			πF.SetLineno(16)
			if πTemp001, πE = πg.ImportModule(πF, "_abcoll"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ß_abcoll.ToObject(), πTemp002); πE != nil {
				continue
			}
			if πTemp003, πE = πg.ResolveGlobal(πF, ß_abcoll); πE != nil {
				continue
			}
			if πTemp004, πE = πg.GetAttr(πF, πTemp003, ß__all__, nil); πE != nil {
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
				if πE = πF.Globals().SetItem(πF, ßname.ToObject(), πTemp003); πE != nil {
					continue
				}
			}
			if πE != nil || !πTemp006 {
				continue
			}
			πF.PushCheckpoint(1)            
			// line 19: globals()[name] = getattr(_abcoll, name)
			πF.SetLineno(19)
			πTemp001 = πF.MakeArgs(2)
			if πTemp003, πE = πg.ResolveGlobal(πF, ß_abcoll); πE != nil {
				continue
			}
			πTemp001[0] = πTemp003
			if πTemp003, πE = πg.ResolveGlobal(πF, ßname); πE != nil {
				continue
			}
			πTemp001[1] = πTemp003
			if πTemp003, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
				continue
			}
			if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp004); πE != nil {
				continue
			}
			if πTemp007, πE = πg.ResolveGlobal(πF, ßglobals); πE != nil {
				continue
			}
			if πTemp008, πE = πTemp007.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πTemp009, πE = πg.ResolveGlobal(πF, ßname); πE != nil {
				continue
			}
			πTemp007 = πTemp009
			if πE = πg.SetItem(πF, πTemp008, πTemp007, πTemp003); πE != nil {
				continue
			}
			continue
		Label2:
			if πE != nil || πR != nil {
				continue
			}
		Label3:
			// line 21: import _collections
			πF.SetLineno(21)
			if πTemp001, πE = πg.ImportModule(πF, "_collections"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ß_collections.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 22: deque = _collections.deque
			πF.SetLineno(22)
			if πTemp002, πE = πg.ResolveGlobal(πF, ß_collections); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßdeque, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßdeque.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 24: import operator
			πF.SetLineno(24)
			if πTemp001, πE = πg.ImportModule(πF, "operator"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ßoperator.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 25: _itemgetter = operator.itemgetter
			πF.SetLineno(25)
			if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßitemgetter, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_itemgetter.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 26: _eq = operator.eq
			πF.SetLineno(26)
			if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßeq, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_eq.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 27: import keyword
			πF.SetLineno(27)
			if πTemp001, πE = πg.ImportModule(πF, "keyword"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ßkeyword.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 28: _iskeyword = keyword.iskeyword
			πF.SetLineno(28)
			if πTemp002, πE = πg.ResolveGlobal(πF, ßkeyword); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßiskeyword, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_iskeyword.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 29: import sys as _sys
			πF.SetLineno(29)
			if πTemp001, πE = πg.ImportModule(πF, "sys"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ß_sys.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 30: import heapq as _heapq
			πF.SetLineno(30)
			if πTemp001, πE = πg.ImportModule(πF, "heapq"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ß_heapq.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 31: import itertools
			πF.SetLineno(31)
			if πTemp001, πE = πg.ImportModule(πF, "itertools"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ßitertools.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 32: _repeat = itertools.repeat
			πF.SetLineno(32)
			if πTemp002, πE = πg.ResolveGlobal(πF, ßitertools); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßrepeat, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_repeat.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 33: _chain = itertools.chain
			πF.SetLineno(33)
			if πTemp002, πE = πg.ResolveGlobal(πF, ßitertools); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßchain, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_chain.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 34: _starmap = itertools.starmap
			πF.SetLineno(34)
			if πTemp002, πE = πg.ResolveGlobal(πF, ßitertools); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßstarmap, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_starmap.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 35: _imap = itertools.imap
			πF.SetLineno(35)
			if πTemp002, πE = πg.ResolveGlobal(πF, ßitertools); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßimap, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_imap.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 38: import thread
			πF.SetLineno(38)
			if πTemp001, πE = πg.ImportModule(πF, "thread"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ßthread.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 39: _get_ident = thread.get_ident
			πF.SetLineno(39)
			if πTemp002, πE = πg.ResolveGlobal(πF, ßthread); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßget_ident, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_get_ident.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 48: class OrderedDict(dict):
			πF.SetLineno(48)
			πTemp001 = make([]*πg.Object, 1)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
				continue
			}
			πTemp001[0] = πTemp004
			πTemp010 = πg.NewDict()
			if πTemp002, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp010.SetItem(πF, ß__module__.ToObject(), πTemp002); πE != nil {
				continue
			}
			_, πE = πg.NewCode("OrderedDict", "/usr/lib/python2.7/collections.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
				var πTemp019 *πg.Dict
				_ = πTemp019
				var πTemp020 *πg.Object
				_ = πTemp020
				var πTemp021 *πg.Object
				_ = πTemp021
				var πTemp022 *πg.Object
				_ = πTemp022
				var πTemp023 *πg.Object
				_ = πTemp023
				var πTemp024 []*πg.Object
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 49: 'Dictionary that remembers insertion order'
					πF.SetLineno(49)
					// line 49: 'Dictionary that remembers insertion order'
					πF.SetLineno(49)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Dictionary that remembers insertion order").ToObject()); πE != nil {
						continue
					}
					// line 60: def __init__(*args, **kwds):
					πF.SetLineno(60)
					πTemp002 = make([]πg.Param, 0)
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/collections.py", πTemp002, πg.CodeFlagVarArg | πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µargs *πg.Object = πArgs[0]; _ = µargs
						var µkwds *πg.Object = πArgs[1]; _ = µkwds
						var µself *πg.Object = πg.UnboundLocal; _ = µself
						var µroot *πg.Object = πg.UnboundLocal; _ = µroot
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
						var πTemp006 []*πg.Object
						_ = πTemp006
						var πTemp007 *πg.BaseException
						_ = πTemp007
						var πTemp008 *πg.Traceback
						_ = πTemp008
						var πTemp009 *πg.Dict
						_ = πTemp009
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 6: goto Label6
							default: panic("unexpected function state")
							}
							// line 61: '''Initialize an ordered dictionary.  The signature is the same as
							πF.SetLineno(61)
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, µargs); πE != nil {
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
							// line 66: if not args:
							πF.SetLineno(66)
						Label1:
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewStr("descriptor '__init__' of 'OrderedDict' object needs an argument").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 67: raise TypeError("descriptor '__init__' of 'OrderedDict' object "
							πF.SetLineno(67)
							πE = πF.Raise(πTemp004, nil, nil)
							continue
							goto Label2
						Label2:
							// line 69: self = args[0]
							πF.SetLineno(69)
							πTemp001 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µargs, πTemp001); πE != nil {
								continue
							}
							µself = πTemp004
							// line 70: args = args[1:]
							πF.SetLineno(70)
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µargs, πTemp001); πE != nil {
								continue
							}
							µargs = πTemp004
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							πTemp003[0] = µargs
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp001, πE = πg.GT(πF, πTemp005, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label3
							}
							goto Label4
							// line 71: if len(args) > 1:
							πF.SetLineno(71)
						Label3:
							πTemp003 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							πTemp006[0] = µargs
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("expected at most 1 arguments, got %d").ToObject(), πTemp005); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 72: raise TypeError('expected at most 1 arguments, got %d' % len(args))
							πF.SetLineno(72)
							πE = πF.Raise(πTemp004, nil, nil)
							continue
							goto Label4
						Label4:
							// line 73: try:
							πF.SetLineno(73)
							πF.PushCheckpoint(6)
							// line 74: self.__root
							πF.SetLineno(74)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß__root, nil); πE != nil {
								continue
							}
							πF.PopCheckpoint()
							goto Label5
						Label6:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp007, πTemp008 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßAttributeError); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsInstance(πF, πTemp007.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label7
							}
							πE = πF.Raise(πTemp007.ToObject(), nil, πTemp008.ToObject())
							continue
							// line 75: except AttributeError:
							πF.SetLineno(75)
						Label7:
							// line 76: self.__root = root = []                     # sentinel node
							πF.SetLineno(76)
							πTemp003 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp003...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß__root, πTemp004); πE != nil {
								continue
							}
							µroot = πTemp001
							// line 77: root[:] = [root, root, None]
							πF.SetLineno(77)
							πTemp003 = make([]*πg.Object, 3)
							if πE = πg.CheckLocal(πF, µroot, "root"); πE != nil {
								continue
							}
							πTemp003[0] = µroot
							if πE = πg.CheckLocal(πF, µroot, "root"); πE != nil {
								continue
							}
							πTemp003[1] = µroot
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003[2] = πTemp001
							πTemp001 = πg.NewList(πTemp003...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µroot, "root"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.SetItem(πF, µroot, πTemp005, πTemp004); πE != nil {
								continue
							}
							// line 78: self.__map = {}
							πF.SetLineno(78)
							πTemp009 = πg.NewDict()
							πTemp001 = πTemp009.ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß__map, πTemp004); πE != nil {
								continue
							}
							πF.RestoreExc(nil, nil)
							goto Label5
						Label5:
							// line 79: self.__update(*args, **kwds)
							πF.SetLineno(79)
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µkwds, "kwds"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß__update, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Invoke(πF, πTemp001, nil, µargs, nil, µkwds); πE != nil {
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
					// line 61: '''Initialize an ordered dictionary.  The signature is the same as
					πF.SetLineno(61)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewStr("Initialize an ordered dictionary.  The signature is the same as\n        regular dictionaries, but keyword arguments are not recommended because\n        their insertion order is arbitrary.\n\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ß__init__); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp004, ß__doc__, πTemp003); πE != nil {
						continue
					}
					// line 81: def __setitem__(self, key, value, dict_setitem=dict.__setitem__):
					πF.SetLineno(81)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "key", Def: nil}
					πTemp002[2] = πg.Param{Name: "value", Def: nil}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßdict); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp004, ß__setitem__, nil); πE != nil {
						continue
					}
					πTemp002[3] = πg.Param{Name: "dict_setitem", Def: πTemp005}
					πTemp003 = πg.NewFunction(πg.NewCode("__setitem__", "/usr/lib/python2.7/collections.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µkey *πg.Object = πArgs[1]; _ = µkey
						var µvalue *πg.Object = πArgs[2]; _ = µvalue
						var µdict_setitem *πg.Object = πArgs[3]; _ = µdict_setitem
						var µroot *πg.Object = πg.UnboundLocal; _ = µroot
						var µlast *πg.Object = πg.UnboundLocal; _ = µlast
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 82: 'od.__setitem__(i, y) <==> od[i]=y'
							πF.SetLineno(82)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Contains(πF, µself, µkey); πE != nil {
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
							// line 85: if key not in self:
							πF.SetLineno(85)
						Label1:
							// line 86: root = self.__root
							πF.SetLineno(86)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß__root, nil); πE != nil {
								continue
							}
							µroot = πTemp001
							// line 87: last = root[0]
							πF.SetLineno(87)
							πTemp001 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µroot, "root"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µroot, πTemp001); πE != nil {
								continue
							}
							µlast = πTemp003
							// line 88: last[1] = root[0] = self.__map[key] = [last, root, key]
							πF.SetLineno(88)
							πTemp004 = make([]*πg.Object, 3)
							if πE = πg.CheckLocal(πF, µlast, "last"); πE != nil {
								continue
							}
							πTemp004[0] = µlast
							if πE = πg.CheckLocal(πF, µroot, "root"); πE != nil {
								continue
							}
							πTemp004[1] = µroot
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp004[2] = µkey
							πTemp001 = πg.NewList(πTemp004...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µlast, "last"); πE != nil {
								continue
							}
							πTemp005 = πg.NewInt(1).ToObject()
							if πE = πg.SetItem(πF, µlast, πTemp005, πTemp003); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µroot, "root"); πE != nil {
								continue
							}
							πTemp005 = πg.NewInt(0).ToObject()
							if πE = πg.SetItem(πF, µroot, πTemp005, πTemp003); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ß__map, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp006 = µkey
							if πE = πg.SetItem(πF, πTemp005, πTemp006, πTemp003); πE != nil {
								continue
							}
							goto Label2
						Label2:
							// line 89: return dict_setitem(self, key, value)
							πF.SetLineno(89)
							πTemp004 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp004[0] = µself
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp004[1] = µkey
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp004[2] = µvalue
							if πE = πg.CheckLocal(πF, µdict_setitem, "dict_setitem"); πE != nil {
								continue
							}
							if πTemp001, πE = µdict_setitem.Call(πF, πTemp004, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__setitem__.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 82: 'od.__setitem__(i, y) <==> od[i]=y'
					πF.SetLineno(82)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πg.NewStr("od.__setitem__(i, y) <==> od[i]=y").ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ß__setitem__); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp005, ß__doc__, πTemp004); πE != nil {
						continue
					}
					// line 91: def __delitem__(self, key, dict_delitem=dict.__delitem__):
					πF.SetLineno(91)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "key", Def: nil}
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßdict); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp005, ß__delitem__, nil); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "dict_delitem", Def: πTemp006}
					πTemp004 = πg.NewFunction(πg.NewCode("__delitem__", "/usr/lib/python2.7/collections.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µkey *πg.Object = πArgs[1]; _ = µkey
						var µdict_delitem *πg.Object = πArgs[2]; _ = µdict_delitem
						var µlink_prev *πg.Object = πg.UnboundLocal; _ = µlink_prev
						var µlink_next *πg.Object = πg.UnboundLocal; _ = µlink_next
						var µ_ *πg.Object = πg.UnboundLocal; _ = µ_
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
							// line 92: 'od.__delitem__(y) <==> del od[y]'
							πF.SetLineno(92)
							// line 95: dict_delitem(self, key)
							πF.SetLineno(95)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp001[1] = µkey
							if πE = πg.CheckLocal(πF, µdict_delitem, "dict_delitem"); πE != nil {
								continue
							}
							if πTemp002, πE = µdict_delitem.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 96: link_prev, link_next, _ = self.__map.pop(key)
							πF.SetLineno(96)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp001[0] = µkey
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß__map, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßpop, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}}}, πTemp002); πE != nil {
								continue
							}
							µlink_prev = πTemp003
							µlink_next = πTemp004
							µ_ = πTemp005
							// line 97: link_prev[1] = link_next                        # update link_prev[NEXT]
							πF.SetLineno(97)
							if πE = πg.CheckLocal(πF, µlink_next, "link_next"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µlink_next); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µlink_prev, "link_prev"); πE != nil {
								continue
							}
							πTemp003 = πg.NewInt(1).ToObject()
							if πE = πg.SetItem(πF, µlink_prev, πTemp003, πTemp002); πE != nil {
								continue
							}
							// line 98: link_next[0] = link_prev                        # update link_next[PREV]
							πF.SetLineno(98)
							if πE = πg.CheckLocal(πF, µlink_prev, "link_prev"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µlink_prev); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µlink_next, "link_next"); πE != nil {
								continue
							}
							πTemp003 = πg.NewInt(0).ToObject()
							if πE = πg.SetItem(πF, µlink_next, πTemp003, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__delitem__.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 92: 'od.__delitem__(y) <==> del od[y]'
					πF.SetLineno(92)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πg.NewStr("od.__delitem__(y) <==> del od[y]").ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ß__delitem__); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp006, ß__doc__, πTemp005); πE != nil {
						continue
					}
					// line 100: def __iter__(self):
					πF.SetLineno(100)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("__iter__", "/usr/lib/python2.7/collections.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µroot *πg.Object = πg.UnboundLocal; _ = µroot
						var µcurr *πg.Object = πg.UnboundLocal; _ = µcurr
						var πTemp001 *πg.Object
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
						return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
							for ; πF.State() >= 0; πF.PopCheckpoint() {
								switch πF.State() {
								case 0:
								case 1: goto Label1
								case 2: goto Label2
								case 4: goto Label4
								default: panic("unexpected function state")
								}
								// line 101: 'od.__iter__() <==> iter(od)'
								πF.SetLineno(101)
								// line 103: root = self.__root
								πF.SetLineno(103)
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.GetAttr(πF, µself, ß__root, nil); πE != nil {
									continue
								}
								µroot = πTemp001
								// line 104: curr = root[1]                                  # start at the first node
								πF.SetLineno(104)
								πTemp001 = πg.NewInt(1).ToObject()
								if πE = πg.CheckLocal(πF, µroot, "root"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.GetItem(πF, µroot, πTemp001); πE != nil {
									continue
								}
								µcurr = πTemp002
								// line 105: while curr is not root:
								πF.SetLineno(105)
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
								if πE = πg.CheckLocal(πF, µcurr, "curr"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µroot, "root"); πE != nil {
									continue
								}
								πTemp001 = πg.GetBool(µcurr != µroot).ToObject()
								if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
									continue
								}
								if πE != nil || !πTemp004 {
									continue
								}
								πF.PushCheckpoint(1)            
								// line 106: yield curr[2]                               # yield the curr[KEY]
								πF.SetLineno(106)
								πTemp001 = πg.NewInt(2).ToObject()
								if πE = πg.CheckLocal(πF, µcurr, "curr"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.GetItem(πF, µcurr, πTemp001); πE != nil {
									continue
								}
								πF.PushCheckpoint(4)
								return πTemp002, nil
							Label4:
								πTemp001 = πSent
								// line 107: curr = curr[1]                              # move to next node
								πF.SetLineno(107)
								πTemp001 = πg.NewInt(1).ToObject()
								if πE = πg.CheckLocal(πF, µcurr, "curr"); πE != nil {
									continue
								}
								if πTemp005, πE = πg.GetItem(πF, µcurr, πTemp001); πE != nil {
									continue
								}
								µcurr = πTemp005
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
					if πE = πClass.SetItem(πF, ß__iter__.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 101: 'od.__iter__() <==> iter(od)'
					πF.SetLineno(101)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, πg.NewStr("od.__iter__() <==> iter(od)").ToObject()); πE != nil {
						continue
					}
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ß__iter__); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp007, ß__doc__, πTemp006); πE != nil {
						continue
					}
					// line 109: def __reversed__(self):
					πF.SetLineno(109)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("__reversed__", "/usr/lib/python2.7/collections.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µroot *πg.Object = πg.UnboundLocal; _ = µroot
						var µcurr *πg.Object = πg.UnboundLocal; _ = µcurr
						var πTemp001 *πg.Object
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
						return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
							for ; πF.State() >= 0; πF.PopCheckpoint() {
								switch πF.State() {
								case 0:
								case 1: goto Label1
								case 2: goto Label2
								case 4: goto Label4
								default: panic("unexpected function state")
								}
								// line 110: 'od.__reversed__() <==> reversed(od)'
								πF.SetLineno(110)
								// line 112: root = self.__root
								πF.SetLineno(112)
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.GetAttr(πF, µself, ß__root, nil); πE != nil {
									continue
								}
								µroot = πTemp001
								// line 113: curr = root[0]                                  # start at the last node
								πF.SetLineno(113)
								πTemp001 = πg.NewInt(0).ToObject()
								if πE = πg.CheckLocal(πF, µroot, "root"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.GetItem(πF, µroot, πTemp001); πE != nil {
									continue
								}
								µcurr = πTemp002
								// line 114: while curr is not root:
								πF.SetLineno(114)
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
								if πE = πg.CheckLocal(πF, µcurr, "curr"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µroot, "root"); πE != nil {
									continue
								}
								πTemp001 = πg.GetBool(µcurr != µroot).ToObject()
								if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
									continue
								}
								if πE != nil || !πTemp004 {
									continue
								}
								πF.PushCheckpoint(1)            
								// line 115: yield curr[2]                               # yield the curr[KEY]
								πF.SetLineno(115)
								πTemp001 = πg.NewInt(2).ToObject()
								if πE = πg.CheckLocal(πF, µcurr, "curr"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.GetItem(πF, µcurr, πTemp001); πE != nil {
									continue
								}
								πF.PushCheckpoint(4)
								return πTemp002, nil
							Label4:
								πTemp001 = πSent
								// line 116: curr = curr[0]                              # move to previous node
								πF.SetLineno(116)
								πTemp001 = πg.NewInt(0).ToObject()
								if πE = πg.CheckLocal(πF, µcurr, "curr"); πE != nil {
									continue
								}
								if πTemp005, πE = πg.GetItem(πF, µcurr, πTemp001); πE != nil {
									continue
								}
								µcurr = πTemp005
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
					if πE = πClass.SetItem(πF, ß__reversed__.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 110: 'od.__reversed__() <==> reversed(od)'
					πF.SetLineno(110)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp007}, πg.NewStr("od.__reversed__() <==> reversed(od)").ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ß__reversed__); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp008, ß__doc__, πTemp007); πE != nil {
						continue
					}
					// line 118: def clear(self):
					πF.SetLineno(118)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("clear", "/usr/lib/python2.7/collections.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µroot *πg.Object = πg.UnboundLocal; _ = µroot
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
							// line 119: 'od.clear() -> None.  Remove all items from od.'
							πF.SetLineno(119)
							// line 120: root = self.__root
							πF.SetLineno(120)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß__root, nil); πE != nil {
								continue
							}
							µroot = πTemp001
							// line 121: root[:] = [root, root, None]
							πF.SetLineno(121)
							πTemp002 = make([]*πg.Object, 3)
							if πE = πg.CheckLocal(πF, µroot, "root"); πE != nil {
								continue
							}
							πTemp002[0] = µroot
							if πE = πg.CheckLocal(πF, µroot, "root"); πE != nil {
								continue
							}
							πTemp002[1] = µroot
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002[2] = πTemp001
							πTemp001 = πg.NewList(πTemp002...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µroot, "root"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.SetItem(πF, µroot, πTemp004, πTemp003); πE != nil {
								continue
							}
							// line 122: self.__map.clear()
							πF.SetLineno(122)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß__map, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßclear, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 123: dict.clear(self)
							πF.SetLineno(123)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp002[0] = µself
							if πTemp001, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßclear, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßclear.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 119: 'od.clear() -> None.  Remove all items from od.'
					πF.SetLineno(119)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp008}, πg.NewStr("od.clear() -> None.  Remove all items from od.").ToObject()); πE != nil {
						continue
					}
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßclear); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp009, ß__doc__, πTemp008); πE != nil {
						continue
					}
					// line 127: def keys(self):
					πF.SetLineno(127)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("keys", "/usr/lib/python2.7/collections.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 128: 'od.keys() -> list of keys in od'
							πF.SetLineno(128)
							// line 129: return list(self)
							πF.SetLineno(129)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
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
					if πE = πClass.SetItem(πF, ßkeys.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 128: 'od.keys() -> list of keys in od'
					πF.SetLineno(128)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp009}, πg.NewStr("od.keys() -> list of keys in od").ToObject()); πE != nil {
						continue
					}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßkeys); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp010, ß__doc__, πTemp009); πE != nil {
						continue
					}
					// line 131: def values(self):
					πF.SetLineno(131)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("values", "/usr/lib/python2.7/collections.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []πg.Param
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
							// line 132: 'od.values() -> list of values in od'
							πF.SetLineno(132)
							// line 133: return [self[key] for key in self]
							πF.SetLineno(133)
							πTemp003 = make([]πg.Param, 0)
							πTemp002 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/collections.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µkey *πg.Object = πg.UnboundLocal; _ = µkey
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
										if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
											continue
										}
										if πTemp001, πE = πg.Iter(πF, µself); πE != nil {
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
											µkey = πTemp004
										}
										if πE != nil || !πTemp003 {
											continue
										}
										πF.PushCheckpoint(1)            
										// line 133: return [self[key] for key in self]
										πF.SetLineno(133)
										if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
											continue
										}
										πTemp004 = µkey
										if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
											continue
										}
										if πTemp005, πE = πg.GetItem(πF, µself, πTemp004); πE != nil {
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
							if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.ListType.Call(πF, πg.Args{πTemp004}, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßvalues.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 132: 'od.values() -> list of values in od'
					πF.SetLineno(132)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp010}, πg.NewStr("od.values() -> list of values in od").ToObject()); πE != nil {
						continue
					}
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ßvalues); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp011, ß__doc__, πTemp010); πE != nil {
						continue
					}
					// line 135: def items(self):
					πF.SetLineno(135)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp010 = πg.NewFunction(πg.NewCode("items", "/usr/lib/python2.7/collections.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []πg.Param
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
							// line 136: 'od.items() -> list of (key, value) pairs in od'
							πF.SetLineno(136)
							// line 137: return [(key, self[key]) for key in self]
							πF.SetLineno(137)
							πTemp003 = make([]πg.Param, 0)
							πTemp002 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/collections.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µkey *πg.Object = πg.UnboundLocal; _ = µkey
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
										if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
											continue
										}
										if πTemp001, πE = πg.Iter(πF, µself); πE != nil {
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
											µkey = πTemp004
										}
										if πE != nil || !πTemp003 {
											continue
										}
										πF.PushCheckpoint(1)            
										// line 137: return [(key, self[key]) for key in self]
										πF.SetLineno(137)
										if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
											continue
										}
										if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
											continue
										}
										πTemp005 = µkey
										if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
											continue
										}
										if πTemp006, πE = πg.GetItem(πF, µself, πTemp005); πE != nil {
											continue
										}
										πTemp004 = πg.NewTuple2(µkey, πTemp006).ToObject()
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
							if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.ListType.Call(πF, πg.Args{πTemp004}, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßitems.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 136: 'od.items() -> list of (key, value) pairs in od'
					πF.SetLineno(136)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πg.NewStr("od.items() -> list of (key, value) pairs in od").ToObject()); πE != nil {
						continue
					}
					if πTemp012, πE = πg.ResolveClass(πF, πClass, nil, ßitems); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp012, ß__doc__, πTemp011); πE != nil {
						continue
					}
					// line 139: def iterkeys(self):
					πF.SetLineno(139)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp011 = πg.NewFunction(πg.NewCode("iterkeys", "/usr/lib/python2.7/collections.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 140: 'od.iterkeys() -> an iterator over the keys in od'
							πF.SetLineno(140)
							// line 141: return iter(self)
							πF.SetLineno(141)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πTemp002, πE = πg.ResolveGlobal(πF, ßiter); πE != nil {
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
					if πE = πClass.SetItem(πF, ßiterkeys.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 140: 'od.iterkeys() -> an iterator over the keys in od'
					πF.SetLineno(140)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp012}, πg.NewStr("od.iterkeys() -> an iterator over the keys in od").ToObject()); πE != nil {
						continue
					}
					if πTemp013, πE = πg.ResolveClass(πF, πClass, nil, ßiterkeys); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp013, ß__doc__, πTemp012); πE != nil {
						continue
					}
					// line 143: def itervalues(self):
					πF.SetLineno(143)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp012 = πg.NewFunction(πg.NewCode("itervalues", "/usr/lib/python2.7/collections.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µk *πg.Object = πg.UnboundLocal; _ = µk
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
								// line 144: 'od.itervalues -> an iterator over the values in od'
								πF.SetLineno(144)
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.Iter(πF, µself); πE != nil {
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
									µk = πTemp004
								}
								if πE != nil || !πTemp003 {
									continue
								}
								πF.PushCheckpoint(1)            
								// line 146: yield self[k]
								πF.SetLineno(146)
								if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
									continue
								}
								πTemp004 = µk
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp005, πE = πg.GetItem(πF, µself, πTemp004); πE != nil {
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
					if πE = πClass.SetItem(πF, ßitervalues.ToObject(), πTemp012); πE != nil {
						continue
					}
					// line 144: 'od.itervalues -> an iterator over the values in od'
					πF.SetLineno(144)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp013}, πg.NewStr("od.itervalues -> an iterator over the values in od").ToObject()); πE != nil {
						continue
					}
					if πTemp014, πE = πg.ResolveClass(πF, πClass, nil, ßitervalues); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp014, ß__doc__, πTemp013); πE != nil {
						continue
					}
					// line 148: def iteritems(self):
					πF.SetLineno(148)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp013 = πg.NewFunction(πg.NewCode("iteritems", "/usr/lib/python2.7/collections.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µk *πg.Object = πg.UnboundLocal; _ = µk
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
								// line 149: 'od.iteritems -> an iterator over the (key, value) pairs in od'
								πF.SetLineno(149)
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.Iter(πF, µself); πE != nil {
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
									µk = πTemp004
								}
								if πE != nil || !πTemp003 {
									continue
								}
								πF.PushCheckpoint(1)            
								// line 151: yield (k, self[k])
								πF.SetLineno(151)
								if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
									continue
								}
								πTemp005 = µk
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp006, πE = πg.GetItem(πF, µself, πTemp005); πE != nil {
									continue
								}
								πTemp004 = πg.NewTuple2(µk, πTemp006).ToObject()
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
					if πE = πClass.SetItem(πF, ßiteritems.ToObject(), πTemp013); πE != nil {
						continue
					}
					// line 149: 'od.iteritems -> an iterator over the (key, value) pairs in od'
					πF.SetLineno(149)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp014}, πg.NewStr("od.iteritems -> an iterator over the (key, value) pairs in od").ToObject()); πE != nil {
						continue
					}
					if πTemp015, πE = πg.ResolveClass(πF, πClass, nil, ßiteritems); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp015, ß__doc__, πTemp014); πE != nil {
						continue
					}
					// line 153: update = MutableMapping.update
					πF.SetLineno(153)
					if πTemp014, πE = πg.ResolveClass(πF, πClass, nil, ßMutableMapping); πE != nil {
						continue
					}
					if πTemp015, πE = πg.GetAttr(πF, πTemp014, ßupdate, nil); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßupdate.ToObject(), πTemp015); πE != nil {
						continue
					}
					// line 155: __update = update # let subclasses override update without breaking __init__
					πF.SetLineno(155)
					if πTemp014, πE = πg.ResolveClass(πF, πClass, nil, ßupdate); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ß__update.ToObject(), πTemp014); πE != nil {
						continue
					}
					// line 157: __marker = object()
					πF.SetLineno(157)
					if πTemp014, πE = πg.ResolveClass(πF, πClass, nil, ßobject); πE != nil {
						continue
					}
					if πTemp015, πE = πTemp014.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ß__marker.ToObject(), πTemp015); πE != nil {
						continue
					}
					// line 159: def pop(self, key, default=__marker):
					πF.SetLineno(159)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "key", Def: nil}
					if πTemp015, πE = πg.ResolveClass(πF, πClass, nil, ß__marker); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "default", Def: πTemp015}
					πTemp014 = πg.NewFunction(πg.NewCode("pop", "/usr/lib/python2.7/collections.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µkey *πg.Object = πArgs[1]; _ = µkey
						var µdefault *πg.Object = πArgs[2]; _ = µdefault
						var µresult *πg.Object = πg.UnboundLocal; _ = µresult
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
							// line 160: '''od.pop(k[,d]) -> v, remove specified key and return the corresponding
							πF.SetLineno(160)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Contains(πF, µself, µkey); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp002).ToObject()
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							goto Label2
							// line 165: if key in self:
							πF.SetLineno(165)
						Label1:
							// line 166: result = self[key]
							πF.SetLineno(166)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp001 = µkey
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µself, πTemp001); πE != nil {
								continue
							}
							µresult = πTemp003
							// line 167: del self[key]
							πF.SetLineno(167)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp001 = µkey
							if πE = πg.DelItem(πF, µself, πTemp001); πE != nil {
								continue
							}
							// line 168: return result
							πF.SetLineno(168)
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							πR = µresult
							continue
							goto Label2
						Label2:
							if πE = πg.CheckLocal(πF, µdefault, "default"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß__marker, nil); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µdefault == πTemp003).ToObject()
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label3
							}
							goto Label4
							// line 169: if default is self.__marker:
							πF.SetLineno(169)
						Label3:
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp004[0] = µkey
							if πTemp001, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 170: raise KeyError(key)
							πF.SetLineno(170)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label4
						Label4:
							// line 171: return default
							πF.SetLineno(171)
							if πE = πg.CheckLocal(πF, µdefault, "default"); πE != nil {
								continue
							}
							πR = µdefault
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßpop.ToObject(), πTemp014); πE != nil {
						continue
					}
					// line 160: '''od.pop(k[,d]) -> v, remove specified key and return the corresponding
					πF.SetLineno(160)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp015}, πg.NewStr("od.pop(k[,d]) -> v, remove specified key and return the corresponding\n        value.  If key is not found, d is returned if given, otherwise KeyError\n        is raised.\n\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp016, πE = πg.ResolveClass(πF, πClass, nil, ßpop); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp016, ß__doc__, πTemp015); πE != nil {
						continue
					}
					// line 173: def setdefault(self, key, default=None):
					πF.SetLineno(173)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "key", Def: nil}
					if πTemp016, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "default", Def: πTemp016}
					πTemp015 = πg.NewFunction(πg.NewCode("setdefault", "/usr/lib/python2.7/collections.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µkey *πg.Object = πArgs[1]; _ = µkey
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
							// line 174: 'od.setdefault(k[,d]) -> od.get(k,d), also set od[k]=d if k not in od'
							πF.SetLineno(174)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Contains(πF, µself, µkey); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp002).ToObject()
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							goto Label2
							// line 175: if key in self:
							πF.SetLineno(175)
						Label1:
							// line 176: return self[key]
							πF.SetLineno(176)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp001 = µkey
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µself, πTemp001); πE != nil {
								continue
							}
							πR = πTemp003
							continue
							goto Label2
						Label2:
							// line 177: self[key] = default
							πF.SetLineno(177)
							if πE = πg.CheckLocal(πF, µdefault, "default"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µdefault); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp003 = µkey
							if πE = πg.SetItem(πF, µself, πTemp003, πTemp001); πE != nil {
								continue
							}
							// line 178: return default
							πF.SetLineno(178)
							if πE = πg.CheckLocal(πF, µdefault, "default"); πE != nil {
								continue
							}
							πR = µdefault
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßsetdefault.ToObject(), πTemp015); πE != nil {
						continue
					}
					// line 174: 'od.setdefault(k[,d]) -> od.get(k,d), also set od[k]=d if k not in od'
					πF.SetLineno(174)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp016}, πg.NewStr("od.setdefault(k[,d]) -> od.get(k,d), also set od[k]=d if k not in od").ToObject()); πE != nil {
						continue
					}
					if πTemp017, πE = πg.ResolveClass(πF, πClass, nil, ßsetdefault); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp017, ß__doc__, πTemp016); πE != nil {
						continue
					}
					// line 180: def popitem(self, last=True):
					πF.SetLineno(180)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					if πTemp017, πE = πg.ResolveClass(πF, πClass, nil, ßTrue); πE != nil {
						continue
					}
					πTemp002[1] = πg.Param{Name: "last", Def: πTemp017}
					πTemp016 = πg.NewFunction(πg.NewCode("popitem", "/usr/lib/python2.7/collections.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µlast *πg.Object = πArgs[1]; _ = µlast
						var µkey *πg.Object = πg.UnboundLocal; _ = µkey
						var µvalue *πg.Object = πg.UnboundLocal; _ = µvalue
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
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
							// line 181: '''od.popitem() -> (k, v), return and remove a (key, value) pair.
							πF.SetLineno(181)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, µself); πE != nil {
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
							// line 185: if not self:
							πF.SetLineno(185)
						Label1:
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewStr("dictionary is empty").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 186: raise KeyError('dictionary is empty')
							πF.SetLineno(186)
							πE = πF.Raise(πTemp004, nil, nil)
							continue
							goto Label2
						Label2:
							// line 187: key = next(reversed(self) if last else iter(self))
							πF.SetLineno(187)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µlast, "last"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, µlast); πE != nil {
								continue
							}
							if !πTemp002 {
								goto Label3
							}
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp005[0] = µself
							if πTemp004, πE = πg.ResolveGlobal(πF, ßreversed); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp001 = πTemp006
							goto Label4
						Label3:
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp005[0] = µself
							if πTemp004, πE = πg.ResolveGlobal(πF, ßiter); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp001 = πTemp006
						Label4:
							πTemp003[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnext); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µkey = πTemp004
							// line 188: value = self.pop(key)
							πF.SetLineno(188)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp003[0] = µkey
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßpop, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µvalue = πTemp004
							// line 189: return key, value
							πF.SetLineno(189)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple2(µkey, µvalue).ToObject()
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
					if πE = πClass.SetItem(πF, ßpopitem.ToObject(), πTemp016); πE != nil {
						continue
					}
					// line 181: '''od.popitem() -> (k, v), return and remove a (key, value) pair.
					πF.SetLineno(181)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp017}, πg.NewStr("od.popitem() -> (k, v), return and remove a (key, value) pair.\n        Pairs are returned in LIFO order if last is true or FIFO order if false.\n\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp018, πE = πg.ResolveClass(πF, πClass, nil, ßpopitem); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp018, ß__doc__, πTemp017); πE != nil {
						continue
					}
					// line 191: def __repr__(self, _repr_running={}):
					πF.SetLineno(191)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp019 = πg.NewDict()
					πTemp018 = πTemp019.ToObject()
					πTemp002[1] = πg.Param{Name: "_repr_running", Def: πTemp018}
					πTemp017 = πg.NewFunction(πg.NewCode("__repr__", "/usr/lib/python2.7/collections.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µ_repr_running *πg.Object = πArgs[1]; _ = µ_repr_running
						var µcall_key *πg.Object = πg.UnboundLocal; _ = µcall_key
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
							case 3: goto Label3
							default: panic("unexpected function state")
							}
							// line 192: 'od.__repr__() <==> repr(od)'
							πF.SetLineno(192)
							// line 193: call_key = id(self), _get_ident()
							πF.SetLineno(193)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp002[0] = µself
							if πTemp003, πE = πg.ResolveGlobal(πF, ßid); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp003, πE = πg.ResolveGlobal(πF, ß_get_ident); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple2(πTemp004, πTemp005).ToObject()
							µcall_key = πTemp001
							if πE = πg.CheckLocal(πF, µcall_key, "call_key"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µ_repr_running, "_repr_running"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Contains(πF, µ_repr_running, µcall_key); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp006).ToObject()
							if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label1
							}
							goto Label2
							// line 194: if call_key in _repr_running:
							πF.SetLineno(194)
						Label1:
							// line 195: return '...'
							πF.SetLineno(195)
							πR = πg.NewStr("...").ToObject()
							continue
							goto Label2
						Label2:
							// line 196: _repr_running[call_key] = 1
							πF.SetLineno(196)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µ_repr_running, "_repr_running"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcall_key, "call_key"); πE != nil {
								continue
							}
							πTemp003 = µcall_key
							if πE = πg.SetItem(πF, µ_repr_running, πTemp003, πTemp001); πE != nil {
								continue
							}
							// line 197: try:
							πF.SetLineno(197)
							πF.PushCheckpoint(3)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, µself); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp006).ToObject()
							if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label4
							}
							goto Label5
							// line 198: if not self:
							πF.SetLineno(198)
						Label4:
							// line 199: return '%s()' % (self.__class__.__name__,)
							πF.SetLineno(199)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß__class__, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ß__name__, nil); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple1(πTemp005).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("%s()").ToObject(), πTemp003); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label5
						Label5:
							// line 200: return '%s(%r)' % (self.__class__.__name__, self.items())
							πF.SetLineno(200)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß__class__, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ß__name__, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßitems, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple2(πTemp005, πTemp007).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("%s(%r)").ToObject(), πTemp003); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							πF.PopCheckpoint()
							goto Label3
						Label3:
							πTemp008, πTemp009 = πF.RestoreExc(nil, nil)
							// line 202: del _repr_running[call_key]
							πF.SetLineno(202)
							if πE = πg.CheckLocal(πF, µ_repr_running, "_repr_running"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcall_key, "call_key"); πE != nil {
								continue
							}
							πTemp001 = µcall_key
							if πE = πg.DelItem(πF, µ_repr_running, πTemp001); πE != nil {
								continue
							}
							if πTemp008 != nil {
								πE = πF.Raise(πTemp008.ToObject(), nil, πTemp009.ToObject())
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
					if πE = πClass.SetItem(πF, ß__repr__.ToObject(), πTemp017); πE != nil {
						continue
					}
					// line 192: 'od.__repr__() <==> repr(od)'
					πF.SetLineno(192)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp018}, πg.NewStr("od.__repr__() <==> repr(od)").ToObject()); πE != nil {
						continue
					}
					if πTemp020, πE = πg.ResolveClass(πF, πClass, nil, ß__repr__); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp020, ß__doc__, πTemp018); πE != nil {
						continue
					}
					// line 204: def __reduce__(self):
					πF.SetLineno(204)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp018 = πg.NewFunction(πg.NewCode("__reduce__", "/usr/lib/python2.7/collections.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µitems *πg.Object = πg.UnboundLocal; _ = µitems
						var µinst_dict *πg.Object = πg.UnboundLocal; _ = µinst_dict
						var µk *πg.Object = πg.UnboundLocal; _ = µk
						var πTemp001 *πg.Object
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
						var πTemp007 bool
						_ = πTemp007
						var πTemp008 bool
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
							// line 205: 'Return state information for pickling'
							πF.SetLineno(205)
							// line 206: items = [[k, self[k]] for k in self]
							πF.SetLineno(206)
							πTemp003 = make([]πg.Param, 0)
							πTemp002 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/collections.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µk *πg.Object = πg.UnboundLocal; _ = µk
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
										if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
											continue
										}
										if πTemp001, πE = πg.Iter(πF, µself); πE != nil {
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
											µk = πTemp004
										}
										if πE != nil || !πTemp003 {
											continue
										}
										πF.PushCheckpoint(1)            
										// line 206: items = [[k, self[k]] for k in self]
										πF.SetLineno(206)
										πTemp005 = make([]*πg.Object, 2)
										if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
											continue
										}
										πTemp005[0] = µk
										if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
											continue
										}
										πTemp004 = µk
										if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
											continue
										}
										if πTemp006, πE = πg.GetItem(πF, µself, πTemp004); πE != nil {
											continue
										}
										πTemp005[1] = πTemp006
										πTemp004 = πg.NewList(πTemp005...).ToObject()
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
							if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.ListType.Call(πF, πg.Args{πTemp004}, nil); πE != nil {
								continue
							}
							µitems = πTemp001
							// line 207: inst_dict = vars(self).copy()
							πF.SetLineno(207)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp005[0] = µself
							if πTemp001, πE = πg.ResolveGlobal(πF, ßvars); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp001, πE = πg.GetAttr(πF, πTemp004, ßcopy, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µinst_dict = πTemp004
							πTemp005 = πF.MakeArgs(1)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßOrderedDict); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp005[0] = πTemp006
							if πTemp004, πE = πg.ResolveGlobal(πF, ßvars); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp001, πE = πg.Iter(πF, πTemp006); πE != nil {
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
							if πTemp004, πE = πg.Next(πF, πTemp001); πE != nil {
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
								µk = πTemp004
							}
							if πE != nil || !πTemp008 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 209: inst_dict.pop(k, None)
							πF.SetLineno(209)
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
								continue
							}
							πTemp005[0] = µk
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp005[1] = πTemp004
							if πE = πg.CheckLocal(πF, µinst_dict, "inst_dict"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µinst_dict, ßpop, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							if πE = πg.CheckLocal(πF, µinst_dict, "inst_dict"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, µinst_dict); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label4
							}
							goto Label5
							// line 210: if inst_dict:
							πF.SetLineno(210)
						Label4:
							// line 211: return (self.__class__, (items,), inst_dict)
							πF.SetLineno(211)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß__class__, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µitems, "items"); πE != nil {
								continue
							}
							πTemp006 = πg.NewTuple1(µitems).ToObject()
							if πE = πg.CheckLocal(πF, µinst_dict, "inst_dict"); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple3(πTemp004, πTemp006, µinst_dict).ToObject()
							πR = πTemp001
							continue
							goto Label5
						Label5:
							// line 212: return self.__class__, (items,)
							πF.SetLineno(212)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß__class__, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µitems, "items"); πE != nil {
								continue
							}
							πTemp006 = πg.NewTuple1(µitems).ToObject()
							πTemp001 = πg.NewTuple2(πTemp004, πTemp006).ToObject()
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
					if πE = πClass.SetItem(πF, ß__reduce__.ToObject(), πTemp018); πE != nil {
						continue
					}
					// line 205: 'Return state information for pickling'
					πF.SetLineno(205)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp020}, πg.NewStr("Return state information for pickling").ToObject()); πE != nil {
						continue
					}
					if πTemp021, πE = πg.ResolveClass(πF, πClass, nil, ß__reduce__); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp021, ß__doc__, πTemp020); πE != nil {
						continue
					}
					// line 214: def copy(self):
					πF.SetLineno(214)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp020 = πg.NewFunction(πg.NewCode("copy", "/usr/lib/python2.7/collections.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 215: 'od.copy() -> a shallow copy of od'
							πF.SetLineno(215)
							// line 216: return self.__class__(self)
							πF.SetLineno(216)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß__class__, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßcopy.ToObject(), πTemp020); πE != nil {
						continue
					}
					// line 215: 'od.copy() -> a shallow copy of od'
					πF.SetLineno(215)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp021}, πg.NewStr("od.copy() -> a shallow copy of od").ToObject()); πE != nil {
						continue
					}
					if πTemp022, πE = πg.ResolveClass(πF, πClass, nil, ßcopy); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp022, ß__doc__, πTemp021); πE != nil {
						continue
					}
					// line 219: def fromkeys(cls, iterable, value=None):
					πF.SetLineno(219)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "cls", Def: nil}
					πTemp002[1] = πg.Param{Name: "iterable", Def: nil}
					if πTemp022, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "value", Def: πTemp022}
					πTemp021 = πg.NewFunction(πg.NewCode("fromkeys", "/usr/lib/python2.7/collections.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µcls *πg.Object = πArgs[0]; _ = µcls
						var µiterable *πg.Object = πArgs[1]; _ = µiterable
						var µvalue *πg.Object = πArgs[2]; _ = µvalue
						var µself *πg.Object = πg.UnboundLocal; _ = µself
						var µkey *πg.Object = πg.UnboundLocal; _ = µkey
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
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							// line 220: '''OD.fromkeys(S[, v]) -> New ordered dictionary with keys from S.
							πF.SetLineno(220)
							// line 224: self = cls()
							πF.SetLineno(224)
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							if πTemp001, πE = µcls.Call(πF, nil, nil); πE != nil {
								continue
							}
							µself = πTemp001
							if πE = πg.CheckLocal(πF, µiterable, "iterable"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, µiterable); πE != nil {
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
								µkey = πTemp004
							}
							if πE != nil || !πTemp003 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 226: self[key] = value
							πF.SetLineno(226)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, µvalue); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp005 = µkey
							if πE = πg.SetItem(πF, µself, πTemp005, πTemp004); πE != nil {
								continue
							}
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 227: return self
							πF.SetLineno(227)
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
					if πE = πClass.SetItem(πF, ßfromkeys.ToObject(), πTemp021); πE != nil {
						continue
					}
					// line 220: '''OD.fromkeys(S[, v]) -> New ordered dictionary with keys from S.
					πF.SetLineno(220)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp022}, πg.NewStr("OD.fromkeys(S[, v]) -> New ordered dictionary with keys from S.\n        If not specified, the value defaults to None.\n\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp023, πE = πg.ResolveClass(πF, πClass, nil, ßfromkeys); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp023, ß__doc__, πTemp022); πE != nil {
						continue
					}
					// line 218: @classmethod
					πF.SetLineno(218)
					πTemp024 = πF.MakeArgs(1)
					if πTemp022, πE = πg.ResolveClass(πF, πClass, nil, ßfromkeys); πE != nil {
						continue
					}
					πTemp024[0] = πTemp022
					if πTemp022, πE = πg.ResolveClass(πF, πClass, nil, ßclassmethod); πE != nil {
						continue
					}
					if πTemp023, πE = πTemp022.Call(πF, πTemp024, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp024)
					if πE = πClass.SetItem(πF, ßfromkeys.ToObject(), πTemp023); πE != nil {
						continue
					}
					// line 229: def __eq__(self, other):
					πF.SetLineno(229)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "other", Def: nil}
					πTemp022 = πg.NewFunction(πg.NewCode("__eq__", "/usr/lib/python2.7/collections.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
						var πTemp006 []*πg.Object
						_ = πTemp006
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 230: '''od.__eq__(y) <==> od==y.  Comparison to another OD is order-sensitive
							πF.SetLineno(230)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							if πTemp002, πE = πg.ResolveGlobal(πF, ßOrderedDict); πE != nil {
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
							// line 234: if isinstance(other, OrderedDict):
							πF.SetLineno(234)
						Label1:
							// line 235: return dict.__eq__(self, other) and all(_imap(_eq, self, other))
							πF.SetLineno(235)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[1] = µother
							if πTemp003, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ß__eq__, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp002 = πTemp003
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp004 {
								goto Label3
							}
							πTemp001 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(3)
							if πTemp003, πE = πg.ResolveGlobal(πF, ß_eq); πE != nil {
								continue
							}
							πTemp006[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp006[1] = µself
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp006[2] = µother
							if πTemp003, πE = πg.ResolveGlobal(πF, ß_imap); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp001[0] = πTemp005
							if πTemp003, πE = πg.ResolveGlobal(πF, ßall); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp002 = πTemp005
						Label3:
							πR = πTemp002
							continue
							goto Label2
						Label2:
							// line 236: return dict.__eq__(self, other)
							πF.SetLineno(236)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[1] = µother
							if πTemp002, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ß__eq__, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__eq__.ToObject(), πTemp022); πE != nil {
						continue
					}
					// line 230: '''od.__eq__(y) <==> od==y.  Comparison to another OD is order-sensitive
					πF.SetLineno(230)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp023}, πg.NewStr("od.__eq__(y) <==> od==y.  Comparison to another OD is order-sensitive\n        while comparison to a regular mapping is order-insensitive.\n\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp025, πE = πg.ResolveClass(πF, πClass, nil, ß__eq__); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp025, ß__doc__, πTemp023); πE != nil {
						continue
					}
					// line 238: def __ne__(self, other):
					πF.SetLineno(238)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "other", Def: nil}
					πTemp023 = πg.NewFunction(πg.NewCode("__ne__", "/usr/lib/python2.7/collections.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
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
							// line 239: 'od.__ne__(y) <==> od!=y'
							πF.SetLineno(239)
							// line 240: return not self == other
							πF.SetLineno(240)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, µself, µother); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp003).ToObject()
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
					if πE = πClass.SetItem(πF, ß__ne__.ToObject(), πTemp023); πE != nil {
						continue
					}
					// line 239: 'od.__ne__(y) <==> od!=y'
					πF.SetLineno(239)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp025}, πg.NewStr("od.__ne__(y) <==> od!=y").ToObject()); πE != nil {
						continue
					}
					if πTemp026, πE = πg.ResolveClass(πF, πClass, nil, ß__ne__); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp026, ß__doc__, πTemp025); πE != nil {
						continue
					}
					// line 244: def viewkeys(self):
					πF.SetLineno(244)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp025 = πg.NewFunction(πg.NewCode("viewkeys", "/usr/lib/python2.7/collections.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 245: "od.viewkeys() -> a set-like object providing a view on od's keys"
							πF.SetLineno(245)
							// line 246: return KeysView(self)
							πF.SetLineno(246)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πTemp002, πE = πg.ResolveGlobal(πF, ßKeysView); πE != nil {
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
					if πE = πClass.SetItem(πF, ßviewkeys.ToObject(), πTemp025); πE != nil {
						continue
					}
					// line 245: "od.viewkeys() -> a set-like object providing a view on od's keys"
					πF.SetLineno(245)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp026}, πg.NewStr("od.viewkeys() -> a set-like object providing a view on od's keys").ToObject()); πE != nil {
						continue
					}
					if πTemp027, πE = πg.ResolveClass(πF, πClass, nil, ßviewkeys); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp027, ß__doc__, πTemp026); πE != nil {
						continue
					}
					// line 248: def viewvalues(self):
					πF.SetLineno(248)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp026 = πg.NewFunction(πg.NewCode("viewvalues", "/usr/lib/python2.7/collections.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 249: "od.viewvalues() -> an object providing a view on od's values"
							πF.SetLineno(249)
							// line 250: return ValuesView(self)
							πF.SetLineno(250)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πTemp002, πE = πg.ResolveGlobal(πF, ßValuesView); πE != nil {
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
					if πE = πClass.SetItem(πF, ßviewvalues.ToObject(), πTemp026); πE != nil {
						continue
					}
					// line 249: "od.viewvalues() -> an object providing a view on od's values"
					πF.SetLineno(249)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp027}, πg.NewStr("od.viewvalues() -> an object providing a view on od's values").ToObject()); πE != nil {
						continue
					}
					if πTemp028, πE = πg.ResolveClass(πF, πClass, nil, ßviewvalues); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp028, ß__doc__, πTemp027); πE != nil {
						continue
					}
					// line 252: def viewitems(self):
					πF.SetLineno(252)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp027 = πg.NewFunction(πg.NewCode("viewitems", "/usr/lib/python2.7/collections.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 253: "od.viewitems() -> a set-like object providing a view on od's items"
							πF.SetLineno(253)
							// line 254: return ItemsView(self)
							πF.SetLineno(254)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πTemp002, πE = πg.ResolveGlobal(πF, ßItemsView); πE != nil {
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
					if πE = πClass.SetItem(πF, ßviewitems.ToObject(), πTemp027); πE != nil {
						continue
					}
					// line 253: "od.viewitems() -> a set-like object providing a view on od's items"
					πF.SetLineno(253)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp028}, πg.NewStr("od.viewitems() -> a set-like object providing a view on od's items").ToObject()); πE != nil {
						continue
					}
					if πTemp029, πE = πg.ResolveClass(πF, πClass, nil, ßviewitems); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp029, ß__doc__, πTemp028); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp010.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp004, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("OrderedDict").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp010.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßOrderedDict.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 261: _class_template = '''\
			πF.SetLineno(261)
			if πE = πF.Globals().SetItem(πF, ß_class_template.ToObject(), πg.NewStr("class {typename}(tuple):\n    '{typename}({arg_list})'\n\n    __slots__ = ()\n\n    _fields = {field_names!r}\n\n    def __new__(_cls, {arg_list}):\n        'Create new instance of {typename}({arg_list})'\n        return _tuple.__new__(_cls, ({arg_list}))\n\n    @classmethod\n    def _make(cls, iterable, new=tuple.__new__, len=len):\n        'Make a new {typename} object from a sequence or iterable'\n        result = new(cls, iterable)\n        if len(result) != {num_fields:d}:\n            raise TypeError('Expected {num_fields:d} arguments, got %d' % len(result))\n        return result\n\n    def __repr__(self):\n        'Return a nicely formatted representation string'\n        return '{typename}({repr_fmt})' % self\n\n    def _asdict(self):\n        'Return a new OrderedDict which maps field names to their values'\n        return OrderedDict(zip(self._fields, self))\n\n    def _replace(_self, **kwds):\n        'Return a new {typename} object replacing specified fields with new values'\n        result = _self._make(map(kwds.pop, {field_names!r}, _self))\n        if kwds:\n            raise ValueError('Got unexpected field names: %r' % kwds.keys())\n        return result\n\n    def __getnewargs__(self):\n        'Return self as a plain tuple.  Used by copy and pickle.'\n        return tuple(self)\n\n    __dict__ = _property(_asdict)\n\n    def __getstate__(self):\n        'Exclude the OrderedDict from pickling'\n        pass\n\n{field_defs}\n").ToObject()); πE != nil {
				continue
			}
			// line 309: _repr_template = '{name}=%r'
			πF.SetLineno(309)
			if πE = πF.Globals().SetItem(πF, ß_repr_template.ToObject(), πg.NewStr("{name}=%r").ToObject()); πE != nil {
				continue
			}
			// line 311: _field_template = '''\
			πF.SetLineno(311)
			if πE = πF.Globals().SetItem(πF, ß_field_template.ToObject(), πg.NewStr("    {name} = _property(_itemgetter({index:d}), doc='Alias for field number {index:d}')\n").ToObject()); πE != nil {
				continue
			}
			// line 417: class Counter(dict):
			πF.SetLineno(417)
			πTemp001 = make([]*πg.Object, 1)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
				continue
			}
			πTemp001[0] = πTemp004
			πTemp010 = πg.NewDict()
			if πTemp002, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp010.SetItem(πF, ß__module__.ToObject(), πTemp002); πE != nil {
				continue
			}
			_, πE = πg.NewCode("Counter", "/usr/lib/python2.7/collections.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
				var πTemp007 *πg.Object
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
				var πTemp016 *πg.Object
				_ = πTemp016
				var πTemp017 *πg.Object
				_ = πTemp017
				var πTemp018 *πg.Object
				_ = πTemp018
				var πTemp019 *πg.Object
				_ = πTemp019
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 418: '''Dict subclass for counting hashable items.  Sometimes called a bag
					πF.SetLineno(418)
					// line 418: '''Dict subclass for counting hashable items.  Sometimes called a bag
					πF.SetLineno(418)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Dict subclass for counting hashable items.  Sometimes called a bag\n    or multiset.  Elements are stored as dictionary keys and their counts\n    are stored as dictionary values.\n\n    >>> c = Counter('abcdeabcdabcaba')  # count elements from a string\n\n    >>> c.most_common(3)                # three most common elements\n    [('a', 5), ('b', 4), ('c', 3)]\n    >>> sorted(c)                       # list all unique elements\n    ['a', 'b', 'c', 'd', 'e']\n    >>> ''.join(sorted(c.elements()))   # list elements with repetitions\n    'aaaaabbbbcccdde'\n    >>> sum(c.values())                 # total of all counts\n    15\n\n    >>> c['a']                          # count of letter 'a'\n    5\n    >>> for elem in 'shazam':           # update counts from an iterable\n    ...     c[elem] += 1                # by adding 1 to each element's count\n    >>> c['a']                          # now there are seven 'a'\n    7\n    >>> del c['b']                      # remove all 'b'\n    >>> c['b']                          # now there are zero 'b'\n    0\n\n    >>> d = Counter('simsalabim')       # make another counter\n    >>> c.update(d)                     # add in the second counter\n    >>> c['a']                          # now there are nine 'a'\n    9\n\n    >>> c.clear()                       # empty the counter\n    >>> c\n    Counter()\n\n    Note:  If a count is set to zero or reduced to zero, it will remain\n    in the counter until the entry is deleted or the counter is cleared:\n\n    >>> c = Counter('aaabbc')\n    >>> c['b'] -= 2                     # reduce the count of 'b' by two\n    >>> c.most_common()                 # 'b' is still in, but its count is zero\n    [('a', 3), ('c', 1), ('b', 0)]\n\n    ").ToObject()); πE != nil {
						continue
					}
					// line 468: def __init__(*args, **kwds):
					πF.SetLineno(468)
					πTemp002 = make([]πg.Param, 0)
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/collections.py", πTemp002, πg.CodeFlagVarArg | πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µargs *πg.Object = πArgs[0]; _ = µargs
						var µkwds *πg.Object = πArgs[1]; _ = µkwds
						var µself *πg.Object = πg.UnboundLocal; _ = µself
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
						var πTemp006 []*πg.Object
						_ = πTemp006
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 469: '''Create a new, empty Counter object.  And if given, count elements
							πF.SetLineno(469)
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, µargs); πE != nil {
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
							// line 479: if not args:
							πF.SetLineno(479)
						Label1:
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewStr("descriptor '__init__' of 'Counter' object needs an argument").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 480: raise TypeError("descriptor '__init__' of 'Counter' object "
							πF.SetLineno(480)
							πE = πF.Raise(πTemp004, nil, nil)
							continue
							goto Label2
						Label2:
							// line 482: self = args[0]
							πF.SetLineno(482)
							πTemp001 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µargs, πTemp001); πE != nil {
								continue
							}
							µself = πTemp004
							// line 483: args = args[1:]
							πF.SetLineno(483)
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µargs, πTemp001); πE != nil {
								continue
							}
							µargs = πTemp004
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							πTemp003[0] = µargs
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp001, πE = πg.GT(πF, πTemp005, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label3
							}
							goto Label4
							// line 484: if len(args) > 1:
							πF.SetLineno(484)
						Label3:
							πTemp003 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							πTemp006[0] = µargs
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("expected at most 1 arguments, got %d").ToObject(), πTemp005); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 485: raise TypeError('expected at most 1 arguments, got %d' % len(args))
							πF.SetLineno(485)
							πE = πF.Raise(πTemp004, nil, nil)
							continue
							goto Label4
						Label4:
							// line 486: super(Counter, self).__init__()
							πF.SetLineno(486)
							πTemp003 = πF.MakeArgs(2)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßCounter); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp003[1] = µself
							if πTemp001, πE = πg.ResolveGlobal(πF, ßsuper); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp001, πE = πg.GetAttr(πF, πTemp004, ß__init__, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 487: self.update(*args, **kwds)
							πF.SetLineno(487)
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µkwds, "kwds"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßupdate, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Invoke(πF, πTemp001, nil, µargs, nil, µkwds); πE != nil {
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
					// line 469: '''Create a new, empty Counter object.  And if given, count elements
					πF.SetLineno(469)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewStr("Create a new, empty Counter object.  And if given, count elements\n        from an input iterable.  Or, initialize the count from another mapping\n        of elements to their counts.\n\n        >>> c = Counter()                           # a new, empty counter\n        >>> c = Counter('gallahad')                 # a new counter from an iterable\n        >>> c = Counter({'a': 4, 'b': 2})           # a new counter from a mapping\n        >>> c = Counter(a=4, b=2)                   # a new counter from keyword args\n\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ß__init__); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp004, ß__doc__, πTemp003); πE != nil {
						continue
					}
					// line 489: def __missing__(self, key):
					πF.SetLineno(489)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "key", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("__missing__", "/usr/lib/python2.7/collections.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µkey *πg.Object = πArgs[1]; _ = µkey
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 490: 'The count of elements not in the Counter is zero.'
							πF.SetLineno(490)
							// line 492: return 0
							πF.SetLineno(492)
							πR = πg.NewInt(0).ToObject()
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß__missing__.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 490: 'The count of elements not in the Counter is zero.'
					πF.SetLineno(490)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πg.NewStr("The count of elements not in the Counter is zero.").ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ß__missing__); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp005, ß__doc__, πTemp004); πE != nil {
						continue
					}
					// line 494: def most_common(self, n=None):
					πF.SetLineno(494)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[1] = πg.Param{Name: "n", Def: πTemp005}
					πTemp004 = πg.NewFunction(πg.NewCode("most_common", "/usr/lib/python2.7/collections.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µn *πg.Object = πArgs[1]; _ = µn
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
						var πTemp006 πg.KWArgs
						_ = πTemp006
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 495: '''List the n most common elements and their counts from the most
							πF.SetLineno(495)
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µn == πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 503: if n is None:
							πF.SetLineno(503)
						Label1:
							// line 504: return sorted(self.iteritems(), key=_itemgetter(1), reverse=True)
							πF.SetLineno(504)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßiteritems, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp002
							πTemp005 = πF.MakeArgs(1)
							πTemp005[0] = πg.NewInt(1).ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_itemgetter); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp006 = πg.KWArgs{
								{"key", πTemp002},
								{"reverse", πTemp001},
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßsorted); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, πTemp006); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πR = πTemp002
							continue
							goto Label2
						Label2:
							// line 505: return _heapq.nlargest(n, self.iteritems(), key=_itemgetter(1))
							πF.SetLineno(505)
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							πTemp004[0] = µn
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßiteritems, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp004[1] = πTemp002
							πTemp005 = πF.MakeArgs(1)
							πTemp005[0] = πg.NewInt(1).ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_itemgetter); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp006 = πg.KWArgs{
								{"key", πTemp002},
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_heapq); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßnlargest, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp004, πTemp006); πE != nil {
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
					if πE = πClass.SetItem(πF, ßmost_common.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 495: '''List the n most common elements and their counts from the most
					πF.SetLineno(495)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πg.NewStr("List the n most common elements and their counts from the most\n        common to the least.  If n is None, then list all element counts.\n\n        >>> Counter('abcdeabcdabcaba').most_common(3)\n        [('a', 5), ('b', 4), ('c', 3)]\n\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßmost_common); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp006, ß__doc__, πTemp005); πE != nil {
						continue
					}
					// line 507: def elements(self):
					πF.SetLineno(507)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("elements", "/usr/lib/python2.7/collections.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 508: '''Iterator over elements repeating each as many times as its count.
							πF.SetLineno(508)
							// line 527: return _chain.from_iterable(_starmap(_repeat, self.iteritems()))
							πF.SetLineno(527)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(2)
							if πTemp003, πE = πg.ResolveGlobal(πF, ß_repeat); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßiteritems, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp002[1] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ß_starmap); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ß_chain); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßfrom_iterable, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßelements.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 508: '''Iterator over elements repeating each as many times as its count.
					πF.SetLineno(508)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, πg.NewStr("Iterator over elements repeating each as many times as its count.\n\n        >>> c = Counter('ABCABC')\n        >>> sorted(c.elements())\n        ['A', 'A', 'B', 'B', 'C', 'C']\n\n        # Knuth's example for prime factors of 1836:  2**2 * 3**3 * 17**1\n        >>> prime_factors = Counter({2: 2, 3: 3, 17: 1})\n        >>> product = 1\n        >>> for factor in prime_factors.elements():     # loop over factors\n        ...     product *= factor                       # and multiply them\n        >>> product\n        1836\n\n        Note, if an element's count has been set to zero or is a negative\n        number, elements() will ignore it.\n\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßelements); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp007, ß__doc__, πTemp006); πE != nil {
						continue
					}
					// line 532: def fromkeys(cls, iterable, v=None):
					πF.SetLineno(532)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "cls", Def: nil}
					πTemp002[1] = πg.Param{Name: "iterable", Def: nil}
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "v", Def: πTemp007}
					πTemp006 = πg.NewFunction(πg.NewCode("fromkeys", "/usr/lib/python2.7/collections.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µcls *πg.Object = πArgs[0]; _ = µcls
						var µiterable *πg.Object = πArgs[1]; _ = µiterable
						var µv *πg.Object = πArgs[2]; _ = µv
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
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("Counter.fromkeys() is undefined.  Use Counter(iterable) instead.").ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNotImplementedError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 535: raise NotImplementedError(
							πF.SetLineno(535)
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
					if πE = πClass.SetItem(πF, ßfromkeys.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 531: @classmethod
					πF.SetLineno(531)
					πTemp008 = πF.MakeArgs(1)
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßfromkeys); πE != nil {
						continue
					}
					πTemp008[0] = πTemp007
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßclassmethod); πE != nil {
						continue
					}
					if πTemp009, πE = πTemp007.Call(πF, πTemp008, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp008)
					if πE = πClass.SetItem(πF, ßfromkeys.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 538: def update(*args, **kwds):
					πF.SetLineno(538)
					πTemp002 = make([]πg.Param, 0)
					πTemp007 = πg.NewFunction(πg.NewCode("update", "/usr/lib/python2.7/collections.py", πTemp002, πg.CodeFlagVarArg | πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µargs *πg.Object = πArgs[0]; _ = µargs
						var µkwds *πg.Object = πArgs[1]; _ = µkwds
						var µself *πg.Object = πg.UnboundLocal; _ = µself
						var µiterable *πg.Object = πg.UnboundLocal; _ = µiterable
						var µself_get *πg.Object = πg.UnboundLocal; _ = µself_get
						var µelem *πg.Object = πg.UnboundLocal; _ = µelem
						var µcount *πg.Object = πg.UnboundLocal; _ = µcount
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
						var πTemp006 []*πg.Object
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
							case 16: goto Label16
							case 18: goto Label18
							case 19: goto Label19
							case 15: goto Label15
							default: panic("unexpected function state")
							}
							// line 539: '''Like dict.update() but add counts instead of replacing them.
							πF.SetLineno(539)
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, µargs); πE != nil {
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
							// line 558: if not args:
							πF.SetLineno(558)
						Label1:
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewStr("descriptor 'update' of 'Counter' object needs an argument").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 559: raise TypeError("descriptor 'update' of 'Counter' object "
							πF.SetLineno(559)
							πE = πF.Raise(πTemp004, nil, nil)
							continue
							goto Label2
						Label2:
							// line 561: self = args[0]
							πF.SetLineno(561)
							πTemp001 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µargs, πTemp001); πE != nil {
								continue
							}
							µself = πTemp004
							// line 562: args = args[1:]
							πF.SetLineno(562)
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µargs, πTemp001); πE != nil {
								continue
							}
							µargs = πTemp004
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							πTemp003[0] = µargs
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp001, πE = πg.GT(πF, πTemp005, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label3
							}
							goto Label4
							// line 563: if len(args) > 1:
							πF.SetLineno(563)
						Label3:
							πTemp003 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							πTemp006[0] = µargs
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("expected at most 1 arguments, got %d").ToObject(), πTemp005); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 564: raise TypeError('expected at most 1 arguments, got %d' % len(args))
							πF.SetLineno(564)
							πE = πF.Raise(πTemp004, nil, nil)
							continue
							goto Label4
						Label4:
							// line 565: iterable = args[0] if args else None
							πF.SetLineno(565)
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, µargs); πE != nil {
								continue
							}
							if !πTemp002 {
								goto Label5
							}
							πTemp004 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µargs, πTemp004); πE != nil {
								continue
							}
							πTemp001 = πTemp005
							goto Label6
						Label5:
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πTemp004
						Label6:
							µiterable = πTemp001
							if πE = πg.CheckLocal(πF, µiterable, "iterable"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µiterable != πTemp004).ToObject()
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label7
							}
							goto Label8
							// line 566: if iterable is not None:
							πF.SetLineno(566)
						Label7:
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µiterable, "iterable"); πE != nil {
								continue
							}
							πTemp003[0] = µiterable
							if πTemp001, πE = πg.ResolveGlobal(πF, ßMapping); πE != nil {
								continue
							}
							πTemp003[1] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp002, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label9
							}
							goto Label10
							// line 567: if isinstance(iterable, Mapping):
							πF.SetLineno(567)
						Label9:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, µself); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label12
							}
							goto Label13
							// line 568: if self:
							πF.SetLineno(568)
						Label12:
							// line 569: self_get = self.get
							πF.SetLineno(569)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßget, nil); πE != nil {
								continue
							}
							µself_get = πTemp001
							if πE = πg.CheckLocal(πF, µiterable, "iterable"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µiterable, ßiteritems, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, πTemp005); πE != nil {
								continue
							}
							πF.PushCheckpoint(16)
							πTemp002 = false
						Label15:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp002 {
								πF.PopCheckpoint()
								goto Label17
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
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp008}}}, πTemp004); πE != nil {
									continue
								}
								µelem = πTemp005
								µcount = πTemp008
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(15)            
							// line 571: self[elem] = self_get(elem, 0) + count
							πF.SetLineno(571)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µelem, "elem"); πE != nil {
								continue
							}
							πTemp003[0] = µelem
							πTemp003[1] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself_get, "self_get"); πE != nil {
								continue
							}
							if πTemp005, πE = µself_get.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πE = πg.CheckLocal(πF, µcount, "count"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Add(πF, πTemp005, µcount); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µelem, "elem"); πE != nil {
								continue
							}
							πTemp008 = µelem
							if πE = πg.SetItem(πF, µself, πTemp008, πTemp005); πE != nil {
								continue
							}
							continue
						Label16:
							if πE != nil || πR != nil {
								continue
							}
						Label17:
							goto Label14
						Label13:
							// line 573: super(Counter, self).update(iterable) # fast path when counter is empty
							πF.SetLineno(573)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µiterable, "iterable"); πE != nil {
								continue
							}
							πTemp003[0] = µiterable
							πTemp006 = πF.MakeArgs(2)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßCounter); πE != nil {
								continue
							}
							πTemp006[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp006[1] = µself
							if πTemp001, πE = πg.ResolveGlobal(πF, ßsuper); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πTemp001, πE = πg.GetAttr(πF, πTemp004, ßupdate, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							goto Label14
						Label14:
							goto Label11
						Label10:
							// line 575: self_get = self.get
							πF.SetLineno(575)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßget, nil); πE != nil {
								continue
							}
							µself_get = πTemp001
							if πE = πg.CheckLocal(πF, µiterable, "iterable"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, µiterable); πE != nil {
								continue
							}
							πF.PushCheckpoint(19)
							πTemp002 = false
						Label18:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp002 {
								πF.PopCheckpoint()
								goto Label20
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
								µelem = πTemp004
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(18)            
							// line 577: self[elem] = self_get(elem, 0) + 1
							πF.SetLineno(577)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µelem, "elem"); πE != nil {
								continue
							}
							πTemp003[0] = µelem
							πTemp003[1] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself_get, "self_get"); πE != nil {
								continue
							}
							if πTemp005, πE = µself_get.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp004, πE = πg.Add(πF, πTemp005, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µelem, "elem"); πE != nil {
								continue
							}
							πTemp008 = µelem
							if πE = πg.SetItem(πF, µself, πTemp008, πTemp005); πE != nil {
								continue
							}
							continue
						Label19:
							if πE != nil || πR != nil {
								continue
							}
						Label20:
							goto Label11
						Label11:
							goto Label8
						Label8:
							if πE = πg.CheckLocal(πF, µkwds, "kwds"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, µkwds); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label21
							}
							goto Label22
							// line 578: if kwds:
							πF.SetLineno(578)
						Label21:
							// line 579: self.update(kwds)
							πF.SetLineno(579)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µkwds, "kwds"); πE != nil {
								continue
							}
							πTemp003[0] = µkwds
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßupdate, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							goto Label22
						Label22:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßupdate.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 539: '''Like dict.update() but add counts instead of replacing them.
					πF.SetLineno(539)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp009}, πg.NewStr("Like dict.update() but add counts instead of replacing them.\n\n        Source can be an iterable, a dictionary, or another Counter instance.\n\n        >>> c = Counter('which')\n        >>> c.update('witch')           # add elements from another iterable\n        >>> d = Counter('watch')\n        >>> c.update(d)                 # add elements from another counter\n        >>> c['h']                      # four 'h' in which, witch, and watch\n        4\n\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßupdate); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp010, ß__doc__, πTemp009); πE != nil {
						continue
					}
					// line 581: def subtract(*args, **kwds):
					πF.SetLineno(581)
					πTemp002 = make([]πg.Param, 0)
					πTemp009 = πg.NewFunction(πg.NewCode("subtract", "/usr/lib/python2.7/collections.py", πTemp002, πg.CodeFlagVarArg | πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µargs *πg.Object = πArgs[0]; _ = µargs
						var µkwds *πg.Object = πArgs[1]; _ = µkwds
						var µself *πg.Object = πg.UnboundLocal; _ = µself
						var µiterable *πg.Object = πg.UnboundLocal; _ = µiterable
						var µself_get *πg.Object = πg.UnboundLocal; _ = µself_get
						var µelem *πg.Object = πg.UnboundLocal; _ = µelem
						var µcount *πg.Object = πg.UnboundLocal; _ = µcount
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
						var πTemp006 []*πg.Object
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
							case 16: goto Label16
							case 12: goto Label12
							case 13: goto Label13
							case 15: goto Label15
							default: panic("unexpected function state")
							}
							// line 582: '''Like dict.update() but subtracts counts instead of replacing them.
							πF.SetLineno(582)
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, µargs); πE != nil {
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
							// line 597: if not args:
							πF.SetLineno(597)
						Label1:
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewStr("descriptor 'subtract' of 'Counter' object needs an argument").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 598: raise TypeError("descriptor 'subtract' of 'Counter' object "
							πF.SetLineno(598)
							πE = πF.Raise(πTemp004, nil, nil)
							continue
							goto Label2
						Label2:
							// line 600: self = args[0]
							πF.SetLineno(600)
							πTemp001 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µargs, πTemp001); πE != nil {
								continue
							}
							µself = πTemp004
							// line 601: args = args[1:]
							πF.SetLineno(601)
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µargs, πTemp001); πE != nil {
								continue
							}
							µargs = πTemp004
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							πTemp003[0] = µargs
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp001, πE = πg.GT(πF, πTemp005, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label3
							}
							goto Label4
							// line 602: if len(args) > 1:
							πF.SetLineno(602)
						Label3:
							πTemp003 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							πTemp006[0] = µargs
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("expected at most 1 arguments, got %d").ToObject(), πTemp005); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 603: raise TypeError('expected at most 1 arguments, got %d' % len(args))
							πF.SetLineno(603)
							πE = πF.Raise(πTemp004, nil, nil)
							continue
							goto Label4
						Label4:
							// line 604: iterable = args[0] if args else None
							πF.SetLineno(604)
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, µargs); πE != nil {
								continue
							}
							if !πTemp002 {
								goto Label5
							}
							πTemp004 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µargs, πTemp004); πE != nil {
								continue
							}
							πTemp001 = πTemp005
							goto Label6
						Label5:
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πTemp004
						Label6:
							µiterable = πTemp001
							if πE = πg.CheckLocal(πF, µiterable, "iterable"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µiterable != πTemp004).ToObject()
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label7
							}
							goto Label8
							// line 605: if iterable is not None:
							πF.SetLineno(605)
						Label7:
							// line 606: self_get = self.get
							πF.SetLineno(606)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßget, nil); πE != nil {
								continue
							}
							µself_get = πTemp001
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µiterable, "iterable"); πE != nil {
								continue
							}
							πTemp003[0] = µiterable
							if πTemp001, πE = πg.ResolveGlobal(πF, ßMapping); πE != nil {
								continue
							}
							πTemp003[1] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp002, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label9
							}
							goto Label10
							// line 607: if isinstance(iterable, Mapping):
							πF.SetLineno(607)
						Label9:
							if πE = πg.CheckLocal(πF, µiterable, "iterable"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µiterable, ßitems, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, πTemp005); πE != nil {
								continue
							}
							πF.PushCheckpoint(13)
							πTemp002 = false
						Label12:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp002 {
								πF.PopCheckpoint()
								goto Label14
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
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp008}}}, πTemp004); πE != nil {
									continue
								}
								µelem = πTemp005
								µcount = πTemp008
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(12)            
							// line 609: self[elem] = self_get(elem, 0) - count
							πF.SetLineno(609)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µelem, "elem"); πE != nil {
								continue
							}
							πTemp003[0] = µelem
							πTemp003[1] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself_get, "self_get"); πE != nil {
								continue
							}
							if πTemp005, πE = µself_get.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πE = πg.CheckLocal(πF, µcount, "count"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Sub(πF, πTemp005, µcount); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µelem, "elem"); πE != nil {
								continue
							}
							πTemp008 = µelem
							if πE = πg.SetItem(πF, µself, πTemp008, πTemp005); πE != nil {
								continue
							}
							continue
						Label13:
							if πE != nil || πR != nil {
								continue
							}
						Label14:
							goto Label11
						Label10:
							if πE = πg.CheckLocal(πF, µiterable, "iterable"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, µiterable); πE != nil {
								continue
							}
							πF.PushCheckpoint(16)
							πTemp002 = false
						Label15:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp002 {
								πF.PopCheckpoint()
								goto Label17
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
								µelem = πTemp004
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(15)            
							// line 612: self[elem] = self_get(elem, 0) - 1
							πF.SetLineno(612)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µelem, "elem"); πE != nil {
								continue
							}
							πTemp003[0] = µelem
							πTemp003[1] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself_get, "self_get"); πE != nil {
								continue
							}
							if πTemp005, πE = µself_get.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp004, πE = πg.Sub(πF, πTemp005, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µelem, "elem"); πE != nil {
								continue
							}
							πTemp008 = µelem
							if πE = πg.SetItem(πF, µself, πTemp008, πTemp005); πE != nil {
								continue
							}
							continue
						Label16:
							if πE != nil || πR != nil {
								continue
							}
						Label17:
							goto Label11
						Label11:
							goto Label8
						Label8:
							if πE = πg.CheckLocal(πF, µkwds, "kwds"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, µkwds); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label18
							}
							goto Label19
							// line 613: if kwds:
							πF.SetLineno(613)
						Label18:
							// line 614: self.subtract(kwds)
							πF.SetLineno(614)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µkwds, "kwds"); πE != nil {
								continue
							}
							πTemp003[0] = µkwds
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßsubtract, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							goto Label19
						Label19:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßsubtract.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 582: '''Like dict.update() but subtracts counts instead of replacing them.
					πF.SetLineno(582)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp010}, πg.NewStr("Like dict.update() but subtracts counts instead of replacing them.\n        Counts can be reduced below zero.  Both the inputs and outputs are\n        allowed to contain zero and negative counts.\n\n        Source can be an iterable, a dictionary, or another Counter instance.\n\n        >>> c = Counter('which')\n        >>> c.subtract('witch')             # subtract elements from another iterable\n        >>> c.subtract(Counter('watch'))    # subtract elements from another counter\n        >>> c['h']                          # 2 in which, minus 1 in witch, minus 1 in watch\n        0\n        >>> c['w']                          # 1 in which, minus 1 in witch, minus 1 in watch\n        -1\n\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ßsubtract); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp011, ß__doc__, πTemp010); πE != nil {
						continue
					}
					// line 616: def copy(self):
					πF.SetLineno(616)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp010 = πg.NewFunction(πg.NewCode("copy", "/usr/lib/python2.7/collections.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 617: 'Return a shallow copy.'
							πF.SetLineno(617)
							// line 618: return self.__class__(self)
							πF.SetLineno(618)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß__class__, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßcopy.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 617: 'Return a shallow copy.'
					πF.SetLineno(617)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πg.NewStr("Return a shallow copy.").ToObject()); πE != nil {
						continue
					}
					if πTemp012, πE = πg.ResolveClass(πF, πClass, nil, ßcopy); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp012, ß__doc__, πTemp011); πE != nil {
						continue
					}
					// line 620: def __reduce__(self):
					πF.SetLineno(620)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp011 = πg.NewFunction(πg.NewCode("__reduce__", "/usr/lib/python2.7/collections.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
						var πTemp006 *πg.Object
						_ = πTemp006
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 621: return self.__class__, (dict(self),)
							πF.SetLineno(621)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß__class__, nil); πE != nil {
								continue
							}
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp004[0] = µself
							if πTemp005, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp003 = πg.NewTuple1(πTemp006).ToObject()
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
					if πE = πClass.SetItem(πF, ß__reduce__.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 623: def __delitem__(self, elem):
					πF.SetLineno(623)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "elem", Def: nil}
					πTemp012 = πg.NewFunction(πg.NewCode("__delitem__", "/usr/lib/python2.7/collections.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µelem *πg.Object = πArgs[1]; _ = µelem
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
						_ = πTemp002
						var πTemp003 []*πg.Object
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
							// line 624: 'Like dict.__delitem__() but does not raise KeyError for missing values.'
							πF.SetLineno(624)
							if πE = πg.CheckLocal(πF, µelem, "elem"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Contains(πF, µself, µelem); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp002).ToObject()
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							goto Label2
							// line 625: if elem in self:
							πF.SetLineno(625)
						Label1:
							// line 626: super(Counter, self).__delitem__(elem)
							πF.SetLineno(626)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µelem, "elem"); πE != nil {
								continue
							}
							πTemp003[0] = µelem
							πTemp004 = πF.MakeArgs(2)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßCounter); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp004[1] = µself
							if πTemp001, πE = πg.ResolveGlobal(πF, ßsuper); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp001, πE = πg.GetAttr(πF, πTemp005, ß__delitem__, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
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
					if πE = πClass.SetItem(πF, ß__delitem__.ToObject(), πTemp012); πE != nil {
						continue
					}
					// line 624: 'Like dict.__delitem__() but does not raise KeyError for missing values.'
					πF.SetLineno(624)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp013}, πg.NewStr("Like dict.__delitem__() but does not raise KeyError for missing values.").ToObject()); πE != nil {
						continue
					}
					if πTemp014, πE = πg.ResolveClass(πF, πClass, nil, ß__delitem__); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp014, ß__doc__, πTemp013); πE != nil {
						continue
					}
					// line 628: def __repr__(self):
					πF.SetLineno(628)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp013 = πg.NewFunction(πg.NewCode("__repr__", "/usr/lib/python2.7/collections.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µitems *πg.Object = πg.UnboundLocal; _ = µitems
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
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, µself); πE != nil {
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
							// line 629: if not self:
							πF.SetLineno(629)
						Label1:
							// line 630: return '%s()' % self.__class__.__name__
							πF.SetLineno(630)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß__class__, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ß__name__, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("%s()").ToObject(), πTemp004); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label2
						Label2:
							// line 631: items = ', '.join(map('%r: %r'.__mod__, self.most_common()))
							πF.SetLineno(631)
							πTemp005 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(2)
							if πTemp001, πE = πg.GetAttr(πF, πg.NewStr("%r: %r").ToObject(), ß__mod__, nil); πE != nil {
								continue
							}
							πTemp006[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßmost_common, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp006[1] = πTemp003
							if πTemp001, πE = πg.ResolveGlobal(πF, ßmap); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp005[0] = πTemp003
							if πTemp001, πE = πg.GetAttr(πF, πg.NewStr(", ").ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µitems = πTemp003
							// line 632: return '%s({%s})' % (self.__class__.__name__, items)
							πF.SetLineno(632)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß__class__, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp004, ß__name__, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µitems, "items"); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple2(πTemp007, µitems).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("%s({%s})").ToObject(), πTemp003); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__repr__.ToObject(), πTemp013); πE != nil {
						continue
					}
					// line 643: def __add__(self, other):
					πF.SetLineno(643)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "other", Def: nil}
					πTemp014 = πg.NewFunction(πg.NewCode("__add__", "/usr/lib/python2.7/collections.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
						var µresult *πg.Object = πg.UnboundLocal; _ = µresult
						var µelem *πg.Object = πg.UnboundLocal; _ = µelem
						var µcount *πg.Object = πg.UnboundLocal; _ = µcount
						var µnewcount *πg.Object = πg.UnboundLocal; _ = µnewcount
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
						var πTemp008 bool
						_ = πTemp008
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 8: goto Label8
							case 9: goto Label9
							case 3: goto Label3
							case 4: goto Label4
							default: panic("unexpected function state")
							}
							// line 644: '''Add counts from two counters.
							πF.SetLineno(644)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp002[0] = µother
							if πTemp003, πE = πg.ResolveGlobal(πF, ßCounter); πE != nil {
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
							// line 650: if not isinstance(other, Counter):
							πF.SetLineno(650)
						Label1:
							// line 651: return NotImplemented
							πF.SetLineno(651)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNotImplemented); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label2
						Label2:
							// line 652: result = Counter()
							πF.SetLineno(652)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßCounter); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µresult = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßitems, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, πTemp004); πE != nil {
								continue
							}
							πF.PushCheckpoint(4)
							πTemp005 = false
						Label3:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp005 {
								πF.PopCheckpoint()
								goto Label5
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
								µelem = πTemp004
								µcount = πTemp007
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(3)            
							// line 654: newcount = count + other[elem]
							πF.SetLineno(654)
							if πE = πg.CheckLocal(πF, µcount, "count"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µelem, "elem"); πE != nil {
								continue
							}
							πTemp004 = µelem
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, µother, πTemp004); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, µcount, πTemp007); πE != nil {
								continue
							}
							µnewcount = πTemp003
							if πE = πg.CheckLocal(πF, µnewcount, "newcount"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GT(πF, µnewcount, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label6
							}
							goto Label7
							// line 655: if newcount > 0:
							πF.SetLineno(655)
						Label6:
							// line 656: result[elem] = newcount
							πF.SetLineno(656)
							if πE = πg.CheckLocal(πF, µnewcount, "newcount"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, µnewcount); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µelem, "elem"); πE != nil {
								continue
							}
							πTemp004 = µelem
							if πE = πg.SetItem(πF, µresult, πTemp004, πTemp003); πE != nil {
								continue
							}
							goto Label7
						Label7:
							continue
						Label4:
							if πE != nil || πR != nil {
								continue
							}
						Label5:
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µother, ßitems, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, πTemp004); πE != nil {
								continue
							}
							πF.PushCheckpoint(9)
							πTemp005 = false
						Label8:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp005 {
								πF.PopCheckpoint()
								goto Label10
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
								µelem = πTemp004
								µcount = πTemp007
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(8)            
							if πE = πg.CheckLocal(πF, µelem, "elem"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.Contains(πF, µself, µelem); πE != nil {
								continue
							}
							πTemp004 = πg.GetBool(!πTemp008).ToObject()
							πTemp003 = πTemp004
							if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if !πTemp006 {
								goto Label11
							}
							if πE = πg.CheckLocal(πF, µcount, "count"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GT(πF, µcount, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							πTemp003 = πTemp004
						Label11:
							if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label12
							}
							goto Label13
							// line 658: if elem not in self and count > 0:
							πF.SetLineno(658)
						Label12:
							// line 659: result[elem] = count
							πF.SetLineno(659)
							if πE = πg.CheckLocal(πF, µcount, "count"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, µcount); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µelem, "elem"); πE != nil {
								continue
							}
							πTemp004 = µelem
							if πE = πg.SetItem(πF, µresult, πTemp004, πTemp003); πE != nil {
								continue
							}
							goto Label13
						Label13:
							continue
						Label9:
							if πE != nil || πR != nil {
								continue
							}
						Label10:
							// line 660: return result
							πF.SetLineno(660)
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
					if πE = πClass.SetItem(πF, ß__add__.ToObject(), πTemp014); πE != nil {
						continue
					}
					// line 644: '''Add counts from two counters.
					πF.SetLineno(644)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp015}, πg.NewStr("Add counts from two counters.\n\n        >>> Counter('abbb') + Counter('bcc')\n        Counter({'b': 4, 'c': 2, 'a': 1})\n\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp016, πE = πg.ResolveClass(πF, πClass, nil, ß__add__); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp016, ß__doc__, πTemp015); πE != nil {
						continue
					}
					// line 662: def __sub__(self, other):
					πF.SetLineno(662)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "other", Def: nil}
					πTemp015 = πg.NewFunction(πg.NewCode("__sub__", "/usr/lib/python2.7/collections.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
						var µresult *πg.Object = πg.UnboundLocal; _ = µresult
						var µelem *πg.Object = πg.UnboundLocal; _ = µelem
						var µcount *πg.Object = πg.UnboundLocal; _ = µcount
						var µnewcount *πg.Object = πg.UnboundLocal; _ = µnewcount
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
						var πTemp008 bool
						_ = πTemp008
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 8: goto Label8
							case 9: goto Label9
							case 3: goto Label3
							case 4: goto Label4
							default: panic("unexpected function state")
							}
							// line 663: ''' Subtract count, but keep only results with positive counts.
							πF.SetLineno(663)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp002[0] = µother
							if πTemp003, πE = πg.ResolveGlobal(πF, ßCounter); πE != nil {
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
							// line 669: if not isinstance(other, Counter):
							πF.SetLineno(669)
						Label1:
							// line 670: return NotImplemented
							πF.SetLineno(670)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNotImplemented); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label2
						Label2:
							// line 671: result = Counter()
							πF.SetLineno(671)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßCounter); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µresult = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßitems, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, πTemp004); πE != nil {
								continue
							}
							πF.PushCheckpoint(4)
							πTemp005 = false
						Label3:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp005 {
								πF.PopCheckpoint()
								goto Label5
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
								µelem = πTemp004
								µcount = πTemp007
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(3)            
							// line 673: newcount = count - other[elem]
							πF.SetLineno(673)
							if πE = πg.CheckLocal(πF, µcount, "count"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µelem, "elem"); πE != nil {
								continue
							}
							πTemp004 = µelem
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, µother, πTemp004); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Sub(πF, µcount, πTemp007); πE != nil {
								continue
							}
							µnewcount = πTemp003
							if πE = πg.CheckLocal(πF, µnewcount, "newcount"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GT(πF, µnewcount, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label6
							}
							goto Label7
							// line 674: if newcount > 0:
							πF.SetLineno(674)
						Label6:
							// line 675: result[elem] = newcount
							πF.SetLineno(675)
							if πE = πg.CheckLocal(πF, µnewcount, "newcount"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, µnewcount); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µelem, "elem"); πE != nil {
								continue
							}
							πTemp004 = µelem
							if πE = πg.SetItem(πF, µresult, πTemp004, πTemp003); πE != nil {
								continue
							}
							goto Label7
						Label7:
							continue
						Label4:
							if πE != nil || πR != nil {
								continue
							}
						Label5:
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µother, ßitems, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, πTemp004); πE != nil {
								continue
							}
							πF.PushCheckpoint(9)
							πTemp005 = false
						Label8:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp005 {
								πF.PopCheckpoint()
								goto Label10
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
								µelem = πTemp004
								µcount = πTemp007
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(8)            
							if πE = πg.CheckLocal(πF, µelem, "elem"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.Contains(πF, µself, µelem); πE != nil {
								continue
							}
							πTemp004 = πg.GetBool(!πTemp008).ToObject()
							πTemp003 = πTemp004
							if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if !πTemp006 {
								goto Label11
							}
							if πE = πg.CheckLocal(πF, µcount, "count"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.LT(πF, µcount, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							πTemp003 = πTemp004
						Label11:
							if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label12
							}
							goto Label13
							// line 677: if elem not in self and count < 0:
							πF.SetLineno(677)
						Label12:
							// line 678: result[elem] = 0 - count
							πF.SetLineno(678)
							if πE = πg.CheckLocal(πF, µcount, "count"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Sub(πF, πg.NewInt(0).ToObject(), µcount); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µelem, "elem"); πE != nil {
								continue
							}
							πTemp007 = µelem
							if πE = πg.SetItem(πF, µresult, πTemp007, πTemp004); πE != nil {
								continue
							}
							goto Label13
						Label13:
							continue
						Label9:
							if πE != nil || πR != nil {
								continue
							}
						Label10:
							// line 679: return result
							πF.SetLineno(679)
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
					if πE = πClass.SetItem(πF, ß__sub__.ToObject(), πTemp015); πE != nil {
						continue
					}
					// line 663: ''' Subtract count, but keep only results with positive counts.
					πF.SetLineno(663)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp016}, πg.NewStr(" Subtract count, but keep only results with positive counts.\n\n        >>> Counter('abbbc') - Counter('bccd')\n        Counter({'b': 2, 'a': 1})\n\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp017, πE = πg.ResolveClass(πF, πClass, nil, ß__sub__); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp017, ß__doc__, πTemp016); πE != nil {
						continue
					}
					// line 681: def __or__(self, other):
					πF.SetLineno(681)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "other", Def: nil}
					πTemp016 = πg.NewFunction(πg.NewCode("__or__", "/usr/lib/python2.7/collections.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
						var µresult *πg.Object = πg.UnboundLocal; _ = µresult
						var µelem *πg.Object = πg.UnboundLocal; _ = µelem
						var µcount *πg.Object = πg.UnboundLocal; _ = µcount
						var µother_count *πg.Object = πg.UnboundLocal; _ = µother_count
						var µnewcount *πg.Object = πg.UnboundLocal; _ = µnewcount
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
						var πTemp008 bool
						_ = πTemp008
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 11: goto Label11
							case 10: goto Label10
							case 3: goto Label3
							case 4: goto Label4
							default: panic("unexpected function state")
							}
							// line 682: '''Union is the maximum of value in either of the input counters.
							πF.SetLineno(682)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp002[0] = µother
							if πTemp003, πE = πg.ResolveGlobal(πF, ßCounter); πE != nil {
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
							// line 688: if not isinstance(other, Counter):
							πF.SetLineno(688)
						Label1:
							// line 689: return NotImplemented
							πF.SetLineno(689)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNotImplemented); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label2
						Label2:
							// line 690: result = Counter()
							πF.SetLineno(690)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßCounter); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µresult = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßitems, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, πTemp004); πE != nil {
								continue
							}
							πF.PushCheckpoint(4)
							πTemp005 = false
						Label3:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp005 {
								πF.PopCheckpoint()
								goto Label5
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
								µelem = πTemp004
								µcount = πTemp007
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(3)            
							// line 692: other_count = other[elem]
							πF.SetLineno(692)
							if πE = πg.CheckLocal(πF, µelem, "elem"); πE != nil {
								continue
							}
							πTemp003 = µelem
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µother, πTemp003); πE != nil {
								continue
							}
							µother_count = πTemp004
							// line 693: newcount = other_count if count < other_count else count
							πF.SetLineno(693)
							if πE = πg.CheckLocal(πF, µcount, "count"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µother_count, "other_count"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.LT(πF, µcount, µother_count); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if !πTemp006 {
								goto Label6
							}
							if πE = πg.CheckLocal(πF, µother_count, "other_count"); πE != nil {
								continue
							}
							πTemp003 = µother_count
							goto Label7
						Label6:
							if πE = πg.CheckLocal(πF, µcount, "count"); πE != nil {
								continue
							}
							πTemp003 = µcount
						Label7:
							µnewcount = πTemp003
							if πE = πg.CheckLocal(πF, µnewcount, "newcount"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GT(πF, µnewcount, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label8
							}
							goto Label9
							// line 694: if newcount > 0:
							πF.SetLineno(694)
						Label8:
							// line 695: result[elem] = newcount
							πF.SetLineno(695)
							if πE = πg.CheckLocal(πF, µnewcount, "newcount"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, µnewcount); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µelem, "elem"); πE != nil {
								continue
							}
							πTemp004 = µelem
							if πE = πg.SetItem(πF, µresult, πTemp004, πTemp003); πE != nil {
								continue
							}
							goto Label9
						Label9:
							continue
						Label4:
							if πE != nil || πR != nil {
								continue
							}
						Label5:
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µother, ßitems, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, πTemp004); πE != nil {
								continue
							}
							πF.PushCheckpoint(11)
							πTemp005 = false
						Label10:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp005 {
								πF.PopCheckpoint()
								goto Label12
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
								µelem = πTemp004
								µcount = πTemp007
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(10)            
							if πE = πg.CheckLocal(πF, µelem, "elem"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.Contains(πF, µself, µelem); πE != nil {
								continue
							}
							πTemp004 = πg.GetBool(!πTemp008).ToObject()
							πTemp003 = πTemp004
							if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if !πTemp006 {
								goto Label13
							}
							if πE = πg.CheckLocal(πF, µcount, "count"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GT(πF, µcount, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							πTemp003 = πTemp004
						Label13:
							if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label14
							}
							goto Label15
							// line 697: if elem not in self and count > 0:
							πF.SetLineno(697)
						Label14:
							// line 698: result[elem] = count
							πF.SetLineno(698)
							if πE = πg.CheckLocal(πF, µcount, "count"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, µcount); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µelem, "elem"); πE != nil {
								continue
							}
							πTemp004 = µelem
							if πE = πg.SetItem(πF, µresult, πTemp004, πTemp003); πE != nil {
								continue
							}
							goto Label15
						Label15:
							continue
						Label11:
							if πE != nil || πR != nil {
								continue
							}
						Label12:
							// line 699: return result
							πF.SetLineno(699)
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
					if πE = πClass.SetItem(πF, ß__or__.ToObject(), πTemp016); πE != nil {
						continue
					}
					// line 682: '''Union is the maximum of value in either of the input counters.
					πF.SetLineno(682)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp017}, πg.NewStr("Union is the maximum of value in either of the input counters.\n\n        >>> Counter('abbb') | Counter('bcc')\n        Counter({'b': 3, 'c': 2, 'a': 1})\n\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp018, πE = πg.ResolveClass(πF, πClass, nil, ß__or__); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp018, ß__doc__, πTemp017); πE != nil {
						continue
					}
					// line 701: def __and__(self, other):
					πF.SetLineno(701)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "other", Def: nil}
					πTemp017 = πg.NewFunction(πg.NewCode("__and__", "/usr/lib/python2.7/collections.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
						var µresult *πg.Object = πg.UnboundLocal; _ = µresult
						var µelem *πg.Object = πg.UnboundLocal; _ = µelem
						var µcount *πg.Object = πg.UnboundLocal; _ = µcount
						var µother_count *πg.Object = πg.UnboundLocal; _ = µother_count
						var µnewcount *πg.Object = πg.UnboundLocal; _ = µnewcount
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 3: goto Label3
							case 4: goto Label4
							default: panic("unexpected function state")
							}
							// line 702: ''' Intersection is the minimum of corresponding counts.
							πF.SetLineno(702)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp002[0] = µother
							if πTemp003, πE = πg.ResolveGlobal(πF, ßCounter); πE != nil {
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
							// line 708: if not isinstance(other, Counter):
							πF.SetLineno(708)
						Label1:
							// line 709: return NotImplemented
							πF.SetLineno(709)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNotImplemented); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label2
						Label2:
							// line 710: result = Counter()
							πF.SetLineno(710)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßCounter); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µresult = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßitems, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, πTemp004); πE != nil {
								continue
							}
							πF.PushCheckpoint(4)
							πTemp005 = false
						Label3:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp005 {
								πF.PopCheckpoint()
								goto Label5
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
								µelem = πTemp004
								µcount = πTemp007
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(3)            
							// line 712: other_count = other[elem]
							πF.SetLineno(712)
							if πE = πg.CheckLocal(πF, µelem, "elem"); πE != nil {
								continue
							}
							πTemp003 = µelem
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µother, πTemp003); πE != nil {
								continue
							}
							µother_count = πTemp004
							// line 713: newcount = count if count < other_count else other_count
							πF.SetLineno(713)
							if πE = πg.CheckLocal(πF, µcount, "count"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µother_count, "other_count"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.LT(πF, µcount, µother_count); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if !πTemp006 {
								goto Label6
							}
							if πE = πg.CheckLocal(πF, µcount, "count"); πE != nil {
								continue
							}
							πTemp003 = µcount
							goto Label7
						Label6:
							if πE = πg.CheckLocal(πF, µother_count, "other_count"); πE != nil {
								continue
							}
							πTemp003 = µother_count
						Label7:
							µnewcount = πTemp003
							if πE = πg.CheckLocal(πF, µnewcount, "newcount"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GT(πF, µnewcount, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label8
							}
							goto Label9
							// line 714: if newcount > 0:
							πF.SetLineno(714)
						Label8:
							// line 715: result[elem] = newcount
							πF.SetLineno(715)
							if πE = πg.CheckLocal(πF, µnewcount, "newcount"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, µnewcount); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µelem, "elem"); πE != nil {
								continue
							}
							πTemp004 = µelem
							if πE = πg.SetItem(πF, µresult, πTemp004, πTemp003); πE != nil {
								continue
							}
							goto Label9
						Label9:
							continue
						Label4:
							if πE != nil || πR != nil {
								continue
							}
						Label5:
							// line 716: return result
							πF.SetLineno(716)
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
					if πE = πClass.SetItem(πF, ß__and__.ToObject(), πTemp017); πE != nil {
						continue
					}
					// line 702: ''' Intersection is the minimum of corresponding counts.
					πF.SetLineno(702)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp018}, πg.NewStr(" Intersection is the minimum of corresponding counts.\n\n        >>> Counter('abbb') & Counter('bcc')\n        Counter({'b': 1})\n\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp019, πE = πg.ResolveClass(πF, πClass, nil, ß__and__); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp019, ß__doc__, πTemp018); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp010.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp004, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("Counter").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp010.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßCounter.ToObject(), πTemp004); πE != nil {
				continue
			}
			if πTemp003, πE = πg.ResolveGlobal(πF, ß__name__); πE != nil {
				continue
			}
			if πTemp002, πE = πg.Eq(πF, πTemp003, ß__main__.ToObject()); πE != nil {
				continue
			}
			if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
				continue
			}
			if πTemp005 {
				goto Label4
			}
			goto Label5
			// line 719: if __name__ == '__main__':
			πF.SetLineno(719)
		Label4:
			// line 720: pass
			πF.SetLineno(720)
			goto Label5
		Label5:
		}
		return nil, πE
	})
	πg.RegisterModule("collections", Code)
}

