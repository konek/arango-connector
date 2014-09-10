
package arango

type Options struct{
  CreateCollection  bool
  WaitForSync       bool
  KeepNull          bool
}

func OptionsDefaults() Options {
  var ret Options

  ret.Init()
  return ret
}

func (o *Options) Init() {
  o.CreateCollection = false
  o.WaitForSync = false
  o.KeepNull = false
}

