package migration

import (
	"database/sql"
	"fmt"

	uuid "github.com/gofrs/uuid"
	"github.com/linkai-io/am/am"
	"github.com/linkai-io/database/pkg/migration"
	"github.com/ory/ladon"
	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(Up00004, Down00004)
}

func Up00004(tx *sql.Tx) error {
	m, err := migration.InitPolicyManager()
	if err != nil {
		return err
	}
	policies := buildPolicies()
	for name, policy := range policies {
		if err := m.Create(policy); err != nil {
			return fmt.Errorf("%s policy failed creation: %s", name, err)
		}
	}
	// This code is executed when the migration is applied.
	return nil
}

func buildPolicies() map[string]ladon.Policy {
	policies := make(map[string]ladon.Policy, 0)

	policies["systemManagementPolicy"] = createSystemPolicy()

	policies["manageOrganizationPolicy"] = migration.CreatePolicy(
		"Manage Customer Organization",
		[]byte("{\"key\":\"manageOrganizationPolicy\"}"),
		//subjects
		[]string{am.OwnerRole},
		//actions
		[]string{"read", "update"},
		//resources
		[]string{am.RNOrganizationManage},
	)

	policies["manageScanGroupPolicy"] = migration.CreatePolicy(
		"Manage Scan Groups",
		[]byte("{\"key\":\"manageScanGroupPolicy\"}"),
		//subjects
		[]string{am.OwnerRole, am.AdminRole, am.AuditorRole},
		//actions
		migration.CRUD,
		//resources
		[]string{am.RNScanGroupGroups},
	)

	policies["readScanGroupPolicy"] = migration.CreatePolicy(
		"Read Only Scan Groups",
		[]byte("{\"key\":\"readScanGroupPolicy\"}"),
		//subjects
		[]string{am.EditorRole, am.ReviewerRole},
		//actions
		migration.READ,
		//resources
		[]string{am.RNScanGroupGroups},
	)

	policies["manageAddressesPolicy"] = migration.CreatePolicy(
		"Manage Addresses",
		[]byte("{\"key\":\"manageAddressesPolicy\"}"),
		//subjects
		[]string{am.OwnerRole, am.AdminRole, am.AuditorRole},
		//actions
		migration.CRUD,
		//resources
		[]string{am.RNAddressAddresses},
	)

	policies["readAddressesPolicy"] = migration.CreatePolicy(
		"Read Only Addresses",
		[]byte("{\"key\":\"readAddressesPolicy\"}"),
		//subjects
		[]string{am.EditorRole, am.ReviewerRole},
		//actions
		migration.READ,
		//resources
		[]string{am.RNAddressAddresses},
	)

	policies["manageFindingsPolicy"] = migration.CreatePolicy(
		"Manage Findings",
		[]byte("{\"key\":\"manageFindingsPolicy\"}"),
		//subjects
		[]string{am.OwnerRole, am.AdminRole, am.AuditorRole},
		//actions
		migration.CRUD,
		//resources
		[]string{am.RNFindingsFindings},
	)

	policies["readFindingsPolicy"] = migration.CreatePolicy(
		"Read Only Findings",
		[]byte("{\"key\":\"readFindingsPolicy\"}"),
		//subjects
		[]string{am.EditorRole, am.ReviewerRole},
		//actions
		migration.READ,
		//resources
		[]string{am.RNFindingsFindings},
	)

	policies["manageEventServicePolicy"] = migration.CreatePolicy(
		"Manage Event Service (create, delete)",
		[]byte("{\"key\":\"manageEventServicePolicy\"}"),
		//subjects
		[]string{am.OwnerRole, am.AdminRole, am.AuditorRole},
		//actions
		migration.CRUD,
		//resources
		[]string{am.RNEventService},
	)

	// editor and reviewer can only read
	policies["readEventServicePolicy"] = migration.CreatePolicy(
		"Read Only Event Service access",
		[]byte("{\"key\":\"readEventServicePolicy\"}"),
		//subjects
		[]string{am.EditorRole, am.ReviewerRole},
		//actions
		migration.READ,
		//resources
		[]string{am.RNEventService},
	)

	policies["manageTagServicePolicy"] = migration.CreatePolicy(
		"Manage Tag Service (create tech stacks, groups, add tags etc)",
		[]byte("{\"key\":\"manageTagServicePolicy\"}"),
		//subjects
		[]string{am.OwnerRole, am.AdminRole, am.AuditorRole, am.EditorRole},
		//actions
		migration.CRUD,
		[]string{am.RNTagServiceStacks, am.RNTagServiceGroups, am.RNTagServiceTagging, am.RNTagServiceCustom},
	)

	// reviewer can only read
	policies["readTagServicePolicy"] = migration.CreatePolicy(
		"Read Only access to Tag Service",
		[]byte("{\"key\":\"readTagServicePolicy\"}"),
		//subjects
		[]string{am.ReviewerRole},
		//actions
		migration.READ,
		[]string{am.RNTagServiceStacks, am.RNTagServiceGroups, am.RNTagServiceTagging},
	)

	// reviewer can create custom tags (only they can see)
	policies["manageTagServiceCustom"] = migration.CreatePolicy(
		"Manage Custom Tags access to Tag Service",
		[]byte("{\"key\":\"manageTagServiceCustom\"}"),
		//subjects
		[]string{am.ReviewerRole},
		//actions
		migration.CRUD,
		[]string{am.RNTagServiceCustom},
	)

	return policies
}

// Down00004 rebuilds the policies to get the keys, and creates a map,
// extract all policies from the db and attempt to match the meta keys
// and then delete those ids that match.
func Down00004(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.

	m, err := migration.InitPolicyManager()
	if err != nil {
		return err
	}

	policies := buildPolicies()

	keys := make(map[string]struct{}, 0)
	for _, p := range policies {
		keys[string(p.GetMeta())] = struct{}{}
	}

	dbPolicies, err := m.GetAll(1000, 0)
	if err != nil {
		return err
	}

	for _, policy := range dbPolicies {
		if policy.GetMeta() == nil {
			continue
		}
		key := string(policy.GetMeta())
		if _, ok := keys[key]; ok {
			if err := m.Delete(policy.GetID()); err != nil {
				return err
			}
		}
	}
	return nil
}

// For allowing system users full access to services via the RNSystem resource.
func createSystemPolicy() ladon.Policy {
	id, err := uuid.NewV4()
	if err != nil {
		panic(err) // should never happen really
	}

	return &ladon.DefaultPolicy{
		ID:          id.String(),
		Description: "System Management Policy",
		Meta:        []byte("{\"key\":\"systemManagementPolicy\"}"),
		Subjects:    []string{am.SystemRole, am.SystemSupportRole},
		Actions:     migration.CRUD,
		Resources:   []string{am.RNSystem},
		Effect:      ladon.AllowAccess,
		Conditions:  ladon.Conditions{},
	}
}
