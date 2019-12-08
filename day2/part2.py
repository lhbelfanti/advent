def run_program(data, i):
    if i >= len(data) or data[i] == 99:
        return data

    num = data[i]
    if num == 1 or num == 2:
        num1 = data[data[i + 1]]
        num2 = data[data[i + 2]]
        pos3 = data[i + 3]
        if num == 1:
            data[pos3] = num1 + num2
        if num == 2:
            data[pos3] = num1 * num2

    return run_program(data, i + 4)


def new_program(noun, verb):
    f = open("input.txt", "r")

    data = []
    for line in f:
        data = line.split(",")

    for i in range(0, len(data)):
        data[i] = int(data[i])

    data[1] = noun
    data[2] = verb
    data = run_program(data, 0)
    return data[0]


def main():
    for i in range(100):
        for j in range(100):
            value = new_program(i, j)
            if value == 19690720:
                print("The noun is: " + str(i) + " and the verb is: " + str(j))
                print("The answer is: " + str((100 * i) + j))


main()
