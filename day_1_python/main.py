from dataclasses import dataclass
from typing import List

@dataclass
class Rotation:
    direction: str
    clicks: int


def parse_rotation(raw_rot: str) -> Rotation:
    direction = raw_rot[0]
    clicks = int(raw_rot[1:])
    return Rotation(direction, clicks)

def parse_puzzle_input(raw_input: str) -> List[Rotation]:
    return [parse_rotation(line) for line in raw_input.splitlines()]

def main():
    print("Hello from day-1-python!")


def rotateDial(dialPos, rotation: Rotation) -> Tuple[int, int]:
    newDialPos = dialPos

    # number of times the clicker has hit zero
    zeroPasses = 0

    for i in range(abs(rotation.clicks)):
        if newDialPos == 0 and i != 0:
            zeroPasses += 1
        if rotation.direction == "L":
            newDialPos -= 1
        elif rotation.direction == "R":
            newDialPos += 1
        
        if newDialPos < 0:
            newDialPos = 99
        elif newDialPos > 99:
            newDialPos = 0

    return (newDialPos, zeroPasses)

if __name__ == "__main__":
    with open("input.txt", "r") as file:
        raw_input = file.read()
        rotations = parse_puzzle_input(raw_input)
        dialPos = 50
        zeroCount = 0
        for rotation in rotations:
            dialPos, zeroPasses = rotateDial(dialPos, rotation)
            zeroCount += zeroPasses
            if dialPos == 0:
                zeroCount += 1
    print(f"Zero count: {zeroCount}")




assert rotateDial(50, Rotation("L", 1)) == (49, 0)  
assert rotateDial(50, Rotation("R", 1)) == (51, 0)
assert rotateDial(50, Rotation("L", 10)) == (40, 0)
assert rotateDial(50, Rotation("R", 10)) == (60, 0)
assert rotateDial(50, Rotation("L", 100)) == (50, 1)
assert rotateDial(50, Rotation("R", 100)) == (50, 1)
assert rotateDial(50, Rotation("L", 101)) == (49, 1)
assert rotateDial(50, Rotation("R", 101)) == (51, 1)
assert rotateDial(50, Rotation("L", 102)) == (48, 1)
assert rotateDial(50, Rotation("R", 102)) == (52, 1)
assert rotateDial(50, Rotation("R", 1000)) == (50, 10)

assert rotateDial(99, Rotation("R", 1)) == (0, 0)
assert rotateDial(0, Rotation("L", 1)) == (99, 0)