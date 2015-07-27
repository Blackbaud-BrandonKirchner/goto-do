package main

import (
  "github.com/go-martini/martini"
  "github.com/martini-contrib/render"
)

type NewRouter() *martini {
  m := martini.Classic()
  m.Use(render.Renderer())

  m.Get("/", func() string {
      return "Test go todo service"
    })

    m.Get("/todo", func(r render.Render) {
      r.HTML(200, "list", nil)
    })
  }

  m.Run()
}
