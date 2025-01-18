
db = db.getSiblingDB('warehouse');

db.createUser({
    user: 'warehouse',
    pwd: 'password',
    roles: [
        {
            role: 'readWrite',
            db: 'warehouse'
        },
    ],
});

db.createCollection("storage");
db.createCollection("product_list");

db.createCollection("clients");
db.createCollection("sales");

db.createCollection("suppliers");
db.createCollection("deliveries");
