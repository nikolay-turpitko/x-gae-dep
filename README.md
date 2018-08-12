# x-gae-dep

This project illustrates usage of `https://github.com/golang/dep` tool (and
vendoring) in Google App Engine project.

It is created to illustrate approach described at
https://stackoverflow.com/a/40118834/2063744.

Tested in the environment as follows:

```console
$ lsb_release -a
No LSB modules are available.
Distributor ID:	Debian
Description:	Debian GNU/Linux 9.5 (stretch)
Release:	9.5
Codename:	stretch
$ go version
go version go1.10.3 linux/amd64
$ dep version
dep:
 version     : devel
 build date  : 
 git hash    : 
 go version  : go1.10.2
 go compiler : gc
 platform    : linux/amd64
 features    : ImportDuringSolve=false
$ goapp version
bash: goapp: command not found
$ gcloud version
Google Cloud SDK 211.0.0
alpha 2018.08.03
app-engine-go 
app-engine-python 1.9.73
beta 2018.08.03
bq 2.0.34
core 2018.08.03
gsutil 4.33
```

**NOTE:** `gcloud app deploy` is currently broken (in regard to vendoring), see
[the issue](https://issuetracker.google.com/issues/38449183).  Workaround,
suggested in the discussion of the issue, is a bit ugly, but the only option
worked for me:

- move `vendor` dir(s) from the `$GOPATH/src` directory (subdirectories in case
  of multiple modules like in this sample) into temporary dir;
- setup a `trap` to move them back after deployment;
- setup `$GOPATH` for `gcloud app deploy` so that it contain temporary dir(s)
  with vendored dependencies;
- execute `gcloud app deploy`, directories should be restored after script
  finished.

See `./deploy` and `./appengine-env` scripts. For Mac OS X you may need to
install `greadlink` with `brew install coreutils`.

# Links to the test app

- https://xgaedep.appspot.com/
- https://a-dot-xgaedep.appspot.com/
- https://b-dot-xgaedep.appspot.com/

# Useful links

- https://github.com/golang/dep
- https://stackoverflow.com/a/40118834/2063744
- https://cloud.google.com/appengine/docs/flexible/go/configuration-files
- https://cloud.google.com/appengine/docs/standard/python/designing-microservice-api
- https://issuetracker.google.com/issues/38449183
