// SPDX-FileCopyrightText: 2022 SAP SE or an SAP affiliate company and Open Component Model contributors.
//
// SPDX-License-Identifier: Apache-2.0

package registry

import (
	"strings"
)

type Key[K any] interface {
	comparable
	IsValid() bool

	GetMediaType() string
	GetArtefactType() string

	SetArtefact(arttype, medtatype string) K
}

type Registry[H any, K Key[K]] struct {
	mappings map[K][]H
}

func NewRegistry[H any, K Key[K]]() *Registry[H, K] {
	return &Registry[H, K]{
		mappings: map[K][]H{},
	}
}

func (p *Registry[H, K]) lookupMedia(key K) []H {
	lookup := key
	for {
		if h, ok := p.mappings[lookup]; ok {
			return h
		}
		if i := strings.LastIndex(lookup.GetMediaType(), "+"); i > 0 {
			lookup = lookup.SetArtefact(lookup.GetArtefactType(), lookup.GetMediaType()[:i])
		} else {
			break
		}
	}
	return nil
}

func (p *Registry[H, K]) GetHandler(key K) []H {
	r := p.mappings[key]
	if r == nil {
		return nil
	}
	return append(r[:0:0], r...)
}

func (p *Registry[H, K]) LookupHandler(key K) []H {
	h := p.lookupMedia(key)
	if len(h) > 0 {
		return h
	}

	mediatype := key.GetMediaType()
	arttype := key.GetArtefactType()
	if mediatype == "" || arttype == "" {
		return h
	}
	if h := p.mappings[key.SetArtefact(key.GetArtefactType(), "")]; len(h) > 0 {
		return h
	}
	return p.lookupMedia(key.SetArtefact("", key.GetMediaType()))
}

func (p *Registry[H, K]) Register(key K, h H) {
	p.mappings[key] = append(p.mappings[key], h)
}
