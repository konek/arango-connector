
package arango

import (
  "fmt"
  "bytes"
  "net/http"
  "io/ioutil"
  "encoding/json"
  "encoding/base64"
)

type Request struct{
  *http.Request
  Arango

  Method    string
  URL       string
  Header    http.Header
  Data      interface{}
  System    bool

  Body      []byte
  Marshaled []byte
}

func (a Arango) NewRequest(method string, url string) *Request {
  return &Request{
    Arango: a,
    Method: method,
    URL: url,
  }
}

func (r *Request) setHeaders(h http.Header) {
  if h != nil {
    for k, vs := range h {
      for _, v := range vs {
        r.Request.Header.Set(k, v)
      }
    }
  }
}

func (r *Request) setCredentials() {
  if len(r.Arango.User) == 0 && len(r.Arango.Passwd) == 0 {
    return
  }
  basic := base64.URLEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", r.Arango.User, r.Arango.Passwd)))
  r.Request.Header.Set("Authorization", fmt.Sprintf("Basic %s", basic))
}

func (r *Request) Do() (*http.Response, error) {
  var buf *bytes.Buffer
  var _url string
  var err error

  if r.System == true {
    _url = fmt.Sprintf("%s/_api/%s", r.Arango.Addr, r.URL)
  } else if len(r.Arango.Db) == 0 {
     _url = fmt.Sprintf("%s/_api/%s", r.Arango.Addr, r.URL)
  } else {
    _url = fmt.Sprintf("%s/_db/%s/_api/%s", r.Arango.Addr, r.Arango.Db, r.URL)
  }
  if r.Data != nil {
    r.Marshaled, err = json.Marshal(r.Data)
    if err != nil {
      return nil, err
    }
    buf = bytes.NewBuffer(r.Marshaled)
  } else {
    buf = nil
  }
  if r.Data != nil {
    r.Request, err = http.NewRequest(r.Method, _url, buf)
  } else {
    r.Request, err = http.NewRequest(r.Method, _url, nil)
  }
  if err != nil {
    return nil, err
  }
  if r.Data != nil {
    r.Request.Header.Set("Content-Type", "application/json")
  }
  r.setCredentials()
  r.setHeaders(r.Header)
  resp, err := r.Arango.Client.Do(r.Request)
  if err != nil {
    return resp, err
  }
  r.Body, err = ioutil.ReadAll(resp.Body)
  if err != nil {
    return resp, err
  }
  if resp.StatusCode == 401 {
    return resp, StandardError{
      Err: true,
      Code: resp.StatusCode,
      ErrorNum: -1,
      ErrorMessage: "unauthorized",
    }
  }
  return resp, err
}

func (r Request) Read(obj interface{}) error {
  return json.Unmarshal(r.Body, obj)
}

