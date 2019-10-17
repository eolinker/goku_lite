#!/usr/bin/env bash
. $(dirname $0)/common.sh

cd ${BasePath}/


VERSION=$(genVersion $1)
folder="${BasePath}/out/node-${VERSION}"
if [[ ! -d "$folder" ]]
then

  ${CMD}/build-node.sh $1
  if [[ "$?" != "0" ]]
  then
    exit 1
  fi
fi
packageApp node $VERSION

cd ${ORGPATH}
