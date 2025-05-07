const fs = require('fs');
const path = '/data/import/generated_products.json'; 

const products = JSON.parse(fs.readFileSync(path, 'utf8'));
db = db.getSiblingDB('products_datas');
db.iron_products.drop(); 
db.iron_products.insertMany(products);