# cd webapp & npm i & npm run build & cp -r dist/* ../api/static & cd ../api & go build -o server & mkdir output & cp server output/server & cp -r static/* 

import os, shutil
import platform


os.chdir('webapp')
os.system('npm i')
os.system('npm run build')
os.chdir('..')

if os.path.isdir('build'):
    os.rmdir('build')
os.mkdir('build')

if os.path.exists('./api/static'):
    os.rmdir('./api/static')

shutil.copytree('./webapp/dist/', './api/static')
os.chdir('api')

binary_name = ''
platform_name = platform.system()

if platform_name == 'Linux' or platform_name == 'Darwin':
    binary_name = 'server'
elif platform_name == 'Windows':
    binary_name = 'server.exe'

os.system('go build -o ' + platform_name)
shutil.copy('server', '../build')
shutil.copytree('./static', '../build/static')