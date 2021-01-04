import json


class Poem:
    def __init__(self, options={}):
        self.load_model(options.get('path'))

    def load_model(self, path):
        with open(path, encoding='utf-8') as fp:
            model = json.load(fp)
            self.queneau = model['verses']
            self.wordlists = model['words']
            self.excludes = model['excludes']
