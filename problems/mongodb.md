# MongoDB

```text
db.contest.aggregate([
    {"$group" : {_id:"$province", count:{$sum:1}}}
])
```

MongoDB 中聚合\(aggregate\)主要用于处理数据\(诸如统计平均值，求和等\)，并返回计算后的数据结果。

有点类似 **SQL** 语句中的 count\(\*\)。

```text
db.COLLECTION_NAME.aggregate(AGGREGATE_OPERATION)
```

```text
db.mycol.aggregate([{$group : {_id : "$by_user", num_tutorial : {$sum : 1}}}])
{
   "result" : [
      {
         "_id" : "runoob.com",
         "num_tutorial" : 2
      },
      {
         "_id" : "Neo4j",
         "num_tutorial" : 1
      }
   ],
   "ok" : 1
}
```

| $sum | 计算总和。 | db.mycol.aggregate\(\[{$group : {\_id : "$by\_user", num\_tutorial : {$sum : "$likes"}}}\]\) |
| :--- | :--- | :--- |
| $avg | 计算平均值 | db.mycol.aggregate\(\[{$group : {\_id : "$by\_user", num\_tutorial : {$avg : "$likes"}}}\]\) |
| $min | 获取集合中所有文档对应值得最小值。 | db.mycol.aggregate\(\[{$group : {\_id : "$by\_user", num\_tutorial : {$min : "$likes"}}}\]\) |
| $max | 获取集合中所有文档对应值得最大值。 | db.mycol.aggregate\(\[{$group : {\_id : "$by\_user", num\_tutorial : {$max : "$likes"}}}\]\) |
| $push | 在结果文档中插入值到一个数组中。 | db.mycol.aggregate\(\[{$group : {\_id : "$by\_user", url : {$push: "$url"}}}\]\) |
| $addToSet | 在结果文档中插入值到一个数组中，但不创建副本。 | db.mycol.aggregate\(\[{$group : {\_id : "$by\_user", url : {$addToSet : "$url"}}}\]\) |
| $first | 根据资源文档的排序获取第一个文档数据。 | db.mycol.aggregate\(\[{$group : {\_id : "$by\_user", first\_url : {$first : "$url"}}}\]\) |
| $last | 根据资源文档的排序获取最后一个文档数据 | db.mycol.aggregate\(\[{$group : {\_id : "$by\_user", last\_url : {$last : "$url"}}}\]\) |



* $project：修改输入文档的结构。可以用来重命名、增加或删除域，也可以用于创建计算结果以及嵌套文档。
* $match：用于过滤数据，只输出符合条件的文档。$match使用MongoDB的标准查询操作。
* $limit：用来限制MongoDB聚合管道返回的文档数。
* $skip：在聚合管道中跳过指定数量的文档，并返回余下的文档。
* $unwind：将文档中的某一个数组类型字段拆分成多条，每条包含数组中的一个值。
* $group：将集合中的文档分组，可用于统计结果。
* $sort：将输入文档排序后输出。
* $geoNear：输出接近某一地理位置的有序文档。、

1、$project实例

```text
db.article.aggregate(
    { $project : {
        title : 1 ,
        author : 1 ,
    }}
 );
```

这样的话结果中就只还有\_id,tilte和author三个字段了，默认情况下\_id字段是被包含的，如果要想不包含\_id话可以这样:

```text
db.article.aggregate(
    { $project : {
        _id : 0 ,
        title : 1 ,
        author : 1
    }});
```

2.$match实例

```text
db.articles.aggregate( [
                        { $match : { score : { $gt : 70, $lte : 90 } } },
                        { $group: { _id: null, count: { $sum: 1 } } }
                       ] );
```

$match用于获取分数大于70小于或等于90记录，然后将符合条件的记录送到下一阶段$group管道操作符进行处理。

3.$skip实例

```text
db.article.aggregate(
    { $skip : 5 });
```

经过$skip管道操作符处理后，前五个文档被"过滤"掉。

[https://www.runoob.com/mongodb/mongodb-aggregate.html](https://www.runoob.com/mongodb/mongodb-aggregate.html)

