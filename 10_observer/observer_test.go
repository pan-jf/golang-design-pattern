package observer

import "testing"

func TestExampleObserver(t *testing.T) {
	// 被观察者
	subject := NewSubject()
	// 观察者
	reader1 := NewReader("reader1")
	reader2 := NewReader("reader2")
	reader3 := NewReader("reader3")
	// 注册监听
	subject.Attach(reader1)
	subject.Attach(reader2)
	subject.Attach(reader3)
	// 被观察者发生改变
	subject.UpdateContext("observer mode")

	// Output:
	// reader1 receive observer mode
	// reader2 receive observer mode
	// reader3 receive observer mode
}
