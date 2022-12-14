// Code generated by qtc from "install.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line install.qtpl:1
package templates

//line install.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line install.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line install.qtpl:2
type InstallPage struct {
	BasePage
	AppUrl string
}

//line install.qtpl:8
func (p *InstallPage) StreamTitle(qw422016 *qt422016.Writer) {
//line install.qtpl:8
	qw422016.N().S(`
	Install page
`)
//line install.qtpl:10
}

//line install.qtpl:10
func (p *InstallPage) WriteTitle(qq422016 qtio422016.Writer) {
//line install.qtpl:10
	qw422016 := qt422016.AcquireWriter(qq422016)
//line install.qtpl:10
	p.StreamTitle(qw422016)
//line install.qtpl:10
	qt422016.ReleaseWriter(qw422016)
//line install.qtpl:10
}

//line install.qtpl:10
func (p *InstallPage) Title() string {
//line install.qtpl:10
	qb422016 := qt422016.AcquireByteBuffer()
//line install.qtpl:10
	p.WriteTitle(qb422016)
//line install.qtpl:10
	qs422016 := string(qb422016.B)
//line install.qtpl:10
	qt422016.ReleaseByteBuffer(qb422016)
//line install.qtpl:10
	return qs422016
//line install.qtpl:10
}

//line install.qtpl:12
func (p *InstallPage) StreamBody(qw422016 *qt422016.Writer) {
//line install.qtpl:12
	qw422016.N().S(`
    <div>
        Successfully connected to apaleo! Your app is installed and available at:
        <a href="`)
//line install.qtpl:15
	qw422016.E().S(p.AppUrl)
//line install.qtpl:15
	qw422016.N().S(`">`)
//line install.qtpl:15
	qw422016.E().S(p.AppUrl)
//line install.qtpl:15
	qw422016.N().S(`</a>
    </div>
`)
//line install.qtpl:17
}

//line install.qtpl:17
func (p *InstallPage) WriteBody(qq422016 qtio422016.Writer) {
//line install.qtpl:17
	qw422016 := qt422016.AcquireWriter(qq422016)
//line install.qtpl:17
	p.StreamBody(qw422016)
//line install.qtpl:17
	qt422016.ReleaseWriter(qw422016)
//line install.qtpl:17
}

//line install.qtpl:17
func (p *InstallPage) Body() string {
//line install.qtpl:17
	qb422016 := qt422016.AcquireByteBuffer()
//line install.qtpl:17
	p.WriteBody(qb422016)
//line install.qtpl:17
	qs422016 := string(qb422016.B)
//line install.qtpl:17
	qt422016.ReleaseByteBuffer(qb422016)
//line install.qtpl:17
	return qs422016
//line install.qtpl:17
}
