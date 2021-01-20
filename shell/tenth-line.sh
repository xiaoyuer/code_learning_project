sed -n "10p" file.txt
grep -n "" file.txt | grep -w '10' | cut -d: -f2
awk '{if(NR==10){print $0}}' file.txt # awk 条件 + 文件