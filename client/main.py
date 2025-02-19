from fastapi import FastAPI
import grpc
import jokes_pb2
import jokes_pb2_grpc

app = FastAPI()

@app.get("/health")
async def health():
    return {"status": "ok"}

@app.get("/jokes")
async def jokes(batch_size: int = 25):
    with grpc.insecure_channel("localhost:5050") as channel:
        stub = jokes_pb2_grpc.JokeServiceStub(channel)
        response = stub.GetJokes(jokes_pb2.GetJokesRequest(batch_size=batch_size))
        return {"result": [{"id": joke.id, "url": joke.url, "value": joke.value} for joke in response
                           .jokes], "total": len(response.jokes)}

    return {"result": []}