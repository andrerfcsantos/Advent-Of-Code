import requests
import sys

if len(sys.argv) != 3:
    print('Wrong number of arguments')
    exit(0)

day, cookie = int(sys.argv[1]),sys.argv[2]
headers = {'session': cookie}
url = f'https://adventofcode.com/2017/day/{day}/input'

session = requests.Session()
resp = session.get(url,cookies=headers)

in_file = open(f'day{day:02}.txt','w')
in_file.write(resp.text)
in_file.close()