package learngo

import "testing"

func Test_SomeFn(t *testing.T) {
	want := "你好，世界。"

	if got := SomeFn(); got != want {
		t.Errorf("SomeFn() = %q, want %q", got, want)
	}
}
