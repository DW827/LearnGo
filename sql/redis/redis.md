### [redis](https://www.runoob.com/redis/redis-data-types.html)

### redis应用
+ 缓存系统，减轻主数据库（MySQL）的压力。
+ 计数场景，比如微博、抖音中的关注数和粉丝数。
+ 热门排行榜，需要排序的场景特别适合使用ZSET。
+ 利用LIST可以实现队列的功能。

### 数据类型
键的类型只能是字符串。
值的类型有五种。
##### 1. string
一个键值对。
```
> set key value
> get key
> del key
>
```

##### 2. hash
包含键值对的集合。
```
> hmset key field value [field value ...]
> hget key field
>
```
##### 3. list
Redis 列表是简单的字符串列表，按照插入顺序排序。你可以添加一个元素到列表的头部（左边）或者尾部（右边）。
```
> lpush key value [value ...]
> rpush key value [value ...]
> lrange key start stop
> lindex key index
>
```

##### 4. set
Redis 的 Set 是 string 类型的无序集合。

集合是通过哈希表实现的，所以添加，删除，查找的复杂度都是 O(1)。
```
> sadd key member [member ...]
> smembers key
> sismember key member
>
```

##### 5. zset
Redis zset 和 set 一样也是string类型元素的集合,且不允许重复的成员。
不同的是每个元素都会关联一个double类型的分数。redis正是通过分数来为集合中的成员进行从小到大的排序。

zset的成员是唯一的,但分数(score)却可以重复。
```
> zadd key score1 member1[score member]
> zrange key start stop [WITHSCORES]
```

##### 6. HyperLogLog
用来做基数统计的算法.
HyperLogLog 只会根据输入元素来计算基数，而不会储存输入元素本身.
```
> pfadd key element [element ...]
> pfadd key element [element ...]
> pfcount key [key ...]
>
```

##### 7. pub/sub
Redis 发布订阅 (pub/sub) 是一种消息通信模式：发送者 (pub) 发送消息，订阅者 (sub) 接收消息。

Redis 客户端可以订阅任意数量的频道。
```
// 创建订阅频道名为 chanannel1
> subscribe channel [channel ...]
redis 127.0.0.1:6379> SUBSCRIBE runoobChat

Reading messages... (press Ctrl-C to quit)
1) "subscribe"
2) "runoobChat"
3) (integer) 1

redis 127.0.0.1:6379> PUBLISH runoobChat "Redis PUBLISH test"

(integer) 1

redis 127.0.0.1:6379> PUBLISH runoobChat "Learn redis by runoob.com"

(integer) 1

# 订阅者的客户端会显示如下消息
 1) "message"
2) "runoobChat"
3) "Redis PUBLISH test"
 1) "message"
2) "runoobChat"
3) "Learn redis by runoob.com"
```

### cmd操作redis
```
// 连接redis
> redis-cli
127.0.0.1:6379>
```

```
// 查看所有的键
> keys *
// 查看键的值类型
> TYPE key
// 删除键
> DEL key
```
```
// 获取 redis 配置
> CONFIG GET *
// 获取 redis 目录
> config get dir
// 在 redis 安装目录中创建dump.rdb文件,创建当前数据库的备份。
> SAVE 
```

### redis事务
以 MULTI 开始一个事务， 然后将多个命令入队到事务中， 最后由 EXEC 命令触发事务， 一并执行事务中的所有命令。
Redis 没有在事务上增加任何维持原子性的机制，所以 Redis 事务的执行并不是原子性的。

事务可以理解为一个打包的批量执行脚本，但批量指令并非原子化的操作，中间某条指令的失败不会导致前面已做指令的回滚，也不会造成后续的指令不做。

### Redis Stream
Redis Stream 主要用于消息队列（MQ，Message Queue）。
Redis 本身是有一个 Redis 发布订阅 (pub/sub) 来实现消息队列的功能，但它有个缺点就是消息无法持久化，如果出现网络断开、Redis 宕机等，消息就会被丢弃。

Redis Stream 提供了消息的持久化和主备复制功能，可以让任何客户端访问任何时刻的数据，并且能记住每一个客户端的访问位置，还能保证消息不丢失。
它有一个消息链表，将所有加入的消息都串起来，每个消息都有一个唯一的 ID 和对应的内容。

使用 XADD 向队列添加消息，如果指定的队列不存在，则创建一个队列，XADD 语法格式：

    XADD key ID field value [field value ...]
key ：队列名称，如果不存在就创建
ID ：消息 id，我们使用 * 表示由 redis 生成，可以自定义，但是要自己保证递增性。
field value ： 记录。