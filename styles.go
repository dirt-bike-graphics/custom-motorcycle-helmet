// Package helmetstyles provides motorcycle helmet shell silhouettes, paintable zones
// and surface finish metadata across 33 helmet styles and 5 rider categories.
//
// Maintained from the open shell catalog powering [Custom Motorcycle Helmet] and its
// interactive [3D motorcycle helmet designer].
//
// [Custom Motorcycle Helmet]: https://custommotorcyclehelmet.com
// [3D motorcycle helmet designer]: https://custommotorcyclehelmet.com/design?utm_source=pkg_go_dev&utm_medium=referral&utm_campaign=go
package helmetstyles

import (
	_ "embed"
	"encoding/json"
	"strings"
	"sync"
)

// RawJSON holds the embedded styles.json catalog bytes.
//
//go:embed styles.json
var RawJSON []byte

// HelmetStyle describes a single motorcycle helmet shell silhouette.
type HelmetStyle struct {
	ID         string `json:"id"`
	Category   string `json:"category"`
	Model      string `json:"model"`
	HasPeak    bool   `json:"hasPeak"`
	HasChinBar bool   `json:"hasChinBar"`
}

type catalogFile struct {
	Updated  string        `json:"updated"`
	Finishes []string      `json:"finishes"`
	Zones    []string      `json:"zones"`
	Styles   []HelmetStyle `json:"styles"`
}

var (
	once    sync.Once
	catalog catalogFile
)

func load() {
	once.Do(func() {
		_ = json.Unmarshal(RawJSON, &catalog)
	})
}

func norm(s string) string {
	var b strings.Builder
	for _, r := range s {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') {
			b.WriteRune(r)
		}
	}
	return strings.ToLower(b.String())
}

// Styles returns all 33 motorcycle helmet shell silhouettes.
func Styles() []HelmetStyle {
	load()
	out := make([]HelmetStyle, len(catalog.Styles))
	copy(out, catalog.Styles)
	return out
}

// Finishes returns the supported surface clearcoat finishes.
func Finishes() []string {
	load()
	out := make([]string, len(catalog.Finishes))
	copy(out, catalog.Finishes)
	return out
}

// FindStyle resolves a helmet silhouette by ID or model name (case and separator insensitive).
func FindStyle(query string) *HelmetStyle {
	load()
	q := norm(query)
	if q == "" {
		return nil
	}
	for i := range catalog.Styles {
		s := &catalog.Styles[i]
		if norm(s.ID) == q || norm(s.Model) == q {
			copyStyle := *s
			return &copyStyle
		}
	}
	for i := range catalog.Styles {
		s := &catalog.Styles[i]
		if strings.Contains(norm(s.Model), q) || strings.Contains(norm(s.ID), q) {
			copyStyle := *s
			return &copyStyle
		}
	}
	return nil
}

// ByCategory returns all helmet silhouettes in the given rider category.
func ByCategory(category string) []HelmetStyle {
	load()
	c := norm(category)
	var out []HelmetStyle
	for _, s := range catalog.Styles {
		if norm(s.Category) == c {
			out = append(out, s)
		}
	}
	return out
}

// PaintableZones returns the paintable shell zones for the matched helmet style.
func PaintableZones(query string) []string {
	s := FindStyle(query)
	if s == nil {
		return nil
	}
	var out []string
	for _, z := range catalog.Zones {
		if !s.HasChinBar && z == "chin" {
			continue
		}
		out = append(out, z)
	}
	return out
}
