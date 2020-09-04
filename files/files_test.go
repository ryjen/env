package files

import (
	"github.com/ryjen/env/context"
	"testing"
)

func init() {
	SetEnvFileName("testenv")
}

func TestLoadNoContextNoLocal(t *testing.T) {

	context.Set("")
	SetLocalExt("")

	vals, err := LoadFromDir("../test")

	if err != nil {
		t.Fatal(err)
	}

	val, ok := vals["TEST_VAR"]

	if !ok || val != "1" {
		t.Fail()
	}
}

func TestLoadNoContext(t *testing.T) {
	context.Set("")
	SetLocalExt("loc")

	vals, err := LoadFromDir("../test")

	if err != nil {
		t.Fatal(err)
	}

	val, ok := vals["TEST_VAR"]

	if !ok || val != "2" {
		t.Fail()
	}
}

func TestLoadNoLocal(t *testing.T) {
	context.Set("dev")
	SetLocalExt("")

	vals, err := LoadFromDir("../test")

	if err != nil {
		t.Fatal(err)
	}

	val, ok := vals["TEST_VAR"]

	if !ok || val != "3" {
		t.Fail()
	}
}

func TestLoad(t *testing.T) {
	context.Set("dev")
	SetLocalExt("loc")

	vals, err := LoadFromDir("../test")

	if err != nil {
		t.Fatal(err)
	}

	val, ok := vals["TEST_VAR"]

	if !ok || val != "4" {
		t.Fail()
	}
}
