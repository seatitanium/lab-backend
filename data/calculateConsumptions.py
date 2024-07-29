from json import load

result = {}
for x in ['ecs', 'oss', 'yundisk']:
    with open(f'{x}-bills.json') as f:
        j = load(f)
        result[x] = 0
        for y in j:
            result[x] += y["CashAmount"]
print(result)