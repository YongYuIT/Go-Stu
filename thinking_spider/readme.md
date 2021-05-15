# DETAIL_TASK
./thinking_spider -t DETAIL_TASK -k "Fruit knife" > log1.txt

## fix "Incorrect string value..." error

ALTER TABLE prod_detail_records DEFAULT CHARACTER SET utf8mb4;
conn: "root:123456@(localhost:3306)/thinking_spider?parseTime=True&loc=Local"