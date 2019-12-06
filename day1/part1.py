f = open("input.txt", "r")

total = 0
for line in f:
    total += (int(line) // 3) - 2

print("Total: " + str(total))
