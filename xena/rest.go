package xena

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
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

func newBaseRPC(config *rpcConf) baseRPC {
	rpc := baseRPC{
		config: config,
		http:   &http.Client{Timeout: config.timeout},
	}
	return rpc
}

type baseRPC struct {
	config *rpcConf
	http   *http.Client
}

func (r *baseRPC) get(query query) (*http.Response, []byte, error) {
	var body []byte
	u, err := query.SetHost(r.config.host)
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
	resp, err := r.http.Do(req)
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

func (r *baseRPC) post(query query, reqBody []byte) (*http.Response, []byte, error) {
	u, err := query.SetHost(r.config.host)
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

func (r *baseRPC) do(req *http.Request) error {
	resp, err := r.http.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode == http.StatusOK {

	}
	return nil
}

func (r *baseRPC) getHeaders(paramsHeaders map[string]string) map[string]string {
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

type rpcConf struct {
	host      string
	timeout   time.Duration
	userAgent string
	logger    Logger
	headers   map[string]string
}

type RestOption func(conf *rpcConf)

func WithRestHost(url string) RestOption {
	return func(conf *rpcConf) {
		conf.host = url
	}
}

func WithRestLogger(logger Logger) RestOption {
	if logger != nil {
		logger = newEmptyLogger()
	}
	return func(conf *rpcConf) {
		conf.logger = logger
	}
}

func WithRestMarketDataHost(conf *rpcConf) {
	conf.host = "https://api.xena.exchange"
}

func WithRestTradingHost(conf *rpcConf) {
	conf.host = "https://api.xena.exchange/trading"
}

func withRestDefaultTimeout(conf *rpcConf) {
	conf.timeout = time.Minute
}

func withRestDefaultLogger(conf *rpcConf) {
	conf.logger = newLogger(true)
}

func WithRestUserAgent(userAgent string) RestOption {
	return func(conf *rpcConf) {
		conf.userAgent = userAgent
	}
}

func withRestHeader(key, value string) RestOption {
	return func(conf *rpcConf) {
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

func NewQuery(path ...string) query {
	return query{
		path:    append([]string{}, path...),
		values:  url.Values{},
		headers: make(map[string]string),
	}
}

func (q query) AddPath(path ...string) query {
	q.path = append(q.path, path...)
	return q
}

func (q query) AddPathf(path ...interface{}) query {
	for _, v := range path {
		q.path = append(q.path, fmt.Sprintf("%s", v))
	}
	return q
}

func (q query) AddQuery(key, value string) query {
	q.values.Add(key, value)
	return q
}

func (q query) AddQueryf(key string, value interface{}) query {
	if value != nil && !reflect.ValueOf(value).IsNil() {
		switch t := value.(type) {
		case *int:
			q.AddQuery(key, fmt.Sprintf("%d", *t))
		case *int64:
			q.AddQuery(key, fmt.Sprintf("%d", *t))
		case *uint64:
			q.AddQuery(key, fmt.Sprintf("%d", *t))
		case *string:
			q.AddQuery(key, fmt.Sprintf("%s", *t))
		case *time.Time:
			q.AddQuery(key, fmt.Sprintf("%d", (*t).UnixNano()))
			//default:
			//	fmt.Printf("%s, %s\n", key, value)
		}
	}
	return q
}

func (q query) AddQueryInt(key string, value int64) query {
	q.values.Add(key, strconv.FormatInt(value, 10))
	return q
}

func (q query) SetHost(host string) (string, error) {
	u, err := url.Parse(host)
	if err != nil {
		return "", err
	}
	u.Path = path.Join(append([]string{u.Path}, q.path...)...)
	u.RawQuery = q.values.Encode()
	return u.String(), nil
}

func (q query) AddHeader(key, value string) query {
	q.headers[key] = value
	return q
}

func (q query) AddSecret(apiSecret string) (query, error) {
	nonce := time.Now().UnixNano()
	payload := fmt.Sprintf("AUTH%d", nonce)

	// Signature - SHA512 + ECDSA
	privKeyData, err := hex.DecodeString(apiSecret)
	if err != nil {
		return q, fmt.Errorf("error: %s on DecodeString", err)
	}

	privKey, err := x509.ParseECPrivateKey(privKeyData)
	if err != nil {
		return q, fmt.Errorf("error: %s on ParseECPrivateKey", err)
	}

	digest := sha256.Sum256([]byte(payload))
	r, s, err := ecdsa.Sign(rand.Reader, privKey, digest[:])
	if err != nil {
		return q, fmt.Errorf("%s on ecdsa.Sign()", err)
	}
	signature := append(r.Bytes(), s.Bytes()...)
	sigHex := hex.EncodeToString(signature)

	return q.AddHeader("X-Auth-Api-Payload", payload).AddHeader("X-Auth-Api-Signature", sigHex).AddHeader("X-Auth-Api-Nonce", strconv.FormatInt(nonce, 10)), nil
}

type xenaError struct {
	Error string `json:"error"`
}
