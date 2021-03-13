package endpoint

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	. "github.com/kamaiu/oanda-go/model"
	"github.com/valyala/bytebufferpool"
	"github.com/valyala/fasthttp"
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

type Headers struct {
	contentType    string
	agent          string
	DatetimeFormat string
	auth           string
}

var client = &fasthttp.Client{
	Name:                          "",
	NoDefaultUserAgentHeader:      true,
	Dial:                          fasthttp.Dial,
	DialDualStack:                 false,
	TLSConfig:                     nil,
	MaxConnsPerHost:               100,
	MaxIdleConnDuration:           fasthttp.DefaultMaxIdleConnDuration,
	MaxConnDuration:               0,
	MaxIdemponentCallAttempts:     fasthttp.DefaultMaxIdemponentCallAttempts,
	ReadBufferSize:                0,
	WriteBufferSize:               0,
	ReadTimeout:                   time.Second * 15,
	WriteTimeout:                  time.Second * 15,
	MaxResponseBodySize:           1024 * 1024 * 5,
	DisableHeaderNamesNormalizing: false,
	DisablePathNormalizing:        false,
	MaxConnWaitTimeout:            time.Second * 5,
	RetryIf:                       nil,
}

type urls struct {
	accounts string
}

type Connection struct {
	hostname       string
	hostnameBytes  []byte
	port           int
	ssl            bool
	token          string
	DatetimeFormat string
	headers        *Headers
	client         *fasthttp.Client
	urls           urls
}

const OANDA_AGENT string = "v20-go/0.9.0"

func NewConnection(token string, live bool) *Connection {
	hostname := ""
	// should we use the live API?
	if live {
		hostname = "https://api-fxtrade.oanda.com/v3"
	} else {
		hostname = "https://api-fxpractice.oanda.com/v3"
	}

	var buffer bytes.Buffer
	// Generate the auth header
	buffer.WriteString("Bearer ")
	buffer.WriteString(token)

	authHeader := buffer.String()
	// Create headers for oanda to be used in requests
	headers := &Headers{
		contentType:    "application/json",
		agent:          OANDA_AGENT,
		DatetimeFormat: "RFC3339",
		auth:           authHeader,
	}
	// Create the Connection object
	connection := &Connection{
		hostname:      hostname,
		hostnameBytes: []byte(hostname),
		port:          443,
		ssl:           true,
		token:         token,
		headers:       headers,
		client:        client,
		urls: urls{
			accounts: fmt.Sprintf("%s/accounts", hostname),
		},
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
		if len(c.headers.agent) > 0 {
			ctx.req.Header.Set(fasthttp.HeaderUserAgent, c.headers.agent)
		}
		ctx.req.Header.Set(fasthttp.HeaderAuthorization, c.headers.auth)
		ctx.req.Header.Set(fasthttp.HeaderContentType, c.headers.contentType)
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

func (c *call) WithDateTime(format AcceptDatetimeFormat) *call {
	c.req.Header.Set("Accept-Datetime-Format", (string)(format))
	return c
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
