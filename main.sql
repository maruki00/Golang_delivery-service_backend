create database delivery;



create table users(
    id int primary key not null auto_increment,
    user_name varchar(100) not null unique,
    full_name varchar(255) not null,
    email varchar(255) not null unique,
    address text not null,
    password text not null,
    type varchar(56) not null default 'costumer'
);

create table auth(
    id int primary key primary key auto_increment,
    email varchar(255) not null unique,
    token text not null unique,
    user_id int not null
);

create table files (
    int primary key not null auto_increment
);

create table orders(
    id int primary key not null auto_increment,

);

create table users_orders(
    id long primary key not null auto_increment,
    
);

create table products(
    id int primary key not null auto_increment,
);

create table feedbacks(
    id int primary key not null auto_increment,
);

create table analytics(
    id int primary key not null auto_increment,
);

create table orders_tracking(
    id int primary key not null auto_increment,
);

create table notifications(
    id int primary key not null auto_increment,
);

create table jobs(
    id int primary key not null auto_increment,
);