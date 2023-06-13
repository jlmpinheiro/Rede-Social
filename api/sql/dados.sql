/*popula tabelas*/

INSERT INTO usuarios (nome, nick, email, senha)
values
("Usuario 1", "u1", "u1@gmail.com","$2a$10$Ty61v2iUzwwlkYZyHrqWg.96wNXsuoyi/rgs4zJlnLWVR5D6tzZDW"),
("Usuario 2", "u2", "u2@gmail.com","$2a$10$Ty61v2iUzwwlkYZyHrqWg.96wNXsuoyi/rgs4zJlnLWVR5D6tzZDW"),
("Usuario 3", "u3", "u3@gmail.com","$2a$10$Ty61v2iUzwwlkYZyHrqWg.96wNXsuoyi/rgs4zJlnLWVR5D6tzZDW");

INSERT INTO seguidores(usuario_id, seguidor_id)
values
(1,2),
(2,3),
(3,1);

INSERT INTO publicacoes(titulo,conteudo,autor_id)
values
("Publicação do usuário 1","conteúdo gravado para usuário 1",1),
("Publicação do usuário 2","conteúdo gravado para usuário 2",2),
("Publicação do usuário 3","conteúdo gravado para usuário 3",3);