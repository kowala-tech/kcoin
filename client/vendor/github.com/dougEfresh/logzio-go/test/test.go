// Copyright © 2017 Douglas Chimento <dchimento@gmail.com>
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
	"fmt"
	"os"
	"time"

	"github.com/dougEfresh/logzio-go"
)

func main() {
	l, err := logzio.New(os.Args[1], logzio.SetDebug(os.Stderr))
	if err != nil {
		panic(err)
	}

	msg := fmt.Sprintf("{ \"%s\": \"%d\"}", "message", time.Now().UnixNano())
	err = l.Send([]byte(msg))
	if err != nil {
		panic(err)
	}
	l.Stop() // logs are buffered on disk. Stop will drain the buffer
	time.Sleep(500 * time.Millisecond)
}
