# {{ .title.name }}

{{ .title.description }}

[![GoDoc][doc-img]][doc] [![Build Status][ci-img]][ci] [![Coverage Status][cov-img]][cov] [![Go Report][report-img]][report]

## Installation
{{ .installation }}

## Quick Start

{{ .quickStart.code}}

{{ .quickStart.description }}

## Usage

{{- range .usages }}

{{.}}

{{- end }}

## Disk queue

Logzio go client uses [https://github.com/syndtr/goleveldb] as a persistent storage.
Every 5 seconds logs are sent to logz.io (if any are available)

## Examples

{{- range .examples }}

{{.}}

{{- end }}


## Prerequisites

go 1.x

## Tests

{{- range .tests }}

{{.}}

{{- end }}

## Deployment

## Contributing
 All PRs are welcome

## Authors

* **Douglas Chimento**  - [{{.user}}][me]

## License

This project is licensed under the Apache License - see the [LICENSE](LICENSE) file for details

## Acknowledgments

* [logzio-java-sender](https://github.com/logzio/logzio-java-sender)

### TODO

[doc-img]: https://godoc.org/github.com/{{.user}}/{{.project}}?status.svg
[doc]: https://godoc.org/github.com/{{.user}}/{{.project}}
[ci-img]: https://travis-ci.org/{{.user}}/{{.project}}.svg?branch=master
[ci]: https://travis-ci.org/{{.user}}/{{.project}}
[cov-img]: https://codecov.io/gh/{{.user}}/{{.project}}/branch/master/graph/badge.svg
[cov]: https://codecov.io/gh/{{.user}}/{{.project}}
[glide.lock]: https://github.com/uber-go/zap/blob/master/glide.lock
[zap]: https://github.com/uber-go/zap
[me]: https://github.com/{{.user}}
[report-img]: https://goreportcard.com/badge/github.com/{{.user}}/{{.project}}
[report]: https://goreportcard.com/report/github.com/{{.user}}/{{.project}}
