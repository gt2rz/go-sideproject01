CREATE TABLE customers(
    id VARCHAR(255) PRIMARY KEY,
    fullname VARCHAR(255) NOT NULL,
    email VARCHAR(255)  NOT NULL,
    phone VARCHAR(255)  NOT NULL
);

INSERT INTO `migro_go`.`customers` (`id`, `fullname`, `email`, `phone`) VALUES ('49adf8b0-7b59-11ed-a1eb-0242ac120002', 'Miguel Gutierrez', 'iam.gt2rz@gmail.com', '123456');
