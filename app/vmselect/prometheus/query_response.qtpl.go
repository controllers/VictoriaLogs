// Code generated by qtc from "query_response.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line app/vmselect/prometheus/query_response.qtpl:1
package prometheus

//line app/vmselect/prometheus/query_response.qtpl:1
import (
	"github.com/VictoriaMetrics/VictoriaMetrics/app/vmselect/netstorage"

	// QueryResponse generates response for /api/v1/query.See https://prometheus.io/docs/prometheus/latest/querying/api/#instant-queries

//line app/vmselect/prometheus/query_response.qtpl:8

	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line app/vmselect/prometheus/query_response.qtpl:8
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line app/vmselect/prometheus/query_response.qtpl:8
func StreamQueryResponse(qw422016 *qt422016.Writer, rs []netstorage.Result) {
//line app/vmselect/prometheus/query_response.qtpl:8
	qw422016.N().S(`{"status":"success","data":{"resultType":"vector","result":[`)
//line app/vmselect/prometheus/query_response.qtpl:14
	if len(rs) > 0 {
//line app/vmselect/prometheus/query_response.qtpl:14
		qw422016.N().S(`{"metric":`)
//line app/vmselect/prometheus/query_response.qtpl:16
		streammetricNameObject(qw422016, &rs[0].MetricName)
//line app/vmselect/prometheus/query_response.qtpl:16
		qw422016.N().S(`,"value":`)
//line app/vmselect/prometheus/query_response.qtpl:17
		streammetricRow(qw422016, rs[0].Timestamps[0], rs[0].Values[0])
//line app/vmselect/prometheus/query_response.qtpl:17
		qw422016.N().S(`}`)
//line app/vmselect/prometheus/query_response.qtpl:19
		rs = rs[1:]

//line app/vmselect/prometheus/query_response.qtpl:20
		for i := range rs {
//line app/vmselect/prometheus/query_response.qtpl:21
			r := &rs[i]

//line app/vmselect/prometheus/query_response.qtpl:21
			qw422016.N().S(`,{"metric":`)
//line app/vmselect/prometheus/query_response.qtpl:23
			streammetricNameObject(qw422016, &r.MetricName)
//line app/vmselect/prometheus/query_response.qtpl:23
			qw422016.N().S(`,"value":`)
//line app/vmselect/prometheus/query_response.qtpl:24
			streammetricRow(qw422016, r.Timestamps[0], r.Values[0])
//line app/vmselect/prometheus/query_response.qtpl:24
			qw422016.N().S(`}`)
//line app/vmselect/prometheus/query_response.qtpl:26
		}
//line app/vmselect/prometheus/query_response.qtpl:27
	}
//line app/vmselect/prometheus/query_response.qtpl:27
	qw422016.N().S(`]}}`)
//line app/vmselect/prometheus/query_response.qtpl:31
}

//line app/vmselect/prometheus/query_response.qtpl:31
func WriteQueryResponse(qq422016 qtio422016.Writer, rs []netstorage.Result) {
//line app/vmselect/prometheus/query_response.qtpl:31
	qw422016 := qt422016.AcquireWriter(qq422016)
//line app/vmselect/prometheus/query_response.qtpl:31
	StreamQueryResponse(qw422016, rs)
//line app/vmselect/prometheus/query_response.qtpl:31
	qt422016.ReleaseWriter(qw422016)
//line app/vmselect/prometheus/query_response.qtpl:31
}

//line app/vmselect/prometheus/query_response.qtpl:31
func QueryResponse(rs []netstorage.Result) string {
//line app/vmselect/prometheus/query_response.qtpl:31
	qb422016 := qt422016.AcquireByteBuffer()
//line app/vmselect/prometheus/query_response.qtpl:31
	WriteQueryResponse(qb422016, rs)
//line app/vmselect/prometheus/query_response.qtpl:31
	qs422016 := string(qb422016.B)
//line app/vmselect/prometheus/query_response.qtpl:31
	qt422016.ReleaseByteBuffer(qb422016)
//line app/vmselect/prometheus/query_response.qtpl:31
	return qs422016
//line app/vmselect/prometheus/query_response.qtpl:31
}
