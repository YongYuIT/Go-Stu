# select asin,page,page_index,ratings,starts,main_pic_url from `key_word_prod_records-bak20210420001` where uuid!='' limit 100
select type1        a_type1,
       type2        b_type2,
       type3        c_type3,
       type4        d_type4,
       `index`      e_index,
       asin         f_asin,
       ratings      g_ratings,
       starts       h_starts,
       price        i_price,
       main_pic_url j_main_pic_url
from new_release_prod_records
order by type1, type2, type3, type4, `index`