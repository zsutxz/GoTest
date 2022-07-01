/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `mc_question` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `question` varchar(1024) DEFAULT NULL,
  `choice_a` varchar(256) DEFAULT NULL,
  `choice_b` varchar(256) DEFAULT NULL,
  `choice_c` varchar(256) DEFAULT NULL,
  `choice_d` varchar(256) DEFAULT NULL,
  `answer` tinyint(2) DEFAULT '0',
  `score` tinyint(2) DEFAULT '0',
  `kind` tinyint(2) DEFAULT '0' COMMENT '类型',
  `grade` tinyint(2) DEFAULT '0' COMMENT '等级',
  `dif` tinyint(2) DEFAULT '0' COMMENT '难度',
  `stat` tinyint(2) DEFAULT '0' COMMENT '状态',
  `describe` varchar(1024) DEFAULT NULL COMMENT '答案描述',
  `owner` int(11) DEFAULT '0'COMMENT '出题人',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;


/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `wxuser` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `openid` varchar(64) DEFAULT NULL,
  `nickname` varchar(32) DEFAULT NULL,
  `city` varchar(32) DEFAULT NULL,
  `province` varchar(32) DEFAULT NULL,
  `country` varchar(32) DEFAULT NULL,
  `gender` tinyint(2) DEFAULT '0',
  `createdate` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;