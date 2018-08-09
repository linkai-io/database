package migration

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	uuid "github.com/gofrs/uuid"
	"github.com/jackc/pgx"
	"github.com/ory/ladon"
	"github.com/pressly/goose"
	"gopkg.linkai.io/v1/repos/am/am"
	"gopkg.linkai.io/v1/repos/am/pkg/auth/ladonauth"
	"gopkg.linkai.io/v1/repos/am/pkg/secrets"
)

var (
	crud = []string{"create", "read", "update", "delete"}
)

func init() {
	goose.AddMigration(Up00004, Down00004)
}

func Up00004(tx *sql.Tx) error {
	m, err := initPolicyManager()
	if err != nil {
		return err
	}
	policies := buildPolicies()
	for name, policy := range policies {
		if err := m.Create(policy); err != nil {
			return fmt.Errorf("%s policy failed creation: %s\n", name, err)
		}
	}
	// This code is executed when the migration is applied.
	return nil
}

func buildPolicies() map[string]ladon.Policy {
	policies := make(map[string]ladon.Policy, 0)

	policies["systemManagementPolicy"] = createSystemPolicy()

	policies["manageOrganizationPolicy"] = createPolicy(
		"Manage Customer Organization",
		[]byte("{\"key\":\"manageOrganizationPolicy\"}"),
		//subjects
		[]string{am.OwnerRole},
		//actions
		[]string{"read", "update"},
		//resources
		[]string{am.RNOrganizationManage},
	)

	policies["manageScanGroupPolicy"] = createPolicy(
		"Manage Scan Groups",
		[]byte("{\"key\":\"manageScanGroupPolicy\"}"),
		//subjects
		[]string{am.OwnerRole, am.AdminRole, am.AuditorRole},
		//actions
		crud,
		//resources
		[]string{am.RNScanGroupGroups},
	)

	policies["readScanGroupPolicy"] = createPolicy(
		"Read Only Scan Groups",
		[]byte("{\"key\":\"readScanGroupPolicy\"}"),
		//subjects
		[]string{am.EditorRole, am.ReviewerRole},
		//actions
		[]string{"read"},
		//resources
		[]string{am.RNScanGroupGroups},
	)

	policies["manageAddressesPolicy"] = createPolicy(
		"Manage Addresses",
		[]byte("{\"key\":\"manageAddressesPolicy\"}"),
		//subjects
		[]string{am.OwnerRole, am.AdminRole, am.AuditorRole},
		//actions
		crud,
		//resources
		[]string{am.RNAddressAddresses},
	)

	policies["readAddressesPolicy"] = createPolicy(
		"Read Only Addresses",
		[]byte("{\"key\":\"readAddressesPolicy\"}"),
		//subjects
		[]string{am.EditorRole, am.ReviewerRole},
		//actions
		[]string{"read"},
		//resources
		[]string{am.RNAddressAddresses},
	)

	policies["manageFindingsPolicy"] = createPolicy(
		"Manage Findings",
		[]byte("{\"key\":\"manageFindingsPolicy\"}"),
		//subjects
		[]string{am.OwnerRole, am.AdminRole, am.AuditorRole},
		//actions
		crud,
		//resources
		[]string{am.RNFindingsFindings},
	)

	policies["readFindingsPolicy"] = createPolicy(
		"Read Only Findings",
		[]byte("{\"key\":\"readFindingsPolicy\"}"),
		//subjects
		[]string{am.EditorRole, am.ReviewerRole},
		//actions
		[]string{"read"},
		//resources
		[]string{am.RNFindingsFindings},
	)

	policies["manageJobServicePolicy"] = createPolicy(
		"Manage Job Service (start, stop, pause, cancel etc)",
		[]byte("{\"key\":\"manageJobServicePolicy\"}"),
		//subjects
		[]string{am.OwnerRole, am.AdminRole, am.AuditorRole},
		//actions
		crud,
		//resources
		[]string{am.RNJobService},
	)

	// editor and reviewer can only read
	policies["readJobServicePolicy"] = createPolicy(
		"Read Only Job Service access",
		[]byte("{\"key\":\"readJobServicePolicy\"}"),
		//subjects
		[]string{am.EditorRole, am.ReviewerRole},
		//actions
		[]string{"read"},
		//resources
		[]string{am.RNJobService},
	)

	policies["manageTagServicePolicy"] = createPolicy(
		"Manage Tag Service (create tech stacks, groups, add tags etc)",
		[]byte("{\"key\":\"manageTagServicePolicy\"}"),
		//subjects
		[]string{am.OwnerRole, am.AdminRole, am.AuditorRole, am.EditorRole},
		//actions
		crud,
		[]string{am.RNTagServiceStacks, am.RNTagServiceGroups, am.RNTagServiceTagging, am.RNTagServiceCustom},
	)

	// reviewer can only read
	policies["readTagServicePolicy"] = createPolicy(
		"Read Only access to Tag Service",
		[]byte("{\"key\":\"readTagServicePolicy\"}"),
		//subjects
		[]string{am.ReviewerRole},
		//actions
		[]string{"read"},
		[]string{am.RNTagServiceStacks, am.RNTagServiceGroups, am.RNTagServiceTagging},
	)

	// reviewer can create custom tags (only they can see)
	policies["manageTagServiceCustom"] = createPolicy(
		"Manage Custom Tags access to Tag Service",
		[]byte("{\"key\":\"manageTagServiceCustom\"}"),
		//subjects
		[]string{am.ReviewerRole},
		//actions
		crud,
		[]string{am.RNTagServiceCustom},
	)

	return policies
}

// Down00004 rebuilds the policies to get the keys, and creates a map,
// extract all policies from the db and attempt to match the meta keys
// and then delete those ids that match.
func Down00004(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.

	m, err := initPolicyManager()
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

func createPolicy(description string, meta []byte, subjects, actions, resources []string) ladon.Policy {
	id, err := uuid.NewV4()
	if err != nil {
		panic(err) // should never happen really
	}

	return &ladon.DefaultPolicy{
		ID:          id.String(),
		Description: description,
		Meta:        meta,
		Subjects:    subjects,
		Actions:     actions,
		Resources:   resources,
		Effect:      ladon.AllowAccess,
		Conditions:  ladon.Conditions{},
	}
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
		Actions:     crud,
		Resources:   []string{am.RNSystem},
		Effect:      ladon.AllowAccess,
		Conditions:  ladon.Conditions{},
	}
}

func initPolicyManager() (*ladonauth.LadonPolicyManager, error) {
	dbsecrets := secrets.NewDBSecrets(os.Getenv("APP_ENV"), os.Getenv("APP_REGION"))
	dbstring, err := dbsecrets.DBString("linkai_admin")
	if err != nil {
		return nil, err
	}

	if dbstring == "" {
		return nil, errors.New("db string not set")
	}

	conf, err := pgx.ParseConnectionString(dbstring)
	if err != nil {
		return nil, err
	}
	p, err := pgx.NewConnPool(pgx.ConnPoolConfig{ConnConfig: conf})
	if err != nil {
		return nil, err
	}
	m := ladonauth.NewPolicyManager(p, "pgx")
	err = m.Init()
	return m, err
}
