import aocd


with open("data/08.txt", "r") as file:
    test_data = file.read().splitlines()

prod_data = aocd.get_data(day=8, year=2021).splitlines()

DIGITS = {
    2: 1,
    3: 7,
    4: 4,
    7: 8,
}


def identify_digit(digit):
    unique_letters = len(list(set(digit)))
    d = DIGITS.get(unique_letters, None)
    return d


def go(data):
    digits = []
    for line in data:
        parts = line.split("|")
        for digit in parts[1].strip().split(" "):
            d = identify_digit(digit)
            if d:
                digits.append(d)

    return len(digits)


print("=============== PART 1 ==================")
print("[TEST]", go(test_data))
print("[PROD]", go(prod_data))
