import aocd

with open("data/05.txt", "r") as file:
    test_data = file.read().splitlines()


prod_data = aocd.get_data(day=5, year=2021).splitlines()


class Point:
    def __init__(self, x, y):
        self.x = x
        self.y = y

    def __repr__(self):
        return f"{self.x},{self.y}"


def init_matrix(x, y, value):
    return [[value] * x for _ in range(y)]


def parse_point(data):
    parts = data.split(",")
    return Point(int(parts[0]), int(parts[1]))


def parse_lines(data):

    lines = []
    for line in data:
        line = line.replace(" ", "")
        parts = line.split("->")

        p1 = parse_point(parts[0])
        p2 = parse_point(parts[1])

        lines.append((p1, p2))

    return lines


def draw_line(matrix, p1, p2):
    if p1.y == p2.y:
        for x in range(min(p1.x, p2.x), max(p1.x, p2.x) + 1):
            matrix[p1.y][x] += 1
    elif p1.x == p2.x:
        for y in range(min(p1.y, p2.y), max(p1.y, p2.y) + 1):
            matrix[y][p1.x] += 1
    return matrix


def count_overlaps(matrix, threshold):
    count = 0
    for row in matrix:
        for col in row:
            if col >= threshold:
                count += 1
    return count


def solve_matrix(m, n, data):
    matrix = init_matrix(m, n, 0)
    lines = parse_lines(data)
    for line in lines:
        draw_line(matrix, line[0], line[1])
    return count_overlaps(matrix, 2)


print("=============== PART 1 ==================")
print("[TEST] Overlaps:", solve_matrix(10, 10, test_data))
print("[PROD] Overlaps:", solve_matrix(1000, 1000, prod_data))
