// Code generated by 'yaegi extract syscall'. DO NOT EDIT.

//go:build go1.18 && !go1.19
// +build go1.18,!go1.19

package unrestricted

import (
	"reflect"
	"syscall"
)

func init() {
	Symbols["syscall/syscall"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"AllThreadsSyscall":  reflect.ValueOf(syscall.AllThreadsSyscall),
		"AllThreadsSyscall6": reflect.ValueOf(syscall.AllThreadsSyscall6),
		"Exec":               reflect.ValueOf(syscall.Exec),
		"Exit":               reflect.ValueOf(syscall.Exit),
		"ForkExec":           reflect.ValueOf(syscall.ForkExec),
		"Kill":               reflect.ValueOf(syscall.Kill),
		"PtraceAttach":       reflect.ValueOf(syscall.PtraceAttach),
		"PtraceCont":         reflect.ValueOf(syscall.PtraceCont),
		"PtraceDetach":       reflect.ValueOf(syscall.PtraceDetach),
		"PtraceGetEventMsg":  reflect.ValueOf(syscall.PtraceGetEventMsg),
		"PtraceGetRegs":      reflect.ValueOf(syscall.PtraceGetRegs),
		"PtracePeekData":     reflect.ValueOf(syscall.PtracePeekData),
		"PtracePeekText":     reflect.ValueOf(syscall.PtracePeekText),
		"PtracePokeData":     reflect.ValueOf(syscall.PtracePokeData),
		"PtracePokeText":     reflect.ValueOf(syscall.PtracePokeText),
		"PtraceSetOptions":   reflect.ValueOf(syscall.PtraceSetOptions),
		"PtraceSetRegs":      reflect.ValueOf(syscall.PtraceSetRegs),
		"PtraceSingleStep":   reflect.ValueOf(syscall.PtraceSingleStep),
		"PtraceSyscall":      reflect.ValueOf(syscall.PtraceSyscall),
		"RawSyscall":         reflect.ValueOf(syscall.RawSyscall),
		"RawSyscall6":        reflect.ValueOf(syscall.RawSyscall6),
		"Reboot":             reflect.ValueOf(syscall.Reboot),
		"Shutdown":           reflect.ValueOf(syscall.Shutdown),
		"StartProcess":       reflect.ValueOf(syscall.StartProcess),
		"Syscall":            reflect.ValueOf(syscall.Syscall),
		"Syscall6":           reflect.ValueOf(syscall.Syscall6),
		"Syscall9":           reflect.ValueOf(syscall.Syscall9),

		// type definitions
		"PtraceRegs": reflect.ValueOf((*syscall.PtraceRegs)(nil)),
	}
}
