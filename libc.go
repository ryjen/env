package main

import (
	"C"
	"github.com/ryjen/env/context"
	"github.com/ryjen/env/files"
)

var version = "0.0.1-alpha"

//export env_context_version
func env_context_version() *C.char {
	return C.CString(version)
}

//export env_files_set
func env_files_set(envFile *C.char, localExt *C.char) {
	files.Set(C.GoString(envFile), C.GoString(localExt))
}

//export env_files_set_filename
func env_files_set_filename(val *C.char) {
	files.SetEnvFileName(C.GoString(val))
}

//export env_files_set_local_ext
func env_files_set_local_ext(val *C.char) {
	files.SetLocalExt(C.GoString(val))
}

//export env_files_load
func env_files_load() {
	files.Load()
}

//export env_files_overload
func env_files_overload() {
	files.Overload()
}

//export env_files_load_dir
func env_files_load_dir(val *C.char) {
	files.LoadFromDir(C.GoString(val))
}

//export env_files_overload_dir
func env_files_overload_dir(val *C.char) {
	files.OverloadFromDir(C.GoString(val))
}

//export env_context_get
func env_context_get() *C.char {
	return C.CString(context.Get())
}

//export env_context_set
func env_context_set(val *C.char) {
	context.Set(C.GoString(val))
}

//export env_context_is
func env_context_is(val *C.char) bool {
	return context.Is(C.GoString(val))
}

//export env_context_is_env
func env_context_is_env() bool {
	return context.IsEnv()
}

func main() {}
