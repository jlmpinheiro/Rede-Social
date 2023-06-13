/*Cria a database [devbook] se ela não existir*/
CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

/*Deleta a tabelas se já existirem*/
DROP TABLE IF EXISTS publicacoes;
DROP TABLE IF EXISTS seguidores;
DROP TABLE IF EXISTS usuarios;


/*Cria a tabela [usuarios]*/
CREATE TABLE usuarios(
    id int auto_increment primary key,
    nome varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    senha varchar(100) not null,
    criadoEm timestamp default current_timestamp()
) ENGINE=INNODB;

/*Cria a tabela [seguidores]*/
CREATE TABLE seguidores(
    usuario_id int not null,        /*cria uma coluna usuario_id do tipo int*/
        FOREIGN KEY (usuario_id)    /*define como uma chave estrangeira, não cria se não existir o id em outra tabela*/
        REFERENCES usuarios(id)     /*tabela de referencia estrangeira*/
        ON DELETE CASCADE,          /*define que se a referência principal, neste caso a tabela usuario, for deletada esta também será...*/

    seguidor_id int not null,
        FOREIGN KEY(seguidor_id)
        REFERENCES usuarios(id)
        ON DELETE CASCADE,

    PRIMARY KEY (usuario_id, seguidor_id)
) ENGINE=INNODB;

/*Cria a tabela [publicações]*/
CREATE TABLE publicacoes(
    id int auto_increment primary key,
    titulo varchar(50) not null,
    conteudo varchar(300) not null,
    autor_Id int not null,
        FOREIGN KEY(autor_id)
        REFERENCES usuarios(id)
        ON DELETE CASCADE,
    curtida int default 0,
    criadoEm timestamp default current_timestamp()
) ENGINE=INNODB;