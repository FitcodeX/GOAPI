-- CREATE DATABASE ToDOList;

USE ToDOList;


CREATE TABLE Users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    Username  VARCHAR(100) NOT NULL,
    password VARCHAR(100) NOT NULL, 
    email VARCHAR(100) NOT NULL
);

CREATE TABLE Tasks (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    status VARCHAR(50),
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL
);


INSERT INTO Users (Username, password, email) VALUES
('alice', 'password123', 'alice@example.com'),
('bob', 'password456', 'bob@example.com'),
('charlie', 'password789', 'charlie@example.com');

INSERT INTO Tasks (user_id, title, description, status, created_at, updated_at) VALUES
(1, 'Task 1', 'Description for task 1', 'Pending', '2024-05-21 10:00:00', '2024-05-21 10:00:00'),
(1, 'Task 2', 'Description for task 2', 'Completed', '2024-05-21 11:00:00', '2024-05-21 12:00:00'),
(2, 'Task 3', 'Description for task 3', 'In Progress', '2024-05-21 09:00:00', '2024-05-21 09:30:00'),
(3, 'Task 4', 'Description for task 4', 'Pending', '2024-05-21 08:00:00', '2024-05-21 08:00:00'),
(3, 'Task 5', 'Description for task 5', 'Completed', '2024-05-21 07:00:00', '2024-05-21 07:30:00');

SELECT * FROM Tasks;