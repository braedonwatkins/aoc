from utils import * 

input = file_parse("test.txt")

def partOne(input: str):
    print() #silly pytest, newlines are for kids!
    lines = line_parse(input) 

    sign = None
    cur = 50
    total = 0

    # loop each line
    for line in lines:
        ## split first char 
        if line[0] == 'L':
            sign = -1
        elif line[0] == 'R':
            sign = 1
        else:
            raise Exception("halt and catch fire")

        ## if L subtract, if R add, mod 100
        num = int(line[1:])
        cur += sign*num
        cur = cur % 100

        if cur == 0:
            total += 1

    return total




    return 3

def partTwo(input: str):
    print()

    return -1

