CREATE EXTENSION IF NOT EXISTS pgcrypto;

-- Insert Admin
INSERT INTO users (role, first_name, last_name, email, password)
VALUES ('Admin', 'Hafiz', 'Arrahman', 'fizrahman47@gmail.com', crypt('password123', gen_salt('bf')));

INSERT INTO userverified (user_id, email, verified)
VALUES (
           (SELECT user_id FROM users WHERE email = 'fizrahman47@gmail.com'),
           'fizrahman47@gmail.com',
           true
       );

-- Insert Member
INSERT INTO users (role, first_name, last_name, email, password)
VALUES ('Member', 'Hafiz', 'Arrahman', 'hfizrrhman@gmail.com', crypt('password123', gen_salt('bf')));

INSERT INTO userverified (user_id, email, verified)
VALUES (
           (SELECT user_id FROM users WHERE email = 'hfizrrhman@gmail.com'),
           'hfizrrhman@gmail.com',
           true
       );