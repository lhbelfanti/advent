def run_program(dat, idx, inp):
    if idx >= len(dat):
        return dat

    num = dat[idx]

    digits = [int(x) for x in str(num)]
    modes = [0, 0, 0]
    if len(digits) == 5:
        num = int(str(digits[3]) + str(digits[4]))
        modes = [digits[2], digits[1], digits[0]]
    elif len(digits) == 4:
        num = int(str(digits[2]) + str(digits[3]))
        modes = [digits[1], digits[0], 0]
    elif len(digits) == 3:
        num = int(str(digits[1]) + str(digits[2]))
        modes = [digits[0], 0, 0]

    if num == 1:
        num1 = dat[dat[idx + 1]] if modes[0] == 0 else dat[idx + 1]
        num2 = dat[dat[idx + 2]] if modes[1] == 0 else dat[idx + 2]
        pos3 = dat[idx + 3]
        dat[pos3] = num1 + num2
        return run_program(dat, idx + 4, inp)
    elif num == 2:
        num1 = dat[dat[idx + 1]] if modes[0] == 0 else dat[idx + 1]
        num2 = dat[dat[idx + 2]] if modes[1] == 0 else dat[idx + 2]
        pos3 = dat[idx + 3]
        dat[pos3] = num1 * num2
        return run_program(dat, idx + 4, inp)
    elif num == 3:
        pos1 = dat[idx + 1]
        dat[pos1] = inp
        return run_program(dat, idx + 2, inp)
    elif num == 4:
        pos1 = dat[dat[idx + 1]] if modes[0] == 0 else dat[idx + 1]
        print(pos1)
        return run_program(dat, idx + 2, inp)
    elif num == 5:
        num1 = dat[dat[idx + 1]] if modes[0] == 0 else dat[idx + 1]
        num2 = dat[dat[idx + 2]] if modes[1] == 0 else dat[idx + 2]
        if num1 != 0:
            return run_program(dat, num2, inp)
        return run_program(dat, idx + 3, inp)
    elif num == 6:
        num1 = dat[dat[idx + 1]] if modes[0] == 0 else dat[idx + 1]
        num2 = dat[dat[idx + 2]] if modes[1] == 0 else dat[idx + 2]
        if num1 == 0:
            return run_program(dat, num2, inp)
        return run_program(dat, idx + 3, inp)
    elif num == 7:
        num1 = dat[dat[idx + 1]] if modes[0] == 0 else dat[idx + 1]
        num2 = dat[dat[idx + 2]] if modes[1] == 0 else dat[idx + 2]
        pos3 = dat[idx + 3]
        if num1 < num2:
            dat[pos3] = 1
        else:
            dat[pos3] = 0
        return run_program(dat, idx + 4, inp)
    elif num == 8:
        num1 = dat[dat[idx + 1]] if modes[0] == 0 else dat[idx + 1]
        num2 = dat[dat[idx + 2]] if modes[1] == 0 else dat[idx + 2]
        pos3 = dat[idx + 3]
        if num1 == num2:
            dat[pos3] = 1
        else:
            dat[pos3] = 0
        return run_program(dat, idx + 4, inp)
    elif num == 99:
        return dat

    return dat


def main():
    f = open("input.txt", "r")

    data = []
    for line in f:
        data = line.split(",")

    for i in range(0, len(data)):
        data[i] = int(data[i])

    input_value = 5
    run_program(data, 0, input_value)


main()
