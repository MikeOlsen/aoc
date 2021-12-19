import aocd

prod_data = aocd.get_data(day=3, year=2021).splitlines()

# fmt: off
test_data = [
    "00100", "11110", "10110", "10111", "10101", "01111",
    "00111", "11100", "10000", "11001", "00010", "01010",
]
# fmt: on


def most_common_bit(list, position):
    zero = 0
    one = 0
    for item in list:
        if int(item[position]) == 0:
            zero += 1
        else:
            one += 1

    return 0 if zero > one else 1


def most_common_bit_list(list):
    bits = len(list[0])
    values = []
    for i in range(0, bits):
        values.append(most_common_bit(list, i))
    return values


def filter_list(list, position, value):
    bits = []
    for item in list:
        if int(item[position]) == value:
            bits.append(item)
    return bits


def invert(list):
    values = []
    for value in list:
        values.append(0 if value == 1 else 1)
    return values


def decimal(bit_list):
    return int("".join([str(bit) for bit in bit_list]), 2)


def get_oxygen_reading(list):
    bits = len(list[0])
    for position in range(0, bits):
        if len(list) == 1:
            break
        bit = most_common_bit(list, position)
        list = filter_list(list, position, bit)
    return decimal(list)


def get_co2_reading(list):
    bits = len(list[0])
    for position in range(0, bits):
        if len(list) == 1:
            break
        bit = 1 if most_common_bit(list, position) == 0 else 0
        list = filter_list(list, position, bit)
    return decimal(list)


def air_report(data):
    o2 = get_oxygen_reading(data)
    co2 = get_co2_reading(data)
    return {
        "oxygen": o2,
        "co2": co2,
        "rating": o2 * co2,
    }


def diagnostics_report(data):
    mcbl = most_common_bit_list(data)
    return {
        "gamma": decimal(mcbl),
        "epsilon": decimal(invert(mcbl)),
        "power": decimal(mcbl) * decimal(invert(mcbl)),
    }


test_report = diagnostics_report(test_data)
print()
print("=============== PART 1 ==================")
print("[TEST] Gamma:  ", test_report["gamma"])
print("[TEST] Epsilon:", test_report["epsilon"])
print("-----------------------------------------")
print("[TEST] Power Consumption:", test_report["power"])

prod_report = diagnostics_report(prod_data)
print("=========================================")
print("[PROD] Gamma:  ", prod_report["gamma"])
print("[PROD] Epsilon:", prod_report["epsilon"])
print("-----------------------------------------")
print("[PROD] Power Consumption:", prod_report["power"])
print("=========================================")

test_air_report = air_report(test_data)
print()
print("=============== PART 2 ==================")
print("[TEST] Oxygen: ", test_air_report["oxygen"])
print("[TEST] CO2:    ", test_air_report["co2"])
print("-----------------------------------------")
print("[TEST] Rating: ", test_air_report["rating"])

prod_air_report = air_report(prod_data)
print()
print("=========================================")
print("[PROD] Oxygen: ", prod_air_report["oxygen"])
print("[PROD] CO2:    ", prod_air_report["co2"])
print("-----------------------------------------")
print("[PROD] Quality rating: ", prod_air_report["rating"])
