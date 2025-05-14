const fs = require('fs');
const path = '/data/import/melting.json';

const meltingData = JSON.parse(fs.readFileSync(path, 'utf8'));
db = db.getSiblingDB('products_datas');
db.melting.drop();
db.melting.insertMany(meltingData);
