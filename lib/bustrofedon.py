#!/usr/bin/env python3
import sys
from bustrofedon.composer import Poem

OPTIONS = {
    'path': '../corpus/model.json'
}


def main():
    poem = Poem(OPTIONS)


if __name__ == '__main__':
    sys.exit(main())
