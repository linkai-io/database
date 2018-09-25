package am

import (
	"context"
)

const (
	RNScanGroupGroups = "lrn:service:scangroup:feature:groups"
)

type GroupStatus int

var (
	GroupStarted GroupStatus = 1
	GroupStopped GroupStatus = 2
)

var GroupStatusMap = map[GroupStatus]string{
	1: "started",
	2: "stopped",
}

type ScanGroupEvent struct {
	EventID          int64  `json:"event_id"`
	OrgID            int    `json:"org_id"`
	GroupID          int64  `json:"group_id"`
	EventUserID      int    `json:"event_user_id"`
	EventTime        int64  `json:"event_time"`
	EventDescription string `json:"event_description"`
	EventFrom        string `json:"event_from"`
}

// ScanGroup is a grouping configuration that has owner related information
type ScanGroup struct {
	OrgID                int                  `json:"org_id"`
	GroupID              int                  `json:"group_id"`
	GroupName            string               `json:"group_name"`
	CreationTime         int64                `json:"creation_time"`
	CreatedBy            int                  `json:"created_by"`
	ModifiedBy           int                  `json:"modified_by"`
	ModifiedTime         int64                `json:"modified_time"`
	OriginalInputS3URL   string               `json:"original_input_s3_url"`
	ModuleConfigurations *ModuleConfiguration `json:"module_configurations"`
	Paused               bool                 `json:"paused"`
	Deleted              bool                 `json:"deleted"`
}

// ScanGroupService manages input lists and configurations for an organization and group. OrgIDs should
// always be returned for ensuring data integrity for requesters
type ScanGroupService interface {
	Init(config []byte) error
	IsAuthorized(ctx context.Context, userContext UserContext, resource, action string) bool
	Get(ctx context.Context, userContext UserContext, groupID int) (oid int, group *ScanGroup, err error)
	GetByName(ctx context.Context, userContext UserContext, groupName string) (oid int, group *ScanGroup, err error)
	Groups(ctx context.Context, userContext UserContext) (oid int, groups []*ScanGroup, err error)
	Create(ctx context.Context, userContext UserContext, newGroup *ScanGroup) (oid int, gid int, err error)
	Update(ctx context.Context, userContext UserContext, group *ScanGroup) (oid int, gid int, err error)
	Delete(ctx context.Context, userContext UserContext, groupID int) (oid int, gid int, err error)
	Pause(ctx context.Context, userContext UserContext, groupID int) (oid int, gid int, err error)
	Resume(ctx context.Context, userContext UserContext, groupID int) (oid int, gid int, err error)
}