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

package file

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"

	"github.com/open-component-model/ocm/cmds/ocm/clictx"
	"github.com/open-component-model/ocm/cmds/ocm/commands/ocmcmds/common/inputs"
	"github.com/open-component-model/ocm/pkg/common/accessio"
	"github.com/open-component-model/ocm/pkg/mime"
	"k8s.io/apimachinery/pkg/util/validation/field"
)

type Handler struct{}

var _ inputs.InputHandler = (*Handler)(nil)

func (h *Handler) Validate(fldPath *field.Path, ctx clictx.Context, input *inputs.BlobInput, inputFilePath string) field.ErrorList {
	allErrs := inputs.ForbidFilePattern(fldPath, input)
	path := fldPath.Child("path")
	if input.Path == "" {
		allErrs = append(allErrs, field.Required(path, "path is required for input"))
	} else {
		fileInfo, filePath, err := inputs.FileInfo(ctx, input.Path, inputFilePath)
		if err != nil {
			allErrs = append(allErrs, field.Invalid(path, filePath, err.Error()))
		} else {
			if fileInfo.Mode()&os.ModeType != 0 {
				allErrs = append(allErrs, field.Invalid(path, filePath, "no regular file"))
			}
		}
	}
	return allErrs
}

func (h *Handler) GetBlob(ctx clictx.Context, input *inputs.BlobInput, inputFilePath string) (accessio.TemporaryBlobAccess, string, error) {
	fs := ctx.FileSystem()
	inputInfo, inputPath, err := inputs.FileInfo(ctx, input.Path, inputFilePath)
	if inputInfo.IsDir() {
		return nil, "", fmt.Errorf("resource type is file but a directory was provided")
	}
	// otherwise just open the file
	inputBlob, err := fs.Open(inputPath)
	if err != nil {
		return nil, "", fmt.Errorf("unable to read input blob from %q: %w", inputPath, err)
	}

	if !input.Compress() {
		inputBlob.Close()
		return accessio.BlobNopCloser(accessio.BlobAccessForFile(input.MediaType, inputPath, fs)), "", nil
	}

	temp, err := accessio.NewTempFile(fs, "", "compressed*.gzip")
	if err != nil {
		return nil, "", err
	}
	defer temp.Close()

	input.SetMediaTypeIfNotDefined(mime.MIME_GZIP)
	gw := gzip.NewWriter(temp.Writer())
	if _, err := io.Copy(gw, inputBlob); err != nil {
		return nil, "", fmt.Errorf("unable to compress input file %q: %w", inputPath, err)
	}
	if err := gw.Close(); err != nil {
		return nil, "", fmt.Errorf("unable to close gzip writer: %w", err)
	}

	return temp.AsBlob(input.MediaType), "", nil
}

func (h *Handler) Usage() string {
	return `
- <code>file</code>

  The path must denote a file relative the the resources file.
  The content is compressed if the <code>compress</code> field
  is set to <code>true</code>.
`
}