// SPDX-FileCopyrightText: 2022 SAP SE or an SAP affiliate company and Open Component Model contributors.
//
// SPDX-License-Identifier: Apache-2.0

package common

import (
	"github.com/open-component-model/ocm/cmds/ocm/commands/ocmcmds/common/handlers/elemhdlr"
	"github.com/open-component-model/ocm/cmds/ocm/pkg/output"
	"github.com/open-component-model/ocm/cmds/ocm/pkg/utils"
	"github.com/open-component-model/ocm/pkg/contexts/clictx"
	"github.com/open-component-model/ocm/pkg/contexts/ocm"
	"github.com/open-component-model/ocm/pkg/contexts/ocm/compdesc"
)

func Elem(e interface{}) *compdesc.Source {
	return e.(*elemhdlr.Object).Element.(*compdesc.Source)
}

////////////////////////////////////////////////////////////////////////////////

type TypeHandler struct {
	*elemhdlr.TypeHandler
}

func NewTypeHandler(octx clictx.OCM, opts *output.Options, repo ocm.Repository, session ocm.Session, compspecs []string, hopts ...elemhdlr.Option) (utils.TypeHandler, error) {
	return elemhdlr.NewTypeHandler(octx, opts, repo, session, ocm.KIND_SOURCE, compspecs, func(access ocm.ComponentVersionAccess) compdesc.ElementAccessor {
		return access.GetDescriptor().Sources
	}, hopts...)
}
