
package arango

type MatchQuery struct{
  Collection  string      `json:"collection"`
  Example     interface{} `json:"example"`
}

func (c Collection) Match(example interface{}) (*Cursor, error) {
  var e StandardError
  ret := c.Cursor()
  q := MatchQuery{
    Collection: c.Name,
    Example: example,
  }

  req := c.NewRequest("PUT", "simple/by-example")
  req.Data = q
  resp, err := req.Do()
  if err != nil {
    return nil, err
  }
  err = req.Read(&e)
  if err != nil {
    return nil, err
  }
  if e.Err == true {
    return nil, e
  }
  if resp.StatusCode != 201 {
    return nil, UnknownError{resp.StatusCode}
  }
  err = req.Read(ret)
  if err != nil {
    return nil, err
  }
  if len(ret.Result_) > 0 {
    ret.HasMore = true
  }
  return ret, nil
}

