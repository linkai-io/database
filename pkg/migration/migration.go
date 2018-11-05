package migration

import (
	"errors"
	"os"

	"github.com/gofrs/uuid"
	"github.com/jackc/pgx"
	"github.com/linkai-io/am/pkg/auth/ladonauth"
	"github.com/linkai-io/am/pkg/secrets"
	"github.com/ory/ladon"
)

var (
	CRUD = []string{"create", "read", "update", "delete"}
	READ = []string{"read"}
)

// GetServicePasswords returns the passwords of the supplied users via the provided DBSecrets
func GetServicePasswords(dbsecrets *secrets.SecretsCache, users []string) (map[string]string, error) {
	userMap := make(map[string]string, 0)
	for _, user := range users {
		password, err := dbsecrets.ServicePassword(user)
		if password == "" {
			return nil, errors.New("empty password for user: " + user)
		}
		userMap[user] = password
		if err != nil {
			return nil, err
		}
	}
	return userMap, nil
}

// InitPolicyManager for managing user policies
func InitPolicyManager() (*ladonauth.LadonPolicyManager, error) {
	dbsecrets := secrets.NewSecretsCache(os.Getenv("APP_ENV"), os.Getenv("APP_REGION"))
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

// CreatePolicy for a ladon.Policy
func CreatePolicy(description string, meta []byte, subjects, actions, resources []string) ladon.Policy {
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
