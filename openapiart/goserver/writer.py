class Writer(object):
    _indents: [str] = []

    @property
    def strings(self) -> [str]:
        return self._strings

    def __init__(self, indent: str):
        self._indent = indent
        self._strings: [str] = []

    def write_line(self, *txt: str) -> "Writer":
        for t in txt:
            self._strings.append("".join(self._indents) + t)
        return self

    def push_indent(self) -> "Writer":
        self._indents.append(self._indent)
        return self

    def pop_indent(self) -> "Writer":
        self._indents.pop()
        return self
