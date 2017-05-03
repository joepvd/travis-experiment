#!/bin/bash

ldd `which curl` | awk '$3~/^\//{print "dpkg -S", $3}'  | sh | awk -F: '{print "dpkg -s", $1}' | sh |  awk '$1=="Package:"{p=$2};$1=="Version:"{print p, $0}'
