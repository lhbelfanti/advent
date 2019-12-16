import math


def get_best_asteroid(asteroids):
	asteroid_id = (0, 0)
	max_asteroids = 0
	for a1 in asteroids:
		values = []
		for a2 in asteroids:
			if a1 == a2:
				continue
			x1, y1 = a1
			x2, y2 = a2
			atan2 = math.atan2(x1 - x2, y1 - y2)
			if atan2 not in values:
				values.append(atan2)
		asteroids_in_radar = len(values)
		if asteroids_in_radar > max_asteroids:
			max_asteroids = asteroids_in_radar
			asteroid_id = a1

	return asteroid_id, max_asteroids


def find_asteroids(data):
	asteroids = []
	for i in range(len(data)):
		for j in range(len(data)):
			d = data[i][j]
			if d == '#':
				asteroids.append((j, i))
	return asteroids


def get_higher_amount_of_asteroids_detected(data):
	asteroids = find_asteroids(data)
	best = get_best_asteroid(asteroids)
	return best


def main():
	f = open("input.txt", "r")

	data = []
	for line in f:
		line = line.rstrip()
		data.append(list(line))

	asteroid = get_higher_amount_of_asteroids_detected(data)

	print("Best is " + str(asteroid[0]) + " with " + str(asteroid[1]) + " other asteroids detected")


main()
