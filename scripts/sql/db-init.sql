DROP TABLE IF EXISTS todos;
CREATE TABLE todos (
  id UUID NOT NULL,
  title STRING,
  description STRING,
  status INT,
  created_at TIMESTAMP
);

INSERT INTO todos (id, title, description, status, created_at)
  VALUES (gen_random_uuid(), 'first hoge', 'iidesune', 0, now());
INSERT INTO todos (id, title, description, status, created_at)
  VALUES (gen_random_uuid(), 'second hoge', 'iidesune', 0, now());
INSERT INTO todos (id, title, description, status, created_at)
  VALUES (gen_random_uuid(), 'third hoge', 'iidesune', 0, now());
INSERT INTO todos (id, title, description, status, created_at)
  VALUES (gen_random_uuid(), 'fourth hoge', 'iidesune', 0, now());
INSERT INTO todos (id, title, description, status, created_at)
  VALUES (gen_random_uuid(), 'fifth hoge', 'iidesune', 0, now());

DROP TABLE IF EXISTS labels;
CREATE TABLE labels (
  id UUID NOT NULL,
  name STRING,
  created_at TIMESTAMP
);
