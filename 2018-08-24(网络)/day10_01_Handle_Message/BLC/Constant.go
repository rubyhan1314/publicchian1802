package BLC

//专程存储一些常量值
const DBName = "blockchain_%s.db" //数据库的名字
/*
数据库名：
blockchain_port.db
blockchian_3000.db
blockchian_3001.db

fmt.Printf("%s",s1)
fmt.Sprintf("%s",s1)-->string
...
 */
const BlockBucketName = "blocks" //定义bucket

//定义程序版本
const NODE_VERSION = 1

//定义命令的长度
const COMMAND_LENGTH = 12

//类型
const BLOCK_TYPE = "block"
const TX_TYPE = "tx"

//具体的命令：
const COMMAND_VERSION = "version"

const COMMAND_GETBLOCKS = "getblocks"

const COMMAND_INV = "inv"

const COMMAND_GETDATA = "getdata"

const COMMAND_BLOCKDATA	 ="blockdata"
