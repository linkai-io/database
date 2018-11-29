package migration

import (
	"database/sql"
	"time"

	uuid "github.com/gofrs/uuid"
	"github.com/linkai-io/am/am"
	"github.com/pressly/goose"
)

const (
	systemOrgName        = "linkai-system"
	systemSupportOrgName = "linkai-support"

	deleteOrgStatement = `select am.delete_org((select organization_id from am.organizations where organization_name=$1));`

	deleteOrgRoles = `delete from am.ladon_role_member where organization_id=(select organization_id from am.organizations where organization_name=$1)`

	createOrgStatement = `with org as (
		insert into am.organizations (
			organization_name, organization_custom_id, user_pool_id, user_pool_client_id, user_pool_client_secret, identity_pool_id, 
			owner_email, first_name, last_name, phone, country, state_prefecture, street, 
			address1, address2, city, postal_code, creation_time, status_id, deleted, subscription_id
		)
		values 
			($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, false, $20)
		returning organization_id
	) 
	insert into am.users (
			organization_id, user_custom_id, email, first_name, last_name, user_status_id, creation_time, deleted
		) 
		values
			( (select org.organization_id from org), $21, $22, $23, $24, $25, $26, false) returning organization_id, user_id;`

	addRoleStatement = `INSERT INTO am.ladon_role (role_id, organization_id, role_name) values ($1,$2,$3)`

	addMemberToRoleStatement = `INSERT INTO am.ladon_role_member (organization_id, role_id, member_id) VALUES 
		($1, (select role_id from am.ladon_role where role_name=$2), $3)`
)

func init() {
	goose.AddMigration(Up00007, Down00007)
}

func Up00007(tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	if err := createSystemOrg(systemOrgName, "system", tx); err != nil {
		return err
	}

	return createSystemOrg(systemSupportOrgName, "support", tx)
}

func createSystemOrg(orgName string, name string, tx *sql.Tx) error {
	stmt, err := tx.Prepare(createOrgStatement)
	if err != nil {
		return err
	}

	coid := generateID()
	cuid := generateID()

	now := time.Now().UnixNano()

	row := stmt.QueryRow(orgName, coid, "empty", "empty", "empty", "empty", name+"@linkai.io", "linkai", name, "+81", "japan",
		name, name, "", "", name, name, now, am.OrgStatusActive, am.SubscriptionSystem,
		cuid, name+"@linkai.io", "linkai", name, am.UserStatusActive, now)

	oid := 0
	uid := 0
	if err := row.Scan(&oid, &uid); err != nil {
		return err
	}
	stmt.Close()

	// Create the system or support role
	stmt, err = tx.Prepare(addRoleStatement)
	if err != nil {
		return err
	}

	roleID := generateID()
	role := am.SystemRole
	if name == "support" {
		role = am.SystemSupportRole
	}
	if _, err := stmt.Exec(roleID, oid, role); err != nil {
		return err
	}
	stmt.Close()

	// now add the owner user to that role
	stmt, err = tx.Prepare(addMemberToRoleStatement)
	if err != nil {
		return err
	}

	if _, err := stmt.Exec(oid, role, uid); err != nil {
		return err
	}
	stmt.Close()
	return nil
}

func Down00007(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	if err := deleteSystemOrg(systemOrgName, tx); err != nil {
		return err
	}

	return deleteSystemOrg(systemSupportOrgName, tx)
}

func deleteSystemOrg(orgName string, tx *sql.Tx) error {
	// deleteOrgStatement will remove the roles as well.
	stmt, err := tx.Prepare(deleteOrgStatement)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(orgName)
	if err != nil {
		return err
	}

	return stmt.Close()
}

func generateID() string {
	id, err := uuid.NewV4()
	if err != nil {
		panic(err) // should never happen really
	}
	return id.String()
}
