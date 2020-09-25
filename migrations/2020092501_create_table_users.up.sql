CREATE TABLE IF NOT EXISTS users(
    id int NOT NULL comment 'ユーザーID',
    name VARCHAR(255) NOT NULL comment '名前',
    age int NOT NULL comment '年齢',
    job_large_type_id  VARCHAR(2) comment '職種大分類ID',
    job_middle_type_id VARCHAR(4) comment '職種中分類ID',
    job_small_type_id VARCHAR(6) comment '職種小分類ID',
    job_name VARCHAR(255) NOT NULL comment '職種名',
    job_term int NOT NULL comment '仕事の期間'
);
