// SPDX-FileCopyrightText: 2022 SAP SE or an SAP affiliate company and Open Component Model contributors.
//
// SPDX-License-Identifier: Apache-2.0

package get_test

import (
	"bytes"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/open-component-model/ocm/cmds/ocm/testhelper"
	. "github.com/open-component-model/ocm/pkg/testutils"

	"github.com/open-component-model/ocm/pkg/common/accessio"
	"github.com/open-component-model/ocm/pkg/contexts/ocm"
)

const ARCH = "/tmp/ca"
const VERSION = "v1"
const COMP = "test.de/x"
const COMP2 = "test.de/y"
const PROVIDER = "mandelsoft"

var _ = Describe("Test Environment", func() {
	var env *TestEnv

	spec, err := ocm.NewGenericAccessSpec("{\"type\":\"git\"}")
	Expect(err).To(Succeed())

	BeforeEach(func() {
		env = NewTestEnv()
	})

	AfterEach(func() {
		env.Cleanup()
	})

	It("lists single source in component archive", func() {
		env.ComponentArchive(ARCH, accessio.FormatDirectory, COMP, VERSION, func() {
			env.Provider(PROVIDER)
			env.Source("testdata", "v1", "git", func() {
				env.Access(spec)
			})
		})

		buf := bytes.NewBuffer(nil)
		Expect(env.CatchOutput(buf).Execute("get", "sources", ARCH)).To(Succeed())
		Expect(buf.String()).To(StringEqualTrimmedWithContext(
			`
NAME     VERSION IDENTITY TYPE
testdata v1               git
`))
	})

	It("lists ambigious source in component archive", func() {
		env.ComponentArchive(ARCH, accessio.FormatDirectory, COMP, VERSION, func() {
			env.Provider(PROVIDER)
			env.Source("testdata", "v1", "git", func() {
				env.Access(spec)
				env.ExtraIdentity("platform", "a")
			})
			env.Source("testdata", "v1", "git", func() {
				env.Access(spec)
				env.ExtraIdentity("platform", "b")
			})
		})

		buf := bytes.NewBuffer(nil)
		Expect(env.CatchOutput(buf).Execute("get", "sources", ARCH)).To(Succeed())
		Expect(buf.String()).To(StringEqualTrimmedWithContext(
			`
NAME     VERSION IDENTITY       TYPE
testdata v1      "platform"="a" git
testdata v1      "platform"="b" git
`))
	})

	It("lists single source in ctf file", func() {
		env.OCMCommonTransport(ARCH, accessio.FormatDirectory, func() {
			env.Component(COMP, func() {
				env.Version(VERSION, func() {
					env.Provider(PROVIDER)
					env.Source("testdata", "v1", "git", func() {
						env.Access(spec)
					})
				})
			})
		})

		buf := bytes.NewBuffer(nil)
		Expect(env.CatchOutput(buf).Execute("get", "sources", ARCH)).To(Succeed())
		Expect(buf.String()).To(StringEqualTrimmedWithContext(
			`
NAME     VERSION IDENTITY TYPE
testdata v1               git
`))
	})

	Context("with closure", func() {
		BeforeEach(func() {
			env.OCMCommonTransport(ARCH, accessio.FormatDirectory, func() {
				env.Component(COMP, func() {
					env.Version(VERSION, func() {
						env.Provider(PROVIDER)
						env.Source("testdata", "v1", "git", func() {
							env.Access(spec)
						})
					})
				})
				env.Component(COMP2, func() {
					env.Version(VERSION, func() {
						env.Provider(PROVIDER)
						env.Source("source", "v1", "git", func() {
							env.Access(spec)
						})
						env.Reference("base", COMP, VERSION)
					})
				})
			})
		})

		It("lists resource closure in ctf file", func() {
			buf := bytes.NewBuffer(nil)
			Expect(env.CatchOutput(buf).Execute("get", "sources", "-r", "--repo", ARCH, COMP2+":"+VERSION)).To(Succeed())
			Expect(buf.String()).To(StringEqualTrimmedWithContext(
				`
REFERENCEPATH              NAME     VERSION IDENTITY TYPE
test.de/y:v1               source   v1               git
test.de/y:v1->test.de/x:v1 testdata v1               git
`))
		})
		It("lists flat tree in ctf file", func() {
			buf := bytes.NewBuffer(nil)
			Expect(env.CatchOutput(buf).Execute("get", "sources", "-o", "tree", "--repo", ARCH, COMP2+":"+VERSION)).To(Succeed())
			Expect(buf.String()).To(StringEqualTrimmedWithContext(
				`
COMPONENTVERSION    NAME   VERSION IDENTITY TYPE
└─ test.de/y:v1                             
   └─               source v1               git
`))
		})

		It("lists resource closure in ctf file", func() {
			buf := bytes.NewBuffer(nil)
			Expect(env.CatchOutput(buf).Execute("get", "sources", "-r", "-o", "tree", "--repo", ARCH, COMP2+":"+VERSION)).To(Succeed())
			Expect(buf.String()).To(StringEqualTrimmedWithContext(
				`
COMPONENTVERSION       NAME     VERSION IDENTITY TYPE
└─ test.de/y:v1                                  
   ├─                  source   v1               git
   └─ test.de/x:v1                               
      └─               testdata v1               git
`))
		})
	})
})
