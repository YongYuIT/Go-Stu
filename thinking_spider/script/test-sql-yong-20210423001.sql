with clean_data_today as (
    select *
    from (
             select *, row_number() over (partition by asin order by created_at desc) as del_index
             from new_release_prod_records) t1
    where del_index = 1
),
     clean_data_yestaday as (
         select *
         from (
                  select *, row_number() over (partition by asin order by created_at desc) as del_index
                  from `new_release_prod_records-bak20210422001`) t1
         where del_index = 1
     )
select grow         a_grow,
       type1        b_type1,
       type2        c_type2,
       type3        d_type3,
       type4        e_type4,
       `index`      f_index,
       asin         g_asin,
       ratings      h_ratings,
       starts       i_starts,
       price        j_price,
       main_pic_url k_main_pic_url
from (
         select today.ratings - yestaday.ratings as grow, today.*
         from clean_data_today today
                  inner join clean_data_yestaday yestaday on today.asin = yestaday.asin
     ) t1
order by grow desc,type1,type2,type3,type4
limit 100