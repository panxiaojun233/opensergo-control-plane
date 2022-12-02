// Copyright 2022, OpenSergo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	v3 "github.com/opensergo/opensergo-control-plane/pkg/proto/router/v3"
	"log"

	"github.com/opensergo/opensergo-control-plane"
)

func main() {
	cp, err := opensergo.NewControlPlane()
	if err != nil {
		log.Fatal(err)
	}
	r := v3.RouteConfiguration{Name: "123"}
	r.Name = "1233"
	err = cp.Start()
	if err != nil {
		log.Fatal(err)
	}
}
