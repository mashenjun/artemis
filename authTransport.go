package artemis

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
)

type AuthTransport struct {
	Ak string
	Sk string
	HeaderMap map[string]struct{}
}


func (at AuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// todo: add auth info according to the artemis docs

	baseStr := req.Method + "\n"
	if req.Header.Get("Accept") != "" {
		baseStr = baseStr + req.Header.Get("Accept") + "\n"
	}else {
		req.Header.Set("Accept", "*/*")
		baseStr = baseStr + "*/*"+ "\n"
	}

	// cal MD5 hash
	if req.Body != nil {
		buf, err := ioutil.ReadAll(req.Body)
		if err != nil {
			return nil, err
		}
		rdr2 := ioutil.NopCloser(bytes.NewBuffer(buf))
		b := md5.Sum(buf)
		md5Hash := base64.StdEncoding.EncodeToString(b[:])
		req.Body = rdr2 // OK since rdr2 implements the io.ReadCloser interface

		baseStr = baseStr + md5Hash + "\n"
	}

	baseStr = baseStr + "application/json;charset=UTF-8" + "\n"
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	if req.Header.Get("Date") != "" {
		baseStr = baseStr + req.Header.Get("Date") + "\n"
	}
	tmp := map[string]string{}
	keys := make([]string, 0) // will be used in header
	for k, _ := range req.Header {
		if _, ok := at.HeaderMap[k]; !ok {
			tmp[strings.ToLower(k)] = strings.ToLower(k)+":"+strings.TrimSpace(req.Header.Get(k))
			keys = append(keys, strings.ToLower(k))
		}
	}
	sort.Strings(keys)
	for _, key := range keys {
		baseStr = baseStr + tmp[key] + "\n"
	}
	baseStr = baseStr + "x-ca-key:"+ at.Ak + "\n"
	req.Header.Set("x-ca-key", at.Ak)
	keys = append(keys, "x-ca-key")
	sort.Strings(keys)
	if req.URL.RawQuery != "" {
		baseStr	= baseStr + req.URL.Path + "?"+ req.URL.RawQuery
	}else {
		baseStr	= baseStr + req.URL.Path
	}
	h := hmac.New(sha256.New, []byte(at.Sk))
	h.Write([]byte(baseStr))
	signed := base64.StdEncoding.EncodeToString(h.Sum(nil))
	req.Header.Set("x-ca-signature", signed)
	req.Header.Set("x-ca-signature-headers", strings.Join(keys, ","))

	// still use default transport to send request
	return http.DefaultTransport.RoundTrip(req)
}

func NewAuthTransport(ak string, sk string) (*AuthTransport) {
	headerMap := map[string]struct{}{
		"X-Ca-Signature": struct{}{},
		"X-Ca-Signature-Headers":struct{}{},
		"Accept":struct{}{},
		"Content-MD5":struct{}{},
		"Content-Type":struct{}{},
		"Date":struct{}{},
		"Content-Length":struct{}{},
		"Server":struct{}{},
		"Connection":struct{}{},
		"Host":struct{}{},
		"Transfer-Encoding":struct{}{},
		"X-Application-Context":struct{}{},
		"Content-Encoding":struct{}{},
	}
	return &AuthTransport{Ak: ak, Sk: sk, HeaderMap:headerMap}
}
