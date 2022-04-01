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

INSERT INTO PUBLICATIONS (TITLE, CONTENT, AUTHOR_ID)
VALUES
("Publicação do Usuário 1", "Essa é a publicação do usuário 1! Viva!!!", 1),
("Publicação do Usuário 2", "Essa é a publicação do usuário 2! Viva!!!", 2),
("Publicação do Usuário 3", "Essa é a publicação do usuário 3! Viva!!!", 3),
("Publicação do Usuário 4", "Essa é a publicação do usuário 4! Viva!!!", 4),
("Publicação do Usuário 5", "Essa é a publicação do usuário 5! Viva!!!", 5);