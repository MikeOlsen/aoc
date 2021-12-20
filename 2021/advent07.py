import aocd

test_data = [16, 1, 2, 0, 4, 2, 7, 1, 2, 14]
prod_data = [int(v) for v in aocd.get_data(day=7, year=2021).split(",")]


def distance(number, point):
    return abs(number - point)


def consumption_report(data):
    dist = [0] * max(data)
    for x in range(max(data)):
        for crab in data:
            dist[x] += distance(x, crab)

    return dist


print("=============== PART 1 ==================")
report = consumption_report(test_data)
print("[TEST]", report)
print("[TEST] Cheapest outcome:", min(report))
report = consumption_report(prod_data)
print("[PROD] Cheapest outcome:", min(report))
