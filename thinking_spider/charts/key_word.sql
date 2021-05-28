with data_clean as (
    # 如果数据重复，保留页面指数较小（较靠前的数据）
    select *
    from (
             select *, row_number() over (partition by asin order by page_num asc) as r_num
             from (
                      # 计算页面指数
                      select *,
                             (page - 1) * 40 + page_index                 as page_num,
                             concat('https://www.amazon.com', detial_url) as main_detail
                      from key_word_prod_records
                      where key_word like '%XXX%'
                        and uuid != ''
                        and asin != ''
                  ) t1
         ) t2
    where r_num = 1
)
select asin         as a_asin,
       page         as b_page,
       page_index   as c_page_index,
       titles       as d_titles,
       ratings      as e_ratings,
       starts       as f_starts,
       price        as g_price,
       price_level  as h_price_level,
       main_pic_url as i_main_pic_url,
       main_detail  as j_detail_url
from data_clean
order by ratings desc