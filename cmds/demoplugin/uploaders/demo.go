// SPDX-FileCopyrightText: 2022 SAP SE or an SAP affiliate company and Open Component Model contributors.
//
// SPDX-License-Identifier: Apache-2.0

package uploaders

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/open-component-model/ocm/cmds/common"
	"github.com/open-component-model/ocm/cmds/demoplugin/accessmethods"
	"github.com/open-component-model/ocm/pkg/contexts/credentials"
	"github.com/open-component-model/ocm/pkg/contexts/oci/identity"
	"github.com/open-component-model/ocm/pkg/contexts/ocm/plugin/ppi"
	"github.com/open-component-model/ocm/pkg/runtime"
)

const (
	NAME    = "demo"
	VERSION = "v1"
)

type TargetSpec struct {
	runtime.ObjectVersionedType `json:",inline"`

	Path string `json:"path"`
}

var types map[string]runtime.TypedObjectDecoder

func init() {
	decoder, err := runtime.NewDirectDecoder(&TargetSpec{})
	if err != nil {
		panic(err)
	}
	types = map[string]runtime.TypedObjectDecoder{NAME + runtime.VersionSeparator + VERSION: decoder}
}

type Uploader struct {
	ppi.UploaderBase
}

var _ ppi.Uploader = (*Uploader)(nil)

func New() ppi.Uploader {
	return &Uploader{
		UploaderBase: ppi.MustNewUploaderBase("demo", "upload temp files"),
	}
}

func (a *Uploader) Decoders() map[string]runtime.TypedObjectDecoder {
	return types
}

func (a *Uploader) ValidateSpecification(p ppi.Plugin, spec ppi.UploadTargetSpec) (*ppi.UploadTargetSpecInfo, error) {
	var info ppi.UploadTargetSpecInfo
	my := spec.(*TargetSpec)

	if strings.HasPrefix(my.Path, "/") {
		return nil, fmt.Errorf("path must be relative (%s)", my.Path)
	}

	info.ConsumerId = credentials.ConsumerIdentity{
		identity.ID_TYPE:       common.CONSUMER_TYPE,
		identity.ID_HOSTNAME:   "localhost",
		identity.ID_PATHPREFIX: my.Path,
	}
	return &info, nil
}

func (a *Uploader) Writer(p ppi.Plugin, arttype, mediatype, hint string, repo ppi.UploadTargetSpec, creds credentials.Credentials) (io.WriteCloser, ppi.AccessSpecProvider, error) {
	var file *os.File
	var err error

	my := repo.(*TargetSpec)

	path := hint
	root := os.TempDir()
	dir := root
	if my.Path != "" {
		root = filepath.Join(root, my.Path)
		if hint == "" {
			path = my.Path
			dir = filepath.Join(dir, path)
		} else {
			path = filepath.Join(my.Path, hint)
			dir = filepath.Join(dir, filepath.Dir(path))
		}
	}

	err = os.MkdirAll(dir, 0o700)
	if err != nil {
		return nil, nil, err
	}

	if hint == "" {
		file, err = os.CreateTemp(root, "demo.*.blob")
	} else {
		file, err = os.OpenFile(filepath.Join(os.TempDir(), path), os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o600)
	}
	if err != nil {
		return nil, nil, err
	}
	writer := NewWriter(file, path, mediatype, hint == "", accessmethods.NAME, accessmethods.VERSION)
	return writer, writer.Specification, nil
}
