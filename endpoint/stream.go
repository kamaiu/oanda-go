package endpoint

import (
	"context"
	"errors"
	"github.com/kamaiu/oanda-go/model"
	"github.com/valyala/bytebufferpool"
	"github.com/valyala/fasthttp"
	"io"
	"net/http"
	"sync"
	"time"
)

const maxLineSize = 1024 * 1024 * 5

var (
	ErrChunkTooBig = errors.New("chunk too big")
)

func (c *Connection) doStream(
	url *bytebufferpool.ByteBuffer,
	handler streamHandler,
) (*Stream, error) {
	u := url.String()
	req, err := http.NewRequest(http.MethodGet, u, nil)
	bytebufferpool.Put(url)
	if err != nil {
		return nil, err
	}

	req.Header.Set(fasthttp.HeaderUserAgent, c.agent)
	req.Header.Set(fasthttp.HeaderAuthorization, c.auth)
	req.Header.Set(fasthttp.HeaderContentType, "application/json")
	req.Header.Set("Accept-Datetime-Format", (string)(model.AcceptDatetimeFormat_RFC3339))
	//req.Header.Set(fasthttp.HeaderContentType, "application/octet-stream")

	resp, err := c.netClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, StatusCodeError{Code: resp.StatusCode}
	}

	ctx, cancel := context.WithCancel(context.Background())
	s := &Stream{
		started: time.Now(),
		req:     req,
		resp:    resp,
		rd:      resp.Body,
		handler: handler,
		ctx:     ctx,
		cancel:  cancel,
	}
	s.wg.Add(1)
	go s.run()
	return s, nil
}

type streamHandler interface {
	handle(message []byte) error

	onClose()
}

type Stream struct {
	started time.Time
	req     *http.Request
	resp    *http.Response
	rd      io.ReadCloser
	handler streamHandler
	ctx     context.Context
	cancel  context.CancelFunc
	closed  bool
	wg      sync.WaitGroup
	mu      sync.Mutex
}

func (s *Stream) Started() time.Time {
	return s.started
}

func (s *Stream) Wait() {
	s.wg.Wait()
}

func (s *Stream) Done() <-chan struct{} {
	return s.ctx.Done()
}

func (s *Stream) push(b []byte) {
	b = jsonObjectTrim(b)
	if len(b) == 0 {
		return
	}
	_ = s.handler.handle(b)
}

func (s *Stream) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.closed {
		return nil
	}
	s.cancel()
	s.closed = true
	defer s.handler.onClose()
	return s.rd.Close()
}

func (s *Stream) run() {
	defer func() {
		s.wg.Done()
		_ = s.Close()
	}()
	var (
		lines = make([][]byte, 0, 8)
		rd    = newStreamReader(s.rd)
		err   error
	)
	for {
		lines, err = rd.next(lines[:0])

		switch len(lines) {
		case 0:
		case 1:
			s.push(lines[0])
		default:
			for _, line := range lines {
				s.push(line)
			}
		}
		if err != nil {
			return
		}
	}
}

type streamReader struct {
	rd io.Reader
	b  []byte
	l  int
}

func newStreamReader(rd io.Reader) *streamReader {
	b := make([]byte, 1024)
	return &streamReader{
		rd: rd,
		b:  b,
		l:  0,
	}
}

func (lr *streamReader) next(frames [][]byte) ([][]byte, error) {
	for {
		// Resize buffer to fit line
		if lr.l > 0 && len(lr.b)-lr.l < 256 {
			newSize := len(lr.b) + 1024
			if newSize >= maxLineSize {
				return nil, ErrChunkTooBig
			}
			newB := make([]byte, newSize)
			copy(newB, lr.b)
			lr.b = newB
		}

		// Read more bytes
		n, err := lr.rd.Read(lr.b[lr.l:])

		count := 0
		// Process more?
		if n > 0 && n < len(lr.b[lr.l:]) {
			idx := lr.l
			sz := idx + n
			mark := 0
			for ; idx < sz; idx++ {
				if lr.b[idx] == '\n' {
					frames = append(frames, lr.b[mark:idx])
					mark = idx + 1
					count++
				}
			}

			if idx > mark {
				if lr.b[idx-1] == '}' {
					frames = append(frames, lr.b[mark:idx])
					lr.l = 0
					count++
				} else {
					copy(lr.b[0:], lr.b[mark:])
					lr.l = idx - mark
				}
			} else {
				lr.l = 0
			}
		}
		if err != nil {
			return frames, err
		}

		if count == 0 {
			continue
		}
		return frames, err
	}
}

func jsonObjectTrim(b []byte) []byte {
	if len(b) < 2 {
		return b
	}
	if b[0] != '{' {
	loop:
		for i := 0; i < len(b); i++ {
			switch b[i] {
			case ' ', '\t', '\r':
				b = b[i:]
			case '{':
				b = b[i:]
				break loop

			default:
				break loop
			}
		}
	}
	if b[len(b)-1] != '}' {
	loop2:
		for i := len(b) - 1; i > -1; i-- {
			switch b[i] {
			case ' ', '\t', '\r':
				b = b[0:i]
			case '}':
				b = b[0 : i+1]
				break loop2

			default:
				break loop2
			}
		}
	}
	return b
}
