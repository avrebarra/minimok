<div class="info" align="left">
  <h1 class="name">🦹🏾‍♀️ minimok</h1>
  postman collections as local documentation server
  <br>
  <br>

[![Godoc][godoc-image]][godoc-url]
[![Report][report-image]][report-url]
[![Tests][tests-image]][tests-url]
[![Coverage][coverage-image]][coverage-url]

</div>


## Installation

Download latest binary release from release page.

You can use this install script to download the latest version:

```sh
# install latest release to /usr/local/bin/
curl https://i.jpillora.com/avrebarra/minimok! | *remove_this* bash
```

```sh
# install specific version
curl https://i.jpillora.com/avrebarra/minimok@{version} | *remove_this* bash
```

## Usage
### CLI 
```bash
# start minimok server. sample config file can be seen at samples/ruleconfig.yaml.
$ minimok start -conf ./ruleconfig.yaml -port 4434
```

```bash
# check for helps
$ minimok -help

minimok v0 - mini mock server

Available commands:

   start   start minimok 

Flags:

  -help
        Get help on the 'minimok' command.
```

[godoc-image]: https://godoc.org/github.com/avrebarra/minimok?status.svg
[godoc-url]: https://godoc.org/github.com/avrebarra/minimok
[report-image]: https://goreportcard.com/badge/github.com/avrebarra/minimok
[report-url]: https://goreportcard.com/report/github.com/avrebarra/minimok
[tests-image]: https://cloud.drone.io/api/badges/avrebarra/minimok/status.svg
[tests-url]: https://cloud.drone.io/avrebarra/minimok
[coverage-image]: https://codecov.io/gh/avrebarra/minimok/graph/badge.svg
[coverage-url]: https://codecov.io/gh/avrebarra/minimok
[sponsor-image]: https://img.shields.io/badge/github-donate-green.svg