
def pascal_case(value: str) -> str:
    pascal_case = ""
    for piece in value.split("_"):
        pascal_case += "{}{}".format(piece[0].upper(), piece[1:])
    return pascal_case
