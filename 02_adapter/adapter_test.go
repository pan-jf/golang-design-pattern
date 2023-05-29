package adapter

import "testing"

func TestAdapter(t *testing.T) {
	adapter := requestAdapter{}
	t.Log(adapter.Request("a")) // group a request impl
	t.Log(adapter.Request("b")) // group b request impl
}
