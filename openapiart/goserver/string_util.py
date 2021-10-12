
def camel_case(value: str) -> str:
    camel_case = ""
    for piece in value.split("_"):
        camel_case += "{}{}".format(piece[0].upper(), piece[1:])
    return camel_case
