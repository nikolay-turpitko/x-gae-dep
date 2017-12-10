package def

import (
	"html/template"
	"net/http"

	"google.golang.org/appengine"
)

func init() {
	http.HandleFunc("/", root)
}

func root(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	if err := tmpl.Execute(w, struct {
		ServiceName            string
		AppID                  string
		Datacenter             string
		DefaultVersionHostname string
		InstanceID             string
		IsDevAppServer         bool
		ModuleName             string
		RequestID              string
		ServerSoftware         string
		VersionID              string
	}{
		ServiceName:            "Default",
		AppID:                  appengine.AppID(c),
		Datacenter:             appengine.Datacenter(c),
		DefaultVersionHostname: appengine.DefaultVersionHostname(c),
		InstanceID:             appengine.InstanceID(),
		IsDevAppServer:         appengine.IsDevAppServer(),
		ModuleName:             appengine.ModuleName(c),
		RequestID:              appengine.RequestID(c),
		ServerSoftware:         appengine.ServerSoftware(),
		VersionID:              appengine.VersionID(c),
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var tmpl = template.Must(template.New("book").Parse(`
<html>
  <head>
    <title>Test</title>
  </head>
  <body>
    <div>{{printf "%+v" .}}</div>
  </body>
</html>
`))
