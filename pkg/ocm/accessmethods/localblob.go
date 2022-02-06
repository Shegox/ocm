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

package accessmethods

import (
	"encoding/json"

	"github.com/gardener/ocm/pkg/ocm/core"
	"github.com/gardener/ocm/pkg/ocm/cpi"
	"github.com/gardener/ocm/pkg/runtime"
)

// LocalBlobType is the access type of a blob local to a component.
const LocalBlobType = "localBlob"
const LocalBlobTypeV1 = LocalBlobType + runtime.VersionSeparator + "v1"

func init() {
	cpi.RegisterAccessType(cpi.NewConvertedAccessSpecType(LocalBlobType, LocalBlobV1))
	cpi.RegisterAccessType(cpi.NewConvertedAccessSpecType(LocalBlobTypeV1, LocalBlobV1))
}

// NewLocalBlobAccessSpecV1 creates a new localFilesystemBlob accessor.
func NewLocalBlobAccessSpecV1(path, name string, mediaType string) *LocalBlobAccessSpec {
	return &LocalBlobAccessSpec{
		ObjectVersionedType: runtime.NewVersionedObjectType(LocalBlobType),
		LocalReference:      path,
		ReferenceName:       name,
		MediaType:           mediaType,
	}
}

// LocalBlobAccessSpec describes the access for a blob on the filesystem.
type LocalBlobAccessSpec struct {
	runtime.ObjectVersionedType
	// LocalReference is the repository local identity of the blob.
	// it is used by the repository implementation to get access
	// to the blob and if therefore specific to a dedicated repository type.
	LocalReference string `json:"localReference"`
	// MediaType is the media type of the object represented by the blob
	MediaType string `json:"mediaType"`

	// ImageReference is an optional access information according the oci
	// distribution spec generated by a repository holding the local blob.
	// it can be used to offer a direct global access besides the access
	// through the component model api.
	ImageReference string `json:"imageReference,omitempty"`
	// ReferenceName is an optional static name the object should be
	// use in a local repository context. It is use by a repository
	// to optionally determine a globally referencable access according
	// to the OCI distribution spec. The result will be stored
	// by the repository in the field ImageReference.
	// The value is typically an OCI repository name optionally
	// followed by a colon ':' and a tag
	ReferenceName string `json:"referenceName,omitempty"`
}

var _ json.Marshaler = &LocalBlobAccessSpec{}

func (a LocalBlobAccessSpec) MarshalJSON() ([]byte, error) {
	return cpi.MarshalConvertedAccessSpec(core.DefaultContext, &a)
}

func (a *LocalBlobAccessSpec) IsLocal(cpi.Context) bool {
	return true
}

func (a *LocalBlobAccessSpec) AccessMethod(c cpi.ComponentVersionAccess) (cpi.AccessMethod, error) {
	return c.AccessMethod(a)
}

////////////////////////////////////////////////////////////////////////////////

type LocalBlobAccessSpecV1 struct {
	runtime.ObjectVersionedType `json:",inline"`
	// LocalReference is the repository local identity of the blob.
	// it is used by the repository implementation to get access
	// to the blob and if therefore specific to a dedicated repository type.
	LocalReference string `json:"localReference"`
	// MediaType is the media type of the object represented by the blob
	MediaType string `json:"mediaType"`

	// ImageReference is an optional access information according the oci
	// distribution spec generated by a repository holding the local blob.
	// it can be used to offer a direct global access besides the access
	// through the component model api.
	ImageReference string `json:"imageReference,omitempty"`
	// ReferenceName is an optional static name the object should be
	// use in a local repository context. It is use by a repository
	// to optionally determine a globally referencable access according
	// to the OCI distribution spec. The result will be stored
	// by the repository in the field ImageReference.
	// The value is typically an OCI repository name optionally
	// followed by a colon ':' and a tag
	ReferenceName string `json:"referenceName,omitempty"`
}

type localblobConverterV1 struct{}

var LocalBlobV1 = cpi.NewAccessSpecVersion(&LocalBlobAccessSpecV1{}, localblobConverterV1{})

func (_ localblobConverterV1) ConvertFrom(object cpi.AccessSpec) (runtime.TypedObject, error) {
	in := object.(*LocalBlobAccessSpec)
	return &LocalBlobAccessSpecV1{
		ObjectVersionedType: runtime.NewVersionedObjectType(in.Type),
		LocalReference:      in.LocalReference,
		ReferenceName:       in.ReferenceName,
		ImageReference:      in.ImageReference,
		MediaType:           in.MediaType,
	}, nil
}

func (_ localblobConverterV1) ConvertTo(object interface{}) (cpi.AccessSpec, error) {
	in := object.(*LocalBlobAccessSpecV1)
	return &LocalBlobAccessSpec{
		ObjectVersionedType: runtime.NewVersionedObjectType(in.Type),
		LocalReference:      in.LocalReference,
		ReferenceName:       in.ReferenceName,
		ImageReference:      in.ImageReference,
		MediaType:           in.MediaType,
	}, nil
}
