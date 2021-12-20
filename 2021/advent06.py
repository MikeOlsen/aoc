import aocd

from tabulate import tabulate

test_data = list = [3, 4, 3, 1, 2]
prod_data = [int(v) for v in aocd.get_data(day=6, year=2021).split(",")]

INIT_LENGTH = 8
RESET_LENGTH = 6


class Swarm:
    def __init__(self, list=[]):
        self.clock = [0] * (INIT_LENGTH + 1)
        for v in list:
            self.clock[v + 1] += 1

    def tick(self):
        reset = self.clock[0]
        self.clock[0] = 0

        # Rotate left
        self.clock = self.clock[1:] + self.clock[:1]

        self.clock[8] += reset
        self.clock[6] += reset

    def count(self):
        return sum(self.clock)

    def __repr__(self):
        days = ["Clock"]
        days.extend([str(_) for _ in range(1, INIT_LENGTH + 1)])
        return tabulate([self.clock], headers=days)


def simulate(data, days, verbose=False):
    swarm = Swarm(data)

    for i in range(days + 1):
        swarm.tick()

        if verbose == 1:
            print("Day:", i)
            print(swarm, end="\n\n")

    return swarm.count()


if __name__ == "__main__":

    print("\n=============== PART 1 ==================")
    print(
        f"------------------------------------------\n"
        f"[TEST] Total: {simulate(test_data, 18, verbose=True)}"
    )
    print("[TEST] Total:", simulate(test_data, 80, verbose=False))
    print("-----------------------------------------")
    print("[PROD] Total:", simulate(prod_data, 80, verbose=False))
    print("-----------------------------------------")

    print("\n=============== PART 2 ==================")
    print(f"[PROD] Total: {simulate(prod_data, 256, verbose=False):,}")
    print("-----------------------------------------")
