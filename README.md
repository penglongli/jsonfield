# jsonfield

If our `JSON` string is very large, but we want to keep only a part of the fields during transmission to shorten the network overhead of transmission.

This little plugin is used to achieve this goal

## Purpose

There is a json string as example:

```json
{"apiVersion":"apps/v1","kind":"Deployment","metadata":{"name":"nginx-deployment","labels":{"app":"nginx"}},"spec":{"replicas":3,"selector":{"matchLabels":{"app":"nginx"}},"template":{"metadata":{"labels":{"app":"nginx"}},"spec":{"containers":[{"name":"nginx","image":"nginx:1.14.2","ports":[{"containerPort":80}]}]}}}}
```

It's so long, and we only want some target field:

```json 
{
	"kind": "Deployment",
	"metadata": {
		"name": "nginx-deployment"
	},
	"spec": {
		"replicas": 3,
		"template": {
			"spec": {
				"containers": [{
					"image": "nginx:1.14.2"
				}]
			}
		}
	}
}
```

## Usage

```go
package main

import (
	"fmt"

	"github.com/penglongli/jsonfield"
)

func main() {
	var str = `{"apiVersion":"apps/v1","kind":"Deployment","metadata":{"name":"nginx-deployment","labels":{"app":"nginx"}},"spec":{"replicas":3,"selector":{"matchLabels":{"app":"nginx"}},"template":{"metadata":{"labels":{"app":"nginx"}},"spec":{"containers":[{"name":"nginx","image":"nginx:1.14.2","ports":[{"containerPort":80}]}]}}}}`
	reservePath := []string{"kind", "metadata.name", "spec.replicas", "spec.template.spec.containers.image"}
	bs, err := jsonfield.ReserveField([]byte(str), reservePath)
	if err != nil {
		panic(err)
	}
	
	fmt.Println(string(bs))
}
```

## Benchmark

```bash
goos: darwin
goarch: arm64
pkg: github.com/penglongli/jsonfield
BenchmarkReserveField
BenchmarkReserveField-10    	  149000	      8418 ns/op
PASS
```
