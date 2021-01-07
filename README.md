<div class="info" align="left">
  <h1 class="name">🦹🏾‍♀️ minimok</h1>
  minimum mocking/proxy server
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
### Create a mock
To create mock, you can define a `configfile.yml` as below:
```yml
minimok:
- name: godoc_mock
  port: 1236
  rules:
    - accept: /
      use_origin: http://godoc.org/
      mock_latency: 
        mode: const   # will constantly/always set response latency
        value: 2000   # for each calls to 2000ms 

    - accept: /mocked/
      mock_response:  # send mocked response
        status: 200
        body: '{"rules":[{"accept":"/","origin":"http://localhost:9797","mock_response":"","mock_latency":{"mode":"swing","value":"2000","swing":"2000"}}]}'
        header: {}
      mock_latency:
        mode: max
        value: 10000
```

And apply it with this command:
```sh
$ ./dist/minimok start -conf ./config.yml
using configfile ./config.yaml
starting up mokserver 'godoc_mock' on http://localhost:1236
```

### Create a proxy port with modified latency
To create mock, you can define a `configfile.yml` as below:
```yml
minimok:
- name: askubuntu_proxy
  port: 1234
  rules:
    - accept: /
      use_origin: https://askubuntu.com/
      mock_latency: 
        mode: swing   # will set response latency 
        value: 2000   # for each calls between 1500ms and 2500ms 
        swing: 500
        hog: early    # (optional) hog request BEFORE delivering to remote target

- name: ubuntu_mock_server_with_regex
  port: 1237
  rules:
    - accept: /{rest:.*} # use gorilla mux syntax see https://github.com/gorilla/mux
      use_origin: https://askubuntu.com/
      mock_latency:
        mode: max     # will set response latency 
        value: 10000  # for each calls between 0ms and 2500ms
        hog: late     # (optional) hog request AFTER delivering to remote target
```

And apply it with this command:
```sh
$ ./dist/minimok start -conf ./config.yml
using configfile ./config.yaml
starting up mokserver 'askubuntu_proxy' on http://localhost:1234
starting up mokserver 'ubuntu_mock_server_with_regex' on http://localhost:1237
```


### Using the CLI 
Run help to check available commands:
```bash
# check for helps
$ minimok -help
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