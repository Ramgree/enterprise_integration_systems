# i got kinda tired of this too

import os

try:
    os.system('docker build -t rentit:1.0 .')
    os.system('docker-compose up --no-start')
    os.system('docker-compose run -d rentit')
    os.system('docker-compose ps')
    os.system('docker-compose run rentit-test')
    os.system('docker-compose down')
except KeyboardInterrupt:
    pass