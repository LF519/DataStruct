# -*- conding:utf-8 -*-

class BitNode(object):
    def __init__(self, data, l_child=None, r_child=None) -> None:
        self.data = data
        self.l_child = l_child
        self.r_child = r_child


def pre_order_traverse(bit_note: BitNode) -> None:
    if bit_note is None:
        return

    print(bit_note.data, end='')
    pre_order_traverse(bit_note=bit_note.l_child)
    pre_order_traverse(bit_note=bit_note.r_child)


def in_order_traverse(bit_note: BitNode) -> None:
    if bit_note is None:
        return

    in_order_traverse(bit_note=bit_note.l_child)
    print(bit_note.data, end='')
    in_order_traverse(bit_note=bit_note.r_child)


def post_order_traverse(bit_note: BitNode) -> None:
    if bit_note is None:
        return

    post_order_traverse(bit_note=bit_note.l_child)
    post_order_traverse(bit_note=bit_note.r_child)
    print(bit_note.data, end='')


if __name__ == "__main__":
    bit_note = BitNode(
        data="A",
        l_child=BitNode(
            data="B",
            l_child=BitNode(
                data="D",
                l_child=BitNode(
                    data="H",
                    r_child=BitNode(
                        data="K"
                    )
                )
            ),
            r_child=BitNode(
                data="E"
            )
        ),
        r_child=BitNode(
            data="C",
            l_child=BitNode(
                data="F",
                l_child=BitNode(
                    data="I"
                )
            ),
            r_child=BitNode(
                data="G",
                r_child=BitNode(
                    data="J"
                )
            )
        )
    )

    pre_order_traverse(bit_note=bit_note)
    print()
    in_order_traverse(bit_note=bit_note)
    print()
    post_order_traverse(bit_note=bit_note)
