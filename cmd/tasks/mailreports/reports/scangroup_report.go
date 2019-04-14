package reports

import "time"

// ScanGroupReportEvent holds the time stamp and the data
// of events (so urls for new websites, hosts for new host events etc)
type ScanGroupReportEvent struct {
	Timestamp time.Time
	Data      []string
}

// ScanGroupData holds the group name and a map of a slice
// of typeid -> events
type ScanGroupData struct {
	Name   string
	Events map[int32][]*ScanGroupReportEvent // key == typeID
}

func (d *ScanGroupData) Add(typeID int32, evt *ScanGroupReportEvent) {
	if d.Events == nil {
		d.Events = make(map[int32][]*ScanGroupReportEvent)
		d.Events[typeID] = make([]*ScanGroupReportEvent, 0)
	}
	d.Events[typeID] = append(d.Events[typeID], evt)
}

func (d *ScanGroupData) Get(typeID int32) []*ScanGroupReportEvent {
	return d.Events[typeID]
}

// ScanGroupReport holds each group's events in a map of group id -> data
type ScanGroupReport struct {
	GroupData map[int]*ScanGroupData
}

func NewScanGroupReport() *ScanGroupReport {
	r := &ScanGroupReport{}
	r.GroupData = make(map[int]*ScanGroupData)
	return r
}

// Add an event for the group specified by group id, if first time, assign the group name
//
func (r *ScanGroupReport) Add(groupID int, sgName string, typeID int32, evt *ScanGroupReportEvent) {
	if _, ok := r.GroupData[groupID]; !ok {
		r.GroupData[groupID] = &ScanGroupData{Name: sgName}
	}
	r.GroupData[groupID].Add(typeID, evt)
}

func (r *ScanGroupReport) Get(groupID int) *ScanGroupData {
	return r.GroupData[groupID]
}
