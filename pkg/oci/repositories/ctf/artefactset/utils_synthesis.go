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

package artefactset

import (
	"io"

	"github.com/gardener/ocm/pkg/common/accessio"
	"github.com/gardener/ocm/pkg/common/accessobj"
	"github.com/gardener/ocm/pkg/errors"
	"github.com/gardener/ocm/pkg/oci/artdesc"
	"github.com/gardener/ocm/pkg/oci/cpi"
	"github.com/mandelsoft/vfs/pkg/osfs"
	"github.com/mandelsoft/vfs/pkg/vfs"
)

type ArtefactBlob interface {
	accessio.BlobAccess
	io.Closer
	FileSystem() vfs.FileSystem
	Path() string
}

type artefactBlob struct {
	accessio.BlobAccess
	temp       vfs.File
	filesystem vfs.FileSystem
}

func (a *artefactBlob) Close() error {
	if a.temp != nil {
		list := errors.ErrListf("synthesized blob")
		list.Add(a.temp.Close())
		list.Add(a.filesystem.Remove(a.temp.Name()))
		a.temp = nil
		return list.Result()
	}
	return nil
}

func (a *artefactBlob) FileSystem() vfs.FileSystem {
	return a.filesystem
}

func (a *artefactBlob) Path() string {
	return a.temp.Name()
}

// SynthesizeArtefactBlob synthesizes an artefact blob incorporation all side artefacts.
// To support extensions like cosign, we need the namespace access her to find
// additionally objects associated by tags.
func SynthesizeArtefactBlob(ns cpi.NamespaceAccess, ref string) (ArtefactBlob, error) {
	art, err := ns.GetArtefact(ref)
	if err != nil {
		return nil, err
	}

	blob, err := art.Blob()
	if err != nil {
		return nil, err
	}
	digest := blob.Digest()

	fs := osfs.New()
	temp, err := vfs.TempFile(fs, "", "artefactblob*.tgz")
	if err != nil {
		return nil, err
	}
	defer func() {
		// cleanup everything, if an error is returned (indicated by valid temp)
		if temp != nil {
			name := temp.Name()
			temp.Close()
			fs.Remove(name)
		}
	}()

	ab := &artefactBlob{
		BlobAccess: accessio.BlobAccessForFile(blob.MimeType()+"+tar+gzip", temp.Name(), fs),
		filesystem: fs,
		temp:       temp,
	}
	_ = art

	set, err := Create(accessobj.ACC_CREATE, "", 0600, accessobj.File(temp), accessobj.FormatTGZ)
	if err != nil {
		return nil, err
	}
	defer set.Close()
	err = TransferArtefact(art, ns, set)
	if err != nil {
		return nil, err
	}

	if !artdesc.IsDigest(ref) {
		err = set.AddTags(digest, ref)
		if err != nil {
			return nil, err
		}
	}
	set.Annotate(TAGS_MAINARTEFACT, digest.String())
	temp = nil
	return ab, nil
}

func TransferArtefact(art cpi.ArtefactAccess, ns cpi.NamespaceAccess, set cpi.ArtefactSink) error {
	if art.GetDescriptor().IsIndex() {
		return TransferIndexToSet(art.IndexAccess(), ns, set)
	} else {
		return TransferManifestToSet(art.ManifestAccess(), ns, set)
	}
}

func TransferIndexToSet(art cpi.IndexAccess, ns cpi.NamespaceAccess, set cpi.ArtefactSink) error {
	for _, l := range art.GetDescriptor().Manifests {
		art, err := art.GetArtefact(l.Digest)
		if err != nil {
			return err
		}
		err = TransferArtefact(art, ns, set)
		if err != nil {
			return err
		}
	}
	_, err := set.AddTaggedArtefact(art)
	return err
}

func TransferManifestToSet(art cpi.ManifestAccess, ns cpi.NamespaceAccess, set cpi.ArtefactSink) error {
	blob, err := art.GetConfigBlob()
	if err != nil {
		return err
	}
	err = set.AddBlob(blob)
	if err != nil {
		return err
	}
	for _, l := range art.GetDescriptor().Layers {
		blob, err = art.GetBlob(l.Digest)
		if err != nil {
			return err
		}
		err = set.AddBlob(blob)
		if err != nil {
			return err
		}
	}
	_, err = set.AddTaggedArtefact(art)
	return err
}