## desmos-server 

A listener for http requests to plot equations in desmos in realtime for educational simulations

### How it works
An http listener listens on port :8080 for incoming plot requests. These requests are then forwaded to the frontend via a websocket connection on port :8081. The frontend calls desmos api to plot the request

### How to Use

Start the listener, by default it runs on port 8080

```
go run main.go
```

Serve the desmos frontend

```
cd static
http-server
```

Send HTTP GET requests from any application with with query params to dynamically plot or remove existing equation on the desmos frontend

```
HTTP GET http://localhost:8080/add?id=<base64 encoded id>&exp=<base64 encoded latex expression>
HTTP GET http://localhost:8080/remove?id=<base64 encoded id>
```

For example, to plot a normal distribution from a python application,

```python
def toB4(s: str):
    return base64.b64encode(s.encode("ascii"))

requests.get(url = "http://localhost:8080/add", params={"id": toB64("norm"), "exp": toB64(r"""\frac{e^{\frac{-x^2}{2}}}{\sqrt{2\pi}}""")})
```

PS - Some safety features like concurrent write protection on websocket are not added yet 
