// SPDX-FileCopyrightText: 2022 SAP SE or an SAP affiliate company and Open Component Model contributors.
//
// SPDX-License-Identifier: Apache-2.0

package download_test

import (
	"bytes"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/open-component-model/ocm/cmds/ocm/testhelper"
	. "github.com/open-component-model/ocm/pkg/testutils"

	"github.com/open-component-model/ocm/pkg/common/accessio"
	"github.com/open-component-model/ocm/pkg/contexts/oci/grammar"
	"github.com/open-component-model/ocm/pkg/contexts/oci/repositories/artefactset"
	"github.com/open-component-model/ocm/pkg/mime"
)

const ARCH = "/tmp/ctf"
const VERSION = "v1"
const NS = "mandelsoft/test"
const OUT = "/tmp/res"

var _ = Describe("Test Environment", func() {
	var env *TestEnv

	BeforeEach(func() {
		env = NewTestEnv()
	})

	AfterEach(func() {
		env.Cleanup()
	})

	It("download single artefact from ctf file", func() {
		env.OCICommonTransport(ARCH, accessio.FormatDirectory, func() {
			env.Namespace(NS, func() {
				env.Manifest(VERSION, func() {
					env.Config(func() {
						env.BlobStringData(mime.MIME_JSON, "{}")
					})
					env.Layer(func() {
						env.BlobStringData(mime.MIME_TEXT, "testdata")
					})
				})
			})
		})

		buf := bytes.NewBuffer(nil)
		Expect(env.CatchOutput(buf).Execute("download", "artefact", "-O", OUT, "--repo", ARCH, NS+grammar.TagSeparator+VERSION)).To(Succeed())
		Expect(buf.String()).To(StringEqualTrimmedWithContext(
			`
/tmp/res: downloaded
`))
		Expect(env.DirExists(OUT)).To(BeTrue())
		tags := ""
		if artefactset.IsOCIDefaultFormat() {
			tags = ",\"org.opencontainers.image.ref.name\":\"v1\""
		}
		sha := "sha256:2c3e2c59e0ac9c99864bf0a9f9727c09f21a66080f9f9b03b36a2dad3cce6ff9"
		Expect(env.ReadFile(OUT + "/" + artefactset.DefaultArtefactSetDescriptorFileName)).To(Equal([]byte("{\"schemaVersion\":2,\"mediaType\":\"application/vnd.oci.image.index.v1+json\",\"manifests\":[{\"mediaType\":\"application/vnd.oci.image.manifest.v1+json\",\"digest\":\"sha256:2c3e2c59e0ac9c99864bf0a9f9727c09f21a66080f9f9b03b36a2dad3cce6ff9\",\"size\":342,\"annotations\":{\"cloud.gardener.ocm/tags\":\"v1\"" + tags + ",\"software.ocm/tags\":\"v1\"}}],\"annotations\":{\"cloud.gardener.ocm/main\":\"" + sha + "\",\"software.ocm/main\":\"" + sha + "\"}}")))
	})
})
