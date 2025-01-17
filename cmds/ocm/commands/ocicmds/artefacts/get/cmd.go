// SPDX-FileCopyrightText: 2022 SAP SE or an SAP affiliate company and Open Component Model contributors.
//
// SPDX-License-Identifier: Apache-2.0

package get

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/open-component-model/ocm/cmds/ocm/commands/common/options/closureoption"
	"github.com/open-component-model/ocm/cmds/ocm/commands/ocicmds/common"
	"github.com/open-component-model/ocm/cmds/ocm/commands/ocicmds/common/handlers/artefacthdlr"
	"github.com/open-component-model/ocm/cmds/ocm/commands/ocicmds/common/options/repooption"
	"github.com/open-component-model/ocm/cmds/ocm/commands/ocicmds/names"
	"github.com/open-component-model/ocm/cmds/ocm/commands/verbs"
	"github.com/open-component-model/ocm/cmds/ocm/pkg/options"
	"github.com/open-component-model/ocm/cmds/ocm/pkg/output"
	"github.com/open-component-model/ocm/cmds/ocm/pkg/processing"
	"github.com/open-component-model/ocm/cmds/ocm/pkg/utils"
	"github.com/open-component-model/ocm/pkg/contexts/clictx"
	"github.com/open-component-model/ocm/pkg/contexts/oci"
)

var (
	Names = names.Artefacts
	Verb  = verbs.Get
)

type Command struct {
	utils.BaseCommand

	Refs []string
}

// NewCommand creates a new artefact command.
func NewCommand(ctx clictx.Context, names ...string) *cobra.Command {
	return utils.SetupCommand(&Command{BaseCommand: utils.NewBaseCommand(ctx, repooption.New(), output.OutputOptions(outputs, &Attached{}, closureoption.New("index")))}, utils.Names(Names, names...)...)
}

func (o *Command) ForName(name string) *cobra.Command {
	return &cobra.Command{
		Use:   "[<options>] {<artefact-reference>}",
		Short: "get artefact version",
		Long: `
Get lists all artefact versions specified, if only a repository is specified
all tagged artefacts are listed.
	`,
		Example: `
$ ocm get artefact ghcr.io/mandelsoft/kubelink
$ ocm get artefact --repo OCIRegistry:ghcr.io mandelsoft/kubelink
`,
	}
}

func (o *Command) Complete(args []string) error {
	if len(args) == 0 && repooption.From(o).Spec == "" {
		return fmt.Errorf("a repository or at least one argument that defines the reference is needed")
	}
	o.Refs = args
	return nil
}

func (o *Command) Run() error {
	session := oci.NewSession(nil)
	defer session.Close()
	err := o.ProcessOnOptions(common.CompleteOptionsWithContext(o.Context, session))
	if err != nil {
		return err
	}
	handler := artefacthdlr.NewTypeHandler(o.Context.OCI(), session, repooption.From(o).Repository)
	return utils.HandleArgs(output.From(o), handler, o.Refs...)
}

/////////////////////////////////////////////////////////////////////////////

func OutputChainFunction() output.ChainFunction {
	return func(opts *output.Options) processing.ProcessChain {
		chain := closureoption.Closure(opts, artefacthdlr.ClosureExplode, artefacthdlr.Sort)
		chain = processing.Append(chain, artefacthdlr.ExplodeAttached, AttachedFrom(opts))
		return processing.Append(chain, artefacthdlr.Clean, options.Or(closureoption.From(opts), output.OutputModeCondition(opts, "tree")))
	}
}

func TableOutput(opts *output.Options, mapping processing.MappingFunction, wide ...string) *output.TableOutput {
	return &output.TableOutput{
		Headers: output.Fields("REGISTRY", "REPOSITORY", "KIND", "TAG", "DIGEST", wide),
		Chain:   OutputChainFunction()(opts),
		Options: opts,
		Mapping: mapping,
	}
}

var outputs = output.NewOutputs(getRegular, output.Outputs{
	"wide": getWide,
	"tree": getTree,
}).AddChainedManifestOutputs(OutputChainFunction())

func getRegular(opts *output.Options) output.Output {
	return closureoption.TableOutput(TableOutput(opts, mapGetRegularOutput)).New()
}

func getWide(opts *output.Options) output.Output {
	return closureoption.TableOutput(TableOutput(opts, mapGetWideOutput, "MIMETYPE", "CONFIGTYPE")).New()
}

func getTree(opts *output.Options) output.Output {
	return output.TreeOutput(TableOutput(opts, mapGetRegularOutput), "NESTING").New()
}

func mapGetRegularOutput(e interface{}) interface{} {
	digest := "unknown"
	p := e.(*artefacthdlr.Object)
	blob, err := p.Artefact.Blob()
	if err == nil {
		digest = blob.Digest().String()
	}
	tag := "-"
	if p.Spec.Tag != nil {
		tag = *p.Spec.Tag
	}
	kind := "-"
	if p.Artefact.IsManifest() {
		kind = "manifest"
	}
	if p.Artefact.IsIndex() {
		kind = "index"
	}
	return []string{p.Spec.UniformRepositorySpec.String(), p.Spec.Repository, kind, tag, digest}
}

func mapGetWideOutput(e interface{}) interface{} {
	p := e.(*artefacthdlr.Object)

	config := "-"
	if p.Artefact.IsManifest() {
		config = p.Artefact.ManifestAccess().GetDescriptor().Config.MediaType
	}
	return output.Fields(mapGetRegularOutput(e), p.Artefact.GetDescriptor().MimeType(), config)
}
