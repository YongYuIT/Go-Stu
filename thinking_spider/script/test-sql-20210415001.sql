select key_word,updated_at,sum
from(
        select  t.key_word,substr(t.updated_at ,1,10) as  updated_at , sum(t.ratings) as sum
        from
            (select key_word,updated_at,ratings ,row_number() over(partition by key_word,updated_at ,t.uuid order by updated_at desc )as num
             from key_word_prod_records_all t
             where t.ratings>=0
                /*and key_word in('Meatball maker','raceless double-sided tape','Spiral hangers')*/
            ) t

        where num =1
        group by t.key_word,substr(t.updated_at ,1,10)
            /*order by key_word,substr(t.updated_at ,1,10)*/
        order by substr(t.updated_at ,1,10),sum
    ) tt
#where sum<3000
order by substr(updated_at ,1,10),sum


;

select * from `key_word_prod_records-bak20210407001` t where t.key_word='Mop slippers'