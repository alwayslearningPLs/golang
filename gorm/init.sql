SELECT "warehouses"."id",
  "warehouses"."created_at",
  "warehouses"."updated_at",
  "warehouses"."deleted_at",
  "warehouses"."name",
  "warehouses"."location",
  "warehouses"."food_id",
  "Food"."id" AS "Food__id",
  "Food"."created_at" AS "Food__created_at",
  "Food"."updated_at" AS "Food__updated_at",
  "Food"."deleted_at" AS "Food__deleted_at",
  "Food"."name" AS "Food__name",
  "Food"."value" AS "Food__value"
FROM "warehouses" LEFT JOIN "foods" "Food" ON "warehouses"."food_id" = "Food"."id" AND "Food"."deleted_at" IS NULL
WHERE "Food"."id"=2 AND "warehouses"."deleted_at" IS NULL;

-- one-on-one relationship example
INSERT INTO companies(name)
VALUES ('viewnext'), ('elecnor'), ('imatia'), ('primeit');

INSERT INTO users(name, company_id)
VALUES ('Ivan', 4), ('Juan', 1), ('Noe', 2), ('Pedro', 3);

-- one-has-many relationship example
INSERT INTO libraries(name)
VALUES ('La nobel'), ('La no nobel');

INSERT INTO books(name, library_id)
VALUES ('The name of the wind', 1),
  ('The mister', 2),
  ('IPV6 security', 1),
  ('Introduccion a la filosofia', 2),
  ('El gato con botas', 1);
