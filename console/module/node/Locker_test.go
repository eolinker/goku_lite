package node

import "testing"

func Test_InstanceLocker_Lock(t *testing.T) {

	l := newInstanceLocker()
	instance := "test"
	if !l.Lock(instance) {
		t.Error("should lock true")
		return
	}

	if l.Lock(instance) {
		t.Error("should lock false")
		return
	}

	l.UnLock(instance)
	if !l.Lock(instance) {
		t.Error("should lock true")
		return
	}

	if l.Lock(instance) {
		t.Error("should lock false")
		return
	}

}
