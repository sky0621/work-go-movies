package client

import "strings"

// Path ...
type Path struct {
	Path string
	ID   string
}

// PathSep ...
const PathSep = "/"

// NewPath ... とにかくシンプルに「/movies/{id}」の形式にのみ対応
func NewPath(p string) *Path {
	applog.debugf("NewPath", "[Before] Path: %s", p)
	var id string
	p = strings.Trim(p, PathSep) // これで前後の「/」が取り除ける
	s := strings.Split(p, PathSep)
	if len(s) > 1 {
		id = s[len(s)-1]
		p = strings.Join(s[:len(s)-1], PathSep)
	}
	applog.debugf("NewPath", "[After] Path: %s, ID: %s", p, id)
	return &Path{Path: p, ID: id}
}

// HasID ...
func (p *Path) HasID() bool {
	return len(p.ID) > 0
}
