with data_clean as (
    # 如果数据重复，保留页面指数较小（较靠前的数据）
    select *
    from (
             select *, row_number() over (partition by asin order by page_num asc) as r_num
             from (
                      # 计算页面指数
                      select *, (page - 1) * 40 + page_index as page_num
                      from `key_word_prod_records-bak20210419001`
                      where uuid != ''
                  ) t1
         ) t2
    where r_num = 1
),
     # 排除大卖家较多的数据
     with_out_big_seller as (
         select *
         from (
                  select key_word, count(asin) as a_count from data_clean where ratings > 1000 group by key_word) t1
         where a_count > 5
     ),
     # 排除评分太低
     start_up_than_3_5 as (
         select *
         from (
                  select key_word, avg(v_starts) as avg_start
                  from (
                           #处理没有评分的项目
#                            select key_word, case when starts < 0 then 0 else starts end as v_starts
#                            from data_clean
#                            where key_word not in (
#                                select key_word
#                                from with_out_big_seller)) t1
                           select key_word, starts as v_starts
                           from data_clean
                           where key_word not in (
                               select key_word
                               from with_out_big_seller)
                             and starts >= 0) t1
                  group by key_word) t2
         where avg_start > 3.5
     ),
     # 排除留评太低
     ratings_rate_up_0_3 as (
         select *
         from (
                  select t1.key_word, t2.all_count / t1.all_count as rate
                  from (
                           select key_word, count(*) as all_count
                           from data_clean
                           where key_word in (select key_word from start_up_than_3_5)
                           group by key_word
                       ) t1
                           join(
                      select key_word, count(*) as all_count
                      from data_clean
                      where key_word in (select key_word from start_up_than_3_5)
                        and ratings > -1
                      group by key_word
                  ) t2 on t1.key_word = t2.key_word) t3
         where rate > 0.3
     )
select key_word, ratings, starts, asin, page, page_index
from data_clean
where key_word in (select key_word from ratings_rate_up_0_3)
order by key_word, ratings desc, asin