CREATE TABLE IF NOT EXISTS molding_batches (
    id INT PRIMARY KEY AUTO_INCREMENT,
    serial_number VARCHAR(50),
    material_type VARCHAR(100),
    cooling_time INT,
    production_date TIMESTAMP
);

LOAD DATA INFILE '/var/lib/mysql-files/molding.csv'
INTO TABLE molding_batches
FIELDS TERMINATED BY ',' 
ENCLOSED BY '"'
LINES TERMINATED BY '\n'
IGNORE 1 LINES
(serial_number, material_type, cooling_time, production_date);
