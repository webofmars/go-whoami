package main

import (
	"log"
	"net/http"
	"os"
	"text/template"
)

func main() {
  // preparing the template
  type RuntimeEnv struct {
    Name string
  }

  hostname, _ := os.Hostname()
  Myself := RuntimeEnv{hostname}

  t, err := template.ParseFiles("./templates/index.html.tmpl")
  if err != nil {
    log.Println(err)
  }
  f, _ := os.Create("html/index.html")
  err = t.ExecuteTemplate(f, "T1", Myself)
  if err != nil {
    log.Println(err)
  }

  // go serve the files generated
  fs := http.FileServer(http.Dir("html"))
  http.Handle("/", noCacheMiddleware(fs))

  log.Println("Listening...")
  http.ListenAndServe(":3000", nil)
}

// noCacheMiddleware adds anti-caching headers so browsers
// (especially Chromium-based ones) never serve stale content.
func noCacheMiddleware(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
    w.Header().Set("Pragma", "no-cache")
    w.Header().Set("Expires", "0")
    next.ServeHTTP(w, r)
  })
}
