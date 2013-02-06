package logplex

import (
	"bytes"
	"io"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	data := `66 <174>1 2012-07-22T00:06:26-00:00 somehost Go console 2 Hi from Go
67 <174>1 2012-07-22T00:06:26-00:00 somehost Go console 10 Hi from Go
67 <174>1 2012-07-22T00:06:26-00:00 somehost Go console 10 Hi from Go
`

	exp := []*Msg{
		{
			174,
			[]byte("2012-07-22T00:06:26-00:00"),
			[]byte("somehost"),
			[]byte("Go"),
			[]byte("console"),
			[]byte("2"),
			[]byte("Hi from Go\n"),
		},
		{
			174,
			[]byte("2012-07-22T00:06:26-00:00"),
			[]byte("somehost"),
			[]byte("Go"),
			[]byte("console"),
			[]byte("10"),
			[]byte("Hi from Go\n"),
		},
		{
			174,
			[]byte("2012-07-22T00:06:26-00:00"),
			[]byte("somehost"),
			[]byte("Go"),
			[]byte("console"),
			[]byte("10"),
			[]byte("Hi from Go\n"),
		},
	}

	b := bytes.NewBufferString(data)
	r := NewReader(b)

	for i, e := range exp {
		t.Logf("EXP %d", i)

		m, err := r.ReadMsg()
		if err != nil {
			t.Errorf("error on %d: %v", i, err)
			continue
		}
		if m.Priority != e.Priority {
			t.Errorf("expected %d, got %d", e.Priority, m.Priority)
		}
		if !reflect.DeepEqual(m.Timestamp, e.Timestamp) {
			t.Errorf("expected %d, got %d", e.Timestamp, m.Timestamp)
		}
		if !reflect.DeepEqual(m.Host, e.Host) {
			t.Errorf("expected %s, got %s", e.Host, m.Host)
		}
		if !reflect.DeepEqual(m.Pid, e.Pid) {
			t.Errorf("expected %s, got %s", e.Pid, m.Pid)
		}
		if !reflect.DeepEqual(m.Id, e.Id) {
			t.Errorf("expected %s, got %s", e.Id, m.Id)
		}
		if !reflect.DeepEqual(m.Msg, e.Msg) {
			t.Errorf("expected %q, got %q", e.Msg, m.Msg)
		}
	}

	_, err := r.ReadMsg()
	if err != io.EOF {
		t.Errorf("expected io.EOF, got %v", err)
	}

	// one more for good measure
	_, err = r.ReadMsg()
	if err != io.EOF {
		t.Errorf("expected io.EOF, got %v", err)
	}
}
