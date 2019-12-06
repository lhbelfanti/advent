def calculateFuel(f, t):
    pf = (int(f) // 3) - 2
    if pf <= 0:
        return t
    return calculateFuel(pf, t + pf)


f = open("input.txt", "r")

total = 0
for line in f:
    total += calculateFuel(int(line), 0)

print("Total: " + str(total))


