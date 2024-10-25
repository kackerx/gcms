create database gcms;

create table t_gcms_user
(
    id         bigint unsigned auto_increment
        primary key,
    created_at datetime(3)             null,
    updated_at datetime(3)             null,
    deleted_at datetime(3)             null,
    username   varchar(255) default '' not null,
    password   varchar(128) default '' not null comment '用户密码',
    age        bigint       default 0  not null comment '年龄',
    constraint uni_t_gcms_user_username
        unique (username)
);

create index idx_t_gcms_user_deleted_at
    on t_gcms_user (deleted_at);