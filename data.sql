CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(100) NOT NULL CHECK (username <> '') UNIQUE,
    email VARCHAR(100) NOT NULL CHECK (email <> '') UNIQUE,
    password VARCHAR(120)
);

INSERT INTO users (username, email, password) VALUES
('budi_santoso', 'budi.santoso@gmail.com', '11111'),
('siti_aminah', 'siti.aminah@gmail.com', '22222'),
('dewi_sari', 'dewi.sari@gmail.com', '33333'),
('agus_pramono', 'agus.pramono@gmail.com', '44444'),
('rina_kusuma', 'rina.kusuma@gmail.com', '55555');

SELECT * FROM users