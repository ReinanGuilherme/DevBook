DROP TABLE IF EXISTS usuariosdevbook;
DROP TABLE IF EXISTS seguidoresdevbook;

CREATE TABLE usuariosdevbook (
    id int auto_increment primary key,
    nome varchar(50) not null unique,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    senha varchar(100) not null,
    criadoEm timestamp default current_timestamp()
) ENGINE=INNODB

CREATE TABLE seguidoresdevbook(
    usuario_id int not null,
    FOREIGN KEY(usuario_id)
    REFERENCES usuariosdevbook(id)
    ON DELETE CASCADE,
    seguidor_id int not null,
    FOREIGN KEY (seguidor_id)
    REFERENCES usuariosdevbook(id)
    ON DELETE CASCADE,

    primary KEY(usuario_id, seguidor_id)
)ENGINE=INNODB

CREATE TABLE publicacoesdevbook(
    id int auto_increment primary key,
    titulo varchar(50) not null,
    conteudo varchar(300) not null,

    autor_id int not null,
    FOREIGN KEY (autor_id)
    REFERENCES usuariosdevbook(id)
    ON DELETE CASCADE,

    curtidas int default 0,
    criadaEm timestamp default current_timestamp
)ENGINE=INNODB