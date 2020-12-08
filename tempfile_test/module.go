package tempfile_test
import (
	πg "github.com/pygolin/runtime"
	// _ "github.com/pygolin/stdlib/tempfile"
	// _ "github.com/pygolin/stdlib/stat"
	// _ "github.com/pygolin/stdlib/os"
	// _ "github.com/pygolin/stdlib/weetest"
)
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/tempfile_test.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ßAssertionError := πg.InternStr("AssertionError")
		ßOSError := πg.InternStr("OSError")
		ßRunTests := πg.InternStr("RunTests")
		ßS_IMODE := πg.InternStr("S_IMODE")
		ßS_ISDIR := πg.InternStr("S_ISDIR")
		ßTestMkdTemp := πg.InternStr("TestMkdTemp")
		ßTestMkdTempDir := πg.InternStr("TestMkdTempDir")
		ßTestMkdTempOSError := πg.InternStr("TestMkdTempOSError")
		ßTestMkdTempPrefixSuffix := πg.InternStr("TestMkdTempPrefixSuffix")
		ßTestMksTemp := πg.InternStr("TestMksTemp")
		ßTestMksTempDir := πg.InternStr("TestMksTempDir")
		ßTestMksTempOSError := πg.InternStr("TestMksTempOSError")
		ßTestMksTempPerms := πg.InternStr("TestMksTempPerms")
		ßTestMksTempPrefixSuffix := πg.InternStr("TestMksTempPrefixSuffix")
		ß__main__ := πg.InternStr("__main__")
		ß__name__ := πg.InternStr("__name__")
		ßbar := πg.InternStr("bar")
		ßchmod := πg.InternStr("chmod")
		ßclose := πg.InternStr("close")
		ßfdopen := πg.InternStr("fdopen")
		ßfoo := πg.InternStr("foo")
		ßfoobar := πg.InternStr("foobar")
		ßgeteuid := πg.InternStr("geteuid")
		ßmkdtemp := πg.InternStr("mkdtemp")
		ßmkstemp := πg.InternStr("mkstemp")
		ßopen := πg.InternStr("open")
		ßos := πg.InternStr("os")
		ßread := πg.InternStr("read")
		ßremove := πg.InternStr("remove")
		ßrmdir := πg.InternStr("rmdir")
		ßst_mode := πg.InternStr("st_mode")
		ßstartswith := πg.InternStr("startswith")
		ßstat := πg.InternStr("stat")
		ßtempfile := πg.InternStr("tempfile")
		ßw := πg.InternStr("w")
		ßweetest := πg.InternStr("weetest")
		ßwrite := πg.InternStr("write")
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
		var πTemp010 *πg.Object
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
			// line 15: import os
			πF.SetLineno(15)
			if πTemp002, πE = πg.ImportModule(πF, "os"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßos.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 16: import stat
			πF.SetLineno(16)
			if πTemp002, πE = πg.ImportModule(πF, "stat"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßstat.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 17: import tempfile
			πF.SetLineno(17)
			if πTemp002, πE = πg.ImportModule(πF, "tempfile"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßtempfile.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 19: import weetest
			πF.SetLineno(19)
			if πTemp002, πE = πg.ImportModule(πF, "weetest"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßweetest.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 22: def TestMkdTemp():
			πF.SetLineno(22)
			πTemp003 = make([]πg.Param, 0)
			πTemp001 = πg.NewFunction(πg.NewCode("TestMkdTemp", "/usr/lib/python2.7/tempfile_test.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µpath *πg.Object = πg.UnboundLocal; _ = µpath
				var µmode *πg.Object = πg.UnboundLocal; _ = µmode
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
					// line 23: path = tempfile.mkdtemp()
					πF.SetLineno(23)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßtempfile); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßmkdtemp, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					µpath = πTemp001
					// line 24: mode = os.stat(path).st_mode
					πF.SetLineno(24)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					πTemp003[0] = µpath
					if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßstat, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßst_mode, nil); πE != nil {
						continue
					}
					µmode = πTemp002
					// line 25: os.rmdir(path)
					πF.SetLineno(25)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					πTemp003[0] = µpath
					if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßrmdir, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					// line 26: assert stat.S_ISDIR(mode), mode
					πF.SetLineno(26)
					if πE = πg.CheckLocal(πF, µmode, "mode"); πE != nil {
						continue
					}
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µmode, "mode"); πE != nil {
						continue
					}
					πTemp003[0] = µmode
					if πTemp001, πE = πg.ResolveGlobal(πF, ßstat); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßS_ISDIR, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πE = πg.Assert(πF, πTemp001, µmode); πE != nil {
						continue
					}
					// line 27: assert stat.S_IMODE(mode) == 0o700, mode
					πF.SetLineno(27)
					if πE = πg.CheckLocal(πF, µmode, "mode"); πE != nil {
						continue
					}
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µmode, "mode"); πE != nil {
						continue
					}
					πTemp003[0] = µmode
					if πTemp002, πE = πg.ResolveGlobal(πF, ßstat); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßS_IMODE, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp001, πE = πg.Eq(πF, πTemp002, πg.NewInt(448).ToObject()); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp001, µmode); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßTestMkdTemp.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 30: def TestMkdTempDir():
			πF.SetLineno(30)
			πTemp003 = make([]πg.Param, 0)
			πTemp004 = πg.NewFunction(πg.NewCode("TestMkdTempDir", "/usr/lib/python2.7/tempfile_test.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µtempdir *πg.Object = πg.UnboundLocal; _ = µtempdir
				var µpath *πg.Object = πg.UnboundLocal; _ = µpath
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 πg.KWArgs
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
					// line 31: tempdir = tempfile.mkdtemp()
					πF.SetLineno(31)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßtempfile); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßmkdtemp, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					µtempdir = πTemp001
					// line 32: path = tempfile.mkdtemp(dir=tempdir)
					πF.SetLineno(32)
					if πE = πg.CheckLocal(πF, µtempdir, "tempdir"); πE != nil {
						continue
					}
					πTemp003 = πg.KWArgs{
						{"dir", µtempdir},
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ßtempfile); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßmkdtemp, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, nil, πTemp003); πE != nil {
						continue
					}
					µpath = πTemp001
					// line 33: os.rmdir(path)
					πF.SetLineno(33)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					πTemp004[0] = µpath
					if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßrmdir, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 34: os.rmdir(tempdir)
					πF.SetLineno(34)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µtempdir, "tempdir"); πE != nil {
						continue
					}
					πTemp004[0] = µtempdir
					if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßrmdir, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 35: assert path.startswith(tempdir)
					πF.SetLineno(35)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µtempdir, "tempdir"); πE != nil {
						continue
					}
					πTemp004[0] = µtempdir
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µpath, ßstartswith, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πE = πg.Assert(πF, πTemp002, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßTestMkdTempDir.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 38: def TestMkdTempOSError():
			πF.SetLineno(38)
			πTemp003 = make([]πg.Param, 0)
			πTemp005 = πg.NewFunction(πg.NewCode("TestMkdTempOSError", "/usr/lib/python2.7/tempfile_test.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µtempdir *πg.Object = πg.UnboundLocal; _ = µtempdir
				var µmode *πg.Object = πg.UnboundLocal; _ = µmode
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
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 πg.KWArgs
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
					case 5: goto Label5
					default: panic("unexpected function state")
					}
					// line 39: tempdir = tempfile.mkdtemp()
					πF.SetLineno(39)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßtempfile); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßmkdtemp, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					µtempdir = πTemp001
					// line 40: os.chmod(tempdir, 0o500)
					πF.SetLineno(40)
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µtempdir, "tempdir"); πE != nil {
						continue
					}
					πTemp003[0] = µtempdir
					πTemp003[1] = πg.NewInt(320).ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßchmod, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßgeteuid, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, πTemp002, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label1
					}
					goto Label2
					// line 42: if os.geteuid() == 0:
					πF.SetLineno(42)
				Label1:
					// line 43: print ('Warning: Cannot reliable test file readonly-ness with Root user')
					πF.SetLineno(43)
					πTemp003 = make([]*πg.Object, 1)
					πTemp003[0] = πg.NewStr("Warning: Cannot reliable test file readonly-ness with Root user").ToObject()
					if πE = πg.Print(πF, πTemp003, true); πE != nil {
						continue
					}
					// line 44: mode = os.stat(tempdir).st_mode
					πF.SetLineno(44)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µtempdir, "tempdir"); πE != nil {
						continue
					}
					πTemp003[0] = µtempdir
					if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßstat, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßst_mode, nil); πE != nil {
						continue
					}
					µmode = πTemp002
					// line 45: assert stat.S_IMODE(mode) == 0o500, ('Wrong file mode "%s" detected' % mode)
					πF.SetLineno(45)
					if πE = πg.CheckLocal(πF, µmode, "mode"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("Wrong file mode \"%s\" detected").ToObject(), µmode); πE != nil {
						continue
					}
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µmode, "mode"); πE != nil {
						continue
					}
					πTemp003[0] = µmode
					if πTemp004, πE = πg.ResolveGlobal(πF, ßstat); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp004, ßS_IMODE, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp006.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp002, πE = πg.Eq(πF, πTemp004, πg.NewInt(320).ToObject()); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp002, πTemp001); πE != nil {
						continue
					}
					goto Label3
				Label2:
					// line 48: try:
					πF.SetLineno(48)
					πF.PushCheckpoint(5)
					// line 49: tempfile.mkdtemp(dir=tempdir)
					πF.SetLineno(49)
					if πE = πg.CheckLocal(πF, µtempdir, "tempdir"); πE != nil {
						continue
					}
					πTemp007 = πg.KWArgs{
						{"dir", µtempdir},
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ßtempfile); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßmkdtemp, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, nil, πTemp007); πE != nil {
						continue
					}
					πF.PopCheckpoint()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßAssertionError); πE != nil {
						continue
					}
					// line 53: raise AssertionError, 'Should not be able to touch 0o500 paths'
					πF.SetLineno(53)
					πE = πF.Raise(πTemp001, πg.NewStr("Should not be able to touch 0o500 paths").ToObject(), nil)
					continue
					goto Label4
				Label5:
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
						goto Label6
					}
					πE = πF.Raise(πTemp008.ToObject(), nil, πTemp009.ToObject())
					continue
					// line 50: except OSError:
					πF.SetLineno(50)
				Label6:
					// line 51: pass
					πF.SetLineno(51)
					πF.RestoreExc(nil, nil)
					goto Label4
				Label4:
					goto Label3
				Label3:
					// line 54: os.rmdir(tempdir)
					πF.SetLineno(54)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µtempdir, "tempdir"); πE != nil {
						continue
					}
					πTemp003[0] = µtempdir
					if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßrmdir, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßTestMkdTempOSError.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 57: def TestMkdTempPrefixSuffix():
			πF.SetLineno(57)
			πTemp003 = make([]πg.Param, 0)
			πTemp006 = πg.NewFunction(πg.NewCode("TestMkdTempPrefixSuffix", "/usr/lib/python2.7/tempfile_test.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µpath *πg.Object = πg.UnboundLocal; _ = µpath
				var πTemp001 πg.KWArgs
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 []*πg.Object
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
					// line 58: path = tempfile.mkdtemp(prefix='foo', suffix='bar')
					πF.SetLineno(58)
					πTemp001 = πg.KWArgs{
						{"prefix", ßfoo.ToObject()},
						{"suffix", ßbar.ToObject()},
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßtempfile); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßmkdtemp, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, nil, πTemp001); πE != nil {
						continue
					}
					µpath = πTemp002
					// line 59: os.rmdir(path)
					πF.SetLineno(59)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					πTemp004[0] = µpath
					if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßrmdir, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 60: assert 'foo' in path
					πF.SetLineno(60)
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Contains(πF, µpath, ßfoo.ToObject()); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(πTemp005).ToObject()
					if πE = πg.Assert(πF, πTemp002, nil); πE != nil {
						continue
					}
					// line 61: assert 'bar' in path
					πF.SetLineno(61)
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Contains(πF, µpath, ßbar.ToObject()); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(πTemp005).ToObject()
					if πE = πg.Assert(πF, πTemp002, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßTestMkdTempPrefixSuffix.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 65: def TestMksTemp():
			πF.SetLineno(65)
			πTemp003 = make([]πg.Param, 0)
			πTemp007 = πg.NewFunction(πg.NewCode("TestMksTemp", "/usr/lib/python2.7/tempfile_test.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µfd *πg.Object = πg.UnboundLocal; _ = µfd
				var µpath *πg.Object = πg.UnboundLocal; _ = µpath
				var µf *πg.Object = πg.UnboundLocal; _ = µf
				var µcontents *πg.Object = πg.UnboundLocal; _ = µcontents
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
					// line 66: fd, path = tempfile.mkstemp()
					πF.SetLineno(66)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßtempfile); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßmkstemp, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp003}}}, πTemp001); πE != nil {
						continue
					}
					µfd = πTemp002
					µpath = πTemp003
					// line 67: f = os.fdopen(fd, 'w')
					πF.SetLineno(67)
					πTemp004 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µfd, "fd"); πE != nil {
						continue
					}
					πTemp004[0] = µfd
					πTemp004[1] = ßw.ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßfdopen, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µf = πTemp001
					// line 68: f.write('foobar')
					πF.SetLineno(68)
					πTemp004 = πF.MakeArgs(1)
					πTemp004[0] = ßfoobar.ToObject()
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µf, ßwrite, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 69: f.close()
					πF.SetLineno(69)
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µf, ßclose, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					// line 70: f = open(path)
					πF.SetLineno(70)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					πTemp004[0] = µpath
					if πTemp001, πE = πg.ResolveGlobal(πF, ßopen); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µf = πTemp002
					// line 71: contents = f.read()
					πF.SetLineno(71)
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µf, ßread, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					µcontents = πTemp002
					// line 72: f.close()
					πF.SetLineno(72)
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µf, ßclose, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					// line 73: os.remove(path)
					πF.SetLineno(73)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					πTemp004[0] = µpath
					if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßremove, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 74: assert contents == 'foobar', contents
					πF.SetLineno(74)
					if πE = πg.CheckLocal(πF, µcontents, "contents"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcontents, "contents"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, µcontents, ßfoobar.ToObject()); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp001, µcontents); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßTestMksTemp.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 77: def TestMksTempDir():
			πF.SetLineno(77)
			πTemp003 = make([]πg.Param, 0)
			πTemp008 = πg.NewFunction(πg.NewCode("TestMksTempDir", "/usr/lib/python2.7/tempfile_test.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µtempdir *πg.Object = πg.UnboundLocal; _ = µtempdir
				var µfd *πg.Object = πg.UnboundLocal; _ = µfd
				var µpath *πg.Object = πg.UnboundLocal; _ = µpath
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 πg.KWArgs
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
					// line 78: tempdir = tempfile.mkdtemp()
					πF.SetLineno(78)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßtempfile); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßmkdtemp, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					µtempdir = πTemp001
					// line 79: fd, path = tempfile.mkstemp(dir=tempdir)
					πF.SetLineno(79)
					if πE = πg.CheckLocal(πF, µtempdir, "tempdir"); πE != nil {
						continue
					}
					πTemp003 = πg.KWArgs{
						{"dir", µtempdir},
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ßtempfile); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßmkstemp, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, nil, πTemp003); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp004}}}, πTemp001); πE != nil {
						continue
					}
					µfd = πTemp002
					µpath = πTemp004
					// line 80: os.close(fd)
					πF.SetLineno(80)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfd, "fd"); πE != nil {
						continue
					}
					πTemp005[0] = µfd
					if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßclose, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					// line 81: os.remove(path)
					πF.SetLineno(81)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					πTemp005[0] = µpath
					if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßremove, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					// line 82: os.rmdir(tempdir)
					πF.SetLineno(82)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µtempdir, "tempdir"); πE != nil {
						continue
					}
					πTemp005[0] = µtempdir
					if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßrmdir, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					// line 83: assert path.startswith(tempdir)
					πF.SetLineno(83)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µtempdir, "tempdir"); πE != nil {
						continue
					}
					πTemp005[0] = µtempdir
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µpath, ßstartswith, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πE = πg.Assert(πF, πTemp002, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßTestMksTempDir.ToObject(), πTemp008); πE != nil {
				continue
			}
			// line 86: def TestMksTempOSError():
			πF.SetLineno(86)
			πTemp003 = make([]πg.Param, 0)
			πTemp009 = πg.NewFunction(πg.NewCode("TestMksTempOSError", "/usr/lib/python2.7/tempfile_test.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µtempdir *πg.Object = πg.UnboundLocal; _ = µtempdir
				var µmode *πg.Object = πg.UnboundLocal; _ = µmode
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
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 πg.KWArgs
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
					case 5: goto Label5
					default: panic("unexpected function state")
					}
					// line 87: tempdir = tempfile.mkdtemp()
					πF.SetLineno(87)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßtempfile); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßmkdtemp, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					µtempdir = πTemp001
					// line 88: os.chmod(tempdir, 0o500)
					πF.SetLineno(88)
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µtempdir, "tempdir"); πE != nil {
						continue
					}
					πTemp003[0] = µtempdir
					πTemp003[1] = πg.NewInt(320).ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßchmod, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßgeteuid, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, πTemp002, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label1
					}
					goto Label2
					// line 90: if os.geteuid() == 0:
					πF.SetLineno(90)
				Label1:
					// line 91: print ('Warning: Cannot reliable test file readonly-ness with Root user')
					πF.SetLineno(91)
					πTemp003 = make([]*πg.Object, 1)
					πTemp003[0] = πg.NewStr("Warning: Cannot reliable test file readonly-ness with Root user").ToObject()
					if πE = πg.Print(πF, πTemp003, true); πE != nil {
						continue
					}
					// line 92: mode = os.stat(tempdir).st_mode
					πF.SetLineno(92)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µtempdir, "tempdir"); πE != nil {
						continue
					}
					πTemp003[0] = µtempdir
					if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßstat, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßst_mode, nil); πE != nil {
						continue
					}
					µmode = πTemp002
					// line 93: assert stat.S_IMODE(mode) == 0o500, ('Wrong file mode "%s" detected' % mode)
					πF.SetLineno(93)
					if πE = πg.CheckLocal(πF, µmode, "mode"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("Wrong file mode \"%s\" detected").ToObject(), µmode); πE != nil {
						continue
					}
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µmode, "mode"); πE != nil {
						continue
					}
					πTemp003[0] = µmode
					if πTemp004, πE = πg.ResolveGlobal(πF, ßstat); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp004, ßS_IMODE, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp006.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp002, πE = πg.Eq(πF, πTemp004, πg.NewInt(320).ToObject()); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp002, πTemp001); πE != nil {
						continue
					}
					goto Label3
				Label2:
					// line 96: try:
					πF.SetLineno(96)
					πF.PushCheckpoint(5)
					// line 97: tempfile.mkstemp(dir=tempdir)
					πF.SetLineno(97)
					if πE = πg.CheckLocal(πF, µtempdir, "tempdir"); πE != nil {
						continue
					}
					πTemp007 = πg.KWArgs{
						{"dir", µtempdir},
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ßtempfile); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßmkstemp, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, nil, πTemp007); πE != nil {
						continue
					}
					πF.PopCheckpoint()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßAssertionError); πE != nil {
						continue
					}
					// line 101: raise AssertionError
					πF.SetLineno(101)
					πE = πF.Raise(πTemp001, nil, nil)
					continue
					goto Label4
				Label5:
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
						goto Label6
					}
					πE = πF.Raise(πTemp008.ToObject(), nil, πTemp009.ToObject())
					continue
					// line 98: except OSError:
					πF.SetLineno(98)
				Label6:
					// line 99: pass
					πF.SetLineno(99)
					πF.RestoreExc(nil, nil)
					goto Label4
				Label4:
					goto Label3
				Label3:
					// line 102: os.rmdir(tempdir)
					πF.SetLineno(102)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µtempdir, "tempdir"); πE != nil {
						continue
					}
					πTemp003[0] = µtempdir
					if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßrmdir, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßTestMksTempOSError.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 105: def TestMksTempPerms():
			πF.SetLineno(105)
			πTemp003 = make([]πg.Param, 0)
			πTemp010 = πg.NewFunction(πg.NewCode("TestMksTempPerms", "/usr/lib/python2.7/tempfile_test.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µfd *πg.Object = πg.UnboundLocal; _ = µfd
				var µpath *πg.Object = πg.UnboundLocal; _ = µpath
				var µmode *πg.Object = πg.UnboundLocal; _ = µmode
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
					// line 106: fd, path = tempfile.mkstemp()
					πF.SetLineno(106)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßtempfile); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßmkstemp, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp003}}}, πTemp001); πE != nil {
						continue
					}
					µfd = πTemp002
					µpath = πTemp003
					// line 107: os.close(fd)
					πF.SetLineno(107)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfd, "fd"); πE != nil {
						continue
					}
					πTemp004[0] = µfd
					if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßclose, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 108: mode = os.stat(path).st_mode
					πF.SetLineno(108)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					πTemp004[0] = µpath
					if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßstat, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßst_mode, nil); πE != nil {
						continue
					}
					µmode = πTemp002
					// line 109: os.remove(path)
					πF.SetLineno(109)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					πTemp004[0] = µpath
					if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßremove, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 110: assert stat.S_IMODE(mode) == 0o600, mode
					πF.SetLineno(110)
					if πE = πg.CheckLocal(πF, µmode, "mode"); πE != nil {
						continue
					}
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µmode, "mode"); πE != nil {
						continue
					}
					πTemp004[0] = µmode
					if πTemp002, πE = πg.ResolveGlobal(πF, ßstat); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßS_IMODE, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp001, πE = πg.Eq(πF, πTemp002, πg.NewInt(384).ToObject()); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp001, µmode); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßTestMksTempPerms.ToObject(), πTemp010); πE != nil {
				continue
			}
			// line 113: def TestMksTempPrefixSuffix():
			πF.SetLineno(113)
			πTemp003 = make([]πg.Param, 0)
			πTemp011 = πg.NewFunction(πg.NewCode("TestMksTempPrefixSuffix", "/usr/lib/python2.7/tempfile_test.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µfd *πg.Object = πg.UnboundLocal; _ = µfd
				var µpath *πg.Object = πg.UnboundLocal; _ = µpath
				var πTemp001 πg.KWArgs
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
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
					default: panic("unexpected function state")
					}
					// line 114: fd, path = tempfile.mkstemp(prefix='foo', suffix='bar')
					πF.SetLineno(114)
					πTemp001 = πg.KWArgs{
						{"prefix", ßfoo.ToObject()},
						{"suffix", ßbar.ToObject()},
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßtempfile); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßmkstemp, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, nil, πTemp001); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp004}}}, πTemp002); πE != nil {
						continue
					}
					µfd = πTemp003
					µpath = πTemp004
					// line 115: os.close(fd)
					πF.SetLineno(115)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfd, "fd"); πE != nil {
						continue
					}
					πTemp005[0] = µfd
					if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßclose, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					// line 116: os.remove(path)
					πF.SetLineno(116)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					πTemp005[0] = µpath
					if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßremove, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					// line 117: assert 'foo' in path
					πF.SetLineno(117)
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.Contains(πF, µpath, ßfoo.ToObject()); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(πTemp006).ToObject()
					if πE = πg.Assert(πF, πTemp002, nil); πE != nil {
						continue
					}
					// line 118: assert 'bar' in path
					πF.SetLineno(118)
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.Contains(πF, µpath, ßbar.ToObject()); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(πTemp006).ToObject()
					if πE = πg.Assert(πF, πTemp002, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßTestMksTempPrefixSuffix.ToObject(), πTemp011); πE != nil {
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
			// line 122: if __name__ == '__main__':
			πF.SetLineno(122)
		Label1:
			// line 123: weetest.RunTests()
			πF.SetLineno(123)
			if πTemp012, πE = πg.ResolveGlobal(πF, ßweetest); πE != nil {
				continue
			}
			if πTemp013, πE = πg.GetAttr(πF, πTemp012, ßRunTests, nil); πE != nil {
				continue
			}
			if πTemp012, πE = πTemp013.Call(πF, nil, nil); πE != nil {
				continue
			}
			goto Label2
		Label2:
		}
		return nil, πE
	})
	πg.RegisterModule("tempfile_test", Code)
}

