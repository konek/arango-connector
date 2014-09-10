
package arango

type CreateDatabaseQuery struct{
  Name  string  `json:"name" binding:"required"`
  Users []User  `json:"users,omitempty"`
}

func (d Database) CreateAuto() error {
  list := make([]User, 1)
  user := d.Arango.User
  passwd := d.Arango.Passwd

  if len(user) == 0 {
    user = "root"
  }

  list[0] = DefaultUser()
  list[0].Username = user
  list[0].Passwd = passwd
  return d.Create(list)
}

func (d Database) Create(users []User) error {
  var e StandardError
  q := CreateDatabaseQuery{
    Name: d.Arango.Db,
    Users: users,
  }

  req := d.NewRequest("POST", "database")
  req.Data = q
  req.System = true
  resp, err := req.Do()
  if err != nil {
    return err
  }
  err = req.Read(&e)
  if err != nil {
    return err
  }
  if e.Err == true {
    return e
  }
  if resp.StatusCode != 201 {
    return UnknownError{resp.StatusCode}
  }
  return nil
}

