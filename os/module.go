package os
import (
	πg "github.com/pygolin/runtime"
	// _ "github.com/pygolin/stdlib/__go__/time"
	// _ "github.com/pygolin/stdlib/stat"
	// _ "github.com/pygolin/stdlib/__go__/io/ioutil"
	// _ "github.com/pygolin/stdlib/__go__/reflect"
	// _ "github.com/pygolin/stdlib/_syscall"
	// _ "github.com/pygolin/stdlib/sys"
	// _ "github.com/pygolin/stdlib/__go__/pygolin"
	// _ "github.com/pygolin/stdlib/__go__/path/filepath"
	// _ "github.com/pygolin/stdlib/__go__/sync"
	// _ "github.com/pygolin/stdlib/__go__/runtime"
	// _ "github.com/pygolin/stdlib/__go__/os"
	// _ "github.com/pygolin/stdlib/__go__/syscall"
	// _ "github.com/pygolin/stdlib/os/path"
)
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/os/__init__.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ßAdd := πg.InternStr("Add")
		ßChdir := πg.InternStr("Chdir")
		ßChmod := πg.InternStr("Chmod")
		ßClose := πg.InternStr("Close")
		ßDone := πg.InternStr("Done")
		ßEnviron := πg.InternStr("Environ")
		ßError := πg.InternStr("Error")
		ßExitStatus := πg.InternStr("ExitStatus")
		ßF_GETFD := πg.InternStr("F_GETFD")
		ßFd := πg.InternStr("Fd")
		ßFiles := πg.InternStr("Files")
		ßGOOS := πg.InternStr("GOOS")
		ßGetegid := πg.InternStr("Getegid")
		ßGeteuid := πg.InternStr("Geteuid")
		ßGetgid := πg.InternStr("Getgid")
		ßGetpid := πg.InternStr("Getpid")
		ßGetuid := πg.InternStr("Getuid")
		ßGetwd := πg.InternStr("Getwd")
		ßIn := πg.InternStr("In")
		ßInterface := πg.InternStr("Interface")
		ßMakeSlice := πg.InternStr("MakeSlice")
		ßMkdir := πg.InternStr("Mkdir")
		ßModTime := πg.InternStr("ModTime")
		ßMode := πg.InternStr("Mode")
		ßName := πg.InternStr("Name")
		ßNewFileFromFD := πg.InternStr("NewFileFromFD")
		ßNone := πg.InternStr("None")
		ßNotImplementedError := πg.InternStr("NotImplementedError")
		ßOSError := πg.InternStr("OSError")
		ßPipe := πg.InternStr("Pipe")
		ßProcAttr := πg.InternStr("ProcAttr")
		ßReadDir := πg.InternStr("ReadDir")
		ßRemove := πg.InternStr("Remove")
		ßSHELL := πg.InternStr("SHELL")
		ßSYS_FCNTL := πg.InternStr("SYS_FCNTL")
		ßS_ISDIR := πg.InternStr("S_ISDIR")
		ßSecond := πg.InternStr("Second")
		ßSeparator := πg.InternStr("Separator")
		ßSignal := πg.InternStr("Signal")
		ßSize := πg.InternStr("Size")
		ßStartProcess := πg.InternStr("StartProcess")
		ßStartThread := πg.InternStr("StartThread")
		ßStat := πg.InternStr("Stat")
		ßStatResult := πg.InternStr("StatResult")
		ßStderr := πg.InternStr("Stderr")
		ßStdin := πg.InternStr("Stdin")
		ßStdout := πg.InternStr("Stdout")
		ßSys := πg.InternStr("Sys")
		ßSyscall := πg.InternStr("Syscall")
		ßToNative := πg.InternStr("ToNative")
		ßType := πg.InternStr("Type")
		ßUnixNano := πg.InternStr("UnixNano")
		ßValueError := πg.InternStr("ValueError")
		ßWNOHANG := πg.InternStr("WNOHANG")
		ßWait := πg.InternStr("Wait")
		ßWait4 := πg.InternStr("Wait4")
		ßWaitGroup := πg.InternStr("WaitGroup")
		ßWaitStatus := πg.InternStr("WaitStatus")
		ß_Popen := πg.InternStr("_Popen")
		ß__doc__ := πg.InternStr("__doc__")
		ß__frame__ := πg.InternStr("__frame__")
		ß__init__ := πg.InternStr("__init__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß_encode_wait_result := πg.InternStr("_encode_wait_result")
		ß_info := πg.InternStr("_info")
		ß_syscall := πg.InternStr("_syscall")
		ß_thread_func := πg.InternStr("_thread_func")
		ßchdir := πg.InternStr("chdir")
		ßchmod := πg.InternStr("chmod")
		ßchr := πg.InternStr("chr")
		ßclose := πg.InternStr("close")
		ßctermid := πg.InternStr("ctermid")
		ßcurdir := πg.InternStr("curdir")
		ßenviron := πg.InternStr("environ")
		ßerror := πg.InternStr("error")
		ßfchdir := πg.InternStr("fchdir")
		ßfdopen := πg.InternStr("fdopen")
		ßfile := πg.InternStr("file")
		ßfloat := πg.InternStr("float")
		ßget := πg.InternStr("get")
		ßgetandroidapilevel := πg.InternStr("getandroidapilevel")
		ßgetcwd := πg.InternStr("getcwd")
		ßgetegid := πg.InternStr("getegid")
		ßgeteuid := πg.InternStr("geteuid")
		ßgetgid := πg.InternStr("getgid")
		ßgetpid := πg.InternStr("getpid")
		ßgetuid := πg.InternStr("getuid")
		ßhasattr := πg.InternStr("hasattr")
		ßinvoke := πg.InternStr("invoke")
		ßk := πg.InternStr("k")
		ßlistdir := πg.InternStr("listdir")
		ßmkdir := πg.InternStr("mkdir")
		ßmode := πg.InternStr("mode")
		ßname := πg.InternStr("name")
		ßnew := πg.InternStr("new")
		ßobject := πg.InternStr("object")
		ßpath := πg.InternStr("path")
		ßpopen := πg.InternStr("popen")
		ßposix := πg.InternStr("posix")
		ßproc := πg.InternStr("proc")
		ßproperty := πg.InternStr("property")
		ßr := πg.InternStr("r")
		ßremove := πg.InternStr("remove")
		ßresult := πg.InternStr("result")
		ßrmdir := πg.InternStr("rmdir")
		ßsep := πg.InternStr("sep")
		ßsplit := πg.InternStr("split")
		ßst_mode := πg.InternStr("st_mode")
		ßst_mtime := πg.InternStr("st_mtime")
		ßst_size := πg.InternStr("st_size")
		ßstat := πg.InternStr("stat")
		ßstat_module := πg.InternStr("stat_module")
		ßsys := πg.InternStr("sys")
		ßunlink := πg.InternStr("unlink")
		ßv := πg.InternStr("v")
		ßvar := πg.InternStr("var")
		ßw := πg.InternStr("w")
		ßwaitpid := πg.InternStr("waitpid")
		ßwg := πg.InternStr("wg")
		var πTemp001 *πg.Object
		_ = πTemp001
		var πTemp002 []*πg.Object
		_ = πTemp002
		var πTemp003 *πg.Object
		_ = πTemp003
		var πTemp004 *πg.Dict
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
		var πTemp011 []πg.Param
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
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			case 1: goto Label1
			case 2: goto Label2
			default: panic("unexpected function state")
			}
			// line 15: """Miscellaneous operating system interfaces."""
			πF.SetLineno(15)
			// line 15: """Miscellaneous operating system interfaces."""
			πF.SetLineno(15)
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Miscellaneous operating system interfaces.").ToObject()); πE != nil {
				continue
			}
			// line 18: from '__go__/io/ioutil' import ReadDir
			πF.SetLineno(18)
			if πTemp002, πE = πg.ImportModule(πF, "__go__/io/ioutil"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßReadDir); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßReadDir.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 19: from '__go__/os' import (Chdir, Chmod, Environ, Getpid as getpid, Getuid as getuid,
			πF.SetLineno(19)
			if πTemp002, πE = πg.ImportModule(πF, "__go__/os"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßChdir); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßChdir.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßChmod); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßChmod.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßEnviron); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßEnviron.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßGetpid); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßgetpid.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßGetuid); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßgetuid.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßGeteuid); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßgeteuid.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßGetgid); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßgetgid.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßGetegid); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßgetegid.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßGetwd); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßGetwd.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßPipe); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßPipe.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßProcAttr); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßProcAttr.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßRemove); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßRemove.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßStartProcess); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßStartProcess.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßStat); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßStat.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßStdout); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßStdout.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßStdin); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßStdin.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßStderr); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßStderr.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßMkdir); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßMkdir.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 23: from '__go__/path/filepath' import Separator
			πF.SetLineno(23)
			if πTemp002, πE = πg.ImportModule(πF, "__go__/path/filepath"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßSeparator); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßSeparator.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 24: from '__go__/pygolin' import (NewFileFromFD, StartThread, ToNative)
			πF.SetLineno(24)
			if πTemp002, πE = πg.ImportModule(πF, "__go__/pygolin"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßNewFileFromFD); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßNewFileFromFD.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßStartThread); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßStartThread.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßToNative); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßToNative.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 25: from '__go__/reflect' import MakeSlice
			πF.SetLineno(25)
			if πTemp002, πE = πg.ImportModule(πF, "__go__/reflect"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßMakeSlice); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßMakeSlice.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 26: from '__go__/runtime' import GOOS
			πF.SetLineno(26)
			if πTemp002, πE = πg.ImportModule(πF, "__go__/runtime"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßGOOS); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßGOOS.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 27: from '__go__/syscall' import (Close, SYS_FCNTL, Syscall, F_GETFD, Wait4,
			πF.SetLineno(27)
			if πTemp002, πE = πg.ImportModule(πF, "__go__/syscall"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßClose); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßClose.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßSYS_FCNTL); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßSYS_FCNTL.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßSyscall); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßSyscall.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßF_GETFD); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßF_GETFD.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßWait4); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßWait4.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßWaitStatus); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßWaitStatus.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßWNOHANG); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßWNOHANG.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 29: from '__go__/sync' import WaitGroup
			πF.SetLineno(29)
			if πTemp002, πE = πg.ImportModule(πF, "__go__/sync"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßWaitGroup); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßWaitGroup.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 30: from '__go__/time' import Second
			πF.SetLineno(30)
			if πTemp002, πE = πg.ImportModule(πF, "__go__/time"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßSecond); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßSecond.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 31: import _syscall
			πF.SetLineno(31)
			if πTemp002, πE = πg.ImportModule(πF, "_syscall"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ß_syscall.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 32: from os import path
			πF.SetLineno(32)
			if πTemp002, πE = πg.ImportModule(πF, "os.path"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πE = πF.Globals().SetItem(πF, ßpath.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 33: import stat as stat_module
			πF.SetLineno(33)
			if πTemp002, πE = πg.ImportModule(πF, "stat"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßstat_module.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 34: import sys
			πF.SetLineno(34)
			if πTemp002, πE = πg.ImportModule(πF, "sys"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßsys.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 37: sep = chr(Separator)
			πF.SetLineno(37)
			πTemp002 = πF.MakeArgs(1)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßSeparator); πE != nil {
				continue
			}
			πTemp002[0] = πTemp001
			if πTemp001, πE = πg.ResolveGlobal(πF, ßchr); πE != nil {
				continue
			}
			if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πF.Globals().SetItem(πF, ßsep.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 38: error = OSError  # pylint: disable=invalid-name
			πF.SetLineno(38)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßOSError); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßerror.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 39: curdir = '.'
			πF.SetLineno(39)
			if πE = πF.Globals().SetItem(πF, ßcurdir.ToObject(), πg.NewStr(".").ToObject()); πE != nil {
				continue
			}
			// line 40: name = 'posix'
			πF.SetLineno(40)
			if πE = πF.Globals().SetItem(πF, ßname.ToObject(), ßposix.ToObject()); πE != nil {
				continue
			}
			// line 43: environ = {}
			πF.SetLineno(43)
			πTemp004 = πg.NewDict()
			πTemp001 = πTemp004.ToObject()
			if πE = πF.Globals().SetItem(πF, ßenviron.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßEnviron); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp003.Call(πF, nil, nil); πE != nil {
				continue
			}
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
				if πE = πF.Globals().SetItem(πF, ßvar.ToObject(), πTemp003); πE != nil {
					continue
				}
			}
			if πE != nil || !πTemp007 {
				continue
			}
			πF.PushCheckpoint(1)            
			// line 45: k, v = var.split('=', 1)
			πF.SetLineno(45)
			πTemp002 = πF.MakeArgs(2)
			πTemp002[0] = πg.NewStr("=").ToObject()
			πTemp002[1] = πg.NewInt(1).ToObject()
			if πTemp003, πE = πg.ResolveGlobal(πF, ßvar); πE != nil {
				continue
			}
			if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßsplit, nil); πE != nil {
				continue
			}
			if πTemp003, πE = πTemp005.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp008}}}, πTemp003); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßk.ToObject(), πTemp005); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßv.ToObject(), πTemp008); πE != nil {
				continue
			}
			// line 46: environ[k] = v
			πF.SetLineno(46)
			if πTemp003, πE = πg.ResolveGlobal(πF, ßv); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp003); πE != nil {
				continue
			}
			if πTemp008, πE = πg.ResolveGlobal(πF, ßenviron); πE != nil {
				continue
			}
			if πTemp010, πE = πg.ResolveGlobal(πF, ßk); πE != nil {
				continue
			}
			πTemp009 = πTemp010
			if πE = πg.SetItem(πF, πTemp008, πTemp009, πTemp005); πE != nil {
				continue
			}
			continue
		Label2:
			if πE != nil || πR != nil {
				continue
			}
		Label3:
			// line 49: def mkdir(path, mode=0o777):
			πF.SetLineno(49)
			πTemp011 = make([]πg.Param, 2)
			πTemp011[0] = πg.Param{Name: "path", Def: nil}
			πTemp011[1] = πg.Param{Name: "mode", Def: πg.NewInt(511).ToObject()}
			πTemp001 = πg.NewFunction(πg.NewCode("mkdir", "/usr/lib/python2.7/os/__init__.py", πTemp011, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µpath *πg.Object = πArgs[0]; _ = µpath
				var µmode *πg.Object = πArgs[1]; _ = µmode
				var µerr *πg.Object = πg.UnboundLocal; _ = µerr
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
					// line 50: err = Mkdir(path, mode)
					πF.SetLineno(50)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					πTemp001[0] = µpath
					if πE = πg.CheckLocal(πF, µmode, "mode"); πE != nil {
						continue
					}
					πTemp001[1] = µmode
					if πTemp002, πE = πg.ResolveGlobal(πF, ßMkdir); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µerr = πTemp003
					if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, µerr); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label1
					}
					goto Label2
					// line 51: if err:
					πF.SetLineno(51)
				Label1:
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µerr, ßError, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					if πTemp002, πE = πg.ResolveGlobal(πF, ßOSError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 52: raise OSError(err.Error())
					πF.SetLineno(52)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
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
			if πE = πF.Globals().SetItem(πF, ßmkdir.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 55: def chdir(path):
			πF.SetLineno(55)
			πTemp011 = make([]πg.Param, 1)
			πTemp011[0] = πg.Param{Name: "path", Def: nil}
			πTemp003 = πg.NewFunction(πg.NewCode("chdir", "/usr/lib/python2.7/os/__init__.py", πTemp011, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µpath *πg.Object = πArgs[0]; _ = µpath
				var µerr *πg.Object = πg.UnboundLocal; _ = µerr
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
					// line 56: err = Chdir(path)
					πF.SetLineno(56)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					πTemp001[0] = µpath
					if πTemp002, πE = πg.ResolveGlobal(πF, ßChdir); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µerr = πTemp003
					if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, µerr); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label1
					}
					goto Label2
					// line 57: if err:
					πF.SetLineno(57)
				Label1:
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µerr, ßError, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					if πTemp002, πE = πg.ResolveGlobal(πF, ßOSError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 58: raise OSError(err.Error())
					πF.SetLineno(58)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
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
			if πE = πF.Globals().SetItem(πF, ßchdir.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 61: def fchdir(fd):
			πF.SetLineno(61)
			πTemp011 = make([]πg.Param, 1)
			πTemp011[0] = πg.Param{Name: "fd", Def: nil}
			πTemp005 = πg.NewFunction(πg.NewCode("fchdir", "/usr/lib/python2.7/os/__init__.py", πTemp011, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µfd *πg.Object = πArgs[0]; _ = µfd
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
					πTemp001[0] = πg.NewStr("Code contributions to implement 'os.fchdir(fd)' are welcome").ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNotImplementedError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 62: raise NotImplementedError("Code contributions to implement 'os.fchdir(fd)' are welcome")
					πF.SetLineno(62)
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
			if πE = πF.Globals().SetItem(πF, ßfchdir.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 65: def ctermid():
			πF.SetLineno(65)
			πTemp011 = make([]πg.Param, 0)
			πTemp008 = πg.NewFunction(πg.NewCode("ctermid", "/usr/lib/python2.7/os/__init__.py", πTemp011, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					πTemp001[0] = πg.NewStr("Code contributions to implement 'os.ctermid()' are welcome").ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNotImplementedError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 66: raise NotImplementedError("Code contributions to implement 'os.ctermid()' are welcome")
					πF.SetLineno(66)
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
			if πE = πF.Globals().SetItem(πF, ßctermid.ToObject(), πTemp008); πE != nil {
				continue
			}
			// line 69: def chmod(filepath, mode):
			πF.SetLineno(69)
			πTemp011 = make([]πg.Param, 2)
			πTemp011[0] = πg.Param{Name: "filepath", Def: nil}
			πTemp011[1] = πg.Param{Name: "mode", Def: nil}
			πTemp009 = πg.NewFunction(πg.NewCode("chmod", "/usr/lib/python2.7/os/__init__.py", πTemp011, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µfilepath *πg.Object = πArgs[0]; _ = µfilepath
				var µmode *πg.Object = πArgs[1]; _ = µmode
				var µerr *πg.Object = πg.UnboundLocal; _ = µerr
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
				var πTemp007 bool
				_ = πTemp007
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 71: err = Chmod(filepath, stat(filepath).st_mode & ~0o777 | mode & 0o777)
					πF.SetLineno(71)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µfilepath, "filepath"); πE != nil {
						continue
					}
					πTemp001[0] = µfilepath
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfilepath, "filepath"); πE != nil {
						continue
					}
					πTemp004[0] = µfilepath
					if πTemp005, πE = πg.ResolveGlobal(πF, ßstat); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp005, πE = πg.GetAttr(πF, πTemp006, ßst_mode, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πg.Invert(πF, πg.NewInt(511).ToObject()); πE != nil {
						continue
					}
					if πTemp003, πE = πg.And(πF, πTemp005, πTemp006); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µmode, "mode"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.And(πF, µmode, πg.NewInt(511).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Or(πF, πTemp003, πTemp005); πE != nil {
						continue
					}
					πTemp001[1] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßChmod); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µerr = πTemp003
					if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsTrue(πF, µerr); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label1
					}
					goto Label2
					// line 72: if err:
					πF.SetLineno(72)
				Label1:
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µerr, ßError, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					if πTemp002, πE = πg.ResolveGlobal(πF, ßOSError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 73: raise OSError(err.Error())
					πF.SetLineno(73)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
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
			if πE = πF.Globals().SetItem(πF, ßchmod.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 76: def close(fd):
			πF.SetLineno(76)
			πTemp011 = make([]πg.Param, 1)
			πTemp011[0] = πg.Param{Name: "fd", Def: nil}
			πTemp010 = πg.NewFunction(πg.NewCode("close", "/usr/lib/python2.7/os/__init__.py", πTemp011, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µfd *πg.Object = πArgs[0]; _ = µfd
				var µerr *πg.Object = πg.UnboundLocal; _ = µerr
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
					// line 77: err = Close(fd)
					πF.SetLineno(77)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfd, "fd"); πE != nil {
						continue
					}
					πTemp001[0] = µfd
					if πTemp002, πE = πg.ResolveGlobal(πF, ßClose); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µerr = πTemp003
					if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, µerr); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label1
					}
					goto Label2
					// line 78: if err:
					πF.SetLineno(78)
				Label1:
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µerr, ßError, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					if πTemp002, πE = πg.ResolveGlobal(πF, ßOSError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 79: raise OSError(err.Error())
					πF.SetLineno(79)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
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
			if πE = πF.Globals().SetItem(πF, ßclose.ToObject(), πTemp010); πE != nil {
				continue
			}
			// line 82: def fdopen(fd, mode='r'):  # pylint: disable=unused-argument
			πF.SetLineno(82)
			πTemp011 = make([]πg.Param, 2)
			πTemp011[0] = πg.Param{Name: "fd", Def: nil}
			πTemp011[1] = πg.Param{Name: "mode", Def: ßr.ToObject()}
			πTemp012 = πg.NewFunction(πg.NewCode("fdopen", "/usr/lib/python2.7/os/__init__.py", πTemp011, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µfd *πg.Object = πArgs[0]; _ = µfd
				var µmode *πg.Object = πArgs[1]; _ = µmode
				var µ_ *πg.Object = πg.UnboundLocal; _ = µ_
				var µerr *πg.Object = πg.UnboundLocal; _ = µerr
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
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 84: _, _, err = Syscall(SYS_FCNTL, fd, F_GETFD, 0)
					πF.SetLineno(84)
					πTemp001 = πF.MakeArgs(4)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßSYS_FCNTL); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πE = πg.CheckLocal(πF, µfd, "fd"); πE != nil {
						continue
					}
					πTemp001[1] = µfd
					if πTemp002, πE = πg.ResolveGlobal(πF, ßF_GETFD); πE != nil {
						continue
					}
					πTemp001[2] = πTemp002
					πTemp001[3] = πg.NewInt(0).ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßSyscall); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}}}, πTemp003); πE != nil {
						continue
					}
					µ_ = πTemp002
					µ_ = πTemp004
					µerr = πTemp005
					if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, µerr); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label1
					}
					goto Label2
					// line 85: if err:
					πF.SetLineno(85)
				Label1:
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µerr, ßError, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					if πTemp002, πE = πg.ResolveGlobal(πF, ßOSError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 86: raise OSError(err.Error())
					πF.SetLineno(86)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					goto Label2
				Label2:
					// line 87: return NewFileFromFD(fd, None)
					πF.SetLineno(87)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µfd, "fd"); πE != nil {
						continue
					}
					πTemp001[0] = µfd
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001[1] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNewFileFromFD); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßfdopen.ToObject(), πTemp012); πE != nil {
				continue
			}
			// line 90: def listdir(p):
			πF.SetLineno(90)
			πTemp011 = make([]πg.Param, 1)
			πTemp011[0] = πg.Param{Name: "p", Def: nil}
			πTemp013 = πg.NewFunction(πg.NewCode("listdir", "/usr/lib/python2.7/os/__init__.py", πTemp011, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µp *πg.Object = πArgs[0]; _ = µp
				var µfiles *πg.Object = πg.UnboundLocal; _ = µfiles
				var µerr *πg.Object = πg.UnboundLocal; _ = µerr
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
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 91: files, err = ReadDir(p)
					πF.SetLineno(91)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
						continue
					}
					πTemp001[0] = µp
					if πTemp002, πE = πg.ResolveGlobal(πF, ßReadDir); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp004}}}, πTemp003); πE != nil {
						continue
					}
					µfiles = πTemp002
					µerr = πTemp004
					if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, µerr); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label1
					}
					goto Label2
					// line 92: if err:
					πF.SetLineno(92)
				Label1:
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µerr, ßError, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					if πTemp002, πE = πg.ResolveGlobal(πF, ßOSError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 93: raise OSError(err.Error())
					πF.SetLineno(93)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					goto Label2
				Label2:
					// line 94: return [x.Name() for x in files]
					πF.SetLineno(94)
					πTemp006 = make([]πg.Param, 0)
					πTemp003 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/os/__init__.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
								if πE = πg.CheckLocal(πF, µfiles, "files"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.Iter(πF, µfiles); πE != nil {
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
									µx = πTemp004
								}
								if πE != nil || !πTemp003 {
									continue
								}
								πF.PushCheckpoint(1)            
								// line 94: return [x.Name() for x in files]
								πF.SetLineno(94)
								if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
									continue
								}
								if πTemp004, πE = πg.GetAttr(πF, µx, ßName, nil); πE != nil {
									continue
								}
								if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßlistdir.ToObject(), πTemp013); πE != nil {
				continue
			}
			// line 97: def getcwd():
			πF.SetLineno(97)
			πTemp011 = make([]πg.Param, 0)
			πTemp014 = πg.NewFunction(πg.NewCode("getcwd", "/usr/lib/python2.7/os/__init__.py", πTemp011, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µdir *πg.Object = πg.UnboundLocal; _ = µdir
				var µerr *πg.Object = πg.UnboundLocal; _ = µerr
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
					// line 98: dir, err = Getwd()
					πF.SetLineno(98)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßGetwd); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp003}}}, πTemp002); πE != nil {
						continue
					}
					µdir = πTemp001
					µerr = πTemp003
					if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, µerr); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label1
					}
					goto Label2
					// line 99: if err:
					πF.SetLineno(99)
				Label1:
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µerr, ßError, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp005[0] = πTemp002
					if πTemp001, πE = πg.ResolveGlobal(πF, ßOSError); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					// line 100: raise OSError(err.Error())
					πF.SetLineno(100)
					πE = πF.Raise(πTemp002, nil, nil)
					continue
					goto Label2
				Label2:
					// line 101: return dir
					πF.SetLineno(101)
					if πE = πg.CheckLocal(πF, µdir, "dir"); πE != nil {
						continue
					}
					πR = µdir
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßgetcwd.ToObject(), πTemp014); πE != nil {
				continue
			}
			// line 104: class _Popen(object):
			πF.SetLineno(104)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp017, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp017
			πTemp004 = πg.NewDict()
			if πTemp015, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ß__module__.ToObject(), πTemp015); πE != nil {
				continue
			}
			_, πE = πg.NewCode("_Popen", "/usr/lib/python2.7/os/__init__.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp004
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []πg.Param
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 106: def __init__(self, command, mode):
					πF.SetLineno(106)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "command", Def: nil}
					πTemp002[2] = πg.Param{Name: "mode", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/os/__init__.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µcommand *πg.Object = πArgs[1]; _ = µcommand
						var µmode *πg.Object = πArgs[2]; _ = µmode
						var µerr *πg.Object = πg.UnboundLocal; _ = µerr
						var µattr *πg.Object = πg.UnboundLocal; _ = µattr
						var µfiles_type *πg.Object = πg.UnboundLocal; _ = µfiles_type
						var µfiles *πg.Object = πg.UnboundLocal; _ = µfiles
						var µfd *πg.Object = πg.UnboundLocal; _ = µfd
						var µargs_type *πg.Object = πg.UnboundLocal; _ = µargs_type
						var µargs *πg.Object = πg.UnboundLocal; _ = µargs
						var µunix_shell *πg.Object = πg.UnboundLocal; _ = µunix_shell
						var µshell *πg.Object = πg.UnboundLocal; _ = µshell
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
						var πTemp007 *πg.Object
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
							// line 107: self.mode = mode
							πF.SetLineno(107)
							if πE = πg.CheckLocal(πF, µmode, "mode"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µmode); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßmode, πTemp001); πE != nil {
								continue
							}
							// line 108: self.result = None
							πF.SetLineno(108)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßresult, πTemp002); πE != nil {
								continue
							}
							// line 109: self.r, self.w, err = Pipe()
							πF.SetLineno(109)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßPipe); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp004}}}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßr, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßw, πTemp003); πE != nil {
								continue
							}
							µerr = πTemp004
							if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, µerr); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							goto Label2
							// line 110: if err:
							πF.SetLineno(110)
						Label1:
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µerr, ßError, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp006[0] = πTemp002
							if πTemp001, πE = πg.ResolveGlobal(πF, ßOSError); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							// line 111: raise OSError(err.Error())
							πF.SetLineno(111)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							goto Label2
						Label2:
							// line 112: attr = ProcAttr.new()
							πF.SetLineno(112)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßProcAttr); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßnew, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µattr = πTemp001
							// line 115: files_type = ToNative(__frame__(), attr.Files).Type()
							πF.SetLineno(115)
							πTemp006 = πF.MakeArgs(2)
							if πTemp001, πE = πg.ResolveGlobal(πF, ß__frame__); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp006[0] = πTemp002
							if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µattr, ßFiles, nil); πE != nil {
								continue
							}
							πTemp006[1] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßToNative); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßType, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µfiles_type = πTemp002
							// line 116: files = MakeSlice(files_type, 3, 3).Interface()
							πF.SetLineno(116)
							πTemp006 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µfiles_type, "files_type"); πE != nil {
								continue
							}
							πTemp006[0] = µfiles_type
							πTemp006[1] = πg.NewInt(3).ToObject()
							πTemp006[2] = πg.NewInt(3).ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßMakeSlice); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßInterface, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µfiles = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßmode, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, πTemp002, ßr.ToObject()); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label3
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßmode, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, πTemp002, ßw.ToObject()); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label4
							}
							goto Label5
							// line 117: if self.mode == 'r':
							πF.SetLineno(117)
						Label3:
							// line 118: fd = self.r.Fd()
							πF.SetLineno(118)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßr, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßFd, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µfd = πTemp001
							// line 119: files[0], files[1], files[2] = Stdin, self.w, Stderr
							πF.SetLineno(119)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßStdin); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßw, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßStderr); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple3(πTemp002, πTemp003, πTemp004).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp004}}}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µfiles, "files"); πE != nil {
								continue
							}
							πTemp007 = πg.NewInt(0).ToObject()
							if πE = πg.SetItem(πF, µfiles, πTemp007, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µfiles, "files"); πE != nil {
								continue
							}
							πTemp002 = πg.NewInt(1).ToObject()
							if πE = πg.SetItem(πF, µfiles, πTemp002, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µfiles, "files"); πE != nil {
								continue
							}
							πTemp002 = πg.NewInt(2).ToObject()
							if πE = πg.SetItem(πF, µfiles, πTemp002, πTemp004); πE != nil {
								continue
							}
							goto Label6
							// line 120: elif self.mode == 'w':
							πF.SetLineno(120)
						Label4:
							// line 121: fd = self.w.Fd()
							πF.SetLineno(121)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßw, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßFd, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µfd = πTemp001
							// line 122: files[0], files[1], files[2] = self.r, Stdout, Stderr
							πF.SetLineno(122)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßr, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßStdout); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßStderr); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple3(πTemp002, πTemp003, πTemp004).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp004}}}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µfiles, "files"); πE != nil {
								continue
							}
							πTemp007 = πg.NewInt(0).ToObject()
							if πE = πg.SetItem(πF, µfiles, πTemp007, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µfiles, "files"); πE != nil {
								continue
							}
							πTemp002 = πg.NewInt(1).ToObject()
							if πE = πg.SetItem(πF, µfiles, πTemp002, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µfiles, "files"); πE != nil {
								continue
							}
							πTemp002 = πg.NewInt(2).ToObject()
							if πE = πg.SetItem(πF, µfiles, πTemp002, πTemp004); πE != nil {
								continue
							}
							goto Label6
						Label5:
							πTemp006 = πF.MakeArgs(2)
							πTemp006[0] = πg.NewStr("invalid popen mode: %r").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßmode, nil); πE != nil {
								continue
							}
							πTemp006[1] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							// line 124: raise ValueError('invalid popen mode: %r', self.mode)
							πF.SetLineno(124)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							goto Label6
						Label6:
							// line 125: attr.Files = files
							πF.SetLineno(125)
							if πE = πg.CheckLocal(πF, µfiles, "files"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µfiles); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µattr, ßFiles, πTemp001); πE != nil {
								continue
							}
							// line 127: args_type = ToNative(__frame__(), StartProcess).Type().In(1)
							πF.SetLineno(127)
							πTemp006 = πF.MakeArgs(1)
							πTemp006[0] = πg.NewInt(1).ToObject()
							πTemp008 = πF.MakeArgs(2)
							if πTemp001, πE = πg.ResolveGlobal(πF, ß__frame__); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp008[0] = πTemp002
							if πTemp001, πE = πg.ResolveGlobal(πF, ßStartProcess); πE != nil {
								continue
							}
							πTemp008[1] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßToNative); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßType, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßIn, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							µargs_type = πTemp002
							// line 128: args = MakeSlice(args_type, 3, 3).Interface()
							πF.SetLineno(128)
							πTemp006 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µargs_type, "args_type"); πE != nil {
								continue
							}
							πTemp006[0] = µargs_type
							πTemp006[1] = πg.NewInt(3).ToObject()
							πTemp006[2] = πg.NewInt(3).ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßMakeSlice); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßInterface, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µargs = πTemp002
							πTemp006 = πF.MakeArgs(2)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							πTemp006[0] = πTemp001
							πTemp006[1] = ßgetandroidapilevel.ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label7
							}
							goto Label8
							// line 132: if hasattr(sys, 'getandroidapilevel'):
							πF.SetLineno(132)
						Label7:
							// line 133: unix_shell = '/system/bin/sh'
							πF.SetLineno(133)
							µunix_shell = πg.NewStr("/system/bin/sh").ToObject()
							goto Label9
						Label8:
							// line 135: unix_shell = '/bin/sh'
							πF.SetLineno(135)
							µunix_shell = πg.NewStr("/bin/sh").ToObject()
							goto Label9
						Label9:
							// line 136: shell = environ.get('SHELL', unix_shell)
							πF.SetLineno(136)
							πTemp006 = πF.MakeArgs(2)
							πTemp006[0] = ßSHELL.ToObject()
							if πE = πg.CheckLocal(πF, µunix_shell, "unix_shell"); πE != nil {
								continue
							}
							πTemp006[1] = µunix_shell
							if πTemp001, πE = πg.ResolveGlobal(πF, ßenviron); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßget, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							µshell = πTemp001
							// line 137: args[0] = shell
							πF.SetLineno(137)
							if πE = πg.CheckLocal(πF, µshell, "shell"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µshell); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							πTemp002 = πg.NewInt(0).ToObject()
							if πE = πg.SetItem(πF, µargs, πTemp002, πTemp001); πE != nil {
								continue
							}
							// line 138: args[1] = '-c'
							πF.SetLineno(138)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewStr("-c").ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							πTemp002 = πg.NewInt(1).ToObject()
							if πE = πg.SetItem(πF, µargs, πTemp002, πTemp001); πE != nil {
								continue
							}
							// line 139: args[2] = command
							πF.SetLineno(139)
							if πE = πg.CheckLocal(πF, µcommand, "command"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µcommand); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							πTemp002 = πg.NewInt(2).ToObject()
							if πE = πg.SetItem(πF, µargs, πTemp002, πTemp001); πE != nil {
								continue
							}
							// line 140: self.proc, err = StartProcess(shell, args, attr)
							πF.SetLineno(140)
							πTemp006 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µshell, "shell"); πE != nil {
								continue
							}
							πTemp006[0] = µshell
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							πTemp006[1] = µargs
							if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
								continue
							}
							πTemp006[2] = µattr
							if πTemp001, πE = πg.ResolveGlobal(πF, ßStartProcess); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp003}}}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßproc, πTemp001); πE != nil {
								continue
							}
							µerr = πTemp003
							if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, µerr); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label10
							}
							goto Label11
							// line 141: if err:
							πF.SetLineno(141)
						Label10:
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µerr, ßError, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp006[0] = πTemp002
							if πTemp001, πE = πg.ResolveGlobal(πF, ßOSError); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							// line 142: raise OSError(err.Error())
							πF.SetLineno(142)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							goto Label11
						Label11:
							// line 143: self.wg = WaitGroup.new()
							πF.SetLineno(143)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßWaitGroup); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßnew, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßwg, πTemp002); πE != nil {
								continue
							}
							// line 144: self.wg.Add(1)
							πF.SetLineno(144)
							πTemp006 = πF.MakeArgs(1)
							πTemp006[0] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßwg, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßAdd, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							// line 145: StartThread(self._thread_func)
							πF.SetLineno(145)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_thread_func, nil); πE != nil {
								continue
							}
							πTemp006[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßStartThread); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							// line 146: self.file = NewFileFromFD(fd, self.close)
							πF.SetLineno(146)
							πTemp006 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µfd, "fd"); πE != nil {
								continue
							}
							πTemp006[0] = µfd
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßclose, nil); πE != nil {
								continue
							}
							πTemp006[1] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNewFileFromFD); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßfile, πTemp001); πE != nil {
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
					// line 148: def _thread_func(self):
					πF.SetLineno(148)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("_thread_func", "/usr/lib/python2.7/os/__init__.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 149: self.result = self.proc.Wait()
							πF.SetLineno(149)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßproc, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßWait, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßresult, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßmode, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, πTemp002, ßr.ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 150: if self.mode == 'r':
							πF.SetLineno(150)
						Label1:
							// line 151: self.w.Close()
							πF.SetLineno(151)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßw, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßClose, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							goto Label2
						Label2:
							// line 152: self.wg.Done()
							πF.SetLineno(152)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßwg, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßDone, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ß_thread_func.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 154: def close(self, _):
					πF.SetLineno(154)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "_", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("close", "/usr/lib/python2.7/os/__init__.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µ_ *πg.Object = πArgs[1]; _ = µ_
						var µstate *πg.Object = πg.UnboundLocal; _ = µstate
						var µerr *πg.Object = πg.UnboundLocal; _ = µerr
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
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßmode, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, πTemp002, ßw.ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 155: if self.mode == 'w':
							πF.SetLineno(155)
						Label1:
							// line 156: self.w.Close()
							πF.SetLineno(156)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßw, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßClose, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							goto Label2
						Label2:
							// line 157: self.wg.Wait()
							πF.SetLineno(157)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßwg, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßWait, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 158: state, err = self.result
							πF.SetLineno(158)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßresult, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp004}}}, πTemp001); πE != nil {
								continue
							}
							µstate = πTemp002
							µerr = πTemp004
							if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, µerr); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label3
							}
							goto Label4
							// line 159: if err:
							πF.SetLineno(159)
						Label3:
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µerr, ßError, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp005[0] = πTemp002
							if πTemp001, πE = πg.ResolveGlobal(πF, ßOSError); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 160: raise OSError(err.Error())
							πF.SetLineno(160)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							goto Label4
						Label4:
							// line 161: return state.Sys()
							πF.SetLineno(161)
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µstate, ßSys, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßclose.ToObject(), πTemp004); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp016, πE = πTemp004.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp016 == nil {
				πTemp016 = πg.TypeType.ToObject()
			}
			if πTemp017, πE = πTemp016.Call(πF, []*πg.Object{πg.NewStr("_Popen").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_Popen.ToObject(), πTemp017); πE != nil {
				continue
			}
			// line 164: def popen(command, mode='r'):
			πF.SetLineno(164)
			πTemp011 = make([]πg.Param, 2)
			πTemp011[0] = πg.Param{Name: "command", Def: nil}
			πTemp011[1] = πg.Param{Name: "mode", Def: ßr.ToObject()}
			πTemp015 = πg.NewFunction(πg.NewCode("popen", "/usr/lib/python2.7/os/__init__.py", πTemp011, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µcommand *πg.Object = πArgs[0]; _ = µcommand
				var µmode *πg.Object = πArgs[1]; _ = µmode
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
					// line 165: return _Popen(command, mode).file
					πF.SetLineno(165)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µcommand, "command"); πE != nil {
						continue
					}
					πTemp001[0] = µcommand
					if πE = πg.CheckLocal(πF, µmode, "mode"); πE != nil {
						continue
					}
					πTemp001[1] = µmode
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_Popen); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßfile, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßpopen.ToObject(), πTemp015); πE != nil {
				continue
			}
			// line 168: def remove(filepath):
			πF.SetLineno(168)
			πTemp011 = make([]πg.Param, 1)
			πTemp011[0] = πg.Param{Name: "filepath", Def: nil}
			πTemp016 = πg.NewFunction(πg.NewCode("remove", "/usr/lib/python2.7/os/__init__.py", πTemp011, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µfilepath *πg.Object = πArgs[0]; _ = µfilepath
				var µerr *πg.Object = πg.UnboundLocal; _ = µerr
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
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					πTemp001 = πF.MakeArgs(1)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfilepath, "filepath"); πE != nil {
						continue
					}
					πTemp002[0] = µfilepath
					if πTemp003, πE = πg.ResolveGlobal(πF, ßstat); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßst_mode, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ßstat_module); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßS_ISDIR, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
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
					// line 169: if stat_module.S_ISDIR(stat(filepath).st_mode):
					πF.SetLineno(169)
				Label1:
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfilepath, "filepath"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Add(πF, πg.NewStr("Operation not permitted: ").ToObject(), µfilepath); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ßOSError); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 170: raise OSError('Operation not permitted: ' + filepath)
					πF.SetLineno(170)
					πE = πF.Raise(πTemp004, nil, nil)
					continue
					goto Label2
				Label2:
					// line 171: err = Remove(filepath)
					πF.SetLineno(171)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfilepath, "filepath"); πE != nil {
						continue
					}
					πTemp001[0] = µfilepath
					if πTemp003, πE = πg.ResolveGlobal(πF, ßRemove); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µerr = πTemp004
					if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, µerr); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label3
					}
					goto Label4
					// line 172: if err:
					πF.SetLineno(172)
				Label3:
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µerr, ßError, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp004
					if πTemp003, πE = πg.ResolveGlobal(πF, ßOSError); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 173: raise OSError(err.Error())
					πF.SetLineno(173)
					πE = πF.Raise(πTemp004, nil, nil)
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
			if πE = πF.Globals().SetItem(πF, ßremove.ToObject(), πTemp016); πE != nil {
				continue
			}
			// line 176: def rmdir(filepath):
			πF.SetLineno(176)
			πTemp011 = make([]πg.Param, 1)
			πTemp011[0] = πg.Param{Name: "filepath", Def: nil}
			πTemp017 = πg.NewFunction(πg.NewCode("rmdir", "/usr/lib/python2.7/os/__init__.py", πTemp011, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µfilepath *πg.Object = πArgs[0]; _ = µfilepath
				var µerr *πg.Object = πg.UnboundLocal; _ = µerr
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
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					πTemp002 = πF.MakeArgs(1)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfilepath, "filepath"); πE != nil {
						continue
					}
					πTemp003[0] = µfilepath
					if πTemp004, πE = πg.ResolveGlobal(πF, ßstat); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp004, πE = πg.GetAttr(πF, πTemp005, ßst_mode, nil); πE != nil {
						continue
					}
					πTemp002[0] = πTemp004
					if πTemp004, πE = πg.ResolveGlobal(πF, ßstat_module); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßS_ISDIR, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp005.Call(πF, πTemp002, nil); πE != nil {
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
					// line 177: if not stat_module.S_ISDIR(stat(filepath).st_mode):
					πF.SetLineno(177)
				Label1:
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfilepath, "filepath"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, πg.NewStr("Operation not permitted: ").ToObject(), µfilepath); πE != nil {
						continue
					}
					πTemp002[0] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßOSError); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					// line 178: raise OSError('Operation not permitted: ' + filepath)
					πF.SetLineno(178)
					πE = πF.Raise(πTemp004, nil, nil)
					continue
					goto Label2
				Label2:
					// line 179: err = Remove(filepath)
					πF.SetLineno(179)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfilepath, "filepath"); πE != nil {
						continue
					}
					πTemp002[0] = µfilepath
					if πTemp001, πE = πg.ResolveGlobal(πF, ßRemove); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µerr = πTemp004
					if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, µerr); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label3
					}
					goto Label4
					// line 180: if err:
					πF.SetLineno(180)
				Label3:
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µerr, ßError, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp002[0] = πTemp004
					if πTemp001, πE = πg.ResolveGlobal(πF, ßOSError); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					// line 181: raise OSError(err.Error())
					πF.SetLineno(181)
					πE = πF.Raise(πTemp004, nil, nil)
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
			if πE = πF.Globals().SetItem(πF, ßrmdir.ToObject(), πTemp017); πE != nil {
				continue
			}
			// line 184: class StatResult(object):
			πF.SetLineno(184)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp020, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp020
			πTemp004 = πg.NewDict()
			if πTemp018, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ß__module__.ToObject(), πTemp018); πE != nil {
				continue
			}
			_, πE = πg.NewCode("StatResult", "/usr/lib/python2.7/os/__init__.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp004
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 186: def __init__(self, info):
					πF.SetLineno(186)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "info", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/os/__init__.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µinfo *πg.Object = πArgs[1]; _ = µinfo
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 187: self._info = info
							πF.SetLineno(187)
							if πE = πg.CheckLocal(πF, µinfo, "info"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µinfo); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_info, πTemp001); πE != nil {
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
					// line 189: def st_mode(self):
					πF.SetLineno(189)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("st_mode", "/usr/lib/python2.7/os/__init__.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
							// line 191: return self._info.Mode()
							πF.SetLineno(191)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_info, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßMode, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßst_mode.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 193: st_mode = property(st_mode)
					πF.SetLineno(193)
					πTemp004 = πF.MakeArgs(1)
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßst_mode); πE != nil {
						continue
					}
					πTemp004[0] = πTemp005
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßproperty); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πE = πClass.SetItem(πF, ßst_mode.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 195: def st_mtime(self):
					πF.SetLineno(195)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("st_mtime", "/usr/lib/python2.7/os/__init__.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
							// line 196: return float(self._info.ModTime().UnixNano()) / Second
							πF.SetLineno(196)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_info, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßModTime, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßUnixNano, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßfloat); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßSecond); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Div(πF, πTemp004, πTemp003); πE != nil {
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
					if πE = πClass.SetItem(πF, ßst_mtime.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 198: st_mtime = property(st_mtime)
					πF.SetLineno(198)
					πTemp004 = πF.MakeArgs(1)
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßst_mtime); πE != nil {
						continue
					}
					πTemp004[0] = πTemp006
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßproperty); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πE = πClass.SetItem(πF, ßst_mtime.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 200: def st_size(self):
					πF.SetLineno(200)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("st_size", "/usr/lib/python2.7/os/__init__.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
							// line 201: return self._info.Size()
							πF.SetLineno(201)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_info, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßSize, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßst_size.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 203: st_size = property(st_size)
					πF.SetLineno(203)
					πTemp004 = πF.MakeArgs(1)
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßst_size); πE != nil {
						continue
					}
					πTemp004[0] = πTemp007
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßproperty); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp007.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πE = πClass.SetItem(πF, ßst_size.ToObject(), πTemp008); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp019, πE = πTemp004.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp019 == nil {
				πTemp019 = πg.TypeType.ToObject()
			}
			if πTemp020, πE = πTemp019.Call(πF, []*πg.Object{πg.NewStr("StatResult").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßStatResult.ToObject(), πTemp020); πE != nil {
				continue
			}
			// line 206: def stat(filepath):
			πF.SetLineno(206)
			πTemp011 = make([]πg.Param, 1)
			πTemp011[0] = πg.Param{Name: "filepath", Def: nil}
			πTemp018 = πg.NewFunction(πg.NewCode("stat", "/usr/lib/python2.7/os/__init__.py", πTemp011, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µfilepath *πg.Object = πArgs[0]; _ = µfilepath
				var µinfo *πg.Object = πg.UnboundLocal; _ = µinfo
				var µerr *πg.Object = πg.UnboundLocal; _ = µerr
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
					// line 207: info, err = Stat(filepath)
					πF.SetLineno(207)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfilepath, "filepath"); πE != nil {
						continue
					}
					πTemp001[0] = µfilepath
					if πTemp002, πE = πg.ResolveGlobal(πF, ßStat); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp004}}}, πTemp003); πE != nil {
						continue
					}
					µinfo = πTemp002
					µerr = πTemp004
					if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, µerr); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label1
					}
					goto Label2
					// line 208: if err:
					πF.SetLineno(208)
				Label1:
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µerr, ßError, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					if πTemp002, πE = πg.ResolveGlobal(πF, ßOSError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 209: raise OSError(err.Error())
					πF.SetLineno(209)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					goto Label2
				Label2:
					// line 210: return StatResult(info)
					πF.SetLineno(210)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µinfo, "info"); πE != nil {
						continue
					}
					πTemp001[0] = µinfo
					if πTemp002, πE = πg.ResolveGlobal(πF, ßStatResult); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßstat.ToObject(), πTemp018); πE != nil {
				continue
			}
			// line 213: unlink = remove
			πF.SetLineno(213)
			if πTemp019, πE = πg.ResolveGlobal(πF, ßremove); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßunlink.ToObject(), πTemp019); πE != nil {
				continue
			}
			// line 216: def waitpid(pid, options):
			πF.SetLineno(216)
			πTemp011 = make([]πg.Param, 2)
			πTemp011[0] = πg.Param{Name: "pid", Def: nil}
			πTemp011[1] = πg.Param{Name: "options", Def: nil}
			πTemp019 = πg.NewFunction(πg.NewCode("waitpid", "/usr/lib/python2.7/os/__init__.py", πTemp011, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µpid *πg.Object = πArgs[0]; _ = µpid
				var µoptions *πg.Object = πArgs[1]; _ = µoptions
				var µstatus *πg.Object = πg.UnboundLocal; _ = µstatus
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
					// line 217: status = WaitStatus.new()
					πF.SetLineno(217)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßWaitStatus); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßnew, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					µstatus = πTemp001
					// line 218: _syscall.invoke(Wait4, pid, status, options, None)
					πF.SetLineno(218)
					πTemp003 = πF.MakeArgs(5)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßWait4); πE != nil {
						continue
					}
					πTemp003[0] = πTemp001
					if πE = πg.CheckLocal(πF, µpid, "pid"); πE != nil {
						continue
					}
					πTemp003[1] = µpid
					if πE = πg.CheckLocal(πF, µstatus, "status"); πE != nil {
						continue
					}
					πTemp003[2] = µstatus
					if πE = πg.CheckLocal(πF, µoptions, "options"); πE != nil {
						continue
					}
					πTemp003[3] = µoptions
					if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp003[4] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_syscall); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßinvoke, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					// line 219: return pid, _encode_wait_result(status)
					πF.SetLineno(219)
					if πE = πg.CheckLocal(πF, µpid, "pid"); πE != nil {
						continue
					}
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µstatus, "status"); πE != nil {
						continue
					}
					πTemp003[0] = µstatus
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_encode_wait_result); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πTemp001 = πg.NewTuple2(µpid, πTemp004).ToObject()
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
			if πE = πF.Globals().SetItem(πF, ßwaitpid.ToObject(), πTemp019); πE != nil {
				continue
			}
			// line 222: def _encode_wait_result(status):
			πF.SetLineno(222)
			πTemp011 = make([]πg.Param, 1)
			πTemp011[0] = πg.Param{Name: "status", Def: nil}
			πTemp020 = πg.NewFunction(πg.NewCode("_encode_wait_result", "/usr/lib/python2.7/os/__init__.py", πTemp011, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µstatus *πg.Object = πArgs[0]; _ = µstatus
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
					// line 223: return status.Signal() | (status.ExitStatus() << 8)
					πF.SetLineno(223)
					if πE = πg.CheckLocal(πF, µstatus, "status"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µstatus, ßSignal, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µstatus, "status"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µstatus, ßExitStatus, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.LShift(πF, πTemp005, πg.NewInt(8).ToObject()); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Or(πF, πTemp003, πTemp002); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ß_encode_wait_result.ToObject(), πTemp020); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("os", Code)
}

