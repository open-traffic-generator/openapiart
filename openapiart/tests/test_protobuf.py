from google.protobuf import json_format


def test_protobuf(config, pb_config, utils):
    """
    1. Take a json without defaults and deserialize to snappi Config object.
    2. serialize the snappi config object so that defaults get added.
    3. Deserialize snappi json to protobuf.
    4. Serialize protobuf to json

    Validation :
        Compare jsons serialized from snappi & protobuf
    """
    # Take a json without defaults and deserialize to snappi Config object
    with open(utils.get_test_config_path("config.json")) as f:
        config.deserialize(f.read())

    # Serialize the snappi config object so that defaults get added
    snappi_json = config.serialize(config.JSON)

    # Deserialize snappi_json to protobuf
    pb_obj = json_format.Parse(snappi_json, pb_config)

    # Serialize protobuf to json
    pb_json = json_format.MessageToJson(
        pb_obj, preserving_proto_field_name=True
    )

    # Compare jsons serialized from snappi & protobuf
    assert utils.compare_json(snappi_json, pb_json)
