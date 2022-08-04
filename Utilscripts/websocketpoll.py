import json
import asyncio
from websockets import connect

async def ws_connect():
    async with connect("ws://127.0.0.1:8545") as ws:
        await ws.send(json.dumps({"id":1, "method": "eth_subscribe", "params": ["newHeads"]}))
        counter = 0
        while True:
            res = await ws.recv()
            jsonResponse = json.loads(res)
            if counter == 0:
                counter += 1
                continue
            blockNumberHex = jsonResponse["params"]["result"]["number"]
            blockNumber = int(blockNumberHex, 16)
            print(blockNumber)

asyncio.run(ws_connect())

