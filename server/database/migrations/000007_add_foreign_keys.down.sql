ALTER TABLE users DROP CONSTRAINT users_role_id_fkey;
ALTER TABLE cola_types DROP CONSTRAINT cola_types_manufacturer_id_fkey;
ALTER TABLE cola_types DROP CONSTRAINT cola_types_package_id_fkey;
ALTER TABLE cola_results DROP CONSTRAINT cola_results_cola_id_fkey;
ALTER TABLE cola_results DROP CONSTRAINT cola_results_user_id_fkey;
