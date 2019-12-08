import re


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
                wire_points.append(pos)
                # print("U " + str(pos))
        elif letter == "D":
            future_pos = pos[1] - moves
            current_pos = pos[1]
            for y in reversed(range(future_pos, current_pos + 1)):
                pos = (pos[0], y)
                wire_points.append(pos)
                # print("D " + str(pos))
        elif letter == "R":
            future_pos = pos[0] + moves
            current_pos = pos[0]
            for x in range(current_pos, future_pos + 1):
                pos = (x, pos[1])
                wire_points.append(pos)
                # print("R " + str(pos))
        elif letter == "L":
            future_pos = pos[0] - moves
            current_pos = pos[0]
            for x in reversed(range(future_pos, current_pos + 1)):
                pos = (x, pos[1])
                wire_points.append(pos)
                # print("L " + str(pos))

    return wire_points


def calculate_intersections(points):
    set_x = frozenset(points[0])
    set_y = frozenset(points[1])
    return set_x & set_y


def manhattan_distance(intersections, start_point):
    min_distance = 1000000
    for i in intersections:
        x1, y1 = i
        x2, y2 = start_point
        if x1 == 0 and y1 == 0:
            continue
        distance = abs(x1 - x2) + abs(y1 - y2)
        # print("Distance: abs(" + str(x1) + " - " + str(x2) + ") + abs(" + str(y1) + " - " + str(y2) + ") = " + str(distance))
        if distance < min_distance:
            min_distance = distance
    return min_distance


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

    print("Calculating Manhattan distance")
    distance = manhattan_distance(intersections, start_point)

    print("The min distance is: " + str(distance))


main()
