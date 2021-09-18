with data_clean as (
    # 如果数据重复，保留页面指数较小（较靠前的数据）
    select *
    from (
             select *, row_number() over (partition by asin order by page_num asc) as r_num
             from (
                      # 计算页面指数
                      select *, (page - 1) * 40 + page_index as page_num
                      from `key_word_prod_records-bak20210918001`
                      where 1 = 1
                        and key_word like '%stop%'
                        and uuid != ''
                        and asin != ''
                  ) t1
         ) t2
    where r_num = 1
),
     clean_detail_data as (
         select prod_detail_records.*
         from data_clean
                  left join prod_detail_records on data_clean.asin = prod_detail_records.asin
         where prod_detail_records.asin is not null
     )
        ,
     igno_char as (
         select replace(desc1, ',', ' ') as clean_desc1, clean_detail_data.*
         from clean_detail_data
     )
        ,
     split_result as (SELECT substring_index(substring_index(a.chain, ' ', b.help_topic_id + 1), ' ', - 1) AS word, asin
                      FROM (select clean_desc1 as chain, asin from igno_char) a
                               JOIN mysql.help_topic b ON b.help_topic_id <
                                                          (length(a.chain) - length(replace(a.chain, ' ', '')) + 1)
     )
select *
from (
         select word, count(word), count(distinct asin) as w_count, count(word) / count(distinct asin)
         from split_result
         group by word) tmp
order by w_count desc