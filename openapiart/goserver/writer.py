
class Writer(object):
    _indents = []  # type [str]

    @property
    def strings(self):
        # type: () -> [str]
        return self._strings

    def __init__(self, indent):
        self._indent = indent
        self._strings = []  # type [str]

    def write_line(self, *txt):
        # type: (str) -> Writer
        for t in txt:
            self._strings.append( ''.join(self._indents) + t)
        return self

    def push_indent(self):
        # type: () -> Writer
        self._indents.append(self._indent)
        return self

    def pop_indent(self):
        # type: () -> Writer
        self._indents.pop()
        return self

