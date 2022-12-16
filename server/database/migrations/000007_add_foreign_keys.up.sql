ALTER TABLE users ADD FOREIGN KEY (role_id) REFERENCES user_roles (id);
ALTER TABLE cola_types ADD FOREIGN KEY (manufacturer_id) REFERENCES manufacturers (id);
ALTER TABLE cola_types ADD FOREIGN KEY (package_id) REFERENCES packages (id);
ALTER TABLE cola_results ADD FOREIGN KEY (user_id) REFERENCES users (id);
ALTER TABLE cola_results ADD FOREIGN KEY (cola_id) REFERENCES cola_types (id);
