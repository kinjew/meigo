#!/bin/sh

set -e
set -x

ps aux | grep meigo

supervisorctl stop meigo
