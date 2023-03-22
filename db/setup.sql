USE zopstore;

CREATE TABLE brands(
                       id INT,
                       name VARCHAR(50) UNIQUE NOT NULL ,
                       PRIMARY KEY (id)
);

INSERT INTO brands VALUES (4,'Nike');

INSERT INTO brands VALUES (5,'Titan');
INSERT INTO brands VALUES (6,'Bru');

INSERT INTO brands VALUES (1,'Maggi');
INSERT INTO brands VALUES (2,'dairy');

CREATE TABLE products(
                         id INT,
                         name VARCHAR(50) NOT NULL ,
                         description VARCHAR(500) NOT NULL ,
                         price INT NOT NULL ,
                         quantity INT NOT NULL ,
                         category varchar(30) NOT NULL ,
                         brand_id INT NOT NULL ,
                         status ENUM('Available','Out of Stock','Discontinued'),
                         PRIMARY KEY (id),
                         FOREIGN KEY (brand_id) REFERENCES brands(id)
);

INSERT INTO products VALUES (3,'sneaker shoes','stylish',1000,3,'shoes',4,'Available');

INSERT INTO products VALUES (4,'Rolex','useful',50000,1,'wristwatch',5,'Discontinued');

INSERT INTO products VALUES (5,'Bru','tasty',100,3,'coffee',6,'Available');
