from day1.main import *
from utils.main import *

def test_one():
    input_data = file_parse("test.txt")  
    expected = 3
    result = partOne(input_data)
    assert result == expected, f"expected result was {expected}, but got {result} instead"

def test_two():
    input_data = file_parse("test.txt")
    expected = -1
    result = partTwo(input_data)
    assert result == expected, f"expected result was {expected}, but got {result} instead"


