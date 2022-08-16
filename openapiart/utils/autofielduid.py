import os
import fnmatch
import jsonpath_ng
from ruamel import yaml
from copy import deepcopy


class AutoFieldUid(object):
    """This utility will use to fill x-field-uid, x-enum and x-include"""

    _FIELD_UID = "x-field-uid"

    def __init__(self, parent_folder):
        self._files = []
        self._parent_folder = parent_folder
        self._include_files = {}
        for r, d, f in os.walk(parent_folder):
            for file in f:
                if fnmatch.fnmatch(file, "*.yaml"):
                    self._files.append(os.path.join(r, file))

    def annotate(self):
        for filename in self._files:
            with open(filename) as fid:
                yobject = yaml.load(
                    fid, Loader=yaml.RoundTripLoader, preserve_quotes=True
                )
            self._annotate_msg_fields(yobject, filename)
            self._annnotate_response_fields(yobject)
            self._dump_file(filename, yobject)

    def _dump_file(self, output_filename, content):
        with open(output_filename, "w") as fp:
            yaml.dump(content, fp, Dumper=yaml.RoundTripDumper)

    def _annnotate_response_fields(self, yobject):
        path_object = yobject.get("paths")
        if path_object is None:
            return
        for response in jsonpath_ng.parse("$..responses").find(path_object):
            rsp_value = response.value
            if "x-include" in rsp_value:
                self._update_x_incude_response(yobject, rsp_value)
            idx = 1
            for code, code_schema in rsp_value.items():
                code_schema.update({"x-field-uid": idx})
                idx += 1

    def _annotate_enum_fields(self, property_object):
        if "type" in property_object:
            type = property_object["type"]
            if type == "string" and "enum" in property_object:
                x_enum = dict()
                for idx, enum in enumerate(property_object["enum"]):
                    x_enum[enum] = {AutoFieldUid._FIELD_UID: idx + 1}
                property_object.update({"x-enum": x_enum})
                property_object.pop("enum")
            elif type == "array" and "items" in property_object:
                self._annotate_enum_fields(property_object["items"])

    def _annotate_msg_fields(self, yobject, filename):
        components_object = yobject.get("components")
        if components_object is None:
            return
        schema_objects = components_object.get("schemas")
        if schema_objects is None:
            return
        for schema_name, schema_object in schema_objects.items():
            # ignore content field as it always contain single value
            if "x-include" in schema_object:
                self._update_x_incude_properties(
                    yobject, schema_object, filename
                )
            if "properties" not in schema_object:
                continue
            id = 0
            for property_name, property_object in schema_object[
                "properties"
            ].items():
                id += 1
                if not isinstance(property_object, dict):
                    print(
                        "schema %s do not have dict of %s"
                        % (schema_name, property_name)
                    )
                    continue
                property_object[AutoFieldUid._FIELD_UID] = id
                self._annotate_enum_fields(property_object)

    def _get_include_response(self, yobject, object_path):
        include_response = yobject
        for node_name in object_path.split("/"):
            if node_name == str():
                continue
            include_response = include_response.get(node_name, {})
        return include_response

    def _update_x_incude_response(self, yobject, rsp_value):
        include_names = rsp_value["x-include"]
        for include_name in include_names:
            file_name, object_path = include_name.split("#")
            if file_name == str():
                include_respones = self._get_include_properties(
                    yobject, object_path
                )
            else:
                file_name = "/".join(
                    [x for x in file_name.split("/") if x != ".."]
                )
                if file_name in self._include_files:
                    file_obj = self._include_files[file_name]
                else:
                    abs_path = os.path.join(self._parent_folder, file_name)
                    with open(abs_path) as fid:
                        file_obj = yaml.load(
                            fid,
                            Loader=yaml.RoundTripLoader,
                            preserve_quotes=True,
                        )
                    self._include_files[file_name] = file_obj
                include_respones = self._get_include_response(
                    file_obj, object_path
                )
            for response_name in include_respones:
                rsp_value.update(
                    {
                        response_name: {
                            "x-include": "{include_name}/{property_name}".format(
                                include_name=include_name,
                                property_name=response_name,
                            )
                        }
                    }
                )
        rsp_value.pop("x-include")

    def _get_include_properties(self, yobject, object_path):
        include_properties = yobject
        for node_name in object_path.split("/"):
            if node_name == str():
                continue
            include_properties = include_properties.get(node_name, {})
        return include_properties.get("properties", {})

    def _merge(self, src, dst):
        for key, value in src.items():
            if key in ["x-include", "properties", "x-field-uid"]:
                continue
            if key not in dst:
                dst[key] = deepcopy(value)
            elif isinstance(value, list):
                for item in value:
                    if item not in dst[key]:
                        dst[key].append(item)
            elif isinstance(value, dict):
                self._merge(value, dst[key])

    def _update_x_incude_properties(self, yobject, schema_object, file_path):
        include_names = schema_object["x-include"]
        properties = schema_object.get("properties")
        if properties is None:
            schema_object.update({"properties": {}})
            properties = schema_object.get("properties")
        for include_name in include_names:
            file_name, object_path = include_name.split("#")
            if file_name == str() or file_name == ".":
                include_properties = self._get_include_properties(
                    yobject, object_path
                )
                include_obj = self._get_include_response(yobject, object_path)
            else:
                if file_name in self._include_files:
                    file_obj = self._include_files[file_name]
                else:
                    (parent_path, _) = os.path.split(file_path)
                    abs_path = os.path.join(parent_path, file_name)
                    if not os.path.exists(abs_path):
                        file_name = "/".join(
                            [x for x in file_name.split("/") if x != ".."]
                        )

                        abs_path = os.path.join(self._parent_folder, file_name)
                    with open(abs_path) as fid:
                        file_obj = yaml.load(
                            fid,
                            Loader=yaml.RoundTripLoader,
                            preserve_quotes=True,
                        )

                    file_schema = self._get_include_response(
                        file_obj, object_path
                    )
                    if "x-include" in file_schema:
                        file_obj = self._update_x_incude_properties(
                            file_obj, file_schema, abs_path
                        )
                    self._include_files[file_name] = file_obj
                include_properties = self._get_include_properties(
                    file_obj, object_path
                )
                include_obj = self._get_include_response(file_obj, object_path)

            self._merge(include_obj, schema_object)
            for property_name in include_properties:
                properties.update(
                    {
                        property_name: {
                            "x-include": "{include_name}/properties/{property_name}".format(
                                include_name=include_name,
                                property_name=property_name,
                            )
                        }
                    }
                )
        schema_object.pop("x-include")
        return yobject


if __name__ == "__main__":
    parent_folder = "D:/OTG/Codebase/models"
    AutoFieldUid(parent_folder).annotate()
