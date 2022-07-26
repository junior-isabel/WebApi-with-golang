CREATE DATABASE IF NOT EXISTS devbook;

use devbook;
DROP TABLE IF EXISTS publicacoes
DROP TABLE IF EXISTS seguidores;
DROP TABLE IF EXISTS usuarios;
CREATE TABLE usuarios (
  id int auto_increment primary key,
  nome varchar(50) not null,
  nick varchar(50) not null unique,
  email varchar(255) not null unique,
  senha varchar(100) not null,
  criadoEm timestamp default curdate()
) ENGINE=INNODB;

CREATE TABLE seguidores (
  usuario_id int not null,
  foreign key (usuario_id) references usuarios(id)
  on delete cascade,
  seguidor_id int not null,
  foreign key (seguidor_id) references usuarios(id) on delete cascade,
  primary key(usuario_id, seguidor_id)
)

CREATE TABLE publicacoes (
  id int auto_increment primary key,
  titulo varchar(20) not null,
  conteudo varchar(300) not null,

  autor_id int not null,
  foreign key (autor_id) references usuarios(id)
  on delete cascade,

  curtidas int default 0,
  criadoEm timestamp default curdate()
)