START TRANSACTION;

DROP TABLE IF EXISTS messages;

-- メッセージ
CREATE TABLE IF NOT EXISTS messages (
  title     VARCHAR(255) NOT NULL comment 'タイトル',
  message   VARCHAR(255) NOT NULL comment 'メッセージ'
)
comment = 'メッセージ';

INSERT INTO `messages`(`title`, `message`) VALUES ('Welcome','Hello World!');
INSERT INTO `messages`(`title`, `message`) VALUES ('使用するRDBMSは','MySQLです。');
INSERT INTO `messages`(`title`, `message`) VALUES ('使用する言語とフレームワークは？','Go言語とechoです。');
COMMIT;
