# i got kinda tired of rebuilding

import os

try:
    os.system('docker build -t rentit:1.0 .')
    os.system('docker-compose up')
except KeyboardInterrupt:
    pass