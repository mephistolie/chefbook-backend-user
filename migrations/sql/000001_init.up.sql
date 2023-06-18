CREATE TABLE users
(
    user_id     uuid PRIMARY KEY NOT NULL,
    first_name  VARCHAR(64)  DEFAULT NULL,
    last_name   VARCHAR(64)  DEFAULT NULL,
    description VARCHAR(150) DEFAULT NULL,
    avatar_id   uuid         DEFAULT NULL
);

CREATE TABLE avatar_uploads
(
    avatar_id uuid PRIMARY KEY                                  NOT NULL DEFAULT gen_random_uuid(),
    user_id   uuid REFERENCES users (user_id) ON DELETE CASCADE NOT NULL UNIQUE
);

CREATE TABLE inbox
(
    message_id uuid PRIMARY KEY         NOT NULL UNIQUE,
    timestamp  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now():: timestamp
);
