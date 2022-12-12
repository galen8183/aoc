def move(source, dest, num):
    dest += source[(len(source) - num):]
    for i in range(0, num): source.pop()

def main(stacks, file):
    data = open(file)

    for line in data:
        line = [(int(n) if n.isdigit() else -1) for n in line.split()]
        move(stacks[line[3] - 1], stacks[line[5] - 1], line[1])

    ans = ''
    for stack in stacks: ans += stack[-1]
    print(ans)
    data.close()

if __name__ == "__main__":
    stacks = [
        ['G', 'T', 'R', 'W'],
        ['G', 'C', 'H', 'P', 'M', 'S', 'V', 'W'],
        ['C', 'L', 'T', 'S', 'G', 'M'],
        ['J', 'H', 'D', 'M', 'W', 'R', 'F'],
        ['P', 'Q', 'L', 'H', 'S', 'W', 'F', 'J'],
        ['P', 'J', 'D', 'N', 'F', 'M', 'S'],
        ['Z', 'B', 'D', 'F', 'G', 'C', 'S', 'J'],
        ['R', 'T', 'B'],
        ['H', 'N', 'W', 'L', 'C']
    ]

    test = [
        ['Z', 'N'],
        ['M', 'C', 'D'],
        ['P']
    ]

    main(stacks, 'input.txt')
