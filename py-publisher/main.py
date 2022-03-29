import zmq
import sys
import time
import os
import json
from datetime import datetime

pub_url = os.getenv("PUB_URL", default = "tcp://*:5555")
context = zmq.Context()
socket = context.socket(zmq.PUB)
socket.bind(pub_url)
print("Accepting connections at %s" % (pub_url), flush=True)

time.sleep(1)

count = int(os.getenv("PUBLISH_COUNT", default = "42"))

for i in range(count):
    text = "from py-publisher %s" % (i)
    message = json.dumps({
        'message': text,
        'timestamp': int(datetime.now().timestamp())
    })
    print("Publishing: %s" % (text), flush=True)
    socket.send_string(message)
    
    time.sleep(1)

