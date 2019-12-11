def find_level(root, node, level):
	if not root:
		return None

	if root.value == node:
		return level + 1

	left = find_level(root.left, node, level + 1)
	right = find_level(root.right, node, level + 1)

	if not left and not right:
		return None
	elif left:
		return left
	else:
		return right


def find_lca_node(root, n1, n2):
	if not root or root.value == n1 or root.value == n2:
		return root

	lca_left = find_lca_node(root.left, n1, n2)
	lca_right = find_lca_node(root.right, n1, n2)

	if lca_left and lca_right:
		return root

	return lca_right if lca_right else lca_left


def find_distance(root, n1, n2):
	level_n1 = find_level(root, n1, 0)
	level_n2 = find_level(root, n2, 0)
	lca = find_lca_node(root, n1, n2)
	level_lca = find_level(root, lca.value, 0)
	distance = str((level_n1 + level_n2 - 2 * level_lca) - 2)
	return distance


def get_next_node(data, curr_node, leaves):
	if curr_node.value in data:
		nodes = data[curr_node.value]
		for i in range(0, len(nodes)):
			n = Node(nodes[i])
			curr_node.add_node(n)
			get_next_node(data, n, leaves)
	else:
		leaves.append(curr_node)

	return leaves


def create_tree(data, root):
	curr_node = root
	leaves = get_next_node(data, curr_node, [])
	return leaves


class Node:
	def __init__(self, value):
		self.value = value
		self.right = None
		self.left = None

	def add_node(self, node):
		if self.right is None:
			self.right = node
		elif self.left is None:
			self.left = node


def main():
	f = open("input.txt", "r")
	data = {}
	for line in f:
		objs = line.split(")")
		if objs[0] not in data:
			data[objs[0]] = []
		data[objs[0]].append(objs[1].rstrip())

	root = Node("COM")
	create_tree(data, root)
	distance = find_distance(root, "YOU", "SAN")
	print("The distance between node YOU and node SAN is: " + distance)


main()
