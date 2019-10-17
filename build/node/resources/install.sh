#!/usr/bin/env bash
cd  $(dirname $0) # 当前位置跳到脚本位置

APP="node"
SOURCE_DIR="$(pwd)"
PRO_DIR="$(basename `pwd`)"
cd ../
ORG_DIR="$(pwd)"
INSTALL_DIR=$1
if [[ "$INSTALL_DIR" = "" ]]
then
    INSTALL_DIR=$(pwd)
else
    if [[ ! -d "$INSTALL_DIR" ]]
    then
      mkdir -p $INSTALL_DIR
      cd $INSTALL_DIR
      INSTALL_DIR="$(pwd)"
      cd $ORG_DIR
    fi
fi

if [[ "$SOURCE_DIR" = "$INSTALL_DIR" ]]
then
    echo "无效的安装目录"
    exit 1
fi



if [[ "$ORG_DIR" != "$INSTALL_DIR" ]]
then
    if [[ -d "$INSTALL_DIR/$PRO_DIR" ]]
    then
        rm -rf "$INSTALL_DIR/$PRO_DIR"
    fi
    cp -a $PRO_DIR "$INSTALL_DIR/$PRO_DIR"
fi

cd $INSTALL_DIR

if [[ -L "$INSTALL_DIR/$APP" ]]
then
    rm -f "$INSTALL_DIR/$APP"
fi

ln -sf $PRO_DIR ./$APP



mkdir -p work

ln -sf $INSTALL_DIR/work $PRO_DIR/

cd  $PRO_DIR
