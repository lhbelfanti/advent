def print_layer(layer, wide, tall):
	line = ""
	for i in range(0, tall):
		for j in range(0, wide):
			c = layer.matrix[i][j]
			line += " " if c == 0 else "█"
		print(line)
		line = ""


def combine_layers(layers, wide, tall):
	combined_layer = Layer(wide, tall)
	for i in range(0, tall):
		for j in range(0, wide):
			color = 2
			for lay in layers:
				c = lay.matrix[i][j]
				if c != 2 and color == 2:
					color = c
			combined_layer.matrix[i][j] = color

	return combined_layer


def create_layers(pixels, wide, tall):
	image_size = wide * tall
	layers = []
	length = len(pixels) + 1
	nums = []
	current = None
	for i in range(0, length):
		if i % image_size == 0:
			if len(nums) > 0:
				current.add_data(nums)
				layers.append(current)
			nums = []
			current = Layer(wide, tall)

		if i != length - 1:
			nums.append(pixels[i])

	return layers


class Layer:
	def __init__(self, wide, tall):
		self.wide = wide
		self.tall = tall
		self.zeros = 0
		self.ones = 0
		self.twos = 0
		self.matrix = [0] * self.tall
		for i in range(0, len(self.matrix)):
			self.matrix[i] = [0] * self.wide

	def add_data(self, nums):
		c = 0
		for i in range(0, self.tall):
			for j in range(0, self.wide):
				n = nums[c]
				self.zeros += 1 if n == 0 else 0
				self.ones += 1 if n == 1 else 0
				self.twos += 1 if n == 2 else 0
				self.matrix[i][j] = n
				c += 1

	def multiply(self):
		return self.ones * self.twos


def main():
	f = open("input.txt", "r")

	wide = 25
	tall = 6

	pixels = []
	line = ""
	for line in f:
		line = line.rstrip()

	for p in line:
		pixels.append(int(p))

	layers = create_layers(pixels, wide, tall)
	layer = combine_layers(layers, wide, tall)
	print_layer(layer, wide, tall)


main()
