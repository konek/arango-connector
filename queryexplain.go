
package arango

type Expr struct{
  Type  string                  `json:"type"`
  Val   string                  `json:"value"`
  Extra map[string]interface{}  `json:"extra"`
}

type PlanElement struct{
  Id              int     `json:"id"`
  Type            string  `json:"type"`
  LoopLevel       int     `json:"loopLevel"`
  ResultVariable  string  `json:"resultVariable"`
  Limit           bool    `json:"limit"`
  Expr            Expr    `json:"expression"`
}

type QueryExplain struct{
  Plan  []PlanElement  `json:"plan"`
}

func (q *Query) Explain() ([]PlanElement, error) {
  var e StandardError
  ret := &QueryExplain{}

  req := q.NewRequest("POST", "explain")
  req.Data = *q
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
  if resp.StatusCode != 200 {
    return nil, UnknownError{resp.StatusCode}
  }
  err = req.Read(&ret)
  if err != nil {
    return nil, err
  }
  return ret.Plan, nil
}

