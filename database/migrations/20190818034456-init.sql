
-- +migrate Up
CREATE TABLE TEST (
  id int ,
  last_name varchar (25),
  first_name varchar (25),
  address varchar (25)
);
-- +migrate Down
DROP TABLE TEST;