// SPDX-FileCopyrightText: 2022 SAP SE or an SAP affiliate company and Open Component Model contributors.
//
// SPDX-License-Identifier: Apache-2.0

package ocirepo_test

import (
	"encoding/json"
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/open-component-model/ocm/pkg/contexts/oci/testhelper"
	. "github.com/open-component-model/ocm/pkg/env"
	. "github.com/open-component-model/ocm/pkg/env/builder"
	. "github.com/open-component-model/ocm/pkg/testutils"

	"github.com/open-component-model/ocm/pkg/common/accessio"
	"github.com/open-component-model/ocm/pkg/common/accessobj"
	"github.com/open-component-model/ocm/pkg/contexts/oci"
	"github.com/open-component-model/ocm/pkg/contexts/oci/artdesc"
	ocictf "github.com/open-component-model/ocm/pkg/contexts/oci/repositories/ctf"
	"github.com/open-component-model/ocm/pkg/contexts/ocm/accessmethods/ociartefact"
	storagecontext "github.com/open-component-model/ocm/pkg/contexts/ocm/blobhandler/oci"
	"github.com/open-component-model/ocm/pkg/contexts/ocm/blobhandler/oci/ocirepo"
	metav1 "github.com/open-component-model/ocm/pkg/contexts/ocm/compdesc/meta/v1"
	"github.com/open-component-model/ocm/pkg/contexts/ocm/cpi"
	"github.com/open-component-model/ocm/pkg/contexts/ocm/repositories/ctf"
	"github.com/open-component-model/ocm/pkg/contexts/ocm/repositories/genericocireg"
	"github.com/open-component-model/ocm/pkg/contexts/ocm/resourcetypes"
	"github.com/open-component-model/ocm/pkg/contexts/ocm/transfer"
	"github.com/open-component-model/ocm/pkg/contexts/ocm/transfer/transferhandler/standard"
	"github.com/open-component-model/ocm/pkg/mime"
)

const ARCH = "/tmp/ctf"
const ARCH2 = "/tmp/ctf2"
const PROVIDER = "mandelsoft"
const VERSION = "v1"
const COMPONENT = "github.com/mandelsoft/test"
const COMPONENT2 = "github.com/mandelsoft/test2"
const OUT = "/tmp/res"
const OCIPATH = "/tmp/oci"
const OCIHOST = "alias"

func FakeOCIRegBaseFunction(ctx *storagecontext.StorageContext) string {
	return "baseurl.io"
}

var _ = Describe("oci artefact transfer", func() {
	var env *Builder
	var ldesc *artdesc.Descriptor

	BeforeEach(func() {
		env = NewBuilder(NewEnvironment())

		env.OCICommonTransport(OCIPATH, accessio.FormatDirectory, func() {
			ldesc = OCIManifest1(env)
		})

		env.OCMCommonTransport(ARCH, accessio.FormatDirectory, func() {
			env.Component(COMPONENT, func() {
				env.Version(VERSION, func() {
					env.Provider(PROVIDER)
					env.Resource("testdata", "", "PlainText", metav1.LocalRelation, func() {
						env.BlobStringData(mime.MIME_TEXT, "testdata")
					})
					env.Resource("artefact", "", resourcetypes.OCI_IMAGE, metav1.LocalRelation, func() {
						env.Access(
							ociartefact.New(oci.StandardOCIRef(OCIHOST+".alias", OCINAMESPACE, OCIVERSION)),
						)
					})
				})
			})
		})

		FakeOCIRepo(env, OCIPATH, OCIHOST)

		_ = ldesc
	})

	AfterEach(func() {
		env.Cleanup()
	})

	It("it should copy a resource by value to a ctf file", func() {

		env.OCMContext().BlobHandlers().Register(ocirepo.NewArtefactHandler(FakeOCIRegBaseFunction),
			cpi.ForRepo(oci.CONTEXT_TYPE, ocictf.Type), cpi.ForMimeType(artdesc.ToContentMediaType(artdesc.MediaTypeImageManifest)))

		src := Must(ctf.Open(env.OCMContext(), accessobj.ACC_READONLY, ARCH, 0, env))
		cv := Must(src.LookupComponentVersion(COMPONENT, VERSION))
		tgt := Must(ctf.Create(env.OCMContext(), accessobj.ACC_WRITABLE|accessobj.ACC_CREATE, OUT, 0700, accessio.FormatDirectory, env))
		defer tgt.Close()

		opts := &standard.Options{}
		opts.SetResourcesByValue(true)
		handler := standard.NewDefaultHandler(opts)

		MustBeSuccessful(transfer.TransferVersion(nil, nil, cv, tgt, handler))
		Expect(env.DirExists(OUT)).To(BeTrue())

		list := Must(tgt.ComponentLister().GetComponents("", true))
		Expect(list).To(Equal([]string{COMPONENT}))
		comp := Must(tgt.LookupComponentVersion(COMPONENT, VERSION))
		Expect(len(comp.GetDescriptor().Resources)).To(Equal(2))
		data := Must(json.Marshal(comp.GetDescriptor().Resources[1].Access))

		fmt.Printf("%s\n", string(data))
		Expect(string(data)).To(StringEqualWithContext("{\"imageReference\":\"baseurl.io/ocm/value:v2.0\",\"type\":\"ociArtefact\"}"))

		ocirepo := tgt.(genericocireg.OCIBasedRepository).OCIRepository()

		art := Must(ocirepo.LookupArtefact(OCINAMESPACE, OCIVERSION))
		defer Close(art, "artefact")

		man := MustBeNonNil(art.ManifestAccess())
		Expect(len(man.GetDescriptor().Layers)).To(Equal(1))
		Expect(man.GetDescriptor().Layers[0].Digest).To(Equal(ldesc.Digest))

		blob := Must(man.GetBlob(ldesc.Digest))
		data = Must(blob.Get())
		Expect(string(data)).To(Equal(OCILAYER))
	})
})
