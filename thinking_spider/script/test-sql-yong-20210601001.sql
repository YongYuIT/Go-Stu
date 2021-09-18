insert into key_word_prod_records_source
select *
from `key_word_prod_records-bak20210918001`
where id in (
    with data_clean as (
        # 如果数据重复，保留页面指数较小（较靠前的数据）
        select *
        from (
                 select *, row_number() over (partition by asin order by page_num asc) as r_num
                 from (
                          # 计算页面指数
                          select *, (page - 1) * 40 + page_index as page_num
                          from `key_word_prod_records-bak20210918001`
                          where key_word like '%XXX%'
                            and uuid != ''
                            and asin != ''
                      ) t1
             ) t2
        where r_num = 1
    )
    select id
    from data_clean)