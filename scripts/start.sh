#!/bin/sh

set -e
set -x

FILEDIR=$(cd `dirname $0` && pwd -P)

cd $FILEDIR/../

pwd

supervisorctl reload
