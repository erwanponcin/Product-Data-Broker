CREATE TABLE IF NOT EXISTS iron_products (
  id INT PRIMARY KEY AUTO_INCREMENT,
  code VARCHAR(50),
  name VARCHAR(100),
  category VARCHAR(100),
  price_per_tonne DECIMAL(10,2),
  stock_tonne DECIMAL(10,2)
);

LOAD DATA INFILE '/var/lib/mysql-files/generated_products.csv'
INTO TABLE iron_products
FIELDS TERMINATED BY ',' 
ENCLOSED BY '"'
LINES TERMINATED BY '\n'
IGNORE 1 LINES
(id,code,name,category,price_per_tonne,stock_tonne);
