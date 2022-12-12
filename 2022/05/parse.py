stacks = """
    [W]         [J]     [J]        
    [V]     [F] [F] [S] [S]        
    [S] [M] [R] [W] [M] [C]        
    [M] [G] [W] [S] [F] [G]     [C]
[W] [P] [S] [M] [H] [N] [F]     [L]
[R] [H] [T] [D] [L] [D] [D] [B] [W]
[T] [C] [L] [H] [Q] [J] [B] [T] [N]
[G] [G] [C] [J] [P] [P] [Z] [R] [H]
 1   2   3   4   5   6   7   8   9 
"""

def parse(stacks):
    parsed = [ [] for i in range(0, 9) ]

    for layer in stacks.split('\n'):
        for i in range(0, len(layer)):
            if layer[i].isalpha():
                parsed[i // 4] = [layer[i]] + parsed[i // 4]

    for stack in parsed: print(stack)
