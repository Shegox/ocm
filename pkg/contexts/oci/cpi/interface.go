// SPDX-FileCopyrightText: 2022 SAP SE or an SAP affiliate company and Open Component Model contributors.
//
// SPDX-License-Identifier: Apache-2.0

package cpi

// This is the Context Provider Interface for credential providers

import (
	"github.com/opencontainers/go-digest"
	ociv1 "github.com/opencontainers/image-spec/specs-go/v1"

	"github.com/open-component-model/ocm/pkg/common/accessio"
	"github.com/open-component-model/ocm/pkg/contexts/datacontext"
	"github.com/open-component-model/ocm/pkg/contexts/oci/internal"
)

const CONTEXT_TYPE = internal.CONTEXT_TYPE

const CommonTransportFormat = internal.CommonTransportFormat

type (
	Context                          = internal.Context
	ContextProvider                  = internal.ContextProvider
	Repository                       = internal.Repository
	RepositorySpecHandlers           = internal.RepositorySpecHandlers
	RepositorySpecHandler            = internal.RepositorySpecHandler
	UniformRepositorySpec            = internal.UniformRepositorySpec
	RepositoryType                   = internal.RepositoryType
	RepositorySpec                   = internal.RepositorySpec
	IntermediateRepositorySpecAspect = internal.IntermediateRepositorySpecAspect
	GenericRepositorySpec            = internal.GenericRepositorySpec
	ArtefactAccess                   = internal.ArtefactAccess
	Artefact                         = internal.Artefact
	ArtefactSource                   = internal.ArtefactSource
	ArtefactSink                     = internal.ArtefactSink
	BlobSource                       = internal.BlobSource
	BlobSink                         = internal.BlobSink
	NamespaceLister                  = internal.NamespaceLister
	NamespaceAccess                  = internal.NamespaceAccess
	ManifestAccess                   = internal.ManifestAccess
	IndexAccess                      = internal.IndexAccess
	BlobAccess                       = internal.BlobAccess
	DataAccess                       = internal.DataAccess
	RepositorySource                 = internal.RepositorySource
)

type Descriptor = ociv1.Descriptor

var DefaultContext = internal.DefaultContext

func New(m ...datacontext.BuilderMode) Context {
	return internal.Builder{}.New(m...)
}

func RegisterRepositoryType(name string, atype RepositoryType) {
	internal.DefaultRepositoryTypeScheme.Register(name, atype)
}

func RegisterRepositorySpecHandler(handler RepositorySpecHandler, types ...string) {
	internal.RegisterRepositorySpecHandler(handler, types...)
}

func ToGenericRepositorySpec(spec RepositorySpec) (*GenericRepositorySpec, error) {
	return internal.ToGenericRepositorySpec(spec)
}

func UniformRepositorySpecForHostURL(typ string, host string) *UniformRepositorySpec {
	return internal.UniformRepositorySpecForHostURL(typ, host)
}

const (
	KIND_OCIARTEFACT = internal.KIND_OCIARTEFACT
	KIND_MEDIATYPE   = accessio.KIND_MEDIATYPE
	KIND_BLOB        = accessio.KIND_BLOB
)

func ErrUnknownArtefact(name, version string) error {
	return internal.ErrUnknownArtefact(name, version)
}

func ErrBlobNotFound(digest digest.Digest) error {
	return accessio.ErrBlobNotFound(digest)
}

func IsErrBlobNotFound(err error) bool {
	return accessio.IsErrBlobNotFound(err)
}
