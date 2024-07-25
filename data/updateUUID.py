import json
from requests import get

with open('history-term-players.json', 'r', -1, 'utf8') as f:
    j = json.load(f)
    result = dict()
    for k, v in j.items():
        result[k] = []
        for u in v:
            res = get("https://api.ashcon.app/mojang/v2/user/" + u['name'])
            try:
                uuid = res.json()['uuid']
            except KeyError:
                print(u['name'] + ' not valid')
                result[k].append({
                    'name': u['name'],
                    'uuid': ''
                })
                continue
            print('got uuid mapping: ' + u['name'] + '->' + uuid)
            result[k].append({
                'name': u['name'],
                'uuid': uuid
            })

print(json.dumps(result))