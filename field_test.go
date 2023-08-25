/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

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
