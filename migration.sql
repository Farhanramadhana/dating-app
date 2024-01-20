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
	image_url VARCHAR(255) NOT NULL,
	created_at TIMESTAMPTZ DEFAULT current_timestamp,
	updated_at TIMESTAMPTZ DEFAULT current_timestamp,
	FOREIGN KEY (user_id) REFERENCES "users"(id)
);

create table swipe_matches (
	id SERIAL PRIMARY KEY,
	first_user_id INT not null,
	second_user_id INT not null,
	is_first_user_like bool null,
	is_second_user_like bool null,
	created_at TIMESTAMPTZ DEFAULT current_timestamp,
	updated_at TIMESTAMPTZ DEFAULT current_timestamp
);


-- Insert dummy data into "users" table
INSERT INTO "users" (id, email, first_name, last_name)
VALUES
    (1001, 'user1@example.com', 'John', 'Doe'),
    (1002, 'user2@example.com', 'Jane', 'Smith'),
    (1003, 'user3@example.com', 'Alice', 'Johnson'),
    (1004, 'user4@example.com', 'Bob', 'Williams'),
    (1005, 'user5@example.com', 'Eva', 'Brown'),
    (1006, 'user6@example.com', 'David', 'Miller'),
    (1007, 'user7@example.com', 'Sara', 'Clark'),
    (1008, 'user8@example.com', 'Michael', 'Davis'),
    (1009, 'user9@example.com', 'Emily', 'Jones'),
    (1010, 'user10@example.com', 'Alex', 'Taylor'),
    (1011, 'user11@example.com', 'Saja', 'Tia');

-- Insert dummy data into "user_profiles" table
INSERT INTO "user_profiles" (user_id, gender, birthdate, gender_preference, is_premium_user)
VALUES
    (1001, 'Male', '1990-05-15', 'Female', true),
    (1002, 'Female', '1985-08-22', 'Male', false),
    (1003, 'Female', '1992-12-10', 'Male', true),
    (1004, 'Male', '1988-04-03', 'Female', false),
    (1005, 'Female', '1995-06-28', 'Male', true),
    (1006, 'Male', '1983-11-20', 'Female', false),
    (1007, 'Female', '1998-03-18', 'Male', true),
    (1008, 'Male', '1991-07-05', 'Female', false),
    (1009, 'Female', '1987-09-12', 'Male', true),
    (1010, 'Male', '1994-01-30', 'Female', false);
	(1011, 'Male', '1994-01-30', 'Female', true);