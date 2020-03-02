package db_try

//gp 用户自定义函数
/*
create schema f_func_test;
grant all on schema f_func_test to yuyong;
create table f_func_test.f_stu_info(
    pkid serial PRIMARY KEY,
    f_stu_name varchar(255),
    f_stu_age varchar(255)
);
grant all on f_func_test.f_stu_info to yuyong;

insert into f_func_test.f_stu_info(f_stu_name,f_stu_age)
values('aaa','111'),
      ('bbb','222'),
      ('ccc','333');

select * from f_func_test.f_stu_info;


create table f_func_test.f_stu_info_cpy(
    pkid serial PRIMARY KEY,
    f_stu_name varchar(255),
    f_stu_age varchar(255)
);
grant all on f_func_test.f_stu_info_cpy to yuyong;

create or replace function copy_data_from_info(start_index numeric, end_index numeric) returns numeric as
$$
    declare
	insert_total numeric,
    insert_start_index numeric;

	begin
        --查询插入的起始位置
        select case when max(pkid) is null then -1 else max(pkid) end into insert_start_index from f_func_test.f_stu_info_cpy;
        insert into f_func_test.f_stu_info_cpy(f_stu_name,f_stu_age)
        select f_stu_name,f_stu_age from f_func_test.f_stu_info where pkid>=start_index and pkid<=end_index;
        select count(*) into insert_total from f_func_test.f_stu_info_cpy where pkid>insert_start_index;
        return insert_total;
    end;
$$ LANGUAGE plpgsql;

select * from f_func_test.f_stu_info_cpy;
select copy_data_from_info(0,2);
select * from f_func_test.f_stu_info_cpy;
*/
