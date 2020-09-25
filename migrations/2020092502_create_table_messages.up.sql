-- メッセージ
CREATE TABLE IF NOT EXISTS messages (
  id int NOT NULL comment 'メッセージID',
  user_id    int NOT NULL comment 'ユーザーID',
  title      VARCHAR(255) NOT NULL comment 'タイトル',
  message    VARCHAR(255) NOT NULL comment 'メッセージ'
)
comment = 'メッセージ';
