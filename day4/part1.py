def process_num(n):
    digits = [int(x) for x in str(n)]
    has_double = False
    never_decrease = True

    prev = digits[0]
    for i in range(1, len(digits)):
        current = digits[i]
        if current == prev:
            has_double = True
        if prev > current:
            never_decrease = False
            break
        prev = current

    return has_double and never_decrease


def get_passwords(range_min, range_max):
    passwords = []

    for i in range(range_min, range_max):
        if process_num(i):
            passwords.append(i)

    return passwords


def main():
    f = open("input.txt", "r")

    data = []
    for line in f:
        data = line.split("-")

    print("Getting passwords")
    passwords = get_passwords(int(data[0]), int(data[1]))

    print("The amount of different passwords is: " + str(len(passwords)))


main()
