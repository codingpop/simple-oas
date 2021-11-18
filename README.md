# simple-oas

A dead-simple OpenApi docs server. By default, it uses [Swagger UI](https://github.com/swagger-api/swagger-ui) to present the docs, but it also suports [Redoc](https://github.com/Redocly/redoc).

## Compatibilty

This server supports OpenAPI specification version 2 or later.

## Install

The easiest way to install is using go install

```
go install github.com/codingpop/simple-oas@latest
```

## Usage

```
simple-oas
  -spec string
    	A file path or url to the json or yml OpenAPI spec (default "./spec.yml")
  -doc string
    Presentation UI (can be "redoc" or "swagger". Default "swagger")
  -host string
    	Host IP to serve using (default "0.0.0.0")
  -port int
    	Port to serve http over (default 9000)
```

## License & Contributing

* SimpleSwag is available as open source under the terms of the [MIT License](http://opensource.org/licenses/MIT).
* Bug reports and pull requests are welcome on GitHub at https://github.com/codingpop/simple-oas

## Inspiration
This project was inspired by [simple-swag](https://github.com/gocardless/simple-swag)
