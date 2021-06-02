with data_clean as (
    with data_clean as (
        # 如果数据重复，保留页面指数较小（较靠前的数据）
        select *
        from (
                 select *, row_number() over (partition by asin order by page_num asc) as r_num
                 from (
                          # 计算页面指数
                          select *, (page - 1) * 40 + page_index as page_num
                          from key_word_prod_records
                          where (/*key_word like '%wine corks%' or*/ key_word like '%mini knife%')
                            and uuid != ''
                            and asin != ''
                      ) t1
             ) t2
        where r_num = 1
    )
    select *
    from (
             select row_number() over (partition by price_level order by page_num asc) as all_index, data_clean.*
             from data_clean) tmp
    where all_index <= 200
),
     igno_char as (
         select replace(titles, ',', ' ') as clean_titles, data_clean.*
         from data_clean
     )
        ,
     split_result as (SELECT substring_index(substring_index(a.chain, ' ', b.help_topic_id + 1), ' ', - 1) AS word, asin
                      FROM (select clean_titles as chain, asin from igno_char) a
                               JOIN mysql.help_topic b ON b.help_topic_id <
                                                          (length(a.chain) - length(replace(a.chain, ' ', '')) + 1)
     )
select *
from (
         select word, count(word), count(distinct asin) as w_count, count(word) / count(distinct asin)
         from split_result
         group by word) tmp
order by w_count desc