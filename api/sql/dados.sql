insert into usuariosdevbook(nome,nick,email,senha)
values
("Usuario1","usuario_1","usuario1@gmail.com","$2a$10$KPByE5DDwJqzvGBk3zI9qu9j67cl2uBXaOlVP4c4t8PqntpqjQcvG"),
("Usuario2","usuario_2","usuario2@gmail.com","$2a$10$KPByE5DDwJqzvGBk3zI9qu9j67cl2uBXaOlVP4c4t8PqntpqjQcvG"),
("Usuario3","usuario_3","usuario3@gmail.com","$2a$10$KPByE5DDwJqzvGBk3zI9qu9j67cl2uBXaOlVP4c4t8PqntpqjQcvG")

insert into seguidoresdevbook(usuario_id,seguidor_id)
values
(1,2),
(3,1),
(1,3)