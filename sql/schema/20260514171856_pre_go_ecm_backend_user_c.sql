-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE IF NOT EXISTS go_db_users (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'Account ID',
  `username` varchar(30) NOT NULL DEFAULT '' COMMENT 'Email',
  `password` varchar(32) NOT NULL DEFAULT '' COMMENT 'Password',
  `created_at` int(11) NOT NULL DEFAULT '0' COMMENT 'Created time',
  `updated_at` int(11) NOT NULL DEFAULT '0' COMMENT 'Updated time',
  `is_active` int(1) NOT NULL DEFAULT '0' COMMENT 'Is active',
 `deleted_at` DATETIME(3) NULL DEFAULT NULL COMMENT 'Deleted time'
  PRIMARY KEY(id),
  KEY `idx_email` (`username`),
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Account';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `go_db_users`;
-- +goose StatementEnd