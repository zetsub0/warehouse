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

db.createCollection("customers");
db.createCollection("sales");

db.createCollection("suppliers");
db.createCollection("deliveries");

customersCl = "customers"

db.customersCl.createIndex(
    {name: "text"},
    {default_language: "russian", language_override: "language"}
)