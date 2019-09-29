DROP TABLE IF EXISTS todos;
CREATE TABLE todos (
  id UUID NOT NULL,
  title STRING,
  description STRING,
  status INT,
  created_at STRING
);
