def parse(stacks):
    parsed = [[] for _ in range(0, int(stacks[-2]))]

    for layer in stacks.split('\n'):
        for i in range(0, len(layer)):
            if layer[i].isalpha():
                parsed[i // 4] = [layer[i]] + parsed[i // 4]

    return parsed

def solve(stacks, moves):
    for line in moves.strip().split('\n'):
        line = [(int(n) if n.isdigit() else -1) for n in line.split()]
        for i in range(0, line[1]):
            stacks[line[5] - 1].append(stacks[line[3] - 1].pop())

    return "".join(stack[-1] for stack in stacks)

if __name__ == "__main__":
    f = open('input.txt')
    data = f.read().split('\n\n')
    f.close()

    stacks = parse(data[0])

    print(solve(stacks, data[1]))

