package context

import (
	"os"
	"testing"
)

func TestGet(t *testing.T) {

	os.Setenv("ENV_CONTEXT", "test")

	ctx := Get()

	if ctx != "test" {
		t.Fail()
	}
}

func TestIs(t *testing.T) {

	os.Setenv("ENV_CONTEXT", "test")

	if !Is("test") {
		t.Fail()
	}
}

func TestIsEnv(t *testing.T) {

	os.Setenv("ENV_CONTEXT", "test")

	if !IsEnv() {
		t.Fail()
	}

	os.Unsetenv("ENV_CONTEXT")

	if IsEnv() {
		t.Fail()
	}
}

func TestSet(t *testing.T) {

	Set("dev")

	if IsEnv() {
		t.Fail()
	}

	if Is("test") {
		t.Fail()
	}

	ctx := Get()

	if ctx != "dev" {
		t.Fail()
	}
}
