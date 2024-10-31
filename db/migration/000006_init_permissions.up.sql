-- Dashboards
INSERT INTO permission_group (id, name) VALUES (1, 'Dashboards');
INSERT INTO permission (id, code, name, permission_group_id) VALUES (1, 'view_dashboard_analytics', 'View Dashboard Analytics', 1);

-- Users
INSERT INTO permission_group (id, name) VALUES (2, 'Users');
INSERT INTO permission (id, code, name, permission_group_id) VALUES (2, 'view_user', 'View User', 2);
INSERT INTO permission (id, code, name, permission_group_id) VALUES (3, 'edit_user', 'Edit User', 2);

-- Roles
INSERT INTO permission_group (id, name) VALUES (3, 'Roles');
INSERT INTO permission (id, code, name, permission_group_id) VALUES (4, 'view_role', 'View Role', 3);
INSERT INTO permission (id, code, name, permission_group_id) VALUES (5, 'edit_role', 'Edit Role', 3);

-- Blogs
INSERT INTO permission_group (id, name) VALUES (4, 'Blogs');
INSERT INTO permission (id, code, name, permission_group_id) VALUES (6, 'view_blog', 'View Blog', 4);
INSERT INTO permission (id, code, name, permission_group_id) VALUES (7, 'edit_blog', 'Edit Blog', 4);

-- Tags
INSERT INTO permission_group (id, name) VALUES (5, 'Tags');
INSERT INTO permission (id, code, name, permission_group_id) VALUES (8, 'view_tag', 'View Tag', 5);
INSERT INTO permission (id, code, name, permission_group_id) VALUES (9, 'edit_tag', 'Edit Tag', 5);
