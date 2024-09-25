import os
import sys
import scripts.hello

if __name__ == '__main__':
    cur_dir = os.getcwd()
    sys.path.append(os.path.join(cur_dir, '/scripts/protoc'))

    scripts.hello.dotest()
