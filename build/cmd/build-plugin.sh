#!/usr/bin/env bash
. $(dirname $0)/common.sh

if [ $# != 0 ] ; then
 for i in $*               #在$*中遍历参数，此时每个参数都是独立的，会遍历$#次
do
    buildPlugin $i
    if [[ "$?" != "0" ]]
        then exit 1
    fi
done
else
    if [ ! -d "${BasePath}/app/plugins/" ];then
    	exit 0
    fi
    for i in $(ls ${BasePath}/app/plugins/)
    do
        buildPlugin $i
        if [[ "$?" != "0" ]]
        then
            exit 1
        fi
   done
fi

exit 0
