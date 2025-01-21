# Lem-In: Digital Ant Farm

## Project Overview

Lem-In is a Go-based program that simulates a digital ant farm. The goal is to find the quickest way to move a specified number of ants across a colony consisting of rooms and tunnels. The program reads input from a file describing the colony's structure and outputs the sequence of ant movements.

---

## Features

- **Pathfinding**: Identifies the quickest paths for ants to traverse the colony.
- **Path Optimization**: Avoids traffic jams and ensures ants follow the shortest paths effectively.
- **Error Handling**: Validates input and handles poorly formatted or invalid data gracefully.

---

## Input Format

The input file should include:

1. **Number of ants**  
   Example: > 0 

2. **Rooms**  
Rooms are defined as `name coord_x coord_y`.  
Example:  
##start 
Room1 23 3 
Room2 16 7 
##end 
Room0 9 5

3. **Tunnels**  
Tunnels are defined as `name1-name2`.  
Example:  
Room1-Room2 
Room0-Room1

4. **Comments**  
Comments start with `#`.  
Example:  
#This is a comment

---

## Output Format

1. The program displays the content of the input file.
2. Each line of moves shows the ants that moved, using the format `Lx-y`, where:
- `x` is the ant number.
- `y` is the destination room.

Example output:
L1-Room1 L2-Room2 L1-Room3 L2-Room4 L3-Room1

---

## Error Handling

The program checks for the following errors:
- Missing `##start` or `##end` rooms.
- Duplicated rooms or links.
- Links to undefined rooms to avoid unending cycles of movement.
- Invalid coordinates or improperly formatted input.

Error messages follow the format:  
ERROR: invalid data format
Examples:  
ERROR: invalid data format, no start room found ERROR: invalid data format, invalid number of ants

---
The progrma ensures that

### Room Rules
- Room names must not start with `L` or `#`.
- Room names must not contain spaces.

### Tunnel Rules
- Tunnels connect only two rooms.
- Each room can have multiple tunnels, but no duplicate connections.

### Movement Rules
- Each room can hold only one ant at a time (except `##start` and `##end`, which can hold multiple ants).
- Tunnels are used only once per turn.
- Ants must take the shortest available path while avoiding traffic jams.

---
## Usage

git clone the program Repository
```bash
$ git clone https://learn.zone01kisumu.ke/git/rodnochieng/lem-in

```

Run the program with a file as an argument:  
```bash
$ go run . <filename>
```

---

Example 1
Input file (test0.txt):

3

##start

1 23 3

2 16 7

##end

0 9 5

0-4

0-6

1-3

4-3

5-2

3-5

4-2

2-1

7-6

7-2

7-4

6-5

Output:

L1-3 L2-2

L1-4 L2-5 L3-3

L1-0 L2-6 L3-4

L2-0 L3-0

Example 2
Input file (test1.txt):

3

##start

0 1 0

##end

1 5 0

2 9 0

3 13 0

0-2

2-3

3-1

Output:

L1-2
L1-3 L2-2
L1-1 L2-3 L3-2
L2-1 L3-3
L3-1

### Development algorithm logic notes
The program first Validates input for:
* Proper formatting.
* Unique and valid room names.
* No duplicate coordinates or tunnels.
* Single start and end rooms.
Then Constructs and defines a graph model(ant colony)that contains the number of ants,start and end room, rooms, paths and ant names.

Find path : the program finds all possible paths between two rooms (start and end) in a graph (colony)
Data structure : Queue- Implements Breadth-First Search BFS traversal.
Path: A slice of strings representing the current traversal/path.
Finds all paths from start to end in a graph, avoiding cycles and ensuring all valid paths are explored

The program then validates and prints the input file content


## License
This project is open-source and available under the MIT License.

## Contributions
Contributions, issues, and feature requests are welcome!
Feel free to check the issues page for known bugs or feature ideas.