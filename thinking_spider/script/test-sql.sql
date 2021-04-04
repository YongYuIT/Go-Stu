select t.key_word, sum(t.`range`) all_sum,avg(t.price),count(*),sum(t.`range`)/count(*)
from key_word_prod_records t
group by t.key_word
order by all_sum desc