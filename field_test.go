package jsonfield

import (
	"testing"
)

var str = `{"apiVersion":"apps/v1","kind":"Deployment","metadata":{"name":"nginx-deployment","labels":{"app":"nginx"}},"spec":{"replicas":3,"selector":{"matchLabels":{"app":"nginx"}},"template":{"metadata":{"labels":{"app":"nginx"}},"spec":{"containers":[{"name":"nginx","image":"nginx:1.14.2","ports":[{"containerPort":80}]}]}}}}`
var reservePath = []string{"kind", "metadata.name", "spec.replicas", "spec.template.spec.containers.image"}

func BenchmarkReserveField(b *testing.B) {
	bs := []byte(str)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = ReserveField(bs, reservePath)
	}
	b.StopTimer()
}
