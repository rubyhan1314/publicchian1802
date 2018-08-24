package BLC

//专程存储一些常量值
const DBName  = "blockchain_%s.db" //数据库的名字
/*
数据库名：
blockchain_port.db
blockchian_3000.db
blockchian_3001.db

fmt.Printf("%s",s1)
	%s,%v,%d,%f,%T,%t,%p..........

fmt.Sprintf("%s",s1)-->string
fmt.Sprintf(DBName,nodeid)-->string
fmt.Sprintf("blockchain_%s.db","3000")-->string:blockchain_3000.db

...
 */
const BlockBucketName = "blocks" //定义bucket
