import sys
from pathlib import Path
from typing import List

def file_parse(path: str) -> str:
    """Read the entire content of a file as a string."""
    try:
        return Path(path).read_text()
    except Exception as e:
        sys.exit(f"Error in file_parse(): {e}")

def line_parse(data: str) -> List[str]:
    """Split a string into lines, removing any trailing newline."""
    return data.rstrip("\n").split("\n")

def grid_parse(data: str) -> List[List[str]]:
    """Convert string data into a 2D grid of characters."""
    lines = line_parse(data)
    return [list(line) for line in lines]


