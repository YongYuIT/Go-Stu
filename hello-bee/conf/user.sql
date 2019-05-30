-- auto-generated definition
create table user
(
    userId     varchar(127)                        not null comment 'uuid',
    userName   varchar(127)                        not null comment '用户名',
    passwd     varchar(127)                        not null comment '密码(MD5)',
    createTime timestamp default CURRENT_TIMESTAMP not null comment '创建时间',
    updateTime timestamp default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '最后更新时间',
    token      varchar(127)                        null comment '最新时间戳',
    constraint user_userId_uindex
        unique (userId)
)
    comment '用户资料列表';

alter table user
    add primary key (userId);

