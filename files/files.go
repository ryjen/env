package files

import (
	"os"
	"strings"

	"github.com/golobby/env"
	"github.com/ryjen/env/context"
)

var (
	// the base environment file
	identifier = ".env"
	// the local environment file extension
	local = "local"
)

// Set configures the name of an environment file and the local extension.
func Set(envFile string, localExt string) {
	SetEnvFileName(envFile)
	SetLocalExt(localExt)
}

// SetEnvFileName configures the name of the environment file
// A value must be a non-empty string.
func SetEnvFileName(value string) {
	if len(value) == 0 {
		return
	}
	identifier = value
}

// SetLocalExt configures the local version of the environment file by appending this value
// Setting an empty string effectively turns off local extension file parsing
// The default value is 'local'
func SetLocalExt(value string) {
	local = value
}

// Load loads all the possible environment files.
// It will NOT overwrite the OS and return loaded values
func Load() map[string]string {
	return load(false)
}

// LoadFromDir is the same as Load but will temporarily set the current working directory
func LoadFromDir(value string) (map[string]string, error) {
	return loadFromDir(value, Load)
}

// Overload loads all the possible environment files.
// It will overwrite the OS and return loaded values.
func Overload() map[string]string {
	return load(true)
}

// OverloadFromDir is the same as Overload but will temporarily set the current working directory
func OverloadFromDir(value string) (map[string]string, error) {
	return loadFromDir(value, Overload)
}

func loadFromDir(value string, loader func() map[string]string) (map[string]string, error) {
	wd, err := os.Getwd()

	if err != nil {
		return nil, err
	}

	if err = os.Chdir(value); err != nil {
		return nil, err
	}

	vals := loader()

	if err = os.Chdir(wd); err != nil {
		return vals, err
	}

	return vals, nil
}

func load(override bool) map[string]string {

	// given an orderd list of env files
	input := fileNames()

	// and an output
	output := make(map[string]string)

	var vs map[string]string
	var err error

	// for each env file...
	for _, f := range input {

		// load the file
		if override {
			vs, err = env.Overload(f)
		} else {
			vs, err = env.Load(f)
		}

		// if the file exists
		if err != nil {
			continue
		}

		// and override or set the values in output
		for k, v := range vs {
			output[k] = v
		}
	}

	return output
}

func fn(parts ...string) string {
	return strings.Join(parts, ".")
}

// gets an ordered list of possible env file names based on configuration
func fileNames() []string {

	// start with non-context env files
	vals := []string{
		identifier,
		fn(identifier, local),
	}

	ctx := context.Get()

	// add context env files
	vals = append(vals, fn(identifier, ctx))
	vals = append(vals, fn(identifier, ctx, local))

	return vals
}
