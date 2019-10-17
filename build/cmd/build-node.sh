#!/usr/bin/env bash

. $(dirname $0)/common.sh
#echo ${BasePath}
#echo ${CMD}
#echo ${Hour}

VERSION=$(genVersion $1)

folder="${BasePath}/out/plugins"
if [[ ! -d "$folder" ]]
then
  ${CMD}/build-plugin.sh
  if [[ "$?" != "0" ]]
    then exit 1
  fi
fi


buildApp node $VERSION


OUTPATH="${BasePath}/out/node-${VERSION}"
mkdir ${OUTPATH}/plugin
cp -a ${BasePath}/build/node/resources/*  ${OUTPATH}/
if [ -d "${BasePath}/out/plugins" ];then 
    cp -a ${BasePath}/out/plugins/*  ${OUTPATH}/plugin/
fi

chmod +x ${OUTPATH}/install.sh ${OUTPATH}/run.sh

cd ${ORGPATH}
