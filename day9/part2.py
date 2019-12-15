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

	output_value = 0
	while idx < len(dat) - 1:
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
			idx += 4
		elif num == 2:
			num1 = get_value(dat, idx + 1, modes[0], rb)
			num2 = get_value(dat, idx + 2, modes[1], rb)
			pos3 = get_index(dat, idx + 3, modes[2], rb)
			dat[pos3] = num1 * num2
			idx += 4
		elif num == 3:
			pos1 = get_index(dat, idx + 1, modes[0], rb)
			dat[pos1] = inp
			idx += 2
		elif num == 4:
			pos1 = get_value(dat, idx + 1, modes[0], rb)
			output_value = pos1
			idx += 2
		elif num == 5:
			num1 = get_value(dat, idx + 1, modes[0], rb)
			num2 = get_value(dat, idx + 2, modes[1], rb)
			if num1 != 0:
				idx = num2
			else:
				idx += 3
		elif num == 6:
			num1 = get_value(dat, idx + 1, modes[0], rb)
			num2 = get_value(dat, idx + 2, modes[1], rb)
			if num1 == 0:
				idx = num2
			else:
				idx += 3
		elif num == 7:
			num1 = get_value(dat, idx + 1, modes[0], rb)
			num2 = get_value(dat, idx + 2, modes[1], rb)
			pos3 = get_index(dat, idx + 3, modes[2], rb)
			if num1 < num2:
				dat[pos3] = 1
			else:
				dat[pos3] = 0
			idx += 4
		elif num == 8:
			num1 = get_value(dat, idx + 1, modes[0], rb)
			num2 = get_value(dat, idx + 2, modes[1], rb)
			pos3 = get_index(dat, idx + 3, modes[2], rb)
			if num1 == num2:
				dat[pos3] = 1
			else:
				dat[pos3] = 0
			idx += 4
		elif num == 9:
			num1 = get_value(dat, idx + 1, modes[0], rb)
			idx += 2
			rb += num1
		elif num == 99:
			return output_value


def main():
	f = open("input.txt", "r")

	data = []
	for line in f:
		data = line.split(",")

	# For some reason the program ask for a position 1000 and it doesn't exist
	# Adding extra 0s like we do with the modes
	for j in range(300):
		data.append(0)

	for i in range(0, len(data)):
		data[i] = int(data[i])

	input_value = 2
	output_value = run_program(data, 0, input_value, 0)
	print("The coordinates of the distress signal are: " + str(output_value))

main()
