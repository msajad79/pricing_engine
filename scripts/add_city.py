import requests
import csv
import os
import json

session = requests.Session()
data = []
with open("./bootcamp/city.csv", 'r', newline='', encoding="utf8") as csvfile:
    reader = csv.reader(csvfile)
    for row in reader:
        data.append({"name":row[2]})

#os.environ['NO_PROXY'] = '127.0.0.1'
res = requests.post("http://127.0.0.1:8080/add/cities", data=json.dumps(data))
print(res.json())

stop