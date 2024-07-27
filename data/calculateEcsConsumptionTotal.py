from json import load

with open('ecs-bills.json') as f:
    j = load(f)
    r = 0
    for x in j:
        r += x["CashAmount"]
    print(r)