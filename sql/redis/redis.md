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
>set key value
>get key
>del key
>
```

##### 2. hash
包含键值对的集合。
```
>hmset key field value [field value ...]
>hget key field
>
```
##### 3. list
Redis 列表是简单的字符串列表，按照插入顺序排序。你可以添加一个元素到列表的头部（左边）或者尾部（右边）。
```
>lpush key value [value ...]
>rpush key value [value ...]
>
```

##### 4. set
Redis 的 Set 是 string 类型的无序集合。

集合是通过哈希表实现的，所以添加，删除，查找的复杂度都是 O(1)。
```
>sadd key member [member ...]
>smembers key
>
```

##### 5. zset
Redis zset 和 set 一样也是string类型元素的集合,且不允许重复的成员。
不同的是每个元素都会关联一个double类型的分数。redis正是通过分数来为集合中的成员进行从小到大的排序。

zset的成员是唯一的,但分数(score)却可以重复。
```
>zadd key score1 member1[score member]
>zrange key start stop [WITHSCORES]
```

### cmd操作redis
```
// 连接redis
>redis-cli
127.0.0.1:6379>
```
```
// 查看所有的键
>keys *
// 查看键的值类型
>TYPE key
// 删除键
>DEL key
// 
```