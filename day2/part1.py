def runProgram(data, i):
    print(i)
    if i >= len(data) or data[i] == 99:
        return data

    num = data[i]
    if num == 1 or num == 2:
        num1 = data[data[i + 1]]
        num2 = data[data[i + 2]]
        pos3 = data[i + 3]
        print("Num1: " + str(num1) + " - Num2: " + str(num2) + " - Pos: " + str(pos3))
        if num == 1:
            data[pos3] = num1 + num2
        if num == 2:
            data[pos3] = num1 * num2

    return runProgram(data, i + 4)


f = open("input.txt", "r")

data = []
for line in f:
    data = line.split(",")

for i in range(0, len(data)):
    data[i] = int(data[i])


data[1] = 12
data[2] = 2
print(len(data))
data = runProgram(data, 0)
print(data)
print("Value at position 0: " + str(data[0]))
