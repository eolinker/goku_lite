#!/usr/bin/env bash

echo $0
. $(dirname $0)/common.sh

#echo ${BasePath}
#echo ${CMD}
#echo ${Hour}

VERSION=$(genVersion $1)
OUTPATH="${BasePath}/out/console-${VERSION}"
buildApp console $VERSION

mkdir  ${OUTPATH}/static
#cp -a ${BasePath}/app/console/static/*  ${OUTPATH}/static/
cp -a ${BasePath}/build/console/resources/*  ${OUTPATH}/

chmod +x ${OUTPATH}/install.sh ${OUTPATH}/run.sh
cd ${ORGPATH}
