import random


VERSE_FORMS = [
    ['a', 'c'],
    ['a', 'b', 'c'],
    ['a', 'b', 'b', 'c'],
    ['a', 'c', 'a', 'c']
]

class Queneau:

    def __init__(self, verse):
        self.verse = verse


    def sample(self):
        verse_form = random.choice(VERSE_FORMS)
        verses = random.sample(self.verse, len(verse_form))
        return [
            verses[index][key] for index, key in enumerate(verse_form)
        ]
