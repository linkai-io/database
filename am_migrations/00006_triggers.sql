-- +goose Up
-- +goose StatementBegin
-- SQL in this section is executed when the migration is applied.
CREATE OR REPLACE FUNCTION update_owner_user() RETURNS TRIGGER AS $body$
BEGIN 
    IF NEW.owner_email <> OLD.owner_email OR NEW.first_name <> OLD.first_name OR NEW.last_name <> OLD.last_name THEN
        UPDATE am.users set email=NEW.owner_email, first_name=NEW.first_name, last_name=NEW.last_name where organization_id=OLD.organization_id and email=OLD.owner_email;
    END IF;
    RETURN NEW;
END;
$body$ language plpgsql;
-- +goose StatementEnd

CREATE TRIGGER owner_user AFTER UPDATE ON am.organizations FOR EACH ROW execute procedure update_owner_user();

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TRIGGER owner_user on am.organizations;
DROP FUNCTION update_owner_user;