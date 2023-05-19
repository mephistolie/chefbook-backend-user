CREATE TABLE users
(
    user_id     uuid PRIMARY KEY NOT NULL UNIQUE DEFAULT gen_random_uuid(),
    first_name  VARCHAR(64)                      DEFAULT NULL,
    last_name   VARCHAR(64)                      DEFAULT NULL,
    description VARCHAR(150)                     DEFAULT NULL,
    avatar      TEXT                             DEFAULT NULL
);

CREATE TABLE inbox
(
    message_id uuid PRIMARY KEY         NOT NULL UNIQUE,
    timestamp  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now():: timestamp
);
