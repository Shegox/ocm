// SPDX-FileCopyrightText: 2022 SAP SE or an SAP affiliate company and Open Component Model contributors.
//
// SPDX-License-Identifier: Apache-2.0

package ociartefact

import (
	"github.com/open-component-model/ocm/pkg/cobrautils/flagsets"
	"github.com/open-component-model/ocm/pkg/contexts/ocm/accessmethods/options"
)

func ConfigHandler() flagsets.ConfigOptionTypeSetHandler {
	return flagsets.NewConfigOptionTypeSetHandler(
		Type, AddConfig,
		options.ReferenceOption,
	)
}

func AddConfig(opts flagsets.ConfigOptions, config flagsets.Config) error {
	flagsets.AddFieldByOptionP(opts, options.ReferenceOption, config, "imageReference")
	return nil
}

var usage = `
This method implements the access of an OCI artefact stored in an OCI registry.
`

var formatV1 = `
The type specific specification fields are:

- **<code>imageReference</code>** *string*

  OCI image/artefact reference following the possible docker schemes:
  - <code>&lt;repo>/&lt;artefact>:&lt;digest>@&lt;tag></code>
  - <code><host>[&lt;port>]/&lt;repo path>/&lt;artefact>:&lt;version>@&lt;tag></code>
`
