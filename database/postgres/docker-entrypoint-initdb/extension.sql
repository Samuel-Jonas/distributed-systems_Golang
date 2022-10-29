CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    user_id uuid NOT NULL DEFAULT uuid_generate_v4(),
    username VARCHAR ( 50 )  UNIQUE NOT NULL,
    pwd VARCHAR ( 50 ) NOT NULL,
    email VARCHAR ( 255 ) UNIQUE NOT NULL,
    addre VARCHAR ( 255 ),
    creation_date TIMESTAMP NOT NULL,
    update_date TIMESTAMP
);

INSERT INTO users(user_id, username, pwd, email, addre, creation_date, update_date) VALUES (uuid_generate_v4(), 'samuel', 'passencrypt', 'samuel@email.com', 'address1', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);