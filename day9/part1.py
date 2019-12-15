def get_value(dat, idx, mode, rb):
	return dat[get_index(dat, idx, mode, rb)]


def get_index(dat, idx, mode, rb):
	if mode == 0:
		return dat[idx]
	elif mode == 1:
		return idx
	elif mode == 2:
		return rb + dat[idx]


def run_program(dat, idx, inp, rb):
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
		num1 = get_value(dat, idx + 1, modes[0], rb)
		num2 = get_value(dat, idx + 2, modes[1], rb)
		pos3 = get_index(dat, idx + 3, modes[2], rb)
		dat[pos3] = num1 + num2
		return run_program(dat, idx + 4, inp, rb)
	elif num == 2:
		num1 = get_value(dat, idx + 1, modes[0], rb)
		num2 = get_value(dat, idx + 2, modes[1], rb)
		pos3 = get_index(dat, idx + 3, modes[2], rb)
		dat[pos3] = num1 * num2
		return run_program(dat, idx + 4, inp, rb)
	elif num == 3:
		pos1 = get_index(dat, idx + 1, modes[0], rb)
		dat[pos1] = inp
		return run_program(dat, idx + 2, inp, rb)
	elif num == 4:
		num1 = get_value(dat, idx + 1, modes[0], rb)
		print(num1)
		return run_program(dat, idx + 2, inp, rb)
	elif num == 5:
		num1 = get_value(dat, idx + 1, modes[0], rb)
		num2 = get_value(dat, idx + 2, modes[1], rb)
		if num1 != 0:
			return run_program(dat, num2, inp, rb)
		else:
			return run_program(dat, idx + 3, inp, rb)
	elif num == 6:
		num1 = get_value(dat, idx + 1, modes[0], rb)
		num2 = get_value(dat, idx + 2, modes[1], rb)
		if num1 == 0:
			return run_program(dat, num2, inp, rb)
		else:
			return run_program(dat, idx + 3, inp, rb)
	elif num == 7:
		num1 = get_value(dat, idx + 1, modes[0], rb)
		num2 = get_value(dat, idx + 2, modes[1], rb)
		pos3 = get_index(dat, idx + 3, modes[2], rb)
		if num1 < num2:
			dat[pos3] = 1
		else:
			dat[pos3] = 0
		return run_program(dat, idx + 4, inp, rb)
	elif num == 8:
		num1 = get_value(dat, idx + 1, modes[0], rb)
		num2 = get_value(dat, idx + 2, modes[1], rb)
		pos3 = get_index(dat, idx + 3, modes[2], rb)
		if num1 == num2:
			dat[pos3] = 1
		else:
			dat[pos3] = 0
		return run_program(dat, idx + 4, inp, rb)
	elif num == 9:
		num1 = get_value(dat, idx + 1, modes[0], rb)
		return run_program(dat, idx + 2, inp, rb + num1)
	elif num == 99:
		return dat


def main():
	f = open("input.txt", "r")

	data = []
	for line in f:
		data = line.split(",")

	# For some reason the program ask for a position 1000 and it doesn't exist
	# Adding extra 0s like we do with the modes
	for j in range(100):
		data.append(0)

	for i in range(0, len(data)):
		data[i] = int(data[i])

	input_value = 1
	run_program(data, 0, input_value, 0)


main()
