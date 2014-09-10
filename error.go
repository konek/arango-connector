
package arango

import (
  "fmt"
)

const (
  E_OK = 200
  E_Created = 201
  E_Accepted = 202
  E_NoContent = 204
  E_Partial = 206

  E_RedirMultiple = 300
  E_RedirPermanent = 301
  E_RedirFound = 302
  E_RedirSeeOther = 303
  E_RedirNotModified = 304
  E_RedirUseProxy = 305
  E_RedirTemporary = 307

  E_BadRequest = 400
  E_Unauthorized = 401
  E_Forbiden = 403
  E_NotFound = 404
  E_NotAllowed = 404
  E_NotAcceptable = 406
  E_ProxyAuth = 407
  E_Timeout = 408
  E_Conflict = 409
  E_Duplicate = 409

  E_Internal = 500
  E_NotImplemented = 501
  E_BadGateway = 502
  E_Unavailable = 503
  E_GatewayTimeout = 504
)

type StandardError struct{
  Err           bool    `json:"error" binding:"required"`
  Code          int     `json:"code"`
  ErrorNum      int     `json:"errorNum"`
  ErrorMessage  string  `json:"errorMessage"`
}

func (e StandardError) Error() string {
  return e.ErrorMessage
}

type UnknownError struct{
  Status  int
}

func (e UnknownError) Error() string {
  return fmt.Sprintf("an unknown error occured : %d", e.Status)
}

