#!/usr/bin/env bash
# Minor utility that prints memory usage.

data=$(awk 'NR' /proc/meminfo)

mt=$(echo -n "$data"|awk '/MemTotal/ {print $2}')
mb_f=$(echo -n "$data"|awk '/MemFree/ {print $2}')
mb_b=$(echo -n "$data"|awk '/Buffers/ {print $2}')
mb_c=$(echo -n "$data"|awk '/^Cached/ {print $2}')
#st=$(echo -n "$data"|awk '/SwapTotal/ {print $2}')
#sf=$(echo -n "$data"|awk '/SwapFree/ {print $2}')

let memfree=$mb_f+$mb_b+$mb_c
let inuse=$mt-$memfree
let usep=$inuse/$mt*100

##awk 'BEGIN { printf "%.0f\n", 10/6 }'

let rmt=$mt/1024
let rmb_f=$mb_f/1024
let rmb_b=$mb_b/1024
let rmb_c=$mb_c/1024
let rmfree=$rmb_f+$rmb_b+$rmb_c
let rinuse=$rmt-$rmfree
let rusep=$rinuse/$rmt*100

echo -n "MEM: ${rinuse}M"

