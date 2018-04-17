#!/bin/sh

NAME=(
"gateway.go"
"monitor_redis"
)
for process in ${NAME[@]}
do
        echo $process
        echo "---------------"   
        ID=`ps -ef|grep $process|grep -v grep|grep -v PPID|awk '{print $2}'`  
        echo $ID    
        for id in $ID  
        do  
                kill -9 $id  
                echo "killed $id"  
        done  
        echo "---------------"
done

nohup go run gateway.go > gateway.log 2>&1 &
nohup go run monitor_redis.go > monitor_redis.log 2>&1 &