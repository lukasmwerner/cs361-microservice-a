# cs361-microservice-a

Pokedex Info Service

UML Sequence Diagram:
```mermaid
sequenceDiagram
    User->>+Service A: Get :id Pokemon
    Service A->>+PokéAPI: HTTP GET Request
    PokéAPI->>+Service A: JSON Data Response
    Service A->>+User: Returns Picture URL, Description

```

Example request (using localhosted version):
```python
import requests

id = input("Enter the id of the pokemon: ")
resp = requests.get("http://localhost:8000/pokemon/"+id)
print(resp.status_code, resp.json())
```
