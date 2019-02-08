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
	goose.AddMigration(Up00023, Down00023)
}

func Up00023(tx *sql.Tx) error {
	m, err := migration.InitPolicyManager()
	if err != nil {
		return err
	}
	policies := buildBigDataPolicies()
	for name, policy := range policies {
		if err := m.Create(policy); err != nil {
			return fmt.Errorf("%s policy failed creation: %s", name, err)
		}
	}
	// This code is executed when the migration is applied.
	return nil
}

func buildBigDataPolicies() map[string]ladon.Policy {
	policies := make(map[string]ladon.Policy, 0)

	policies["readWriteBigData"] = migration.CreatePolicy(
		"Allow users to read/write BigData from local db",
		[]byte("{\"key\":\"readWriteBigData\"}"),
		//subjects
		[]string{am.OwnerRole, am.AdminRole, am.ReviewerRole, am.AuditorRole, am.EditorRole},
		//actions
		migration.CREATE_READ,
		//resources
		[]string{am.RNBigData},
	)

	return policies
}

// Down00023 rebuilds the policies to get the keys, and creates a map,
// extract all policies from the db and attempt to match the meta keys
// and then delete those ids that match.
func Down00023(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.

	m, err := migration.InitPolicyManager()
	if err != nil {
		return err
	}

	policies := buildBigDataPolicies()

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
