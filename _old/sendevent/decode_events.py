#!/usr/bin/python
# coding: utf-8
import re
import sys

# [   43319.628481] /dev/input/event1: 0003 0039 ffffffff
# 48470-342082: /dev/input/event1: 0000 0000 00000000
_re = re.compile(r'[^\d]*(?P<sec>\d+)[.-](?P<msec>\d+)[:\]] (?P<device>[^:]+):'
' (?P<class>[0-9a-f]+) (?P<event>[0-9a-f]+) (?P<params>[0-9a-f]+)')
T_FIX = 0.1

last_time = None

if __name__ == '__main__':
    if len(sys.argv) < 2:
        print USAGE
        sys.exit(1)
    
    print '#!/bin/sh'
    input_fn = sys.argv[1]
    for line in open(input_fn, 'rt'):
        m = _re.match(line)
        if m is not None:
            d = m.groupdict()
            cur_time = float(d['sec']) + float(d['msec'][:2])/100
            if last_time is not None:
                diff_time = (cur_time - last_time)
                if diff_time > 0.2:
                    print 'sleep %.2f' % (diff_time-T_FIX,)
            last_time = cur_time
            print 'sendevent', d['device'], int(d['class'], 16), \
                int(d['event'], 16), int(d['params'], 16)
        else:
            print '#', line.strip('\n\r\t ')
