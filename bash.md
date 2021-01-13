# bash

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

