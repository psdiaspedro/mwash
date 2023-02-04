CREATE DATABASE IF NOT EXISTS mwash;
USE mwash;

DROP TABLE IF EXISTS usuarios;
DROP TABLE IF EXISTS propriedades;
DROP TABLE IF EXISTS agendamentos;

CREATE TABLE usuarios(
    id int auto_increment primary key,
    nome varchar(50) not null,
    email varchar(50) not null unique,
    senha varchar(100) not null,
    contato varchar(50) not null,
    admin boolean not null
) ENGINE=INNODB;

CREATE TABLE propriedades(
    id int auto_increment primary key,
    cliente_id int not null,
    FOREIGN KEY (cliente_id)
    REFERENCES usuarios(id)
    ON DELETE CASCADE,
    cidade varchar(30) not null,
    bairro varchar(30) not null,
    CEP varchar(15) not null,
    logadouro varchar(50) not null,
    numero varchar(10) not null,
    complemento varchar(100) default "",
    senha varchar(50) default "",
    acomodacao varchar (100) default "",
    wifi varchar(100) default "",
    outros varchar(100) default "",
    observacoes varchar(100) default "",
    cor int not null
) ENGINE=INNODB;  

CREATE TABLE agendamentos(
    id int auto_increment primary key,
    propriedade_id int not null,
    FOREIGN KEY (propriedade_id)
    REFERENCES propriedades(id)
    ON DELETE CASCADE,
    dia_agendamento DATE not null,
    checkin TIME (0),
    checkout TIME (0) not null,
    observacoes varchar(100) default ""
) ENGINE=INNODB;
