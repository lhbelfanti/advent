def reduce_leaves(leaves):
	should_reduce = False
	total = 0
	for i in range(0, len(leaves)):
		leaf = leaves[i]
		if leaf is not None and leaf.orbits == 0:
			c = get_orbits(leaf)
			leaf.orbits = c
			total += c
			if leaf.parent is not None:
				should_reduce = True
			leaves[i] = leaf.parent

	if should_reduce:
		return reduce_leaves(leaves) + total

	return 0


def get_orbits(node):
	if node is None or node.parent is None:
		return 0
	return get_orbits(node.parent) + 1


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
		self.parent = None
		self.right = None
		self.left = None
		self.orbits = 0

	def add_node(self, node):
		if self.right is None:
			self.right = node
		elif self.left is None:
			self.left = node
		node.parent = self


def main():
	f = open("input.txt", "r")
	data = {}
	for line in f:
		objs = line.split(")")
		if objs[0] not in data:
			data[objs[0]] = []
		data[objs[0]].append(objs[1].rstrip())

	root = Node("COM")
	leaves = create_tree(data, root)

	total = reduce_leaves(leaves)

	print("The total number of direct and indirect orbits is: " + str(total))


main()
