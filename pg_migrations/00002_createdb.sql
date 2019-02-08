-- +goose Up 
-- +goose NO TRANSACTION
-- CREATE DATABASE linkai WITH OWNER = linkai_admin TEMPLATE = template1 ENCODING = 'UTF8' LC_COLLATE = 'en_US.UTF-8' LC_CTYPE = 'en_US.UTF-8' CONNECTION LIMIT = -1;
CREATE DATABASE linkai WITH OWNER = linkai_admin TEMPLATE = template1 ENCODING = 'UTF8' CONNECTION LIMIT = -1;

-- +goose Down 
-- +goose NO TRANSACTION
-- SQL in this section is executed when the migration is rolled back.
DROP DATABASE linkai;
