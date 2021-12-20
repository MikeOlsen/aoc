import aocd

test_data = list = [3, 4, 3, 1, 2]
prod_data = [int(v) for v in aocd.get_data(day=6, year=2021).split(",")]

INIT_LENGTH = 8
RESET_LENGTH = 6


class Fish:
    def __init__(self, clock=INIT_LENGTH):
        self.timer = clock

    def tick(self, rate=1):
        self.timer -= rate

        if self.timer < 0:
            self.timer = RESET_LENGTH
            return Fish()

    def __repr__(self):
        return str(self.timer)


class Swarm:
    def __init__(self, list=[]):
        self.fishes = []
        for item in list:
            self.fishes.append(Fish(item))

    def append(self, fish):
        if fish:
            self.fishes.append(fish)

    def extend(self, children):
        self.fishes.extend(children.fishes)

    def __repr__(self):
        return ", ".join(str(v) for v in self.fishes)


def simulate(data, days, verbose=1):
    swarm = Swarm(data)

    for i in range(days):
        children = Swarm()
        for fish in swarm.fishes:
            children.append(fish.tick())
        swarm.extend(children)

        if verbose == 1:
            print(f"Day {i+1}:\t {len(swarm.fishes)} -> {len(children.fishes)}")
        elif verbose >= 2:
            print(
                f"Day {i+1}:\t {swarm} ->"
                f" total: {len(swarm.fishes)}"
                f" new: {len(children.fishes)}"
            )

    return len(swarm.fishes)


if __name__ == "__main__":

    print()
    print("=============== PART 1 ==================")
    print("[TEST] Total:", simulate(test_data, 18, verbose=2))
    print("-----------------------------------------")
    print("[TEST] Total:", simulate(test_data, 80, verbose=1))
    print("-----------------------------------------")
    print("[PROD] Total:", simulate(prod_data, 80, verbose=1))
    print("-----------------------------------------")
