// Code generated by qtc from "home.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line home.qtpl:1
package templates

//line home.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line home.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line home.qtpl:2
type HomePage struct {
	BasePage
}

//line home.qtpl:7
func (p *HomePage) StreamTitle(qw422016 *qt422016.Writer) {
//line home.qtpl:7
	qw422016.N().S(`
	Home page
`)
//line home.qtpl:9
}

//line home.qtpl:9
func (p *HomePage) WriteTitle(qq422016 qtio422016.Writer) {
//line home.qtpl:9
	qw422016 := qt422016.AcquireWriter(qq422016)
//line home.qtpl:9
	p.StreamTitle(qw422016)
//line home.qtpl:9
	qt422016.ReleaseWriter(qw422016)
//line home.qtpl:9
}

//line home.qtpl:9
func (p *HomePage) Title() string {
//line home.qtpl:9
	qb422016 := qt422016.AcquireByteBuffer()
//line home.qtpl:9
	p.WriteTitle(qb422016)
//line home.qtpl:9
	qs422016 := string(qb422016.B)
//line home.qtpl:9
	qt422016.ReleaseByteBuffer(qb422016)
//line home.qtpl:9
	return qs422016
//line home.qtpl:9
}

//line home.qtpl:11
func (p *HomePage) StreamBody(qw422016 *qt422016.Writer) {
//line home.qtpl:11
	qw422016.N().S(`
	<h1>Home page</h1>
`)
//line home.qtpl:13
}

//line home.qtpl:13
func (p *HomePage) WriteBody(qq422016 qtio422016.Writer) {
//line home.qtpl:13
	qw422016 := qt422016.AcquireWriter(qq422016)
//line home.qtpl:13
	p.StreamBody(qw422016)
//line home.qtpl:13
	qt422016.ReleaseWriter(qw422016)
//line home.qtpl:13
}

//line home.qtpl:13
func (p *HomePage) Body() string {
//line home.qtpl:13
	qb422016 := qt422016.AcquireByteBuffer()
//line home.qtpl:13
	p.WriteBody(qb422016)
//line home.qtpl:13
	qs422016 := string(qb422016.B)
//line home.qtpl:13
	qt422016.ReleaseByteBuffer(qb422016)
//line home.qtpl:13
	return qs422016
//line home.qtpl:13
}