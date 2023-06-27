import requests

url = "http://localhost:8080/hello"

num_requests = 0
num_errors = 0

while num_errors < 5:
    response = requests.get(url)
    num_requests += 1

    if response.status_code == 429:
        num_errors += 1
        print(f"Request {num_requests}: Error 429 - Too Many Requests")
    elif response.status_code == 200:
        print(f"Request {num_requests}: Success")

print("Maximum number of errors reached. Closing the client.")