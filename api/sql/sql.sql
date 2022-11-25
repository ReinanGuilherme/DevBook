CREATE TABLE usuariosdevbook (
    id int auto_increment primary key,
    nome varchar(50) not null unique,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    senha varchar(20) not null,
    criadoEm timestamp default current_timestamp()
) ENGINE=INNODB