def run_program(dat, idx):
    if idx >= len(dat) or dat[idx] == 99:
        return dat

    num = dat[idx]
    if num == 1 or num == 2:
        num1 = dat[dat[idx + 1]]
        num2 = dat[dat[idx + 2]]
        pos3 = dat[idx + 3]
        if num == 1:
            dat[pos3] = num1 + num2
        if num == 2:
            dat[pos3] = num1 * num2

    return run_program(dat, idx + 4)


def main():
    f = open("input.txt", "r")

    data = []
    for line in f:
        data = line.split(",")

    for i in range(0, len(data)):
        data[i] = int(data[i])

    data[1] = 12
    data[2] = 2
    data = run_program(data, 0)
    print("Value at position 0: " + str(data[0]))


main()
