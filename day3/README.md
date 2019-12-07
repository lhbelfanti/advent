## --- Day 3: Crossed Wires ---

### --- Part One ---

The gravity assist was successful, and you're well on your way to the Venus refuelling station. During the rush back on Earth, the fuel management system wasn't completely installed, so that's next on the priority list.

Opening the front panel reveals a jumble of wires. Specifically, two wires are connected to a central port and extend outward on a grid. You trace the path each wire takes as it leaves the central port, one wire per line of text (your puzzle input).

The wires twist and turn, but the two wires occasionally cross paths. To fix the circuit, you need to find the intersection point closest to the central port. Because the wires are on a grid, use the Manhattan distance for this measurement. While the wires do technically cross right at the central port where they both start, this point does not count, nor does a wire count as crossing with itself.

For example, if the first wire's path is R8,U5,L5,D3, then starting from the central port (o), it goes right 8, up 5, left 5, and finally down 3:

<img src="https://raw.githubusercontent.com/lhbelfanti/adventure-of-code-2019/master/day3/example1.png?token=AB7KT6BLJL45BRUVM2EJQUS56UFS2" width="150" height="200" />

Then, if the second wire's path is U7,R6,D4,L4, it goes up 7, right 6, down 4, and left 4:

<img src="https://raw.githubusercontent.com/lhbelfanti/adventure-of-code-2019/master/day3/example2.png?token=AB7KT6H5MMB67XHQOK2HOWS56UFW4" width="150" height="200" />


These wires cross at two locations (marked X), but the lower-left one is closer to the central port: its distance is 3 + 3 = 6.

Here are a few more examples:

- R75,D30,R83,U83,L12,D49,R71,U7,L72
- U62,R66,U55,R34,D71,R55,D58,R83 = distance 159
- R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51
- U98,R91,D20,R16,D67,R40,U7,R15,U6,R7 = distance 135

What is the Manhattan distance from the central port to the closest intersection?