package migration

import (
	"database/sql"
	"fmt"

	"github.com/linkai-io/am/am"
	"github.com/linkai-io/database/pkg/migration"
	"github.com/ory/ladon"
	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(Up00019, Down00019)
}

func Up00019(tx *sql.Tx) error {
	m, err := migration.InitPolicyManager()
	if err != nil {
		return err
	}
	policies := buildUserOrgPolicies()
	for name, policy := range policies {
		if err := m.Create(policy); err != nil {
			return fmt.Errorf("%s policy failed creation: %s", name, err)
		}
	}
	// This code is executed when the migration is applied.
	return nil
}

func buildUserOrgPolicies() map[string]ladon.Policy {
	policies := make(map[string]ladon.Policy, 0)

	policies["manageUsersPolicy"] = migration.CreatePolicy(
		"Allow Owner/Admin to Manage Organization's Users",
		[]byte("{\"key\":\"manageUsersPolicy\"}"),
		//subjects
		[]string{am.OwnerRole, am.AdminRole},
		//actions
		migration.CRUD,
		//resources
		[]string{am.RNUserManage},
	)

	policies["manageUserSelfPolicy"] = migration.CreatePolicy(
		"Allow User to Update their user information",
		[]byte("{\"key\":\"manageUserSelfPolicy\"}"),
		//subjects
		[]string{am.OwnerRole, am.AdminRole, am.ReviewerRole, am.AuditorRole, am.EditorRole},
		//actions
		[]string{"update"},
		//resources
		[]string{am.RNUserSelf},
	)

	return policies
}

// Down00019 rebuilds the policies to get the keys, and creates a map,
// extract all policies from the db and attempt to match the meta keys
// and then delete those ids that match.
func Down00019(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.

	m, err := migration.InitPolicyManager()
	if err != nil {
		return err
	}

	policies := buildUserOrgPolicies()

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
