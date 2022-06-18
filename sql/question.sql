/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `mc_question` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `question` varchar(1024) DEFAULT NULL,
  `choicea` varchar(256) DEFAULT NULL,
  `choiceb` varchar(256) DEFAULT NULL,
  `choicec` varchar(256) DEFAULT NULL,
  `choiced` varchar(256) DEFAULT NULL,
  `answer` tinyint(2) DEFAULT '0',
  `score` tinyint(2) DEFAULT '0',
  `kind` tinyint(2) DEFAULT '0' COMMENT '类型',
  `grade` tinyint(2) DEFAULT '0' COMMENT '等级',
  `dif` tinyint(2) DEFAULT '0' COMMENT '难度',
  `stat` tinyint(2) DEFAULT '0' COMMENT '状态',
  `owner` int(11) DEFAULT '0'COMMENT '出题人',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;