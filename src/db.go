func DB() martini.Handler {
  session, err := mgo.Dial("mongodb://localhost")
  if err != nil { panic(err) }
  
  return func(c martini.Context) {
    s := session.Clone()
    c.Map(s.DB("todo"))
    defer d.Close()
    c.Next()
  }
}
