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