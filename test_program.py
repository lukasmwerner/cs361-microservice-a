import requests

id = input("Enter the id of the pokemon: ")
resp = requests.get("http://localhost:8000/pokemon/"+id)
print(resp.status_code, resp.json())
