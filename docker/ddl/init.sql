START TRANSACTION;

DROP TABLE IF EXISTS messages;
DROP TABLE IF EXISTS users;

-- メッセージ
CREATE TABLE IF NOT EXISTS messages (
  title     VARCHAR(255) NOT NULL comment 'タイトル',
  message   VARCHAR(255) NOT NULL comment 'メッセージ'
)
comment = 'メッセージ';

INSERT INTO `messages`(`title`, `message`) VALUES ('Welcome','Hello World!');
INSERT INTO `messages`(`title`, `message`) VALUES ('使用するRDBMSは','MySQLです。');
INSERT INTO `messages`(`title`, `message`) VALUES ('使用する言語とフレームワークは？','Go言語とechoです。');

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

INSERT INTO users (id, name, age, job_large_type_id, job_middle_type_id, job_small_type_id, job_name, job_term) VALUES (1, 'ジョブ一郎', 32, '01', '0101', '010101', 'Web系エンジニア（フロントエンド）', 5);
INSERT INTO users (id, name, age, job_large_type_id, job_middle_type_id, job_small_type_id, job_name, job_term) VALUES (2, 'ジョブ二郎', 33, '01', '0101', '010102', 'Web系エンジニア（バックエンド）', 5);
INSERT INTO users (id, name, age, job_large_type_id, job_middle_type_id, job_small_type_id, job_name, job_term) VALUES (3, 'ジョブ三郎', 34, '01', '0101', '010103', 'Web系エンジニア（インフラ）', 5);
INSERT INTO users (id, name, age, job_large_type_id, job_middle_type_id, job_small_type_id, job_name, job_term) VALUES (4, 'ジョブ四郎', 35, '01', '0101', null, 'エンジニア', 5);
INSERT INTO users (id, name, age, job_large_type_id, job_middle_type_id, job_small_type_id, job_name, job_term) VALUES (5, 'ジョブ五郎', 36, '02', '0202', '020201', '小学校教師', 5);
INSERT INTO users (id, name, age, job_large_type_id, job_middle_type_id, job_small_type_id, job_name, job_term) VALUES (6, 'ジョブ六郎', 37, '02', '0202', '020201', '中学校教師', 5);
INSERT INTO users (id, name, age, job_large_type_id, job_middle_type_id, job_small_type_id, job_name, job_term) VALUES (7, 'ジョブ七郎', 38, '02', '0202', null, '教師', 5);
INSERT INTO users (id, name, age, job_large_type_id, job_middle_type_id, job_small_type_id, job_name, job_term) VALUES (8, 'ジョブ八郎', 39, '03', null, null, '営業', 5);
INSERT INTO users (id, name, age, job_large_type_id, job_middle_type_id, job_small_type_id, job_name, job_term) VALUES (9, 'ジョブ九郎', 40, '03', '0303', null, '不動産営業', 5);
INSERT INTO users (id, name, age, job_large_type_id, job_middle_type_id, job_small_type_id, job_name, job_term) VALUES (10, 'ジョブ十郎', 20, '03', '0303', '030301', 'サポート', 5);

COMMIT;
