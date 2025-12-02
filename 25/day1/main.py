from utils import * 


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

# def partTwo(text_input: str):
#     print()
#     lines = line_parse(text_input)
#     cur = 50 
#     clicks = 0
#
#     for line in lines:
#         sign = -1 if line[0] == 'L' else 1
#         num = int(line[1:])
#         next = (cur + sign*num) % 100
#
#         if sign == -1:
#             clicks += ((100-cur)%100 + num)//100
#         else:
#             clicks += (cur + num)//100
#
#         # print(next, clicks)
#
#         cur = next
#
#
#
#     return clicks

def partTwo(text_input: str):
    print()
    lines = line_parse(text_input)
    cur = 50 
    clicks = 0

    for line in lines:
        sign = -1 if line[0] == 'L' else 1
        num = int(line[1:])
        next = (cur + sign*num)

        if next == 0:
            clicks += 1

        while next > 99:
            if next == 0:
                clicks += 1

            next -= 100

            clicks += 1
        while next < 0:
            next += 100

            if next == 0:
                clicks += 1


            if cur == 0:
                cur = next
            else:
                clicks += 1

        # print(line, next, clicks)


        cur = next

    return clicks
