package cmp

import (
	"fmt"
	"strings"

	"github.com/google/go-cmp/cmp"
)

// DiffReporter is a simple custom reporter that only records differences
// detected during comparison.
type DiffReporter struct {
	path  cmp.Path
	diffs []string
}

// PushStep ...
func (r *DiffReporter) PushStep(ps cmp.PathStep) {
	r.path = append(r.path, ps)
}

// Report ...
func (r *DiffReporter) Report(rs cmp.Result) {
	if !rs.Equal() {
		vx, vy := r.path.Last().Values()
		r.diffs = append(r.diffs, fmt.Sprintf("%#v:\n\t-: %+v\n\t+: %+v\n", r.path, vx, vy))
	}
}

// PopStep ...
func (r *DiffReporter) PopStep() {
	r.path = r.path[:len(r.path)-1]
}

// String ...
func (r *DiffReporter) String() string {
	return strings.Join(r.diffs, "\n")
}

// Diff ...
func Diff(x, y interface{}) string {
	var r DiffReporter
	cmp.Equal(x, y, cmp.Reporter(&r))
	return r.String()
}
