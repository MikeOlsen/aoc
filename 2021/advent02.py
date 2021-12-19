import aocd

test_data = ["forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"]
prod_data = aocd.get_data(day=2, year=2021).splitlines()

# fmt: off
walk_mapping = {
    "forward":  {"x": 1, "y":  0},
    "down":     {"x": 0, "y":  1},
    "up":       {"x": 0, "y": -1},
}
# fmt: on


def walk(directions, verbose=False):
    x = 0
    y = 0
    for item in directions:
        parts = item.split(" ")
        direction = parts[0]
        distance = parts[1]

        _x = walk_mapping[direction]["x"] * int(distance)
        _y = walk_mapping[direction]["y"] * int(distance)

        if verbose:
            print(f"{item}    \t -> \t x: {_x}, y: {_y}")

        x += _x
        y += _y

    return (x, y, x * y)


# fmt: off
steer_mapping = {
    "forward": {"x": 1, "y": 1, "a":  0},
    "down":    {"x": 0, "y": 0, "a":  1},
    "up":      {"x": 0, "y": 0, "a": -1},
}
# fmt: on


def steer(directions, verbose=False):
    x = 0
    y = 0
    a = 0
    for item in directions:
        parts = item.split(" ")
        direction = parts[0]
        distance = parts[1]

        _x = steer_mapping[direction]["x"] * int(distance)
        _y = steer_mapping[direction]["y"] * int(distance) * a
        _a = steer_mapping[direction]["a"] * int(distance)

        if verbose:
            print(  # Print calculation steps
                f"{item}    \t-> \t"
                f"x: {_x}, y: {_y}, a: {_a}\t->\t"
                f"x: {x+_x}, y: {y+_y}, a: {a+_a}"
            )

        x += _x
        y += _y
        a += _a

    return (x, y, x * y)


# First part
print("=============== PART 1 ==================")

test = walk(test_data, verbose=True)
prod = walk(prod_data, verbose=False)

print()
print(f"[TEST]: x: {test[0]},  \t y: {test[1]},  \t sum: {test[2]}")
print(f"[PROD]: x: {prod[0]},\t y: {prod[1]},\t sum: {prod[2]}")


# Second part
print("=============== PART 2 ==================")

test = steer(test_data, verbose=True)
prod = steer(prod_data, verbose=False)

print()
print(f"[TEST]: x: {test[0]},  \t y: {test[1]},  \t sum: {test[2]}")
print(f"[PROD]: x: {prod[0]},\t y: {prod[1]},\t sum: {prod[2]}")
