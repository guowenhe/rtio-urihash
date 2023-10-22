/*
*
* Copyright 2023 RTIO authors.
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*      http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
*
 */

package main

import (
	"flag"
	"fmt"
	"hash/crc32"
)

func main() {

	uri := flag.String("u", "", "uri string, example: /uri/example")
	x := flag.Bool("x", false, "display digest with hex")
	flag.Parse()

	if len(*uri) == 0 {

		fmt.Println("parameter uri empty, Usage of urihash:")
		flag.PrintDefaults()
		return
	}

	d := crc32.ChecksumIEEE([]byte(*uri))
	if *x {
		fmt.Printf("URI: %s, CRC: 0x%x\n", *uri, d)
	} else {
		fmt.Printf("URI: %s CRC: %d\n", *uri, d)
	}

}
