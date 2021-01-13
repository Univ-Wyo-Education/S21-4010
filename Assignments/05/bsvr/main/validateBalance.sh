#!/bin/bash

#
# 3 argumens
#
# $1 - file with before balance
# $2 - change to balance - up or down
# $3 - balance fiel after - do calculation to see if it adds up.
#

bef=$(cat $1 | awk '{print $4}')
aft=$(cat $3 | awk '{print $4}')
delta=$2
let "calc=$bef+($delta)"
if (( $calc == $aft )) ; then
	# echo "All is good"
	:
else
	# echo "Nope"
	echo "FAIL - Expected $aft - got $bef + ($delta), values should match, they did not."
	exit 1
fi

