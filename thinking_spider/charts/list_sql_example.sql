with clean_data_today as (
    select *
    from (
             select *, row_number() over (partition by asin order by created_at desc) as del_index
             from new_release_prod_records
             where task_index = 1620354621) t1
    where del_index = 1
),
     clean_data_yestaday as (
         select *
         from (
                  select *, row_number() over (partition by asin order by created_at desc) as del_index
                  from new_release_prod_records
                  where task_index = 1620270648) t1
         where del_index = 1
     )
select grow         a_grow,
       index_impv   b_index_impv,
       type1        c_type1,
       type2        d_type2,
       type3        e_type3,
       type4        f_type4,
       type5        g_type5,
       `index`      h_index,
       asin         i_asin,
       ratings      j_ratings,
       starts       k_starts,
       price        l_price,
       main_pic_url m_main_pic_url
from (
         select today.ratings - yestaday.ratings as grow, today.`index` - yestaday.`index` as index_impv, today.*
         from clean_data_today today
                  inner join clean_data_yestaday yestaday on today.asin = yestaday.asin
     ) t1
order by grow desc, index_impv desc, type1, type2, type3, type4, type5
limit 100