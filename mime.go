// +build linux darwin freebsd

package magic

// Copyright 2013 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// #cgo CFLAGS: -I/usr/local/include
// #cgo LDFLAGS: -lmagic -L/usr/local/lib
// #include <stdlib.h>
// #include <magic.h>
import "C"
import (
	"unsafe"
)

// DetectContentType implements the algorithm described
// at https://mimesniff.spec.whatwg.org/ to determine the
// Content-Type of the given data. It considers at most the
// first 512 bytes of data. DetectContentType always returns
// a valid MIME type: if it cannot determine a more specific one, it
// returns "application/octet-stream".
func DetectContentType(data []byte) string {
	db := C.magic_open(C.int(0))

	if db != nil {
		// free the library
		defer C.magic_close(db)

		// proceed with detection
		if code := C.magic_setflags(db, C.int(C.MAGIC_MIME_TYPE)); code == 0 {
			if code := C.magic_load(db, nil); code == 0 {
				var (
					ptr  = unsafe.Pointer(&data[0])
					size = C.size_t(len(data))
				)

				if out := C.magic_buffer(db, ptr, size); out != nil {
					return C.GoString(out)
				}
			}
		}
	}

	return "application/octet-stream"
}
