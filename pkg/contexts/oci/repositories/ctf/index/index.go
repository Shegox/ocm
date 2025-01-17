// SPDX-FileCopyrightText: 2022 SAP SE or an SAP affiliate company and Open Component Model contributors.
//
// SPDX-License-Identifier: Apache-2.0

package index

import (
	"sort"
	"strings"
	"sync"

	"github.com/opencontainers/go-digest"
	"github.com/opencontainers/image-spec/specs-go"

	"github.com/open-component-model/ocm/pkg/contexts/oci/cpi"
)

type RepositoryIndex struct {
	lock         sync.RWMutex
	byDigest     map[digest.Digest][]*ArtefactMeta
	byRepository map[string]map[string]*ArtefactMeta
}

func NewMeta(repo string, tag string, digest digest.Digest) *ArtefactMeta {
	return &ArtefactMeta{
		Repository: repo,
		Tag:        tag,
		Digest:     digest,
	}
}

func NewRepositoryIndex() *RepositoryIndex {
	return &RepositoryIndex{
		byDigest:     map[digest.Digest][]*ArtefactMeta{},
		byRepository: map[string]map[string]*ArtefactMeta{},
	}
}

func (r *RepositoryIndex) RepositoryList() []string {
	result := []string{}
	for k := range r.byRepository {
		result = append(result, k)
	}
	return result
}

func (r *RepositoryIndex) AddTagsFor(repo string, digest digest.Digest, tags ...string) error {
	r.lock.Lock()
	defer r.lock.Unlock()

	a := r.getArtefactInfo(repo, digest.String())
	if a == nil {
		return cpi.ErrUnknownArtefact(repo, digest.String())
	}
	for _, tag := range tags {
		n := *a
		n.Tag = tag
		r.addArtefactInfo(n)
	}
	return nil
}

func (r *RepositoryIndex) AddArtefactInfo(n *ArtefactMeta) {
	r.lock.Lock()
	defer r.lock.Unlock()
	r.addArtefactInfo(*n)
}

func (r *RepositoryIndex) addArtefactInfo(m ArtefactMeta) {
	repos := r.byRepository[m.Repository]
	if len(repos) == 0 {
		repos = map[string]*ArtefactMeta{}
		r.byRepository[m.Repository] = repos
	}

	list := r.byDigest[m.Digest]
	if list == nil {
		list = []*ArtefactMeta{&m}
	} else {
		for _, e := range list {
			if m.Repository == e.Repository && m.Digest == e.Digest {
				if e.Tag == "" || e.Tag == m.Tag {
					e.Tag = m.Tag
					if e.Tag != "" {
						repos[m.Tag] = e
					}
					return
				}
			}
		}
		list = append(list, &m)
	}
	r.byDigest[m.Digest] = list

	repos["@"+m.Digest.String()] = &m
	if m.Tag != "" {
		repos[m.Tag] = &m
	}
}

func (r *RepositoryIndex) HasArtefact(repo, tag string) bool {
	r.lock.RLock()
	defer r.lock.RUnlock()
	repos := r.byRepository[repo]
	if repos == nil {
		return false
	}
	m := repos[tag]
	return m != nil
}

func (r *RepositoryIndex) GetTags(repo string) []string {
	r.lock.RLock()
	defer r.lock.RUnlock()

	repos := r.byRepository[repo]
	if repos == nil {
		return nil
	}
	result := []string{}
	digests := map[digest.Digest]bool{}
	for t, a := range repos {
		if !strings.HasPrefix(t, "@") {
			result = append(result, t)
			digests[a.Digest] = true
		} else if !digests[a.Digest] {
			digests[a.Digest] = false
		}
	}
	/* TODO: how to query untagged entries at api level
	for d, found := range digests {
		if !found {
			result = append(result, "@"+d.String())
		}
	}
	*/
	return result
}

func (r *RepositoryIndex) GetArtefacts(repo string) []string {
	r.lock.RLock()
	defer r.lock.RUnlock()

	repos := r.byRepository[repo]
	if repos == nil {
		return nil
	}
	result := []string{}
	for t := range repos {
		result = append(result, t)
	}
	return result
}

func (r *RepositoryIndex) GetArtefactInfos(digest digest.Digest) []*ArtefactMeta {
	r.lock.RLock()
	defer r.lock.RUnlock()
	return r.byDigest[digest]
}

func (r *RepositoryIndex) GetArtefactInfo(repo, reference string) *ArtefactMeta {
	r.lock.RLock()
	defer r.lock.RUnlock()
	return r.getArtefactInfo(repo, reference)
}

func (r *RepositoryIndex) getArtefactInfo(repo, reference string) *ArtefactMeta {
	repos := r.byRepository[repo]
	if repos == nil {
		return nil
	}
	m := repos[reference]
	if m == nil && !strings.HasPrefix(reference, "@") {
		m = repos["@"+reference]
	}
	if m == nil {
		return nil
	}
	result := *m
	return &result
}

func (r *RepositoryIndex) GetDescriptor() *ArtefactIndex {
	r.lock.RLock()
	defer r.lock.RUnlock()
	index := &ArtefactIndex{
		Versioned: specs.Versioned{SchemaVersion},
	}

	repos := make([]string, len(r.byRepository))
	i := 0
	for repo := range r.byRepository {
		repos[i] = repo
		i++
	}
	sort.Strings(repos)
	for _, name := range repos {
		repo := r.byRepository[name]
		versions := make([]string, len(repo))
		i := 0
		for vers := range repo {
			versions[i] = vers
			i++
		}
		sort.Strings(versions)

		for _, name := range versions {
			vers := repo[name]
			if "@"+vers.Digest.String() != name || vers.Tag == "" {
				d := &ArtefactMeta{
					Repository: vers.Repository,
					Tag:        vers.Tag,
					Digest:     vers.Digest,
				}
				index.Index = append(index.Index, *d)
			}
		}
	}
	return index
}
