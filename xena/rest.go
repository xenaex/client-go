package xena

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"reflect"
	"strconv"
	"time"
)

const userAgent = "xena/go"

func newBaseREST(config *restConf) baseREST {
	rest := baseREST{
		config: config,
		http:   &http.Client{Timeout: config.timeout},
	}
	return rest
}

type baseREST struct {
	config *restConf
	http   *http.Client
}

func (r *baseREST) get(query query) (*http.Response, []byte, error) {
	var body []byte
	u, err := query.setHost(r.config.host)
	if err != nil {
		return nil, nil, err
	}
	r.config.logger.Debugf("call %s", u)
	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}
	for k, v := range r.getHeaders(query.headers) {
		req.Header.Add(k, v)
	}
	//r.config.logger.Debugf("%s", req.Header)
	st := time.Now()
	resp, err := r.http.Do(req)
	if time.Now().Sub(st) > time.Second {
		r.config.logger.Debugf("long request %s to %s", time.Now().Sub(st), u)
	}
	if err == nil && (resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusBadRequest || resp.StatusCode == http.StatusUnauthorized) {
		body, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return resp, nil, err
		}
		//r.config.logger.Debugf("body: %s", body)
		if resp.StatusCode == http.StatusBadRequest {
			xenaError := xenaError{}
			err := json.Unmarshal(body, &xenaError)
			if err != nil {
				return resp, nil, err
			}
			if len(xenaError.Error) > 0 {
				return resp, nil, fmt.Errorf("%s", xenaError.Error)
			}
			return resp, nil, fmt.Errorf("%s", body)
		}
	}
	return resp, body, err
}

func (r *baseREST) post(query query, reqBody []byte) (*http.Response, []byte, error) {
	u, err := query.setHost(r.config.host)
	if err != nil {
		return nil, nil, err
	}
	r.config.logger.Debugf("call %s", u)
	req, err := http.NewRequest(http.MethodPost, u, bytes.NewReader(reqBody))
	if err != nil {
		return nil, nil, err
	}
	for k, v := range r.getHeaders(query.headers) {
		req.Header.Add(k, v)
	}
	resp, err := r.http.Do(req)
	var respBody []byte
	if err == nil && (resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusBadRequest) {
		respBody, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return resp, nil, err
		}
		if resp.StatusCode == http.StatusBadRequest {
			xenaError := xenaError{}
			err := json.Unmarshal(respBody, &xenaError)
			if err != nil {
				return resp, nil, err
			}
			if len(xenaError.Error) > 0 {
				return resp, nil, fmt.Errorf("%s", xenaError.Error)
			}
			return resp, nil, fmt.Errorf("%s", respBody)
		}
	}
	return resp, respBody, err
}

func (r *baseREST) do(req *http.Request) error {
	resp, err := r.http.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode == http.StatusOK {

	}
	return nil
}

func (r *baseREST) getHeaders(paramsHeaders map[string]string) map[string]string {
	headers := make(map[string]string, 2)
	headers["Accept"] = "application/json"
	headers["User-Agent"] = r.config.userAgent
	for k, v := range r.config.headers {
		headers[k] = v
	}
	for k, v := range paramsHeaders {
		headers[k] = v
	}
	return headers
}

type restConf struct {
	host      string
	timeout   time.Duration
	userAgent string
	logger    Logger
	headers   map[string]string
}

type RestOption func(conf *restConf)

//WithRestHost sets the url for connection.
func WithRestHost(url string) RestOption {
	return func(conf *restConf) {
		conf.host = url
	}
}

//WithRestLogger sets custom logger.
func WithRestLogger(logger Logger) RestOption {
	if logger != nil {
		logger = newEmptyLogger()
	}
	return func(conf *restConf) {
		conf.logger = logger
	}
}

//WithRestMarketDataHost sets the url for connection to market data api.
func WithRestMarketDataHost(conf *restConf) {
	conf.host = "https://api.xena.exchange"
}

//WithRestMarketDataHost sets the url for connection to trading api.
func WithRestTradingHost(conf *restConf) {
	conf.host = "https://api.xena.exchange/trading"
}

func withRestDefaultTimeout(conf *restConf) {
	conf.timeout = time.Minute
}

func withRestDefaultLogger(conf *restConf) {
	conf.logger = newLogger(true)
}

//WithRestUserAgent sets user agent.
func WithRestUserAgent(userAgent string) RestOption {
	return func(conf *restConf) {
		conf.userAgent = userAgent
	}
}

func withRestHeader(key, value string) RestOption {
	return func(conf *restConf) {
		if conf.headers == nil {
			conf.headers = make(map[string]string)
		}
		conf.headers[key] = value
	}
}

type query struct {
	path    []string
	values  url.Values
	headers map[string]string
}

func newQuery(path ...string) query {
	return query{
		path:    append([]string{}, path...),
		values:  url.Values{},
		headers: make(map[string]string),
	}
}

func (q query) addPath(path ...string) query {
	q.path = append(q.path, path...)
	return q
}

func (q query) addPathf(path ...interface{}) query {
	for _, v := range path {
		q.path = append(q.path, fmt.Sprintf("%s", v))
	}
	return q
}

func (q query) addQuery(key, value string) query {
	q.values.Add(key, value)
	return q
}

func (q query) addQueryf(key string, value interface{}) query {
	if value != nil && !reflect.ValueOf(value).IsNil() {
		switch t := value.(type) {
		case *int:
			q.addQuery(key, fmt.Sprintf("%d", *t))
		case *int64:
			q.addQuery(key, fmt.Sprintf("%d", *t))
		case *uint64:
			q.addQuery(key, fmt.Sprintf("%d", *t))
		case *string:
			q.addQuery(key, fmt.Sprintf("%s", *t))
		case *time.Time:
			q.addQuery(key, fmt.Sprintf("%d", (*t).UnixNano()))
			//default:
			//	fmt.Printf("%s, %s\n", key, value)
		}
	}
	return q
}

func (q query) addQueryInt(key string, value int64) query {
	q.values.Add(key, strconv.FormatInt(value, 10))
	return q
}

func (q query) setHost(host string) (string, error) {
	u, err := url.Parse(host)
	if err != nil {
		return "", err
	}
	u.Path = path.Join(append([]string{u.Path}, q.path...)...)
	u.RawQuery = q.values.Encode()
	return u.String(), nil
}

func (q query) addHeader(key, value string) query {
	q.headers[key] = value
	return q
}

func (q query) addSecret(apiSecret string) (query, error) {
	nonce, payload, sigHex, err := sing(apiSecret)
	if err != nil {
		return q, fmt.Errorf("%s on query.sing()", err)
	}
	return q.addHeader("X-Auth-Api-Payload", payload).addHeader("X-Auth-Api-Signature", sigHex).addHeader("X-Auth-Api-Nonce", strconv.FormatInt(nonce, 10)), nil
}

type xenaError struct {
	Error string `json:"error"`
}
