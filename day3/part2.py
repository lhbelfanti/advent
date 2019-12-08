import re


def should_be_added(pos, wire_points):
    if len(wire_points) == 0:
        return True

    x1, y1 = wire_points[-1]
    x2, y2 = pos
    if x1 == x2 and y1 == y2:
        return False

    return True


def calculate_points(wire, start_point):
    wire_points = []
    pos = start_point
    for point in wire:
        p = re.split("([UDRL])", point)
        letter = p[1]
        moves = int(p[2])
        # print(p)
        if letter == "U":
            future_pos = pos[1] + moves
            current_pos = pos[1]
            for y in range(current_pos, future_pos + 1):
                pos = (pos[0], y)
                if should_be_added(pos, wire_points):
                    wire_points.append(pos)
        elif letter == "D":
            future_pos = pos[1] - moves
            current_pos = pos[1]
            for y in reversed(range(future_pos, current_pos + 1)):
                pos = (pos[0], y)
                if should_be_added(pos, wire_points):
                    wire_points.append(pos)
        elif letter == "R":
            future_pos = pos[0] + moves
            current_pos = pos[0]
            for x in range(current_pos, future_pos + 1):
                pos = (x, pos[1])
                if should_be_added(pos, wire_points):
                    wire_points.append(pos)
        elif letter == "L":
            future_pos = pos[0] - moves
            current_pos = pos[0]
            for x in reversed(range(future_pos, current_pos + 1)):
                pos = (x, pos[1])
                if should_be_added(pos, wire_points):
                    wire_points.append(pos)

    return wire_points


def calculate_intersections(points):
    set_x = frozenset(points[0])
    set_y = frozenset(points[1])
    return set_x & set_y


def count_steps(points, x1, y1):
    counter = 0
    for i in range(0, len(points)):
        x2, y2 = points[i]
        if x1 == x2 and y1 == y2:
            break
        counter += 1

    return counter


def calculate_steps(intersections, points):
    steps = []
    p0 = points[0]
    p1 = points[1]
    for intersection in intersections:
        x1, y1 = intersection
        c1 = count_steps(p0, x1, y1)
        c2 = count_steps(p1, x1, y1)
        if c1 + c2 > 0:
            steps.append(c1 + c2)

    return min(steps)


def main():
    f = open("input.txt", "r")

    wires = []
    for line in f:
        line = line.rstrip("\n")
        wires.append(line.split(","))

    print("Calculating wires points")
    start_point = (0, 0)
    points = []
    for wire in wires:
        points.append(calculate_points(wire, start_point))

    print("Calculating intersections")
    intersections = calculate_intersections(points)

    print("Calculating Steps")
    min_steps = calculate_steps(intersections, points)

    print("The min amount of steps are: " + str(min_steps))


main()
