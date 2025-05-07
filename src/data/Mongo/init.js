const fs = require('fs');
const path = '/data/import/generated_products.json'; // Chemin du fichier JSON

// Lire le fichier JSON et l'insérer dans MongoDB
const products = JSON.parse(fs.readFileSync(path, 'utf8')); // Utilisation de fs.readFileSync pour lire le fichier
db = db.getSiblingDB('products_datas');
db.iron_products.drop();  // Si tu veux supprimer la collection avant de la réinitialiser
db.iron_products.insertMany(products);  // Insérer les produits dans la collection
