-- User Table
CREATE TABLE "users" (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    created_at TIMESTAMPTZ DEFAULT current_timestamp,
    updated_at TIMESTAMPTZ DEFAULT current_timestamp
);

-- Password Table
CREATE TABLE user_passwords (
    id SERIAL PRIMARY KEY,
    user_id INT UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT current_timestamp,
    updated_at TIMESTAMPTZ DEFAULT current_timestamp,
    FOREIGN KEY (user_id) REFERENCES "users"(id)
);

-- ExternalLogin Table
CREATE TABLE external_login (
    id SERIAL PRIMARY KEY,
    user_id INT UNIQUE,
    login_type VARCHAR(50) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT current_timestamp,
    updated_at TIMESTAMPTZ DEFAULT current_timestamp,
    FOREIGN KEY (user_id) REFERENCES "users"(id)
);

create table user_profiles (
	id SERIAL PRIMARY KEY,
	user_id INT UNIQUE,
	gender VARCHAR(10) NOT NULL,
	birthdate date,
	gender_preference VARCHAR(10),
	is_premium_user bool default false,
	created_at TIMESTAMPTZ DEFAULT current_timestamp,
	updated_at TIMESTAMPTZ DEFAULT current_timestamp,
	FOREIGN KEY (user_id) REFERENCES "users"(id)
);

create table user_images (
	id SERIAL PRIMARY KEY,
	user_id INT,
	image_url VARCHAR(10) NOT NULL,
	created_at TIMESTAMPTZ DEFAULT current_timestamp,
	updated_at TIMESTAMPTZ DEFAULT current_timestamp,
	FOREIGN KEY (user_id) REFERENCES "users"(id)
)
