
package arango

type User struct{
  Username  string  `json:"username"`
  Passwd    string  `json:"passwd"`
  Active    bool    `json:"active"`
  Extra     map[string]interface{}  `json:"extra"`
}

func DefaultUser() User {
  return User{
    Username: "",
    Passwd: "",
    Active: true,
    Extra: make(map[string]interface{}),
  }
}

