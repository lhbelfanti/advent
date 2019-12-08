def calculate_fuel(fuel, t):
    pf = (int(fuel) // 3) - 2
    if pf <= 0:
        return t
    return calculate_fuel(pf, t + pf)


def main():
    f = open("input.txt", "r")

    total = 0
    for line in f:
        total += calculate_fuel(int(line), 0)

    print("Total: " + str(total))


main()
