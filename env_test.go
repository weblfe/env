package env

import "testing"

func TestEnvMethod_Get(t *testing.T) {
	var (
		env   = GetEnv()
		key   = "test.name"
		value = "{.Name}"
	)
	env.Put(key, value)
	env.Put("name", "env")

	if value != env.Get(key) {
		t.Error("env value not equal!")
	}
}

func TestEnvMethod_GetIntOf(t *testing.T) {
	var (
		env   = GetEnv()
		key   = "test.int"
		value =1
	)
	env.Put(key, value)
	if value != env.GetIntOf(key,-1) {
		t.Error("env value not equal!")
	}
}
