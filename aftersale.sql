-- MySQL dump 10.13  Distrib 5.7.35, for Linux (x86_64)
--
-- Host: localhost    Database: aftersale
-- ------------------------------------------------------
-- Server version	5.7.35-0ubuntu0.18.04.2-log

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `clientAddr`
--

DROP TABLE IF EXISTS `clientAddr`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `clientAddr` (
  `cid` int(4) NOT NULL AUTO_INCREMENT,
  `username` varchar(16) NOT NULL COMMENT '客户名称',
  `phone` varchar(11) DEFAULT NULL COMMENT '联系方式',
  `address` varchar(128) DEFAULT NULL COMMENT '收货地址',
  PRIMARY KEY (`cid`)
) ENGINE=InnoDB AUTO_INCREMENT=24 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `clientAddr`
--

LOCK TABLES `clientAddr` WRITE;
/*!40000 ALTER TABLE `clientAddr` DISABLE KEYS */;
INSERT INTO `clientAddr` VALUES (19,'测试','13724571522','测试测试测试222222'),(20,'测试账户2','13724571588',NULL),(21,'元红标','13682409834',NULL),(22,'dsbgrnr','13724571525',NULL),(23,'jgjktyjrfbn','13724571523',NULL);
/*!40000 ALTER TABLE `clientAddr` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `ldss`
--

DROP TABLE IF EXISTS `ldss`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `ldss` (
  `lid` int(11) NOT NULL AUTO_INCREMENT,
  `date` varchar(64) DEFAULT NULL COMMENT '时间日期',
  `ldssnum` varchar(16) NOT NULL COMMENT '物流单号',
  `lgss` varchar(32) NOT NULL COMMENT '物流名称',
  `mateInfo` varchar(32) NOT NULL COMMENT '物料信息',
  `amount` int(3) NOT NULL COMMENT '数量',
  `paid` int(3) NOT NULL DEFAULT '0' COMMENT '实收数量',
  `pstate` int(1) NOT NULL DEFAULT '0' COMMENT '实收状态,1为返回到客户,2为客户以确定',
  `sentnum` int(3) NOT NULL DEFAULT '0' COMMENT '已发数量',
  `bad_rea` varchar(128) NOT NULL COMMENT '不良因素',
  `status` int(1) DEFAULT '0' COMMENT '状态0为返厂1为维修2为以返',
  `remark` varchar(512) DEFAULT '' COMMENT '备注，用来记录收货详细信息',
  `pro_ret` varchar(64) DEFAULT '' COMMENT '处理结果',
  `ret_date` varchar(64) DEFAULT '' COMMENT '返回时间',
  `ret_ldss` varchar(64) DEFAULT '' COMMENT '返回单号',
  `ret_ldssname` varchar(64) DEFAULT '' COMMENT '返回快递公司',
  `ret_remark` varchar(64) DEFAULT '' COMMENT '返回备注内容',
  `shipNum` int(3) DEFAULT '0' COMMENT '  发货次数',
  `sale_id` int(2) DEFAULT NULL COMMENT '关联的销售员id',
  `c_id` int(3) DEFAULT NULL COMMENT '关联的客户地址id',
  PRIMARY KEY (`lid`),
  KEY `sale_id` (`sale_id`),
  KEY `c_id` (`c_id`),
  CONSTRAINT `ldss_ibfk_1` FOREIGN KEY (`sale_id`) REFERENCES `sale` (`sid`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `ldss_ibfk_2` FOREIGN KEY (`c_id`) REFERENCES `clientAddr` (`cid`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ldss`
--

LOCK TABLES `ldss` WRITE;
/*!40000 ALTER TABLE `ldss` DISABLE KEYS */;
INSERT INTO `ldss` VALUES (1,'2021-09-13 05:35:55','ccc4444444','fsddfgfd','测试物料',99,99,2,99,'ok',0,'','sdfdssdf ','2021-10-08 11:46:50','ccc4444455555','侧身','都是非法的',NULL,22,19),(2,'2021-09-16 04:11:12','cccc99999999','百达','测试2测试2',88,88,2,88,'OK',0,'','','2021-10-11 11:49:57','cccc65656565','zfsdfsd','dsfsgdfbfd',NULL,26,19),(3,'2021-09-13 05:05:49','ccccc8888888888','测测测测测测测测','东方闪电',89,80,2,80,'对符合规范哪个服南方姑娘',0,'sdvdfvdfbvfdg','交由产线处理，并以处理完毕','2021-09-24 08:45:03','ceshi777777','发货测试','讲话稿茶树菇讲话稿水电费第三方别的地方',NULL,21,19),(4,'2021-09-16 04:12:26','ccccc88888','饮水','时代大厦',66,66,2,66,'的深V单独',0,'ghm,o.,io.myuytmnrbrty','交由维修师傅进行维修处理，并以处理OK','2021-10-11 11:46:57','cccc999999','dscsdvdsf','scsdvdfv',NULL,22,19),(5,'2021-09-23 11:49:25','ccccc9999999999','侧二测9999','测999',99,99,2,0,'按时；代理商的付款计划',2,'lhkjhkjjhgjhgjhghj','ghj,jh y','','','','',NULL,14,19),(6,'2021-09-23 11:50:02','cccc10101010','测101010','测测测1010101',88,88,2,0,'胜多负少的',1,'6878786465456465','','','','','',NULL,16,19),(7,'2021-09-23 11:50:36','ccc1111','测测测11111','测测测测1111',11,11,2,0,'爱是飞洒',1,'ululygrhjtykubrnyu,tuymt','','','','','',NULL,15,19),(8,'2021-09-23 11:51:04','cccc222222','测测测测222222','测测测测测测测测2',22,22,2,10,'四渡赤水',0,'gfjtgmjty','dfhngmyumui','2021-10-22 10:58:46','ccc7878787878','sdfdfbdf','fngmghmtymghmgfgdbfg',1,16,19),(9,'2021-09-23 11:54:48','cccc3333333','测333','测测测测33333',33,33,2,30,'正常深V是',0,'xcbfdngh','gnfmy,78tyk8y','2021-10-22 10:48:46','8989898989','oioiojkjjj','sdfgetbtnyt',2,17,19),(10,'2021-09-23 11:55:13','cccc44444','测测44','测测测测测44',44,44,2,44,'是的是的发送到',0,'sdcsvdgbfg','okokokokoko','2021-10-15 10:24:41','cscscs3333333333','ceshiceshi','fgnghmhmyumty',1,18,19),(11,'2021-09-23 11:55:37','ccc5555555555','测测55555','测测测测测5',55,55,2,55,'是多少大V上的',0,'sdgdgbf','dfhbtmut,8ll,8nbef','2021-10-13 08:50:28','77777777','方便','dbfdnghnmhg',1,19,19),(12,'2021-09-23 11:56:08','cccc666666666666','测66','测测测测测6',66,66,2,66,'是大幅度发给',0,'afdbfngmyum','dfvdfbfgntyjy7','2021-10-13 08:45:04','8888888888','fafa ','dfbfdnmyu,iui',1,20,19),(13,'2021-09-23 11:56:37','ccc7777777777','测77','测测测测7',77,77,2,77,'电饭锅电饭锅和',1,'方便是大概发给你不','fgngmh,mui,uimtb','2021-10-12 06:59:10','9999999999','就要','的觉得符合大概比较地方考虑和是大概脚后跟脚后跟',2,21,19);
/*!40000 ALTER TABLE `ldss` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `retInfo`
--

DROP TABLE IF EXISTS `retInfo`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `retInfo` (
  `retId` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `retDate` varchar(64) NOT NULL DEFAULT '' COMMENT '返回时间',
  `retLdsNum` varchar(16) NOT NULL DEFAULT '' COMMENT '返回快递单号',
  `retLdsName` varchar(32) NOT NULL DEFAULT '' COMMENT '返回物流名称',
  `retMate` varchar(128) NOT NULL DEFAULT '' COMMENT '返回商品名称',
  `retAmount` int(3) NOT NULL DEFAULT '0' COMMENT '返回数量',
  `retMark` varchar(256) DEFAULT NULL COMMENT '返回备注信息',
  `ldss_id` int(11) DEFAULT NULL COMMENT '关联的 ldss 表 lid，用于根据 lid 查询数据',
  PRIMARY KEY (`retId`),
  KEY `retInfo_FK` (`ldss_id`),
  CONSTRAINT `retInfo_FK` FOREIGN KEY (`ldss_id`) REFERENCES `ldss` (`lid`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `retInfo`
--

LOCK TABLES `retInfo` WRITE;
/*!40000 ALTER TABLE `retInfo` DISABLE KEYS */;
INSERT INTO `retInfo` VALUES (5,'2021-10-12 05:55:20','65654564','你好','测测测测7',77,'我们都好',13),(6,'2021-10-12 06:59:10','9999999999','就要','测测测测7',77,'的觉得符合大概比较地方考虑和是大概脚后跟脚后跟',13),(7,'2021-10-13 08:45:04','8888888888','fafa ','测测测测测6',66,'dfbfdnmyu,iui',12),(8,'2021-10-13 08:50:28','77777777','方便','测测测测测5',55,'dbfdnghnmhg',11),(9,'2021-10-15 10:24:41','cscscs3333333333','ceshiceshi','测测测测测44',44,'fgnghmhmyumty',10),(10,'2021-10-15 10:25:14','cscscs3333333333','dsvfdvdfb','测测测测33333',20,'fnytkyu',9),(11,'2021-10-22 10:48:46','8989898989','oioiojkjjj','测测测测33333',10,'sdfgetbtnyt',9),(12,'2021-10-22 10:58:46','ccc7878787878','sdfdfbdf','测测测测测测测测2',10,'fngmghmtymghmgfgdbfg',8);
/*!40000 ALTER TABLE `retInfo` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `returnGoods`
--

DROP TABLE IF EXISTS `returnGoods`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `returnGoods` (
  `rid` int(11) NOT NULL AUTO_INCREMENT,
  `ldsnum` varchar(16) DEFAULT NULL COMMENT '快递单号',
  `ldsname` varchar(32) DEFAULT NULL COMMENT '快递名称',
  `matename` varchar(32) DEFAULT NULL COMMENT '商品名称',
  `reason` varchar(256) DEFAULT NULL COMMENT '退货原因',
  `renum` int(3) DEFAULT '0' COMMENT '退货数量',
  `paid` int(3) DEFAULT '0' COMMENT '实收数量',
  `method` varchar(256) DEFAULT NULL COMMENT '处理方式',
  `stat` int(1) DEFAULT NULL COMMENT '单据状态，0为返厂，1为已接收，2为需销售处理后需要发货，3为处理成功',
  `crtdate` varchar(64) DEFAULT NULL COMMENT '创建时间',
  `dealdate` varchar(64) DEFAULT NULL COMMENT '处理时间',
  `sendnum` int(3) DEFAULT '0' COMMENT '已换数量',
  `ssid` int(11) DEFAULT '0' COMMENT '关联的销售id',
  `ccid` int(11) DEFAULT '0' COMMENT '关联的客户id',
  PRIMARY KEY (`rid`),
  KEY `returnGoods_ibfk_1` (`ssid`),
  KEY `returnGoods_ibfk_2` (`ccid`),
  CONSTRAINT `returnGoods_ibfk_1` FOREIGN KEY (`ssid`) REFERENCES `sale` (`sid`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `returnGoods_ibfk_2` FOREIGN KEY (`ccid`) REFERENCES `clientAddr` (`cid`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `returnGoods`
--

LOCK TABLES `returnGoods` WRITE;
/*!40000 ALTER TABLE `returnGoods` DISABLE KEYS */;
INSERT INTO `returnGoods` VALUES (1,'ccc555555555','sdfdsfds','ghj,oi.iop/o','dfbntmyumyu',88,88,'从VB给你发',3,'2021-09-14 10:16:06','2021-09-15 11:31:44',60,20,20),(2,'ccc2222222','顺牛','加基森','开展客户电话公司单号',90,90,'以旧换新',2,'2021-09-14 05:32:47','2021-09-16 04:41:53',0,15,19),(3,'ssss111111111','顺意','最后一次测试','最后一次测试最后一次测试',66,66,'接受退货退款',3,'2021-09-16 04:14:02','2021-09-16 04:32:22',0,21,19),(4,'443432','德邦','3080','测试',2,2,'12',2,'2021-09-17 09:55:21','2021-09-17 10:10:10',0,14,21),(5,'cs33333333333','测试3','测试测测测3','感觉被忽悠了，不想要了',33,33,NULL,1,'2021-10-08 05:29:24',NULL,0,18,19);
/*!40000 ALTER TABLE `returnGoods` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sale`
--

DROP TABLE IF EXISTS `sale`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sale` (
  `sid` int(2) NOT NULL AUTO_INCREMENT,
  `saleName` varchar(16) DEFAULT NULL COMMENT '销售员名称',
  `phone` varchar(11) DEFAULT NULL COMMENT '联系方式',
  PRIMARY KEY (`sid`)
) ENGINE=InnoDB AUTO_INCREMENT=33 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sale`
--

LOCK TABLES `sale` WRITE;
/*!40000 ALTER TABLE `sale` DISABLE KEYS */;
INSERT INTO `sale` VALUES (14,'陈小兵','13922481979'),(15,'李珊珊','13632860056'),(16,'林建立','13432792743'),(17,'余燕珊','13750062003'),(18,'黄静','13794992054'),(19,'王芷君','15818264340'),(20,'蒋镇霞','18028961720'),(21,'吴思燕','18038273039'),(22,'庄雪怡','13077250526'),(23,'康磊','18025907087'),(24,'郭伟','18025907087'),(25,'庞辉','15920648872'),(26,'李广升','13713251729'),(27,'傅宝封','13263885312'),(28,'李冬杏','13432379768'),(29,'程途','18025294803'),(30,'刘燕娟','13087361713'),(31,'陈文君','18373636739'),(32,'徐锦珊','18002936913');
/*!40000 ALTER TABLE `sale` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user` (
  `uid` int(4) NOT NULL AUTO_INCREMENT,
  `username` varchar(32) NOT NULL COMMENT '用户名',
  `phone` varchar(16) NOT NULL COMMENT '联系方式',
  `passwd` varchar(32) NOT NULL COMMENT '用户名',
  `utype` int(1) NOT NULL DEFAULT '0' COMMENT '用户类型,默认0为管理员用户，1为售后用户，2为仓库用户',
  PRIMARY KEY (`uid`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (1,'lianli','18888888888','60ee9f82873aff3d52d7425920e0c8aa',0),(2,'shouhou','18025907087','2472bdbf546c5f39df99121851453330',1),(3,'cangku','17777777777','2a333ecd917a2d0d65d0192c44cea24c',2),(4,'test','10000000000','2a333ecd917a2d0d65d0192c44cea24c',2),(5,'zf','10000000000','2a333ecd917a2d0d65d0192c44cea24c',2),(6,'zhufeng','10000000000','2a333ecd917a2d0d65d0192c44cea24c',2);
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2021-10-22 17:52:34
