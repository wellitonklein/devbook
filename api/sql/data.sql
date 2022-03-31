USE devbook;

INSERT INTO USERS (NAME, NICK, EMAIL, PASSWORD)
VALUES
("Usuário 1", "user_1", "usuario1@gmail.com", "$2a$10$lCSU.hwSGXGWUzuUVNUQEukYhQl2X.LltJeFd6mIJZAnm2.QsXg0y"),
("Usuário 2", "user_2", "usuario2@gmail.com", "$2a$10$lCSU.hwSGXGWUzuUVNUQEukYhQl2X.LltJeFd6mIJZAnm2.QsXg0y"),
("Usuário 3", "user_3", "usuario3@gmail.com", "$2a$10$lCSU.hwSGXGWUzuUVNUQEukYhQl2X.LltJeFd6mIJZAnm2.QsXg0y"),
("Usuário 4", "user_4", "usuario4@gmail.com", "$2a$10$lCSU.hwSGXGWUzuUVNUQEukYhQl2X.LltJeFd6mIJZAnm2.QsXg0y"),
("Usuário 5", "user_5", "usuario5@gmail.com", "$2a$10$lCSU.hwSGXGWUzuUVNUQEukYhQl2X.LltJeFd6mIJZAnm2.QsXg0y");

INSERT INTO FOLLOWERS (USER_ID, FOLLOWER_ID)
VALUES
(1, 2),
(2, 3),
(3, 4),
(5, 1);