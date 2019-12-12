from itertools import permutations


def run_program(dat, idx, inp, out):
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
		return run_program(dat, idx + 4, inp, out)
	elif num == 2:
		num1 = dat[dat[idx + 1]] if modes[0] == 0 else dat[idx + 1]
		num2 = dat[dat[idx + 2]] if modes[1] == 0 else dat[idx + 2]
		pos3 = dat[idx + 3]
		dat[pos3] = num1 * num2
		return run_program(dat, idx + 4, inp, out)
	elif num == 3:
		pos1 = dat[idx + 1]
		dat[pos1] = inp.pop(0)
		return run_program(dat, idx + 2, inp, out)
	elif num == 4:
		pos1 = dat[dat[idx + 1]] if modes[0] == 0 else dat[idx + 1]
		out = pos1
		return [out, idx + 2, False]
	# return run_program(dat, idx + 2, inp, out)
	elif num == 5:
		num1 = dat[dat[idx + 1]] if modes[0] == 0 else dat[idx + 1]
		num2 = dat[dat[idx + 2]] if modes[1] == 0 else dat[idx + 2]
		if num1 != 0:
			return run_program(dat, num2, inp, out)
		return run_program(dat, idx + 3, inp, out)
	elif num == 6:
		num1 = dat[dat[idx + 1]] if modes[0] == 0 else dat[idx + 1]
		num2 = dat[dat[idx + 2]] if modes[1] == 0 else dat[idx + 2]
		if num1 == 0:
			return run_program(dat, num2, inp, out)
		return run_program(dat, idx + 3, inp, out)
	elif num == 7:
		num1 = dat[dat[idx + 1]] if modes[0] == 0 else dat[idx + 1]
		num2 = dat[dat[idx + 2]] if modes[1] == 0 else dat[idx + 2]
		pos3 = dat[idx + 3]
		if num1 < num2:
			dat[pos3] = 1
		else:
			dat[pos3] = 0
		return run_program(dat, idx + 4, inp, out)
	elif num == 8:
		num1 = dat[dat[idx + 1]] if modes[0] == 0 else dat[idx + 1]
		num2 = dat[dat[idx + 2]] if modes[1] == 0 else dat[idx + 2]
		pos3 = dat[idx + 3]
		if num1 == num2:
			dat[pos3] = 1
		else:
			dat[pos3] = 0
		return run_program(dat, idx + 4, inp, out)
	elif num == 99:
		return [out, idx, True]


def run_loop(amplifiers, data, permutation, amp_input, amp_d, use_phase_input=False):
	amps_data = [{}] * len(amplifiers) if amp_d is None else amp_d

	for i in range(0, len(amplifiers)):
		if "data" not in amps_data[i]:
			d = data.copy()
			amps_data[i] = {"data": d, "idx": 0}

		d = amps_data[i]["data"]
		idx = amps_data[i]["idx"]

		phase_setting = permutation[i]
		program_input = [phase_setting, amp_input] if use_phase_input else [amp_input]

		out_data = run_program(d, idx, program_input, amp_input)
		amp_input = out_data[0]
		amps_data[i] = {"data": d, "idx": out_data[1]}

		if i == len(amplifiers) - 1:
			if not out_data[2]:
				return run_loop(amplifiers, data, permutation, amp_input, amps_data)

	return amp_input


def get_highest_signal(data, input_value):
	amplifiers = [5, 6, 7, 8, 9]
	perm = permutations(amplifiers)
	total = {}
	for p in list(perm):
		amp_input = input_value
		out = run_loop(amplifiers, data, p, amp_input, None, True)
		setting = ''.join(map(str, p))
		total[setting] = out

	max_value = (0, "")
	for seq in total:
		if total[seq] > max_value[0]:
			max_value = (total[seq], seq)

	print("Max thruster signal: " + str(total[max_value[1]]))


def main():
	f = open("input.txt", "r")

	data = []
	for line in f:
		data = line.split(",")

	for i in range(0, len(data)):
		data[i] = int(data[i])

	input_value = 0
	get_highest_signal(data, input_value)


main()
