# x-gae-dep

This project illustrates usage of `https://github.com/golang/dep` tool (and
vendoring) in Google App Engine project.

It is created to try approach described at
https://stackoverflow.com/a/40118834/2063744.

Tested in the environment as follows:

```bash
$ lsb_release -a
No LSB modules are available.
Distributor ID:	Ubuntu
Description:	Ubuntu 16.04.3 LTS
Release:	16.04
Codename:	xenial
$ go version
go version go1.9.2 linux/amd64
$ dep version
dep:
 version     : devel
 build date  : 
 git hash    : 
 go version  : go1.9.2
 go compiler : gc
 platform    : linux/amd64
$ goapp version
go version 1.8.3 (appengine-1.9.63) linux/amd64
$ gcloud version
Google Cloud SDK 182.0.0
app-engine-go 
app-engine-python 1.9.63
bq 2.0.27
core 2017.12.01
gsutil 4.28
```

To use `dep` and `goapp` you'll need to setup worksapce (GOPATH, GOROOT, PATH
environment variables for current project). Easy way to do it is via `source`
shell command. Fix paths, according to your local directory layout in the
`./appengine-env` file and use it to setup environment in a new shell.

```bash
source ./appengine-env
# dep init/ensure/status:
# - default service
cd src/def
dep ensure
# service A
cd ../a
dep ensure
# service B
cd ../b
dep ensure
# dep status:
dep status
dep status -dot | dot -T png | display
# Run app locally:
cd ../..
goapp version
# Use next line to build service (check it has no compilation errors).
goapp build gaedep/def gaedep/a gaedep/b
# Serve all modules
goapp serve ./appengine/def/app.yaml ./appengine/a/a.yaml ./appengine/b/b.yaml
# Deploy app. Note: gcloud tool is currently broken (attempts to deploy vendor dir), goapp works.
# gcloud app deploy --project xgaedep ./appengine/def/app.yaml ./appengine/a/a.yaml ./appengine/b/b.yaml
goapp deploy -application xgaedep -version v1 ./appengine/def/app.yaml ./appengine/a/a.yaml ./appengine/b/b.yaml
```

# Useful links

- https://github.com/golang/dep
- https://stackoverflow.com/a/40118834/2063744
- https://cloud.google.com/appengine/docs/flexible/go/configuration-files
- https://cloud.google.com/appengine/docs/standard/python/designing-microservice-api
