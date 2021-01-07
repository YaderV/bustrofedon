#!/usr/bin/env python3
import sys
from bustrofedon.composer import Poem

PATH = '../corpus/model.json'

OPTIONS = {
    'path': PATH
}


def main():
    Poem(OPTIONS)


if __name__ == '__main__':
    sys.exit(main())
