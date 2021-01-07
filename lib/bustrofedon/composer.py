import json
from bustrofedon.queneau import Queneau


class Poem:
    def __init__(self, options={}):
        self.load_model(options.get('path'))

    def load_model(self, path):
        with open(path, encoding='utf-8') as fp:
            model = json.load(fp)
            self.queneau = Queneau(model['verses'])
            print(self.queneau.sample())
            self.wordlists = model['words']
            self.excludes = model['excludes']
