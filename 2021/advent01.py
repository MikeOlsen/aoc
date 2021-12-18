"""
As the submarine drops below the surface of the ocean,
it automatically performs a sonar sweep of the nearby sea floor.
On a small screen, the sonar sweep report (your puzzle input) appears:
each line is a measurement of the sea floor depth as the sweep looks
further and further away from the submarine.
"""
import aocd

test_data = [199, 200, 208, 210, 200, 207, 240, 269, 260, 263]
prod_data = [int(i) for i in aocd.get_data(day=1, year=2021).splitlines()]


def sliding_window(iterable, window_size):
    for i in range(len(iterable) - window_size + 1):
        yield iterable[i : i + window_size]


# First part
def count_increase(measurements, verbose=False):
    count = 0
    for window in sliding_window(measurements, 2):
        if window[0] < window[1]:
            count += 1
            if verbose:
                print(f"{window} (+)")
        elif verbose:
            if window[0] > window[1]:
                print(f"{window} (-)")
            else:
                print(f"{window} (unchanged)")
    return count


print("=============== PART 1 ==================")
print("Test:", count_increase(test_data, verbose=True))
print("Prod:", count_increase(prod_data, verbose=False))

# Second part
def count_window_increase(measurements, window_size, verbose=False):
    count = 0
    previous = sum(measurements[0:window_size])
    for window in sliding_window(measurements, window_size):
        current = sum(window)
        if current > previous:
            count += 1
            if verbose:
                print(f"{window} -> {current} (+)")
        elif verbose:
            if current < previous:
                print(f"{window} -> {current} (-)")
            else:
                print(f"{window} -> {current} (no change)")
        previous = current
    return count


print("=============== PART 2 ==================")
print("Test:", count_window_increase(test_data, 3, verbose=True))
print("Prod:", count_window_increase(prod_data, 3, verbose=False))
