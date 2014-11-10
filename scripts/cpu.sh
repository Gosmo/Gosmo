#!/bin/sh
# Note this script does not work properly
# unless you let it run a few rounds.

PREV_TOTAL=0
PREV_IDLE=0

counter=0
limitsecs=10

# Sum all the CPU usage percentages and then calc
# the average in that time. Silly thing to do.
# Will be changed to continues data spewing,
# this is for demonstrative purposes.
total=0

while (( counter < limitsecs )); do
    CPU=$(awk '/^cpu / {print substr($0, index($0,$2))}' /proc/stat)
    IDLE=$(echo $CPU|awk '{print $4}')
    # Calculate the total CPU time.
    FUN_TOTAL=$(echo $CPU|awk 'func sum(s) {
                                split(s,a);
                                r=0;
                                for(i in a) {
                                    r+=a[i];
                                }
                                return r}
                                {print sum($0)}')
    # Calculate the CPU usage since we last checked.
    (( DIFF_IDLE  = IDLE - PREV_IDLE ))
    (( DIFF_TOTAL =  FUN_TOTAL -  PREV_TOTAL ))
    (( DIFF_USAGE = (1000 * (DIFF_TOTAL - DIFF_IDLE) / DIFF_TOTAL + 5) / 10 ))
    # printf "\rCPU: %d%% \b\b" $DIFF_USAGE
    (( total += DIFF_USAGE ))
    
    # Remember the total and idle CPU times for the next check.
    PREV_TOTAL="$FUN_TOTAL"
    PREV_IDLE="$IDLE"

    (( counter++ ))
    sleep 1
done

(( total /= limitsecs ))
printf "%d" $total
