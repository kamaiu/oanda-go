package endpoint

import (
	"bytes"
	"encoding/json"
	"errors"
	. "github.com/kamaiu/oanda-go/model"
	"github.com/valyala/bytebufferpool"
	"github.com/valyala/fasthttp"
	"net/http"
	"strconv"
	"time"
	"unsafe"
)

const (
	maxRedirectsCount = 16
	acceptEncoding    = "gzip, deflate, br"
)

var (
	ErrNilRequest = errors.New("nil request")
)

type StatusCodeError struct {
	Code int
}

func (s StatusCodeError) Error() string {
	return "status code " + strconv.Itoa(s.Code)
}

type Connection struct {
	hostname       string
	streamingHost  string
	hostnameBytes  []byte
	port           int
	ssl            bool
	token          string
	auth           string
	DatetimeFormat string
	agent          string
	client         *fasthttp.HostClient
	netClient      *http.Client
}

const DefaultUserAgent string = "oanda-go/0.9.0"

func NewConnection(token string, live bool) *Connection {
	hostname := ""
	streamingHost := ""
	// should we use the live API?
	if live {
		hostname = "https://api-fxtrade.oanda.com"
		streamingHost = "https://stream-fxtrade.oanda.com"
	} else {
		hostname = "https://api-fxpractice.oanda.com"
		streamingHost = "https://stream-fxpractice.oanda.com"
	}

	var auth bytes.Buffer
	// Generate the auth header
	auth.WriteString("Bearer ")
	auth.WriteString(token)

	// Create the Connection object
	connection := &Connection{
		hostname:      hostname,
		hostnameBytes: []byte(hostname),
		streamingHost: streamingHost,
		port:          443,
		ssl:           true,
		token:         token,
		auth:          auth.String(),
		agent:         DefaultUserAgent,
		client: &fasthttp.HostClient{
			Addr:                          hostname,
			Name:                          "oanda",
			NoDefaultUserAgentHeader:      true,
			Dial:                          fasthttp.Dial,
			DialDualStack:                 false,
			IsTLS:                         true,
			MaxConns:                      120,
			MaxConnDuration:               0,
			MaxIdleConnDuration:           time.Minute * 5,
			MaxIdemponentCallAttempts:     fasthttp.DefaultMaxIdemponentCallAttempts,
			ReadBufferSize:                1024 * 64,
			WriteBufferSize:               1024 * 64,
			ReadTimeout:                   time.Second * 15,
			WriteTimeout:                  time.Second * 15,
			MaxResponseBodySize:           fasthttp.DefaultMaxRequestBodySize,
			DisableHeaderNamesNormalizing: false,
			DisablePathNormalizing:        false,
			MaxConnWaitTimeout:            time.Second * 5,
			RetryIf:                       nil,
		},
		netClient: &http.Client{},
	}
	return connection
}

type call struct {
	req  *fasthttp.Request
	resp *fasthttp.Response
}

func newCall(
	c *Connection,
	method string,
	url *bytebufferpool.ByteBuffer,
	timeFormat AcceptDatetimeFormat,
) *call {
	ctx := &call{
		req:  fasthttp.AcquireRequest(),
		resp: fasthttp.AcquireResponse(),
	}
	ctx.req.Header.SetMethod(method)
	ctx.req.SetRequestURIBytes(url.B)
	bytebufferpool.Put(url)
	if c != nil {
		if len(c.agent) > 0 {
			ctx.req.Header.Set(fasthttp.HeaderUserAgent, c.agent)
		}
		ctx.req.Header.Set(fasthttp.HeaderAuthorization, c.auth)
		ctx.req.Header.Set(fasthttp.HeaderContentType, "application/json")
	}
	ctx.req.Header.Set(fasthttp.HeaderAcceptEncoding, acceptEncoding)
	// Set time format
	if len(timeFormat) > 0 {
		ctx.req.Header.Set("Accept-Datetime-Format", (string)(timeFormat))
	}
	return ctx
}

func (c *call) complete(unmarshaller json.Unmarshaler) error {
	defer c.release()
	err := fasthttp.DoRedirects(c.req, c.resp, maxRedirectsCount)
	if err != nil {
		return err
	}
	statusCode := c.resp.StatusCode()
	if statusCode != fasthttp.StatusOK {
		return StatusCodeError{Code: statusCode}
	}
	var body []byte
	body, err = readBody(c.resp)
	err = unmarshaller.UnmarshalJSON(body)
	return err
}

func (c *call) release() {
	fasthttp.ReleaseRequest(c.req)
	fasthttp.ReleaseResponse(c.resp)
}

func doGET(
	conn *Connection,
	url *bytebufferpool.ByteBuffer, // owned
	timeFormat AcceptDatetimeFormat,
	resp json.Unmarshaler,
) (statusCode int, err error) {
	ctx := newCall(conn, fasthttp.MethodGet, url, timeFormat)
	defer ctx.release()

	err = fasthttp.DoRedirects(ctx.req, ctx.resp, maxRedirectsCount)
	if err != nil {
		return 0, err
	}
	statusCode = ctx.resp.StatusCode()
	if statusCode != fasthttp.StatusOK {
		return statusCode, StatusCodeError{Code: statusCode}
	}

	var body []byte
	body, err = readBody(ctx.resp)
	err = resp.UnmarshalJSON(body)
	return statusCode, err
}

func readBody(r *fasthttp.Response) ([]byte, error) {
	if r == nil {
		return nil, nil
	}
	encoding := r.Header.Peek(fasthttp.HeaderContentEncoding)
	var body []byte
	var err error
	if len(encoding) > 0 {
		switch *(*string)(unsafe.Pointer(&encoding)) {
		case "deflate":
			body, err = r.BodyInflate()
		case "brotli":
			body, err = r.BodyUnbrotli()
		case "gzip":
			body, err = r.BodyGunzip()
		default:
			body = r.Body()
		}
	} else {
		body = r.Body()
	}
	return body, err
}
