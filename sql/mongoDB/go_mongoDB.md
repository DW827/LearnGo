### 安装mongoDB和驱动
```
go get github.com/mongodb/mongo
go get github.com/mongodb/mongo-go-driver
```

### go连接mongodb
```go
package main

import (
    "contetx"
    "fmt"
    "log"

    "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
    // 设置客户端连接配置
    clientOption := options.Client().ApplyURI("mongodb://localhost:27017")

    // 连接MongoDB
    client, err := mongo.Connect(context.TODO(), clientOption)
    if err != nil {
        // 打印err，然后退出应用程序
        log.Fatal(err)
    }

    // 检查连接
    err = client.Ping(context.TODO(), nil)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("connect to mongodb !")

    // 指定获取要操作的数据集
    collection := client.Database("q1mi").Collection("student")

    // 断开连接
    err = client.Disconnect(context.TODO())
    if err != nil {
    	log.Fatal(err)
    }
    fmt.Println("Connection to MongoDB closed.")
}

```

### 连接池模式
```go
import (
    "context"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToDB(uri, name string, timeout time.Duration, num uint64) (*mongo.Database, err) {
    ctx, cancel := context.WithTimeout(context.Background(), timeout)
    defer cancel()

    o := options.Client().ApplyURI(uri)     // 连接设置
    o.SetMaxPoolSize(num)       // 连接池

    client, err := mongo.Connect(ctx, o)    // 连接
    if err != nil {
        return nil, err
    }
    return client.Database(name), nil
}
```

### BSON
MongoDB中的JSON文档存储在名为BSON(二进制编码的JSON)的二进制表示中。与其他将JSON数据存储为简单字符串和数字的数据库不同，BSON编码扩展了JSON表示，使其包含额外的类型，如int、long、date、浮点数和decimal128。这使得应用程序更容易可靠地处理、排序和比较数据。

### CRUD
```go
// 学生模型
type Student struct {
    Name    string
    Age     int
}

s1 := Student{"Tom", 12}
s2 := Student{"Mike", 13}
s3 := Student{"Bob", 11}
```
```go
// 插入文档
// collection.InsertOne()插入一条文档记录
insertResult, err := collection.InsertOne(context.TODO(), s1)
if err != nil {
	log.Fatal(err)
}

fmt.Println("Inserted a single document: ", insertResult.InsertedID)

// collection.InsertMany()插入文档多条记录
students := []interface{}{s2, s3}
insertManyResult, err := collection.InsertMany(context.TODO(), students)
if err != nil {
	log.Fatal(err)
}
fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)
```
```go
// 更新文档
// updateone()方法允许你更新单个文档记录
// 一个筛选器文档来匹配数据库中的文档
// 用来查找name字段与’张三’或’李四’匹配的文档:
filter := bson.D{{
	"name",
	bson.D{{
		"$in",
		bson.A{"张三", "李四"},
	}},
}}
// 用来查找name字段与'Tom'匹配的文档:
filter := bson.D{{"name", "Tom"}}

// 一个更新文档来描述更新操作
update := bson.D{
    {"$inc", bson.D{
        {"name", "Tomy"},
    }},
}

// 更新
updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
if err != nil {
	log.Fatal(err)
}
fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
```
```go
// 查找文档
// collection.FindOne()返回一个可以解码为值的结果。
// 创建一个Student变量用来接收查询的结果
var result Student
err = collection.FindOne(context.TODO(), filter).Decode(&result)
if err != nil {
	log.Fatal(err)
}
fmt.Printf("Found a single document: %+v\n", result)

// collection.Find()返回一个游标,用来查找多个文档。
// 查询多个
// 将选项传递给Find()
// options包设置一个限制以便只返回两个文档。
findOptions := options.Find()
findOptions.SetLimit(2)

// 定义一个切片用来存储查询结果
var results []*Student

// 把bson.D{{}}作为一个filter来匹配所有文档
cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
if err != nil {
	log.Fatal(err)
}

// 查找多个文档返回一个游标
// 遍历游标允许我们一次解码一个文档
for cur.Next(context.TODO()) {
	// 创建一个值，将单个文档解码为该值
	var elem Student
	err := cur.Decode(&elem)
	if err != nil {
		log.Fatal(err)
	}
	results = append(results, &elem)
}

if err := cur.Err(); err != nil {
	log.Fatal(err)
}

// 完成后关闭游标
cur.Close(context.TODO())
fmt.Printf("Found multiple documents (array of pointers): %#v\n", results)
```
```go
// collection.DeleteOne()或collection.DeleteMany()删除文档。
// 删除名字是小黄的那个
deleteResult1, err := collection.DeleteOne(context.TODO(), bson.D{{"name","小黄"}})
if err != nil {
	log.Fatal(err)
}
fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult1.DeletedCount)

// 删除所有
deleteResult2, err := collection.DeleteMany(context.TODO(), bson.D{{}})
if err != nil {
	log.Fatal(err)
}
fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult2.DeletedCount)
```