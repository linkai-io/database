package reports

import (
	"testing"
	"time"
)

func TestQuery(t *testing.T) {
	query, args, err := buildGetReportQuery(1, 1, false, time.Now().Add(-24*time.Hour))
	t.Logf("%s %#v %s\n", query, args, err)

}
