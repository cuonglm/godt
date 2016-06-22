#godt - Go Docker Tag

godt -- Go library for getting docker image tags

[![Build Status](https://travis-ci.org/Gnouc/godt.svg?branch=master)](https://travis-ci.org/Gnouc/godt)
[![Go Report Card](https://goreportcard.com/badge/github.com/Gnouc/godt)](https://goreportcard.com/report/github.com/Gnouc/godt)

#Why godt?

At this time, there's no way to get image tags from docker client. We can use any HTTP client to get tags from docker hub:
```sh
curl https://registry.hub.docker.com//v1/repositories/fedora/tags | python -mjson.tool
```

It's too complicated from user perspective, so godt come in for simplification

#Installation
```sh
go get -u github.com/Gnouc/godt
```

#Usage
```sh
import "github.com/Gnouc/godt"
```

Example can be seen in `bin/docker-tags-v1`

```sh
$ go build docker-tags-v1.go
$ docker-tags-v1 -image alpine
latest
2.6
2.7
3.1
3.2
3.3
edge
```

#Environment variables

You can use `GODT_HUB_API_VERSION` and `GODT_HUB_URL` environment variable to change the docker hub version and url. Example with `docker-tags-v1`:

```sh
$ GODT_HUB_API_VERSION=2 ./docker-tags-v1
The requested URL (/v2//tags/list) was not found on this server.
```

By default, `GODT_HUB_API_VERSION` is `1` and `GODT_HUB_URL` is `https://registry.hub.docker.com`

#Note

Currently, authentication is not supported.

#Author

Cuong Manh Le <cuong.manhle.vn@gmail.com>

#License

See [LICENSE](https://github.com/Gnouc/godt/blob/master/LICENSE)
