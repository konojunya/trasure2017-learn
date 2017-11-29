-- +migrate Up
CREATE TABLE `daily_report` (
  `report_id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `title` varchar(256) NOT NULL COMMENT 'title',
  `body` TEXT COMMENT 'article body',
  `created` timestamp NOT NULL DEFAULT NOW() COMMENT 'when created',
  `updated` timestamp NOT NULL DEFAULT NOW() ON UPDATE CURRENT_TIMESTAMP COMMENT 'when last updated',
  PRIMARY KEY (`report_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COMMENT='list of reports';

-- +migrate Down
DROP TABLE daily_report;
