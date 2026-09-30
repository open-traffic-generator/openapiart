from openapiart.requirements import _resolve_required_packages


def test_resolve_required_packages_preserves_curated_order():
    orig_packages = [
        "PyYAML",
        "requests",
        "grpcio-tools~=1.80.0",
        "protobuf~=6.33.6",
        "grpcio~=1.80.0",
        "typing_extensions>=4.0.0",
    ]
    test_packages = ["urllib3", "semantic_version", "sanity"]
    new_packages = [
        "semantic_version",
        "grpcio",
        "protobuf",
        "requests",
        "grpcio-tools",
        "PyYAML",
        "urllib3",
        "typing_extensions",
        "sanity",
    ]

    assert _resolve_required_packages(
        new_packages,
        orig_packages,
        test_packages,
        ignored_packages=["sanity", "typing_extensions"],
    ) == [
        "PyYAML",
        "requests",
        "grpcio-tools~=1.80.0",
        "protobuf~=6.33.6",
        "grpcio~=1.80.0",
        "urllib3",
        "semantic_version",
    ]


def test_resolve_required_packages_matches_package_names_not_substrings():
    assert _resolve_required_packages(
        new_packages=["grpcio"],
        orig_packages=["grpcio-tools~=1.80.0", "grpcio~=1.80.0"],
        test_packages=[],
    ) == ["grpcio~=1.80.0"]
