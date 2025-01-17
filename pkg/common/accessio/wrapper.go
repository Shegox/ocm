// SPDX-FileCopyrightText: 2022 SAP SE or an SAP affiliate company and Open Component Model contributors.
//
// SPDX-License-Identifier: Apache-2.0

package accessio

import (
	"io"

	"github.com/opencontainers/go-digest"
)

type Writer interface {
	io.Writer
	io.WriterAt
}

type DataWriter interface {
	WriteTo(Writer) (int64, digest.Digest, error)
}

////////////////////////////////////////////////////////////////////////////////

type dataAccessWriter struct {
	access DataAccess
}

func NewDataAccessWriter(acc DataAccess) DataWriter {
	return &dataAccessWriter{acc}
}

func (d *dataAccessWriter) WriteTo(w Writer) (int64, digest.Digest, error) {
	r, err := d.access.Reader()
	if err != nil {
		return BLOB_UNKNOWN_SIZE, BLOB_UNKNOWN_DIGEST, err
	}
	defer r.Close()
	dr := NewDefaultDigestReader(r)
	if err != nil {
		return BLOB_UNKNOWN_SIZE, BLOB_UNKNOWN_DIGEST, err
	}
	_, err = io.Copy(w, dr)
	if err != nil {
		return BLOB_UNKNOWN_SIZE, BLOB_UNKNOWN_DIGEST, err
	}
	return dr.Size(), dr.Digest(), err
}

type writerAtWrapper struct {
	writer func(w io.WriterAt) error
}

func NewWriteAtWriter(at func(w io.WriterAt) error) DataWriter {
	return &writerAtWrapper{at}
}

func (d *writerAtWrapper) WriteTo(w Writer) (int64, digest.Digest, error) {
	return BLOB_UNKNOWN_SIZE, BLOB_UNKNOWN_DIGEST, d.writer(w)
}
