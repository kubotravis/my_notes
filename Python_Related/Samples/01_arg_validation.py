#!/usr/bin/env python
import sys

# Usage function
def usage():
    print('Please supply atleast one ARG to this script')
    sys.exit(1)

# Get the total no.of arg passed
arglen=len(sys.argv)-1

if arglen == 0:
    usage()
else:
    print("you have supplied totally"), arglen, ("arguments")
    print("Let's loop thought the each arguments")
    count = 1
    while count <= arglen:
        print("The"), count,("ARG is"), sys.argv[count]
        count = count+1

print("---END---")
