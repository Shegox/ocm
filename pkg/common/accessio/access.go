// Copyright 2022 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package accessio

import (
	"io"

	"github.com/opencontainers/go-digest"
)

//  DataAccess describes the access to sequence of bytes
type DataAccess interface {
	// Get returns the content of the blob as byte array
	Get() ([]byte, error)
	// Reader returns a reader to incrementally access the blob content
	Reader() (io.ReadCloser, error)
}

//  BlobAccess describes the access to a blob
type BlobAccess interface {
	DataAccess

	// MimeType returns the mime type of the blob
	MimeType() string
	// Digest returns the blob digest
	Digest() digest.Digest
	// Size returns the blob size
	Size() int64
}