from math import hypot, pi, atan2


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
			at = atan2(x1 - x2, y1 - y2)
			if at not in values:
				values.append(at)
		asteroids_in_radar = len(values)
		if asteroids_in_radar > max_asteroids:
			max_asteroids = asteroids_in_radar
			asteroid_id = a1

	return asteroid_id


def angle(a, b):
	return atan2(b[0] - a[0], a[1] - b[1]) % (2 * pi)


def rank_asteroids(asteroids, a):
	ranks = {}
	for i, b in enumerate(asteroids):
		s = 0
		for c in asteroids[:i]:
			s += int(angle(a, b) == angle(a, c))
		ranks[b] = s

	return ranks


def get_xth_asteroid_vaporized(asteroids, best, th):
	asteroids.remove(best)
	asteroids.sort(key=lambda a: hypot(a[0] - best[0], a[1] - best[1]))
	ranks = rank_asteroids(asteroids, best)
	x, y = sorted(asteroids, key=lambda b: (ranks[b], angle(best, b)))[th]
	return x, y


def main():
	f = open("input.txt", "r")
	lines = list(f)
	asteroids = []
	for y in range(len(lines)):
		for x in range(len(lines[0])):
			if lines[y][x] == '#':
				asteroids.append((x, y))

	best = get_best_asteroid(asteroids)
	asteroid = get_xth_asteroid_vaporized(asteroids, best, 199)
	value = (asteroid[0] * 100) + asteroid[1]
	print("The 200th asteroid that will be vaporized is: " + str(value))


main()
