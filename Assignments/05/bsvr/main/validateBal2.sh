#!/bin/bash

#
# 3 argumens
#
# $1 - file with before balance
# $2 - change to balance - up or down
# $3 - balance fiel after - do calculation to see if it adds up.
#

bef=$(cat $1 | awk '{print $4}')
total=$2
if (( $bef == $total )) ; then
	# echo "All is good"
	:
else
	# echo "Nope"
	echo "FAIL - Looking for total of $total, got $bef"
	exit 1
fi

