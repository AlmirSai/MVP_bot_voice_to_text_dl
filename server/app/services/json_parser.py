import json


def parse_json(content: bytes):
    try:
        data = json.loads(content.decode("utf-8"))
        return data
    except json.JSONDecodeError:
        raise ValueError("Invalid JSON file")
