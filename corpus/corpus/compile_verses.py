#!/usr/bin/env python3
import argparse
import codecs
import json
import os
import re
import sys
from glob import glob


OUTPUT_FILENAME = '../model/verses.json'

RE_CONTAINS_URL = re.compile('https?://')
RE_WEAK_LINE = re.compile(' (and|by|from|is|of|that|the|with)\\n',
                          flags=re.IGNORECASE)

REPLACE = (
    # convert smart quotes
    (re.compile(u'\u2018|\u2019'), "'"),
    (re.compile(u'\u201C|\u201D'), '"'),

    # remove double quotes
    (re.compile('"'), ''),

    # em dash
    (re.compile(u'--|\u2013|\\B-|-\\B|~'), u'\u2014'),
    (re.compile(u'\u2014'), u' \u2014'),

    # ellipsis
    (re.compile(r'\.( \.|\.)+'), u'\u2026'),
    (re.compile(u'\\b \u2026'), u'\u2026'),
)


# Objetos a, b, c para versos
output_haiku = []
# lista de todos los haikus con el formato verso 1 / verso 2
output_texts = []


# Evita leer la primera linea si esta empieza
# de esta manera // foo
def parse_credit(line):
    line = line.strip()
    return line[2:].lstrip() if line.startswith('//') else None


def convert_dirs(dirs, export=False):
    # Iterar directorios
    for dirname in dirs:
        # Glob devuelve los archivos encontrados que cumplen
        # la condición, en este caso directorio/*.txt
        # por eso el uso de os.path.join
        for fn in glob(os.path.join(dirname, '*.txt')):
            print('Reading {}'.format(fn))
            input_lines = None
            # Codecs.open es una alternativa con mayores opciones
            # a open o el módulo io
            with codecs.open(fn, encoding='utf-8') as fp:
                # cada linea como un elemento de la lista
                input_lines = fp.readlines()

            # find a credit line for the file, if one exists
            file_credit = parse_credit(input_lines[0])
            if file_credit:
                input_lines.pop(0)

            # strip trailing whitespace and comments
            input_lines = (line.split('#', 1)[0].strip()
                           for line in input_lines
                           if not line.lstrip().startswith('#'))

            # select all unique verses,
            # where each verse is separated by an empty line
            unique_haiku = set()
            # Iterar una lista de todos los haikus
            # primero se lee el objeto string con join
            # luego se crea una lista divida por linea vacía
            # la separación de los haikus
            for haiku in '\n'.join(input_lines).split('\n\n'):
                haiku = haiku.strip().lower()

                # Si contiene algun formato de url
                if RE_CONTAINS_URL.search(haiku):
                    raise ValueError('haiku includes a url: {}'
                                     .format(haiku))
                # TODO
                # No entiendo
                elif RE_WEAK_LINE.search(haiku):
                    raise ValueError('haiku contains a weak line: {}'
                                     .format(haiku))
                for regex, repl in REPLACE:
                    haiku = regex.sub(repl, haiku)
                if haiku:
                    unique_haiku.add(haiku)
            print('    %d unique haiku' % len(unique_haiku))

            # separate the lines into first, middle, and last buckets
            for haiku in unique_haiku:
                haiku_lines = haiku.split('\n')
                credit = file_credit
                line_count = len(haiku_lines)
                # Lista de todos los versos
                output_lines = []
                # Lista de los primeros, segundos y terceros versos
                buckets = [[], [], []]
                for i, line in enumerate(haiku_lines):
                    # Determina si el verso es el primero el
                    # terceo o el segundo
                    bucket = 0 if i == 0 else (2 if i == line_count - 1 else 1)

                    # comprobar si el verso esta compuesto
                    # por mas de una plabra
                    tokens = line.split()
                    if not tokens:
                        continue
                    line = ' '.join(tokens)
                    output_lines.append(line)
                    buckets[bucket].append(line)
                if output_lines:
                    obj = {
                        'a': tuple(buckets[0]),
                        'b': tuple(buckets[1]),
                        'c': tuple(buckets[2]),
                    }
                    if credit:
                        obj['source'] = credit
                    output_haiku.append(obj)
                    output_texts.append(' / '.join(output_lines))

    if export:
        print('Writing {}'.format(OUTPUT_FILENAME))
        with codecs.open(OUTPUT_FILENAME, 'w', encoding='utf-8') as fp:
            json.dump(output_haiku, fp,
                      indent=2, sort_keys=True, ensure_ascii=False)

        maxlen = max(len(text) for text in output_texts)
        print('Wrote {} poems with maximum length {}'
              .format(len(output_haiku), maxlen))

    return output_haiku


def main():
    parser = argparse.ArgumentParser()
    parser.add_argument('srcdir', nargs='*')
    args = parser.parse_args()

    if not args.srcdir:
        parser.print_help()
        return 1

    convert_dirs(args.srcdir, export=True)


if __name__ == '__main__':
    sys.exit(main())
