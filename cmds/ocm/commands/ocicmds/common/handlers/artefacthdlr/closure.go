// SPDX-FileCopyrightText: 2022 SAP SE or an SAP affiliate company and Open Component Model contributors.
//
// SPDX-License-Identifier: Apache-2.0

package artefacthdlr

import (
	"github.com/open-component-model/ocm/cmds/ocm/pkg/output"
	"github.com/open-component-model/ocm/pkg/common"
	"github.com/open-component-model/ocm/pkg/contexts/oci"
	"github.com/open-component-model/ocm/pkg/out"
)

////////////////////////////////////////////////////////////////////////////////

func ClosureExplode(opts *output.Options, e interface{}) []interface{} {
	return traverse(common.History{}, e.(*Object), opts.Context)
}

func traverse(hist common.History, o *Object, octx out.Context) []output.Object {
	blob, _ := o.Artefact.Blob()
	key := common.NewNameVersion("", blob.Digest().String())
	if err := hist.Add(oci.KIND_OCIARTEFACT, key); err != nil {
		return nil
	}
	result := []output.Object{o}
	if o.Artefact.IsIndex() {
		refs := o.Artefact.IndexAccess().GetDescriptor().Manifests

		found := map[common.NameVersion]bool{}
		for _, ref := range refs {
			key := common.NewNameVersion("", ref.Digest.String())
			if found[key] {
				continue // skip same ref wit different attributes for recursion
			}
			found[key] = true
			nested, err := o.Namespace.GetArtefact(key.GetVersion())
			if err != nil {
				out.Errf(octx, "Warning: lookup nested artefact %q [%s]: %s\n", ref.Digest, hist, err)
			}
			obj := &Object{
				History: hist.Copy(),
				Key:     Key(nested),
				Spec: oci.RefSpec{
					UniformRepositorySpec: o.Spec.UniformRepositorySpec,
					Repository:            o.Spec.Repository,
					Digest:                &ref.Digest,
				},
				Namespace: o.Namespace,
				Artefact:  nested,
			}
			result = append(result, traverse(hist, obj, octx)...)
		}
	}
	output.Print(result, "traverse %s", blob.Digest())
	return result
}
