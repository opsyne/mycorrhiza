// Code generated by qtc from "view.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line tree/view.qtpl:1
package tree

//line tree/view.qtpl:1
import "sort"

//line tree/view.qtpl:2
import "path"

//line tree/view.qtpl:3
import "github.com/bouncepaw/mycorrhiza/util"

//line tree/view.qtpl:5
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line tree/view.qtpl:5
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line tree/view.qtpl:5
func StreamTreeHTML(qw422016 *qt422016.Writer) {
//line tree/view.qtpl:5
	qw422016.N().S(`
`)
//line tree/view.qtpl:6
}

//line tree/view.qtpl:6
func WriteTreeHTML(qq422016 qtio422016.Writer) {
//line tree/view.qtpl:6
	qw422016 := qt422016.AcquireWriter(qq422016)
//line tree/view.qtpl:6
	StreamTreeHTML(qw422016)
//line tree/view.qtpl:6
	qt422016.ReleaseWriter(qw422016)
//line tree/view.qtpl:6
}

//line tree/view.qtpl:6
func TreeHTML() string {
//line tree/view.qtpl:6
	qb422016 := qt422016.AcquireByteBuffer()
//line tree/view.qtpl:6
	WriteTreeHTML(qb422016)
//line tree/view.qtpl:6
	qs422016 := string(qb422016.B)
//line tree/view.qtpl:6
	qt422016.ReleaseByteBuffer(qb422016)
//line tree/view.qtpl:6
	return qs422016
//line tree/view.qtpl:6
}

// Subhyphae links are recursive. It may end up looking like that if drawn with
// pseudographics:
// ╔══════════════╗
// ║Foo           ║   The presented hyphae are foo and foo/bar
// ║╔════════════╗║
// ║║Bar         ║║
// ║╚════════════╝║
// ╚══════════════╝

//line tree/view.qtpl:16
func streamchildHTML(qw422016 *qt422016.Writer, c *child) {
//line tree/view.qtpl:16
	qw422016.N().S(`
`)
//line tree/view.qtpl:18
	sort.Slice(c.children, func(i, j int) bool {
		return c.children[i].name < c.children[j].name
	})

//line tree/view.qtpl:21
	qw422016.N().S(`
<li class="subhyphae__entry">
	<a class="subhyphae__link" href="/hypha/`)
//line tree/view.qtpl:23
	qw422016.E().S(c.name)
//line tree/view.qtpl:23
	qw422016.N().S(`">
		`)
//line tree/view.qtpl:24
	qw422016.E().S(util.BeautifulName(path.Base(c.name)))
//line tree/view.qtpl:24
	qw422016.N().S(`
	</a>
`)
//line tree/view.qtpl:26
	if len(c.children) > 0 {
//line tree/view.qtpl:26
		qw422016.N().S(`
	<ul>
	`)
//line tree/view.qtpl:28
		for _, child := range c.children {
//line tree/view.qtpl:28
			qw422016.N().S(`
		`)
//line tree/view.qtpl:29
			qw422016.N().S(childHTML(&child))
//line tree/view.qtpl:29
			qw422016.N().S(`
	`)
//line tree/view.qtpl:30
		}
//line tree/view.qtpl:30
		qw422016.N().S(`
	</ul>
`)
//line tree/view.qtpl:32
	}
//line tree/view.qtpl:32
	qw422016.N().S(`
</li>
`)
//line tree/view.qtpl:34
}

//line tree/view.qtpl:34
func writechildHTML(qq422016 qtio422016.Writer, c *child) {
//line tree/view.qtpl:34
	qw422016 := qt422016.AcquireWriter(qq422016)
//line tree/view.qtpl:34
	streamchildHTML(qw422016, c)
//line tree/view.qtpl:34
	qt422016.ReleaseWriter(qw422016)
//line tree/view.qtpl:34
}

//line tree/view.qtpl:34
func childHTML(c *child) string {
//line tree/view.qtpl:34
	qb422016 := qt422016.AcquireByteBuffer()
//line tree/view.qtpl:34
	writechildHTML(qb422016, c)
//line tree/view.qtpl:34
	qs422016 := string(qb422016.B)
//line tree/view.qtpl:34
	qt422016.ReleaseByteBuffer(qb422016)
//line tree/view.qtpl:34
	return qs422016
//line tree/view.qtpl:34
}

//line tree/view.qtpl:37
func streamsiblingHTML(qw422016 *qt422016.Writer, s *sibling) {
//line tree/view.qtpl:37
	qw422016.N().S(`
<li class="relative-hyphae__entry">
	<a class="relative-hyphae__link" href="/hypha/`)
//line tree/view.qtpl:39
	qw422016.E().S(s.name)
//line tree/view.qtpl:39
	qw422016.N().S(`">
		`)
//line tree/view.qtpl:40
	qw422016.E().S(util.BeautifulName(path.Base(s.name)))
//line tree/view.qtpl:40
	qw422016.N().S(`
		<span class="relative-hyphae__count">
		`)
//line tree/view.qtpl:42
	if s.directSubhyphaeCount > 0 {
//line tree/view.qtpl:42
		qw422016.N().S(`
			<span class="relative-hyphae__direct-count">
				`)
//line tree/view.qtpl:44
		qw422016.N().D(s.directSubhyphaeCount)
//line tree/view.qtpl:44
		qw422016.N().S(`
			</span>
		`)
//line tree/view.qtpl:46
	}
//line tree/view.qtpl:46
	qw422016.N().S(`
		`)
//line tree/view.qtpl:47
	if s.indirectSubhyphaeCount > 0 {
//line tree/view.qtpl:47
		qw422016.N().S(`
			<span class="relative-hyphae__indirect-count">
				(`)
//line tree/view.qtpl:49
		qw422016.N().D(s.indirectSubhyphaeCount)
//line tree/view.qtpl:49
		qw422016.N().S(`)
			</span>
		`)
//line tree/view.qtpl:51
	}
//line tree/view.qtpl:51
	qw422016.N().S(`
		</span>
	</a>
</li>
`)
//line tree/view.qtpl:55
}

//line tree/view.qtpl:55
func writesiblingHTML(qq422016 qtio422016.Writer, s *sibling) {
//line tree/view.qtpl:55
	qw422016 := qt422016.AcquireWriter(qq422016)
//line tree/view.qtpl:55
	streamsiblingHTML(qw422016, s)
//line tree/view.qtpl:55
	qt422016.ReleaseWriter(qw422016)
//line tree/view.qtpl:55
}

//line tree/view.qtpl:55
func siblingHTML(s *sibling) string {
//line tree/view.qtpl:55
	qb422016 := qt422016.AcquireByteBuffer()
//line tree/view.qtpl:55
	writesiblingHTML(qb422016, s)
//line tree/view.qtpl:55
	qs422016 := string(qb422016.B)
//line tree/view.qtpl:55
	qt422016.ReleaseByteBuffer(qb422016)
//line tree/view.qtpl:55
	return qs422016
//line tree/view.qtpl:55
}