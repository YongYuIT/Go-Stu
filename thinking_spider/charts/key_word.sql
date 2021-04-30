with data_clean as (
    # 如果数据重复，保留页面指数较小（较靠前的数据）
    select *
    from (
             select *, row_number() over (partition by asin order by page_num asc) as r_num
             from (
                      # 计算页面指数
                      select *, (page - 1) * 40 + page_index as page_num
                      from key_word_prod_records
                      where uuid != ''
                  ) t1
         ) t2
    where r_num = 1
)
select asin         as a_asin,
       titles       as b_titles,
       ratings      as c_ratings,
       starts       as d_starts,
       price        as e_price,
       main_pic_url as f_main_pic_url
from data_clean
order by ratings desc