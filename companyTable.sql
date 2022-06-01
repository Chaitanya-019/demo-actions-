DROP TABLE IF EXISTS company;
CREATE TABLE company(
    id VARCHAR(3) NOT NULL,
    name  VARCHAR(500) NOT NULL,
    creator  VARCHAR(500) NOT NULL
);

INSERT INTO company VALUES 
("1","apple","steve jobs"),
("2","zip2","elon musk");