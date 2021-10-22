# 售后追踪平台系统

## 1、安装`MySQL5.7`版本数据库及创建表

#### 1.1、安装数据库方法自行百度，安装后需将 root 密码设置为 root

#### 1.2、登录到数据库创建数据库

```sql
$
CREATE
database aftersale CHARACTER SET utf8 COLLATE utf8_general_ci;
```

#### 1.3、切换到创建好的数据库中创建表

- 切换数据库命令：`use databaseName`

##### 1.3.1、clientAddr 表

```sql
$
CREATE TABLE `clientAddr`
(
    `cid`      int(4) NOT NULL AUTO_INCREMENT,
    `username` varchar(16) NOT NULL COMMENT '客户名称',
    `phone`    varchar(11)  DEFAULT NULL COMMENT '联系方式',
    `address`  varchar(128) DEFAULT NULL COMMENT '收货地址',
    PRIMARY KEY (`cid`)
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8;
```

##### 1.3.2、sale表

```sql
$
CREATE TABLE `sale`
(
    `sid`      int(2) NOT NULL AUTO_INCREMENT,
    `saleName` varchar(16) DEFAULT NULL COMMENT '销售员名称',
    `phone`    varchar(11) DEFAULT NULL COMMENT '联系方式',
    PRIMARY KEY (`sid`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8;
```

##### 1.3.3、ldss表
```sql
$
CREATE TABLE `ldss`
(
    `lid`          int NOT NULL AUTO_INCREMENT,
    `date`         varchar(64)  DEFAULT NULL COMMENT '时间日期',
    `ldssnum`      varchar(16)  NOT NULL COMMENT '物流单号',
    `lgss`         varchar(32)  NOT NULL COMMENT '物流名称',
    `mateInfo`     varchar(32)  NOT NULL COMMENT '物料信息',
    `amount`       int(3) NOT NULL COMMENT '数量',
    `bad_rea`      varchar(128) NOT NULL COMMENT '不良因素',
    `status`       int(1) DEFAULT '0' COMMENT '状态0为返厂1为维修2为以返',
    `remark`       varchar(512) default "" comment '备注.，用来记录收货详细信息',
    `pro_ret`      varchar(64)  default "" comment '处理结果',
    `ret_date`     varchar(64)  default "" comment '返回时间',
    `ret_ldss`     varchar(64)  default "" comment '返回单号',
    `ret_ldssname` varchar(64)  default "" comment '返回快递公司',
    `ret_remark`   varchar(64)  default "" comment '返回备注内容',
    `shipNum`      int(3) default 0 comment '发货次数',
    `sale_id`      int(2) DEFAULT NULL COMMENT '关联的销售员id',
    `c_id`         int(3) DEFAULT NULL COMMENT '关联的客户地址id',
    PRIMARY KEY (`lid`),
    KEY            `sale_id` (`sale_id`),
    KEY            `c_id` (`c_id`),
    CONSTRAINT `ldss_ibfk_1` FOREIGN KEY (`sale_id`) REFERENCES `sale` (`sid`),
    CONSTRAINT `ldss_ibfk_2` FOREIGN KEY (`c_id`) REFERENCES `clientAddr` (`cid`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
```

##### 1.3.4、插入销售员数据

```sql
$
insert into sale(saleName) values('陈小兵'),('李珊珊'),('林建立'),('余燕珊'),('黄静'),('王芷君'),('蒋镇霞'),('吴思燕'),('庄雪怡'),('康磊'),('郭伟'),('庞辉'),('其它');
```

##### 1.3.5、新增字段

```sql
$
ALTER TABLE ldss
    ADD pstate int(1) not null default 0 comment '实收状态,1为返回到客户,2为客户以确定' AFTER paid;
$
alter table ldss
    add paid int(3) not null default 0 comment '实收数量' after amount;
$
alter table ldss
    add sentnum int(3) not null default 0 comment '已发数量' after pstate;
```

##### 1.3.6 新增表及插入数据

```sql
$
create table user
(
    uid      int(4) not null primary key auto_increment,
    username varchar(32) not null comment '用户名',
    phone    varchar(16) not null comment '联系方式',
    passwd   varchar(32) not null comment '用户名'
)engine=InnoDB default charset=utf8;
$
insert into user(username,phone,passwd) value('lianli','18888888888','d956363bfb7caaaa264d4e7a0b095969');
```

## 2、安装Go环境

#### 2.1、执行下面命令下载 go 安装包

```bash
$ wget https://studygolang.com/dl/golang/go1.16.7.linux-amd64.tar.gz
```

#### 2.2、执行下面命令解压缩下载好的文件到 `/usr/local` 目录下

```bash
$ tar -C /usr/local/ -zxvf go1.16.7.linux-amd64.tar.gz
```

#### 2.3、配置环境变量

- 在`/etc/profile` 或 `~/.bashrc` 文件中追加如下内容：

```shell
export GOROOT=/usr/local/go
export GOPATH=/home/用户名/data/code
export PATH=$PATH:$GOROOT/bin
export PATH=$PATH:$GOPATH/bin
```

> **注：GOPATH 的目录以自身目录情况为准**

执行下面命令查看版本号，输出如下结果则说明安装成功：

```bash
$ go version
go version go1.16.6 linux/amd64
```

## 3、安装部署

#### 3.1、将 afterSaleSys 文件打包(打包命令如下)复制到需要部署的设备上

```bash
tar -zcvf 需要打包的文件夹.tar.gz
```

#### 3.2、在需要部署的机器上解压文件，命令如下：

```bash
tar -zxvf xxxxx.tar.gz
```

#### 3.3、切换到加压后的目录下执行如下命令:

```bash
$ cd 解压后的文件目录  # 切换目录
$ go mod tidy
```

> 注：除了 view、cfg、static 目录外都需在对应目录下执行 `go mod tidy` 命令检查相关依赖

#### 3.4、退回到 afterSaleSys 目录下执行如下命令运行程序

```bash
$ go run main.go
```

至此整个程序部署完成了，若以为此服务器配置过路由映射及绑定过域名即可通过以下两个地址打开进行操作(没有做映射和绑定过域名的自行搜索绑定方法)：

```bash
http://xxx.xxx:6969/          # 客户操作地址链接
http://xxx.xxx:6969/admin     # 后台管理操作地址链接
```

### 后台管理登录账号密码

```bash
# 管理员用户
username: lianli
passwd: lianl666/*-
# 售后用户
username: shouhou
passwd: shouh888***
# 仓库用户
username: cangku
passwd: cangk789///
```

### 注意事项

每次更新后更新到线上前需修改 `utils/readIni.go` 文件中的配置文件目录为 `/root/code/src/afterSaleSys/cfg/config.ini`，否则会出现初始化数据库错误现象。
并修改微信消息推送中配置文件的路径，修改路径`wechatMsgPush/util/readIni.go`, 修改为：`/root/code/src/afterSaleSys/wechatMsgPush/conf/wxconf.ini`

## 不足

- 登录注册无需输入密码，后期有时间再补上。
    - 没有短信验证码登录功能。
- 后台管理无法增加或修改销售员信息。
- 接收工单与处理完毕发货之间缺少中间操作(`如维修员车间维修后更改工单状态`)环节。
- 后台管理页面分页功能尚未实现。
- 其它不足。

## 新增内容

- user 表新增字段

```sql
$
alter table user
    add utype int(1) not null default 0 comment '用户类型,默认0为管理员用户，1为售后用户，2为仓库用户';
```

- 创建 retuenGoods 表，语法如下：

```sql
$
create table
    returnGoods(
          rid      int not null primary key auto_increment,
          ldsnum   varchar(16)  default null comment '快递单号',
          ldsname  varchar(64)  default null comment '快递名称',
          matename varchar(32)  default null comment '商品名称',
          reason   varchar(256) default null comment '退货原因',
          renum    int(3) default 0 comment '退货数量',
          paid     int(3) default 0 comment '实收数量',
          method   varchar(256) default null comment '处理方式',
          stat     int(1) default null comment '单据状态，0为返厂，1为已接收，2为需销售处理，3为处理成功',
          crtdate  varchar(64)  default null comment '创建时间',
          dealdate varchar(64)  default null comment '处理时间',
          sendnum  int(3) default 0 comment '返回数量',
          ssid     int          default 0 comment '关联的销售id',
          foreign key (ssid) references sale (sid),
          ccid     int          default 0 comment '关联的客户id',
          foreign key (ccid) references clientAddr (cid)
    )engine=InnoDB default charset=utf8 collate=utf8_general_ci;
```

###### 创建记录返回信息表，如下：
```sql
create table retInfo(
            retId int primary key auto_increment comment '主键',
            retDate varchar(64) not null default '' comment '返回时间',
            retLdsNum varchar(16) not null default '' comment '返回快递单号',
            retLdsName varchar(32) not null default '' comment '返回物流名称', 
            retMate varchar(128) not null default '' comment '返回商品名称', 
            retAmount int(3) not null default 0 comment '返回数量',
            retMark varchar(256) comment '返回备注信息',
            ldss_id  int   comment '关联的 ldss 表 lid，用于根据 lid 查询数据',
            foreign key (ldss_id) references ldss (lid)
        );
```

##### 本次更新需要进 wechatMsgPush 目录中执行 go mod tidy / go get 下载相关依赖项
lsof -i:6969 | kill -9 `awk '{}print $2'`
nohup go run main.go >> /data/aftersys.log & 2>&1  

<font color=orange>sale 销售表中- sid=16 的已经清空。下次新增销售时可对该编号进行重新修改添加。添加方法如下：</font>
```sql
update sale set saleName='姓名',phone='电话' where sid=16;
```