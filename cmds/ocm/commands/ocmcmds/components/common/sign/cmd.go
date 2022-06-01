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

package sign

import (
	"fmt"

	"github.com/open-component-model/ocm/cmds/ocm/commands/common/options/signoption"
	"github.com/open-component-model/ocm/cmds/ocm/pkg/output"
	"github.com/open-component-model/ocm/pkg/common"
	"github.com/open-component-model/ocm/pkg/contexts/ocm/signing"
	"github.com/open-component-model/ocm/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/open-component-model/ocm/cmds/ocm/clictx"
	ocmcommon "github.com/open-component-model/ocm/cmds/ocm/commands/ocmcmds/common"
	"github.com/open-component-model/ocm/cmds/ocm/commands/ocmcmds/common/handlers/comphdlr"
	"github.com/open-component-model/ocm/cmds/ocm/commands/ocmcmds/common/options/repooption"
	"github.com/open-component-model/ocm/cmds/ocm/pkg/utils"
	"github.com/open-component-model/ocm/pkg/contexts/ocm"
)

type SignatureCommand struct {
	utils.BaseCommand
	Refs []string
	op   *Operation
}

type Operation struct {
	op      string
	sign    bool
	example string
	terms   []string
}

func NewOperation(op string, sign bool, terms []string, example string) *Operation {
	return &Operation{
		op:      op,
		sign:    sign,
		example: example,
		terms:   terms,
	}
}

// NewSigningCommand creates a new ctf command.
func NewSigningCommand(ctx clictx.Context, op *Operation, names ...string) *cobra.Command {
	return utils.SetupCommand(&SignatureCommand{op: op, BaseCommand: utils.NewBaseCommand(ctx, repooption.New(), signoption.New(op.sign))}, names...)
}

func (o *SignatureCommand) ForName(name string) *cobra.Command {
	return &cobra.Command{
		Use:   "[<options>] {<component-reference>}",
		Short: o.op.op + " component version",
		Long: `
` + o.op.op + ` specified component versions. 
`,
		Example: o.op.example,
	}
}

func (o *SignatureCommand) Complete(args []string) error {
	o.Refs = args
	if len(args) == 0 && repooption.From(o).Spec == "" {
		return fmt.Errorf("a repository or at least one argument that defines the reference is needed")
	}
	return nil
}

func (o *SignatureCommand) Run() error {
	session := ocm.NewSession(nil)
	defer session.Close()

	err := o.ProcessOnOptions(ocmcommon.CompleteOptionsWithContext(o, session))
	if err != nil {
		return err
	}
	sign := signoption.From(o)
	repo := repooption.From(o).Repository
	handler := comphdlr.NewTypeHandler(o.Context.OCM(), session, repo)
	sopts := signing.NewOptions(sign, signing.Resolver(repo))
	err = sopts.Complete(nil)
	if err != nil {
		return err
	}
	return utils.HandleOutput(NewAction(o.op.terms, o, sopts), handler, utils.StringElemSpecs(o.Refs...)...)
}

/////////////////////////////////////////////////////////////////////////////

type action struct {
	desc    []string
	cmd     *SignatureCommand
	printer common.Printer
	state   common.WalkingState
	sopts   *signing.Options
	errlist *errors.ErrorList
}

var _ output.Output = (*action)(nil)

func NewAction(desc []string, cmd *SignatureCommand, sopts *signing.Options) output.Output {
	return &action{
		desc:    desc,
		cmd:     cmd,
		printer: common.NewPrinter(cmd.Context.StdOut()),
		state:   common.NewWalkingState(),
		sopts:   sopts,
		errlist: errors.ErrListf(desc[1]),
	}
}

func (a *action) Add(e interface{}) error {
	cv := comphdlr.Elem(e)
	d, err := signing.Apply(a.printer, &a.state, cv, a.sopts)
	a.errlist.Add(err)
	if err == nil {
		a.printer.Printf("successfully %s %s:%s (digest %s:%s)\n", a.desc[0], cv.GetName(), cv.GetVersion(), d.HashAlgorithm, d.Value)
	} else {
		a.printer.Printf("failed %s %s:%s: %s\n", a.desc[1], cv.GetName(), cv.GetVersion(), err)
	}
	return nil
}

func (a *action) Close() error {
	return nil
}

func (a *action) Out() error {
	if a.errlist.Len() > 0 {
		a.printer.Printf("finished with %d error(s)\n", a.errlist.Len())
	}
	return a.errlist.Result()
}