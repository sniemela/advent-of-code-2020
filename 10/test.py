bag = [0]
with open('test-input.txt') as fp:
    line = fp.readline()
    while line:
        bag.append(int(line.strip()))
        line = fp.readline()

bag.sort()
bag.append(bag[-1:][0]+3)
countDict = {}
for i in range(1, len(bag)):
    diff = bag[i] - bag[i - 1]
    if diff in countDict:
        countDict[diff] += 1
    else:
        countDict[diff] = 1

arrange = [1]+[0]*bag[-1]
print(arrange)

print(bag)

for i in bag[1:]:
    arrange[i] = arrange[i-3] + arrange[i-2] + arrange[i-1]

print(arrange)

print(f'Part 1\n{countDict[1] * countDict[3]}\n\nPart 2\n{arrange[-1]}')