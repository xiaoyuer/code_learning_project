# bash

### transport file

```text
awk '{
    for (i=1;i<=NF;i++){
        if (NR==1){
            res[i]=$i
        }
        else{
            res[i]=res[i]" "$i
        }
    }
}END{
    for(j=1;j<=NF;j++){
        print res[j]
    }
}' file.txt


#NF是当前行的field字段数；NR是正在处理的当前行数。
```

```text
sed -n "10p" file.txt
grep -n "" file.txt | grep -w '10' | cut -d: -f2
awk '{if(NR==10){print $0}}' file.txt # awk 条件 + 文件
```

```text
#grep -P '^([0-9]{3}-|\([0-9]{3}\) )[0-9]{3}-[0-9]{4}$' file.txt
awk '/^([0-9]{3}-|\([0-9]{3}\) )[0-9]{3}-[0-9]{4}$/' file.txt
#gawk '/^([0-9]{3}-|\([0-9]{3}\) )[0-9]{3}-[0-9]{4}$/' file.txt
```

```text
cat words.txt | xargs -n 1 | sort | uniq -c | sort -nr | awk '{print $2" "$1}'


#xargs 分割字符串 -n 1表示每行输出一个 可以加-d指定分割符

#要使用uniq统计词频需要被统计文本相同字符前后在一起，所以先排序 

#uniq -c 表示同时输出出现次数

#sort -nr 其中-n表示把数字当做真正的数字处理(当数字被当做字符串处理，会出现11比2小的情况)
```



```text
cat words.txt |tr -s ' ' '\n' |sort|uniq -c|sort -r|awk '{print $2,$1}'

# 1、首先cat命令查看words.txt
# 2、tr -s ' ' '\n'将空格都替换为换行 实现分词
# 3、sort排序 将分好的词按照顺序排序
# 4、uniq -c 统计重复次数（此步骤与上一步息息相关，-c原理是字符串相同则加一，如果不进行先排序的话将无法统计数目）
# 5、sort -r 将数目倒序排列
# 6、awk '{print $2,$1}' 将词频和词语调换位置打印出来
```

## **Match Character Set: \[...\]**

```text
$ echo d[aeiou]g
```

```text
$ echo d[a-o]g
```

## **Match Inverse Character Set: \[^...\] or \[!...\]**

To find matches of the form`d?g` but with no vowel:-

```text
$ echo d[^aeiou]g
d0g
```

```text
$ echo report[!1–3].txt
report4.txt report5.txt
```

 match all text files that _don’t_ end with a number, this pattern does the job:-

```text
$ echo *[^1-9].txt
index.txt report.txt
```

## **Brace Expansion: {...}**

```text
$ echo d{a,e,i,u,o}g
dag deg dig dug dog
```



```text
$ echo d{a..z}g
dag dbg dcg ddg deg dfg dgg dhg dig djg dkg dlg dmg dng dog dpg dqg drg dsg dtg dug dvg dwg dxg dyg dzg
```

## create file 

```text
$ touch report{6..20}.txt
```

## **Alternation: {a,b}**

```text
$ echo {cat,d*}
cat dawg dg dig dog doug dug
```

```text
$ ls {*.jpg,*.jpeg,*.png}
family1.jpg family2.jpeg holiday1.jpg holiday2.jpg
family1.png family3.jpeg holiday1.png
```

```text
$ ls *.{j{p,pe}g,png}
family1.jpg family2.jpeg holiday1.jpg holiday2.jpg
family1.png family3.jpeg holiday1.png
```

```text
$ echo {j{p,pe}g,png}
jpg jpeg png
```

```text
$ echo .{mp{3..4},m4{a,b,p,v}}
.mp3 .mp4 .m4a .m4b .m4p .m4v
```

