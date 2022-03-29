// Code generated by qtc from "readers.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/readers.qtpl:1
package views

//line views/readers.qtpl:1
import "net/http"

//line views/readers.qtpl:2
import "strings"

//line views/readers.qtpl:3
import "path"

//line views/readers.qtpl:4
import "os"

//line views/readers.qtpl:6
import "github.com/bouncepaw/mycorrhiza/cfg"

//line views/readers.qtpl:7
import "github.com/bouncepaw/mycorrhiza/hyphae"

//line views/readers.qtpl:8
import "github.com/bouncepaw/mycorrhiza/l18n"

//line views/readers.qtpl:9
import "github.com/bouncepaw/mycorrhiza/mimetype"

//line views/readers.qtpl:10
import "github.com/bouncepaw/mycorrhiza/tree"

//line views/readers.qtpl:11
import "github.com/bouncepaw/mycorrhiza/user"

//line views/readers.qtpl:12
import "github.com/bouncepaw/mycorrhiza/util"

//line views/readers.qtpl:13
import "github.com/bouncepaw/mycorrhiza/viewutil"

//line views/readers.qtpl:15
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/readers.qtpl:15
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/readers.qtpl:15
func StreamMediaMenu(qw422016 *qt422016.Writer, rq *http.Request, h hyphae.Hypha, u *user.User) {
//line views/readers.qtpl:15
	qw422016.N().S(`
`)
//line views/readers.qtpl:17
	lc := l18n.FromRequest(rq)

//line views/readers.qtpl:18
	qw422016.N().S(`
<div class="layout">
<main class="main-width media-tab">
	<h1>`)
//line views/readers.qtpl:21
	qw422016.N().S(lc.Get("ui.media_title", &l18n.Replacements{"name": beautifulLink(h.CanonicalName())}))
//line views/readers.qtpl:21
	qw422016.N().S(`</h1>
	`)
//line views/readers.qtpl:22
	switch h.(type) {
//line views/readers.qtpl:23
	case *hyphae.MediaHypha:
//line views/readers.qtpl:23
		qw422016.N().S(`
		<p class="explanation">`)
//line views/readers.qtpl:24
		qw422016.E().S(lc.Get("ui.media_tip"))
//line views/readers.qtpl:24
		qw422016.N().S(` <a href="/help/en/media" class="shy-link">`)
//line views/readers.qtpl:24
		qw422016.E().S(lc.Get("ui.media_what_is"))
//line views/readers.qtpl:24
		qw422016.N().S(`</a></p>
	`)
//line views/readers.qtpl:25
	default:
//line views/readers.qtpl:25
		qw422016.N().S(`
		<p class="explanation">`)
//line views/readers.qtpl:26
		qw422016.E().S(lc.Get("ui.media_empty"))
//line views/readers.qtpl:26
		qw422016.N().S(` <a href="/help/en/media" class="shy-link">`)
//line views/readers.qtpl:26
		qw422016.E().S(lc.Get("ui.media_what_is"))
//line views/readers.qtpl:26
		qw422016.N().S(`</a></p>
	`)
//line views/readers.qtpl:27
	}
//line views/readers.qtpl:27
	qw422016.N().S(`

	<section class="amnt-grid">
	`)
//line views/readers.qtpl:30
	switch h := h.(type) {
//line views/readers.qtpl:31
	case *hyphae.MediaHypha:
//line views/readers.qtpl:31
		qw422016.N().S(`
		`)
//line views/readers.qtpl:33
		mime := mimetype.FromExtension(path.Ext(h.MediaFilePath()))
		fileinfo, err := os.Stat(h.MediaFilePath())

//line views/readers.qtpl:34
		qw422016.N().S(`
		`)
//line views/readers.qtpl:35
		if err == nil {
//line views/readers.qtpl:35
			qw422016.N().S(`
		<fieldset class="amnt-menu-block">
			<legend class="modal__title modal__title_small">`)
//line views/readers.qtpl:37
			qw422016.E().S(lc.Get("ui.media_stat"))
//line views/readers.qtpl:37
			qw422016.N().S(`</legend>
			<p class="modal__confirmation-msg"><b>`)
//line views/readers.qtpl:38
			qw422016.E().S(lc.Get("ui.media_stat_size"))
//line views/readers.qtpl:38
			qw422016.N().S(`</b> `)
//line views/readers.qtpl:38
			qw422016.E().S(lc.GetPlural64("ui.media_size_value", fileinfo.Size()))
//line views/readers.qtpl:38
			qw422016.N().S(`</p>
			<p><b>`)
//line views/readers.qtpl:39
			qw422016.E().S(lc.Get("ui.media_stat_mime"))
//line views/readers.qtpl:39
			qw422016.N().S(`</b> `)
//line views/readers.qtpl:39
			qw422016.E().S(mime)
//line views/readers.qtpl:39
			qw422016.N().S(`</p>
		</fieldset>
		`)
//line views/readers.qtpl:41
		}
//line views/readers.qtpl:41
		qw422016.N().S(`

		`)
//line views/readers.qtpl:43
		if strings.HasPrefix(mime, "image/") {
//line views/readers.qtpl:43
			qw422016.N().S(`
		<fieldset class="amnt-menu-block">
			<legend class="modal__title modal__title_small">`)
//line views/readers.qtpl:45
			qw422016.E().S(lc.Get("ui.media_include"))
//line views/readers.qtpl:45
			qw422016.N().S(`</legend>
			<p class="modal__confirmation-msg">`)
//line views/readers.qtpl:46
			qw422016.E().S(lc.Get("ui.media_include_tip"))
//line views/readers.qtpl:46
			qw422016.N().S(`</p>
			<pre class="codeblock"><code>img { `)
//line views/readers.qtpl:47
			qw422016.E().S(h.CanonicalName())
//line views/readers.qtpl:47
			qw422016.N().S(` }</code></pre>
		</fieldset>
		`)
//line views/readers.qtpl:49
		}
//line views/readers.qtpl:49
		qw422016.N().S(`
	`)
//line views/readers.qtpl:50
	}
//line views/readers.qtpl:50
	qw422016.N().S(`

	`)
//line views/readers.qtpl:52
	if u.CanProceed("upload-binary") {
//line views/readers.qtpl:52
		qw422016.N().S(`
	<form action="/upload-binary/`)
//line views/readers.qtpl:53
		qw422016.E().S(h.CanonicalName())
//line views/readers.qtpl:53
		qw422016.N().S(`"
			method="post" enctype="multipart/form-data"
			class="upload-binary modal amnt-menu-block">
		<fieldset class="modal__fieldset">
			<legend class="modal__title modal__title_small">`)
//line views/readers.qtpl:57
		qw422016.E().S(lc.Get("ui.media_new"))
//line views/readers.qtpl:57
		qw422016.N().S(`</legend>
			<p class="modal__confirmation-msg">`)
//line views/readers.qtpl:58
		qw422016.E().S(lc.Get("ui.media_new_tip"))
//line views/readers.qtpl:58
		qw422016.N().S(`</p>
			<label for="upload-binary__input"></label>
			<input type="file" id="upload-binary__input" name="binary">

			<button type="submit" class="btn stick-to-bottom" value="Upload">`)
//line views/readers.qtpl:62
		qw422016.E().S(lc.Get("ui.media_upload"))
//line views/readers.qtpl:62
		qw422016.N().S(`</button>
		</fieldset>
	</form>
	`)
//line views/readers.qtpl:65
	}
//line views/readers.qtpl:65
	qw422016.N().S(`


	`)
//line views/readers.qtpl:68
	switch h := h.(type) {
//line views/readers.qtpl:69
	case *hyphae.MediaHypha:
//line views/readers.qtpl:69
		qw422016.N().S(`
		`)
//line views/readers.qtpl:70
		if u.CanProceed("remove-media") {
//line views/readers.qtpl:70
			qw422016.N().S(`
		<form action="/remove-media/`)
//line views/readers.qtpl:71
			qw422016.E().S(h.CanonicalName())
//line views/readers.qtpl:71
			qw422016.N().S(`" method="post" class="modal amnt-menu-block" method="POST">
			<fieldset class="modal__fieldset">
				<legend class="modal__title modal__title_small">`)
//line views/readers.qtpl:73
			qw422016.E().S(lc.Get("ui.media_remove"))
//line views/readers.qtpl:73
			qw422016.N().S(`</legend>
				<p class="modal__confirmation-msg">`)
//line views/readers.qtpl:74
			qw422016.E().S(lc.Get("ui.media_remove_tip"))
//line views/readers.qtpl:74
			qw422016.N().S(`</p>
				<button type="submit" class="btn" value="Remove media">`)
//line views/readers.qtpl:75
			qw422016.E().S(lc.Get("ui.media_remove_button"))
//line views/readers.qtpl:75
			qw422016.N().S(`</button>
			</fieldset>
		</form>
		`)
//line views/readers.qtpl:78
		}
//line views/readers.qtpl:78
		qw422016.N().S(`
	`)
//line views/readers.qtpl:79
	}
//line views/readers.qtpl:79
	qw422016.N().S(`

	</section>
</main>
</div>
`)
//line views/readers.qtpl:84
}

//line views/readers.qtpl:84
func WriteMediaMenu(qq422016 qtio422016.Writer, rq *http.Request, h hyphae.Hypha, u *user.User) {
//line views/readers.qtpl:84
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/readers.qtpl:84
	StreamMediaMenu(qw422016, rq, h, u)
//line views/readers.qtpl:84
	qt422016.ReleaseWriter(qw422016)
//line views/readers.qtpl:84
}

//line views/readers.qtpl:84
func MediaMenu(rq *http.Request, h hyphae.Hypha, u *user.User) string {
//line views/readers.qtpl:84
	qb422016 := qt422016.AcquireByteBuffer()
//line views/readers.qtpl:84
	WriteMediaMenu(qb422016, rq, h, u)
//line views/readers.qtpl:84
	qs422016 := string(qb422016.B)
//line views/readers.qtpl:84
	qt422016.ReleaseByteBuffer(qb422016)
//line views/readers.qtpl:84
	return qs422016
//line views/readers.qtpl:84
}

// If `contents` == "", a helpful message is shown instead.
//
// If you rename .prevnext, change the docs too.

//line views/readers.qtpl:89
func StreamHypha(qw422016 *qt422016.Writer, meta viewutil.Meta, h hyphae.Hypha, contents string) {
//line views/readers.qtpl:89
	qw422016.N().S(`
`)
//line views/readers.qtpl:91
	siblings, subhyphae, prevHyphaName, nextHyphaName := tree.Tree(h.CanonicalName())
	lc := meta.Lc

//line views/readers.qtpl:93
	qw422016.N().S(`
<div class="layout">
<main class="main-width">
	<section id="hypha">
		`)
//line views/readers.qtpl:97
	if meta.U.CanProceed("edit") {
//line views/readers.qtpl:97
		qw422016.N().S(`
		<div class="btn btn_navititle">
			<a class="btn__link_navititle" href="/edit/`)
//line views/readers.qtpl:99
		qw422016.E().S(h.CanonicalName())
//line views/readers.qtpl:99
		qw422016.N().S(`">`)
//line views/readers.qtpl:99
		qw422016.E().S(lc.Get("ui.edit_link"))
//line views/readers.qtpl:99
		qw422016.N().S(`</a>
		</div>
		`)
//line views/readers.qtpl:101
	}
//line views/readers.qtpl:101
	qw422016.N().S(`

		`)
//line views/readers.qtpl:103
	if cfg.UseAuth && util.IsProfileName(h.CanonicalName()) && meta.U.Name == strings.TrimPrefix(h.CanonicalName(), cfg.UserHypha+"/") {
//line views/readers.qtpl:103
		qw422016.N().S(`
		<div class="btn btn_navititle">
			<a class="btn__link_navititle" href="/logout">`)
//line views/readers.qtpl:105
		qw422016.E().S(lc.Get("ui.logout_link"))
//line views/readers.qtpl:105
		qw422016.N().S(`</a>
		</div>
		`)
//line views/readers.qtpl:107
		if meta.U.Group == "admin" {
//line views/readers.qtpl:107
			qw422016.N().S(`
		<div class="btn btn_navititle">
			<a class="btn__link_navititle" href="/admin">`)
//line views/readers.qtpl:109
			qw422016.E().S(lc.Get("ui.admin_panel"))
//line views/readers.qtpl:109
			qw422016.N().S(`<a>
		</div>
		`)
//line views/readers.qtpl:111
		}
//line views/readers.qtpl:111
		qw422016.N().S(`
		`)
//line views/readers.qtpl:112
	}
//line views/readers.qtpl:112
	qw422016.N().S(`

		`)
//line views/readers.qtpl:114
	qw422016.N().S(NaviTitle(h))
//line views/readers.qtpl:114
	qw422016.N().S(`
		`)
//line views/readers.qtpl:115
	switch h.(type) {
//line views/readers.qtpl:116
	case *hyphae.EmptyHypha:
//line views/readers.qtpl:116
		qw422016.N().S(`
				`)
//line views/readers.qtpl:117
		streamnonExistentHyphaNotice(qw422016, h, meta.U, meta.Lc)
//line views/readers.qtpl:117
		qw422016.N().S(`
			`)
//line views/readers.qtpl:118
	default:
//line views/readers.qtpl:118
		qw422016.N().S(`
				`)
//line views/readers.qtpl:119
		qw422016.N().S(contents)
//line views/readers.qtpl:119
		qw422016.N().S(`
		`)
//line views/readers.qtpl:120
	}
//line views/readers.qtpl:120
	qw422016.N().S(`
	</section>
	<section class="prevnext">
		`)
//line views/readers.qtpl:123
	if prevHyphaName != "" {
//line views/readers.qtpl:123
		qw422016.N().S(`
		<a class="prevnext__el prevnext__prev" href="/hypha/`)
//line views/readers.qtpl:124
		qw422016.E().S(prevHyphaName)
//line views/readers.qtpl:124
		qw422016.N().S(`" rel="prev">← `)
//line views/readers.qtpl:124
		qw422016.E().S(util.BeautifulName(path.Base(prevHyphaName)))
//line views/readers.qtpl:124
		qw422016.N().S(`</a>
		`)
//line views/readers.qtpl:125
	}
//line views/readers.qtpl:125
	qw422016.N().S(`
		`)
//line views/readers.qtpl:126
	if nextHyphaName != "" {
//line views/readers.qtpl:126
		qw422016.N().S(`
		<a class="prevnext__el prevnext__next" href="/hypha/`)
//line views/readers.qtpl:127
		qw422016.E().S(nextHyphaName)
//line views/readers.qtpl:127
		qw422016.N().S(`" rel="next">`)
//line views/readers.qtpl:127
		qw422016.E().S(util.BeautifulName(path.Base(nextHyphaName)))
//line views/readers.qtpl:127
		qw422016.N().S(` →</a>
		`)
//line views/readers.qtpl:128
	}
//line views/readers.qtpl:128
	qw422016.N().S(`
	</section>
`)
//line views/readers.qtpl:130
	StreamSubhyphae(qw422016, subhyphae, meta.Lc)
//line views/readers.qtpl:130
	qw422016.N().S(`
	<section id="hypha-bottom">
   		`)
//line views/readers.qtpl:132
	streamhyphaInfo(qw422016, meta, h)
//line views/readers.qtpl:132
	qw422016.N().S(`
	</section>
</main>
`)
//line views/readers.qtpl:135
	qw422016.N().S(categoryCard(meta, h.CanonicalName()))
//line views/readers.qtpl:135
	qw422016.N().S(`
`)
//line views/readers.qtpl:136
	streamsiblingHyphae(qw422016, siblings, meta.Lc)
//line views/readers.qtpl:136
	qw422016.N().S(`
</div>
`)
//line views/readers.qtpl:138
	streamviewScripts(qw422016)
//line views/readers.qtpl:138
	qw422016.N().S(`
`)
//line views/readers.qtpl:139
}

//line views/readers.qtpl:139
func WriteHypha(qq422016 qtio422016.Writer, meta viewutil.Meta, h hyphae.Hypha, contents string) {
//line views/readers.qtpl:139
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/readers.qtpl:139
	StreamHypha(qw422016, meta, h, contents)
//line views/readers.qtpl:139
	qt422016.ReleaseWriter(qw422016)
//line views/readers.qtpl:139
}

//line views/readers.qtpl:139
func Hypha(meta viewutil.Meta, h hyphae.Hypha, contents string) string {
//line views/readers.qtpl:139
	qb422016 := qt422016.AcquireByteBuffer()
//line views/readers.qtpl:139
	WriteHypha(qb422016, meta, h, contents)
//line views/readers.qtpl:139
	qs422016 := string(qb422016.B)
//line views/readers.qtpl:139
	qt422016.ReleaseByteBuffer(qb422016)
//line views/readers.qtpl:139
	return qs422016
//line views/readers.qtpl:139
}

//line views/readers.qtpl:141
func StreamRevision(qw422016 *qt422016.Writer, rq *http.Request, lc *l18n.Localizer, h hyphae.Hypha, contents, revHash string) {
//line views/readers.qtpl:141
	qw422016.N().S(`
<div class="layout">
<main class="main-width">
	<section>
		<p>`)
//line views/readers.qtpl:145
	qw422016.E().S(lc.Get("ui.revision_warning"))
//line views/readers.qtpl:145
	qw422016.N().S(` <a href="/rev-text/`)
//line views/readers.qtpl:145
	qw422016.E().S(revHash)
//line views/readers.qtpl:145
	qw422016.N().S(`/`)
//line views/readers.qtpl:145
	qw422016.E().S(h.CanonicalName())
//line views/readers.qtpl:145
	qw422016.N().S(`">`)
//line views/readers.qtpl:145
	qw422016.E().S(lc.Get("ui.revision_link"))
//line views/readers.qtpl:145
	qw422016.N().S(`</a></p>
		`)
//line views/readers.qtpl:146
	qw422016.N().S(NaviTitle(h))
//line views/readers.qtpl:146
	qw422016.N().S(`
		`)
//line views/readers.qtpl:147
	qw422016.N().S(contents)
//line views/readers.qtpl:147
	qw422016.N().S(`
	</section>
</main>
</div>
`)
//line views/readers.qtpl:151
	streamviewScripts(qw422016)
//line views/readers.qtpl:151
	qw422016.N().S(`
`)
//line views/readers.qtpl:152
}

//line views/readers.qtpl:152
func WriteRevision(qq422016 qtio422016.Writer, rq *http.Request, lc *l18n.Localizer, h hyphae.Hypha, contents, revHash string) {
//line views/readers.qtpl:152
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/readers.qtpl:152
	StreamRevision(qw422016, rq, lc, h, contents, revHash)
//line views/readers.qtpl:152
	qt422016.ReleaseWriter(qw422016)
//line views/readers.qtpl:152
}

//line views/readers.qtpl:152
func Revision(rq *http.Request, lc *l18n.Localizer, h hyphae.Hypha, contents, revHash string) string {
//line views/readers.qtpl:152
	qb422016 := qt422016.AcquireByteBuffer()
//line views/readers.qtpl:152
	WriteRevision(qb422016, rq, lc, h, contents, revHash)
//line views/readers.qtpl:152
	qs422016 := string(qb422016.B)
//line views/readers.qtpl:152
	qt422016.ReleaseByteBuffer(qb422016)
//line views/readers.qtpl:152
	return qs422016
//line views/readers.qtpl:152
}

//line views/readers.qtpl:154
func streamviewScripts(qw422016 *qt422016.Writer) {
//line views/readers.qtpl:154
	qw422016.N().S(`
`)
//line views/readers.qtpl:155
	for _, scriptPath := range cfg.ViewScripts {
//line views/readers.qtpl:155
		qw422016.N().S(`
<script src="`)
//line views/readers.qtpl:156
		qw422016.E().S(scriptPath)
//line views/readers.qtpl:156
		qw422016.N().S(`"></script>
`)
//line views/readers.qtpl:157
	}
//line views/readers.qtpl:157
	qw422016.N().S(`
`)
//line views/readers.qtpl:158
}

//line views/readers.qtpl:158
func writeviewScripts(qq422016 qtio422016.Writer) {
//line views/readers.qtpl:158
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/readers.qtpl:158
	streamviewScripts(qw422016)
//line views/readers.qtpl:158
	qt422016.ReleaseWriter(qw422016)
//line views/readers.qtpl:158
}

//line views/readers.qtpl:158
func viewScripts() string {
//line views/readers.qtpl:158
	qb422016 := qt422016.AcquireByteBuffer()
//line views/readers.qtpl:158
	writeviewScripts(qb422016)
//line views/readers.qtpl:158
	qs422016 := string(qb422016.B)
//line views/readers.qtpl:158
	qt422016.ReleaseByteBuffer(qb422016)
//line views/readers.qtpl:158
	return qs422016
//line views/readers.qtpl:158
}
