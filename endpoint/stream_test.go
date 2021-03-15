package endpoint

import (
	"bytes"
	"io"
	"testing"
)

func TestLineReader(t *testing.T) {
	lines := make([][]byte, 0, 8)
	frame := bytes.NewBufferString
	frames := newFrameReader(
		frame("{\"type\":\"HEARTBEAT\"}\n{\"type\":\"HEARTBEAT\"}\n{\"type\":\"HEARTBEAT\"}\n{\"type\":\"TX\"}\n"),
		frame("{\"type\":\"HEARTBEAT\"}"),
		// Test fragmented frame
		frame("{\"type\":\"TX\""),
		frame("}\n"),
	)

	rd := newStreamReader(frames)
	lines, err := rd.next(lines[:0])
	if err != nil {
		t.Fatal(err)
	}
	if len(lines) != 4 {
		t.Fatal("expected 4 lines")
	}

	lines, err = rd.next(lines[:0])
	if err != nil {
		t.Fatal(err)
	}
	if len(lines) != 1 {
		t.Fatal("expected 1 lines")
	}

	lines, err = rd.next(lines[:0])
	if err != nil {
		t.Fatal(err)
	}
	if len(lines) != 1 {
		t.Fatal("expected 1 lines")
	}
	if string(lines[0]) != "{\"type\":\"TX\"}" {
		t.Fatal("did not expect: " + string(lines[0]))
	}

	lines, err = rd.next(lines[:0])
	if err != io.EOF {
		t.Fatal("not EOF")
	}
}

func TestJsonObjectTrim(t *testing.T) {
	type test struct {
		b      []byte
		expect []byte
	}
	tests := []test{
		{[]byte("{}"), []byte("{}")},
		{[]byte(" {} "), []byte("{}")},
		{[]byte("{} "), []byte("{}")},
		{[]byte(" {"), []byte("{")},
		{[]byte(" } "), []byte(" }")},
		{[]byte(" { a "), []byte("{ a")},
		{[]byte("\t \t  \r"), []byte("")},
	}
	for _, item := range tests {
		if string(jsonObjectTrim(item.b)) != string(item.expect) {
			t.Fatal("expected: '" + string(item.expect) + "'   instead: '" + string(jsonObjectTrim(item.b)) + "'")
		}
	}
}

type frameReader struct {
	frames  []io.Reader
	current io.Reader
}

func newFrameReader(bufs ...io.Reader) *frameReader {
	return &frameReader{frames: bufs}
}

func (m *frameReader) Read(b []byte) (n int, err error) {
	if m.current == nil {
		if len(m.frames) == 0 {
			return 0, io.EOF
		}
		m.current = m.frames[0]
		m.frames = m.frames[1:]
	}

	for {
		n, err = m.current.Read(b)
		if n > 0 {
			if err != nil {
				if err == io.EOF {
					if len(m.frames) == 0 {
						return n, err
					}
					m.current = m.frames[0]
					m.frames = m.frames[1:]
				}
				return n, err
			}
			return n, nil
		}
		if err != nil {
			if err == io.EOF {
				if len(m.frames) == 0 {
					return n, err
				}
				m.current = m.frames[0]
				m.frames = m.frames[1:]
			}
		}
	}
}
