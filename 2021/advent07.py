import aocd

test_data = [16, 1, 2, 0, 4, 2, 7, 1, 2, 14]
prod_data = [int(v) for v in aocd.get_data(day=7, year=2021).split(",")]


def drive(start, end, exponential):
    if exponential:
        return sum([*range(abs(start - end) + 1)])
    else:
        return abs(start - end)


def consumption_report(data, exponential=False):
    dist = [0] * max(data)
    for x in range(max(data)):
        for crab in data:
            dist[x] += drive(x, crab, exponential)
    return dist


print("=============== PART 1 ==================")
report = consumption_report(test_data)
print("[TEST]", report)
print("[TEST] Cheapest outcome:", min(report))
report = consumption_report(prod_data)
print("[PROD] Cheapest outcome:", min(report))

print("=============== PART 2 ==================")
report = consumption_report(test_data, exponential=True)
print("[TEST]", report)
print("[TEST] Cheapest outcome:", min(report))
report = consumption_report(prod_data, exponential=True)
print("[PROD] Cheapest outcome:", min(report))
