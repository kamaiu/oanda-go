package endpoint

import (
	"encoding/json"
	"errors"
	. "github.com/kamaiu/oanda-go/model"
	"github.com/valyala/bytebufferpool"
	"github.com/valyala/fasthttp"
	"net"
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
	host           string
	hostStreaming  string
	port           int
	ssl            bool
	token          string
	auth           string
	DatetimeFormat string
	agent          string
	restClient     *fasthttp.HostClient
	streamClient   *http.Client
}

const DefaultUserAgent string = "oanda-go/0.9.0"

func NewConnection(token string, live bool) *Connection {
	host := ""
	fastHttpHost := ""
	hostStreaming := ""
	if live {
		fastHttpHost = "api-fxtrade.oanda.com:443"
		host = "https://api-fxtrade.oanda.com"
		hostStreaming = "https://stream-fxtrade.oanda.com"
	} else {
		fastHttpHost = "api-fxpractice.oanda.com:443"
		host = "https://api-fxpractice.oanda.com"
		hostStreaming = "https://stream-fxpractice.oanda.com"
	}

	// Create the Connection object
	connection := &Connection{
		host:          host,
		hostStreaming: hostStreaming,
		port:          443,
		ssl:           true,
		token:         token,
		auth:          "Bearer " + token,
		agent:         DefaultUserAgent,
		// HTTP client used for REST endpoints
		restClient: &fasthttp.HostClient{
			Addr:                          fastHttpHost,
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
		// HTTP client used for streaming endpoints
		// - Transactions
		// - Pricing
		streamClient: &http.Client{
			Transport: &http.Transport{
				DialContext: (&net.Dialer{
					Timeout:   30 * time.Second,
					KeepAlive: 30 * time.Second,
				}).DialContext,
				ForceAttemptHTTP2:     false,
				MaxIdleConns:          20,
				IdleConnTimeout:       90 * time.Second,
				TLSHandshakeTimeout:   10 * time.Second,
				ExpectContinueTimeout: 1 * time.Second,
			},
		},
	}
	return connection
}

type call struct {
	conn *Connection
	req  *fasthttp.Request
	resp *fasthttp.Response
}

func newCall(
	conn *Connection,
	method string,
	url *bytebufferpool.ByteBuffer,
	timeFormat AcceptDatetimeFormat,
) *call {
	ctx := &call{
		conn: conn,
		req:  fasthttp.AcquireRequest(),
		resp: fasthttp.AcquireResponse(),
	}
	ctx.req.Header.SetMethod(method)
	ctx.req.SetRequestURIBytes(url.B)
	// Release url buffer
	bytebufferpool.Put(url)
	if conn != nil {
		if len(conn.agent) > 0 {
			ctx.req.Header.Set(fasthttp.HeaderUserAgent, conn.agent)
		}
		ctx.req.Header.Set(fasthttp.HeaderAuthorization, conn.auth)
	}

	ctx.req.Header.Set(fasthttp.HeaderContentType, "application/json")
	ctx.req.Header.Set(fasthttp.HeaderAcceptEncoding, acceptEncoding)
	// Set time format
	if len(timeFormat) > 0 {
		ctx.req.Header.Set("Accept-Datetime-Format", (string)(timeFormat))
	}
	return ctx
}

func (c *call) complete(unmarshaller json.Unmarshaler) error {
	defer c.release()
	err := c.conn.restClient.DoRedirects(c.req, c.resp, maxRedirectsCount)
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

	err = conn.restClient.DoRedirects(ctx.req, ctx.resp, maxRedirectsCount)
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

func readBody(r *fasthttp.Response) (body []byte, err error) {
	if r == nil {
		return nil, nil
	}
	encoding := r.Header.Peek(fasthttp.HeaderContentEncoding)
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
