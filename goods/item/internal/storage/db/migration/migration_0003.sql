create table if not exists "item"
(
    id uuid
    constraint item_pk
    primary key,
    created_at          timestamp(6) NOT NULL,
    updated_at          timestamp(6) NOT NULL,
    category_id uuid not null ,
    brand varchar(255) not null,
    "name" varchar(255) not null,
    description varchar(255) not null ,
    image_link varchar (255) not null ,
    price decimal not null,

    constraint fk_category foreign key (category_id) references category(id) on delete cascade
    );

-- category - 47deec22-ac0d-430d-a9b8-3ac810db49f4
INSERT INTO "item" (id, created_at, updated_at, category_id, brand, name, description, image_link, price)
SELECT '16da1598-3181-4cbc-89ad-d564d4d368eb', NOW(), NOW() ,'47deec22-ac0d-430d-a9b8-3ac810db49f4',
       'Yiota',  'Y-1002', 'lorem', 'https://www.gophergala.com/assets/img/fancy_gopher_renee.jpg', 200000
WHERE NOT EXISTS (SELECT 1 FROM "item" WHERE id='16da1598-3181-4cbc-89ad-d564d4d368eb');

INSERT INTO "item" (id, created_at, updated_at, category_id, brand, name, description, image_link, price)
SELECT 'c88bfe16-d0de-4bfc-b867-0173e6075e73', NOW(), NOW() ,'47deec22-ac0d-430d-a9b8-3ac810db49f4',
       'Zota',  'Z-1002', 'lorem', 'https://www.gophergala.com/assets/img/fancy_gopher_renee.jpg', 20001
WHERE NOT EXISTS (SELECT 1 FROM "item" WHERE id='c88bfe16-d0de-4bfc-b867-0173e6075e73');


INSERT INTO "item" (id, created_at, updated_at, category_id, brand, name, description, image_link, price)
SELECT '54a9d607-b05b-4f3e-bf04-7f252ab1c387', NOW(), NOW() ,'47deec22-ac0d-430d-a9b8-3ac810db49f4',
       'Oma',  'O-1002', 'lorem', 'https://www.gophergala.com/assets/img/fancy_gopher_renee.jpg', 30001
WHERE NOT EXISTS (SELECT 1 FROM "item" WHERE id='54a9d607-b05b-4f3e-bf04-7f252ab1c387');

-- category - 16da1598-3181-4cbc-89ad-d564d4d368eb
INSERT INTO "item" (id, created_at, updated_at, category_id, brand, name, description, image_link, price)
SELECT '85a42c7e-4fa0-4b73-a33a-41b3342b5e7b', NOW(), NOW() ,'16da1598-3181-4cbc-89ad-d564d4d368eb',
       'Z-phone',  'Z-002', 'lorem', 'https://www.gophergala.com/assets/img/fancy_gopher_renee.jpg', 35000
WHERE NOT EXISTS (SELECT 1 FROM "item" WHERE id='85a42c7e-4fa0-4b73-a33a-41b3342b5e7b');

INSERT INTO "item" (id, created_at, updated_at, category_id, brand, name, description, image_link, price)
SELECT 'bd95688e-1e73-42f7-aa1a-2edd4f10125f', NOW(), NOW() ,'16da1598-3181-4cbc-89ad-d564d4d368eb',
       'Z-phone',  'Z-005', 'lorem', 'https://www.gophergala.com/assets/img/fancy_gopher_renee.jpg', 45000
WHERE NOT EXISTS (SELECT 1 FROM "item" WHERE id='bd95688e-1e73-42f7-aa1a-2edd4f10125f');

INSERT INTO "item" (id, created_at, updated_at, category_id, brand, name, description, image_link, price)
SELECT 'cf15831d-63de-4c82-ba78-79d5a960cc01', NOW(), NOW() ,'16da1598-3181-4cbc-89ad-d564d4d368eb',
       'Z-phone',  'Z-009', 'lorem', 'https://www.gophergala.com/assets/img/fancy_gopher_renee.jpg', 55000
WHERE NOT EXISTS (SELECT 1 FROM "item" WHERE id='cf15831d-63de-4c82-ba78-79d5a960cc01');
