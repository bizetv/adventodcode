from sys import stdin
from re import findall

ops = {
    'AND'    : lambda a, b: a & b,
    'OR'     : lambda a, b: a | b,
    'RSHIFT' : lambda a, b: a >> b,
    'LSHIFT' : lambda a, b: a << b,
    'NOT'    : lambda a: ~a
}

def evaluate(wire, wires):
    if wire.isdigit():
        return int(wire)

    if isinstance(wires[wire], int):
        return wires[wire]

    tokens = findall("([^\s]+)", wires[wire])
    wires[wire] = {
        1: lambda a: evaluate(a, wires),
        2: lambda op, a: ops[op](evaluate(a, wires)),
        3: lambda a, op, b: ops[op](evaluate(a, wires), evaluate(b, wires))
    }[len(tokens)](*tokens)

    return wires[wire]

def day7(input):
    wires = {k:v for v, k in findall("(.*) -> ([^\s]*)", input)}
    part1 = evaluate('a', wires.copy())
    wires['b'] = str(part1)
    part2 = evaluate('a', wires.copy())
    return part1, part2

print(day7(stdin.read()))
