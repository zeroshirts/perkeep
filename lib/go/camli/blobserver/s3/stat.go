/*
Copyright 2011 Google Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package s3

import (
	"log"
	"os"

	"camli/blobref"
)

var _ = log.Printf

func (sto *s3Storage) StatBlobs(dest chan<- blobref.SizedBlobRef, blobs []*blobref.BlobRef, waitSeconds int) os.Error {
	// TODO: do n stats in parallel
	for _, br := range blobs {
		size, err := sto.s3Client.Stat(br.String(), sto.bucket)
		log.Printf("stat of %s: %d, %v", br.String(), size, err)
		if err == nil {
			dest <- blobref.SizedBlobRef{BlobRef: br, Size: size}
		} else {
			// TODO: handle
		}
	}
	return nil
}
