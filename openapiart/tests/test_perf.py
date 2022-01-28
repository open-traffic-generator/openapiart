import sys
import os

sys.path.append(os.path.join(os.path.dirname(__file__), "..", "..", "art"))
import sanity
import time


def test_perf(api):
    start = time.time()
    config = sanity.Api().prefix_config()
    config.a = "abc"
    config.b = 0.2
    config.c = 20
    config.required_object.e_a = 10.2
    config.required_object.e_b = 33.1
    config.response = config.STATUS_200
    config.optional_object.e_a = 11.2
    config.optional_object.e_b = 21.1
    config.d_values = [config.A]
    config.e.e_a = 1.1
    config.e.e_b = 2.1
    config.f.f_a = "FA"
    config.g.add()
    config.h = False
    config.i = "1010100"
    j = config.j.add()
    j.j_a.e_a = 10.1
    j.j_a.e_b = 2.1
    j.j_b.f_b = 3.1
    config.k.e_object.e_a = 1.0
    config.k.e_object.e_b = 0.1
    config.k.f_object.f_a
    config.l.double = 1.1
    config.l.float = 2.1
    config.l.hex = "0x123"
    config.l.integer = 23
    config.l.ipv4 = "10.1.1.1"
    config.l.ipv6 = "::"
    config.l.mac = "aa:bb:cc:11:22:33"
    config.l.string_param = "abcd"
    config.ieee_802_1qbb = True
    config.space_1 = 10
    config.full_duplex_100_mb = 20
    config.mandatory.required_param = "required"
    config.list_of_string_values = ["abc", "efg"]
    config.list_of_integer_values = [1, 2, 3, 4]
    config.level.l1_p1.l2_p1.l3_p1 = "L3P1"
    config.ipv4_pattern.ipv4.value = "10.1.1.1"
    config.ipv6_pattern.ipv6.values = ["::", "::1"]
    config.mac_pattern.mac.increment.start = "aa:aa:bb:bb:cc:cc"
    config.mac_pattern.mac.increment.step = "00:01:00:00:10:00"
    config.mac_pattern.mac.increment.count = 10
    config.integer_pattern.integer.decrement.start = 200
    config.integer_pattern.integer.decrement.step = 2
    config.integer_pattern.integer.decrement.count = 100
    config.checksum_pattern.checksum.generated = (
        config.checksum_pattern.checksum.GOOD
    )
    config.validate()
    end = time.time()
    print("Time taken for the manual config %f ms" % ((end - start) * 1000))

    start = time.time()
    json = config.serialize()
    end = time.time()
    print("Time elapsed to serialize to Json %f ms" % ((end - start) * 1000))

    start = time.time()
    c1 = sanity.Api().prefix_config()
    c1.deserialize(json)
    end = time.time()
    print(
        "Time elapsed to deserialize from Json %f ms" % ((end - start) * 1000)
    )

    start = time.time()
    yaml = config.serialize(config.YAML)
    end = time.time()
    print("Time elapsed to serialize to yaml %f ms" % ((end - start) * 1000))

    start = time.time()
    c1 = sanity.Api().prefix_config()
    c1.deserialize(yaml)
    end = time.time()
    print(
        "Time elapsed to deserialize from yaml %f ms" % ((end - start) * 1000)
    )

    start = time.time()
    dt = config.serialize(config.DICT)
    end = time.time()
    print("Time elapsed to serialize to DICT %f ms" % ((end - start) * 1000))

    start = time.time()
    c1 = sanity.Api().prefix_config()
    c1.deserialize(dt)
    end = time.time()
    print(
        "Time elapsed to deserialize from DICT %f ms" % ((end - start) * 1000)
    )

    start = time.time()
    api.set_config(config)
    end = time.time()
    print(
        "Time elapsed on set_config http call %f ms" % ((end - start) * 1000)
    )
