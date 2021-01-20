cat words.txt | xargs -n 1 | sort | uniq -c | sort -nr | awk '{print $2" "$1}'


#xargs 分割字符串 -n 1表示每行输出一个 可以加-d指定分割符

#要使用uniq统计词频需要被统计文本相同字符前后在一起，所以先排序 uniq -c 表示同时输出出现次数

#sort -nr 其中-n表示把数字当做真正的数字处理(当数字被当做字符串处理，会出现11比2小的情况)

cat words.txt |tr -s ' ' '\n' |sort|uniq -c|sort -r|awk '{print $2,$1}'

# 1、首先cat命令查看words.txt
# 2、tr -s ' ' '\n'将空格都替换为换行 实现分词
# 3、sort排序 将分好的词按照顺序排序
# 4、uniq -c 统计重复次数（此步骤与上一步息息相关，-c原理是字符串相同则加一，如果不进行先排序的话将无法统计数目）
# 5、sort -r 将数目倒序排列
# 6、awk '{print $2,$1}' 将词频和词语调换位置打印出来

