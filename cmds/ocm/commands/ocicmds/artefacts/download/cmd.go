// SPDX-FileCopyrightText: 2022 SAP SE or an SAP affiliate company and Open Component Model contributors.
//
// SPDX-License-Identifier: Apache-2.0

package download

import (
	"strings"

	"github.com/mandelsoft/vfs/pkg/vfs"
	"github.com/spf13/cobra"

	"github.com/open-component-model/ocm/cmds/ocm/commands/common/options/destoption"
	"github.com/open-component-model/ocm/cmds/ocm/commands/common/options/formatoption"
	"github.com/open-component-model/ocm/cmds/ocm/commands/ocicmds/common"
	"github.com/open-component-model/ocm/cmds/ocm/commands/ocicmds/common/handlers/artefacthdlr"
	"github.com/open-component-model/ocm/cmds/ocm/commands/ocicmds/common/options/repooption"
	"github.com/open-component-model/ocm/cmds/ocm/commands/ocicmds/names"
	"github.com/open-component-model/ocm/cmds/ocm/commands/verbs"
	"github.com/open-component-model/ocm/cmds/ocm/pkg/output"
	"github.com/open-component-model/ocm/cmds/ocm/pkg/utils"
	"github.com/open-component-model/ocm/pkg/common/accessio"
	"github.com/open-component-model/ocm/pkg/common/accessobj"
	"github.com/open-component-model/ocm/pkg/contexts/clictx"
	"github.com/open-component-model/ocm/pkg/contexts/oci"
	"github.com/open-component-model/ocm/pkg/contexts/oci/repositories/artefactset"
	"github.com/open-component-model/ocm/pkg/errors"
	"github.com/open-component-model/ocm/pkg/out"
)

var (
	Names = names.Artefacts
	Verb  = verbs.Download
)

type Command struct {
	utils.BaseCommand

	Refs []string
}

// NewCommand creates a new download command.
func NewCommand(ctx clictx.Context, names ...string) *cobra.Command {
	return utils.SetupCommand(&Command{BaseCommand: utils.NewBaseCommand(ctx, repooption.New(), output.OutputOptions(outputs, destoption.New(), &formatoption.Option{}))}, utils.Names(Names, names...)...)
}

func (o *Command) ForName(name string) *cobra.Command {
	return &cobra.Command{
		Use:   "[<options>]  {<artefact>} ",
		Args:  cobra.MinimumNArgs(1),
		Short: "download oci artefacts",
		Long: `
Download artefacts from an OCI registry. The result is stored in
artefact set format, without the repository part

The files are named according to the artefact repository name.
`,
	}
}

func (o *Command) Complete(args []string) error {
	var err error
	o.Refs = args
	return err
}

func (o *Command) Run() error {
	session := oci.NewSession(nil)
	defer session.Close()

	err := o.ProcessOnOptions(common.CompleteOptionsWithContext(o.Context, session))
	if err != nil {
		return err
	}

	hdlr := artefacthdlr.NewTypeHandler(o.Context.OCI(), session, repooption.From(o).Repository)
	return utils.HandleArgs(output.From(o), hdlr, o.Refs...)
}

////////////////////////////////////////////////////////////////////////////////

var outputs = output.NewOutputs(getDownload)

func getDownload(opts *output.Options) output.Output {
	return &download{opts: opts}
}

type download struct {
	data artefacthdlr.Objects
	opts *output.Options
}

func (d *download) Add(e interface{}) error {
	d.data = append(d.data, e.(*artefacthdlr.Object))
	return nil
}

func (d *download) Close() error {
	return nil
}

func (d *download) Out() error {
	list := errors.ErrListf("downloading artefacts")
	dest := destoption.From(d.opts)
	if len(d.data) == 0 {
		out.Outf(d.opts.Context, "no artefacts found\n")
	}
	if len(d.data) == 1 {
		f := dest.Destination
		e := d.data[0]
		if f == "" {
			f = composePath(dest, e)
		}
		return d.Save(e, f)
	} else {
		for _, e := range d.data {
			f := composePath(dest, e)
			err := d.Save(e, f)
			if err != nil {
				list.Add(err)
				out.Outf(d.opts.Context, "%s failed: %s\n", f, err)
			}
		}
	}
	return list.Result()
}

func composePath(dest *destoption.Option, e *artefacthdlr.Object) string {
	f := e.Spec.UniformRepositorySpec.String()
	f = strings.ReplaceAll(f, "::", "-")
	f = vfs.Join(dest.PathFilesystem, f, e.Spec.Repository)
	if dest.Destination != "" {
		return vfs.Join(dest.PathFilesystem, dest.Destination, f)
	}
	return f
}

func (d *download) Save(o *artefacthdlr.Object, f string) error {
	dest := destoption.From(d.opts)
	art := o.Artefact
	dir := vfs.Dir(dest.PathFilesystem, f)

	err := dest.PathFilesystem.MkdirAll(dir, 0o770)
	if err != nil {
		return err
	}

	blob, err := art.Blob()
	if err != nil {
		return err
	}
	digest := blob.Digest()

	format := formatoption.From(d.opts)
	set, err := artefactset.Create(accessobj.ACC_CREATE, f, format.Mode(), format.Format, accessio.PathFileSystem(dest.PathFilesystem))
	if err != nil {
		return err
	}
	defer set.Close()
	err = artefactset.TransferArtefact(art, set)
	if err != nil {
		return err
	}

	if o.Spec.Tag != nil {
		err = set.AddTags(digest, *o.Spec.Tag)
		if err != nil {
			return err
		}
	}
	set.Annotate(artefactset.MAINARTEFACT_ANNOTATION, digest.String())
	set.Annotate(artefactset.LEGACY_MAINARTEFACT_ANNOTATION, digest.String())

	if err == nil {
		out.Outf(d.opts.Context, "%s: downloaded\n", f)
	}
	return err
}
