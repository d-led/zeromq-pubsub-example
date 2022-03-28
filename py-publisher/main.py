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
print("Accepting connections at %s" % (pub_url))

for i in range(42):
    text = "Message %s" % (i)
    message = json.dumps({
        'message': text,
        'timestamp': int(datetime.now().timestamp())
    })
    print("Publishing: %s" % (text))
    socket.send_string(message, zmq.NOBLOCK)
    
    time.sleep(1)

