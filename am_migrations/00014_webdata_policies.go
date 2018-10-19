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
	goose.AddMigration(Up00014, Down00014)
}

func Up00014(tx *sql.Tx) error {
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

func buildUp00014Policies() map[string]ladon.Policy {
	policies := make(map[string]ladon.Policy, 0)

	policies["manageWebDataResponsesPolicy"] = migration.CreatePolicy(
		"Manage Web Data Responses",
		[]byte("{\"key\":\"manageWebDataResponsesPolicy\"}"),
		//subjects
		[]string{am.OwnerRole, am.AdminRole, am.AuditorRole},
		//actions
		migration.CRUD,
		//resources
		[]string{am.RNWebDataResponses},
	)

	policies["readWebDataResponsesPolicy"] = migration.CreatePolicy(
		"Read Only Web Data Certificates",
		[]byte("{\"key\":\"readAddressesPolicy\"}"),
		//subjects
		[]string{am.EditorRole, am.ReviewerRole},
		//actions
		migration.READ,
		//resources
		[]string{am.RNWebDataResponses},
	)

	policies["manageWebDataCertificatesPolicy"] = migration.CreatePolicy(
		"Manage Web Data Certificates",
		[]byte("{\"key\":\"manageWebDataCertificatesPolicy\"}"),
		//subjects
		[]string{am.OwnerRole, am.AdminRole, am.AuditorRole},
		//actions
		migration.CRUD,
		//resources
		[]string{am.RNWebDataCertificates},
	)

	policies["readWebDataCertificatesPolicy"] = migration.CreatePolicy(
		"Read Only Web Data Certificates",
		[]byte("{\"key\":\"readWebDataCertificatesPolicy\"}"),
		//subjects
		[]string{am.EditorRole, am.ReviewerRole},
		//actions
		migration.READ,
		//resources
		[]string{am.RNWebDataCertificates},
	)

	policies["manageWebDataSnapshotsPolicy"] = migration.CreatePolicy(
		"Manage Web Data Snapshots",
		[]byte("{\"key\":\"manageWebDataSnapshotsPolicy\"}"),
		//subjects
		[]string{am.OwnerRole, am.AdminRole, am.AuditorRole},
		//actions
		migration.CRUD,
		//resources
		[]string{am.RNWebDataSnapshots},
	)

	policies["readWebDataSnapshotsPolicy"] = migration.CreatePolicy(
		"Read Only Web Data Snapshots",
		[]byte("{\"key\":\"readWebDataSnapshotsPolicy\"}"),
		//subjects
		[]string{am.EditorRole, am.ReviewerRole},
		//actions
		migration.READ,
		//resources
		[]string{am.RNWebDataSnapshots},
	)

	return policies
}

// Down00014 rebuilds the policies to get the keys, and creates a map,
// extract all policies from the db and attempt to match the meta keys
// and then delete those ids that match.
func Down00014(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.

	m, err := migration.InitPolicyManager()
	if err != nil {
		return err
	}

	policies := buildUp00014Policies()

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
